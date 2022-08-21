package v1alpha1

type FSOperation string

var (
	FSOperationRM = NFSOperation("rm")
	FSOperationMV = NFSOperation("mv")
)

type FSOp struct {
	Operation   *FSOperation `json:"operation,omitempty"`
	Bucket      *string      `json:"bucket,omitempty"`
	Key         *string      `json:"key,omitempty"`
	KeyTemplate *string      `json:"keyTemplate,omitempty"`
}

type FS struct {
	Folder    *string `json:"folder,omitempty"`
	Key       *string `json:"key,omitempty"`
	KeyPrefix *string `json:"keyPrefix,omitempty"`
	KeyRegex  *string `json:"keyRegex,omitempty"`
	ClearOp   *NFSOp  `json:"clearOp,omitempty"`
	FailOp    *NFSOp  `json:"failOp,omitempty"`
}
