package a

// StringAliasB is a string without a MinLength.
type StringAliasB string

// StringAliasWithMinLengthB is a string with a MinLength.
// +kubebuilder:validation:MinLength:=512
type StringAliasWithMinLengthB string
