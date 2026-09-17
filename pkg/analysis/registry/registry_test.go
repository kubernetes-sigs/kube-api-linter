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
package registry_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"golang.org/x/tools/go/analysis"

	"sigs.k8s.io/kube-api-linter/pkg/analysis/commentstart"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/conditions"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/jsontags"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/nobools"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/optionalorrequired"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/registry"
	"sigs.k8s.io/kube-api-linter/pkg/config"
	"sigs.k8s.io/kube-api-linter/pkg/markers"
)

var _ = Describe("Registry", func() {
	var r registry.Registry

	BeforeEach(func() {
		r = registry.NewRegistry()

		// Register a selection of linters to test the registry functionality.
		r.RegisterLinter(conditions.Initializer())
		r.RegisterLinter(jsontags.Initializer())
		r.RegisterLinter(optionalorrequired.Initializer())
		r.RegisterLinter(nobools.Initializer())
		r.RegisterLinter(commentstart.Initializer())
	})

	Context("DefaultLinters", func() {
		It("should return the default linters", func() {
			Expect(r.DefaultLinters().UnsortedList()).To(ConsistOf(
				"conditions",
				"jsontags",
				"optionalorrequired",
				"commentstart",
			))
		})
	})

	Context("AllLinters", func() {
		It("should return the all known linters", func() {
			Expect(r.AllLinters().UnsortedList()).To(ConsistOf(
				"conditions",
				"jsontags",
				"optionalorrequired",
				"nobools",
				"commentstart",
			))
		})
	})

	Context("InitializeLinters", func() {
		type initLintersTableInput struct {
			config        config.Linters
			lintersConfig config.LintersConfig

			expectedLinters []string
		}

		DescribeTable("Initialize Linters", func(in initLintersTableInput) {
			linters, err := r.InitializeLinters(in.config, in.lintersConfig)
			Expect(err).NotTo(HaveOccurred())

			toLinterNames := func(a []*analysis.Analyzer) []string {
				names := []string{}

				for _, linter := range a {
					names = append(names, linter.Name)
				}

				return names
			}

			Expect(linters).To(WithTransform(toLinterNames, ConsistOf(in.expectedLinters)))
		},
			Entry("Empty config", initLintersTableInput{
				config:          config.Linters{},
				lintersConfig:   config.LintersConfig{},
				expectedLinters: []string{"conditions", "jsontags", "optionalorrequired", "commentstart"},
			}),
			Entry("With wildcard enabled linters", initLintersTableInput{
				config: config.Linters{
					Enable: []string{config.Wildcard},
				},
				lintersConfig:   config.LintersConfig{},
				expectedLinters: []string{"conditions", "jsontags", "optionalorrequired", "nobools", "commentstart"},
			}),
			Entry("With wildcard enabled linters and a disabled linter", initLintersTableInput{
				config: config.Linters{
					Enable:  []string{config.Wildcard},
					Disable: []string{"jsontags"},
				},
				lintersConfig:   config.LintersConfig{},
				expectedLinters: []string{"conditions", "optionalorrequired", "nobools", "commentstart"},
			}),
			Entry("With wildcard disabled linters", initLintersTableInput{
				config: config.Linters{
					Disable: []string{config.Wildcard},
				},
				lintersConfig:   config.LintersConfig{},
				expectedLinters: []string{},
			}),
			Entry("With wildcard disabled linters and an enabled linter", initLintersTableInput{
				config: config.Linters{
					Disable: []string{config.Wildcard},
					Enable:  []string{"jsontags"},
				},
				lintersConfig:   config.LintersConfig{},
				expectedLinters: []string{"jsontags"},
			}),
		)
	})

	Context("Config validation", func() {
		type validateLintersConfigTableInput struct {
			linters     config.Linters
			config      config.LintersConfig
			expectedErr string
		}

		DescribeTable("Validate Linters Configuration through Initialization", func(in validateLintersConfigTableInput) {
			_, err := r.InitializeLinters(in.linters, in.config)
			if len(in.expectedErr) > 0 {
				Expect(err).To(MatchError(in.expectedErr))
			} else {
				Expect(err).To(Not(HaveOccurred()), "No errors were expected")
			}
		},
			Entry("Empty config", validateLintersConfigTableInput{
				linters:     config.Linters{},
				config:      config.LintersConfig{},
				expectedErr: "",
			}),

			Entry("With a valid JSONTagsConfig JSONTagRegex", validateLintersConfigTableInput{
				linters: config.Linters{},
				config: config.LintersConfig{
					"jsontags": jsontags.JSONTagsConfig{
						JSONTagRegex: "^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*)*$",
					},
				},
			}),
			Entry("With an invalid JSONTagsConfig JSONTagRegex", validateLintersConfigTableInput{
				config: config.LintersConfig{
					"jsontags": jsontags.JSONTagsConfig{
						JSONTagRegex: "^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*",
					},
				},
				expectedErr: "error validating linters config: lintersConfig.jsontags.jsonTagRegex: Invalid value: \"^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*\": invalid regex: error parsing regexp: missing closing ): `^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*`",
			}),
			Entry("With a valid JSONTagsConfig JSONTagRegex (legacy field name)", validateLintersConfigTableInput{
				linters: config.Linters{},
				config: config.LintersConfig{
					"jsonTags": jsontags.JSONTagsConfig{
						JSONTagRegex: "^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*)*$",
					},
				},
			}),
			Entry("With an invalid JSONTagsConfig JSONTagRegex (legacy field name)", validateLintersConfigTableInput{
				config: config.LintersConfig{
					"jsonTags": jsontags.JSONTagsConfig{
						JSONTagRegex: "^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*",
					},
				},
				expectedErr: "error validating linters config: lintersConfig.jsontags.jsonTagRegex: Invalid value: \"^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*\": invalid regex: error parsing regexp: missing closing ): `^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*`",
			}),

			Entry("With a valid CommentStartConfig", validateLintersConfigTableInput{
				linters: config.Linters{},
				config: config.LintersConfig{
					"commentstart": commentstart.Config{
						ExcludePrefixes: []string{"TODO:"},
					},
				},
			}),
			Entry("With an invalid CommentStartConfig - duplicate prefix", validateLintersConfigTableInput{
				config: config.LintersConfig{
					"commentstart": commentstart.Config{
						ExcludePrefixes: []string{"TODO:", "TODO:"},
					},
				},
				expectedErr: "error validating linters config: lintersConfig.commentstart.excludePrefixes[1]: Duplicate value: \"TODO:\"",
			}),
			Entry("With an invalid CommentStartConfig - duplicates built-in default", validateLintersConfigTableInput{
				config: config.LintersConfig{
					"commentstart": commentstart.Config{
						ExcludePrefixes: []string{"Deprecated:"},
					},
				},
				expectedErr: "error validating linters config: lintersConfig.commentstart.excludePrefixes[0]: Invalid value: \"Deprecated:\": is already a built-in default prefix",
			}),

			Entry("With a valid OptionalOrRequiredConfig (legacy field name)", validateLintersConfigTableInput{
				config: config.LintersConfig{
					"optionalorrequired": optionalorrequired.OptionalOrRequiredConfig{
						PreferredOptionalMarker: markers.OptionalMarker,
						PreferredRequiredMarker: markers.RequiredMarker,
					},
				},
				expectedErr: "",
			}),
			Entry("With kubebuilder preferred markers (legacy field name)", validateLintersConfigTableInput{
				config: config.LintersConfig{
					"optionalorrequired": optionalorrequired.OptionalOrRequiredConfig{
						PreferredOptionalMarker: markers.KubebuilderOptionalMarker,
						PreferredRequiredMarker: markers.KubebuilderRequiredMarker,
					},
				},
				expectedErr: "",
			}),
			Entry("With invalid preferred optional marker (legacy field name)", validateLintersConfigTableInput{
				config: config.LintersConfig{
					"optionalorrequired": optionalorrequired.OptionalOrRequiredConfig{
						PreferredOptionalMarker: "invalid",
					},
				},
				expectedErr: "error validating linters config: lintersConfig.optionalorrequired.preferredOptionalMarker: Invalid value: \"invalid\": invalid value, must be one of \"optional\", \"kubebuilder:validation:Optional\" or omitted",
			}),
			Entry("With invalid preferred required marker (legacy field name)", validateLintersConfigTableInput{
				config: config.LintersConfig{
					"optionalorrequired": optionalorrequired.OptionalOrRequiredConfig{
						PreferredRequiredMarker: "invalid",
					},
				},
				expectedErr: "error validating linters config: lintersConfig.optionalorrequired.preferredRequiredMarker: Invalid value: \"invalid\": invalid value, must be one of \"required\", \"kubebuilder:validation:Required\" or omitted",
			}),
			Entry("With a valid OptionalOrRequiredConfig (legacy field name)", validateLintersConfigTableInput{
				config: config.LintersConfig{
					"optionalOrRequired": optionalorrequired.OptionalOrRequiredConfig{
						PreferredOptionalMarker: markers.OptionalMarker,
						PreferredRequiredMarker: markers.RequiredMarker,
					},
				},
				expectedErr: "",
			}),
			Entry("With kubebuilder preferred markers (legacy field name)", validateLintersConfigTableInput{
				config: config.LintersConfig{
					"optionalOrRequired": optionalorrequired.OptionalOrRequiredConfig{
						PreferredOptionalMarker: markers.KubebuilderOptionalMarker,
						PreferredRequiredMarker: markers.KubebuilderRequiredMarker,
					},
				},
				expectedErr: "",
			}),
			Entry("With invalid preferred optional marker (legacy field name)", validateLintersConfigTableInput{
				config: config.LintersConfig{
					"optionalOrRequired": optionalorrequired.OptionalOrRequiredConfig{
						PreferredOptionalMarker: "invalid",
					},
				},
				expectedErr: "error validating linters config: lintersConfig.optionalorrequired.preferredOptionalMarker: Invalid value: \"invalid\": invalid value, must be one of \"optional\", \"kubebuilder:validation:Optional\" or omitted",
			}),
			Entry("With invalid preferred required marker (legacy field name)", validateLintersConfigTableInput{
				config: config.LintersConfig{
					"optionalOrRequired": optionalorrequired.OptionalOrRequiredConfig{
						PreferredRequiredMarker: "invalid",
					},
				},
				expectedErr: "error validating linters config: lintersConfig.optionalorrequired.preferredRequiredMarker: Invalid value: \"invalid\": invalid value, must be one of \"required\", \"kubebuilder:validation:Required\" or omitted",
			}),

			// Tests for disabled linters with configuration
			Entry("With config for explicitly disabled linter should not error", validateLintersConfigTableInput{
				linters: config.Linters{
					Disable: []string{"jsontags"},
				},
				config: config.LintersConfig{
					"jsontags": jsontags.JSONTagsConfig{
						JSONTagRegex: "^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*)*$",
					},
				},
				expectedErr: "",
			}),
			Entry("With config for disabled linter when all disabled with wildcard should not error", validateLintersConfigTableInput{
				linters: config.Linters{
					Disable: []string{config.Wildcard},
				},
				config: config.LintersConfig{
					"jsontags": jsontags.JSONTagsConfig{
						JSONTagRegex: "^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*)*$",
					},
				},
				expectedErr: "",
			}),
			Entry("With config for non-existent linter should error", validateLintersConfigTableInput{
				config: config.LintersConfig{
					"nonexistent": map[string]any{
						"someOption": "value",
					},
				},
				expectedErr: "error validating linters config: lintersConfig.nonexistent: Invalid value: \"nonexistent\": unknown linter",
			}),
			Entry("With config for non-configurable linter should error", validateLintersConfigTableInput{
				config: config.LintersConfig{
					"nobools": map[string]any{
						"someOption": "value",
					},
				},
				expectedErr: "error validating linters config: lintersConfig.nobools: Invalid value: \"nobools\": linter is not configurable",
			}),
		)
	})
})
