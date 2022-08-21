package v1alpha1

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
