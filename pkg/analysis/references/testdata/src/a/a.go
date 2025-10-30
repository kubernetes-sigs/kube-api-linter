package a

// TestWithPolicyAllowRefAndRefs tests the linter with PolicyAllowRefAndRefs
// In this mode, Reference/References are flagged ANYWHERE in field names, but Ref/Refs are allowed
type TestWithPolicyAllowRefAndRefs struct {
	// Fields ending with Reference should be flagged
	NodeReference string `json:"nodeReference"` // want `naming convention "reference-to-ref": field NodeReference: field names should use 'Ref' instead of 'Reference'`

	ConfigReference string `json:"configReference"` // want `naming convention "reference-to-ref": field ConfigReference: field names should use 'Ref' instead of 'Reference'`

	// Fields ending with References should be flagged
	NodeReferences []string `json:"nodeReferences"` // want `naming convention "references-to-refs": field NodeReferences: field names should use 'Refs' instead of 'References'`

	ConfigReferences []string `json:"configReferences"` // want `naming convention "references-to-refs": field ConfigReferences: field names should use 'Refs' instead of 'References'`

	// Fields with Reference at beginning should be flagged
	ReferenceCount int `json:"referenceCount"` // want `naming convention "reference-to-ref": field ReferenceCount: field names should use 'Ref' instead of 'Reference'`

	ReferenceData string `json:"referenceData"` // want `naming convention "reference-to-ref": field ReferenceData: field names should use 'Ref' instead of 'Reference'`

	// Fields with References at beginning should be flagged
	ReferencesCount int `json:"referencesCount"` // want `naming convention "references-to-refs": field ReferencesCount: field names should use 'Refs' instead of 'References'`

	ReferencesData []string `json:"referencesData"` // want `naming convention "references-to-refs": field ReferencesData: field names should use 'Refs' instead of 'References'`

	// Fields with Reference in middle should be flagged
	CrossReferenceID string `json:"crossReferenceID"` // want `naming convention "reference-to-ref": field CrossReferenceID: field names should use 'Ref' instead of 'Reference'`

	// Fields with References in middle should be flagged
	CrossReferencesMap map[string]string `json:"crossReferencesMap"` // want `naming convention "references-to-refs": field CrossReferencesMap: field names should use 'Refs' instead of 'References'`

	// Past tense "Referenced" should be flagged (has "Reference")
	Referenced bool `json:"referenced"` // want `naming convention "reference-to-ref": field Referenced: field names should use 'Ref' instead of 'Reference'`

	// Fields with Ref/Refs anywhere are ALLOWED in this mode (no diagnostics expected)
	NodeRef string `json:"nodeRef"`

	ConfigRef string `json:"configRef"`

	NodeRefs []string `json:"nodeRefs"`

	ConfigRefs []string `json:"configRefs"`

	RefCount int `json:"refCount"`

	RefsData []string `json:"refsData"`

	InternalRefData string `json:"internalRefData"`

	InternalRefsData []string `json:"internalRefsData"`

	// Normal fields should not be flagged (no Reference/References/Ref/Refs)
	Name string `json:"name"`

	Namespace string `json:"namespace"`

	// Edge cases - Preference contains "reference" and WILL be flagged with (?i)
	// This is intentional to catch all variations including malformed casing
	PreferenceType string `json:"preferenceType"` // want `naming convention "reference-to-ref": field PreferenceType: field names should use 'Ref' instead of 'Reference'`

	Preferences map[string]string `json:"preferences,omitempty"` // want `naming convention "references-to-refs": field Preferences: field names should use 'Refs' instead of 'References'`

	// These don't contain capital Reference
	Referral string `json:"referral"`

	Referee string `json:"referee"`
}
