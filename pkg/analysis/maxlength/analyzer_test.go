package maxlength_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/maxlength"
)

func TestMaxLength(t *testing.T) {
	testdata := analysistest.TestData()

	analysistest.Run(t, testdata, maxlength.Analyzer, "a")
}
