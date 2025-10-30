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
		true,
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
	fieldErrors = append(fieldErrors, validatePolicy(cfg.Policy, fldPath)...)

	// Get default marker rules for validation
	defaultRules := DefaultMarkerRules()

	// Validate override marker rules
	fieldErrors = append(fieldErrors, validateOverrideMarkers(cfg.OverrideMarkers, defaultRules, fldPath)...)

	// Validate custom marker rules
	fieldErrors = append(fieldErrors, validateCustomMarkers(cfg.CustomMarkers, defaultRules, fldPath)...)

	return fieldErrors
}

func validatePolicy(policy MarkerScopePolicy, fldPath *field.Path) field.ErrorList {
	if policy != "" && policy != MarkerScopePolicyWarn && policy != MarkerScopePolicySuggestFix {
		return field.ErrorList{
			field.Invalid(fldPath.Child("policy"), policy,
				fmt.Sprintf("invalid policy, must be one of: %q, %q", MarkerScopePolicyWarn, MarkerScopePolicySuggestFix)),
		}
	}

	return nil
}

func validateOverrideMarkers(rules []MarkerScopeRule, defaultRules map[string]MarkerScopeRule, fldPath *field.Path) field.ErrorList {
	fieldErrors := field.ErrorList{}

	for i, rule := range rules {
		markerRulePath := fldPath.Child("overrideMarkers").Index(i)

		// Validate that identifier is not empty
		if rule.Identifier == "" {
			fieldErrors = append(fieldErrors, field.Required(markerRulePath.Child("identifier"), "marker identifier is required"))

			continue
		}

		// Validate that override marker exists in default rules
		if _, exists := defaultRules[rule.Identifier]; !exists {
			fieldErrors = append(fieldErrors, field.Invalid(markerRulePath.Child("identifier"), rule.Identifier,
				"override marker must be a built-in marker; use customMarkers for custom markers"))

			continue
		}

		if err := validateMarkerRule(rule); err != nil {
			fieldErrors = append(fieldErrors, field.Invalid(markerRulePath, rule, err.Error()))
		}
	}

	return fieldErrors
}

func validateCustomMarkers(rules []MarkerScopeRule, defaultRules map[string]MarkerScopeRule, fldPath *field.Path) field.ErrorList {
	fieldErrors := field.ErrorList{}

	for i, rule := range rules {
		markerRulePath := fldPath.Child("customMarkers").Index(i)

		// Validate that identifier is not empty
		if rule.Identifier == "" {
			fieldErrors = append(fieldErrors, field.Required(markerRulePath.Child("identifier"), "marker identifier is required"))

			continue
		}

		// Validate that custom marker does not exist in default rules
		if _, exists := defaultRules[rule.Identifier]; exists {
			fieldErrors = append(fieldErrors, field.Invalid(markerRulePath.Child("identifier"), rule.Identifier,
				"custom marker cannot be a built-in marker; use overrideMarkers to override built-in markers"))

			continue
		}

		if err := validateMarkerRule(rule); err != nil {
			fieldErrors = append(fieldErrors, field.Invalid(markerRulePath, rule, err.Error()))
		}
	}

	return fieldErrors
}

func validateMarkerRule(rule MarkerScopeRule) error {
	// Validate scope constraint
	if rule.Scope == "" {
		return errScopeRequired
	}

	// Validate that scope is a valid value
	switch rule.Scope {
	case FieldScope, TypeScope, AnyScope:
		// Valid scope
	default:
		return &invalidScopeConstraintError{scope: string(rule.Scope)}
	}

	// Validate named type constraint if present
	if !isValidNamedTypeConstraint(rule.NamedTypeConstraint) {
		return &invalidNamedTypeConstraintError{constraint: string(rule.NamedTypeConstraint)}
	}

	// Validate type constraint if present
	if rule.TypeConstraint != nil {
		if err := validateTypeConstraint(rule.TypeConstraint); err != nil {
			return &invalidTypeConstraintError{
				err: err,
			}
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
			return &invalidSchemaTypeError{schemaType: string(st)}
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
	case SchemaTypeInteger, SchemaTypeString, SchemaTypeBoolean, SchemaTypeArray, SchemaTypeObject:
		return true
	default:
		return false
	}
}

func isValidNamedTypeConstraint(ntc NamedTypeConstraint) bool {
	// Empty is valid (defaults to AllowField)
	if ntc == "" {
		return true
	}

	switch ntc {
	case NamedTypeConstraintAllowField, NamedTypeConstraintRequireTypeDefinition:
		return true
	default:
		return false
	}
}
