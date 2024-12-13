package integers

import (
	"errors"
	"fmt"
	"go/ast"

	"github.com/JoelSpeed/kal/pkg/analysis/helpers/extractjsontags"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const name = "integers"

var (
	errCouldNotGetInspector = errors.New("could not get inspector")
)

// Analyzer is the analyzer for the integers package.
// It checks that no struct fields or type aliases are `int`, or unsigned integers.
var Analyzer = &analysis.Analyzer{
	Name:     name,
	Doc:      "All integers should be explicit about their size, int32 and int64 should be used over plain int. Unsigned ints are not allowed.",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer, extractjsontags.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, errCouldNotGetInspector
	}

	// Filter to structs so that we can iterate over fields in a struct.
	// Filter typespecs so that we can look at type aliases.
	nodeFilter := []ast.Node{
		(*ast.StructType)(nil),
		(*ast.TypeSpec)(nil),
	}

	// Preorder visits all the nodes of the AST in depth-first order. It calls
	// f(n) for each node n before it visits n's children.
	//
	// We use the filter defined above, ensuring we only look at struct fields and type declarations.
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch typ := n.(type) {
		case *ast.StructType:
			if typ.Fields == nil {
				return
			}

			for _, field := range typ.Fields.List {
				checkField(pass, field)
			}
		case *ast.TypeSpec:
			checkTypeSpec(pass, typ, typ, "type")
		}
	})

	return nil, nil //nolint:nilnil
}

func checkField(pass *analysis.Pass, field *ast.Field) {
	if field == nil || len(field.Names) == 0 || field.Names[0] == nil {
		return
	}

	fieldName := field.Names[0].Name
	prefix := fmt.Sprintf("field %s", fieldName)

	checkTypeExpr(pass, field.Type, field, prefix)
}

func checkTypeSpec(pass *analysis.Pass, tSpec *ast.TypeSpec, node ast.Node, prefix string) {
	if tSpec.Name == nil {
		return
	}

	typeName := tSpec.Name.Name
	prefix = fmt.Sprintf("%s %s", prefix, typeName)

	checkTypeExpr(pass, tSpec.Type, node, prefix)
}

func checkTypeExpr(pass *analysis.Pass, typeExpr ast.Expr, node ast.Node, prefix string) {
	switch typ := typeExpr.(type) {
	case *ast.Ident:
		checkIdent(pass, typ, node, prefix)
	case *ast.StarExpr:
		checkTypeExpr(pass, typ.X, node, fmt.Sprintf("%s pointer", prefix))
	case *ast.ArrayType:
		checkTypeExpr(pass, typ.Elt, node, fmt.Sprintf("%s array element", prefix))
	}
}

// checkIdent looks for known type of integers that do not match the allowed `int32` or `int64` requirements.
// It will also identify fields that are using type aliases and determine if these type aliases use invalid
// types of integers, and highlight this.
func checkIdent(pass *analysis.Pass, ident *ast.Ident, node ast.Node, prefix string) {
	switch ident.Name {
	case "int32", "int64":
		// Valid cases
	case "int", "int8", "int16":
		pass.Reportf(node.Pos(), "%s should not use an int, int8 or int16. Use int32 or int64 depending on bounding requirements", prefix)
	case "uint", "uint8", "uint16", "uint32", "uint64":
		pass.Reportf(node.Pos(), "%s should not use unsigned integers, use only int32 or int64 and apply validation to ensure the value is positive", prefix)
	}

	if ident.Obj == nil || ident.Obj.Decl == nil {
		return
	}

	tSpec, ok := ident.Obj.Decl.(*ast.TypeSpec)
	if !ok {
		return
	}

	// The field is using a type alias, check if the alias is an int.
	checkTypeSpec(pass, tSpec, node, fmt.Sprintf("%s type", prefix))
}
