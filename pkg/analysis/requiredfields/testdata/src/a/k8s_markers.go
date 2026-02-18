package a

// TestK8sMarkers tests declarative validation markers
type TestK8sMarkers struct {
	// +k8s:required
	String string `json:"string"` // want "field TestK8sMarkers.String has \\+k8s:required but should also have \\+required marker" "field TestK8sMarkers.String should have the omitempty tag." "field TestK8sMarkers.String has a valid zero value \\(\"\"\\), but the validation is not complete \\(e.g. minimum length\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +k8s:required
	StringWithOmitEmpty string `json:"stringWithOmitEmpty,omitempty"` // want "field TestK8sMarkers.StringWithOmitEmpty has \\+k8s:required but should also have \\+required marker" "field TestK8sMarkers.StringWithOmitEmpty has a valid zero value \\(\"\"\\), but the validation is not complete \\(e.g. minimum length\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +k8s:required
	// +k8s:minLength=1
	StringWithMinLength string `json:"stringWithMinLength"` // want "field TestK8sMarkers.StringWithMinLength has \\+k8s:required but should also have \\+required marker" "field TestK8sMarkers.StringWithMinLength should have the omitempty tag."

	// +k8s:required
	// +k8s:minLength=1
	StringWithMinLengthWithOmitEmpty string `json:"stringWithMinLengthWithOmitEmpty,omitempty"` // want "field TestK8sMarkers.StringWithMinLengthWithOmitEmpty has \\+k8s:required but should also have \\+required marker"

	// +k8s:required
	// +k8s:minLength=1
	StringPtrWithMinLength *string `json:"stringPtrWithMinLength"` // want "field TestK8sMarkers.StringPtrWithMinLength has \\+k8s:required but should also have \\+required marker" "field TestK8sMarkers.StringPtrWithMinLength should have the omitempty tag." "field TestK8sMarkers.StringPtrWithMinLength does not allow the zero value. The field does not need to be a pointer."

	// +k8s:required
	// +k8s:minLength=1
	StringPtrWithMinLengthWithOmitEmpty *string `json:"stringPtrWithMinLengthWithOmitEmpty,omitempty"` // want "field TestK8sMarkers.StringPtrWithMinLengthWithOmitEmpty has \\+k8s:required but should also have \\+required marker" "field TestK8sMarkers.StringPtrWithMinLengthWithOmitEmpty does not allow the zero value. The field does not need to be a pointer."

	// +k8s:required
	// +k8s:minLength=0
	StringWithMinLength0 string `json:"stringWithMinLength0"` // want "field TestK8sMarkers.StringWithMinLength0 has \\+k8s:required but should also have \\+required marker" "field TestK8sMarkers.StringWithMinLength0 should have the omitempty tag." "field TestK8sMarkers.StringWithMinLength0 has a valid zero value \\(\"\"\\) and should be a pointer."

	// +k8s:required
	// +k8s:minLength=0
	StringWithMinLength0WithOmitEmpty string `json:"stringWithMinLength0WithOmitEmpty,omitempty"` // want "field TestK8sMarkers.StringWithMinLength0WithOmitEmpty has \\+k8s:required but should also have \\+required marker" "field TestK8sMarkers.StringWithMinLength0WithOmitEmpty has a valid zero value \\(\"\"\\) and should be a pointer."

	// +k8s:required
	// +k8s:minimum=1
	IntegerWithMinimum int `json:"integerWithMinimum"` // want "field TestK8sMarkers.IntegerWithMinimum has \\+k8s:required but should also have \\+required marker" "field TestK8sMarkers.IntegerWithMinimum should have the omitempty tag."

	// +k8s:required
	// +k8s:minimum=1
	IntegerWithMinimumWithOmitEmpty int `json:"integerWithMinimumWithOmitEmpty,omitempty"` // want "field TestK8sMarkers.IntegerWithMinimumWithOmitEmpty has \\+k8s:required but should also have \\+required marker"

	// +k8s:required
	// +k8s:minimum=1
	IntegerPtrWithMinimum *int `json:"integerPtrWithMinimum"` // want "field TestK8sMarkers.IntegerPtrWithMinimum has \\+k8s:required but should also have \\+required marker" "field TestK8sMarkers.IntegerPtrWithMinimum should have the omitempty tag." "field TestK8sMarkers.IntegerPtrWithMinimum does not allow the zero value. The field does not need to be a pointer."

	// +k8s:required
	// +k8s:minimum=1
	IntegerPtrWithMinimumWithOmitEmpty *int `json:"integerPtrWithMinimumWithOmitEmpty,omitempty"` // want "field TestK8sMarkers.IntegerPtrWithMinimumWithOmitEmpty has \\+k8s:required but should also have \\+required marker" "field TestK8sMarkers.IntegerPtrWithMinimumWithOmitEmpty does not allow the zero value. The field does not need to be a pointer."

	// +k8s:required
	// +k8s:minimum=0
	// +k8s:maximum=10
	IntegerWithMinMax int `json:"integerWithMinMax"` // want "field TestK8sMarkers.IntegerWithMinMax has \\+k8s:required but should also have \\+required marker" "field TestK8sMarkers.IntegerWithMinMax should have the omitempty tag." "field TestK8sMarkers.IntegerWithMinMax has a valid zero value \\(0\\) and should be a pointer."

	// +k8s:required
	// +k8s:minimum=0
	// +k8s:maximum=10
	IntegerWithMinMaxWithOmitEmpty int `json:"integerWithMinMaxWithOmitEmpty,omitempty"` // want "field TestK8sMarkers.IntegerWithMinMaxWithOmitEmpty has \\+k8s:required but should also have \\+required marker" "field TestK8sMarkers.IntegerWithMinMaxWithOmitEmpty has a valid zero value \\(0\\) and should be a pointer."

	// +k8s:required
	// +k8s:minItems=1
	ArrayWithMinItems []string `json:"arrayWithMinItems"` // want "field TestK8sMarkers.ArrayWithMinItems has \\+k8s:required but should also have \\+required marker" "field TestK8sMarkers.ArrayWithMinItems should have the omitempty tag."

	// +k8s:required
	// +k8s:minItems=1
	ArrayWithMinItemsWithOmitEmpty []string `json:"arrayWithMinItemsWithOmitEmpty,omitempty"` // want "field TestK8sMarkers.ArrayWithMinItemsWithOmitEmpty has \\+k8s:required but should also have \\+required marker"

	// +k8s:required
	// +k8s:minItems=0
	ArrayWithMinItems0 []string `json:"arrayWithMinItems0"` // want "field TestK8sMarkers.ArrayWithMinItems0 has \\+k8s:required but should also have \\+required marker" "field TestK8sMarkers.ArrayWithMinItems0 should have the omitempty tag."

	// +k8s:required
	// +k8s:minItems=0
	ArrayWithMinItems0WithOmitEmpty []string `json:"arrayWithMinItems0WithOmitEmpty,omitempty"` // want "field TestK8sMarkers.ArrayWithMinItems0WithOmitEmpty has \\+k8s:required but should also have \\+required marker"
}

