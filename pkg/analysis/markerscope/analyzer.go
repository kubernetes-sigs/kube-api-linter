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
	"go/types"
	"maps"
	"slices"
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

// TODO: SuggestFix.
func init() {
	// Register all markers we want to validate scope for
	defaults := DefaultMarkerRules()
	markers := make([]string, 0, len(defaults))

	for marker := range defaults {
		markers = append(markers, marker)
	}

	markershelper.DefaultRegistry().Register(markers...)
}

type analyzer struct {
	markerRules map[string]MarkerScopeRule
	policy      MarkerScopePolicy
}

// newAnalyzer creates a new analyzer.
func newAnalyzer(cfg *MarkerScopeConfig) *analysis.Analyzer {
	if cfg == nil {
		cfg = &MarkerScopeConfig{}
	}

	a := &analyzer{
		markerRules: mergeMarkerRules(DefaultMarkerRules(), cfg.MarkerRules),
		policy:      cfg.Policy,
	}

	// Register all markers (both default and custom) with the markers helper
	// This must be done before the analyzer runs because the markers helper
	// analyzer needs to know about these markers
	for marker := range a.markerRules {
		markershelper.DefaultRegistry().Register(marker)
	}

	// Set default policy if not specified
	if a.policy == "" {
		a.policy = MarkerScopePolicyWarn
	}

	return &analysis.Analyzer{
		Name:             name,
		Doc:              "Validates that markers are applied in the correct scope.",
		Run:              a.run,
		Requires:         []*analysis.Analyzer{inspect.Analyzer, markershelper.Analyzer},
		RunDespiteErrors: true,
	}
}

// mergeMarkerRules merges custom marker rules with default marker rules.
// Custom rules take precedence over default rules for the same marker.
func mergeMarkerRules(defaults, custom map[string]MarkerScopeRule) map[string]MarkerScopeRule {
	merged := make(map[string]MarkerScopeRule, len(defaults)+len(custom))

	// Copy all default rules
	maps.Copy(merged, defaults)

	// Override with custom rules
	maps.Copy(merged, custom)

	return merged
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

	return nil, nil //nolint:nilnil
}

// checkFieldMarkers checks markers on fields for violations.
func (a *analyzer) checkFieldMarkers(pass *analysis.Pass, field *ast.Field, markersAccess markershelper.Markers) {
	fieldMarkers := markersAccess.FieldMarkers(field)

	for _, marker := range fieldMarkers.UnsortedList() {
		rule, ok := a.markerRules[marker.Identifier]
		if !ok {
			// No rule defined for this marker, skip validation
			continue
		}

		// Check if FieldScope is allowed
		if !rule.Scope.Allows(FieldScope) {
			var message string

			var fixes []analysis.SuggestedFix

			if rule.Scope == TypeScope {
				message = fmt.Sprintf("marker %q can only be applied to types", marker.Identifier)

				if a.policy == MarkerScopePolicySuggestFix {
					fixes = a.suggestMoveToFieldsIfCompatible(pass, field, marker, rule)
				}
			} else {
				// This shouldn't happen in practice, but handle it gracefully
				message = fmt.Sprintf("marker %q cannot be applied to fields", marker.Identifier)
			}

			pass.Report(analysis.Diagnostic{
				Pos:            marker.Pos,
				End:            marker.End,
				Message:        message,
				SuggestedFixes: fixes,
			})

			continue
		}

		// Check type constraints if present
		if rule.TypeConstraint != nil {
			if err := a.validateFieldTypeConstraint(pass, field, rule.TypeConstraint); err != nil {
				if a.policy == MarkerScopePolicySuggestFix {
					pass.Report(analysis.Diagnostic{
						Pos:            marker.Pos,
						End:            marker.End,
						Message:        fmt.Sprintf("marker %q: %s", marker.Identifier, err),
						SuggestedFixes: a.suggestMoveToFieldsIfCompatible(pass, field, marker, rule),
					})
				} else {
					pass.Report(analysis.Diagnostic{
						Pos:     marker.Pos,
						End:     marker.End,
						Message: fmt.Sprintf("marker %q: %s", marker.Identifier, err),
					})
				}
			}
		}
	}
}

// checkTypeMarkers checks markers on types for violations.
func (a *analyzer) checkTypeMarkers(pass *analysis.Pass, genDecl *ast.GenDecl, markersAccess markershelper.Markers) {
	if len(genDecl.Specs) == 0 {
		return
	}

	for i := range genDecl.Specs {
		typeSpec, ok := genDecl.Specs[i].(*ast.TypeSpec)
		if !ok {
			continue
		}

		a.checkSingleTypeMarkers(pass, typeSpec, markersAccess)
	}
}

