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
package forbiddenmarkers_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/forbiddenmarkers"
	"sigs.k8s.io/kube-api-linter/pkg/config"
)

func TestWithConfiguration(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, forbiddenmarkers.NewAnalyzer(config.ForbiddenMarkersConfig{
		Markers: []config.ForbiddenMarker{
			{
				Identifier: "custom:forbidden",
			},
			{
				Identifier: "custom:AttrNoValues",
				Attributes: []config.ForbiddenMarkerAttribute{
					{
						Attribute: "fruit",
					},
				},
			},
			{
				Identifier: "custom:AttrValues",
				Attributes: []config.ForbiddenMarkerAttribute{
					{
						Attribute: "fruit",
						Values: []string{
							"apple",
							"orange",
							"banana",
						},
					},
				},
			},
			{
				Identifier: "custom:AttrsNoValues",
				Attributes: []config.ForbiddenMarkerAttribute{
					{
						Attribute: "fruit",
					},
					{
						Attribute: "color",
					},
				},
			},
			{
				Identifier: "custom:AttrsValues",
				Attributes: []config.ForbiddenMarkerAttribute{
					{
						Attribute: "fruit",
						Values: []string{
							"apple",
							"orange",
							"banana",
						},
					},
					{
						Attribute: "color",
						Values: []string{
							"red",
							"blue",
							"green",
						},
					},
				},
			},
		},
	}), "a/...")
}
