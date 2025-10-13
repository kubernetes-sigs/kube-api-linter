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

// ============================================================================
// Field-only markers (FieldScope)
// These should ERROR when placed on types
// ============================================================================

// +optional // want `marker "optional" can only be applied to fields`
// +required // want `marker "required" can only be applied to fields`
// +nullable // want `marker "nullable" can only be applied to fields`
type InvalidFieldOnlyOnType struct {
	Name string `json:"name"`
}

type FieldOnlyMarkersTest struct {
	// Valid field-only markers
	// +optional
	// +required
	// +k8s:optional
	// +k8s:required
	// +nullable
	// +kubebuilder:default="default"
	// +kubebuilder:validation:Example="example"
	// +kubebuilder:validation:EmbeddedResource
	// +kubebuilder:validation:Schemaless
	ValidFieldOnlyMarkers string `json:"validFieldOnlyMarkers"`
}

// ============================================================================
// Type-only markers (TypeScope)
// These should ERROR when placed on fields
// ============================================================================

type TypeOnlyMarkersTest struct {
	// +kubebuilder:validation:items:ExactlyOneOf={field1,field2} // want `marker "kubebuilder:validation:items:ExactlyOneOf" can only be applied to types`
	// +kubebuilder:validation:items:AtMostOneOf={field1,field2} // want `marker "kubebuilder:validation:items:AtMostOneOf" can only be applied to types`
	// +kubebuilder:validation:items:AtLeastOneOf={field1,field2} // want `marker "kubebuilder:validation:items:AtLeastOneOf" can only be applied to types`
	InvalidTypeOnlyOnField string `json:"invalidTypeOnlyOnField"`
}

// +kubebuilder:validation:items:ExactlyOneOf={Field1,Field2}
// +kubebuilder:validation:items:AtMostOneOf={Field1,Field2}
// +kubebuilder:validation:items:AtLeastOneOf={Field1,Field2}
type ValidTypeOnlyMarkers struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
}

// ============================================================================
// AnyScope markers - can be on both fields and types
// ============================================================================

// +kubebuilder:pruning:PreserveUnknownFields
// +kubebuilder:title="My Title"
type AnyScopeOnType struct {
	Name string `json:"name"`
}

type AnyScopeOnFieldTest struct {
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:title="Field Title"
	ValidAnyScopeField map[string]string `json:"validAnyScopeField"`
}

// ============================================================================
// Numeric markers (AnyScope, integer/number types)
// ============================================================================

// +kubebuilder:validation:Minimum=0
// +kubebuilder:validation:Maximum=100
// +kubebuilder:validation:ExclusiveMinimum=false
// +kubebuilder:validation:ExclusiveMaximum=false
// +kubebuilder:validation:MultipleOf=5
type NumericType int32

type NumericMarkersFieldTest struct {
	// Valid: numeric markers on numeric types
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:validation:ExclusiveMinimum=false
	// +kubebuilder:validation:ExclusiveMaximum=false
	// +kubebuilder:validation:MultipleOf=5
	ValidNumericField int32 `json:"validNumericField"`

	// +kubebuilder:validation:Minimum=0.0
	// +kubebuilder:validation:Maximum=1.0
	ValidFloatField float64 `json:"validFloatField"`

	// Invalid: numeric marker on string field
	// +kubebuilder:validation:Minimum=0 // want `marker "kubebuilder:validation:Minimum": type string is not allowed \(expected one of: \[integer number\]\)`
	InvalidMinimumOnString string `json:"invalidMinimumOnString"`

	// Invalid: numeric marker on bool field
	// +kubebuilder:validation:Maximum=100 // want `marker "kubebuilder:validation:Maximum": type boolean is not allowed \(expected one of: \[integer number\]\)`
	InvalidMaximumOnBool bool `json:"invalidMaximumOnBool"`
}

// ============================================================================
// String markers (AnyScope, string types)
// ============================================================================

// +kubebuilder:validation:Pattern="^[a-z]+$"
// +kubebuilder:validation:MinLength=1
// +kubebuilder:validation:MaxLength=100
type StringType string

