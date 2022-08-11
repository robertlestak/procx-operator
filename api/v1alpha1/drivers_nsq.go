package v1alpha1

type NSQ struct {
	NsqLookupdAddress *string `json:"nsqLookupdAddress,omitempty"`
	NsqdAddress       *string `json:"nsqdAddress,omitempty"`
	Topic             *string `json:"topic,omitempty"`
	Channel           *string `json:"channel,omitempty"`
	// TLS
	TLSSecretName *string `json:"tlsSecretName,omitempty"`
	EnableTLS     *bool   `json:"enableTLS,omitempty"`
	TLSInsecure   *bool   `json:"tlsInsecure,omitempty"`
	TLSCert       *string `json:"tlsCert,omitempty"`
	TLSKey        *string `json:"tlsKey,omitempty"`
	TLSCA         *string `json:"tlsCA,omitempty"`
}
