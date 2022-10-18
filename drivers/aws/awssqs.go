package aws

import (
	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

type AWSSQS struct {
	Region              *string `json:"region,omitempty"`
	QueueURL            *string `json:"queueURL,omitempty"`
	RoleARN             *string `json:"roleARN,omitempty"`
	AuthRoleARN         *string `json:"authRoleARN,omitempty"`
	AccessKeySecretName *string `json:"accessKeySecretName,omitempty"`
	QueueLength         *string `json:"queueLength,omitempty"`
	IdentityOwner       *string `json:"identityOwner,omitempty"`
	PodIdentityProvider *string `json:"podIdentityProvider,omitempty"`
}

func (d *AWSSQS) ConfigSecret() map[string]string {
	secData := map[string]string{}
	if d.Region != nil && *d.Region != "" {
		secData["PROCX_AWS_REGION"] = *d.Region
	}
	if d.QueueURL != nil && *d.QueueURL != "" {
		secData["PROCX_AWS_SQS_QUEUE_URL"] = *d.QueueURL
	}
	if d.RoleARN != nil && *d.RoleARN != "" {
		secData["PROCX_AWS_ROLE_ARN"] = *d.RoleARN
	}
	if d.AuthRoleARN != nil && *d.AuthRoleARN != "" {
		secData["PROCX_AWS_AUTH_ROLE_ARN"] = *d.AuthRoleARN
	}
	return secData
}

func (d *AWSSQS) KedaSupport() bool {
	return true
}

func (d *AWSSQS) HasAuth() bool {
	return true
}

func (d *AWSSQS) KedaScalerName() string {
	return "aws-sqs-queue"
}

func (d *AWSSQS) TriggerAuth(name string) *kedav1alpha1.TriggerAuthenticationSpec {
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
	if d.AuthRoleARN != nil && *d.AuthRoleARN != "" {
		s.SecretTargetRef = append(s.SecretTargetRef, kedav1alpha1.AuthSecretTargetRef{
			Name:      name,
			Parameter: "awsRoleArn",
			Key:       "PROCX_AWS_AUTH_ROLE_ARN",
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

func (d *AWSSQS) Metadata() map[string]string {
	md := map[string]string{}
	if d.RoleARN != nil && *d.RoleARN != "" {
		md["awsRoleArn"] = *d.RoleARN
	}
	if d.AuthRoleARN != nil && *d.AuthRoleARN != "" {
		md["awsRoleArn"] = *d.AuthRoleARN
	}
	if d.QueueURL != nil && *d.QueueURL != "" {
		md["queueURL"] = *d.QueueURL
	}
	if d.Region != nil && *d.Region != "" {
		md["awsRegion"] = *d.Region
	}
	if d.QueueLength == nil {
		md["queueLength"] = "1"
	} else {
		md["queueLength"] = *d.QueueLength
	}
	if d.IdentityOwner == nil {
		md["identityOwner"] = "pod"
	} else {
		md["identityOwner"] = *d.IdentityOwner
	}
	return md
}

func (d *AWSSQS) ContainerEnv() []corev1.EnvFromSource {
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

func (d *AWSSQS) VolumeMounts() []corev1.VolumeMount {
	return nil
}

func (d *AWSSQS) Volumes() []corev1.Volume {
	return nil
}
