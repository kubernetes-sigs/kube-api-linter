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
package duplicatemarkers

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"

	kalerrors "sigs.k8s.io/kube-api-linter/pkg/analysis/errors"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/extractjsontags"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/inspector"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/markers"
)

const (
	name = "duplicatemarkers"
)

// Analyzer is the analyzer for the duplicatemarkers package.
// It checks for duplicate markers on struct fields.
var Analyzer = &analysis.Analyzer{
	Name:     name,
	Doc:      "Check for duplicate markers on struct fields.",
	Run:      run,
	Requires: []*analysis.Analyzer{inspector.Analyzer},
}

func run(pass *analysis.Pass) (any, error) {
	inspect, ok := pass.ResultOf[inspector.Analyzer].(inspector.Inspector)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetInspector
	}

	inspect.InspectFields(func(field *ast.Field, _ []ast.Node, _ extractjsontags.FieldTagInfo, markersAccess markers.Markers) {
		if len(field.Names) == 0 {
			return
		}
		checkField(pass, field, markersAccess)
	})

	return nil, nil //nolint:nilnil
}

func checkField(pass *analysis.Pass, field *ast.Field, markersAccess markers.Markers) {
	set := markersAccess.FieldMarkers(field)

	fieldName := field.Names[0].Name

	for _, marker := range set.UnsortedList() {
		// TODO: Add check whether the marker is a duuplicate or not.
		pass.Report(analysis.Diagnostic{
			Pos:     field.Pos(),
			Message: fmt.Sprintf("%s has duplicated markers %s", fieldName, marker.String()),
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: fmt.Sprintf("should remove `// +%s`", marker.String()),
					TextEdits: []analysis.TextEdit{
						{
							Pos:     marker.Pos,
							End:     marker.End,
							NewText: nil,
						},
					},
				},
			},
		})
	}
}
