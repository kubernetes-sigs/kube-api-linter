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

type GeneralType string

// Valid: Enum on type
// +kubebuilder:validation:Enum=A;B;C
type EnumType string

// Valid: Format on type
// +kubebuilder:validation:Format=email
type FormatType string

// Valid: Type on type
// +kubebuilder:validation:Type=string
type TypeType string

// Valid: XValidation on type
// +kubebuilder:validation:XValidation:rule="self.size() > 0"
type XValidationType string

type GeneralMarkersFieldTest struct {
	// Valid: Enum marker on string field
	// +kubebuilder:validation:Enum=A;B;C
	ValidEnum string `json:"validEnum"`

	// Valid: Format marker on string field
	// +kubebuilder:validation:Format=email
	ValidFormat string `json:"validFormat"`

	// Valid: Type marker on string field
	// +kubebuilder:validation:Type=string
	ValidType string `json:"validType"`

	// Valid: XValidation marker on string field
	// +kubebuilder:validation:XValidation:rule="self.size() > 0"
	ValidXValidation string `json:"validXValidation"`

	// Valid: All general markers on string field
	// +kubebuilder:validation:Enum=A;B;C
	// +kubebuilder:validation:Format=email
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:XValidation:rule="self.size() > 0"
	ValidAllGeneralMarkers string `json:"validAllGeneralMarkers"`

	// Valid: Using named type with general markers
	ValidGeneralTyped GeneralType `json:"validGeneralTyped"`

	// Valid: Using EnumType
	ValidEnumTyped EnumType `json:"validEnumTyped"`

	// Valid: Using FormatType
	ValidFormatTyped FormatType `json:"validFormatTyped"`

	// Valid: Using TypeType
	ValidTypeTyped TypeType `json:"validTypeTyped"`

	// Valid: Using XValidationType
	ValidXValidationTyped XValidationType `json:"validXValidationTyped"`

	// Valid: General markers can be applied to any type
	// +kubebuilder:validation:Enum=1;2;3
	ValidEnumOnInt int32 `json:"validEnumOnInt"`

	// Valid: Format can be applied to any type
	// +kubebuilder:validation:Format=int32
	ValidFormatOnInt int32 `json:"validFormatOnInt"`

	// Invalid: Enum marker on named type
	// +kubebuilder:validation:Enum=A;B;C
	InvalidEnumOnGeneralType GeneralType `json:"invalidEnumOnGeneralType"`

	// Invalid: Format marker on named type
	// +kubebuilder:validation:Format=email
	InvalidFormatOnGeneralType GeneralType `json:"invalidFormatOnGeneralType"`

	// Invalid: Type marker on named type
	// +kubebuilder:validation:Type=string
	InvalidTypeOnGeneralType GeneralType `json:"invalidTypeOnGeneralType"`

	// Valid: XValidation marker on named type (allowed on both field and type)
	// +kubebuilder:validation:XValidation:rule="self.size() > 0"
	ValidXValidationOnGeneralType GeneralType `json:"validXValidationOnGeneralType"`
}
