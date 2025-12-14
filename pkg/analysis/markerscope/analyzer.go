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

	// Convert override and custom marker lists to maps
	overrideRules := markerRulesListToMap(cfg.OverrideMarkers)
	customRules := markerRulesListToMap(cfg.CustomMarkers)

	// Merge rules:
	// 1. Start with default built-in marker rules
	// 2. Apply overrides (replaces default rules for built-in markers)
	// 3. Add custom markers (new markers not in defaults)
	// Note: Validation ensures overrideMarkers only contains built-in markers
	// and customMarkers only contains non-built-in markers, so no conflicts.
	rules := defaultMarkerRules()
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
	a.checkMarkers(
		markersAccess.FieldMarkers(field).UnsortedList(),
		FieldScope,
		func(marker markershelper.Marker, rule MarkerScopeRule) {
			a.reportFieldScopeViolation(pass, field, marker, rule)
		},
		func(marker markershelper.Marker, rule MarkerScopeRule) {
			a.checkFieldTypeConstraintViolation(pass, field, marker, rule)
		},
	)
}

// checkTypeSpecMarkers checks markers on a type spec for violations.
func (a *analyzer) checkTypeSpecMarkers(pass *analysis.Pass, typeSpec *ast.TypeSpec, markersAccess markershelper.Markers) {
	a.checkMarkers(
		markersAccess.TypeMarkers(typeSpec).UnsortedList(),
		TypeScope,
		func(marker markershelper.Marker, rule MarkerScopeRule) {
			a.reportTypeScopeViolation(pass, typeSpec, marker, rule)
		},
		func(marker markershelper.Marker, rule MarkerScopeRule) {
			a.checkTypeConstraintViolation(pass, typeSpec, marker, rule)
		},
	)
}

// checkMarkers is a common function for checking markers against rules.
func (a *analyzer) checkMarkers(
	unsortedMarkers []markershelper.Marker,
	scope ScopeConstraint,
	reportScopeViolation func(marker markershelper.Marker, rule MarkerScopeRule),
	checkTypeConstraint func(marker markershelper.Marker, rule MarkerScopeRule),
) {
	markers := sortMarkersByPosition(unsortedMarkers)

	for _, marker := range markers {
		rule, ok := a.markerRules[marker.Identifier]
		if !ok {
			// No rule defined for this marker, skip validation
			continue
		}

		// Check if scope is allowed
		if !rule.AllowsScope(scope) {
			reportScopeViolation(marker, rule)
		}

		// Check type constraints if present
		checkTypeConstraint(marker, rule)
	}
}

// reportFieldScopeViolation reports a scope violation for a field marker.
func (a *analyzer) reportFieldScopeViolation(pass *analysis.Pass, field *ast.Field, marker markershelper.Marker, rule MarkerScopeRule) {
	a.reportScopeViolation(pass, marker, rule, TypeScope, "types", "fields", func() []analysis.SuggestedFix {
		return a.suggestMoveToTypeDefinition(pass, field, marker, rule)
	})
}

// checkFieldTypeConstraintViolation checks and reports type constraint violations for field markers.
func (a *analyzer) checkFieldTypeConstraintViolation(pass *analysis.Pass, field *ast.Field, marker markershelper.Marker, rule MarkerScopeRule) {
	// First, validate the type constraint
	if err := a.validateFieldTypeConstraint(pass, field, rule); err != nil {
		a.reportTypeConstraintViolation(pass, marker, err)
		return
	}

	// Then, check if marker should be on type definition instead of field
	if typeName := a.shouldBeOnTypeDefinition(pass, field, rule); typeName != "" {
		a.reportShouldBeOnTypeDefinition(pass, field, marker, rule, typeName)
	}
}

// shouldBeOnTypeDefinition checks if a marker should be on the type definition instead of the field.
// Returns the type name if the marker should be on the type definition, empty string otherwise.
func (a *analyzer) shouldBeOnTypeDefinition(pass *analysis.Pass, field *ast.Field, rule MarkerScopeRule) string {
	if rule.NamedTypeConstraint != NamedTypeConstraintOnTypeOnly || !rule.AllowsScope(TypeScope) {
		return ""
	}

	tv, ok := pass.TypesInfo.Types[field.Type]
	if !ok {
		return ""
	}

	namedType, ok := tv.Type.(*types.Named)
	if !ok {
		return ""
	}

	return namedType.Obj().Name()
}

