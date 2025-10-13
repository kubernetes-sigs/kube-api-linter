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

import (
	"sigs.k8s.io/kube-api-linter/pkg/markers"
)

// ScopeConstraint defines where a marker is allowed to be placed using bit flags.
type ScopeConstraint uint8

const (
	// FieldScope indicates the marker can be placed on fields.
	FieldScope ScopeConstraint = 1 << iota
	// TypeScope indicates the marker can be placed on type definitions.
	TypeScope

	// AnyScope indicates the marker can be placed on either fields or types.
	AnyScope = FieldScope | TypeScope
)

// String returns a human-readable representation of the scope constraint.
func (s ScopeConstraint) String() string {
	switch s {
	case FieldScope:
		return "field"
	case TypeScope:
		return "type"
	case AnyScope:
		return "any"
	default:
		return "unknown"
	}
}

// Allows checks if the given scope is allowed by this constraint.
func (s ScopeConstraint) Allows(scope ScopeConstraint) bool {
	return s&scope != 0
}

// SchemaType represents OpenAPI schema types that markers can target.
type SchemaType string

const (
	// SchemaTypeInteger represents integer types (int, int32, int64, uint, etc.)
	SchemaTypeInteger SchemaType = "integer"
	// SchemaTypeNumber represents floating-point types (float32, float64)
	SchemaTypeNumber SchemaType = "number"
	// SchemaTypeString represents string types
	SchemaTypeString SchemaType = "string"
	// SchemaTypeBoolean represents boolean types
	SchemaTypeBoolean SchemaType = "boolean"
	// SchemaTypeArray represents array/slice types
	SchemaTypeArray SchemaType = "array"
	// SchemaTypeObject represents struct/map types
	SchemaTypeObject SchemaType = "object"
)

// TypeConstraint defines what types a marker can be applied to.
// NOTE: This constraint is only used when the marker is placed on a field (not TypeScope).
// Type-level markers (TypeScope) do not use type constraints.
type TypeConstraint struct {
	// AllowedSchemaTypes specifies the allowed OpenAPI schema types.
	// If nil or empty, any type is allowed.
	// Maps to JSONSchemaProps.Type (integer, number, string, boolean, array, object)
	AllowedSchemaTypes []SchemaType

	// ElementConstraint specifies constraints on slice/array element types.
	// Only applies when AllowSlice or AllowArray is true.
	ElementConstraint *TypeConstraint
}

// MarkerScopeRule defines comprehensive scope validation rules for a marker.
type MarkerScopeRule struct {
	// Scope specifies where the marker can be placed (field vs type).
	Scope ScopeConstraint

	// TypeConstraint specifies what types the marker can be applied to.
	// NOTE: This is used for both field and type scopes, but typically only enforced
	// when Scope includes FieldScope. For TypeScope-only markers, this is usually nil.
	// If nil, no type constraint is enforced (any type is allowed).
	TypeConstraint *TypeConstraint
}

// MarkerScopePolicy defines how the linter should handle violations.
type MarkerScopePolicy string

const (
	// MarkerScopePolicyWarn only reports warnings without suggesting fixes.
	MarkerScopePolicyWarn MarkerScopePolicy = "warn"

	// MarkerScopePolicySuggestFix reports warnings and suggests automatic fixes.
	MarkerScopePolicySuggestFix MarkerScopePolicy = "suggest_fix"
)

// MarkerScopeConfig contains configuration for marker scope validation.
type MarkerScopeConfig struct {
	// MarkerRules maps marker names to their scope rules with scope and type constraints.
	// If a marker is not in this map, no scope validation is performed.
	MarkerRules map[string]MarkerScopeRule `json:"markerRules,omitempty"`

	// Policy determines whether to suggest fixes or just warn.
	Policy MarkerScopePolicy `json:"policy,omitempty"`
}

