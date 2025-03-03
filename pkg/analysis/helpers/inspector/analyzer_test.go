package inspector_test

import (
	"errors"
	"go/ast"
	"testing"

	"github.com/JoelSpeed/kal/pkg/analysis/helpers/extractjsontags"
	"github.com/JoelSpeed/kal/pkg/analysis/helpers/inspector"
	"github.com/JoelSpeed/kal/pkg/analysis/helpers/markers"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestInspector(t *testing.T) {
	testdata := analysistest.TestData()

	analysistest.Run(t, testdata, testAnalyzer, "a")
}

var errCouldNotGetInspector = errors.New("could not get inspector")

var testAnalyzer = &analysis.Analyzer{
	Name:     "test",
	Doc:      "tests the inspector analyzer",
	Run:      run,
	Requires: []*analysis.Analyzer{inspector.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect, ok := pass.ResultOf[inspector.Analyzer].(inspector.Inspector)
	if !ok {
		return nil, errCouldNotGetInspector
	}

	inspect.InspectFields(func(field *ast.Field, stack []ast.Node, jsonTagInfo extractjsontags.FieldTagInfo, markersAccess markers.Markers) {
		var fieldName string
		if len(field.Names) > 0 {
			fieldName = field.Names[0].Name
		} else if ident, ok := field.Type.(*ast.Ident); ok {
			fieldName = ident.Name
		}

		pass.Reportf(field.Pos(), "field: %v", fieldName)

		if jsonTagInfo.Name != "" {
			pass.Reportf(field.Pos(), "json tag: %v", jsonTagInfo.Name)
		}
	})

	return nil, nil //nolint:nilnil
}
