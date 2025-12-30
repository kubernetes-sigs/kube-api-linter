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
	"cmp"
	"fmt"
	"go/ast"
	"go/types"
	"maps"
	"slices"

	"golang.org/x/tools/go/analysis"

	kalerrors "sigs.k8s.io/kube-api-linter/pkg/analysis/errors"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/extractjsontags"
	inspectorhelper "sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/inspector"
	markershelper "sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/markers"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils"
	"sigs.k8s.io/kube-api-linter/pkg/markers"
)

const (
	name = "markerscope"
)

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

	// Convert custom marker list to map
	customRules := markerRulesListToMap(cfg.CustomMarkers)

	// Merge rules:
	// 1. Start with default built-in marker rules
	// 2. Apply custom markers (replaces default rules if identifier matches, otherwise adds new marker)
	rules := defaultMarkerRules()
	maps.Copy(rules, customRules) // Override or add markers

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
		rules or add custom markers using customMarkers configuration.
		`,
		Run:              a.run,
		Requires:         []*analysis.Analyzer{inspectorhelper.Analyzer},
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
		cfg.Policy = MarkerScopePolicySuggestFix
	}
}

func (a *analyzer) run(pass *analysis.Pass) (any, error) {
	inspector, ok := pass.ResultOf[inspectorhelper.Analyzer].(inspectorhelper.Inspector)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetInspector
	}

	// Check field markers
	inspector.InspectFields(func(field *ast.Field, _ extractjsontags.FieldTagInfo, markersAccess markershelper.Markers, _ string) {
		a.checkFieldMarkers(pass, field, markersAccess)
	})

	// Check type markers
	inspector.InspectTypeSpec(func(typeSpec *ast.TypeSpec, markersAccess markershelper.Markers) {
		a.checkTypeSpecMarkers(pass, typeSpec, markersAccess)
	})

	return nil, nil //nolint:nilnil
}

// sortMarkersByPosition sorts markers by their position to ensure consistent ordering.
func sortMarkersByPosition(markers []markershelper.Marker) []markershelper.Marker {
	slices.SortFunc(markers, func(a, b markershelper.Marker) int {
		return cmp.Compare(a.Pos, b.Pos)
	})

	return markers
}

// checkFieldMarkers checks markers on fields for violations.
func (a *analyzer) checkFieldMarkers(pass *analysis.Pass, field *ast.Field, markersAccess markershelper.Markers) {
	fieldMarkers := markersAccess.FieldMarkers(field)

	for _, marker := range sortMarkersByPosition(fieldMarkers.UnsortedList()) {
		rule, ok := a.markerRules[marker.Identifier]
		if !ok {
			// No rule defined for this marker, skip validation
			continue
		}

		a.checkAllowedScope(pass, field, marker, FieldScope, rule)

		a.checkTypeConstraintViolation(pass, field, marker, rule, fieldMarkers)
	}
}

// checkTypeSpecMarkers checks markers on a type spec for violations.
func (a *analyzer) checkTypeSpecMarkers(pass *analysis.Pass, typeSpec *ast.TypeSpec, markersAccess markershelper.Markers) {
	typeMarkers := markersAccess.TypeMarkers(typeSpec)

	for _, marker := range sortMarkersByPosition(typeMarkers.UnsortedList()) {
		rule, ok := a.markerRules[marker.Identifier]
		if !ok {
			// No rule defined for this marker, skip validation
			continue
		}

		a.checkAllowedScope(pass, typeSpec, marker, TypeScope, rule)

		a.checkTypeConstraintViolation(pass, typeSpec, marker, rule, typeMarkers)
	}
}

// reportShouldBeOnTypeDefinition reports that a marker should be on the type definition.
func (a *analyzer) reportShouldBeOnTypeDefinition(pass *analysis.Pass, field *ast.Field, marker markershelper.Marker, typeName string) {
	var fixes []analysis.SuggestedFix

	if a.policy == MarkerScopePolicySuggestFix {
		ident := utils.ExtractIdent(field.Type)
		if ident != nil {
			if fieldTypeSpec, ok := utils.LookupTypeSpec(pass, ident); ok {
				fixes = a.buildMoveToTypeDefinitionFix(pass, fieldTypeSpec, marker)
			}
		}
	}

	pass.Report(analysis.Diagnostic{
		Pos:            marker.Pos,
		End:            marker.End,
		Message:        fmt.Sprintf("marker %q: marker should be declared on the type definition of %s instead of the field", marker.Identifier, typeName),
		SuggestedFixes: fixes,
	})
}

// handleTypeConstraintValidation handles type constraint validation and reporting.
// Returns true if validation passed (or should continue checking), false if error was reported.
func (a *analyzer) handleTypeConstraintValidation(
	pass *analysis.Pass,
	goType types.Type,
	tc *TypeConstraint,
	marker markershelper.Marker,
	markerSet markershelper.MarkerSet,
) bool {
	// Check if there's a validation:Type marker that overrides the schema type
	schemaTypeOverride := getSchemaTypeFromMarker(markerSet, markers.KubebuilderTypeMarker)
	// Check if there's a validation:items:Type marker that overrides the array element schema type
	itemsSchemaTypeOverride := getSchemaTypeFromMarker(markerSet, markers.KubebuilderItemsTypeMarker)

	if err := handleTypeAgainstConstraint(goType, tc, schemaTypeOverride, itemsSchemaTypeOverride); err != nil {
		a.reportTypeConstraintViolation(pass, marker, err)
		return false
	}
	return true
}

// handleShouldBeOnTypeDefinition handles the check and reporting for markers
// that should be on type definition instead of field.
func (a *analyzer) handleShouldBeOnTypeDefinition(
	pass *analysis.Pass,
	field *ast.Field,
	marker markershelper.Marker,
	rule MarkerScopeRule,
) {
	if rule.NamedTypeConstraint != NamedTypeConstraintOnTypeOnly || !rule.AllowsScope(TypeScope) {
		return
	}

	tv, ok := pass.TypesInfo.Types[field.Type]
	if !ok {
		return
	}

	namedType, ok := tv.Type.(*types.Named)
	if !ok {
		return
	}

	typeName := namedType.Obj().Name()
	a.reportShouldBeOnTypeDefinition(pass, field, marker, typeName)
}

// reportTypeConstraintViolation reports a type constraint violation with suggested fix to remove the marker.
func (a *analyzer) reportTypeConstraintViolation(pass *analysis.Pass, marker markershelper.Marker, err error) {
	var fixes []analysis.SuggestedFix

	if a.policy == MarkerScopePolicySuggestFix {
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

	pass.Report(analysis.Diagnostic{
		Pos:            marker.Pos,
		End:            marker.End,
		Message:        fmt.Sprintf("marker %q: %s", marker.Identifier, err),
		SuggestedFixes: fixes,
	})
}

// reportScopeViolation is a common function for reporting scope violations.
func (a *analyzer) reportScopeViolation(
	pass *analysis.Pass,
	marker markershelper.Marker,
	rule MarkerScopeRule,
	alternateScope ScopeConstraint,
	fixes []analysis.SuggestedFix,
) {
	var message string

	if rule.AllowsScope(alternateScope) {
		message = fmt.Sprintf("marker %q can only be applied to %s", marker.Identifier, alternateScope.PluralName())
	} else {
		message = fmt.Sprintf("marker %q cannot be applied to %s", marker.Identifier, alternateScope.Opposite().PluralName())
	}

	pass.Report(analysis.Diagnostic{
		Pos:            marker.Pos,
		End:            marker.End,
		Message:        message,
		SuggestedFixes: fixes,
	})
}

// checkTypeConstraintViolation checks and reports type constraint violations.
// Supports both *ast.Field and *ast.TypeSpec nodes.
func (a *analyzer) checkTypeConstraintViolation(
	pass *analysis.Pass,
	node ast.Node,
	marker markershelper.Marker,
	rule MarkerScopeRule,
	markerSet markershelper.MarkerSet,
) {
	// Extract types.Type based on node type
	var goType types.Type

	switch n := node.(type) {
	case *ast.Field:
		// Field-side logic: get type from field.Type
		tv, ok := pass.TypesInfo.Types[n.Type]
		if !ok {
			return
		}
		goType = tv.Type

		// Handle type constraint validation
		if !a.handleTypeConstraintValidation(pass, goType, rule.TypeConstraint, marker, markerSet) {
			return // Validation failed, error already reported
		}

		// Handle field-specific: check if marker should be on type definition
		a.handleShouldBeOnTypeDefinition(pass, n, marker, rule)

	case *ast.TypeSpec:
		// Type-side logic: get type from type spec name
		obj := pass.TypesInfo.Defs[n.Name]
		if obj == nil {
			return
		}
		typeName, ok := obj.(*types.TypeName)
		if !ok {
			return
		}
		goType = typeName.Type()

		// Handle type constraint validation
		a.handleTypeConstraintValidation(pass, goType, rule.TypeConstraint, marker, markerSet)

	default:
		return
	}
}

// handleTypeAgainstConstraint validates that a Go type satisfies the type constraint.
// If schemaTypeOverride is not empty, it will be used instead of deriving the schema type from the Go type.
// If itemsSchemaTypeOverride is not empty, it will be used for array element type checking.
func handleTypeAgainstConstraint(t types.Type, tc *TypeConstraint, schemaTypeOverride SchemaType, itemsSchemaTypeOverride SchemaType) error {
	// Get the schema type from the Go type, or use the override if provided
	schemaType := schemaTypeOverride
	if schemaType == "" {
		schemaType = getSchemaType(t)
	}

	if tc == nil {
		return nil
	}

	// Check if the schema type is allowed
	if len(tc.AllowedSchemaTypes) > 0 {
		if !slices.Contains(tc.AllowedSchemaTypes, schemaType) {
			//nolint:err113 // This is a valid error message
			return fmt.Errorf("type %s is not allowed (expected one of: %v)", schemaType, tc.AllowedSchemaTypes)
		}
	}

	// Validate element constraint for arrays/slices
	if tc.ElementConstraint != nil && schemaType == SchemaTypeArray {
		elemType := utils.UnwrapType(t)
		if elemType != nil {
			// Use items type override for element constraint if provided
			if err := handleTypeAgainstConstraint(elemType, tc.ElementConstraint, itemsSchemaTypeOverride, ""); err != nil {
				return fmt.Errorf("array element: %w", err)
			}
		}
	}

	return nil
}

func (a *analyzer) extractMarkerText(marker markershelper.Marker) string {
	return marker.RawComment + "\n"
}

// buildMoveToTypeDefinitionFix builds a suggested fix to move a marker to a type definition.
// Assumes the type definition exists and is valid (validation must be done by caller).
func (a *analyzer) buildMoveToTypeDefinitionFix(pass *analysis.Pass, typeSpec *ast.TypeSpec, marker markershelper.Marker) []analysis.SuggestedFix {
	var edits []analysis.TextEdit

	// Remove marker from current location (including the newline)
	edits = append(edits, analysis.TextEdit{
		Pos: marker.Pos,
		End: marker.End + 1,
	})

	// Add marker to the line before the type definition
	markerText := a.extractMarkerText(marker)
	file := pass.Fset.File(typeSpec.Pos())
	if file != nil {
		lineStart := file.LineStart(file.Line(typeSpec.Pos()))
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

// buildMoveToFieldFix builds a suggested fix to move a marker to field definitions.
// Assumes the field usages exist and are valid (validation must be done by caller).
func (a *analyzer) buildMoveToFieldFix(pass *analysis.Pass, fieldTypeSpecs []*ast.Field, marker markershelper.Marker) []analysis.SuggestedFix {
	var edits []analysis.TextEdit

	// Remove marker from current location (including the newline)
	edits = append(edits, analysis.TextEdit{
		Pos: marker.Pos,
		End: marker.End + 1,
	})

	// Add marker to each field usage
	for _, fieldTypeSpec := range fieldTypeSpecs {
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

// buildScopeViolationFix builds fixes for scope violations.
// Returns suggested fixes if they can be built, nil otherwise.
func (a *analyzer) buildScopeViolationFix(
	pass *analysis.Pass,
	node ast.Node,
	marker markershelper.Marker,
	appliedScope ScopeConstraint,
	targetScope ScopeConstraint,
	rule MarkerScopeRule,
) []analysis.SuggestedFix {
	// Only build fix if policy allows
	if a.policy != MarkerScopePolicySuggestFix {
		return nil
	}

	// Only build fix if target scope is allowed
	if !rule.AllowsScope(targetScope) {
		return nil
	}

	// Build fix based on scope transition
	switch {
	case appliedScope == FieldScope && targetScope == TypeScope:
		// Moving from field to type
		if field, ok := node.(*ast.Field); ok {
			ident := utils.ExtractIdent(field.Type)
			if ident != nil {
				if typeSpec, ok := utils.LookupTypeSpec(pass, ident); ok {
					return a.buildMoveToTypeDefinitionFix(pass, typeSpec, marker)
				}
			}
		}
	case appliedScope == TypeScope && targetScope == FieldScope:
		// Moving from type to field
		if typeSpec, ok := node.(*ast.TypeSpec); ok {
			fieldTypeSpecs := utils.LookupTypeSpecUsage(pass, typeSpec)
			if len(fieldTypeSpecs) > 0 {
				return a.buildMoveToFieldFix(pass, fieldTypeSpecs, marker)
			}
		}
	}

	return nil
}

// handleScopeViolation handles scope violations by building fixes and reporting.
func (a *analyzer) handleScopeViolation(
	pass *analysis.Pass,
	node ast.Node,
	marker markershelper.Marker,
	appliedScope ScopeConstraint,
	rule MarkerScopeRule,
) {
	// Determine the alternate scope
	alternateScope := appliedScope.Opposite()

	// Build fixes if possible
	fixes := a.buildScopeViolationFix(pass, node, marker, appliedScope, alternateScope, rule)

	// Report the violation
	a.reportScopeViolation(pass, marker, rule, alternateScope, fixes)
}

// checkAllowedScope checks if a marker is applied to an allowed scope and reports violations.
func (a *analyzer) checkAllowedScope(
	pass *analysis.Pass,
	node ast.Node,
	marker markershelper.Marker,
	appliedScope ScopeConstraint,
	rule MarkerScopeRule,
) {
	// Check if applied scope is allowed
	if rule.AllowsScope(appliedScope) {
		// Applied scope is allowed, no violation
		return
	}

	// Handle the violation (build + report)
	a.handleScopeViolation(pass, node, marker, appliedScope, rule)
}
