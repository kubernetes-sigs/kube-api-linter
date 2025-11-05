package a

type OptionalOrRequiredTestStruct struct {
	NoMarkers string // want "field OptionalOrRequiredTestStruct.NoMarkers must be marked as kubebuilder:validation:Optional or kubebuilder:validation:Required"

	// noOptionalOrRequiredMarker is a field with no optional or required marker.
	// +enum
	// +kubebuilder:validation:Enum=Foo;Bar
	NoOptionalOrRequiredMarker string // want "field OptionalOrRequiredTestStruct.NoOptionalOrRequiredMarker must be marked as kubebuilder:validation:Optional or kubebuilder:validation:Required"

	// +optional
	// +required
	MarkedOpitonalAndRequired string // want "field OptionalOrRequiredTestStruct.MarkedOpitonalAndRequired must not be marked as both optional and required"

	// +optional
	// +kubebuilder:validation:Optional
	MarkedOptionalAndKubeBuilderOptional string // want "field OptionalOrRequiredTestStruct.MarkedOptionalAndKubeBuilderOptional should use only the marker kubebuilder:validation:Optional, optional is not required"

	// +required
	// +kubebuilder:validation:Required
	MarkedRequiredAndKubeBuilderRequired string // want "field OptionalOrRequiredTestStruct.MarkedRequiredAndKubeBuilderRequired should use only the marker kubebuilder:validation:Required, required is not required"

	// +kubebuilder:validation:Optional
	KubebuilderOptionalMarker string

	// +kubebuilder:validation:Required
	KubebuilderRequiredMarker string

	// +optional
	OptionalMarker string // want "field OptionalOrRequiredTestStruct.OptionalMarker should use marker kubebuilder:validation:Optional instead of optional"

	// +required
	RequiredMarker string // want "field OptionalOrRequiredTestStruct.RequiredMarker should use marker kubebuilder:validation:Required instead of required"

}
