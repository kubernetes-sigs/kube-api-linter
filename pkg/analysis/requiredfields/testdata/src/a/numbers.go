package a

type TestNumbers struct {
	// +required
	Int int `json:"int"` // want "field TestNumbers.Int should have the omitempty tag." "field TestNumbers.Int has a valid zero value \\(0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +required
	IntWithOmitEmpty int `json:"intWithOmitEmpty,omitempty"` // want "field TestNumbers.IntWithOmitEmpty has a valid zero value \\(0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +required
	// +kubebuilder:validation:Minimum=1
	IntWithPositiveMinimum int `json:"intWithPositiveMinimum"` // want "field TestNumbers.IntWithPositiveMinimum should have the omitempty tag."

	// +required
	// +kubebuilder:validation:Minimum=1
	IntWithPositiveMinimumWithOmitEmpty int `json:"intWithPositiveMinimumWithOmitEmpty,omitempty"`

	// +required
	// +kubebuilder:validation:Minimum=0
	IntWithZeroMinimum int `json:"intWithZeroMinimum"` // want "field TestNumbers.IntWithZeroMinimum should have the omitempty tag." "field TestNumbers.IntWithZeroMinimum has a valid zero value \\(0\\) and should be a pointer."

	// +required
	// +kubebuilder:validation:Minimum=0
	IntWithZeroMinimumWithOmitEmpty int `json:"intWithZeroMinimumWithOmitEmpty,omitempty"` // want "field TestNumbers.IntWithZeroMinimumWithOmitEmpty has a valid zero value \\(0\\) and should be a pointer."

	// +required
	// +kubebuilder:validation:Minimum=-1
	IntWithNegativeMinimum int `json:"intWithNegativeMinimum"` // want "field TestNumbers.IntWithNegativeMinimum should have the omitempty tag." "field TestNumbers.IntWithNegativeMinimum has a valid zero value \\(0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +required
	// +kubebuilder:validation:Minimum=-1
	IntWithNegativeMinimumWithOmitEmpty int `json:"intWithNegativeMinimumWithOmitEmpty,omitempty"` // want "field TestNumbers.IntWithNegativeMinimumWithOmitEmpty has a valid zero value \\(0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +required
	// +kubebuilder:validation:Maximum=1
	IntWithPositiveMaximum int `json:"intWithPositiveMaximum"` // want "field TestNumbers.IntWithPositiveMaximum should have the omitempty tag." "field TestNumbers.IntWithPositiveMaximum has a valid zero value \\(0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +required
	// +kubebuilder:validation:Maximum=1
	IntWithPositiveMaximumWithOmitEmpty int `json:"intWithPositiveMaximumWithOmitEmpty,omitempty"` // want "field TestNumbers.IntWithPositiveMaximumWithOmitEmpty has a valid zero value \\(0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +required
	// +kubebuilder:validation:Maximum=0
	IntWithZeroMaximum int `json:"intWithZeroMaximum"` // want "field TestNumbers.IntWithZeroMaximum should have the omitempty tag." "field TestNumbers.IntWithZeroMaximum has a valid zero value \\(0\\) and should be a pointer."

	// +required
	// +kubebuilder:validation:Maximum=0
	IntWithZeroMaximumWithOmitEmpty int `json:"intWithZeroMaximumWithOmitEmpty,omitempty"` // want "field TestNumbers.IntWithZeroMaximumWithOmitEmpty has a valid zero value \\(0\\) and should be a pointer."

	// +required
	// +kubebuilder:validation:Maximum=-1
	IntWithNegativeMaximum int `json:"intWithNegativeMaximum"` // want "field TestNumbers.IntWithNegativeMaximum should have the omitempty tag."

	// +required
	// +kubebuilder:validation:Maximum=-1
	IntWithNegativeMaximumWithOmitEmpty int `json:"intWithNegativeMaximumWithOmitEmpty,omitempty"`

	// +required
	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	IntWithRangeIncludingZero int `json:"intWithRangeIncludingZero"` // want "field TestNumbers.IntWithRangeIncludingZero should have the omitempty tag." "field TestNumbers.IntWithRangeIncludingZero has a valid zero value \\(0\\) and should be a pointer."

	// +required
	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	IntWithRangeIncludingZeroWithOmitEmpty int `json:"intWithRangeIncludingZeroWithOmitEmpty,omitempty"` // want "field TestNumbers.IntWithRangeIncludingZeroWithOmitEmpty has a valid zero value \\(0\\) and should be a pointer."

	// +required
	IntPtr *int `json:"intPtr"` // want "field TestNumbers.IntPtr should have the omitempty tag."

	// +required
	IntPtrWithOmitEmpty *int `json:"intPtrWithOmitEmpty,omitempty"`

	// +required
	// +kubebuilder:validation:Minimum=1
	IntPtrWithPositiveMinimum *int `json:"intPtrWithPositiveMinimum"` // want "field TestNumbers.IntPtrWithPositiveMinimum should have the omitempty tag." "field TestNumbers.IntPtrWithPositiveMinimum does not allow the zero value. The field does not need to be a pointer."

	// +required
	// +kubebuilder:validation:Minimum=1
	IntPtrWithPositiveMinimumWithOmitEmpty *int `json:"intPtrWithPositiveMinimumWithOmitEmpty,omitempty"` // want "field TestNumbers.IntPtrWithPositiveMinimumWithOmitEmpty does not allow the zero value. The field does not need to be a pointer."

	// +required
	// +kubebuilder:validation:Minimum=0
	IntPtrWithZeroMinimum *int `json:"intPtrWithZeroMinimum"` // want "field TestNumbers.IntPtrWithZeroMinimum should have the omitempty tag."

	// +required
	// +kubebuilder:validation:Minimum=0
	IntPtrWithZeroMinimumWithOmitEmpty *int `json:"intPtrWithZeroMinimumWithOmitEmpty,omitempty"`

	// +required
	// +kubebuilder:validation:Minimum=-1
	IntPtrWithNegativeMinimum *int `json:"intPtrWithNegativeMinimum"` // want "field TestNumbers.IntPtrWithNegativeMinimum should have the omitempty tag."

	// +required
	// +kubebuilder:validation:Minimum=-1
	IntPtrWithNegativeMinimumWithOmitEmpty *int `json:"intPtrWithNegativeMinimumWithOmitEmpty,omitempty"`

	// +required
	// +kubebuilder:validation:Maximum=1
	IntPtrWithPositiveMaximum *int `json:"intPtrWithPositiveMaximum"` // want "field TestNumbers.IntPtrWithPositiveMaximum should have the omitempty tag."

	// +required
	// +kubebuilder:validation:Maximum=1
	IntPtrWithPositiveMaximumWithOmitEmpty *int `json:"intPtrWithPositiveMaximumWithOmitEmpty,omitempty"`

	// +required
	// +kubebuilder:validation:Maximum=0
	IntPtrWithZeroMaximum *int `json:"intPtrWithZeroMaximum"` // want "field TestNumbers.IntPtrWithZeroMaximum should have the omitempty tag."

	// +required
	// +kubebuilder:validation:Maximum=0
	IntPtrWithZeroMaximumWithOmitEmpty *int `json:"intPtrWithZeroMaximumWithOmitEmpty,omitempty"`

	// +required
	// +kubebuilder:validation:Maximum=-1
	IntPtrWithNegativeMaximum *int `json:"intPtrWithNegativeMaximum"` // want "field TestNumbers.IntPtrWithNegativeMaximum should have the omitempty tag." "field TestNumbers.IntPtrWithNegativeMaximum does not allow the zero value. The field does not need to be a pointer."

	// +required
	// +kubebuilder:validation:Maximum=-1
	IntPtrWithNegativeMaximumWithOmitEmpty *int `json:"intPtrWithNegativeMaximumWithOmitEmpty,omitempty"` // want "field TestNumbers.IntPtrWithNegativeMaximumWithOmitEmpty does not allow the zero value. The field does not need to be a pointer."

	// +required
	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	IntPtrWithRangeIncludingZero *int `json:"intPtrWithRangeIncludingZero"` // want "field TestNumbers.IntPtrWithRangeIncludingZero should have the omitempty tag."

	// +required
	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	IntPtrWithRangeIncludingZeroWithOmitEmpty *int `json:"intPtrWithRangeIncludingZeroWithOmitEmpty,omitempty"`

	// +required
	Int32 int32 `json:"int32"` // want "field TestNumbers.Int32 should have the omitempty tag." "field TestNumbers.Int32 has a valid zero value \\(0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +required
	// +kubebuilder:validation:Minimum=1
	Int32WithPositiveMinimum int32 `json:"int32WithPositiveMinimum"` // want "field TestNumbers.Int32WithPositiveMinimum should have the omitempty tag."

	// +required
	// +kubebuilder:validation:Minimum=0
	Int32WithZeroMinimum int32 `json:"int32WithZeroMinimum"` // want "field TestNumbers.Int32WithZeroMinimum should have the omitempty tag." "field TestNumbers.Int32WithZeroMinimum has a valid zero value \\(0\\) and should be a pointer."

	// +required
	// +kubebuilder:validation:Minimum=-1
	Int32WithNegativeMinimum int32 `json:"int32WithNegativeMinimum"` // want "field TestNumbers.Int32WithNegativeMinimum should have the omitempty tag." "field TestNumbers.Int32WithNegativeMinimum has a valid zero value \\(0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +required
	// +kubebuilder:validation:Maximum=1
	Int32WithPositiveMaximum int32 `json:"int32WithPositiveMaximum"` // want "field TestNumbers.Int32WithPositiveMaximum should have the omitempty tag." "field TestNumbers.Int32WithPositiveMaximum has a valid zero value \\(0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +required
	// +kubebuilder:validation:Maximum=0
	Int32WithZeroMaximum int32 `json:"int32WithZeroMaximum"` // want "field TestNumbers.Int32WithZeroMaximum should have the omitempty tag." "field TestNumbers.Int32WithZeroMaximum has a valid zero value \\(0\\) and should be a pointer."

	// +required
	// +kubebuilder:validation:Maximum=-1
	Int32WithNegativeMaximum int32 `json:"int32WithNegativeMaximum"` // want "field TestNumbers.Int32WithNegativeMaximum should have the omitempty tag."

	// +required
	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	Int32WithRangeIncludingZero int32 `json:"int32WithRangeIncludingZero"` // want "field TestNumbers.Int32WithRangeIncludingZero should have the omitempty tag." "field TestNumbers.Int32WithRangeIncludingZero has a valid zero value \\(0\\) and should be a pointer."

	// +required
	Int64 int64 `json:"int64"` // want "field TestNumbers.Int64 should have the omitempty tag." "field TestNumbers.Int64 has a valid zero value \\(0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +required
	// +kubebuilder:validation:Minimum=1
	Int64WithPositiveMinimum int64 `json:"int64WithPositiveMinimum"` // want "field TestNumbers.Int64WithPositiveMinimum should have the omitempty tag."

	// +required
	// +kubebuilder:validation:Minimum=0
	Int64WithZeroMinimum int64 `json:"int64WithZeroMinimum"` // want "field TestNumbers.Int64WithZeroMinimum should have the omitempty tag." "field TestNumbers.Int64WithZeroMinimum has a valid zero value \\(0\\) and should be a pointer."

	// +required
	// +kubebuilder:validation:Minimum=-1
	Int64WithNegativeMinimum int64 `json:"int64WithNegativeMinimum"` // want "field TestNumbers.Int64WithNegativeMinimum should have the omitempty tag." "field TestNumbers.Int64WithNegativeMinimum has a valid zero value \\(0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +required
	// +kubebuilder:validation:Maximum=1
	Int64WithPositiveMaximum int64 `json:"int64WithPositiveMaximum"` // want "field TestNumbers.Int64WithPositiveMaximum should have the omitempty tag." "field TestNumbers.Int64WithPositiveMaximum has a valid zero value \\(0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +required
	// +kubebuilder:validation:Maximum=0
	Int64WithZeroMaximum int64 `json:"int64WithZeroMaximum"` // want "field TestNumbers.Int64WithZeroMaximum should have the omitempty tag." "field TestNumbers.Int64WithZeroMaximum has a valid zero value \\(0\\) and should be a pointer."

	// +required
	// +kubebuilder:validation:Maximum=-1
	Int64WithNegativeMaximum int64 `json:"int64WithNegativeMaximum"` // want "field TestNumbers.Int64WithNegativeMaximum should have the omitempty tag."

	// +required
	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	Int64WithRangeIncludingZero int64 `json:"int64WithRangeIncludingZero"` // want "field TestNumbers.Int64WithRangeIncludingZero should have the omitempty tag." "field TestNumbers.Int64WithRangeIncludingZero has a valid zero value \\(0\\) and should be a pointer."

	// +required
	Float32 float32 `json:"float32"` // want "field TestNumbers.Float32 should have the omitempty tag." "field TestNumbers.Float32 has a valid zero value \\(0.0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +required
	Float32WithOmitEmpty float32 `json:"float32WithOmitEmpty,omitempty"` // want "field TestNumbers.Float32WithOmitEmpty has a valid zero value \\(0.0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +required
	// +kubebuilder:validation:Minimum=1
	Float32WithPositiveMinimum float32 `json:"float32WithPositiveMinimum"` // want "field TestNumbers.Float32WithPositiveMinimum should have the omitempty tag."

	// +required
	// +kubebuilder:validation:Minimum=1
	Float32WithPositiveMinimumWithOmitEmpty float32 `json:"float32WithPositiveMinimumWithOmitEmpty,omitempty"`

	// +required
	// +kubebuilder:validation:Minimum=0
	Float32WithZeroMinimum float32 `json:"float32WithZeroMinimum"` // want "field TestNumbers.Float32WithZeroMinimum should have the omitempty tag." "field TestNumbers.Float32WithZeroMinimum has a valid zero value \\(0.0\\) and should be a pointer."

	// +required
	// +kubebuilder:validation:Minimum=0
	Float32WithZeroMinimumWithOmitEmpty float32 `json:"float32WithZeroMinimumWithOmitEmpty,omitempty"` // want "field TestNumbers.Float32WithZeroMinimumWithOmitEmpty has a valid zero value \\(0.0\\) and should be a pointer."

	// +required
	// +kubebuilder:validation:Minimum=-1
	Float32WithNegativeMinimum float32 `json:"float32WithNegativeMinimum"` // want "field TestNumbers.Float32WithNegativeMinimum should have the omitempty tag." "field TestNumbers.Float32WithNegativeMinimum has a valid zero value \\(0.0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +required
	// +kubebuilder:validation:Minimum=-1
	Float32WithNegativeMinimumWithOmitEmpty float32 `json:"float32WithNegativeMinimumWithOmitEmpty,omitempty"` // want "field TestNumbers.Float32WithNegativeMinimumWithOmitEmpty has a valid zero value \\(0.0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +required
	// +kubebuilder:validation:Maximum=1
	Float32WithPositiveMaximum float32 `json:"float32WithPositiveMaximum"` // want "field TestNumbers.Float32WithPositiveMaximum should have the omitempty tag." "field TestNumbers.Float32WithPositiveMaximum has a valid zero value \\(0.0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +required
	// +kubebuilder:validation:Maximum=1
	Float32WithPositiveMaximumWithOmitEmpty float32 `json:"float32WithPositiveMaximumWithOmitEmpty,omitempty"` // want "field TestNumbers.Float32WithPositiveMaximumWithOmitEmpty has a valid zero value \\(0.0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +required
	// +kubebuilder:validation:Maximum=0
	Float32WithZeroMaximum float32 `json:"float32WithZeroMaximum"` // want "field TestNumbers.Float32WithZeroMaximum should have the omitempty tag." "field TestNumbers.Float32WithZeroMaximum has a valid zero value \\(0.0\\) and should be a pointer."

	// +required
	// +kubebuilder:validation:Maximum=0
	Float32WithZeroMaximumWithOmitEmpty float32 `json:"float32WithZeroMaximumWithOmitEmpty,omitempty"` // want "field TestNumbers.Float32WithZeroMaximumWithOmitEmpty has a valid zero value \\(0.0\\) and should be a pointer."

	// +required
	// +kubebuilder:validation:Maximum=-1
	Float32WithNegativeMaximum float32 `json:"float32WithNegativeMaximum"` // want "field TestNumbers.Float32WithNegativeMaximum should have the omitempty tag."

	// +required
	// +kubebuilder:validation:Maximum=-1
	Float32WithNegativeMaximumWithOmitEmpty float32 `json:"float32WithNegativeMaximumWithOmitEmpty,omitempty"`

	// +required
	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	Float32WithRangeIncludingZero float32 `json:"float32WithRangeIncludingZero"` // want "field TestNumbers.Float32WithRangeIncludingZero should have the omitempty tag." "field TestNumbers.Float32WithRangeIncludingZero has a valid zero value \\(0.0\\) and should be a pointer."

	// +required
	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	Float32WithRangeIncludingZeroWithOmitEmpty float32 `json:"float32WithRangeIncludingZeroWithOmitEmpty,omitempty"` // want "field TestNumbers.Float32WithRangeIncludingZeroWithOmitEmpty has a valid zero value \\(0.0\\) and should be a pointer."

	// +required
	Float64 float64 `json:"float64"` // want "field TestNumbers.Float64 should have the omitempty tag." "field TestNumbers.Float64 has a valid zero value \\(0.0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +required
	// +kubebuilder:validation:Minimum=1
	Float64WithPositiveMinimum float64 `json:"float64WithPositiveMinimum"` // want "field TestNumbers.Float64WithPositiveMinimum should have the omitempty tag."

	// +required
	// +kubebuilder:validation:Minimum=0
	Float64WithZeroMinimum float64 `json:"float64WithZeroMinimum"` // want "field TestNumbers.Float64WithZeroMinimum should have the omitempty tag." "field TestNumbers.Float64WithZeroMinimum has a valid zero value \\(0.0\\) and should be a pointer."

	// +required
	// +kubebuilder:validation:Minimum=-1
	Float64WithNegativeMinimum float64 `json:"float64WithNegativeMinimum"` // want "field TestNumbers.Float64WithNegativeMinimum should have the omitempty tag." "field TestNumbers.Float64WithNegativeMinimum has a valid zero value \\(0.0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +required
	// +kubebuilder:validation:Maximum=1
	Float64WithPositiveMaximum float64 `json:"float64WithPositiveMaximum"` // want "field TestNumbers.Float64WithPositiveMaximum should have the omitempty tag." "field TestNumbers.Float64WithPositiveMaximum has a valid zero value \\(0.0\\), but the validation is not complete \\(e.g. minimum/maximum\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// +required
	// +kubebuilder:validation:Maximum=0
	Float64WithZeroMaximum float64 `json:"float64WithZeroMaximum"` // want "field TestNumbers.Float64WithZeroMaximum should have the omitempty tag." "field TestNumbers.Float64WithZeroMaximum has a valid zero value \\(0.0\\) and should be a pointer."

	// +required
	// +kubebuilder:validation:Maximum=-1
	Float64WithNegativeMaximum float64 `json:"float64WithNegativeMaximum"` // want "field TestNumbers.Float64WithNegativeMaximum should have the omitempty tag."

	// +required
	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	Float64WithRangeIncludingZero float64 `json:"float64WithRangeIncludingZero"` // want "field TestNumbers.Float64WithRangeIncludingZero should have the omitempty tag." "field TestNumbers.Float64WithRangeIncludingZero has a valid zero value \\(0.0\\) and should be a pointer."
}
