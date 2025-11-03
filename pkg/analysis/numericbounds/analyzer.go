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
package numericbounds

import (
	"errors"
	"fmt"
	"go/ast"
	"strconv"

	"golang.org/x/tools/go/analysis"
	kalerrors "sigs.k8s.io/kube-api-linter/pkg/analysis/errors"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/extractjsontags"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/inspector"
	markershelper "sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/markers"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils"
	"sigs.k8s.io/kube-api-linter/pkg/markers"
)

const name = "numericbounds"

// JavaScript safe integer bounds (2^53 - 1 and -(2^53 - 1))
const (
	maxSafeInt = 9007199254740991  // 2^53 - 1
	minSafeInt = -9007199254740991 // -(2^53 - 1)
)

var errMarkerMissingValue = errors.New("marker value not found")

// Analyzer is the analyzer for the numericbounds package.
// It checks that numeric fields have appropriate bounds validation markers.
var Analyzer = &analysis.Analyzer{
	Name:     name,
	Doc:      "Checks that numeric fields (int32, int64) have appropriate minimum and maximum bounds validation markers",
	Run:      run,
	Requires: []*analysis.Analyzer{inspector.Analyzer},
}

func run(pass *analysis.Pass) (any, error) {
	inspect, ok := pass.ResultOf[inspector.Analyzer].(inspector.Inspector)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetInspector
	}

	inspect.InspectFields(func(field *ast.Field, _ extractjsontags.FieldTagInfo, markersAccess markershelper.Markers) {
		checkField(pass, field, markersAccess)
	})

	return nil, nil //nolint:nilnil
}

func checkField(pass *analysis.Pass, field *ast.Field, markersAccess markershelper.Markers) {
	fieldName := utils.FieldName(field)
	if fieldName == "" {
		return
	}

	// Unwrap pointers and slices to get the underlying type
	fieldType, isSlice := unwrapType(field.Type)

	// Get the underlying numeric type identifier (int32 or int64)
	ident := getNumericTypeIdent(pass, fieldType)
	if ident == nil {
		return
	}

	// Only check int32 and int64 types
	if ident.Name != "int32" && ident.Name != "int64" {
		return
	}

	fieldMarkers := utils.TypeAwareMarkerCollectionForField(pass, markersAccess, field)

	// Determine which markers to look for based on whether the field is a slice
	minMarker, maxMarker := getMarkerNames(isSlice)

	// Get minimum and maximum marker values
	minimum, minErr := getMarkerNumericValue(fieldMarkers, minMarker)
	maximum, maxErr := getMarkerNumericValue(fieldMarkers, maxMarker)

	// Check if markers are missing
	minMissing := errors.Is(minErr, errMarkerMissingValue)
	maxMissing := errors.Is(maxErr, errMarkerMissingValue)

	// Report any invalid marker values (e.g., non-numeric values)
	if minErr != nil && !minMissing {
		pass.Reportf(field.Pos(), "field %s has an invalid minimum marker: %v", fieldName, minErr)
		return
	}
	if maxErr != nil && !maxMissing {
		pass.Reportf(field.Pos(), "field %s has an invalid maximum marker: %v", fieldName, maxErr)
		return
	}

	// Report if both markers are missing
	if minMissing && maxMissing {
		pass.Reportf(field.Pos(), "field %s of type %s should have minimum and maximum bounds validation markers", fieldName, ident.Name)
		return
	}

	// Report if only one marker is present
	if minMissing {
		pass.Reportf(field.Pos(), "field %s of type %s has maximum but is missing minimum bounds validation marker", fieldName, ident.Name)
		return
	}
	if maxMissing {
		pass.Reportf(field.Pos(), "field %s of type %s has minimum but is missing maximum bounds validation marker", fieldName, ident.Name)
		return
	}

	// For int64 fields, check if bounds are within JavaScript safe integer range
	checkJavaScriptSafeBounds(pass, field, fieldName, ident.Name, minimum, maximum)
}

// getNumericTypeIdent returns the identifier for int32 or int64 types.
// It handles type aliases by looking up the underlying type.
// Note: This function expects pointers and slices to already be unwrapped.
func getNumericTypeIdent(pass *analysis.Pass, expr ast.Expr) *ast.Ident {
	ident, ok := expr.(*ast.Ident)
	if !ok {
		return nil
	}

	// Check if it's a basic int32 or int64 type
	if ident.Name == "int32" || ident.Name == "int64" {
		return ident
	}

	// Check if it's a type alias to int32 or int64
	if !utils.IsBasicType(pass, ident) {
		typeSpec, ok := utils.LookupTypeSpec(pass, ident)
		if ok {
			return getNumericTypeIdent(pass, typeSpec.Type)
		}
	}

	return nil
}

// unwrapType unwraps pointers and slices to get the underlying type.
// Returns the unwrapped type and a boolean indicating if it's a slice.
func unwrapType(expr ast.Expr) (ast.Expr, bool) {
	isSlice := false

	// Unwrap pointer if present (e.g., *int32)
	if starExpr, ok := expr.(*ast.StarExpr); ok {
		expr = starExpr.X
	}

	// Check if it's a slice and unwrap (e.g., []int32)
	if arrayType, ok := expr.(*ast.ArrayType); ok {
		isSlice = true
		expr = arrayType.Elt

		// Handle pointer inside slice (e.g., []*int32)
		if starExpr, ok := expr.(*ast.StarExpr); ok {
			expr = starExpr.X
		}
	}

	return expr, isSlice
}

// getMarkerNames returns the appropriate minimum and maximum marker names
// based on whether the field is a slice.
func getMarkerNames(isSlice bool) (minMarker, maxMarker string) {
	if isSlice {
		return markers.KubebuilderItemsMinimumMarker, markers.KubebuilderItemsMaximumMarker
	}
	return markers.KubebuilderMinimumMarker, markers.KubebuilderMaximumMarker
}

// getMarkerNumericValue extracts the numeric value from the first instance of the marker with the given name.
func getMarkerNumericValue(markerSet markershelper.MarkerSet, markerName string) (float64, error) {
	markerList := markerSet.Get(markerName)
	if len(markerList) == 0 {
		return 0, errMarkerMissingValue
	}

	marker := markerList[0]
	rawValue, ok := marker.Expressions[""]
	if !ok {
		return 0, errMarkerMissingValue
	}

	// Parse as float64 using strconv for better error handling
	value, err := strconv.ParseFloat(rawValue, 64)
	if err != nil {
		return 0, fmt.Errorf("error converting value to number: %w", err)
	}

	return value, nil
}

// checkJavaScriptSafeBounds checks if int64 bounds are within JavaScript safe integer range.
func checkJavaScriptSafeBounds(pass *analysis.Pass, field *ast.Field, fieldName, typeName string, minimum, maximum float64) {
	if typeName != "int64" {
		return
	}

	if minimum < minSafeInt || maximum > maxSafeInt {
		pass.Reportf(field.Pos(),
			"field %s of type int64 has bounds [%d, %d] that exceed safe integer range [%d, %d]. Consider using a string type to avoid precision loss in JavaScript clients",
			fieldName, int64(minimum), int64(maximum), minSafeInt, maxSafeInt)
	}
}
