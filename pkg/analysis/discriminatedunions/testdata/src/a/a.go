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

// +union
type ValidLegacyUnion struct {
	// +unionDiscriminator
	// +required
	Type string `json:"type"`

	// +unionMember
	// +optional
	Foo *string `json:"foo,omitempty"`
}

// +union
type ValidLegacyUnionOptionalMember struct {
	// +unionDiscriminator
	// +required
	Type string `json:"type"`

	// +unionMember,optional
	Foo *string `json:"foo,omitempty"`
}

// +union
type ValidDeclarativeUnion struct {
	// +k8s:unionDiscriminator
	// +k8s:required
	Type string `json:"type"`

	// +k8s:unionMember
	// +k8s:optional
	Foo *string `json:"foo,omitempty"`
}

// +union
type MissingDiscriminatorUnion struct { // want "type MissingDiscriminatorUnion is marked as a discriminated union but has no discriminator field"
	// +unionMember
	// +optional
	Foo *string `json:"foo,omitempty"`
}

// +union
type MultipleDiscriminatorUnion struct { // want "type MultipleDiscriminatorUnion is marked as a discriminated union but has 2 discriminator fields"
	// +unionDiscriminator
	// +required
	One string `json:"one"`

	// +unionDiscriminator
	// +required
	Two string `json:"two"`

	// +unionMember
	// +optional
	Foo *string `json:"foo,omitempty"`
}

// +union
type DiscriminatorNotRequiredUnion struct {
	// +unionDiscriminator
	Kind string `json:"kind"` // want "discriminator field DiscriminatorNotRequiredUnion.Kind must be marked as required"

	// +unionMember
	// +optional
	Foo *string `json:"foo,omitempty"`
}

// +union
type MemberNotOptionalUnion struct {
	// +unionDiscriminator
	// +required
	Kind string `json:"kind"`

	// +unionMember
	Foo *string `json:"foo,omitempty"` // want "union member field MemberNotOptionalUnion.Foo must be marked as optional"
}

// +union
type NonMemberFieldUnion struct {
	// +unionDiscriminator
	// +required
	Kind string `json:"kind"`

	// +unionMember
	// +optional
	Foo *string `json:"foo,omitempty"`

	Other string `json:"other,omitempty"` // want "field NonMemberFieldUnion.Other is not a union discriminator/member in union type NonMemberFieldUnion"
}

type UnionWithoutTypeMarkerWithK8sMarkers struct {
	// +k8s:unionDiscriminator
	// +k8s:required
	Kind string `json:"kind"`

	// +k8s:unionMember
	Foo *string `json:"foo,omitempty"` // want "union member field UnionWithoutTypeMarkerWithK8sMarkers.Foo must be marked as optional"
}
