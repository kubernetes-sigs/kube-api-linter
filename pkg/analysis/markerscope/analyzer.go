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
package markerscope

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	kalerrors "sigs.k8s.io/kube-api-linter/pkg/analysis/errors"
	markershelper "sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/markers"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils"
)

const (
	name = "markerscope"
)

func init() {
	// Register all markers we want to validate scope for
	defaults := DefaultMarkerScopes()
	markers := make([]string, 0, len(defaults))
	for marker := range defaults {
		markers = append(markers, marker)
	}
	markershelper.DefaultRegistry().Register(markers...)
}

type analyzer struct {
	markerScopes map[string]MarkerScope
	policy       MarkerScopePolicy
}

// newAnalyzer creates a new analyzer.
func newAnalyzer(cfg *MarkerScopeConfig) *analysis.Analyzer {
	if cfg == nil {
		cfg = &MarkerScopeConfig{}
	}

	a := &analyzer{
		markerScopes: DefaultMarkerScopes(),
		policy:       cfg.Policy,
	}

	// Set default policy if not specified
	if a.policy == "" {
		a.policy = MarkerScopePolicySuggestFix
	}

	return &analysis.Analyzer{
		Name:             name,
		Doc:              "Validates that markers are applied in the correct scope.",
		Run:              a.run,
		Requires:         []*analysis.Analyzer{inspect.Analyzer, markershelper.Analyzer},
		RunDespiteErrors: true,
	}
}

func (a *analyzer) run(pass *analysis.Pass) (any, error) {
	inspect, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetInspector
	}

	markersAccess, ok := pass.ResultOf[markershelper.Analyzer].(markershelper.Markers)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetMarkers
	}

	// Check field markers and type markers
	nodeFilter := []ast.Node{
		(*ast.Field)(nil),
		(*ast.GenDecl)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch node := n.(type) {
		case *ast.Field:
			a.checkFieldMarkers(pass, node, markersAccess)
		case *ast.GenDecl:
			a.checkTypeMarkers(pass, node, markersAccess)
		}
	})

	return nil, nil
}

// reportAndRemoveMarker reports an error and suggests removing the marker
func (a *analyzer) reportAndRemoveMarker(pass *analysis.Pass, marker markershelper.Marker, allowedScope string) {
	suggestedFixes := []analysis.SuggestedFix{}

	// Only add suggested fixes if policy allows it
	if a.policy == MarkerScopePolicySuggestFix {
		suggestedFixes = []analysis.SuggestedFix{
			{
				Message: fmt.Sprintf("Remove marker %q", marker.Identifier),
				TextEdits: []analysis.TextEdit{
					a.removeMarker(marker),
				},
			},
		}
	}

	pass.Report(analysis.Diagnostic{
		Pos:            marker.Pos,
		End:            marker.End,
		Message:        fmt.Sprintf("marker %q can only be applied to %s", marker.Identifier, allowedScope),
		SuggestedFixes: suggestedFixes,
	})
}

// reportAndMoveMarkerToType reports an error and suggests moving the marker to the type definition
func (a *analyzer) reportAndMoveMarkerToType(pass *analysis.Pass, marker markershelper.Marker, typeSpec *ast.TypeSpec) {
	suggestedFixes := []analysis.SuggestedFix{}

	// Only add suggested fixes if policy allows it
	if a.policy == MarkerScopePolicySuggestFix {
		insertPos := a.getTypeInsertPos(pass, typeSpec)
		if insertPos != token.NoPos {
			suggestedFixes = []analysis.SuggestedFix{
				{
					Message: fmt.Sprintf("Move marker %q to type %s", marker.Identifier, typeSpec.Name.Name),
					TextEdits: []analysis.TextEdit{
						// Remove from field
						a.removeMarker(marker),
						// Add to struct type
						{
							Pos:     insertPos,
							End:     insertPos,
							NewText: []byte(fmt.Sprintf("// +%s\n", cleanMarkerString(marker))),
						},
					},
				},
				{
					Message: fmt.Sprintf("Remove marker %q", marker.Identifier),
					TextEdits: []analysis.TextEdit{
						a.removeMarker(marker),
					},
				},
			}
		}
	}

	pass.Report(analysis.Diagnostic{
		Pos:            marker.Pos,
		End:            marker.End,
		Message:        fmt.Sprintf("marker %q can only be applied to type definitions or object fields (struct/map)", marker.Identifier),
		SuggestedFixes: suggestedFixes,
	})
}