// TestK8sMarkersWithRequired tests that fields with both +k8s:required and +required don't trigger warning
type TestK8sMarkersWithRequired struct {
	// +required
	// +k8s:required
	// +k8s:minLength=1
	StringWithBothMarkers string `json:"stringWithBothMarkers,omitempty"`

	// +kubebuilder:validation:Required
	// +k8s:required
	// +k8s:minLength=1
	StringWithKubebuilderRequired string `json:"stringWithKubebuilderRequired,omitempty"`
}

// K8s DV Enum types (enum is defined at type level, const values become enum values)

// K8sEnumNoEmpty is an enum type without empty string value
// +k8s:enum
type K8sEnumNoEmpty string

const (
	K8sEnumNoEmptyA K8sEnumNoEmpty = "A"
	K8sEnumNoEmptyB K8sEnumNoEmpty = "B"
	K8sEnumNoEmptyC K8sEnumNoEmpty = "C"
)

// K8sEnumWithEmpty is an enum type with empty string value
// +k8s:enum
type K8sEnumWithEmpty string

const (
	K8sEnumWithEmptyA     K8sEnumWithEmpty = "A"
	K8sEnumWithEmptyB     K8sEnumWithEmpty = "B"
	K8sEnumWithEmptyEmpty K8sEnumWithEmpty = ""
)

