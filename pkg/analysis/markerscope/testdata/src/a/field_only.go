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

// Invalid: Field-only markers on type
// +optional // want `marker "optional" can only be applied to fields`
// +required // want `marker "required" can only be applied to fields`
// +nullable // want `marker "nullable" can only be applied to fields`
type InvalidFieldOnlyOnType struct {
	Name string `json:"name"`
}

// Invalid: kubebuilder:default on type
// +kubebuilder:default="default" // want `marker "kubebuilder:default" can only be applied to fields`
type InvalidDefaultOnType string

// Invalid: kubebuilder:example on type
// +kubebuilder:example="example" // want `marker "kubebuilder:example" can only be applied to fields`
type InvalidExampleOnType string

// Invalid: kubebuilder:validation:EmbeddedResource on type
// +kubebuilder:validation:EmbeddedResource // want `marker "kubebuilder:validation:EmbeddedResource" can only be applied to fields`
type InvalidEmbeddedResourceOnType struct {
	Data map[string]interface{} `json:"data"`
}

// Invalid: kubebuilder:validation:Schemaless on type
// +kubebuilder:validation:Schemaless // want `marker "kubebuilder:validation:Schemaless" can only be applied to fields`
type InvalidSchemalessOnType struct {
	Data map[string]interface{} `json:"data"`
}

type FieldOnlyMarkersTest struct {
	// Valid: optional marker
	// +optional
	ValidOptional string `json:"validOptional,omitempty"`

	// Valid: required marker
	// +required
	ValidRequired string `json:"validRequired"`

	// Valid: k8s:optional marker
	// +k8s:optional
	ValidK8sOptional string `json:"validK8sOptional,omitempty"`

	// Valid: k8s:required marker
	// +k8s:required
	ValidK8sRequired string `json:"validK8sRequired"`

	// Valid: nullable marker
	// +nullable
	ValidNullable *string `json:"validNullable"`

	// Valid: kubebuilder:default marker
	// +kubebuilder:default="default"
	ValidDefault string `json:"validDefault"`

	// Valid: kubebuilder:example marker
	// +kubebuilder:example="example"
	ValidExample string `json:"validExample"`

	// Valid: kubebuilder:validation:EmbeddedResource marker
	// +kubebuilder:validation:EmbeddedResource
	ValidEmbeddedResource map[string]interface{} `json:"validEmbeddedResource"`

	// Valid: kubebuilder:validation:Schemaless marker
	// +kubebuilder:validation:Schemaless
	ValidSchemaless map[string]interface{} `json:"validSchemaless"`

	// Valid: All field-only markers combined
	// +optional
	// +nullable
	// +kubebuilder:default="combined"
	// +kubebuilder:example="combined-example"
	ValidAllFieldOnlyMarkers *string `json:"validAllFieldOnlyMarkers,omitempty"`
}
