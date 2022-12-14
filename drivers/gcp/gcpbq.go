package gcp

import (
	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

type GCPBQ struct {
	ProjectID     string  `json:"projectId"`
	RetrieveField *string `json:"retrieveField,omitempty"`
	RetrieveQuery *string `json:"retrieveQuery,omitempty"`
	ClearQuery    *string `json:"clearQuery,omitempty"`
	FailQuery     *string `json:"failQuery,omitempty"`
}

func (d *GCPBQ) ConfigSecret() map[string]string {
	secData := map[string]string{}
	if d.ProjectID != "" {
		secData["PROCX_GCP_PROJECT_ID"] = d.ProjectID
	}
	if d.RetrieveField != nil {
		secData["PROCX_GCP_BQ_RETRIEVE_FIELD"] = *d.RetrieveField
	}
	if d.RetrieveQuery != nil {
		secData["PROCX_GCP_BQ_RETRIEVE_QUERY"] = *d.RetrieveQuery
	}
	if d.ClearQuery != nil {
		secData["PROCX_GCP_BQ_CLEAR_QUERY"] = *d.ClearQuery
	}
	if d.FailQuery != nil {
		secData["PROCX_GCP_BQ_FAIL_QUERY"] = *d.FailQuery
	}
	return secData
}

func (d *GCPBQ) KedaSupport() bool {
	return false
}

func (d *GCPBQ) HasAuth() bool {
	return true
}

func (d *GCPBQ) KedaScalerName() string {
	return ""
}

func (d *GCPBQ) TriggerAuth(name string) *kedav1alpha1.TriggerAuthenticationSpec {
	s := &kedav1alpha1.TriggerAuthenticationSpec{}
	return s
}

func (d *GCPBQ) Metadata() map[string]string {
	md := map[string]string{}
	return md
}

func (d *GCPBQ) ContainerEnv() []corev1.EnvFromSource {
	envFrom := []corev1.EnvFromSource{}
	return envFrom
}

func (d *GCPBQ) VolumeMounts() []corev1.VolumeMount {
	return nil
}

func (d *GCPBQ) Volumes() []corev1.Volume {
	return nil
}
