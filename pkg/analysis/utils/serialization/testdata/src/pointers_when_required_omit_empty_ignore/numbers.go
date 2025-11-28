package a

type TestNumbers struct {
	Int int `json:"int"`

	IntWithOmitEmpty int `json:"intWithOmitEmpty,omitempty"` // want "field TestNumbers.IntWithOmitEmpty has a valid zero value \\(0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +kubebuilder:validation:Minimum=1
	IntWithPositiveMinimum int `json:"intWithPositiveMinimum"` // want "field TestNumbers.IntWithPositiveMinimum does not allow the zero value. It must have the omitempty tag."

	// +kubebuilder:validation:Minimum=1
	IntWithPositiveMinimumWithOmitEmpty int `json:"intWithPositiveMinimumWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:Minimum=0
	IntWithZeroMinimum int `json:"intWithZeroMinimum"`

	// +kubebuilder:validation:Minimum=0
	IntWithZeroMinimumWithOmitEmpty int `json:"intWithZeroMinimumWithOmitEmpty,omitempty"` // want "field TestNumbers.IntWithZeroMinimumWithOmitEmpty has a valid zero value \\(0\\) and should be a pointer."

	// +kubebuilder:validation:Minimum=-1
	IntWithNegativeMinimum int `json:"intWithNegativeMinimum"`

	// +kubebuilder:validation:Minimum=-1
	IntWithNegativeMinimumWithOmitEmpty int `json:"intWithNegativeMinimumWithOmitEmpty,omitempty"` // want "field TestNumbers.IntWithNegativeMinimumWithOmitEmpty has a valid zero value \\(0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +kubebuilder:validation:Maximum=1
	IntWithPositiveMaximum int `json:"intWithPositiveMaximum"`

	// +kubebuilder:validation:Maximum=1
	IntWithPositiveMaximumWithOmitEmpty int `json:"intWithPositiveMaximumWithOmitEmpty,omitempty"` // want "field TestNumbers.IntWithPositiveMaximumWithOmitEmpty has a valid zero value \\(0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +kubebuilder:validation:Maximum=0
	IntWithZeroMaximum int `json:"intWithZeroMaximum"`

	// +kubebuilder:validation:Maximum=0
	IntWithZeroMaximumWithOmitEmpty int `json:"intWithZeroMaximumWithOmitEmpty,omitempty"` // want "field TestNumbers.IntWithZeroMaximumWithOmitEmpty has a valid zero value \\(0\\) and should be a pointer."

	// +kubebuilder:validation:Maximum=-1
	IntWithNegativeMaximum int `json:"intWithNegativeMaximum"` // want "field TestNumbers.IntWithNegativeMaximum does not allow the zero value. It must have the omitempty tag."

	// +kubebuilder:validation:Maximum=-1
	IntWithNegativeMaximumWithOmitEmpty int `json:"intWithNegativeMaximumWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	IntWithRangeIncludingZero int `json:"intWithRangeIncludingZero"`

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	IntWithRangeIncludingZeroWithOmitEmpty int `json:"intWithRangeIncludingZeroWithOmitEmpty,omitempty"` // want "field TestNumbers.IntWithRangeIncludingZeroWithOmitEmpty has a valid zero value \\(0\\) and should be a pointer."

	IntPtr *int `json:"intPtr"` // want "field TestNumbers.IntPtr does not have omitempty and allows the zero value. The field does not need to be a pointer."

	IntPtrWithOmitEmpty *int `json:"intPtrWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:Minimum=1
	IntPtrWithPositiveMinimum *int `json:"intPtrWithPositiveMinimum"` // want "field TestNumbers.IntPtrWithPositiveMinimum does not allow the zero value. The field does not need to be a pointer." "field TestNumbers.IntPtrWithPositiveMinimum does not allow the zero value. It must have the omitempty tag."

	// +kubebuilder:validation:Minimum=1
	IntPtrWithPositiveMinimumWithOmitEmpty *int `json:"intPtrWithPositiveMinimumWithOmitEmpty,omitempty"` // want "field TestNumbers.IntPtrWithPositiveMinimumWithOmitEmpty does not allow the zero value. The field does not need to be a pointer."

	// +kubebuilder:validation:Minimum=0
	IntPtrWithZeroMinimum *int `json:"intPtrWithZeroMinimum"` // want "field TestNumbers.IntPtrWithZeroMinimum does not have omitempty and allows the zero value. The field does not need to be a pointer."

	// +kubebuilder:validation:Minimum=0
	IntPtrWithZeroMinimumWithOmitEmpty *int `json:"intPtrWithZeroMinimumWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:Minimum=-1
	IntPtrWithNegativeMinimum *int `json:"intPtrWithNegativeMinimum"` // want "field TestNumbers.IntPtrWithNegativeMinimum does not have omitempty and allows the zero value. The field does not need to be a pointer."

	// +kubebuilder:validation:Minimum=-1
	IntPtrWithNegativeMinimumWithOmitEmpty *int `json:"intPtrWithNegativeMinimumWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:Maximum=1
	IntPtrWithPositiveMaximum *int `json:"intPtrWithPositiveMaximum"` // want "field TestNumbers.IntPtrWithPositiveMaximum does not have omitempty and allows the zero value. The field does not need to be a pointer."

	// +kubebuilder:validation:Maximum=1
	IntPtrWithPositiveMaximumWithOmitEmpty *int `json:"intPtrWithPositiveMaximumWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:Maximum=0
	IntPtrWithZeroMaximum *int `json:"intPtrWithZeroMaximum"` // want "field TestNumbers.IntPtrWithZeroMaximum does not have omitempty and allows the zero value. The field does not need to be a pointer."

	// +kubebuilder:validation:Maximum=0
	IntPtrWithZeroMaximumWithOmitEmpty *int `json:"intPtrWithZeroMaximumWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:Maximum=-1
	IntPtrWithNegativeMaximum *int `json:"intPtrWithNegativeMaximum"` // want "field TestNumbers.IntPtrWithNegativeMaximum does not allow the zero value. It must have the omitempty tag." "field TestNumbers.IntPtrWithNegativeMaximum does not allow the zero value. The field does not need to be a pointer."

	// +kubebuilder:validation:Maximum=-1
	IntPtrWithNegativeMaximumWithOmitEmpty *int `json:"intPtrWithNegativeMaximumWithOmitEmpty,omitempty"` // want "field TestNumbers.IntPtrWithNegativeMaximumWithOmitEmpty does not allow the zero value. The field does not need to be a pointer."

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	IntPtrWithRangeIncludingZero *int `json:"intPtrWithRangeIncludingZero"` // want "field TestNumbers.IntPtrWithRangeIncludingZero does not have omitempty and allows the zero value. The field does not need to be a pointer."

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	IntPtrWithRangeIncludingZeroWithOmitEmpty *int `json:"intPtrWithRangeIncludingZeroWithOmitEmpty,omitempty"`

	Int32 int32 `json:"int32"`

	// +kubebuilder:validation:Minimum=1
	Int32WithPositiveMinimum int32 `json:"int32WithPositiveMinimum"` // want "field TestNumbers.Int32WithPositiveMinimum does not allow the zero value. It must have the omitempty tag."

	// +kubebuilder:validation:Minimum=0
	Int32WithZeroMinimum int32 `json:"int32WithZeroMinimum"`

	// +kubebuilder:validation:Minimum=-1
	Int32WithNegativeMinimum int32 `json:"int32WithNegativeMinimum"`

	// +kubebuilder:validation:Maximum=1
	Int32WithPositiveMaximum int32 `json:"int32WithPositiveMaximum"`

	// +kubebuilder:validation:Maximum=0
	Int32WithZeroMaximum int32 `json:"int32WithZeroMaximum"`

	// +kubebuilder:validation:Maximum=-1
	Int32WithNegativeMaximum int32 `json:"int32WithNegativeMaximum"` // want "field TestNumbers.Int32WithNegativeMaximum does not allow the zero value. It must have the omitempty tag."

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	Int32WithRangeIncludingZero int32 `json:"int32WithRangeIncludingZero"`

	Int64 int64 `json:"int64"`

	// +kubebuilder:validation:Minimum=1
	Int64WithPositiveMinimum int64 `json:"int64WithPositiveMinimum"` // want "field TestNumbers.Int64WithPositiveMinimum does not allow the zero value. It must have the omitempty tag."

	// +kubebuilder:validation:Minimum=0
	Int64WithZeroMinimum int64 `json:"int64WithZeroMinimum"`

	// +kubebuilder:validation:Minimum=-1
	Int64WithNegativeMinimum int64 `json:"int64WithNegativeMinimum"`

	// +kubebuilder:validation:Maximum=1
	Int64WithPositiveMaximum int64 `json:"int64WithPositiveMaximum"`

	// +kubebuilder:validation:Maximum=0
	Int64WithZeroMaximum int64 `json:"int64WithZeroMaximum"`

	// +kubebuilder:validation:Maximum=-1
	Int64WithNegativeMaximum int64 `json:"int64WithNegativeMaximum"` // want "field TestNumbers.Int64WithNegativeMaximum does not allow the zero value. It must have the omitempty tag."

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	Int64WithRangeIncludingZero int64 `json:"int64WithRangeIncludingZero"`

	Float32 float32 `json:"float32"`

	Float32WithOmitEmpty float32 `json:"float32WithOmitEmpty,omitempty"` // want "field TestNumbers.Float32WithOmitEmpty has a valid zero value \\(0.0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +kubebuilder:validation:Minimum=1
	Float32WithPositiveMinimum float32 `json:"float32WithPositiveMinimum"` // want "field TestNumbers.Float32WithPositiveMinimum does not allow the zero value. It must have the omitempty tag."

	// +kubebuilder:validation:Minimum=1
	Float32WithPositiveMinimumWithOmitEmpty float32 `json:"float32WithPositiveMinimumWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:Minimum=0
	Float32WithZeroMinimum float32 `json:"float32WithZeroMinimum"`

	// +kubebuilder:validation:Minimum=0
	Float32WithZeroMinimumWithOmitEmpty float32 `json:"float32WithZeroMinimumWithOmitEmpty,omitempty"` // want "field TestNumbers.Float32WithZeroMinimumWithOmitEmpty has a valid zero value \\(0.0\\) and should be a pointer."

	// +kubebuilder:validation:Minimum=-1
	Float32WithNegativeMinimum float32 `json:"float32WithNegativeMinimum"`

	// +kubebuilder:validation:Minimum=-1
	Float32WithNegativeMinimumWithOmitEmpty float32 `json:"float32WithNegativeMinimumWithOmitEmpty,omitempty"` // want "field TestNumbers.Float32WithNegativeMinimumWithOmitEmpty has a valid zero value \\(0.0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +kubebuilder:validation:Maximum=1
	Float32WithPositiveMaximum float32 `json:"float32WithPositiveMaximum"`

	// +kubebuilder:validation:Maximum=1
	Float32WithPositiveMaximumWithOmitEmpty float32 `json:"float32WithPositiveMaximumWithOmitEmpty,omitempty"` // want "field TestNumbers.Float32WithPositiveMaximumWithOmitEmpty has a valid zero value \\(0.0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +kubebuilder:validation:Maximum=0
	Float32WithZeroMaximum float32 `json:"float32WithZeroMaximum"`

	// +kubebuilder:validation:Maximum=0
	Float32WithZeroMaximumWithOmitEmpty float32 `json:"float32WithZeroMaximumWithOmitEmpty,omitempty"` // want "field TestNumbers.Float32WithZeroMaximumWithOmitEmpty has a valid zero value \\(0.0\\) and should be a pointer."

	// +kubebuilder:validation:Maximum=-1
	Float32WithNegativeMaximum float32 `json:"float32WithNegativeMaximum"` // want "field TestNumbers.Float32WithNegativeMaximum does not allow the zero value. It must have the omitempty tag."

	// +kubebuilder:validation:Maximum=-1
	Float32WithNegativeMaximumWithOmitEmpty float32 `json:"float32WithNegativeMaximumWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	Float32WithRangeIncludingZero float32 `json:"float32WithRangeIncludingZero"`

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	Float32WithRangeIncludingZeroWithOmitEmpty float32 `json:"float32WithRangeIncludingZeroWithOmitEmpty,omitempty"` // want "field TestNumbers.Float32WithRangeIncludingZeroWithOmitEmpty has a valid zero value \\(0.0\\) and should be a pointer."

	Float64 float64 `json:"float64"`

	// +kubebuilder:validation:Minimum=1
	Float64WithPositiveMinimum float64 `json:"float64WithPositiveMinimum"` // want "field TestNumbers.Float64WithPositiveMinimum does not allow the zero value. It must have the omitempty tag."

	// +kubebuilder:validation:Minimum=0
	Float64WithZeroMinimum float64 `json:"float64WithZeroMinimum"`

	// +kubebuilder:validation:Minimum=-1
	Float64WithNegativeMinimum float64 `json:"float64WithNegativeMinimum"`

	// +kubebuilder:validation:Maximum=1
	Float64WithPositiveMaximum float64 `json:"float64WithPositiveMaximum"`

	// +kubebuilder:validation:Maximum=0
	Float64WithZeroMaximum float64 `json:"float64WithZeroMaximum"`

	// +kubebuilder:validation:Maximum=-1
	Float64WithNegativeMaximum float64 `json:"float64WithNegativeMaximum"` // want "field TestNumbers.Float64WithNegativeMaximum does not allow the zero value. It must have the omitempty tag."

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	Float64WithRangeIncludingZero float64 `json:"float64WithRangeIncludingZero"`
}