type StringMarkersFieldTest struct {
	// Valid: string markers on string field
	// +kubebuilder:validation:Pattern="^[a-z]+$"
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=100
	ValidStringField string `json:"validStringField"`

	// Invalid: string marker on int field
	// +kubebuilder:validation:Pattern="^[0-9]+$" // want `marker "kubebuilder:validation:Pattern": type integer is not allowed \(expected one of: \[string\]\)`
	InvalidPatternOnInt int32 `json:"invalidPatternOnInt"`

	// Invalid: string marker on array field
	// +kubebuilder:validation:MinLength=5 // want `marker "kubebuilder:validation:MinLength": type array is not allowed \(expected one of: \[string\]\)`
	InvalidMinLengthOnArray []string `json:"invalidMinLengthOnArray"`
}

// ============================================================================
// Array markers (AnyScope, array types)
// ============================================================================

// +kubebuilder:validation:MinItems=1
// +kubebuilder:validation:MaxItems=10
// +kubebuilder:validation:UniqueItems=true
type StringArray []string

type ArrayMarkersFieldTest struct {
	// Valid: array markers on array field
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=10
	// +kubebuilder:validation:UniqueItems=true
	ValidArrayField []string `json:"validArrayField"`

	// Invalid: array marker on string field
	// +kubebuilder:validation:MinItems=1 // want `marker "kubebuilder:validation:MinItems": type string is not allowed \(expected one of: \[array\]\)`
	InvalidMinItemsOnString string `json:"invalidMinItemsOnString"`

	// Invalid: array marker on object field
	// +kubebuilder:validation:MaxItems=10 // want `marker "kubebuilder:validation:MaxItems": type object is not allowed \(expected one of: \[array\]\)`
	InvalidMaxItemsOnObject map[string]string `json:"invalidMaxItemsOnObject"`
}

// ============================================================================
// Object markers (AnyScope, object types)
// ============================================================================

// +kubebuilder:validation:MinProperties=1
// +kubebuilder:validation:MaxProperties=10
type ObjectType struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
}

type ObjectMarkersFieldTest struct {
	// Valid: object markers on map field
	// +kubebuilder:validation:MinProperties=1
	// +kubebuilder:validation:MaxProperties=10
	ValidObjectField map[string]string `json:"validObjectField"`

	// Invalid: object marker on string field
	// +kubebuilder:validation:MinProperties=2 // want `marker "kubebuilder:validation:MinProperties": type string is not allowed \(expected one of: \[object\]\)`
	InvalidMinPropertiesOnString string `json:"invalidMinPropertiesOnString"`

	// Invalid: object marker on array field
	// +kubebuilder:validation:MaxProperties=5 // want `marker "kubebuilder:validation:MaxProperties": type array is not allowed \(expected one of: \[object\]\)`
	InvalidMaxPropertiesOnArray []string `json:"invalidMaxPropertiesOnArray"`
}

// ============================================================================
// General markers (AnyScope, any type)
// ============================================================================

// +kubebuilder:validation:Enum=A;B;C
// +kubebuilder:validation:Format=email
// +kubebuilder:validation:Type=string
// +kubebuilder:validation:XValidation:rule="self.size() > 0"
type GeneralType string

type GeneralMarkersFieldTest struct {
	// +kubebuilder:validation:Enum=A;B;C
	// +kubebuilder:validation:Format=email
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:XValidation:rule="self.size() > 0"
	ValidGeneralField string `json:"validGeneralField"`
}

// ============================================================================
// Server-Side Apply topology markers (AnyScope)
// ============================================================================

// +listType=map
type ItemList []Item

// +mapType=granular
type ConfigMap map[string]string

// +structType=atomic
type AtomicStruct struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
}

type TopologyMarkersFieldTest struct {
	// +listType=map
	// +listMapKey=name
	ValidListMarkers []Item `json:"validListMarkers"`

	// +mapType=granular
	ValidMapType map[string]string `json:"validMapType"`

	// +structType=atomic
	ValidStruct EmbeddedStruct `json:"validStruct"`
}

