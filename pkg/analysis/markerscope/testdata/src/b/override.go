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
package b

// ============================================================================
// Override marker tests
// These test that customMarkers configuration works correctly for overriding built-in markers.
//
// The TestAnalyzerWithCustomAndOverrideMarkers configures:
// - optional: [Field, Type] (overrides default FieldScope)
// - required: [Field, Type] (overrides default FieldScope)
// ============================================================================

// Built-in markers should work with their default scope on fields
type BuiltInMarkerTest struct {
	// Valid: optional is allowed on fields ([Field, Type] includes FieldScope)
	// +optional
	ValidOptionalField string `json:"validOptionalField"`

	// Valid: required is allowed on fields ([Field, Type] includes FieldScope)
	// +required
	ValidRequiredField string `json:"validRequiredField"`
}

// Built-in markers on types should now be VALID with overridden [Field, Type]
// +optional
type ValidOptionalOnType struct {
	Name string `json:"name"`
}

// +required
type ValidRequiredOnType struct {
	Name string `json:"name"`
}

// Test that non-overridden markers still follow default rules
type NonOverriddenMarkerTest struct {
	// Valid: nullable is FieldScope by default (not overridden)
	// +nullable
	ValidNullableField *string `json:"validNullableField"`
}

// +nullable // want `marker "nullable" can only be applied to fields`
type InvalidNullableOnType struct {
	Name string `json:"name"`
}
