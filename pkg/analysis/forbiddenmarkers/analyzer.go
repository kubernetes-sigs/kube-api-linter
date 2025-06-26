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
package forbiddenmarkers

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	kalerrors "sigs.k8s.io/kube-api-linter/pkg/analysis/errors"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/extractjsontags"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/inspector"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/markers"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils"
	"sigs.k8s.io/kube-api-linter/pkg/config"
)

const name = "forbiddenmarkers"

type analyzer struct {
	forbiddenMarkers []string
}

// NewAnalyzer creates a new analysis.Analyzer for the forbiddenmarkers
// linter based on the provided config.ForbiddenMarkersConfig.
func NewAnalyzer(cfg config.ForbiddenMarkersConfig) *analysis.Analyzer {
	a := &analyzer{
		forbiddenMarkers: cfg.Markers,
	}

	return &analysis.Analyzer{
		Name:     name,
		Doc:      "Check that no forbidden markers are present on types and fields.",
		Run:      a.run,
		Requires: []*analysis.Analyzer{inspector.Analyzer},
	}
}

func (a *analyzer) run(pass *analysis.Pass) (any, error) {
	inspect, ok := pass.ResultOf[inspector.Analyzer].(inspector.Inspector)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetInspector
	}

	inspect.InspectFields(func(field *ast.Field, stack []ast.Node, _ extractjsontags.FieldTagInfo, markersAccess markers.Markers) {
		checkField(pass, field, markersAccess, a.forbiddenMarkers)
	})

	inspect.InspectTypeSpec(func(typeSpec *ast.TypeSpec, markersAccess markers.Markers) {
		checkType(pass, typeSpec, markersAccess, a.forbiddenMarkers)
	})

	return nil, nil //nolint:nilnil
}

func checkField(pass *analysis.Pass, field *ast.Field, markersAccess markers.Markers, forbiddenMarkers []string) {
	if field == nil || len(field.Names) == 0 {
		return
	}

	markers := utils.TypeAwareMarkerCollectionForField(pass, markersAccess, field)
	check(markers, forbiddenMarkers, reportField(pass, field))
}

func checkType(pass *analysis.Pass, typeSpec *ast.TypeSpec, markersAccess markers.Markers, forbiddenMarkers []string) {
	if typeSpec == nil {
		return
	}

	markers := markersAccess.TypeMarkers(typeSpec)
	check(markers, forbiddenMarkers, reportType(pass, typeSpec))
}

func check(markerSet markers.MarkerSet, forbiddenMarkers []string, reportFunc func(marker markers.Marker)) {
	for _, marker := range forbiddenMarkers {
		marks := markerSet.Get(marker)
		for _, mark := range marks {
			reportFunc(mark)
		}
	}
}

func reportField(pass *analysis.Pass, field *ast.Field) func(marker markers.Marker) {
	return func(marker markers.Marker) {
		pass.Report(analysis.Diagnostic{
			Pos:     field.Pos(),
			Message: fmt.Sprintf("field %s has forbidden marker %q", field.Names[0].Name, marker.Identifier),
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: fmt.Sprintf("remove forbidden marker %q", marker.Identifier),
					TextEdits: []analysis.TextEdit{
						{
							Pos: marker.Pos,
							End: marker.End + 1,
						},
					},
				},
			},
		})
	}
}

func reportType(pass *analysis.Pass, typeSpec *ast.TypeSpec) func(marker markers.Marker) {
	return func(marker markers.Marker) {
		pass.Report(analysis.Diagnostic{
			Pos:     typeSpec.Pos(),
			Message: fmt.Sprintf("type %s has forbidden marker %q", typeSpec.Name, marker.Identifier),
		})
	}
}
