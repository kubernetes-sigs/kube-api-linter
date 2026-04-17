package b

// Test with allowlist configuration

// Valid enum type with proper marker (+enum so constants are validated for PascalCase)
// +enum
type ExecutableName string

const (
	ExecKubectl ExecutableName = "kubectl" // Valid: in allowlist
	ExecDocker  ExecutableName = "docker"  // Valid: in allowlist
	ExecHelm    ExecutableName = "helm"    // Valid: in allowlist
	ExecUnknown ExecutableName = "unknown" // want "enum value \"unknown\" should be PascalCase"
)

// Regular enum with PascalCase
// +enum
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
