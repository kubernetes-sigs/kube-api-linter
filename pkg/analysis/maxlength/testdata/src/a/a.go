package a

type MaxLength struct {
	// +kubebuilder:validation:MaxLength:=256
	StringWithMaxLength string

	// +kubebuilder:validation:MaxLength:=128
	StringAliasWithMaxLengthOnField StringAlias

	StringAliasWithMaxLengthOnAlias StringAliasWithMaxLength

	StringAliasFromAnotherFile StringAliasB // want "field StringAliasFromAnotherFile type StringAliasB must have a maximum length, add kubebuilder:validation:MaxLength marker"

	// +kubebuilder:validation:MaxLength:=128
	StringAliasFromAnotherFileWithMaxLengthOnField StringAliasB

	StringAliasWithMaxLengthFromAnotherFile StringAliasWithMaxLengthB

	StringWithoutMaxLength string // want "field StringWithoutMaxLength must have a maximum length, add kubebuilder:validation:MaxLength marker"

	StringAliasWithoutMaxLength StringAlias // want "field StringAliasWithoutMaxLength type StringAlias must have a maximum length, add kubebuilder:validation:MaxLength marker"

	// +kubebuilder:validation:Enum:="A";"B";"C"
	EnumString string

	EnumStringAlias EnumStringAlias

	// +kubebuilder:validation:Format:=duration
	DurationFormat string

	// +kubebuilder:validation:Format:=date-time
	DateTimeFormat string

	// +kubebuilder:validation:Format:=date
	DateFormat string

	// +kubebuilder:validation:MaxItems:=256
	ArrayWithMaxItems []int

	ArrayWithoutMaxItems []int // want "field ArrayWithoutMaxItems must have a maximum items, add kubebuilder:validation:MaxItems"

	ByteSlice []byte // want "field ByteSlice must have a maximum length, add kubebuilder:validation:MaxLength marker"

	// +kubebuilder:validation:MaxLength:=512
	ByteSliceWithMaxLength []byte

	ByteSliceAlias ByteSliceAlias // want "field ByteSliceAlias type ByteSliceAlias must have a maximum length, add kubebuilder:validation:MaxLength marker"

	// +kubebuilder:validation:MaxLength:=512
	ByteSliceAliasWithMaxLength ByteSliceAlias

	ByteSliceAliasWithMaxLengthOnAlias ByteSliceAliasWithMaxLength

	// +kubebuilder:validation:MaxItems:=128
	StringArrayWithMaxItemsWithoutMaxElementLength []string // want "field StringArrayWithMaxItemsWithoutMaxElementLength array element must have a maximum length, add kubebuilder:validation:items:MaxLength"

	StringArrayWithoutMaxItemsWithoutMaxElementLength []string // want "field StringArrayWithoutMaxItemsWithoutMaxElementLength must have a maximum items, add kubebuilder:validation:MaxItems" "field StringArrayWithoutMaxItemsWithoutMaxElementLength array element must have a maximum length, add kubebuilder:validation:items:MaxLength"

	// +kubebuilder:validation:MaxItems:=64
	// +kubebuilder:validation:items:MaxLength:=64
	StringArrayWithMaxItemsAndMaxElementLength []string

	// +kubebuilder:validation:items:MaxLength:=512
	StringArrayWithoutMaxItemsWithMaxElementLength []string // want  "field StringArrayWithoutMaxItemsWithMaxElementLength must have a maximum items, add kubebuilder:validation:MaxItems marker"

	// +kubebuilder:validation:MaxItems:=128
	StringAliasArrayWithMaxItemsWithoutMaxElementLength []StringAlias // want "field StringAliasArrayWithMaxItemsWithoutMaxElementLength array element type StringAlias must have a maximum length, add kubebuilder:validation:MaxLength marker"

	StringAliasArrayWithoutMaxItemsWithoutMaxElementLength []StringAlias // want "field StringAliasArrayWithoutMaxItemsWithoutMaxElementLength must have a maximum items, add kubebuilder:validation:MaxItems" "field StringAliasArrayWithoutMaxItemsWithoutMaxElementLength array element type StringAlias must have a maximum length, add kubebuilder:validation:MaxLength"

	// +kubebuilder:validation:MaxItems:=64
	// +kubebuilder:validation:items:MaxLength:=64
	StringAliasArrayWithMaxItemsAndMaxElementLength []StringAlias

	// +kubebuilder:validation:items:MaxLength:=512
	StringAliasArrayWithoutMaxItemsWithMaxElementLength []StringAlias // want  "field StringAliasArrayWithoutMaxItemsWithMaxElementLength must have a maximum items, add kubebuilder:validation:MaxItems"

	// +kubebuilder:validation:MaxItems:=64
	StringAliasArrayWithMaxItemsAndMaxElementLengthOnAlias []StringAliasWithMaxLength

	StringAliasArrayWithoutMaxItemsWithMaxElementLengthOnAlias []StringAliasWithMaxLength // want  "field StringAliasArrayWithoutMaxItemsWithMaxElementLengthOnAlias must have a maximum items, add kubebuilder:validation:MaxItems"

	Struct struct {
		// +kubebuilder:validation:MaxLength:=256
		StringWithMaxLength string

		StringWithoutMaxLength string // want "field StringWithoutMaxLength must have a maximum length, add kubebuilder:validation:MaxLength marker"
	} `json:"struct"`

	// +optional
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	AllOf []JSONSchemaProps `json:"allOf,omitempty"`
}

// StringAlias is a string without a MaxLength.
type StringAlias string

// StringAliasWithMaxLength is a string with a MaxLength.
// +kubebuilder:validation:MaxLength:=512
type StringAliasWithMaxLength string

// EnumStringAlias is a string alias that is an enum.
// +kubebuilder:validation:Enum:="A";"B";"C"
type EnumStringAlias string

// ByteSliceAlias is a byte slice without a MaxLength.
type ByteSliceAlias []byte

// ByteSliceAliasWithMaxLength is a byte slice with a MaxLength.
// +kubebuilder:validation:MaxLength:=512
type ByteSliceAliasWithMaxLength []byte

// JSONSchemaProps is a placeholder for the JSON schema properties.
type JSONSchemaProps struct{}
