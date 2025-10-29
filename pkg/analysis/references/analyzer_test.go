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
package references_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/references"
)

func TestAllowRefAndRefs(t *testing.T) {
	testdata := analysistest.TestData()

	cfg := &references.Config{
		Policy: references.PolicyAllowRefAndRefs,
	}

	analyzer, err := references.Initializer().Init(cfg)
	if err != nil {
		t.Fatalf("initializing references linter: %v", err)
	}

	analysistest.RunWithSuggestedFixes(t, testdata, analyzer, "a")
}

func TestEmptyConfig(t *testing.T) {
	testdata := analysistest.TestData()

	// Test with empty config - should default to ForbidRefAndRefs behavior
	cfg := &references.Config{}

	analyzer, err := references.Initializer().Init(cfg)
	if err != nil {
		t.Fatalf("initializing references linter: %v", err)
	}

	// With default config (empty Policy), it should default to ForbidRefAndRefs behavior
	// So we test with folder 'b' which has the same expectations
	analysistest.RunWithSuggestedFixes(t, testdata, analyzer, "b")
}

func TestForbidRefAndRefs(t *testing.T) {
	testdata := analysistest.TestData()

	cfg := &references.Config{
		Policy: references.PolicyForbidRefAndRefs,
	}

	analyzer, err := references.Initializer().Init(cfg)
	if err != nil {
		t.Fatalf("initializing references linter: %v", err)
	}

	analysistest.RunWithSuggestedFixes(t, testdata, analyzer, "b")
}
