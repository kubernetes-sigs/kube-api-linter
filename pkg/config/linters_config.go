package config

// LintersConfig contains configuration for individual linters.
type LintersConfig struct {
	// jsonTags contains configuration for the jsontags linter.
	JSONTags JSONTagsConfig `json:"jsonTags"`

	// optionalOrRequired contains configuration for the optionalorrequired linter.
	OptionalOrRequired OptionalOrRequiredConfig `json:"optionalOrRequired"`

	// requiredFields contains configuration for the requiredfields linter.
	RequiredFields RequiredFieldsConfig `json:"requiredFields"`
}

// JSONTagsConfig contains configuration for the jsontags linter.
type JSONTagsConfig struct {
	// jsonTagRegex is the regular expression used to validate that json tags are in a particular format.
	// By default, the regex used is "^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*)*$" and is used to check for
	// camel case like string.
	JSONTagRegex string `json:"jsonTagRegex"`
}

// OptionalOrRequiredConfig contains configuration for the optionalorrequired linter.
type OptionalOrRequiredConfig struct {
	// preferredOptionalMarker is the preferred marker to use for optional fields.
	// If this field is not set, the default value is "optional".
	// Valid values are "optional" and "kubebuilder:validation:Optional".
	PreferredOptionalMarker string `json:"preferredOptionalMarker"`

	// preferredRequiredMarker is the preferred marker to use for required fields.
	// If this field is not set, the default value is "required".
	// Valid values are "required" and "kubebuilder:validation:Required".
	PreferredRequiredMarker string `json:"preferredRequiredMarker"`
}

// RequiredFieldPointerPolicy is the policy for pointers in required fields.
type RequiredFieldPointerPolicy string

const (
	// RequiredFieldPointerWarn indicates that the linter will emit a warning if a required field is a pointer.
	RequiredFieldPointerWarn RequiredFieldPointerPolicy = "Warn"

	// RequiredFieldPointerSuggestFix indicates that the linter will emit a warning if a required field is a pointer and suggest a fix.
	RequiredFieldPointerSuggestFix RequiredFieldPointerPolicy = "SuggestFix"
)

// RequiredFieldsConfig contains configuration for the requiredfields linter.
type RequiredFieldsConfig struct {
	// pointerPolicy is the policy for pointers in required fields.
	// Valid values are "Warn" and "SuggestFix".
	// When set to "Warn", the linter will emit a warning if a required field is a pointer.
	// When set to "SuggestFix", the linter will emit a warning if a required field is a pointer and suggest a fix.
	// When otherwise not specified, the default value is "SuggestFix".
	PointerPolicy RequiredFieldPointerPolicy `json:"pointerPolicy"`
}
