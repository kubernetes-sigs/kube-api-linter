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
numericbounds is an analyzer that checks numeric fields have appropriate bounds validation markers.

According to Kubernetes API conventions, numeric fields should have bounds checking to prevent
values that are too small, negative (when not intended), or too large.

The analyzer checks that int32, int64, float32, and float64 fields have both minimum and
maximum bounds markers. It supports both kubebuilder markers (+kubebuilder:validation:Minimum/Maximum)
and k8s declarative validation markers (+k8s:minimum/maximum).

For slices of numeric types, the analyzer checks the element type for items:Minimum and items:Maximum markers.

Type aliases are resolved and checked. Pointer types are unwrapped and validated.

Bounds values are validated to be within the type's range:
  - int32: full int32 range (±2^31-1)
  - int64: JavaScript-safe range (±2^53-1) per Kubernetes API conventions
  - float32/float64: within their respective ranges

For int64 fields, Kubernetes API conventions enforce JavaScript-safe bounds (±2^53-1)
to ensure compatibility with JavaScript clients and prevent precision loss.

For arrays of numeric types, the minimum/maximum of each element can be set using
+kubebuilder:validation:items:Minimum and +kubebuilder:validation:items:Maximum markers.
Alternatively, if the array uses a numeric type alias, the markers can be placed on the
alias type definition itself.
*/
package numericbounds
