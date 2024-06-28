package controllers

import (
	"context"
	"reflect"

	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	procxv1alpha1 "github.com/robertlestak/procx-operator/api/v1alpha1"
	batchv1 "k8s.io/api/batch/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func (r *ProcXReconciler) updateJob(ctx context.Context, m *procxv1alpha1.ProcX, found *kedav1alpha1.ScaledJob) error {
	// if image has changed, update the image in the scaled job
	var hasChanges bool
	dep, err := r.scaledJobForProcX(m)
	if err != nil {
		return err
	}
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

func (r *ProcXReconciler) handleJob(ctx context.Context, job *procxv1alpha1.ProcX) (*ctrl.Result, error) {
	l := log.Log.WithValues("procx", job.Name)
	l.Info("Reconciling ProcX")
	found := &kedav1alpha1.ScaledJob{}
	err := r.Get(ctx, types.NamespacedName{Name: deployName(job.Name), Namespace: job.Namespace}, found)
	if err != nil {
		l.Info("Error getting Job", "error", err)
		if k8serrors.IsNotFound(err) {
			l.Info("Creating Job")
			// Define and create a new deployment.
			dep, err := r.scaledJobForProcX(job)
			if err != nil {
				l.Info("Error creating Job", "error", err)
				return &ctrl.Result{}, err
			}
			if err = r.Create(ctx, dep); err != nil {
				l.Info("Error creating Job", "error", err)
				return &ctrl.Result{}, err
			}
			job.Status.Status = "Pending"
			if err := r.Status().Update(ctx, job); err != nil {
				l.Info("Error updating ProcX status", "error", err)
				return &ctrl.Result{}, err
			}
			l.Info("Created Job")
			return &ctrl.Result{Requeue: true}, nil
		} else {
			l.Info("Error getting Job", "error", err)
			return &ctrl.Result{}, err
		}
	} else {
		l.Info("Found Job")
		if err := r.updateJob(ctx, job, found); err != nil {
			l.Info("Error updating Job", "error", err)
			return &ctrl.Result{}, err
		}
	}
	l.Info("job handled")
	return nil, nil
}

func (r *ProcXReconciler) triggerForProcX(m *procxv1alpha1.ProcX) (*kedav1alpha1.ScaleTriggers, error) {
	trigger := &kedav1alpha1.ScaleTriggers{
		Type:     Driver(m).KedaScalerName(),
		Metadata: Driver(m).Metadata(),
	}
	if Driver(m).HasAuth() {
		trigger.AuthenticationRef = &kedav1alpha1.ScaledObjectAuthRef{
			Name: deployName(m.Name),
		}
	}
	return trigger, nil
}

func (r *ProcXReconciler) scaledJobForProcX(m *procxv1alpha1.ProcX) (*kedav1alpha1.ScaledJob, error) {
	lbls := labelsForApp(m)

	scaledJob := &kedav1alpha1.ScaledJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployName(m.Name),
			Namespace: m.Namespace,
			Labels:    lbls,
		},
	}

	scaledJob.Spec.JobTargetRef = &batchv1.JobSpec{
		Template: r.podTemplateForProcX(m),
	}
	if m.Spec.BackoffLimit != nil {
		scaledJob.Spec.JobTargetRef.BackoffLimit = m.Spec.BackoffLimit
	}
	if m.Spec.FailedJobsHistoryLimit != nil {
		scaledJob.Spec.FailedJobsHistoryLimit = m.Spec.FailedJobsHistoryLimit
	} else {
		i := int32(10)
		scaledJob.Spec.FailedJobsHistoryLimit = &i
	}
	if m.Spec.SuccessfulJobsHistoryLimit != nil {
		scaledJob.Spec.SuccessfulJobsHistoryLimit = m.Spec.SuccessfulJobsHistoryLimit
	} else {
		i := int32(3)
		scaledJob.Spec.SuccessfulJobsHistoryLimit = &i
	}
	trigger, err := r.triggerForProcX(m)
	if err != nil {
		return nil, err
	}
	pi := m.Spec.PollingInterval
	if pi == nil {
		i := int32(30)
		pi = &i
	}
	scaledJob.Spec.Triggers = append(scaledJob.Spec.Triggers, *trigger)
	scaledJob.Spec.PollingInterval = pi
	if m.Spec.MaxReplicaCount != nil {
		scaledJob.Spec.MaxReplicaCount = m.Spec.MaxReplicaCount
	}

	// NOTE: calling SetControllerReference, and setting owner references in
	// general, is important as it allows deleted objects to be garbage collected.
	controllerutil.SetControllerReference(m, scaledJob, r.Scheme)
	return scaledJob, nil
}

