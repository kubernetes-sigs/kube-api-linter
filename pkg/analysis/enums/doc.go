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

1. String type aliases that have associated constants must be annotated with an enum marker:
  - +enum (plain marker, recommended primary idiom)
  - +k8s:enum for declarative validation (used in Kubernetes core API types)
  - +kubebuilder:validation:Enum=Value1;Value2 for CRD validation (used in projects with CustomResourceDefinitions)

2. Enum constant values must follow PascalCase when using auto-discovery mode.
Valid: "Pending", "Running", "HTTP", "HTTPS" (acronyms allowed).
Invalid: "pending", "phase_pending", "Phase-Failed" (snake_case/kebab-case).

3. (Optional) When KubebuilderEnumPolicy is RequireTypeAlias: String fields without enum markers
should use type aliases instead of plain string types.

The linter only flags type aliases that have constants defined, avoiding false positives
for generic string wrapper types.

# Enum Marker Types

The linter recognizes three enum markers:

+enum (Primary Kubernetes idiom):
- Plain marker, the canonical/recommended approach
- Auto-discovers constants and validates them as PascalCase

+k8s:enum (Declarative validation for core APIs):
- Used in Kubernetes core API types (in-tree APIs)
- Auto-discovers constants and validates them as PascalCase

+kubebuilder:validation:Enum (CRD OpenAPI schema validation):
- Used in projects with CustomResourceDefinitions
- Processed by controller-gen to generate OpenAPI schema validation
- REQUIRES explicit values: +kubebuilder:validation:Enum=Value1;Value2;Value3
- Does NOT auto-discover constants
- The explicit values in the marker are validated as PascalCase

Examples:

Auto-discovery (validates constants):

	// +enum                    ← Primary Kubernetes idiom (recommended)
	// +k8s:enum                ← Alternative for core APIs
	type Phase string
	const (
		PhasePending Phase = "Pending"  ← These must be PascalCase
	)

Explicit values (validates marker values, not constants):

	// +kubebuilder:validation:Enum=Pending;Running;Failed  ← For CRD schema validation
	type Phase string
	const (
		helper Phase = "helper"  ← Constants not checked; marker values are validated
	)

# Examples

Good:

	// +enum  (or // +kubebuilder:validation:Enum for CRDs, // +k8s:enum for core APIs)
	type Phase string
	const (
		PhasePending Phase = "Pending"
		PhaseRunning Phase = "Running"
	)
	type MySpec struct {
		Phase Phase
	}

Bad:

	// String type alias with constants but missing enum marker
	type Phase string
	const (
		phase_pending Phase = "pending"      // Should be "Pending"
		phase_running Phase = "phase_running" // Should be "PhaseRunning"
		Phase_Failed  Phase = "Phase-Failed"  // Should be "PhaseFailed"
	)

Note: Acronyms (HTTP, HTTPS, API) are allowed. The linter only flags type aliases with
constants, not all string types.

# Configuration

Configuration options:

	linterConfig:
	  enums:
	    # Values exempt from PascalCase validation
	    allowlist:
	      - kubectl
	      - docker
	    # Require type aliases (RequireTypeAlias, default) or allow plain strings (AllowPlainString)
	    kubebuilderEnumPolicy: RequireTypeAlias

# Rationale

Using type aliases with enum markers instead of raw strings provides several benefits:
- API schemas can explicitly list valid enum values
- Better validation at both the schema and runtime level
- Improved documentation and API evolution
- Stronger type safety in generated clients
- Clearer intent for API consumers

The PascalCase convention for enum values aligns with Kubernetes API conventions and
improves readability and consistency across the ecosystem.

The distinction between CRD validation markers and declarative validation markers allows
the linter to work correctly in both CRD-based projects (using controller-gen) and
Kubernetes core API development (using declarative validation).

Note: This linter is enabled by default to ensure string types with constants follow enum
conventions. It only flags types that have associated constants, minimizing false positives.
*/
package enums
