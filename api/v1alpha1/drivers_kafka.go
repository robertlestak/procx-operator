package v1alpha1

type KafkaSaslType string

var (
	KafkaSaslTypePlain = KafkaSaslType("plain")
	KafkaSaslTypeScram = KafkaSaslType("scram")
)

type Kafka struct {
	Brokers *[]string
	Group   *string
	Topic   *string
	// TLS
	EnableTLS   *bool
	TLSInsecure *bool
	TLSCert     *string
	TLSKey      *string
	TLSCA       *string
	// SASL
	EnableSASL                 *bool
	SaslType                   *KafkaSaslType
	Username                   *string
	Password                   *string
	PasswordSecretName         *string
	LagThreshold               *string
	ActivationThreshold        *string
	OffsetResetPolicy          *string
	AllowIdleConsumers         *bool
	ScaleToZeroOnInvalidOffset *bool
	Version                    *string
}
