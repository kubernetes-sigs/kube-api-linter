package b

type MaxLength struct {
	// +k8s:maxLength=256
	StringWithMaxLength string

	StringWithoutMaxLength string // want "field MaxLength.StringWithoutMaxLength must have a maximum length, add k8s:maxLength marker"

	// +k8s:maxItems=256
	ArrayWithMaxItems []int

	ArrayWithoutMaxItems []int // want "field MaxLength.ArrayWithoutMaxItems must have a maximum items, add k8s:maxItems marker"

	// +k8s:maxLength=512
	ByteSliceWithMaxLength []byte

	ByteSliceWithoutMaxLength []byte // want "field MaxLength.ByteSliceWithoutMaxLength must have a maximum length, add k8s:maxLength marker"
}
