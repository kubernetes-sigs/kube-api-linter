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

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Optional
	MarkedWithKubeBuilderOptionalTwice string // want "field MarkedWithKubeBuilderOptionalTwice should use marker optional instead of kubebuilder:validation:Optional"

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Optional
	// +optional
	MarkedWithKubeBuilderOptionalTwiceAndOptional string // want "field MarkedWithKubeBuilderOptionalTwiceAndOptional should use only the marker optional, kubebuilder:validation:Optional is not required"

	// MarkedWithK8sRequiredAndKubeBuilderRequired is a field with both k8s and kubebuilder required markers.
	// The k8s versions of the markers are currently in addition to other markers so this is accepted.
	// +k8s:required
	// +kubebuilder:validation:Required
	MarkedWithK8sRequiredAndKubeBuilderRequired string // want "field MarkedWithK8sRequiredAndKubeBuilderRequired should use marker required instead of kubebuilder:validation:Required"

	// MarkedWithK8sRequiredAndRequired is a field with both k8s and required markers.
	// The k8s versions of the markers are currently in addition to other markers so this is accepted.
	// +k8s:required
	// +required
	MarkedWithK8sRequiredAndRequired string

	// MarkedWithK8sOptionalAndKubeBuilderOptional is a field with both k8s and kubebuilder optional markers.
	// The k8s versions of the markers are currently in addition to other markers so this is accepted.
	// +k8s:optional
	// +kubebuilder:validation:Optional
	MarkedWithK8sOptionalAndKubeBuilderOptional string // want "field MarkedWithK8sOptionalAndKubeBuilderOptional should use marker optional instead of kubebuilder:validation:Optional"

	// MarkedWithK8sOptionalAndOptional is a field with both k8s and optional markers.
	// The k8s versions of the markers are currently in addition to other markers so this is accepted.
	// +k8s:optional
	// +optional
	MarkedWithK8sOptionalAndOptional string

	// MarkedWithK8sOptionalAndRequired is a field with both k8s and required markers.
	// The k8s versions of the markers are currently in addition to other markers, but they should match the semantics.
	// +k8s:optional
	// +required
	MarkedWithK8sOptionalAndRequired string // want "field MarkedWithK8sOptionalAndRequired must not be marked as both k8s:optional and required"

	// MarkedWithK8sRequiredAndOptional is a field with both k8s and optional markers.
	// The k8s versions of the markers are currently in addition to other markers, but they should match the semantics.
	// +k8s:required
	// +optional
	MarkedWithK8sRequiredAndOptional string // want "field MarkedWithK8sRequiredAndOptional must not be marked as both optional and k8s:required"

	// MarkedWithK8sRequiredAndK8sOptional is a field with both k8s and kubebuilder optional markers.
	// +k8s:required
	// +k8s:optional
	MarkedWithK8sRequiredAndK8sOptional string // want "field MarkedWithK8sRequiredAndK8sOptional must be marked as optional or required" "field MarkedWithK8sRequiredAndK8sOptional must not be marked as both k8s:optional and k8s:required"

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

type Interface interface {
	InaccessibleFunction() string
}
