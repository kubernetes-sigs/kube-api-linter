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
package presets_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"sigs.k8s.io/kube-api-linter/pkg/config/presets"
)

var _ = Describe("Registry", func() {
	var r presets.Registry

	BeforeEach(func() {
		r = presets.NewRegistry()
	})

	Context("Register and Get", func() {
		It("should return a registered preset", func() {
			r.Register("TestPreset", presets.Preset{"linterA", "linterB"})

			p, ok := r.Get("TestPreset")
			Expect(ok).To(BeTrue())
			Expect([]string(p)).To(ConsistOf("linterA", "linterB"))
		})

		It("should return false for an unknown preset", func() {
			_, ok := r.Get("Unknown")
			Expect(ok).To(BeFalse())
		})
	})

	Context("All", func() {
		It("should return all registered preset names", func() {
			r.Register("Alpha", presets.Preset{"a"})
			r.Register("Beta", presets.Preset{"b"})

			Expect(r.All()).To(ConsistOf("Alpha", "Beta"))
		})

		It("should return empty for no presets", func() {
			Expect(r.All()).To(BeEmpty())
		})
	})

	Context("Custom preset registration (fork extensibility)", func() {
		It("should allow registering a custom preset and retrieving it", func() {
			custom := presets.Preset{"customLinterA", "customLinterB", "customLinterC"}
			r.Register("MyOrgPreset", custom)

			p, ok := r.Get("MyOrgPreset")
			Expect(ok).To(BeTrue())
			Expect([]string(p)).To(ConsistOf("customLinterA", "customLinterB", "customLinterC"))
		})
	})

	Context("DefaultRegistry", func() {
		It("should return a non-nil registry", func() {
			Expect(presets.DefaultRegistry()).NotTo(BeNil())
		})
	})
})
