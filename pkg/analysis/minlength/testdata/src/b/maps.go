package b

type Maps struct {
	// +k8s:minProperties=1
	InlineMapWithMinProperties map[string]string `json:"inlineMapWithMinProperties"`

	InlineMapWithoutMinProperties map[string]string `json:"inlineMapWithoutMinProperties"` // want "field Maps.InlineMapWithoutMinProperties must have a minimum properties, add k8s:minProperties marker"

	// +k8s:minProperties=0
	InlineMapWithMinPropertiesZero map[string]string `json:"inlineMapWithMinPropertiesZero"`

	MapWithMinProperties MapWithMinProperties `json:"mapWithMinProperties"`

	MapWithoutMinProperties MapWithoutMinProperties `json:"mapWithoutMinProperties"` // want "field Maps.MapWithoutMinProperties type MapWithoutMinProperties must have a minimum properties, add k8s:minProperties marker"

	MapWithMinPropertiesZero MapWithMinPropertiesZero `json:"mapWithMinPropertiesZero"`
}

// +k8s:minProperties=0
type MapWithMinProperties map[string]string

type MapWithoutMinProperties map[string]string

// +k8s:minProperties=0
type MapWithMinPropertiesZero map[string]string
