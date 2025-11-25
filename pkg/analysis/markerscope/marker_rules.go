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

// defaultMarkerRules returns the default marker scope rules with type constraints.
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
func defaultMarkerRules() map[string]MarkerScopeRule {
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
		markers.OptionalMarker:                    {Scopes: []ScopeConstraint{FieldScope}, TypeConstraint: nil},
		markers.RequiredMarker:                    {Scopes: []ScopeConstraint{FieldScope}, TypeConstraint: nil},
		markers.K8sOptionalMarker:                 {Scopes: []ScopeConstraint{FieldScope}, TypeConstraint: nil},
		markers.K8sRequiredMarker:                 {Scopes: []ScopeConstraint{FieldScope}, TypeConstraint: nil},
		markers.NullableMarker:                    {Scopes: []ScopeConstraint{FieldScope}, TypeConstraint: nil},
		markers.DefaultMarker:                     {Scopes: []ScopeConstraint{FieldScope}, TypeConstraint: nil},
		markers.KubebuilderDefaultMarker:          {Scopes: []ScopeConstraint{FieldScope}, TypeConstraint: nil},
		markers.KubebuilderExampleMarker:          {Scopes: []ScopeConstraint{FieldScope}, TypeConstraint: nil},
		markers.KubebuilderEmbeddedResourceMarker: {Scopes: []ScopeConstraint{FieldScope}, TypeConstraint: nil},
		markers.KubebuilderSchemaLessMarker:       {Scopes: []ScopeConstraint{FieldScope}, TypeConstraint: nil},
	}

	maps.Copy(rules, fieldOnlyMarkers)
}

// addTypeOnlyMarkers adds type-only markers for object-level validation and CRD generation.
func addTypeOnlyMarkers(rules map[string]MarkerScopeRule) {
	typeOnlyMarkers := map[string]MarkerScopeRule{
		// Type-only markers (object-level validation and CRD generation)
		markers.KubebuilderValidationItemsExactlyOneOfMarker: {Scopes: []ScopeConstraint{TypeScope}, TypeConstraint: nil},
		markers.KubebuilderValidationItemsAtMostOneOfMarker:  {Scopes: []ScopeConstraint{TypeScope}, TypeConstraint: nil},
		markers.KubebuilderValidationItemsAtLeastOneOfMarker: {Scopes: []ScopeConstraint{TypeScope}, TypeConstraint: nil},
	}

	maps.Copy(rules, typeOnlyMarkers)
}

// addFieldOrTypeMarkers adds markers that can be applied to both fields and types.
func addFieldOrTypeMarkers(rules map[string]MarkerScopeRule) {
	fieldOrTypeMarkers := map[string]MarkerScopeRule{
		// field-or-type markers
		markers.KubebuilderPruningPreserveUnknownFieldsMarker: {Scopes: []ScopeConstraint{FieldScope, TypeScope}, TypeConstraint: nil},
		markers.KubebuilderTitleMarker:                        {Scopes: []ScopeConstraint{FieldScope, TypeScope}, TypeConstraint: nil},
	}

	maps.Copy(rules, fieldOrTypeMarkers)
}

// addNumericMarkers adds numeric validation markers for integer and number types.
func addNumericMarkers(rules map[string]MarkerScopeRule) {
	numericMarkers := map[string]MarkerScopeRule{
		// numeric markers (field or type, integer or number types)
		markers.KubebuilderMinimumMarker: {
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeInteger},
			},
		},
		markers.KubebuilderMaximumMarker: {
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeInteger},
			},
		},
		markers.KubebuilderExclusiveMaximumMarker: {
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeInteger},
			},
		},
		markers.KubebuilderExclusiveMinimumMarker: {
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeInteger},
			},
		},
		markers.KubebuilderMultipleOfMarker: {
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
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
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeObject},
			},
		},
		markers.KubebuilderMaxPropertiesMarker: {
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
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
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeString},
			},
		},
		markers.KubebuilderMinLengthMarker: {
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeString},
			},
		},
		markers.KubebuilderMaxLengthMarker: {
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
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
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
			},
		},
		markers.KubebuilderMaxItemsMarker: {
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
			},
		},
		markers.KubebuilderUniqueItemsMarker: {
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
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
			Scopes: []ScopeConstraint{FieldScope, TypeScope},
		},
		markers.KubebuilderFormatMarker: {
			Scopes: []ScopeConstraint{FieldScope, TypeScope},
		},
		markers.KubebuilderTypeMarker: {
			Scopes: []ScopeConstraint{FieldScope, TypeScope},
		},
		markers.KubebuilderXValidationMarker: {
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
		},
	}

	maps.Copy(rules, generalMarkers)
}

// addSSATopologyMarkers adds Server-Side Apply topology markers.
func addSSATopologyMarkers(rules map[string]MarkerScopeRule) {
	ssaMarkers := map[string]MarkerScopeRule{
		// Server-Side Apply topology markers
		markers.ListTypeMarker: {
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
			},
		},
		markers.ListMapKeyMarker: {
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
			},
		},
		markers.MapTypeMarker: {
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeObject},
			},
		},
		markers.StructTypeMarker: {
			Scopes: []ScopeConstraint{FieldScope, TypeScope},
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
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeInteger},
				},
			},
		},
		markers.KubebuilderItemsMinimumMarker: {
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeInteger},
				},
			},
		},
		markers.KubebuilderItemsExclusiveMaximumMarker: {
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeInteger},
				},
			},
		},
		markers.KubebuilderItemsExclusiveMinimumMarker: {
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeInteger},
				},
			},
		},
		markers.KubebuilderItemsMultipleOfMarker: {
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
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
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeString},
				},
			},
		},
		markers.KubebuilderItemsMaxLengthMarker: {
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeString},
				},
			},
		},
		markers.KubebuilderItemsPatternMarker: {
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
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
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				},
			},
		},
		markers.KubebuilderItemsMinItemsMarker: {
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				},
			},
		},
		markers.KubebuilderItemsUniqueItemsMarker: {
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
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
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				ElementConstraint: &TypeConstraint{
					AllowedSchemaTypes: []SchemaType{SchemaTypeObject},
				},
			},
		},
		markers.KubebuilderItemsMaxPropertiesMarker: {
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
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
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				// Enum can apply to any element type
				ElementConstraint: nil,
			},
		},
		markers.KubebuilderItemsFormatMarker: {
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				// Format can apply to various types
				ElementConstraint: nil,
			},
		},
		markers.KubebuilderItemsTypeMarker: {
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
			NamedTypeConstraint: NamedTypeConstraintRequireTypeDefinition,
			TypeConstraint: &TypeConstraint{
				AllowedSchemaTypes: []SchemaType{SchemaTypeArray},
				// Type marker can override any element type
				ElementConstraint: nil,
			},
		},
		markers.KubebuilderItemsXValidationMarker: {
			Scopes:              []ScopeConstraint{FieldScope, TypeScope},
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
