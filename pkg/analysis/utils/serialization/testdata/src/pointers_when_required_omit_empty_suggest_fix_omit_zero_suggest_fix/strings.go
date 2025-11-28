package a

type TestStrings struct {
	String string `json:"string"` // want "field TestStrings.String should have the omitempty tag." "field TestStrings.String has a valid zero value \\(\"\"\\), but the validation is not complete \\(e.g. minimum length\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	StringWithOmitEmpty string `json:"stringWithOmitEmpty,omitempty"` // want "field TestStrings.StringWithOmitEmpty has a valid zero value \\(\"\"\\), but the validation is not complete \\(e.g. minimum length\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	StringPtr *string `json:"stringPtr"` // want "field TestStrings.StringPtr should have the omitempty tag."

	StringPtrWithOmitEmpty *string `json:"stringPtrWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:MinLength=1
	StringWithMinLength string `json:"stringWithMinLength"` // want "field TestStrings.StringWithMinLength should have the omitempty tag."

	// +kubebuilder:validation:MinLength=1
	StringWithMinLengthWithOmitEmpty string `json:"stringWithMinLengthWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:MinLength=1
	StringPtrWithMinLength *string `json:"stringPtrWithMinLength"` // want "field TestStrings.StringPtrWithMinLength should have the omitempty tag." "field TestStrings.StringPtrWithMinLength does not allow the zero value. The field does not need to be a pointer."

	// +kubebuilder:validation:MinLength=1
	StringPtrWithMinLengthWithOmitEmpty *string `json:"stringPtrWithMinLengthWithOmitEmpty,omitempty"` // want "field TestStrings.StringPtrWithMinLengthWithOmitEmpty does not allow the zero value. The field does not need to be a pointer."

	// +kubebuilder:validation:MinLength=0
	StringWithMinLength0 string `json:"stringWithMinLength0"` // want "field TestStrings.StringWithMinLength0 should have the omitempty tag." "field TestStrings.StringWithMinLength0 has a valid zero value \\(\"\"\\) and should be a pointer."

	// +kubebuilder:validation:MinLength=0
	StringWithMinLength0WithOmitEmpty string `json:"stringWithMinLength0WithOmitEmpty,omitempty"` // want "field TestStrings.StringWithMinLength0WithOmitEmpty has a valid zero value \\(\"\"\\) and should be a pointer."

	// +kubebuilder:validation:MinLength=0
	StringPtrWithMinLength0 *string `json:"stringPtrWithMinLength0"` // want "field TestStrings.StringPtrWithMinLength0 should have the omitempty tag."

	// +kubebuilder:validation:MinLength=0
	StringPtrWithMinLength0WithOmitEmpty *string `json:"stringPtrWithMinLength0WithOmitEmpty,omitempty"`

	// +kubebuilder:validation:Enum=a;b;c
	EnumString string `json:"enumString"` // want "field TestStrings.EnumString should have the omitempty tag."

	// +kubebuilder:validation:Enum=a;b;c
	EnumStringWithOmitEmpty string `json:"enumStringWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:Enum=a;b;c
	EnumStringPtr *string `json:"enumStringPtr"` // want "field TestStrings.EnumStringPtr should have the omitempty tag." "field TestStrings.EnumStringPtr does not allow the zero value. The field does not need to be a pointer."

	// +kubebuilder:validation:Enum=a;b;c
	EnumStringPtrWithOmitEmpty *string `json:"enumStringPtrWithOmitEmpty,omitempty"` // want "field TestStrings.EnumStringPtrWithOmitEmpty does not allow the zero value. The field does not need to be a pointer."

	// +kubebuilder:validation:Enum=a;b;c;""
	EnumValidEmptytring string `json:"enumValidEmptytring"` // want "field TestStrings.EnumValidEmptytring should have the omitempty tag." "field TestStrings.EnumValidEmptytring has a valid zero value \\(\"\"\\) and should be a pointer."

	// +kubebuilder:validation:Enum=a;b;c;""
	EnumValidEmptyStringWithOmitEmpty string `json:"enumValidEmptyStringWithOmitEmpty,omitempty"` // want "field TestStrings.EnumValidEmptyStringWithOmitEmpty has a valid zero value \\(\"\"\\) and should be a pointer."

	// +kubebuilder:validation:Enum=a;b;c;""
	EnumValidEmptyStringPtr *string `json:"enumValidEmptyStringPtr"` // want "field TestStrings.EnumValidEmptyStringPtr should have the omitempty tag."

	// +kubebuilder:validation:Enum=a;b;c;""
	EnumValidEmptyStringPtrWithOmitEmpty *string `json:"enumValidEmptyStringPtrWithOmitEmpty,omitempty"`
}