// checkSingleTypeMarkers checks markers on a single type for violations.
func (a *analyzer) checkSingleTypeMarkers(pass *analysis.Pass, typeSpec *ast.TypeSpec, markersAccess markershelper.Markers) {
	typeMarkers := markersAccess.TypeMarkers(typeSpec)

	for _, marker := range typeMarkers.UnsortedList() {
		rule, ok := a.markerRules[marker.Identifier]
		if !ok {
			// No rule defined for this marker, skip validation
			continue
		}

		// Check if TypeScope is allowed
		if !rule.Scope.Allows(TypeScope) {
			a.reportTypeScopeViolation(pass, typeSpec, marker, rule)
			continue
		}

		// Check type constraints if present
		if rule.TypeConstraint != nil {
			a.checkTypeConstraintViolation(pass, typeSpec, marker, rule)
		}
	}
}

// reportTypeScopeViolation reports a scope violation for a type marker.
func (a *analyzer) reportTypeScopeViolation(pass *analysis.Pass, typeSpec *ast.TypeSpec, marker markershelper.Marker, rule MarkerScopeRule) {
	var message string

	var fixes []analysis.SuggestedFix

	if rule.Scope == FieldScope {
		message = fmt.Sprintf("marker %q can only be applied to fields", marker.Identifier)

		if a.policy == MarkerScopePolicySuggestFix {
			fixes = a.suggestMoveToField(pass, typeSpec, marker, rule)
		}
	} else {
		message = fmt.Sprintf("marker %q cannot be applied to types", marker.Identifier)
	}

	pass.Report(analysis.Diagnostic{
		Pos:            marker.Pos,
		End:            marker.End,
		Message:        message,
		SuggestedFixes: fixes,
	})
}

// checkTypeConstraintViolation checks and reports type constraint violations.
func (a *analyzer) checkTypeConstraintViolation(pass *analysis.Pass, typeSpec *ast.TypeSpec, marker markershelper.Marker, rule MarkerScopeRule) {
	if err := a.validateTypeSpecTypeConstraint(pass, typeSpec, rule.TypeConstraint); err != nil {
		var fixes []analysis.SuggestedFix

		if a.policy == MarkerScopePolicySuggestFix {
			fixes = a.suggestMoveToField(pass, typeSpec, marker, rule)
		}

		message := fmt.Sprintf("marker %q: %s", marker.Identifier, err)
		pass.Report(analysis.Diagnostic{
			Pos:            marker.Pos,
			End:            marker.End,
			Message:        message,
			SuggestedFixes: fixes,
		})
	}
}

// validateFieldTypeConstraint validates that a field's type matches the type constraint.
func (a *analyzer) validateFieldTypeConstraint(pass *analysis.Pass, field *ast.Field, tc *TypeConstraint) error {
	// Get the type of the field
	tv, ok := pass.TypesInfo.Types[field.Type]
	if !ok {
		return nil // Skip if we can't determine the type
	}

	if err := validateTypeAgainstConstraint(tv.Type, tc); err != nil {
		return err
	}

	return nil
}

// validateTypeSpecTypeConstraint validates that a type spec's type matches the type constraint.
func (a *analyzer) validateTypeSpecTypeConstraint(pass *analysis.Pass, typeSpec *ast.TypeSpec, tc *TypeConstraint) error {
	// Get the type of the type spec
	obj := pass.TypesInfo.Defs[typeSpec.Name]
	if obj == nil {
		return nil // Skip if we can't determine the type
	}

	typeName, ok := obj.(*types.TypeName)
	if !ok {
		return nil
	}

	return validateTypeAgainstConstraint(typeName.Type(), tc)
}

// validateTypeAgainstConstraint validates that a Go type satisfies the type constraint.
func validateTypeAgainstConstraint(t types.Type, tc *TypeConstraint) error {
	if tc == nil {
		return nil
	}

	// Get the schema type from the Go type
	schemaType := getSchemaType(t)

	// Check if the schema type is allowed
	if len(tc.AllowedSchemaTypes) > 0 {
		if !slices.Contains(tc.AllowedSchemaTypes, schemaType) {
			return fmt.Errorf("%w: type %s (expected one of: %v)", errTypeNotAllowed, schemaType, tc.AllowedSchemaTypes)
		}
	}

	// Validate element constraint for arrays/slices
	if tc.ElementConstraint != nil && schemaType == SchemaTypeArray {
		elemType := getElementType(t)
		if elemType != nil {
			if err := validateTypeAgainstConstraint(elemType, tc.ElementConstraint); err != nil {
				return fmt.Errorf("array element: %w", err)
			}
		}
	}

	return nil
}

