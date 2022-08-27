package v1alpha1

type ElasticsearchCloseOp string

var (
	ElasticsearchCloseOpDelete   = ElasticsearchCloseOp("delete")
	ElasticsearchCloseOpPut      = ElasticsearchCloseOp("put")
	ElasticsearchCloseOpMergePut = ElasticsearchCloseOp("merge-put")
	ElasticsearchCloseOpMove     = ElasticsearchCloseOp("move")
)

type Elasticsearch struct {
	Address                             *string `json:"address,omitempty"`
	Username                            *string `json:"username,omitempty"`
	Password                            *string `json:"password,omitempty"`
	PasswordSecretName                  *string `json:"passwordSecretName,omitempty"`
	RetrieveSearchTemplate              *string `json:"retrieveSearchTemplate,omitempty"`
	RetrieveSearchParams                *string `json:"retrieveSearchParams,omitempty"`
	RetrieveSearchValueLocation         *string `json:"retrieveSearchValueLocation,omitempty"`
	RetrieveSearchTargetValue           *string `json:"retrieveSearchTargetValue,omitempty"`
	RetrieveSearchActivationTargetValue *string `json:"retrieveSearchActivationTargetValue,omitempty"`
	// TLS
	EnableTLS     *bool                 `json:"enableTLS,omitempty"`
	TLSInsecure   *bool                 `json:"tlsInsecure,omitempty"`
	TLSCert       *string               `json:"tlsCert,omitempty"`
	TLSKey        *string               `json:"tlsKey,omitempty"`
	TLSCA         *string               `json:"tlsCA,omitempty"`
	TLSSecretName *string               `json:"tlsSecretName,omitempty"`
	RetrieveIndex *string               `json:"retrieveIndex,omitempty"`
	RetrieveQuery *string               `json:"retrieveQuery,omitempty"`
	ClearDoc      *string               `json:"clearDoc,omitempty"`
	ClearIndex    *string               `json:"clearIndex,omitempty"`
	ClearOp       *ElasticsearchCloseOp `json:"clearOp,omitempty"`
	FailDoc       *string               `json:"failDoc,omitempty"`
	FailIndex     *string               `json:"failIndex,omitempty"`
	FailOp        *ElasticsearchCloseOp `json:"failOp,omitempty"`
	Key           *string               `json:"key,omitempty"`
}
