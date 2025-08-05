package a

type TestBools struct {
	Bool bool `json:"bool"` // want "field Bool should have the omitempty tag." "field Bool should be a pointer."

	BoolWithOmitEmpty bool `json:"boolWithOmitEmpty,omitempty"` // want "field BoolWithOmitEmpty should be a pointer."

	BoolPtr *bool `json:"boolPtr"` // want "field BoolPtr should have the omitempty tag."

	BoolPtrWithOmitEmpty *bool `json:"boolPtrWithOmitEmpty,omitempty"`
}
