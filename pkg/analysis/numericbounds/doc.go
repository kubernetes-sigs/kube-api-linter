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
numericbounds is an analyzer that checks for proper bounds validation on numeric fields.

According to Kubernetes API conventions, numeric fields should have appropriate bounds
checking to prevent values that are too small, negative (when not intended), or too large.

This analyzer ensures that:
  - int32 and int64 fields have both minimum and maximum bounds markers
  - For slices of numeric types, the analyzer checks for items:Minimum and items:Maximum markers
  - Type aliases to int32 or int64 are also checked
  - Pointer types (e.g., *int32, []*int64) are unwrapped and validated
  - int64 fields with values outside the JavaScript safe integer range (-(2^53-1) to (2^53-1))
    are flagged, as they may cause precision loss in JavaScript clients

The analyzer checks for the presence of +kubebuilder:validation:Minimum and
+kubebuilder:validation:Maximum markers on numeric fields, or the items: variants for slices.

For int64 fields, if the bounds exceed the JavaScript safe integer range of
[-9007199254740991, 9007199254740991], the analyzer suggests using a string type instead
to avoid precision loss in JavaScript environments.

Examples of valid and invalid code:

Valid:

	type Example struct {
		// +kubebuilder:validation:Minimum=0
		// +kubebuilder:validation:Maximum=100
		Count int32
	}

Invalid:

	type Example struct {
		Count int32 // Missing minimum and maximum markers
	}
*/
package numericbounds
