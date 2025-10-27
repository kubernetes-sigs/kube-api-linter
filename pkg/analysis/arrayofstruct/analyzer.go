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
package arrayofstruct

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

const name = "arrayofstruct"

// Analyzer is the analyzer for the arrayofstruct package.
// It checks that arrays containing structs have at least one required field.
var Analyzer = &analysis.Analyzer{
	Name:     name,
	Doc:      "Arrays containing structs must have at least one required field to prevent ambiguous YAML representations",
	Run:      run,
	Requires: []*analysis.Analyzer{inspector.Analyzer},
}

func run(pass *analysis.Pass) (any, error) {
	inspect, ok := pass.ResultOf[inspector.Analyzer].(inspector.Inspector)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetInspector
	}

	inspect.InspectFields(func(field *ast.Field, stack []ast.Node, jsonTagInfo extractjsontags.FieldTagInfo, markersAccess markershelper.Markers) {
		checkField(pass, field, markersAccess)
	})

	return nil, nil //nolint:nilnil
}

func checkField(pass *analysis.Pass, field *ast.Field, markersAccess markershelper.Markers) {
	// Get the element type of the array
	elementType := getArrayElementType(pass, field)
	if elementType == nil {
		return
	}

	// Check if this is an array of objects (not primitives)
	if !isObjectType(pass, elementType) {
		return
	}

	// Handle pointer types (e.g., []*MyStruct)
	if starExpr, ok := elementType.(*ast.StarExpr); ok {
		elementType = starExpr.X
	}

	// Get the struct type definition
	structType := getStructType(pass, elementType)
	if structType == nil {
		return
	}

	// Check if at least one field in the struct has a required marker
	if hasRequiredField(structType, markersAccess) {
		return
	}

	// Report the issue with a suggested fix
	reportArrayOfStructIssue(pass, field, structType, markersAccess)
}

// getArrayElementType extracts the element type from an array field.
// Returns nil if the field is not an array.
func getArrayElementType(pass *analysis.Pass, field *ast.Field) ast.Expr {
	switch fieldType := field.Type.(type) {
	case *ast.ArrayType:
		return fieldType.Elt
	case *ast.Ident:
		// For type aliases to arrays, we need to resolve the underlying type
		typeSpec, ok := utils.LookupTypeSpec(pass, fieldType)
		if !ok {
			return nil
		}

		arrayType, ok := typeSpec.Type.(*ast.ArrayType)
		if !ok {
			return nil
		}

		return arrayType.Elt
	default:
		return nil
	}
}

// reportArrayOfStructIssue reports a diagnostic for an array of structs without required fields.
func reportArrayOfStructIssue(pass *analysis.Pass, field *ast.Field, structType *ast.StructType, markersAccess markershelper.Markers) {
	fieldName := utils.FieldName(field)
	structName := utils.GetStructNameForField(pass, field)

	var prefix string
	if structName != "" {
		prefix = fmt.Sprintf("field %s in struct %s", fieldName, structName)
	} else {
		prefix = fmt.Sprintf("field %s", fieldName)
	}

	message := fmt.Sprintf("%s is an array of structs, but the struct has no required fields. At least one field should be marked as %s to prevent ambiguous YAML configurations", prefix, markers.RequiredMarker)

	// Create suggested fix to add +required marker and remove +optional marker from the first field
	suggestedFix := createSuggestedFix(structType, markersAccess)

	if suggestedFix != nil {
		pass.Report(analysis.Diagnostic{
			Pos:            field.Pos(),
			Message:        message,
			SuggestedFixes: []analysis.SuggestedFix{*suggestedFix},
		})
	} else {
		pass.Reportf(field.Pos(), "%s", message)
	}
}

// isObjectType checks if the given expression represents an object type (not a primitive).
func isObjectType(pass *analysis.Pass, expr ast.Expr) bool {
	switch et := expr.(type) {
	case *ast.StructType:
		// Inline struct definition
		return true
	case *ast.Ident:
		// Check if it's a basic type
		if utils.IsBasicType(pass, et) {
			return false
		}
		// It's a named type, check if it's a struct
		typeSpec, ok := utils.LookupTypeSpec(pass, et)
		if !ok {
			// Might be from another package, assume it's an object
			return true
		}
		// Recursively check the underlying type
		return isObjectType(pass, typeSpec.Type)
	case *ast.StarExpr:
		// Pointer to something, check what it points to
		return isObjectType(pass, et.X)
	case *ast.SelectorExpr:
		// Type from another package, assume it's an object
		return true
	default:
		return false
	}
}

