package a

// StringAliasDVMaxLengthB is a string type with the DV maxLength marker on the alias (second file).
// +k8s:maxLength=512
type StringAliasDVMaxLengthB string

// StringAliasNoDVMaxLengthB is a string type without any max-length marker.
type StringAliasNoDVMaxLengthB string
