package utils_test

import (
	"errors"
	"go/ast"
	"testing"

	"github.com/JoelSpeed/kal/pkg/analysis/utils"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
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
		Run: func(pass *analysis.Pass) (interface{}, error) {
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
