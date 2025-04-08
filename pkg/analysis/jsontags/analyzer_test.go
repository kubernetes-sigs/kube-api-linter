package jsontags_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/jsontags"
	"sigs.k8s.io/kube-api-linter/pkg/config"
)

func TestDefaultConfiguration(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := jsontags.Initializer().Init(config.LintersConfig{})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.Run(t, testdata, a, "a")
}

func TestAlternativeRegex(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := jsontags.Initializer().Init(config.LintersConfig{
		JSONTags: config.JSONTagsConfig{
			JSONTagRegex: "^[a-z][a-z]*(?:[A-Z][a-z0-9]+)*[a-z0-9]?$",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.Run(t, testdata, a, "b")
}
