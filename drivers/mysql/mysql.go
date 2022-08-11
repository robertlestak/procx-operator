package mysql

import (
	"strconv"
	"strings"

	"github.com/robertlestak/procx-operator/internal/utils"
	"github.com/robertlestak/procx/pkg/schema"

	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"

	corev1 "k8s.io/api/core/v1"
)

type MySQL struct {
	Host               string           `json:"host"`
	Port               int              `json:"port"`
	User               string           `json:"user"`
	Password           string           `json:"password"`
	PasswordSecretName *string          `json:"passwordSecretName,omitempty"`
	DBName             string           `json:"dbName"`
	QueryReturnsKey    *bool            `json:"queryReturnsKey"`
	RetrieveQuery      *schema.SqlQuery `json:"retrieveQuery"`
	FailureQuery       *schema.SqlQuery `json:"failureQuery"`
	ClearQuery         *schema.SqlQuery `json:"clearQuery"`
	Key                *string          `json:"key"`
	TargetQueryValue   *int             `json:"targetQueryValue"`
	ScaleQuery         *string          `json:"scaleQuery"`
}

func (d *MySQL) ConfigSecret() map[string]string {
	secData := map[string]string{}
	if d != nil {
		secData["PROCX_MYSQL_HOST"] = d.Host
		secData["PROCX_MYSQL_PORT"] = strconv.Itoa(int(d.Port))
		secData["PROCX_MYSQL_USER"] = d.User
		secData["PROCX_MYSQL_PASSWORD"] = d.Password
		secData["PROCX_MYSQL_DATABASE"] = d.DBName
		if d.QueryReturnsKey != nil {
			secData["PROCX_MYSQL_QUERY_KEY"] = strconv.FormatBool(*d.QueryReturnsKey)
		}
		secData["PROCX_MYSQL_RETRIEVE_QUERY"] = d.RetrieveQuery.Query
		if d.RetrieveQuery.Params != nil {
			secData["PROCX_MYSQL_RETRIEVE_PARAMS"] = strings.Join(utils.AnySliceToString(d.RetrieveQuery.Params), ",")
		}
		secData["PROCX_MYSQL_CLEAR_QUERY"] = d.ClearQuery.Query
		if d.ClearQuery.Params != nil {
			secData["PROCX_MYSQL_CLEAR_PARAMS"] = strings.Join(utils.AnySliceToString(d.ClearQuery.Params), ",")
		}
		secData["PROCX_MYSQL_FAIL_QUERY"] = d.FailureQuery.Query
		if d.FailureQuery.Params != nil {
			secData["PROCX_MYSQL_FAIL_PARAMS"] = strings.Join(utils.AnySliceToString(d.FailureQuery.Params), ",")
		}
	}
	return secData
}

func (d *MySQL) Metadata() map[string]string {
	md := map[string]string{
		"host":     d.Host,
		"port":     strconv.Itoa(d.Port),
		"username": d.User,
		"dbName":   d.DBName,
		"query":    *d.ScaleQuery,
	}
	if d.TargetQueryValue != nil {
		md["queryValue"] = strconv.Itoa(*d.TargetQueryValue)
	}
	return md
}

func (d *MySQL) HasAuth() bool {
	return d.PasswordSecretName != nil && d.Password != ""
}

func (d *MySQL) KedaSupport() bool {
	return true
}

func (d *MySQL) KedaScalerName() string {
	return "mysql"
}

func (d *MySQL) ContainerEnv() []corev1.EnvFromSource {
	var envFrom []corev1.EnvFromSource
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

func (d *MySQL) TriggerAuth(name string) *kedav1alpha1.TriggerAuthenticationSpec {
	s := &kedav1alpha1.TriggerAuthenticationSpec{}
	if d.PasswordSecretName != nil {
		s.SecretTargetRef = append(s.SecretTargetRef, kedav1alpha1.AuthSecretTargetRef{
			Name:      *d.PasswordSecretName,
			Key:       "PROCX_MYSQL_PASSWORD",
			Parameter: "password",
		})
	} else if d.Password != "" {
		s.SecretTargetRef = append(s.SecretTargetRef, kedav1alpha1.AuthSecretTargetRef{
			Name:      name,
			Key:       "PROCX_MYSQL_PASSWORD",
			Parameter: "password",
		})
	}
	return s
}

func (d *MySQL) VolumeMounts() []corev1.VolumeMount {
	return nil
}

func (d *MySQL) Volumes() []corev1.Volume {
	return nil
}
