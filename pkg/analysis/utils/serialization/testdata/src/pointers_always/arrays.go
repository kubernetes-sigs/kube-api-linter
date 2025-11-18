package a

type TestArrays struct {
	Array []string `json:"array"` // want "field TestArrays.Array should have the omitempty tag."

	ArrayWithOmitEmpty []string `json:"arrayWithOmitEmpty,omitempty"`

	ArrayPtr []*string `json:"arrayPtr"` // want "field TestArrays.ArrayPtr should have the omitempty tag."

	ArrayWithOmitEmptyPtr []*string `json:"arrayWithOmitEmptyPtr,omitempty"`

	// +kubebuilder:validation:MinItems=1
	ArrayWithPositiveMinItems []string `json:"arrayWithPositiveMinItems"` // want "field TestArrays.ArrayWithPositiveMinItems should have the omitempty tag."

	// +kubebuilder:validation:MinItems=1
	ArrayWithPositiveMinItemsWithOmitEmpty []string `json:"arrayWithPositiveMinItemsWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:MinItems=0
	ArrayWithZeroMinItems []string `json:"arrayWithZeroMinItems"` // want "field TestArrays.ArrayWithZeroMinItems should have the omitempty tag." "field TestArrays.ArrayWithZeroMinItems with MinItems=0/MinProperties=0, underlying type should be a pointer to distinguish nil \\(unset\\) from empty."

	// +kubebuilder:validation:MinItems=0
	ArrayWithZeroMinItemsWithOmitEmpty []string `json:"arrayWithZeroMinItemsWithOmitEmpty,omitempty"` // want "field TestArrays.ArrayWithZeroMinItemsWithOmitEmpty with MinItems=0/MinProperties=0, underlying type should be a pointer to distinguish nil \\(unset\\) from empty."

	ByteArray []byte `json:"byteArray"` // want "field TestArrays.ByteArray should have the omitempty tag."

	ByteArrayWithOmitEmpty []byte `json:"byteArrayWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:MinLength=1
	ByteArrayWithMinLength []byte `json:"byteArrayWithMinLength"` // want "field TestArrays.ByteArrayWithMinLength should have the omitempty tag."

	// +kubebuilder:validation:MinLength=1
	ByteArrayWithMinLengthWithOmitEmpty []byte `json:"byteArrayWithMinLengthWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:MinLength=0
	ByteArrayWithMinLength0 []byte `json:"byteArrayWithMinLength0"` // want "field TestArrays.ByteArrayWithMinLength0 should have the omitempty tag."

	// +kubebuilder:validation:MinLength=0
	ByteArrayWithMinLength0WithOmitEmpty []byte `json:"byteArrayWithMinLength0WithOmitEmpty,omitempty"`

	PtrArray *[]string `json:"ptrArray"` // want "field TestArrays.PtrArray should have the omitempty tag." "field TestArrays.PtrArray underlying type does not need to be a pointer. The pointer should be removed."

	PtrArrayWithOmitEmpty *[]string `json:"ptrArrayWithOmitEmpty,omitempty"` // want "field TestArrays.PtrArrayWithOmitEmpty underlying type does not need to be a pointer. The pointer should be removed."
}
