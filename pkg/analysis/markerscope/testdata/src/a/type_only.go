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

// Valid: ExactlyOneOf on type
// +kubebuilder:validation:items:ExactlyOneOf={Field1,Field2}
type ValidExactlyOneOfType struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
}

// Valid: AtMostOneOf on type
// +kubebuilder:validation:items:AtMostOneOf={Field1,Field2}
type ValidAtMostOneOfType struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
}

// Valid: AtLeastOneOf on type
// +kubebuilder:validation:items:AtLeastOneOf={Field1,Field2}
type ValidAtLeastOneOfType struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
}

// Valid: All type-only markers combined
// +kubebuilder:validation:items:ExactlyOneOf={Field1,Field2}
// +kubebuilder:validation:items:AtMostOneOf={Field3,Field4}
// +kubebuilder:validation:items:AtLeastOneOf={Field5,Field6}
type ValidAllTypeOnlyMarkers struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
	Field3 string `json:"field3"`
	Field4 string `json:"field4"`
	Field5 string `json:"field5"`
	Field6 string `json:"field6"`
}

type TypeOnlyMarkersTest struct {
	// Invalid: ExactlyOneOf on field
	// +kubebuilder:validation:items:ExactlyOneOf={field1,field2} // want `marker "kubebuilder:validation:items:ExactlyOneOf" can only be applied to types`
	InvalidExactlyOneOfOnField string `json:"invalidExactlyOneOfOnField"`

	// Invalid: AtMostOneOf on field
	// +kubebuilder:validation:items:AtMostOneOf={field1,field2} // want `marker "kubebuilder:validation:items:AtMostOneOf" can only be applied to types`
	InvalidAtMostOneOfOnField string `json:"invalidAtMostOneOfOnField"`

	// Invalid: AtLeastOneOf on field
	// +kubebuilder:validation:items:AtLeastOneOf={field1,field2} // want `marker "kubebuilder:validation:items:AtLeastOneOf" can only be applied to types`
	InvalidAtLeastOneOfOnField string `json:"invalidAtLeastOneOfOnField"`

	// Invalid: All type-only markers on field
	// +kubebuilder:validation:items:ExactlyOneOf={field1,field2} // want `marker "kubebuilder:validation:items:ExactlyOneOf" can only be applied to types`
	// +kubebuilder:validation:items:AtMostOneOf={field1,field2} // want `marker "kubebuilder:validation:items:AtMostOneOf" can only be applied to types`
	// +kubebuilder:validation:items:AtLeastOneOf={field1,field2} // want `marker "kubebuilder:validation:items:AtLeastOneOf" can only be applied to types`
	InvalidAllTypeOnlyOnField string `json:"invalidAllTypeOnlyOnField"`

	// Valid: Using type with type-only markers
	ValidExactlyOneOf ValidExactlyOneOfType `json:"validExactlyOneOf"`

	// Valid: Using type with type-only markers
	ValidAtMostOneOf ValidAtMostOneOfType `json:"validAtMostOneOf"`

	// Valid: Using type with type-only markers
	ValidAtLeastOneOf ValidAtLeastOneOfType `json:"validAtLeastOneOf"`

	// Valid: Using type with all type-only markers
	ValidAllTypeOnly ValidAllTypeOnlyMarkers `json:"validAllTypeOnly"`
}
