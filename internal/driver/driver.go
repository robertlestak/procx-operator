package driver

import (
	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

type Driver interface {
	ConfigSecret() map[string]string
	Metadata() map[string]string
	HasAuth() bool
	KedaSupport() bool
	KedaScalerName() string
	ContainerEnv() []corev1.EnvFromSource
	TriggerAuth(string) *kedav1alpha1.TriggerAuthenticationSpec
	VolumeMounts() []corev1.VolumeMount
	Volumes() []corev1.Volume
}
