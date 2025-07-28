package a

type TestStruct struct {
	// Valid field with only marker1
	// +marker1
	ValidMarker1Field string `json:"validMarker1Field"`

	// Valid field with marker1 that has a value
	// +marker1:=someValue
	ValidMarker1WithValueField string `json:"validMarker1WithValueField"`

	// Valid field with marker1 and marker2 (both in same conflict set)
	// +marker1
	// +marker2
	ValidMarker1And2Field string `json:"validMarker1And2Field"`

	// Conflict: marker1 vs marker3 (both with values)
	// +marker1:=value1
	// +marker3:=value2
	ConflictWithValuesField string `json:"conflictWithValuesField"` // want "field ConflictWithValuesField has conflicting markers: test_conflict: \\{\\[marker1\\], \\[marker3\\]\\}. Test markers conflict with each other"

	// Conflict: marker1 vs marker3 (mixed with and without values)
	// +marker1
	// +marker3:=someValue
	ConflictMixedField string `json:"conflictMixedField"` // want "field ConflictMixedField has conflicting markers: test_conflict: \\{\\[marker1\\], \\[marker3\\]\\}. Test markers conflict with each other"

	// Multiple conflicts with multiple markers in each set:
	// +marker1
	// +marker2
	// +marker3
	// +marker4
	MultipleConflictsField string `json:"multipleConflictsField"` // want "field MultipleConflictsField has conflicting markers: test_conflict: \\{\\[marker1 marker2\\], \\[marker3 marker4\\]\\}. Test markers conflict with each other"

	// Three-way conflict: marker5 vs marker7 vs marker9
	// +marker5
	// +marker7
	// +marker9
	ThreeWayConflictField string `json:"threeWayConflictField"` // want "field ThreeWayConflictField has conflicting markers: three_way_conflict: \\{\\[marker5\\], \\[marker7\\], \\[marker9\\]\\}. Three-way conflict between marker sets"

	// Three-way conflict with values
	// +marker6:=value1
	// +marker8:=value2
	// +marker10:=value3
	ThreeWayConflictWithValuesField string `json:"threeWayConflictWithValuesField"` // want "field ThreeWayConflictWithValuesField has conflicting markers: three_way_conflict: \\{\\[marker6\\], \\[marker8\\], \\[marker10\\]\\}. Three-way conflict between marker sets"

	// Valid field with markers from same set in three-way conflict
	// +marker5
	// +marker6
	ValidThreeWaySameSetField string `json:"validThreeWaySameSetField"`

	// Three-way conflict with multiple markers from each set
	// +marker5
	// +marker6
	// +marker7
	// +marker8
	// +marker9
	// +marker10
	ThreeWayMultipleMarkersField string `json:"threeWayMultipleMarkersField"` // want "field ThreeWayMultipleMarkersField has conflicting markers: three_way_conflict: \\{\\[marker5 marker6\\], \\[marker7 marker8\\], \\[marker10 marker9\\]\\}. Three-way conflict between marker sets"

	// Three-way conflict with only subset of sets triggered (sets 1 and 2 only)
	// +marker5
	// +marker7
	SubsetThreeWayConflictField string `json:"subsetThreeWayConflictField"` // want "field SubsetThreeWayConflictField has conflicting markers: three_way_conflict: \\{\\[marker5\\], \\[marker7\\]\\}. Three-way conflict between marker sets"
}
