package a

type TestStructs struct {
	// structWithMinProperties is a struct with no required fields, but marked with a min properties.
	// +required
	StructWithMinProperties StructWithMinProperties `json:"structWithMinProperties,omitempty,omitzero"`

	// structWithOmittedRequiredField is a struct with an omitted required field.
	// +required
	StructWithOmittedRequiredField StructWithOmittedRequiredField `json:"structWithOmittedRequiredField,omitempty,omitzero"`
}

// +kubebuilder:validation:MinProperties=1
type StructWithMinProperties struct {
	// map is a map of string to string.
	// +kubebuilder:validation:MinProperties=1
	// +optional
	Map map[string]string `json:"map,omitempty"`
}

// Struct with an omitted required field.
// The zero value of the struct is `{}` which is not valid because it does not satisfy the required marker on the string field.
type StructWithOmittedRequiredField struct {
	// string is a string.
	// +kubebuilder:validation:MinLength:=1
	// +required
	String string `json:"string,omitempty"`
}
