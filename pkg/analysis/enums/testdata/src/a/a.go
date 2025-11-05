package a

// Valid enum type with proper marker
// +kubebuilder:validation:Enum
type Phase string

const (
	PhasePending   Phase = "Pending"   // Valid PascalCase
	PhaseRunning   Phase = "Running"   // Valid PascalCase
	PhaseSucceeded Phase = "Succeeded" // Valid PascalCase
	PhaseFailed    Phase = "Failed"    // Valid PascalCase
)

// Alternative marker format
// +k8s:enum
type State string

const (
	StateActive   State = "Active"   // Valid PascalCase
	StateInactive State = "Inactive" // Valid PascalCase
)

// Invalid: Missing +enum marker
// This type doesn't have +enum marker, so it will be flagged when used in fields
type Status string

const (
	StatusReady    Status = "Ready"
	StatusNotReady Status = "NotReady"
)

// Invalid: Type with +enum but not string
// +kubebuilder:validation:Enum
type InvalidEnumType int // want "type InvalidEnumType has \\+enum marker but underlying type is not string"

// Invalid enum values (not PascalCase)
// +kubebuilder:validation:Enum
type BadPhase string

const (
	phase_pending   BadPhase = "pending"      // want "enum value \"pending\" should be PascalCase"
	PHASE_RUNNING   BadPhase = "RUNNING"      // want "enum value \"RUNNING\" should be PascalCase"
	phase_succeeded BadPhase = "succeeded"    // want "enum value \"succeeded\" should be PascalCase"
	Phase_Failed    BadPhase = "Phase-Failed" // want "enum value \"Phase-Failed\" should be PascalCase"
)

// Test struct with fields
type MySpec struct {
	// Valid: uses type alias with +enum
	Phase Phase

	// Valid: uses type alias with +enum
	State State

	// Invalid: plain string without +enum
	PlainString string // want "field PlainString uses plain string without \\+enum marker"

	// Invalid: type alias without +enum marker
	Status Status // want "field Status uses type Status which appears to be an enum but is missing \\+enum marker"

	// Valid: pointer to enum type
	PhasePtr *Phase

	// Valid: slice of enum type
	Phases []Phase

	// Invalid: plain string slice
	PlainStrings []string // want "field PlainStrings array element uses plain string without \\+enum marker"
}

// Test pointer fields
type PointerSpec struct {
	PhasePtr    *Phase
	StatusPtr   *Status // want "field StatusPtr uses type Status which appears to be an enum but is missing \\+enum marker"
	PlainStrPtr *string // want "field PlainStrPtr uses plain string without \\+enum marker"
}

// Test array fields
type ArraySpec struct {
	Phases        []Phase
	Statuses      []Status // want "field Statuses array element uses type Status which appears to be an enum but is missing \\+enum marker"
	PlainStrArray []string // want "field PlainStrArray array element uses plain string without \\+enum marker"
}

// Embedded field test
type EmbeddedSpec struct {
	MySpec
	Phase Phase
}

// Edge case: Field with +enum marker directly on the field (should be allowed as exception)
type DirectMarkerSpec struct {
	// +kubebuilder:validation:Enum
	DirectEnum string // Valid: has +enum marker directly on field
}

// Edge case: Enum values with numbers (should be valid PascalCase)
// +kubebuilder:validation:Enum
type Priority string

const (
	Priority1 Priority = "Priority1" // Valid: PascalCase with number
	Priority2 Priority = "Priority2" // Valid: PascalCase with number
)

// Edge case: Single letter uppercase (edge case for all-caps check)
// +kubebuilder:validation:Enum
type Level string

const (
	LevelA Level = "A" // Valid: single uppercase letter
	LevelB Level = "B" // Valid: single uppercase letter
)

// Edge case: Map with enum types (maps are allowed, not enforced to use enum types)
type MapSpec struct {
	// Valid: map with enum value type
	PhaseMap map[string]Phase

	// Valid: map with plain string value (maps are not enforced to use enums)
	PlainMap map[string]string
}
