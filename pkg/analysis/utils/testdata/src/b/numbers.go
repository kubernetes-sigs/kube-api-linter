package b

type ZeroValueTestNumbers struct {
	Int int // want "zero value is valid" "validation is not complete"

	// +kubebuilder:validation:Minimum=1
	IntWithPositiveMinimum int // want "zero value is not valid" "validation is complete"

	// +kubebuilder:validation:Minimum=0
	IntWithZeroMinimum int // want "zero value is valid" "validation is complete"

	// +kubebuilder:validation:Minimum=-1
	IntWithNegativeMinimum int // want "zero value is valid" "validation is not complete"

	// +kubebuilder:validation:Maximum=1
	IntWithPositiveMaximum int // want "zero value is valid" "validation is not complete"

	// +kubebuilder:validation:Maximum=0
	IntWithZeroMaximum int // want "zero value is valid" "validation is complete"

	// +kubebuilder:validation:Maximum=-1
	IntWithNegativeMaximum int // want "zero value is not valid" "validation is complete"

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	IntWithRangeIncludingZero int // want "zero value is valid" "validation is complete"

	IntPtr *int // want "zero value is valid" "validation is not complete"

	// +kubebuilder:validation:Minimum=1
	IntPtrWithPositiveMinimum *int // want "zero value is not valid" "validation is complete"

	// +kubebuilder:validation:Minimum=0
	IntPtrWithZeroMinimum *int // want "zero value is valid" "validation is complete"

	// +kubebuilder:validation:Minimum=-1
	IntPtrWithNegativeMinimum *int // want "zero value is valid" "validation is not complete"

	// +kubebuilder:validation:Maximum=1
	IntPtrWithPositiveMaximum *int // want "zero value is valid" "validation is not complete"

	// +kubebuilder:validation:Maximum=0
	IntPtrWithZeroMaximum *int // want "zero value is valid" "validation is complete"

	// +kubebuilder:validation:Maximum=-1
	IntPtrWithNegativeMaximum *int // want "zero value is not valid" "validation is complete"

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	IntPtrWithRangeIncludingZero *int // want "zero value is valid" "validation is complete"

	Int32 int32 // want "zero value is valid" "validation is not complete"

	// +kubebuilder:validation:Minimum=1
	Int32WithPositiveMinimum int32 // want "zero value is not valid" "validation is complete"

	// +kubebuilder:validation:Minimum=0
	Int32WithZeroMinimum int32 // want "zero value is valid" "validation is complete"

	// +kubebuilder:validation:Minimum=-1
	Int32WithNegativeMinimum int32 // want "zero value is valid" "validation is not complete"

	// +kubebuilder:validation:Maximum=1
	Int32WithPositiveMaximum int32 // want "zero value is valid" "validation is not complete"

	// +kubebuilder:validation:Maximum=0
	Int32WithZeroMaximum int32 // want "zero value is valid" "validation is complete"

	// +kubebuilder:validation:Maximum=-1
	Int32WithNegativeMaximum int32 // want "zero value is not valid" "validation is complete"

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	Int32WithRangeIncludingZero int32 // want "zero value is valid" "validation is complete"

	Int64 int64 // want "zero value is valid" "validation is not complete"

	// +kubebuilder:validation:Minimum=1
	Int64WithPositiveMinimum int64 // want "zero value is not valid" "validation is complete"

	// +kubebuilder:validation:Minimum=0
	Int64WithZeroMinimum int64 // want "zero value is valid" "validation is complete"

	// +kubebuilder:validation:Minimum=-1
	Int64WithNegativeMinimum int64 // want "zero value is valid" "validation is not complete"

	// +kubebuilder:validation:Maximum=1
	Int64WithPositiveMaximum int64 // want "zero value is valid" "validation is not complete"

	// +kubebuilder:validation:Maximum=0
	Int64WithZeroMaximum int64 // want "zero value is valid" "validation is complete"

	// +kubebuilder:validation:Maximum=-1
	Int64WithNegativeMaximum int64 // want "zero value is not valid" "validation is complete"

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	Int64WithRangeIncludingZero int64 // want "zero value is valid" "validation is complete"

	Float32 float32 // want "zero value is valid" "validation is not complete"

	// +kubebuilder:validation:Minimum=1
	Float32WithPositiveMinimum float32 // want "zero value is not valid" "validation is complete"

	// +kubebuilder:validation:Minimum=0
	Float32WithZeroMinimum float32 // want "zero value is valid" "validation is complete"

	// +kubebuilder:validation:Minimum=-1
	Float32WithNegativeMinimum float32 // want "zero value is valid" "validation is not complete"

	// +kubebuilder:validation:Maximum=1
	Float32WithPositiveMaximum float32 // want "zero value is valid" "validation is not complete"

	// +kubebuilder:validation:Maximum=0
	Float32WithZeroMaximum float32 // want "zero value is valid" "validation is complete"

	// +kubebuilder:validation:Maximum=-1
	Float32WithNegativeMaximum float32 // want "zero value is not valid" "validation is complete"

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	Float32WithRangeIncludingZero float32 // want "zero value is valid" "validation is complete"

	Float64 float64 // want "zero value is valid" "validation is not complete"

	// +kubebuilder:validation:Minimum=1
	Float64WithPositiveMinimum float64 // want "zero value is not valid" "validation is complete"

	// +kubebuilder:validation:Minimum=0
	Float64WithZeroMinimum float64 // want "zero value is valid" "validation is complete"

	// +kubebuilder:validation:Minimum=-1
	Float64WithNegativeMinimum float64 // want "zero value is valid" "validation is not complete"

	// +kubebuilder:validation:Maximum=1
	Float64WithPositiveMaximum float64 // want "zero value is valid" "validation is not complete"

	// +kubebuilder:validation:Maximum=0
	Float64WithZeroMaximum float64 // want "zero value is valid" "validation is complete"

	// +kubebuilder:validation:Maximum=-1
	Float64WithNegativeMaximum float64 // want "zero value is not valid" "validation is complete"

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=1
	Float64WithRangeIncludingZero float64 // want "zero value is valid" "validation is complete"
}
