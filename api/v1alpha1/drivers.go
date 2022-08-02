package v1alpha1

type SqlQuery struct {
	Query  string   `json:"query"`
	Params []string `json:"params"`
}

type AWSDynamoDB struct {
	Region                         *string `json:"region,omitempty"`
	RoleARN                        *string `json:"roleARN,omitempty"`
	Table                          *string `json:"table,omitempty"`
	QueryKeyJSONPath               *string `json:"queryKeyJSONPath,omitempty"`
	DataJSONPath                   *string `json:"dataJSONPath,omitempty"`
	RetrieveQuery                  *string `json:"retrieveQuery,omitempty"`
	ClearQuery                     *string `json:"clearQuery,omitempty"`
	FailQuery                      *string `json:"failQuery,omitempty"`
	AccessKeySecretName            *string `json:"accessKeySecretName,omitempty"`
	IdentityOwner                  *string `json:"identityOwner,omitempty"`
	PodIdentityProvider            *string `json:"podIdentityProvider,omitempty"`
	ScaleTargetValue               *string `json:"scaleTargetValue,omitempty"`
	ScaleExpressionAttributeNames  *string `json:"scaleExpressionAttributeNames,omitempty"`
	ScaleKeyConditionExpression    *string `json:"scaleKeyConditionExpression,omitempty"`
	ScaleExpressionAttributeValues *string `json:"scaleExpressionAttributeValues,omitempty"`
}

type AWSSQS struct {
	Region              *string `json:"region,omitempty"`
	QueueURL            *string `json:"queueURL,omitempty"`
	RoleARN             *string `json:"roleARN,omitempty"`
	AccessKeySecretName *string `json:"accessKeySecretName,omitempty"`
	QueueLength         *string `json:"queueLength,omitempty"`
	IdentityOwner       *string `json:"identityOwner,omitempty"`
	PodIdentityProvider *string `json:"podIdentityProvider,omitempty"`
}

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

type GCPPubSub struct {
	ProjectID             string  `json:"projectId"`
	SubscriptionName      string  `json:"subscriptionName"`
	CredentialsSecretName *string `json:"credentialsSecretName"`
	Mode                  *string `json:"mode,omitempty"`
	Value                 *string `json:"value,omitempty"`
	PodIdentityProvider   *string `json:"podIdentityProvider,omitempty"`
}

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

type Postgres struct {
	Host               string    `json:"host"`
	Port               int       `json:"port"`
	User               string    `json:"user"`
	Password           string    `json:"password"`
	PasswordSecretName *string   `json:"passwordSecretName,omitempty"`
	DBName             string    `json:"dbName"`
	SSLMode            string    `json:"sslMode"`
	QueryReturnsKey    *bool     `json:"queryReturnsKey"`
	RetrieveQuery      *SqlQuery `json:"retrieveQuery"`
	FailureQuery       *SqlQuery `json:"failureQuery"`
	ClearQuery         *SqlQuery `json:"clearQuery"`
	Key                *string   `json:"key"`
	TargetQueryValue   *int      `json:"targetQueryValue"`
	ScaleQuery         *string   `json:"scaleQuery"`
}

type RabbitMQ struct {
	URL           string  `json:"url,omitempty"`
	URLSecretName *string `json:"urlSecretName,omitempty"`
	Queue         string  `json:"queue"`
	Mode          *string `json:"mode,omitempty"`
	Value         *string `json:"value,omitempty"`
}

type RedisList struct {
	Host               string  `json:"host"`
	Port               int     `json:"port"`
	Password           *string `json:"password,omitempty"`
	PasswordSecretName *string `json:"passwordSecretName,omitempty"`
	Key                string  `json:"key"`
	ListLength         *string `json:"listLength,omitempty"`
}
