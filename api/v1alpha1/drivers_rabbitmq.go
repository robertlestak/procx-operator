package v1alpha1

type RabbitMQ struct {
	URL           string  `json:"url,omitempty"`
	URLSecretName *string `json:"urlSecretName,omitempty"`
	Queue         string  `json:"queue"`
	Mode          *string `json:"mode,omitempty"`
	Value         *string `json:"value,omitempty"`
}
