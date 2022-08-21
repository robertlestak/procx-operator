package v1alpha1

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
