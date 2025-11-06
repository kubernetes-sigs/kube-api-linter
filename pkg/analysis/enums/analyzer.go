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

const (
	name = "enums"
)

type analyzer struct {
	config *Config
}

func newAnalyzer(cfg *Config) *analysis.Analyzer {
	a := &analyzer{config: cfg}

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
	inspect.InspectFields(func(field *ast.Field, _ extractjsontags.FieldTagInfo, markersAccess markershelper.Markers, _ string) {
		a.checkField(pass, field, markersAccess)
	})

	// Check type declarations for +enum markers
	inspect.InspectTypeSpec(func(typeSpec *ast.TypeSpec, markersAccess markershelper.Markers) {
		a.checkTypeSpec(pass, typeSpec, markersAccess)
	})
	a.checkConstValues(pass)

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

	prefix := buildFieldPrefix(fieldName, isArray)

	if ident.Name == "string" && utils.IsBasicType(pass, ident) {
		a.checkPlainStringField(pass, field, markersAccess, prefix)

		return
	}

	a.checkTypeAliasField(pass, field, ident, markersAccess, prefix)
}

func buildFieldPrefix(fieldName string, isArray bool) string {
	if isArray {
		return fmt.Sprintf("field %s array element", fieldName)
	}

	return fmt.Sprintf("field %s", fieldName)
}

func (a *analyzer) checkPlainStringField(pass *analysis.Pass, field *ast.Field, markersAccess markershelper.Markers, prefix string) {
	if !hasEnumMarker(markersAccess.FieldMarkers(field)) {
		pass.Reportf(field.Pos(),
			"%s uses plain string without +enum marker. Enumerated fields should use a type alias with +enum marker",
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

	if !hasEnumMarker(markersAccess.TypeMarkers(typeSpec)) {
		pass.Reportf(field.Pos(),
			"%s uses type %s which appears to be an enum but is missing +enum marker (kubebuilder:validation:Enum)",
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
			"type %s has +enum marker but underlying type is not string",
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

	if typeSpec == nil || !hasEnumMarkerOnTypeSpec(pass, typeSpec) {
		return nil
	}

	return typeSpec
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

func isStringTypeAlias(pass *analysis.Pass, typeSpec *ast.TypeSpec) bool {
	underlyingType := unwrapType(typeSpec.Type)

	ident, ok := underlyingType.(*ast.Ident)

	if !ok {
		return false
	}

	return ident.Name == "string" && utils.IsBasicType(pass, ident)
}

func hasEnumMarker(markerSet markershelper.MarkerSet) bool {
	return markerSet.Has(markers.KubebuilderEnumMarker) || markerSet.Has(markers.K8sEnumMarker)
}

func hasEnumMarkerOnTypeSpec(pass *analysis.Pass, typeSpec *ast.TypeSpec) bool {
	for _, file := range pass.Files {
		if genDecl := findGenDeclForSpec(file, typeSpec); genDecl != nil {
			return hasEnumMarkerInDoc(genDecl.Doc)
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

func hasEnumMarkerInDoc(doc *ast.CommentGroup) bool {
	if doc == nil {
		return false
	}

	for _, comment := range doc.List {
		text := comment.Text
		if strings.Contains(text, markers.KubebuilderEnumMarker) || strings.Contains(text, markers.K8sEnumMarker) {
			return true
		}
	}

	return false
}

// isInAllowlist checks if a value is in the configured allowlist.
func (a *analyzer) isInAllowlist(value string) bool {
	if a.config == nil {
		return false
	}

	return slices.Contains(a.config.Allowlist, value)
}

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

func isPascalCase(s string) bool {
	if len(s) == 0 {
		return false
	}

	if !unicode.IsUpper(rune(s[0])) {
		return false
	}

	hasLower := false

	for _, r := range s {
		if r == '_' || r == '-' {
			return false
		}

		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}

		if unicode.IsLower(r) {
			hasLower = true
		}
	}

	return len(s) == 1 || hasLower
}
