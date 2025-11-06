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

package preferredmarkers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/initializer"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/preferredmarkers"
)

var _ = Describe("preferredmarkers initializer", func() {
	Context("config validation", func() {
		type testCase struct {
			config      *preferredmarkers.Config
			expectedErr string
		}

		DescribeTable("should validate the provided config", func(in testCase) {
			ci, ok := preferredmarkers.Initializer().(initializer.ConfigurableAnalyzerInitializer)
			Expect(ok).To(BeTrue())

			errs := ci.ValidateConfig(in.config, field.NewPath("preferredmarkers"))
			if len(in.expectedErr) > 0 {
				Expect(errs.ToAggregate()).To(MatchError(in.expectedErr))
			} else {
				Expect(errs).To(HaveLen(0), "No errors were expected")
			}
		},
			Entry("With a valid preferredmarkers configuration", testCase{
				config: &preferredmarkers.Config{
					Markers: []preferredmarkers.Marker{
						{
							PreferredIdentifier: "k8s:optional",
							EquivalentIdentifiers: []preferredmarkers.EquivalentIdentifier{
								{Identifier: "kubebuilder:validation:Optional"},
							},
						},
					},
				},
				expectedErr: "",
			}),
			Entry("With a valid preferredmarkers configuration with multiple equivalents", testCase{
				config: &preferredmarkers.Config{
					Markers: []preferredmarkers.Marker{
						{
							PreferredIdentifier: "custom:preferred",
							EquivalentIdentifiers: []preferredmarkers.EquivalentIdentifier{
								{Identifier: "custom:old"},
								{Identifier: "custom:deprecated"},
								{Identifier: "custom:legacy"},
							},
						},
					},
				},
				expectedErr: "",
			}),
			Entry("With a valid preferredmarkers configuration with multiple markers", testCase{
				config: &preferredmarkers.Config{
					Markers: []preferredmarkers.Marker{
						{
							PreferredIdentifier: "k8s:optional",
							EquivalentIdentifiers: []preferredmarkers.EquivalentIdentifier{
								{Identifier: "kubebuilder:validation:Optional"},
							},
						},
						{
							PreferredIdentifier: "k8s:required",
							EquivalentIdentifiers: []preferredmarkers.EquivalentIdentifier{
								{Identifier: "kubebuilder:validation:Required"},
							},
						},
					},
				},
				expectedErr: "",
			}),
			Entry("With an invalid preferredmarkers configuration, duplicate preferred markers", testCase{
				config: &preferredmarkers.Config{
					Markers: []preferredmarkers.Marker{
						{
							PreferredIdentifier: "k8s:optional",
							EquivalentIdentifiers: []preferredmarkers.EquivalentIdentifier{
								{Identifier: "kubebuilder:validation:Optional"},
							},
						},
						{
							PreferredIdentifier: "k8s:optional",
							EquivalentIdentifiers: []preferredmarkers.EquivalentIdentifier{
								{Identifier: "custom:optional"},
							},
						},
					},
				},
				expectedErr: "preferredmarkers.markers[1].preferredIdentifier: Duplicate value: \"k8s:optional\"",
			}),
			Entry("With an invalid preferredmarkers configuration, duplicate equivalent identifiers within a marker", testCase{
				config: &preferredmarkers.Config{
					Markers: []preferredmarkers.Marker{
						{
							PreferredIdentifier: "k8s:optional",
							EquivalentIdentifiers: []preferredmarkers.EquivalentIdentifier{
								{Identifier: "kubebuilder:validation:Optional"},
								{Identifier: "kubebuilder:validation:Optional"},
							},
						},
					},
				},
				expectedErr: "preferredmarkers.markers[0].equivalentIdentifiers[1].identifier: Duplicate value: \"kubebuilder:validation:Optional\"",
			}),
			Entry("With an invalid preferredmarkers configuration, duplicate equivalent identifiers across markers", testCase{
				config: &preferredmarkers.Config{
					Markers: []preferredmarkers.Marker{
						{
							PreferredIdentifier: "k8s:optional",
							EquivalentIdentifiers: []preferredmarkers.EquivalentIdentifier{
								{Identifier: "custom:optional"},
							},
						},
						{
							PreferredIdentifier: "k8s:required",
							EquivalentIdentifiers: []preferredmarkers.EquivalentIdentifier{
								{Identifier: "custom:optional"},
							},
						},
					},
				},
				expectedErr: "preferredmarkers.markers[1].equivalentIdentifiers[0].identifier: Duplicate value: \"custom:optional\"",
			}),
			Entry("With an invalid preferredmarkers configuration, equivalent identifier is same as preferred identifier", testCase{
				config: &preferredmarkers.Config{
					Markers: []preferredmarkers.Marker{
						{
							PreferredIdentifier: "k8s:optional",
							EquivalentIdentifiers: []preferredmarkers.EquivalentIdentifier{
								{Identifier: "kubebuilder:validation:Optional"},
							},
						},
						{
							PreferredIdentifier: "k8s:required",
							EquivalentIdentifiers: []preferredmarkers.EquivalentIdentifier{
								{Identifier: "k8s:optional"},
							},
						},
					},
				},
				expectedErr: "preferredmarkers.markers[1].equivalentIdentifiers[0].identifier: Invalid value: \"k8s:optional\": equivalent identifier cannot be the same as a preferred identifier",
			}),
			Entry("With an invalid preferredmarkers configuration, preferred identifier in its own equivalent identifiers list", testCase{
				config: &preferredmarkers.Config{
					Markers: []preferredmarkers.Marker{
						{
							PreferredIdentifier: "k8s:optional",
							EquivalentIdentifiers: []preferredmarkers.EquivalentIdentifier{
								{Identifier: "kubebuilder:validation:Optional"},
								{Identifier: "k8s:optional"},
							},
						},
					},
				},
				expectedErr: "preferredmarkers.markers[0].equivalentIdentifiers[1].identifier: Invalid value: \"k8s:optional\": equivalent identifier cannot be the same as a preferred identifier",
			}),
			Entry("With an invalid preferredmarkers configuration, no equivalent identifiers", testCase{
				config: &preferredmarkers.Config{
					Markers: []preferredmarkers.Marker{
						{
							PreferredIdentifier:   "k8s:optional",
							EquivalentIdentifiers: []preferredmarkers.EquivalentIdentifier{},
						},
					},
				},
				expectedErr: "preferredmarkers.markers[0].equivalentIdentifiers: Required value: must contain at least one equivalent identifier",
			}),
			Entry("With a nil config", testCase{
				config:      nil,
				expectedErr: "preferredmarkers: Required value: configuration is required for the preferredmarkers linter when it is enabled",
			}),
			Entry("With an invalid preferredmarkers configuration with no markers", testCase{
				config: &preferredmarkers.Config{
					Markers: []preferredmarkers.Marker{},
				},
				expectedErr: "preferredmarkers.markers: Required value: must contain at least one preferred marker",
			}),
		)
	})
})
