package c

type DisableBuiltInConflictsStruct struct {
	// This field has built-in conflicts but they should be ignored when DisableBuiltInConflicts is true
	// +optional
	// +required
	BuiltInConflictIgnoredField string `json:"builtInConflictIgnoredField"`

	// This field has custom conflicts that should still be detected
	// +custom:marker1
	// +custom:marker3
	CustomConflictDetectedField string `json:"customConflictDetectedField"` // want "field CustomConflictDetectedField has conflicting markers: custom_conflict: \\[custom:marker1\\] and \\[custom:marker3\\]. Custom markers conflict with each other"

	// This field has both built-in and custom conflicts, but only custom conflicts should be reported
	// +optional
	// +required
	// +custom:marker1
	// +custom:marker3
	MixedConflictsField string `json:"mixedConflictsField"` // want "field MixedConflictsField has conflicting markers: custom_conflict: \\[custom:marker1\\] and \\[custom:marker3\\]. Custom markers conflict with each other"
}
