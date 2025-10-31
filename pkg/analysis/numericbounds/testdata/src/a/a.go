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
	NoBounds int32 // want "field NoBounds of type int32 should have minimum and maximum bounds validation markers"
}

// InvalidInt64NoBounds should have bounds markers
type InvalidInt64NoBounds struct {
	NoBounds int64 // want "field NoBounds of type int64 should have minimum and maximum bounds validation markers"
}

// InvalidInt32OnlyMin should have maximum marker
type InvalidInt32OnlyMin struct {
	// +kubebuilder:validation:Minimum=0
	OnlyMin int32 // want "field OnlyMin of type int32 has minimum but is missing maximum bounds validation marker"
}

// InvalidInt32OnlyMax should have minimum marker
type InvalidInt32OnlyMax struct {
	// +kubebuilder:validation:Maximum=100
	OnlyMax int32 // want "field OnlyMax of type int32 has maximum but is missing minimum bounds validation marker"
}

// InvalidInt64OnlyMin should have maximum marker
type InvalidInt64OnlyMin struct {
	// +kubebuilder:validation:Minimum=0
	OnlyMin int64 // want "field OnlyMin of type int64 has minimum but is missing maximum bounds validation marker"
}

// InvalidInt64OnlyMax should have minimum marker
type InvalidInt64OnlyMax struct {
	// +kubebuilder:validation:Maximum=100
	OnlyMax int64 // want "field OnlyMax of type int64 has maximum but is missing minimum bounds validation marker"
}

// InvalidInt64ExceedsJSMaxBounds has maximum that exceeds JavaScript safe integer range
type InvalidInt64ExceedsJSMaxBounds struct {
	// +kubebuilder:validation:Minimum=-1000
	// +kubebuilder:validation:Maximum=9007199254740992
	UnsafeMax int64 // want "field UnsafeMax of type int64 has bounds \\[-1000, 9007199254740992\\] that exceed safe integer range \\[-9007199254740991, 9007199254740991\\]\\. Consider using a string type to avoid precision loss in JavaScript clients"
}

// InvalidInt64ExceedsJSMinBounds has minimum that exceeds JavaScript safe integer range
type InvalidInt64ExceedsJSMinBounds struct {
	// +kubebuilder:validation:Minimum=-9007199254740992
	// +kubebuilder:validation:Maximum=1000
	UnsafeMin int64 // want "field UnsafeMin of type int64 has bounds \\[-9007199254740992, 1000\\] that exceed safe integer range \\[-9007199254740991, 9007199254740991\\]\\. Consider using a string type to avoid precision loss in JavaScript clients"
}

// InvalidInt64ExceedsJSBothBounds has both bounds exceeding JavaScript safe integer range
type InvalidInt64ExceedsJSBothBounds struct {
	// +kubebuilder:validation:Minimum=-9007199254740992
	// +kubebuilder:validation:Maximum=9007199254740992
	UnsafeBoth int64 // want "field UnsafeBoth of type int64 has bounds \\[-9007199254740992, 9007199254740992\\] that exceed safe integer range \\[-9007199254740991, 9007199254740991\\]\\. Consider using a string type to avoid precision loss in JavaScript clients"
}

// IgnoredStringField should not be checked
type IgnoredStringField struct {
	Name string
}

// IgnoredBoolField should not be checked
type IgnoredBoolField struct {
	Enabled bool
}

// IgnoredFloat64Field should not be checked
type IgnoredFloat64Field struct {
	Value float64
}

// MixedFields has both valid and invalid fields
type MixedFields struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	ValidCount int32

	InvalidCount int32 // want "field InvalidCount of type int32 should have minimum and maximum bounds validation markers"

	Name string

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1000
	ValidValue int64

	InvalidValue int64 // want "field InvalidValue of type int64 should have minimum and maximum bounds validation markers"
}

// PointerFields with pointers should also be checked
type PointerFields struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	ValidPointerWithBounds *int32

	InvalidPointer *int32 // want "field InvalidPointer of type int32 should have minimum and maximum bounds validation markers"

	InvalidPointer64 *int64 // want "field InvalidPointer64 of type int64 should have minimum and maximum bounds validation markers"
}

// SliceFields with slices should check the element type
type SliceFields struct {
	ValidSlice []string

	InvalidSlice []int32 // want "field InvalidSlice of type int32 should have minimum and maximum bounds validation markers"

	InvalidSlice64 []int64 // want "field InvalidSlice64 of type int64 should have minimum and maximum bounds validation markers"

	// +kubebuilder:validation:items:Minimum=0
	// +kubebuilder:validation:items:Maximum=100
	ValidSliceWithBounds []int32
}

// TypeAliasFields with type aliases should be checked
type Int32Alias int32
type Int64Alias int64

type TypeAliasFields struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	ValidAlias Int32Alias

	InvalidAlias Int32Alias // want "field InvalidAlias of type int32 should have minimum and maximum bounds validation markers"

	InvalidAlias64 Int64Alias // want "field InvalidAlias64 of type int64 should have minimum and maximum bounds validation markers"
}

// PointerSliceFields with pointer slices should also be checked
type PointerSliceFields struct {
	InvalidPointerSlice []*int32 // want "field InvalidPointerSlice of type int32 should have minimum and maximum bounds validation markers"
}
