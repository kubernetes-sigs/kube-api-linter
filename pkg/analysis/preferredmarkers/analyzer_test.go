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
package preferredmarkers

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestWithConfiguration(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, newAnalyzer(&Config{
		Markers: []Marker{
			{
				PreferredIdentifier: "k8s:optional",
				EquivalentIdentifiers: []EquivalentIdentifier{
					{Identifier: "kubebuilder:validation:Optional"},
					{Identifier: "custom:optional"},
				},
			},
			{
				PreferredIdentifier: "k8s:required",
				EquivalentIdentifiers: []EquivalentIdentifier{
					{Identifier: "kubebuilder:validation:Required"},
				},
			},
			{
				PreferredIdentifier: "custom:preferred",
				EquivalentIdentifiers: []EquivalentIdentifier{
					{Identifier: "custom:old"},
					{Identifier: "custom:deprecated"},
				},
			},
		},
	}), "a/...")
}