// getSchemaType converts a Go type to an OpenAPI schema type.
//
//nolint:cyclop // This function has many cases for different Go types
func getSchemaType(t types.Type) SchemaType {
	// Unwrap pointer types
	if ptr, ok := t.(*types.Pointer); ok {
		t = ptr.Elem()
	}

	// Unwrap named types to get underlying type
	if named, ok := t.(*types.Named); ok {
		t = named.Underlying()
	}

	switch ut := t.Underlying().(type) {
	case *types.Basic:
		switch ut.Kind() {
		case types.Bool:
			return SchemaTypeBoolean
		case types.Int, types.Int8, types.Int16, types.Int32, types.Int64,
			types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64:
			return SchemaTypeInteger
		case types.Float32, types.Float64:
			return SchemaTypeNumber
		case types.String:
			return SchemaTypeString
		case types.Invalid, types.Uintptr, types.Complex64, types.Complex128,
			types.UnsafePointer, types.UntypedBool, types.UntypedInt, types.UntypedRune,
			types.UntypedFloat, types.UntypedComplex, types.UntypedString, types.UntypedNil:
			// These types are not supported in OpenAPI schemas
			return ""
		}
	case *types.Slice, *types.Array:
		return SchemaTypeArray
	case *types.Map, *types.Struct:
		return SchemaTypeObject
	}

	return ""
}

// getElementType returns the element type of an array or slice.
func getElementType(t types.Type) types.Type {
	// Unwrap pointer types
	if ptr, ok := t.(*types.Pointer); ok {
		t = ptr.Elem()
	}

	// Unwrap named types to get underlying type
	if named, ok := t.(*types.Named); ok {
		t = named.Underlying()
	}

	switch ut := t.(type) {
	case *types.Slice:
		return ut.Elem()
	case *types.Array:
		return ut.Elem()
	}

	return nil
}

// extractIdent extracts an *ast.Ident from an ast.Expr, unwrapping pointers and arrays.
func extractIdent(expr ast.Expr) *ast.Ident {
	switch e := expr.(type) {
	case *ast.Ident:
		return e
	case *ast.StarExpr:
		return extractIdent(e.X)
	case *ast.ArrayType:
		return extractIdent(e.Elt)
	default:
		return nil
	}
}

func (a *analyzer) suggestMoveToField(pass *analysis.Pass, typeSpec *ast.TypeSpec, marker markershelper.Marker, rule MarkerScopeRule) []analysis.SuggestedFix {
	// Only suggest moving to field if FieldScope is allowed
	if !rule.Scope.Allows(FieldScope) {
		return nil
	}

	fieldTypeSpecs := utils.LookupFieldsUsingType(pass, typeSpec)
	fmt.Println("fieldTypeSpecs", fieldTypeSpecs)

	var edits []analysis.TextEdit

	// Remove marker from current field (including the newline)
	edits = append(edits, analysis.TextEdit{
		Pos: marker.Pos,
		End: marker.End + 1,
	})

	for _, fieldTypeSpec := range fieldTypeSpecs {
		// Add marker to the line before the type definition
		markerText := a.extractMarkerText(marker)

		file := pass.Fset.File(fieldTypeSpec.Pos())
		if file != nil {
			lineStart := file.LineStart(file.Line(fieldTypeSpec.Pos()))
			edits = append(edits, analysis.TextEdit{
				Pos:     lineStart,
				End:     lineStart,
				NewText: []byte(markerText),
			})
		}
	}

	return []analysis.SuggestedFix{
		{
			Message:   "Move marker to field definition",
			TextEdits: edits,
		},
	}
}

// suggestMoveToFieldsIfCompatible generates suggested fixes to move a marker from type to compatible fields.
func (a *analyzer) suggestMoveToFieldsIfCompatible(pass *analysis.Pass, field *ast.Field, marker markershelper.Marker, rule MarkerScopeRule) []analysis.SuggestedFix {
	// Only suggest moving to type if TypeScope is allowed
	if !rule.Scope.Allows(TypeScope) {
		return nil
	}

	// Extract identifier from field type
	ident := extractIdent(field.Type)
	if ident == nil {
		return nil
	}

	fieldTypeSpec, ok := utils.LookupTypeSpec(pass, ident)
	if !ok {
		return nil
	}

	var edits []analysis.TextEdit

	// Remove marker from current field (including the newline)
	edits = append(edits, analysis.TextEdit{
		Pos: marker.Pos,
		End: marker.End + 1,
	})

	// Add marker to the line before the type definition
	markerText := a.extractMarkerText(marker)

	file := pass.Fset.File(fieldTypeSpec.Pos())
	if file != nil {
		lineStart := file.LineStart(file.Line(fieldTypeSpec.Pos()))
		edits = append(edits, analysis.TextEdit{
			Pos:     lineStart,
			End:     lineStart,
			NewText: []byte(markerText),
		})
	}

	return []analysis.SuggestedFix{
		{
			Message:   "Move marker to type definition",
			TextEdits: edits,
		},
	}
}

func (a *analyzer) extractMarkerText(marker markershelper.Marker) string {
	return strings.Split(marker.RawComment, " //")[0] + "\n"
}
