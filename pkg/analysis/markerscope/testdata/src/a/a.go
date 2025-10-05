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
package a

// +kubebuilder:validation:MinProperties=1
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type ValidTypeMarkers struct {
	Name string `json:"name"`
}

// +optional // want `marker "optional" can only be applied to fields, not type definitions`
// +required // want `marker "required" can only be applied to fields, not type definitions`
// +kubebuilder:validation:Optional // want `marker "kubebuilder:validation:Optional" can only be applied to fields, not type definitions`
type InvalidTypeMarkers struct {
	Name string `json:"name"`
}

type FieldMarkerTest struct {
	// +optional
	// +kubebuilder:validation:Optional
	ValidOptionalField string `json:"validOptionalField,omitempty"`

	// +required
	// +kubebuilder:validation:Required
	ValidRequiredField string `json:"validRequiredField"`

	// +kubebuilder:validation:MinProperties=1 // want `marker "kubebuilder:validation:MinProperties" can only be applied to type definitions, not fields`
	InvalidMinPropertiesField map[string]string `json:"invalidMinPropertiesField"`

	// +kubebuilder:object:root=true // want `marker "kubebuilder:object:root" can only be applied to type definitions, not fields`
	InvalidRootField string `json:"invalidRootField"`

	// +kubebuilder:subresource:status // want `marker "kubebuilder:subresource:status" can only be applied to type definitions, not fields`
	InvalidStatusField string `json:"invalidStatusField"`

	// +kubebuilder:validation:MaxProperties=10 // want `marker "kubebuilder:validation:MaxProperties" can only be applied to type definitions, not fields`
	InvalidMaxPropertiesField map[string]string `json:"invalidMaxPropertiesField"`
}

// Test markers that can be on both fields and types
// +kubebuilder:default="default-value"
// +kubebuilder:validation:MinLength=1
// +kubebuilder:validation:MaxLength=100
type ValidBothMarkers struct {
	// +kubebuilder:default="field-default"
	// +kubebuilder:validation:MinLength=5
	// +kubebuilder:validation:MaxLength=50
	Name string `json:"name"`
}

// Test array field markers
type ArrayFieldTest struct {
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=10
	// +kubebuilder:validation:UniqueItems=true
	// +listType=map
	// +listMapKey=name
	ValidArrayField []Item `json:"validArrayField"`
}

type Item struct {
	Name string `json:"name"`
}

// Test custom enum and format markers
// +kubebuilder:validation:Enum=TypeA;TypeB;TypeC
// +kubebuilder:validation:Format=email
type EnumType string

type EnumFieldTest struct {
	// +kubebuilder:validation:Enum=FieldA;FieldB;FieldC
	// +kubebuilder:validation:Format=ipv4
	EnumField string `json:"enumField"`
}
