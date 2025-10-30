package a

type TestMaps struct {
	// map is a map of string to string.
	// +required
	Map map[string]string `json:"map,omitempty"`

	// mapWithPositiveMinProperties is a map with MinProperties set to 1.
	// +kubebuilder:validation:MinProperties=1
	// +required
	MapWithPositiveMinProperties map[string]string `json:"mapWithPositiveMinProperties,omitempty"`

	// mapWithZeroMinProperties is a map with MinProperties set to 0.
	// +kubebuilder:validation:MinProperties=0
	// +required
	MapWithZeroMinProperties map[string]string `json:"mapWithZeroMinProperties,omitempty"`
}
