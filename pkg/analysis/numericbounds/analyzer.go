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

const (
	name = "numericbounds"
)

// Type bounds for validation
const (
	maxInt32   = 2147483647                                    // 2^31 - 1
	minInt32   = -2147483648                                   // -2^31
	maxFloat32 = 3.40282346638528859811704183484516925440e+38  // max float32
	minFloat32 = -3.40282346638528859811704183484516925440e+38 // min float32
)

// JavaScript safe integer bounds (2^53 - 1 and -(2^53 - 1))
const (
	maxSafeInt32 = 2147483647        // 2^31 - 1 (fits in JS Number)
	minSafeInt32 = -2147483648       // -2^31 (fits in JS Number)
	maxSafeInt64 = 9007199254740991  // 2^53 - 1 (max safe integer in JS)
	minSafeInt64 = -9007199254740991 // -(2^53 - 1) (min safe integer in JS)
)

var errMarkerMissingValue = errors.New("marker value not found")

// Analyzer is the analyzer for the numericbounds package.
// It checks that numeric fields have appropriate bounds validation markers.
var Analyzer = &analysis.Analyzer{
	Name:     name,
	Doc:      "Checks that numeric fields (int32, int64, float32, float64) have appropriate minimum and maximum bounds validation markers",
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

	// Get the underlying numeric type identifier (int32, int64, float32, float64)
	ident := getNumericTypeIdent(pass, fieldType)
	if ident == nil {
		return
	}

	// Only check int32, int64, float32, and float64 types
	if ident.Name != "int32" && ident.Name != "int64" && ident.Name != "float32" && ident.Name != "float64" {
		return
	}

	// Create type description that clarifies array element types
	typeDesc := ident.Name
	if isSlice {
		typeDesc = fmt.Sprintf("array element type of %s", ident.Name)
	}

	fieldMarkers := utils.TypeAwareMarkerCollectionForField(pass, markersAccess, field)

	// Determine which markers to look for based on whether the field is a slice
	minMarkers, maxMarkers := getMarkerNames(isSlice)

	// Get minimum and maximum marker values
	minimum, minErr := getMarkerNumericValue(fieldMarkers, minMarkers)
	maximum, maxErr := getMarkerNumericValue(fieldMarkers, maxMarkers)

	// Check if markers are missing
	minMissing := errors.Is(minErr, errMarkerMissingValue)
	maxMissing := errors.Is(maxErr, errMarkerMissingValue)

	// Report any invalid marker values (e.g., non-numeric values)
	if minErr != nil && !minMissing {
		pass.Reportf(field.Pos(), "field %s has an invalid minimum marker: %v", fieldName, minErr)
	}
	if maxErr != nil && !maxMissing {
		pass.Reportf(field.Pos(), "field %s has an invalid maximum marker: %v", fieldName, maxErr)
	}

	// Report if markers are missing
	if minMissing {
		pass.Reportf(field.Pos(), "field %s %s is missing minimum bounds validation marker", fieldName, typeDesc)
	}
	if maxMissing {
		pass.Reportf(field.Pos(), "field %s %s is missing maximum bounds validation marker", fieldName, typeDesc)
	}

	// If any markers are missing or invalid, don't continue with bounds checks
	if minErr != nil || maxErr != nil {
		return
	}

	// Validate bounds are within the type's range
	checkBoundsWithinTypeRange(pass, field, fieldName, typeDesc, minimum, maximum)

	// For int64 fields, check if bounds are within JavaScript safe integer range
	checkJavaScriptSafeBounds(pass, field, fieldName, typeDesc, minimum, maximum)
}

