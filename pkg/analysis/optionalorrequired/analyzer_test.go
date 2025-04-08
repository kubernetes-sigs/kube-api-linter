package optionalorrequired_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/optionalorrequired"
	"sigs.k8s.io/kube-api-linter/pkg/config"
)

func TestDefaultConfiguration(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := optionalorrequired.Initializer().Init(config.LintersConfig{})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.RunWithSuggestedFixes(t, testdata, a, "a")
}

func TestSwappedMarkerPriority(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := optionalorrequired.Initializer().Init(config.LintersConfig{
		OptionalOrRequired: config.OptionalOrRequiredConfig{
			PreferredOptionalMarker: optionalorrequired.KubebuilderOptionalMarker,
			PreferredRequiredMarker: optionalorrequired.KubebuilderRequiredMarker,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.RunWithSuggestedFixes(t, testdata, a, "b")
}
