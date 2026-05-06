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

func initAnalyzer(config *MaxLengthConfig) (*analysis.Analyzer, error) {
	return newAnalyzer(config), nil
}

// validateConfig is used to validate the configuration in the config.MaxLengthConfig struct.
func validateConfig(config *MaxLengthConfig, fldPath *field.Path) field.ErrorList {
	if config == nil {
		return field.ErrorList{}
	}

	fieldErrors := field.ErrorList{}

	switch config.PreferredMaxLengthMarker {
	case "", markers.KubebuilderMaxLengthMarker, markers.K8sMaxLengthMarker:
	default:
		fieldErrors = append(fieldErrors, field.Invalid(fldPath.Child("preferredMaxLengthMarker"), config.PreferredMaxLengthMarker, fmt.Sprintf("invalid value, must be one of %q, %q or omitted", markers.KubebuilderMaxLengthMarker, markers.K8sMaxLengthMarker)))
	}

	switch config.PreferredMaxItemsMarker {
	case "", markers.KubebuilderMaxItemsMarker, markers.K8sMaxItemsMarker:
	default:
		fieldErrors = append(fieldErrors, field.Invalid(fldPath.Child("preferredMaxItemsMarker"), config.PreferredMaxItemsMarker, fmt.Sprintf("invalid value, must be one of %q, %q or omitted", markers.KubebuilderMaxItemsMarker, markers.K8sMaxItemsMarker)))
	}

	switch config.PreferredMaxPropertiesMarker {
	case "", markers.KubebuilderMaxPropertiesMarker, markers.K8sMaxPropertiesMarker:
	default:
		fieldErrors = append(fieldErrors, field.Invalid(fldPath.Child("preferredMaxPropertiesMarker"), config.PreferredMaxPropertiesMarker, fmt.Sprintf("invalid value, must be one of %q, %q or omitted", markers.KubebuilderMaxPropertiesMarker, markers.K8sMaxPropertiesMarker)))
	}

	switch config.PreferredMaximumMarker {
	case "", markers.KubebuilderMaximumMarker, markers.K8sMaximumMarker:
	default:
		fieldErrors = append(fieldErrors, field.Invalid(fldPath.Child("preferredMaximumMarker"), config.PreferredMaximumMarker, fmt.Sprintf("invalid value, must be one of %q, %q or omitted", markers.KubebuilderMaximumMarker, markers.K8sMaximumMarker)))
	}

	return fieldErrors
}
