package rabbitmq

import (
	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"

	corev1 "k8s.io/api/core/v1"
)

type RabbitMQ struct {
	URL           string  `json:"url,omitempty"`
	URLSecretName *string `json:"urlSecretName,omitempty"`
	Queue         string  `json:"queue"`
	Mode          *string `json:"mode,omitempty"`
	Value         *string `json:"value,omitempty"`
}

func (d *RabbitMQ) ConfigSecret() map[string]string {
	secData := map[string]string{}
	if d != nil {
		if d.URL != "" {
			secData["PROCX_RABBITMQ_URL"] = d.URL
		}
		secData["PROCX_RABBITMQ_QUEUE"] = d.Queue
	}
	return secData
}

func (d *RabbitMQ) Metadata() map[string]string {
	md := map[string]string{
		"queueName": d.Queue,
	}
	if d.URL != "" {
		md["host"] = d.URL
	}
	if d.Mode == nil {
		md["mode"] = "QueueLength"
	} else {
		md["mode"] = *d.Mode
	}
	if d.Value == nil {
		md["value"] = "1"
	} else {
		md["value"] = *d.Value
	}
	return md
}

func (d *RabbitMQ) HasAuth() bool {
	if d.URLSecretName != nil {
		return true
	}
	return false
}

func (d *RabbitMQ) KedaSupport() bool {
	return true
}

func (d *RabbitMQ) KedaScalerName() string {
	return "rabbitmq"
}

func (d *RabbitMQ) ContainerEnv() []corev1.EnvFromSource {
	var envFrom []corev1.EnvFromSource
	falseVal := false
	if d.URLSecretName != nil {
		envFrom = append(envFrom, corev1.EnvFromSource{
			SecretRef: &corev1.SecretEnvSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: *d.URLSecretName,
				},
				Optional: &falseVal,
			},
		})
	}
	return envFrom
}

func (d *RabbitMQ) TriggerAuth(string) *kedav1alpha1.TriggerAuthenticationSpec {
	s := &kedav1alpha1.TriggerAuthenticationSpec{}
	if d.URLSecretName != nil {
		s.SecretTargetRef = append(s.SecretTargetRef, kedav1alpha1.AuthSecretTargetRef{
			Name:      *d.URLSecretName,
			Key:       "PROCX_RABBITMQ_URL",
			Parameter: "host",
		})
	}
	return s
}

func (d *RabbitMQ) VolumeMounts() []corev1.VolumeMount {
	return nil
}

func (d *RabbitMQ) Volumes() []corev1.Volume {
	return nil
}
