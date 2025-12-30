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

// TypeOverrideTest tests that Type marker can override the schema type
// for type constraint validation in real-world scenarios
type TypeOverrideTest struct {
	// Valid: []byte (treated as string by default) with MinLength
	// This is a common case - []byte is represented as base64 string in OpenAPI
	// +kubebuilder:validation:MinLength=1
	ValidByteSliceMinLength []byte `json:"validByteSliceMinLength"`

	// Valid: []byte explicitly marked as string with MinLength
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:MinLength=1
	ValidByteSliceAsString []byte `json:"validByteSliceAsString"`

	// Valid: []byte overridden as array allows MinItems (array marker)
	// This allows treating []byte as actual byte array instead of base64 string
	// +kubebuilder:validation:Type=array
	// +kubebuilder:validation:MinItems=1
	ValidByteSliceAsArray []byte `json:"validByteSliceAsArray"`

	// Invalid: []byte without Type=array override cannot use MinItems (array marker)
	// Note: []byte is treated as string by default, so MinItems is not allowed
	// +kubebuilder:validation:MinItems=1 // want `marker "kubebuilder:validation:MinItems": type string is not allowed \(expected one of: \[array\]\)`
	InvalidByteSliceMinItems []byte `json:"invalidByteSliceMinItems"`

	// Real-world case: IntOrString-like type that can be either int or string
	// Valid: Custom type overridden as string to allow string validation
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern="^[0-9]+$"
	ValidIntOrStringAsString int32 `json:"validIntOrStringAsString"`

	// Real-world case: Port number stored as string but validated as integer
	// Valid: string overridden as integer to allow numeric validation
	// +kubebuilder:validation:Type=integer
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	ValidPortAsInteger string `json:"validPortAsInteger"`
}

// Real-world type override on type definition
// IntOrString represents a value that can be an int or string (like k8s.io/apimachinery/pkg/util/intstr.IntOrString)
// Valid: Type override on custom type to allow string validation
// +kubebuilder:validation:Type=string
// +kubebuilder:validation:Pattern="^[0-9]+%?$"
type IntOrString int32

type TypeOverrideOnTypeUsageTest struct {
	// Valid: Using IntOrString type with Type override
	ValidIntOrString IntOrString `json:"validIntOrString"`

	// Invalid: Cannot apply string marker to IntOrString field (underlying type is int32)
	// The Type=string override is on the type definition, not visible when checking field constraints
	// +kubebuilder:validation:MaxLength=10 // want `marker "kubebuilder:validation:MaxLength": type integer is not allowed \(expected one of: \[string\]\)`
	InvalidMaxLengthOnIntOrString IntOrString `json:"invalidMaxLengthOnIntOrString"`
}
