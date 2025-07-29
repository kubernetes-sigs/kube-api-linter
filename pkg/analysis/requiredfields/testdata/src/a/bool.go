package a

type TestBools struct {
	// +kubebuilder:validation:Required
	Bool bool `json:"bool"` // want "field Bool should have the omitempty tag." "field Bool has a valid zero value \\(false\\) and should be a pointer."

	// +kubebuilder:validation:Required
	BoolWithOmitEmpty bool `json:"boolWithOmitEmpty,omitempty"` // want "field BoolWithOmitEmpty has a valid zero value \\(false\\) and should be a pointer."

	// +kubebuilder:validation:Required
	BoolPtr *bool `json:"boolPtr"` // want "field BoolPtr should have the omitempty tag."

	// +kubebuilder:validation:Required
	BoolPtrWithOmitEmpty *bool `json:"boolPtrWithOmitEmpty,omitempty"`
}
