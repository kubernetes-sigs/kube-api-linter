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

// Valid: PreserveUnknownFields on type
// +kubebuilder:pruning:PreserveUnknownFields
type ValidPreserveUnknownFieldsType struct {
	Name string `json:"name"`
}

// Valid: title on type
// +kubebuilder:title="My Title"
type ValidTitleType struct {
	Name string `json:"name"`
}

// Valid: All [Field, Type] markers on type
// +kubebuilder:pruning:PreserveUnknownFields
// +kubebuilder:title="Combined Title"
type ValidAllFieldOrTypeType struct {
	Name string `json:"name"`
}

type NoMarkerAllFieldOrTypeType struct {
	Name string `json:"name"`
}

type FieldOrTypeOnFieldTest struct {
	// Valid: PreserveUnknownFields on field
	// +kubebuilder:pruning:PreserveUnknownFields
	ValidPreserveUnknownFields map[string]string `json:"validPreserveUnknownFields"`

	// Valid: title on field
	// +kubebuilder:title="Field Title"
	ValidTitle map[string]string `json:"validTitle"`

	// Valid: All [Field, Type] markers on field
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:title="Combined Field Title"
	ValidAllFieldOrTypeMarkers map[string]interface{} `json:"validAllFieldOrTypeMarkers"`

	// Valid: Using type with [Field, Type] markers
	ValidPreserveUnknownFieldsTyped ValidPreserveUnknownFieldsType `json:"validPreserveUnknownFieldsTyped"`

	// Valid: Using type with title marker
	ValidTitleTyped ValidTitleType `json:"validTitleTyped"`

	// Valid: Using type with all [Field, Type] markers
	ValidAllFieldOrTypeTyped ValidAllFieldOrTypeType `json:"validAllFieldOrTypeTyped"`

	// Valid: Using type with no [Field, Type] markers
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:title="Field Combined Title"
	NoMarkerAllFieldOrTypeType NoMarkerAllFieldOrTypeType `json:"noMarkerAllFieldOrTypeType"`
}
