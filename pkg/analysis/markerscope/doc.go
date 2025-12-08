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

// Package markerscope provides a linter that validates markers are applied in the correct scope
// and to compatible data types.
//
// # Scope Validation
//
// Some markers are only valid when applied to specific Go constructs:
//   - Field-only markers: optional, required, nullable
//   - Type/Struct-only markers: MinProperties, MaxProperties, kubebuilder:object:root, kubebuilder:subresource:status
//   - Field or Type markers: default, MinLength, MaxLength, etc.
//
// # Type Constraint Validation
//
// Markers are also validated for type correctness to ensure they are applied to compatible data types:
//   - Numeric markers (Minimum, Maximum, MultipleOf) must be applied to integer or number types
//   - String markers (Pattern, MinLength, MaxLength) must be applied to string types
//   - Array markers (MinItems, MaxItems, UniqueItems) must be applied to array types
//   - Object markers (MinProperties, MaxProperties) must be applied to object types (struct/map)
//
// For example, applying kubebuilder:validation:Maximum to a string field will be flagged as an error
// since Maximum is only valid for numeric types.
//
// # Array Element Type Constraints
//
// For array types, element-level constraints can be specified using items: prefix markers
// (e.g., items:Minimum, items:Pattern). These validate the array element types rather than
// the array itself.
//
// This linter ensures markers are applied in their appropriate contexts and to compatible types
// to prevent configuration errors and improve API consistency.
package markerscope
