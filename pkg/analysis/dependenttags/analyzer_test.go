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

package dependenttags_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/dependenttags"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	cfg := dependenttags.Config{
		Rules: []dependenttags.Rule{
			{
				Identifier: "k8s:unionMember",
				Type:       dependenttags.DependencyTypeAll,
				Dependents: []string{"k8s:optional"},
			},
			{
				Identifier: "listType",
				Type:       dependenttags.DependencyTypeAll,
				Dependents: []string{"k8s:listType"},
			},
			{
				Identifier: "example:any",
				Type:       dependenttags.DependencyTypeAny,
				Dependents: []string{"dep1", "dep2"},
			},
		},
	}
	analyzer, err := dependenttags.Initializer().Init(&cfg)

	if err != nil {
		t.Fatalf("failed to initialize analyzer: %v", err)
	}

	analysistest.Run(t, testdata, analyzer, "a")
}
