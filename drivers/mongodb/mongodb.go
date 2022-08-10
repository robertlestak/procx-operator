package mongodb

import (
	"strconv"

	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

type MongoDB struct {
	Host               string  `json:"host"`
	Port               int     `json:"port"`
	User               string  `json:"user"`
	Password           *string `json:"password"`
	PasswordSecretName *string `json:"passwordSecretName,omitempty"`
	DBName             string  `json:"dbName"`
	Collection         string  `json:"collection"`
	ScaleQuery         *string `json:"scaleQuery,omitempty"`
	QueryValue         *string `json:"queryValue,omitempty"`
	RetrieveQuery      *string `json:"retrieveQuery"`
	FailureQuery       *string `json:"failureQuery"`
	ClearQuery         *string `json:"clearQuery"`
	Key                *string `json:"key"`
}

func (d *MongoDB) ConfigSecret() map[string]string {
	secData := map[string]string{}
	secData["PROCX_MONGO_HOST"] = d.Host
	secData["PROCX_MONGO_PORT"] = strconv.Itoa(int(d.Port))
	secData["PROCX_MONGO_USER"] = d.User
	if d.Password != nil {
		secData["PROCX_MONGO_PASSWORD"] = *d.Password
	}
	secData["PROCX_MONGO_DATABASE"] = d.DBName
	secData["PROCX_MONGO_COLLECTION"] = d.Collection

	if d.RetrieveQuery != nil {
		secData["PROCX_MONGO_RETRIEVE_QUERY"] = *d.RetrieveQuery
	}
	if d.ClearQuery != nil {
		secData["PROCX_MONGO_CLEAR_QUERY"] = *d.ClearQuery
	}
	if d.FailureQuery != nil {
		secData["PROCX_MONGO_FAIL_QUERY"] = *d.FailureQuery
	}
	return secData
}

func (d *MongoDB) Metadata() map[string]string {
	md := map[string]string{
		"host":       d.Host,
		"port":       strconv.Itoa(d.Port),
		"username":   d.User,
		"dbName":     d.DBName,
		"collection": d.Collection,
		"query":      *d.ScaleQuery,
	}
	if d.QueryValue != nil {
		md["queryValue"] = *d.QueryValue
	}
	return md
}

func (d *MongoDB) HasAuth() bool {
	if d.User != "" {
		return true
	}
	if d.PasswordSecretName != nil {
		return true
	}
	if d.Password != nil {
		return true
	}
	return false
}

func (d *MongoDB) KedaSupport() bool {
	return true
}

func (d *MongoDB) KedaScalerName() string {
	return "mongodb"
}

func (d *MongoDB) ContainerEnv() []corev1.EnvFromSource {
	envFrom := []corev1.EnvFromSource{}
	falseVal := false
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

func (d *MongoDB) TriggerAuth(name string) *kedav1alpha1.TriggerAuthenticationSpec {
	s := &kedav1alpha1.TriggerAuthenticationSpec{}
	if d.PasswordSecretName != nil {
		s.SecretTargetRef = append(s.SecretTargetRef, kedav1alpha1.AuthSecretTargetRef{
			Name:      *d.PasswordSecretName,
			Key:       "PROCX_MONGO_PASSWORD",
			Parameter: "password",
		})
	} else if d.Password != nil && *d.Password != "" {
		s.SecretTargetRef = append(s.SecretTargetRef, kedav1alpha1.AuthSecretTargetRef{
			Name:      name,
			Key:       "PROCX_MONGO_PASSWORD",
			Parameter: "password",
		})
	}
	return s
}
