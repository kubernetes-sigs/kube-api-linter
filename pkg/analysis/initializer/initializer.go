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
package initializer

import (
	"golang.org/x/tools/go/analysis"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/kube-api-linter/pkg/config"
)

// InitializerFunc is a function that initializes an Analyzer.
type InitializerFunc func(config.LintersConfig) (*analysis.Analyzer, error)

// ValidateFunc is a function that validates the configuration for an Analyzer.
type ValidateFunc func(any, *field.Path) field.ErrorList

// AnalyzerInitializer is used to initialize analyzers.
type AnalyzerInitializer interface {
	// Name returns the name of the analyzer initialized by this initializer.
	Name() string

	// Init returns the newly initialized analyzer.
	// It will be passed the complete LintersConfig and is expected to rely only on its own configuration.
	Init(config.LintersConfig) (*analysis.Analyzer, error)

	// IsConfigurable determines whether or not the initializer expects to be provided a config.
	// When true, the initializer should also match the ConfigurableAnalyzerInitializer interface.
	IsConfigurable() bool

	// Default determines whether the inializer intializes an analyzer that should be
	// on by default, or not.
	Default() bool
}

// ConfigurableAnalyzerInitializer is an analyzer initializer that also has a configuration.
// This means it can validate its config.
type ConfigurableAnalyzerInitializer interface {
	AnalyzerInitializer

	// ValidateConfig will be called during the config validation phase and is used to validate
	// the provided config for the linter.
	ValidateConfig(any, *field.Path) field.ErrorList
}

// NewInitializer construct a new initializer for initializing an Analyzer.
func NewInitializer(name string, initFunc InitializerFunc, isDefault bool) AnalyzerInitializer {
	return initializer{
		name:      name,
		initFunc:  initFunc,
		isDefault: isDefault,
	}
}

// NewConfigurableInitializer constructs a new initializer for intializing a
// configurable Analyzer.
func NewConfigurableInitializer(name string, initFunc InitializerFunc, isDefault bool, validateFunc ValidateFunc) ConfigurableAnalyzerInitializer {
	return configurableInitializer{
		initializer: initializer{
			name:      name,
			initFunc:  initFunc,
			isDefault: isDefault,
		},
		validateFunc: validateFunc,
	}
}

type initializer struct {
	name      string
	initFunc  InitializerFunc
	isDefault bool
}

// Name returns the name of the initializer.
func (i initializer) Name() string {
	return i.name
}

// Init returns a newly initializr analyzer.
func (i initializer) Init(cfg config.LintersConfig) (*analysis.Analyzer, error) {
	return i.initFunc(cfg)
}

// IsConfigurable determines whether or not to expect this initializer to
// be able to be configured with custom configuration.
func (i initializer) IsConfigurable() bool {
	return false
}

// Default determines whether this initializer should be enabled by default or not.
func (i initializer) Default() bool {
	return i.isDefault
}

type configurableInitializer struct {
	initializer

	validateFunc ValidateFunc
}

// IsConfigurable determines whether or not to expect this initializer to
// be able to be configured with custom configuration.
func (i configurableInitializer) IsConfigurable() bool {
	return true
}

// ValidateConfig validates the configuration for the initializer.
func (i configurableInitializer) ValidateConfig(cfg any, fld *field.Path) field.ErrorList {
	return i.validateFunc(cfg, fld)
}
