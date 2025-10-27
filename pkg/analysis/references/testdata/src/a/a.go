package a

type TestSpec struct {
	NodeReference string `json:"nodeReference"` // want "field NodeReference should use 'Ref' instead of 'Reference'"

	NodeReferences []string `json:"nodeReferences"` // want "field NodeReferences should use 'Refs' instead of 'References'"

	// AllowedRef field should be valid
	AllowedRef string `json:"allowedRef"`

	// AllowedRefs field should be valid
	AllowedRefs []string `json:"allowedRefs"`

	// ValidField should be valid
	ValidField string `json:"validField"`

	// ParentRef should be valid
	ParentRef string `json:"parentRef"`

	// ParentRefs should be valid
	ParentRefs []string `json:"parentRefs"`
}
