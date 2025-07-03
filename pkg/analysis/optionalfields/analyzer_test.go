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
package optionalfields_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/optionalfields"
)

func TestDefaultConfiguration(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := optionalfields.Initializer().Init(&optionalfields.OptionalFieldsConfig{})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.RunWithSuggestedFixes(t, testdata, a, "a")
}

func TestWhenRequiredPreferenceConfiguration(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := optionalfields.Initializer().Init(&optionalfields.OptionalFieldsConfig{
		Pointers: optionalfields.OptionalFieldsPointers{
			Preference: optionalfields.OptionalFieldsPointerPreferenceWhenRequired,
		},
		OmitZero: optionalfields.OptionalFieldsOmitZero{
			Policy: optionalfields.OptionalFieldsOmitZeroPolicyForbid,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.RunWithSuggestedFixes(t, testdata, a, "b")
}

func TestWhenRequiredWithOmitEmptyIgnorePreferenceConfiguration(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := optionalfields.Initializer().Init(&optionalfields.OptionalFieldsConfig{
		Pointers: optionalfields.OptionalFieldsPointers{
			Preference: optionalfields.OptionalFieldsPointerPreferenceWhenRequired,
		},
		OmitEmpty: optionalfields.OptionalFieldsOmitEmpty{
			Policy: optionalfields.OptionalFieldsOmitEmptyPolicyIgnore,
		},
		OmitZero: optionalfields.OptionalFieldsOmitZero{
			Policy: optionalfields.OptionalFieldsOmitZeroPolicyForbid,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.RunWithSuggestedFixes(t, testdata, a, "c")
}

func TestWithSuggestFixForOmitZero(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := optionalfields.Initializer().Init(&optionalfields.OptionalFieldsConfig{
		Pointers: optionalfields.OptionalFieldsPointers{
			Preference: optionalfields.OptionalFieldsPointerPreferenceWhenRequired,
			Policy:     optionalfields.OptionalFieldsPointerPolicySuggestFix,
		},
		OmitEmpty: optionalfields.OptionalFieldsOmitEmpty{
			Policy: optionalfields.OptionalFieldsOmitEmptyPolicySuggestFix,
		},
		OmitZero: optionalfields.OptionalFieldsOmitZero{
			Policy: optionalfields.OptionalFieldsOmitZeroPolicySuggestFix,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.RunWithSuggestedFixes(t, testdata, a, "d")
}