// reportAndMoveMarkerToField reports an error and suggests moving the marker to a field
func (a *analyzer) reportAndMoveMarkerToField(pass *analysis.Pass, marker markershelper.Marker, targetField *ast.Field) {
	suggestedFixes := []analysis.SuggestedFix{}

	// Only add suggested fixes if policy allows it
	if a.policy == MarkerScopePolicySuggestFix {
		insertPos := a.getFieldInsertPos(targetField)
		fieldName := utils.FieldName(targetField)
		if fieldName == "" {
			fieldName = "field"
		}

		suggestedFixes = []analysis.SuggestedFix{
			{
				Message: fmt.Sprintf("Move marker %q to field %s", marker.Identifier, fieldName),
				TextEdits: []analysis.TextEdit{
					// Remove from type
					a.removeMarker(marker),
					// Add to field
					{
						Pos:     insertPos,
						End:     insertPos,
						NewText: []byte(fmt.Sprintf("// +%s\n\t", cleanMarkerString(marker))),
					},
				},
			},
			{
				Message: fmt.Sprintf("Remove marker %q", marker.Identifier),
				TextEdits: []analysis.TextEdit{
					a.removeMarker(marker),
				},
			},
		}
	}

	pass.Report(analysis.Diagnostic{
		Pos:            marker.Pos,
		End:            marker.End,
		Message:        fmt.Sprintf("marker %q can only be applied to fields", marker.Identifier),
		SuggestedFixes: suggestedFixes,
	})
}

// checkFieldMarkers checks markers on fields for violations
func (a *analyzer) checkFieldMarkers(pass *analysis.Pass, field *ast.Field, markersAccess markershelper.Markers) {
	fieldMarkers := markersAccess.FieldMarkers(field)

	for _, marker := range fieldMarkers.UnsortedList() {
		scope, ok := a.markerScopes[marker.Identifier]
		if !ok {
			// No scope defined for this marker, skip validation
			continue
		}

		switch scope {
		case ScopeType:
			a.reportAndRemoveMarker(pass, marker, "type definitions")
		case ScopeTypeOrObjectField:
			// Check if field type is an object type (struct or map) 
			if !isMapType(field.Type) && !utils.IsStructType(pass, field.Type) {
				// Check if it's a struct type and we can move the marker there
				typeSpec := a.getStructTypeSpec(pass, field.Type)

				if typeSpec != nil {
					// We can move the marker to the struct definition
					a.reportAndMoveMarkerToType(pass, marker, typeSpec)
				} else {
					// Can't find struct type, just offer removal
					a.reportAndRemoveMarker(pass, marker, "type definitions or object fields (struct/map)")
				}
			}
		}
	}
}

// checkTypeMarkers checks markers on types for violations
func (a *analyzer) checkTypeMarkers(pass *analysis.Pass, genDecl *ast.GenDecl, markersAccess markershelper.Markers) {
	if len(genDecl.Specs) == 0 {
		return
	}

	for i := range genDecl.Specs {
		typeSpec, ok := genDecl.Specs[i].(*ast.TypeSpec)
		if !ok {
			continue
		}

		typeMarkers := markersAccess.TypeMarkers(typeSpec)

		for _, marker := range typeMarkers.UnsortedList() {
			scope, ok := a.markerScopes[marker.Identifier]
			if !ok {
				// No scope defined for this marker, skip validation
				continue
			}

			if scope == ScopeField {
				// For field-only markers on types, try to move to a field using this type
				targetField := a.findFieldUsingType(pass, typeSpec.Name.Name)
				if targetField != nil {
					a.reportAndMoveMarkerToField(pass, marker, targetField)
				} else {
					// No suitable field found, just suggest removal
					a.reportAndRemoveMarker(pass, marker, "fields")
				}
			}
		}
	}
}

