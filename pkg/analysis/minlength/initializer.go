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
package minlength

import (
	"fmt"

	"golang.org/x/tools/go/analysis"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/initializer"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/registry"
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
		false, // No longer CRD only, but keeping as disabled-by-default for backwards compatibility because the default behavior is for CRD-based suggestions
		validateConfig,
	)
}

func initAnalyzer(c *Config) (*analysis.Analyzer, error) {
	return newAnalyzer(c), nil
}

func validateConfig(c *Config, fldPath *field.Path) field.ErrorList {
	if c == nil {
		return field.ErrorList{}
	}

	errs := field.ErrorList{}

	switch c.PreferredSuggestionMarkerType {
	case "", PreferredSuggestionMarkerTypeKubebuilder, PreferredSuggestionMarkerTypeDeclarativeValidation:
	default:
		errs = append(errs, field.Invalid(fldPath.Child("preferredSuggestionMarkerType"), c.PreferredSuggestionMarkerType, fmt.Sprintf("must be one of %s, %s or omitted.", PreferredSuggestionMarkerTypeKubebuilder, PreferredSuggestionMarkerTypeDeclarativeValidation)))
	}

	return errs
}
