package a

import "externaltypes"

// TestExternalTypesWithNonZeroMin tests that pointers to external types
// are still flagged when MinItems/MinProperties is NOT zero.
type TestExternalTypesWithNonZeroMin struct {
	// Pointer to external slice with MinItems=1 should be flagged
	// +kubebuilder:validation:MinItems=1
	PtrExternalSliceWithMinOne *externaltypes.StringSlice `json:"ptrExternalSliceWithMinOne,omitempty"` // want "field TestExternalTypesWithNonZeroMin.PtrExternalSliceWithMinOne underlying type does not need to be a pointer. The pointer should be removed."

	// Pointer to external map with MinProperties=1 should be flagged
	// +kubebuilder:validation:MinProperties=1
	PtrExternalMapWithMinOne *externaltypes.StringMap `json:"ptrExternalMapWithMinOne,omitempty"` // want "field TestExternalTypesWithNonZeroMin.PtrExternalMapWithMinOne underlying type does not need to be a pointer. The pointer should be removed."

	// Pointer to external slice without MinItems should be flagged
	PtrExternalSliceNoMin *externaltypes.StringSlice `json:"ptrExternalSliceNoMin,omitempty"` // want "field TestExternalTypesWithNonZeroMin.PtrExternalSliceNoMin underlying type does not need to be a pointer. The pointer should be removed."

	// Pointer to external map without MinProperties should be flagged
	PtrExternalMapNoMin *externaltypes.StringMap `json:"ptrExternalMapNoMin,omitempty"` // want "field TestExternalTypesWithNonZeroMin.PtrExternalMapNoMin underlying type does not need to be a pointer. The pointer should be removed."
}

// TestExternalTypesNoPointer tests that non-pointer external types are valid.
type TestExternalTypesNoPointer struct {
	// Non-pointer external slice is valid
	ExternalSlice externaltypes.StringSlice `json:"externalSlice,omitempty"`

	// Non-pointer external map is valid
	ExternalMap externaltypes.StringMap `json:"externalMap,omitempty"`

	// Non-pointer external ResourceList (map type) is valid
	ExternalResourceList externaltypes.ResourceList `json:"externalResourceList,omitempty"`
}
