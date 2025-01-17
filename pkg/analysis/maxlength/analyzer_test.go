package maxlength_test

import (
	"testing"

	"github.com/JoelSpeed/kal/pkg/analysis/maxlength"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestMaxLength(t *testing.T) {
	testdata := analysistest.TestData()

	analysistest.Run(t, testdata, maxlength.Analyzer, "a")
}
