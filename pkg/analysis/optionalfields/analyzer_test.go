package optionalfields_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	requiredfields "sigs.k8s.io/kube-api-linter/pkg/analysis/optionalfields"
	"sigs.k8s.io/kube-api-linter/pkg/config"
)

func TestDefaultConfiguration(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := requiredfields.Initializer().Init(config.LintersConfig{})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.RunWithSuggestedFixes(t, testdata, a, "a")
}

func TestWhenRequiredPreferenceConfiguration(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := requiredfields.Initializer().Init(config.LintersConfig{
		OptionalFields: config.OptionalFieldsConfig{
			Pointers: config.OptionalFieldsPointers{
				Preference: config.OptionalFieldsPointerPreferenceWhenRequired,
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.RunWithSuggestedFixes(t, testdata, a, "b")
}
