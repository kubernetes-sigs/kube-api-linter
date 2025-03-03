package A

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
