package nobools_test

import (
	"testing"

	"github.com/JoelSpeed/kal/pkg/analysis/nobools"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, nobools.Analyzer, "a")
}
