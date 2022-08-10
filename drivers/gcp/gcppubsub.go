package gcp

import (
	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

type GCPPubSub struct {
	ProjectID             string  `json:"projectId"`
	SubscriptionName      string  `json:"subscriptionName"`
	CredentialsSecretName *string `json:"credentialsSecretName"`
	Mode                  *string `json:"mode,omitempty"`
	Value                 *string `json:"value,omitempty"`
	PodIdentityProvider   *string `json:"podIdentityProvider,omitempty"`
}

func (d *GCPPubSub) ConfigSecret() map[string]string {
	secData := map[string]string{}
	if d != nil {
		secData["PROCX_GCP_PROJECT_ID"] = d.ProjectID
		secData["PROCX_GCP_PUBSUB_SUBSCRIPTION"] = d.SubscriptionName
	}
	return secData
}

func (d *GCPPubSub) Metadata() map[string]string {
	md := map[string]string{
		"subscriptionName": d.SubscriptionName,
	}
	if d.Mode == nil {
		md["mode"] = "SubscriptionSize"
	} else {
		md["mode"] = *d.Mode
	}
	if d.Value == nil {
		md["value"] = "1"
	} else {
		md["value"] = *d.Value
	}
	return md
}

func (d *GCPPubSub) HasAuth() bool {
	return true
}

func (d *GCPPubSub) KedaSupport() bool {
	return true
}

func (d *GCPPubSub) KedaScalerName() string {
	return "gcp-pubsub"
}

func (d *GCPPubSub) ContainerEnv() []corev1.EnvFromSource {
	return []corev1.EnvFromSource{}
}

func (d *GCPPubSub) TriggerAuth(name string) *kedav1alpha1.TriggerAuthenticationSpec {
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
