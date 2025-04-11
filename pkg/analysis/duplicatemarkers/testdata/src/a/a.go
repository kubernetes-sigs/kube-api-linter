package a

// +kubebuilder:validation:Enum=foo;bar;baz
// +optional
// +kubebuilder:validation:Enum=foo;bar;baz
type Enum string // want "Enum has duplicated markers kubebuilder:validation:Enum"

// It must be ignored since it is not a type
// +required
var Variable string

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:object:root=true
type DuplicateMarkerSpec struct { // want "DuplicateMarkerSpec has duplicated markers kubebuilder:object:root"
	// +kubebuilder:validation:Required
	Foo string `json:"foo"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MaxLength=10
	DuplicatedFoo string `json:"duplicatedFoo"` // want "DuplicatedFoo has duplicated markers kubebuilder:validation:Required"
}
