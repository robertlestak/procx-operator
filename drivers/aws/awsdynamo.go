package aws

import (
	"strconv"

	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

type AWSDynamoDB struct {
	Region                         *string `json:"region,omitempty"`
	RoleARN                        *string `json:"roleARN,omitempty"`
	Table                          *string `json:"table,omitempty"`
	RetrieveQuery                  *string `json:"retrieveQuery,omitempty"`
	ClearQuery                     *string `json:"clearQuery,omitempty"`
	FailQuery                      *string `json:"failQuery,omitempty"`
	IncludeNextToken               *bool   `json:"includeNextToken,omitempty"`
	Limit                          *int64  `json:"limit,omitempty"`
	NextToken                      *string `json:"nextToken,omitempty"`
	RetrieveField                  *string `json:"retrieveField,omitempty"`
	UnmarshalJSON                  *bool   `json:"unmarshalJSON,omitempty"`
	AccessKeySecretName            *string `json:"accessKeySecretName,omitempty"`
	IdentityOwner                  *string `json:"identityOwner,omitempty"`
	PodIdentityProvider            *string `json:"podIdentityProvider,omitempty"`
	ScaleTargetValue               *string `json:"scaleTargetValue,omitempty"`
	ScaleExpressionAttributeNames  *string `json:"scaleExpressionAttributeNames,omitempty"`
	ScaleKeyConditionExpression    *string `json:"scaleKeyConditionExpression,omitempty"`
	ScaleExpressionAttributeValues *string `json:"scaleExpressionAttributeValues,omitempty"`
}

func (d *AWSDynamoDB) ConfigSecret() map[string]string {
	secData := map[string]string{}
	if d.Region != nil && *d.Region != "" {
		secData["PROCX_AWS_REGION"] = *d.Region
	}
	if d.RoleARN != nil && *d.RoleARN != "" {
		secData["PROCX_AWS_ROLE_ARN"] = *d.RoleARN
	}
	if d.Table != nil && *d.Table != "" {
		secData["PROCX_AWS_DYNAMODB_TABLE"] = *d.Table
	}
	if d.IncludeNextToken != nil && *d.IncludeNextToken {
		secData["PROCX_AWS_DYNAMO_INCLUDE_NEXT_TOKEN"] = "true"
	}
	if d.Limit != nil && *d.Limit > 0 {
		secData["PROCX_AWS_DYNAMO_LIMIT"] = strconv.Itoa(int(*d.Limit))
	}
	if d.NextToken != nil && *d.NextToken != "" {
		secData["PROCX_AWS_DYNAMO_NEXT_TOKEN"] = *d.NextToken
	}
	if d.RetrieveField != nil && *d.RetrieveField != "" {
		secData["PROCX_AWS_DYNAMO_RETRIEVE_FIELD"] = *d.RetrieveField
	}
	if d.RetrieveQuery != nil {
		secData["PROCX_AWS_DYNAMO_RETRIEVE_QUERY"] = *d.RetrieveQuery
	}
	if d.ClearQuery != nil {
		secData["PROCX_AWS_DYNAMO_CLEAR_QUERY"] = *d.ClearQuery
	}
	if d.FailQuery != nil {
		secData["PROCX_AWS_DYNAMO_FAIL_QUERY"] = *d.FailQuery
	}
	if d.UnmarshalJSON != nil && *d.UnmarshalJSON {
		secData["PROCX_AWS_DYNAMO_UNMARSHAL_JSON"] = "true"
	}
	return secData
}

func (d *AWSDynamoDB) KedaSupport() bool {
	return true
}

func (d *AWSDynamoDB) HasAuth() bool {
	return true
}

func (d *AWSDynamoDB) KedaScalerName() string {
	return "aws-dynamodb"
}

func (d *AWSDynamoDB) TriggerAuth(name string) *kedav1alpha1.TriggerAuthenticationSpec {
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

func (d *AWSDynamoDB) Metadata() map[string]string {
	md := map[string]string{
		"expressionAttributeNames":  *d.ScaleExpressionAttributeNames,
		"expressionAttributeValues": *d.ScaleExpressionAttributeValues,
		"keyConditionExpression":    *d.ScaleKeyConditionExpression,
	}
	if d.RoleARN != nil && *d.RoleARN != "" {
		md["awsRoleArn"] = *d.RoleARN
	}
	if d.Table != nil && *d.Table != "" {
		md["tableName"] = *d.Table
	}
	if d.Region != nil && *d.Region != "" {
		md["awsRegion"] = *d.Region
	}
	if d.ScaleTargetValue == nil {
		md["targetValue"] = "1"
	} else {
		md["targetValue"] = *d.ScaleTargetValue
	}
	if d.IdentityOwner == nil {
		md["identityOwner"] = "pod"
	} else {
		md["identityOwner"] = *d.IdentityOwner
	}
	return md
}

func (d *AWSDynamoDB) ContainerEnv() []corev1.EnvFromSource {
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

func (d *AWSDynamoDB) VolumeMounts() []corev1.VolumeMount {
	return nil
}

func (d *AWSDynamoDB) Volumes() []corev1.Volume {
	return nil
}
