package elasticsearch

import (
	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

type ElasticsearchCloseOp string

var (
	ElasticsearchCloseOpDelete   = ElasticsearchCloseOp("delete")
	ElasticsearchCloseOpPut      = ElasticsearchCloseOp("put")
	ElasticsearchCloseOpMergePut = ElasticsearchCloseOp("merge-put")
	ElasticsearchCloseOpMove     = ElasticsearchCloseOp("move")
)

type Elasticsearch struct {
	Address            *string `json:"address,omitempty"`
	Username           *string `json:"username,omitempty"`
	Password           *string `json:"password,omitempty"`
	PasswordSecretName *string `json:"passwordSecretName,omitempty"`
	// TLS
	EnableTLS                           *bool                 `json:"enableTLS,omitempty"`
	TLSInsecure                         *bool                 `json:"tlsInsecure,omitempty"`
	TLSCert                             *string               `json:"tlsCert,omitempty"`
	TLSKey                              *string               `json:"tlsKey,omitempty"`
	TLSCA                               *string               `json:"tlsCA,omitempty"`
	RetrieveIndex                       *string               `json:"retrieveIndex,omitempty"`
	RetrieveSearchTemplate              *string               `json:"retrieveSearchTemplate,omitempty"`
	RetrieveSearchParams                *string               `json:"retrieveSearchParams,omitempty"`
	RetrieveSearchValueLocation         *string               `json:"retrieveSearchValueLocation,omitempty"`
	RetrieveSearchTargetValue           *string               `json:"retrieveSearchTargetValue,omitempty"`
	RetrieveSearchActivationTargetValue *string               `json:"retrieveSearchActivationTargetValue,omitempty"`
	RetrieveQuery                       *string               `json:"retrieveQuery,omitempty"`
	ClearQuery                          *string               `json:"clearQuery,omitempty"`
	ClearIndex                          *string               `json:"clearIndex,omitempty"`
	ClearOp                             *ElasticsearchCloseOp `json:"clearOp,omitempty"`
	FailQuery                           *string               `json:"failQuery,omitempty"`
	FailIndex                           *string               `json:"failIndex,omitempty"`
	FailOp                              *ElasticsearchCloseOp `json:"failOp,omitempty"`
	Key                                 *string               `json:"key,omitempty"`
}

func (d *Elasticsearch) ConfigSecret() map[string]string {
	secData := map[string]string{}
	if d.Password != nil && *d.Password != "" {
		secData["PROCX_ELASTICSEARCH_PASSWORD"] = *d.Password
	}
	if d.Username != nil && *d.Username != "" {
		secData["PROCX_ELASTICSEARCH_USERNAME"] = *d.Username
	}
	if d.Address != nil && *d.Address != "" {
		secData["PROCX_ELASTICSEARCH_ADDRESS"] = *d.Address
	}
	if d.EnableTLS != nil && *d.EnableTLS {
		secData["PROCX_ELASTICSEARCH_ENABLE_TLS"] = "true"
	}
	if d.TLSInsecure != nil && *d.TLSInsecure {
		secData["PROCX_ELASTICSEARCH_TLS_SKIP_VERIFY"] = "true"
	}
	if d.TLSCert != nil && *d.TLSCert != "" {
		secData["PROCX_ELASTICSEARCH_TLS_CERT_FILE"] = *d.TLSCert
	}
	if d.TLSKey != nil && *d.TLSKey != "" {
		secData["PROCX_ELASTICSEARCH_TLS_KEY_FILE"] = *d.TLSKey
	}
	if d.TLSCA != nil && *d.TLSCA != "" {
		secData["PROCX_ELASTICSEARCH_TLS_CA_FILE"] = *d.TLSCA
	}
	if d.RetrieveIndex != nil && *d.RetrieveIndex != "" {
		secData["PROCX_ELASTICSEARCH_RETRIEVE_INDEX"] = *d.RetrieveIndex
	}
	if d.RetrieveQuery != nil && *d.RetrieveQuery != "" {
		secData["PROCX_ELASTICSEARCH_RETRIEVE_QUERY"] = *d.RetrieveQuery
	}
	if d.ClearQuery != nil && *d.ClearQuery != "" {
		secData["PROCX_ELASTICSEARCH_CLEAR_QUERY"] = *d.ClearQuery
	}
	if d.ClearIndex != nil && *d.ClearIndex != "" {
		secData["PROCX_ELASTICSEARCH_CLEAR_INDEX"] = *d.ClearIndex
	}
	if d.FailQuery != nil && *d.FailQuery != "" {
		secData["PROCX_ELASTICSEARCH_FAIL_QUERY"] = *d.FailQuery
	}
	if d.FailIndex != nil && *d.FailIndex != "" {
		secData["PROCX_ELASTICSEARCH_FAIL_INDEX"] = *d.FailIndex
	}
	if d.ClearOp != nil && *d.ClearOp != "" {
		secData["PROCX_ELASTICSEARCH_CLEAR_OP"] = string(*d.ClearOp)
	}
	if d.FailOp != nil && *d.FailOp != "" {
		secData["PROCX_ELASTICSEARCH_FAIL_OP"] = string(*d.FailOp)
	}
	return secData
}

func (d *Elasticsearch) Metadata() map[string]string {
	md := map[string]string{
		"addresses": *d.Address,
		"username":  *d.Username,
		"index":     *d.RetrieveIndex,
	}
	if d.RetrieveSearchTemplate != nil {
		md["searchTemplate"] = *d.RetrieveSearchTemplate
	}
	if d.RetrieveSearchParams != nil {
		md["params"] = *d.RetrieveSearchParams
	}
	if d.RetrieveSearchValueLocation != nil {
		md["valueLocation"] = *d.RetrieveSearchValueLocation
	}
	if d.RetrieveSearchTargetValue != nil {
		md["targetValue"] = *d.RetrieveSearchTargetValue
	}
	if d.RetrieveSearchActivationTargetValue != nil {
		md["activationTargetValue"] = *d.RetrieveSearchActivationTargetValue
	}
	if d.TLSInsecure != nil && *d.TLSInsecure {
		md["unsafeSsl"] = "true"
	}
	return md
}

func (d *Elasticsearch) HasAuth() bool {
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

func (d *Elasticsearch) KedaSupport() bool {
	return true
}

func (d *Elasticsearch) KedaScalerName() string {
	return "elasticsearch"
}

func (d *Elasticsearch) ContainerEnv() []corev1.EnvFromSource {
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

func (d *Elasticsearch) TriggerAuth(name string) *kedav1alpha1.TriggerAuthenticationSpec {
	s := &kedav1alpha1.TriggerAuthenticationSpec{}
	if d.PasswordSecretName != nil {
		s.SecretTargetRef = append(s.SecretTargetRef, kedav1alpha1.AuthSecretTargetRef{
			Name:      *d.PasswordSecretName,
			Key:       "PROCX_ELASTICSEARCH_PASSWORD",
			Parameter: "password",
		})
	} else if d.Password != nil && *d.Password != "" {
		s.SecretTargetRef = append(s.SecretTargetRef, kedav1alpha1.AuthSecretTargetRef{
			Name:      name,
			Key:       "PROCX_ELASTICSEARCH_PASSWORD",
			Parameter: "password",
		})
	}
	return s
}
