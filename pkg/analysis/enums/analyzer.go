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
	"slices"
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

const name = "enums"

func init() {
	// Register enum markers so they can be parsed
	markershelper.DefaultRegistry().Register(
		markers.KubebuilderEnumMarker,
		markers.K8sEnumMarker,
	)
}

type analyzer struct {
	config *Config
}

func newAnalyzer(cfg *Config) *analysis.Analyzer {
	a := &analyzer{config: cfg}

	return &analysis.Analyzer{
		Name:     name,
		Doc:      "Enforces that string type aliases with constants have enum markers (+kubebuilder:validation:Enum for CRDs, +k8s:enum for core APIs) and that enum values follow PascalCase conventions",
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
	inspect.InspectFields(func(field *ast.Field, _ extractjsontags.FieldTagInfo, markersAccess markershelper.Markers, qualifiedFieldName string) {
		a.checkField(pass, field, markersAccess, qualifiedFieldName)
	})

	// Check type declarations for enum markers
	inspect.InspectTypeSpec(func(typeSpec *ast.TypeSpec, markersAccess markershelper.Markers) {
		a.checkTypeSpec(pass, typeSpec, markersAccess)
	})

	a.checkConstValues(pass)

	return nil, nil //nolint:nilnil
}

func (a *analyzer) checkField(pass *analysis.Pass, field *ast.Field, markersAccess markershelper.Markers, qualifiedFieldName string) {
	if qualifiedFieldName == "" {
		return
	}

	fieldType, isArray := unwrapTypeWithArrayTracking(field.Type)
	ident, ok := fieldType.(*ast.Ident)

	if !ok {
		return
	}

	prefix := buildFieldPrefix(qualifiedFieldName, isArray)

	if ident.Name == "string" && utils.IsBasicType(pass, ident) {
		if a.config != nil && a.config.RequireTypeAliasForEnums {
			a.checkPlainStringField(pass, field, markersAccess, prefix)
		}

		return
	}

	a.checkTypeAliasField(pass, field, ident, markersAccess, prefix)
}

func (a *analyzer) checkPlainStringField(pass *analysis.Pass, field *ast.Field, markersAccess markershelper.Markers, prefix string) {
	if !hasEnumMarker(markersAccess.FieldMarkers(field)) {
		pass.Reportf(field.Pos(),
			"%s uses plain string type. Consider using a type alias with an enum marker (+kubebuilder:validation:Enum or +k8s:enum)",
			prefix)
	}
}

func (a *analyzer) checkTypeAliasField(pass *analysis.Pass, field *ast.Field, ident *ast.Ident, markersAccess markershelper.Markers, prefix string) {
	if utils.IsBasicType(pass, ident) {
		return
	}

	typeSpec, ok := utils.LookupTypeSpec(pass, ident)
	if !ok || !isStringTypeAlias(pass, typeSpec) {
		return
	}
	// Only check if this type has constants (indicating enum usage)
	if !typeHasConstants(pass, typeSpec.Name.Name) {
		return
	}
	// Check for enum markers (CRD validation or declarative validation)
	if !hasEnumMarker(markersAccess.TypeMarkers(typeSpec)) {
		pass.Reportf(field.Pos(),
			"%s uses type %s which appears to be an enum but is missing an enum marker (+kubebuilder:validation:Enum or +k8s:enum)",
			prefix, typeSpec.Name.Name)
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

	if !isStringTypeAlias(pass, typeSpec) {
		pass.Reportf(typeSpec.Pos(),
			"type %s has enum marker but underlying type is not string",
			typeSpec.Name.Name)
	}
}

func (a *analyzer) checkConstValues(pass *analysis.Pass) {
	for _, file := range pass.Files {
		for _, decl := range file.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok || genDecl.Tok.String() != "const" {
				continue
			}

			for _, spec := range genDecl.Specs {
				if valueSpec, ok := spec.(*ast.ValueSpec); ok {
					a.checkConstSpec(pass, valueSpec)
				}
			}
		}
	}
}

func (a *analyzer) checkConstSpec(pass *analysis.Pass, valueSpec *ast.ValueSpec) {
	for i, name := range valueSpec.Names {
		a.validateEnumConstant(pass, name, valueSpec, i)
	}
}

func (a *analyzer) validateEnumConstant(pass *analysis.Pass, name *ast.Ident, valueSpec *ast.ValueSpec, index int) {
	if name == nil || index >= len(valueSpec.Values) {
		return
	}

	typeSpec := a.getEnumTypeSpec(pass, name)
	if typeSpec == nil {
		return
	}

	// Extract and validate the enum value
	basicLit, ok := valueSpec.Values[index].(*ast.BasicLit)
	if !ok {
		return
	}

	strValue := strings.Trim(basicLit.Value, `"`)

	if !a.isInAllowlist(strValue) && !isPascalCase(strValue) {
		pass.Reportf(basicLit.Pos(),
			"enum value %q should be PascalCase (e.g., \"PhasePending\", \"StateRunning\")",
			strValue)
	}
}

func (a *analyzer) getEnumTypeSpec(pass *analysis.Pass, name *ast.Ident) *ast.TypeSpec {
	constObj, ok := pass.TypesInfo.ObjectOf(name).(*types.Const)
	if !ok {
		return nil
	}

	namedType, ok := constObj.Type().(*types.Named)

	if !ok || namedType.Obj().Pkg() == nil || namedType.Obj().Pkg() != pass.Pkg {
		return nil
	}

	typeSpec := findTypeSpecByName(pass, namedType.Obj().Name())
	if typeSpec == nil || !usesAutoDiscovery(pass, typeSpec) {
		return nil
	}

	return typeSpec
}

func (a *analyzer) isInAllowlist(value string) bool {
	if a.config == nil {
		return false
	}

	return slices.Contains(a.config.Allowlist, value)
}

// Helper functions below this line.

// buildFieldPrefix constructs a human-readable prefix for error messages.
func buildFieldPrefix(fieldName string, isArray bool) string {
	if isArray {
		return fmt.Sprintf("field %s array element", fieldName)
	}

	return fmt.Sprintf("field %s", fieldName)
}

// unwrapType removes pointer and array wrappers to get the underlying type.
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

// isStringTypeAlias checks if a type spec is a string type alias.
func isStringTypeAlias(pass *analysis.Pass, typeSpec *ast.TypeSpec) bool {
	underlyingType := unwrapType(typeSpec.Type)
	ident, ok := underlyingType.(*ast.Ident)

	if !ok {
		return false
	}

	// Both checks are needed: name check is fast, IsBasicType handles edge cases
	// where 'string' might be redefined (rare but possible)
	return ident.Name == "string" && utils.IsBasicType(pass, ident)
}

// hasEnumMarker checks if a marker set contains enum markers
// (either CRD validation marker or declarative validation marker).
func hasEnumMarker(markerSet markershelper.MarkerSet) bool {
	return markerSet.Has(markers.KubebuilderEnumMarker) ||
		markerSet.Has(markers.K8sEnumMarker)
}

// usesAutoDiscovery checks if a type uses auto-discovery mode for enum values.
// Auto-discovery mode validates that constants follow PascalCase conventions.
//
// Returns true for:
//   - +k8s:enum (declarative validation marker, always auto-discovers)
//   - +kubebuilder:validation:Enum without explicit values (CRD validation marker in auto-discovery mode)
//
// Returns false for:
//   - +kubebuilder:validation:Enum=value1;value2 (CRD validation marker with explicit values)
func usesAutoDiscovery(pass *analysis.Pass, typeSpec *ast.TypeSpec) bool {
	for _, file := range pass.Files {
		genDecl := findGenDeclForSpec(file, typeSpec)
		if genDecl == nil || genDecl.Doc == nil {
			continue
		}

		for _, comment := range genDecl.Doc.List {
			text := comment.Text
			// Must be an actual marker (starts with "// +")
			if !strings.HasPrefix(text, "// +") {
				continue
			}

			markerContent := strings.TrimPrefix(text, "// +")

			// +k8s:enum always uses auto-discovery
			if strings.HasPrefix(markerContent, markers.K8sEnumMarker) {
				return true
			}
			// +kubebuilder:validation:Enum without explicit values uses auto-discovery
			if strings.HasPrefix(markerContent, markers.KubebuilderEnumMarker) {
				afterMarker := markerContent[len(markers.KubebuilderEnumMarker):]
				// If there's an "=" or ":=" immediately after, it has explicit values
				trimmed := strings.TrimSpace(afterMarker)
				if strings.HasPrefix(trimmed, "=") || strings.HasPrefix(trimmed, ":=") {
					return false // Explicit values mode
				}

				return true // Auto-discovery mode
			}
		}
	}

	return false
}

func findGenDeclForSpec(file *ast.File, typeSpec *ast.TypeSpec) *ast.GenDecl {
	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		for _, spec := range genDecl.Specs {
			if spec == typeSpec {
				return genDecl
			}
		}
	}

	return nil
}

