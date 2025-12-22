package e

// This test file tests the behavior when OmitEmpty.Policy is Ignore.
// When set to Ignore, the linter does not check for omitempty at all.

type E struct {
	// GoodField is correctly configured with +default, +optional, and omitempty.
	// +optional
	// +default="value"
	GoodField string `json:"goodField,omitempty"`

	// MissingOmitEmpty has a default and optional but no omitempty.
	// When policy is Ignore, this should NOT emit any warning.
	// +optional
	// +default="value"
	MissingOmitEmpty string `json:"missingOmitEmpty"`

	// StructFieldMissingBoth is a struct field missing both omitempty and omitzero.
	// When OmitEmpty policy is Ignore, only omitzero warning should be emitted.
	// +optional
	// +default={}
	StructFieldMissingBoth NestedStruct `json:"structFieldMissingBoth"` // want "field E.StructFieldMissingBoth has a default value but does not have omitzero in its json tag"

	// StructFieldWithOmitEmpty has omitempty but missing omitzero.
	// +optional
	// +default={}
	StructFieldWithOmitEmpty NestedStruct `json:"structFieldWithOmitEmpty,omitempty"` // want "field E.StructFieldWithOmitEmpty has a default value but does not have omitzero in its json tag"
}

type NestedStruct struct {
	// Field is a nested field.
	// +optional
	Field string `json:"field,omitempty"`
}
