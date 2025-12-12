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

package defaults_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/defaults"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/initializer"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils/serialization"
	"sigs.k8s.io/kube-api-linter/pkg/markers"
)

var _ = Describe("defaults initializer", func() {
	Context("config validation", func() {
		type testCase struct {
			config      defaults.DefaultsConfig
			expectedErr string
		}

		DescribeTable("should validate the provided config", func(in testCase) {
			ci, ok := defaults.Initializer().(initializer.ConfigurableAnalyzerInitializer)
			Expect(ok).To(BeTrue())

			errs := ci.ValidateConfig(&in.config, field.NewPath("defaults"))
			if len(in.expectedErr) > 0 {
				Expect(errs.ToAggregate()).To(MatchError(in.expectedErr))
			} else {
				Expect(errs).To(HaveLen(0), "No errors were expected")
			}
		},
			Entry("With a valid DefaultsConfig using default marker", testCase{
				config: defaults.DefaultsConfig{
					PreferredDefaultMarker: markers.DefaultMarker,
				},
				expectedErr: "",
			}),
			Entry("With kubebuilder:default preferred marker", testCase{
				config: defaults.DefaultsConfig{
					PreferredDefaultMarker: markers.KubebuilderDefaultMarker,
				},
				expectedErr: "",
			}),
			Entry("With invalid preferred default marker", testCase{
				config: defaults.DefaultsConfig{
					PreferredDefaultMarker: "invalid",
				},
				expectedErr: "defaults.preferredDefaultMarker: Invalid value: \"invalid\": invalid value, must be one of \"default\", \"kubebuilder:default\" or omitted",
			}),
			Entry("With SuggestFix omitempty policy", testCase{
				config: defaults.DefaultsConfig{
					OmitEmpty: defaults.DefaultsOmitEmpty{
						Policy: serialization.OmitEmptyPolicySuggestFix,
					},
				},
				expectedErr: "",
			}),
			Entry("With Warn omitempty policy", testCase{
				config: defaults.DefaultsConfig{
					OmitEmpty: defaults.DefaultsOmitEmpty{
						Policy: serialization.OmitEmptyPolicyWarn,
					},
				},
				expectedErr: "",
			}),
			Entry("With Ignore omitempty policy", testCase{
				config: defaults.DefaultsConfig{
					OmitEmpty: defaults.DefaultsOmitEmpty{
						Policy: serialization.OmitEmptyPolicyIgnore,
					},
				},
				expectedErr: "",
			}),
			Entry("With invalid omitempty policy", testCase{
				config: defaults.DefaultsConfig{
					OmitEmpty: defaults.DefaultsOmitEmpty{
						Policy: "invalid",
					},
				},
				expectedErr: "defaults.omitempty.policy: Invalid value: \"invalid\": invalid value, must be one of \"Ignore\", \"Warn\", \"SuggestFix\" or omitted",
			}),
			Entry("With SuggestFix omitzero policy", testCase{
				config: defaults.DefaultsConfig{
					OmitZero: defaults.DefaultsOmitZero{
						Policy: serialization.OmitZeroPolicySuggestFix,
					},
				},
				expectedErr: "",
			}),
			Entry("With Warn omitzero policy", testCase{
				config: defaults.DefaultsConfig{
					OmitZero: defaults.DefaultsOmitZero{
						Policy: serialization.OmitZeroPolicyWarn,
					},
				},
				expectedErr: "",
			}),
			Entry("With Forbid omitzero policy", testCase{
				config: defaults.DefaultsConfig{
					OmitZero: defaults.DefaultsOmitZero{
						Policy: serialization.OmitZeroPolicyForbid,
					},
				},
				expectedErr: "",
			}),
			Entry("With invalid omitzero policy", testCase{
				config: defaults.DefaultsConfig{
					OmitZero: defaults.DefaultsOmitZero{
						Policy: "invalid",
					},
				},
				expectedErr: "defaults.omitzero.policy: Invalid value: \"invalid\": invalid value, must be one of \"Forbid\", \"Warn\", \"SuggestFix\" or omitted",
			}),
		)
	})
})
