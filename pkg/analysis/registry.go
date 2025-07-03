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
package analysis

import (
	"fmt"
	"strings"

	"golang.org/x/tools/go/analysis"
	"gopkg.in/yaml.v3"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/commentstart"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/conditions"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/duplicatemarkers"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/initializer"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/integers"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/jsontags"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/maxlength"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/nobools"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/nofloats"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/nomaps"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/nophase"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/optionalfields"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/optionalorrequired"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/requiredfields"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/ssatags"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/statusoptional"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/statussubresource"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/uniquemarkers"
	"sigs.k8s.io/kube-api-linter/pkg/config"

	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// Registry is used to fetch and initialize analyzers.
type Registry interface {
	// DefaultLinters returns the names of linters that are enabled by default.
	DefaultLinters() sets.Set[string]

	// AllLinters returns the names of all registered linters.
	AllLinters() sets.Set[string]

	// InitializeLinters returns a set of newly initialized linters based on the
	// provided configuration.
	InitializeLinters(config.Linters, config.LintersConfig) ([]*analysis.Analyzer, error)

	// ValidateLintersConfig validates the provided linters config
	// against the set or registered linters.
	ValidateLintersConfig(config.Linters, config.LintersConfig, *field.Path) field.ErrorList
}

type registry struct {
	initializers []initializer.AnalyzerInitializer
}

// NewRegistry returns a new registry, from which analyzers can be fetched.
func NewRegistry() Registry {
	return &registry{
		initializers: []initializer.AnalyzerInitializer{
			conditions.Initializer(),
			commentstart.Initializer(),
			duplicatemarkers.Initializer(),
			integers.Initializer(),
			jsontags.Initializer(),
			maxlength.Initializer(),
			nobools.Initializer(),
			nofloats.Initializer(),
			nomaps.Initializer(),
			nophase.Initializer(),
			optionalfields.Initializer(),
			optionalorrequired.Initializer(),
			requiredfields.Initializer(),
			ssatags.Initializer(),
			statusoptional.Initializer(),
			statussubresource.Initializer(),
			uniquemarkers.Initializer(),
		},
	}
}

// DefaultLinters returns the list of linters that are registered
// as being enabled by default.
func (r *registry) DefaultLinters() sets.Set[string] {
	defaultLinters := sets.New[string]()

	for _, initializer := range r.initializers {
		if initializer.Default() {
			defaultLinters.Insert(initializer.Name())
		}
	}

	return defaultLinters
}

// AllLinters returns the list of all known linters that are known
// to the registry.
func (r *registry) AllLinters() sets.Set[string] {
	linters := sets.New[string]()

	for _, initializer := range r.initializers {
		linters.Insert(initializer.Name())
	}

	return linters
}

// InitializeLinters returns a list of initialized linters based on the provided config.
func (r *registry) InitializeLinters(cfg config.Linters, lintersCfg config.LintersConfig) ([]*analysis.Analyzer, error) {
	analyzers := []*analysis.Analyzer{}
	errs := []error{}

	for _, init := range r.getEnabledInitializers(cfg) {
		var linterConfig any

		if init.IsConfigurable() {
			var err error

			linterConfig, err = getLinterTypedConfig(init, lintersCfg)
			if err != nil {
				errs = append(errs, fmt.Errorf("failed to get linter config: %w", err))
			}
		}

		a, err := init.Init(linterConfig)
		if err != nil {
			errs = append(errs, fmt.Errorf("failed to initialize linter %s: %w", init.Name(), err))
			continue
		}

		analyzers = append(analyzers, a)
	}

	return analyzers, kerrors.NewAggregate(errs)
}

// ValidateLintersConfig validates the provided linters config
// against the set or registered linters.
func (r *registry) ValidateLintersConfig(cfg config.Linters, lintersCfg config.LintersConfig, fieldPath *field.Path) field.ErrorList {
	fieldErrors := field.ErrorList{}
	validatedLinters := sets.New[string]()

	for _, init := range r.getEnabledInitializers(cfg) {
		if init.IsConfigurable() {
			linterConfig, err := getLinterTypedConfig(init, lintersCfg)
			if err != nil {
				fieldErrors = append(fieldErrors, field.Invalid(fieldPath.Child(init.Name()), linterConfig, err.Error()))
				continue
			}

			ci, ok := init.(initializer.ConfigurableAnalyzerInitializer)
			if !ok {
				panic(fmt.Sprintf("Analyzer %s claims to be configurable but does not implement the ConfigurableAnalyzerInitializer interface", init.Name()))
			}

			fieldErrors = append(fieldErrors, ci.ValidateConfig(linterConfig, fieldPath.Child(init.Name()))...)

			validatedLinters.Insert(init.Name())
		}
	}

	fieldErrors = append(fieldErrors, validateUnusedLinters(lintersCfg, validatedLinters, fieldPath)...)

	return fieldErrors
}

// getEnabledInitializers returns the initializers that are enabled by the config.
// It returns a list of initializers that are enabled by the config.
func (r *registry) getEnabledInitializers(cfg config.Linters) []initializer.AnalyzerInitializer {
	enabled := sets.New(cfg.Enable...)
	disabled := sets.New(cfg.Disable...)

	allEnabled := enabled.Len() == 1 && enabled.Has(config.Wildcard)
	allDisabled := disabled.Len() == 1 && disabled.Has(config.Wildcard)

	initializers := []initializer.AnalyzerInitializer{}

	for _, init := range r.initializers {
		if !disabled.Has(init.Name()) && (allEnabled || enabled.Has(init.Name()) || !allDisabled && init.Default()) {
			initializers = append(initializers, init)
		}
	}

	return initializers
}

// getLinterTypedConfig returns the typed config for a linter.
func getLinterTypedConfig(init initializer.AnalyzerInitializer, lintersCfg config.LintersConfig) (any, error) {
	ci, ok := init.(initializer.ConfigurableAnalyzerInitializer)
	if !ok {
		panic(fmt.Sprintf("Analyzer %s claims to be configurable but does not implement the ConfigurableAnalyzerInitializer interface", init.Name()))
	}

	rawConfig, ok := getConfigByName(init.Name(), lintersCfg)
	if !ok {
		return ci.ConfigType(), nil
	}

	linterConfig := ci.ConfigType()

	if err := yaml.Unmarshal(rawConfig, linterConfig); err != nil {
		return nil, fmt.Errorf("error reading config for linter %q: %w", ci.Name(), err)
	}

	return linterConfig, nil
}

// getConfigByName returns the config for a linter by name.
// It returns the config as a byte slice and a boolean indicating if the config was found.
// It also supports backwards compatibility with early configuration.
// We use to have camelCased config names, but now it is all lowercase matched on the linter name.
// TODO(@JoelSpeed): Remove the strings.ToLower in a future release with a release note about the change.
func getConfigByName(name string, lintersCfg config.LintersConfig) ([]byte, bool) {
	rawConfig, ok := lintersCfg[name]
	if ok {
		return rawConfig, true
	}

	// Hack to allow backwards compatibility with early configuration.
	// We use to have camelCased config names, but now it is all lowercase matched on the linter name.
	// TODO(@JoelSpeed): Remove this in a future release with a release note about the change.
	for k, v := range lintersCfg {
		if strings.ToLower(k) == name {
			return v, true
		}
	}

	return nil, false
}

// validateUnusedLinters validates that all linters in the config are enabled.
// It returns a list of errors for each linter that is not enabled.
func validateUnusedLinters(lintersCfg config.LintersConfig, validatedLinters sets.Set[string], fieldPath *field.Path) field.ErrorList {
	fieldErrors := field.ErrorList{}

	for name := range lintersCfg {
		// Hack to allow backwards compatibility with early configuration.
		// We use to have camelCased config names, but now it is all lowercase matched on the linter name.
		// TODO(@JoelSpeed): Remove the strings.ToLower in a future release with a release note about the change.
		if !validatedLinters.Has(strings.ToLower(name)) {
			fieldErrors = append(fieldErrors, field.Invalid(fieldPath.Child(name), nil, "linter is not enabled"))
		}
	}

	return fieldErrors
}
