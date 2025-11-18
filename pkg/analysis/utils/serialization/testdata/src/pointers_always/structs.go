package a

type TestStructs struct {
	// StructWithAllOptionalFields has a zero value of {}, which is valid because all fields are optional.

	StructWithAllOptionalFields StructWithAllOptionalFields `json:"structWithAllOptionalFields"` // want "field TestStructs.StructWithAllOptionalFields should have the omitempty tag." "field TestStructs.StructWithAllOptionalFields should be a pointer."

	StructWithAllOptionalFieldsWithOmitEmpty StructWithAllOptionalFields `json:"structWithAllOptionalFieldsWithOmitEmpty,omitempty"` // want "field TestStructs.StructWithAllOptionalFieldsWithOmitEmpty should be a pointer."

	StructPtrWithAllOptionalFields *StructWithAllOptionalFields `json:"structPtrWithAllOptionalFields"` // want "field TestStructs.StructPtrWithAllOptionalFields should have the omitempty tag."

	StructPtrWithAllOptionalFieldsWithOmitEmpty *StructWithAllOptionalFields `json:"structPtrWithAllOptionalFieldsWithOmitEmpty,omitempty"`

	// StructWithMinProperties has a zero value of {}, which is not valid because the MinProperties marker is not satisfied.

	StructWithMinProperties StructWithMinProperties `json:"structWithMinProperties"` // want "field TestStructs.StructWithMinProperties should have the omitempty tag." "field TestStructs.StructWithMinProperties should be a pointer."

	StructWithMinPropertiesWithOmitEmpty StructWithMinProperties `json:"structWithMinPropertiesWithOmitEmpty,omitempty"` // want "field TestStructs.StructWithMinPropertiesWithOmitEmpty should be a pointer."

	StructPtrWithMinProperties *StructWithMinProperties `json:"structPtrWithMinProperties"` // want "field TestStructs.StructPtrWithMinProperties should have the omitempty tag."

	StructPtrWithMinPropertiesWithOmitEmpty *StructWithMinProperties `json:"structPtrWithMinPropertiesWithOmitEmpty,omitempty"`

	// StructWithNonOmittedFields has a zero value of {"string":"", "int":0}, which is valid because all fields are required.

	StructWithNonOmittedFields StructWithNonOmittedFields `json:"structWithNonOmittedFields"` // want "field TestStructs.StructWithNonOmittedFields should have the omitempty tag." "field TestStructs.StructWithNonOmittedFields should be a pointer."

	StructWithNonOmittedFieldsWithOmitEmpty StructWithNonOmittedFields `json:"structWithNonOmittedFieldsWithOmitEmpty,omitempty"` // want "field TestStructs.StructWithNonOmittedFieldsWithOmitEmpty should be a pointer."

	StructPtrWithNonOmittedFields *StructWithNonOmittedFields `json:"structPtrWithNonOmittedFields"` // want "field TestStructs.StructPtrWithNonOmittedFields should have the omitempty tag."

	StructPtrWithNonOmittedFieldsWithOmitEmpty *StructWithNonOmittedFields `json:"structPtrWithNonOmittedFieldsWithOmitEmpty,omitempty"`

	// StructWithNonOmittedFieldsAndMinProperties has a zero value of {"string":"", "int":0}, which is valid because the MinProperties marker is satisfied.

	StructWithNonOmittedFieldsAndMinProperties StructWithNonOmittedFieldsAndMinProperties `json:"structWithOneNonOmittedFieldAndMinProperties"` // want "field TestStructs.StructWithNonOmittedFieldsAndMinProperties should have the omitempty tag." "field TestStructs.StructWithNonOmittedFieldsAndMinProperties should be a pointer."

	StructWithNonOmittedFieldsAndMinPropertiesWithOmitEmpty StructWithNonOmittedFieldsAndMinProperties `json:"structWithOneNonOmittedFieldAndMinPropertiesAndOmitEmpty,omitempty"` // want "field TestStructs.StructWithNonOmittedFieldsAndMinPropertiesWithOmitEmpty should be a pointer."

	StructPtrWithNonOmittedFieldsAndMinProperties *StructWithNonOmittedFieldsAndMinProperties `json:"structPtrWithNonOmittedFieldsAndMinProperties"` // want "field TestStructs.StructPtrWithNonOmittedFieldsAndMinProperties should have the omitempty tag."

	StructPtrWithNonOmittedFieldsAndMinPropertiesWithOmitEmpty *StructWithNonOmittedFieldsAndMinProperties `json:"structPtrWithNonOmittedFieldsAndMinPropertiesWithOmitEmpty,omitempty"`

	// StructWithOneNonOmittedFieldAndMinProperties has a zero value of {"string":""}, which is not valid because the MinProperties marker is not satisfied.

	StructWithOneNonOmittedFieldAndMinProperties StructWithOneNonOmittedFieldAndMinProperties `json:"structWithOneNonOmittedFieldAndMinPropertiesAndOmitEmpty"` // want "field TestStructs.StructWithOneNonOmittedFieldAndMinProperties should have the omitempty tag." "field TestStructs.StructWithOneNonOmittedFieldAndMinProperties should be a pointer."

	StructWithOneNonOmittedFieldAndMinPropertiesWithOmitEmpty StructWithOneNonOmittedFieldAndMinProperties `json:"structWithOneNonOmittedFieldAndMinPropertiesAndOmitEmpty,omitempty"` // want "field TestStructs.StructWithOneNonOmittedFieldAndMinPropertiesWithOmitEmpty should be a pointer."

	StructPtrWithOneNonOmittedFieldAndMinProperties *StructWithOneNonOmittedFieldAndMinProperties `json:"structPtrWithOneNonOmittedFieldAndMinProperties"` // want "field TestStructs.StructPtrWithOneNonOmittedFieldAndMinProperties should have the omitempty tag."

	StructPtrWithOneNonOmittedFieldAndMinPropertiesWithOmitEmpty *StructWithOneNonOmittedFieldAndMinProperties `json:"structPtrWithOneNonOmittedFieldAndMinPropertiesWithOmitEmpty,omitempty"`

	// StructWithOmittedRequiredField has a zero value of {}, which is not valid because the required marker is not satisfied.

	StructWithOmittedRequiredField StructWithOmittedRequiredField `json:"structWithOmittedRequiredField"` // want "field TestStructs.StructWithOmittedRequiredField should have the omitempty tag." "field TestStructs.StructWithOmittedRequiredField should be a pointer."

	StructWithOmittedRequiredFieldWithOmitEmpty StructWithOmittedRequiredField `json:"structWithOmittedRequiredFieldWithOmitEmpty,omitempty"` // want "field TestStructs.StructWithOmittedRequiredFieldWithOmitEmpty should be a pointer."

	StructPtrWithOmittedRequiredField *StructWithOmittedRequiredField `json:"structPtrWithOmittedRequiredField"` // want "field TestStructs.StructPtrWithOmittedRequiredField should have the omitempty tag."

	StructPtrWithOmittedRequiredFieldWithOmitEmpty *StructWithOmittedRequiredField `json:"structPtrWithOmittedRequiredFieldWithOmitEmpty,omitempty"`
}

