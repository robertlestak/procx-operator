package kafka

import (
	"fmt"
	"strings"

	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

type KafkaSaslType string

var (
	KafkaSaslTypePlain = KafkaSaslType("plain")
	KafkaSaslTypeScram = KafkaSaslType("scram")
)

type Kafka struct {
	Brokers *[]string `json:"brokers,omitempty"`
	Group   *string   `json:"group,omitempty"`
	Topic   *string   `json:"topic,omitempty"`
	// TLS
	TLSSecretName *string `json:"tlsSecretName,omitempty"`
	EnableTLS     *bool   `json:"enableTLS,omitempty"`
	TLSInsecure   *bool   `json:"tlsInsecure,omitempty"`
	TLSCert       *string `json:"tlsCert,omitempty"`
	TLSKey        *string `json:"tlsKey,omitempty"`
	TLSCA         *string `json:"tlsCA,omitempty"`
	// SASL
	EnableSASL                 *bool          `json:"enableSASL,omitempty"`
	SaslType                   *KafkaSaslType `json:"saslType,omitempty"`
	Username                   *string        `json:"username,omitempty"`
	Password                   *string        `json:"password,omitempty"`
	PasswordSecretName         *string        `json:"passwordSecretName,omitempty"`
	LagThreshold               *string        `json:"lagThreshold,omitempty"`
	ActivationThreshold        *string        `json:"activationThreshold,omitempty"`
	OffsetResetPolicy          *string        `json:"offsetResetPolicy,omitempty"`
	AllowIdleConsumers         *bool          `json:"allowIdleConsumers,omitempty"`
	ScaleToZeroOnInvalidOffset *bool          `json:"scaleToZeroOnInvalidOffset,omitempty"`
	Version                    *string        `json:"version,omitempty"`
}

func (d *Kafka) ConfigSecret() map[string]string {
	secData := map[string]string{}
	if d.Password != nil && *d.Password != "" {
		secData["PROCX_KAFKA_SASL_PASSWORD"] = *d.Password
	}
	if d.Username != nil && *d.Username != "" {
		secData["PROCX_KAFKA_SASL_USERNAME"] = *d.Username
	}
	if d.Brokers != nil && len(*d.Brokers) > 0 {
		secData["PROCX_KAFKA_BROKERS"] = strings.Join(*d.Brokers, ",")
	}
	if d.EnableTLS != nil && *d.EnableTLS {
		secData["PROCX_KAFKA_ENABLE_TLS"] = "true"
	}
	if d.TLSInsecure != nil && *d.TLSInsecure {
		secData["PROCX_KAFKA_TLS_INSECURE"] = "true"
	}
	if d.TLSCert != nil && *d.TLSCert != "" {
		secData["PROCX_KAFKA_TLS_CERT_FILE"] = *d.TLSCert
	}
	if d.TLSKey != nil && *d.TLSKey != "" {
		secData["PROCX_KAFKA_TLS_KEY_FILE"] = *d.TLSKey
	}
	if d.TLSCA != nil && *d.TLSCA != "" {
		secData["PROCX_KAFKA_TLS_CA_FILE"] = *d.TLSCA
	}
	if d.EnableSASL != nil && *d.EnableSASL {
		secData["PROCX_KAFKA_ENABLE_SASL"] = "true"
	}
	if d.SaslType != nil {
		secData["PROCX_KAFKA_SASL_TYPE"] = string(*d.SaslType)
	}
	if d.Group != nil && *d.Group != "" {
		secData["PROCX_KAFKA_GROUP"] = *d.Group
	}
	if d.Topic != nil && *d.Topic != "" {
		secData["PROCX_KAFKA_TOPIC"] = *d.Topic
	}
	return secData
}

func (d *Kafka) Metadata() map[string]string {
	md := map[string]string{}
	if d.Brokers != nil && len(*d.Brokers) > 0 {
		md["bootstrapServers"] = strings.Join(*d.Brokers, ",")
	}
	if d.Group != nil && *d.Group != "" {
		md["consumerGroup"] = *d.Group
	}
	if d.Topic != nil && *d.Topic != "" {
		md["topic"] = *d.Topic
	}
	if d.LagThreshold != nil && *d.LagThreshold != "" {
		md["lagThreshold"] = *d.LagThreshold
	}
	if d.ActivationThreshold != nil && *d.ActivationThreshold != "" {
		md["activationThreshold"] = *d.ActivationThreshold
	}
	if d.OffsetResetPolicy != nil && *d.OffsetResetPolicy != "" {
		md["offsetResetPolicy"] = *d.OffsetResetPolicy
	}
	if d.AllowIdleConsumers != nil {
		md["allowIdleConsumers"] = fmt.Sprintf("%v", *d.AllowIdleConsumers)
	}
	if d.ScaleToZeroOnInvalidOffset != nil {
		md["scaleToZeroOnInvalidOffset"] = fmt.Sprintf("%v", *d.ScaleToZeroOnInvalidOffset)
	}
	if d.Version != nil && *d.Version != "" {
		md["version"] = *d.Version
	}
	return md
}

func (d *Kafka) HasAuth() bool {
	if d.Username != nil && *d.Username != "" {
		return true
	}
	if d.PasswordSecretName != nil {
		return true
	}
	if d.Password != nil {
		return true
	}
	return false
}

func (d *Kafka) KedaSupport() bool {
	return true
}

func (d *Kafka) KedaScalerName() string {
	return "kafka"
}

func (d *Kafka) ContainerEnv() []corev1.EnvFromSource {
	envFrom := []corev1.EnvFromSource{}
	falseVal := false
	if d.PasswordSecretName != nil {
		envFrom = append(envFrom, corev1.EnvFromSource{
			SecretRef: &corev1.SecretEnvSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: *d.PasswordSecretName,
				},
				Optional: &falseVal,
			},
		})
	}
	return envFrom
}

func (d *Kafka) TriggerAuth(name string) *kedav1alpha1.TriggerAuthenticationSpec {
	s := &kedav1alpha1.TriggerAuthenticationSpec{}
	if d.PasswordSecretName != nil {
		s.SecretTargetRef = append(s.SecretTargetRef, kedav1alpha1.AuthSecretTargetRef{
			Name:      *d.PasswordSecretName,
			Key:       "PROCX_KAFKA_SASL_PASSWORD",
			Parameter: "password",
		})
	} else if d.Password != nil && *d.Password != "" {
		s.SecretTargetRef = append(s.SecretTargetRef, kedav1alpha1.AuthSecretTargetRef{
			Name:      name,
			Key:       "PROCX_KAFKA_SASL_PASSWORD",
			Parameter: "password",
		})
	}
	return s
}

func (d *Kafka) VolumeMounts() []corev1.VolumeMount {
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

func (d *Kafka) Volumes() []corev1.Volume {
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
