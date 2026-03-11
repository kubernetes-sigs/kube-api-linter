package a

import "externaltypes"

type A struct {
	// required field should not be picked up.
	// +required
	RequiredField string `json:"requiredField,omitempty"`

	// pointerString is a pointer string field.
	// +optional
	PointerString *string `json:"pointerString,omitempty"`

	// pointerInt is a pointer int field.
	// +optional
	PointerInt *int `json:"pointerInt,omitempty"`

	// pointerStruct is a pointer struct field.
	// +optional
	PointerStruct *B `json:"pointerStruct,omitempty"`

	// string is a string field.
	// +optional
	String string `json:"string,omitempty"` // want "field A.String should be a pointer."

	// NonOmittedString is a string field without omitempty
	// +optional
	NonOmittedString string `json:"nonOmittedString"` // want "field A.NonOmittedString should be a pointer." "field A.NonOmittedString should have the omitempty tag."

	// int is an int field.
	// +optional
	Int int `json:"int,omitempty"` // want "field A.Int should be a pointer."

	// nonOmittedInt is an int field without omitempty
	// +optional
	NonOmittedInt int `json:"nonOmittedInt"` // want "field A.NonOmittedInt should be a pointer." "field A.NonOmittedInt should have the omitempty tag."

	// struct is a struct field.
	// +optional
	Struct B `json:"struct,omitempty"` // want "field A.Struct should be a pointer."

	// nonOmittedStruct is a struct field without omitempty.
	// +optional
	NonOmittedStruct B `json:"nonOmittedStruct"` // want "field A.NonOmittedStruct should be a pointer." "field A.NonOmittedStruct should have the omitempty tag."

	// structWithMinProperties is a struct field with a minimum number of properties.
	// +kubebuilder:validation:MinProperties=1
	// +optional
	StructWithMinProperties B `json:"structWithMinProperties,omitempty"` // want "field A.StructWithMinProperties should be a pointer."

	// structWithMinPropertiesOnStruct is a struct field with a minimum number of properties on the struct.
	// +optional
	StructWithMinPropertiesOnStruct D `json:"structWithMinPropertiesOnStruct,omitempty"` // want "field A.StructWithMinPropertiesOnStruct should be a pointer."

	// slice is a slice field.
	// +optional
	Slice []string `json:"slice,omitempty"`

	// map is a map field.
	// +optional
	Map map[string]string `json:"map,omitempty"`

	// PointerSlice is a pointer slice field.
	// +optional
	PointerSlice *[]string `json:"pointerSlice,omitempty"` // want "field A.PointerSlice underlying type does not need to be a pointer. The pointer should be removed."

	// PointerMap is a pointer map field.
	// +optional
	PointerMap *map[string]string `json:"pointerMap,omitempty"` // want "field A.PointerMap underlying type does not need to be a pointer. The pointer should be removed."

	// PointerPointerString is a double pointer string field.
	// +optional
	DoublePointerString **string `json:"doublePointerString,omitempty"` // want "field A.DoublePointerString underlying type does not need to be a pointer. The pointer should be removed."

	// PointerStringAlias is a pointer string alias field.
	// +optional
	PointerStringAlias *StringAlias `json:"pointerStringAlias,omitempty"`

	// PointerStringAliasFromAnotherFile is a pointer string alias field.
	// It proves that we can use types defined in other files.
	// +optional
	StringAliasFromAnotherFile *StringAliasFromAnotherFile `json:"pointerStringAliasFromAnotherFile,omitempty"`

	// PointerIntAlias is a pointer int alias field.
	// +optional
	PointerIntAlias *IntAlias `json:"pointerIntAlias,omitempty"`

	// PointerFloatAlias is a pointer float alias field.
	// +optional
	PointerFloatAlias *FloatAlias `json:"pointerFloatAlias,omitempty"`

	// PointerBoolAlias is a pointer bool alias field.
	// +optional
	PointerBoolAlias *BoolAlias `json:"pointerBoolAlias,omitempty"`

	// PointerSliceAlias is a pointer slice alias field.
	// +optional
	PointerSliceAlias *SliceAlias `json:"pointerSliceAlias,omitempty"` // want "field A.PointerSliceAlias underlying type does not need to be a pointer. The pointer should be removed."

	// PointerMapAlias is a pointer map alias field.
	// +optional
	PointerMapAlias *MapAlias `json:"pointerMapAlias,omitempty"` // want "field A.PointerMapAlias underlying type does not need to be a pointer. The pointer should be removed."

	// StringAliasWithEnum is a string alias field with enum validation.
	// With the "Always" pointer preference, optional fields should be pointers regardless of zero value validity.
	// +optional
	StringAliasWithEnum StringAliasWithEnum `json:"stringAliasWithEnum,omitempty"` // want "field A.StringAliasWithEnum should be a pointer."

	// StringAliasWithEnumPointer is a pointer string alias field with enum validation.
	// This is correctly a pointer since the zero value is not valid.
	// +optional
	StringAliasWithEnumPointer *StringAliasWithEnum `json:"stringAliasWithEnumPointer,omitempty"`

	// StringAliasWithEnumNoOmitEmpty is a string alias field with enum validation and no omitempty.
	// +optional
	StringAliasWithEnumNoOmitEmpty *StringAliasWithEnum `json:"stringAliasWithEnumNoOmitEmpty"` // want "field A.StringAliasWithEnumNoOmitEmpty should have the omitempty tag."

	// ExternalResourceList is a named map type from an external package.
	// This should NOT require a pointer since the underlying type is a map.
	// +optional
	ExternalResourceList externaltypes.ResourceList `json:"externalResourceList,omitempty"`

	// ExternalStringSlice is a named slice type from an external package.
	// This should NOT require a pointer since the underlying type is a slice.
	// +optional
	ExternalStringSlice externaltypes.StringSlice `json:"externalStringSlice,omitempty"`

	// ExternalStringMap is a named map type from an external package.
	// This should NOT require a pointer since the underlying type is a map.
	// +optional
	ExternalStringMap externaltypes.StringMap `json:"externalStringMap,omitempty"`

	// PointerExternalResourceList is a pointer to a named map type from an external package.
	// This should report that the pointer is unnecessary since the underlying type is a map.
	// +optional
	PointerExternalResourceList *externaltypes.ResourceList `json:"pointerExternalResourceList,omitempty"` // want "field A.PointerExternalResourceList underlying type does not need to be a pointer. The pointer should be removed."

	// PointerExternalStringSlice is a pointer to a named slice type from an external package.
	// This should report that the pointer is unnecessary since the underlying type is a slice.
	// +optional
	PointerExternalStringSlice *externaltypes.StringSlice `json:"pointerExternalStringSlice,omitempty"` // want "field A.PointerExternalStringSlice underlying type does not need to be a pointer. The pointer should be removed."

	// ExternalStringAlias is a string alias from an external package.
	// This should require a pointer since the underlying type is a string.
	// +optional
	ExternalStringAlias externaltypes.StringAlias `json:"externalStringAlias,omitempty"` // want "field A.ExternalStringAlias should be a pointer."

	// ExternalIntAlias is an int alias from an external package.
	// This should require a pointer since the underlying type is an int.
	// +optional
	ExternalIntAlias externaltypes.IntAlias `json:"externalIntAlias,omitempty"` // want "field A.ExternalIntAlias should be a pointer."

	// ExternalBoolAlias is a bool alias from an external package.
	// This should require a pointer since the underlying type is a bool.
	// +optional
	ExternalBoolAlias externaltypes.BoolAlias `json:"externalBoolAlias,omitempty"` // want "field A.ExternalBoolAlias should be a pointer."

	// ExternalFloatAlias is a float alias from an external package.
	// This should require a pointer since the underlying type is a float.
	// +optional
	ExternalFloatAlias externaltypes.FloatAlias `json:"externalFloatAlias,omitempty"` // want "field A.ExternalFloatAlias should be a pointer."

	// ExternalStructType is a struct type from an external package.
	// This should require a pointer since the underlying type is a struct.
	// +optional
	ExternalStructType externaltypes.StructType `json:"externalStructType,omitempty"` // want "field A.ExternalStructType should be a pointer."

	// ExternalStructWithNonOmittedField is a struct type with non-omitted fields from an external package.
	// With "Always" pointer preference, this should still require a pointer.
	// +optional
	ExternalStructWithNonOmittedField externaltypes.StructTypeWithNonOmittedField `json:"externalStructWithNonOmittedField,omitempty"` // want "field A.ExternalStructWithNonOmittedField should be a pointer."

	// PointerExternalStringAlias is a pointer to a string alias from an external package.
	// This is valid - string aliases should be pointers.
	// +optional
	PointerExternalStringAlias *externaltypes.StringAlias `json:"pointerExternalStringAlias,omitempty"`

	// PointerExternalIntAlias is a pointer to an int alias from an external package.
	// This is valid - int aliases should be pointers.
	// +optional
	PointerExternalIntAlias *externaltypes.IntAlias `json:"pointerExternalIntAlias,omitempty"`

	// PointerExternalBoolAlias is a pointer to a bool alias from an external package.
	// This is valid - bool aliases should be pointers.
	// +optional
	PointerExternalBoolAlias *externaltypes.BoolAlias `json:"pointerExternalBoolAlias,omitempty"`

	// PointerExternalStructType is a pointer to a struct type from an external package.
	// This is valid - struct types should be pointers.
	// +optional
	PointerExternalStructType *externaltypes.StructType `json:"pointerExternalStructType,omitempty"`
}

type B struct {
	// pointerString is a pointer string field.
	// +optional
	PointerString *string `json:"pointerString,omitempty"`
}

// +kubebuilder:validation:MinProperties=1
type D struct {
	// string is a string field.
	// +optional
	String *string `json:"string,omitempty"`

	// stringWithMinLength1 with minimum length is a string field.
	// +kubebuilder:validation:MinLength=1
	// +optional
	StringWithMinLength1 *string `json:"stringWithMinLength1,omitempty"`
}

type StringAlias string

type IntAlias int

type FloatAlias float64

type BoolAlias bool

type SliceAlias []string

type MapAlias map[string]string

// StringAliasWithEnum is a string alias with enum validation.
// The zero value ("") is not in the enum, making it invalid.
// +kubebuilder:validation:Enum=value1;value2
type StringAliasWithEnum string
