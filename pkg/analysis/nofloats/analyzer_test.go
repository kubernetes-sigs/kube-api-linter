package nofloats_test

import (
	"testing"

	"github.com/JoelSpeed/kal/pkg/analysis/nofloats"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, nofloats.Analyzer, "a")
}
