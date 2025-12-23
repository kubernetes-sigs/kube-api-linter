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
package serialization_test

import (
	"errors"
	"go/ast"
	"testing"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/extractjsontags"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/inspector"
	markershelper "sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/markers"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils/serialization"
)

var (
	errCouldNotGetInspector = errors.New("could not get inspector")
)

func TestPointersAlways(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, testSerializationAnalyzer(&serialization.Config{
		Pointers: serialization.PointersConfig{
			Policy:     serialization.PointersPolicySuggestFix,
			Preference: serialization.PointersPreferenceAlways,
		},
		OmitEmpty: serialization.OmitEmptyConfig{
			Policy: serialization.OmitEmptyPolicySuggestFix,
		},
		OmitZero: serialization.OmitZeroConfig{
			Policy: serialization.OmitZeroPolicySuggestFix, // This should make no difference as the pointer policy is always.
		},
	}), "pointers_always")
}

func TestPointersWhenRequired(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, testSerializationAnalyzer(&serialization.Config{
		Pointers: serialization.PointersConfig{
			Policy:     serialization.PointersPolicySuggestFix,
			Preference: serialization.PointersPreferenceWhenRequired,
		},
		OmitEmpty: serialization.OmitEmptyConfig{
			Policy: serialization.OmitEmptyPolicySuggestFix,
		},
		OmitZero: serialization.OmitZeroConfig{
			Policy: serialization.OmitZeroPolicyForbid, // This is the legacy behaviour before omitzero was introduced.
		},
	}), "pointers_when_required")
}

func TestPointersWhenRequiredOmitEmptyIgnore(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, testSerializationAnalyzer(&serialization.Config{
		Pointers: serialization.PointersConfig{
			Policy:     serialization.PointersPolicySuggestFix,
			Preference: serialization.PointersPreferenceWhenRequired,
		},
		OmitEmpty: serialization.OmitEmptyConfig{
			Policy: serialization.OmitEmptyPolicyIgnore,
		},
		OmitZero: serialization.OmitZeroConfig{
			Policy: serialization.OmitZeroPolicyForbid, // This is the legacy behaviour before omitzero was introduced.
		},
	}), "pointers_when_required_omit_empty_ignore")
}

func TestPointersWhenRequiredOmitEmptyIgnoreOmitZeroSuggestFix(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, testSerializationAnalyzer(&serialization.Config{
		Pointers: serialization.PointersConfig{
			Policy:     serialization.PointersPolicySuggestFix,
			Preference: serialization.PointersPreferenceWhenRequired,
		},
		OmitEmpty: serialization.OmitEmptyConfig{
			Policy: serialization.OmitEmptyPolicyIgnore,
		},
		OmitZero: serialization.OmitZeroConfig{
			Policy: serialization.OmitZeroPolicySuggestFix,
		},
	}), "pointers_when_required_omit_empty_ignore_omit_zero_suggest_fix")
}

func TestPointersWhenRequiredOmitEmptySuggestFixOmitZeroSuggestFix(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, testSerializationAnalyzer(&serialization.Config{
		Pointers: serialization.PointersConfig{
			Policy:     serialization.PointersPolicySuggestFix,
			Preference: serialization.PointersPreferenceWhenRequired,
		},
		OmitEmpty: serialization.OmitEmptyConfig{
			Policy: serialization.OmitEmptyPolicySuggestFix,
		},
		OmitZero: serialization.OmitZeroConfig{
			Policy: serialization.OmitZeroPolicySuggestFix,
		},
	}), "pointers_when_required_omit_empty_suggest_fix_omit_zero_suggest_fix")
}

func testSerializationAnalyzer(cfg *serialization.Config) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "test",
		Doc:      "test",
		Requires: []*analysis.Analyzer{inspector.Analyzer, extractjsontags.Analyzer},
		Run: func(pass *analysis.Pass) (any, error) {
			inspect, ok := pass.ResultOf[inspector.Analyzer].(inspector.Inspector)
			if !ok {
				return nil, errCouldNotGetInspector
			}

			inspect.InspectFields(func(field *ast.Field, jsonTagInfo extractjsontags.FieldTagInfo, markersAccess markershelper.Markers, qualifiedFieldName string) {
				serialization.New(cfg).Check(pass, field, markersAccess, jsonTagInfo, qualifiedFieldName)
			})

			return nil, nil
		},
	}
}
