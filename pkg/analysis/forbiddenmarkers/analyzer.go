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
	"cmp"
	"fmt"
	"go/ast"
	"slices"

	"golang.org/x/tools/go/analysis"
	kalerrors "sigs.k8s.io/kube-api-linter/pkg/analysis/errors"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/inspector"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/markers"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils"
)

const name = "forbiddenmarkers"

type analyzer struct {
	forbiddenMarkers []Marker
}

// NewAnalyzer creates a new analysis.Analyzer for the forbiddenmarkers
// linter based on the provided config.ForbiddenMarkersConfig.
func newAnalyzer(cfg *Config) *analysis.Analyzer {
	a := &analyzer{
		forbiddenMarkers: cfg.Markers,
	}

	analyzer := &analysis.Analyzer{
		Name:     name,
		Doc:      "Check that no forbidden markers are present on types and fields.",
		Run:      a.run,
		Requires: []*analysis.Analyzer{inspector.Analyzer},
	}

	for _, marker := range a.forbiddenMarkers {
		markers.DefaultRegistry().Register(marker.Identifier)
	}

	return analyzer
}

func (a *analyzer) run(pass *analysis.Pass) (any, error) {
	inspect, ok := pass.ResultOf[inspector.Analyzer].(inspector.Inspector)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetInspector
	}

	for f := range inspect.Fields() {
		checkField(pass, f.Field, f.Markers, a.forbiddenMarkers, f.QualifiedFieldName)
	}

	for ts := range inspect.TypeSpecs() {
		checkType(pass, ts.TypeSpec, ts.Markers, a.forbiddenMarkers)
	}

	return nil, nil //nolint:nilnil
}

func checkField(pass *analysis.Pass, field *ast.Field, markersAccess markers.Markers, forbiddenMarkers []Marker, qualifiedFieldName string) {
	if field == nil || len(field.Names) == 0 {
		return
	}

	markers := utils.TypeAwareMarkerCollectionForField(pass, markersAccess, field)
	check(markers, forbiddenMarkers, reportField(pass, field, qualifiedFieldName))
}

func checkType(pass *analysis.Pass, typeSpec *ast.TypeSpec, markersAccess markers.Markers, forbiddenMarkers []Marker) {
	if typeSpec == nil {
		return
	}

	markers := markersAccess.TypeMarkers(typeSpec)
	check(markers, forbiddenMarkers, reportType(pass, typeSpec))
}

func check(markerSet markers.MarkerSet, forbiddenMarkers []Marker, reportFunc func(marker markers.Marker)) {
	for _, marker := range forbiddenMarkers {
		marks := markersWithIdentifier(markerSet, marker.Identifier)
		for _, mark := range marks {
			if len(marker.RuleSets) == 0 {
				reportFunc(mark)
				continue
			}

			for _, ruleSet := range marker.RuleSets {
				if markerMatchesAttributeRules(mark, ruleSet.Attributes...) {
					reportFunc(mark)
				}
			}
		}
	}
}

// markersWithIdentifier returns all markers in the marker set with the given identifier,
// including markers nested within a declarative validation chain.
//
// Declarative validation markers may wrap other markers, for example
// `+k8s:alpha(since: "1.36")=+k8s:customValidation`. The MarkerSet indexes markers
// only by their outermost identifier (k8s:alpha in this example), with the nested
// marker stored in Payload.Marker. To ensure forbidden markers cannot be bypassed by
// nesting them, each marker chain is walked and every link in the chain is compared
// against the identifier.
//
// Nested markers share the Pos, End and RawComment of the comment they were parsed from,
// so reporting a nested marker removes the whole comment line.
//
// The returned markers are sorted by position so that diagnostics are deterministic,
// as nested matches may originate from markers stored under different identifiers.
func markersWithIdentifier(markerSet markers.MarkerSet, identifier string) []markers.Marker {
	matches := []markers.Marker{}

	for _, topLevel := range markerSet.UnsortedList() {
		for m := &topLevel; m != nil; m = m.Payload.Marker {
			if m.Identifier == identifier {
				matches = append(matches, *m)
			}
		}
	}

	slices.SortStableFunc(matches, func(a, b markers.Marker) int {
		return cmp.Compare(a.Pos, b.Pos)
	})

	return matches
}

func markerMatchesAttributeRules(marker markers.Marker, attrRules ...MarkerAttribute) bool {
	for _, attrRule := range attrRules {
		// if the marker doesn't contain the attribute for a specified rule it fails the AND
		// operation.
		val, ok := marker.Arguments[attrRule.Name]
		if !ok {
			return false
		}

		// if the value doesn't match one of the forbidden ones, this marker is not forbidden
		if len(attrRule.Values) > 0 && !slices.Contains(attrRule.Values, val) {
			return false
		}
	}

	return true
}

func reportField(pass *analysis.Pass, field *ast.Field, qualifiedFieldName string) func(marker markers.Marker) {
	return func(marker markers.Marker) {
		pass.Report(analysis.Diagnostic{
			Pos:     field.Pos(),
			Message: fmt.Sprintf("field %s has forbidden marker %q", qualifiedFieldName, marker.String()),
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: fmt.Sprintf("remove forbidden marker %q", marker.String()),
					TextEdits: []analysis.TextEdit{
						{
							Pos: marker.Pos,
							End: marker.End,
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
			Message: fmt.Sprintf("type %s has forbidden marker %q", typeSpec.Name, marker.String()),
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: fmt.Sprintf("remove forbidden marker %q", marker.String()),
					TextEdits: []analysis.TextEdit{
						{
							Pos: marker.Pos,
							End: marker.End,
						},
					},
				},
			},
		})
	}
}
