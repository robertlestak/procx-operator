package http

import (
	"strconv"
	"strings"

	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

type HTTPRequest struct {
	Method                *string            `json:"method,omitempty"`
	URL                   string             `json:"url"`
	ContentType           *string            `json:"contentType,omitempty"`
	SuccessfulStatusCodes *[]int             `json:"successfulStatusCodes,omitempty"`
	Headers               *map[string]string `json:"headers,omitempty"`
}

type RetrieveRequest struct {
	HTTPRequest      `json:",inline"`
	KeyJSONSelector  *string `json:"keyJsonSelector,omitempty"`
	WorkJSONSelector *string `json:"workJsonSelector,omitempty"`
}

type HTTP struct {
	EnableTLS       *bool            `json:"enableTLS,omitempty"`
	TLSCA           *string          `json:"tlsCA,omitempty"`
	TLSCert         *string          `json:"tlsCert,omitempty"`
	TLSKey          *string          `json:"tlsKey,omitempty"`
	TLSSecretName   *string          `json:"tlsSecretName,omitempty"`
	TLSInsecure     *bool            `json:"tlsInsecure,omitempty"`
	RetrieveRequest *RetrieveRequest `json:"retrieveRequest,omitempty"`
	ClearRequest    *HTTPRequest     `json:"clearRequest,omitempty"`
	FailRequest     *HTTPRequest     `json:"failRequest,omitempty"`
	Key             *string          `json:"key,omitempty"`
}

func (d *HTTP) KedaSupport() bool {
	return false
}

func (d *HTTP) ConfigSecret() map[string]string {
	secData := map[string]string{}
	if d.EnableTLS != nil && *d.EnableTLS {
		secData["PROCX_HTTP_ENABLE_TLS"] = "true"
	}
	if d.TLSCA != nil {
		secData["PROCX_HTTP_TLS_CA_FILE"] = *d.TLSCA
	}
	if d.TLSCert != nil {
		secData["PROCX_HTTP_TLS_CERT_FILE"] = *d.TLSCert
	}
	if d.TLSKey != nil {
		secData["PROCX_HTTP_TLS_KEY_FILE"] = *d.TLSKey
	}
	if d.TLSInsecure != nil {
		secData["PROCX_HTTP_TLS_INSECURE"] = strconv.FormatBool(*d.TLSInsecure)
	}
	if d.RetrieveRequest != nil {
		if d.RetrieveRequest.Method != nil {
			secData["PROCX_HTTP_RETRIEVE_METHOD"] = *d.RetrieveRequest.Method
		}
		if d.RetrieveRequest.URL != "" {
			secData["PROCX_HTTP_RETRIEVE_URL"] = d.RetrieveRequest.URL
		}
		if d.RetrieveRequest.ContentType != nil {
			secData["PROCX_HTTP_RETRIEVE_CONTENT_TYPE"] = *d.RetrieveRequest.ContentType
		}
		if d.RetrieveRequest.SuccessfulStatusCodes != nil {
			var sc []string
			for _, s := range *d.RetrieveRequest.SuccessfulStatusCodes {
				sc = append(sc, strconv.Itoa(s))
			}
			secData["PROCX_HTTP_RETRIEVE_SUCCESSFUL_STATUS_CODES"] = strings.Join(sc, ",")
		}
		if d.RetrieveRequest.Headers != nil {
			var hd string
			for k, v := range *d.RetrieveRequest.Headers {
				hd += k + "=" + v + ","
			}
			// remove last comma
			hd = hd[:len(hd)-1]
			secData["PROCX_HTTP_RETRIEVE_HEADERS"] = hd
		}
		if d.RetrieveRequest.KeyJSONSelector != nil {
			secData["PROCX_HTTP_RETRIEVE_KEY_JSON_SELECTOR"] = *d.RetrieveRequest.KeyJSONSelector
		}
		if d.RetrieveRequest.WorkJSONSelector != nil {
			secData["PROCX_HTTP_RETRIEVE_WORK_JSON_SELECTOR"] = *d.RetrieveRequest.WorkJSONSelector
		}
	}
	if d.ClearRequest != nil {
		if d.ClearRequest.Method != nil {
			secData["PROCX_HTTP_CLEAR_METHOD"] = *d.ClearRequest.Method
		}
		if d.ClearRequest.URL != "" {
			secData["PROCX_HTTP_CLEAR_URL"] = d.ClearRequest.URL
		}
		if d.ClearRequest.ContentType != nil {
			secData["PROCX_HTTP_CLEAR_CONTENT_TYPE"] = *d.ClearRequest.ContentType
		}
		if d.ClearRequest.SuccessfulStatusCodes != nil {
			var sc []string
			for _, s := range *d.ClearRequest.SuccessfulStatusCodes {
				sc = append(sc, strconv.Itoa(s))
			}
			secData["PROCX_HTTP_CLEAR_SUCCESSFUL_STATUS_CODES"] = strings.Join(sc, ",")
		}
		if d.ClearRequest.Headers != nil {
			var hd string
			for k, v := range *d.ClearRequest.Headers {
				hd += k + "=" + v + ","
			}
			// remove last comma
			hd = hd[:len(hd)-1]
			secData["PROCX_HTTP_CLEAR_HEADERS"] = hd
		}
	}
	if d.FailRequest != nil {
		if d.FailRequest.Method != nil {
			secData["PROCX_HTTP_FAIL_METHOD"] = *d.FailRequest.Method
		}
		if d.FailRequest.URL != "" {
			secData["PROCX_HTTP_FAIL_URL"] = d.FailRequest.URL
		}
		if d.FailRequest.ContentType != nil {
			secData["PROCX_HTTP_FAIL_CONTENT_TYPE"] = *d.FailRequest.ContentType
		}
		if d.FailRequest.SuccessfulStatusCodes != nil {
			var sc []string
			for _, s := range *d.FailRequest.SuccessfulStatusCodes {
				sc = append(sc, strconv.Itoa(s))
			}
			secData["PROCX_HTTP_FAIL_SUCCESSFUL_STATUS_CODES"] = strings.Join(sc, ",")
		}
		if d.FailRequest.Headers != nil {
			var hd string
			for k, v := range *d.FailRequest.Headers {
				hd += k + "=" + v + ","
			}
			// remove last comma
			hd = hd[:len(hd)-1]
			secData["PROCX_HTTP_FAIL_HEADERS"] = hd
		}
	}
	return secData
}

func (d *HTTP) TriggerAuth(name string) *kedav1alpha1.TriggerAuthenticationSpec {
	s := &kedav1alpha1.TriggerAuthenticationSpec{}
	return s
}

func (d *HTTP) HasAuth() bool {
	return false
}

func (d *HTTP) KedaScalerName() string {
	return ""
}

func (d *HTTP) ContainerEnv() []corev1.EnvFromSource {
	envFrom := []corev1.EnvFromSource{}
	return envFrom
}

func (d *HTTP) Metadata() map[string]string {
	md := map[string]string{}
	return md
}

func (d *HTTP) VolumeMounts() []corev1.VolumeMount {
	if d.TLSSecretName == nil || *d.TLSSecretName == "" {
		return nil
	}
	v := corev1.VolumeMount{
		Name:      *d.TLSSecretName,
		MountPath: "/etc/procx/tls",
		ReadOnly:  true,
	}
	return []corev1.VolumeMount{v}
}

func (d *HTTP) Volumes() []corev1.Volume {
	if d.TLSSecretName == nil || *d.TLSSecretName == "" {
		return nil
	}
	v := corev1.Volume{
		Name: *d.TLSSecretName,
		VolumeSource: corev1.VolumeSource{
			Secret: &corev1.SecretVolumeSource{
				SecretName: *d.TLSSecretName,
			},
		},
	}
	return []corev1.Volume{v}
}
