package b

type ZeroValueTestStrings struct {
	String string // want "zero value is valid" "validation is not complete"

	StringPtr *string // want "zero value is valid" "validation is not complete"

	// +kubebuilder:validation:MinLength=1
	StringWithMinLength string // want "zero value is not valid" "validation is complete"

	// +kubebuilder:validation:MinLength=1
	StringPtrWithMinLength *string // want "zero value is not valid" "validation is complete"

	// +kubebuilder:validation:MinLength=0
	StringWithMinLength0 string // want "zero value is valid" "validation is complete"

	// +kubebuilder:validation:Enum=a;b;c
	EnumString string // want "zero value is not valid" "validation is complete"

	// +kubebuilder:validation:Enum=a;b;c
	EnumStringPtr *string // want "zero value is not valid" "validation is complete"

	// +kubebuilder:validation:Enum=a;b;c;""
	EnumValidEmptytring string // want "zero value is valid" "validation is complete"

	// +kubebuilder:validation:Enum=a;b;c;""
	EnumValidEmptyStringPtr *string // want "zero value is valid" "validation is complete"
}
