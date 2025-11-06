package a

// TestWithPolicyPreferAbbreviatedReference tests the linter with PolicyPreferAbbreviatedReference
// In this mode, Reference/References are flagged at START or END of field names, but Ref/Refs are allowed
type TestWithPolicyPreferAbbreviatedReference struct {
	// Fields ending with Reference should be flagged
	NodeReference string `json:"nodeReference"` // want `naming convention "reference-to-ref": field NodeReference: field names should use 'Ref' instead of 'Reference'`

	ConfigReference string `json:"configReference"` // want `naming convention "reference-to-ref": field ConfigReference: field names should use 'Ref' instead of 'Reference'`

	// Fields ending with References should be flagged
	NodeReferences []string `json:"nodeReferences"` // want `naming convention "reference-to-ref": field NodeReferences: field names should use 'Ref' instead of 'Reference'`

	ConfigReferences []string `json:"configReferences"` // want `naming convention "reference-to-ref": field ConfigReferences: field names should use 'Ref' instead of 'Reference'`

	// Fields with Reference at beginning should be flagged
	ReferenceCount int `json:"referenceCount"` // want `naming convention "reference-to-ref": field ReferenceCount: field names should use 'Ref' instead of 'Reference'`

	ReferenceData string `json:"referenceData"` // want `naming convention "reference-to-ref": field ReferenceData: field names should use 'Ref' instead of 'Reference'`

	// Fields with References at beginning should be flagged
	ReferencesCount int `json:"referencesCount"` // want `naming convention "reference-to-ref": field ReferencesCount: field names should use 'Ref' instead of 'Reference'`

	ReferencesData []string `json:"referencesData"` // want `naming convention "reference-to-ref": field ReferencesData: field names should use 'Ref' instead of 'Reference'`

	// Fields with Reference in middle are NOT flagged (only start/end boundaries)
	CrossReferenceID string `json:"crossReferenceID"`

	// Fields with References in middle are NOT flagged (only start/end boundaries)
	CrossReferencesMap map[string]string `json:"crossReferencesMap"`

	// Past tense "Referenced" at end is flagged
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

	// Edge cases - Preference contains "eference" in middle and is NOT flagged (only boundaries)
	PreferenceType string `json:"preferenceType"`

	Preferences map[string]string `json:"preferences,omitempty"`

	// These don't contain capital Reference
	Referral string `json:"referral"`

	Referee string `json:"referee"`
}
