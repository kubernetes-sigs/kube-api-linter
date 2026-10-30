package a

// CRDSpec tests that the CustomResource preset enables CRD-specific linters
// (maxlength) and does not enable BuiltIn-only linters (nonpointerstructs).
type CRDSpec struct {
	// name is a string field without maxLength marker.
	// The maxlength linter (in CustomResource) should flag this.
	// +required
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name,omitempty"` // want "field CRDSpec.Name must have a maximum length, add kubebuilder:validation:MaxLength marker"

	// config is a non-pointer struct.
	// The nonpointerstructs linter (in BuiltIn, NOT in CustomResource) should NOT flag this.
	// +required
	Config CRDConfig `json:"config,omitempty,omitzero"`
}

type CRDConfig struct {
	// replicas is an integer.
	// +required
	// +kubebuilder:validation:Minimum=1
	Replicas int32 `json:"replicas,omitempty"`
}
