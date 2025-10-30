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
	"maps"

	"sigs.k8s.io/kube-api-linter/pkg/markers"
)

// ScopeConstraint defines where a marker is allowed to be placed.
type ScopeConstraint string

const (
	// FieldScope indicates the marker can be placed on fields.
	FieldScope ScopeConstraint = "Field"
	// TypeScope indicates the marker can be placed on type definitions.
	TypeScope ScopeConstraint = "Type"
	// AnyScope indicates the marker can be placed on either fields or types.
	AnyScope ScopeConstraint = "Any"
)

// Allows checks if the given scope is allowed by this constraint.
func (s ScopeConstraint) Allows(scope ScopeConstraint) bool {
	if s == AnyScope {
		return true
	}

	return s == scope
}

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

// NamedTypeConstraint specifies how markers should be applied to named types.
type NamedTypeConstraint string

const (
	// NamedTypeConstraintAllowField allows markers on fields with named types.
	// The marker can be placed on the field even if the field uses a named type.
	NamedTypeConstraintAllowField NamedTypeConstraint = "AllowField"

	// NamedTypeConstraintRequireTypeDefinition requires markers to be on the type definition.
	// When a field uses a named type, the marker must be placed on the type definition instead.
	NamedTypeConstraintRequireTypeDefinition NamedTypeConstraint = "RequireTypeDefinition"
)

// MarkerScopeRule defines comprehensive scope validation rules for a marker.
type MarkerScopeRule struct {
	// Identifier is the marker identifier (e.g., "optional", "kubebuilder:validation:Minimum").
	Identifier string `json:"identifier,omitempty"`

	// Scope specifies where the marker can be placed (field vs type).
	Scope ScopeConstraint

	// NamedTypeConstraint specifies how markers should be applied to named types.
	// When a field uses a named type (e.g., type CustomInt int32), this determines
	// whether the marker can be on the field or must be on the type definition.
	// If empty, defaults to AllowField (marker can be placed on either field or type).
	NamedTypeConstraint NamedTypeConstraint `json:"namedTypeConstraint,omitempty"`

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
	MarkerScopePolicyWarn MarkerScopePolicy = "Warn"

	// MarkerScopePolicySuggestFix reports warnings and suggests automatic fixes.
	MarkerScopePolicySuggestFix MarkerScopePolicy = "SuggestFix"
)

// MarkerScopeConfig contains configuration for marker scope validation.
type MarkerScopeConfig struct {
	// OverrideMarkers is a list of marker rules that override default rules for built-in markers.
	// Use this to customize the behavior of standard kubebuilder/controller-runtime markers.
	//
	// Example: Override the built-in "optional" marker
	//   overrideMarkers:
	//     - identifier: "optional"
	//       scope: Field
	OverrideMarkers []MarkerScopeRule `json:"overrideMarkers,omitempty"`

	// CustomMarkers is a list of marker rules for custom markers not included in the default rules.
	// Use this to add validation for your own custom markers.
	//
	// Example: Add a custom marker
	//   customMarkers:
	//     - identifier: "mycompany:validation:CustomMarker"
	//       scope: Any
	//       typeConstraint:
	//         allowedSchemaTypes: ["string"]
	CustomMarkers []MarkerScopeRule `json:"customMarkers,omitempty"`

	// Policy determines whether to suggest fixes or just warn.
	Policy MarkerScopePolicy `json:"policy,omitempty"`
}

// DefaultMarkerRules returns the default marker scope rules with type constraints.
// These rules are based on kubebuilder markers and k8s declarative validation markers.
//
// Users can override these rules or add custom markers by providing a MarkerScopeConfig
// with MarkerRules that will be merged with (and take precedence over) these defaults.
//
// Note: This function currently covers validation and SSA markers with type and struct constraints.
// Markers from crd.go (e.g., resource, subresource) and pkg.go (e.g., groupName, versionName)
// are not included as they don't have type or struct constraints and are out of scope for
// this linter's current validation capabilities.
//
// ref: https://github.com/kubernetes-sigs/controller-tools/blob/v0.19.0/pkg/crd/markers/
func DefaultMarkerRules() map[string]MarkerScopeRule {
	rules := make(map[string]MarkerScopeRule)

	addFieldOnlyMarkers(rules)
	addTypeOnlyMarkers(rules)
	addFieldOrTypeMarkers(rules)
	addNumericMarkers(rules)
	addObjectMarkers(rules)
	addStringMarkers(rules)
	addArrayMarkers(rules)
	addGeneralMarkers(rules)
	addSSATopologyMarkers(rules)
	addArrayItemsMarkers(rules)

	return rules
}

