package a

// ValidInt32WithBounds has proper bounds validation
type ValidInt32WithBounds struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	Count int32
}

// ValidInt64WithBounds has proper bounds validation
type ValidInt64WithBounds struct {
	// +kubebuilder:validation:Minimum=-1000
	// +kubebuilder:validation:Maximum=1000
	Value int64
}

// ValidInt64WithJSSafeBounds has bounds within JavaScript safe integer range
type ValidInt64WithJSSafeBounds struct {
	// +kubebuilder:validation:Minimum=-9007199254740991
	// +kubebuilder:validation:Maximum=9007199254740991
	SafeValue int64
}

// InvalidInt32NoBounds should have bounds markers
type InvalidInt32NoBounds struct {
	NoBounds int32 // want "field NoBounds is missing minimum bounds validation marker" "field NoBounds is missing maximum bounds validation marker"
}

// InvalidInt64NoBounds should have bounds markers
type InvalidInt64NoBounds struct {
	NoBounds int64 // want "field NoBounds is missing minimum bounds validation marker" "field NoBounds is missing maximum bounds validation marker"
}

// InvalidInt32OnlyMin should have maximum marker
type InvalidInt32OnlyMin struct {
	// +kubebuilder:validation:Minimum=0
	OnlyMin int32 // want "field OnlyMin is missing maximum bounds validation marker"
}

// InvalidInt32OnlyMax should have minimum marker
type InvalidInt32OnlyMax struct {
	// +kubebuilder:validation:Maximum=100
	OnlyMax int32 // want "field OnlyMax is missing minimum bounds validation marker"
}

// InvalidInt64OnlyMin should have maximum marker
type InvalidInt64OnlyMin struct {
	// +kubebuilder:validation:Minimum=0
	OnlyMin int64 // want "field OnlyMin is missing maximum bounds validation marker"
}

// InvalidInt64OnlyMax should have minimum marker
type InvalidInt64OnlyMax struct {
	// +kubebuilder:validation:Maximum=100
	OnlyMax int64 // want "field OnlyMax is missing minimum bounds validation marker"
}

// InvalidInt64ExceedsJSMaxBounds has maximum that exceeds JavaScript safe integer range
type InvalidInt64ExceedsJSMaxBounds struct {
	// +kubebuilder:validation:Minimum=-1000
	// +kubebuilder:validation:Maximum=9007199254740992
	UnsafeMax int64 // want "field UnsafeMax has maximum bound 9\\.007199254740992e\\+15 that is outside the JavaScript-safe int64 range \\[-9007199254740991, 9007199254740991\\]\\. Consider using a string type to avoid precision loss in JavaScript clients"
}

// InvalidInt64ExceedsJSMinBounds has minimum that exceeds JavaScript safe integer range
type InvalidInt64ExceedsJSMinBounds struct {
	// +kubebuilder:validation:Minimum=-9007199254740992
	// +kubebuilder:validation:Maximum=1000
	UnsafeMin int64 // want "field UnsafeMin has minimum bound -9\\.007199254740992e\\+15 that is outside the JavaScript-safe int64 range \\[-9007199254740991, 9007199254740991\\]\\. Consider using a string type to avoid precision loss in JavaScript clients"
}

// InvalidInt64ExceedsJSBothBounds has both bounds exceeding JavaScript safe integer range
type InvalidInt64ExceedsJSBothBounds struct {
	// +kubebuilder:validation:Minimum=-9007199254740992
	// +kubebuilder:validation:Maximum=9007199254740992
	UnsafeBoth int64 // want "field UnsafeBoth has minimum bound -9\\.007199254740992e\\+15 that is outside the JavaScript-safe int64 range \\[-9007199254740991, 9007199254740991\\]\\. Consider using a string type to avoid precision loss in JavaScript clients" "field UnsafeBoth has maximum bound 9\\.007199254740992e\\+15 that is outside the JavaScript-safe int64 range \\[-9007199254740991, 9007199254740991\\]\\. Consider using a string type to avoid precision loss in JavaScript clients"
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
	Value float64 // want "field Value is missing minimum bounds validation marker" "field Value is missing maximum bounds validation marker"
}

// InvalidFloat32NoBounds should have bounds markers
type InvalidFloat32NoBounds struct {
	Ratio float32 // want "field Ratio is missing minimum bounds validation marker" "field Ratio is missing maximum bounds validation marker"
}

// MixedFields has both valid and invalid fields
type MixedFields struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	ValidCount int32

	InvalidCount int32 // want "field InvalidCount is missing minimum bounds validation marker" "field InvalidCount is missing maximum bounds validation marker"

	Name string

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1000
	ValidValue int64

	InvalidValue int64 // want "field InvalidValue is missing minimum bounds validation marker" "field InvalidValue is missing maximum bounds validation marker"
}

// PointerFields with pointers should also be checked
type PointerFields struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	ValidPointerWithBounds *int32

	InvalidPointer *int32 // want "field InvalidPointer pointer is missing minimum bounds validation marker" "field InvalidPointer pointer is missing maximum bounds validation marker"

	InvalidPointer64 *int64 // want "field InvalidPointer64 pointer is missing minimum bounds validation marker" "field InvalidPointer64 pointer is missing maximum bounds validation marker"
}

// SliceFields with slices should check the element type
type SliceFields struct {
	ValidSlice []string

	InvalidSlice []int32 // want "field InvalidSlice array element is missing minimum bounds validation marker" "field InvalidSlice array element is missing maximum bounds validation marker"

	InvalidSlice64 []int64 // want "field InvalidSlice64 array element is missing minimum bounds validation marker" "field InvalidSlice64 array element is missing maximum bounds validation marker"

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

	InvalidAlias Int32Alias // want "field InvalidAlias type Int32Alias is missing minimum bounds validation marker" "field InvalidAlias type Int32Alias is missing maximum bounds validation marker"

	InvalidAlias64 Int64Alias // want "field InvalidAlias64 type Int64Alias is missing minimum bounds validation marker" "field InvalidAlias64 type Int64Alias is missing maximum bounds validation marker"

	InvalidAliasFloat32 Float32Alias // want "field InvalidAliasFloat32 type Float32Alias is missing minimum bounds validation marker" "field InvalidAliasFloat32 type Float32Alias is missing maximum bounds validation marker"

	InvalidAliasFloat64 Float64Alias // want "field InvalidAliasFloat64 type Float64Alias is missing minimum bounds validation marker" "field InvalidAliasFloat64 type Float64Alias is missing maximum bounds validation marker"

	// Valid: bounds are on the type alias itself
	ValidBoundedAlias BoundedInt32Alias

	ValidBoundedAlias64 BoundedInt64Alias

	ValidBoundedFloat32 BoundedFloat32Alias

	ValidBoundedFloat64 BoundedFloat64Alias
}

// PointerSliceFields with pointer slices should also be checked
type PointerSliceFields struct {
	InvalidPointerSlice []*int32 // want "field InvalidPointerSlice array element pointer is missing minimum bounds validation marker" "field InvalidPointerSlice array element pointer is missing maximum bounds validation marker"
}

// K8sDeclarativeValidation with k8s declarative validation markers should work
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
	OutOfRange int32 // want "field OutOfRange has minimum bound -3e\\+09 that is outside the valid int32 range" "field OutOfRange has maximum bound 3e\\+09 that is outside the valid int32 range"
}

// MixedMarkersKubebuilderAndK8s can use either kubebuilder or k8s markers
type MixedMarkersKubebuilderAndK8s struct {
	// +kubebuilder:validation:Minimum=0
	// +k8s:maximum=100
	MixedMarkers int32
}
