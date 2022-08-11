package v1alpha1

type KafkaSaslType string

var (
	KafkaSaslTypePlain = KafkaSaslType("plain")
	KafkaSaslTypeScram = KafkaSaslType("scram")
)

type Kafka struct {
	Brokers *[]string `json:"brokers,omitempty"`
	Group   *string   `json:"group,omitempty"`
	Topic   *string   `json:"topic,omitempty"`
	// TLS
	TLSSecretName *string `json:"tlsSecretName,omitempty"`
	EnableTLS     *bool   `json:"enableTLS,omitempty"`
	TLSInsecure   *bool   `json:"tlsInsecure,omitempty"`
	TLSCert       *string `json:"tlsCert,omitempty"`
	TLSKey        *string `json:"tlsKey,omitempty"`
	TLSCA         *string `json:"tlsCA,omitempty"`
	// SASL
	EnableSASL                 *bool          `json:"enableSASL,omitempty"`
	SaslType                   *KafkaSaslType `json:"saslType,omitempty"`
	Username                   *string        `json:"username,omitempty"`
	Password                   *string        `json:"password,omitempty"`
	PasswordSecretName         *string        `json:"passwordSecretName,omitempty"`
	LagThreshold               *string        `json:"lagThreshold,omitempty"`
	ActivationThreshold        *string        `json:"activationThreshold,omitempty"`
	OffsetResetPolicy          *string        `json:"offsetResetPolicy,omitempty"`
	AllowIdleConsumers         *bool          `json:"allowIdleConsumers,omitempty"`
	ScaleToZeroOnInvalidOffset *bool          `json:"scaleToZeroOnInvalidOffset,omitempty"`
	Version                    *string        `json:"version,omitempty"`
}
