package a

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type JSONTagTestStruct struct {
	NoJSONTag            string // want "field NoJSONTag is missing json tag"
	EmptyJSONTag         string `json:""`                        // want "field EmptyJSONTag has empty json tag"
	NonCamelCaseJSONTag  string `json:"non_camel_case_json_tag"` // want "field NonCamelCaseJSONTag json tag does not match pattern \"\\^\\[a-z\\]\\[a-z0-9\\]\\*\\(\\?:\\[A-Z\\]\\[a-z0-9\\]\\*\\)\\*\\$\": non_camel_case_json_tag"
	WithHyphensJSONTag   string `json:"with-hyphens-json-tag"`   // want "field WithHyphensJSONTag json tag does not match pattern \"\\^\\[a-z\\]\\[a-z0-9\\]\\*\\(\\?:\\[A-Z\\]\\[a-z0-9\\]\\*\\)\\*\\$\": with-hyphens-json-tag"
	PascalCaseJSONTag    string `json:"PascalCaseJSONTag"`       // want "field PascalCaseJSONTag json tag does not match pattern \"\\^\\[a-z\\]\\[a-z0-9\\]\\*\\(\\?:\\[A-Z\\]\\[a-z0-9\\]\\*\\)\\*\\$\": PascalCaseJSONTag"
	NonTerminatedJSONTag string `json:"nonTerminatedJSONTag`     // want "field NonTerminatedJSONTag is missing json tag"
	XMLTag               string `xml:"xmlTag"`                   // want "field XMLTag is missing json tag"
	InlineJSONTag        string `json:",inline"`
	ValidJSONTag         string `json:"validJsonTag"`
	ValidOptionalJSONTag string `json:"validOptionalJsonTag,omitempty"`
	JSONTagWithID        string `json:"jsonTagWithID"`
	JSONTagWithTTL       string `json:"jsonTagWithTTL"`
	JSONTagWithGiB       string `json:"jsonTagWithGiB"`
	Ignored              string `json:"-"`

	IgnoredAnonymousStruct struct {
		// This field should be ignored since the parent field is ignored.
		A string `json:""`
	} `json:"-"`

	A `json:",inline"`
	B `json:"bar,omitempty"`
	C             // want "embedded field C is missing json tag"
	D `json:""`   // want "embedded field D has empty json tag"
	E `json:"e-"` // want "embedded field E json tag does not match pattern \"\\^\\[a-z\\]\\[a-z0-9\\]\\*\\(\\?:\\[A-Z\\]\\[a-z0-9\\]\\*\\)\\*\\$\": e-"

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
}

type A struct{}

func (A) DoNothing() {}

type B struct{}

type C struct{}

type D struct{}

type E struct{}

type Interface interface {
	InaccessibleFunction() string
}

// ValidList is a properly tagged list type
type ValidList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []A `json:"items"`
}

// InvalidList is a list type with missing JSON tag on Items field
type InvalidList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []A // want "field Items is missing json tag"
}

// InvalidListEmptyTag is a list type with empty JSON tag on Items field
type InvalidListEmptyTag struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []A `json:""` // want "field Items has empty json tag"
}

// InvalidListSnakeCase is a list type with snake_case JSON tag on Items field
type InvalidListSnakeCase struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []A `json:"item_list"` // want "field Items json tag does not match pattern \"\\^\\[a-z\\]\\[a-z0-9\\]\\*\\(\\?:\\[A-Z\\]\\[a-z0-9\\]\\*\\)\\*\\$\": item_list"
}

// InvalidListKebabCase is a list type with kebab-case JSON tag on Items field
type InvalidListKebabCase struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []A `json:"item-list"` // want "field Items json tag does not match pattern \"\\^\\[a-z\\]\\[a-z0-9\\]\\*\\(\\?:\\[A-Z\\]\\[a-z0-9\\]\\*\\)\\*\\$\": item-list"
}

// InvalidListMissingMetadataTag is a list type with missing JSON tag on ListMeta
type InvalidListMissingMetadataTag struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta // want "embedded field  is missing json tag"
	Items           []A `json:"items"`
}

// InvalidListWrongMetadataTag is a list type with invalid JSON tag on ListMeta
type InvalidListWrongMetadataTag struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"meta_data"` // want "embedded field  json tag does not match pattern \"\\^\\[a-z\\]\\[a-z0-9\\]\\*\\(\\?:\\[A-Z\\]\\[a-z0-9\\]\\*\\)\\*\\$\": meta_data"
	Items           []A `json:"items"`
}

// ListWithAdditionalFields is a list type with extra fields beyond the standard 3
type ListWithAdditionalFields struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ListMeta   `json:"metadata"`
	Items             []A    `json:"items"`
	ExtraField        string `json:"extraField"`
	MissingTagField   string // want "field MissingTagField is missing json tag"
	InvalidTagField   string `json:"invalid_tag"` // want "field InvalidTagField json tag does not match pattern \"\\^\\[a-z\\]\\[a-z0-9\\]\\*\\(\\?:\\[A-Z\\]\\[a-z0-9\\]\\*\\)\\*\\$\": invalid_tag"
}

// ListWithIgnoredField is a list type with an ignored field
type ListWithIgnoredField struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []A    `json:"items"`
	IgnoredField    string `json:"-"`
}

// Resource is a test resource type
type Resource struct {
	Field string `json:"field"`
}

// ResourceList is the exact scenario from issue #147
// This would serialize incorrectly without proper JSON tags
type ResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Resource // want "field Items is missing json tag"
}

// ValidResourceList is the corrected version
type ValidResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Resource `json:"items"`
}

// ListWithPascalCaseTag tests that PascalCase tags are caught
type ListWithPascalCaseTag struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []A `json:"Items"` // want "field Items json tag does not match pattern \"\\^\\[a-z\\]\\[a-z0-9\\]\\*\\(\\?:\\[A-Z\\]\\[a-z0-9\\]\\*\\)\\*\\$\": Items"
}

// NotActuallyAList should still be linted even though it has TypeMeta and ListMeta
// but is not a proper 3-field list type
type NotActuallyAList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           A      // want "field Items is missing json tag"
	SomeOtherField  string `json:"otherField"`
}