// getNumericTypeIdent returns the identifier for numeric types (int32, int64, float32, float64).
// It handles type aliases by looking up the underlying type.
// Note: This function expects pointers and slices to already be unwrapped.
func getNumericTypeIdent(pass *analysis.Pass, expr ast.Expr) *ast.Ident {
	ident, ok := expr.(*ast.Ident)
	if !ok {
		return nil
	}

	// Check if it's a basic numeric type we care about
	if ident.Name == "int32" || ident.Name == "int64" || ident.Name == "float32" || ident.Name == "float64" {
		return ident
	}

	// Check if it's a type alias to a numeric type
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
// When the field is a slice, we extract the element type since that's what
// needs bounds validation (e.g., []int32 -> int32). The isSlice flag allows
// the caller to report errors with "array element type" for clarity.
func unwrapType(expr ast.Expr) (ast.Expr, bool) {
	isSlice := false

	// Unwrap pointer if present (e.g., *int32)
	if starExpr, ok := expr.(*ast.StarExpr); ok {
		expr = starExpr.X
	}

	// Check if it's a slice and unwrap to get element type (e.g., []int32 -> int32)
	if arrayType, ok := expr.(*ast.ArrayType); ok {
		isSlice = true
		expr = arrayType.Elt

		// Handle pointer inside slice (e.g., []*int32 -> int32)
		if starExpr, ok := expr.(*ast.StarExpr); ok {
			expr = starExpr.X
		}
	}

	return expr, isSlice
}

// getMarkerNames returns the appropriate minimum and maximum marker names
// based on whether the field is a slice.
// Returns both kubebuilder and k8s declarative validation markers.
func getMarkerNames(isSlice bool) (minMarkers, maxMarkers []string) {
	if isSlice {
		return []string{markers.KubebuilderItemsMinimumMarker}, []string{markers.KubebuilderItemsMaximumMarker}
	}
	return []string{markers.KubebuilderMinimumMarker, markers.K8sMinimumMarker}, []string{markers.KubebuilderMaximumMarker, markers.K8sMaximumMarker}
}

// getMarkerNumericValue extracts the numeric value from the first instance of any of the given marker names.
// Checks multiple marker names to support both kubebuilder and k8s declarative validation markers.
func getMarkerNumericValue(markerSet markershelper.MarkerSet, markerNames []string) (float64, error) {
	for _, markerName := range markerNames {
		markerList := markerSet.Get(markerName)
		if len(markerList) == 0 {
			continue
		}

		marker := markerList[0]
		rawValue, ok := marker.Expressions[""]
		if !ok {
			continue
		}

		// Parse as float64 using strconv for better error handling
		value, err := strconv.ParseFloat(rawValue, 64)
		if err != nil {
			return 0, fmt.Errorf("error converting value to number: %w", err)
		}

		return value, nil
	}

	return 0, errMarkerMissingValue
}

// checkBoundsWithinTypeRange validates that the bounds are within the valid range for the type.
func checkBoundsWithinTypeRange(pass *analysis.Pass, field *ast.Field, fieldName, typeDesc string, minimum, maximum float64) {
	// Extract the actual type name from typeDesc (e.g., "array element type of int32" -> "int32")
	typeName := extractTypeName(typeDesc)

	switch typeName {
	case "int32":
		if minimum < minInt32 || minimum > maxInt32 {
			pass.Reportf(field.Pos(), "field %s %s has minimum bound %v that is outside the valid int32 range [%d, %d]", fieldName, typeDesc, minimum, minInt32, maxInt32)
		}
		if maximum < minInt32 || maximum > maxInt32 {
			pass.Reportf(field.Pos(), "field %s %s has maximum bound %v that is outside the valid int32 range [%d, %d]", fieldName, typeDesc, maximum, minInt32, maxInt32)
		}
	case "float32":
		if minimum < minFloat32 || minimum > maxFloat32 {
			pass.Reportf(field.Pos(), "field %s %s has minimum bound %v that is outside the valid float32 range", fieldName, typeDesc, minimum)
		}
		if maximum < minFloat32 || maximum > maxFloat32 {
			pass.Reportf(field.Pos(), "field %s %s has maximum bound %v that is outside the valid float32 range", fieldName, typeDesc, maximum)
		}
	}
}

// checkJavaScriptSafeBounds checks if int64 bounds are within JavaScript safe integer range.
func checkJavaScriptSafeBounds(pass *analysis.Pass, field *ast.Field, fieldName, typeDesc string, minimum, maximum float64) {
	// Extract the actual type name from typeDesc
	typeName := extractTypeName(typeDesc)

	if typeName != "int64" {
		return
	}

	if minimum < minSafeInt64 || maximum > maxSafeInt64 {
		pass.Reportf(field.Pos(),
			"field %s %s has bounds [%d, %d] that exceed safe integer range [%d, %d]. Consider using a string type to avoid precision loss in JavaScript clients",
			fieldName, typeDesc, int64(minimum), int64(maximum), minSafeInt64, maxSafeInt64)
	}
}

// extractTypeName extracts the base type name from a type description.
// E.g., "array element type of int32" -> "int32", "int64" -> "int64"
func extractTypeName(typeDesc string) string {
	// Check if it's an array element type description
	const prefix = "array element type of "
	if len(typeDesc) > len(prefix) && typeDesc[:len(prefix)] == prefix {
		return typeDesc[len(prefix):]
	}
	return typeDesc
}
