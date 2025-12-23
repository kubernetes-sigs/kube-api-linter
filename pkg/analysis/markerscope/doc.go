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
//   - Field-only markers: optional, required, kubebuilder:validation:XValidation
//   - Type-only markers: kubebuilder:validation:MinProperties, kubebuilder:validation:MaxProperties
//   - Field or Type markers: kubebuilder:default, kubebuilder:validation:MinLength, kubebuilder:validation:MaxLength, etc.
//
// # Type Constraint Validation
//
// Markers are also validated for type correctness to ensure they are applied to compatible data types:
//   - Numeric markers (kubebuilder:validation:Minimum, kubebuilder:validation:Maximum, kubebuilder:validation:MultipleOf) must be applied to integer and number types
//   - String markers (kubebuilder:validation:Pattern, kubebuilder:validation:MinLength, kubebuilder:validation:MaxLength) must be applied to string types
//   - Array markers (kubebuilder:validation:MinItems, kubebuilder:validation:MaxItems, kubebuilder:validation:UniqueItems) must be applied to array types
//   - Object markers (kubebuilder:validation:MinProperties, kubebuilder:validation:MaxProperties) must be applied to object types (struct/map)
//
// For example, applying kubebuilder:validation:Maximum to a string field will be flagged as an error
// since kubebuilder:validation:Maximum is only valid for numeric types.
//
// # Array Element Type Constraints
//
// For array types, element-level constraints can be specified using kubebuilder:validation:items: prefix markers
// (e.g., kubebuilder:validation:items:Minimum, kubebuilder:validation:items:Pattern). These validate the array
// element types rather than the array itself.
//
// # Custom Markers Configuration
//
// The analyzer includes 100+ built-in kubebuilder marker rules. You can customize the behavior using
// the customMarkers configuration field, which supports two use cases:
//
//  1. Overriding built-in markers: Change the validation rules for existing markers
//     by providing a marker rule with the same identifier as a built-in marker.
//
//  2. Adding custom markers: Define validation rules for your own custom markers
//     that are not included in the default rules.
//
// Example configuration to override a built-in marker:
//
//	customMarkers:
//	  - identifier: "optional"
//	    scopes: [Field, Type]  # Override default [Field] to allow on both
//
// Example configuration to add a custom marker:
//
//	customMarkers:
//	  - identifier: "mycompany:validation:CustomMarker"
//	    scopes: [Field]
//	    typeConstraint:
//	      allowedSchemaTypes: ["string"]
//
// This linter ensures markers are applied in their appropriate contexts and to compatible types
// to prevent configuration errors and improve API consistency.
package markerscope
