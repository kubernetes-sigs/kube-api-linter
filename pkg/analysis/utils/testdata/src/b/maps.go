package b

type ZeroValueTestMaps struct {
	Map map[string]string // want "zero value is valid" "validation is not complete"

	MapPtr *map[string]string // want "zero value is valid" "validation is not complete"

	// +kubebuilder:validation:MinProperties=1
	MapWithPositiveMinProperties map[string]string // want "zero value is not valid" "validation is complete"

	// +kubebuilder:validation:MinProperties=0
	MapWithZeroMinProperties map[string]string // want "zero value is valid" "validation is complete"
}
