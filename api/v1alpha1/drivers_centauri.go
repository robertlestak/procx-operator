package v1alpha1

type Centauri struct {
	URL                  string  `json:"url"`
	PrivateKey           *[]byte `json:"privateKey,omitempty"`
	PrivateKeySecretName *string `json:"privateKeySecretName,omitempty"`
	TLSSecretName        *string `json:"tlsSecretName,omitempty"`
	Channel              *string `json:"channel,omitempty"`
	Key                  *string `json:"key,omitempty"`
}
