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

package requiredfields_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/initializer"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/requiredfields"
)

var _ = Describe("requiredfields initializer", func() {
	Context("config validation", func() {
		type testCase struct {
			config      requiredfields.RequiredFieldsConfig
			expectedErr string
		}

		DescribeTable("should validate the provided config", func(in testCase) {
			ci, ok := requiredfields.Initializer().(initializer.ConfigurableAnalyzerInitializer)
			Expect(ok).To(BeTrue())

			errs := ci.ValidateConfig(&in.config, field.NewPath("requiredfields"))
			if len(in.expectedErr) > 0 {
				Expect(errs.ToAggregate()).To(MatchError(in.expectedErr))
			} else {
				Expect(errs).To(HaveLen(0), "No errors were expected")
			}
		},
			Entry("With a valid RequiredFieldsConfig: omitted", testCase{
				config:      requiredfields.RequiredFieldsConfig{},
				expectedErr: "",
			}),
			Entry("With a valid RequiredFieldsConfig: Pointers: Policy: SuggestFix", testCase{
				config: requiredfields.RequiredFieldsConfig{
					Pointers: requiredfields.RequiredFieldsPointers{
						Policy: requiredfields.RequiredFieldsPointerPolicySuggestFix,
					},
				},
				expectedErr: "",
			}),
			Entry("With a valid RequiredFieldsConfig: Pointers: Policy: Warn", testCase{
				config: requiredfields.RequiredFieldsConfig{
					Pointers: requiredfields.RequiredFieldsPointers{
						Policy: requiredfields.RequiredFieldsPointerPolicyWarn,
					},
				},
				expectedErr: "",
			}),
			Entry("With an invalid RequiredFieldsConfig: Pointers: Policy: invalid", testCase{
				config: requiredfields.RequiredFieldsConfig{
					Pointers: requiredfields.RequiredFieldsPointers{
						Policy: "invalid",
					},
				},
				expectedErr: "requiredfields.pointers.policy: Invalid value: \"invalid\": invalid value, must be one of \"SuggestFix\", \"Warn\" or omitted",
			}),
			Entry("With a valid RequiredFieldsConfig: OmitEmpty: Policy: Ignore", testCase{
				config: requiredfields.RequiredFieldsConfig{
					OmitEmpty: requiredfields.RequiredFieldsOmitEmpty{
						Policy: requiredfields.RequiredFieldsOmitEmptyPolicyIgnore,
					},
				},
			}),
			Entry("With a valid RequiredFieldsConfig: OmitEmpty: Policy: SuggestFix", testCase{
				config: requiredfields.RequiredFieldsConfig{
					OmitEmpty: requiredfields.RequiredFieldsOmitEmpty{
						Policy: requiredfields.RequiredFieldsOmitEmptyPolicySuggestFix,
					},
				},
			}),
			Entry("With a valid RequiredFieldsConfig: OmitEmpty: Policy: Warn", testCase{
				config: requiredfields.RequiredFieldsConfig{
					OmitEmpty: requiredfields.RequiredFieldsOmitEmpty{
						Policy: requiredfields.RequiredFieldsOmitEmptyPolicyWarn,
					},
				},
			}),
			Entry("With an invalid RequiredFieldsConfig: OmitEmpty: Policy: invalid", testCase{
				config: requiredfields.RequiredFieldsConfig{
					OmitEmpty: requiredfields.RequiredFieldsOmitEmpty{
						Policy: "invalid",
					},
				},
				expectedErr: "requiredfields.omitempty.policy: Invalid value: \"invalid\": invalid value, must be one of \"Ignore\", \"Warn\", \"SuggestFix\" or omitted",
			}),
			Entry("With a valid RequiredFieldsConfig: OmitZero: Policy: SuggestFix", testCase{
				config: requiredfields.RequiredFieldsConfig{
					OmitZero: requiredfields.RequiredFieldsOmitZero{
						Policy: requiredfields.RequiredFieldsOmitZeroPolicySuggestFix,
					},
				},
			}),
			Entry("With a valid RequiredFieldsConfig: OmitZero: Policy: Warn", testCase{
				config: requiredfields.RequiredFieldsConfig{
					OmitZero: requiredfields.RequiredFieldsOmitZero{
						Policy: requiredfields.RequiredFieldsOmitZeroPolicyWarn,
					},
				},
			}),
			Entry("With a valid RequiredFieldsConfig: OmitZero: Policy: Forbid", testCase{
				config: requiredfields.RequiredFieldsConfig{
					OmitZero: requiredfields.RequiredFieldsOmitZero{
						Policy: requiredfields.RequiredFieldsOmitZeroPolicyForbid,
					},
				},
			}),
			Entry("With an invalid RequiredFieldsConfig: OmitZero: Policy: invalid", testCase{
				config: requiredfields.RequiredFieldsConfig{
					OmitZero: requiredfields.RequiredFieldsOmitZero{
						Policy: "invalid",
					},
				},
				expectedErr: "requiredfields.omitzero.policy: Invalid value: \"invalid\": invalid value, must be one of \"SuggestFix\", \"Warn\", \"Forbid\" or omitted",
			}),
		)
	})
})
