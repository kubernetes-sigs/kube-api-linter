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
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	kalerrors "sigs.k8s.io/kube-api-linter/pkg/analysis/errors"
	markershelper "sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/markers"
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

	a := &analyzer{
		markerRules: DefaultMarkerRules(),
		policy:      cfg.Policy,
	}

	// Override with custom rules if provided
	if cfg.MarkerRules != nil {
		for marker, rule := range cfg.MarkerRules {
			a.markerRules[marker] = rule
		}
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

// reportScopeViolation reports a scope violation error
func (a *analyzer) reportScopeViolation(pass *analysis.Pass, marker markershelper.Marker, rule MarkerScopeRule) {
	var allowedScopes []string
	if rule.Scope&FieldScope != 0 {
		allowedScopes = append(allowedScopes, "fields")
	}
	if rule.Scope&TypeScope != 0 {
		allowedScopes = append(allowedScopes, "types")
	}

	scopeMsg := strings.Join(allowedScopes, " or ")
	if len(allowedScopes) == 0 {
		scopeMsg = "unknown scope"
	}

	pass.Report(analysis.Diagnostic{
		Pos:     marker.Pos,
		End:     marker.End,
		Message: fmt.Sprintf("marker %q can only be applied to %s", marker.Identifier, scopeMsg),
	})
}

// checkFieldMarkers checks markers on fields for violations
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
			a.reportScopeViolation(pass, marker, rule)
			continue
		}

		// Check type constraints if present
		if rule.TypeConstraint != nil {
			if err := a.validateFieldTypeConstraint(pass, field, rule.TypeConstraint); err != nil {
				pass.Report(analysis.Diagnostic{
					Pos:     marker.Pos,
					End:     marker.End,
					Message: fmt.Sprintf("marker %q: %s", marker.Identifier, err),
				})
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
			rule, ok := a.markerRules[marker.Identifier]
			if !ok {
				// No rule defined for this marker, skip validation
				continue
			}

			// Check if TypeScope is allowed
			if !rule.Scope.Allows(TypeScope) {
				a.reportScopeViolation(pass, marker, rule)
				continue
			}

			// Check type constraints if present
			if rule.TypeConstraint != nil {
				if err := a.validateTypeSpecTypeConstraint(pass, typeSpec, rule.TypeConstraint); err != nil {
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

// validateFieldTypeConstraint validates that a field's type matches the type constraint
func (a *analyzer) validateFieldTypeConstraint(pass *analysis.Pass, field *ast.Field, tc *TypeConstraint) error {
	// Get the type of the field
	tv, ok := pass.TypesInfo.Types[field.Type]
	if !ok {
		return nil // Skip if we can't determine the type
	}

	return validateTypeAgainstConstraint(tv.Type, tc)
}

// validateTypeSpecTypeConstraint validates that a type spec's type matches the type constraint
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

// validateTypeAgainstConstraint validates that a Go type satisfies the type constraint
func validateTypeAgainstConstraint(t types.Type, tc *TypeConstraint) error {
	if tc == nil {
		return nil
	}

	// Get the schema type from the Go type
	schemaType := getSchemaType(t)

	// Check if the schema type is allowed
	if len(tc.AllowedSchemaTypes) > 0 {
		allowed := false
		for _, allowedType := range tc.AllowedSchemaTypes {
			if schemaType == allowedType {
				allowed = true
				break
			}
		}
		if !allowed {
			return fmt.Errorf("type %s is not allowed (expected one of: %v)", schemaType, tc.AllowedSchemaTypes)
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

// getSchemaType converts a Go type to an OpenAPI schema type
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
		}
	case *types.Slice, *types.Array:
		return SchemaTypeArray
	case *types.Map, *types.Struct:
		return SchemaTypeObject
	}

	return ""
}

// getElementType returns the element type of an array or slice
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
