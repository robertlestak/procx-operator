package nfs

import (
	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

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

func (d *NFS) ConfigSecret() map[string]string {
	secData := map[string]string{}
	if d.Host != nil && *d.Host != "" {
		secData["PROCX_NFS_HOST"] = *d.Host
	}
	if d.Target != nil && *d.Target != "" {
		secData["PROCX_NFS_TARGET"] = *d.Target
	}
	if d.Folder != nil && *d.Folder != "" {
		secData["PROCX_NFS_FOLDER"] = *d.Folder
	}
	if d.Key != nil && *d.Key != "" {
		secData["PROCX_NFS_KEY"] = *d.Key
	}
	if d.KeyPrefix != nil && *d.KeyPrefix != "" {
		secData["PROCX_NFS_KEY_PREFIX"] = *d.KeyPrefix
	}
	if d.KeyRegex != nil && *d.KeyRegex != "" {
		secData["PROCX_NFS_KEY_REGEX"] = *d.KeyRegex
	}
	if d.MountPath != nil && *d.MountPath != "" {
		secData["PROCX_NFS_MOUNT_PATH"] = *d.MountPath
	}
	if d.ClearOp != nil {
		if d.ClearOp.Operation != nil && *d.ClearOp.Operation != "" {
			secData["PROCX_NFS_CLEAR_OP"] = string(*d.ClearOp.Operation)
		}
		if d.ClearOp.Bucket != nil && *d.ClearOp.Bucket != "" {
			secData["PROCX_NFS_CLEAR_FOLDER"] = *d.ClearOp.Bucket
		}
		if d.ClearOp.Key != nil && *d.ClearOp.Key != "" {
			secData["PROCX_NFS_CLEAR_KEY"] = *d.ClearOp.Key
		}
		if d.ClearOp.KeyTemplate != nil && *d.ClearOp.KeyTemplate != "" {
			secData["PROCX_NFS_CLEAR_KEY_TEMPLATE"] = *d.ClearOp.KeyTemplate
		}
	}
	if d.FailOp != nil {
		if d.FailOp.Operation != nil && *d.FailOp.Operation != "" {
			secData["PROCX_NFS_FAIL_OP"] = string(*d.FailOp.Operation)
		}
		if d.FailOp.Bucket != nil && *d.FailOp.Bucket != "" {
			secData["PROCX_NFS_FAIL_FOLDER"] = *d.FailOp.Bucket
		}
		if d.FailOp.Key != nil && *d.FailOp.Key != "" {
			secData["PROCX_NFS_FAIL_KEY"] = *d.FailOp.Key
		}
		if d.FailOp.KeyTemplate != nil && *d.FailOp.KeyTemplate != "" {
			secData["PROCX_NFS_FAIL_KEY_TEMPLATE"] = *d.FailOp.KeyTemplate
		}
	}
	return secData
}

func (d *NFS) Metadata() map[string]string {
	md := map[string]string{}
	return md
}

func (d *NFS) HasAuth() bool {
	return false
}

func (d *NFS) KedaSupport() bool {
	return false
}

func (d *NFS) KedaScalerName() string {
	return ""
}

func (d *NFS) ContainerEnv() []corev1.EnvFromSource {
	envFrom := []corev1.EnvFromSource{}
	return envFrom
}

func (d *NFS) TriggerAuth(name string) *kedav1alpha1.TriggerAuthenticationSpec {
	s := &kedav1alpha1.TriggerAuthenticationSpec{}
	return s
}
