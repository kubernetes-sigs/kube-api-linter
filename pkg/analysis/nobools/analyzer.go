package nobools

import (
	"errors"
	"go/ast"

	"github.com/JoelSpeed/kal/pkg/analysis/helpers/structfield"
	"github.com/JoelSpeed/kal/pkg/analysis/utils"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const name = "nobools"

var (
	errCouldNotGetInspector   = errors.New("could not get inspector")
	errCouldNotGetStructField = errors.New("could not get struct field")
)

// Analyzer is the analyzer for the nobools package.
// It checks that no struct fields are `bool`.
var Analyzer = &analysis.Analyzer{
	Name:     name,
	Doc:      "Boolean values cannot evolve over time, use an enum with meaningful values instead",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer, structfield.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, errCouldNotGetInspector
	}

	structField, ok := pass.ResultOf[structfield.Analyzer].(structfield.StructField)
	if !ok {
		return nil, errCouldNotGetStructField
	}

	// Filter to fields so that we can look at fields within structs.
	// Filter typespecs so that we can look at type aliases.
	nodeFilter := []ast.Node{
		(*ast.Field)(nil),
		(*ast.TypeSpec)(nil),
	}

	typeChecker := utils.NewTypeChecker(checkBool)

	// Preorder visits all the nodes of the AST in depth-first order. It calls
	// f(n) for each node n before it visits n's children.
	//
	// We use the filter defined above, ensuring we only look at struct fields and type declarations.
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		if field, ok := n.(*ast.Field); ok {
			// Do not inspect fields that are not part of a struct.
			if structType := structField.StructForField(field); structType == nil {
				return
			}
		}

		typeChecker.CheckNode(pass, n)
	})

	return nil, nil //nolint:nilnil
}

func checkBool(pass *analysis.Pass, ident *ast.Ident, node ast.Node, prefix string) {
	if ident.Name == "bool" {
		pass.Reportf(node.Pos(), "%s should not use a bool. Use a string type with meaningful constant values as an enum.", prefix)
	}
}
