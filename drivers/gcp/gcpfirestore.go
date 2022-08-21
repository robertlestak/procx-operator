package gcp

import (
	"encoding/json"
	"fmt"

	"cloud.google.com/go/firestore"
	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

type FirestoreOp string

var (
	FirestoreRMOp     = FirestoreOp("rm")
	FirestoreMVOp     = FirestoreOp("mv")
	FirestoreUpdateOp = FirestoreOp("update")
)

type GCPFirestoreQuery struct {
	Path    *string              `json:"path,omitempty"`
	Op      *string              `json:"op,omitempty"`
	Value   interface{}          `json:"value,omitempty"`
	OrderBy *string              `json:"orderBy,omitempty"`
	Order   *firestore.Direction `json:"order,omitempty"`
}

type GCPFirestore struct {
	RetrieveCollection      *string                 `json:"retrieveCollection,omitempty"`
	RetrieveDocument        *string                 `json:"retrieveDocument,omitempty"`
	RetrieveQuery           *GCPFirestoreQuery      `json:"retrieveQuery,omitempty"`
	RetrieveDocumentJSONKey *string                 `json:"retrieveDocumentJSONKey,omitempty"`
	ClearOp                 *FirestoreOp            `json:"clearOp,omitempty"`
	ClearUpdate             *map[string]interface{} `json:"clearUpdate,omitempty"`
	ClearCollection         *string                 `json:"clearCollection,omitempty"`
	FailOp                  *FirestoreOp            `json:"failOp,omitempty"`
	FailUpdate              *map[string]interface{} `json:"failUpdate,omitempty"`
	FailCollection          *string                 `json:"failCollection,omitempty"`
	ProjectID               string                  `json:"projectId"`
}

func (d *GCPFirestore) ConfigSecret() map[string]string {
	secData := map[string]string{}
	if d.ProjectID != "" {
		secData["PROCX_GCP_PROJECT_ID"] = d.ProjectID
	}
	if d.RetrieveCollection != nil {
		secData["PROCX_GCP_FIRESTORE_RETRIEVE_COLLECTION"] = *d.RetrieveCollection
	}
	if d.RetrieveDocument != nil {
		secData["PROCX_GCP_FIRESTORE_RETRIEVE_DOCUMENT"] = *d.RetrieveDocument
	}
	if d.RetrieveQuery != nil {
		secData["PROCX_GCP_FIRESTORE_RETRIEVE_QUERY_PATH"] = *d.RetrieveQuery.Path
		secData["PROCX_GCP_FIRESTORE_RETRIEVE_QUERY_OP"] = *d.RetrieveQuery.Op
		secData["PROCX_GCP_FIRESTORE_RETRIEVE_QUERY_VALUE"] = string(d.RetrieveQuery.Value.([]byte))
		secData["PROCX_GCP_FIRESTORE_RETRIEVE_QUERY_ORDER_BY"] = *d.RetrieveQuery.OrderBy
		secData["PROCX_GCP_FIRESTORE_RETRIEVE_QUERY_ORDER"] = fmt.Sprint(*d.RetrieveQuery.Order)
	}
	if d.RetrieveDocumentJSONKey != nil {
		secData["PROCX_GCP_FIRESTORE_RETRIEVE_DOCUMENT_JSON_KEY"] = *d.RetrieveDocumentJSONKey
	}
	if d.ClearOp != nil {
		secData["PROCX_GCP_FIRESTORE_CLEAR_OP"] = string(*d.ClearOp)
	}
	if d.ClearUpdate != nil {
		jd, err := json.Marshal(d.ClearUpdate)
		if err != nil {
			return nil
		}
		secData["PROCX_GCP_FIRESTORE_CLEAR_UPDATE"] = string(jd)
	}
	if d.ClearCollection != nil {
		secData["PROCX_GCP_FIRESTORE_CLEAR_COLLECTION"] = *d.ClearCollection
	}
	if d.FailOp != nil {
		secData["PROCX_GCP_FIRESTORE_FAIL_OP"] = string(*d.FailOp)
	}
	if d.FailUpdate != nil {
		jd, err := json.Marshal(d.FailUpdate)
		if err != nil {
			return nil
		}
		secData["PROCX_GCP_FIRESTORE_FAIL_UPDATE"] = string(jd)
	}
	if d.FailCollection != nil {
		secData["PROCX_GCP_FIRESTORE_FAIL_COLLECTION"] = *d.FailCollection
	}
	return secData
}

func (d *GCPFirestore) KedaSupport() bool {
	return false
}

func (d *GCPFirestore) HasAuth() bool {
	return true
}

func (d *GCPFirestore) KedaScalerName() string {
	return ""
}

func (d *GCPFirestore) TriggerAuth(name string) *kedav1alpha1.TriggerAuthenticationSpec {
	s := &kedav1alpha1.TriggerAuthenticationSpec{}
	return s
}

func (d *GCPFirestore) Metadata() map[string]string {
	md := map[string]string{}
	return md
}

func (d *GCPFirestore) ContainerEnv() []corev1.EnvFromSource {
	envFrom := []corev1.EnvFromSource{}
	return envFrom
}

func (d *GCPFirestore) VolumeMounts() []corev1.VolumeMount {
	return nil
}

func (d *GCPFirestore) Volumes() []corev1.Volume {
	return nil
}
