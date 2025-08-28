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
requiredfields is a linter to check that fields that are marked as required are marshalled properly.
The linter will check for fields that are marked as required using the +required marker, or the +kubebuilder:validation:Required marker.

Required fields should have omitempty tags to prevent "mess" in the encoded object. Fields are not pointers in general.

Where the zero value for a field is not a valid value, the field does not need to be a pointer as the zero value could never be admitted.
Where the zero value for a field is a valid value (e.g. the empty string, or 0), the field should be a pointer to distinguish between unset and zero value states.

Required fields should always have omitempty tags to prevent "mess" in the encoded object, regardless of whether they are pointers.
*/
package requiredfields
