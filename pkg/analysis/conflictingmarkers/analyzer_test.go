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
package conflictingmarkers_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/conflictingmarkers"
)

func TestDefaultConfiguration(t *testing.T) {
	testdata := analysistest.TestData()

	initializer := conflictingmarkers.Initializer()

	analyzer, err := initializer.Init(&conflictingmarkers.ConflictingMarkersConfig{})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.Run(t, testdata, analyzer, "a")
}

func TestCustomConfiguration(t *testing.T) {
	testdata := analysistest.TestData()

	config := &conflictingmarkers.ConflictingMarkersConfig{
		CustomConflicts: []conflictingmarkers.ConflictSet{
			{
				Name:        "custom_conflict",
				SetA:        []string{"custom:marker1", "custom:marker2"},
				SetB:        []string{"custom:marker3", "custom:marker4"},
				Description: "Custom markers conflict with each other",
			},
		},
	}

	initializer := conflictingmarkers.Initializer()

	analyzer, err := initializer.Init(config)
	if err != nil {
		t.Fatal(err)
	}

	analysistest.Run(t, testdata, analyzer, "b")
}
