package a

type TestBools struct {
	Bool bool `json:"bool"` // want "field TestBools.Bool should have the omitempty tag." "field TestBools.Bool should be a pointer."

	BoolWithOmitEmpty bool `json:"boolWithOmitEmpty,omitempty"` // want "field TestBools.BoolWithOmitEmpty should be a pointer."

	BoolPtr *bool `json:"boolPtr"` // want "field TestBools.BoolPtr should have the omitempty tag."

	BoolPtrWithOmitEmpty *bool `json:"boolPtrWithOmitEmpty,omitempty"`
}
