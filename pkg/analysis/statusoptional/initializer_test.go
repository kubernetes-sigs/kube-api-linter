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

package statusoptional_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/initializer"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/statusoptional"
	"sigs.k8s.io/kube-api-linter/pkg/markers"
)

var _ = Describe("statusoptional initializer", func() {
	Context("config validation", func() {
		type testCase struct {
			config      statusoptional.StatusOptionalConfig
			expectedErr string
		}

		DescribeTable("should validate the provided config", func(in testCase) {
			ci, ok := statusoptional.Initializer().(initializer.ConfigurableAnalyzerInitializer)
			Expect(ok).To(BeTrue())

			errs := ci.ValidateConfig(&in.config, field.NewPath("statusoptional"))
			if len(in.expectedErr) > 0 {
				Expect(errs.ToAggregate()).To(MatchError(in.expectedErr))
			} else {
				Expect(errs).To(HaveLen(0), "No errors were expected")
			}
		},
			Entry("With a valid StatusOptionalConfig", testCase{
				config: statusoptional.StatusOptionalConfig{
					PreferredOptionalMarker: markers.OptionalMarker,
				},
				expectedErr: "",
			}),
			Entry("With a valid StatusOptionalConfig: k8s optional marker", testCase{
				config: statusoptional.StatusOptionalConfig{
					PreferredOptionalMarker: markers.K8sOptionalMarker,
				},
				expectedErr: "",
			}),
			Entry("With a valid StatusOptionalConfig: kubebuilder optional marker", testCase{
				config: statusoptional.StatusOptionalConfig{
					PreferredOptionalMarker: markers.KubebuilderOptionalMarker,
				},
			}),
			Entry("With an invalid StatusOptionalConfig", testCase{
				config: statusoptional.StatusOptionalConfig{
					PreferredOptionalMarker: "invalid",
				},
				expectedErr: "statusoptional.preferredOptionalMarker: Invalid value: \"invalid\": invalid value, must be one of \"optional\", \"kubebuilder:validation:Optional\", \"k8s:optional\" or omitted",
			}),
		)
	})
})
