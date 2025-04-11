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
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	kalerrors "sigs.k8s.io/kube-api-linter/pkg/analysis/errors"
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
	Requires: []*analysis.Analyzer{inspect.Analyzer, markers.Analyzer},
}

func run(pass *analysis.Pass) (any, error) {
	inspect, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetInspector
	}

	markersAccess, ok := pass.ResultOf[markers.Analyzer].(markers.Markers)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetMarkers
	}

	nodeFilter := []ast.Node{
		(*ast.Field)(nil),
		(*ast.TypeSpec)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.Field:
			checkField(pass, n, markersAccess)
		case *ast.TypeSpec:
			checkTypeSpec(pass, n, markersAccess)
		}
	})

	return nil, nil //nolint:nilnil
}

func checkField(pass *analysis.Pass, field *ast.Field, markersAccess markers.Markers) {
	set := markersAccess.FieldMarkers(field)

	fieldName := field.Names[0].Name

	countMap := make(map[string]int)
	for _, marker := range set.UnsortedList() {
		countMap[marker.RawComment]++
	}

	for _, marker := range set.UnsortedList() {
		if countMap[marker.RawComment] > 1 {
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

			countMap[marker.RawComment]--
		}
	}
}

func checkTypeSpec(pass *analysis.Pass, tSpec *ast.TypeSpec, markersAccess markers.Markers) {
	name := tSpec.Name.Name

	set := markersAccess.TypeMarkers(tSpec)

	countMap := make(map[string]int)
	for _, marker := range set.UnsortedList() {
		countMap[marker.RawComment]++
	}

	for _, marker := range set.UnsortedList() {
		if countMap[marker.RawComment] > 1 {
			pass.Report(analysis.Diagnostic{
				Pos:     tSpec.Pos(),
				Message: fmt.Sprintf("%s has duplicated markers %s", name, marker.String()),
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

			countMap[marker.RawComment]--
		}
	}
}
