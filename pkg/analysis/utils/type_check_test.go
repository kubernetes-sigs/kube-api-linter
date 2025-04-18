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
	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils"
)

var (
	errCouldNotGetInspector = errors.New("could not get inspector")
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, testAnalyzer(), "a")
}

func testAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "test",
		Doc:      "test",
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Run: func(pass *analysis.Pass) (any, error) {
			inspect, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
			if !ok {
				return nil, errCouldNotGetInspector
			}

			// Filter to structs so that we can iterate over fields in a struct.
			nodeFilter := []ast.Node{
				(*ast.Field)(nil),
				(*ast.TypeSpec)(nil),
			}

			typeChecker := utils.NewTypeChecker(func(pass *analysis.Pass, ident *ast.Ident, node ast.Node, prefix string) {
				if ident.Name == "string" {
					pass.Reportf(node.Pos(), "%s is a string", prefix)
				}
			})

			inspect.Preorder(nodeFilter, func(n ast.Node) {
				typeChecker.CheckNode(pass, n)
			})

			return nil, nil
		},
	}
}
