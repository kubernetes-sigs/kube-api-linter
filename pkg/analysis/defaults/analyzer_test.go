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
package defaults_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/defaults"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils/serialization"
	"sigs.k8s.io/kube-api-linter/pkg/markers"
)

func TestDefaultConfiguration(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := defaults.Initializer().Init(&defaults.DefaultsConfig{})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.RunWithSuggestedFixes(t, testdata, a, "a")
}

func TestKubebuilderDefaultMarkerPreferred(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := defaults.Initializer().Init(&defaults.DefaultsConfig{
		PreferredDefaultMarker: markers.KubebuilderDefaultMarker,
	})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.RunWithSuggestedFixes(t, testdata, a, "b")
}

func TestOmitZeroForbidden(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := defaults.Initializer().Init(&defaults.DefaultsConfig{
		OmitZero: defaults.DefaultsOmitZero{
			Policy: serialization.OmitZeroPolicyForbid,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.RunWithSuggestedFixes(t, testdata, a, "c")
}

func TestOmitEmptyWarn(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := defaults.Initializer().Init(&defaults.DefaultsConfig{
		OmitEmpty: defaults.DefaultsOmitEmpty{
			Policy: serialization.OmitEmptyPolicyWarn,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.RunWithSuggestedFixes(t, testdata, a, "d")
}

func TestOmitEmptyIgnore(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := defaults.Initializer().Init(&defaults.DefaultsConfig{
		OmitEmpty: defaults.DefaultsOmitEmpty{
			Policy: serialization.OmitEmptyPolicyIgnore,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.RunWithSuggestedFixes(t, testdata, a, "e")
}

func TestOmitZeroWarn(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := defaults.Initializer().Init(&defaults.DefaultsConfig{
		OmitZero: defaults.DefaultsOmitZero{
			Policy: serialization.OmitZeroPolicyWarn,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.RunWithSuggestedFixes(t, testdata, a, "f")
}
