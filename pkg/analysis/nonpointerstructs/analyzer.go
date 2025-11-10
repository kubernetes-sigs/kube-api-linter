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
package nonpointerstructs

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	kalerrors "sigs.k8s.io/kube-api-linter/pkg/analysis/errors"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/extractjsontags"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/inspector"
	markershelper "sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/markers"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils"
	"sigs.k8s.io/kube-api-linter/pkg/markers"
)

const name = "nonpointerstructs"

func newAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     name,
		Doc:      "Checks that non-pointer structs that contain required fields are marked as required. Non-pointer structs that contain no required fields are marked as optional.",
		Run:      run,
		Requires: []*analysis.Analyzer{inspector.Analyzer},
	}
}

func run(pass *analysis.Pass) (any, error) {
	inspect, ok := pass.ResultOf[inspector.Analyzer].(inspector.Inspector)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetInspector
	}

	inspect.InspectFields(func(field *ast.Field, jsonTagInfo extractjsontags.FieldTagInfo, markersAccess markershelper.Markers, qualifiedFieldName string) {
		checkField(pass, field, markersAccess, jsonTagInfo, qualifiedFieldName)
	})

	return nil, nil //nolint:nilnil
}

func checkField(pass *analysis.Pass, field *ast.Field, markersAccess markershelper.Markers, jsonTagInfo extractjsontags.FieldTagInfo, qualifiedFieldName string) {
	if field.Type == nil {
		return
	}

	if jsonTagInfo.Inline {
		return
	}

	structType, ok := asNonPointerStruct(pass, field.Type)
	if !ok {
		return
	}

	hasRequiredField := hasRequiredField(structType, markersAccess)
	isOptional := utils.IsFieldOptional(field, markersAccess)
	isRequired := utils.IsFieldRequired(field, markersAccess)

	switch {
	case hasRequiredField && isRequired, !hasRequiredField && isOptional:
		// This is the desired case.
	case hasRequiredField:
		pass.Reportf(field.Pos(), "field %s is a non-pointer struct with required fields. It must be marked as required.", qualifiedFieldName)
	case !hasRequiredField:
		pass.Reportf(field.Pos(), "field %s is a non-pointer struct with no required fields. It must be marked as optional.", qualifiedFieldName)
	}
}

func asNonPointerStruct(pass *analysis.Pass, field ast.Expr) (*ast.StructType, bool) {
	switch typ := field.(type) {
	case *ast.StructType:
		return typ, true
	case *ast.Ident:
		typeSpec, ok := utils.LookupTypeSpec(pass, typ)
		if !ok {
			return nil, false
		}

		return asNonPointerStruct(pass, typeSpec.Type)
	default:
		return nil, false
	}
}

func hasRequiredField(structType *ast.StructType, markersAccess markershelper.Markers) bool {
	for _, field := range structType.Fields.List {
		if utils.IsFieldRequired(field, markersAccess) {
			return true
		}
	}

	structMarkers := markersAccess.StructMarkers(structType)

	if structMarkers.Has(markers.KubebuilderMinPropertiesMarker) && !structMarkers.HasWithValue(fmt.Sprintf("%s=0", markers.KubebuilderMinPropertiesMarker)) {
		// A non-zero min properties marker means that the struct is validated to have at least one field.
		// This means it can be treated the same as having a required field.
		return true
	}

	return false
}
