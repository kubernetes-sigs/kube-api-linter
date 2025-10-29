package a

// Valid markers that should not trigger any errors

// +kubebuilder:validation:MaxLength:=256
type ValidKubebuilderMarker string

// +required
// +optional
type ValidNonKubebuilderMarkers string

// +kubebuilder:object:root:=true
// +kubebuilder:subresource:status
type ValidKubebuilderObject struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength:=1
	ValidField string `json:"validField"`

	// +optional
	// +default:value="test"
	ValidOptionalField string `json:"validOptionalField"`
}

// Invalid markers that should trigger errors

// + kubebuilder:validation:MaxLength:=256 // want "marker should not have space after '\\+' symbol"
type SpacingIssueKubebuilder string

// + required // want "marker should not have space after '\\+' symbol"
type SpacingIssueNonKubebuilder string

// +kubebuilder:validation:MaxLength=256
type KubebuilderWrongSyntax string

// +kubebuilder:validation:MinLength=1
// +kubebuilder:validation:Format=date-time
type MultipleKubebuilderWrongSyntax struct {
	Field string `json:"field"`
}

// +default:value:="test"
type NonKubebuilderWrongSyntax string

// +required:=true
// +optional:=false
type MultipleNonKubebuilderWrongSyntax struct {
	Field string `json:"field"`
}

// Common typos

// +kubebuidler:validation:MaxLength:=256 // want "possible typo: 'kubebuidler' should be 'kubebuilder'"
type KubebuilderTypo1 string

// +kubebuiler:validation:Required // want "possible typo: 'kubebuiler' should be 'kubebuilder'"
type KubebuilderTypo2 string

// +kubebulider:object:root:=true // want "possible typo: 'kubebulider' should be 'kubebuilder'"
type KubebuilderTypo3 string

// +kubbuilder:validation:MinLength:=1 // want "possible typo: 'kubbuilder' should be 'kubebuilder'"
type KubebuilderTypo4 string

// +kubebulder:subresource:status // want "possible typo: 'kubebulder' should be 'kubebuilder'"
type KubebuilderTypo5 string

// +optinal // want "possible typo: 'optinal' should be 'optional'"
type OptionalTypo string

// +requied // want "possible typo: 'requied' should be 'required'"
type RequiredTypo string

// +requird // want "possible typo: 'requird' should be 'required'"
type RequiredTypo2 string

// +nullabel // want "possible typo: 'nullabel' should be 'nullable'"
type NullableTypo string

// +kubebuilder:validaton:MaxLength:=256 // want "possible typo: 'validaton' should be 'validation'"
type ValidationTypo string

// +kubebuilder:valdiation:Required // want "possible typo: 'valdiation' should be 'validation'"
type ValidationTypo2 string

// +defualt:value="test" // want "possible typo: 'defualt' should be 'default'"
type DefaultTypo string

// +defult:value="test" // want "possible typo: 'defult' should be 'default'"
type DefaultTypo2 string

// +kubebuilder:exampl:="test" // want "possible typo: 'exampl' should be 'example'"
type ExampleTypo string

// +kubebuilder:examle:="test" // want "possible typo: 'examle' should be 'example'"
type ExampleTypo2 string

// Missing space after // prefix

// +kubebuilder:validation:MaxLength:=256 // want "marker should have space after '//' comment prefix"
type MissingSpaceKubebuilder string

// +required // want "marker should have space after '//' comment prefix"
type MissingSpaceNonKubebuilder string

// +optional // want "marker should have space after '//' comment prefix"
type MissingSpaceOptional string

// +kubebuilder:object:root:=true // want "marker should have space after '//' comment prefix"
type MissingSpaceKubebuilderObject string

// +default:value="test" // want "marker should have space after '//' comment prefix"
type MissingSpaceDefault string

// Complex cases with multiple issues

// + kubebuidler:validaton:MaxLength=256 // want "marker should not have space after '\\+' symbol" "possible typo: 'kubebuidler' should be 'kubebuilder'" "possible typo: 'validaton' should be 'validation'"
type MultipleIssues string

// +kubebuidler:validaton:MaxLength=256 // want "marker should have space after '//' comment prefix" "possible typo: 'kubebuidler' should be 'kubebuilder'" "possible typo: 'validaton' should be 'validation'"
type MultipleIssuesWithMissingSpace string

// +requied // want "marker should have space after '//' comment prefix" "possible typo: 'requied' should be 'required'"
type MissingSpaceAndTypo string

// +defualt:value:="test" // want "marker should have space after '//' comment prefix" "possible typo: 'defualt' should be 'default'"
type MissingSpaceDefaultTypoWrongSyntax string

// +kubebuilder:validation:MaxLength:=256
type ComplexValidStruct struct {
	// + requied // want "marker should not have space after '\\+' symbol" "possible typo: 'requied' should be 'required'"
	InvalidField1 string `json:"invalidField1"`

	//+kubebuilder:validation:Required // want "marker should have space after '//' comment prefix"
	FieldWithMissingSpace string `json:"fieldWithMissingSpace"`

	//+optional // want "marker should have space after '//' comment prefix"
	AnotherFieldWithMissingSpace string `json:"anotherFieldWithMissingSpace"`

	// +kubebuilder:validation:Required
	ValidField string `json:"validField"`
}

type NoLintMarker //nolint

