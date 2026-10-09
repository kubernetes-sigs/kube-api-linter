package a

// CustomPresetSpec tests that a custom preset (TestOnlyJSONTags) enables
// only the jsontags linter. Other linters like conditions, maxlength, etc.
// should NOT fire because they are not in the preset.
type CustomPresetSpec struct {
	MissingJSONTag string // want "field CustomPresetSpec.MissingJSONTag is missing json tag"

	// validField has a proper json tag and should not trigger jsontags.
	ValidField string `json:"validField"`

	// stringWithoutMaxLength would trigger maxlength if it were enabled,
	// but maxlength is NOT in the TestOnlyJSONTags preset.
	StringWithoutMaxLength string `json:"stringWithoutMaxLength"`
}
