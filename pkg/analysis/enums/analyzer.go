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
package enums

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
	kalerrors "sigs.k8s.io/kube-api-linter/pkg/analysis/errors"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/extractjsontags"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/inspector"
	markershelper "sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/markers"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils"
	"sigs.k8s.io/kube-api-linter/pkg/markers"
)

const (
	name = "enums"
)

type analyzer struct {
	config *Config
}

// newAnalyzer creates a new analysis.Analyzer for the enums linter based on the provided config.
func newAnalyzer(cfg *Config) *analysis.Analyzer {
	a := &analyzer{
		config: cfg,
	}

	return &analysis.Analyzer{
		Name:     name,
		Doc:      "Enforces that enumerated fields use type aliases with +enum marker and have PascalCase values",
		Run:      a.run,
		Requires: []*analysis.Analyzer{inspector.Analyzer},
	}
}

func (a *analyzer) run(pass *analysis.Pass) (any, error) {
	inspect, ok := pass.ResultOf[inspector.Analyzer].(inspector.Inspector)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetInspector
	}

	// Check struct fields for proper enum usage
	inspect.InspectFields(func(field *ast.Field, _ extractjsontags.FieldTagInfo, markersAccess markershelper.Markers) {
		a.checkField(pass, field, markersAccess)
	})

	// Check type declarations for +enum markers
	inspect.InspectTypeSpec(func(typeSpec *ast.TypeSpec, markersAccess markershelper.Markers) {
		a.checkTypeSpec(pass, typeSpec, markersAccess)
	})

	// Check const values for PascalCase
	a.checkConstValues(pass, inspect)

	return nil, nil //nolint:nilnil
}

func (a *analyzer) checkField(pass *analysis.Pass, field *ast.Field, markersAccess markershelper.Markers) {
	fieldName := utils.FieldName(field)
	if fieldName == "" {
		return
	}
	// Get the underlying type, unwrapping pointers and arrays
	fieldType, isArray := unwrapTypeWithArrayTracking(field.Type)
	ident, ok := fieldType.(*ast.Ident)
	if !ok {
		return
	}

	// Build appropriate prefix for error messages
	prefix := fmt.Sprintf("field %s", fieldName)
	if isArray {
		prefix = fmt.Sprintf("field %s array element", fieldName)
	}

	// Check if it's a basic string type
	if ident.Name == "string" && utils.IsBasicType(pass, ident) {
		// Check if the field has an enum marker directly
		fieldMarkers := markersAccess.FieldMarkers(field)
		if !hasEnumMarker(fieldMarkers) {
			pass.Reportf(field.Pos(),
				"%s uses plain string without +enum marker. Enumerated fields should use a type alias with +enum marker",
				prefix)
		}
		return
	}
	// If it's a type alias, check that the alias has an enum marker
	if !utils.IsBasicType(pass, ident) {
		typeSpec, ok := utils.LookupTypeSpec(pass, ident)
		if !ok {
			return
		}
		// Check if the underlying type is string
		if !isStringTypeAlias(pass, typeSpec) {
			return
		}
		// Check if the type has an enum marker
		typeMarkers := markersAccess.TypeMarkers(typeSpec)
		if !hasEnumMarker(typeMarkers) {
			pass.Reportf(field.Pos(),
				"%s uses type %s which appears to be an enum but is missing +enum marker (kubebuilder:validation:Enum)",
				prefix, typeSpec.Name.Name)
		}
	}
}

func (a *analyzer) checkTypeSpec(pass *analysis.Pass, typeSpec *ast.TypeSpec, markersAccess markershelper.Markers) {
	if typeSpec.Name == nil {
		return
	}
	typeMarkers := markersAccess.TypeMarkers(typeSpec)
	if !hasEnumMarker(typeMarkers) {
		return
	}

	// Has +enum marker, verify it's a string type
	if !isStringTypeAlias(pass, typeSpec) {
		pass.Reportf(typeSpec.Pos(),
			"type %s has +enum marker but underlying type is not string",
			typeSpec.Name.Name)
	}
}

func (a *analyzer) checkConstValues(pass *analysis.Pass, inspect inspector.Inspector) {
	// We need to check const declarations, but the inspector helper doesn't
	// have a method for this, so we'll iterate through files manually
	for _, file := range pass.Files {
		for _, decl := range file.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok || genDecl.Tok.String() != "const" {
				continue
			}
			for _, spec := range genDecl.Specs {
				valueSpec, ok := spec.(*ast.ValueSpec)
				if !ok {
					continue
				}
				a.checkConstSpec(pass, valueSpec)
			}
		}
	}
}

