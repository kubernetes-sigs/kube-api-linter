package b

type ZeroValueTestArrays struct {
	Array []string // want "zero value is valid" "validation is not complete"

	ArrayPtr []*string // want "zero value is valid" "validation is not complete"

	// +kubebuilder:validation:MinItems=1
	ArrayWithPositiveMinItems []string // want "zero value is not valid" "validation is complete"

	// +kubebuilder:validation:MinItems=0
	ArrayWithZeroMinItems []string // want "zero value is valid" "validation is complete"

	ByteArray []byte // want "zero value is valid" "validation is not complete"

	// +kubebuilder:validation:MinLength=1
	ByteArrayWithMinLength []byte // want "zero value is not valid" "validation is complete"

	// +kubebuilder:validation:MinLength=0
	ByteArrayWithMinLength0 []byte // want "zero value is valid" "validation is complete"
}
