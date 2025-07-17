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
package conflictingmarkers

// ConflictingMarkersConfig contains the configuration for the conflictingmarkers linter.
type ConflictingMarkersConfig struct {
	// CustomConflicts allows users to define their own sets of conflicting markers.
	// Each entry defines a conflict between two sets of markers.
	CustomConflicts []ConflictSet `json:"customConflicts"`

	// DisableBuiltInConflicts allows users to opt-out of built-in conflict detection.
	// When set to true, only custom conflicts defined in CustomConflicts will be checked.
	// Built-in conflicts include optional_vs_required and default_vs_required checks.
	DisableBuiltInConflicts bool `json:"disableBuiltInConflicts"`
}

// ConflictSet represents a conflict between two sets of markers.
// Markers within each set are mutually exclusive with markers in the other set.
type ConflictSet struct {
	// Name is a human-readable name for this conflict set.
	// This name will appear in diagnostic messages to identify the type of conflict,
	Name string `json:"name"`
	// SetA contains the first set of markers that conflict with markers in SetB.
	// These markers are mutually exclusive with markers in setB.
	// The linter will emit a diagnostic when a field has markers from both setA and setB.
	SetA []string `json:"setA"`
	// SetB contains the second set of markers that conflict with markers in SetA.
	// These markers are mutually exclusive with markers in setA.
	// The linter will emit a diagnostic when a field has markers from both setA and setB.
	SetB []string `json:"setB"`
	// Description provides a description of why these markers conflict.
	// The linter will include this description in the diagnostic message when a conflict is detected.
	Description string `json:"description"`
}
