package a

import "externaltypes"

// TestExternalTypesWithMinZeroWhenRequired tests named slice/map types from external packages
// with MinItems=0/MinProperties=0 markers using WhenRequired policy.
type TestExternalTypesWithMinZeroWhenRequired struct {
	// Pointer to external slice with MinItems=0 is valid (distinguishes nil from empty)
	// +kubebuilder:validation:MinItems=0
	PtrExternalSliceWithMinZero *externaltypes.StringSlice `json:"ptrExternalSliceWithMinZero,omitempty"`

	// Pointer to external map with MinProperties=0 is valid (distinguishes nil from empty)
	// +kubebuilder:validation:MinProperties=0
	PtrExternalMapWithMinZero *externaltypes.StringMap `json:"ptrExternalMapWithMinZero,omitempty"`

	// Non-pointer external slice with MinItems=0 is valid with WhenRequired policy
	// +kubebuilder:validation:MinItems=0
	ExternalSliceWithMinZero externaltypes.StringSlice `json:"externalSliceWithMinZero,omitempty"`

	// Non-pointer external map with MinProperties=0 is valid with WhenRequired policy
	// +kubebuilder:validation:MinProperties=0
	ExternalMapWithMinZero externaltypes.StringMap `json:"externalMapWithMinZero,omitempty"`
}
