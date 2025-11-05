package b

// Test with allowlist configuration

// Valid enum type with proper marker
// +kubebuilder:validation:Enum
type ExecutableName string

const (
	ExecKubectl ExecutableName = "kubectl" // Valid: in allowlist
	ExecDocker  ExecutableName = "docker"  // Valid: in allowlist
	ExecHelm    ExecutableName = "helm"    // Valid: in allowlist
	ExecUnknown ExecutableName = "unknown" // want "enum value \"unknown\" should be PascalCase"
)

// Regular enum with PascalCase (should still work)
// +kubebuilder:validation:Enum
type Status string

const (
	StatusReady    Status = "Ready"    // Valid PascalCase
	StatusNotReady Status = "NotReady" // Valid PascalCase
)

// Test struct
type MyConfig struct {
	Executable ExecutableName
	Status     Status
}
