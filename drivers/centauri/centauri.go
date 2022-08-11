package centauri

import (
	"encoding/base64"

	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

type Centauri struct {
	URL                  string  `json:"url"`
	PrivateKey           *[]byte `json:"privateKey,omitempty"`
	PrivateKeySecretName *string `json:"privateKeySecretName,omitempty"`
	TLSSecretName        *string `json:"tlsSecretName,omitempty"`
	Channel              *string `json:"channel,omitempty"`
	Key                  *string `json:"key,omitempty"`
}

func (d *Centauri) ConfigSecret() map[string]string {
	secData := map[string]string{}
	if d.URL != "" {
		secData["PROCX_CENTAURI_URL"] = d.URL
	}
	if d.Channel != nil {
		secData["PROCX_CENTAURI_CHANNEL"] = *d.Channel
	}
	if d.PrivateKey != nil && len(*d.PrivateKey) > 0 {
		bd := base64.StdEncoding.EncodeToString(*d.PrivateKey)
		secData["PROCX_CENTAURI_KEY_BASE64"] = bd
	}
	return secData
}

func (d *Centauri) Metadata() map[string]string {
	md := map[string]string{}
	return md
}

func (d *Centauri) HasAuth() bool {
	return true
}

func (d *Centauri) KedaSupport() bool {
	return false
}

func (d *Centauri) KedaScalerName() string {
	return ""
}

func (d *Centauri) ContainerEnv() []corev1.EnvFromSource {
	envFrom := []corev1.EnvFromSource{}
	falseVal := false
	if d.PrivateKeySecretName != nil && *d.PrivateKeySecretName != "" {
		envFrom = append(envFrom, corev1.EnvFromSource{
			SecretRef: &corev1.SecretEnvSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: *d.PrivateKeySecretName,
				},
				Optional: &falseVal,
			},
		})
	}
	return envFrom
}

func (d *Centauri) TriggerAuth(name string) *kedav1alpha1.TriggerAuthenticationSpec {
	s := &kedav1alpha1.TriggerAuthenticationSpec{}
	return s
}

func (d *Centauri) VolumeMounts() []corev1.VolumeMount {
	if d.TLSSecretName == nil || *d.TLSSecretName == "" {
		return nil
	}
	v := corev1.VolumeMount{
		Name:      *d.TLSSecretName,
		MountPath: "/etc/procx/tls",
		ReadOnly:  true,
	}
	return []corev1.VolumeMount{v}
}

func (d *Centauri) Volumes() []corev1.Volume {
	if d.TLSSecretName == nil || *d.TLSSecretName == "" {
		return nil
	}
	v := corev1.Volume{
		Name: *d.TLSSecretName,
		VolumeSource: corev1.VolumeSource{
			Secret: &corev1.SecretVolumeSource{
				SecretName: *d.TLSSecretName,
			},
		},
	}
	return []corev1.Volume{v}
}
