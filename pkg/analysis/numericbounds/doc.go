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
  - int32, int64, float32, and float64 fields have both minimum and maximum bounds markers
  - For slices of numeric types, the analyzer checks the element type for items:Minimum and items:Maximum markers
    (e.g., []int32 checks if the int32 elements have bounds, not the array itself)
  - Type aliases to numeric types are recursively resolved and checked
  - Pointer types (e.g., *int32, []*int64) are unwrapped and validated
  - Both kubebuilder and k8s declarative validation markers are supported
  - Bounds values are validated to be within the type's range (e.g., int32 bounds must fit in int32)
  - int64 fields with values outside the JavaScript safe integer range (±2^53-1)
    are flagged, as they may cause precision loss in JavaScript clients

The analyzer checks for the presence of +kubebuilder:validation:Minimum and
+kubebuilder:validation:Maximum markers on numeric fields, or the items: variants for slices.
It also supports +k8s:minimum and +k8s:maximum for declarative validation.

For int64 fields, if the bounds exceed the JavaScript safe integer range of
[-9007199254740991, 9007199254740991], the analyzer suggests using a string type instead
to avoid precision loss in JavaScript environments.

# Examples

## Valid: Numeric field with proper bounds markers

	type Example struct {
		// +kubebuilder:validation:Minimum=0
		// +kubebuilder:validation:Maximum=100
		Count int32
	}

## Valid: Int64 field with JavaScript-safe bounds

	type Example struct {
		// +kubebuilder:validation:Minimum=-9007199254740991
		// +kubebuilder:validation:Maximum=9007199254740991
		Timestamp int64
	}

## Valid: Float field with bounds

	type Example struct {
		// +kubebuilder:validation:Minimum=0.0
		// +kubebuilder:validation:Maximum=100.0
		Ratio float32
	}

## Valid: Slice with items bounds

	type Example struct {
		// +kubebuilder:validation:items:Minimum=1
		// +kubebuilder:validation:items:Maximum=65535
		Ports []int32
	}

## Valid: Using k8s declarative validation markers

	type Example struct {
		// +k8s:minimum=0
		// +k8s:maximum=100
		Count int32
	}

## Invalid: Missing bounds markers

	type Example struct {
		Count int32 // want: should have minimum and maximum bounds validation markers
	}

## Invalid: Only one bound specified

	type Example struct {
		// +kubebuilder:validation:Minimum=0
		Count int32 // want: has minimum but is missing maximum bounds validation marker
	}

## Invalid: Int64 with bounds exceeding JavaScript safe range

	type Example struct {
		// +kubebuilder:validation:Minimum=-10000000000000000
		// +kubebuilder:validation:Maximum=10000000000000000
		LargeNumber int64 // want: bounds exceed JavaScript safe integer range
	}

## Invalid: Int32 bounds outside valid int32 range

	type Example struct {
		// +kubebuilder:validation:Minimum=-3000000000
		// +kubebuilder:validation:Maximum=3000000000
		Count int32 // want: bounds outside valid int32 range
	}
*/
package numericbounds
