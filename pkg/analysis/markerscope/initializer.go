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

	// Validate marker rules
	for marker, rule := range cfg.MarkerRules {
		if err := validateMarkerRule(rule); err != nil {
			fieldErrors = append(fieldErrors, field.Invalid(fldPath.Child("markerRules", marker), rule, err.Error()))
		}
	}

	return fieldErrors
}

func validateMarkerRule(rule MarkerScopeRule) error {
	// Validate scope constraint
	if rule.Scope == 0 {
		return errScopeNonZero
	}

	// Validate that scope is a valid combination of FieldScope and/or TypeScope
	validScopes := FieldScope | TypeScope
	if rule.Scope&^validScopes != 0 {
		return errInvalidScopeBits
	}

	// Validate type constraint if present
	if rule.TypeConstraint != nil {
		if err := validateTypeConstraint(rule.TypeConstraint); err != nil {
			return fmt.Errorf("invalid type constraint: %w", err)
		}
	}

	return nil
}

func validateTypeConstraint(tc *TypeConstraint) error {
	if tc == nil {
		return nil
	}

	// Validate schema types if specified
	for _, st := range tc.AllowedSchemaTypes {
		if !isValidSchemaType(st) {
			return fmt.Errorf("%w: %q", errInvalidSchemaType, st)
		}
	}

	// Validate element constraint recursively
	if tc.ElementConstraint != nil {
		if err := validateTypeConstraint(tc.ElementConstraint); err != nil {
			return fmt.Errorf("invalid element constraint: %w", err)
		}
	}

	return nil
}

func isValidSchemaType(st SchemaType) bool {
	switch st {
	case SchemaTypeInteger, SchemaTypeNumber, SchemaTypeString, SchemaTypeBoolean, SchemaTypeArray, SchemaTypeObject:
		return true
	default:
		return false
	}
}
