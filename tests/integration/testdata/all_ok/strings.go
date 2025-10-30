package a

type TestStrings struct {
	// string is a string.
	// +required
	// +kubebuilder:validation:MinLength=1
	String string `json:"string,omitempty"`

	// enumString is a string with an enum validation.
	// +required
	// +kubebuilder:validation:Enum=a;b;c
	EnumString string `json:"enumString,omitempty"`
}
