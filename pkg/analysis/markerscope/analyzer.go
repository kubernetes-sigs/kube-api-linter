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
	"errors"
	"fmt"
	"go/ast"
	"go/types"
	"maps"
	"slices"

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

	// Apply default configuration
	defaultConfig(cfg)

	// Convert override and custom marker lists to maps
	overrideRules := markerRulesListToMap(cfg.OverrideMarkers)
	customRules := markerRulesListToMap(cfg.CustomMarkers)

	// Merge rules:
	// 1. Start with default built-in marker rules
	// 2. Apply overrides (replaces default rules for built-in markers)
	// 3. Add custom markers (new markers not in defaults)
	// Note: Validation ensures overrideMarkers only contains built-in markers
	// and customMarkers only contains non-built-in markers, so no conflicts.
	rules := DefaultMarkerRules()
	maps.Copy(rules, overrideRules) // Override built-in markers
	maps.Copy(rules, customRules)   // Add custom markers

	a := &analyzer{
		markerRules: rules,
		policy:      cfg.Policy,
	}

	// Register all markers (both default and custom) with the markers helper
	// This must be done before the analyzer runs because the markers helper
	// analyzer needs to know about these markers
	for marker := range a.markerRules {
		markershelper.DefaultRegistry().Register(marker)
	}

	return &analysis.Analyzer{
		Name: name,
		Doc: `Validates that markers are applied in the correct scope and to compatible data types.
		This analyzer performs two levels of validation:
		1. Scope validation - ensures markers are placed on the correct location (field vs type)
		2. Type constraint validation - ensures markers are applied to compatible data types
		The analyzer includes 100+ built-in kubebuilder marker rules. You can override built-in marker
		rules using overrideMarkers configuration, or add custom markers using customMarkers configuration.
		`,
		Run:              a.run,
		Requires:         []*analysis.Analyzer{inspect.Analyzer, markershelper.Analyzer},
		RunDespiteErrors: true,
	}
}

// markerRulesListToMap converts a list of marker rules to a map keyed by marker identifier.
func markerRulesListToMap(rules []MarkerScopeRule) map[string]MarkerScopeRule {
	result := make(map[string]MarkerScopeRule, len(rules))

	for _, rule := range rules {
		if rule.Identifier != "" {
			result[rule.Identifier] = rule
		}
	}

	return result
}

// defaultConfig applies default values to the configuration.
func defaultConfig(cfg *MarkerScopeConfig) {
	// Set default policy if not specified
	if cfg.Policy == "" {
		cfg.Policy = MarkerScopePolicyWarn
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
			a.reportFieldScopeViolation(pass, field, marker, rule)
			continue
		}

		// Check type constraints if present
		a.checkFieldTypeConstraintViolation(pass, field, marker, rule)
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
		a.checkTypeConstraintViolation(pass, typeSpec, marker, rule)
	}
}

// reportFieldScopeViolation reports a scope violation for a field marker.
func (a *analyzer) reportFieldScopeViolation(pass *analysis.Pass, field *ast.Field, marker markershelper.Marker, rule MarkerScopeRule) {
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
}

// checkFieldTypeConstraintViolation checks and reports type constraint violations for field markers.
func (a *analyzer) checkFieldTypeConstraintViolation(pass *analysis.Pass, field *ast.Field, marker markershelper.Marker, rule MarkerScopeRule) {
	if err := a.validateFieldTypeConstraint(pass, field, rule); err != nil {
		var fixes []analysis.SuggestedFix

		if a.policy == MarkerScopePolicySuggestFix {
			// Check if this is a "should be on type definition" error
			var moveErr *markerShouldBeOnTypeDefinitionError
			if errors.As(err, &moveErr) {
				// Suggest moving to type definition
				fixes = a.suggestMoveToFieldsIfCompatible(pass, field, marker, rule)
			} else {
				// Type constraint violation - suggest removing the marker
				fixes = []analysis.SuggestedFix{
					{
						Message: "Remove invalid marker",
						TextEdits: []analysis.TextEdit{
							{
								Pos: marker.Pos,
								End: marker.End + 1, // Include newline
							},
						},
					},
				}
			}
		}

		pass.Report(analysis.Diagnostic{
			Pos:            marker.Pos,
			End:            marker.End,
			Message:        fmt.Sprintf("marker %q: %s", marker.Identifier, err),
			SuggestedFixes: fixes,
		})
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
			// Check if this is a "should be on field" error (though validateTypeSpecTypeConstraint doesn't return this)
			// For consistency with checkFieldMarkers, we check the error type
			var moveErr *markerShouldBeOnTypeDefinitionError
			if errors.As(err, &moveErr) {
				// This shouldn't happen for type specs, but handle it for consistency
				fixes = a.suggestMoveToField(pass, typeSpec, marker, rule)
			} else {
				// Type constraint violation - suggest removing the marker
				fixes = []analysis.SuggestedFix{
					{
						Message: "Remove invalid marker",
						TextEdits: []analysis.TextEdit{
							{
								Pos: marker.Pos,
								End: marker.End + 1, // Include newline
							},
						},
					},
				}
			}
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
func (a *analyzer) validateFieldTypeConstraint(pass *analysis.Pass, field *ast.Field, rule MarkerScopeRule) error {
	// Get the type of the field
	tv, ok := pass.TypesInfo.Types[field.Type]
	if !ok {
		return nil // Skip if we can't determine the type
	}

	if err := validateTypeAgainstConstraint(tv.Type, rule.TypeConstraint); err != nil {
		return err
	}

	// Check if the marker should be on the type definition instead of the field
	if rule.NamedTypeConstraint == NamedTypeConstraintRequireTypeDefinition && rule.Scope == AnyScope {
		namedType, ok := tv.Type.(*types.Named)
		if ok {
			return &markerShouldBeOnTypeDefinitionError{typeName: namedType.Obj().Name()}
		}
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
	// Get the schema type from the Go type
	schemaType := getSchemaType(t)

	if tc == nil {
		return nil
	}

	// Check if the schema type is allowed
	if len(tc.AllowedSchemaTypes) > 0 {
		if !slices.Contains(tc.AllowedSchemaTypes, schemaType) {
			return &typeNotAllowedError{schemaType: schemaType, allowedTypes: tc.AllowedSchemaTypes}
		}
	}

	// Validate element constraint for arrays/slices
	if tc.ElementConstraint != nil && schemaType == SchemaTypeArray {
		elemType := getElementType(t)
		if elemType != nil {
			if err := validateTypeAgainstConstraint(elemType, tc.ElementConstraint); err != nil {
				return &invalidElementConstraintError{err: err}
			}
		}
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
	return marker.RawComment + "\n"
}
