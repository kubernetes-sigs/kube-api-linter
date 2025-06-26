package forbiddenmarkers_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/forbiddenmarkers"
	"sigs.k8s.io/kube-api-linter/pkg/config"
)

func TestWithConfiguration(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, forbiddenmarkers.NewAnalyzer(config.ForbiddenMarkersConfig{
		Markers: []string{
			"forbidden",
		},
	}), "a/...")
}
