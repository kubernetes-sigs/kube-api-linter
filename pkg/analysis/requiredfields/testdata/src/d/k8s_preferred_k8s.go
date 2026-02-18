package d

// TestPreferredK8sRequired tests that the linter does not suggest adding another marker
// when preferredRequiredMarker is set to "k8s:required" and the field already has +k8s:required.
type TestPreferredK8sRequired struct {
	// +k8s:required
	// +k8s:minLength=1
	OnlyK8sRequired string `json:"onlyK8sRequired,omitempty"`

	// +required
	// +k8s:required
	// +k8s:minLength=1
	WithRequired string `json:"withRequired,omitempty"`

	// +kubebuilder:validation:Required
	// +k8s:required
	// +k8s:minLength=1
	WithKubebuilderRequired string `json:"withKubebuilderRequired,omitempty"`
}
