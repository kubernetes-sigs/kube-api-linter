package a

type TestMaps struct {
	// +required
	Map map[string]string `json:"map"` // want "field Map should have the omitempty tag."

	// +required
	MapWithOmitEmpty map[string]string `json:"mapWithOmitEmpty,omitempty"`

	// +required
	MapPtr *map[string]string `json:"mapPtr"` // want "field MapPtr should have the omitempty tag." "field MapPtr underlying type does not need to be a pointer. The pointer should be removed."

	// +required
	MapPtrWithOmitEmpty *map[string]string `json:"mapPtrWithOmitEmpty,omitempty"` // want "field MapPtrWithOmitEmpty underlying type does not need to be a pointer. The pointer should be removed."

	// +kubebuilder:validation:MinProperties=1
	// +required
	MapWithPositiveMinProperties map[string]string `json:"mapWithPositiveMinProperties"` // want "field MapWithPositiveMinProperties should have the omitempty tag."

	// +kubebuilder:validation:MinProperties=1
	// +required
	MapWithPositiveMinPropertiesWithOmitEmpty map[string]string `json:"mapWithPositiveMinPropertiesWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:MinProperties=0
	// +required
	MapWithZeroMinProperties map[string]string `json:"mapWithZeroMinProperties"` // want "field MapWithZeroMinProperties should have the omitempty tag."

	// +kubebuilder:validation:MinProperties=0
	// +required
	MapWithZeroMinPropertiesWithOmitEmpty map[string]string `json:"mapWithZeroMinPropertiesWithOmitEmpty,omitempty"`
}
