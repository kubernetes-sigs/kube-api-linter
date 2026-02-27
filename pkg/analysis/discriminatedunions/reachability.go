/*
Copyright 2026 The Kubernetes Authors.

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

package discriminatedunions

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	inspectorhelper "sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/inspector"
	markersconsts "sigs.k8s.io/kube-api-linter/pkg/markers"
)

func crdReachableStructs(pass *analysis.Pass, inspect inspectorhelper.Inspector, fields map[*ast.Field]unionField) map[*ast.StructType]bool {
	typesByObject := map[types.Object]*ast.TypeSpec{}

	var pending []ast.Expr

	for ts := range inspect.TypeSpecs() {
		spec := ts.TypeSpec
		typesByObject[pass.TypesInfo.Defs[spec.Name]] = spec

		if ts.Markers.TypeMarkers(spec).HasWithValue(markersconsts.KubebuilderRootMarker + "=true") {
			pending = append(pending, spec.Type)
		}
	}

	reachable := map[*ast.StructType]bool{}
	visited := map[ast.Expr]bool{}

	for len(pending) > 0 {
		expr := pending[len(pending)-1]
		pending = pending[:len(pending)-1]

		if visited[expr] {
			continue
		}

		visited[expr] = true

		if structType, ok := expr.(*ast.StructType); ok {
			reachable[structType] = true
		}

		pending = append(pending, referencedTypes(pass, expr, typesByObject, fields)...)
	}

	return reachable
}

func referencedTypes(pass *analysis.Pass, expr ast.Expr, localTypes map[types.Object]*ast.TypeSpec, fields map[*ast.Field]unionField) []ast.Expr {
	switch typ := expr.(type) {
	case *ast.Ident:
		if spec := localTypes[pass.TypesInfo.ObjectOf(typ)]; spec != nil {
			return []ast.Expr{spec.Type}
		}
	case *ast.StarExpr:
		return []ast.Expr{typ.X}
	case *ast.ArrayType:
		return []ast.Expr{typ.Elt}
	case *ast.MapType:
		return []ast.Expr{typ.Value}
	case *ast.ParenExpr:
		return []ast.Expr{typ.X}
	case *ast.StructType:
		return referencedStructTypes(typ, fields)
	}

	return nil
}

func referencedStructTypes(typ *ast.StructType, fields map[*ast.Field]unionField) []ast.Expr {
	var children []ast.Expr

	for _, field := range typ.Fields.List {
		if len(field.Names) > 0 && !field.Names[0].IsExported() {
			continue
		}

		if _, ok := fields[field]; ok {
			children = append(children, field.Type)
		}
	}

	return children
}
