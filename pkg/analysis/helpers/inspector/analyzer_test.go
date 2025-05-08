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
package inspector_test

import (
	"errors"
	"go/ast"
	"testing"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/extractjsontags"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/inspector"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/markers"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils"
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

func run(pass *analysis.Pass) (any, error) {
	inspect, ok := pass.ResultOf[inspector.Analyzer].(inspector.Inspector)
	if !ok {
		return nil, errCouldNotGetInspector
	}

	inspect.InspectFields(func(field *ast.Field, stack []ast.Node, jsonTagInfo extractjsontags.FieldTagInfo, markersAccess markers.Markers) {
		pass.Reportf(field.Pos(), "field: %v", utils.FieldName(field))

		if jsonTagInfo.Name != "" {
			pass.Reportf(field.Pos(), "json tag: %v", jsonTagInfo.Name)
		}
	})

	return nil, nil //nolint:nilnil
}
