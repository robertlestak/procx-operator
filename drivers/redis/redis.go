package redis

import (
	"strconv"

	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"

	corev1 "k8s.io/api/core/v1"
)

type RedisList struct {
	Host               string  `json:"host"`
	Port               int     `json:"port"`
	Password           *string `json:"password,omitempty"`
	PasswordSecretName *string `json:"passwordSecretName,omitempty"`
	Key                string  `json:"key"`
	ListLength         *string `json:"listLength,omitempty"`
}

func (d *RedisList) ConfigSecret() map[string]string {
	secData := map[string]string{}
	if d != nil {
		secData["PROCX_REDIS_HOST"] = d.Host
		secData["PROCX_REDIS_KEY"] = d.Key
		secData["PROCX_REDIS_PORT"] = strconv.Itoa(int(d.Port))
		if d.Password != nil {
			secData["PROCX_REDIS_PASSWORD"] = *d.Password
		}
	}
	return secData
}

func (d *RedisList) Metadata() map[string]string {
	md := map[string]string{
		"address":  d.Host + ":" + strconv.Itoa(d.Port),
		"listName": d.Key,
	}
	if d.ListLength == nil {
		md["listLength"] = "1"
	} else {
		md["listLength"] = *d.ListLength
	}
	return md
}

func (d *RedisList) HasAuth() bool {
	if d.Password != nil {
		return true
	}
	if d.PasswordSecretName != nil {
		return true
	}
	return false
}

func (d *RedisList) KedaSupport() bool {
	return true
}

func (d *RedisList) KedaScalerName() string {
	return "redis"
}

func (d *RedisList) ContainerEnv() []corev1.EnvFromSource {
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

func (d *RedisList) TriggerAuth(name string) *kedav1alpha1.TriggerAuthenticationSpec {
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
