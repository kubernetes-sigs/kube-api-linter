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
package conflictingmarkers

import (
	"golang.org/x/tools/go/analysis"
	"k8s.io/apimachinery/pkg/util/sets"
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

// initAnalyzer returns the initialized Analyzer.
func initAnalyzer(cfg *ConflictingMarkersConfig) (*analysis.Analyzer, error) {
	return newAnalyzer(cfg), nil
}

// validateConfig validates the configuration in the config.ConflictingMarkersConfig struct.
func validateConfig(cfg *ConflictingMarkersConfig, fldPath *field.Path) field.ErrorList {
	if cfg == nil {
		return field.ErrorList{}
	}

	fieldErrors := field.ErrorList{}
	nameSet := sets.New[string]()

	for i, conflictSet := range cfg.CustomConflicts {
		if nameSet.Has(conflictSet.Name) {
			fieldErrors = append(fieldErrors, field.Invalid(fldPath.Child("customConflicts").Index(i).Child("name"), conflictSet.Name, "repeated value, names must be unique"))
			continue
		}

		fieldErrors = append(fieldErrors, validateConflictSet(conflictSet, fldPath.Child("customConflicts").Index(i))...)

		nameSet.Insert(conflictSet.Name)
	}

	return fieldErrors
}

func validateConflictSet(conflictSet ConflictSet, fldPath *field.Path) field.ErrorList {
	fieldErrors := field.ErrorList{}

	if conflictSet.Name == "" {
		fieldErrors = append(fieldErrors, field.Required(fldPath.Child("name"), "name is required"))
	}

	if conflictSet.Description == "" {
		fieldErrors = append(fieldErrors, field.Required(fldPath.Child("description"), "description is required"))
	}

	if len(conflictSet.SetA) == 0 {
		fieldErrors = append(fieldErrors, field.Required(fldPath.Child("setA"), "setA cannot be empty"))
	}

	if len(conflictSet.SetB) == 0 {
		fieldErrors = append(fieldErrors, field.Required(fldPath.Child("setB"), "setB cannot be empty"))
	}

	// Check for overlapping markers between sets
	setA := sets.New(conflictSet.SetA...)
	setB := sets.New(conflictSet.SetB...)

	if intersection := setA.Intersection(setB); intersection.Len() > 0 {
		fieldErrors = append(fieldErrors, field.Invalid(fldPath, conflictSet, "sets cannot contain overlapping markers"))
	}

	return fieldErrors
}
