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
package nomaps_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/nomaps"
	"sigs.k8s.io/kube-api-linter/pkg/config"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := nomaps.Initializer().Init(config.LintersConfig{})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.Run(t, testdata, a, "a")
}

func TestWithEnforce(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := nomaps.Initializer().Init(config.LintersConfig{
		NoMaps: config.NoMapsConfig{
			Policy: config.NoMapsEnforce,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.Run(t, testdata, a, "b")
}

func TestWithAllowStringToStringMaps(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := nomaps.Initializer().Init(config.LintersConfig{
		NoMaps: config.NoMapsConfig{
			Policy: config.NoMapsAllowStringToStringMaps,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.Run(t, testdata, a, "c")
}

func TestWithIgnore(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := nomaps.Initializer().Init(config.LintersConfig{
		NoMaps: config.NoMapsConfig{
			Policy: config.NoMapsIgnore,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.Run(t, testdata, a, "d")
}
