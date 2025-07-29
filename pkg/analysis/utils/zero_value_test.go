/*
Copyright 2025 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package utils_test

import (
	"errors"
	"go/ast"
	"testing"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/extractjsontags"
	markershelper "sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/markers"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils"
)

var (
	errCouldNotGetMarkers = errors.New("could not get markers")
)

func TestZeroValueWithoutOmitZero(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, testZeroValueAnalyzer(false), "b")
}

func TestZeroValueWithOmitZero(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, testZeroValueAnalyzer(true), "c")
}

func testZeroValueAnalyzer(considerOmitZero bool) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "test",
		Doc:      "test",
		Requires: []*analysis.Analyzer{inspect.Analyzer, markershelper.Analyzer, extractjsontags.Analyzer},
		Run: func(pass *analysis.Pass) (any, error) {
			inspect, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
			if !ok {
				return nil, errCouldNotGetInspector
			}

			markers, ok := pass.ResultOf[markershelper.Analyzer].(markershelper.Markers)
			if !ok {
				return nil, errCouldNotGetMarkers
			}

			// Filter to structs so that we can iterate over fields in a struct.
			nodeFilter := []ast.Node{
				(*ast.Field)(nil),
			}

			inspect.Preorder(nodeFilter, func(n ast.Node) {
				field, ok := n.(*ast.Field)
				if !ok {
					return
				}

				zeroValueValid, complete := utils.IsZeroValueValid(pass, field, field.Type, markers, considerOmitZero)
				if !zeroValueValid {
					pass.Reportf(field.Pos(), "zero value is not valid")
				} else {
					pass.Reportf(field.Pos(), "zero value is valid")
				}
				if !complete {
					pass.Reportf(field.Pos(), "validation is not complete")
				} else {
					pass.Reportf(field.Pos(), "validation is complete")
				}
			})

			return nil, nil
		},
	}
}
