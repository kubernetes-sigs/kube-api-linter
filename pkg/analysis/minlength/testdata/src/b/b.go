package b

// StringAliasB is a string without a MinLength.
type StringAliasB string

// StringAliasWithMinLengthB is a string with a MinLength.
// +k8s:minLength=512
type StringAliasWithMinLengthB string
