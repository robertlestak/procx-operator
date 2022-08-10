package v1alpha1

type MySQL struct {
	Host               string    `json:"host"`
	Port               int       `json:"port"`
	User               string    `json:"user"`
	Password           string    `json:"password"`
	PasswordSecretName *string   `json:"passwordSecretName,omitempty"`
	DBName             string    `json:"dbName"`
	QueryReturnsKey    *bool     `json:"queryReturnsKey"`
	RetrieveQuery      *SqlQuery `json:"retrieveQuery"`
	FailureQuery       *SqlQuery `json:"failureQuery"`
	ClearQuery         *SqlQuery `json:"clearQuery"`
	Key                *string   `json:"key"`
	TargetQueryValue   *int      `json:"targetQueryValue"`
	ScaleQuery         *string   `json:"scaleQuery"`
}
