package main

import (
	"golang.org/x/tools/go/analysis/multichecker"
	kalanalysis "sigs.k8s.io/kube-api-linter/pkg/analysis"
	"sigs.k8s.io/kube-api-linter/pkg/config"
)

func main() {
	analyzers, err := kalanalysis.NewRegistry().InitializeLinters(config.Linters{}, config.LintersConfig{})
	if err != nil {
		panic(err)
	}

	multichecker.Main(
		analyzers...,
	)
}
