package f

// This test file tests the behavior when OmitZero.Policy is Warn.
// When set to Warn, the linter emits a warning for struct fields without omitzero but does not suggest a fix.

type F struct {
	// GoodField is correctly configured with +default, +optional, and omitempty.
	// +optional
	// +default="value"
	GoodField string `json:"goodField,omitempty"`

	// MissingOmitEmpty has a default and optional but no omitempty.
	// OmitEmpty policy is still SuggestFix by default.
	// +optional
	// +default="value"
	MissingOmitEmpty string `json:"missingOmitEmpty"` // want "field F.MissingOmitEmpty has a default value but does not have omitempty in its json tag"

	// StructFieldMissingOmitZero is a struct field with default but missing omitzero.
	// When OmitZero policy is Warn, this should emit a warning without a fix.
	// +optional
	// +default={}
	StructFieldMissingOmitZero NestedStruct `json:"structFieldMissingOmitZero,omitempty"` // want "field F.StructFieldMissingOmitZero has a default value but does not have omitzero in its json tag"

	// StructFieldWithOmitZero is correctly configured with omitzero only.
	// omitempty is not required when omitzero is present (modernize linter would complain if both were present).
	// +optional
	// +default={}
	StructFieldWithOmitZero NestedStruct `json:"structFieldWithOmitZero,omitzero"`
}

type NestedStruct struct {
	// Field is a nested field.
	// +optional
	Field string `json:"field,omitempty"`
}
