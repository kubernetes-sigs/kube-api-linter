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
maxlength is an analyzer that checks that all string fields have a maximum length,
all array fields have a maximum number of items, and all map fields have a maximum
number of properties.

String fields that are not otherwise bound in length, through being an enum or
formatted in a certain way, should have a maximum length.
This ensures that CEL validations on the field are not overly costly in terms of
time and memory.

Array fields should have a maximum number of items.
This ensures that any CEL validations on the field are not overly costly in terms
of time and memory. Where arrays are used to represent a list of structures, CEL
rules may exist within the array. Limiting the array length ensures the cardinality
of the rules within the array is not unbounded.

Map fields (with string keys) should have a maximum number of properties.
This ensures that CEL validations on the map are not overly costly.

For strings, the maximum length can be set using the `kubebuilder:validation:MaxLength`
or `k8s:maxLength` tag.

For byte slices ([]byte), the maximum length can be set using the
`kubebuilder:validation:MaxLength` or `k8s:maxBytes` tag. Note that in k8s declarative
validation, `k8s:maxBytes` constrains the byte count whereas `k8s:maxLength` constrains
the Unicode character count — `k8s:maxBytes` is the semantically correct tag for []byte.

For arrays, the maximum number of items can be set using the `kubebuilder:validation:MaxItems`
or `k8s:maxItems` tag.

For arrays of strings, the maximum length of each string can be set using the
`kubebuilder:validation:items:MaxLength` tag on the array field itself.
Or, if the array uses a string type alias, the `kubebuilder:validation:MaxLength` tag can
be used on the alias.

For maps with string keys, the maximum number of properties can be set using the
`kubebuilder:validation:MaxProperties` or `k8s:maxProperties` tag.

The preferred marker cited in diagnostics is configurable via MaxLengthConfig.
*/
package maxlength
