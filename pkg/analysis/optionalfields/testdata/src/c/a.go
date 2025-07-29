package a

type A struct {
	// required field should not be picked up.
	// +required
	RequiredField string `json:"requiredField,omitempty"`

	// pointerString is a pointer string field.
	// +optional
	PointerString *string `json:"pointerString,omitempty"`

	// pointerStringWithoutOmitEmpty is a pointer string field without omitempty/
	// +optional
	PointerStringWithoutOmitEmpty *string `json:"pointerStringWithoutOmitEmpty"` // want "field PointerStringWithoutOmitEmpty is optional, without omitempty and allows the zero value. The field does not need to be a pointer."

	// pointerStringWithMinLength1 with minimum length is a pointer string field.
	// +kubebuilder:validation:MinLength=1
	// +optional
	PointerStringWithMinLength1 *string `json:"pointerStringWithMinLength1,omitempty"` // want "field PointerStringWithMinLength1 is optional and does not allow the zero value. The field does not need to be a pointer."

	// pointerStringWithMinLength1WithoutOmitEmpty with minimum length is a pointer string field without omitempty.
	// +kubebuilder:validation:MinLength=1
	// +optional
	PointerStringWithMinLength1WithoutOmitEmpty *string `json:"pointerStringWithMinLength1WithoutOmitEmpty"` // want "field PointerStringWithMinLength1WithoutOmitEmpty is optional and does not allow the zero value. The field does not need to be a pointer." "field PointerStringWithMinLength1WithoutOmitEmpty is optional and does not allow the zero value. It must have the omitempty tag."

	// pointerStringWithMinLength0 with minimum length is a pointer string field.
	// +kubebuilder:validation:MinLength=0
	// +optional
	PointerStringWithMinLength0 *string `json:"pointerStringWithMinLength0,omitempty"`

	// pointerStringWithMinLength0WithoutOmitempty with minimum length is a pointer string field without omitempty.
	// +kubebuilder:validation:MinLength=0
	// +optional
	PointerStringWithMinLength0WithoutOmitempty *string `json:"pointerStringWithMinLength0WithoutOmitempty"` // want "field PointerStringWithMinLength0WithoutOmitempty is optional, without omitempty and allows the zero value. The field does not need to be a pointer."

	// pointerInt is a pointer int field.
	// +optional
	PointerInt *int `json:"pointerInt,omitempty"`

	// pointerIntWithoutOmitEmpty is a pointer int field with omitempty.
	// +optional
	PointerIntWithoutOmitEmpty *int `json:"pointerIntWithoutOmitEmpty"` // want "field PointerIntWithoutOmitEmpty is optional, without omitempty and allows the zero value. The field does not need to be a pointer."

	// pointerIntWithMinValue1 with minimum value is a pointer int field.
	// +kubebuilder:validation:Minimum=1
	// +optional
	PointerIntWithMinValue1 *int `json:"pointerIntWithMinValue1,omitempty"` // want "field PointerIntWithMinValue1 is optional and does not allow the zero value. The field does not need to be a pointer."

	// pointerIntWithMinValue1WithoutOmitEmpty with minimum value is a pointer int field without omitempty.
	// +kubebuilder:validation:Minimum=1
	// +optional
	PointerIntWithMinValue1WithoutOmitEmpty *int `json:"pointerIntWithMinValue1WithoutOmitEmpty"` // want "field PointerIntWithMinValue1WithoutOmitEmpty is optional and does not allow the zero value. The field does not need to be a pointer." "field PointerIntWithMinValue1WithoutOmitEmpty is optional and does not allow the zero value. It must have the omitempty tag."

	// pointerIntWithMinValue0 with minimum value is a pointer int field.
	// +kubebuilder:validation:Minimum=0
	// +optional
	PointerIntWithMinValue0 *int `json:"pointerIntWithMinValue0,omitempty"`

	// pointerIntWithMinValue0WithoutOmitEmpty with minimum value is a pointer int field without omitempty.
	// +kubebuilder:validation:Minimum=0
	// +optional
	PointerIntWithMinValue0WithoutOmitEmpty *int `json:"pointerIntWithMinValue0WithoutOmitEmpty"` // want "field PointerIntWithMinValue0WithoutOmitEmpty is optional, without omitempty and allows the zero value. The field does not need to be a pointer."

	// pointerIntWithNegativeMaximumValue with negative maximum value is a pointer int field.
	// +kubebuilder:validation:Maximum=-1
	// +optional
	PointerIntWithNegativeMaximumValue *int `json:"pointerIntWithNegativeMaximumValue,omitempty"` // want "field PointerIntWithNegativeMaximumValue is optional and does not allow the zero value. The field does not need to be a pointer."

	// pointerIntWithNegativeMaximumValueWithoutOmitEmpty with negative maximum value is a pointer int field without omitempty.
	// +kubebuilder:validation:Maximum=-1
	// +optional
	PointerIntWithNegativeMaximumValueWithoutOmitEmpty *int `json:"pointerIntWithNegativeMaximumValueWithoutOmitEmpty"` // want "field PointerIntWithNegativeMaximumValueWithoutOmitEmpty is optional and does not allow the zero value. The field does not need to be a pointer." "field PointerIntWithNegativeMaximumValueWithoutOmitEmpty is optional and does not allow the zero value. It must have the omitempty tag."

	// pointerIntWithNegativeMinimumValue with negative minimum value is a pointer int field.
	// +kubebuilder:validation:Minimum=-1
	// +optional
	PointerIntWithNegativeMinimumValue *int `json:"pointerIntWithNegativeMinimumValue,omitempty"`

	// pointerIntWithPositiveMaximumValue with positive maximum value is a pointer int field.
	// +kubebuilder:validation:Maximum=1
	// +optional
	PointerIntWithPositiveMaximumValue *int `json:"pointerIntWithPositiveMaximumValue,omitempty"`

	// pointerIntWithPositiveMaximumValueWithoutOmitEmpty with positive maximum value is a pointer int field without omitempty.
	// +kubebuilder:validation:Maximum=1
	// +optional
	PointerIntWithPositiveMaximumValueWithoutOmitEmpty *int `json:"pointerIntWithPositiveMaximumValueWithoutOmitEmpty"` // want "field PointerIntWithPositiveMaximumValueWithoutOmitEmpty is optional, without omitempty and allows the zero value. The field does not need to be a pointer."

	// pointerIntWithRange is a pointer int field with a range of values including 0.
	// +kubebuilder:validation:Minimum=-10
	// +kubebuilder:validation:Maximum=10
	// +optional
	PointerIntWithRange *int `json:"pointerIntWithRange,omitempty"`

	// pointerIntWithRangeWithoutOmitEmpty is a pointer int field with a range of values including 0 wihtout omitempty.
	// +kubebuilder:validation:Minimum=-10
	// +kubebuilder:validation:Maximum=10
	// +optional
	PointerIntWithRangeWithoutOmitEmpty *int `json:"pointerIntWithRangeWithoutOmitEmpty"` // want "field PointerIntWithRangeWithoutOmitEmpty is optional, without omitempty and allows the zero value. The field does not need to be a pointer."

	// pointerStruct is a pointer struct field.
	// +optional
	PointerStruct *B `json:"pointerStruct,omitempty"`

	// pointerStructWithoutOmitEmpty is a pointer struct field without omitempty.
	// +optional
	PointerStructWithoutOmitEmpty *B `json:"pointerStructWithoutOmitEmpty"` // want "field PointerStructWithoutOmitEmpty is optional, without omitempty and allows the zero value. The field does not need to be a pointer."

	// string is a string field.
	// +optional
	String string `json:"string,omitempty"` // want "field String is optional and should be a pointer" "field String is optional and has a valid zero value \\(\"\"\\), but the validation is not complete \\(e.g. minimum length\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// stringWithoutOmitEmpty is a string field without omitempty.
	// +optional
	StringWithoutOmitEmpty string `json:"stringWithoutOmitEmpty"`

	// stringWithMinLength1 with minimum length is a string field.
	// +kubebuilder:validation:MinLength=1
	// +optional
	StringWithMinLength1 string `json:"stringWithMinLength1,omitempty"`

	// stringWithMinLength1WithoutOmitEmpty with minimum length is a string field without omitempty.
	// +kubebuilder:validation:MinLength=1
	// +optional
	StringWithMinLength1WithoutOmitEmpty string `json:"stringWithMinLength1WithoutOmitEmpty"` // want "field StringWithMinLength1WithoutOmitEmpty is optional and does not allow the zero value. It must have the omitempty tag."

	// stringWithMinLength0 with minimum length is a string field.
	// +kubebuilder:validation:MinLength=0
	// +optional
	StringWithMinLength0 string `json:"stringWithMinLength0,omitempty"` // want "field StringWithMinLength0 is optional and should be a pointer"

	// EnumString is a string field with an enum.
	// +kubebuilder:validation:Enum=foo;bar;baz
	// +optional
	EnumString string `json:"enumString,omitempty"`

	// EnumStringWithEmptyValue is a string field with an enum, also allowing the empty string.
	// +kubebuilder:validation:Enum=foo;bar;baz;""
	// +optional
	EnumStringWithEmptyValue string `json:"enumStringWithEmptyValue,omitempty"` // want "field EnumStringWithEmptyValue is optional and should be a pointer"

	// stringWithMinLength0WithoutOmitEmpty with minimum length is a string field without omitempty.
	// +kubebuilder:validation:MinLength=0
	// +optional
	StringWithMinLength0WithoutOmitEmpty string `json:"stringWithMinLength0WithoutOmitEmpty"`

	// int is an int field.
	// +optional
	Int int `json:"int,omitempty"` // want "field Int is optional and should be a pointer" "field Int is optional and has a valid zero value \\(0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// intWithMinValue1 with minimum value is an int field.
	// +kubebuilder:validation:Minimum=1
	// +optional
	IntWithMinValue1 int `json:"intWithMinValue1,omitempty"`

	// intWithMinValue1WithoutOmitEmpty with minimum value is an int field without omitempty.
	// +kubebuilder:validation:Minimum=1
	// +optional
	IntWithMinValue1WithoutOmitEmpty int `json:"intWithMinValue1WithoutOmitEmpty"` // want "field IntWithMinValue1WithoutOmitEmpty is optional and does not allow the zero value. It must have the omitempty tag."

	// intWithMinValue0 with minimum value is an int field.
	// +kubebuilder:validation:Minimum=0
	// +optional
	IntWithMinValue0 int `json:"intWithMinValue0,omitempty"` // want "field IntWithMinValue0 is optional and should be a pointer"

	// intWithMinValue0WithoutOmitEmpty with minimum value is an int field without omitempty.
	// +kubebuilder:validation:Minimum=0
	// +optional
	IntWithMinValue0WithoutOmitEmpty int `json:"intWithMinValue0WithoutOmitEmpty"`

	// intWithNegativeMaximumValue with negative maximum value is an int field.
	// +kubebuilder:validation:Maximum=-1
	// +optional
	IntWithNegativeMaximumValue int `json:"intWithNegativeMaximumValue,omitempty"`

	// intWithNegativeMaximumValueWithoutOmitEmpty with negative maximum value is an int field without omitempty.
	// +kubebuilder:validation:Maximum=-1
	// +optional
	IntWithNegativeMaximumValueWithoutOmitEmpty int `json:"intWithNegativeMaximumValueWithoutOmitEmpty"` // want "field IntWithNegativeMaximumValueWithoutOmitEmpty is optional and does not allow the zero value. It must have the omitempty tag."

	// intWithNegativeMinimumValue with negative minimum value is an int field.
	// +kubebuilder:validation:Minimum=-1
	// +optional
	IntWithNegativeMinimumValue int `json:"intWithNegativeMinimumValue,omitempty"` // want "field IntWithNegativeMinimumValue is optional and should be a pointer" "field IntWithNegativeMinimumValue is optional and has a valid zero value \\(0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// intWithNegativeMinimumValueWithoutOmitEmpty with negative minimum value is an int field without omitempty.
	// +kubebuilder:validation:Minimum=-1
	// +optional
	IntWithNegativeMinimumValueWithoutOmitEmpty int `json:"intWithNegativeMinimumValueWithoutOmitEmpty"`

	// intWithPositiveMaximumValue with positive maximum value is an int field.
	// +kubebuilder:validation:Maximum=1
	// +optional
	IntWithPositiveMaximumValue int `json:"intWithPositiveMaximumValue,omitempty"` // want "field IntWithPositiveMaximumValue is optional and should be a pointer" "field IntWithPositiveMaximumValue is optional and has a valid zero value \\(0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// intWithPositiveMaximumValueWithoutOmitEmpty with positive maximum value is an int field without omitempty.
	// +kubebuilder:validation:Maximum=1
	// +optional
	IntWithPositiveMaximumValueWithoutOmitEmpty int `json:"intWithPositiveMaximumValueWithoutOmitEmpty"`

	// intWithRange is an int field with a range of values including 0.
	// +kubebuilder:validation:Minimum=-10
	// +kubebuilder:validation:Maximum=10
	// +optional
	IntWithRange int `json:"intWithRange,omitempty"` // want "field IntWithRange is optional and should be a pointer"

	// intWithRangeWithoutOmitEmpty is an int field with a range of values including 0 without omitempty.
	// +kubebuilder:validation:Minimum=-10
	// +kubebuilder:validation:Maximum=10
	// +optional
	IntWithRangeWithoutOmitEmpty int `json:"intWithRangeWithoutOmitEmpty"`

	// intWithInvalidMinimumValue with invalid minimum value is an int field.
	// +kubebuilder:validation:Minimum=foo
	// +optional
	IntWithInvalidMinimumValue int `json:"intWithInvalidMinimumValue,omitempty"` // want "field IntWithInvalidMinimumValue has an invalid minimum marker: error getting marker value: error converting value to number: strconv.ParseFloat: parsing \\\"foo\\\": invalid syntax"

	// intWithInvalidMaximumValue with invalid maximum value is an int field.
	// +kubebuilder:validation:Maximum=foo
	// +optional
	IntWithInvalidMaximumValue int `json:"intWithInvalidMaximumValue,omitempty"` // want "field IntWithInvalidMaximumValue has an invalid maximum marker: error getting marker value: error converting value to number: strconv.ParseFloat: parsing \\\"foo\\\": invalid syntax"

	// float is a float field.
	// +optional
	Float float64 `json:"float,omitempty"` // want "field Float is optional and should be a pointer" "field Float is optional and has a valid zero value \\(0.0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// floatWithoutOmitEmpty is a float field without omitempty.
	// +optional
	FloatWithoutOmitEmpty float64 `json:"floatWithoutOmitEmpty"`

	// floatWithMinValue1 with minimum value is a float field.
	// +kubebuilder:validation:Minimum=1.0
	// +optional
	FloatWithMinValue1 float64 `json:"floatWithMinValue1,omitempty"`

	// floatWithMinValue1WithoutOmitEmpty with minimum value is a float field without omitempty.
	// +kubebuilder:validation:Minimum=1.0
	// +optional
	FloatWithMinValue1WithoutOmitEmpty float64 `json:"floatWithMinValue1WithoutOmitEmpty"` // want "field FloatWithMinValue1WithoutOmitEmpty is optional and does not allow the zero value. It must have the omitempty tag."

	// floatWithMinValue0 with minimum value is a float field.
	// +kubebuilder:validation:Minimum=0.0
	// +optional
	FloatWithMinValue0 float64 `json:"floatWithMinValue0,omitempty"` // want "field FloatWithMinValue0 is optional and should be a pointer"

	// floatWithMinValue0WithoutOmitEmpty with minimum value is a float field without omitempty.
	// +kubebuilder:validation:Minimum=0.0
	// +optional
	FloatWithMinValue0WithoutOmitEmpty float64 `json:"floatWithMinValue0WithoutOmitEmpty"`

	// floatWithNegativeMaximumValue with negative maximum value is a float field.
	// +kubebuilder:validation:Maximum=-1.0
	// +optional
	FloatWithNegativeMaximumValue float64 `json:"floatWithNegativeMaximumValue,omitempty"`

	// floatWithNegativeMaximumValueWithoutOmitEmpty with negative maximum value is a float field without omitempty.
	// +kubebuilder:validation:Maximum=-1.0
	// +optional
	FloatWithNegativeMaximumValueWithoutOmitEmpty float64 `json:"floatWithNegativeMaximumValueWithoutOmitEmpty"` // want "field FloatWithNegativeMaximumValueWithoutOmitEmpty is optional and does not allow the zero value. It must have the omitempty tag."

	// floatWithNegativeMinimumValue with negative minimum value is a float field.
	// +kubebuilder:validation:Minimum=-1.0
	// +optional
	FloatWithNegativeMinimumValue float64 `json:"floatWithNegativeMinimumValue,omitempty"` // want "field FloatWithNegativeMinimumValue is optional and should be a pointer" "field FloatWithNegativeMinimumValue is optional and has a valid zero value \\(0.0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// floatWithNegativeMinimumValueWithoutOmitEmpty with negative minimum value is a float field without omitempty.
	// +kubebuilder:validation:Minimum=-1.0
	// +optional
	FloatWithNegativeMinimumValueWithoutOmitEmpty float64 `json:"floatWithNegativeMinimumValueWithoutOmitEmpty"`

	// floatWithPositiveMaximumValue with positive maximum value is a float field.
	// +kubebuilder:validation:Maximum=1.0
	// +optional
	FloatWithPositiveMaximumValue float64 `json:"floatWithPositiveMaximumValue,omitempty"` // want "field FloatWithPositiveMaximumValue is optional and should be a pointer" "field FloatWithPositiveMaximumValue is optional and has a valid zero value \\(0.0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// floatWithPositiveMaximumValueWithoutOmitEmpty with positive maximum value is a float field without omitempty.
	// +kubebuilder:validation:Maximum=1.0
	// +optional
	FloatWithPositiveMaximumValueWithoutOmitEmpty float64 `json:"floatWithPositiveMaximumValueWithoutOmitEmpty"`

	// floatWithRange is a float field with a range of values including 0.
	// +kubebuilder:validation:Minimum=-10.0
	// +kubebuilder:validation:Maximum=10.0
	// +optional
	FloatWithRange float64 `json:"floatWithRange,omitempty"` // want "field FloatWithRange is optional and should be a pointer"

	// floatWithRangeWithoutOmitEmpty is a float field with a range of values including 0 without omitempty.
	// +kubebuilder:validation:Minimum=-10.0
	// +kubebuilder:validation:Maximum=10.0
	// +optional
	FloatWithRangeWithoutOmitEmpty float64 `json:"floatWithRangeWithoutOmitEmpty"`

	// floatWithInvalidMinimumValue with invalid minimum value is a float field.
	// +kubebuilder:validation:Minimum=foo
	// +optional
	FloatWithInvalidMinimumValue float64 `json:"floatWithInvalidMinimumValue,omitempty"` // want "field FloatWithInvalidMinimumValue has an invalid minimum marker: error getting marker value: error converting value to number: strconv.ParseFloat: parsing \\\"foo\\\": invalid syntax"

	// floatWithInvalidMaximumValue with invalid maximum value is a float field.
	// +kubebuilder:validation:Maximum=foo
	// +optional
	FloatWithInvalidMaximumValue float64 `json:"floatWithInvalidMaximumValue,omitempty"` // want "field FloatWithInvalidMaximumValue has an invalid maximum marker: error getting marker value: error converting value to number: strconv.ParseFloat: parsing \\\"foo\\\": invalid syntax"

	// structWithOptionalFields is a struct field.
	// +optional
	StructWithOptionalFields B `json:"structWithOptionalFields,omitempty"` // want "field StructWithOptionalFields is optional and should be a pointer" "field StructWithOptionalFields is optional and has a valid zero value \\(\\{\\}\\), but the validation is not complete \\(e.g. min properties/adding required fields\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// structWithOptionalFieldsWithoutOmitEmpty is a struct field without omitempty.
	// +optional
	StructWithOptionalFieldsWithoutOmitEmpty B `json:"structWithOptionalFieldsWithoutOmitEmpty"`

	// structWithMinProperties is a struct field with a minimum number of properties.
	// +kubebuilder:validation:MinProperties=1
	// +optional
	StructWithMinProperties B `json:"structWithMinProperties,omitempty"` // want "field StructWithMinProperties is optional and should be a pointer"

	// structWithMinPropertiesWithoutOmitEmpty is a struct field with a minimum number of properties without omitempty.
	// +kubebuilder:validation:MinProperties=1
	// +optional
	StructWithMinPropertiesWithoutOmitEmpty B `json:"structWithMinPropertiesWithoutOmitEmpty"` // want "field StructWithMinPropertiesWithoutOmitEmpty is optional and should be a pointer" "field StructWithMinPropertiesWithoutOmitEmpty is optional and does not allow the zero value. It must have the omitempty tag."

	// structWithMinPropertiesOnStruct is a struct field with a minimum number of properties on the struct.
	// +optional
	StructWithMinPropertiesOnStruct D `json:"structWithMinPropertiesOnStruct,omitempty"` // want "field StructWithMinPropertiesOnStruct is optional and should be a pointer"

	// structWithMinPropertiesOnStructWithoutOmitEmpty is a struct field with a minimum number of properties on the struct without omitempty.
	// +optional
	StructWithMinPropertiesOnStructWithoutOmitEmpty D `json:"structWithMinPropertiesOnStructWithoutOmitEmpty"` // want "field StructWithMinPropertiesOnStructWithoutOmitEmpty is optional and should be a pointer" "field StructWithMinPropertiesOnStructWithoutOmitEmpty is optional and does not allow the zero value. It must have the omitempty tag."

	// structWithRequiredFields is a struct field.
	// +optional
	StructWithRequiredFields C `json:"structWithRequiredFields,omitempty"` // want "field StructWithRequiredFields is optional and should be a pointer"

	// structWithRequiredFieldsWithoutOmitEmpty is a struct field without omitempty.
	// +optional
	StructWithRequiredFieldsWithoutOmitEmpty E `json:"structWithRequiredFieldsWithoutOmitEmpty"`

	// structWithRequiredFieldsWithNonZeroAllowedValuesWithoutOmitEmpty is a struct field with required fields but where the zero values are not valid.
	// +optional
	StructWithRequiredFieldsWithNonZeroAllowedValuesWithoutOmitEmpty C `json:"structWithRequiredFieldsWithNonZeroAllowedValuesWithoutOmitEmpty"` // want "field StructWithRequiredFieldsWithNonZeroAllowedValuesWithoutOmitEmpty is optional and should be a pointer" "field StructWithRequiredFieldsWithNonZeroAllowedValuesWithoutOmitEmpty is optional and does not allow the zero value. It must have the omitempty tag."

	// structWithNonZeroByteArray is a struct field with a non-zero byte array.
	// +optional
	StructWithNonZeroByteArray F `json:"structWithNonZeroByteArray"` // want "field StructWithNonZeroByteArray is optional and should be a pointer" "field StructWithNonZeroByteArray is optional and does not allow the zero value. It must have the omitempty tag."

	// structWithInvalidEmptyEnum is a struct field with an invalid empty enum.
	// +optional
	StructWithInvalidEmptyEnum G `json:"structWithInvalidEmptyEnum"` // want "field StructWithInvalidEmptyEnum is optional and should be a pointer" "field StructWithInvalidEmptyEnum is optional and does not allow the zero value. It must have the omitempty tag."

	// pointerStructWithOptionalFields is a pointer struct field.
	// +optional
	PointerStructWithOptionalFields *B `json:"pointerStructWithOptionalFields,omitempty"`

	// pointerStructWithOptionalFieldsWithoutOmitEmpty is a pointer struct field without omitempty.
	// +optional
	PointerStructWithOptionalFieldsWithoutOmitEmpty *B `json:"pointerStructWithOptionalFieldsWithoutOmitEmpty"` // want "field PointerStructWithOptionalFieldsWithoutOmitEmpty is optional, without omitempty and allows the zero value. The field does not need to be a pointer."

	// pointerStructWithRequiredFields is a pointer struct field.
	// +optional
	PointerStructWithRequiredFields *C `json:"pointerStructWithRequiredFields,omitempty"`

	// pointerStructWithRequiredFieldsWithoutOmitEmpty is a pointer struct field without omitempty.
	// +optional
	PointerStructWithRequiredFieldsWithoutOmitEmpty *C `json:"pointerStructWithRequiredFieldsWithoutOmitEmpty"` // want "field PointerStructWithRequiredFieldsWithoutOmitEmpty is optional and does not allow the zero value. It must have the omitempty tag."

	// pointerStructWithRequiredFromAnotherFile is a pointer struct field.
	// +optional
	PointerStructWithRequiredFromAnotherFile *StructWithRequiredField `json:"pointerStructWithRequiredFromAnotherFile,omitempty"`

	// bool is a boolean field.
	// +optional
	Bool bool `json:"bool,omitempty"` // want "field Bool is optional and should be a pointer"

	// boolWithoutOmitEmpty is a boolean field without omitempty.
	// +optional
	BoolWithoutOmitEmpty bool `json:"boolWithoutOmitEmpty"`

	// boolPointer is a pointer boolean field.
	// +optional
	BoolPointer *bool `json:"boolPointer,omitempty"`

	// boolPointerWithoutOmitEmpty is a pointer boolean field without omitempty.
	// +optional
	BoolPointerWithoutOmitEmpty *bool `json:"boolPointerWithoutOmitEmpty"` // want "field BoolPointerWithoutOmitEmpty is optional, without omitempty and allows the zero value. The field does not need to be a pointer."

	// slice is a slice field.
	// +optional
	Slice []string `json:"slice,omitempty"`

	// sliceWithMinItems is a slice field with a minimum number of items.
	// +kubebuilder:validation:MinItems=1
	// +optional
	SliceWithMinItems []string `json:"sliceWithMinItems,omitempty"`

	// sliceWithMinItemsWithoutOmitEmpty is a slice field with a minimum number of items without omitempty.
	// +kubebuilder:validation:MinItems=1
	// +optional
	SliceWithMinItemsWithoutOmitEmpty []string `json:"sliceWithMinItemsWithoutOmitEmpty"` // want "field SliceWithMinItemsWithoutOmitEmpty is optional and does not allow the zero value. It must have the omitempty tag."

	// map is a map field.
	// +optional
	Map map[string]string `json:"map,omitempty"`

	// mapWithMinProperties is a map field with a minimum number of properties.
	// +kubebuilder:validation:MinProperties=1
	// +optional
	MapWithMinProperties map[string]string `json:"mapWithMinProperties,omitempty"`

	// mapWithMinPropertiesWithoutOmitEmpty is a map field with a minimum number of properties without omitempty.
	// +kubebuilder:validation:MinProperties=1
	// +optional
	MapWithMinPropertiesWithoutOmitEmpty map[string]string `json:"mapWithMinPropertiesWithoutOmitEmpty"` // want "field MapWithMinPropertiesWithoutOmitEmpty is optional and does not allow the zero value. It must have the omitempty tag."

	// PointerSlice is a pointer slice field.
	// +optional
	PointerSlice *[]string `json:"pointerSlice,omitempty"` // want "field PointerSlice is optional but the underlying type does not need to be a pointer. The pointer should be removed."

	// PointerMap is a pointer map field.
	// +optional
	PointerMap *map[string]string `json:"pointerMap,omitempty"` // want "field PointerMap is optional but the underlying type does not need to be a pointer. The pointer should be removed."

	// PointerPointerString is a double pointer string field.
	// +optional
	DoublePointerString **string `json:"doublePointerString,omitempty"` // want "field DoublePointerString is optional but the underlying type does not need to be a pointer. The pointer should be removed."

	// StringAliasWithEnum is a string alias field with enum validation.
	// The zero value ("") is not in the enum, so this should NOT be a pointer.
	// +optional
	StringAliasWithEnum StringAliasWithEnum `json:"stringAliasWithEnum,omitempty"`

	// StringAliasWithEnumPointer is a pointer string alias field with enum validation.
	// This should NOT be a pointer since the zero value is not valid.
	// +optional
	StringAliasWithEnumPointer *StringAliasWithEnum `json:"stringAliasWithEnumPointer,omitempty"` // want "field StringAliasWithEnumPointer is optional and does not allow the zero value. The field does not need to be a pointer."

	// StringAliasWithEnumNoOmitEmpty is a string alias field with enum validation and no omitempty.
	// +optional
	StringAliasWithEnumNoOmitEmpty StringAliasWithEnum `json:"stringAliasWithEnumNoOmitEmpty"` // want "field StringAliasWithEnumNoOmitEmpty is optional and does not allow the zero value. It must have the omitempty tag."

	// StringAliasWithEnumEmptyValue is a string alias field with enum validation and empty value.
	// +optional
	StringAliasWithEnumEmptyValue *StringAliasWithEnumEmptyValue `json:"stringAliasWithEnumEmptyValue,omitempty"`

	// structWithValidOmitZero is a struct field with a minimum number of properties on the struct so not a valid zero value.
	// +optional
	StructWithValidOmitZero *D `json:"structWithValidOmitZero,omitzero,omitempty"` // want "field StructWithValidOmitZero has the omitzero tag, but by policy is not allowed. The omitzero tag should be removed."

	// structWithOnlyOmitZero is a struct field with a minimum number of properties on the struct so not a valid zero value.
	// +optional
	StructWithOnlyOmitZero *D `json:"structWithOnlyOmitZero,omitzero"` // want "field StructWithOnlyOmitZero has the omitzero tag, but by policy is not allowed. The omitzero tag should be removed." "field StructWithOnlyOmitZero is optional and does not allow the zero value. It must have the omitempty tag."
}

