package b

// TestWithPolicyForbidRefAndRefs tests the linter with PolicyForbidRefAndRefs (strict mode)
// In this mode, Reference/References AND Ref/Refs (at end) are flagged
type TestWithPolicyForbidRefAndRefs struct {
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

	// Fields ending with Ref should be FORBIDDEN in this mode
	NodeRef string `json:"nodeRef"` // want `naming convention "forbid-ref": field NodeRef: should not use 'Ref'`

	ConfigRef string `json:"configRef"` // want `naming convention "forbid-ref": field ConfigRef: should not use 'Ref'`

	// Fields ending with Refs should be FORBIDDEN in this mode
	NodeRefs []string `json:"nodeRefs"` // want `naming convention "forbid-refs": field NodeRefs: should not use 'Refs'`

	ConfigRefs []string `json:"configRefs"` // want `naming convention "forbid-refs": field ConfigRefs: should not use 'Refs'`

	// Normal fields should not be flagged (no Reference/References/Ref/Refs)
	Name string `json:"name"`

	Namespace string `json:"namespace"`

	// Edge cases - Preference contains "reference" and will be flagged
	PreferenceType string `json:"preferenceType"` // want `naming convention "reference-to-ref": field PreferenceType: field names should use 'Ref' instead of 'Reference'`

	Preferences map[string]string `json:"preferences,omitempty"` // want `naming convention "references-to-refs": field Preferences: field names should use 'Refs' instead of 'References'`
}