func (r *ProcXReconciler) scaledObjectForProcX(m *procxv1alpha1.ProcX) (*kedav1alpha1.ScaledObject, error) {
	lbls := labelsForApp(m)
	trigger, err := r.triggerForProcX(m)
	if err != nil {
		return nil, err
	}
	coolDown := m.Spec.CoolDownPeriod
	if coolDown == nil {
		i := int32(300)
		coolDown = &i
	}
	pi := m.Spec.PollingInterval
	if pi == nil {
		i := int32(30)
		pi = &i
	}
	scaledObject := &kedav1alpha1.ScaledObject{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployName(m.Name),
			Namespace: m.Namespace,
			Labels:    lbls,
		},
		Spec: kedav1alpha1.ScaledObjectSpec{
			ScaleTargetRef: &kedav1alpha1.ScaleTarget{
				Kind: "Deployment",
				Name: deployName(m.Name),
			},
			Triggers:        []kedav1alpha1.ScaleTriggers{*trigger},
			CooldownPeriod:  coolDown,
			PollingInterval: pi,
		},
	}
	if m.Spec.MinReplicaCount != nil {
		scaledObject.Spec.MinReplicaCount = m.Spec.MinReplicaCount
	}
	if m.Spec.MaxReplicaCount != nil {
		scaledObject.Spec.MaxReplicaCount = m.Spec.MaxReplicaCount
	}

	// NOTE: calling SetControllerReference, and setting owner references in
	// general, is important as it allows deleted objects to be garbage collected.
	controllerutil.SetControllerReference(m, scaledObject, r.Scheme)
	return scaledObject, nil
}

func (r *ProcXReconciler) updateScaledObject(ctx context.Context, job *procxv1alpha1.ProcX, found *kedav1alpha1.ScaledObject) error {
	var hasChanges bool
	dep, err := r.scaledObjectForProcX(job)
	if err != nil {
		return err
	}
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

func (r *ProcXReconciler) handleScaledObject(ctx context.Context, job *procxv1alpha1.ProcX) (*ctrl.Result, error) {
	l := log.Log.WithValues("procx", job.Name)
	l.Info("Reconciling ProcX")
	found := &kedav1alpha1.ScaledObject{}
	err := r.Get(ctx, types.NamespacedName{Name: deployName(job.Name), Namespace: job.Namespace}, found)
	if err != nil {
		l.Info("Error getting ScaledObject", "error", err)
		if k8serrors.IsNotFound(err) {
			l.Info("Creating ScaledObject")
			// Define and create a new deployment.
			scaledObject, err := r.scaledObjectForProcX(job)
			if err != nil {
				return &ctrl.Result{}, err
			}
			if err = r.Create(ctx, scaledObject); err != nil {
				l.Info("Error creating ScaledObject", "error", err)
				return &ctrl.Result{}, err
			}
			l.Info("Created ScaledObject")
			return &ctrl.Result{Requeue: true}, nil
		} else {
			l.Info("Error getting ScaledObject", "error", err)
			return &ctrl.Result{}, err
		}
	} else {
		l.Info("Found ScaledObject")
		// update the scaled object if there are any changes
		if err := r.updateScaledObject(ctx, job, found); err != nil {
			l.Info("Error updating ScaledObject", "error", err)
			return &ctrl.Result{}, err
		}
		l.Info("Updated ScaledObject")
	}
	if job.Status.ScaledObjectStatus == nil {
		job.Status.ScaledObjectStatus = &procxv1alpha1.ScaledObjectStatus{}
	}
	job.Status.ScaledObjectStatus.LastActiveTime = found.Status.LastActiveTime
	job.Status.ScaledObjectStatus.Health = &found.Status.Health
	job.Status.ScaledObjectStatus.PausedReplicaCount = found.Status.PausedReplicaCount
	l.Info("scaledObject handled")
	return nil, nil
}
