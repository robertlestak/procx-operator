package v1alpha1

type Cassandra struct {
	Hosts              []string  `json:"hosts"`
	User               *string   `json:"user"`
	Password           *string   `json:"password"`
	PasswordSecretName *string   `json:"passwordSecretName,omitempty"`
	Keyspace           string    `json:"keyspace"`
	Consistency        string    `json:"consistency"`
	QueryReturnsKey    *bool     `json:"queryReturnsKey"`
	RetrieveQuery      *SqlQuery `json:"retrieveQuery"`
	FailureQuery       *SqlQuery `json:"failureQuery"`
	ClearQuery         *SqlQuery `json:"clearQuery"`
	ScaleQuery         *string   `json:"scaleQuery"`
	TargetQueryValue   *string   `json:"targetQueryValue"`
	Key                *string   `json:"key"`
}
