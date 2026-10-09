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
package base_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/tools/go/analysis"

	"sigs.k8s.io/kube-api-linter/pkg/config/presets"
	"sigs.k8s.io/kube-api-linter/pkg/plugin/base"
	_ "sigs.k8s.io/kube-api-linter/pkg/registration"
)

func analyzerNames(analyzers []*analysis.Analyzer) []string {
	names := make([]string, 0, len(analyzers))
	for _, a := range analyzers {
		names = append(names, a.Name)
	}

	return names
}

var _ = Describe("Presets end-to-end", func() {
	Context("Custom preset", func() {
		BeforeEach(func() {
			presets.DefaultRegistry().Register("MyCustomPreset", presets.Preset{
				"conditions",
				"jsontags",
				"maxlength",
			})
		})

		It("should enable only the preset's linters", func() {
			plugin, err := base.New(map[string]any{
				"preset": "MyCustomPreset",
			})
			Expect(err).NotTo(HaveOccurred())

			analyzers, err := plugin.BuildAnalyzers()
			Expect(err).NotTo(HaveOccurred())
			Expect(analyzerNames(analyzers)).To(ConsistOf("conditions", "jsontags", "maxlength"))
		})

		It("should allow user enable to add to the preset", func() {
			plugin, err := base.New(map[string]any{
				"preset": "MyCustomPreset",
				"linters": map[string]any{
					"enable": []any{"nobools"},
				},
			})
			Expect(err).NotTo(HaveOccurred())

			analyzers, err := plugin.BuildAnalyzers()
			Expect(err).NotTo(HaveOccurred())
			Expect(analyzerNames(analyzers)).To(ConsistOf("conditions", "jsontags", "maxlength", "nobools"))
		})

		It("should allow user disable to remove from the preset", func() {
			plugin, err := base.New(map[string]any{
				"preset": "MyCustomPreset",
				"linters": map[string]any{
					"disable": []any{"jsontags"},
				},
			})
			Expect(err).NotTo(HaveOccurred())

			analyzers, err := plugin.BuildAnalyzers()
			Expect(err).NotTo(HaveOccurred())
			Expect(analyzerNames(analyzers)).To(ConsistOf("conditions", "maxlength"))
		})
	})

	Context("Built-in presets", func() {
		It("should work with BuiltIn preset", func() {
			plugin, err := base.New(map[string]any{
				"preset": "BuiltIn",
			})
			Expect(err).NotTo(HaveOccurred())

			analyzers, err := plugin.BuildAnalyzers()
			Expect(err).NotTo(HaveOccurred())

			builtIn, ok := presets.DefaultRegistry().Get("BuiltIn")
			Expect(ok).To(BeTrue())
			Expect(analyzerNames(analyzers)).To(ConsistOf([]string(builtIn)))
		})

		It("should work with CustomResource preset", func() {
			plugin, err := base.New(map[string]any{
				"preset": "CustomResource",
			})
			Expect(err).NotTo(HaveOccurred())

			analyzers, err := plugin.BuildAnalyzers()
			Expect(err).NotTo(HaveOccurred())

			cr, ok := presets.DefaultRegistry().Get("CustomResource")
			Expect(ok).To(BeTrue())
			Expect(analyzerNames(analyzers)).To(ConsistOf([]string(cr)))
		})
	})

	Context("Unknown preset", func() {
		It("should return an error", func() {
			plugin, err := base.New(map[string]any{
				"preset": "DoesNotExist",
			})
			Expect(err).NotTo(HaveOccurred())

			_, err = plugin.BuildAnalyzers()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("unknown preset"))
			Expect(err.Error()).To(ContainSubstring("DoesNotExist"))
		})
	})

	Context("No preset (backward compatibility)", func() {
		It("should use default linters when no preset is set", func() {
			plugin, err := base.New(map[string]any{})
			Expect(err).NotTo(HaveOccurred())

			analyzers, err := plugin.BuildAnalyzers()
			Expect(err).NotTo(HaveOccurred())

			builtIn, ok := presets.DefaultRegistry().Get("BuiltIn")
			Expect(ok).To(BeTrue())
			Expect(analyzerNames(analyzers)).To(ConsistOf([]string(builtIn)))
		})
	})
})
