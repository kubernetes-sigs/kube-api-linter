package b

type CustomConflictStruct struct {
	// Valid field with only custom marker1
	// +custom:marker1
	ValidCustomMarker1Field string `json:"validCustomMarker1Field"`

	// Valid field with custom marker1 that has a value
	// +custom:marker1:=someValue
	ValidCustomMarker1WithValueField string `json:"validCustomMarker1WithValueField"`

	// Conflict: custom marker1 vs custom marker3 (both with values)
	// +custom:marker1:=value1
	// +custom:marker3:=value2
	CustomConflictWithValuesField string `json:"customConflictWithValuesField"` // want "field CustomConflictWithValuesField has conflicting markers: custom_conflict: \\[custom:marker1\\] and \\[custom:marker3\\]. Custom markers conflict with each other"

	// Conflict: custom marker1 vs custom marker3 (mixed with and without values)
	// +custom:marker1
	// +custom:marker3:=someValue
	CustomConflictMixedField string `json:"customConflictMixedField"` // want "field CustomConflictMixedField has conflicting markers: custom_conflict: \\[custom:marker1\\] and \\[custom:marker3\\]. Custom markers conflict with each other"

	// Multiple conflicts with multiple markers in each set:
	// - set A: +custom:marker1, +custom:marker2
	// - set B: +custom:marker3, +custom:marker4
	// +custom:marker1
	// +custom:marker2
	// +custom:marker3
	// +custom:marker4
	MultipleCustomConflictsField string `json:"multipleCustomConflictsField"` // want "field MultipleCustomConflictsField has conflicting markers: custom_conflict: \\[custom:marker1 custom:marker2\\] and \\[custom:marker3 custom:marker4\\]. Custom markers conflict with each other"

	// Test that custom markers don't disable built-in marker conflict detection
	// This field has both custom marker conflicts AND built-in marker conflicts
	// +custom:marker1
	// +custom:marker3
	// +optional
	// +required
	CustomAndBuiltinConflictsField string `json:"customAndBuiltinConflictsField"` // want "field CustomAndBuiltinConflictsField has conflicting markers: custom_conflict: \\[custom:marker1\\] and \\[custom:marker3\\]. Custom markers conflict with each other" "field CustomAndBuiltinConflictsField has conflicting markers: optional_vs_required: \\[optional\\] and \\[required\\]. A field cannot be both optional and required"
}
