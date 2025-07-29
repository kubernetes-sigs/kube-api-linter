package a

type TestArrays struct {
	// +required
	Array []string `json:"array"` // want "field Array should have the omitempty tag."

	// +required
	ArrayWithOmitEmpty []string `json:"arrayWithOmitEmpty,omitempty"`

	// +required
	ArrayPtr []*string `json:"arrayPtr"` // want "field ArrayPtr should have the omitempty tag."

	// +required
	ArrayWithOmitEmptyPtr []*string `json:"arrayWithOmitEmptyPtr,omitempty"`

	// This is not picked up as the field is marked as optional.
	// +kubebuilder:validation:MinItems=1
	// +required
	ArrayWithPositiveMinItems []string `json:"arrayWithPositiveMinItems"` // want "field ArrayWithPositiveMinItems should have the omitempty tag."

	// This is not picked up as the field is marked as optional.
	// +kubebuilder:validation:MinItems=1
	// +optional
	OptionalArrayWithPositiveMinItems []string `json:"optionalArrayWithPositiveMinItems"`

	// +kubebuilder:validation:MinItems=1
	// +required
	ArrayWithPositiveMinItemsWithOmitEmpty []string `json:"arrayWithPositiveMinItemsWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:MinItems=0
	// +required
	ArrayWithZeroMinItems []string `json:"arrayWithZeroMinItems"` // want "field ArrayWithZeroMinItems should have the omitempty tag."

	// +kubebuilder:validation:MinItems=0
	// +optional
	ArrayWithZeroMinItemsWithOmitEmpty []string `json:"arrayWithZeroMinItemsWithOmitEmpty,omitempty"`

	// +optional
	ByteArray []byte `json:"byteArray"`

	ByteArrayWithOmitEmpty []byte `json:"byteArrayWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:MinLength=1
	// +required
	ByteArrayWithMinLength []byte `json:"byteArrayWithMinLength"` // want "field ByteArrayWithMinLength should have the omitempty tag."

	// +kubebuilder:validation:MinLength=1
	// +required
	ByteArrayWithMinLengthWithOmitEmpty []byte `json:"byteArrayWithMinLengthWithOmitEmpty,omitempty"`

	// neither required or empty so no issue reported
	// +kubebuilder:validation:MinLength=0
	ByteArrayWithMinLength0 []byte `json:"byteArrayWithMinLength0"`

	// +kubebuilder:validation:MinLength=0
	ByteArrayWithMinLength0WithOmitEmpty []byte `json:"byteArrayWithMinLength0WithOmitEmpty,omitempty"`

	// +required
	PtrArray *[]string `json:"ptrArray"` // want "field PtrArray should have the omitempty tag." "field PtrArray underlying type does not need to be a pointer. The pointer should be removed."

	// +required
	PtrArrayWithOmitEmpty *[]string `json:"ptrArrayWithOmitEmpty,omitempty"` // want "field PtrArrayWithOmitEmpty underlying type does not need to be a pointer. The pointer should be removed."
}
