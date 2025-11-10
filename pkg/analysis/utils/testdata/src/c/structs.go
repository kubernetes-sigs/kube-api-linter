package c

type ZeroValueTestStructs struct {
	StructWithOmittedRequiredField StructWithOmittedRequiredField `json:"structWithOmittedRequiredField,omitempty"` // want "zero value is not valid" "validation is complete"

	StructWithoutOmitZeroFieldsAndMinProperties StructWithoutOmitZeroFieldsAndMinProperties `json:"structWithoutOmitZeroFieldsAndMinProperties,omitempty"` // want "zero value is valid" "validation is complete"

	StructWithOmitZeroFieldsAndMinProperties StructWithOmitZeroFieldsAndMinProperties `json:"structWithOmitZeroFieldsAndMinProperties,omitempty"` // want "zero value is not valid" "validation is complete"

	StructWithOmitZeroAndRequiredFieldsAndMinProperties StructWithOmitZeroAndRequiredFieldsAndMinProperties `json:"structWithOmitZeroAndRequiredFieldsAndMinProperties,omitempty"` // want "zero value is valid" "validation is complete"

	// RecursiveStructs demonstrates that the zero value checker stops when it finds recursive structs

	RecursiveStruct RecursiveStructA `json:"recursiveStruct` // want "zero value is valid" "validation is complete"

	RecursiveStructWithOmitEmpty RecursiveStructA `json:"recursiveStructWithOmitEmpty,omitempty"` // want "zero value is valid" "validation is complete"

	RecursiveStructPointer *RecursiveStructA `json:"recursiveStructPointer` // want "zero value is valid" "validation is complete"

	RecursiveStructPointerWithOmitEmpty *RecursiveStructA `json:"recursiveStructPointerWithOmitEmpty,omitempty"` // want "zero value is valid" "validation is complete"

	SelfRecursiveStruct SelfRecursiveStruct `json:"selfRecursiveStruct` // want "zero value is valid" "validation is complete"

	SelfRecursiveStructWithOmitEmpty SelfRecursiveStruct `json:"selfRecursiveStructWithOmitEmpty,omitempty"` // want "zero value is valid" "validation is complete"

	SelfRecursiveStructPointer *SelfRecursiveStruct `json:"selfRecursiveStructPointer` // want "zero value is valid" "validation is complete"

	SelfRecursiveStructPointerWithOmitEmpty *SelfRecursiveStruct `json:"selfRecursiveStructPointerWithOmitEmpty,omitempty"` // want "zero value is valid" "validation is complete"
}

// StructWithOmittedRequiredField
// The zero value of the struct is `{}` which is not valid because it does not satisfy the required marker on the string field.
type StructWithOmittedRequiredField struct {
	// +required
	String string `json:"string,omitempty"` // want "zero value is valid" "validation is not complete"
}

type StructWithAllOptionalFields struct {
	// +optional
	String string `json:"string,omitempty"` // want "zero value is valid" "validation is not complete"

	// +optional
	StringPtr *string `json:"stringPtr,omitempty"` // want "zero value is valid" "validation is not complete"

	// +optional
	Int int `json:"int,omitempty"` // want "zero value is valid" "validation is not complete"

	// +optional
	IntPtr *int `json:"intPtr,omitempty"` // want "zero value is valid" "validation is not complete"
}

// StructWithoutOmitZeroFieldsAndMinProperties is a Struct having a struct field without omitzero or omitempty tag and minProperties marker.
// Because there is no omitempty or omitzero, and the zero value is valid, the zero value here is `{"structWithAllOptionalFields":{}}`.
// This means the MinProperties marker is satisfied even when the object is the zero value.
// +kubebuilder:validation:MinProperties=1
type StructWithoutOmitZeroFieldsAndMinProperties struct {
	// +optional
	StructWithAllOptionalFields StructWithAllOptionalFields `json:"structWithAllOptionalFields"` // want "zero value is valid" "validation is not complete"
}

// StructWithOmitZeroFieldsAndMinProperties is a Struct having a struct field with omitzero and minProperties marker.
// Because there is a omitzero tag, and the zero value is {}, which is not valid because it does not satisfy the MinProperties marker.
// +kubebuilder:validation:MinProperties=1
type StructWithOmitZeroFieldsAndMinProperties struct {
	// +optional
	StructWithAllOptionalFields StructWithAllOptionalFields `json:"structWithAllOptionalFields,omitzero"` // want "zero value is valid" "validation is not complete"
}

// StructWithOmitZeroAndRequiredFieldsAndMinProperties is a Struct having a struct field with omitzero and minProperties marker.
// Because there is a omitzero tag, and the zero value is {}, which is not valid because it does not satisfy the MinProperties marker.
// +kubebuilder:validation:MinProperties=1
type StructWithOmitZeroAndRequiredFieldsAndMinProperties struct {
	// +optional
	StructWithAllOptionalFields StructWithAllOptionalFields `json:"structWithAllOptionalFields,omitzero"` // want "zero value is valid" "validation is not complete"

	// +required
	RequiredStructWithAllOptionalFields StructWithAllOptionalFields `json:"requiredStructWithAllOptionalFields"` // want "zero value is valid" "validation is not complete"
}

// RecursiveStructA is paired with RecursiveStructB to demonstrate a recursive relationship for zero value checking.
type RecursiveStructA struct {
	// +optional
	String string `json:"string,omitempty"` // want "zero value is valid" "validation is not complete"

	RecursiveStructB *RecursiveStructB `json:"recursiveStructB,omitempty` // want "zero value is valid" "validation is complete"
}

// RecursiveStructB is paired with RecursiveStructA to demonstrate a recursive relationship for zero value checking.
type RecursiveStructB struct {
	// +optional
	String string `json:"string,omitempty"` // want "zero value is valid" "validation is not complete"

	RecursiveStructA *RecursiveStructA `json:"recursiveStructB` // want "zero value is valid" "validation is complete"

	RecursiveStructAWithOmitEmpty *RecursiveStructA `json:"recursiveStructAWithOmitEmpty,omitempty"` // want "zero value is valid" "validation is complete"
}

type SelfRecursiveStruct struct {
	// +optional
	String string `json:"string,omitempty"` // want "zero value is valid" "validation is not complete"

	SelfRecursiveStruct *SelfRecursiveStruct `json:"selfRecursiveStruct` // want "zero value is valid" "validation is complete"

	SelfRecursiveStructWithOmitEmpty *SelfRecursiveStruct `json:"selfRecursiveStructWithOmitEmpty,omitempty"` // want "zero value is valid" "validation is complete"
}
