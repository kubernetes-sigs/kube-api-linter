package a

type Maps struct {
	// +kubebuilder:validation:MinProperties:=1
	InlineMapWithMinProperties map[string]string `json:"inlineMapWithMinProperties"`

	InlineMapWithoutMinProperties map[string]string `json:"inlineMapWithoutMinProperties"` // want "field InlineMapWithoutMinProperties must have a minimum properties, add kubebuilder:validation:MinProperties marker"

	// +kubebuilder:validation:MinProperties:=0
	InlineMapWithMinPropertiesZero map[string]string `json:"inlineMapWithMinPropertiesZero"`

	MapWithMinProperties MapWithMinProperties `json:"mapWithMinProperties"`

	MapWithoutMinProperties MapWithoutMinProperties `json:"mapWithoutMinProperties"` // want "field MapWithoutMinProperties type MapWithoutMinProperties must have a minimum properties, add kubebuilder:validation:MinProperties marker"

	MapWithMinPropertiesZero MapWithMinPropertiesZero `json:"mapWithMinPropertiesZero"`
}

// +kubebuilder:validation:MinProperties:=1
type MapWithMinProperties map[string]string

type MapWithoutMinProperties map[string]string

// +kubebuilder:validation:MinProperties:=0
type MapWithMinPropertiesZero map[string]string
