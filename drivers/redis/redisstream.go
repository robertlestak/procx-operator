package redis

import (
	"fmt"
	"strings"

	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"

	corev1 "k8s.io/api/core/v1"
)

type RedisStreamOp string

var (
	RedisStreamOpAck = RedisStreamOp("ack")
	RedisStreamOpDel = RedisStreamOp("del")
)

type RedisStream struct {
	Host               string    `json:"host"`
	Port               string    `json:"port"`
	Password           *string   `json:"password,omitempty"`
	PasswordSecretName *string   `json:"passwordSecretName,omitempty"`
	ConsumerName       *string   `json:"consumerName,omitempty"`
	ConsumerGroup      *string   `json:"consumerGroup,omitempty"`
	Key                *string   `json:"key"`
	ValueKeys          *[]string `json:"valueKeys,omitempty"`
	MessageID          *string
	ClearOp            *RedisStreamOp `json:"clearOp,omitempty"`
	FailOp             *RedisStreamOp `json:"failOp,omitempty"`
	// TLS
	EnableTLS           *bool   `json:"enableTLS,omitempty"`
	TLSInsecure         *bool   `json:"tlsInsecure,omitempty"`
	TLSCert             *string `json:"tlsCert,omitempty"`
	TLSKey              *string `json:"tlsKey,omitempty"`
	TLSCA               *string `json:"tlsCA,omitempty"`
	PendingEntriesCount *string `json:"pendingEntriesCount,omitempty"`
}

func (d *RedisStream) ConfigSecret() map[string]string {
	secData := map[string]string{}
	if d != nil {
		secData["PROCX_REDIS_HOST"] = d.Host
		secData["PROCX_REDIS_KEY"] = *d.Key
		secData["PROCX_REDIS_PORT"] = d.Port
		if d.Password != nil {
			secData["PROCX_REDIS_PASSWORD"] = *d.Password
		}
		if d.ConsumerName != nil {
			secData["PROCX_REDIS_STREAM_CONSUMER_NAME"] = *d.ConsumerName
		}
		if d.ConsumerGroup != nil {
			secData["PROCX_REDIS_STREAM_CONSUMER_GROUP"] = *d.ConsumerGroup
		}
		if d.ValueKeys != nil && len(*d.ValueKeys) > 0 {
			secData["PROCX_REDIS_STREAM_VALUE_KEYS"] = strings.Join(*d.ValueKeys, ",")
		}
		if d.ClearOp != nil {
			secData["PROCX_REDIS_STREAM_CLEAR_OP"] = string(*d.ClearOp)
		}
		if d.FailOp != nil {
			secData["PROCX_REDIS_STREAM_FAIL_OP"] = string(*d.FailOp)
		}
		if d.EnableTLS != nil && *d.EnableTLS {
			secData["PROCX_REDIS_STREAM_ENABLE_TLS"] = "true"
		}
		if d.TLSInsecure != nil && *d.TLSInsecure {
			secData["PROCX_REDIS_STREAM_TLS_INSECURE"] = "true"
		}
		if d.TLSCert != nil && *d.TLSCert != "" {
			secData["PROCX_REDIS_STREAM_TLS_CERT_FILE"] = *d.TLSCert
		}
		if d.TLSKey != nil && *d.TLSKey != "" {
			secData["PROCX_REDIS_STREAM_TLS_KEY_FILE"] = *d.TLSKey
		}
		if d.TLSCA != nil && *d.TLSCA != "" {
			secData["PROCX_REDIS_STREAM_TLS_CA_FILE"] = *d.TLSCA
		}
	}
	return secData
}

func (d *RedisStream) Metadata() map[string]string {
	md := map[string]string{}
	md["address"] = d.Host + ":" + d.Port
	if d.ConsumerName != nil {
		md["consumerName"] = *d.ConsumerName
	}
	if d.ConsumerGroup != nil {
		md["consumerGroup"] = *d.ConsumerGroup
	}
	if d.Key != nil {
		md["stream"] = *d.Key
	}
	if d.EnableTLS != nil {
		md["enableTLS"] = fmt.Sprintf("%v", *d.EnableTLS)
	}
	if d.PendingEntriesCount != nil && *d.PendingEntriesCount != "" {
		md["pendingEntriesCount"] = *d.PendingEntriesCount
	}
	return md
}

func (d *RedisStream) HasAuth() bool {
	if d.Password != nil {
		return true
	}
	if d.PasswordSecretName != nil {
		return true
	}
	return false
}

func (d *RedisStream) KedaSupport() bool {
	return true
}

func (d *RedisStream) KedaScalerName() string {
	return "redis-streams"
}

func (d *RedisStream) ContainerEnv() []corev1.EnvFromSource {
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

func (d *RedisStream) TriggerAuth(name string) *kedav1alpha1.TriggerAuthenticationSpec {
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
