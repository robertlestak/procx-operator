package gcp

import (
	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

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

func (d *GCPGCS) ConfigSecret() map[string]string {
	secData := map[string]string{}
	if d.ProjectID != nil {
		secData["PROCX_GCP_PROJECT_ID"] = *d.ProjectID
	}
	if d.Bucket != nil {
		secData["PROCX_GCP_GCS_BUCKET"] = *d.Bucket
	}
	if d.Key != nil {
		secData["PROCX_GCP_GCS_KEY"] = *d.Key
	}
	if d.KeyRegex != nil {
		secData["PROCX_GCP_GCS_KEY_REGEX"] = *d.KeyRegex
	}
	if d.KeyPrefix != nil {
		secData["PROCX_GCP_GCS_KEY_PREFIX"] = *d.KeyPrefix
	}
	if d.ClearOp != nil {
		if d.ClearOp.Operation != nil {
			secData["PROCX_GCP_GCS_CLEAR_OP"] = string(*d.ClearOp.Operation)
		}
		if d.ClearOp.Bucket != nil {
			secData["PROCX_GCP_GCS_CLEAR_BUCKET"] = *d.ClearOp.Bucket
		}
		if d.ClearOp.Key != nil {
			secData["PROCX_GCP_GCS_CLEAR_KEY"] = *d.ClearOp.Key
		}
		if d.ClearOp.KeyTemplate != nil {
			secData["PROCX_GCP_GCS_CLEAR_KEY_TEMPLATE"] = *d.ClearOp.KeyTemplate
		}
	}
	if d.FailOp != nil {
		if d.FailOp.Operation != nil {
			secData["PROCX_GCP_GCS_FAIL_OP"] = string(*d.FailOp.Operation)
		}
		if d.FailOp.Bucket != nil {
			secData["PROCX_GCP_GCS_FAIL_BUCKET"] = *d.FailOp.Bucket
		}
		if d.FailOp.Key != nil {
			secData["PROCX_GCP_GCS_FAIL_KEY"] = *d.FailOp.Key
		}
		if d.FailOp.KeyTemplate != nil {
			secData["PROCX_GCP_GCS_FAIL_KEY_TEMPLATE"] = *d.FailOp.KeyTemplate
		}
	}
	return secData
}

func (d *GCPGCS) Metadata() map[string]string {
	md := map[string]string{
		"bucketName": *d.Bucket,
	}
	if d.TargetObjectCount != nil && *d.TargetObjectCount != "" {
		md["targetObjectCount"] = *d.TargetObjectCount
	}
	if d.ActivationObjectCount != nil && *d.ActivationObjectCount != "" {
		md["activationObjectCount"] = *d.ActivationObjectCount
	}
	if d.MaxBucketItemsToScan != nil && *d.MaxBucketItemsToScan != "" {
		md["maxBucketItemsToScan"] = *d.MaxBucketItemsToScan
	}
	return md
}

func (d *GCPGCS) HasAuth() bool {
	return true
}

func (d *GCPGCS) KedaSupport() bool {
	return true
}

func (d *GCPGCS) KedaScalerName() string {
	return "gcp-storage"
}

func (d *GCPGCS) ContainerEnv() []corev1.EnvFromSource {
	return []corev1.EnvFromSource{}
}

func (d *GCPGCS) TriggerAuth(name string) *kedav1alpha1.TriggerAuthenticationSpec {
	s := &kedav1alpha1.TriggerAuthenticationSpec{}
	if d.CredentialsSecretName != nil {
		s.SecretTargetRef = append(s.SecretTargetRef, kedav1alpha1.AuthSecretTargetRef{
			Name:      *d.CredentialsSecretName,
			Key:       "GOOGLE_APPLICATION_CREDENTIALS_JSON",
			Parameter: "GoogleApplicationCredentials",
		})
	}
	if d.PodIdentityProvider != nil {
		s.PodIdentity = &kedav1alpha1.AuthPodIdentity{
			Provider: kedav1alpha1.PodIdentityProvider(*d.PodIdentityProvider),
		}
	}
	return s
}
