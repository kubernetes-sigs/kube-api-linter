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
package uniquemarkers

import (
	"golang.org/x/tools/go/analysis"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/initializer"
	"sigs.k8s.io/kube-api-linter/pkg/config"
)

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

// Init returns the intialized Analyzer.
func initAnalyzer(cfg config.LintersConfig) (*analysis.Analyzer, error) {
	return newAnalyzer(cfg.UniqueMarkers), nil
}

// validateConfig validates the configuration in the config.UniqueMarkersConfig struct.
func validateConfig(cfg any, fldPath *field.Path) field.ErrorList {
	umc, ok := cfg.(config.UniqueMarkersConfig)
	if !ok {
		return field.ErrorList{field.InternalError(fldPath, initializer.NewIncorrectTypeError(cfg))}
	}

	fieldErrors := field.ErrorList{}
	identifierSet := sets.New[string]()

	for i, marker := range umc.CustomMarkers {
		if identifierSet.Has(marker.Identifier) {
			fieldErrors = append(fieldErrors, field.Invalid(fldPath.Child("customMarkers").Index(i).Child("identifier"), marker.Identifier, "repeated value, values must be unique"))
			continue
		}

		fieldErrors = append(fieldErrors, validateUniqueMarker(marker, fldPath.Child("customMarkers").Index(i))...)

		identifierSet.Insert(marker.Identifier)
	}

	return fieldErrors
}

func validateUniqueMarker(um config.UniqueMarker, fldPath *field.Path) field.ErrorList {
	fieldErrors := field.ErrorList{}
	attrSet := sets.New[string]()

	for i, attr := range um.Attributes {
		if attrSet.Has(attr) {
			fieldErrors = append(fieldErrors, field.Invalid(fldPath.Child("attributes").Index(i), attr, "repeated value, values must be unique"))
		}
	}

	return fieldErrors
}
