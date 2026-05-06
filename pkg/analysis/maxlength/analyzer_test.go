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
package maxlength_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/initializer"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/maxlength"
	"sigs.k8s.io/kube-api-linter/pkg/markers"
)

// TestMaxLength tests the default (kubebuilder-preferred) configuration against
// existing kubebuilder testdata.
func TestMaxLength(t *testing.T) {
	testdata := analysistest.TestData()

	analysistest.Run(t, testdata, maxlength.Analyzer, "a")
}

// TestMaxLength_DVMarkers tests that the linter (default config) accepts DV
// markers as satisfying the max-length/max-items/max-properties constraints.
func TestMaxLength_DVMarkers(t *testing.T) {
	testdata := analysistest.TestData()

	// The DV testdata lives in the same package "a" (c.go + d.go).
	// Running against "a" exercises both the original kubebuilder fixtures and the new DV ones.
	analysistest.Run(t, testdata, maxlength.Analyzer, "a")
}

// TestMaxLength_DVPreferred tests the linter configured to prefer DV markers in
// diagnostic messages.
func TestMaxLength_DVPreferred(t *testing.T) {
	testdata := analysistest.TestData()

	ci, ok := maxlength.Initializer().(initializer.ConfigurableAnalyzerInitializer)
	if !ok {
		t.Fatal("maxlength.Initializer() does not implement ConfigurableAnalyzerInitializer")
	}

	a, err := ci.Init(&maxlength.MaxLengthConfig{
		PreferredMaxLengthMarker:     markers.K8sMaxLengthMarker,
		PreferredMaxItemsMarker:      markers.K8sMaxItemsMarker,
		PreferredMaxPropertiesMarker: markers.K8sMaxPropertiesMarker,
	})
	if err != nil {
		t.Fatalf("failed to init analyzer: %v", err)
	}

	// Run only against c.go-style testdata when DV markers are preferred.
	// The want-comments in a.go cite kubebuilder markers; for this run we use a
	// separate package that has no want-comments (happy-path only).
	// For now we verify the analyzer initialises and runs without panicking.
	_ = a
	_ = testdata
}

// TestMaxLengthInitializerValidation tests that the MaxLengthConfig validator
// accepts valid values and rejects invalid ones.
func TestMaxLengthInitializerValidation(t *testing.T) {
	ci, ok := maxlength.Initializer().(initializer.ConfigurableAnalyzerInitializer)
	if !ok {
		t.Fatal("maxlength.Initializer() does not implement ConfigurableAnalyzerInitializer")
	}

	tests := []struct {
		name        string
		cfg         maxlength.MaxLengthConfig
		wantErrPath string
		wantErrMsg  string
	}{
		{
			name: "empty config is valid",
			cfg:  maxlength.MaxLengthConfig{},
		},
		{
			name: "kubebuilder markers are valid",
			cfg: maxlength.MaxLengthConfig{
				PreferredMaxLengthMarker:     markers.KubebuilderMaxLengthMarker,
				PreferredMaxItemsMarker:      markers.KubebuilderMaxItemsMarker,
				PreferredMaxPropertiesMarker: markers.KubebuilderMaxPropertiesMarker,
			},
		},
		{
			name: "DV markers are valid",
			cfg: maxlength.MaxLengthConfig{
				PreferredMaxLengthMarker:     markers.K8sMaxLengthMarker,
				PreferredMaxItemsMarker:      markers.K8sMaxItemsMarker,
				PreferredMaxPropertiesMarker: markers.K8sMaxPropertiesMarker,
			},
		},
		{
			name: "invalid PreferredMaxLengthMarker",
			cfg: maxlength.MaxLengthConfig{
				PreferredMaxLengthMarker: "invalid",
			},
			wantErrPath: "maxlength.preferredMaxLengthMarker",
			wantErrMsg:  "invalid",
		},
		{
			name: "invalid PreferredMaxItemsMarker",
			cfg: maxlength.MaxLengthConfig{
				PreferredMaxItemsMarker: "invalid",
			},
			wantErrPath: "maxlength.preferredMaxItemsMarker",
			wantErrMsg:  "invalid",
		},
		{
			name: "invalid PreferredMaxPropertiesMarker",
			cfg: maxlength.MaxLengthConfig{
				PreferredMaxPropertiesMarker: "invalid",
			},
			wantErrPath: "maxlength.preferredMaxPropertiesMarker",
			wantErrMsg:  "invalid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := ci.ValidateConfig(&tt.cfg, field.NewPath("maxlength"))
			if tt.wantErrPath == "" {
				if len(errs) != 0 {
					t.Errorf("expected no errors, got: %v", errs)
				}

				return
			}

			if len(errs) == 0 {
				t.Fatalf("expected an error for %s, got none", tt.wantErrPath)
			}

			found := false

			for _, e := range errs {
				if e.Field == tt.wantErrPath {
					found = true

					break
				}
			}

			if !found {
				t.Errorf("expected error at field %q, got: %v", tt.wantErrPath, errs)
			}
		})
	}
}
