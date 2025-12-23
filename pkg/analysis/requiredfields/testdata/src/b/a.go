package c

type Foo struct {
	// +required
	RequiredStructWithRequiredFieldWithoutOmitZero StructWithRequiredField `json:"requiredStructWithRequiredFieldWithoutOmitZero"` // want "field Foo.RequiredStructWithRequiredFieldWithoutOmitZero does not allow the zero value. It must have the omitzero tag."

	// +required
	RequiredStructWithRequiredFieldWithOmitZero StructWithRequiredField `json:"requiredStructWithRequiredFieldWithOmitZero,omitzero"`
}

type StructWithRequiredField struct {
	// Does not allow the zero value.
	// +required
	// +kubebuilder:validation:Enum=A;B;C
	EnumStringWithoutOmitEmpty string `json:"enumStringWithoutOmitEmpty"` // want "field StructWithRequiredField.EnumStringWithoutOmitEmpty does not allow the zero value. It must have the omitempty tag."
}