func (a *analyzer) checkConstSpec(pass *analysis.Pass, valueSpec *ast.ValueSpec) {
	for i, name := range valueSpec.Names {
		if name == nil {
			continue
		}
		// Get the type of the constant
		obj := pass.TypesInfo.ObjectOf(name)
		if obj == nil {
			continue
		}
		constObj, ok := obj.(*types.Const)
		if !ok {
			continue
		}
		// Check if the type is a named type (potential enum)
		namedType, ok := constObj.Type().(*types.Named)
		if !ok {
			continue
		}
		// Check if the type is in the current package
		if namedType.Obj().Pkg() == nil || namedType.Obj().Pkg() != pass.Pkg {
			continue
		}
		// Find the type spec for this named type
		typeSpec := findTypeSpecByName(pass, namedType.Obj().Name())
		if typeSpec == nil {
			continue
		}
		// Check if this type has an enum marker
		if !hasEnumMarkerOnTypeSpec(pass, typeSpec) {
			continue
		}
		// This is an enum constant, validate the value
		if i >= len(valueSpec.Values) {
			continue
		}
		value := valueSpec.Values[i]
		basicLit, ok := value.(*ast.BasicLit)
		if !ok {
			continue
		}
		// Extract the string value (remove quotes)
		strValue := strings.Trim(basicLit.Value, `"`)
		// Check if it's in the allowlist
		if a.isInAllowlist(strValue) {
			continue
		}
		// Validate PascalCase
		if !isPascalCase(strValue) {
			pass.Reportf(basicLit.Pos(),
				"enum value %q should be PascalCase (e.g., \"PhasePending\", \"StateRunning\")",
				strValue)
		}
	}
}

// unwrapType removes pointer and array wrappers to get the underlying type
func unwrapType(expr ast.Expr) ast.Expr {
	switch t := expr.(type) {
	case *ast.StarExpr:
		return unwrapType(t.X)
	case *ast.ArrayType:
		return unwrapType(t.Elt)
	default:
		return expr
	}
}

// unwrapTypeWithArrayTracking removes pointer and array wrappers to get the underlying type
// and tracks whether an array was encountered during unwrapping.
func unwrapTypeWithArrayTracking(expr ast.Expr) (ast.Expr, bool) {
	isArray := false
	for {
		switch t := expr.(type) {
		case *ast.StarExpr:
			expr = t.X
		case *ast.ArrayType:
			expr = t.Elt
			isArray = true
		default:
			return expr, isArray
		}
	}
}

// isStringTypeAlias checks if a type spec is an alias for string
func isStringTypeAlias(pass *analysis.Pass, typeSpec *ast.TypeSpec) bool {
	underlyingType := unwrapType(typeSpec.Type)
	ident, ok := underlyingType.(*ast.Ident)
	if !ok {
		return false
	}
	return ident.Name == "string" && utils.IsBasicType(pass, ident)
}

// hasEnumMarker checks if a marker set contains an enum marker
func hasEnumMarker(markerSet markershelper.MarkerSet) bool {
	return markerSet.Has(markers.KubebuilderEnumMarker) || markerSet.Has(markers.K8sEnumMarker)
}

// hasEnumMarkerOnTypeSpec checks if a type spec has an enum marker by checking its doc comments
func hasEnumMarkerOnTypeSpec(pass *analysis.Pass, typeSpec *ast.TypeSpec) bool {
	for _, file := range pass.Files {
		for _, decl := range file.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}
			for _, spec := range genDecl.Specs {
				if spec == typeSpec {
					if genDecl.Doc != nil {
						for _, comment := range genDecl.Doc.List {
							if strings.Contains(comment.Text, markers.KubebuilderEnumMarker) ||
								strings.Contains(comment.Text, markers.K8sEnumMarker) {
								return true
							}
						}
					}
					return false
				}
			}
		}
	}
	return false
}

// isInAllowlist checks if a value is in the configured allowlist
func (a *analyzer) isInAllowlist(value string) bool {
	if a.config == nil {
		return false
	}
	for _, allowed := range a.config.Allowlist {
		if value == allowed {
			return true
		}
	}
	return false
}

// findTypeSpecByName searches through the AST files to find a TypeSpec by name
func findTypeSpecByName(pass *analysis.Pass, typeName string) *ast.TypeSpec {
	for _, file := range pass.Files {
		for _, decl := range file.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}
			for _, spec := range genDecl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				if typeSpec.Name != nil && typeSpec.Name.Name == typeName {
					return typeSpec
				}
			}
		}
	}
	return nil
}

// isPascalCase validates that a string follows PascalCase convention
// PascalCase: FirstLetterUpperCase, no underscores, no hyphens
func isPascalCase(s string) bool {
	if len(s) == 0 {
		return false
	}
	// First character must be uppercase
	if !unicode.IsUpper(rune(s[0])) {
		return false
	}
	// Check for invalid characters (underscores, hyphens)
	for _, r := range s {
		if r == '_' || r == '-' {
			return false
		}
		// Only allow letters and digits
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	// Should not be all uppercase (that's SCREAMING_SNAKE_CASE or similar)
	allUpper := true
	for _, r := range s[1:] {
		if unicode.IsLower(r) {
			allUpper = false
			break
		}
	}
	if allUpper && len(s) > 1 {
		return false
	}
	return true
}
