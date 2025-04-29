package a

// StringAliasB is a string without a MaxLength.
type StringAliasB string

// StringAliasWithMaxLengthB is a string with a MaxLength.
// +kubebuilder:validation:MaxLength:=512
type StringAliasWithMaxLengthB string
