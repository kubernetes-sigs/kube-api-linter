package b

type JSONTagTestStruct struct {
	NoJSONTag            string // want "field JSONTagTestStruct.NoJSONTag is missing json tag"
	EmptyJSONTag         string `json:""`                        // want "field JSONTagTestStruct.EmptyJSONTag has empty json tag"
	NonCamelCaseJSONTag  string `json:"non_camel_case_json_tag"` // want "field JSONTagTestStruct.NonCamelCaseJSONTag json tag does not match pattern \"\\^\\[a-z\\]\\[a-z\\]\\*\\(\\?\\:\\[A-Z\\]\\[a-z0-9\\]\\+\\)\\*\\[a-z0-9\\]\\?\\$\": non_camel_case_json_tag"
	WithHyphensJSONTag   string `json:"with-hyphens-json-tag"`   // want "field JSONTagTestStruct.WithHyphensJSONTag json tag does not match pattern \"\\^\\[a-z\\]\\[a-z\\]\\*\\(\\?\\:\\[A-Z\\]\\[a-z0-9\\]\\+\\)\\*\\[a-z0-9\\]\\?\\$\": with-hyphens-json-tag"
	PascalCaseJSONTag    string `json:"PascalCaseJSONTag"`       // want "field JSONTagTestStruct.PascalCaseJSONTag json tag does not match pattern \"\\^\\[a-z\\]\\[a-z\\]\\*\\(\\?\\:\\[A-Z\\]\\[a-z0-9\\]\\+\\)\\*\\[a-z0-9\\]\\?\\$\": PascalCaseJSONTag"
	NonTerminatedJSONTag string `json:"nonTerminatedJSONTag`     // want "field JSONTagTestStruct.NonTerminatedJSONTag is missing json tag"
	XMLTag               string `xml:"xmlTag"`                   // want "field JSONTagTestStruct.XMLTag is missing json tag"
	InlineJSONTag        string `json:",inline"`
	ValidJSONTag         string `json:"validJsonTag"`
	ValidOptionalJSONTag string `json:"validOptionalJsonTag,omitempty"`
	JSONTagWithID        string `json:"jsonTagWithID"`  // want "field JSONTagTestStruct.JSONTagWithID json tag does not match pattern \"\\^\\[a-z\\]\\[a-z\\]\\*\\(\\?\\:\\[A-Z\\]\\[a-z0-9\\]\\+\\)\\*\\[a-z0-9\\]\\?\\$\": jsonTagWithID"
	JSONTagWithTTL       string `json:"jsonTagWithTTL"` // want "field JSONTagTestStruct.JSONTagWithTTL json tag does not match pattern \"\\^\\[a-z\\]\\[a-z\\]\\*\\(\\?\\:\\[A-Z\\]\\[a-z0-9\\]\\+\\)\\*\\[a-z0-9\\]\\?\\$\": jsonTagWithTTL"
	JSONTagWithGiB       string `json:"jsonTagWithGiB"` // want "field JSONTagTestStruct.JSONTagWithGiB json tag does not match pattern \"\\^\\[a-z\\]\\[a-z\\]\\*\\(\\?\\:\\[A-Z\\]\\[a-z0-9\\]\\+\\)\\*\\[a-z0-9\\]\\?\\$\": jsonTagWithGiB"
}
