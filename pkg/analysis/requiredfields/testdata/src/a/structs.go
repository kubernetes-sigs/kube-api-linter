package a

type TestStructs struct {
	// StructWithAllOptionalFields has a zero value of {}, which is valid because all fields are optional.

	// +required
	StructWithAllOptionalFields StructWithAllOptionalFields `json:"structWithAllOptionalFields"` // want "field TestStructs.StructWithAllOptionalFields should have the omitempty tag." "field TestStructs.StructWithAllOptionalFields has a valid zero value \\({}\\), but the validation is not complete \\(e.g. min properties/adding required fields\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +required
	StructWithAllOptionalFieldsWithOmitEmpty StructWithAllOptionalFields `json:"structWithAllOptionalFieldsWithOmitEmpty,omitempty"` // want "field TestStructs.StructWithAllOptionalFieldsWithOmitEmpty has a valid zero value \\({}\\), but the validation is not complete \\(e.g. min properties/adding required fields\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +required
	StructPtrWithAllOptionalFields *StructWithAllOptionalFields `json:"structPtrWithAllOptionalFields"` // want "field TestStructs.StructPtrWithAllOptionalFields should have the omitempty tag."

	// +required
	StructPtrWithAllOptionalFieldsWithOmitEmpty *StructWithAllOptionalFields `json:"structPtrWithAllOptionalFieldsWithOmitEmpty,omitempty"`

	// StructWithMinProperties has a zero value of {}, which is not valid because the MinProperties marker is not satisfied.

	// +required
	StructWithMinProperties StructWithMinProperties `json:"structWithMinProperties"` // want "field TestStructs.StructWithMinProperties does not allow the zero value. It must have the omitzero tag."

	// +required
	StructWithMinPropertiesWithOmitEmpty StructWithMinProperties `json:"structWithMinPropertiesWithOmitEmpty,omitempty"` // want "field TestStructs.StructWithMinPropertiesWithOmitEmpty does not allow the zero value. It must have the omitzero tag."

	// +required
	StructPtrWithMinProperties *StructWithMinProperties `json:"structPtrWithMinProperties"` // want "field TestStructs.StructPtrWithMinProperties does not allow the zero value. It must have the omitzero tag." "field TestStructs.StructPtrWithMinProperties does not allow the zero value. The field does not need to be a pointer."

	// +required
	StructPtrWithMinPropertiesWithOmitEmpty *StructWithMinProperties `json:"structPtrWithMinPropertiesWithOmitEmpty,omitempty"` // want "field TestStructs.StructPtrWithMinPropertiesWithOmitEmpty does not allow the zero value. It must have the omitzero tag." "field TestStructs.StructPtrWithMinPropertiesWithOmitEmpty does not allow the zero value. The field does not need to be a pointer."

	// StructWithNonOmittedFields has a zero value of {"string":"", "int":0}, which is valid because all fields are required.

	// +required
	StructWithNonOmittedFields StructWithNonOmittedFields `json:"structWithNonOmittedFields"` // want "field TestStructs.StructWithNonOmittedFields should have the omitempty tag." "field TestStructs.StructWithNonOmittedFields has a valid zero value \\({\"string\": \"\", \"int\": 0}\\) and should be a pointer."

	// +required
	StructWithNonOmittedFieldsWithOmitEmpty StructWithNonOmittedFields `json:"structWithNonOmittedFieldsWithOmitEmpty,omitempty"` // want "field TestStructs.StructWithNonOmittedFieldsWithOmitEmpty has a valid zero value \\({\"string\": \"\", \"int\": 0}\\) and should be a pointer."

	// +required
	StructPtrWithNonOmittedFields *StructWithNonOmittedFields `json:"structPtrWithNonOmittedFields"` // want "field TestStructs.StructPtrWithNonOmittedFields should have the omitempty tag."

	// +required
	StructPtrWithNonOmittedFieldsWithOmitEmpty *StructWithNonOmittedFields `json:"structPtrWithNonOmittedFieldsWithOmitEmpty,omitempty"`

	// StructWithNonOmittedFieldsAndMinProperties has a zero value of {"string":"", "int":0}, which is valid because the MinProperties marker is satisfied.

	// +required
	StructWithNonOmittedFieldsAndMinProperties StructWithNonOmittedFieldsAndMinProperties `json:"structWithOneNonOmittedFieldAndMinProperties"` // want "field TestStructs.StructWithNonOmittedFieldsAndMinProperties should have the omitempty tag." "field TestStructs.StructWithNonOmittedFieldsAndMinProperties has a valid zero value \\({\"string\": \"\", \"int\": 0}\\) and should be a pointer."

	// +required
	StructWithNonOmittedFieldsAndMinPropertiesWithOmitEmpty StructWithNonOmittedFieldsAndMinProperties `json:"structWithOneNonOmittedFieldAndMinPropertiesAndOmitEmpty,omitempty"` // want "field TestStructs.StructWithNonOmittedFieldsAndMinPropertiesWithOmitEmpty has a valid zero value \\({\"string\": \"\", \"int\": 0}\\) and should be a pointer."

	// +required
	StructPtrWithNonOmittedFieldsAndMinProperties *StructWithNonOmittedFieldsAndMinProperties `json:"structPtrWithNonOmittedFieldsAndMinProperties"` // want "field TestStructs.StructPtrWithNonOmittedFieldsAndMinProperties should have the omitempty tag."

	// +required
	StructPtrWithNonOmittedFieldsAndMinPropertiesWithOmitEmpty *StructWithNonOmittedFieldsAndMinProperties `json:"structPtrWithNonOmittedFieldsAndMinPropertiesWithOmitEmpty,omitempty"`

	// StructWithOneNonOmittedFieldAndMinProperties has a zero value of {"string":""}, which is not valid because the MinProperties marker is not satisfied.

	// +required
	StructWithOneNonOmittedFieldAndMinProperties StructWithOneNonOmittedFieldAndMinProperties `json:"structWithOneNonOmittedFieldAndMinPropertiesAndOmitEmpty"` // want "field TestStructs.StructWithOneNonOmittedFieldAndMinProperties does not allow the zero value. It must have the omitzero tag."

	// +required
	StructWithOneNonOmittedFieldAndMinPropertiesWithOmitEmpty StructWithOneNonOmittedFieldAndMinProperties `json:"structWithOneNonOmittedFieldAndMinPropertiesAndOmitEmpty,omitempty"` // want "field TestStructs.StructWithOneNonOmittedFieldAndMinPropertiesWithOmitEmpty does not allow the zero value. It must have the omitzero tag."

	// +required
	StructPtrWithOneNonOmittedFieldAndMinProperties *StructWithOneNonOmittedFieldAndMinProperties `json:"structPtrWithOneNonOmittedFieldAndMinProperties"` // want "field TestStructs.StructPtrWithOneNonOmittedFieldAndMinProperties does not allow the zero value. It must have the omitzero tag." "field TestStructs.StructPtrWithOneNonOmittedFieldAndMinProperties does not allow the zero value. The field does not need to be a pointer."

	// +required
	StructPtrWithOneNonOmittedFieldAndMinPropertiesWithOmitEmpty *StructWithOneNonOmittedFieldAndMinProperties `json:"structPtrWithOneNonOmittedFieldAndMinPropertiesWithOmitEmpty,omitempty"` // want "field TestStructs.StructPtrWithOneNonOmittedFieldAndMinPropertiesWithOmitEmpty does not allow the zero value. It must have the omitzero tag." "field TestStructs.StructPtrWithOneNonOmittedFieldAndMinPropertiesWithOmitEmpty does not allow the zero value. The field does not need to be a pointer."

	// StructWithOmittedRequiredField has a zero value of {}, which is not valid because the required marker is not satisfied.

	// +required
	StructWithOmittedRequiredField StructWithOmittedRequiredField `json:"structWithOmittedRequiredField"` // want "field TestStructs.StructWithOmittedRequiredField does not allow the zero value. It must have the omitzero tag."

	// +required
	StructWithOmittedRequiredFieldWithOmitEmpty StructWithOmittedRequiredField `json:"structWithOmittedRequiredFieldWithOmitEmpty,omitempty"` // want "field TestStructs.StructWithOmittedRequiredFieldWithOmitEmpty does not allow the zero value. It must have the omitzero tag."

	// +required
	StructPtrWithOmittedRequiredField *StructWithOmittedRequiredField `json:"structPtrWithOmittedRequiredField"` // want "field TestStructs.StructPtrWithOmittedRequiredField does not allow the zero value. It must have the omitzero tag." "field TestStructs.StructPtrWithOmittedRequiredField does not allow the zero value. The field does not need to be a pointer."

	// +required
	StructPtrWithOmittedRequiredFieldWithOmitEmpty *StructWithOmittedRequiredField `json:"structPtrWithOmittedRequiredFieldWithOmitEmpty,omitempty"` // want "field TestStructs.StructPtrWithOmittedRequiredFieldWithOmitEmpty does not allow the zero value. It must have the omitzero tag." "field TestStructs.StructPtrWithOmittedRequiredFieldWithOmitEmpty does not allow the zero value. The field does not need to be a pointer."

	// StructWithExactlyOneOf has a zero value of {}, which is not valid because the ExactlyOneOf marker requires exactly one field to be set.

	// +required
	StructWithExactlyOneOf StructWithExactlyOneOf `json:"structWithExactlyOneOf"` // want "field TestStructs.StructWithExactlyOneOf does not allow the zero value. It must have the omitzero tag."

	// +required
	StructWithExactlyOneOfWithOmitEmpty StructWithExactlyOneOf `json:"structWithExactlyOneOfWithOmitEmpty,omitempty"` // want "field TestStructs.StructWithExactlyOneOfWithOmitEmpty does not allow the zero value. It must have the omitzero tag."

	// +required
	StructPtrWithExactlyOneOf *StructWithExactlyOneOf `json:"structPtrWithExactlyOneOf"` // want "field TestStructs.StructPtrWithExactlyOneOf does not allow the zero value. It must have the omitzero tag." "field TestStructs.StructPtrWithExactlyOneOf does not allow the zero value. The field does not need to be a pointer."

	// +required
	StructPtrWithExactlyOneOfWithOmitEmpty *StructWithExactlyOneOf `json:"structPtrWithExactlyOneOfWithOmitEmpty,omitempty"` // want "field TestStructs.StructPtrWithExactlyOneOfWithOmitEmpty does not allow the zero value. It must have the omitzero tag." "field TestStructs.StructPtrWithExactlyOneOfWithOmitEmpty does not allow the zero value. The field does not need to be a pointer."

	// StructWithAtLeastOneOf has a zero value of {}, which is not valid because the AtLeastOneOf marker requires at least one field to be set.

	// +required
	StructWithAtLeastOneOf StructWithAtLeastOneOf `json:"structWithAtLeastOneOf"` // want "field TestStructs.StructWithAtLeastOneOf does not allow the zero value. It must have the omitzero tag."

	// +required
	StructWithAtLeastOneOfWithOmitEmpty StructWithAtLeastOneOf `json:"structWithAtLeastOneOfWithOmitEmpty,omitempty"` // want "field TestStructs.StructWithAtLeastOneOfWithOmitEmpty does not allow the zero value. It must have the omitzero tag."

	// +required
	StructPtrWithAtLeastOneOf *StructWithAtLeastOneOf `json:"structPtrWithAtLeastOneOf"` // want "field TestStructs.StructPtrWithAtLeastOneOf does not allow the zero value. It must have the omitzero tag." "field TestStructs.StructPtrWithAtLeastOneOf does not allow the zero value. The field does not need to be a pointer."

	// +required
	StructPtrWithAtLeastOneOfWithOmitEmpty *StructWithAtLeastOneOf `json:"structPtrWithAtLeastOneOfWithOmitEmpty,omitempty"` // want "field TestStructs.StructPtrWithAtLeastOneOfWithOmitEmpty does not allow the zero value. It must have the omitzero tag." "field TestStructs.StructPtrWithAtLeastOneOfWithOmitEmpty does not allow the zero value. The field does not need to be a pointer."

	// StructWithExactlyOneOfOneWithoutOmitEmpty has one union member without omitempty.
	// That member serializes as null at zero value; has(null) = false in CEL.
	// ExactlyOneOf is still violated.

	// +required
	StructWithExactlyOneOfOneWithoutOmitEmpty StructWithExactlyOneOfOneWithoutOmitEmpty `json:"structWithExactlyOneOfOneWithoutOmitEmpty"` // want "field TestStructs.StructWithExactlyOneOfOneWithoutOmitEmpty does not allow the zero value. It must have the omitzero tag."

	// StructWithExactlyOneOfNonUnionNonOmitted has a non-union field without omitempty
	// alongside union fields. The non-union field does not satisfy ExactlyOneOf.

	// +required
	StructWithExactlyOneOfNonUnionNonOmitted StructWithExactlyOneOfNonUnionNonOmitted `json:"structWithExactlyOneOfNonUnionNonOmitted"` // want "field TestStructs.StructWithExactlyOneOfNonUnionNonOmitted does not allow the zero value. It must have the omitzero tag."

	// StructWithMinPropertiesAndExactlyOneOf: MinProperties=1 takes precedence in
	// resolveEffectiveMinProperties; ExactlyOneOf is redundant. Behavior is same as MinProperties=1.

	// +required
	StructWithMinPropertiesAndExactlyOneOf StructWithMinPropertiesAndExactlyOneOf `json:"structWithMinPropertiesAndExactlyOneOf"` // want "field TestStructs.StructWithMinPropertiesAndExactlyOneOf does not allow the zero value. It must have the omitzero tag."
}

