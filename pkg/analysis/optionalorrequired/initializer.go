package optionalorrequired

import (
	"golang.org/x/tools/go/analysis"
	"sigs.k8s.io/kube-api-linter/pkg/config"
)

// Initializer returns the AnalyzerInitializer for this
// Analyzer so that it can be added to the registry.
func Initializer() initializer {
	return initializer{}
}

// intializer implements the AnalyzerInitializer interface.
type initializer struct{}

// Name returns the name of the Analyzer.
func (initializer) Name() string {
	return name
}

// Init returns the intialized Analyzer.
func (initializer) Init(cfg config.LintersConfig) (*analysis.Analyzer, error) {
	return newAnalyzer(cfg.OptionalOrRequired), nil
}

// Default determines whether this Analyzer is on by default, or not.
func (initializer) Default() bool {
	return true
}
