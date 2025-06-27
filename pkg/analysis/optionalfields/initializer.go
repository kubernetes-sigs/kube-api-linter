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
package optionalfields

import (
	"fmt"

	"golang.org/x/tools/go/analysis"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/kube-api-linter/pkg/config"
)

// Initializer returns the AnalyzerInitializer for this
// Analyzer so that it can be added to the registry.
func Initializer() initializer {
	return initializer{}
}

// intializer implements the AnalyzerInitializer interface.
type initializer struct{}

// Name returns the name of the Analyzer.
func (initializer) Name() string {
	return name
}

// Init returns the intialized Analyzer.
func (initializer) Init(cfg config.LintersConfig) (*analysis.Analyzer, error) {
	return newAnalyzer(cfg.OptionalFields), nil
}

// ValidateConfig validates the configuration in the config.OptionalFieldsConfig struct.
func (initializer) ValidateConfig(cfg any, fldPath *field.Path) field.ErrorList {
	ofc, ok := cfg.(config.OptionalFieldsConfig)
	if !ok {
		return field.ErrorList{field.InternalError(fldPath, fmt.Errorf("incorrect type for passed configuration: %T", cfg))}
	}

	fieldErrors := field.ErrorList{}

	fieldErrors = append(fieldErrors, validateOptionFieldsPointers(ofc.Pointers, fldPath.Child("pointers"))...)
	fieldErrors = append(fieldErrors, validateOptionFieldsOmitEmpty(ofc.OmitEmpty, fldPath.Child("omitEmpty"))...)

	return fieldErrors
}

// validateOptionFieldsPointers is used to validate the configuration in the config.OptionalFieldsPointers struct.
func validateOptionFieldsPointers(opc config.OptionalFieldsPointers, fldPath *field.Path) field.ErrorList {
	fieldErrors := field.ErrorList{}

	switch opc.Preference {
	case "", config.OptionalFieldsPointerPreferenceAlways, config.OptionalFieldsPointerPreferenceWhenRequired:
	default:
		fieldErrors = append(fieldErrors, field.Invalid(fldPath.Child("preference"), opc.Preference, fmt.Sprintf("invalid value, must be one of %q, %q or omitted", config.OptionalFieldsPointerPreferenceAlways, config.OptionalFieldsPointerPreferenceWhenRequired)))
	}

	switch opc.Policy {
	case "", config.OptionalFieldsPointerPolicySuggestFix, config.OptionalFieldsPointerPolicyWarn:
	default:
		fieldErrors = append(fieldErrors, field.Invalid(fldPath.Child("policy"), opc.Policy, fmt.Sprintf("invalid value, must be one of %q, %q or omitted", config.OptionalFieldsPointerPolicySuggestFix, config.OptionalFieldsPointerPolicyWarn)))
	}

	return fieldErrors
}

// validateOptionFieldsOmitEmpty is used to validate the configuration in the config.OptionalFieldsOmitEmpty struct.
func validateOptionFieldsOmitEmpty(oec config.OptionalFieldsOmitEmpty, fldPath *field.Path) field.ErrorList {
	fieldErrors := field.ErrorList{}

	switch oec.Policy {
	case "", config.OptionalFieldsOmitEmptyPolicyIgnore, config.OptionalFieldsOmitEmptyPolicyWarn, config.OptionalFieldsOmitEmptyPolicySuggestFix:
	default:
		fieldErrors = append(fieldErrors, field.Invalid(fldPath.Child("policy"), oec.Policy, fmt.Sprintf("invalid value, must be one of %q, %q, %q or omitted", config.OptionalFieldsOmitEmptyPolicyIgnore, config.OptionalFieldsOmitEmptyPolicyWarn, config.OptionalFieldsOmitEmptyPolicySuggestFix)))
	}

	return fieldErrors
}

// Default determines whether this Analyzer is on by default, or not.
func (initializer) Default() bool {
	return true
}

// IsConfigurable determines whether or not the Analyzer provides configuration options.
func (initializer) IsConfigurable() bool {
	return true
}
