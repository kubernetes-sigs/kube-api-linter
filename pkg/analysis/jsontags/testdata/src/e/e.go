package e

type JSONTagRegexAndFieldNameMismatch struct {
	FieldName string `json:"field_name,omitempty"` // want "json tag does not match pattern" "field JSONTagRegexAndFieldNameMismatch.FieldName json tag should match the camelCase field name \"fieldName\": got \"field_name\""
}
