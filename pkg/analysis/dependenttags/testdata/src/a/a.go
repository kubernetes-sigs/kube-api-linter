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

type Union struct {
	// +k8s:unionMember
	InvalidMember string // want "field Union.InvalidMember with marker \\+k8s:unionMember is missing required marker\\(s\\): \\+k8s:optional"

	// +k8s:unionMember
	// +k8s:optional
	ValidMember string
}

type List struct {
	// +listType
	InvalidList []string // want "field List.InvalidList with marker \\+listType is missing required marker\\(s\\): \\+k8s:listType"

	// +listType
	// +k8s:listType
	ValidList []string
}

type AnyOf struct {
	// +example:any
	InvalidAny string // want "field AnyOf.InvalidAny with marker \\+example:any requires at least one of the following markers, but none were found: \\+dep1, \\+dep2"

	// +example:any
	// +dep1
	ValidAny1 string

	// +example:any
	// +dep2
	ValidAny2 string
}

// +dep1
type MyString string

type TypeAlias struct {
	// +example:any
	InvalidAlias string // want "field TypeAlias.InvalidAlias with marker \\+example:any requires at least one of the following markers, but none were found: \\+dep1, \\+dep2"

	// +example:any
	ValidAlias MyString
}

type ListMap struct {
	// +listType=map
	InvalidListMap []string // want "field ListMap.InvalidListMap with marker \\+listType=map is missing required marker\\(s\\): \\+listMapKey"

	// +listType=map
	// +listMapKey
	ValidListMap []string
}

type Monotonic struct {
	// FAIL: missing both k8s:minimum and the any-dep
	// +k8s:monotonic
	MissingAll int64 // want "field Monotonic.MissingAll with marker \\+k8s:monotonic is missing required marker\\(s\\): \\+k8s:minimum" "field Monotonic.MissingAll with marker \\+k8s:monotonic requires at least one of the following markers, but none were found: \\+k8s:required, \\+k8s:update=NoUnset"

	// FAIL: has k8s:minimum but wrong update value — any-dep is absent
	// +k8s:monotonic
	// +k8s:minimum=0
	// +k8s:update=SomethingElse
	WrongUpdateValue *int64 // want "field Monotonic.WrongUpdateValue with marker \\+k8s:monotonic requires at least one of the following markers, but none were found: \\+k8s:required, \\+k8s:update=NoUnset"

	// PASS: k8s:minimum present, k8s:required satisfies any-dep
	// +k8s:monotonic
	// +k8s:minimum=0
	// +k8s:required
	ValidWithRequired int64

	// PASS: k8s:minimum present, k8s:update=NoUnset satisfies any-dep
	// +k8s:monotonic
	// +k8s:minimum=0
	// +k8s:update=NoUnset
	ValidWithUpdateNoUnset *int64

	// FAIL: bare k8s:update should NOT satisfy k8s:update=NoUnset
	// +k8s:monotonic
	// +k8s:minimum=0
	// +k8s:update
	WrongUpdateNoValue *int64 // want "field Monotonic.WrongUpdateNoValue with marker \\+k8s:monotonic requires at least one of the following markers, but none were found: \\+k8s:required, \\+k8s:update=NoUnset"
}
