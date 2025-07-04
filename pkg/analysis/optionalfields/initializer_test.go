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

package optionalfields_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/initializer"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/optionalfields"
)

var _ = Describe("optionalfields initializer", func() {
	Context("config validation", func() {
		type testCase struct {
			config      optionalfields.OptionalFieldsConfig
			expectedErr string
		}

		DescribeTable("should validate the provided config", func(in testCase) {
			ci, ok := optionalfields.Initializer().(initializer.ConfigurableAnalyzerInitializer)
			Expect(ok).To(BeTrue())

			errs := ci.ValidateConfig(&in.config, field.NewPath("optionalfields"))
			if len(in.expectedErr) > 0 {
				Expect(errs.ToAggregate()).To(MatchError(in.expectedErr))
			} else {
				Expect(errs).To(HaveLen(0), "No errors were expected")
			}
		},
			Entry("With a valid OptionalFieldsConfig", testCase{
				config: optionalfields.OptionalFieldsConfig{
					Pointers: optionalfields.OptionalFieldsPointers{
						Preference: "",
						Policy:     "",
					},
					OmitEmpty: optionalfields.OptionalFieldsOmitEmpty{
						Policy: "",
					},
				},
				expectedErr: "",
			}),
			Entry("With a valid OptionalFieldsConfig: Pointer Preference Always", testCase{
				config: optionalfields.OptionalFieldsConfig{
					Pointers: optionalfields.OptionalFieldsPointers{
						Preference: optionalfields.OptionalFieldsPointerPreferenceAlways,
					},
				},
				expectedErr: "",
			}),
			Entry("With a valid OptionalFieldsConfig: Pointer Preference WhenRequired", testCase{
				config: optionalfields.OptionalFieldsConfig{
					Pointers: optionalfields.OptionalFieldsPointers{
						Preference: optionalfields.OptionalFieldsPointerPreferenceWhenRequired,
					},
				},
				expectedErr: "",
			}),
			Entry("With an invalid OptionalFieldsConfig: Pointer Preference", testCase{
				config: optionalfields.OptionalFieldsConfig{
					Pointers: optionalfields.OptionalFieldsPointers{
						Preference: "invalid",
					},
				},
				expectedErr: "optionalfields.pointers.preference: Invalid value: \"invalid\": invalid value, must be one of \"Always\", \"WhenRequired\" or omitted",
			}),
			Entry("With a valid OptionalFieldsConfig: Pointer Policy SuggestFix", testCase{
				config: optionalfields.OptionalFieldsConfig{
					Pointers: optionalfields.OptionalFieldsPointers{
						Policy: optionalfields.OptionalFieldsPointerPolicySuggestFix,
					},
				},
				expectedErr: "",
			}),
			Entry("With a valid OptionalFieldsConfig: Pointer Policy Warn", testCase{
				config: optionalfields.OptionalFieldsConfig{
					Pointers: optionalfields.OptionalFieldsPointers{
						Policy: optionalfields.OptionalFieldsPointerPolicyWarn,
					},
				},
				expectedErr: "",
			}),
			Entry("With an invalid OptionalFieldsConfig: Pointer Policy", testCase{
				config: optionalfields.OptionalFieldsConfig{
					Pointers: optionalfields.OptionalFieldsPointers{
						Policy: "invalid",
					},
				},
				expectedErr: "optionalfields.pointers.policy: Invalid value: \"invalid\": invalid value, must be one of \"SuggestFix\", \"Warn\" or omitted",
			}),
			Entry("With a valid OptionalFieldsConfig: OmitEmpty Policy Ignore", testCase{
				config: optionalfields.OptionalFieldsConfig{
					OmitEmpty: optionalfields.OptionalFieldsOmitEmpty{
						Policy: optionalfields.OptionalFieldsOmitEmptyPolicyIgnore,
					},
				},
				expectedErr: "",
			}),
			Entry("With a valid OptionalFieldsConfig: OmitEmpty Policy Warn", testCase{
				config: optionalfields.OptionalFieldsConfig{
					OmitEmpty: optionalfields.OptionalFieldsOmitEmpty{
						Policy: optionalfields.OptionalFieldsOmitEmptyPolicyWarn,
					},
				},
				expectedErr: "",
			}),
			Entry("With a valid OptionalFieldsConfig: OmitEmpty Policy SuggestFix", testCase{
				config: optionalfields.OptionalFieldsConfig{
					OmitEmpty: optionalfields.OptionalFieldsOmitEmpty{
						Policy: optionalfields.OptionalFieldsOmitEmptyPolicySuggestFix,
					},
				},
				expectedErr: "",
			}),
			Entry("With an invalid OptionalFieldsConfig: OmitEmpty Policy", testCase{
				config: optionalfields.OptionalFieldsConfig{
					OmitEmpty: optionalfields.OptionalFieldsOmitEmpty{
						Policy: "invalid",
					},
				},
				expectedErr: "optionalfields.omitEmpty.policy: Invalid value: \"invalid\": invalid value, must be one of \"Ignore\", \"Warn\", \"SuggestFix\" or omitted",
			}),
		)
	})
})
