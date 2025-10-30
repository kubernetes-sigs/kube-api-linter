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

package dependenttags_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/dependenttags"
)

var _ = Describe("dependenttags initializer", func() {
	Context("config validation", func() {
		type testCase struct {
			config      dependenttags.Config
			expectedErr string
		}

		DescribeTable("should validate the provided config", func(in testCase) {
			ci := dependenttags.Initializer()

			errs := ci.ValidateConfig(&in.config, field.NewPath("dependenttags"))
			if len(in.expectedErr) > 0 {
				Expect(errs.ToAggregate()).To(MatchError(in.expectedErr))
			} else {
				Expect(errs).To(HaveLen(0), "No errors were expected")
			}
		},
			Entry("with a valid config", testCase{
				config: dependenttags.Config{
					Rules: []dependenttags.Rule{
						{
							Identifier: "k8s:unionMember",
							Type:       dependenttags.DependencyTypeAll,
							Dependents: []string{"k8s:optional"},
						},
					},
				},
				expectedErr: "",
			}),
			Entry("with missing type", testCase{
				config: dependenttags.Config{
					Rules: []dependenttags.Rule{
						{
							Identifier: "k8s:unionMember",
							Dependents: []string{"k8s:optional"},
						},
					},
				},
				expectedErr: fmt.Sprintf("dependenttags.rules[0].type: Required value: type must be explicitly set to '%s' or '%s'", dependenttags.DependencyTypeAll, dependenttags.DependencyTypeAny),
			}),
			Entry("with empty rules", testCase{
				config: dependenttags.Config{
					Rules: []dependenttags.Rule{},
				},
				expectedErr: "dependenttags.rules: Invalid value: []dependenttags.Rule{}: rules cannot be empty",
			}),
			Entry("with missing identifier", testCase{
				config: dependenttags.Config{
					Rules: []dependenttags.Rule{
						{
							Type:       dependenttags.DependencyTypeAll,
							Dependents: []string{"k8s:optional"},
						},
					},
				},
				expectedErr: "dependenttags.rules[0].identifier: Invalid value: \"\": identifier marker cannot be empty",
			}),
			Entry("with missing dependents", testCase{
				config: dependenttags.Config{
					Rules: []dependenttags.Rule{
						{
							Identifier: "k8s:unionMember",
							Type:       dependenttags.DependencyTypeAll,
						},
					},
				},
				expectedErr: "dependenttags.rules[0].dependents: Invalid value: []string(nil): dependents list cannot be empty",
			}),
			Entry("with invalid type", testCase{
				config: dependenttags.Config{
					Rules: []dependenttags.Rule{
						{
							Identifier: "k8s:unionMember",
							Type:       "invalid",
							Dependents: []string{"k8s:optional"},
						},
					},
				},
				expectedErr: fmt.Sprintf(`dependenttags.rules[0].type: Invalid value: "invalid": type must be '%s' or '%s'`, dependenttags.DependencyTypeAll, dependenttags.DependencyTypeAny),
			}),
		)
	})
})
