package nsq

import (
	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

type NSQ struct {
	NsqLookupdAddress *string `json:"nsqLookupdAddress,omitempty"`
	NsqdAddress       *string `json:"nsqdAddress,omitempty"`
	Topic             *string `json:"topic,omitempty"`
	Channel           *string `json:"channel,omitempty"`
	// TLS
	EnableTLS   *bool   `json:"enableTLS,omitempty"`
	TLSInsecure *bool   `json:"tlsInsecure,omitempty"`
	TLSCert     *string `json:"tlsCert,omitempty"`
	TLSKey      *string `json:"tlsKey,omitempty"`
	TLSCA       *string `json:"tlsCA,omitempty"`
}

func (d *NSQ) ConfigSecret() map[string]string {
	secData := map[string]string{}
	if d.NsqLookupdAddress != nil && *d.NsqLookupdAddress != "" {
		secData["PROCX_NSQ_NSQLOOKUPD_ADDRESS"] = *d.NsqLookupdAddress
	}
	if d.NsqdAddress != nil && *d.NsqdAddress != "" {
		secData["PROCX_NSQ_NSQD_ADDRESS"] = *d.NsqdAddress
	}
	if d.Topic != nil && *d.Topic != "" {
		secData["PROCX_NSQ_TOPIC"] = *d.Topic
	}
	if d.Channel != nil && *d.Channel != "" {
		secData["PROCX_NSQ_CHANNEL"] = *d.Channel
	}
	if d.EnableTLS != nil && *d.EnableTLS {
		secData["PROCX_NSQ_ENABLE_TLS"] = "true"
	}
	if d.TLSInsecure != nil && *d.TLSInsecure {
		secData["PROCX_NSQ_TLS_INSECURE"] = "true"
	}
	if d.TLSCert != nil && *d.TLSCert != "" {
		secData["PROCX_NSQ_TLS_CERT_FILE"] = *d.TLSCert
	}
	if d.TLSKey != nil && *d.TLSKey != "" {
		secData["PROCX_NSQ_TLS_KEY_FILE"] = *d.TLSKey
	}
	if d.TLSCA != nil && *d.TLSCA != "" {
		secData["PROCX_NSQ_TLS_CA_FILE"] = *d.TLSCA
	}
	return secData
}

func (d *NSQ) Metadata() map[string]string {
	md := map[string]string{}
	return md
}

func (d *NSQ) HasAuth() bool {
	return false
}

func (d *NSQ) KedaSupport() bool {
	return false
}

func (d *NSQ) KedaScalerName() string {
	return ""
}

func (d *NSQ) ContainerEnv() []corev1.EnvFromSource {
	envFrom := []corev1.EnvFromSource{}
	return envFrom
}

func (d *NSQ) TriggerAuth(name string) *kedav1alpha1.TriggerAuthenticationSpec {
	s := &kedav1alpha1.TriggerAuthenticationSpec{}
	return s
}

func (d *NSQ) VolumeMounts() []corev1.VolumeMount {
	return nil
}

func (d *NSQ) Volumes() []corev1.Volume {
	return nil
}
