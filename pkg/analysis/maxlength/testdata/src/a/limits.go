package a

type Limits struct {
	// +k8s:maxLength=256
	StringWithK8sMaxLength string

	// +k8s:maxBytes=512
	StringWithK8sMaxBytes string

	// +k8s:maxItems=128
	ArrayWithK8sMaxItems []int

	// +k8s:maxProperties=64
	MapWithK8sMaxProperties map[string]string

	// +k8s:maximum=100
	IntWithK8sMaximum int32

	// +kubebuilder:validation:MaxProperties=32
	MapWithKubebuilderMaxProperties map[string]int

	// +kubebuilder:validation:Maximum=50
	IntWithKubebuilderMaximum int64

	// Missing limits
	StringWithoutLimit string // want "field Limits.StringWithoutLimit must have a maximum length, add kubebuilder:validation:MaxLength marker"
	ArrayWithoutLimit  []int  // want "field Limits.ArrayWithoutLimit must have a maximum items, add kubebuilder:validation:MaxItems marker"
	MapWithoutLimit    map[string]string // want "field Limits.MapWithoutLimit must have a maximum number of properties, add kubebuilder:validation:MaxProperties marker"
	IntWithoutLimit    int32 // want "field Limits.IntWithoutLimit must have a maximum value, add kubebuilder:validation:Maximum marker"
}
