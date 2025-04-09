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
package validation_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"sigs.k8s.io/kube-api-linter/pkg/config"
	"sigs.k8s.io/kube-api-linter/pkg/validation"

	"k8s.io/apimachinery/pkg/util/validation/field"
)

var _ = Describe("Linters", func() {
	type validateLintersTableInput struct {
		config      config.Linters
		expectedErr string
	}

	DescribeTable("Validate Linters Configuration", func(in validateLintersTableInput) {
		errs := validation.ValidateLinters(in.config, field.NewPath("linters"))
		if len(in.expectedErr) > 0 {
			Expect(errs.ToAggregate()).To(MatchError(in.expectedErr))
		} else {
			Expect(errs).To(HaveLen(0), "No errors were expected")
		}

	},
		Entry("Empty config", validateLintersTableInput{
			config:      config.Linters{},
			expectedErr: "",
		}),
		Entry("With duplicate enabled linter names", validateLintersTableInput{
			config: config.Linters{
				Enable: []string{"jsontags", "jsontags"},
			},
			expectedErr: "linters.enable: Invalid value: []string{\"jsontags\", \"jsontags\"}: values in 'enable' must be unique",
		}),
		Entry("With a wildcard enable value", validateLintersTableInput{
			config: config.Linters{
				Enable: []string{"*"},
			},
			expectedErr: "",
		}),
		Entry("With a wildcard enable value and a disable value", validateLintersTableInput{
			config: config.Linters{
				Enable:  []string{"*"},
				Disable: []string{"jsontags"},
			},
			expectedErr: "",
		}),
		Entry("With a wildcard enable value and an additional value", validateLintersTableInput{
			config: config.Linters{
				Enable: []string{"*", "jsontags"},
			},
			expectedErr: "linters.enable: Invalid value: []string{\"*\", \"jsontags\"}: wildcard ('*') must not be specified with other values",
		}),
		Entry("With valid enabled linter names", validateLintersTableInput{
			config: config.Linters{
				Enable: []string{"jsontags", "commentstart", "optionalorrequired"},
			},
			expectedErr: "",
		}),
		Entry("With an invalid enabled linter name", validateLintersTableInput{
			config: config.Linters{
				Enable: []string{"jsontags", "invalid"},
			},
			expectedErr: "linters.enable: Invalid value: []string{\"jsontags\", \"invalid\"}: unknown linters: invalid",
		}),

		Entry("With duplicate disabled linter names", validateLintersTableInput{
			config: config.Linters{
				Disable: []string{"jsontags", "jsontags"},
			},
			expectedErr: "linters.disable: Invalid value: []string{\"jsontags\", \"jsontags\"}: values in 'disable' must be unique",
		}),
		Entry("With a wildcard disable value", validateLintersTableInput{
			config: config.Linters{
				Disable: []string{"*"},
			},
			expectedErr: "",
		}),
		Entry("With a wildcard disable value and a enable value", validateLintersTableInput{
			config: config.Linters{
				Enable:  []string{"jsontags"},
				Disable: []string{"*"},
			},
			expectedErr: "",
		}),
		Entry("With a wildcard disable value and an additional value", validateLintersTableInput{
			config: config.Linters{
				Disable: []string{"*", "jsontags"},
			},
			expectedErr: "linters.disable: Invalid value: []string{\"*\", \"jsontags\"}: wildcard ('*') must not be specified with other values",
		}),
		Entry("With valid disable linter names", validateLintersTableInput{
			config: config.Linters{
				Disable: []string{"jsontags", "commentstart", "optionalorrequired"},
			},
			expectedErr: "",
		}),
		Entry("With an invalid disabled linter name", validateLintersTableInput{
			config: config.Linters{
				Disable: []string{"jsontags", "invalid"},
			},
			expectedErr: "linters.disable: Invalid value: []string{\"jsontags\", \"invalid\"}: unknown linters: invalid",
		}),
		Entry("With a value in both enable and disable", validateLintersTableInput{
			config: config.Linters{
				Enable:  []string{"jsontags"},
				Disable: []string{"jsontags"},
			},
			expectedErr: "linters: Invalid value: config.Linters{Enable:[]string{\"jsontags\"}, Disable:[]string{\"jsontags\"}}: values in 'enable' and 'disable may not overlap, overlapping values: jsontags",
		}),
	)
})
