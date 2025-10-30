package a

// Valid cases - arrays of structs with at least one required field

type ValidStruct struct {
	Items []ValidItem
}

type ValidItem struct {
	// +required
	Name string

	// +optional
	Description string
}

type ValidStructWithKubebuilderRequired struct {
	Items []ValidItemKubebuilder
}

type ValidItemKubebuilder struct {
	// +kubebuilder:validation:Required
	Name string

	// +optional
	Description string
}

type ValidStructWithK8sRequired struct {
	Items []ValidItemK8s
}

type ValidItemK8s struct {
	// +k8s:required
	Name string

	// +optional
	Description string
}

type ValidStructWithPointer struct {
	Items []*ValidPointerItem
}

type ValidPointerItem struct {
	// +required
	ID string
}

// Valid cases - arrays of primitives (not checked by this linter)

type ValidPrimitiveArray struct {
	Strings []string
	Ints    []int
}

// Invalid cases - arrays of structs without any required fields

type InvalidStruct struct {
	Items []InvalidItem // want "InvalidStruct.Items is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
}

type InvalidItem struct {
	// +optional
	Name string

	// +optional
	Description string
}

type InvalidStructWithPointer struct {
	Items []*InvalidPointerItem // want "InvalidStructWithPointer.Items is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
}

type InvalidPointerItem struct {
	// +optional
	ID string
}

type InvalidStructWithInlineStruct struct {
	// Inline struct definitions should also be checked
	Items []struct { // want "InvalidStructWithInlineStruct.Items is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
		// +optional
		Name string
	}
}

// Invalid case - struct with no markers at all

type InvalidStructNoMarkers struct {
	Items []InvalidItemNoMarkers // want "InvalidStructNoMarkers.Items is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
}

type InvalidItemNoMarkers struct {
	Name        string
	Description string
}

// Valid case - array of structs where all fields are required

type ValidAllRequired struct {
	Items []AllRequiredItem
}

type AllRequiredItem struct {
	// +required
	Name string

	// +required
	Description string
}

// Test with type alias

type ItemAlias = InvalidItem

type InvalidStructWithAlias struct {
	Items []ItemAlias // want "InvalidStructWithAlias.Items is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
}

// Test with array type alias

type ArrayAlias = []InvalidItem

type InvalidStructWithArrayAlias struct {
	Items ArrayAlias // want "InvalidStructWithArrayAlias.Items is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
}

// Valid case with multiple array fields

type ValidMultipleArrays struct {
	ValidItems   []ValidItem
	InvalidItems []InvalidItem // want "ValidMultipleArrays.InvalidItems is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
}

// Valid case - type alias to basic type (should not trigger linter)

type StringAlias = string

type ValidStructWithBasicTypeAlias struct {
	// This should not trigger the linter because StringAlias resolves to string, a basic type
	Items []StringAlias
}

// Valid case - custom type wrapping basic type (should not trigger linter)

type CustomString string

type ValidStructWithCustomBasicType struct {
	// This should not trigger the linter because CustomString is based on string, a basic type
	Items []CustomString
}
