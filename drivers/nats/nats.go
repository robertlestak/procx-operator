package nats

import (
	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

type NATS struct {
	URL                *string `json:"url,omitempty"`
	Subject            *string `json:"subject,omitempty"`
	QueueGroup         *string `json:"queueGroup,omitempty"`
	CredsFile          *string `json:"credsFile,omitempty"`
	JWTFile            *string `json:"jwtFile,omitempty"`
	NKeyFile           *string `json:"nkeyFile,omitempty"`
	Username           *string `json:"username,omitempty"`
	Password           *string `json:"password,omitempty"`
	PasswordSecretName *string `json:"passwordSecretName,omitempty"`
	Token              *string `json:"token,omitempty"`
	TokenSecretName    *string `json:"tokenSecretName,omitempty"`
	EnableTLS          *bool   `json:"enableTLS,omitempty"`
	TLSInsecure        *bool   `json:"tlsInsecure,omitempty"`
	TLSCA              *string `json:"tlsCA,omitempty"`
	TLSCert            *string `json:"tlsCert,omitempty"`
	TLSKey             *string `json:"tlsKey,omitempty"`
	ClearResponse      *string `json:"clearResponse,omitempty"`
	FailResponse       *string `json:"failResponse,omitempty"`
}

func (d *NATS) ConfigSecret() map[string]string {
	secData := map[string]string{}
	if d.URL != nil && *d.URL != "" {
		secData["PROCX_NATS_URL"] = *d.URL
	}
	if d.Subject != nil && *d.Subject != "" {
		secData["PROCX_NATS_SUBJECT"] = *d.Subject
	}
	if d.QueueGroup != nil && *d.QueueGroup != "" {
		secData["PROCX_NATS_QUEUE_GROUP"] = *d.QueueGroup
	}
	if d.CredsFile != nil && *d.CredsFile != "" {
		secData["PROCX_NATS_CREDS_FILE"] = *d.CredsFile
	}
	if d.JWTFile != nil && *d.JWTFile != "" {
		secData["PROCX_NATS_JWT_FILE"] = *d.JWTFile
	}
	if d.NKeyFile != nil && *d.NKeyFile != "" {
		secData["PROCX_NATS_NKEY_FILE"] = *d.NKeyFile
	}
	if d.Username != nil && *d.Username != "" {
		secData["PROCX_NATS_USERNAME"] = *d.Username
	}
	if d.Password != nil && *d.Password != "" {
		secData["PROCX_NATS_PASSWORD"] = *d.Password
	}
	if d.Token != nil && *d.Token != "" {
		secData["PROCX_NATS_TOKEN"] = *d.Token
	}
	if d.ClearResponse != nil && *d.ClearResponse != "" {
		secData["PROCX_NATS_CLEAR_RESPONSE"] = *d.ClearResponse
	}
	if d.FailResponse != nil && *d.FailResponse != "" {
		secData["PROCX_NATS_FAIL_RESPONSE"] = *d.FailResponse
	}
	if d.EnableTLS != nil && *d.EnableTLS {
		secData["PROCX_NATS_ENABLE_TLS"] = "true"
	}
	if d.TLSInsecure != nil && *d.TLSInsecure {
		secData["PROCX_NATS_TLS_INSECURE"] = "true"
	}
	if d.TLSCert != nil && *d.TLSCert != "" {
		secData["PROCX_NATS_TLS_CERT_FILE"] = *d.TLSCert
	}
	if d.TLSKey != nil && *d.TLSKey != "" {
		secData["PROCX_NATS_TLS_KEY_FILE"] = *d.TLSKey
	}
	if d.TLSCA != nil && *d.TLSCA != "" {
		secData["PROCX_NATS_TLS_CA_FILE"] = *d.TLSCA
	}
	return secData
}

func (d *NATS) Metadata() map[string]string {
	md := map[string]string{}
	return md
}

func (d *NATS) HasAuth() bool {
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

func (d *NATS) KedaSupport() bool {
	return false
}

func (d *NATS) KedaScalerName() string {
	return ""
}

func (d *NATS) ContainerEnv() []corev1.EnvFromSource {
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
	if d.TokenSecretName != nil {
		envFrom = append(envFrom, corev1.EnvFromSource{
			SecretRef: &corev1.SecretEnvSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: *d.TokenSecretName,
				},
				Optional: &falseVal,
			},
		})
	}
	return envFrom
}

func (d *NATS) TriggerAuth(name string) *kedav1alpha1.TriggerAuthenticationSpec {
	s := &kedav1alpha1.TriggerAuthenticationSpec{}
	if d.PasswordSecretName != nil {
		s.SecretTargetRef = append(s.SecretTargetRef, kedav1alpha1.AuthSecretTargetRef{
			Name:      *d.PasswordSecretName,
			Key:       "PROCX_NATS_PASSWORD",
			Parameter: "password",
		})
	} else if d.Password != nil && *d.Password != "" {
		s.SecretTargetRef = append(s.SecretTargetRef, kedav1alpha1.AuthSecretTargetRef{
			Name:      name,
			Key:       "PROCX_NATS_PASSWORD",
			Parameter: "password",
		})
	}
	if d.TokenSecretName != nil {
		s.SecretTargetRef = append(s.SecretTargetRef, kedav1alpha1.AuthSecretTargetRef{
			Name:      *d.TokenSecretName,
			Key:       "PROCX_NATS_TOKEN",
			Parameter: "token",
		})
	} else if d.Token != nil && *d.Token != "" {
		s.SecretTargetRef = append(s.SecretTargetRef, kedav1alpha1.AuthSecretTargetRef{
			Name:      name,
			Key:       "PROCX_NATS_TOKEN",
			Parameter: "token",
		})
	}
	return s
}
