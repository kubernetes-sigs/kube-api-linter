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
	return newAnalyzer(cfg.NoMaps), nil
}

// IsConfigurable determines whether or not the Analyzer provides configuration options.
func (initializer) IsConfigurable() bool {
	return true
}

// ValidateConfig is used to validate the configuration in the config.NoMapsConfig struct.
func (initializer) ValidateConfig(cfg any, fldPath *field.Path) field.ErrorList {
	nmc, ok := cfg.(config.NoMapsConfig)
	if !ok {
		return field.ErrorList{field.InternalError(fldPath, fmt.Errorf("incorrect type for passed configuration: %T", cfg))}
	}

	fieldErrors := field.ErrorList{}

	switch nmc.Policy {
	case config.NoMapsEnforce, config.NoMapsAllowStringToStringMaps, config.NoMapsIgnore, "":
		// Valid values
	default:
		fieldErrors = append(fieldErrors, field.Invalid(fldPath.Child("policy"), nmc.Policy, fmt.Sprintf("invalid value, must be one of %q, %q, %q or omitted", config.NoMapsEnforce, config.NoMapsAllowStringToStringMaps, config.NoMapsIgnore)))
	}

	return fieldErrors
}

// Default determines whether this Analyzer is on by default, or not.
func (initializer) Default() bool {
	return true
}