type StructWithAllOptionalFields struct {
	// +optional
	String string `json:"string,omitempty"`

	// +optional
	StringPtr *string `json:"stringPtr,omitempty"`

	// +optional
	Int int `json:"int,omitempty"`

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
	String string `json:"string"` // want "field StructWithNonOmittedFields.String has a valid zero value \\(\"\"\\), but the validation is not complete \\(e.g. minimum length\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer." "field StructWithNonOmittedFields.String should have the omitempty tag."

	// +required
	Int int32 `json:"int"` // want "field StructWithNonOmittedFields.Int has a valid zero value \\(0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer." "field StructWithNonOmittedFields.Int should have the omitempty tag."
}

// Struct with non-omitted fields and minProperties marker.
// Because there is no omitempty, and the zero values are valid, the zero value here is `{"string:"", "int":0}`.
// This means the MinProperties marker is satisfied even when the object is the zero value.
// +kubebuilder:validation:MinProperties=2
type StructWithNonOmittedFieldsAndMinProperties struct {
	// +required
	String string `json:"string"` // want "field StructWithNonOmittedFieldsAndMinProperties.String has a valid zero value \\(\"\"\\), but the validation is not complete \\(e.g. minimum length\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer." "field StructWithNonOmittedFieldsAndMinProperties.String should have the omitempty tag."

	// +required
	Int int32 `json:"int"` // want "field StructWithNonOmittedFieldsAndMinProperties.Int has a valid zero value \\(0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer." "field StructWithNonOmittedFieldsAndMinProperties.Int should have the omitempty tag."
}

