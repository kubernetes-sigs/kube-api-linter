package b

type MinLength struct {
	// +k8s:minLength=256
	StringWithMinLength string

	// +k8s:minLength=256
	StringPointerWithMinLength *string

	// +k8s:minLength=128
	StringAliasWithMinLengthOnField StringAlias

	StringAliasWithMinLengthOnAlias StringAliasWithMinLength

	StringAliasFromAnotherFile StringAliasB // want "field MinLength.StringAliasFromAnotherFile type StringAliasB must have a minimum length, add k8s:minLength marker"

	// +k8s:minLength=128
	StringAliasFromAnotherFileWithMinLengthOnField StringAliasB

	StringAliasWithMinLengthFromAnotherFile StringAliasWithMinLengthB

	StringWithoutMinLength string // want "field MinLength.StringWithoutMinLength must have a minimum length, add k8s:minLength marker"

	StringPointerWithoutMinLength *string // want "field MinLength.StringPointerWithoutMinLength must have a minimum length, add k8s:minLength marker"

	StringAliasWithoutMinLength StringAlias // want "field MinLength.StringAliasWithoutMinLength type StringAlias must have a minimum length, add k8s:minLength marker"

	// +k8s:enum
	EnumString string

	// +k8s:enum
	EnumStringPointer *string

	EnumStringAlias EnumStringAlias

	// +k8s:format=duration
	DurationFormat string

	// +k8s:format=date-time
	DateTimeFormat string

	// +k8s:format=date
	DateFormat string

	// +k8s:minItems=256
	ArrayWithMinItems []int

	ArrayWithoutMinItems []int // want "field MinLength.ArrayWithoutMinItems must have a minimum items, add k8s:minItems marker"

	ByteSlice []byte // want "field MinLength.ByteSlice must have a minimum length, add k8s:minLength marker"

	// +k8s:minLength=512
	ByteSliceWithMinLength []byte

	ByteSliceAlias ByteSliceAlias // want "field MinLength.ByteSliceAlias type ByteSliceAlias must have a minimum length, add k8s:minLength marker"

	// +k8s:minLength=512
	ByteSliceAliasWithMinLength ByteSliceAlias

	ByteSliceAliasWithMinLengthOnAlias ByteSliceAliasWithMinLength

	// +k8s:minItems=128
	StringArrayWithMinItemsWithoutMinElementLength []string // want "field MinLength.StringArrayWithMinItemsWithoutMinElementLength array element must have a minimum length, add k8s:eachVal=\\+k8s:minLength marker"

	StringArrayWithoutMinItemsWithoutMinElementLength []string // want "field MinLength.StringArrayWithoutMinItemsWithoutMinElementLength must have a minimum items, add k8s:minItems marker" "field MinLength.StringArrayWithoutMinItemsWithoutMinElementLength array element must have a minimum length, add k8s:eachVal=\\+k8s:minLength marker"

	// +k8s:minItems=64
	// +k8s:eachVal=+k8s:minLength=64
	StringArrayWithMinItemsAndMinElementLength []string

	// +k8s:eachVal=+k8s:minLength=512
	StringArrayWithoutMinItemsWithMinElementLength []string // want  "field MinLength.StringArrayWithoutMinItemsWithMinElementLength must have a minimum items, add k8s:minItems marker"

	// +k8s:minItems=128
	StringAliasArrayWithMinItemsWithoutMinElementLength []StringAlias // want "field MinLength.StringAliasArrayWithMinItemsWithoutMinElementLength array element type StringAlias must have a minimum length, add k8s:minLength marker"

	StringAliasArrayWithoutMinItemsWithoutMinElementLength []StringAlias // want "field MinLength.StringAliasArrayWithoutMinItemsWithoutMinElementLength must have a minimum items, add k8s:minItems marker" "field MinLength.StringAliasArrayWithoutMinItemsWithoutMinElementLength array element type StringAlias must have a minimum length, add k8s:minLength marker"

	// +k8s:minItems=64
	// +k8s:eachVal=+k8s:minLength=64
	StringAliasArrayWithMinItemsAndMinElementLength []StringAlias

	// +k8s:eachVal=+k8s:minLength=512
	StringAliasArrayWithoutMinItemsWithMinElementLength []StringAlias // want  "field MinLength.StringAliasArrayWithoutMinItemsWithMinElementLength must have a minimum items, add k8s:minItems marker"

	// +k8s:minItems=64
	StringAliasArrayWithMinItemsAndMinElementLengthOnAlias []StringAliasWithMinLength

	StringAliasArrayWithoutMinItemsWithMinElementLengthOnAlias []StringAliasWithMinLength // want  "field MinLength.StringAliasArrayWithoutMinItemsWithMinElementLengthOnAlias must have a minimum items, add k8s:minItems marker"

	InlineStruct struct { // want "field MinLength.InlineStruct must have a minimum properties, add k8s:minProperties marker"
		// +k8s:minLength=256
		StringWithMinLength string

		StringWithoutMinLength string // want "field StringWithoutMinLength must have a minimum length, add k8s:minLength marker"
	} `json:"inlineStruct"`

	// +k8s:minProperties=1
	InlineStructWithMinProperties struct{} `json:"inlineStructWithMinProperties"`

	InlineStructWithARequiredField struct {
		// +k8s:minLength=256
		// +required
		StringWithMinLength string
	} `json:"inlineStructWithARequiredField`

	StructWithoutMinProperties StructWithoutMinProperties `json:"structWithoutMinProperties` // want "field MinLength.StructWithoutMinProperties type StructWithoutMinProperties must have a minimum properties, add k8s:minProperties marker"

	StructWithMinProperties StructWithMinProperties `json:"structWithMinProperties`

	StructWithARequiredField StructWithARequiredField `json:"structWithARequiredField"`

	StructWithExactlyOneOf StructWithARequiredField `json:"structWithExactlyOneOf"`

	StructWithMalformedMinProperties StructWithMalformedMinProperties `json:"structWithMalformedMinProperties` // want "could not get min properties for struct: invalid format for minimum properties marker: error getting marker value: error converting value to number: strconv.ParseFloat: parsing \\\"abc\\\": invalid syntax"

	// +k8s:minItems=1
	StructArrayWithMinProperties []StructWithMinProperties `json:"structArrayWithMinProperties"`

	// +k8s:minItems=1
	StructArrayWithARequiredField []StructWithARequiredField `json:"structArrayWithARequiredField"`

	// +k8s:minItems=1
	StructArrayWithMalformedMinProperties []StructWithMalformedMinProperties `json:"structArrayWithMalformedMinProperties` // want "could not get min properties for struct: invalid format for minimum properties marker: error getting marker value: error converting value to number: strconv.ParseFloat: parsing \\\"abc\\\": invalid syntax"

	// +k8s:minItems=1
	StructArrayWithoutMinProperties []StructWithoutMinProperties `json:"structArrayWithoutMinProperties` // want "field MinLength.StructArrayWithoutMinProperties array element type StructWithoutMinProperties must have a minimum properties, add k8s:minProperties marker"
}

// StringAlias is a string without a MinLength.
type StringAlias string

// StringAliasWithMinLength is a string with a MinLength.
// +k8s:minLength=512
type StringAliasWithMinLength string

// EnumStringAlias is a string alias that is an enum.
// +k8s:enum
type EnumStringAlias string

// ByteSliceAlias is a byte slice without a MinLength.
type ByteSliceAlias []byte

// ByteSliceAliasWithMinLength is a byte slice with a MinLength.
// +k8s:minLength=512
type ByteSliceAliasWithMinLength []byte

type StructWithoutMinProperties struct {
	// +k8s:minLength=256
	StringWithMinLength string
}

// +k8s:minProperties=1
type StructWithMinProperties struct {
	// +k8s:minLength=256
	StringWithMinLength string
}

type StructWithARequiredField struct {
	// +k8s:minLength=256
	// +required
	StringWithMinLength string
}

// +k8s:minProperties=abc
type StructWithMalformedMinProperties struct {
	// +k8s:minLength=256
	StringWithMinLength string
}
