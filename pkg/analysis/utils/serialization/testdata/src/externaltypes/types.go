package externaltypes

// ResourceList is a named type with underlying type map.
// This simulates corev1.ResourceList from the k8s.io/api package.
type ResourceList map[string]string

// StringSlice is a named type with underlying type slice.
type StringSlice []string

// StringMap is a named type with underlying type map.
type StringMap map[string]string

// StringAlias is a named type with underlying type string.
type StringAlias string

// IntAlias is a named type with underlying type int.
type IntAlias int

// BoolAlias is a named type with underlying type bool.
type BoolAlias bool

// FloatAlias is a named type with underlying type float64.
type FloatAlias float64

// StructType is a named struct type from an external package.
// All fields have omitempty, so the zero value is {}.
type StructType struct {
	// Field is a string field.
	Field string `json:"field,omitempty"`
}

// StructTypeWithNonOmittedField is a named struct type with a non-omitempty field.
// The zero value would include the non-omitted field: {"name": ""}.
type StructTypeWithNonOmittedField struct {
	// Name is a required field without omitempty.
	Name string `json:"name"`
	// Description is an optional field with omitempty.
	Description string `json:"description,omitempty"`
}
