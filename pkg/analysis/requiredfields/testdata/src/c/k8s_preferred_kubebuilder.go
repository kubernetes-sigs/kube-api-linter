package c

// TestPreferredKubebuilderRequired tests that the linter suggests +kubebuilder:validation:Required
// when preferredRequiredMarker is set to "kubebuilder:validation:Required".
type TestPreferredKubebuilderRequired struct {
	// +k8s:required
	// +k8s:minLength=1
	OnlyK8sRequired string `json:"onlyK8sRequired,omitempty"` // want "field TestPreferredKubebuilderRequired.OnlyK8sRequired has \\+k8s:required but should also have \\+kubebuilder:validation:Required marker"

	// +kubebuilder:validation:Required
	// +k8s:required
	// +k8s:minLength=1
	WithKubebuilderRequired string `json:"withKubebuilderRequired,omitempty"`

	// +required
	// +k8s:required
	// +k8s:minLength=1
	WithRequired string `json:"withRequired,omitempty"`
}
