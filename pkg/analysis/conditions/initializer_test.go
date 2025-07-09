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

package conditions_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/conditions"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/initializer"
)

var _ = Describe("conditions initializer", func() {
	Context("config validation", func() {
		type testCase struct {
			config      conditions.ConditionsConfig
			expectedErr string
		}

		DescribeTable("should validate the provided config", func(in testCase) {
			ci, ok := conditions.Initializer().(initializer.ConfigurableAnalyzerInitializer)
			Expect(ok).To(BeTrue())

			errs := ci.ValidateConfig(&in.config, field.NewPath("conditions"))
			if len(in.expectedErr) > 0 {
				Expect(errs.ToAggregate()).To(MatchError(in.expectedErr))
			} else {
				Expect(errs).To(HaveLen(0), "No errors were expected")
			}
		},
			Entry("With a valid ConditionsConfig", testCase{
				config: conditions.ConditionsConfig{
					IsFirstField: "",
					UseProtobuf:  "",
				},
				expectedErr: "",
			}),
			Entry("With a valid ConditionsConfig IsFirstField: Warn", testCase{
				config: conditions.ConditionsConfig{
					IsFirstField: conditions.ConditionsFirstFieldWarn,
				},
				expectedErr: "",
			}),
			Entry("With a valid ConditionsConfig IsFirstField: Ignore", testCase{
				config: conditions.ConditionsConfig{
					IsFirstField: conditions.ConditionsFirstFieldIgnore,
				},
				expectedErr: "",
			}),
			Entry("With an invalid ConditionsConfig IsFirstField", testCase{
				config: conditions.ConditionsConfig{
					IsFirstField: "invalid",
				},
				expectedErr: "conditions.isFirstField: Invalid value: \"invalid\": invalid value, must be one of \"Warn\", \"Ignore\" or omitted",
			}),
			Entry("With a valid ConditionsConfig UseProtobuf: SuggestFix", testCase{
				config: conditions.ConditionsConfig{
					UseProtobuf: conditions.ConditionsUseProtobufSuggestFix,
				},
				expectedErr: "",
			}),
			Entry("With a valid ConditionsConfig UseProtobuf: Warn", testCase{
				config: conditions.ConditionsConfig{
					UseProtobuf: conditions.ConditionsUseProtobufWarn,
				},
				expectedErr: "",
			}),
			Entry("With a valid ConditionsConfig UseProtobuf: Ignore", testCase{
				config: conditions.ConditionsConfig{
					UseProtobuf: conditions.ConditionsUseProtobufIgnore,
				},
				expectedErr: "",
			}),
			Entry("With a valid ConditionsConfig UseProtobuf: Forbid", testCase{
				config: conditions.ConditionsConfig{
					UseProtobuf: conditions.ConditionsUseProtobufForbid,
				},
				expectedErr: "",
			}),
			Entry("With an invalid ConditionsConfig UseProtobuf", testCase{
				config: conditions.ConditionsConfig{
					UseProtobuf: "invalid",
				},
				expectedErr: "conditions.useProtobuf: Invalid value: \"invalid\": invalid value, must be one of \"SuggestFix\", \"Warn\", \"Ignore\", \"Forbid\" or omitted",
			}),
			Entry("With a valid ConditionsConfig UsePatchStrategy: SuggestFix", testCase{
				config: conditions.ConditionsConfig{
					UsePatchStrategy: conditions.ConditionsUsePatchStrategySuggestFix,
				},
				expectedErr: "",
			}),
			Entry("With a valid ConditionsConfig UsePatchStrategy: Warn", testCase{
				config: conditions.ConditionsConfig{
					UsePatchStrategy: conditions.ConditionsUsePatchStrategyWarn,
				},
				expectedErr: "",
			}),
			Entry("With a valid ConditionsConfig UsePatchStrategy: Ignore", testCase{
				config: conditions.ConditionsConfig{
					UsePatchStrategy: conditions.ConditionsUsePatchStrategyIgnore,
				},
				expectedErr: "",
			}),
			Entry("With a valid ConditionsConfig UsePatchStrategy: Forbid", testCase{
				config: conditions.ConditionsConfig{
					UsePatchStrategy: conditions.ConditionsUsePatchStrategyForbid,
				},
				expectedErr: "",
			}),
			Entry("With an invalid ConditionsConfig UsePatchStrategy", testCase{
				config: conditions.ConditionsConfig{
					UsePatchStrategy: "invalid",
				},
				expectedErr: "conditions.usePatchStrategy: Invalid value: \"invalid\": invalid value, must be one of \"SuggestFix\", \"Warn\", \"Ignore\", \"Forbid\" or omitted",
			}),
		)
	})
})
