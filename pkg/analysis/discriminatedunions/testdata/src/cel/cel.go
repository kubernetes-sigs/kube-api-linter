/*
Copyright 2026 The Kubernetes Authors.

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

package cel

// +kubebuilder:object:root=true
type Root struct {
	Valid   *Valid                `json:"valid"`
	Missing *Missing              `json:"missing"`
	Slice   []SliceMember         `json:"slice"`
	Array   [1]ArrayMember        `json:"array"`
	Map     map[string]*MapMember `json:"map"`
	Alias   Alias                 `json:"alias"`
	Cycle   *Cycle                `json:"cycle"`
	Inline  struct {
		Nested *InlineMember `json:"nested"`
	} `json:"inline"`
	Named           NamedUnions            `json:"named"`
	EmptyName       EmptyMemberName        `json:"emptyName"`
	WrongEmptyName  WrongEmptyMemberName   `json:"wrongEmptyName"`
	Reversed        Reversed               `json:"reversed"`
	Malformed       Malformed              `json:"malformed"`
	Invalid         InvalidDiscriminator   `json:"invalid"`
	Multiple        MultipleDiscriminators `json:"multiple"`
	NotRequired     NotRequired            `json:"notRequired"`
	NotOptional     NotOptional            `json:"notOptional"`
	Empty           DiscriminatorOnly      `json:"empty"`
	Undiscriminated Undiscriminated        `json:"undiscriminated"`
	// +kubebuilder:validation:Schemaless
	Schemaless *Unreachable `json:"schemaless"`
	Ignored    *Unreachable `json:"-"`
	private    *Unreachable
}

// +union
// +kubebuilder:validation:XValidation:rule="self.kind != 'Other'",message="unrelated rule"
// +kubebuilder:validation:XValidation:rule="(has(self.kind) && self.kind == 'Payload') ? has(self.renamed) : !has(self.renamed)"
// +kubebuilder:validation:XValidation:rule="has(self.kind) && self.kind == \"Maybe\" ? true : !has(self.maybe)"
type Valid struct {
	// +unionDiscriminator
	// +required
	Type string `json:"kind"`
	// +unionMember
	// +optional
	Payload *string `json:"renamed,omitempty"`
	// +unionMember,optional
	// +optional
	Maybe *string `json:"maybe,omitempty"`
}

// +union
type Missing struct {
	// +unionDiscriminator
	// +required
	Kind string `json:"kind"`
	// +unionMember
	// +optional
	Foo *string `json:"foo,omitempty"` // want "union member field Missing.Foo.*expected has\\(self.kind\\).*self.kind == \"Foo\""
}

// +union
type SliceMember struct {
	// +unionDiscriminator
	// +required
	Kind string `json:"kind"`
	// +unionMember
	// +optional
	Foo *string `json:"foo,omitempty"` // want "union member field SliceMember.Foo.*requires a recognized"
}

// +union
type ArrayMember struct {
	// +unionDiscriminator
	// +required
	Kind string `json:"kind"`
	// +unionMember
	// +optional
	Foo *string `json:"foo,omitempty"` // want "union member field ArrayMember.Foo.*requires a recognized"
}

// +union
type MapMember struct {
	// +unionDiscriminator
	// +required
	Kind string `json:"kind"`
	// +unionMember
	// +optional
	Foo *string `json:"foo,omitempty"` // want "union member field MapMember.Foo.*requires a recognized"
}

type Alias = NamedWrapper
type NamedWrapper []*AliasMember

// +union
type AliasMember struct {
	// +unionDiscriminator
	// +required
	Kind string `json:"kind"`
	// +unionMember
	// +optional
	Foo *string `json:"foo,omitempty"` // want "union member field AliasMember.Foo.*requires a recognized"
}

type Cycle struct {
	Next  *Cycle   `json:"next"`
	Union *Missing `json:"union"`
}

// +union
type InlineMember struct {
	// +unionDiscriminator
	// +required
	Kind string `json:"kind"`
	// +unionMember
	// +optional
	Foo *string `json:"foo,omitempty"` // want "union member field InlineMember.Foo.*requires a recognized"
}

// +kubebuilder:validation:XValidation:rule="has(self.kind) && self.kind == 'HTTP' ? has(self.transport) : !has(self.transport)"
// +kubebuilder:validation:XValidation:rule="has(self.__namespace__) && self.__namespace__ == 'Bearer' ? has(self.access__dash__token) : !has(self.access__dash__token)"
type NamedUnions struct {
	// +k8s:unionDiscriminator(union: "transport")
	// +k8s:required
	Kind string `json:"kind"`
	// +k8s:unionMember(union: "transport", memberName: "HTTP")
	// +k8s:optional
	Transport *string `json:"transport,omitempty"`
	// +k8s:unionDiscriminator(union: "auth")
	// +k8s:required
	Namespace string `json:"namespace"`
	// +k8s:unionMember(union: "auth", memberName: "Bearer")
	// +k8s:optional
	Token *string `json:"access-token,omitempty"`
}

// +kubebuilder:validation:XValidation:rule="has(self.kind) && self.kind == \"\" ? has(self.foo) : !has(self.foo)"
type EmptyMemberName struct {
	// +k8s:unionDiscriminator
	// +k8s:required
	Kind string `json:"kind"`
	// +k8s:unionMember(memberName: "")
	// +k8s:optional
	Foo *string `json:"foo,omitempty"`
}

// +kubebuilder:validation:XValidation:rule="has(self.kind) && self.kind == 'Foo' ? has(self.foo) : !has(self.foo)"
type WrongEmptyMemberName struct {
	// +k8s:unionDiscriminator
	// +k8s:required
	Kind string `json:"kind"`
	// +k8s:unionMember(memberName: "")
	// +k8s:optional
	Foo *string `json:"foo,omitempty"` // want "union member field WrongEmptyMemberName.Foo.*expected has\\(self.kind\\).*self.kind == \"\""
}

// +union
// +kubebuilder:validation:XValidation:rule="has(self.kind) && self.kind == 'Foo' ? !has(self.foo) : has(self.foo)"
// +kubebuilder:validation:XValidation:rule="has(self.kind) && self.kind == 'Maybe' ? !has(self.maybe) : true"
type Reversed struct {
	// +unionDiscriminator
	// +required
	Kind string `json:"kind"`
	// +unionMember
	// +optional
	Foo *string `json:"foo,omitempty"` // want "union member field Reversed.Foo.*requires a recognized"
	// +unionMember,optional
	// +optional
	Maybe *string `json:"maybe,omitempty"` // want "union member field Reversed.Maybe.*expected .*\\? true : !has\\(self.maybe\\)"
}

// +union
// +kubebuilder:validation:XValidation:rule="has(self.kind) && self.kind == 'Foo' ? has(self.foo) : !has(self.foo"
type Malformed struct {
	// +unionDiscriminator
	// +required
	Kind string `json:"kind"`
	// +unionMember
	// +optional
	Foo *string `json:"foo,omitempty"` // want "union member field Malformed.Foo.*requires a recognized"
}

// +union
type InvalidDiscriminator struct { // want "has no discriminator field"
	// +unionMember
	// +optional
	Foo *string `json:"foo,omitempty"`
}

// +union
type MultipleDiscriminators struct { // want "has 2 discriminator fields"
	// +unionDiscriminator
	// +required
	First string `json:"first"`
	// +unionDiscriminator
	// +required
	Second string `json:"second"`
	// +unionMember
	// +optional
	Foo *string `json:"foo,omitempty"`
}

// +union
type NotRequired struct {
	// +unionDiscriminator
	Kind string `json:"kind"` // want "discriminator field NotRequired.Kind must be marked as required"
	// +unionMember
	// +optional
	Foo *string `json:"foo,omitempty"`
}

// +union
type NotOptional struct {
	// +unionDiscriminator
	// +required
	Kind string `json:"kind"`
	// +unionMember,optional
	Foo *string `json:"foo,omitempty"` // want "union member field NotOptional.Foo must be marked as optional"
}

// +union
type DiscriminatorOnly struct {
	// +unionDiscriminator
	// +required
	Kind string `json:"kind"`
}

type Undiscriminated struct {
	// +k8s:unionMember
	// +k8s:optional
	Foo *string `json:"foo,omitempty"`
}

// +union
type Unreachable struct {
	// +unionDiscriminator
	// +required
	Kind string `json:"kind"`
	// +unionMember
	// +optional
	Foo *string `json:"foo,omitempty"`
}

// +kubebuilder:object:root=false
type NotRoot struct {
	Union *Unreachable `json:"union"`
}

// +kubebuilder:object:root:=true
// +union
type RootUnion struct {
	// +unionDiscriminator
	// +required
	Kind string `json:"kind"`
	// +unionMember
	// +optional
	Foo *string `json:"foo,omitempty"` // want "union member field RootUnion.Foo.*requires a recognized"
}

type TypeMeta struct{}
type ListMeta struct{}

// +kubebuilder:object:root=true
type RootList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty"`
	Items    []ListMember `json:"items"`
}

// +union
type ListMember struct {
	// +unionDiscriminator
	// +required
	Kind string `json:"kind"`
	// +unionMember
	// +optional
	Foo *string `json:"foo,omitempty"` // want "union member field ListMember.Foo.*requires a recognized"
}
