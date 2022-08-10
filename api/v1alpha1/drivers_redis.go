package v1alpha1

type RedisList struct {
	Host               string  `json:"host"`
	Port               int     `json:"port"`
	Password           *string `json:"password,omitempty"`
	PasswordSecretName *string `json:"passwordSecretName,omitempty"`
	Key                string  `json:"key"`
	ListLength         *string `json:"listLength,omitempty"`
}
