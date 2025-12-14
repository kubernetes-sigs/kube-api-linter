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
					Expect(errs.ToAggregate().Error()).To(Equal(in.expectedErr))
				} else {
					Expect(errs).To(HaveLen(0), "No errors were expected")
				}
			},

			Entry("With nil config", testCase{
				config:      markerscope.MarkerScopeConfig{},
				expectedErr: "",
			}),

			Entry("With valid Warn policy", testCase{
				config: markerscope.MarkerScopeConfig{
					Policy: markerscope.MarkerScopePolicyWarn,
				},
				expectedErr: "",
			}),

			Entry("With valid SuggestFix policy", testCase{
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
					CustomMarkers: []markerscope.MarkerScopeRule{
						{
							Identifier: "custom:marker",
							Scopes:     []markerscope.ScopeConstraint{markerscope.FieldScope},
						},
					},
				},
				expectedErr: "",
			}),

			Entry("With marker rule having empty scope", testCase{
				config: markerscope.MarkerScopeConfig{
					CustomMarkers: []markerscope.MarkerScopeRule{
						{
							Identifier: "custom:marker",
							Scopes:     []markerscope.ScopeConstraint{},
						},
					},
				},
				expectedErr: `markerscope.customMarkers[0].scopes: Required value: scope is required`,
			}),

			Entry("With marker rule having invalid scope value", testCase{
				config: markerscope.MarkerScopeConfig{
					CustomMarkers: []markerscope.MarkerScopeRule{
						{
							Identifier: "custom:marker",
							Scopes:     []markerscope.ScopeConstraint{"invalid"},
						},
					},
				},
				expectedErr: `markerscope.customMarkers[0].scopes[0]: Invalid value: "invalid": invalid scope: "invalid" (must be one of: Field, Type)`,
			}),

			Entry("With marker rule having invalid schema type", testCase{
				config: markerscope.MarkerScopeConfig{
					CustomMarkers: []markerscope.MarkerScopeRule{
						{
							Identifier: "custom:marker",
							Scopes:     []markerscope.ScopeConstraint{markerscope.FieldScope},
							TypeConstraint: &markerscope.TypeConstraint{
								AllowedSchemaTypes: []markerscope.SchemaType{"invalid-type"},
							},
						},
					},
				},
				expectedErr: `markerscope.customMarkers[0].typeConstraint.allowedSchemaTypes[0]: Invalid value: "invalid-type": invalid schema type: "invalid-type"`,
			}),

			Entry("With valid type constraint with string type", testCase{
				config: markerscope.MarkerScopeConfig{
					CustomMarkers: []markerscope.MarkerScopeRule{
						{
							Identifier: "custom:string-marker",
							Scopes:     []markerscope.ScopeConstraint{markerscope.FieldScope},
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
					CustomMarkers: []markerscope.MarkerScopeRule{
						{
							Identifier: "custom:integer-marker",
							Scopes:     []markerscope.ScopeConstraint{markerscope.FieldScope},
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
					CustomMarkers: []markerscope.MarkerScopeRule{
						{
							Identifier: "custom:numeric-string-marker",
							Scopes:     []markerscope.ScopeConstraint{markerscope.FieldScope},
							TypeConstraint: &markerscope.TypeConstraint{
								AllowedSchemaTypes: []markerscope.SchemaType{
									markerscope.SchemaTypeInteger,
									markerscope.SchemaTypeString,
								},
							},
						},
					},
				},
				expectedErr: "",
			}),

			Entry("With valid type constraint with element constraint", testCase{
				config: markerscope.MarkerScopeConfig{
					CustomMarkers: []markerscope.MarkerScopeRule{
						{
							Identifier: "custom:string-array",
							Scopes:     []markerscope.ScopeConstraint{markerscope.FieldScope},
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
					CustomMarkers: []markerscope.MarkerScopeRule{
						{
							Identifier: "custom:invalid-array",
							Scopes:     []markerscope.ScopeConstraint{markerscope.FieldScope},
							TypeConstraint: &markerscope.TypeConstraint{
								AllowedSchemaTypes: []markerscope.SchemaType{markerscope.SchemaTypeArray},
								ElementConstraint: &markerscope.TypeConstraint{
									AllowedSchemaTypes: []markerscope.SchemaType{"invalid-type"},
								},
							},
						},
					},
				},
				expectedErr: `markerscope.customMarkers[0].typeConstraint.elementConstraint.allowedSchemaTypes[0]: Invalid value: "invalid-type": invalid schema type: "invalid-type"`,
			}),

			Entry("With both Field and Type scopes", testCase{
				config: markerscope.MarkerScopeConfig{
					CustomMarkers: []markerscope.MarkerScopeRule{
						{
							Identifier: "custom:flexible-marker",
							Scopes:     []markerscope.ScopeConstraint{markerscope.FieldScope, markerscope.TypeScope},
						},
					},
				},
				expectedErr: "",
			}),

			Entry("With override marker for built-in marker", testCase{
				config: markerscope.MarkerScopeConfig{
					OverrideMarkers: []markerscope.MarkerScopeRule{
						{
							Identifier: "optional",
							Scopes:     []markerscope.ScopeConstraint{markerscope.FieldScope, markerscope.TypeScope}, // Override default [FieldScope]
						},
					},
				},
				expectedErr: "",
			}),

			Entry("With override marker for non-built-in marker", testCase{
				config: markerscope.MarkerScopeConfig{
					OverrideMarkers: []markerscope.MarkerScopeRule{
						{
							Identifier: "custom:nonexistent",
							Scopes:     []markerscope.ScopeConstraint{markerscope.FieldScope},
						},
					},
				},
				expectedErr: `markerscope.overrideMarkers[0].identifier: Invalid value: "custom:nonexistent": override marker must be a built-in marker; use customMarkers for custom markers`,
			}),

			Entry("With custom marker for built-in marker", testCase{
				config: markerscope.MarkerScopeConfig{
					CustomMarkers: []markerscope.MarkerScopeRule{
						{
							Identifier: "optional", // Built-in marker
							Scopes:     []markerscope.ScopeConstraint{markerscope.FieldScope, markerscope.TypeScope},
						},
					},
				},
				expectedErr: `markerscope.customMarkers[0].identifier: Invalid value: "optional": custom marker cannot be a built-in marker; use overrideMarkers to override built-in markers`,
			}),

			Entry("With both override and custom markers", testCase{
				config: markerscope.MarkerScopeConfig{
					OverrideMarkers: []markerscope.MarkerScopeRule{
						{
							Identifier: "optional",
							Scopes:     []markerscope.ScopeConstraint{markerscope.FieldScope, markerscope.TypeScope},
						},
					},
					CustomMarkers: []markerscope.MarkerScopeRule{
						{
							Identifier: "custom:marker",
							Scopes:     []markerscope.ScopeConstraint{markerscope.FieldScope},
						},
					},
				},
				expectedErr: "",
			}),
		)
	})

	Context("analyzer initialization", func() {
		It("should initialize analyzer with nil config", func() {
			analyzer, err := markerscope.Initializer().Init(&markerscope.MarkerScopeConfig{})
			Expect(err).ToNot(HaveOccurred())
			Expect(analyzer).ToNot(BeNil())
		})

		It("should initialize analyzer with custom markers", func() {
			cfg := &markerscope.MarkerScopeConfig{
				Policy: markerscope.MarkerScopePolicyWarn,
				CustomMarkers: []markerscope.MarkerScopeRule{
					{
						Identifier: "custom:marker",
						Scopes:     []markerscope.ScopeConstraint{markerscope.FieldScope},
					},
				},
			}
			analyzer, err := markerscope.Initializer().Init(cfg)
			Expect(err).ToNot(HaveOccurred())
			Expect(analyzer).ToNot(BeNil())
		})

		It("should initialize analyzer with override markers", func() {
			cfg := &markerscope.MarkerScopeConfig{
				Policy: markerscope.MarkerScopePolicyWarn,
				OverrideMarkers: []markerscope.MarkerScopeRule{
					{
						Identifier: "optional",
						Scopes:     []markerscope.ScopeConstraint{markerscope.FieldScope, markerscope.TypeScope}, // Override default [FieldScope]
					},
				},
			}
			analyzer, err := markerscope.Initializer().Init(cfg)
			Expect(err).ToNot(HaveOccurred())
			Expect(analyzer).ToNot(BeNil())
		})

		It("should initialize analyzer with both override and custom markers", func() {
			cfg := &markerscope.MarkerScopeConfig{
				Policy: markerscope.MarkerScopePolicySuggestFix,
				OverrideMarkers: []markerscope.MarkerScopeRule{
					{
						Identifier: "optional",
						Scopes:     []markerscope.ScopeConstraint{markerscope.FieldScope, markerscope.TypeScope},
					},
				},
				CustomMarkers: []markerscope.MarkerScopeRule{
					{
						Identifier: "custom:validation:MyMarker",
						Scopes:     []markerscope.ScopeConstraint{markerscope.FieldScope},
						TypeConstraint: &markerscope.TypeConstraint{
							AllowedSchemaTypes: []markerscope.SchemaType{markerscope.SchemaTypeString},
						},
					},
				},
			}
			analyzer, err := markerscope.Initializer().Init(cfg)
			Expect(err).ToNot(HaveOccurred())
			Expect(analyzer).ToNot(BeNil())
		})
	})
})
