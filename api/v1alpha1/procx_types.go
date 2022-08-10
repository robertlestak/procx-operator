/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	"github.com/robertlestak/procx/pkg/drivers"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type ScalerType string

var (
	ScalerTypeDeployment = ScalerType("Deployment")
	ScalerTypeJob        = ScalerType("Job")
)

// ProcXSpec defines the desired state of ProcX
type ProcXSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Enum=Deployment;Job
	ScalerType *ScalerType `json:"scalerType,omitempty"`
	// +kubebuilder:validation:Enum=aws-sqs;cassandra;gcp-pubsub;mongodb;mysql;postgres;rabbitmq;redis-list;redis-pubsub
	DriverName         drivers.DriverName             `json:"driver"`
	AWSDynamoDB        *AWSDynamoDB                   `json:"awsDynamoDB,omitempty"`
	AWSSQS             *AWSSQS                        `json:"awsSQS,omitempty"`
	AWSS3              *AWSS3                         `json:"awsS3,omitempty"`
	Cassandra          *Cassandra                     `json:"cassandra,omitempty"`
	Elasticsearch      *Elasticsearch                 `json:"elasticsearch,omitempty"`
	GCPPubSub          *GCPPubSub                     `json:"gcpPubSub,omitempty"`
	GCPGCS             *GCPGCS                        `json:"gcpGCS,omitempty"`
	Kafka              *Kafka                         `json:"kafka,omitempty"`
	MongoDB            *MongoDB                       `json:"mongodb,omitempty"`
	MySQL              *MySQL                         `json:"mysql,omitempty"`
	Postgres           *Postgres                      `json:"postgres,omitempty"`
	RabbitMQ           *RabbitMQ                      `json:"rabbitmq,omitempty"`
	RedisList          *RedisList                     `json:"redisList,omitempty"`
	Image              string                         `json:"image"`
	HostEnv            *bool                          `json:"hostEnv,omitempty"`
	Daemon             *bool                          `json:"daemon,omitempty"`
	BackoffLimit       *int32                         `json:"backoffLimit,omitempty"`
	MinReplicaCount    *int32                         `json:"minReplicas,omitempty"`
	MaxReplicaCount    *int32                         `json:"maxReplicas,omitempty"`
	CoolDownPeriod     *int32                         `json:"coolDownPeriod,omitempty"`
	PollingInterval    *int32                         `json:"pollingInterval,omitempty"`
	ServiceAccountName *string                        `json:"serviceAccountName,omitempty"`
	CommonLabels       *map[string]string             `json:"commonLabels,omitempty"`
	ImagePullSecrets   *[]corev1.LocalObjectReference `json:"imagePullSecrets,omitempty"`
	EnvSecretNames     *[]string                      `json:"envSecretNames,omitempty"`
	Resources          *corev1.ResourceRequirements   `json:"resources,omitempty"`
	VolumeMounts       *[]corev1.VolumeMount          `json:"volumeMounts,omitempty"`
	Volumes            *[]corev1.Volume               `json:"volumes,omitempty"`
	PodTemplate        *corev1.PodTemplateSpec        `json:"podTemplate,omitempty"`
}

type DeployStatus struct {
	Replicas            int32 `json:"replicas"`
	AvailableReplicas   int32 `json:"availableReplicas"`
	UnavailableReplicas int32 `json:"unavailableReplicas"`
	UpdatedReplicas     int32 `json:"updatedReplicas"`
	ReadyReplicas       int32 `json:"readyReplicas"`
}

type ScaledObjectStatus struct {
	LastActiveTime     *metav1.Time                          `json:"lastActiveTime,omitempty"`
	Health             *map[string]kedav1alpha1.HealthStatus `json:"health,omitempty"`
	PausedReplicaCount *int32                                `json:"pausedReplicaCount,omitempty"`
}

// ProcXStatus defines the observed state of ProcX
type ProcXStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Status             string              `json:"status"`
	Pods               []string            `json:"pods,omitempty"`
	DeployStatus       *DeployStatus       `json:"deployment,omitempty"`
	ScaledObjectStatus *ScaledObjectStatus `json:"scaledObject,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ProcX is the Schema for the ProcXs API
type ProcX struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProcXSpec   `json:"spec,omitempty"`
	Status ProcXStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ProcXList contains a list of ProcX
type ProcXList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProcX `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ProcX{}, &ProcXList{})
}
