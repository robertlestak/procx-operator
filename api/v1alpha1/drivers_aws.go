package v1alpha1

type AWSS3Operation string

var (
	AWSS3OperationRM = AWSS3Operation("rm")
	AWSS3OperationMV = AWSS3Operation("mv")
)

type AWSS3Op struct {
	Operation   AWSS3Operation `json:"operation"`
	Bucket      *string        `json:"bucket,omitempty"`
	Key         *string        `json:"key,omitempty"`
	KeyTemplate *string        `json:"keyTemplate,omitempty"`
}

type AWSS3 struct {
	Region              *string  `json:"region,omitempty"`
	RoleARN             *string  `json:"roleARN,omitempty"`
	Bucket              string   `json:"bucket"`
	Key                 *string  `json:"key,omitempty"`
	KeyRegex            *string  `json:"keyRegex,omitempty"`
	KeyPrefix           *string  `json:"keyPrefix,omitempty"`
	ClearOp             *AWSS3Op `json:"clearOp,omitempty"`
	FailOp              *AWSS3Op `json:"failOp,omitempty"`
	AccessKeySecretName *string  `json:"accessKeySecretName,omitempty"`
	IdentityOwner       *string  `json:"identityOwner,omitempty"`
	PodIdentityProvider *string  `json:"podIdentityProvider,omitempty"`
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