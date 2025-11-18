package a

type TestStrings struct {
	String string `json:"string"` // want "field TestStrings.String should have the omitempty tag." "field TestStrings.String should be a pointer."

	StringWithOmitEmpty string `json:"stringWithOmitEmpty,omitempty"` // want "field TestStrings.StringWithOmitEmpty should be a pointer."

	StringPtr *string `json:"stringPtr"` // want "field TestStrings.StringPtr should have the omitempty tag."

	StringPtrWithOmitEmpty *string `json:"stringPtrWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:MinLength=1
	StringWithMinLength string `json:"stringWithMinLength"` // want "field TestStrings.StringWithMinLength should have the omitempty tag." "field TestStrings.StringWithMinLength should be a pointer."

	// +kubebuilder:validation:MinLength=1
	StringWithMinLengthWithOmitEmpty string `json:"stringWithMinLengthWithOmitEmpty,omitempty"` // want "field TestStrings.StringWithMinLengthWithOmitEmpty should be a pointer."

	// +kubebuilder:validation:MinLength=1
	StringPtrWithMinLength *string `json:"stringPtrWithMinLength"` // want "field TestStrings.StringPtrWithMinLength should have the omitempty tag."

	// +kubebuilder:validation:MinLength=1
	StringPtrWithMinLengthWithOmitEmpty *string `json:"stringPtrWithMinLengthWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:MinLength=0
	StringWithMinLength0 string `json:"stringWithMinLength0"` // want "field TestStrings.StringWithMinLength0 should have the omitempty tag." "field TestStrings.StringWithMinLength0 should be a pointer."

	// +kubebuilder:validation:MinLength=0
	StringWithMinLength0WithOmitEmpty string `json:"stringWithMinLength0WithOmitEmpty,omitempty"` // want "field TestStrings.StringWithMinLength0WithOmitEmpty should be a pointer."

	// +kubebuilder:validation:MinLength=0
	StringPtrWithMinLength0 *string `json:"stringPtrWithMinLength0"` // want "field TestStrings.StringPtrWithMinLength0 should have the omitempty tag."

	// +kubebuilder:validation:MinLength=0
	StringPtrWithMinLength0WithOmitEmpty *string `json:"stringPtrWithMinLength0WithOmitEmpty,omitempty"`

	// +kubebuilder:validation:Enum=a;b;c
	EnumString string `json:"enumString"` // want "field TestStrings.EnumString should have the omitempty tag." "field TestStrings.EnumString should be a pointer."

	// +kubebuilder:validation:Enum=a;b;c
	EnumStringWithOmitEmpty string `json:"enumStringWithOmitEmpty,omitempty"` // want "field TestStrings.EnumStringWithOmitEmpty should be a pointer."

	// +kubebuilder:validation:Enum=a;b;c
	EnumStringPtr *string `json:"enumStringPtr"` // want "field TestStrings.EnumStringPtr should have the omitempty tag."

	// +kubebuilder:validation:Enum=a;b;c
	EnumStringPtrWithOmitEmpty *string `json:"enumStringPtrWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:Enum=a;b;c;""
	EnumValidEmptytring string `json:"enumValidEmptytring"` // want "field TestStrings.EnumValidEmptytring should have the omitempty tag." "field TestStrings.EnumValidEmptytring should be a pointer."

	// +kubebuilder:validation:Enum=a;b;c;""
	EnumValidEmptyStringWithOmitEmpty string `json:"enumValidEmptyStringWithOmitEmpty,omitempty"` // want "field TestStrings.EnumValidEmptyStringWithOmitEmpty should be a pointer."

	// +kubebuilder:validation:Enum=a;b;c;""
	EnumValidEmptyStringPtr *string `json:"enumValidEmptyStringPtr"` // want "field TestStrings.EnumValidEmptyStringPtr should have the omitempty tag."

	// +kubebuilder:validation:Enum=a;b;c;""
	EnumValidEmptyStringPtrWithOmitEmpty *string `json:"enumValidEmptyStringPtrWithOmitEmpty,omitempty"`
}
