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

import "time"

// RawMessage is similar to json.RawMessage (type RawMessage []byte)
// Used to store arbitrary JSON configuration data.
type RawMessage []byte

// Duration is similar to metav1.Duration which wraps time.Duration
// It is serialized as a string (e.g., "1m", "30s").
type Duration struct {
	time.Duration
}

// TypeOverrideValidTest tests valid cases where Type marker overrides the schema type
type TypeOverrideValidTest struct {
	// Example 1: RawMessage ([]byte alias) with Type=object
	// RawMessage stores arbitrary JSON configuration data.
	// Since it doesn't have a known schema, use Type=object to override.
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Type=object
	// +optional
	Config RawMessage `json:"config,omitempty"`

	// Example 2: Duration with Type=string
	// Duration is serialized as a string (e.g., "1m", "30s").
	// Use Type=string to enable string validation markers like Pattern.
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern="^([0-9]+(\\.[0-9]+)?(ms|s|m|h))+$"
	// +kubebuilder:default:="1m"
	// +optional
	Interval *Duration `json:"interval,omitempty"`
}

// TypeOverrideInvalidTest tests invalid cases without proper Type override
type TypeOverrideInvalidTest struct {
	// Invalid: RawMessage without Type=object override cannot use object markers
	// RawMessage is []byte which is treated as string by default
	// +kubebuilder:validation:MinProperties=1 // want `marker "kubebuilder:validation:MinProperties": type string is not allowed \(expected one of: \[object\]\)`
	// +optional
	InvalidConfig RawMessage `json:"invalidConfig,omitempty"`

	// Invalid: Duration without Type=string override cannot use string markers
	// Duration is a struct type which is treated as object by default
	// +kubebuilder:validation:Pattern="^([0-9]+(\\.[0-9]+)?(ms|s|m|h))+$" // want `marker "kubebuilder:validation:Pattern": type object is not allowed \(expected one of: \[string\]\)`
	// +optional
	InvalidInterval *Duration `json:"invalidInterval,omitempty"`
}
