package v1alpha1

type GCPPubSub struct {
	ProjectID             string  `json:"projectId"`
	SubscriptionName      string  `json:"subscriptionName"`
	CredentialsSecretName *string `json:"credentialsSecretName"`
	Mode                  *string `json:"mode,omitempty"`
	Value                 *string `json:"value,omitempty"`
	PodIdentityProvider   *string `json:"podIdentityProvider,omitempty"`
}

type GCPGCSOperation string

var (
	GCPGCSOperationRM = GCPGCSOperation("rm")
	GCPGCSOperationMV = GCPGCSOperation("mv")
)

type GCPGCSOp struct {
	Operation   *GCPGCSOperation `json:"operation,omitempty"`
	Bucket      *string          `json:"bucket,omitempty"`
	Key         *string          `json:"key,omitempty"`
	KeyTemplate *string          `json:"keyTemplate,omitempty"`
}

type GCPGCS struct {
	ProjectID             *string   `json:"projectId,omitempty"`
	Bucket                *string   `json:"bucket,omitempty"`
	Key                   *string   `json:"key,omitempty"`
	KeyRegex              *string   `json:"keyRegex,omitempty"`
	KeyPrefix             *string   `json:"keyPrefix,omitempty"`
	ClearOp               *GCPGCSOp `json:"clearOp,omitempty"`
	FailOp                *GCPGCSOp `json:"failOp,omitempty"`
	CredentialsSecretName *string   `json:"credentialsSecretName"`
	PodIdentityProvider   *string   `json:"podIdentityProvider,omitempty"`
	TargetObjectCount     *string   `json:"targetObjectCount,omitempty"`
	ActivationObjectCount *string   `json:"activationObjectCount,omitempty"`
	MaxBucketItemsToScan  *string   `json:"maxBucketItemsToScan,omitempty"`
}
