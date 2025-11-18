package a

type TestNumbers struct {
	Int int `json:"int"` // want "field TestNumbers.Int should have the omitempty tag." "field TestNumbers.Int should be a pointer."

	IntWithOmitEmpty int `json:"intWithOmitEmpty,omitempty"` // want "field TestNumbers.IntWithOmitEmpty should be a pointer."

	// +kubebuilder:validation:Minimum=1
	IntWithPositiveMinimum int `json:"intWithPositiveMinimum"` // want "field TestNumbers.IntWithPositiveMinimum should have the omitempty tag." "field TestNumbers.IntWithPositiveMinimum should be a pointer."

	// +kubebuilder:validation:Minimum=1
	IntWithPositiveMinimumWithOmitEmpty int `json:"intWithPositiveMinimumWithOmitEmpty,omitempty"` // want "field TestNumbers.IntWithPositiveMinimumWithOmitEmpty should be a pointer."

	// +kubebuilder:validation:Minimum=0
	IntWithZeroMinimum int `json:"intWithZeroMinimum"` // want "field TestNumbers.IntWithZeroMinimum should have the omitempty tag." "field TestNumbers.IntWithZeroMinimum should be a pointer."

	// +kubebuilder:validation:Minimum=0
	IntWithZeroMinimumWithOmitEmpty int `json:"intWithZeroMinimumWithOmitEmpty,omitempty"` // want "field TestNumbers.IntWithZeroMinimumWithOmitEmpty should be a pointer."

	// +kubebuilder:validation:Minimum=-1
	IntWithNegativeMinimum int `json:"intWithNegativeMinimum"` // want "field TestNumbers.IntWithNegativeMinimum should have the omitempty tag." "field TestNumbers.IntWithNegativeMinimum should be a pointer."

	// +kubebuilder:validation:Minimum=-1
	IntWithNegativeMinimumWithOmitEmpty int `json:"intWithNegativeMinimumWithOmitEmpty,omitempty"` // want "field TestNumbers.IntWithNegativeMinimumWithOmitEmpty should be a pointer."

	// +kubebuilder:validation:Maximum=1
	IntWithPositiveMaximum int `json:"intWithPositiveMaximum"` // want "field TestNumbers.IntWithPositiveMaximum should have the omitempty tag." "field TestNumbers.IntWithPositiveMaximum should be a pointer."

	// +kubebuilder:validation:Maximum=1
	IntWithPositiveMaximumWithOmitEmpty int `json:"intWithPositiveMaximumWithOmitEmpty,omitempty"` // want "field TestNumbers.IntWithPositiveMaximumWithOmitEmpty should be a pointer."

	// +kubebuilder:validation:Maximum=0
	IntWithZeroMaximum int `json:"intWithZeroMaximum"` // want "field TestNumbers.IntWithZeroMaximum should have the omitempty tag." "field TestNumbers.IntWithZeroMaximum should be a pointer."

	// +kubebuilder:validation:Maximum=0
	IntWithZeroMaximumWithOmitEmpty int `json:"intWithZeroMaximumWithOmitEmpty,omitempty"` // want "field TestNumbers.IntWithZeroMaximumWithOmitEmpty should be a pointer."

	// +kubebuilder:validation:Maximum=-1
	IntWithNegativeMaximum int `json:"intWithNegativeMaximum"` // want "field TestNumbers.IntWithNegativeMaximum should have the omitempty tag." "field TestNumbers.IntWithNegativeMaximum should be a pointer."

	// +kubebuilder:validation:Maximum=-1
	IntWithNegativeMaximumWithOmitEmpty int `json:"intWithNegativeMaximumWithOmitEmpty,omitempty"` // want "field TestNumbers.IntWithNegativeMaximumWithOmitEmpty should be a pointer."

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	IntWithRangeIncludingZero int `json:"intWithRangeIncludingZero"` // want "field TestNumbers.IntWithRangeIncludingZero should have the omitempty tag." "field TestNumbers.IntWithRangeIncludingZero should be a pointer."

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	IntWithRangeIncludingZeroWithOmitEmpty int `json:"intWithRangeIncludingZeroWithOmitEmpty,omitempty"` // want "field TestNumbers.IntWithRangeIncludingZeroWithOmitEmpty should be a pointer."

	IntPtr *int `json:"intPtr"` // want "field TestNumbers.IntPtr should have the omitempty tag."

	IntPtrWithOmitEmpty *int `json:"intPtrWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:Minimum=1
	IntPtrWithPositiveMinimum *int `json:"intPtrWithPositiveMinimum"` // want "field TestNumbers.IntPtrWithPositiveMinimum should have the omitempty tag."

	// +kubebuilder:validation:Minimum=1
	IntPtrWithPositiveMinimumWithOmitEmpty *int `json:"intPtrWithPositiveMinimumWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:Minimum=0
	IntPtrWithZeroMinimum *int `json:"intPtrWithZeroMinimum"` // want "field TestNumbers.IntPtrWithZeroMinimum should have the omitempty tag."

	// +kubebuilder:validation:Minimum=0
	IntPtrWithZeroMinimumWithOmitEmpty *int `json:"intPtrWithZeroMinimumWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:Minimum=-1
	IntPtrWithNegativeMinimum *int `json:"intPtrWithNegativeMinimum"` // want "field TestNumbers.IntPtrWithNegativeMinimum should have the omitempty tag."

	// +kubebuilder:validation:Minimum=-1
	IntPtrWithNegativeMinimumWithOmitEmpty *int `json:"intPtrWithNegativeMinimumWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:Maximum=1
	IntPtrWithPositiveMaximum *int `json:"intPtrWithPositiveMaximum"` // want "field TestNumbers.IntPtrWithPositiveMaximum should have the omitempty tag."

	// +kubebuilder:validation:Maximum=1
	IntPtrWithPositiveMaximumWithOmitEmpty *int `json:"intPtrWithPositiveMaximumWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:Maximum=0
	IntPtrWithZeroMaximum *int `json:"intPtrWithZeroMaximum"` // want "field TestNumbers.IntPtrWithZeroMaximum should have the omitempty tag."

	// +kubebuilder:validation:Maximum=0
	IntPtrWithZeroMaximumWithOmitEmpty *int `json:"intPtrWithZeroMaximumWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:Maximum=-1
	IntPtrWithNegativeMaximum *int `json:"intPtrWithNegativeMaximum"` // want "field TestNumbers.IntPtrWithNegativeMaximum should have the omitempty tag."

	// +kubebuilder:validation:Maximum=-1
	IntPtrWithNegativeMaximumWithOmitEmpty *int `json:"intPtrWithNegativeMaximumWithOmitEmpty,omitempty"`

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	IntPtrWithRangeIncludingZero *int `json:"intPtrWithRangeIncludingZero"` // want "field TestNumbers.IntPtrWithRangeIncludingZero should have the omitempty tag."

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	IntPtrWithRangeIncludingZeroWithOmitEmpty *int `json:"intPtrWithRangeIncludingZeroWithOmitEmpty,omitempty"`

	Int32 int32 `json:"int32"` // want "field TestNumbers.Int32 should have the omitempty tag." "field TestNumbers.Int32 should be a pointer."

	// +kubebuilder:validation:Minimum=1
	Int32WithPositiveMinimum int32 `json:"int32WithPositiveMinimum"` // want "field TestNumbers.Int32WithPositiveMinimum should have the omitempty tag." "field TestNumbers.Int32WithPositiveMinimum should be a pointer."

	// +kubebuilder:validation:Minimum=0
	Int32WithZeroMinimum int32 `json:"int32WithZeroMinimum"` // want "field TestNumbers.Int32WithZeroMinimum should have the omitempty tag." "field TestNumbers.Int32WithZeroMinimum should be a pointer."

	// +kubebuilder:validation:Minimum=-1
	Int32WithNegativeMinimum int32 `json:"int32WithNegativeMinimum"` // want "field TestNumbers.Int32WithNegativeMinimum should have the omitempty tag." "field TestNumbers.Int32WithNegativeMinimum should be a pointer."

	// +kubebuilder:validation:Maximum=1
	Int32WithPositiveMaximum int32 `json:"int32WithPositiveMaximum"` // want "field TestNumbers.Int32WithPositiveMaximum should have the omitempty tag." "field TestNumbers.Int32WithPositiveMaximum should be a pointer."

	// +kubebuilder:validation:Maximum=0
	Int32WithZeroMaximum int32 `json:"int32WithZeroMaximum"` // want "field TestNumbers.Int32WithZeroMaximum should have the omitempty tag." "field TestNumbers.Int32WithZeroMaximum should be a pointer."

	// +kubebuilder:validation:Maximum=-1
	Int32WithNegativeMaximum int32 `json:"int32WithNegativeMaximum"` // want "field TestNumbers.Int32WithNegativeMaximum should have the omitempty tag." "field TestNumbers.Int32WithNegativeMaximum should be a pointer."

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	Int32WithRangeIncludingZero int32 `json:"int32WithRangeIncludingZero"` // want "field TestNumbers.Int32WithRangeIncludingZero should have the omitempty tag." "field TestNumbers.Int32WithRangeIncludingZero should be a pointer."

	Int64 int64 `json:"int64"` // want "field TestNumbers.Int64 should have the omitempty tag." "field TestNumbers.Int64 should be a pointer."

	// +kubebuilder:validation:Minimum=1
	Int64WithPositiveMinimum int64 `json:"int64WithPositiveMinimum"` // want "field TestNumbers.Int64WithPositiveMinimum should have the omitempty tag." "field TestNumbers.Int64WithPositiveMinimum should be a pointer."

	// +kubebuilder:validation:Minimum=0
	Int64WithZeroMinimum int64 `json:"int64WithZeroMinimum"` // want "field TestNumbers.Int64WithZeroMinimum should have the omitempty tag." "field TestNumbers.Int64WithZeroMinimum should be a pointer."

	// +kubebuilder:validation:Minimum=-1
	Int64WithNegativeMinimum int64 `json:"int64WithNegativeMinimum"` // want "field TestNumbers.Int64WithNegativeMinimum should have the omitempty tag." "field TestNumbers.Int64WithNegativeMinimum should be a pointer."

	// +kubebuilder:validation:Maximum=1
	Int64WithPositiveMaximum int64 `json:"int64WithPositiveMaximum"` // want "field TestNumbers.Int64WithPositiveMaximum should have the omitempty tag." "field TestNumbers.Int64WithPositiveMaximum should be a pointer."

	// +kubebuilder:validation:Maximum=0
	Int64WithZeroMaximum int64 `json:"int64WithZeroMaximum"` // want "field TestNumbers.Int64WithZeroMaximum should have the omitempty tag." "field TestNumbers.Int64WithZeroMaximum should be a pointer."

	// +kubebuilder:validation:Maximum=-1
	Int64WithNegativeMaximum int64 `json:"int64WithNegativeMaximum"` // want "field TestNumbers.Int64WithNegativeMaximum should have the omitempty tag." "field TestNumbers.Int64WithNegativeMaximum should be a pointer."

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	Int64WithRangeIncludingZero int64 `json:"int64WithRangeIncludingZero"` // want "field TestNumbers.Int64WithRangeIncludingZero should have the omitempty tag." "field TestNumbers.Int64WithRangeIncludingZero should be a pointer."

	Float32 float32 `json:"float32"` // want "field TestNumbers.Float32 should have the omitempty tag." "field TestNumbers.Float32 should be a pointer."

	Float32WithOmitEmpty float32 `json:"float32WithOmitEmpty,omitempty"` // want "field TestNumbers.Float32WithOmitEmpty should be a pointer."

	// +kubebuilder:validation:Minimum=1
	Float32WithPositiveMinimum float32 `json:"float32WithPositiveMinimum"` // want "field TestNumbers.Float32WithPositiveMinimum should have the omitempty tag." "field TestNumbers.Float32WithPositiveMinimum should be a pointer."

	// +kubebuilder:validation:Minimum=1
	Float32WithPositiveMinimumWithOmitEmpty float32 `json:"float32WithPositiveMinimumWithOmitEmpty,omitempty"` // want "field TestNumbers.Float32WithPositiveMinimumWithOmitEmpty should be a pointer."

	// +kubebuilder:validation:Minimum=0
	Float32WithZeroMinimum float32 `json:"float32WithZeroMinimum"` // want "field TestNumbers.Float32WithZeroMinimum should have the omitempty tag." "field TestNumbers.Float32WithZeroMinimum should be a pointer."

	// +kubebuilder:validation:Minimum=0
	Float32WithZeroMinimumWithOmitEmpty float32 `json:"float32WithZeroMinimumWithOmitEmpty,omitempty"` // want "field TestNumbers.Float32WithZeroMinimumWithOmitEmpty should be a pointer."

	// +kubebuilder:validation:Minimum=-1
	Float32WithNegativeMinimum float32 `json:"float32WithNegativeMinimum"` // want "field TestNumbers.Float32WithNegativeMinimum should have the omitempty tag." "field TestNumbers.Float32WithNegativeMinimum should be a pointer."

	// +kubebuilder:validation:Minimum=-1
	Float32WithNegativeMinimumWithOmitEmpty float32 `json:"float32WithNegativeMinimumWithOmitEmpty,omitempty"` // want "field TestNumbers.Float32WithNegativeMinimumWithOmitEmpty should be a pointer."

	// +kubebuilder:validation:Maximum=1
	Float32WithPositiveMaximum float32 `json:"float32WithPositiveMaximum"` // want "field TestNumbers.Float32WithPositiveMaximum should have the omitempty tag." "field TestNumbers.Float32WithPositiveMaximum should be a pointer."

	// +kubebuilder:validation:Maximum=1
	Float32WithPositiveMaximumWithOmitEmpty float32 `json:"float32WithPositiveMaximumWithOmitEmpty,omitempty"` // want "field TestNumbers.Float32WithPositiveMaximumWithOmitEmpty should be a pointer."

	// +kubebuilder:validation:Maximum=0
	Float32WithZeroMaximum float32 `json:"float32WithZeroMaximum"` // want "field TestNumbers.Float32WithZeroMaximum should have the omitempty tag." "field TestNumbers.Float32WithZeroMaximum should be a pointer."

	// +kubebuilder:validation:Maximum=0
	Float32WithZeroMaximumWithOmitEmpty float32 `json:"float32WithZeroMaximumWithOmitEmpty,omitempty"` // want "field TestNumbers.Float32WithZeroMaximumWithOmitEmpty should be a pointer."

	// +kubebuilder:validation:Maximum=-1
	Float32WithNegativeMaximum float32 `json:"float32WithNegativeMaximum"` // want "field TestNumbers.Float32WithNegativeMaximum should have the omitempty tag." "field TestNumbers.Float32WithNegativeMaximum should be a pointer."

	// +kubebuilder:validation:Maximum=-1
	Float32WithNegativeMaximumWithOmitEmpty float32 `json:"float32WithNegativeMaximumWithOmitEmpty,omitempty"` // want "field TestNumbers.Float32WithNegativeMaximumWithOmitEmpty should be a pointer."

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	Float32WithRangeIncludingZero float32 `json:"float32WithRangeIncludingZero"` // want "field TestNumbers.Float32WithRangeIncludingZero should have the omitempty tag." "field TestNumbers.Float32WithRangeIncludingZero should be a pointer."

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	Float32WithRangeIncludingZeroWithOmitEmpty float32 `json:"float32WithRangeIncludingZeroWithOmitEmpty,omitempty"` // want "field TestNumbers.Float32WithRangeIncludingZeroWithOmitEmpty should be a pointer."

	Float64 float64 `json:"float64"` // want "field TestNumbers.Float64 should have the omitempty tag." "field TestNumbers.Float64 should be a pointer."

	// +kubebuilder:validation:Minimum=1
	Float64WithPositiveMinimum float64 `json:"float64WithPositiveMinimum"` // want "field TestNumbers.Float64WithPositiveMinimum should have the omitempty tag." "field TestNumbers.Float64WithPositiveMinimum should be a pointer."

	// +kubebuilder:validation:Minimum=0
	Float64WithZeroMinimum float64 `json:"float64WithZeroMinimum"` // want "field TestNumbers.Float64WithZeroMinimum should have the omitempty tag." "field TestNumbers.Float64WithZeroMinimum should be a pointer."

	// +kubebuilder:validation:Minimum=-1
	Float64WithNegativeMinimum float64 `json:"float64WithNegativeMinimum"` // want "field TestNumbers.Float64WithNegativeMinimum should have the omitempty tag." "field TestNumbers.Float64WithNegativeMinimum should be a pointer."

	// +kubebuilder:validation:Maximum=1
	Float64WithPositiveMaximum float64 `json:"float64WithPositiveMaximum"` // want "field TestNumbers.Float64WithPositiveMaximum should have the omitempty tag." "field TestNumbers.Float64WithPositiveMaximum should be a pointer."

	// +kubebuilder:validation:Maximum=0
	Float64WithZeroMaximum float64 `json:"float64WithZeroMaximum"` // want "field TestNumbers.Float64WithZeroMaximum should have the omitempty tag." "field TestNumbers.Float64WithZeroMaximum should be a pointer."

	// +kubebuilder:validation:Maximum=-1
	Float64WithNegativeMaximum float64 `json:"float64WithNegativeMaximum"` // want "field TestNumbers.Float64WithNegativeMaximum should have the omitempty tag." "field TestNumbers.Float64WithNegativeMaximum should be a pointer."

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	Float64WithRangeIncludingZero float64 `json:"float64WithRangeIncludingZero"` // want "field TestNumbers.Float64WithRangeIncludingZero should have the omitempty tag." "field TestNumbers.Float64WithRangeIncludingZero should be a pointer."
}
