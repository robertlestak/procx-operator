package v1alpha1

import "cloud.google.com/go/firestore"

type GCPBQ struct {
	ProjectID     string  `json:"projectId"`
	RetrieveField *string `json:"retrieveField,omitempty"`
	RetrieveQuery *string `json:"retrieveQuery,omitempty"`
	ClearQuery    *string `json:"clearQuery,omitempty"`
	FailQuery     *string `json:"failQuery,omitempty"`
}

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

type FirestoreOp string

var (
	FirestoreRMOp     = FirestoreOp("rm")
	FirestoreMVOp     = FirestoreOp("mv")
	FirestoreUpdateOp = FirestoreOp("update")
)

type GCPFirestoreQuery struct {
	Path    *string              `json:"path,omitempty"`
	Op      *string              `json:"op,omitempty"`
	Value   *string              `json:"value,omitempty"`
	OrderBy *string              `json:"orderBy,omitempty"`
	Order   *firestore.Direction `json:"order,omitempty"`
}

type GCPFirestore struct {
	RetrieveCollection      *string            `json:"retrieveCollection,omitempty"`
	RetrieveDocument        *string            `json:"retrieveDocument,omitempty"`
	RetrieveQuery           *GCPFirestoreQuery `json:"retrieveQuery,omitempty"`
	RetrieveDocumentJSONKey *string            `json:"retrieveDocumentJSONKey,omitempty"`
	ClearOp                 *FirestoreOp       `json:"clearOp,omitempty"`
	ClearUpdate             *map[string]string `json:"clearUpdate,omitempty"`
	ClearCollection         *string            `json:"clearCollection,omitempty"`
	FailOp                  *FirestoreOp       `json:"failOp,omitempty"`
	FailUpdate              *map[string]string `json:"failUpdate,omitempty"`
	FailCollection          *string            `json:"failCollection,omitempty"`
	ProjectID               string             `json:"projectId"`
}
