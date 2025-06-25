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

package namingconventions_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/initializer"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/namingconventions"
)

var _ = Describe("namingconventions initializer", func() {
	Context("config validation", func() {
		type testCase struct {
			config      *namingconventions.Config
			expectedErr string
		}

		DescribeTable("should validate the provided config", func(in testCase) {
			ci, ok := namingconventions.Initializer().(initializer.ConfigurableAnalyzerInitializer)
			Expect(ok).To(BeTrue())

			errs := ci.ValidateConfig(in.config, field.NewPath("namingconventions"))
			if len(in.expectedErr) > 0 {
				Expect(errs.ToAggregate()).To(MatchError(in.expectedErr))
			} else {
				Expect(errs).To(HaveLen(0), "No errors were expected")
			}
		},
			Entry("With a valid namingconventions configuration", testCase{
				config: &namingconventions.Config{
					Conventions: []namingconventions.Convention{
						{
							Name:             "nothing",
							ViolationMatcher: "(?i)thing",
							Operation:        namingconventions.OperationDrop,
							Message:          "no fields should have any variations of the word 'thing' in their name",
						},
					},
				},
				expectedErr: "",
			}),
			Entry("With a nil config", testCase{
				config:      nil,
				expectedErr: "namingconventions: Required value: configuration is required for the namingconventions linter when it is enabled",
			}),
			Entry("With an invalid namingconventions configuration, duplicate convention name", testCase{
				config: &namingconventions.Config{
					Conventions: []namingconventions.Convention{
						{
							Name:             "nothing",
							ViolationMatcher: "(?i)thing",
							Operation:        namingconventions.OperationDrop,
							Message:          "no fields should have any variations of the word 'thing' in their name",
						},
						{
							Name:             "nothing",
							ViolationMatcher: "(?i)thing",
							Operation:        namingconventions.OperationDrop,
							Message:          "no fields should have any variations of the word 'thing' in their name",
						},
					},
				},
				expectedErr: "namingconventions.conventions[1].name: Duplicate value: \"nothing\"",
			}),
			Entry("With an invalid namingconventions configuration with empty convention name", testCase{
				config: &namingconventions.Config{
					Conventions: []namingconventions.Convention{
						{
							ViolationMatcher: "(?i)thing",
							Operation:        namingconventions.OperationDrop,
							Message:          "no fields should have any variations of the word 'thing' in their name",
						},
					},
				},
				expectedErr: "namingconventions.conventions[0].name: Required value: name is required",
			}),
			Entry("With an invalid namingconventions configuration with an empty violationMatcher ", testCase{
				config: &namingconventions.Config{
					Conventions: []namingconventions.Convention{
						{
							Name:      "nothing",
							Operation: namingconventions.OperationDrop,
							Message:   "no fields should have any variations of the word 'thing' in their name",
						},
					},
				},
				expectedErr: "namingconventions.conventions[0].violationMatcher: Required value: violationMatcher is required",
			}),
			Entry("With an invalid namingconventions configuration with a violationMatcher that doesn't compile", testCase{
				config: &namingconventions.Config{
					Conventions: []namingconventions.Convention{
						{
							Name:             "nothing",
							ViolationMatcher: "!&*@^#(*!@&^$",
							Operation:        namingconventions.OperationDrop,
							Message:          "no fields should have any variations of the word 'thing' in their name",
						},
					},
				},
				expectedErr: "namingconventions.conventions[0].violationMatcher: Invalid value: \"!&*@^#(*!@&^$\": violationMatcher regular expression failed to compile: error parsing regexp: missing argument to repetition operator: `*`",
			}),
			Entry("With an invalid namingconventions configuration with an empty message", testCase{
				config: &namingconventions.Config{
					Conventions: []namingconventions.Convention{
						{
							Name:             "nothing",
							ViolationMatcher: "(?i)thing",
							Operation:        namingconventions.OperationDrop,
						},
					},
				},
				expectedErr: "namingconventions.conventions[0].message: Required value: message is required",
			}),
			Entry("With an invalid namingconventions configuration with an empty operation", testCase{
				config: &namingconventions.Config{
					Conventions: []namingconventions.Convention{
						{
							Name:             "nothing",
							ViolationMatcher: "(?i)thing",
							Message:          "no fields should have any variations of the word 'thing' in their name",
						},
					},
				},
				expectedErr: "namingconventions.conventions[0].operation: Required value: operation is required",
			}),
			Entry("With an invalid namingconventions configuration with an unknown operation", testCase{
				config: &namingconventions.Config{
					Conventions: []namingconventions.Convention{
						{
							Name:             "nothing",
							ViolationMatcher: "(?i)thing",
							Operation:        "Unknown",
							Message:          "no fields should have any variations of the word 'thing' in their name",
						},
					},
				},
				expectedErr: "namingconventions.conventions[0].operation: Invalid value: \"Unknown\": operation must be one of \"Inform\", \"Drop\", \"DropField\", or \"Replace\"",
			}),
			Entry("With an invalid namingconventions configuration with a replace when operation is not 'Replace'", testCase{
				config: &namingconventions.Config{
					Conventions: []namingconventions.Convention{
						{
							Name:             "nothing",
							ViolationMatcher: "(?i)thing",
							Operation:        namingconventions.OperationDrop,
							Message:          "no fields should have any variations of the word 'thing' in their name",
							Replace:          "item",
						},
					},
				},
				expectedErr: "namingconventions.conventions[0].replace: Invalid value: \"item\": replace must be specified when operation is 'Replace' and is forbidden otherwise",
			}),
			Entry("With an invalid namingconventions configuration with no replace when operation is 'Replace'", testCase{
				config: &namingconventions.Config{
					Conventions: []namingconventions.Convention{
						{
							Name:             "nothing",
							ViolationMatcher: "(?i)thing",
							Operation:        namingconventions.OperationReplace,
							Message:          "no fields should have any variations of the word 'thing' in their name",
						},
					},
				},
				expectedErr: "namingconventions.conventions[0].replace: Invalid value: \"\": replace must be specified when operation is 'Replace' and is forbidden otherwise",
			}),
			Entry("With an invalid namingconventions configuration where replacement string matches violationMatcher", testCase{
				config: &namingconventions.Config{
					Conventions: []namingconventions.Convention{
						{
							Name:             "nothing",
							ViolationMatcher: "(?i)thing",
							Operation:        namingconventions.OperationReplace,
							Message:          "no fields should have any variations of the word 'thing' in their name",
							Replace:          "anotherthing",
						},
					},
				},
				expectedErr: "namingconventions.conventions[0].replace: Invalid value: \"anotherthing\": replace must not match violationMatcher",
			}),
		)
	})
})
