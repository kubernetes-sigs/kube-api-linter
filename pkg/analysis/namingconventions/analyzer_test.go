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
package namingconventions_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/namingconventions"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()

	cfg := &namingconventions.Config{
		Conventions: []namingconventions.Convention{
			{
				Name:             "nofruit",
				ViolationMatcher: "(?i)fruit",
				Operation:        namingconventions.OperationDrop,
				Message:          "no fields should contain any variations of the word 'fruit' in their name.",
			},
			{
				Name:             "preferbehaviour",
				ViolationMatcher: "(?i)behavior",
				Operation:        namingconventions.OperationReplace,
				Message:          "prefer the use of the word 'behaviour' instead of 'behavior'.",
				Replace:          "Behaviour",
			},
			{
				Name:             "nounsupported",
				ViolationMatcher: "(?i)unsupported",
				Operation:        namingconventions.OperationDropField,
				Message:          "no fields allowing for unsupported behaviors allowed",
			},
			{
				Name:             "notest",
				ViolationMatcher: "(?i)test",
				Operation:        namingconventions.OperationInform,
				Message:          "no temporary test fields",
			},
		},
	}

	analyzer, err := namingconventions.Initializer().Init(cfg)
	if err != nil {
		t.Fatalf("initializing namingconventions linter: %v", err)
	}

	analysistest.RunWithSuggestedFixes(t, testdata, analyzer, "a")
}
