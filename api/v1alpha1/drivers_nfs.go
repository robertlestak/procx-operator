package v1alpha1

type NFSOperation string

var (
	NFSOperationRM = NFSOperation("rm")
	NFSOperationMV = NFSOperation("mv")
)

type NFSOp struct {
	Operation   *NFSOperation `json:"operation,omitempty"`
	Bucket      *string       `json:"bucket,omitempty"`
	Key         *string       `json:"key,omitempty"`
	KeyTemplate *string       `json:"keyTemplate,omitempty"`
}

type NFS struct {
	Host      *string `json:"host,omitempty"`
	Target    *string `json:"target,omitempty"`
	Folder    *string `json:"folder,omitempty"`
	Key       *string `json:"key,omitempty"`
	KeyPrefix *string `json:"keyPrefix,omitempty"`
	KeyRegex  *string `json:"keyRegex,omitempty"`
	MountPath *string `json:"mountPath,omitempty"`
	ClearOp   *NFSOp  `json:"clearOp,omitempty"`
	FailOp    *NFSOp  `json:"failOp,omitempty"`
}
