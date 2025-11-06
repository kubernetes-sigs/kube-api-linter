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
package uniquemarkers_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/uniquemarkers"
)

func TestWithDefaults(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := uniquemarkers.Initializer().Init(&uniquemarkers.UniqueMarkersConfig{})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.Run(t, testdata, a, "a/...")
}

func TestWithConfiguration(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := uniquemarkers.Initializer().Init(&uniquemarkers.UniqueMarkersConfig{
		CustomMarkers: []uniquemarkers.UniqueMarker{
			{
				Identifier: "custom:SomeCustomMarker",
			},
			{
				Identifier: "custom:OtherMarker",
				Attributes: []string{
					"attribute",
				},
			},
			{
				Identifier: "custom:MultiMarker",
				Attributes: []string{
					"fruit",
					"color",
					"country",
				},
			},
			{
				Identifier: "k8s:maxLength",
			},
			{
				Identifier: "k8s:uniqueMarkerArguments",
				Attributes: []string{
					"fruit",
				},
			},
			{
				Identifier: "k8s:uniqueMarkerUnnamedArguments",
				Attributes: []string{
					"",
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.Run(t, testdata, a, "b/...")
}
