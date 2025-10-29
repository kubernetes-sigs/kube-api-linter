package a

type Test struct {
	// This field has both required and default markers - should be flagged
	// +kubebuilder:validation:Required
	// +kubebuilder:default:=testValue
	ConflictedField string `json:"conflictedField"` // want `field Test.ConflictedField has conflicting markers: default_or_required: \{\[kubebuilder:default\], \[kubebuilder:validation:Required\]\}. A field with a default value cannot be required. A required field must be provided by the user, so a default value is not meaningful.`

	// This field has only required marker - should not be flagged
	// +kubebuilder:validation:Required
	RequiredOnlyField string `json:"requiredOnlyField"`

	// This field has only default marker - should not be flagged
	// +kubebuilder:default:=defaultValue
	DefaultOnlyField string `json:"defaultOnlyField"`

	// This field uses KAL required marker with default - should be flagged
	// +required
	// +default:=anotherValue
	KALConflictedField string `json:"kalConflictedField"` // want `field Test.KALConflictedField has conflicting markers: default_or_required: \{\[default\], \[required\]\}. A field with a default value cannot be required. A required field must be provided by the user, so a default value is not meaningful.`

	// This field uses k8s required marker with kubebuilder default - should be flagged
	// +k8s:required
	// +kubebuilder:default:=yetanotherValue
	K8sConflictedField string `json:"k8sConflictedField"` // want `field Test.K8sConflictedField has conflicting markers: default_or_required: \{\[kubebuilder:default\], \[k8s:required\]\}. A field with a default value cannot be required. A required field must be provided by the user, so a default value is not meaningful.`

	// This field has neither marker - should not be flagged
	NormalField string `json:"normalField"`

	// Multiple fields to test various combinations
	ValidField1 string `json:"validField1"` // no markers

	// +required
	ValidField2 string `json:"validField2"` // only required

	// +kubebuilder:default:=value
	ValidField3 string `json:"validField3"` // only default

	// +kubebuilder:validation:Required
	ValidField4 string `json:"validField4"` // only required (kubebuilder format)
}
