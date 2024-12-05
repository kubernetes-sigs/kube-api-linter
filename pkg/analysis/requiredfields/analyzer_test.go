package requiredfields_test

import (
	"testing"

	"github.com/JoelSpeed/kal/pkg/analysis/requiredfields"
	"github.com/JoelSpeed/kal/pkg/config"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestDefaultConfiguration(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := requiredfields.Initializer().Init(config.LintersConfig{})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.RunWithSuggestedFixes(t, testdata, a, "a")
}