// Struct with one non-omitted field, and one omitted field and minProperties marker.
// The zero value of the struct is `{"string:""}` which is not valid because it does not satisfy the MinProperties marker.
// +kubebuilder:validation:MinProperties=2
type StructWithOneNonOmittedFieldAndMinProperties struct {
	// +required
	String string `json:"string"` // want "field StructWithOneNonOmittedFieldAndMinProperties.String has a valid zero value \\(\"\"\\), but the validation is not complete \\(e.g. minimum length\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer." "field StructWithOneNonOmittedFieldAndMinProperties.String should have the omitempty tag."

	// +optional
	Int int32 `json:"int,omitempty"`
}

// Struct with an omitted required field.
// The zero value of the struct is `{}` which is not valid because it does not satisfy the required marker on the string field.
type StructWithOmittedRequiredField struct {
	// +required
	String string `json:"string,omitempty"` // want "field StructWithOmittedRequiredField.String has a valid zero value \\(\"\"\\), but the validation is not complete \\(e.g. minimum length\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."
}

// StructWithExactlyOneOf has ExactlyOneOf across both fields, all with omitempty.
// Zero value {} violates ExactlyOneOf — no union member is set.
// +kubebuilder:validation:ExactlyOneOf=serviceKeyRef;tokenRef
type StructWithExactlyOneOf struct {
	// +optional
	ServiceKeyRef *string `json:"serviceKeyRef,omitempty"`
	// +optional
	TokenRef *string `json:"tokenRef,omitempty"`
}

