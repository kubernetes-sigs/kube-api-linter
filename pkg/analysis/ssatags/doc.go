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

// Package ssatags provides an analyzer to enforce proper Server-Side Apply (SSA) tags on array fields.
//
// This analyzer ensures that all array fields in Kubernetes API objects have the appropriate
// listType markers (atomic, set, or map) for proper Server-Side Apply behavior.
//
// The analyzer checks for:
//
// 1. Missing listType markers on array fields
// 2. Invalid listType values (must be atomic, set, or map)
// 3. Missing listMapKey markers for listType=map arrays
//
// Example usage:
//
//	// +kubebuilder:listType=atomic
//	Items []string `json:"items,omitempty"`
//
//	// +kubebuilder:listType=map
//	// +kubebuilder:listMapKey=name
//	NamedItems []NamedItem `json:"namedItems,omitempty"`
package ssatags