// findTypeSpecByName finds a type spec by its name in the pass's files.
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

// typeHasConstants checks if any constants are defined for the given type name.
func typeHasConstants(pass *analysis.Pass, typeName string) bool {
	for _, file := range pass.Files {
		for _, decl := range file.Decls {
			genDecl, ok := decl.(*ast.GenDecl)

			if !ok || genDecl.Tok.String() != "const" {
				continue
			}

			for _, spec := range genDecl.Specs {
				valueSpec, ok := spec.(*ast.ValueSpec)
				if !ok || valueSpec.Type == nil {
					continue
				}
				// Check if the const has this type
				if ident, ok := valueSpec.Type.(*ast.Ident); ok && ident.Name == typeName {
					return true
				}
			}
		}
	}

	return false
}

// isPascalCase checks if a string follows PascalCase naming convention.
// Allows: PascalCase (Running), all-uppercase acronyms (HTTP, API), single letters (A).
// Rejects: lowercase start (running), snake_case (phase_pending), kebab-case (phase-pending).
func isPascalCase(s string) bool {
	if len(s) == 0 {
		return false
	}
	// Must start with uppercase
	if !unicode.IsUpper(rune(s[0])) {
		return false
	}

	// No underscores or hyphens allowed (these indicate snake_case/kebab-case)
	for _, r := range s {
		if r == '_' || r == '-' {
			return false
		}
		// Allow letters, digits, and "+" (for signal names like SIGRTMIN+1)
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '+' {
			return false
		}
	}
	// Valid: starts with uppercase, no underscores/hyphens
	// Accepts: PascalCase, HTTP, HTTPS, SIGRTMIN+1, etc.
	return true
}
