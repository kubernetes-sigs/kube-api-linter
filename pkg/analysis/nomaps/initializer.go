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
package nomaps

import (
	"fmt"

	"golang.org/x/tools/go/analysis"
	"k8s.io/apimachinery/pkg/util/validation/field"
	kalanalysis "sigs.k8s.io/kube-api-linter/pkg/analysis"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/initializer"
)

func init() {
	kalanalysis.DefaultRegistry().RegisterLinter(Initializer())
}

// Initializer returns the AnalyzerInitializer for this
// Analyzer so that it can be added to the registry.
func Initializer() initializer.AnalyzerInitializer {
	return initializer.NewConfigurableInitializer(
		name,
		initAnalyzer,
		true,
		func() any { return &NoMapsConfig{} },
		validateConfig,
	)
}

// Init returns the intialized Analyzer.
func initAnalyzer(cfg any) (*analysis.Analyzer, error) {
	nmc, ok := cfg.(*NoMapsConfig)
	if !ok {
		return nil, fmt.Errorf("failed to initialize no maps analyzer: %w", initializer.NewIncorrectTypeError(cfg))
	}

	return newAnalyzer(nmc), nil
}

// validateConfig is used to validate the configuration in the config.NoMapsConfig struct.
func validateConfig(cfg any, fldPath *field.Path) field.ErrorList {
	nmc, ok := cfg.(*NoMapsConfig)
	if !ok {
		return field.ErrorList{field.InternalError(fldPath, initializer.NewIncorrectTypeError(cfg))}
	}

	fieldErrors := field.ErrorList{}

	switch nmc.Policy {
	case NoMapsEnforce, NoMapsAllowStringToStringMaps, NoMapsIgnore, "":
		// Valid values
	default:
		fieldErrors = append(fieldErrors, field.Invalid(fldPath.Child("policy"), nmc.Policy, fmt.Sprintf("invalid value, must be one of %q, %q, %q or omitted", NoMapsEnforce, NoMapsAllowStringToStringMaps, NoMapsIgnore)))
	}

	return fieldErrors
}