// addFieldOnlyMarkers adds field-only markers based on controller-tools validation.go.
func addFieldOnlyMarkers(rules map[string]MarkerScopeRule) {
	fieldOnlyMarkers := map[string]MarkerScopeRule{
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
	}

	maps.Copy(rules, fieldOnlyMarkers)
}

// addTypeOnlyMarkers adds type-only markers for object-level validation and CRD generation.
func addTypeOnlyMarkers(rules map[string]MarkerScopeRule) {
	typeOnlyMarkers := map[string]MarkerScopeRule{
		// Type-only markers (object-level validation and CRD generation)
		markers.KubebuilderValidationItemsExactlyOneOfMarker: {Scope: TypeScope, TypeConstraint: nil},
		markers.KubebuilderValidationItemsAtMostOneOfMarker:  {Scope: TypeScope, TypeConstraint: nil},
		markers.KubebuilderValidationItemsAtLeastOneOfMarker: {Scope: TypeScope, TypeConstraint: nil},
	}

	maps.Copy(rules, typeOnlyMarkers)
}

// addFieldOrTypeMarkers adds markers that can be applied to both fields and types.
func addFieldOrTypeMarkers(rules map[string]MarkerScopeRule) {
	fieldOrTypeMarkers := map[string]MarkerScopeRule{
		// field-or-type markers
		markers.KubebuilderPruningPreserveUnknownFieldsMarker: {Scope: AnyScope, TypeConstraint: nil},
		markers.KubebuilderTitleMarker:                        {Scope: AnyScope, TypeConstraint: nil},
	}

	maps.Copy(rules, fieldOrTypeMarkers)
}

// addNumericMarkers adds numeric validation markers for integer and number types.
func addNumericMarkers(rules map[string]MarkerScopeRule) {
	numericMarkers := map[string]MarkerScopeRule{
		// numeric markers (field or type, integer or number types)
		markers.KubebuilderMinimumMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeInteger},
			},
		},
		markers.KubebuilderMaximumMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeInteger},
			},
		},
		markers.KubebuilderExclusiveMaximumMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeInteger},
			},
		},
		markers.KubebuilderExclusiveMinimumMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeInteger},
			},
		},
		markers.KubebuilderMultipleOfMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeInteger},
			},
		},
	}

	maps.Copy(rules, numericMarkers)
}

// addObjectMarkers adds object validation markers for struct and map types.
func addObjectMarkers(rules map[string]MarkerScopeRule) {
	objectMarkers := map[string]MarkerScopeRule{
		// object markers (field or type, object types)
		markers.KubebuilderMinPropertiesMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeObject},
			},
		},
		markers.KubebuilderMaxPropertiesMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeObject},
			},
		},
	}

	maps.Copy(rules, objectMarkers)
}

// addStringMarkers adds string validation markers.
func addStringMarkers(rules map[string]MarkerScopeRule) {
	stringMarkers := map[string]MarkerScopeRule{
		// string markers (field or type, string types)
		markers.KubebuilderPatternMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeString},
			},
		},
		markers.KubebuilderMinLengthMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeString},
			},
		},
		markers.KubebuilderMaxLengthMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeString},
			},
		},
	}

	maps.Copy(rules, stringMarkers)
}

// addArrayMarkers adds array validation markers.
func addArrayMarkers(rules map[string]MarkerScopeRule) {
	arrayMarkers := map[string]MarkerScopeRule{
		// array markers (field or type, array types)
		markers.KubebuilderMinItemsMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
			},
		},
		markers.KubebuilderMaxItemsMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
			},
		},
		markers.KubebuilderUniqueItemsMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
			},
		},
	}

	maps.Copy(rules, arrayMarkers)
}

// addGeneralMarkers adds general markers that can apply to any type.
func addGeneralMarkers(rules map[string]MarkerScopeRule) {
	generalMarkers := map[string]MarkerScopeRule{
		// general markers (field or type, any type)
		markers.KubebuilderEnumMarker: {
			Scope: AnyScope,
		},
		markers.KubebuilderFormatMarker: {
			Scope: AnyScope,
		},
		markers.KubebuilderTypeMarker: {
			Scope: AnyScope,
		},
		markers.KubebuilderXValidationMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
		},
	}

	maps.Copy(rules, generalMarkers)
}