type EmbeddedStruct struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
}

type Item struct {
	Name string `json:"name"`
}

// ============================================================================
// Array items markers (AnyScope, array with element constraints)
// ============================================================================

// +kubebuilder:validation:items:Maximum=100
// +kubebuilder:validation:items:Minimum=0
// +kubebuilder:validation:items:MultipleOf=5
type NumericArrayType []int32

// +kubebuilder:validation:items:Pattern="^[a-z]+$"
// +kubebuilder:validation:items:MinLength=1
// +kubebuilder:validation:items:MaxLength=50
type StringArrayType []string

// +kubebuilder:validation:items:MinItems=1
// +kubebuilder:validation:items:MaxItems=5
type NestedArrayType [][]string

// +kubebuilder:validation:items:MinProperties=1
// +kubebuilder:validation:items:MaxProperties=5
type ObjectArrayType []map[string]string

type ArrayItemsMarkersFieldTest struct {
	// Valid: Numeric element constraints
	// +kubebuilder:validation:items:Maximum=100
	// +kubebuilder:validation:items:Minimum=0
	// +kubebuilder:validation:items:MultipleOf=5
	// +kubebuilder:validation:items:ExclusiveMaximum=false
	// +kubebuilder:validation:items:ExclusiveMinimum=false
	ValidNumericArrayItems []int32 `json:"validNumericArrayItems"`

	// Valid: String element constraints
	// +kubebuilder:validation:items:Pattern="^[a-z]+$"
	// +kubebuilder:validation:items:MinLength=1
	// +kubebuilder:validation:items:MaxLength=50
	ValidStringArrayItems []string `json:"validStringArrayItems"`

	// Valid: Nested array constraints
	// +kubebuilder:validation:items:MinItems=1
	// +kubebuilder:validation:items:MaxItems=5
	// +kubebuilder:validation:items:UniqueItems=true
	ValidNestedArrayItems [][]string `json:"validNestedArrayItems"`

	// Valid: Object element constraints
	// +kubebuilder:validation:items:MinProperties=1
	// +kubebuilder:validation:items:MaxProperties=5
	ValidObjectArrayItems []map[string]string `json:"validObjectArrayItems"`

	// Valid: General items markers
	// +kubebuilder:validation:items:Enum=A;B;C
	// +kubebuilder:validation:items:Format=uuid
	// +kubebuilder:validation:items:Type=string
	// +kubebuilder:validation:items:XValidation:rule="self != ''"
	ValidGeneralArrayItems []string `json:"validGeneralArrayItems"`

	// Invalid: items:Maximum on string array (element type mismatch)
	// +kubebuilder:validation:items:Maximum=100 // want `marker "kubebuilder:validation:items:Maximum": array element: type string is not allowed \(expected one of: \[integer number\]\)`
	InvalidItemsMaximumOnStringArray []string `json:"invalidItemsMaximumOnStringArray"`

	// Invalid: items:Pattern on int array (element type mismatch)
	// +kubebuilder:validation:items:Pattern="^[0-9]+$" // want `marker "kubebuilder:validation:items:Pattern": array element: type integer is not allowed \(expected one of: \[string\]\)`
	InvalidItemsPatternOnIntArray []int32 `json:"invalidItemsPatternOnIntArray"`

	// Invalid: items:MinProperties on string array (element type mismatch)
	// +kubebuilder:validation:items:MinProperties=1 // want `marker "kubebuilder:validation:items:MinProperties": array element: type string is not allowed \(expected one of: \[object\]\)`
	InvalidItemsMinPropertiesOnStringArray []string `json:"invalidItemsMinPropertiesOnStringArray"`

	// Invalid: items marker on non-array field
	// +kubebuilder:validation:items:Maximum=100 // want `marker "kubebuilder:validation:items:Maximum": type string is not allowed \(expected one of: \[array\]\)`
	InvalidItemsMarkerOnNonArray string `json:"invalidItemsMarkerOnNonArray"`
}
