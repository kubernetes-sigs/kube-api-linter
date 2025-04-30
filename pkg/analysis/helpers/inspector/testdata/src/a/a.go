package A

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	String string
)

const (
	Int int = 0
)

type A struct {
	Field string `json:"field"` // want "field: Field" "json tag: field"

	B `json:"b"` // want "field: B" "json tag: b"

	C `json:",inline"` // want "field: C"

	D `json:"-"`

	E struct { // want "field: E" "json tag: e"
		Field string `json:"field"` // want "field: Field" "json tag: field"
	} `json:"e"`

	F struct {
		Field string `json:"field"`
	} `json:"-"`
}

func (A) DoNothing() {}

type B struct {
	Field string `json:"field"` // want "field: Field" "json tag: field"
}

type (
	C struct {
		Field string `json:"field"` // want "field: Field" "json tag: field"
	}

	D struct {
		Field string `json:"field"` // want "field: Field" "json tag: field"
	}
)

func Foo() {
	type Bar struct {
		Field string
	}
}

type Bar interface {
	Name() string
}

var Var = struct {
	Field string
}{
	Field: "field",
}

// AItems is a list of A.
// This represents the Items types in Kubernetes APIs.
// We don't need to lint this type as it doesn't affect the API behaviour in general.
type AItems struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []A `json:"items"`
}

type NotItems struct {
	metav1.TypeMeta `json:",inline"`            // want "field: "
	metav1.ListMeta `json:"metadata,omitempty"` // want "field: " "json tag: metadata"

	Items A `json:"items"` // want "field: Items" "json tag: items"
}

type NotItemsWrongMetadata struct {
	metav1.TypeMeta   `json:",inline"`            // want "field: "
	metav1.ObjectMeta `json:"metadata,omitempty"` // want "field: " "json tag: metadata"

	Items []A `json:"items"` // want "field: Items" "json tag: items"
}

type NotItemsWrongTypeMeta struct {
	metav1.ObjectMeta `json:",inline"`            // want "field: "
	metav1.ListMeta   `json:"metadata,omitempty"` // want "field: " "json tag: metadata"

	Items []A `json:"items"` // want "field: Items" "json tag: items"
}
