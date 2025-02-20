// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package controllers

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/cluster-api/util"
	"sigs.k8s.io/cluster-api/util/patch"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	infrav1 "github.com/talos-systems/sidero/app/caps-controller-manager/api/v1alpha3"
	"github.com/talos-systems/sidero/app/caps-controller-manager/pkg/constants"
)

// MetalClusterReconciler reconciles a MetalCluster object.
type MetalClusterReconciler struct {
	client.Client
	Log      logr.Logger
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

// +kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=metalclusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=metalclusters/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=cluster.x-k8s.io,resources=clusters;clusters/status,verbs=get;list;watch

func (r *MetalClusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (_ ctrl.Result, err error) {
	log := r.Log.WithValues("metalcluster", req.NamespacedName)

	// Fetch the metalCluster instance

	metalCluster := &infrav1.MetalCluster{}

	err = r.Get(ctx, req.NamespacedName, metalCluster)
	if apierrors.IsNotFound(err) {
		return reconcile.Result{}, nil
	}

	if err != nil {
		return ctrl.Result{}, err
	}

	log = log.WithName(metalCluster.APIVersion)

	// Fetch the Cluster
	cluster, err := util.GetOwnerCluster(ctx, r.Client, metalCluster.ObjectMeta)
	if err != nil {
		return ctrl.Result{}, err
	}

	if cluster == nil {
		log.Info("Cluster Controller has not yet set OwnerRef")
		return ctrl.Result{RequeueAfter: constants.DefaultRequeueAfter}, nil
	}

	log = log.WithName(fmt.Sprintf("cluster=%s", cluster.Name))

	// Initialize the patch helper
	patchHelper, err := patch.NewHelper(metalCluster, r.Client)
	if err != nil {
		return ctrl.Result{}, err
	}

	// Always attempt to Patch the metalCluster object and status after each reconciliation.
	defer func() {
		if e := patchHelper.Patch(ctx, metalCluster); err != nil {
			log.Error(err, "failed to patch metalCluster")

			if e == nil {
				err = e
			}
		}
	}()

	// If the MetalCluster doesn't have our finalizer, add it.
	controllerutil.AddFinalizer(metalCluster, infrav1.ClusterFinalizer)

	// Handle deleted machines
	if !metalCluster.ObjectMeta.DeletionTimestamp.IsZero() {
		log.Info("deleting cluster")

		// Cluster is deleted so remove the finalizer.
		controllerutil.RemoveFinalizer(metalCluster, infrav1.ClusterFinalizer)

		return ctrl.Result{}, nil
	}

	metalCluster.Status.Ready = true

	return ctrl.Result{}, nil
}

func (r *MetalClusterReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager, options controller.Options) error {
	return ctrl.NewControllerManagedBy(mgr).
		WithOptions(options).
		For(&infrav1.MetalCluster{}).
		Complete(r)
}
