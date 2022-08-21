package pulsar

import (
	"encoding/json"
	"fmt"
	"strings"

	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

type Pulsar struct {
	Address             string             `json:"address"`
	Subscription        *string            `json:"subscription,omitempty"`
	Topic               *string            `json:"topic,omitempty"`
	TopicsPattern       *string            `json:"topicsPattern,omitempty"`
	Topics              []string           `json:"topics,omitempty"`
	AuthToken           *string            `json:"authToken,omitempty"`
	AuthTokenSecretName *string            `json:"authTokenSecretName,omitempty"`
	AuthTokenFile       *string            `json:"authTokenFile,omitempty"`
	AuthCertPath        *string            `json:"authCertPath,omitempty"`
	AuthKeyPath         *string            `json:"authKeyPath,omitempty"`
	AuthOAuthParams     *map[string]string `json:"authOAuthParams,omitempty"`
	// TLS
	TLSSecretName              *string `json:"tlsSecretName,omitempty"`
	TLSTrustCertsFilePath      *string `json:"tlsTrustCertsFilePath,omitempty"`
	TLSAllowInsecureConnection *bool   `json:"tlsAllowInsecureConnection,omitempty"`
	TLSValidateHostname        *bool   `json:"tlsValidateHostname,omitempty"`
}

func (d *Pulsar) ConfigSecret() map[string]string {
	secData := map[string]string{}
	if d.Address != "" {
		secData["PROCX_PULSAR_ADDRESS"] = d.Address
	}
	if d.AuthToken != nil {
		secData["PROCX_PULSAR_AUTH_TOKEN"] = *d.AuthToken
	}
	if d.Subscription != nil {
		secData["PROCX_PULSAR_SUBSCRIPTION"] = *d.Subscription
	}
	if d.Topic != nil {
		secData["PROCX_PULSAR_TOPIC"] = *d.Topic
	}
	if d.TopicsPattern != nil {
		secData["PROCX_PULSAR_TOPICS_PATTERN"] = *d.TopicsPattern
	}
	if len(d.Topics) > 0 {
		secData["PROCX_PULSAR_TOPICS"] = strings.Join(d.Topics, ",")
	}
	if d.AuthTokenFile != nil {
		secData["PROCX_PULSAR_AUTH_TOKEN_FILE"] = *d.AuthTokenFile
	}
	if d.AuthCertPath != nil {
		secData["PROCX_PULSAR_AUTH_CERT_FILE"] = *d.AuthCertPath
	}
	if d.AuthKeyPath != nil {
		secData["PROCX_PULSAR_AUTH_KEY_FILE"] = *d.AuthKeyPath
	}
	if d.AuthOAuthParams != nil {
		jd, err := json.Marshal(*d.AuthOAuthParams)
		if err != nil {
			return nil
		}
		secData["PROCX_PULSAR_AUTH_OAUTH_PARAMS"] = string(jd)
	}
	if d.TLSTrustCertsFilePath != nil {
		secData["PROCX_PULSAR_TLS_TRUST_CERTS_FILE"] = *d.TLSTrustCertsFilePath
	}
	if d.TLSAllowInsecureConnection != nil {
		secData["PROCX_PULSAR_TLS_ALLOW_INSECURE_CONNECTION"] = fmt.Sprintf("%v", *d.TLSAllowInsecureConnection)
	}
	if d.TLSValidateHostname != nil {
		secData["PROCX_PULSAR_TLS_VALIDATE_HOSTNAME"] = fmt.Sprintf("%v", *d.TLSValidateHostname)
	}
	return secData
}

func (d *Pulsar) Metadata() map[string]string {
	md := map[string]string{}
	return md
}

func (d *Pulsar) HasAuth() bool {
	return d.AuthTokenSecretName != nil || d.AuthTokenFile != nil || d.AuthCertPath != nil || d.AuthKeyPath != nil || d.AuthOAuthParams != nil
}

func (d *Pulsar) KedaSupport() bool {
	return false
}

func (d *Pulsar) KedaScalerName() string {
	return ""
}

func (d *Pulsar) ContainerEnv() []corev1.EnvFromSource {
	envFrom := []corev1.EnvFromSource{}
	falseVal := false
	if d.AuthTokenSecretName != nil {
		envFrom = append(envFrom, corev1.EnvFromSource{
			SecretRef: &corev1.SecretEnvSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: *d.AuthTokenSecretName,
				},
				Optional: &falseVal,
			},
		})
	}
	return envFrom
}

func (d *Pulsar) TriggerAuth(name string) *kedav1alpha1.TriggerAuthenticationSpec {
	s := &kedav1alpha1.TriggerAuthenticationSpec{}
	return s
}

func (d *Pulsar) VolumeMounts() []corev1.VolumeMount {
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

func (d *Pulsar) Volumes() []corev1.Volume {
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
