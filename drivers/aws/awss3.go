package aws

import (
	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

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

func (d *AWSS3) ConfigSecret() map[string]string {
	secData := map[string]string{}
	if d.Region != nil && *d.Region != "" {
		secData["PROCX_AWS_REGION"] = *d.Region
	}
	if d.RoleARN != nil && *d.RoleARN != "" {
		secData["PROCX_AWS_ROLE_ARN"] = *d.RoleARN
	}
	if d.Bucket != "" {
		secData["PROCX_AWS_S3_BUCKET"] = d.Bucket
	}
	if d.Key != nil && *d.Key != "" {
		secData["PROCX_AWS_S3_KEY"] = *d.Key
	}
	if d.KeyRegex != nil && *d.KeyRegex != "" {
		secData["PROCX_AWS_S3_KEY_REGEX"] = *d.KeyRegex
	}
	if d.KeyPrefix != nil && *d.KeyPrefix != "" {
		secData["PROCX_AWS_S3_KEY_PREFIX"] = *d.KeyPrefix
	}
	if d.ClearOp != nil {
		if d.ClearOp.Operation != "" {
			secData["PROCX_AWS_S3_CLEAR_OP"] = string(d.ClearOp.Operation)
		}
		if d.ClearOp.Bucket != nil && *d.ClearOp.Bucket != "" {
			secData["PROCX_AWS_S3_CLEAR_BUCKET"] = *d.ClearOp.Bucket
		}
		if d.ClearOp.Key != nil && *d.ClearOp.Key != "" {
			secData["PROCX_AWS_S3_CLEAR_KEY"] = *d.ClearOp.Key
		}
		if d.ClearOp.KeyTemplate != nil && *d.ClearOp.KeyTemplate != "" {
			secData["PROCX_AWS_S3_CLEAR_KEY_TEMPLATE"] = *d.ClearOp.KeyTemplate
		}
	}
	if d.FailOp != nil {
		if d.FailOp.Operation != "" {
			secData["PROCX_AWS_S3_FAIL_OP"] = string(d.FailOp.Operation)
		}
		if d.FailOp.Bucket != nil && *d.FailOp.Bucket != "" {
			secData["PROCX_AWS_S3_FAIL_BUCKET"] = *d.FailOp.Bucket
		}
		if d.FailOp.Key != nil && *d.FailOp.Key != "" {
			secData["PROCX_AWS_S3_FAIL_KEY"] = *d.FailOp.Key
		}
		if d.FailOp.KeyTemplate != nil && *d.FailOp.KeyTemplate != "" {
			secData["PROCX_AWS_S3_FAIL_KEY_TEMPLATE"] = *d.FailOp.KeyTemplate
		}
	}
	return secData
}

func (d *AWSS3) KedaSupport() bool {
	return false
}

func (d *AWSS3) HasAuth() bool {
	return true
}

func (d *AWSS3) KedaScalerName() string {
	return ""
}

func (d *AWSS3) TriggerAuth(name string) *kedav1alpha1.TriggerAuthenticationSpec {
	s := &kedav1alpha1.TriggerAuthenticationSpec{}
	if d.PodIdentityProvider != nil {
		s.PodIdentity = &kedav1alpha1.AuthPodIdentity{
			Provider: kedav1alpha1.PodIdentityProvider(*d.PodIdentityProvider),
		}
	}
	if d.RoleARN != nil && *d.RoleARN != "" {
		s.SecretTargetRef = append(s.SecretTargetRef, kedav1alpha1.AuthSecretTargetRef{
			Name:      name,
			Parameter: "awsRoleArn",
			Key:       "PROCX_AWS_ROLE_ARN",
		})
	}
	if d.AccessKeySecretName != nil {
		s.SecretTargetRef = append(s.SecretTargetRef, kedav1alpha1.AuthSecretTargetRef{
			Name:      *d.AccessKeySecretName,
			Parameter: "awsAccessKeyID",
			Key:       "AWS_ACCESS_KEY_ID",
		})
		s.SecretTargetRef = append(s.SecretTargetRef, kedav1alpha1.AuthSecretTargetRef{
			Name:      *d.AccessKeySecretName,
			Parameter: "awsSecretAccessKey",
			Key:       "AWS_SECRET_ACCESS_KEY",
		})
	}
	return s
}

func (d *AWSS3) Metadata() map[string]string {
	md := map[string]string{}
	if d.RoleARN != nil && *d.RoleARN != "" {
		md["awsRoleArn"] = *d.RoleARN
	}
	if d.IdentityOwner == nil {
		md["identityOwner"] = "pod"
	} else {
		md["identityOwner"] = *d.IdentityOwner
	}
	return md
}

func (d *AWSS3) ContainerEnv() []corev1.EnvFromSource {
	falseVal := false
	envFrom := []corev1.EnvFromSource{}
	if d.AccessKeySecretName != nil {
		envFrom = append(envFrom, corev1.EnvFromSource{
			SecretRef: &corev1.SecretEnvSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: *d.AccessKeySecretName,
				},
				Optional: &falseVal,
			},
		})
	}
	return envFrom
}
