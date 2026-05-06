package a

// DVMaxLength exercises the DV marker variants of the maxlength linter.
// DV-only markers should be accepted without requiring the kubebuilder form.
type DVMaxLength struct {
	// Only DV maxLength — should NOT lint.
	// +k8s:maxLength=256
	StringWithDVMaxLength string

	// DV maxBytes for a []byte field — should NOT lint.
	// +k8s:maxBytes=512
	ByteSliceWithDVMaxBytes []byte

	// DV maxItems for an array — should NOT lint.
	// +k8s:maxItems=128
	ArrayWithDVMaxItems []int

	// DV maxProperties for a map[string]V — should NOT lint.
	// +k8s:maxProperties=64
	MapWithDVMaxProperties map[string]string

	// Both kubebuilder and DV markers — should NOT lint (either satisfies).
	// +kubebuilder:validation:MaxLength:=256
	// +k8s:maxLength=256
	StringWithBothMarkers string

	// Both kubebuilder and DV maxItems — should NOT lint.
	// +kubebuilder:validation:MaxItems:=32
	// +k8s:maxItems=32
	ArrayWithBothMaxItemsMarkers []int

	// +k8s:enum
	// DV enum marker should exempt the string from requiring a max-length.
	DVEnumString string

	// +k8s:format=date-time
	// DV format:=date-time should exempt the string from requiring a max-length.
	DVDateTimeString string

	// +k8s:format=date
	// DV format:=date should exempt the string from requiring a max-length.
	DVDateString string

	// +k8s:format=duration
	// DV format:=duration should exempt the string from requiring a max-length.
	DVDurationString string

	// No marker on string — SHOULD lint.
	StringWithNoMaxLength string // want `field DVMaxLength.StringWithNoMaxLength must have a maximum length, add kubebuilder:validation:MaxLength marker`

	// No marker on []byte — SHOULD lint.
	ByteSliceWithNoMax []byte // want `field DVMaxLength.ByteSliceWithNoMax must have a maximum length, add kubebuilder:validation:MaxLength marker`

	// No marker on array — SHOULD lint.
	ArrayWithNoMaxItems []int // want `field DVMaxLength.ArrayWithNoMaxItems must have a maximum items, add kubebuilder:validation:MaxItems marker`

	// No marker on map[string]V — SHOULD lint.
	MapWithNoMaxProperties map[string]string // want `field DVMaxLength.MapWithNoMaxProperties must have a maximum properties, add kubebuilder:validation:MaxProperties marker`

	// Non-string-keyed map — should NOT lint.
	// DV +k8s:maxProperties only supports string-keyed maps; linter mirrors this constraint.
	MapWithIntKey map[int]string

	// DV maxLength applied to []byte — SHOULD lint.
	// +k8s:maxLength (counts chars) is not the correct DV tag for []byte;
	// +k8s:maxBytes (counts bytes) should be used instead.
	// +k8s:maxLength=256
	ByteSliceWithWrongDVMarker []byte // want `field DVMaxLength.ByteSliceWithWrongDVMarker must have a maximum length, add kubebuilder:validation:MaxLength marker`
}

// StringAliasDVMaxLength is a string type with the DV maxLength marker on the alias.
// +k8s:maxLength=512
type StringAliasDVMaxLength string

// DVMaxLengthWithAliases exercises DV markers on string-alias array elements.
type DVMaxLengthWithAliases struct {
	// Array of DV-max-length string aliases — should NOT lint on the element length.
	// +kubebuilder:validation:MaxItems:=64
	AliasArrayWithMaxItems []StringAliasDVMaxLength

	// Array without MaxItems marker — SHOULD lint.
	AliasArrayWithoutMaxItems []StringAliasDVMaxLength // want `field DVMaxLengthWithAliases.AliasArrayWithoutMaxItems must have a maximum items, add kubebuilder:validation:MaxItems marker`
}
