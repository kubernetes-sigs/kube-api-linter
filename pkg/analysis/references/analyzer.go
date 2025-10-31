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
package references

import (
	"errors"
	"fmt"

	"golang.org/x/tools/go/analysis"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/initializer"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/namingconventions"
)

const (
	name = "references"
	doc  = "Enforces that fields use Ref/Refs and not Reference/References"
)

var errUnexpectedInitializerType = errors.New("expected namingconventions.Initializer() to be of type initializer.ConfigurableAnalyzerInitializer, but was not")

// newAnalyzer creates a new analyzer for the references linter that is a wrapper around the namingconventions linter.
func newAnalyzer(cfg *Config) *analysis.Analyzer {
	if cfg == nil {
		cfg = &Config{}
	}

	// Default to PreferAbbreviatedReference if no policy is specified
	policy := cfg.Policy
	if policy == "" {
		policy = PolicyPreferAbbreviatedReference
	}

	// Build the namingconventions config based on the policy
	ncConfig := &namingconventions.Config{
		Conventions: buildConventions(policy),
	}

	// Get the configurable initializer for namingconventions
	configInit, ok := namingconventions.Initializer().(initializer.ConfigurableAnalyzerInitializer)
	if !ok {
		panic(fmt.Errorf("getting initializer: %w", errUnexpectedInitializerType))
	}

	// Validate generated namingconventions configuration
	errs := configInit.ValidateConfig(ncConfig, field.NewPath("references"))
	if err := errs.ToAggregate(); err != nil {
		panic(fmt.Errorf("references linter has an invalid namingconventions configuration: %w", err))
	}

	// Initialize the wrapped analyzer
	analyzer, err := configInit.Init(ncConfig)
	if err != nil {
		panic(fmt.Errorf("references linter encountered an error initializing wrapped namingconventions analyzer: %w", err))
	}

	analyzer.Name = name
	analyzer.Doc = doc

	return analyzer
}

// buildConventions creates the naming conventions based on the policy.
func buildConventions(policy Policy) []namingconventions.Convention {
	switch policy {
	case PolicyPreferAbbreviatedReference:
		// Replace "Reference" or "References" with "Ref" or "Refs"
		// Using a single regex with optional 's' capture group to handle both cases
		return []namingconventions.Convention{
			{
				Name:             "reference-to-ref",
				ViolationMatcher: "(?i)reference(s?)",
				Operation:        namingconventions.OperationReplacement,
				Replacement:      "Ref$1",
				Message:          "field names should use 'Ref' instead of 'Reference'",
			},
		}

	case PolicyNoReferences:
		// Drop any reference-related words from field names
		// Using a regex that matches Ref/Refs/Reference/References at start or end
		// At start: matches when followed by an uppercase letter, preserving that letter
		// At end: matches at the end of the field name
		return []namingconventions.Convention{
			{
				Name:             "no-references",
				ViolationMatcher: "^([Rr]ef(?:erence)?s?)([A-Z])|([Rr]ef(?:erence)?s?)$",
				Operation:        namingconventions.OperationReplacement,
				Replacement:      "$2",
				Message:          "field names should not contain reference-related words",
			},
		}

	default:
		// Should not happen due to validation, but return PreferAbbreviatedReference conventions as fallback
		return []namingconventions.Convention{
			{
				Name:             "reference-to-ref",
				ViolationMatcher: "(?i)reference(s?)",
				Operation:        namingconventions.OperationReplacement,
				Replacement:      "Ref$1",
				Message:          "field names should use 'Ref' instead of 'Reference'",
			},
		}
	}
}
