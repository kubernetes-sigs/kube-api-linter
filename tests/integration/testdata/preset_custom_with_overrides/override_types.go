package a

// OverrideSpec tests preset + user enable/disable overrides.
// Config: preset TestOnlyJSONTags (jsontags only), enable maxlength, disable jsontags.
// Result: only maxlength should fire, jsontags should NOT fire.
type OverrideSpec struct {
	// missingJSONTag would trigger jsontags, but jsontags is disabled by user override.
	// No jsontags diagnostic expected.
	// maxlength also fires on bare strings.
	MissingJSONTag string // want "field OverrideSpec.MissingJSONTag must have a maximum length, add kubebuilder:validation:MaxLength marker"

	// name has a json tag but no maxLength marker.
	// maxlength is enabled by user override, so it should flag this.
	// +required
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name,omitempty"` // want "field OverrideSpec.Name must have a maximum length, add kubebuilder:validation:MaxLength marker"
}
