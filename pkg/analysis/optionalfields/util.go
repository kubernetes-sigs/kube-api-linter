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
package optionalfields

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/extractjsontags"
	markershelper "sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/markers"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils"
	"sigs.k8s.io/kube-api-linter/pkg/config"
	"sigs.k8s.io/kube-api-linter/pkg/markers"
)

// isStarExpr checks if the expression is a pointer type.
// If it is, it returns the expression inside the pointer.
func isStarExpr(expr ast.Expr) (bool, ast.Expr) {
	if ptrType, ok := expr.(*ast.StarExpr); ok {
		return true, ptrType.X
	}

	return false, expr
}

// isPointerType checks if the expression is a pointer type.
// This is for types that are always implemented as pointers and therefore should
// not be the underlying type of a star expr.
func isPointerType(expr ast.Expr) bool {
	switch expr.(type) {
	case *ast.StarExpr, *ast.MapType, *ast.ArrayType:
		return true
	default:
		return false
	}
}

// structContainsRequiredFields checks whether the struct has any required fields.
// Having a required field means that `{}` is not a valid entity for the struct,
// and therefore the struct should be a pointer when optional.
func structContainsRequiredFields(structType *ast.StructType, markersAccess markershelper.Markers) bool {
	if structType == nil {
		return false
	}

	for _, field := range structType.Fields.List {
		fieldMarkers := markersAccess.FieldMarkers(field)

		if isFieldRequired(fieldMarkers) {
			return true
		}
	}

	return false
}

// isFieldRequired checks if a field has a required marker.
func isFieldRequired(fieldMarkers markershelper.MarkerSet) bool {
	return fieldMarkers.Has(requiredMarker) || fieldMarkers.Has(kubebuilderRequiredMarker)
}

// isFieldOptional checks if a field has an optional marker.
func isFieldOptional(fieldMarkers markershelper.MarkerSet) bool {
	return fieldMarkers.Has(optionalMarker) || fieldMarkers.Has(kubebuilderOptionalMarker)
}

// reportShouldAddPointer adds an analysis diagnostic that explains that a pointer should be added.
// Where the pointer policy is suggest fix, it also adds a suggested fix to add the pointer.
func reportShouldAddPointer(pass *analysis.Pass, field *ast.Field, pointerPolicy config.OptionalFieldsPointerPolicy, fieldName, messageFmt string) {
	switch pointerPolicy {
	case config.OptionalFieldsPointerPolicySuggestFix:
		pass.Report(analysis.Diagnostic{
			Pos:     field.Pos(),
			Message: fmt.Sprintf(messageFmt, fieldName),
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "should make the field a pointer",
					TextEdits: []analysis.TextEdit{
						{
							Pos:     field.Pos() + token.Pos(len(fieldName)+1),
							NewText: []byte("*"),
						},
					},
				},
			},
		})
	case config.OptionalFieldsPointerPolicyWarn:
		pass.Reportf(field.Pos(), messageFmt, fieldName)
	default:
		panic(fmt.Sprintf("unknown pointer policy: %s", pointerPolicy))
	}
}

// reportShouldRemovePointer adds an analysis diagnostic that explains that a pointer should be removed.
// Where the pointer policy is suggest fix, it also adds a suggested fix to remove the pointer.
func reportShouldRemovePointer(pass *analysis.Pass, field *ast.Field, pointerPolicy config.OptionalFieldsPointerPolicy, fieldName, messageFmt string) {
	switch pointerPolicy {
	case config.OptionalFieldsPointerPolicySuggestFix:
		pass.Report(analysis.Diagnostic{
			Pos:     field.Pos(),
			Message: fmt.Sprintf(messageFmt, fieldName),
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "should remove the pointer",
					TextEdits: []analysis.TextEdit{
						{
							Pos: field.Pos() + token.Pos(len(fieldName)+1),
							End: field.Pos() + token.Pos(len(fieldName)+2),
						},
					},
				},
			},
		})
	case config.OptionalFieldsPointerPolicyWarn:
		pass.Reportf(field.Pos(), messageFmt, fieldName)
	default:
		panic(fmt.Sprintf("unknown pointer policy: %s", pointerPolicy))
	}
}

