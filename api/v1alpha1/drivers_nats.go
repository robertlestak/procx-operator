package v1alpha1

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
	TLSSecretName      *string `json:"tlsSecretName,omitempty"`
	EnableTLS          *bool   `json:"enableTLS,omitempty"`
	TLSInsecure        *bool   `json:"tlsInsecure,omitempty"`
	TLSCA              *string `json:"tlsCA,omitempty"`
	TLSCert            *string `json:"tlsCert,omitempty"`
	TLSKey             *string `json:"tlsKey,omitempty"`
	ClearResponse      *string `json:"clearResponse,omitempty"`
	FailResponse       *string `json:"failResponse,omitempty"`
}
