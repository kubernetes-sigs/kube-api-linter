/*
Copyright 2025 The Kubernetes Authors.

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
package maxlength

import (
	"fmt"

	"golang.org/x/tools/go/analysis"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/initializer"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/registry"
	"sigs.k8s.io/kube-api-linter/pkg/markers"
)

func init() {
	registry.DefaultRegistry().RegisterLinter(Initializer())
}

// Initializer returns the AnalyzerInitializer for this
// Analyzer so that it can be added to the registry.
func Initializer() initializer.AnalyzerInitializer {
	return initializer.NewConfigurableInitializer(
		name,
		initAnalyzer,
		false, // For now, CRD only, and so not on by default.
		validateConfig,
	)
}

func initAnalyzer(cfg *MaxLengthConfig) (*analysis.Analyzer, error) {
	return newAnalyzer(cfg), nil
}

// validateConfig validates the configuration in the MaxLengthConfig struct.
func validateConfig(cfg *MaxLengthConfig, fldPath *field.Path) field.ErrorList {
	if cfg == nil {
		return field.ErrorList{}
	}

	errs := field.ErrorList{}

	validateMarkerChoice(
		fldPath.Child("preferredMaxLengthMarker"),
		cfg.PreferredMaxLengthMarker,
		markers.KubebuilderMaxLengthMarker,
		markers.K8sMaxLengthMarker,
		&errs,
	)
	validateMarkerChoice(
		fldPath.Child("preferredMaxItemsMarker"),
		cfg.PreferredMaxItemsMarker,
		markers.KubebuilderMaxItemsMarker,
		markers.K8sMaxItemsMarker,
		&errs,
	)
	validateMarkerChoice(
		fldPath.Child("preferredMaxPropertiesMarker"),
		cfg.PreferredMaxPropertiesMarker,
		markers.KubebuilderMaxPropertiesMarker,
		markers.K8sMaxPropertiesMarker,
		&errs,
	)

	return errs
}

// validateMarkerChoice validates that value is either empty, the kubebuilder form, or the k8s DV form.
func validateMarkerChoice(fld *field.Path, value, kubebuilder, k8s string, errs *field.ErrorList) {
	switch value {
	case "", kubebuilder, k8s:
		// Valid.
	default:
		*errs = append(*errs, field.Invalid(
			fld,
			value,
			fmt.Sprintf("invalid value, must be one of %q, %q or omitted", kubebuilder, k8s),
		))
	}
}