// reportShouldRemoveAllInstancesOfIntegerMarker adds an analysis diagnostic that explains that a marker should be removed.
// This function is used to find non-zero valued markers, and suggest that they are removed when the field is not a pointer.
// This is used for markers like MinLength and MinProperties.
func reportShouldRemoveAllInstancesOfIntegerMarker(pass *analysis.Pass, field *ast.Field, markersAccess markershelper.Markers, markerName, fieldName, messageFmt string) {
	fieldMarkers := markersAccess.FieldMarkers(field)

	for _, marker := range fieldMarkers.Get(markerName) {
		markerValue, err := getMarkerIntegerValue(marker)
		if err != nil {
			pass.Reportf(marker.Pos, "invalid value for %s marker: %v", markerName, err)
			return
		}

		if markerValue > 0 {
			reportShouldRemoveMarker(pass, field, marker, fieldName, messageFmt)
		}
	}
}

// reportShouldAddMarker adds an analysis diagnostic that explains that a marker should be removed.
// This is used where we see a marker that would conflict with a field that lacks omitempty.
func reportShouldRemoveMarker(pass *analysis.Pass, field *ast.Field, marker markershelper.Marker, fieldName, messageFmt string) {
	pass.Report(analysis.Diagnostic{
		Pos:     field.Pos(),
		Message: fmt.Sprintf(messageFmt, fieldName),
		SuggestedFixes: []analysis.SuggestedFix{
			{
				Message: fmt.Sprintf("should remove the marker: %s", marker.RawComment),
				TextEdits: []analysis.TextEdit{
					{
						Pos: marker.Pos,
						End: marker.End + 1,
					},
				},
			},
		},
	})
}

// reportShouldAddOmitEmpty adds an analysis diagnostic that explains that an omitempty tag should be added.
func reportShouldAddOmitEmpty(pass *analysis.Pass, field *ast.Field, fieldName, messageFmt string, fieldTagInfo extractjsontags.FieldTagInfo) {
	pass.Report(analysis.Diagnostic{
		Pos:     field.Pos(),
		Message: fmt.Sprintf(messageFmt, fieldName),
		SuggestedFixes: []analysis.SuggestedFix{
			{
				Message: fmt.Sprintf("should add 'omitempty' to the field tag for field %s", fieldName),
				TextEdits: []analysis.TextEdit{
					{
						Pos:     fieldTagInfo.Pos + token.Pos(len(fieldTagInfo.Name)),
						NewText: []byte(",omitempty"),
					},
				},
			},
		},
	})
}

// getMarkerIntegerValueByName extracts the numeric value from the first
// instace of the marker with the given name.
// Works for markers like MaxLength, MinLength, etc.
func getMarkerIntegerValueByName(marker markershelper.MarkerSet, markerName string) (*int, error) {
	markerList := marker.Get(markerName)
	if len(markerList) == 0 {
		return nil, errMarkerMissingValue
	}

	markerValue, err := getMarkerIntegerValue(markerList[0])
	if err != nil {
		return nil, fmt.Errorf("error getting marker value: %w", err)
	}

	return &markerValue, nil
}

// getMarkerIntegerValue extracts a numeric value from the default
// value of a marker.
// Works for markers like MaxLength, MinLength, etc.
func getMarkerIntegerValue(marker markershelper.Marker) (int, error) {
	rawValue, ok := marker.Expressions[""]
	if !ok {
		return 0, errMarkerMissingValue
	}

	value, err := strconv.Atoi(rawValue)
	if err != nil {
		return 0, fmt.Errorf("error converting value to integer: %w", err)
	}

	return value, nil
}

// getMarkerFloatValueByName extracts the numeric value from the first
// instace of the marker with the given name.
// Works for markers like MaxLength, MinLength, etc.
func getMarkerFloatValueByName(marker markershelper.MarkerSet, markerName string) (*float64, error) {
	markerList := marker.Get(markerName)
	if len(markerList) == 0 {
		return nil, errMarkerMissingValue
	}

	markerValue, err := getMarkerFloatValue(markerList[0])
	if err != nil {
		return nil, fmt.Errorf("error getting marker value: %w", err)
	}

	return &markerValue, nil
}

