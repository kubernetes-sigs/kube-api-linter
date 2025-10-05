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

import "sigs.k8s.io/kube-api-linter/pkg/markers"

// MarkerScope defines where a marker is allowed to be placed.
type MarkerScope string

const (
	// ScopeField indicates the marker can only be placed on fields.
	ScopeField MarkerScope = "field"

	// ScopeType indicates the marker can only be placed on type definitions.
	ScopeType MarkerScope = "type"

	// ScopeFieldOrType indicates the marker can be placed on either fields or types.
	ScopeFieldOrType MarkerScope = "field_or_type"

	// ScopeTypeOrObjectField indicates the marker can be placed on type definitions or object fields (struct/map).
	ScopeTypeOrObjectField MarkerScope = "type_or_object_field"
)

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
	// Markers maps marker names to their allowed scopes.
	// If a marker is not in this map, no scope validation is performed.
	Markers map[string]MarkerScope `json:"markers,omitempty"`

	// Policy determines whether to suggest fixes or just warn.
	Policy MarkerScopePolicy `json:"policy,omitempty"`
}

// DefaultMarkerScopes returns the default marker scope configurations.
// ref: https://github.com/kubernetes-sigs/controller-tools/blob/main/pkg/crd/markers/validation.go
func DefaultMarkerScopes() map[string]MarkerScope {
	return map[string]MarkerScope{
		// Field-only markers (based on controller-tools validation.go)
		markers.OptionalMarker:                    ScopeField,
		markers.RequiredMarker:                    ScopeField,
		markers.K8sOptionalMarker:                 ScopeField,
		markers.K8sRequiredMarker:                 ScopeField,
		markers.KubebuilderOptionalMarker:         ScopeField,
		markers.KubebuilderRequiredMarker:         ScopeField,
		markers.NullableMarker:                    ScopeField,
		markers.DefaultMarker:                     ScopeField,
		markers.KubebuilderDefaultMarker:          ScopeField,
		markers.KubebuilderExampleMarker:          ScopeField,
		"kubebuilder:validation:EmbeddedResource": ScopeField,
		markers.KubebuilderSchemaLessMarker:       ScopeField,

		// Type-only markers (object-level validation and CRD generation)
		"kubebuilder:validation:items:ExactlyOneOf": ScopeType,
		"kubebuilder:validation:items:AtMostOneOf":  ScopeType,
		"kubebuilder:validation:items:AtLeastOneOf": ScopeType,
		markers.KubebuilderRootMarker:               ScopeType,
		markers.KubebuilderStatusSubresourceMarker:  ScopeType,

		// field-and-type markers
		"kubebuilder:pruning:PreserveUnknownFields": ScopeFieldOrType,
		"kubebuilder:title":                         ScopeFieldOrType,

		// numeric markers
		markers.KubebuilderMinimumMarker:          ScopeField,
		markers.KubebuilderMaximumMarker:          ScopeField,
		markers.KubebuilderExclusiveMaximumMarker: ScopeField,
		markers.KubebuilderExclusiveMinimumMarker: ScopeField,
		markers.KubebuilderMultipleOfMarker:       ScopeField,

		// object markers
		markers.KubebuilderMinPropertiesMarker: ScopeTypeOrObjectField,
		markers.KubebuilderMaxPropertiesMarker: ScopeTypeOrObjectField,

		// string markers
		markers.KubebuilderPatternMarker:   ScopeField,
		markers.KubebuilderMinLengthMarker: ScopeField,
		markers.KubebuilderMaxLengthMarker: ScopeField,

		// array markers
		markers.KubebuilderMinItemsMarker:    ScopeField,
		markers.KubebuilderMaxItemsMarker:    ScopeField,
		markers.KubebuilderUniqueItemsMarker: ScopeField,

		// general markers
		markers.KubebuilderEnumMarker:        ScopeField,
		markers.KubebuilderFormatMarker:      ScopeField,
		markers.KubebuilderTypeMarker:        ScopeField,
		markers.KubebuilderXValidationMarker: ScopeField,

		// Array/slice field markers (Server-Side Apply related)
		markers.KubebuilderListTypeMarker:   ScopeFieldOrType,
		markers.KubebuilderListMapKeyMarker: ScopeFieldOrType,

		// Array items markers (field-only, apply to array elements)
		markers.KubebuilderItemsMaxItemsMarker:         ScopeField,
		markers.KubebuilderItemsMaximumMarker:          ScopeField,
		markers.KubebuilderItemsMinItemsMarker:         ScopeField,
		markers.KubebuilderItemsMinLengthMarker:        ScopeField,
		markers.KubebuilderItemsMinimumMarker:          ScopeField,
		markers.KubebuilderItemsMaxLengthMarker:        ScopeField,
		markers.KubebuilderItemsEnumMarker:             ScopeField,
		markers.KubebuilderItemsFormatMarker:           ScopeField,
		markers.KubebuilderItemsExclusiveMaximumMarker: ScopeField,
		markers.KubebuilderItemsExclusiveMinimumMarker: ScopeField,
		markers.KubebuilderItemsMultipleOfMarker:       ScopeField,
		markers.KubebuilderItemsPatternMarker:          ScopeField,
		markers.KubebuilderItemsTypeMarker:             ScopeField,
		markers.KubebuilderItemsUniqueItemsMarker:      ScopeField,
		markers.KubebuilderItemsXValidationMarker:      ScopeField,
		markers.KubebuilderItemsMinPropertiesMarker:    ScopeTypeOrObjectField,
		markers.KubebuilderItemsMaxPropertiesMarker:    ScopeTypeOrObjectField,
	}
}
