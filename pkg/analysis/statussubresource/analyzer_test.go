package statussubresource_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/statussubresource"
	"sigs.k8s.io/kube-api-linter/pkg/config"
)

func TestStatusSubresourceAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	initializer := statussubresource.Initializer()

	analyzer, err := initializer.Init(config.LintersConfig{})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.RunWithSuggestedFixes(t, testdata, analyzer, "a")
}
