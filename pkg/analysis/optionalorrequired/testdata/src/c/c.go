package c

type OptionalOrRequiredTestStruct struct {
	RequiredEnumField RequiredEnum // want "field RequiredEnumField must be marked as optional or required"

	KubebuilderRequiredEnumField KubeBuilderRequiredEnum // want "field KubebuilderRequiredEnumField must be marked as optional or required"

	OptionalEnumField OptionalEnum // want "field OptionalEnumField must be marked as optional or required"

	KubebuilderOptionalEnumField KubeBuilderOptionalEnum // want "field KubebuilderOptionalEnumField must be marked as optional or required"
}

// +required
// +kubebuilder:validation:Enum=Foo;Bar;Baz
type RequiredEnum string // want "type RequiredEnum should not be marked as required"

// +kubebuilder:validation:Required
// +kubebuilder:validation:Enum=Foo;Bar;Baz
type KubeBuilderRequiredEnum string // want "type KubeBuilderRequiredEnum should not be marked as kubebuilder:validation:Required"

// +optional
// +kubebuilder:validation:Enum=Foo;Bar;Baz
type OptionalEnum string // want "type OptionalEnum should not be marked as optional"

// +kubebuilder:validation:Optional
// +kubebuilder:validation:Enum=Foo;Bar;Baz
type KubeBuilderOptionalEnum string // want "type KubeBuilderOptionalEnum should not be marked as kubebuilder:validation:Optional"