type StructWithAllOptionalFields struct {
	// +optional
	String string `json:"string,omitempty"` // want "field StructWithAllOptionalFields.String should be a pointer."

	// +optional
	StringPtr *string `json:"stringPtr,omitempty"`

	// +optional
	Int int `json:"int,omitempty"` // want "field StructWithAllOptionalFields.Int should be a pointer."

	// +optional
	IntPtr *int `json:"intPtr,omitempty"`
}

// +kubebuilder:validation:MinProperties=1
type StructWithMinProperties struct {
	// +kubebuilder:validation:MinProperties=1
	// +optional
	Map map[string]string `json:"map,omitempty"`
}

type StructWithNonOmittedFields struct {
	// +required
	String string `json:"string"` // want "field StructWithNonOmittedFields.String should be a pointer." "field StructWithNonOmittedFields.String should have the omitempty tag."

	// +required
	Int int32 `json:"int"` // want "field StructWithNonOmittedFields.Int should be a pointer." "field StructWithNonOmittedFields.Int should have the omitempty tag."
}

// Struct with non-omitted fields and minProperties marker.
// Because there is no omitempty, and the zero values are valid, the zero value here is `{"string:"", "int":0}`.
// This means the MinProperties marker is satisfied even when the object is the zero value.
// +kubebuilder:validation:MinProperties=2
type StructWithNonOmittedFieldsAndMinProperties struct {
	// +required
	String string `json:"string"` // want "field StructWithNonOmittedFieldsAndMinProperties.String should be a pointer." "field StructWithNonOmittedFieldsAndMinProperties.String should have the omitempty tag."

	// +required
	Int int32 `json:"int"` // want "field StructWithNonOmittedFieldsAndMinProperties.Int should be a pointer." "field StructWithNonOmittedFieldsAndMinProperties.Int should have the omitempty tag."
}

// Struct with one non-omitted field, and one omitted field and minProperties marker.
// The zero value of the struct is `{"string:""}` which is not valid because it does not satisfy the MinProperties marker.
// +kubebuilder:validation:MinProperties=2
type StructWithOneNonOmittedFieldAndMinProperties struct {
	// +required
	String string `json:"string"` // want "field StructWithOneNonOmittedFieldAndMinProperties.String should be a pointer." "field StructWithOneNonOmittedFieldAndMinProperties.String should have the omitempty tag."

	// +optional
	Int int32 `json:"int,omitempty"` // want "field StructWithOneNonOmittedFieldAndMinProperties.Int should be a pointer."
}

// Struct with an omitted required field.
// The zero value of the struct is `{}` which is not valid because it does not satisfy the required marker on the string field.
type StructWithOmittedRequiredField struct {
	// +required
	String string `json:"string,omitempty"` // want "field StructWithOmittedRequiredField.String should be a pointer."
}
