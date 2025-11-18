package a

type TestPointerToSliceWithMinZeroAlways struct {
	// Pointer to slice with MinItems=0 allows distinguishing:
	// - nil (unset, use defaults)
	// - [] (explicitly empty)
	// +kubebuilder:validation:MinItems=0
	PtrArrayWithZeroMinItems *[]string `json:"ptrArrayWithZeroMinItems,omitempty"`

	// +kubebuilder:validation:MinItems=0
	PtrArrayWithZeroMinItemsNoOmitEmpty *[]string `json:"ptrArrayWithZeroMinItemsNoOmitEmpty"` // want "field TestPointerToSliceWithMinZeroAlways.PtrArrayWithZeroMinItemsNoOmitEmpty should have the omitempty tag."
}

type TestPointerToMapWithMinZeroAlways struct {
	// Pointer to map with MinProperties=0 allows distinguishing:
	// - nil (unset, use defaults)
	// - {} (explicitly empty)
	// +kubebuilder:validation:MinProperties=0
	MapPtrWithZeroMinProperties *map[string]string `json:"mapPtrWithZeroMinProperties,omitempty"`

	// +kubebuilder:validation:MinProperties=0
	MapPtrWithZeroMinPropertiesNoOmitEmpty *map[string]string `json:"mapPtrWithZeroMinPropertiesNoOmitEmpty"` // want "field TestPointerToMapWithMinZeroAlways.MapPtrWithZeroMinPropertiesNoOmitEmpty should have the omitempty tag."

	// Non-pointer version should suggest adding pointer
	// +kubebuilder:validation:MinProperties=0
	MapWithZeroMinPropertiesNoPtr map[string]string `json:"mapWithZeroMinPropertiesNoPtr,omitempty"` // want "field TestPointerToMapWithMinZeroAlways.MapWithZeroMinPropertiesNoPtr with MinItems=0/MinProperties=0, underlying type should be a pointer to distinguish nil \\(unset\\) from empty."
}

// Test that slices WITHOUT pointers are flagged when MinItems is zero
type TestSliceWithMinZeroAlways struct {
	// +kubebuilder:validation:MinItems=0
	ArrayWithZeroMinItemsNoPtr []string `json:"arrayWithZeroMinItemsNoPtr,omitempty"` // want "field TestSliceWithMinZeroAlways.ArrayWithZeroMinItemsNoPtr with MinItems=0/MinProperties=0, underlying type should be a pointer to distinguish nil \\(unset\\) from empty."
}

// Test that pointers ARE still flagged when MinItems/MinProperties is NOT zero
type TestPointerToSliceWithNonZeroMinAlways struct {
	// +kubebuilder:validation:MinItems=1
	PtrArrayWithNonZeroMinItems *[]string `json:"ptrArrayWithNonZeroMinItems,omitempty"` // want "field TestPointerToSliceWithNonZeroMinAlways.PtrArrayWithNonZeroMinItems underlying type does not need to be a pointer. The pointer should be removed."

	// No MinItems validation
	PtrArrayWithoutMinItems *[]string `json:"ptrArrayWithoutMinItems,omitempty"` // want "field TestPointerToSliceWithNonZeroMinAlways.PtrArrayWithoutMinItems underlying type does not need to be a pointer. The pointer should be removed."
}

type TestPointerToMapWithNonZeroMinAlways struct {
	// +kubebuilder:validation:MinProperties=1
	MapPtrWithNonZeroMinProperties *map[string]string `json:"mapPtrWithNonZeroMinProperties,omitempty"` // want "field TestPointerToMapWithNonZeroMinAlways.MapPtrWithNonZeroMinProperties underlying type does not need to be a pointer. The pointer should be removed."

	// No MinProperties validation
	MapPtrWithoutMinProperties *map[string]string `json:"mapPtrWithoutMinProperties,omitempty"` // want "field TestPointerToMapWithNonZeroMinAlways.MapPtrWithoutMinProperties underlying type does not need to be a pointer. The pointer should be removed."
}
