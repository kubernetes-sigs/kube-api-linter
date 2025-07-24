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
			Entry("With a valid config with single conflict", testCase{
				config: conflictingmarkers.ConflictingMarkersConfig{
					Conflicts: []conflictingmarkers.ConflictSet{
						{
							Name:        "test_conflict",
							Sets:        [][]string{{"marker1"}, {"marker2"}},
							Description: "Test conflict",
						},
					},
				},
				expectedErr: "",
			}),
			Entry("With a valid config with multiple conflicts", testCase{
				config: conflictingmarkers.ConflictingMarkersConfig{
					Conflicts: []conflictingmarkers.ConflictSet{
						{
							Name:        "test_conflict",
							Sets:        [][]string{{"marker1"}, {"marker2"}},
							Description: "Test conflict",
						},
						{
							Name:        "another_conflict",
							Sets:        [][]string{{"marker3", "marker4"}, {"marker5"}},
							Description: "Another test conflict",
						},
					},
				},
				expectedErr: "",
			}),
			Entry("With empty conflicts list", testCase{
				config: conflictingmarkers.ConflictingMarkersConfig{
					Conflicts: []conflictingmarkers.ConflictSet{},
				},
				expectedErr: "conflictingmarkers.conflicts: Required value: at least one conflict set is required",
			}),
			Entry("With missing name", testCase{
				config: conflictingmarkers.ConflictingMarkersConfig{
					Conflicts: []conflictingmarkers.ConflictSet{
						{
							Name:        "",
							Sets:        [][]string{{"marker1"}, {"marker2"}},
							Description: "Test conflict",
						},
					},
				},
				expectedErr: "conflictingmarkers.conflicts[0].name: Required value: name is required",
			}),
			Entry("With missing description", testCase{
				config: conflictingmarkers.ConflictingMarkersConfig{
					Conflicts: []conflictingmarkers.ConflictSet{
						{
							Name: "test_conflict",
							Sets: [][]string{{"marker1"}, {"marker2"}},
						},
					},
				},
				expectedErr: "conflictingmarkers.conflicts[0].description: Required value: description is required",
			}),
			Entry("With insufficient sets (only 1 set)", testCase{
				config: conflictingmarkers.ConflictingMarkersConfig{
					Conflicts: []conflictingmarkers.ConflictSet{
						{
							Name:        "test_conflict",
							Sets:        [][]string{{"marker1"}},
							Description: "Test conflict",
						},
					},
				},
				expectedErr: "conflictingmarkers.conflicts[0].sets: Required value: at least 2 sets are required",
			}),
			Entry("With empty set", testCase{
				config: conflictingmarkers.ConflictingMarkersConfig{
					Conflicts: []conflictingmarkers.ConflictSet{
						{
							Name:        "test_conflict",
							Sets:        [][]string{{"marker1"}, {}},
							Description: "Test conflict",
						},
					},
				},
				expectedErr: "conflictingmarkers.conflicts[0].sets[1]: Required value: set cannot be empty",
			}),
			Entry("With overlapping markers between sets", testCase{
				config: conflictingmarkers.ConflictingMarkersConfig{
					Conflicts: []conflictingmarkers.ConflictSet{
						{
							Name:        "test_conflict",
							Sets:        [][]string{{"marker1", "marker2"}, {"marker2", "marker3"}},
							Description: "Test conflict",
						},
					},
				},
				expectedErr: "conflictingmarkers.conflicts[0].sets: Invalid value: conflictingmarkers.ConflictSet{Name:\"test_conflict\", Sets:[][]string{[]string{\"marker1\", \"marker2\"}, []string{\"marker2\", \"marker3\"}}, Description:\"Test conflict\"}: sets 1 and 2 cannot contain overlapping markers: [marker2]",
			}),
			Entry("With duplicate conflict names", testCase{
				config: conflictingmarkers.ConflictingMarkersConfig{
					Conflicts: []conflictingmarkers.ConflictSet{
						{
							Name:        "duplicate_name",
							Sets:        [][]string{{"marker1"}, {"marker2"}},
							Description: "First conflict",
						},
						{
							Name:        "duplicate_name",
							Sets:        [][]string{{"marker3"}, {"marker4"}},
							Description: "Second conflict",
						},
					},
				},
				expectedErr: "conflictingmarkers.conflicts[1].name: Duplicate value: \"duplicate_name\"",
			}),
			Entry("With three-way conflict", testCase{
				config: conflictingmarkers.ConflictingMarkersConfig{
					Conflicts: []conflictingmarkers.ConflictSet{
						{
							Name:        "three_way_conflict",
							Sets:        [][]string{{"marker1"}, {"marker2"}, {"marker3"}},
							Description: "Three-way conflict",
						},
					},
				},
				expectedErr: "",
			}),
		)
	})
})
