package a

type TestStrings struct {
	String string `json:"string"` // want "field String should have the omitempty tag." "field String should be a pointer."

	StringWithOmitEmpty string `json:"stringWithOmitEmpty,omitempty"` // want "field StringWithOmitEmpty should be a pointer."

	StringPtr *string `json:"stringPtr"` // want "field StringPtr should have the omitempty tag."

	StringPtrWithOmitEmpty *string `json:"stringPtrWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:MinLength=1
	StringWithMinLength string `json:"stringWithMinLength"` // want "field StringWithMinLength should have the omitempty tag." "field StringWithMinLength should be a pointer."

	// +kubebuilder:validation:MinLength=1
	StringWithMinLengthWithOmitEmpty string `json:"stringWithMinLengthWithOmitEmpty,omitempty"` // want "field StringWithMinLengthWithOmitEmpty should be a pointer."

	// +kubebuilder:validation:MinLength=1
	StringPtrWithMinLength *string `json:"stringPtrWithMinLength"` // want "field StringPtrWithMinLength should have the omitempty tag."

	// +kubebuilder:validation:MinLength=1
	StringPtrWithMinLengthWithOmitEmpty *string `json:"stringPtrWithMinLengthWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:MinLength=0
	StringWithMinLength0 string `json:"stringWithMinLength0"` // want "field StringWithMinLength0 should have the omitempty tag." "field StringWithMinLength0 should be a pointer."

	// +kubebuilder:validation:MinLength=0
	StringWithMinLength0WithOmitEmpty string `json:"stringWithMinLength0WithOmitEmpty,omitempty"` // want "field StringWithMinLength0WithOmitEmpty should be a pointer."

	// +kubebuilder:validation:MinLength=0
	StringPtrWithMinLength0 *string `json:"stringPtrWithMinLength0"` // want "field StringPtrWithMinLength0 should have the omitempty tag."

	// +kubebuilder:validation:MinLength=0
	StringPtrWithMinLength0WithOmitEmpty *string `json:"stringPtrWithMinLength0WithOmitEmpty,omitempty"`

	// +kubebuilder:validation:Enum=a;b;c
	EnumString string `json:"enumString"` // want "field EnumString should have the omitempty tag." "field EnumString should be a pointer."

	// +kubebuilder:validation:Enum=a;b;c
	EnumStringWithOmitEmpty string `json:"enumStringWithOmitEmpty,omitempty"` // want "field EnumStringWithOmitEmpty should be a pointer."

	// +kubebuilder:validation:Enum=a;b;c
	EnumStringPtr *string `json:"enumStringPtr"` // want "field EnumStringPtr should have the omitempty tag."

	// +kubebuilder:validation:Enum=a;b;c
	EnumStringPtrWithOmitEmpty *string `json:"enumStringPtrWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:Enum=a;b;c;""
	EnumValidEmptytring string `json:"enumValidEmptytring"` // want "field EnumValidEmptytring should have the omitempty tag." "field EnumValidEmptytring should be a pointer."

	// +kubebuilder:validation:Enum=a;b;c;""
	EnumValidEmptyStringWithOmitEmpty string `json:"enumValidEmptyStringWithOmitEmpty,omitempty"` // want "field EnumValidEmptyStringWithOmitEmpty should be a pointer."

	// +kubebuilder:validation:Enum=a;b;c;""
	EnumValidEmptyStringPtr *string `json:"enumValidEmptyStringPtr"` // want "field EnumValidEmptyStringPtr should have the omitempty tag."

	// +kubebuilder:validation:Enum=a;b;c;""
	EnumValidEmptyStringPtrWithOmitEmpty *string `json:"enumValidEmptyStringPtrWithOmitEmpty,omitempty"`
}
