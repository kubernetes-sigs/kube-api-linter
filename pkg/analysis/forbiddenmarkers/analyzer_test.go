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
package forbiddenmarkers

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestWithConfiguration(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, newAnalyzer(&Config{
		Markers: []Marker{
			{
				Identifier: "custom:forbidden",
			},
			{
				Identifier: "custom:AttrNoValues",
				RuleSets: []RuleSet{
					{
						Attributes: []MarkerAttribute{
							{
								Name: "fruit",
							},
						},
					},
				},
			},
			{
				Identifier: "custom:AttrValues",
				RuleSets: []RuleSet{
					{
						Attributes: []MarkerAttribute{
							{
								Name: "fruit",
								Values: []string{
									"apple",
									"orange",
									"banana",
								},
							},
						},
					},
				},
			},
			{
				Identifier: "custom:AttrsNoValues",
				RuleSets: []RuleSet{
					{
						Attributes: []MarkerAttribute{
							{
								Name: "fruit",
							},
							{
								Name: "color",
							},
						},
					},
				},
			},
			{
				Identifier: "custom:AttrsValues",
				RuleSets: []RuleSet{
					{
						Attributes: []MarkerAttribute{
							{
								Name: "fruit",
								Values: []string{
									"apple",
									"orange",
									"banana",
								},
							},
							{
								Name: "color",
								Values: []string{
									"red",
									"blue",
									"green",
								},
							},
						},
					},
				},
			},
			{
				// No blue or green apples, but any othe apples allowed
				// No red, blue, or green oranges, but any other oranges allowed
				// No bananas allowed
				Identifier: "custom:MultiRuleSet",
				RuleSets: []RuleSet{
					{
						Attributes: []MarkerAttribute{
							{
								Name: "fruit",
								Values: []string{
									"banana",
								},
							},
						},
					},
					{
						Attributes: []MarkerAttribute{
							{
								Name: "fruit",
								Values: []string{
									"apple",
								},
							},
							{
								Name: "color",
								Values: []string{
									"blue",
									"green",
								},
							},
						},
					},
					{
						Attributes: []MarkerAttribute{
							{
								Name: "fruit",
								Values: []string{
									"orange",
								},
							},
							{
								Name: "color",
								Values: []string{
									"blue",
									"green",
									"red",
								},
							},
						},
					},
				},
			},
		},
	}), "a/...")
}
