package v1alpha1

type RedisList struct {
	Host               string  `json:"host"`
	Port               int     `json:"port"`
	Password           *string `json:"password,omitempty"`
	PasswordSecretName *string `json:"passwordSecretName,omitempty"`
	Key                string  `json:"key"`
	ListLength         *string `json:"listLength,omitempty"`
	// TLS
	TLSSecretName *string `json:"tlsSecretName,omitempty"`
	EnableTLS     *bool   `json:"enableTLS,omitempty"`
	TLSInsecure   *bool   `json:"tlsInsecure,omitempty"`
	TLSCert       *string `json:"tlsCert,omitempty"`
	TLSKey        *string `json:"tlsKey,omitempty"`
	TLSCA         *string `json:"tlsCA,omitempty"`
}

type RedisPubSub struct {
	Host               string  `json:"host"`
	Port               string  `json:"port"`
	Password           *string `json:"password,omitempty"`
	PasswordSecretName *string `json:"passwordSecretName,omitempty"`
	Key                *string `json:"key,omitempty"`
	// TLS
	TLSSecretName *string `json:"tlsSecretName,omitempty"`
	EnableTLS     *bool   `json:"enableTLS,omitempty"`
	TLSInsecure   *bool   `json:"tlsInsecure,omitempty"`
	TLSCert       *string `json:"tlsCert,omitempty"`
	TLSKey        *string `json:"tlsKey,omitempty"`
	TLSCA         *string `json:"tlsCA,omitempty"`
}

type RedisStreamOp string

var (
	RedisStreamOpAck = RedisStreamOp("ack")
	RedisStreamOpDel = RedisStreamOp("del")
)

type RedisStream struct {
	Host               string         `json:"host"`
	Port               string         `json:"port"`
	Password           *string        `json:"password,omitempty"`
	PasswordSecretName *string        `json:"passwordSecretName,omitempty"`
	ConsumerName       *string        `json:"consumerName,omitempty"`
	ConsumerGroup      *string        `json:"consumerGroup,omitempty"`
	Key                *string        `json:"key"`
	ValueKeys          *[]string      `json:"valueKeys,omitempty"`
	MessageID          *string        `json:"messageID,omitempty"`
	ClearOp            *RedisStreamOp `json:"clearOp,omitempty"`
	FailOp             *RedisStreamOp `json:"failOp,omitempty"`
	// TLS
	TLSSecretName       *string `json:"tlsSecretName,omitempty"`
	EnableTLS           *bool   `json:"enableTLS,omitempty"`
	TLSInsecure         *bool   `json:"tlsInsecure,omitempty"`
	TLSCert             *string `json:"tlsCert,omitempty"`
	TLSKey              *string `json:"tlsKey,omitempty"`
	TLSCA               *string `json:"tlsCA,omitempty"`
	PendingEntriesCount *string `json:"pendingEntriesCount,omitempty"`
}
