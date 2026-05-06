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
//   - Field-only markers: optional, required, nullable, default, kubebuilder:default, kubebuilder:example, kubebuilder:validation:EmbeddedResource, kubebuilder:validation:Schemaless
//   - Type-only markers: kubebuilder:validation:items:ExactlyOneOf, kubebuilder:validation:items:AtMostOneOf, kubebuilder:validation:items:AtLeastOneOf
//   - Field or Type markers: kubebuilder:validation:Minimum, kubebuilder:validation:Maximum, kubebuilder:validation:MinLength, kubebuilder:validation:MaxLength, kubebuilder:validation:MinProperties, kubebuilder:validation:MaxProperties, etc.
//
// # Type Constraint Validation
//
// Markers are also validated for type correctness to ensure they are applied to compatible data types:
//   - Numeric markers (Minimum, Maximum, MultipleOf, ExclusiveMinimum, ExclusiveMaximum) require integer or number types
//   - String markers (Pattern, MinLength, MaxLength, Format) require string types
//   - Array markers (MinItems, MaxItems, UniqueItems) require array types
//   - Object markers (MinProperties, MaxProperties) require object types (struct or map)
//
// For example, applying kubebuilder:validation:Maximum to a string field will be flagged as an error
// since Maximum is only valid for numeric types.
//
// # Array Element Type Constraints
//
// For array types, element-level constraints can be specified using kubebuilder:validation:items: prefix markers
// (e.g., kubebuilder:validation:items:Minimum, kubebuilder:validation:items:Pattern). These validate the array
// element types rather than the array itself.
//
// # Named Type Preference
//
// For markers that can be applied to both fields and types, the namedTypePreference setting controls
// where the marker should be declared when used with named types:
//
//   - AllowTypeOrField (default): Marker can be declared on either the field or the type definition
//   - OnTypeOnly: Marker must be declared on the type definition, not on fields using that type
//
// Most built-in kubebuilder validation markers use OnTypeOnly to encourage consistent marker placement
// on type definitions. For example:
//
//	// Valid: marker on type definition
//	// +kubebuilder:validation:Minimum=0
//	type Port int32
//
//	type Service struct {
//	    Port Port `json:"port"`
//	}
//
//	// Invalid with OnTypeOnly: marker on field using named type
//	type Port int32
//
//	type Service struct {
//	    // +kubebuilder:validation:Minimum=0  // Error: should be on Port type definition
//	    Port Port `json:"port"`
//	}
//
// You can set a global namedTypePreference for uniform behavior across all markers, or override it
// per marker using the marker's namedTypePreference field.
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
//	lintersConfig:
//	  markerscope:
//	    namedTypePreference: OnTypeOnly  # Global setting for all markers
//	    customMarkers:
//	      - identifier: "optional"
//	        scopes: [Field, Type]  # Override default [Field] to allow on both
//
// Example configuration to add a custom marker with type constraints:
//
//	customMarkers:
//	  - identifier: "mycompany:validation:CustomMarker"
//	    scopes: [Field, Type]
//	    namedTypePreference: OnTypeOnly  # Require on type definition for named types
//	    typeConstraint:
//	      allowedSchemaTypes: ["string"]
//
// This linter ensures markers are applied in their appropriate contexts and to compatible types
// to prevent configuration errors and improve API consistency.
package markerscope
