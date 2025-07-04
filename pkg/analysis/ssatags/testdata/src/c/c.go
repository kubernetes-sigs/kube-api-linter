package c

type SSATagsSpec struct {
	// +kubebuilder:listType=atomic
	AtomicList []string `json:"atomicList,omitempty"`

	// +kubebuilder:listType=set
	SetList []string `json:"setList,omitempty"` // want "listType=set is forbidden, use listType=atomic or listType=map instead"

	// +kubebuilder:listType=map
	// +kubebuilder:listMapKey=name
	MapList []string `json:"mapList,omitempty"`

	// Non-array fields should be ignored
	SingleValue string `json:"singleValue,omitempty"`

	// MissingListType is an array field without a listType marker.
	MissingListType []string `json:"missingListType,omitempty"` // want "MissingListType should have a listType marker \\(atomic, set, or map\\)"

	// InvalidListType is an array field with an invalid listType marker.
	// +kubebuilder:listType=invalid
	InvalidListType []string `json:"invalidListType,omitempty"` // want "InvalidListType has invalid listType \"invalid\", must be one of: atomic, set, map"

	// MissingListMapKey is a map-type array field without a listMapKey marker.
	// +kubebuilder:listType=map
	MissingListMapKey []string `json:"missingListMapKey,omitempty"` // want "MissingListMapKey with listType=map must have at least one listMapKey marker"
}
