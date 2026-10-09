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
package inspector

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"iter"

	astinspector "golang.org/x/tools/go/ast/inspector"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/extractjsontags"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/markers"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils"
	markersconsts "sigs.k8s.io/kube-api-linter/pkg/markers"
)

// Field is the data yielded for each struct field during inspection.
type Field struct {
	// Field is the AST node of the field being inspected.
	Field *ast.Field

	// JSONTagInfo holds the parsed JSON tag information for the field.
	JSONTagInfo extractjsontags.FieldTagInfo

	// Markers provides access to the markers for the package being inspected.
	Markers markers.Markers

	// QualifiedFieldName is the field name qualified with its struct name (e.g. "MyStruct.MyField").
	QualifiedFieldName string
}

// TypeSpec is the data yielded for each type spec during inspection.
type TypeSpec struct {
	// TypeSpec is the AST node of the type spec being inspected.
	TypeSpec *ast.TypeSpec

	// Markers provides access to the markers for the package being inspected.
	Markers markers.Markers
}

// Inspector is an interface that allows for the inspection of fields in structs.
type Inspector interface {
	// Fields returns an iterator over fields in structs, ignoring any struct that is not a type
	// declaration, and any field that is ignored and therefore would not be included in the CRD spec.
	Fields() iter.Seq[Field]

	// FieldsIncludingListTypes returns an iterator over fields in structs, including list types.
	// Unlike Fields, it does not skip fields in list type structs.
	FieldsIncludingListTypes() iter.Seq[Field]

	// TypeSpecs returns an iterator over the type specs in the package.
	TypeSpecs() iter.Seq[TypeSpec]
}

// inspector implements the Inspector interface.
type inspector struct {
	inspector *astinspector.Inspector
	jsonTags  extractjsontags.StructFieldTags
	markers   markers.Markers
}

// newInspector creates a new inspector.
func newInspector(astinspector *astinspector.Inspector, jsonTags extractjsontags.StructFieldTags, markers markers.Markers) Inspector {
	return &inspector{
		inspector: astinspector,
		jsonTags:  jsonTags,
		markers:   markers,
	}
}

// Fields returns an iterator over fields in structs, ignoring any struct that is not a type declaration,
// and any field that is ignored and therefore would not be included in the CRD spec.
func (i *inspector) Fields() iter.Seq[Field] {
	return i.fields(true)
}

// FieldsIncludingListTypes returns an iterator over fields in structs, including list types.
// Unlike Fields, it does not skip fields in list type structs.
func (i *inspector) FieldsIncludingListTypes() iter.Seq[Field] {
	return i.fields(false)
}

// fields is a shared implementation for field iteration.
// The skipListTypes parameter controls whether list type structs should be skipped.
func (i *inspector) fields(skipListTypes bool) iter.Seq[Field] {
	return func(yield func(Field) bool) {
		// Cursor.Preorder yields a Cursor per field, which provides O(1) access to enclosing
		// nodes (Parent/Enclosing) without manual stack bookkeeping.
		for c := range i.inspector.Root().Preorder((*ast.Field)(nil)) {
			field, ok := c.Node().(*ast.Field)
			if !ok {
				continue
			}

			// The field's parent is the enclosing *ast.FieldList; its parent is the type
			// (struct, interface, func signature, ...) that owns the field.
			structCursor := c.Parent().Parent()

			structType, ok := structCursor.Node().(*ast.StructType)
			if !ok || !isTopLevelTypeDecl(c) {
				continue
			}

			if skipListTypes && utils.IsKubernetesListType(structType, "") {
				continue
			}

			// Skip ignored or schemaless fields, as well as any field nested within one.
			if i.isFieldSkipped(c) {
				continue
			}

			if !i.yieldFieldWithRecovery(field, qualifiedFieldName(field, structCursor), yield) {
				return
			}
		}
	}
}