// getStructType resolves the given expression to a struct type,
// following type aliases and handling inline structs.
func getStructType(pass *analysis.Pass, expr ast.Expr) *ast.StructType {
	switch et := expr.(type) {
	case *ast.StructType:
		// Inline struct definition
		return et
	case *ast.Ident:
		// Named struct type or type alias
		typeSpec, ok := utils.LookupTypeSpec(pass, et)
		if !ok {
			// This might be a type from another package or a built-in type
			// In this case, we can't inspect it, so we return nil
			return nil
		}

		// Check if the typeSpec.Type is a struct
		if structType, ok := typeSpec.Type.(*ast.StructType); ok {
			return structType
		}

		// If not a struct, it might be an alias to another type
		// Recursively resolve it
		return getStructType(pass, typeSpec.Type)
	case *ast.SelectorExpr:
		// Type from another package, we can't inspect it
		return nil
	default:
		return nil
	}
}

// createSuggestedFix creates a suggested fix that adds a +required marker
// and removes any +optional markers from the first field in the struct.
func createSuggestedFix(structType *ast.StructType, markersAccess markershelper.Markers) *analysis.SuggestedFix {
	if structType.Fields == nil || len(structType.Fields.List) == 0 {
		return nil
	}

	firstField := structType.Fields.List[0]
	fieldMarkers := markersAccess.FieldMarkers(firstField)

	// Remove all optional markers and track which ones were removed
	textEdits, removedMarkers := removeOptionalMarkers(fieldMarkers)

	// Add the +required marker before the field
	textEdits = append(textEdits, analysis.TextEdit{
		Pos:     firstField.Pos(),
		End:     firstField.Pos(),
		NewText: fmt.Appendf(nil, "// +%s\n\t", markers.RequiredMarker),
	})

	return &analysis.SuggestedFix{
		Message:   buildFixMessage(removedMarkers),
		TextEdits: textEdits,
	}
}

// removeOptionalMarkers removes all optional marker types from a field
// and returns the text edits and a list of marker names that were removed.
func removeOptionalMarkers(fieldMarkers markershelper.MarkerSet) ([]analysis.TextEdit, []string) {
	var (
		textEdits      []analysis.TextEdit
		removedMarkers []string
	)

	// Define all optional marker types to remove
	optionalMarkerTypes := []string{
		markers.OptionalMarker,
		markers.KubebuilderOptionalMarker,
		markers.K8sOptionalMarker,
	}

	for _, markerType := range optionalMarkerTypes {
		markerList := fieldMarkers[markerType]
		if len(markerList) > 0 {
			removedMarkers = append(removedMarkers, markerType)

			for _, marker := range markerList {
				textEdits = append(textEdits, analysis.TextEdit{
					Pos:     marker.Pos,
					End:     marker.End + 1, // +1 to include the newline
					NewText: nil,
				})
			}
		}
	}

	return textEdits, removedMarkers
}

// buildFixMessage constructs the suggested fix message based on which markers were removed.
func buildFixMessage(removedMarkers []string) string {
	if len(removedMarkers) == 0 {
		return fmt.Sprintf("Add `// +%s` marker to the first field", markers.RequiredMarker)
	}

	if len(removedMarkers) == 1 {
		return fmt.Sprintf("Add `// +%s` marker and remove `// +%s` marker from the first field",
			markers.RequiredMarker, removedMarkers[0])
	}

	return fmt.Sprintf("Add `// +%s` marker and remove optional markers from the first field",
		markers.RequiredMarker)
}

// hasRequiredField checks if at least one field in the struct has a required marker.
func hasRequiredField(structType *ast.StructType, markersAccess markershelper.Markers) bool {
	if structType.Fields == nil {
		return false
	}

	for _, field := range structType.Fields.List {
		fieldMarkers := markersAccess.FieldMarkers(field)

		// Check for any of the required markers
		if fieldMarkers.Has(markers.RequiredMarker) ||
			fieldMarkers.Has(markers.KubebuilderRequiredMarker) ||
			fieldMarkers.Has(markers.K8sRequiredMarker) {
			return true
		}
	}

	return false
}