type B struct {
	// pointerString is a pointer string field.
	// +kubebuilder:validation:MinLength=0
	// +optional
	PointerString *string `json:"pointerString,omitempty"`
}

type C struct {
	// string is a string field.
	// +required
	// +kubebuilder:validation:MinLength=1
	String string `json:"string"`
}

// +kubebuilder:validation:MinProperties=1
type D struct {
	// string is a string field.
	// +optional
	String string `json:"string,omitempty"` // want "field String is optional and should be a pointer" "field String is optional and has a valid zero value \\(\"\"\\), but the validation is not complete \\(e.g. minimum length\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// stringWithMinLength1 with minimum length is a string field.
	// +kubebuilder:validation:MinLength=1
	// +optional
	StringWithMinLength1 string `json:"stringWithMinLength1,omitempty"`
}

// E is a struct with required fields but where the zero values are valid.
// In this case the struct does not need to be a pointer when there is no omitempty.
type E struct {
	// string is a string field.
	// +required
	// +kubebuilder:validation:MinLength=0
	String string `json:"string"`

	// stringWithMinLength1 with minimum length is a string field.
	// +optional
	// +kubebuilder:validation:MinLength=1
	StringWithMinLength1 string `json:"stringWithMinLength1,omitempty"`

	// enumWithOmitEmpty is an enum field with omitempty.
	// It does not need to allow the zero value.
	// +optional
	// +kubebuilder:validation:Enum=foo;bar
	EnumWithOmitEmpty string `json:"enumWithOmitEmpty,omitempty"`

	// enumWithoutOmitEmpty is an enum field without omitempty.
	// It does need to allow the zero value.
	// +optional
	// +kubebuilder:validation:Enum=foo;bar;""
	EnumWithoutOmitEmpty string `json:"enumWithoutOmitEmpty"`

	// int is an int field.
	// +required
	// +kubebuilder:validation:Minimum=0
	Int int `json:"int"`

	// intWithMinValue1 with minimum value is an int field.
	// +optional
	// +kubebuilder:validation:Minimum=1
	IntWithMinValue1 int `json:"intWithMinValue1,omitempty"`

	// intWithMinValue1WithoutOmitEmpty with minimum value is an int field without omitempty.
	// +optional
	// +kubebuilder:validation:Maximum=1
	IntWithNoMinimum int `json:"intWithNoMinimum"`

	// intWithNoMaximum with no maximum value is an int field.
	// +optional
	// +kubebuilder:validation:Minimum=-1
	IntWithNoMaximum int `json:"intWithNoMaximum"`

	// float is a float field.
	// +required
	// +kubebuilder:validation:Minimum=0.0
	Float float64 `json:"float"`

	// floatWithMinValue1 with minimum value is a float field.
	// +optional
	// +kubebuilder:validation:Minimum=1.0
	FloatWithMinValue1 float64 `json:"floatWithMinValue1,omitempty"`

	// B is a struct with optional fields.
	B B `json:"b"`

	// ByteArry is a byte array field.
	// +optional
	ByteArray []byte `json:"byteArray"`
}

type F struct {
	// nonZeroByteArray is a byte array field that does not allow the zero value.
	// +kubebuilder:validation:MinLength=1
	// +optional
	NonZeroByteArray []byte `json:"nonZeroByteArray"` // want "field NonZeroByteArray is optional and does not allow the zero value. It must have the omitempty tag."
}

type G struct {
	// enumWithoutOmitEmpty is an enum field without omitempty.
	// It does need to allow the zero value, else the parent zero value is not valid.
	// +optional
	// +kubebuilder:validation:Enum=foo;bar
	EnumWithoutOmitEmpty string `json:"enumWithoutOmitEmpty"` // want "field EnumWithoutOmitEmpty is optional and does not allow the zero value. It must have the omitempty tag."
}

// StringAliasWithEnum is a string alias with enum validation.
// The zero value ("") is not in the enum, making it invalid.
// +kubebuilder:validation:Enum=value1;value2
type StringAliasWithEnum string

// StringAliasWithEnumEmptyValue is a string alias with enum validation and empty value.
// +kubebuilder:validation:Enum=value1;value2;""
type StringAliasWithEnumEmptyValue string
