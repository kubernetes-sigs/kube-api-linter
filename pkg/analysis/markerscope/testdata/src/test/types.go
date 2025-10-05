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
package test

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +kubebuilder:object:root=true

// TestResource is the Schema for the testresources API
type TestResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TestResourceSpec   `json:"spec,omitempty"`
	Status TestResourceStatus `json:"status,omitempty"`
}

// TestResourceSpec defines the desired state of TestResource
type TestResourceSpec struct {
	// Valid: MinProperties on map field
	// +kubebuilder:validation:MinProperties=1
	ConfigMap map[string]string `json:"configMap,omitempty"`

	// Valid: MinProperties on struct field
	Settings []SettingsStruct `json:"settings,omitempty"`

	// Invalid: MinProperties on slice field (should cause error)
	Items []string `json:"items,omitempty"`

	// Invalid: MinProperties on string field (should cause error)
	Name string `json:"name,omitempty"`
}

// +kubebuilder:validation:MinProperties=2
type SettingsStruct struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// TestResourceStatus defines the observed state of TestResource
type TestResourceStatus struct {
	Ready bool `json:"ready,omitempty"`
}

// +kubebuilder:object:root=true

// TestResourceList contains a list of TestResource
type TestResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TestResource `json:"items"`
}
