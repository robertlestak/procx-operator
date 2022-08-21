package cassandra

import (
	"strings"

	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	"github.com/robertlestak/procx-operator/internal/utils"
	"github.com/robertlestak/procx/pkg/schema"
	corev1 "k8s.io/api/core/v1"
)

type Cassandra struct {
	Hosts              []string         `json:"hosts"`
	User               *string          `json:"user"`
	Password           *string          `json:"password"`
	PasswordSecretName *string          `json:"passwordSecretName,omitempty"`
	Keyspace           string           `json:"keyspace"`
	Consistency        string           `json:"consistency"`
	RetrieveField      *string          `json:"retrieveField,omitempty"`
	RetrieveQuery      *schema.SqlQuery `json:"retrieveQuery"`
	FailureQuery       *schema.SqlQuery `json:"failureQuery"`
	ClearQuery         *schema.SqlQuery `json:"clearQuery"`
	ScaleQuery         *string          `json:"scaleQuery"`
	TargetQueryValue   *string          `json:"targetQueryValue"`
}

func (d *Cassandra) KedaSupport() bool {
	return true
}

func (d *Cassandra) ConfigSecret() map[string]string {
	secData := map[string]string{}
	secData["PROCX_CASSANDRA_HOSTS"] = strings.Join(d.Hosts, ",")
	if d.User != nil {
		secData["PROCX_CASSANDRA_USER"] = *d.User
	}
	if d.Password != nil {
		secData["PROCX_CASSANDRA_PASSWORD"] = *d.Password
	}
	secData["PROCX_CASSANDRA_KEYSPACE"] = d.Keyspace
	if d.RetrieveField != nil {
		secData["PROCX_CASSANDRA_RETRIEVE_FIELD"] = *d.RetrieveField
	}
	secData["PROCX_CASSANDRA_RETRIEVE_QUERY"] = d.RetrieveQuery.Query
	if d.RetrieveQuery.Params != nil {
		secData["PROCX_CASSANDRA_RETRIEVE_PARAMS"] = strings.Join(utils.AnySliceToString(d.RetrieveQuery.Params), ",")
	}
	secData["PROCX_CASSANDRA_CLEAR_QUERY"] = d.ClearQuery.Query
	if d.ClearQuery.Params != nil {
		secData["PROCX_CASSANDRA_CLEAR_PARAMS"] = strings.Join(utils.AnySliceToString(d.ClearQuery.Params), ",")
	}
	secData["PROCX_CASSANDRA_FAIL_QUERY"] = d.FailureQuery.Query
	if d.FailureQuery.Params != nil {
		secData["PROCX_CASSANDRA_FAIL_PARAMS"] = strings.Join(utils.AnySliceToString(d.FailureQuery.Params), ",")
	}
	return secData
}

func (d *Cassandra) TriggerAuth(name string) *kedav1alpha1.TriggerAuthenticationSpec {
	s := &kedav1alpha1.TriggerAuthenticationSpec{}
	if d.PasswordSecretName != nil {
		s.SecretTargetRef = append(s.SecretTargetRef, kedav1alpha1.AuthSecretTargetRef{
			Name:      *d.PasswordSecretName,
			Key:       "PROCX_CASSANDRA_PASSWORD",
			Parameter: "password",
		})
	} else if *d.Password != "" {
		s.SecretTargetRef = append(s.SecretTargetRef, kedav1alpha1.AuthSecretTargetRef{
			Name:      name,
			Key:       "PROCX_CASSANDRA_PASSWORD",
			Parameter: "password",
		})
	}
	return s
}

func (d *Cassandra) HasAuth() bool {
	return d.PasswordSecretName != nil && *d.Password != ""
}

func (d *Cassandra) KedaScalerName() string {
	return "cassandra"
}

func (d *Cassandra) ContainerEnv() []corev1.EnvFromSource {
	falseVal := false
	envFrom := []corev1.EnvFromSource{}
	if d.PasswordSecretName != nil {
		envFrom = append(envFrom, corev1.EnvFromSource{
			SecretRef: &corev1.SecretEnvSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: *d.PasswordSecretName,
				},
				Optional: &falseVal,
			},
		})
	}
	return envFrom
}

func (d *Cassandra) Metadata() map[string]string {
	md := map[string]string{
		"clusterIPAddress": d.Hosts[0],
		"consistency":      d.Consistency,
		"keyspace":         d.Keyspace,
		"query":            *d.ScaleQuery,
	}
	if d.User != nil {
		md["username"] = *d.User
	}
	if d.TargetQueryValue != nil {
		md["targetQueryValue"] = *d.TargetQueryValue
	}
	return md
}

func (d *Cassandra) VolumeMounts() []corev1.VolumeMount {
	return nil
}

func (d *Cassandra) Volumes() []corev1.Volume {
	return nil
}
