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
package markerscope

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
		false, // Not enabled by default
		validateConfig,
	)
}

func initAnalyzer(cfg *MarkerScopeConfig) (*analysis.Analyzer, error) {
	return newAnalyzer(cfg), nil
}

// validateConfig validates the configuration in the MarkerScopeConfig struct.
func validateConfig(cfg *MarkerScopeConfig, fldPath *field.Path) field.ErrorList {
	if cfg == nil {
		return field.ErrorList{}
	}

	fieldErrors := field.ErrorList{}

	// Validate policy
	if cfg.Policy != "" && cfg.Policy != MarkerScopePolicyWarn && cfg.Policy != MarkerScopePolicySuggestFix {
		fieldErrors = append(fieldErrors, field.Invalid(fldPath.Child("policy"), cfg.Policy,
			fmt.Sprintf("invalid policy, must be one of: %q, %q", MarkerScopePolicyWarn, MarkerScopePolicySuggestFix)))
	}

	// Validate marker scopes
	for marker, scope := range cfg.Markers {
		if err := validateScope(scope); err != nil {
			fieldErrors = append(fieldErrors, field.Invalid(fldPath.Child("markers", marker), scope, err.Error()))
		}
	}

	return fieldErrors
}

func validateScope(scope MarkerScope) error {
	switch scope {
	case ScopeField, ScopeType, ScopeFieldOrType, ScopeTypeOrObjectField:
		return nil
	default:
		return fmt.Errorf("invalid scope %q, must be one of: %q, %q, %q, %q", scope, ScopeField, ScopeType, ScopeFieldOrType, ScopeTypeOrObjectField)
	}
}
