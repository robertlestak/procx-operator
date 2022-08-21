package v1alpha1

type HTTPRequest struct {
	Method                *string            `json:"method,omitempty"`
	URL                   string             `json:"url"`
	ContentType           *string            `json:"contentType,omitempty"`
	SuccessfulStatusCodes *[]int             `json:"successfulStatusCodes,omitempty"`
	Headers               *map[string]string `json:"headers,omitempty"`
}

type RetrieveRequest struct {
	HTTPRequest      `json:",inline"`
	KeyJSONSelector  *string `json:"keyJsonSelector,omitempty"`
	WorkJSONSelector *string `json:"workJsonSelector,omitempty"`
}

type HTTP struct {
	EnableTLS       *bool            `json:"enableTLS,omitempty"`
	TLSCA           *string          `json:"tlsCA,omitempty"`
	TLSCert         *string          `json:"tlsCert,omitempty"`
	TLSKey          *string          `json:"tlsKey,omitempty"`
	TLSSecretName   *string          `json:"tlsSecretName,omitempty"`
	TLSInsecure     *bool            `json:"tlsInsecure,omitempty"`
	RetrieveRequest *RetrieveRequest `json:"retrieveRequest,omitempty"`
	ClearRequest    *HTTPRequest     `json:"clearRequest,omitempty"`
	FailRequest     *HTTPRequest     `json:"failRequest,omitempty"`
	Key             *string          `json:"key,omitempty"`
}
