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

// +kubebuilder:validation:MinItems=1
// +kubebuilder:validation:MaxItems=10
// +kubebuilder:validation:UniqueItems=true
type StringArray []string

// +kubebuilder:validation:MinItems=0
type IntegerArray []int32

// +kubebuilder:validation:MaxItems=100
type BooleanArray []bool

// Type definitions with invalid markers
// +kubebuilder:validation:MinItems=1 // want `marker "kubebuilder:validation:MinItems": type string is not allowed \(expected one of: \[array\]\)`
type InvalidArrayMarkerOnStringType string

// +kubebuilder:validation:MaxItems=10 // want `marker "kubebuilder:validation:MaxItems": type object is not allowed \(expected one of: \[array\]\)`
type InvalidArrayMarkerOnMapType map[string]string

type ArrayMarkersFieldTest struct {
	// Valid: MinItems marker on array field
	// +kubebuilder:validation:MinItems=1
	ValidMinItems []string `json:"validMinItems"`

	// Valid: MaxItems marker on array field
	// +kubebuilder:validation:MaxItems=10
	ValidMaxItems []string `json:"validMaxItems"`

	// Valid: UniqueItems marker on array field
	// +kubebuilder:validation:UniqueItems=true
	ValidUniqueItems []string `json:"validUniqueItems"`

	// Valid: All array markers combined on array field
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=10
	// +kubebuilder:validation:UniqueItems=true
	ValidAllArrayMarkers []string `json:"validAllArrayMarkers"`

	// Invalid: MinItems marker on string field
	// +kubebuilder:validation:MinItems=1 // want `marker "kubebuilder:validation:MinItems": type string is not allowed \(expected one of: \[array\]\)`
	InvalidMinItemsOnString string `json:"invalidMinItemsOnString"`

	// Invalid: MaxItems marker on map field
	// +kubebuilder:validation:MaxItems=10 // want `marker "kubebuilder:validation:MaxItems": type object is not allowed \(expected one of: \[array\]\)`
	InvalidMaxItemsOnMap map[string]string `json:"invalidMaxItemsOnMap"`

	// Invalid: UniqueItems marker on integer field
	// +kubebuilder:validation:UniqueItems=true // want `marker "kubebuilder:validation:UniqueItems": type integer is not allowed \(expected one of: \[array\]\)`
	InvalidUniqueItemsOnInteger int32 `json:"invalidUniqueItemsOnInteger"`

	// Invalid: MinItems marker on boolean field
	// +kubebuilder:validation:MinItems=1 // want `marker "kubebuilder:validation:MinItems": type boolean is not allowed \(expected one of: \[array\]\)`
	InvalidMinItemsOnBoolean bool `json:"invalidMinItemsOnBoolean"`

	// Invalid: MaxItems marker on struct field
	// +kubebuilder:validation:MaxItems=10 // want `marker "kubebuilder:validation:MaxItems": type object is not allowed \(expected one of: \[array\]\)`
	InvalidMaxItemsOnStruct ArrayItem `json:"invalidMaxItemsOnStruct"`

	// Valid: Using named array type with markers
	ValidNamedArrayType StringArray `json:"validNamedArrayType"`

	// Valid: Using named array type (IntegerArray)
	ValidIntegerArrayType IntegerArray `json:"validIntegerArrayType"`

	// Valid: Using named array type (BooleanArray)
	ValidBooleanArrayType BooleanArray `json:"validBooleanArrayType"`

	// Invalid: Field marker on named array type (should use type definition)
	// +kubebuilder:validation:MinItems=5 // want `marker "kubebuilder:validation:MinItems": marker should be declared on the type definition of StringArray instead of the field`
	InvalidFieldMarkerOnNamedArray StringArray `json:"invalidFieldMarkerOnNamedArray"`

	// Invalid: Using invalid named type with array marker
	InvalidNamedStringType InvalidArrayMarkerOnStringType `json:"invalidNamedStringType"`

	// Invalid: Using invalid named type with array marker
	InvalidNamedMapType InvalidArrayMarkerOnMapType `json:"invalidNamedMapType"`
}

type ArrayItem struct {
	Name string `json:"name"`
}
