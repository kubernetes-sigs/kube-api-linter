package a

type TestArrays struct {
	// array is an array of strings.
	// +required
	// +listType=atomic
	Array []string `json:"array,omitempty"`

	// arrayPtr is an array of string pointers.
	// +required
	// +listType=set
	ArrayPtr []*string `json:"arrayPtr,omitempty"`

	// arrayWithPositiveMinItems is an array with MinItems set to 1.
	// +kubebuilder:validation:MinItems=1
	// +optional
	// +listType=atomic
	ArrayWithPositiveMinItems []string `json:"arrayWithPositiveMinItems,omitempty"`

	// arrayWithZeroMinItems is an array with MinItems set to 0.
	// +kubebuilder:validation:MinItems=0
	// +required
	// +listType=atomic
	ArrayWithZeroMinItems []string `json:"arrayWithZeroMinItems,omitempty"`

	// byteArray is an array of bytes.
	// +optional
	ByteArray []byte `json:"byteArray,omitempty"`

	// byteArrayWithMinLength is an array of bytes with MinLength set to 1.
	// +kubebuilder:validation:MinLength=1
	// +required
	ByteArrayWithMinLength []byte `json:"byteArrayWithMinLength,omitempty"`

	// byteArrayWithMinLength0 is an array of bytes with MinLength set to 0.
	// +kubebuilder:validation:MinLength=0
	// +optional
	ByteArrayWithMinLength0 []byte `json:"byteArrayWithMinLength0,omitempty"`
}
