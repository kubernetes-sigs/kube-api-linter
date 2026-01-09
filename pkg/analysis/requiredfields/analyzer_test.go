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
package requiredfields_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/requiredfields"
)

func TestDefaultConfiguration(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := requiredfields.Initializer().Init(&requiredfields.RequiredFieldsConfig{})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.RunWithSuggestedFixes(t, testdata, a, "a")
}

func TestWithOmitEmptyPolicyIgnore(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := requiredfields.Initializer().Init(&requiredfields.RequiredFieldsConfig{
		OmitEmpty: requiredfields.RequiredFieldsOmitEmpty{
			Policy: requiredfields.RequiredFieldsOmitEmptyPolicyIgnore,
		},
		OmitZero: requiredfields.RequiredFieldsOmitZero{
			Policy: requiredfields.RequiredFieldsOmitZeroPolicySuggestFix,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.RunWithSuggestedFixes(t, testdata, a, "b")
}
