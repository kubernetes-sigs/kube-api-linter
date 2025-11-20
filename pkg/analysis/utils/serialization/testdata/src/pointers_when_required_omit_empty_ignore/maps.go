package a

type TestMaps struct {
	Map map[string]string `json:"map"`

	MapWithOmitEmpty map[string]string `json:"mapWithOmitEmpty,omitempty"`

	MapPtr *map[string]string `json:"mapPtr"` // want "field TestMaps.MapPtr does not have omitempty and allows the zero value. The field does not need to be a pointer."

	MapPtrWithOmitEmpty *map[string]string `json:"mapPtrWithOmitEmpty,omitempty"` // want "field TestMaps.MapPtrWithOmitEmpty underlying type does not need to be a pointer. The pointer should be removed."

	// +kubebuilder:validation:MinProperties=1
	MapWithPositiveMinProperties map[string]string `json:"mapWithPositiveMinProperties"` // want "field TestMaps.MapWithPositiveMinProperties does not allow the zero value. It must have the omitempty tag."

	// +kubebuilder:validation:MinProperties=1
	MapWithPositiveMinPropertiesWithOmitEmpty map[string]string `json:"mapWithPositiveMinPropertiesWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:MinProperties=0
	MapWithZeroMinProperties map[string]string `json:"mapWithZeroMinProperties"`

	// +kubebuilder:validation:MinProperties=0
	MapWithZeroMinPropertiesWithOmitEmpty map[string]string `json:"mapWithZeroMinPropertiesWithOmitEmpty,omitempty"`
}
