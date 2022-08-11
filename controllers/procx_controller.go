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

package controllers

import (
	"context"
	"errors"

	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	procxv1alpha1 "github.com/robertlestak/procx-operator/api/v1alpha1"
	"github.com/robertlestak/procx-operator/internal/driver"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// ProcXReconciler reconciles a ProcX object
type ProcXReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=procx.k8s.lestak.sh,resources=procxes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=procx.k8s.lestak.sh,resources=procxes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=procx.k8s.lestak.sh,resources=procxes/finalizers,verbs=update

//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=batch,resources=jobs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=keda.sh,resources=scaledobjects,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=keda.sh,resources=scaledjobs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=keda.sh,resources=triggerauthentications,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ProcX object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.1/pkg/reconcile
func (r *ProcXReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	job := &procxv1alpha1.ProcX{}
	if err := r.Get(ctx, req.NamespacedName, job); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	l := log.Log.WithValues("procx", job.Name)
	if job.Spec.ScalerType == nil {
		dt := procxv1alpha1.ScalerTypeDeployment
		job.Spec.ScalerType = &dt
	} else {
		dt := *job.Spec.ScalerType
		if dt != procxv1alpha1.ScalerTypeDeployment && dt != procxv1alpha1.ScalerTypeJob {
			l.Info("ScalerType must be either Deployment or Job")
			return ctrl.Result{}, nil
		}
		job.Spec.ScalerType = &dt
		// if this is a job, do not run daemon - we want to run single job to completion
		if job.Spec.ScalerType != nil && *job.Spec.ScalerType == procxv1alpha1.ScalerTypeJob {
			fval := false
			job.Spec.Daemon = &fval
		}
	}
	l.Info("Reconciling Procx")
	driver := Driver(job)
	if driver == nil {
		l.Error(errors.New("Driver is not set"), "Driver is not set")
		return ctrl.Result{}, nil
	}
	if driver.KedaSupport() {
		return r.KedaReconciler(ctx, req, job, driver)
	} else {
		// this driver currently does not have a corresponding keda scaler
		// so we will need to create a deployment, and set the DAEMON flag to true
		// this will run the procx as a single deployment and poll the queue for messages
		// this is not as efficient as a keda scaler, but is the simplest way to get the job running
		return r.DeploymentReconciler(ctx, req, job, driver)
	}
}

func (r *ProcXReconciler) DeploymentReconciler(ctx context.Context, req ctrl.Request, job *procxv1alpha1.ProcX, driver driver.Driver) (ctrl.Result, error) {
	l := log.Log.WithValues("procx", job.Name)
	// as there is no keda support for this driver, we will need to create a deployment
	// and set the DAEMON flag to true
	trueval := true
	job.Spec.Daemon = &trueval
	sres, err := r.handleConfigSecret(ctx, job)
	if err != nil {
		return ctrl.Result{}, err
	}
	if sres != nil {
		return ctrl.Result{}, nil
	}
	res, err := r.handleDeployment(ctx, job)
	if err != nil {
		l.Info("Error handling Deployment", "error", err)
		return ctrl.Result{}, err
	}
	if res != nil {
		l.Info("Requeueing Procx", "result", res)
		return *res, nil
	}
	l.Info("listing pods")
	podList := &corev1.PodList{}
	listOpts := []client.ListOption{
		client.InNamespace(job.Namespace),
		client.MatchingLabels(labelsForApp(job)),
	}
	if err = r.List(ctx, podList, listOpts...); err != nil {
		l.Info("Error listing pods", "error", err)
		return ctrl.Result{}, err
	}
	podNames := getPodNames(podList.Items)
	job.Status.Pods = podNames
	job.Status.Status = "Ready"
	if err := r.Status().Update(ctx, job); err != nil {
		l.Info("Error updating Procx status", "error", err)
		return ctrl.Result{}, err
	}
	l.Info("Reconciled Procx", "status", job.Status)
	return ctrl.Result{}, nil
}

func (r *ProcXReconciler) KedaReconciler(ctx context.Context, req ctrl.Request, job *procxv1alpha1.ProcX, driver driver.Driver) (ctrl.Result, error) {
	l := log.Log.WithValues("procx", job.Name)
	sres, err := r.handleConfigSecret(ctx, job)
	if err != nil {
		return ctrl.Result{}, err
	}
	if sres != nil {
		return ctrl.Result{}, nil
	}
	if driver.HasAuth() {
		// handle trigger auth
		ares, err := r.handleTriggerAuth(ctx, job)
		if err != nil {
			return ctrl.Result{}, err
		}
		if ares != nil {
			return ctrl.Result{}, nil
		}
	}
	if *job.Spec.ScalerType == procxv1alpha1.ScalerTypeDeployment {
		l.Info("Reconciling Deployment")
		res, err := r.handleDeployment(ctx, job)
		if err != nil {
			l.Info("Error handling Deployment", "error", err)
			return ctrl.Result{}, err
		}
		if res != nil {
			l.Info("Requeueing Procx", "result", res)
			return *res, nil
		}
		sores, err := r.handleScaledObject(ctx, job)
		if err != nil {
			l.Info("Error handling ScaledObject", "error", err)
			return ctrl.Result{}, err
		}
		if sores != nil {
			l.Info("Requeueing Procx", "result", sores)
			return *sores, nil
		}
	} else if *job.Spec.ScalerType == procxv1alpha1.ScalerTypeJob {
		res, err := r.handleJob(ctx, job)
		if err != nil {
			l.Info("Error handling Job", "error", err)
			return ctrl.Result{}, err
		}
		if res != nil {
			l.Info("Requeueing Procx", "result", res)
			return *res, nil
		}
	}
	l.Info("listing pods")
	podList := &corev1.PodList{}
	listOpts := []client.ListOption{
		client.InNamespace(job.Namespace),
		client.MatchingLabels(labelsForApp(job)),
	}
	if err = r.List(ctx, podList, listOpts...); err != nil {
		l.Info("Error listing pods", "error", err)
		return ctrl.Result{}, err
	}
	podNames := getPodNames(podList.Items)
	job.Status.Pods = podNames
	job.Status.Status = "Ready"
	if err := r.Status().Update(ctx, job); err != nil {
		l.Info("Error updating Procx status", "error", err)
		return ctrl.Result{}, err
	}
	l.Info("Reconciled Procx", "status", job.Status)
	return ctrl.Result{}, nil
}

func getPodNames(list []corev1.Pod) []string {
	var pods []string
	for _, pod := range list {
		pods = append(pods, pod.Name)
	}
	return pods
}

func labelsForApp(m *procxv1alpha1.ProcX) map[string]string {
	if m.Spec.CommonLabels == nil {
		cl := make(map[string]string)
		m.Spec.CommonLabels = &cl
	}
	return labels.Merge(*m.Spec.CommonLabels, map[string]string{
		"procx": m.Name,
	})
}

func deployName(name string) string {
	return "procx-" + name
}

// SetupWithManager sets up the controller with the Manager.
func (r *ProcXReconciler) SetupWithManager(mgr ctrl.Manager) error {
	if err := kedav1alpha1.AddToScheme(mgr.GetScheme()); err != nil {
		return err
	}
	return ctrl.NewControllerManagedBy(mgr).
		For(&procxv1alpha1.ProcX{}).
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.Secret{}).
		Owns(&kedav1alpha1.ScaledObject{}).
		Owns(&kedav1alpha1.ScaledJob{}).
		Owns(&kedav1alpha1.TriggerAuthentication{}).
		Complete(r)
}
