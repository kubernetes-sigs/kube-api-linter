package a

type MinLength struct {
	// +kubebuilder:validation:MinLength:=256
	StringWithMinLength string

	// +kubebuilder:validation:MinLength:=256
	StringPointerWithMinLength *string

	// +kubebuilder:validation:MinLength:=128
	StringAliasWithMinLengthOnField StringAlias

	StringAliasWithMinLengthOnAlias StringAliasWithMinLength

	StringAliasFromAnotherFile StringAliasB // want "field StringAliasFromAnotherFile type StringAliasB must have a minimum length, add kubebuilder:validation:MinLength marker"

	// +kubebuilder:validation:MinLength:=128
	StringAliasFromAnotherFileWithMinLengthOnField StringAliasB

	StringAliasWithMinLengthFromAnotherFile StringAliasWithMinLengthB

	StringWithoutMinLength string // want "field StringWithoutMinLength must have a minimum length, add kubebuilder:validation:MinLength marker"

	StringPointerWithoutMinLength *string // want "field StringPointerWithoutMinLength must have a minimum length, add kubebuilder:validation:MinLength marker"

	StringAliasWithoutMinLength StringAlias // want "field StringAliasWithoutMinLength type StringAlias must have a minimum length, add kubebuilder:validation:MinLength marker"

	// +kubebuilder:validation:Enum:="A";"B";"C"
	EnumString string

	// +kubebuilder:validation:Enum:="A";"B";"C"
	EnumStringPointer *string

	EnumStringAlias EnumStringAlias

	// +kubebuilder:validation:Format:=duration
	DurationFormat string

	// +kubebuilder:validation:Format:=date-time
	DateTimeFormat string

	// +kubebuilder:validation:Format:=date
	DateFormat string

	// +kubebuilder:validation:MinItems:=256
	ArrayWithMinItems []int

	ArrayWithoutMinItems []int // want "field ArrayWithoutMinItems must have a minimum items, add kubebuilder:validation:MinItems"

	ByteSlice []byte // want "field ByteSlice must have a minimum length, add kubebuilder:validation:MinLength marker"

	// +kubebuilder:validation:MinLength:=512
	ByteSliceWithMinLength []byte

	ByteSliceAlias ByteSliceAlias // want "field ByteSliceAlias type ByteSliceAlias must have a minimum length, add kubebuilder:validation:MinLength marker"

	// +kubebuilder:validation:MinLength:=512
	ByteSliceAliasWithMinLength ByteSliceAlias

	ByteSliceAliasWithMinLengthOnAlias ByteSliceAliasWithMinLength

	// +kubebuilder:validation:MinItems:=128
	StringArrayWithMinItemsWithoutMinElementLength []string // want "field StringArrayWithMinItemsWithoutMinElementLength array element must have a minimum length, add kubebuilder:validation:items:MinLength"

	StringArrayWithoutMinItemsWithoutMinElementLength []string // want "field StringArrayWithoutMinItemsWithoutMinElementLength must have a minimum items, add kubebuilder:validation:MinItems" "field StringArrayWithoutMinItemsWithoutMinElementLength array element must have a minimum length, add kubebuilder:validation:items:MinLength"

	// +kubebuilder:validation:MinItems:=64
	// +kubebuilder:validation:items:MinLength:=64
	StringArrayWithMinItemsAndMinElementLength []string

	// +kubebuilder:validation:items:MinLength:=512
	StringArrayWithoutMinItemsWithMinElementLength []string // want  "field StringArrayWithoutMinItemsWithMinElementLength must have a minimum items, add kubebuilder:validation:MinItems marker"

	// +kubebuilder:validation:MinItems:=128
	StringAliasArrayWithMinItemsWithoutMinElementLength []StringAlias // want "field StringAliasArrayWithMinItemsWithoutMinElementLength array element type StringAlias must have a minimum length, add kubebuilder:validation:MinLength marker"

	StringAliasArrayWithoutMinItemsWithoutMinElementLength []StringAlias // want "field StringAliasArrayWithoutMinItemsWithoutMinElementLength must have a minimum items, add kubebuilder:validation:MinItems" "field StringAliasArrayWithoutMinItemsWithoutMinElementLength array element type StringAlias must have a minimum length, add kubebuilder:validation:MinLength"

	// +kubebuilder:validation:MinItems:=64
	// +kubebuilder:validation:items:MinLength:=64
	StringAliasArrayWithMinItemsAndMinElementLength []StringAlias

	// +kubebuilder:validation:items:MinLength:=512
	StringAliasArrayWithoutMinItemsWithMinElementLength []StringAlias // want  "field StringAliasArrayWithoutMinItemsWithMinElementLength must have a minimum items, add kubebuilder:validation:MinItems"

	// +kubebuilder:validation:MinItems:=64
	StringAliasArrayWithMinItemsAndMinElementLengthOnAlias []StringAliasWithMinLength

	StringAliasArrayWithoutMinItemsWithMinElementLengthOnAlias []StringAliasWithMinLength // want  "field StringAliasArrayWithoutMinItemsWithMinElementLengthOnAlias must have a minimum items, add kubebuilder:validation:MinItems"

	InlineStruct struct { // want "field InlineStruct must have a minimum properties, add kubebuilder:validation:MinProperties marker"
		// +kubebuilder:validation:MinLength:=256
		StringWithMinLength string

		StringWithoutMinLength string // want "field StringWithoutMinLength must have a minimum length, add kubebuilder:validation:MinLength marker"
	} `json:"inlineStruct"`

	// +kubebuilder:validation:MinProperties:=1
	InlineStructWithMinProperties struct{} `json:"inlineStructWithMinProperties"`

	InlineStructWithARequiredField struct {
		// +kubebuilder:validation:MinLength:=256
		// +required
		StringWithMinLength string
	} `json:"inlineStructWithARequiredField`

	StructWithoutMinProperties StructWithoutMinProperties `json:"structWithoutMinProperties` // want "field StructWithoutMinProperties type StructWithoutMinProperties must have a minimum properties, add kubebuilder:validation:MinProperties marker"

	StructWithMinProperties StructWithMinProperties `json:"structWithMinProperties`

	StructWithARequiredField StructWithARequiredField `json:"structWithARequiredField"`

	StructWithMalformedMinProperties StructWithMalformedMinProperties `json:"structWithMalformedMinProperties` // want "could not get min properties for struct: invalid format for minimum properties marker: error getting marker value: error converting value to number: strconv.ParseFloat: parsing \\\"abc\\\": invalid syntax"

	// +kubebuilder:validation:MinItems:=1
	StructArrayWithMinProperties []StructWithMinProperties `json:"structArrayWithMinProperties"`

	// +kubebuilder:validation:MinItems:=1
	StructArrayWithARequiredField []StructWithARequiredField `json:"structArrayWithARequiredField"`

	// +kubebuilder:validation:MinItems:=1
	StructArrayWithMalformedMinProperties []StructWithMalformedMinProperties `json:"structArrayWithMalformedMinProperties` // want "could not get min properties for struct: invalid format for minimum properties marker: error getting marker value: error converting value to number: strconv.ParseFloat: parsing \\\"abc\\\": invalid syntax"

	// +kubebuilder:validation:MinItems:=1
	StructArrayWithoutMinProperties []StructWithoutMinProperties `json:"structArrayWithoutMinProperties` // want "field StructArrayWithoutMinProperties array element type StructWithoutMinProperties must have a minimum properties, add kubebuilder:validation:MinProperties marker"
}

// StringAlias is a string without a MinLength.
type StringAlias string

// StringAliasWithMinLength is a string with a MinLength.
// +kubebuilder:validation:MinLength:=512
type StringAliasWithMinLength string

// EnumStringAlias is a string alias that is an enum.
// +kubebuilder:validation:Enum:="A";"B";"C"
type EnumStringAlias string

// ByteSliceAlias is a byte slice without a MinLength.
type ByteSliceAlias []byte

// ByteSliceAliasWithMinLength is a byte slice with a MinLength.
// +kubebuilder:validation:MinLength:=512
type ByteSliceAliasWithMinLength []byte

type StructWithoutMinProperties struct {
	// +kubebuilder:validation:MinLength:=256
	StringWithMinLength string
}

// +kubebuilder:validation:MinProperties:=1
type StructWithMinProperties struct {
	// +kubebuilder:validation:MinLength:=256
	StringWithMinLength string
}

type StructWithARequiredField struct {
	// +kubebuilder:validation:MinLength:=256
	// +required
	StringWithMinLength string
}

// +kubebuilder:validation:MinProperties:=abc
type StructWithMalformedMinProperties struct {
	// +kubebuilder:validation:MinLength:=256
	StringWithMinLength string
}
