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

/*
duplicatemarkers is an analyzer that checks for duplicate markers in the API types.

It reports only if the marker and value together are completely unique now. For example, the `duplicatemarkers` will not diagnose the field has kubebuilder:validation:MaxLength=10 and kubebuilder:validation:MaxLength=11

Example:

	type Foo struct {
		// +kubebuilder:validation:MaxLength=10
		// +kubebuilder:validation:MaxLength=11
		Field string `json:"field"`
	}

	// +kubebuilder:validation:MaxLength=10
	// +kubebuilder:validation:MaxLength=11
	type Bar string

About the duplicated markers which has different value, it requires specific rule for each marker, these are processed by its corresponding linter.
*/
package duplicatemarkers
