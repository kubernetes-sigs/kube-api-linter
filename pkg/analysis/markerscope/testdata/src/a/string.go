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
package a

type StringType string

// Valid: Pattern on type
// +kubebuilder:validation:Pattern="^[a-z]+$"
type PatternType string

// Valid: MinLength on type
// +kubebuilder:validation:MinLength=1
type MinLengthType string

// Valid: MaxLength on type
// +kubebuilder:validation:MaxLength=100
type MaxLengthType string

// Invalid: Pattern marker on integer type
// +kubebuilder:validation:Pattern="^[0-9]+$" // want `marker "kubebuilder:validation:Pattern": type integer is not allowed \(expected one of: \[string\]\)`
type InvalidPatternOnIntType int32

// Invalid: MinLength marker on boolean type
// +kubebuilder:validation:MinLength=5 // want `marker "kubebuilder:validation:MinLength": type boolean is not allowed \(expected one of: \[string\]\)`
type InvalidMinLengthOnBoolType bool

// Invalid: MaxLength marker on array type
// +kubebuilder:validation:MaxLength=10 // want `marker "kubebuilder:validation:MaxLength": type array is not allowed \(expected one of: \[string\]\)`
type InvalidMaxLengthOnArrayType []string

type StringMarkersFieldTest struct {
	// Valid: Pattern marker on string field
	// +kubebuilder:validation:Pattern="^[a-z]+$"
	ValidPattern string `json:"validPattern"`

	// Valid: MinLength marker on string field
	// +kubebuilder:validation:MinLength=1
	ValidMinLength string `json:"validMinLength"`

	// Valid: MaxLength marker on string field
	// +kubebuilder:validation:MaxLength=100
	ValidMaxLength string `json:"validMaxLength"`

	// Valid: All string markers on string field
	// +kubebuilder:validation:Pattern="^[a-z]+$"
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=100
	ValidAllStringMarkers string `json:"validAllStringMarkers"`

	// Valid: Using PatternType
	ValidPatternTyped PatternType `json:"validPatternTyped"`

	// Valid: Using MinLengthType
	ValidMinLengthTyped MinLengthType `json:"validMinLengthTyped"`

	// Valid: Using MaxLengthType
	ValidMaxLengthTyped MaxLengthType `json:"validMaxLengthTyped"`

	// Invalid: Pattern marker on integer field
	// +kubebuilder:validation:Pattern="^[0-9]+$" // want `marker "kubebuilder:validation:Pattern": type integer is not allowed \(expected one of: \[string\]\)`
	InvalidPatternOnInt int32 `json:"invalidPatternOnInt"`

	// Invalid: MinLength marker on boolean field
	// +kubebuilder:validation:MinLength=5 // want `marker "kubebuilder:validation:MinLength": type boolean is not allowed \(expected one of: \[string\]\)`
	InvalidMinLengthOnBool bool `json:"invalidMinLengthOnBool"`

	// Invalid: MaxLength marker on array field
	// +kubebuilder:validation:MaxLength=10 // want `marker "kubebuilder:validation:MaxLength": type array is not allowed \(expected one of: \[string\]\)`
	InvalidMaxLengthOnArray []string `json:"invalidMaxLengthOnArray"`

	// Invalid: MinLength marker on named type
	// +kubebuilder:validation:MinLength=1 // want `marker "kubebuilder:validation:MinLength": marker should be declared on the type definition of StringType instead of the field`
	InvalidMinLengthOnMinLengthType StringType `json:"invalidMinLengthOnMinLengthType"`

	// Invalid: MaxLength marker on named type
	// +kubebuilder:validation:MaxLength=100 // want `marker "kubebuilder:validation:MaxLength": marker should be declared on the type definition of StringType instead of the field`
	InvalidMaxLengthOnMaxLengthType StringType `json:"invalidMaxLengthOnMaxLengthType"`

	// Invalid: Pattern marker on named type
	// +kubebuilder:validation:Pattern="^[0-9]+$" // want `marker "kubebuilder:validation:Pattern": marker should be declared on the type definition of StringType instead of the field`
	InvalidPatternOnIntTyped StringType `json:"invalidPatternOnIntTyped"`

	// Invalid: Using invalid named type
	InvalidMinLengthOnBoolTyped InvalidMinLengthOnBoolType `json:"invalidMinLengthOnBoolTyped"`

	// Invalid: Using invalid named type
	InvalidMaxLengthOnArrayTyped InvalidMaxLengthOnArrayType `json:"invalidMaxLengthOnArrayTyped"`
}
