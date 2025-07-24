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

func TestConflictingMarkersAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()

	config := &conflictingmarkers.ConflictingMarkersConfig{
		Conflicts: []conflictingmarkers.ConflictSet{
			{
				Name:        "test_conflict",
				Sets:        [][]string{{"marker1", "marker2"}, {"marker3", "marker4"}},
				Description: "Test markers conflict with each other",
			},
			{
				Name:        "three_way_conflict",
				Sets:        [][]string{{"marker5", "marker6"}, {"marker7", "marker8"}, {"marker9", "marker10"}},
				Description: "Three-way conflict between marker sets",
			},
		},
	}

	initializer := conflictingmarkers.Initializer()

	analyzer, err := initializer.Init(config)
	if err != nil {
		t.Fatal(err)
	}

	analysistest.Run(t, testdata, analyzer, "a")
}
