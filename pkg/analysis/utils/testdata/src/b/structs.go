package b

type ZeroValueTestStructs struct {
	StructWithAllOptionalFields StructWithAllOptionalFields `json:"structWithAllOptionalFields,omitempty"` // want "zero value is valid" "validation is not complete"

	StructWithMinProperties StructWithMinProperties `json:"structWithMinProperties,omitempty"` // want "zero value is not valid" "validation is complete"

	StructWithNonOmittedFieldsAndMinProperties StructWithNonOmittedFieldsAndMinProperties `json:"structWithOneNonOmittedFieldAndMinProperties,omitempty"` // want "zero value is valid" "validation is complete"

	StructWithOneNonOmittedFieldAndMinProperties StructWithOneNonOmittedFieldAndMinProperties `json:"structWithOneNonOmittedFieldAndMinPropertiesAndOmitEmpty,omitempty"` // want "zero value is not valid" "validation is complete"

	StructWithOmittedRequiredField StructWithOmittedRequiredField `json:"structWithOmittedRequiredField,omitempty"` // want "zero value is not valid" "validation is complete"
}

type StructWithAllOptionalFields struct {
	// +optional
	String string `json:"string,omitempty"` // want "zero value is valid" "validation is not complete"

	// +optional
	StringPtr *string `json:"stringPtr,omitempty"` // want "zero value is valid" "validation is not complete"

	// +optional
	Int int `json:"int,omitempty"` // want "zero value is valid" "validation is not complete"

	// +optional
	IntPtr *int `json:"intPtr,omitempty"` // want "zero value is valid" "validation is not complete"
}

// +kubebuilder:validation:MinProperties=1
type StructWithMinProperties struct {
	// +kubebuilder:validation:MinProperties=1
	// +optional
	Map map[string]string `json:"map,omitempty"` // want "zero value is not valid" "validation is complete"
}

type StructWithNonOmittedFields struct {
	// +required
	String string `json:"string"` // want "zero value is valid" "validation is not complete"

	// +required
	Int int32 `json:"int"` // want "zero value is valid" "validation is not complete"
}

// Struct with non-omitted fields and minProperties marker.
// Because there is no omitempty, and the zero values are valid, the zero value here is `{"string:"", "int":0}`.
// This means the MinProperties marker is satisfied even when the object is the zero value.
// +kubebuilder:validation:MinProperties=2
type StructWithNonOmittedFieldsAndMinProperties struct {
	// +required
	String string `json:"string"` // want "zero value is valid" "validation is not complete"

	// +required
	Int int32 `json:"int"` // want "zero value is valid" "validation is not complete"
}

// Struct with one non-omitted field, and one omitted field and minProperties marker.
// The zero value of the struct is `{"string:""}` which is not valid because it does not satisfy the MinProperties marker.
// +kubebuilder:validation:MinProperties=2
type StructWithOneNonOmittedFieldAndMinProperties struct {
	// +required
	String string `json:"string"` // want "zero value is valid" "validation is not complete"

	// +optional
	Int int32 `json:"int,omitempty"` // want "zero value is valid" "validation is not complete"
}

// Struct with an omitted required field.
// The zero value of the struct is `{}` which is not valid because it does not satisfy the required marker on the string field.
type StructWithOmittedRequiredField struct {
	// +required
	String string `json:"string,omitempty"` // want "zero value is valid" "validation is not complete"
}
