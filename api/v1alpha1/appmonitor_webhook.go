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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var appmonitorlog = logf.Log.WithName("appmonitor-resource")

func (r *AppMonitor) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-app-monitoring-galexrt-moe-v1alpha1-appmonitor,mutating=true,failurePolicy=fail,groups=app-monitoring.galexrt.moe,resources=appmonitors,verbs=create;update,versions=v1alpha1,name=mappmonitor.kb.io

var _ webhook.Defaulter = &AppMonitor{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *AppMonitor) Default() {
	appmonitorlog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
	if r.Spec.PrometheusImage == "" {
		r.Spec.PrometheusImage = "docker.io/prom/prometheus:v1.2.3"
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-app-monitoring-galexrt-moe-v1alpha1-appmonitor,mutating=false,failurePolicy=fail,groups=app-monitoring.galexrt.moe,resources=appmonitors,versions=v1alpha1,name=vappmonitor.kb.io

var _ webhook.Validator = &AppMonitor{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *AppMonitor) ValidateCreate() error {
	appmonitorlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *AppMonitor) ValidateUpdate(old runtime.Object) error {
	appmonitorlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *AppMonitor) ValidateDelete() error {
	appmonitorlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
