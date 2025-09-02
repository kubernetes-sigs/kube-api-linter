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
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/forbiddenmarkers"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/initializer"
)

var _ = Describe("forbiddenmarkers initializer", func() {
	Context("config validation", func() {
		type testCase struct {
			config      *forbiddenmarkers.Config
			expectedErr string
		}

		DescribeTable("should validate the provided config", func(in testCase) {
			ci, ok := forbiddenmarkers.Initializer().(initializer.ConfigurableAnalyzerInitializer)
			Expect(ok).To(BeTrue())

			errs := ci.ValidateConfig(in.config, field.NewPath("forbiddenmarkers"))
			if len(in.expectedErr) > 0 {
				Expect(errs.ToAggregate()).To(MatchError(in.expectedErr))
			} else {
				Expect(errs).To(HaveLen(0), "No errors were expected")
			}
		},
			Entry("With a valid forbiddenmarkers configuration", testCase{
				config: &forbiddenmarkers.Config{
					Markers: []forbiddenmarkers.Marker{
						{
							Identifier: "custom:forbidden",
						},
					},
				},
				expectedErr: "",
			}),
			Entry("With an invalid forbiddenmarkers configuration, duplicate markers", testCase{
				config: &forbiddenmarkers.Config{
					Markers: []forbiddenmarkers.Marker{
						{
							Identifier: "custom:forbidden",
						},
						{
							Identifier: "custom:forbidden",
						},
					},
				},
				expectedErr: "forbiddenmarkers.markers[1].identifier: Duplicate value: \"custom:forbidden\"",
			}),
			Entry("With a valid forbiddenmarkers configuration with ruleSets", testCase{
				config: &forbiddenmarkers.Config{
					Markers: []forbiddenmarkers.Marker{
						{
							Identifier: "custom:forbidden",
							RuleSets: []forbiddenmarkers.RuleSet{
								{
									Attributes: []forbiddenmarkers.MarkerAttribute{
										{
											Name: "fruit",
										},
									},
								},
							},
						},
					},
				},
				expectedErr: "",
			}),
			Entry("With an invalid forbiddenmarkers configuration with duplicate attributes in a single ruleSet", testCase{
				config: &forbiddenmarkers.Config{
					Markers: []forbiddenmarkers.Marker{
						{
							Identifier: "custom:forbidden",
							RuleSets: []forbiddenmarkers.RuleSet{
								{
									Attributes: []forbiddenmarkers.MarkerAttribute{
										{
											Name: "fruit",
										},
										{
											Name: "fruit",
										},
									},
								},
							},
						},
					},
				},
				expectedErr: "forbiddenmarkers.markers[0].ruleSets[0].attributes[1].name: Duplicate value: \"fruit\"",
			}),
			Entry("With a valid forbiddenmarkers configuration with a ruleSet with attributes with values", testCase{
				config: &forbiddenmarkers.Config{
					Markers: []forbiddenmarkers.Marker{
						{
							Identifier: "custom:forbidden",
							RuleSets: []forbiddenmarkers.RuleSet{
								{
									Attributes: []forbiddenmarkers.MarkerAttribute{
										{
											Name: "fruit",
											Values: []string{
												"apple",
												"banana",
											},
										},
									},
								},
							},
						},
					},
				},
				expectedErr: "",
			}),
			Entry("With an invalid forbiddenmarkers configuration with a ruleSet with attributes with duplicate values", testCase{
				config: &forbiddenmarkers.Config{
					Markers: []forbiddenmarkers.Marker{
						{
							Identifier: "custom:forbidden",
							RuleSets: []forbiddenmarkers.RuleSet{
								{
									Attributes: []forbiddenmarkers.MarkerAttribute{
										{
											Name: "fruit",
											Values: []string{
												"apple",
												"apple",
											},
										},
									},
								},
							},
						},
					},
				},
				expectedErr: "forbiddenmarkers.markers[0].ruleSets[0].attributes[0].values[1]: Duplicate value: \"apple\"",
			}),
			Entry("With a nil config", testCase{
				config:      nil,
				expectedErr: "forbiddenmarkers: Required value: configuration is required for the forbiddenmarkers linter when it is enabled",
			}),
			Entry("With an invalid forbiddenmarkers configuration with empty ruleSet specified", testCase{
				config: &forbiddenmarkers.Config{
					Markers: []forbiddenmarkers.Marker{
						{
							Identifier: "custom:forbidden",
							RuleSets: []forbiddenmarkers.RuleSet{
								{},
							},
						},
					},
				},
				expectedErr: "forbiddenmarkers.markers[0].ruleSets[0].attributes: Required value: must contain at least one attribute",
			}),
		)
	})
})