// DefaultMarkerRules returns the default marker scope rules with type constraints.
// ref: https://github.com/kubernetes-sigs/controller-tools/blob/v0.19.0/pkg/crd/markers/
func DefaultMarkerRules() map[string]MarkerScopeRule {
	return map[string]MarkerScopeRule{
		// Field-only markers (based on controller-tools validation.go)
		markers.OptionalMarker:                    {Scope: FieldScope, TypeConstraint: nil},
		markers.RequiredMarker:                    {Scope: FieldScope, TypeConstraint: nil},
		markers.K8sOptionalMarker:                 {Scope: FieldScope, TypeConstraint: nil},
		markers.K8sRequiredMarker:                 {Scope: FieldScope, TypeConstraint: nil},
		markers.NullableMarker:                    {Scope: FieldScope, TypeConstraint: nil},
		markers.DefaultMarker:                     {Scope: FieldScope, TypeConstraint: nil},
		markers.KubebuilderDefaultMarker:          {Scope: FieldScope, TypeConstraint: nil},
		markers.KubebuilderExampleMarker:          {Scope: FieldScope, TypeConstraint: nil},
		markers.KubebuilderEmbeddedResourceMarker: {Scope: FieldScope, TypeConstraint: nil},
		markers.KubebuilderSchemaLessMarker:       {Scope: FieldScope, TypeConstraint: nil},

		// Type-only markers (object-level validation and CRD generation)
		markers.KubebuilderValidationItemsExactlyOneOfMarker: {Scope: TypeScope, TypeConstraint: nil},
		markers.KubebuilderValidationItemsAtMostOneOfMarker:  {Scope: TypeScope, TypeConstraint: nil},
		markers.KubebuilderValidationItemsAtLeastOneOfMarker: {Scope: TypeScope, TypeConstraint: nil},

		// field-or-type markers
		markers.KubebuilderPruningPreserveUnknownFieldsMarker: {Scope: AnyScope, TypeConstraint: nil},
		markers.KubebuilderTitleMarker:                        {Scope: AnyScope, TypeConstraint: nil},

		// numeric markers (field or type, integer or number types)
		markers.KubebuilderMinimumMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeInteger, SchemaTypeNumber},
			},
		},
		markers.KubebuilderMaximumMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeInteger, SchemaTypeNumber},
			},
		},
		markers.KubebuilderExclusiveMaximumMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeInteger, SchemaTypeNumber},
			},
		},
		markers.KubebuilderExclusiveMinimumMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeInteger, SchemaTypeNumber},
			},
		},
		markers.KubebuilderMultipleOfMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeInteger, SchemaTypeNumber},
			},
		},

		// object markers (field or type, object types)
		markers.KubebuilderMinPropertiesMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeObject},
			},
		},
		markers.KubebuilderMaxPropertiesMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeObject},
			},
		},

		// string markers (field or type, string types)
		markers.KubebuilderPatternMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeString},
			},
		},
		markers.KubebuilderMinLengthMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeString},
			},
		},
		markers.KubebuilderMaxLengthMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeString},
			},
		},

		// array markers (field or type, array types)
		markers.KubebuilderMinItemsMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
			},
		},
		markers.KubebuilderMaxItemsMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
			},
		},
		markers.KubebuilderUniqueItemsMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
			},
		},

		// general markers (field or type, any type)
		markers.KubebuilderEnumMarker:        {Scope: AnyScope, TypeConstraint: nil},
		markers.KubebuilderFormatMarker:      {Scope: AnyScope, TypeConstraint: nil},
		markers.KubebuilderTypeMarker:        {Scope: AnyScope, TypeConstraint: nil},
		markers.KubebuilderXValidationMarker: {Scope: AnyScope, TypeConstraint: nil},

		// Server-Side Apply topology markers
		markers.KubebuilderListTypeMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
			},
		},
		markers.KubebuilderListMapKeyMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
			},
		},
		markers.KubebuilderMapTypeMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeObject},
			},
		},
		markers.KubebuilderStructTypeMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeObject},
			},
		},

		// Array items markers (field or type, apply to array elements)
		// These validate the ELEMENTS of arrays, not the arrays themselves
		markers.KubebuilderItemsMaxItemsMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				},
			},
		},
		markers.KubebuilderItemsMaximumMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeInteger, SchemaTypeNumber},
				},
			},
		},
		markers.KubebuilderItemsMinItemsMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				},
			},
		},
		markers.KubebuilderItemsMinLengthMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeString},
				},
			},
		},
		markers.KubebuilderItemsMinimumMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeInteger, SchemaTypeNumber},
				},
			},
		},
		markers.KubebuilderItemsMaxLengthMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeString},
				},
			},
		},
		markers.KubebuilderItemsEnumMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				// Enum can apply to any element type
				ElementConstraint: nil,
			},
		},
		markers.KubebuilderItemsFormatMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				// Format can apply to various types
				ElementConstraint: nil,
			},
		},
		markers.KubebuilderItemsExclusiveMaximumMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeInteger, SchemaTypeNumber},
				},
			},
		},
		markers.KubebuilderItemsExclusiveMinimumMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeInteger, SchemaTypeNumber},
				},
			},
		},
		markers.KubebuilderItemsMultipleOfMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeInteger, SchemaTypeNumber},
				},
			},
		},
		markers.KubebuilderItemsPatternMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeString},
				},
			},
		},
		markers.KubebuilderItemsTypeMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				// Type marker can override any element type
				ElementConstraint: nil,
			},
		},
		markers.KubebuilderItemsUniqueItemsMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				},
			},
		},
		markers.KubebuilderItemsXValidationMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				// CEL validation can apply to any element type
				ElementConstraint: nil,
			},
		},
		markers.KubebuilderItemsMinPropertiesMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeObject},
				},
			},
		},
		markers.KubebuilderItemsMaxPropertiesMarker: {
			Scope: AnyScope,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeObject},
				},
			},
		},
		// TODO crd.go
		// TODO package.go
	}
}
