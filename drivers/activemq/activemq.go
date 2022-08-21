package activemq

import (
	"strconv"

	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

type ActiveMQ struct {
	Address string  `json:"address"`
	Type    *string `json:"type,omitempty"`
	Name    *string `json:"name,omitempty"`
	// TLS
	TLSSecretName *string `json:"tlsSecretName,omitempty"`
	EnableTLS     *bool   `json:"enableTLS,omitempty"`
	TLSInsecure   *bool   `json:"tlsInsecure,omitempty"`
	TLSCert       *string `json:"tlsCert,omitempty"`
	TLSKey        *string `json:"tlsKey,omitempty"`
	TLSCA         *string `json:"tlsCA,omitempty"`
}

func (d *ActiveMQ) KedaSupport() bool {
	return false
}

func (d *ActiveMQ) ConfigSecret() map[string]string {
	secData := map[string]string{}
	if d.Address != "" {
		secData["PROCX_ACTIVEMQ_ADDRESS"] = d.Address
	}
	if d.Type != nil {
		secData["PROCX_ACTIVEMQ_TYPE"] = *d.Type
	}
	if d.Name != nil {
		secData["PROCX_ACTIVEMQ_NAME"] = *d.Name
	}
	if d.EnableTLS != nil {
		secData["PROCX_ACTIVEMQ_ENABLE_TLS"] = strconv.FormatBool(*d.EnableTLS)
	}
	if d.TLSInsecure != nil {
		secData["PROCX_ACTIVEMQ_TLS_INSECURE"] = strconv.FormatBool(*d.TLSInsecure)
	}
	if d.TLSCert != nil {
		secData["PROCX_ACTIVEMQ_TLS_CERT_FILE"] = *d.TLSCert
	}
	if d.TLSKey != nil {
		secData["PROCX_ACTIVEMQ_TLS_KEY_FILE"] = *d.TLSKey
	}
	if d.TLSCA != nil {
		secData["PROCX_ACTIVEMQ_TLS_CA_FILE"] = *d.TLSCA
	}
	return secData
}

func (d *ActiveMQ) TriggerAuth(name string) *kedav1alpha1.TriggerAuthenticationSpec {
	s := &kedav1alpha1.TriggerAuthenticationSpec{}
	return s
}

func (d *ActiveMQ) HasAuth() bool {
	return false
}

func (d *ActiveMQ) KedaScalerName() string {
	return "activemq"
}

func (d *ActiveMQ) ContainerEnv() []corev1.EnvFromSource {
	envFrom := []corev1.EnvFromSource{}
	return envFrom
}

func (d *ActiveMQ) Metadata() map[string]string {
	md := map[string]string{}
	return md
}

func (d *ActiveMQ) VolumeMounts() []corev1.VolumeMount {
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

func (d *ActiveMQ) Volumes() []corev1.Volume {
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
