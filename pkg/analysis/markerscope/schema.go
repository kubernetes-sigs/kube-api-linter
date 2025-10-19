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
package markerscope

import "go/types"

// SchemaType represents OpenAPI schema types that markers can target.
type SchemaType string

const (
	// SchemaTypeInteger represents integer types (int, int32, int64, uint, etc.)
	SchemaTypeInteger SchemaType = "integer"
	// SchemaTypeNumber represents floating-point types (float32, float64).
	SchemaTypeNumber SchemaType = "number"
	// SchemaTypeString represents string types.
	SchemaTypeString SchemaType = "string"
	// SchemaTypeBoolean represents boolean types.
	SchemaTypeBoolean SchemaType = "boolean"
	// SchemaTypeArray represents array/slice types.
	SchemaTypeArray SchemaType = "array"
	// SchemaTypeObject represents struct/map types.
	SchemaTypeObject SchemaType = "object"
)

// getSchemaType converts a Go type to its corresponding OpenAPI schema type.
func getSchemaType(t types.Type) SchemaType {
	t = unwrapType(t)

	switch ut := t.Underlying().(type) {
	case *types.Basic:
		return getBasicTypeSchema(ut)
	case *types.Slice, *types.Array:
		return SchemaTypeArray
	case *types.Map, *types.Struct:
		return SchemaTypeObject
	}

	return ""
}

// unwrapType unwraps pointer and named types to get the underlying type.
func unwrapType(t types.Type) types.Type {
	// Unwrap pointer types
	if ptr, ok := t.(*types.Pointer); ok {
		t = ptr.Elem()
	}

	// Unwrap named types to get underlying type
	if named, ok := t.(*types.Named); ok {
		t = named.Underlying()
	}

	return t
}

// getBasicTypeSchema returns the schema type for a basic Go type.
func getBasicTypeSchema(bt *types.Basic) SchemaType {
	switch bt.Kind() {
	case types.Bool:
		return SchemaTypeBoolean
	case types.Int, types.Int8, types.Int16, types.Int32, types.Int64,
		types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64:
		return SchemaTypeInteger
	case types.Float32, types.Float64:
		return SchemaTypeNumber
	case types.String:
		return SchemaTypeString
	case types.Invalid, types.Uintptr, types.Complex64, types.Complex128,
		types.UnsafePointer, types.UntypedBool, types.UntypedInt, types.UntypedRune,
		types.UntypedFloat, types.UntypedComplex, types.UntypedString, types.UntypedNil:
		// These types are not supported in OpenAPI schemas
		return ""
	default:
		return ""
	}
}

// getElementType returns the element type of an array or slice.
func getElementType(t types.Type) types.Type {
	// Unwrap pointer types
	if ptr, ok := t.(*types.Pointer); ok {
		t = ptr.Elem()
	}

	// Unwrap named types
	if named, ok := t.(*types.Named); ok {
		t = named.Underlying()
	}

	switch ut := t.Underlying().(type) {
	case *types.Slice:
		return ut.Elem()
	case *types.Array:
		return ut.Elem()
	}

	return nil
}

// getUnderlyingType recursively unwraps type to find the underlying type.
func getUnderlyingType(expr types.Type) types.Type {
	switch t := expr.(type) {
	case *types.Pointer:
		return getUnderlyingType(t.Elem())
	case *types.Named:
		return getUnderlyingType(t.Underlying())
	case *types.Alias:
		return getUnderlyingType(t.Underlying())
	default:
		return expr
	}
}
