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
