package a

// ValidInt32WithBound has proper bounds validation
type ValidInt32WithBound struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	Count int32
}

// ValidInt64WithBound has proper bounds validation
type ValidInt64WithBound struct {
	// +kubebuilder:validation:Minimum=-1000
	// +kubebuilder:validation:Maximum=1000
	Value int64
}

// ValidInt64WithJSSafeBound has bounds within JavaScript safe integer range
type ValidInt64WithJSSafeBound struct {
	// +kubebuilder:validation:Minimum=-9007199254740991
	// +kubebuilder:validation:Maximum=9007199254740991
	SafeValue int64
}

// InvalidInt32NoBound should have bounds markers
type InvalidInt32NoBound struct {
	NoBound int32 // want "InvalidInt32NoBound.NoBound is missing minimum bound validation marker" "InvalidInt32NoBound.NoBound is missing maximum bound validation marker"
}

// InvalidInt64NoBound should have bounds markers
type InvalidInt64NoBound struct {
	NoBound int64 // want "InvalidInt64NoBound.NoBound is missing minimum bound validation marker" "InvalidInt64NoBound.NoBound is missing maximum bound validation marker"
}

// InvalidInt32OnlyMin should have maximum marker
type InvalidInt32OnlyMin struct {
	// +kubebuilder:validation:Minimum=0
	OnlyMin int32 // want "InvalidInt32OnlyMin.OnlyMin is missing maximum bound validation marker"
}

// InvalidInt32OnlyMax should have minimum marker
type InvalidInt32OnlyMax struct {
	// +kubebuilder:validation:Maximum=100
	OnlyMax int32 // want "InvalidInt32OnlyMax.OnlyMax is missing minimum bound validation marker"
}

// InvalidInt64OnlyMin should have maximum marker
type InvalidInt64OnlyMin struct {
	// +kubebuilder:validation:Minimum=0
	OnlyMin int64 // want "InvalidInt64OnlyMin.OnlyMin is missing maximum bound validation marker"
}

// InvalidInt64OnlyMax should have minimum marker
type InvalidInt64OnlyMax struct {
	// +kubebuilder:validation:Maximum=100
	OnlyMax int64 // want "InvalidInt64OnlyMax.OnlyMax is missing minimum bound validation marker"
}

// InvalidInt64ExceedsJSMaxBounds has maximum that exceeds JavaScript safe integer range
type InvalidInt64ExceedsJSMaxBounds struct {
	// +kubebuilder:validation:Minimum=-1000
	// +kubebuilder:validation:Maximum=9007199254740992
	UnsafeMax int64 // want "InvalidInt64ExceedsJSMaxBounds.UnsafeMax has maximum bound 9\\.007199254740992e\\+15 that is outside the JavaScript-safe int64 range \\[-9007199254740991, 9007199254740991\\]\\. Consider using a string type to avoid precision loss in JavaScript clients"
}

// InvalidInt64ExceedsJSMinBounds has minimum that exceeds JavaScript safe integer range
type InvalidInt64ExceedsJSMinBounds struct {
	// +kubebuilder:validation:Minimum=-9007199254740992
	// +kubebuilder:validation:Maximum=1000
	UnsafeMin int64 // want "InvalidInt64ExceedsJSMinBounds.UnsafeMin has minimum bound -9\\.007199254740992e\\+15 that is outside the JavaScript-safe int64 range \\[-9007199254740991, 9007199254740991\\]\\. Consider using a string type to avoid precision loss in JavaScript clients"
}

// InvalidInt64ExceedsJSBothBounds has both bounds exceeding JavaScript safe integer range
type InvalidInt64ExceedsJSBothBounds struct {
	// +kubebuilder:validation:Minimum=-9007199254740992
	// +kubebuilder:validation:Maximum=9007199254740992
	UnsafeBoth int64 // want "InvalidInt64ExceedsJSBothBounds.UnsafeBoth has minimum bound -9\\.007199254740992e\\+15 that is outside the JavaScript-safe int64 range \\[-9007199254740991, 9007199254740991\\]\\. Consider using a string type to avoid precision loss in JavaScript clients" "InvalidInt64ExceedsJSBothBounds.UnsafeBoth has maximum bound 9\\.007199254740992e\\+15 that is outside the JavaScript-safe int64 range \\[-9007199254740991, 9007199254740991\\]\\. Consider using a string type to avoid precision loss in JavaScript clients"
}

// IgnoredStringField should not be checked
type IgnoredStringField struct {
	Name string
}

// IgnoredBoolField should not be checked
type IgnoredBoolField struct {
	Enabled bool
}

// ValidFloat64WithBounds has proper bounds validation
type ValidFloat64WithBounds struct {
	// +kubebuilder:validation:Minimum=-1000.5
	// +kubebuilder:validation:Maximum=1000.5
	Value float64
}

// ValidFloat32WithBounds has proper bounds validation
type ValidFloat32WithBounds struct {
	// +kubebuilder:validation:Minimum=0.0
	// +kubebuilder:validation:Maximum=100.0
	Ratio float32
}

// InvalidFloat64NoBounds should have bounds markers
type InvalidFloat64NoBounds struct {
	Value float64 // want "InvalidFloat64NoBounds.Value is missing minimum bound validation marker" "InvalidFloat64NoBounds.Value is missing maximum bound validation marker"
}

// InvalidFloat32NoBounds should have bounds markers
type InvalidFloat32NoBounds struct {
	Ratio float32 // want "InvalidFloat32NoBounds.Ratio is missing minimum bound validation marker" "InvalidFloat32NoBounds.Ratio is missing maximum bound validation marker"
}

