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
package markerscope

import "slices"

// ScopeConstraint defines where a marker is allowed to be placed.
type ScopeConstraint string

const (
	// FieldScope indicates the marker can be placed on fields.
	FieldScope ScopeConstraint = "Field"
	// TypeScope indicates the marker can be placed on type definitions.
	TypeScope ScopeConstraint = "Type"
)

// AllowsScope checks if the given scope is allowed by this rule.
func (r MarkerScopeRule) AllowsScope(scope ScopeConstraint) bool {
	return slices.Contains(r.Scopes, scope)
}

// TypeConstraint specifies what types a marker can be applied to.
type TypeConstraint struct {
	// allowedSchemaTypes specifies the allowed OpenAPI schema types.
	// If nil or empty, any type is allowed.
	// Maps to JSONSchemaProps.Type (integer, string, boolean, array, object)
	AllowedSchemaTypes []SchemaType `json:"allowedSchemaTypes,omitempty"`

	// elementConstraint specifies constraints on slice/array element types.
	// Only applies when allowedSchemaTypes includes array.
	ElementConstraint *TypeConstraint `json:"elementConstraint,omitempty"`
}

// NamedTypeConstraint specifies how markers should be applied to named types.
type NamedTypeConstraint string

const (
	// NamedTypeConstraintAllowTypeOrField allows markers on either the field or the type definition.
	// The marker can be placed on the field even if the field uses a named type.
	NamedTypeConstraintAllowTypeOrField NamedTypeConstraint = "AllowTypeOrField"

	// NamedTypeConstraintOnTypeOnly requires markers to be on the type definition only.
	// When a field uses a named type, the marker must be placed on the type definition instead.
	NamedTypeConstraintOnTypeOnly NamedTypeConstraint = "OnTypeOnly"
)

// MarkerScopeRule defines comprehensive scope validation rules for a marker.
type MarkerScopeRule struct {
	// identifier is the marker identifier (e.g., "optional", "kubebuilder:validation:Minimum").
	Identifier string `json:"identifier,omitempty"`

	// scopes specifies where the marker can be placed (Field, Type, or both).
	Scopes []ScopeConstraint `json:"scopes,omitempty"`

	// namedTypeConstraint specifies how markers should be applied to named types.
	// When a field uses a named type (e.g., type CustomInt int32), this determines
	// whether the marker can be on the field or must be on the type definition.
	// Valid values: AllowTypeOrField, OnTypeOnly.
	// If empty, defaults to AllowTypeOrField.
	NamedTypeConstraint NamedTypeConstraint `json:"namedTypeConstraint,omitempty"`

	// typeConstraint specifies what types the marker can be applied to.
	// This is used for both Field and Type scope markers.
	// If nil, no type constraint is enforced (any type is allowed).
	TypeConstraint *TypeConstraint `json:"typeConstraint,omitempty"`
}

// MarkerScopePolicy defines how the linter should handle violations.
type MarkerScopePolicy string

const (
	// MarkerScopePolicyWarn only reports warnings without suggesting fixes.
	MarkerScopePolicyWarn MarkerScopePolicy = "Warn"

	// MarkerScopePolicySuggestFix reports warnings and suggests automatic fixes.
	MarkerScopePolicySuggestFix MarkerScopePolicy = "SuggestFix"
)

// MarkerScopeConfig contains configuration for marker scope validation.
type MarkerScopeConfig struct {
	// customMarkers is a list of marker rules for custom markers.
	// This can include both:
	// 1. New custom markers not included in the default rules
	// 2. Overrides for built-in markers (if the identifier matches a built-in marker)
	//
	// Example: Override the built-in "optional" marker
	//   customMarkers:
	//     - identifier: "optional"
	//       scope: Field
	//
	// Example: Add a custom marker
	//   customMarkers:
	//     - identifier: "mycompany:validation:CustomMarker"
	//       scopes: [Field, Type]
	//       typeConstraint:
	//         allowedSchemaTypes: ["string"]
	CustomMarkers []MarkerScopeRule `json:"customMarkers,omitempty"`

	// policy determines whether to suggest fixes or just warn.
	Policy MarkerScopePolicy `json:"policy,omitempty"`
}