// TestK8sEnumOnField tests that +k8s:enum on a field (not a type) is not treated as enum
type TestK8sEnumOnField struct {
	// +k8s:required
	// +k8s:enum
	NotAnEnumType string `json:"notAnEnumType"` // want "field TestK8sEnumOnField.NotAnEnumType has \\+k8s:required but should also have \\+required marker" "field TestK8sEnumOnField.NotAnEnumType should have the omitempty tag." "field TestK8sEnumOnField.NotAnEnumType has a valid zero value \\(\"\"\\), but the validation is not complete \\(e.g. minimum length\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +k8s:required
	// +k8s:enum
	NotAnEnumTypeWithOmitEmpty string `json:"notAnEnumTypeWithOmitEmpty,omitempty"` // want "field TestK8sEnumOnField.NotAnEnumTypeWithOmitEmpty has \\+k8s:required but should also have \\+required marker" "field TestK8sEnumOnField.NotAnEnumTypeWithOmitEmpty has a valid zero value \\(\"\"\\), but the validation is not complete \\(e.g. minimum length\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."
}

// TestK8sEnum tests K8s DV enum fields
type TestK8sEnum struct {
	// +k8s:required
	EnumNoEmpty K8sEnumNoEmpty `json:"enumNoEmpty"` // want "field TestK8sEnum.EnumNoEmpty has \\+k8s:required but should also have \\+required marker" "field TestK8sEnum.EnumNoEmpty should have the omitempty tag."

	// +k8s:required
	EnumNoEmptyWithOmitEmpty K8sEnumNoEmpty `json:"enumNoEmptyWithOmitEmpty,omitempty"` // want "field TestK8sEnum.EnumNoEmptyWithOmitEmpty has \\+k8s:required but should also have \\+required marker"

	// +k8s:required
	EnumNoEmptyPtr *K8sEnumNoEmpty `json:"enumNoEmptyPtr"` // want "field TestK8sEnum.EnumNoEmptyPtr has \\+k8s:required but should also have \\+required marker" "field TestK8sEnum.EnumNoEmptyPtr should have the omitempty tag." "field TestK8sEnum.EnumNoEmptyPtr does not allow the zero value. The field does not need to be a pointer."

	// +k8s:required
	EnumNoEmptyPtrWithOmitEmpty *K8sEnumNoEmpty `json:"enumNoEmptyPtrWithOmitEmpty,omitempty"` // want "field TestK8sEnum.EnumNoEmptyPtrWithOmitEmpty has \\+k8s:required but should also have \\+required marker" "field TestK8sEnum.EnumNoEmptyPtrWithOmitEmpty does not allow the zero value. The field does not need to be a pointer."

	// +k8s:required
	EnumWithEmpty K8sEnumWithEmpty `json:"enumWithEmpty"` // want "field TestK8sEnum.EnumWithEmpty has \\+k8s:required but should also have \\+required marker" "field TestK8sEnum.EnumWithEmpty should have the omitempty tag." "field TestK8sEnum.EnumWithEmpty has a valid zero value \\(\"\"\\) and should be a pointer."

	// +k8s:required
	EnumWithEmptyWithOmitEmpty K8sEnumWithEmpty `json:"enumWithEmptyWithOmitEmpty,omitempty"` // want "field TestK8sEnum.EnumWithEmptyWithOmitEmpty has \\+k8s:required but should also have \\+required marker" "field TestK8sEnum.EnumWithEmptyWithOmitEmpty has a valid zero value \\(\"\"\\) and should be a pointer."

	// +k8s:required
	EnumWithEmptyPtr *K8sEnumWithEmpty `json:"enumWithEmptyPtr"` // want "field TestK8sEnum.EnumWithEmptyPtr has \\+k8s:required but should also have \\+required marker" "field TestK8sEnum.EnumWithEmptyPtr should have the omitempty tag."

	// +k8s:required
	EnumWithEmptyPtrWithOmitEmpty *K8sEnumWithEmpty `json:"enumWithEmptyPtrWithOmitEmpty,omitempty"` // want "field TestK8sEnum.EnumWithEmptyPtrWithOmitEmpty has \\+k8s:required but should also have \\+required marker"
}
