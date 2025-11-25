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
package markerscope_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/markerscope"
)

func TestAnalyzerWithDefaultConfig(t *testing.T) {
	testdata := analysistest.TestData()
	// Test with nil config - should use all defaults:
	// - Policy: Warn
	// - OverrideMarkers: empty (use built-in defaults)
	// - CustomMarkers: empty
	analyzer, err := markerscope.Initializer().Init(&markerscope.MarkerScopeConfig{})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.Run(t, testdata, analyzer, "a")
}

func TestAnalyzerSuggestFixes(t *testing.T) {
	testdata := analysistest.TestData()
	cfg := &markerscope.MarkerScopeConfig{
		Policy: markerscope.MarkerScopePolicySuggestFix,
	}

	analyzer, err := markerscope.Initializer().Init(cfg)
	if err != nil {
		t.Fatal(err)
	}

	analysistest.RunWithSuggestedFixes(t, testdata, analyzer, "a")
}

func TestAnalyzerWithCustomAndOverrideMarkers(t *testing.T) {
	testdata := analysistest.TestData()
	cfg := &markerscope.MarkerScopeConfig{
		Policy: markerscope.MarkerScopePolicyWarn,
		OverrideMarkers: []markerscope.MarkerScopeRule{
			// Override built-in "optional" to allow on types (default is FieldScope only)
			{
				Identifier: "optional",
				Scopes:     []markerscope.ScopeConstraint{markerscope.FieldScope, markerscope.TypeScope},
			},
			// Override built-in "required" to allow on types (default is FieldScope only)
			{
				Identifier: "required",
				Scopes:     []markerscope.ScopeConstraint{markerscope.FieldScope, markerscope.TypeScope},
			},
		},
		CustomMarkers: []markerscope.MarkerScopeRule{
			// Custom field-only marker
			{
				Identifier: "custom:field-only",
				Scopes:     []markerscope.ScopeConstraint{markerscope.FieldScope},
			},
			// Custom type-only marker
			{
				Identifier: "custom:type-only",
				Scopes:     []markerscope.ScopeConstraint{markerscope.TypeScope},
			},
			// Custom marker with string type constraint
			{
				Identifier: "custom:string-only",
				Scopes:     []markerscope.ScopeConstraint{markerscope.FieldScope},
				TypeConstraint: &markerscope.TypeConstraint{
					AllowedSchemaTypes: []markerscope.SchemaType{
						markerscope.SchemaTypeString,
					},
				},
			},
			// Custom marker with integer type constraint
			{
				Identifier: "custom:integer-only",
				Scopes:     []markerscope.ScopeConstraint{markerscope.FieldScope},
				TypeConstraint: &markerscope.TypeConstraint{
					AllowedSchemaTypes: []markerscope.SchemaType{
						markerscope.SchemaTypeInteger,
					},
				},
			},
			// Custom marker with array of strings constraint
			{
				Identifier: "custom:string-array",
				Scopes:     []markerscope.ScopeConstraint{markerscope.FieldScope},
				TypeConstraint: &markerscope.TypeConstraint{
					AllowedSchemaTypes: []markerscope.SchemaType{
						markerscope.SchemaTypeArray,
					},
					ElementConstraint: &markerscope.TypeConstraint{
						AllowedSchemaTypes: []markerscope.SchemaType{
							markerscope.SchemaTypeString,
						},
					},
				},
			},
		},
	}

	analyzer, err := markerscope.Initializer().Init(cfg)
	if err != nil {
		t.Fatal(err)
	}

	analysistest.Run(t, testdata, analyzer, "b")
}
