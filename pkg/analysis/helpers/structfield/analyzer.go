package structfield

import (
	"errors"
	"go/ast"
	"reflect"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var (
	errCouldNotGetInspector      = errors.New("could not get inspector")
	errCouldNotCreateStructField = errors.New("could not create new structField")
)

// StructField is used to determine if an *ast.Field belongs to a struct.
type StructField interface {
	StructForField(*ast.Field) *ast.StructType
}

type structField struct {
	fieldToStruct map[*ast.Field]*ast.StructType
}

func newStructField() StructField {
	return &structField{
		fieldToStruct: make(map[*ast.Field]*ast.StructType),
	}
}

// StructForField returns the struct that the field belongs to.
func (s *structField) StructForField(field *ast.Field) *ast.StructType {
	return s.fieldToStruct[field]
}

// Analyzer is the analyzer for the structfield package.
// It creates a way to check if a particular *ast.Field belongs to a struct.
var Analyzer = &analysis.Analyzer{
	Name:       "strucfield",
	Doc:        "Iterates over all fields in structs and creates a mapping of the field to the struct.",
	Run:        run,
	Requires:   []*analysis.Analyzer{inspect.Analyzer},
	ResultType: reflect.TypeOf(newStructField()),
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, errCouldNotGetInspector
	}

	// Filter to structs so that we can iterate over fields in a struct.
	nodeFilter := []ast.Node{
		(*ast.StructType)(nil),
	}

	results, ok := newStructField().(*structField)
	if !ok {
		return nil, errCouldNotCreateStructField
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		sTyp, ok := n.(*ast.StructType)
		if !ok {
			return
		}

		if sTyp.Fields == nil {
			return
		}

		for _, field := range sTyp.Fields.List {
			results.fieldToStruct[field] = sTyp
		}
	})

	return results, nil
}
