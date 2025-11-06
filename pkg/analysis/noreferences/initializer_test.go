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
package noreferences_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/initializer"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/noreferences"
)

var _ = Describe("noreferences initializer", func() {
	Context("config validation", func() {
		type testCase struct {
			config      *noreferences.Config
			expectedErr string
		}

		DescribeTable("should validate the provided config", func(in testCase) {
			ci, ok := noreferences.Initializer().(initializer.ConfigurableAnalyzerInitializer)
			Expect(ok).To(BeTrue())

			errs := ci.ValidateConfig(in.config, field.NewPath("noreferences"))
			if len(in.expectedErr) > 0 {
				Expect(errs.ToAggregate()).To(MatchError(in.expectedErr))
			} else {
				Expect(errs).To(HaveLen(0), "No errors were expected")
			}
		},
			Entry("With a valid noreferences configuration with policy=NoReferences", testCase{
				config: &noreferences.Config{
					Policy: noreferences.PolicyNoReferences,
				},
				expectedErr: "",
			}),
			Entry("With a valid noreferences configuration with policy=PreferAbbreviatedReference", testCase{
				config: &noreferences.Config{
					Policy: noreferences.PolicyPreferAbbreviatedReference,
				},
				expectedErr: "",
			}),
			Entry("With a nil config", testCase{
				config:      nil,
				expectedErr: "",
			}),
			Entry("With an empty config", testCase{
				config:      &noreferences.Config{},
				expectedErr: "",
			}),
			Entry("With an invalid policy value", testCase{
				config: &noreferences.Config{
					Policy: "InvalidPolicy",
				},
				expectedErr: "noreferences.policy: Unsupported value: \"InvalidPolicy\": supported values: \"PreferAbbreviatedReference\", \"NoReferences\"",
			}),
		)
	})
})
