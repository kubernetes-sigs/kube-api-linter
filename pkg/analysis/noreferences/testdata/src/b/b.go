package b

// TestWithPolicyNoReferences tests the linter with PolicyNoReferences (strict mode)
// In this mode, all reference-related words (Reference/References/Ref/Refs) are warned about but not removed
type TestWithPolicyNoReferences struct {
	// Fields ending with Reference should be flagged
	NodeReference string `json:"nodeReference"` // want `naming convention "no-references": field TestWithPolicyNoReferences.NodeReference: field names should not contain reference-related words`

	// Fields ending with References should be flagged
	ConfigReferences []string `json:"configReferences"` // want `naming convention "no-references": field TestWithPolicyNoReferences.ConfigReferences: field names should not contain reference-related words`

	// Fields with Reference at beginning should be flagged
	ReferenceCount int `json:"referenceCount"` // want `naming convention "no-references": field TestWithPolicyNoReferences.ReferenceCount: field names should not contain reference-related words`

	// Fields with References at beginning should be flagged
	ReferencesData []string `json:"referencesData"` // want `naming convention "no-references": field TestWithPolicyNoReferences.ReferencesData: field names should not contain reference-related words`

	// Fields ending with Ref should be flagged
	PodRef string `json:"podRef"` // want `naming convention "no-references": field TestWithPolicyNoReferences.PodRef: field names should not contain reference-related words`

	// Fields ending with Refs should be flagged
	ContainerRefs []string `json:"containerRefs"` // want `naming convention "no-references": field TestWithPolicyNoReferences.ContainerRefs: field names should not contain reference-related words`

	// Fields starting with Ref should be flagged
	RefStatus string `json:"refStatus"` // want `naming convention "no-references": field TestWithPolicyNoReferences.RefStatus: field names should not contain reference-related words`

	// Fields starting with Refs should be flagged
	RefsTotal int `json:"refsTotal"` // want `naming convention "no-references": field TestWithPolicyNoReferences.RefsTotal: field names should not contain reference-related words`

	// Normal fields should not be flagged
	Name string `json:"name"`

	Namespace string `json:"namespace"`

	// Edge cases - "Preference" contains "eference" but is NOT flagged
	// because it doesn't start with Ref/Reference and doesn't end with capital-R Reference
	PreferenceType string `json:"preferenceType"`

	Preferences map[string]string `json:"preferences,omitempty"`
}
