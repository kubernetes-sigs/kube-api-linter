package a

type TestNumbers struct {
	// int32 is an Int32.
	// +kubebuilder:validation:Minimum=1
	// +required
	Int32 int32 `json:"int32,omitempty"`

	// int64 is an Int64.
	// +kubebuilder:validation:Minimum=1
	// +required
	Int64 int64 `json:"int64,omitempty"`
}
