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
package b

// ============================================================================
// Custom marker tests
// These test custom marker configurations
// ============================================================================

// Custom marker that should only apply to fields
type CustomFieldMarkerTest struct {
	// +custom:field-only
	ValidFieldMarker string `json:"validFieldMarker"`
}

// +custom:field-only // want `marker "custom:field-only" can only be applied to fields`
type InvalidCustomFieldMarkerOnType struct {
	Name string `json:"name"`
}

// Custom marker that should only apply to types
// +custom:type-only
type ValidCustomTypeMarker struct {
	Name string `json:"name"`
}

type CustomTypeMarkerTest struct {
	// +custom:type-only // want `marker "custom:type-only" can only be applied to types`
	InvalidTypeMarker string `json:"invalidTypeMarker"`
}

// Custom marker with type constraints
type CustomTypeConstraintTest struct {
	// Valid: string type field with string-only custom marker
	// +custom:string-only
	ValidStringField string `json:"validStringField"`

	// Invalid: integer type field with string-only custom marker
	// +custom:string-only // want `marker "custom:string-only": type integer is not allowed \(expected one of: \[string\]\)`
	InvalidIntegerField int `json:"invalidIntegerField"`

	// Valid: integer type field with integer-only custom marker
	// +custom:integer-only
	ValidIntegerField int `json:"validIntegerField"`

	// Invalid: string type field with integer-only custom marker
	// +custom:integer-only // want `marker "custom:integer-only": type string is not allowed \(expected one of: \[integer\]\)`
	InvalidStringField string `json:"invalidStringField"`
}

// Custom marker with array element type constraints
type CustomArrayConstraintTest struct {
	// Valid: array of strings with string element constraint
	// +custom:string-array
	ValidStringArray []string `json:"validStringArray"`

	// Invalid: array of integers with string element constraint
	// +custom:string-array // want `marker "custom:string-array": array element: type integer is not allowed \(expected one of: \[string\]\)`
	InvalidIntegerArray []int `json:"invalidIntegerArray"`

	// Invalid: not an array type with array constraint
	// +custom:string-array // want `marker "custom:string-array": type string is not allowed \(expected one of: \[array\]\)`
	InvalidNonArray string `json:"invalidNonArray"`
}
