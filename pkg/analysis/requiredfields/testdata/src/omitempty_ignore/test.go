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
package omitempty_ignore

// TestOmitEmptyIgnore tests that when omitempty policy is set to Ignore,
// fields that don't allow the zero value still report diagnostics because
// this is a correctness issue (the field must have omitempty to work correctly).
type TestOmitEmptyIgnore struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// This field does not allow the zero value (empty string) and does not have omitempty.
	// Even with Ignore policy, this must be reported as it's a correctness issue.
	RequiredFieldWithoutOmitEmpty string `json:"requiredField"` // want "field TestOmitEmptyIgnore.RequiredFieldWithoutOmitEmpty does not allow the zero value. It must have the omitempty tag."

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// This field has omitempty, so it should work normally.
	RequiredFieldWithOmitEmpty string `json:"requiredFieldWithOmitEmpty,omitempty"`
}