// isMapType checks if the given expression is a map type
func isMapType(expr ast.Expr) bool {
	_, ok := expr.(*ast.MapType)
	return ok
}

// getStructTypeSpec finds the type spec for a given field type if it's a struct
func (a *analyzer) getStructTypeSpec(pass *analysis.Pass, expr ast.Expr) *ast.TypeSpec {
	// Check if it's a struct type first
	if !utils.IsStructType(pass, expr) {
		return nil
	}

	// Get the identifier from the expression
	var ident *ast.Ident
	switch t := expr.(type) {
	case *ast.Ident:
		ident = t
	case *ast.StarExpr:
		// Handle pointer types
		if identExpr, ok := t.X.(*ast.Ident); ok {
			ident = identExpr
		}
	default:
		return nil
	}

	if ident == nil {
		return nil
	}

	// Use utils.LookupTypeSpec to find the type declaration
	typeSpec, ok := utils.LookupTypeSpec(pass, ident)
	if !ok {
		return nil
	}

	return typeSpec
}

// findFieldUsingType finds a field that uses the given type name
func (a *analyzer) findFieldUsingType(pass *analysis.Pass, typeName string) *ast.Field {
	inspect, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil
	}

	var result *ast.Field

	// Find field nodes that use this type
	nodeFilter := []ast.Node{
		(*ast.Field)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		field, ok := n.(*ast.Field)
		if !ok || result != nil {
			return
		}

		// Check if this field uses the type we're looking for
		fieldTypeName := getFieldTypeName(field.Type)
		if fieldTypeName == typeName {
			result = field
		}
	})

	return result
}

// getTypeInsertPos calculates the insertion position for a type declaration
func (a *analyzer) getTypeInsertPos(pass *analysis.Pass, typeSpec *ast.TypeSpec) token.Pos {
	inspect, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return token.NoPos
	}

	var genDecl *ast.GenDecl
	nodeFilter := []ast.Node{
		(*ast.GenDecl)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		if genDecl != nil {
			return
		}
		gd, ok := n.(*ast.GenDecl)
		if !ok || gd.Tok != token.TYPE {
			return
		}

		for _, spec := range gd.Specs {
			if spec == typeSpec {
				genDecl = gd
				return
			}
		}
	})

	if genDecl == nil {
		return token.NoPos
	}

	// Calculate insertion position (before the type declaration)
	if genDecl.Doc != nil && len(genDecl.Doc.List) > 0 {
		// Insert after existing comments
		lastComment := genDecl.Doc.List[len(genDecl.Doc.List)-1]
		return lastComment.End() + 1 // After last comment + newline
	}
	// Insert before the type keyword
	return genDecl.Pos()
}

// getFieldInsertPos calculates the insertion position for a field
func (a *analyzer) getFieldInsertPos(field *ast.Field) token.Pos {
	// Calculate insertion position (before the field)
	if field.Doc != nil && len(field.Doc.List) > 0 {
		// Insert after existing field comments
		lastComment := field.Doc.List[len(field.Doc.List)-1]
		return lastComment.End() + 1 // After last comment + newline
	}
	// Insert before the field
	return field.Pos()
}

// getFieldTypeName extracts the type name from a field type expression
func getFieldTypeName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		// Handle pointer types
		if ident, ok := t.X.(*ast.Ident); ok {
			return ident.Name
		}
	}
	return ""
}

// cleanMarkerString returns the marker string without any // want comments
func cleanMarkerString(marker markershelper.Marker) string {
	markerStr := marker.String()
	// Remove any // want comment suffix
	if idx := strings.Index(markerStr, " //"); idx != -1 {
		markerStr = markerStr[:idx]
	}
	return markerStr
}

// removeMarker creates a TextEdit that removes a marker
func (a *analyzer) removeMarker(marker markershelper.Marker) analysis.TextEdit {
	// For now, just remove the marker text itself
	// The go/analysis framework will handle line cleanup automatically in many cases
	return analysis.TextEdit{
		Pos:     marker.Pos,
		End:     marker.End + 1,
		NewText: []byte(""),
	}
}
