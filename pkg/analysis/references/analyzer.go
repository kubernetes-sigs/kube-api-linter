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
	"golang.org/x/tools/go/analysis"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/namingconventions"
)

const name = "references"

// newAnalyzer creates a new analysis.Analyzer for the references linter.
// The references linter is implemented as a wrapper around the namingconventions
// linter with fixed configuration based on the policy.
func newAnalyzer(cfg *Config) *analysis.Analyzer {
	if cfg == nil {
		cfg = &Config{}
	}

	// Default to ForbidRefAndRefs if no policy is specified
	policy := cfg.Policy
	if policy == "" {
		policy = PolicyForbidRefAndRefs
	}

	// Build naming conventions based on policy
	conventions := buildConventions(policy)

	// Create namingconventions config
	ncConfig := &namingconventions.Config{
		Conventions: conventions,
	}

	// Create the underlying namingconventions analyzer
	ncAnalyzer := namingconventions.NewAnalyzer(ncConfig)

	// Wrap it to return our name
	analyzer := &analysis.Analyzer{
		Name:     name,
		Doc:      "Enforces that fields use Ref/Refs and not Reference/References",
		Run:      ncAnalyzer.Run,
		Requires: ncAnalyzer.Requires,
	}

	return analyzer
}

// buildConventions creates the naming conventions based on the policy
func buildConventions(policy Policy) []namingconventions.Convention {
	conventions := []namingconventions.Convention{
		// Match "References" anywhere (case-sensitive, capital R)
		{
			Name:             "references-to-refs",
			ViolationMatcher: "References",
			Operation:        namingconventions.OperationReplacement,
			Replacement:      "Refs",
			Message:          "field names should use 'Refs' instead of 'References'",
		},
		// Match "Reference" but not when part of "References"
		// Match Reference followed by: end-of-string OR non-'s' character
		{
			Name:             "reference-to-ref",
			ViolationMatcher: "Reference([^s]|$)",
			Operation:        namingconventions.OperationReplacement,
			Replacement:      "Ref$1",
			Message:          "field names should use 'Ref' instead of 'Reference'",
		},
	}

	// If policy is ForbidRefAndRefs, add conventions to forbid Ref/Refs anywhere
	// Exclude patterns already handled by Reference/References above
	if policy == PolicyForbidRefAndRefs {
		conventions = append(conventions,
			// Match "Refs" but not when part of "References"
			namingconventions.Convention{
				Name:             "forbid-refs",
				ViolationMatcher: "Refs([^a-z]|$)",
				Operation:        namingconventions.OperationInform,
				Message:          "should not use 'Refs'",
			},
			// Match "Ref" but not when part of "Reference", "References", or "Refs"
			namingconventions.Convention{
				Name:             "forbid-ref",
				ViolationMatcher: "Ref([^ers]|$)",
				Operation:        namingconventions.OperationInform,
				Message:          "should not use 'Ref'",
			},
		)
	}

	return conventions
}
