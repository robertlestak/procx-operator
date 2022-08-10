package redis

import (
	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"

	corev1 "k8s.io/api/core/v1"
)

type RedisPubSub struct {
	Host               string  `json:"host"`
	Port               string  `json:"port"`
	Password           *string `json:"password,omitempty"`
	PasswordSecretName *string `json:"passwordSecretName,omitempty"`
	Key                *string `json:"key,omitempty"`
	// TLS
	EnableTLS   *bool   `json:"enableTLS,omitempty"`
	TLSInsecure *bool   `json:"tlsInsecure,omitempty"`
	TLSCert     *string `json:"tlsCert,omitempty"`
	TLSKey      *string `json:"tlsKey,omitempty"`
	TLSCA       *string `json:"tlsCA,omitempty"`
}

func (d *RedisPubSub) ConfigSecret() map[string]string {
	secData := map[string]string{}
	if d != nil {
		secData["PROCX_REDIS_HOST"] = d.Host
		secData["PROCX_REDIS_KEY"] = *d.Key
		secData["PROCX_REDIS_PORT"] = d.Port
		if d.Password != nil {
			secData["PROCX_REDIS_PASSWORD"] = *d.Password
		}
		if d.EnableTLS != nil && *d.EnableTLS {
			secData["PROCX_REDIS_STREAM_ENABLE_TLS"] = "true"
		}
		if d.TLSInsecure != nil && *d.TLSInsecure {
			secData["PROCX_REDIS_STREAM_TLS_INSECURE"] = "true"
		}
		if d.TLSCert != nil {
			secData["PROCX_REDIS_STREAM_TLS_CERT_FILE"] = *d.TLSCert
		}
		if d.TLSKey != nil {
			secData["PROCX_REDIS_STREAM_TLS_KEY_FILE"] = *d.TLSKey
		}
		if d.TLSCA != nil {
			secData["PROCX_REDIS_STREAM_TLS_CA_FILE"] = *d.TLSCA
		}
	}
	return secData
}

func (d *RedisPubSub) Metadata() map[string]string {
	md := map[string]string{}
	return md
}

func (d *RedisPubSub) HasAuth() bool {
	if d.Password != nil {
		return true
	}
	if d.PasswordSecretName != nil {
		return true
	}
	return false
}

func (d *RedisPubSub) KedaSupport() bool {
	return false
}

func (d *RedisPubSub) KedaScalerName() string {
	return ""
}

func (d *RedisPubSub) ContainerEnv() []corev1.EnvFromSource {
	var envFrom []corev1.EnvFromSource
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

func (d *RedisPubSub) TriggerAuth(name string) *kedav1alpha1.TriggerAuthenticationSpec {
	s := &kedav1alpha1.TriggerAuthenticationSpec{}
	if d.PasswordSecretName != nil {
		s.SecretTargetRef = append(s.SecretTargetRef, kedav1alpha1.AuthSecretTargetRef{
			Name:      *d.PasswordSecretName,
			Key:       "PROCX_REDIS_PASSWORD",
			Parameter: "password",
		})
	}
	if d.Password != nil {
		s.SecretTargetRef = append(s.SecretTargetRef, kedav1alpha1.AuthSecretTargetRef{
			Name:      name,
			Key:       "PROCX_REDIS_PASSWORD",
			Parameter: "password",
		})
	}
	return s
}
