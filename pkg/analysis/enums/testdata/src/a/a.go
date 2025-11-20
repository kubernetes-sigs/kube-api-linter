package a

// Valid: +kubebuilder:validation:Enum without explicit values (auto-discovers constants)
// +kubebuilder:validation:Enum=Pending;Running;Succeeded;Failed
type Phase string

const (
	PhasePending   Phase = "Pending"   // Valid PascalCase
	PhaseRunning   Phase = "Running"   // Valid PascalCase
	PhaseSucceeded Phase = "Succeeded" // Valid PascalCase
	PhaseFailed    Phase = "Failed"    // Valid PascalCase
)

// Valid: +k8s:enum always auto-discovers constants
// +k8s:enum
type State string

const (
	StateActive   State = "Active"   // Valid PascalCase
	StateInactive State = "Inactive" // Valid PascalCase
)

// Valid: +kubebuilder:validation:Enum with explicit values (doesn't check constants)
// +kubebuilder:validation:Enum=Pending;helper
type ExplicitPhase string

// These constants are helpers - not validated because marker has explicit values
const (
	ExplicitPending ExplicitPhase = "Pending"
	explicit_helper ExplicitPhase = "helper" // Not flagged - explicit values mode
)

// Invalid: Type alias without enum marker (will be flagged when used in fields)
type Status string

const (
	StatusReady    Status = "Ready"
	StatusNotReady Status = "NotReady"
)

// Invalid: Type with kubebuilder:validation:Enum marker but underlying type is not string
// +kubebuilder:validation:Enum
type InvalidEnumType int // want "type InvalidEnumType has enum marker but underlying type is not string"

// Invalid enum values (not PascalCase) - using auto-discovery
// +kubebuilder:validation:Enum
type BadPhase string

const (
	phase_pending BadPhase = "pending"      // want "enum value \"pending\" should be PascalCase"
	phase_failed  BadPhase = "Phase-Failed" // want "enum value \"Phase-Failed\" should be PascalCase"
)

// Valid: Acronyms and all-caps are allowed
// +kubebuilder:validation:Enum
type Protocol string

const (
	ProtocolHTTP  Protocol = "HTTP"  // Valid: acronym
	ProtocolHTTPS Protocol = "HTTPS" // Valid: acronym
	ProtocolTCP   Protocol = "TCP"   // Valid: acronym
)

// Test struct with fields
type MySpec struct {
	// Valid: uses type alias with enum marker
	Phase Phase

	// Valid: uses type alias with enum marker
	State State

	// Valid: plain string (not required to be enum unless RequireTypeAliasForEnums is enabled)
	PlainString string

	// Invalid: type alias without enum marker
	Status Status // want "field MySpec.Status uses type Status which appears to be an enum but is missing an enum marker \\(\\+kubebuilder:validation:Enum or \\+k8s:enum\\)"

	// Valid: pointer to enum type
	PhasePtr *Phase

	// Valid: slice of enum type
	Phases []Phase

	// Valid: plain string slice (not required to be enum)
	PlainStrings []string

	// Valid: explicit enum type
	Explicit ExplicitPhase
}

// Test pointer fields
type PointerSpec struct {
	PhasePtr  *Phase
	StatusPtr *Status // want "field PointerSpec.StatusPtr uses type Status which appears to be an enum but is missing an enum marker \\(\\+kubebuilder:validation:Enum or \\+k8s:enum\\)"
}

// Test array fields
type ArraySpec struct {
	Phases   []Phase
	Statuses []Status // want "field ArraySpec.Statuses array element uses type Status which appears to be an enum but is missing an enum marker \\(\\+kubebuilder:validation:Enum or \\+k8s:enum\\)"
}

// Embedded field test
type EmbeddedSpec struct {
	MySpec
	Phase Phase
}

// Edge case: Field with enum marker directly on the field (allowed as exception)
type DirectMarkerSpec struct {
	// +kubebuilder:validation:Enum
	DirectEnum string // Valid: has enum marker directly on field
}

// Edge case: Enum values with numbers (should be valid PascalCase)
// +kubebuilder:validation:Enum=Priority1;Priority2
type Priority string

const (
	Priority1 Priority = "Priority1" // Valid: PascalCase with number
	Priority2 Priority = "Priority2" // Valid: PascalCase with number
)

// Valid: Single letter and all-caps values
// +kubebuilder:validation:Enum=A;B;API
type Level string

const (
	LevelA   Level = "A"   // Valid: single letter
	LevelB   Level = "B"   // Valid: single letter
	LevelAPI Level = "API" // Valid: acronym
)

// Edge case: Map with enum types (maps are allowed, not enforced to use enum types)
type MapSpec struct {
	// Valid: map with enum value type
	PhaseMap map[string]Phase

	// Valid: map with plain string value (maps are not enforced to use enums)
	PlainMap map[string]string
}
