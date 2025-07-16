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

package conflictingmarkers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/conflictingmarkers"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/initializer"
)

var _ = Describe("conflictingmarkers initializer", func() {
	Context("config validation", func() {
		type testCase struct {
			config      conflictingmarkers.ConflictingMarkersConfig
			expectedErr string
		}

		DescribeTable("should validate the provided config", func(in testCase) {
			ci, ok := conflictingmarkers.Initializer().(initializer.ConfigurableAnalyzerInitializer)
			Expect(ok).To(BeTrue())

			errs := ci.ValidateConfig(&in.config, field.NewPath("conflictingmarkers"))
			if len(in.expectedErr) > 0 {
				Expect(errs.ToAggregate()).To(MatchError(in.expectedErr))
			} else {
				Expect(errs).To(HaveLen(0), "No errors were expected")
			}
		},
			Entry("With a valid empty config", testCase{
				config: conflictingmarkers.ConflictingMarkersConfig{
					CustomConflicts: []conflictingmarkers.ConflictSet{},
				},
				expectedErr: "",
			}),
			Entry("With a valid custom conflict", testCase{
				config: conflictingmarkers.ConflictingMarkersConfig{
					CustomConflicts: []conflictingmarkers.ConflictSet{
						{
							Name:        "test_conflict",
							SetA:        []string{"marker1"},
							SetB:        []string{"marker2"},
							Description: "Test conflict",
						},
					},
				},
				expectedErr: "",
			}),
			Entry("With missing name", testCase{
				config: conflictingmarkers.ConflictingMarkersConfig{
					CustomConflicts: []conflictingmarkers.ConflictSet{
						{
							Name:        "",
							SetA:        []string{"marker1"},
							SetB:        []string{"marker2"},
							Description: "Test conflict",
						},
					},
				},
				expectedErr: "conflictingmarkers.customConflicts[0].name: Required value: name is required",
			}),
			Entry("With overlapping markers", testCase{
				config: conflictingmarkers.ConflictingMarkersConfig{
					CustomConflicts: []conflictingmarkers.ConflictSet{
						{
							Name:        "test_conflict",
							SetA:        []string{"marker1", "marker2"},
							SetB:        []string{"marker2", "marker3"},
							Description: "Test conflict",
						},
					},
				},
				expectedErr: "conflictingmarkers.customConflicts[0]: Invalid value: conflictingmarkers.ConflictSet{Name:\"test_conflict\", SetA:[]string{\"marker1\", \"marker2\"}, SetB:[]string{\"marker2\", \"marker3\"}, Description:\"Test conflict\"}: sets cannot contain overlapping markers",
			}),
			Entry("With missing description", testCase{
				config: conflictingmarkers.ConflictingMarkersConfig{
					CustomConflicts: []conflictingmarkers.ConflictSet{
						{
							Name: "test_conflict",
							SetA: []string{"marker1"},
							SetB: []string{"marker2"},
						},
					},
				},
				expectedErr: "conflictingmarkers.customConflicts[0].description: Required value: description is required",
			}),
		)
	})
})
