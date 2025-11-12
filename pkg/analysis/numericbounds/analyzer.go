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

	"golang.org/x/exp/constraints"
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

// Type bounds for validation.
const (
	maxInt32   = 2147483647                                      // 2^31 - 1
	minInt32   = -2147483648                                     // -2^31
	maxFloat32 = 3.40282346638528859811704183484516925440e+38    // max float32
	minFloat32 = -3.40282346638528859811704183484516925440e+38   // min float32
	maxFloat64 = 1.797693134862315708145274237317043567981e+308  // max float64
	minFloat64 = -1.797693134862315708145274237317043567981e+308 // min float64
)

// JavaScript safe integer bounds for int64 (±2^53-1).
// Per Kubernetes API conventions, int64 fields should use bounds within this range
// to ensure compatibility with JavaScript clients.
const (
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

	inspect.InspectFields(func(field *ast.Field, _ extractjsontags.FieldTagInfo, markersAccess markershelper.Markers, qualifiedFieldName string) {
		// Create TypeChecker with closure capturing markersAccess and qualifiedFieldName
		typeChecker := utils.NewTypeChecker(func(pass *analysis.Pass, ident *ast.Ident, node ast.Node, _ string) {
			checkNumericType(pass, ident, node, markersAccess, qualifiedFieldName)
		})

		typeChecker.CheckNode(pass, field)
	})

	return nil, nil //nolint:nilnil
}

//nolint:cyclop
func checkNumericType(pass *analysis.Pass, ident *ast.Ident, node ast.Node, markersAccess markershelper.Markers, qualifiedFieldName string) {
	// Only check int32, int64, float32, and float64 types
	if ident.Name != "int32" && ident.Name != "int64" && ident.Name != "float32" && ident.Name != "float64" {
		return
	}

	field, ok := node.(*ast.Field)
	if !ok {
		return
	}

	fieldMarkers := utils.TypeAwareMarkerCollectionForField(pass, markersAccess, field)

	// Check if this is an array/slice field
	isSlice := utils.IsArrayTypeOrAlias(pass, field)

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
		pass.Reportf(field.Pos(), "%s has an invalid minimum marker: %v", qualifiedFieldName, minErr)
	}

	if maxErr != nil && !maxMissing {
		pass.Reportf(field.Pos(), "%s has an invalid maximum marker: %v", qualifiedFieldName, maxErr)
	}

	// Report if markers are missing
	if minMissing {
		pass.Reportf(field.Pos(), "%s is missing minimum bound validation marker", qualifiedFieldName)
	}

	if maxMissing {
		pass.Reportf(field.Pos(), "%s is missing maximum bound validation marker", qualifiedFieldName)
	}

	// If any markers are missing or invalid, don't continue with bounds checks
	if minErr != nil || maxErr != nil {
		return
	}

	// Validate bounds are within the type's valid range
	checkBoundsWithinTypeRange(pass, field, qualifiedFieldName, ident.Name, minimum, maximum)
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
// Precedence: Markers checked in the order provided and first valid value found is returned.
// We require a valid numeric value (not just marker presence) for both minimum and maximum markers.
func getMarkerNumericValue(markerSet markershelper.MarkerSet, markerNames []string) (float64, error) {
	for _, markerName := range markerNames {
		markerList := markerSet.Get(markerName)
		if len(markerList) == 0 {
			continue
		}

		// Use the exported utils.GetMarkerNumericValue function to parse the marker value
		value, err := utils.GetMarkerNumericValue[float64](markerList[0])
		if err != nil {
			if errors.Is(err, errMarkerMissingValue) {
				continue
			}

			return 0, fmt.Errorf("error getting marker value: %w", err)
		}

		return value, nil
	}

	return 0, errMarkerMissingValue
}

// checkBoundsWithinTypeRange validates that the bounds are within the valid range for the type.
// For int64, enforces JavaScript-safe bounds as per Kubernetes API conventions to ensure
// compatibility with JavaScript clients.
func checkBoundsWithinTypeRange(pass *analysis.Pass, field *ast.Field, prefix, typeName string, minimum, maximum float64) {
	switch typeName {
	case "int32":
		checkBoundInRange(pass, field, prefix, minimum, minInt32, maxInt32, "minimum", "int32")
		checkBoundInRange(pass, field, prefix, maximum, minInt32, maxInt32, "maximum", "int32")
	case "int64":
		// K8s API conventions enforce JavaScript-safe bounds for int64 (±2^53-1)
		checkBoundInRange(pass, field, prefix, minimum, int64(minSafeInt64), int64(maxSafeInt64), "minimum", "JavaScript-safe int64",
			"Consider using a string type to avoid precision loss in JavaScript clients")
		checkBoundInRange(pass, field, prefix, maximum, int64(minSafeInt64), int64(maxSafeInt64), "maximum", "JavaScript-safe int64",
			"Consider using a string type to avoid precision loss in JavaScript clients")
	case "float32":
		checkBoundInRange(pass, field, prefix, minimum, minFloat32, maxFloat32, "minimum", "float32")
		checkBoundInRange(pass, field, prefix, maximum, minFloat32, maxFloat32, "maximum", "float32")
	case "float64":
		checkBoundInRange(pass, field, prefix, minimum, minFloat64, maxFloat64, "minimum", "float64")
		checkBoundInRange(pass, field, prefix, maximum, minFloat64, maxFloat64, "maximum", "float64")
	}
}

// checkBoundInRange checks if a bound value is within the valid range.
// Uses generics to work with both integer and float types.
func checkBoundInRange[T constraints.Integer | constraints.Float](pass *analysis.Pass, field *ast.Field, prefix string, value float64, minBound, maxBound T, boundType, typeName string, extraMsg ...string) {
	if value < float64(minBound) || value > float64(maxBound) {
		msg := fmt.Sprintf("%s has %s bound %%v that is outside the %s range [%%v, %%v]", prefix, boundType, typeName)
		if len(extraMsg) > 0 {
			msg += ". " + extraMsg[0]
		}

		pass.Reportf(field.Pos(), msg, value, minBound, maxBound)
	}
}
