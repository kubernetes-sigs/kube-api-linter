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

	"golang.org/x/tools/go/analysis"
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
//
//nolint:cyclop // Temporary while we refactor linter validation and registration into the registry.
func (r *registry) InitializeLinters(cfg config.Linters, lintersCfg config.LintersConfig) ([]*analysis.Analyzer, error) {
	analyzers := []*analysis.Analyzer{}
	errs := []error{}

	enabled := sets.New(cfg.Enable...)
	disabled := sets.New(cfg.Disable...)

	allEnabled := enabled.Len() == 1 && enabled.Has(config.Wildcard)
	allDisabled := disabled.Len() == 1 && disabled.Has(config.Wildcard)

	for _, init := range r.initializers {
		if init.IsConfigurable() {
			_, ok := init.(initializer.ConfigurableAnalyzerInitializer)
			if !ok {
				panic(fmt.Sprintf("Analyzer %s claims to be configurable but does not implement the ConfigurableAnalyzerInitializer interface", init.Name()))
			}
		}

		if !disabled.Has(init.Name()) && (allEnabled || enabled.Has(init.Name()) || !allDisabled && init.Default()) {
			a, err := init.Init(lintersCfg)
			if err != nil {
				errs = append(errs, fmt.Errorf("failed to initialize linter %s: %w", init.Name(), err))
				continue
			}

			analyzers = append(analyzers, a)
		}
	}

	return analyzers, kerrors.NewAggregate(errs)
}
