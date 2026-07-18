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
package maxlength

// MaxLengthConfig contains configuration for the maxlength linter.
type MaxLengthConfig struct {
	// PreferredMaxLengthMarker is the preferred marker to cite in diagnostics for
	// string maximum-length violations.
	// If not set, the default value is "kubebuilder:validation:MaxLength".
	// Valid values are "kubebuilder:validation:MaxLength" and "k8s:maxLength".
	PreferredMaxLengthMarker string `json:"preferredMaxLengthMarker"`

	// PreferredMaxItemsMarker is the preferred marker to cite in diagnostics for
	// array maximum-items violations.
	// If not set, the default value is "kubebuilder:validation:MaxItems".
	// Valid values are "kubebuilder:validation:MaxItems" and "k8s:maxItems".
	PreferredMaxItemsMarker string `json:"preferredMaxItemsMarker"`

	// PreferredMaxPropertiesMarker is the preferred marker to cite in diagnostics for
	// map maximum-properties violations.
	// If not set, the default value is "kubebuilder:validation:MaxProperties".
	// Valid values are "kubebuilder:validation:MaxProperties" and "k8s:maxProperties".
	PreferredMaxPropertiesMarker string `json:"preferredMaxPropertiesMarker"`
}