// StructWithAtLeastOneOf has AtLeastOneOf across both fields, all with omitempty.
// +kubebuilder:validation:AtLeastOneOf=serviceKeyRef;tokenRef
type StructWithAtLeastOneOf struct {
	// +optional
	ServiceKeyRef *string `json:"serviceKeyRef,omitempty"`
	// +optional
	TokenRef *string `json:"tokenRef,omitempty"`
}

// StructWithExactlyOneOfOneWithoutOmitEmpty has one union member without omitempty.
// That member serializes as null at zero value; has(null) = false in CEL.
// ExactlyOneOf is still violated.
// +kubebuilder:validation:ExactlyOneOf=serviceKeyRef;tokenRef
type StructWithExactlyOneOfOneWithoutOmitEmpty struct {
	// +optional
	ServiceKeyRef *string `json:"serviceKeyRef"` // no omitempty — serializes as null
	// +optional
	TokenRef *string `json:"tokenRef,omitempty"`
}

// StructWithExactlyOneOfNonUnionNonOmitted has a non-union field without omitempty
// alongside union fields. The non-union field does not satisfy ExactlyOneOf.
// +kubebuilder:validation:ExactlyOneOf=fieldA;fieldB
type StructWithExactlyOneOfNonUnionNonOmitted struct {
	// Non-union, non-omitted optional field.
	Name   string  `json:"name"` // no omitempty, no +required
	// +optional
	FieldA *string `json:"fieldA,omitempty"`
	// +optional
	FieldB *string `json:"fieldB,omitempty"`
}

// StructWithMinPropertiesAndExactlyOneOf has both explicit MinProperties=1 and ExactlyOneOf.
// resolveEffectiveMinProperties returns the explicit MinProperties=1 (not synthesized from union).
// The behavior should be identical to MinProperties=1 alone.
// +kubebuilder:validation:MinProperties=1
// +kubebuilder:validation:ExactlyOneOf=fieldA;fieldB
type StructWithMinPropertiesAndExactlyOneOf struct {
	// +optional
	FieldA *string `json:"fieldA,omitempty"`
	// +optional
	FieldB *string `json:"fieldB,omitempty"`
}