// getMarkerFloatValue extracts a numeric value from the default
// value of a marker.
// Works for markers like MaxLength, MinLength, etc.
func getMarkerFloatValue(marker markershelper.Marker) (float64, error) {
	rawValue, ok := marker.Expressions[""]
	if !ok {
		return 0, errMarkerMissingValue
	}

	value, err := strconv.ParseFloat(rawValue, 64)
	if err != nil {
		return 0, fmt.Errorf("error converting value to float: %w", err)
	}

	return value, nil
}

// structHasGreaterThanZeroMaxProperties checks if the struct has a minProperties marker.
// It inspects the minProperties marker on the struct type itself.
func structHasGreaterThanZeroMinProperties(structType *ast.StructType, structMarkers markershelper.MarkerSet) (bool, error) {
	if structType == nil {
		return false, nil
	}

	for _, marker := range structMarkers.Get(minPropertiesMarker) {
		markerValue, err := getMarkerIntegerValue(marker)
		if err != nil {
			return false, fmt.Errorf("error getting marker value: %w", err)
		}

		if markerValue > 0 {
			return true, nil
		}
	}

	return false, nil
}

func integerRangeIncludesZero(minimum, maximum *int) bool {
	return ptr.Deref(minimum, -1) == 0 ||
		ptr.Deref(maximum, -1) == 0 ||
		ptr.Deref(minimum, 0) < 0 && ptr.Deref(maximum, 0) > 0
}

func floatRangeIncludesZero(minimum, maximum *float64) bool {
	return ptr.Deref(minimum, -1) == 0 ||
		ptr.Deref(maximum, -1) == 0 ||
		ptr.Deref(minimum, 0) < 0 && ptr.Deref(maximum, 0) > 0
}

func isZeroValueValid(pass *analysis.Pass, field *ast.Field, typeExpr ast.Expr, markersAccess markershelper.Markers, fieldTagInfo extractjsontags.FieldTagInfo) bool {
	if fieldTagInfo.OmitEmpty {
		// If the field is omitted, we can use a zero value.
		// For structs, if they aren't a pointer another error will be raised.
		return true
	}

	isPointer, underlyingType := isStarExpr(typeExpr)
	if isPointer {
		// The field is a pointer without omitempty, so we cannot use a zero value unless the field is nullable.
		return markersAccess.FieldMarkers(field).Has(markers.NullableMarker)
	}

	switch t := underlyingType.(type) {
	case *ast.StructType:
		// For structs, we have to check if there are any non-omitted fields, that do not accept a zero value.
		return isStructZeroValueValid(pass, t, markersAccess)
	case *ast.Ident:
		return isIdentZeroValueValid(pass, field, t, markersAccess)
	case *ast.MapType:
		return isMapZeroValueValid(field, markersAccess)
	case *ast.ArrayType:
		// For arrays, we can use a zero value if the array is not required to have a minimum number of items.
		return isArrayZeroValueValid(field, t, markersAccess)
	}

	// For other types, we assume that zero value is valid.
	return true
}

func isStructZeroValueValid(pass *analysis.Pass, structType *ast.StructType, markersAccess markershelper.Markers) bool {
	if structType == nil {
		return true
	}

	jsonTagInfo, ok := pass.ResultOf[extractjsontags.Analyzer].(extractjsontags.StructFieldTags)
	if !ok {
		panic("could not get struct field tags from pass result")
	}

	for _, field := range structType.Fields.List {
		fieldTagInfo := jsonTagInfo.FieldTags(field)

		if !isZeroValueValid(pass, field, field.Type, markersAccess, fieldTagInfo) {
			return false
		}
	}

	return true
}

func isIdentZeroValueValid(pass *analysis.Pass, field *ast.Field, ident *ast.Ident, markersAccess markershelper.Markers) bool {
	if ident == nil {
		return true
	}

	// Check if the identifier is a known type that can have a zero value.
	switch ident.Name {
	case "string":
		return isStringZeroValueValid(field, markersAccess)
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64":
		return isIntegerZeroValueValid(field, markersAccess)
	case "float32", "float64":
		return isFloatZeroValueValid(field, markersAccess)
	case "bool":
		// For bool, we can always use a zero value.
		return true
	}

	// If the ident isn't one of the above, check the underlying type spec.
	typeSpec, ok := utils.LookupTypeSpec(pass, ident)
	if !ok {
		return false
	}

	jsonTagInfo, ok := pass.ResultOf[extractjsontags.Analyzer].(extractjsontags.StructFieldTags)
	if !ok {
		panic("could not get struct field tags from pass result")
	}

	return isZeroValueValid(pass, field, typeSpec.Type, markersAccess, jsonTagInfo.FieldTags(field))
}

