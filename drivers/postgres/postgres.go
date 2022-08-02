package postgres

import (
	"strconv"
	"strings"

	"github.com/robertlestak/procx-operator/internal/utils"
	"github.com/robertlestak/procx/pkg/schema"

	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"

	corev1 "k8s.io/api/core/v1"
)

type Postgres struct {
	Host               string           `json:"host"`
	Port               int              `json:"port"`
	User               string           `json:"user"`
	Password           string           `json:"password"`
	PasswordSecretName *string          `json:"passwordSecretName,omitempty"`
	DBName             string           `json:"dbName"`
	SSLMode            string           `json:"sslMode"`
	QueryReturnsKey    *bool            `json:"queryReturnsKey"`
	RetrieveQuery      *schema.SqlQuery `json:"retrieveQuery"`
	FailureQuery       *schema.SqlQuery `json:"failureQuery"`
	ClearQuery         *schema.SqlQuery `json:"clearQuery"`
	Key                *string          `json:"key"`
	TargetQueryValue   *int             `json:"targetQueryValue"`
	ScaleQuery         *string          `json:"scaleQuery"`
}

func (d *Postgres) ConfigSecret() map[string]string {
	secData := map[string]string{}
	if d != nil {
		secData["PROCX_PSQL_HOST"] = d.Host
		secData["PROCX_PSQL_PORT"] = strconv.Itoa(int(d.Port))
		secData["PROCX_PSQL_USER"] = d.User
		secData["PROCX_PSQL_PASSWORD"] = d.Password
		secData["PROCX_PSQL_DATABASE"] = d.DBName
		secData["PROCX_PSQL_SSL_MODE"] = d.SSLMode
		if d.QueryReturnsKey != nil {
			secData["PROCX_PSQL_QUERY_KEY"] = strconv.FormatBool(*d.QueryReturnsKey)
		}
		secData["PROCX_PSQL_RETRIEVE_QUERY"] = d.RetrieveQuery.Query
		if d.RetrieveQuery.Params != nil {
			secData["PROCX_PSQL_RETRIEVE_PARAMS"] = strings.Join(utils.AnySliceToString(d.RetrieveQuery.Params), ",")
		}
		secData["PROCX_PSQL_CLEAR_QUERY"] = d.ClearQuery.Query
		if d.ClearQuery.Params != nil {
			secData["PROCX_PSQL_CLEAR_PARAMS"] = strings.Join(utils.AnySliceToString(d.ClearQuery.Params), ",")
		}
		secData["PROCX_PSQL_FAIL_QUERY"] = d.FailureQuery.Query
		if d.FailureQuery.Params != nil {
			secData["PROCX_PSQL_FAIL_PARAMS"] = strings.Join(utils.AnySliceToString(d.FailureQuery.Params), ",")
		}
	}
	return secData
}

func (d *Postgres) Metadata() map[string]string {
	md := map[string]string{
		"host":     d.Host,
		"port":     strconv.Itoa(d.Port),
		"userName": d.User,
		"dbName":   d.DBName,
		"sslmode":  d.SSLMode,
		"query":    *d.ScaleQuery,
	}
	if d.TargetQueryValue != nil {
		md["targetQueryValue"] = strconv.Itoa(*d.TargetQueryValue)
	}
	return md
}

func (d *Postgres) HasAuth() bool {
	return d.Password != "" && d.PasswordSecretName != nil
}

func (d *Postgres) KedaSupport() bool {
	return true
}

func (d *Postgres) KedaScalerName() string {
	return "postgresql"
}

func (d *Postgres) ContainerEnv() []corev1.EnvFromSource {
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

func (d *Postgres) TriggerAuth(name string) *kedav1alpha1.TriggerAuthenticationSpec {
	s := &kedav1alpha1.TriggerAuthenticationSpec{}
	if d.PasswordSecretName != nil {
		s.SecretTargetRef = append(s.SecretTargetRef, kedav1alpha1.AuthSecretTargetRef{
			Name:      *d.PasswordSecretName,
			Key:       "PROCX_PSQL_PASSWORD",
			Parameter: "password",
		})
	} else if d.Password != "" {
		s.SecretTargetRef = append(s.SecretTargetRef, kedav1alpha1.AuthSecretTargetRef{
			Name:      name,
			Key:       "PROCX_PSQL_PASSWORD",
			Parameter: "password",
		})
	}
	return s
}