// MixedFields has both valid and invalid fields
type MixedFields struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	ValidCount int32

	InvalidCount int32 // want "MixedFields.InvalidCount is missing minimum bound validation marker" "MixedFields.InvalidCount is missing maximum bound validation marker"

	Name string

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1000
	ValidValue int64

	InvalidValue int64 // want "MixedFields.InvalidValue is missing minimum bound validation marker" "MixedFields.InvalidValue is missing maximum bound validation marker"
}

// PointerFields with pointers should also be checked
type PointerFields struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	ValidPointerWithBounds *int32

	InvalidPointer *int32 // want "PointerFields.InvalidPointer is missing minimum bound validation marker" "PointerFields.InvalidPointer is missing maximum bound validation marker"

	InvalidPointer64 *int64 // want "PointerFields.InvalidPointer64 is missing minimum bound validation marker" "PointerFields.InvalidPointer64 is missing maximum bound validation marker"
}

// SliceFields with slices should check the element type
type SliceFields struct {
	ValidSlice []string

	InvalidSlice []int32 // want "SliceFields.InvalidSlice is missing minimum bound validation marker" "SliceFields.InvalidSlice is missing maximum bound validation marker"

	InvalidSlice64 []int64 // want "SliceFields.InvalidSlice64 is missing minimum bound validation marker" "SliceFields.InvalidSlice64 is missing maximum bound validation marker"

	// +kubebuilder:validation:items:Minimum=0
	// +kubebuilder:validation:items:Maximum=100
	ValidSliceWithBounds []int32
}

// TypeAliasFields with type aliases should be checked
type Int32Alias int32
type Int64Alias int64
type Float32Alias float32
type Float64Alias float64

// Type aliases with bounds on the type itself
// +kubebuilder:validation:Minimum=0
// +kubebuilder:validation:Maximum=255
type BoundedInt32Alias int32

// +kubebuilder:validation:Minimum=-1000
// +kubebuilder:validation:Maximum=1000
type BoundedInt64Alias int64

// +kubebuilder:validation:Minimum=0.0
// +kubebuilder:validation:Maximum=1.0
type BoundedFloat32Alias float32

// +kubebuilder:validation:Minimum=-100.5
// +kubebuilder:validation:Maximum=100.5
type BoundedFloat64Alias float64

type TypeAliasFields struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	ValidAlias Int32Alias

	InvalidAlias Int32Alias // want "TypeAliasFields.InvalidAlias is missing minimum bound validation marker" "TypeAliasFields.InvalidAlias is missing maximum bound validation marker"

	InvalidAlias64 Int64Alias // want "TypeAliasFields.InvalidAlias64 is missing minimum bound validation marker" "TypeAliasFields.InvalidAlias64 is missing maximum bound validation marker"

	InvalidAliasFloat32 Float32Alias // want "TypeAliasFields.InvalidAliasFloat32 is missing minimum bound validation marker" "TypeAliasFields.InvalidAliasFloat32 is missing maximum bound validation marker"

	InvalidAliasFloat64 Float64Alias // want "TypeAliasFields.InvalidAliasFloat64 is missing minimum bound validation marker" "TypeAliasFields.InvalidAliasFloat64 is missing maximum bound validation marker"

	// Valid: bounds are on the type alias itself
	ValidBoundedAlias BoundedInt32Alias

	ValidBoundedAlias64 BoundedInt64Alias

	ValidBoundedFloat32 BoundedFloat32Alias

	ValidBoundedFloat64 BoundedFloat64Alias
}

// PointerSliceFields with pointer slices should also be checked
type PointerSliceFields struct {
	InvalidPointerSlice []*int32 // want "PointerSliceFields.InvalidPointerSlice is missing minimum bound validation marker" "PointerSliceFields.InvalidPointerSlice is missing maximum bound validation marker"
}

// K8sDeclarativeValidation with k8s declarative validation markers
type K8sDeclarativeValidation struct {
	// +k8s:minimum=0
	// +k8s:maximum=100
	ValidWithK8sMarkers int32

	// +k8s:minimum=-1000
	// +k8s:maximum=1000
	ValidInt64WithK8s int64
}

// InvalidInt32BoundsOutOfRange has int32 bounds outside the valid int32 range
type InvalidInt32BoundsOutOfRange struct {
	// +kubebuilder:validation:Minimum=-3000000000
	// +kubebuilder:validation:Maximum=3000000000
	OutOfRange int32 // want "InvalidInt32BoundsOutOfRange.OutOfRange has minimum bound -3e\\+09 that is outside the int32 range \\[-2147483648, 2147483647\\]" "InvalidInt32BoundsOutOfRange.OutOfRange has maximum bound 3e\\+09 that is outside the int32 range \\[-2147483648, 2147483647\\]"
}

// MixedMarkersKubebuilderAndK8s can use either kubebuilder or k8s markers
type MixedMarkersKubebuilderAndK8s struct {
	// +kubebuilder:validation:Minimum=0
	// +k8s:maximum=100
	MixedMarkers int32
}

// InvalidMarkerNonNumericMin has a non-numeric minimum value
type InvalidMarkerNonNumericMin struct {
	// +kubebuilder:validation:Minimum=invalid
	// +kubebuilder:validation:Maximum=100
	BadMin int32 // want "InvalidMarkerNonNumericMin.BadMin has an invalid minimum marker"
}

// InvalidMarkerNonNumericMax has a non-numeric maximum value
type InvalidMarkerNonNumericMax struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=notanumber
	BadMax int64 // want "InvalidMarkerNonNumericMax.BadMax has an invalid maximum marker"
}