// isFieldSkipped reports whether the field at the cursor, or any field enclosing it,
// should be skipped (ignored or schemaless). Fields nested within a skipped field are
// themselves skipped, matching a traversal that prunes skipped subtrees.
func (i *inspector) isFieldSkipped(c astinspector.Cursor) bool {
	for fc := range c.Enclosing((*ast.Field)(nil)) {
		field, ok := fc.Node().(*ast.Field)
		if ok && i.shouldSkipField(field) {
			return true
		}
	}

	return false
}

// isTopLevelTypeDecl reports whether the cursor is enclosed by a top-level type declaration,
// i.e. a `type` GenDecl that is a direct child of the file. This excludes types declared
// within functions or non-type (var/const) declarations.
func isTopLevelTypeDecl(c astinspector.Cursor) bool {
	for genDecl := range c.Enclosing((*ast.GenDecl)(nil)) {
		decl, ok := genDecl.Node().(*ast.GenDecl)
		if !ok || decl.Tok != token.TYPE {
			return false
		}

		_, ok = genDecl.Parent().Node().(*ast.File)

		return ok
	}

	return false
}

// qualifiedFieldName returns the field name qualified with its struct name (e.g. "MyStruct.MyField").
// The struct name is taken from the enclosing *ast.TypeSpec when the struct is a named type;
// fields of anonymous (inline) structs are left unqualified.
func qualifiedFieldName(field *ast.Field, structCursor astinspector.Cursor) string {
	name := utils.FieldName(field)
	if name == "" {
		name = types.ExprString(field.Type)
	}

	if typeSpec, ok := structCursor.Parent().Node().(*ast.TypeSpec); ok {
		name = fmt.Sprintf("%s.%s", typeSpec.Name.Name, name)
	}

	return name
}

// shouldSkipField checks if a field should be skipped.
func (i *inspector) shouldSkipField(field *ast.Field) bool {
	tagInfo := i.jsonTags.FieldTags(field)
	if tagInfo.Ignored {
		return true
	}

	markerSet := i.markers.FieldMarkers(field)

	return isSchemalessType(markerSet)
}

// yieldFieldWithRecovery yields a field to the consumer with panic recovery.
// It returns whether the iteration should proceed.
func (i *inspector) yieldFieldWithRecovery(field *ast.Field, qualifiedFieldName string, yield func(Field) bool) (proceed bool) {
	tagInfo := i.jsonTags.FieldTags(field)

	defer func() {
		if r := recover(); r != nil {
			// If the consumer panics, we recover and log information that will help identify the issue.
			debug := printDebugInfo(field)
			panic(fmt.Sprintf("%s %v", debug, r)) // Re-panic to propagate the error.
		}
	}()

	return yield(Field{
		Field:              field,
		JSONTagInfo:        tagInfo,
		Markers:            i.markers,
		QualifiedFieldName: qualifiedFieldName,
	})
}

// TypeSpecs returns an iterator over the type specs in the package.
func (i *inspector) TypeSpecs() iter.Seq[TypeSpec] {
	return func(yield func(TypeSpec) bool) {
		// All yields *ast.TypeSpec nodes directly, so no node-type assertion is needed.
		for typeSpec := range astinspector.All[*ast.TypeSpec](i.inspector) {
			if !yield(TypeSpec{TypeSpec: typeSpec, Markers: i.markers}) {
				return
			}
		}
	}
}

func isSchemalessType(markerSet markers.MarkerSet) bool {
	// Check if the field is marked as schemaless.
	schemalessMarker := markerSet.Get(markersconsts.KubebuilderSchemaLessMarker)
	return len(schemalessMarker) > 0
}

// printDebugInfo prints debug information about the field that caused a panic during inspection.
// This function is designed to allow us to help identify which fields are causing issues during inspection.
func printDebugInfo(field *ast.Field) string {
	var debug string

	debug += fmt.Sprintf("Panic observed while inspecting field: %v (type: %v)\n", utils.FieldName(field), field.Type)

	return debug
}
