package controllers

import (
	"context"
	"reflect"

	procxv1alpha1 "github.com/robertlestak/procx-operator/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func (r *ProcXReconciler) envForContainer(m *procxv1alpha1.ProcX) []corev1.EnvFromSource {
	envFrom := []corev1.EnvFromSource{}
	cs := r.configSecretForProcX(m)
	envFrom = append(envFrom, corev1.EnvFromSource{
		SecretRef: &corev1.SecretEnvSource{
			LocalObjectReference: corev1.LocalObjectReference{
				Name: cs.Name,
			},
		},
	})
	nenv := Driver(m).ContainerEnv()
	envFrom = append(envFrom, nenv...)
	if m.Spec.EnvSecretNames != nil {
		falseVal := false
		for _, name := range *m.Spec.EnvSecretNames {
			envFrom = append(envFrom, corev1.EnvFromSource{
				SecretRef: &corev1.SecretEnvSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: name,
					},
					Optional: &falseVal,
				},
			})
		}
	}
	return envFrom
}

func (r *ProcXReconciler) podTemplateForProcX(m *procxv1alpha1.ProcX) corev1.PodTemplateSpec {
	lbls := labelsForApp(m)
	cont := corev1.Container{
		Image: m.Spec.Image,
		Name:  "procx",
	}

	if m.Spec.VolumeMounts != nil {
		cont.VolumeMounts = *m.Spec.VolumeMounts
	}

	cont.EnvFrom = r.envForContainer(m)

	podTemplate := corev1.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			Labels: lbls,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{cont},
		},
	}

	if m.Spec.PodTemplate != nil {
		podTemplate = *m.Spec.PodTemplate
		podTemplate.ObjectMeta.Labels = lbls
		for k, v := range m.Spec.PodTemplate.ObjectMeta.Labels {
			podTemplate.ObjectMeta.Labels[k] = v
		}
		for _, c := range podTemplate.Spec.Containers {
			if c.Name == "procx" {
				c.EnvFrom = append(c.EnvFrom, r.envForContainer(m)...)
				c.Image = m.Spec.Image
				if m.Spec.VolumeMounts != nil {
					c.VolumeMounts = *m.Spec.VolumeMounts
				}
			}
		}
	}

	if m.Spec.Resources != nil {
		for _, c := range podTemplate.Spec.Containers {
			if c.Name == "procx" {
				c.Resources = *m.Spec.Resources
			}
		}
	}

	if m.Spec.Volumes != nil {
		podTemplate.Spec.Volumes = *m.Spec.Volumes
	}

	if m.Spec.ServiceAccountName != nil {
		podTemplate.Spec.ServiceAccountName = *m.Spec.ServiceAccountName
	}

	if m.Spec.ImagePullSecrets != nil {
		podTemplate.Spec.ImagePullSecrets = append(podTemplate.Spec.ImagePullSecrets, *m.Spec.ImagePullSecrets...)
	}
	return podTemplate
}

func (r *ProcXReconciler) deploymentForProcX(m *procxv1alpha1.ProcX) *appsv1.Deployment {
	lbls := labelsForApp(m)
	replicas := int32(0)

	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployName(m.Name),
			Namespace: m.Namespace,
			Labels:    lbls,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: lbls,
			},
		},
	}

	dep.Spec.Template = r.podTemplateForProcX(m)

	// NOTE: calling SetControllerReference, and setting owner references in
	// general, is important as it allows deleted objects to be garbage collected.
	controllerutil.SetControllerReference(m, dep, r.Scheme)
	return dep
}

func (r *ProcXReconciler) updateDeployment(ctx context.Context, m *procxv1alpha1.ProcX, found *appsv1.Deployment) error {
	// if image has changed, update the image in the scaled job
	var hasChanges bool
	dep := r.deploymentForProcX(m)
	// compare dep to found to see if there are any changes, ignore the status
	// compare dep to found to see if there are any changes, ignore the status
	if !reflect.DeepEqual(dep.Spec, found.Spec) {
		hasChanges = true
		found.Spec = dep.Spec
	}
	if hasChanges {
		// save the updated scaled job
		if err := r.Update(ctx, found); err != nil {
			return err
		}
	}
	return nil
}

func (r *ProcXReconciler) handleDeployment(ctx context.Context, job *procxv1alpha1.ProcX) (*ctrl.Result, error) {
	l := log.Log.WithValues("procx", job.Name)
	l.Info("Reconciling ProcX")
	found := &appsv1.Deployment{}
	err := r.Get(ctx, types.NamespacedName{Name: deployName(job.Name), Namespace: job.Namespace}, found)
	if err != nil {
		l.Info("Error getting Deployment", "error", err)
		if k8serrors.IsNotFound(err) {
			l.Info("Creating Deployment")
			// Define and create a new deployment.
			dep := r.deploymentForProcX(job)
			if err = r.Create(ctx, dep); err != nil {
				l.Info("Error creating Deployment", "error", err)
				return &ctrl.Result{}, err
			}
			job.Status.Status = "Pending"
			if err := r.Status().Update(ctx, job); err != nil {
				l.Info("Error updating ProcX status", "error", err)
				return &ctrl.Result{}, err
			}
			l.Info("Created Deployment")
			return &ctrl.Result{Requeue: true}, nil
		} else {
			l.Info("Error getting Deployment", "error", err)
			return &ctrl.Result{}, err
		}
	} else {
		l.Info("Found Deployment")
		if err := r.updateDeployment(ctx, job, found); err != nil {
			l.Info("Error updating Deployment", "error", err)
			return &ctrl.Result{}, err
		}
	}
	if job.Status.DeployStatus == nil {
		job.Status.DeployStatus = &procxv1alpha1.DeployStatus{}
	}
	job.Status.DeployStatus.Replicas = found.Status.Replicas
	job.Status.DeployStatus.AvailableReplicas = found.Status.AvailableReplicas
	job.Status.DeployStatus.UnavailableReplicas = found.Status.UnavailableReplicas
	job.Status.DeployStatus.UpdatedReplicas = found.Status.UpdatedReplicas
	job.Status.DeployStatus.ReadyReplicas = found.Status.ReadyReplicas
	l.Info("deployment handled")
	return nil, nil
}
