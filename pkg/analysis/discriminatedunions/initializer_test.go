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

package discriminatedunions_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/discriminatedunions"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/initializer"
)

var _ = Describe("discriminatedunions initializer", func() {
	Context("config validation", func() {
		type testCase struct {
			config      discriminatedunions.Config
			expectedErr string
		}

		DescribeTable("should validate the provided config", func(in testCase) {
			ci, ok := discriminatedunions.Initializer().(initializer.ConfigurableAnalyzerInitializer)
			Expect(ok).To(BeTrue())

			errs := ci.ValidateConfig(&in.config, field.NewPath("discriminatedunions"))

			if len(in.expectedErr) > 0 {
				Expect(errs.ToAggregate()).To(MatchError(in.expectedErr))
			} else {
				Expect(errs).To(HaveLen(0), "No errors were expected")
			}
		},
			Entry("with default config", testCase{
				config:      discriminatedunions.Config{},
				expectedErr: "",
			}),
			Entry("with explicit Forbid", testCase{
				config:      discriminatedunions.Config{NonMemberFields: discriminatedunions.NonMemberFieldsForbid},
				expectedErr: "",
			}),
			Entry("with explicit Allow", testCase{
				config:      discriminatedunions.Config{NonMemberFields: discriminatedunions.NonMemberFieldsAllow},
				expectedErr: "",
			}),
			Entry("with invalid policy", testCase{
				config:      discriminatedunions.Config{NonMemberFields: "invalid"},
				expectedErr: "discriminatedunions.nonMemberFields: Invalid value: \"invalid\": invalid value, must be one of \"Forbid\", \"Allow\" or omitted",
			}),
		)
	})
})