// reportShouldBeOnTypeDefinition reports that a marker should be on the type definition.
func (a *analyzer) reportShouldBeOnTypeDefinition(pass *analysis.Pass, field *ast.Field, marker markershelper.Marker, rule MarkerScopeRule, typeName string) {
	var fixes []analysis.SuggestedFix

	if a.policy == MarkerScopePolicySuggestFix {
		fixes = a.suggestMoveToTypeDefinition(pass, field, marker, rule)
	}

	pass.Report(analysis.Diagnostic{
		Pos:            marker.Pos,
		End:            marker.End,
		Message:        fmt.Sprintf("marker %q: marker should be declared on the type definition of %s instead of the field", marker.Identifier, typeName),
		SuggestedFixes: fixes,
	})
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

// reportTypeScopeViolation reports a scope violation for a type marker.
func (a *analyzer) reportTypeScopeViolation(pass *analysis.Pass, typeSpec *ast.TypeSpec, marker markershelper.Marker, rule MarkerScopeRule) {
	a.reportScopeViolation(pass, marker, rule, FieldScope, "fields", "types", func() []analysis.SuggestedFix {
		return a.suggestMoveToField(pass, typeSpec, marker, rule)
	})
}

// reportScopeViolation is a common function for reporting scope violations.
func (a *analyzer) reportScopeViolation(
	pass *analysis.Pass,
	marker markershelper.Marker,
	rule MarkerScopeRule,
	alternateScope ScopeConstraint,
	alternateScopeName string,
	appliedScopeName string,
	suggestFix func() []analysis.SuggestedFix,
) {
	var (
		message string
		fixes   []analysis.SuggestedFix
	)

	if rule.AllowsScope(alternateScope) {
		message = fmt.Sprintf("marker %q can only be applied to %s", marker.Identifier, alternateScopeName)

		if a.policy == MarkerScopePolicySuggestFix {
			fixes = suggestFix()
		}
	} else {
		message = fmt.Sprintf("marker %q cannot be applied to %s", marker.Identifier, appliedScopeName)
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
		a.reportTypeSpecTypeConstraintViolation(pass, marker, err)
	}
}

// reportTypeSpecTypeConstraintViolation reports a type constraint violation on a type spec with suggested fix to remove the marker.
func (a *analyzer) reportTypeSpecTypeConstraintViolation(pass *analysis.Pass, marker markershelper.Marker, err error) {
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

// validateFieldTypeConstraint validates that a field's type matches the type constraint.
func (a *analyzer) validateFieldTypeConstraint(pass *analysis.Pass, field *ast.Field, rule MarkerScopeRule) error {
	// Get the type of the field
	tv, ok := pass.TypesInfo.Types[field.Type]
	if !ok {
		return nil // Skip if we can't determine the type
	}

	return validateTypeAgainstConstraint(tv.Type, rule.TypeConstraint)
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
			//nolint:err113 // This is a valid error message
			return fmt.Errorf("type %s is not allowed (expected one of: %v)", schemaType, tc.AllowedSchemaTypes)
		}
	}

	// Validate element constraint for arrays/slices
	if tc.ElementConstraint != nil && schemaType == SchemaTypeArray {
		elemType := utils.UnwrapType(t)
		if elemType != nil {
			if err := validateTypeAgainstConstraint(elemType, tc.ElementConstraint); err != nil {
				return fmt.Errorf("array element: %w", err)
			}
		}
	}

	return nil
}

func (a *analyzer) suggestMoveToField(pass *analysis.Pass, typeSpec *ast.TypeSpec, marker markershelper.Marker, rule MarkerScopeRule) []analysis.SuggestedFix {
	// Only suggest moving to field if FieldScope is allowed
	if !rule.AllowsScope(FieldScope) {
		return nil
	}

	fieldTypeSpecs := utils.LookupTypeSpecUsage(pass, typeSpec)

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

// suggestMoveToTypeDefinition generates suggested fixes to move a marker from a field to its type definition.
func (a *analyzer) suggestMoveToTypeDefinition(pass *analysis.Pass, field *ast.Field, marker markershelper.Marker, rule MarkerScopeRule) []analysis.SuggestedFix {
	// Only suggest moving to type if TypeScope is allowed
	if !rule.AllowsScope(TypeScope) {
		return nil
	}

	// Extract identifier from field type
	ident := utils.ExtractIdent(field.Type)
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