// addSSATopologyMarkers adds Server-Side Apply topology markers.
func addSSATopologyMarkers(rules map[string]MarkerScopeRule) {
	ssaMarkers := map[string]MarkerScopeRule{
		// Server-Side Apply topology markers
		markers.KubebuilderListTypeMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
			},
		},
		markers.KubebuilderListMapKeyMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
			},
		},
		markers.KubebuilderMapTypeMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
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
	}

	maps.Copy(rules, ssaMarkers)
}

// addArrayItemsMarkers adds array items markers that validate array elements.
// These validate the ELEMENTS of arrays, not the arrays themselves.
func addArrayItemsMarkers(rules map[string]MarkerScopeRule) {
	addArrayItemsNumericMarkers(rules)
	addArrayItemsStringMarkers(rules)
	addArrayItemsArrayMarkers(rules)
	addArrayItemsObjectMarkers(rules)
	addArrayItemsGeneralMarkers(rules)
}

// addArrayItemsNumericMarkers adds items markers for numeric array elements.
func addArrayItemsNumericMarkers(rules map[string]MarkerScopeRule) {
	itemsNumericMarkers := map[string]MarkerScopeRule{
		markers.KubebuilderItemsMaximumMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeInteger},
				},
			},
		},
		markers.KubebuilderItemsMinimumMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeInteger},
				},
			},
		},
		markers.KubebuilderItemsExclusiveMaximumMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeInteger},
				},
			},
		},
		markers.KubebuilderItemsExclusiveMinimumMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeInteger},
				},
			},
		},
		markers.KubebuilderItemsMultipleOfMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeInteger},
				},
			},
		},
	}

	maps.Copy(rules, itemsNumericMarkers)
}

// addArrayItemsStringMarkers adds items markers for string array elements.
func addArrayItemsStringMarkers(rules map[string]MarkerScopeRule) {
	itemsStringMarkers := map[string]MarkerScopeRule{
		markers.KubebuilderItemsMinLengthMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeString},
				},
			},
		},
		markers.KubebuilderItemsMaxLengthMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeString},
				},
			},
		},
		markers.KubebuilderItemsPatternMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeString},
				},
			},
		},
	}

	maps.Copy(rules, itemsStringMarkers)
}

// addArrayItemsArrayMarkers adds items markers for array-of-arrays.
func addArrayItemsArrayMarkers(rules map[string]MarkerScopeRule) {
	itemsArrayMarkers := map[string]MarkerScopeRule{
		markers.KubebuilderItemsMaxItemsMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				},
			},
		},
		markers.KubebuilderItemsMinItemsMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				},
			},
		},
		markers.KubebuilderItemsUniqueItemsMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				},
			},
		},
	}

	maps.Copy(rules, itemsArrayMarkers)
}

// addArrayItemsObjectMarkers adds items markers for arrays of objects.
func addArrayItemsObjectMarkers(rules map[string]MarkerScopeRule) {
	itemsObjectMarkers := map[string]MarkerScopeRule{
		markers.KubebuilderItemsMinPropertiesMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeObject},
				},
			},
		},
		markers.KubebuilderItemsMaxPropertiesMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeObject},
				},
			},
		},
	}

	maps.Copy(rules, itemsObjectMarkers)
}

// addArrayItemsGeneralMarkers adds general items markers that apply to any element type.
func addArrayItemsGeneralMarkers(rules map[string]MarkerScopeRule) {
	itemsGeneralMarkers := map[string]MarkerScopeRule{
		markers.KubebuilderItemsEnumMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				// Enum can apply to any element type
				ElementConstraint: nil,
			},
		},
		markers.KubebuilderItemsFormatMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				// Format can apply to various types
				ElementConstraint: nil,
			},
		},
		markers.KubebuilderItemsTypeMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				// Type marker can override any element type
				ElementConstraint: nil,
			},
		},
		markers.KubebuilderItemsXValidationMarker: {
			Scope:               AnyScope,
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				// CEL validation can apply to any element type
				ElementConstraint: nil,
			},
		},
	}

	maps.Copy(rules, itemsGeneralMarkers)
}
