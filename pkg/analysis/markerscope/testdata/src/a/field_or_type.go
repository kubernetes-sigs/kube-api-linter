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

// FieldOrType markers that can be applied to both fields and types.
//
// Note: The Type marker (kubebuilder:validation:Type) is tested separately in type_override.go.
// This file focuses on testing that FieldOrType markers (Enum, Format, PreserveUnknownFields, Title, XValidation)
// have no type constraints, so they do not produce errors, warnings, or suggested
// fixes regardless of the underlying type they are applied to.

// Enum markers
type GeneralType string

// Valid: Enum on type
// +kubebuilder:validation:Enum=A;B;C
type EnumType string

// Format markers
// Valid: Format on type
// +kubebuilder:validation:Format=email
type FormatType string

// PreserveUnknownFields markers
// Valid: PreserveUnknownFields on type
// +kubebuilder:pruning:PreserveUnknownFields
type ValidPreserveUnknownFieldsType struct {
	Name string `json:"name"`
}

// Title markers
// Valid: title on type
// +kubebuilder:title="My Title"
type ValidTitleType struct {
	Name string `json:"name"`
}

// XValidation markers
// Valid: XValidation on type
// +kubebuilder:validation:XValidation:rule="self.size() > 0"
type XValidationType string

// Valid: All [Field, Type] markers on type
// +kubebuilder:validation:Enum=A;B;C
// +kubebuilder:validation:Format=email
// +kubebuilder:pruning:PreserveUnknownFields
// +kubebuilder:title="Combined Title"
// +kubebuilder:validation:XValidation:rule="self.size() > 0"
type ValidAllFieldOrTypeType struct {
	Name string `json:"name"`
}

// Types without markers for testing field markers
type NoMarkerAllFieldOrTypeType struct {
	Name string `json:"name"`
}

type FieldOrTypeMarkersTest struct {
	// Valid: Enum marker on string field
	// +kubebuilder:validation:Enum=A;B;C
	ValidEnum string `json:"validEnum"`

	// Valid: Format marker on string field
	// +kubebuilder:validation:Format=email
	ValidFormat string `json:"validFormat"`

	// Valid: PreserveUnknownFields on field
	// +kubebuilder:pruning:PreserveUnknownFields
	ValidPreserveUnknownFields map[string]string `json:"validPreserveUnknownFields"`

	// Valid: title on field
	// +kubebuilder:title="Field Title"
	ValidTitle map[string]string `json:"validTitle"`

	// Valid: XValidation marker on string field
	// +kubebuilder:validation:XValidation:rule="self.size() > 0"
	ValidXValidation string `json:"validXValidation"`

	// Valid: All [Field, Type] markers on field
	// +kubebuilder:validation:Enum=A;B;C
	// +kubebuilder:validation:Format=email
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:title="Combined Field Title"
	// +kubebuilder:validation:XValidation:rule="self.size() > 0"
	ValidAllFieldOrTypeMarkers map[string]interface{} `json:"validAllFieldOrTypeMarkers"`

	// Valid: Using EnumType
	ValidEnumTyped EnumType `json:"validEnumTyped"`

	// Valid: Using FormatType
	ValidFormatTyped FormatType `json:"validFormatTyped"`

	// Valid: Using type with [Field, Type] markers
	ValidPreserveUnknownFieldsTyped ValidPreserveUnknownFieldsType `json:"validPreserveUnknownFieldsTyped"`

	// Valid: Using type with title marker
	ValidTitleTyped ValidTitleType `json:"validTitleTyped"`

	// Valid: Using XValidationType
	ValidXValidationTyped XValidationType `json:"validXValidationTyped"`

	// Valid: Using type with all [Field, Type] markers
	ValidAllFieldOrTypeTyped ValidAllFieldOrTypeType `json:"validAllFieldOrTypeTyped"`

	// Valid: Using named type with general markers
	ValidGeneralTyped GeneralType `json:"validGeneralTyped"`

	// Valid: General markers can be applied to any type
	// +kubebuilder:validation:Enum=1;2;3
	ValidEnumOnInt int32 `json:"validEnumOnInt"`

	// Valid: Format can be applied to any type
	// +kubebuilder:validation:Format=int32
	ValidFormatOnInt int32 `json:"validFormatOnInt"`

	// Valid: Enum marker on named type
	// +kubebuilder:validation:Enum=A;B;C
	ValidEnumOnGeneralType GeneralType `json:"validEnumOnGeneralType"`

	// Valid: Format marker on named type
	// +kubebuilder:validation:Format=email
	ValidFormatOnGeneralType GeneralType `json:"validFormatOnGeneralType"`

	// Valid: XValidation marker on named type (allowed on both field and type)
	// +kubebuilder:validation:XValidation:rule="self.size() > 0"
	ValidXValidationOnGeneralType GeneralType `json:"validXValidationOnGeneralType"`

	// Valid: Using type with no [Field, Type] markers
	// +kubebuilder:validation:Enum=A;B;C
	// +kubebuilder:validation:Format=email
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:title="Field Combined Title"
	// +kubebuilder:validation:XValidation:rule="self.size() > 0"
	NoMarkerAllFieldOrTypeType NoMarkerAllFieldOrTypeType `json:"noMarkerAllFieldOrTypeType"`
}
