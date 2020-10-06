/*


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

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appmonitoringv1alpha1 "github.com/galexrt/application-monitoring-operator/api/v1alpha1"
)

// AppMonitorReconciler reconciles a AppMonitor object
type AppMonitorReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=app-monitoring.galexrt.moe,resources=appmonitors,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=app-monitoring.galexrt.moe,resources=appmonitors/status,verbs=get;update;patch

func (r *AppMonitorReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("appmonitor", req.NamespacedName)

	// your logic here
	// TODO Create Prometheus object based on
	// TODO If alerting is enabled, then create Alertmanager object based on

	return ctrl.Result{}, nil
}

func (r *AppMonitorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appmonitoringv1alpha1.AppMonitor{}).
		Complete(r)
}
