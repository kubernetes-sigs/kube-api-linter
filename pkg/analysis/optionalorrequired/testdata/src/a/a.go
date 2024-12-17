package a

type OptionalOrRequiredTestStruct struct {
	NoMarkers string // want "field NoMarkers must be marked as optional or required"

	// noOptionalOrRequiredMarker is a field with no optional or required marker.
	// +enum
	// +kubebuilder:validation:Enum=Foo;Bar
	NoOptionalOrRequiredMarker string // want "field NoOptionalOrRequiredMarker must be marked as optional or required"

	// +optional
	// +required
	MarkedOpitonalAndRequired string // want "field MarkedOpitonalAndRequired must not be marked as both optional and required"

	// +optional
	// +kubebuilder:validation:Optional
	MarkedOptionalAndKubeBuilderOptional string // want "field MarkedOptionalAndKubeBuilderOptional should use only the marker optional, kubebuilder:validation:Optional is not required"

	// +required
	// +kubebuilder:validation:Required
	MarkedRequiredAndKubeBuilderRequired string // want "field MarkedRequiredAndKubeBuilderRequired should use only the marker required, kubebuilder:validation:Required is not required"

	// +kubebuilder:validation:Optional
	KubebuilderOptionalMarker string // want "field KubebuilderOptionalMarker should use marker optional instead of kubebuilder:validation:Optional"

	// +kubebuilder:validation:Required
	KubebuilderRequiredMarker string // want "field KubebuilderRequiredMarker should use marker required instead of kubebuilder:validation:Required"

	// +optional
	OptionalMarker string

	// +required
	RequiredMarker string

	A `json:",inline"`

	B `json:"b,omitempty"` // want "embedded field B must be marked as optional or required"

	// +optional
	C `json:"c,omitempty"`
}

type A struct{}

// DoNothing is used to check that the analyser doesn't report on methods.
func (A) DoNothing() {}

type B struct{}

type C struct{}
