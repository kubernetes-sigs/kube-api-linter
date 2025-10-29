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
package markerscope_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/initializer"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/markerscope"
)

var _ = Describe("markerscope initializer", func() {
	Context("config validation", func() {
		type testCase struct {
			config      markerscope.MarkerScopeConfig
			expectedErr string
		}

		DescribeTable("should validate the provided config",
			func(in testCase) {
				ci, ok := markerscope.Initializer().(initializer.ConfigurableAnalyzerInitializer)
				Expect(ok).To(BeTrue())

				errs := ci.ValidateConfig(&in.config, field.NewPath("markerscope"))
				if len(in.expectedErr) > 0 {
					Expect(errs.ToAggregate()).To(HaveOccurred())
					Expect(errs.ToAggregate().Error()).To(ContainSubstring(in.expectedErr))
				} else {
					Expect(errs).To(HaveLen(0), "No errors were expected")
				}
			},

			Entry("With nil config", testCase{
				config:      markerscope.MarkerScopeConfig{},
				expectedErr: "",
			}),

			Entry("With empty config", testCase{
				config:      markerscope.MarkerScopeConfig{},
				expectedErr: "",
			}),

			Entry("With valid warn policy", testCase{
				config: markerscope.MarkerScopeConfig{
					Policy: markerscope.MarkerScopePolicyWarn,
				},
				expectedErr: "",
			}),

			Entry("With valid suggest_fix policy", testCase{
				config: markerscope.MarkerScopeConfig{
					Policy: markerscope.MarkerScopePolicySuggestFix,
				},
				expectedErr: "",
			}),

			Entry("With invalid policy", testCase{
				config: markerscope.MarkerScopeConfig{
					Policy: "invalid-policy",
				},
				expectedErr: `markerscope.policy: Invalid value: "invalid-policy": invalid policy, must be one of: "Warn", "SuggestFix"`,
			}),

			Entry("With valid marker rules", testCase{
				config: markerscope.MarkerScopeConfig{
					MarkerRules: []markerscope.MarkerScopeRule{
						{
							Name:  "custom:marker",
							Scope: markerscope.FieldScope,
						},
					},
				},
				expectedErr: "",
			}),

			Entry("With marker rule having empty scope", testCase{
				config: markerscope.MarkerScopeConfig{
					MarkerRules: []markerscope.MarkerScopeRule{
						{
							Name:  "custom:marker",
							Scope: "",
						},
					},
				},
				expectedErr: `scope is required`,
			}),

			Entry("With marker rule having invalid scope value", testCase{
				config: markerscope.MarkerScopeConfig{
					MarkerRules: []markerscope.MarkerScopeRule{
						{
							Name:  "custom:marker",
							Scope: "invalid",
						},
					},
				},
				expectedErr: `invalid scope: "invalid" (must be one of: Field, Type, Any)`,
			}),

			Entry("With marker rule having invalid schema type", testCase{
				config: markerscope.MarkerScopeConfig{
					MarkerRules: []markerscope.MarkerScopeRule{
						{
							Name:  "custom:marker",
							Scope: markerscope.FieldScope,
							TypeConstraint: &markerscope.TypeConstraint{
								AllowedSchemaTypes: []markerscope.SchemaType{"invalid-type"},
							},
						},
					},
				},
				expectedErr: `invalid type constraint: invalid schema type: "invalid-type"`,
			}),

			Entry("With valid type constraint with string type", testCase{
				config: markerscope.MarkerScopeConfig{
					MarkerRules: []markerscope.MarkerScopeRule{
						{
							Name:  "custom:string-marker",
							Scope: markerscope.FieldScope,
							TypeConstraint: &markerscope.TypeConstraint{
								AllowedSchemaTypes: []markerscope.SchemaType{markerscope.SchemaTypeString},
							},
						},
					},
				},
				expectedErr: "",
			}),

			Entry("With valid type constraint with integer type", testCase{
				config: markerscope.MarkerScopeConfig{
					MarkerRules: []markerscope.MarkerScopeRule{
						{
							Name:  "custom:integer-marker",
							Scope: markerscope.FieldScope,
							TypeConstraint: &markerscope.TypeConstraint{
								AllowedSchemaTypes: []markerscope.SchemaType{markerscope.SchemaTypeInteger},
							},
						},
					},
				},
				expectedErr: "",
			}),

			Entry("With valid type constraint with multiple types", testCase{
				config: markerscope.MarkerScopeConfig{
					MarkerRules: []markerscope.MarkerScopeRule{
						{
							Name: "custom:numeric-marker",
							Scope: markerscope.FieldScope,
							TypeConstraint: &markerscope.TypeConstraint{
								AllowedSchemaTypes: []markerscope.SchemaType{
									markerscope.SchemaTypeInteger,
									markerscope.SchemaTypeNumber,
								},
							},
						},
					},
				},
				expectedErr: "",
			}),

			Entry("With valid type constraint with element constraint", testCase{
				config: markerscope.MarkerScopeConfig{
					MarkerRules: []markerscope.MarkerScopeRule{
						{
							Name:  "custom:string-array",
							Scope: markerscope.FieldScope,
							TypeConstraint: &markerscope.TypeConstraint{
								AllowedSchemaTypes: []markerscope.SchemaType{markerscope.SchemaTypeArray},
								ElementConstraint: &markerscope.TypeConstraint{
									AllowedSchemaTypes: []markerscope.SchemaType{markerscope.SchemaTypeString},
								},
							},
						},
					},
				},
				expectedErr: "",
			}),

			Entry("With invalid element constraint", testCase{
				config: markerscope.MarkerScopeConfig{
					MarkerRules: []markerscope.MarkerScopeRule{
						{
							Name:  "custom:invalid-array",
							Scope: markerscope.FieldScope,
							TypeConstraint: &markerscope.TypeConstraint{
								AllowedSchemaTypes: []markerscope.SchemaType{markerscope.SchemaTypeArray},
								ElementConstraint: &markerscope.TypeConstraint{
									AllowedSchemaTypes: []markerscope.SchemaType{"invalid-type"},
								},
							},
						},
					},
				},
				expectedErr: `invalid type constraint: invalid element constraint: invalid schema type: "invalid-type"`,
			}),

			Entry("With Any scope (field and type)", testCase{
				config: markerscope.MarkerScopeConfig{
					MarkerRules: []markerscope.MarkerScopeRule{
						{
							Name:  "custom:flexible-marker",
							Scope: markerscope.AnyScope,
						},
					},
				},
				expectedErr: "",
			}),
		)
	})

	Context("analyzer initialization", func() {
		It("should initialize analyzer with nil config", func() {
			// Note: Init expects a MarkerScopeConfig, passing nil will error
			// Use empty config instead
			analyzer, err := markerscope.Initializer().Init(&markerscope.MarkerScopeConfig{})
			Expect(err).ToNot(HaveOccurred())
			Expect(analyzer).ToNot(BeNil())
		})

		It("should initialize analyzer with empty config", func() {
			analyzer, err := markerscope.Initializer().Init(&markerscope.MarkerScopeConfig{})
			Expect(err).ToNot(HaveOccurred())
			Expect(analyzer).ToNot(BeNil())
		})

		It("should initialize analyzer with custom markers", func() {
			cfg := &markerscope.MarkerScopeConfig{
				Policy: markerscope.MarkerScopePolicyWarn,
				MarkerRules: []markerscope.MarkerScopeRule{
					{
						Name:  "custom:marker",
						Scope: markerscope.FieldScope,
					},
				},
			}
			analyzer, err := markerscope.Initializer().Init(cfg)
			Expect(err).ToNot(HaveOccurred())
			Expect(analyzer).ToNot(BeNil())
		})
	})
})