// isStringZeroValueValid checks if a string field can have a zero value.
// This would be true when either there is no minimum length marker, or when the minimmum length marker is set to 0.
func isStringZeroValueValid(field *ast.Field, markersAccess markershelper.Markers) bool {
	fieldMarkers := markersAccess.FieldMarkers(field)

	if stringFieldIsEnum(fieldMarkers) {
		return enumFieldAllowsEmpty(fieldMarkers)
	}

	return !fieldMarkers.Has(markers.KubebuilderMinLengthMarker) || fieldMarkers.HasWithValue(fmt.Sprintf("%s=0", markers.KubebuilderMinLengthMarker))
}

// isIntegerZeroValueValid checks if an integer field can have a zero value.
func isIntegerZeroValueValid(field *ast.Field, markersAccess markershelper.Markers) bool {
	fieldMarkers := markersAccess.FieldMarkers(field)

	minimum, err := getMarkerIntegerValueByName(fieldMarkers, markers.KubebuilderMinimumMarker)
	if err != nil && !errors.Is(err, errMarkerMissingValue) {
		return false
	}

	maximum, err := getMarkerIntegerValueByName(fieldMarkers, markers.KubebuilderMaximumMarker)
	if err != nil && !errors.Is(err, errMarkerMissingValue) {
		return false
	}

	return ptr.Deref(minimum, -1) <= 0 && ptr.Deref(maximum, 1) >= 0
}

// isFloatZeroValueValid checks if a float field can have a zero value.
func isFloatZeroValueValid(field *ast.Field, markersAccess markershelper.Markers) bool {
	fieldMarkers := markersAccess.FieldMarkers(field)

	minimum, err := getMarkerFloatValueByName(fieldMarkers, markers.KubebuilderMinimumMarker)
	if err != nil && !errors.Is(err, errMarkerMissingValue) {
		return false
	}

	maximum, err := getMarkerFloatValueByName(fieldMarkers, markers.KubebuilderMaximumMarker)
	if err != nil && !errors.Is(err, errMarkerMissingValue) {
		return false
	}

	return ptr.Deref(minimum, -1) <= 0 && ptr.Deref(maximum, 1) >= 0
}

// isMapZeroValueValid checks if a map field can have a zero value.
// For maps, this means there is no minProperties marker, or the minProperties marker is set to 0.
func isMapZeroValueValid(field *ast.Field, markersAccess markershelper.Markers) bool {
	fieldMarkers := markersAccess.FieldMarkers(field)

	return !fieldMarkers.Has(markers.KubebuilderMinPropertiesMarker) || fieldMarkers.HasWithValue(fmt.Sprintf("%s=0", markers.KubebuilderMinPropertiesMarker))
}

// isArrayZeroValueValid checks if an array field can have a zero value.
func isArrayZeroValueValid(field *ast.Field, arrayType *ast.ArrayType, markersAccess markershelper.Markers) bool {
	// Arrays of bytes are special cased and treated as strings.
	if ident, ok := arrayType.Elt.(*ast.Ident); ok && ident.Name == "byte" {
		return isStringZeroValueValid(field, markersAccess)
	}

	fieldMarkers := markersAccess.FieldMarkers(field)

	// For arrays, we can use a zero value if the array is not required to have a minimum number of items.
	minItems, err := getMarkerIntegerValueByName(fieldMarkers, markers.KubebuilderMinItemsMarker)
	if err != nil && !errors.Is(err, errMarkerMissingValue) {
		return false
	}

	return minItems == nil || *minItems == 0
}

func stringFieldIsEnum(fieldMarkers markershelper.MarkerSet) bool {
	// Check if the field has a kubebuilder enum marker.
	return fieldMarkers.Has(enumMarker)
}

func enumFieldAllowsEmpty(fieldMarkers markershelper.MarkerSet) bool {
	// Check if the field has a kubebuilder enum marker with an empty value.
	enumMarker := fieldMarkers.Get(enumMarker)
	if len(enumMarker) == 0 {
		return false
	}

	for _, marker := range enumMarker {
		return slices.Contains(strings.Split(marker.Expressions[""], ";"), "\"\"")
	}

	return false
}
