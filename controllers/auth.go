package controllers

import (
	"context"
	"reflect"

	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	procxv1alpha1 "github.com/robertlestak/procx-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func (r *ProcXReconciler) handleConfigSecret(ctx context.Context, job *procxv1alpha1.ProcX) (*ctrl.Result, error) {
	l := log.Log.WithValues("procx", job.Name)
	l.Info("Reconciling ProcX")
	found := &corev1.Secret{}
	var err error
	err = r.Get(ctx, types.NamespacedName{Name: deployName(job.Name), Namespace: job.Namespace}, found)
	if err != nil {
		l.Info("Error getting secret", "error", err)
		if k8serrors.IsNotFound(err) {
			l.Info("Creating secret")
			// Define and create a new secret.
			dep := r.configSecretForProcX(job)
			if err = r.Create(ctx, &dep); err != nil {
				l.Info("Error creating secret", "error", err)
				return &ctrl.Result{}, err
			}
			l.Info("Created secret")
			return nil, nil
		} else {
			l.Info("Error getting secret", "error", err)
			return &ctrl.Result{}, err
		}
	}
	l.Info("secret handled")
	return nil, nil
}

func (r *ProcXReconciler) configSecretForProcX(m *procxv1alpha1.ProcX) corev1.Secret {
	secData := map[string]string{}
	secData["PROCX_DRIVER"] = string(m.Spec.DriverName)
	if m.Spec.HostEnv != nil && *m.Spec.HostEnv {
		secData["PROCX_HOSTENV"] = "true"
	}
	nd := Driver(m).ConfigSecret()
	// merge the config secret data with the existing data
	for k, v := range nd {
		secData[k] = v
	}
	if m.Spec.Daemon != nil && *m.Spec.Daemon {
		secData["PROCX_DAEMON"] = "true"
	}
	if m.Spec.PassWorkAsArg != nil && *m.Spec.PassWorkAsArg {
		secData["PROCX_PASS_WORK_AS_ARG"] = "true"
	}
	if m.Spec.PassWorkAsStdin != nil && *m.Spec.PassWorkAsStdin {
		secData["PROCX_PASS_WORK_AS_STDIN"] = "true"
	}
	if m.Spec.PayloadFile != nil && *m.Spec.PayloadFile != "" {
		secData["PROCX_PAYLOAD_FILE"] = *m.Spec.PayloadFile
	}
	if m.Spec.KeepPayloadFile != nil && *m.Spec.KeepPayloadFile {
		secData["PROCX_KEEP_PAYLOAD_FILE"] = "true"
	}
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployName(m.Name),
			Namespace: m.Namespace,
			Labels:    labelsForApp(m),
		},
		StringData: secData,
	}
	return *secret
}

func (r *ProcXReconciler) triggerAuthForProcX(m *procxv1alpha1.ProcX) (*kedav1alpha1.TriggerAuthentication, error) {
	lbls := labelsForApp(m)
	triggerAuth := &kedav1alpha1.TriggerAuthentication{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployName(m.Name),
			Namespace: m.Namespace,
			Labels:    lbls,
		},
		Spec: *Driver(m).TriggerAuth(deployName(m.Name)),
	}

	// NOTE: calling SetControllerReference, and setting owner references in
	// general, is important as it allows deleted objects to be garbage collected.
	controllerutil.SetControllerReference(m, triggerAuth, r.Scheme)
	return triggerAuth, nil
}

func (r *ProcXReconciler) updateTriggerAuth(ctx context.Context, job *procxv1alpha1.ProcX, found *kedav1alpha1.TriggerAuthentication) error {
	var hasChanges bool
	triggerAuth, err := r.triggerAuthForProcX(job)
	if err != nil {
		return err
	}
	// set the resourceversion to the one from the found object
	triggerAuth.ResourceVersion = found.ResourceVersion
	// set the owner reference to the one from the found object
	triggerAuth.OwnerReferences = found.OwnerReferences

	// compare triggerAuth to found to see if there are any changes, ignore the status
	if !reflect.DeepEqual(triggerAuth.Spec, found.Spec) {
		hasChanges = true
		found.Spec = triggerAuth.Spec
	}
	if hasChanges {
		// save the updated trigger auth
		if err := r.Update(ctx, found); err != nil {
			return err
		}
	}
	return nil
}

func (r *ProcXReconciler) handleTriggerAuth(ctx context.Context, job *procxv1alpha1.ProcX) (*ctrl.Result, error) {
	l := log.Log.WithValues("procx", job.Name)
	l.Info("Reconciling ProcX")
	found := &kedav1alpha1.TriggerAuthentication{}
	err := r.Get(ctx, types.NamespacedName{Name: deployName(job.Name), Namespace: job.Namespace}, found)
	if err != nil {
		l.Info("Error getting TriggerAuthentication", "error", err)
		if k8serrors.IsNotFound(err) {
			l.Info("Creating TriggerAuthentication")
			// Define and create a new deployment.
			triggerAuth, err := r.triggerAuthForProcX(job)
			if err != nil {
				return &ctrl.Result{}, err
			}
			if err = r.Create(ctx, triggerAuth); err != nil {
				l.Info("Error creating TriggerAuthentication", "error", err)
				return &ctrl.Result{}, err
			}
			l.Info("Created TriggerAuthentication")
			return &ctrl.Result{Requeue: true}, nil
		} else {
			l.Info("Error getting TriggerAuthentication", "error", err)
			return &ctrl.Result{}, err
		}
	} else {
		l.Info("Found TriggerAuthentication")
		// update the trigger authentication if there are any changes
		if err := r.updateTriggerAuth(ctx, job, found); err != nil {
			l.Info("Error updating TriggerAuthentication", "error", err)
			return &ctrl.Result{}, err
		}
		l.Info("Updated TriggerAuthentication")
	}
	l.Info("TriggerAuthentication handled")
	return nil, nil
}
