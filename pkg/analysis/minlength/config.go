package minlength

// Config is the set of configuration options
// for the minlength linting rule.
type Config struct {
	// preferredSuggestionMarkerType is an optional
	// field used to specify the marker style you prefer
	// from suggestions made by the minlength linter.
	//
	// Allowed values are Kubebuilder, DeclarativeValidation and omitted.
	//
	// When set to Kubebuilder, the minlength linter will suggest the
	// use of Kubebuilder-style minimum constraint markers in its suggestions.
	// For example, `kubebuilder:validation:Minimum`.
	//
	// When set to DeclarativeValidation, the minlength linter will
	// suggest the use of Declarative Validation-style minimum constraint markers
	// in its suggestions.
	// For example, `k8s:minimum`.
	//
	// When omitted, this defaults to Kubebuilder.
	PreferredSuggestionMarkerType PreferredSuggestionMarkerType `json:"preferredSuggestionMarkerType,omitempty"`
}

// PreferredSuggestionMarkerType is a representation of the different
// marker types that can be returned as part of the suggestions
// for markers to add to a field/type.
type PreferredSuggestionMarkerType string

const (
	// PreferredSuggestionMarkerTypeKubebuilder is used to tell
	// the minlength linter that Kubebuilder-style markers are preferred
	// for enforcing minimum constraints and should be used in suggestions.
	PreferredSuggestionMarkerTypeKubebuilder PreferredSuggestionMarkerType = "Kubebuilder"

	// PreferredSuggestionMarkerTypeDeclarativeValidation is used to tell
	// the minlength linter that Declarative Validation-style markers are preferred
	// for enforcing minimum constraints and should be used in suggestions.
	PreferredSuggestionMarkerTypeDeclarativeValidation PreferredSuggestionMarkerType = "DeclarativeValidation"
)
