package a

type TestMaps struct {
	Map map[string]string `json:"map"` // want "field Map should have the omitempty tag."

	MapWithOmitEmpty map[string]string `json:"mapWithOmitEmpty,omitempty"`

	MapPtr *map[string]string `json:"mapPtr"` // want "field MapPtr should have the omitempty tag." "field MapPtr underlying type does not need to be a pointer. The pointer should be removed."

	MapPtrWithOmitEmpty *map[string]string `json:"mapPtrWithOmitEmpty,omitempty"` // want "field MapPtrWithOmitEmpty underlying type does not need to be a pointer. The pointer should be removed."

	// +kubebuilder:validation:MinProperties=1
	MapWithPositiveMinProperties map[string]string `json:"mapWithPositiveMinProperties"` // want "field MapWithPositiveMinProperties should have the omitempty tag."

	// +kubebuilder:validation:MinProperties=1
	MapWithPositiveMinPropertiesWithOmitEmpty map[string]string `json:"mapWithPositiveMinPropertiesWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:MinProperties=0
	MapWithZeroMinProperties map[string]string `json:"mapWithZeroMinProperties"` // want "field MapWithZeroMinProperties should have the omitempty tag." "field MapWithZeroMinProperties with MinItems=0/MinProperties=0, underlying type should be a pointer to distinguish nil \\(unset\\) from empty."

	// +kubebuilder:validation:MinProperties=0
	MapWithZeroMinPropertiesWithOmitEmpty map[string]string `json:"mapWithZeroMinPropertiesWithOmitEmpty,omitempty"` // want "field MapWithZeroMinPropertiesWithOmitEmpty with MinItems=0/MinProperties=0, underlying type should be a pointer to distinguish nil \\(unset\\) from empty."
}
