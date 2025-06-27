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
	forbiddenMarkers []config.ForbiddenMarker
}

// ForbiddenMarkersOptions is a function that configures the
// forbiddenmarkers analysis.Analyzer
type ForbiddenMarkersOption func(a *analysis.Analyzer)

// WithName sets the name of the forbiddenmarkers analysis.Analyzer
func WithName(name string) ForbiddenMarkersOption {
	return func(a *analysis.Analyzer) {
		a.Name = name
	}
}

// WithDoc sets the doc string of the forbiddenmarkers analysis.Analyzer
func WithDoc(doc string) ForbiddenMarkersOption {
	return func(a *analysis.Analyzer) {
		a.Doc = doc
	}
}

// NewAnalyzer creates a new analysis.Analyzer for the forbiddenmarkers
// linter based on the provided config.ForbiddenMarkersConfig.
func NewAnalyzer(cfg config.ForbiddenMarkersConfig, opts ...ForbiddenMarkersOption) *analysis.Analyzer {
	a := &analyzer{
		forbiddenMarkers: cfg.Markers,
	}

	analyzer := &analysis.Analyzer{
		Name:     name,
		Doc:      "Check that no forbidden markers are present on types and fields.",
		Run:      a.run,
		Requires: []*analysis.Analyzer{inspector.Analyzer},
	}

	for _, opt := range opts {
		opt(analyzer)
	}

	return analyzer
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

func checkField(pass *analysis.Pass, field *ast.Field, markersAccess markers.Markers, forbiddenMarkers []config.ForbiddenMarker) {
	if field == nil || len(field.Names) == 0 {
		return
	}

	markers := utils.TypeAwareMarkerCollectionForField(pass, markersAccess, field)
	check(markers, forbiddenMarkers, reportField(pass, field))
}

func checkType(pass *analysis.Pass, typeSpec *ast.TypeSpec, markersAccess markers.Markers, forbiddenMarkers []config.ForbiddenMarker) {
	if typeSpec == nil {
		return
	}

	markers := markersAccess.TypeMarkers(typeSpec)
	check(markers, forbiddenMarkers, reportType(pass, typeSpec))
}

func check(markerSet markers.MarkerSet, forbiddenMarkers []config.ForbiddenMarker, reportFunc func(marker markers.Marker)) {
	for _, marker := range forbiddenMarkers {
		marks := markerSet.Get(marker.Identifier)
		for _, mark := range marks {
			if markerMatchesAttributeRules(mark, marker.Attributes...) {
				reportFunc(mark)
			}
		}
	}
}

// TODO: this should probably return some representation of the marker that is failing the
// attribute rules so that it can be returned to users helpfully.
func markerMatchesAttributeRules(marker markers.Marker, attrRules ...config.ForbiddenMarkerAttribute) bool {
	matchesAll := true
	for _, attrRule := range attrRules {
		if val, ok := marker.Expressions[attrRule.Attribute]; ok {
			// if no values are specified, that means the existence match is enough
			// and we can continue to the next rule
			if len(attrRule.Values) == 0 {
				continue
			}

			// if the value doesn't match one of the forbidden ones, this marker is not forbidden
			matchesOneValue := false
			for _, value := range attrRule.Values {
				if val == value {
					matchesOneValue = true
					break
				}
			}

			if !matchesOneValue {
				matchesAll = false
				break
			}
		}
		// if the marker doesn't contain the attribute for a specified rule it fails the AND
		// operation.
		matchesAll = false
		break
	}

	return matchesAll
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
