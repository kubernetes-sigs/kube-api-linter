/*
Copyright 2025 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
enums is an analyzer that enforces proper usage of enumerated fields in Kubernetes APIs.

Enumerated fields provide better API evolution and self-documentation compared to plain strings.
By using type aliases with explicit enum markers, the API schema can include valid values and
provide better validation at the schema level.

# Rules

The linter checks for three main patterns:
1. Fields must use type aliases, not plain strings: String fields that represent enums should
use a type alias with an +enum marker rather than a raw string type.
2. Type aliases must have +enum marker: Type aliases used for enumerated values must be
annotated with either // +kubebuilder:validation:Enum or // +k8s:enum.
3. Enum values must be PascalCase: Constant values for enums should follow PascalCase naming
(e.g., "PhasePending", "StateRunning") rather than lowercase, snake_case, or SCREAMING_SNAKE_CASE.

# Examples

Good:

	// +kubebuilder:validation:Enum
	type Phase string
	const (
		PhasePending   Phase = "Pending"
		PhaseRunning   Phase = "Running"
		PhaseSucceeded Phase = "Succeeded"
	)
	type MySpec struct {
		Phase Phase
	}

Bad:

	// Missing +enum marker
	type Phase string
	type MySpec struct {
		// Plain string without type alias
		Phase string
	}
	// Values not PascalCase
	const (
		phase_pending Phase = "pending"      // Should be "Pending"
		PHASE_RUNNING Phase = "RUNNING"      // Should be "Running"
	)

# Configuration

The linter supports an allowlist for enum values that should be exempt from PascalCase
validation, such as command-line executable names:

	linterConfig:
	  enums:
	    allowlist:
	      - kubectl
	      - docker
	      - helm

# Rationale

Using type aliases for enums instead of raw strings provides several benefits:
- API schemas can explicitly list valid enum values
- Better validation at both the schema and runtime level
- Improved documentation and API evolution
- Stronger type safety in generated clients
- Clearer intent for API consumers

The PascalCase convention for enum values aligns with Kubernetes API conventions and
improves readability and consistency across the ecosystem.

Note: This linter is disabled by default as enum usage is recommended but not strictly
required for all Kubernetes APIs. Enable it in your configuration to enforce these conventions.
*/
package enums
