package v1alpha1

type MongoDB struct {
	Host               string  `json:"host"`
	Port               int     `json:"port"`
	User               string  `json:"user"`
	Password           *string `json:"password"`
	PasswordSecretName *string `json:"passwordSecretName,omitempty"`
	DBName             string  `json:"dbName"`
	Collection         string  `json:"collection"`
	ScaleQuery         *string `json:"scaleQuery,omitempty"`
	QueryValue         *string `json:"queryValue,omitempty"`
	RetrieveQuery      *string `json:"retrieveQuery"`
	FailureQuery       *string `json:"failureQuery"`
	ClearQuery         *string `json:"clearQuery"`
	Key                *string `json:"key"`
}
