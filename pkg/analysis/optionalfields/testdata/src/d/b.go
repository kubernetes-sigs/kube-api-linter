package a

type StructWithRequiredField struct {
	// string is a string field.
	// +required
	String string `json:"string"`
}
