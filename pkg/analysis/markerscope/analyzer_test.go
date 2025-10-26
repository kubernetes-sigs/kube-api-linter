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
package markerscope

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzerWarnOnly(t *testing.T) {
	testdata := analysistest.TestData()
	cfg := &MarkerScopeConfig{
		Policy: MarkerScopePolicyWarn,
	}
	analyzer := newAnalyzer(cfg)
	analysistest.Run(t, testdata, analyzer, "a")
}

func TestAnalyzerSuggestFixes(t *testing.T) {
	testdata := analysistest.TestData()
	cfg := &MarkerScopeConfig{
		Policy: MarkerScopePolicySuggestFix,
	}
	analyzer := newAnalyzer(cfg)
	analysistest.RunWithSuggestedFixes(t, testdata, analyzer, "a")
}

func TestAnalyzerWithCustomMarkers(t *testing.T) {
	testdata := analysistest.TestData()
	cfg := &MarkerScopeConfig{
		Policy: MarkerScopePolicyWarn,
		MarkerRules: map[string]MarkerScopeRule{
			// Custom field-only marker
			"custom:field-only": {
				Scope: FieldScope,
			},
			// Custom type-only marker
			"custom:type-only": {
				Scope: TypeScope,
			},
			// Custom marker with string type constraint
			"custom:string-only": {
				Scope: FieldScope,
				TypeConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeString},
				},
			},
			// Custom marker with integer type constraint
			"custom:integer-only": {
				Scope: FieldScope,
				TypeConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeInteger},
				},
			},
			// Custom marker with array of strings constraint
			"custom:string-array": {
				Scope: FieldScope,
				TypeConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
					ElementConstraint: &TypeConstraint{
						AllowedSchemaTypes: []SchemaType{SchemaTypeString},
					},
				},
			},
		},
	}
	analyzer := newAnalyzer(cfg)
	analysistest.Run(t, testdata, analyzer, "b")
}
