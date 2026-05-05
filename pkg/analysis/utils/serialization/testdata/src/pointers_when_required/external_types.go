package a

import "externaltypes"

// TestExternalTypesWithNonZeroMin tests that pointers to external types
// are still flagged when MinItems/MinProperties is NOT zero.
type TestExternalTypesWithNonZeroMin struct {
	// Pointer to external slice with MinItems=1 should be flagged
	// +kubebuilder:validation:MinItems=1
	PtrExternalSliceWithMinOne *externaltypes.StringSlice `json:"ptrExternalSliceWithMinOne,omitempty"` // want "field TestExternalTypesWithNonZeroMin.PtrExternalSliceWithMinOne does not allow the zero value. The field does not need to be a pointer."

	// Pointer to external map with MinProperties=1 should be flagged
	// +kubebuilder:validation:MinProperties=1
	PtrExternalMapWithMinOne *externaltypes.StringMap `json:"ptrExternalMapWithMinOne,omitempty"` // want "field TestExternalTypesWithNonZeroMin.PtrExternalMapWithMinOne does not allow the zero value. The field does not need to be a pointer."

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

// TestExternalBasicTypesWhenRequired tests external basic type aliases
// with the WhenRequired pointer policy.
type TestExternalBasicTypesWhenRequired struct {
	// External string with MinLength=1 - zero value not valid, pointer not needed
	// +kubebuilder:validation:MinLength=1
	// +optional
	ExternalStringWithMinLength1 externaltypes.StringAlias `json:"externalStringWithMinLength1,omitempty"`

	// External string without MinLength - zero value valid but validation incomplete, needs pointer
	// +optional
	ExternalStringNoMinLength externaltypes.StringAlias `json:"externalStringNoMinLength,omitempty"` // want "field TestExternalBasicTypesWhenRequired.ExternalStringNoMinLength has a valid zero value \\(\"\"\\), but the validation is not complete \\(e.g. minimum length\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// External int with Minimum=1 - zero value not valid, pointer not needed
	// +kubebuilder:validation:Minimum=1
	// +optional
	ExternalIntWithMinimum1 externaltypes.IntAlias `json:"externalIntWithMinimum1,omitempty"`

	// External int without Minimum - zero value valid but validation incomplete, needs pointer
	// +optional
	ExternalIntNoMinimum externaltypes.IntAlias `json:"externalIntNoMinimum,omitempty"` // want "field TestExternalBasicTypesWhenRequired.ExternalIntNoMinimum has a valid zero value \\(0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// External float with Minimum=1.0 - zero value not valid, pointer not needed
	// +kubebuilder:validation:Minimum=1.0
	// +optional
	ExternalFloatWithMinimum1 externaltypes.FloatAlias `json:"externalFloatWithMinimum1,omitempty"`

	// External float without Minimum - zero value valid but validation incomplete, needs pointer
	// +optional
	ExternalFloatNoMinimum externaltypes.FloatAlias `json:"externalFloatNoMinimum,omitempty"` // want "field TestExternalBasicTypesWhenRequired.ExternalFloatNoMinimum has a valid zero value \\(0.0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// External bool - always has valid zero value (false), needs pointer
	// +optional
	ExternalBool externaltypes.BoolAlias `json:"externalBool,omitempty"` // want "field TestExternalBasicTypesWhenRequired.ExternalBool has a valid zero value \\(false\\) and should be a pointer."

	// Pointer to external string with MinLength=1 - doesn't need pointer since zero value not valid
	// +kubebuilder:validation:MinLength=1
	// +optional
	PtrExternalStringWithMinLength1 *externaltypes.StringAlias `json:"ptrExternalStringWithMinLength1,omitempty"` // want "field TestExternalBasicTypesWhenRequired.PtrExternalStringWithMinLength1 does not allow the zero value. The field does not need to be a pointer."

	// Pointer to external int with Minimum=1 - doesn't need pointer since zero value not valid
	// +kubebuilder:validation:Minimum=1
	// +optional
	PtrExternalIntWithMinimum1 *externaltypes.IntAlias `json:"ptrExternalIntWithMinimum1,omitempty"` // want "field TestExternalBasicTypesWhenRequired.PtrExternalIntWithMinimum1 does not allow the zero value. The field does not need to be a pointer."

	// Pointer to external bool - valid, bool zero value (false) is a legitimate value
	// +optional
	PtrExternalBool *externaltypes.BoolAlias `json:"ptrExternalBool,omitempty"`
}

// TestExternalStructWhenRequired tests external struct types with the WhenRequired pointer policy.
type TestExternalStructWhenRequired struct {
	// External struct with MinProperties=1 - zero value not valid, pointer not needed
	// +kubebuilder:validation:MinProperties=1
	// +optional
	ExternalStructWithMinProperties1 externaltypes.StructType `json:"externalStructWithMinProperties1,omitempty"`

	// External struct (all omitempty fields) without MinProperties - zero value valid but incomplete
	// +optional
	ExternalStructNoMinProperties externaltypes.StructType `json:"externalStructNoMinProperties,omitempty"` // want "field TestExternalStructWhenRequired.ExternalStructNoMinProperties has a valid zero value \\(\\{\\}\\), but the validation is not complete \\(e.g. min properties/adding required fields\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// External struct with non-omitted fields - zero value valid and validation complete, needs pointer
	// +optional
	ExternalStructWithNonOmittedField externaltypes.StructTypeWithNonOmittedField `json:"externalStructWithNonOmittedField,omitempty"` // want "field TestExternalStructWhenRequired.ExternalStructWithNonOmittedField has a valid zero value \\(\\{\"name\": \"\"\\}\\) and should be a pointer."

	// Pointer to external struct with MinProperties=1 - doesn't need pointer
	// +kubebuilder:validation:MinProperties=1
	// +optional
	PtrExternalStructWithMinProperties1 *externaltypes.StructType `json:"ptrExternalStructWithMinProperties1,omitempty"` // want "field TestExternalStructWhenRequired.PtrExternalStructWithMinProperties1 does not allow the zero value. The field does not need to be a pointer."

	// Pointer to external struct without MinProperties - valid, can't verify so pointer is safe
	// +optional
	PtrExternalStructNoMinProperties *externaltypes.StructType `json:"ptrExternalStructNoMinProperties,omitempty"`
}
