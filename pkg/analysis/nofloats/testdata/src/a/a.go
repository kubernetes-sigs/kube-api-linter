package a

type Floats struct {
	ValidString string

	ValidMap map[string]string

	ValidInt32 int32

	ValidInt64 int64

	InvalidFloat32 float32 // want "field InvalidFloat32 should not use a float value because they cannot be reliably round-tripped."

	InvalidFloat64 float64 // want "field InvalidFloat64 should not use a float value because they cannot be reliably round-tripped."

	InvalidFloat32Ptr *float32 // want "field InvalidFloat32Ptr pointer should not use a float value because they cannot be reliably round-tripped."

	InvalidFloat64Ptr *float64 // want "field InvalidFloat64Ptr pointer should not use a float value because they cannot be reliably round-tripped."

	InvalidFloat32Slice []float32 // want "field InvalidFloat32Slice array element should not use a float value because they cannot be reliably round-tripped."

	InvalidFloat64Slice []float64 // want "field InvalidFloat64Slice array element should not use a float value because they cannot be reliably round-tripped."

	InvalidFloat32PtrSlice []*float32 // want "field InvalidFloat32PtrSlice array element pointer should not use a float value because they cannot be reliably round-tripped."

	InvalidFloat64PtrSlice []*float64 // want "field InvalidFloat64PtrSlice array element pointer should not use a float value because they cannot be reliably round-tripped."

	InvalidFloat32Alias Float32Alias // want "field InvalidFloat32Alias type Float32Alias should not use a float value because they cannot be reliably round-tripped."

	InvalidFloat64Alias Float64Alias // want "field InvalidFloat64Alias type Float64Alias should not use a float value because they cannot be reliably round-tripped."

	InvalidFloat32PtrAlias *Float32Alias // want "field InvalidFloat32PtrAlias pointer type Float32Alias should not use a float value because they cannot be reliably round-tripped."

	InvalidFloat64PtrAlias *Float64Alias // want "field InvalidFloat64PtrAlias pointer type Float64Alias should not use a float value because they cannot be reliably round-tripped."

	InvalidFloat32SliceAlias []Float32Alias // want "field InvalidFloat32SliceAlias array element type Float32Alias should not use a float value because they cannot be reliably round-tripped."

	InvalidFloat64SliceAlias []Float64Alias // want "field InvalidFloat64SliceAlias array element type Float64Alias should not use a float value because they cannot be reliably round-tripped."

	InvalidFloat32PtrSliceAlias []*Float32Alias // want "field InvalidFloat32PtrSliceAlias array element pointer type Float32Alias should not use a float value because they cannot be reliably round-tripped."

	InvalidFloat64PtrSliceAlias []*Float64Alias // want "field InvalidFloat64PtrSliceAlias array element pointer type Float64Alias should not use a float value because they cannot be reliably round-tripped."

	InvalidMapStringToFloat32 map[string]float32 // want "field InvalidMapStringToFloat32 map value should not use a float value because they cannot be reliably round-tripped."

	InvalidMapStringToFloat64 map[string]float64 // want "field InvalidMapStringToFloat64 map value should not use a float value because they cannot be reliably round-tripped."

	InvalidMapStringToFloat32Ptr map[string]*float32 // want "field InvalidMapStringToFloat32Ptr map value pointer should not use a float value because they cannot be reliably round-tripped."

	InvalidMapStringToFloat64Ptr map[string]*float64 // want "field InvalidMapStringToFloat64Ptr map value pointer should not use a float value because they cannot be reliably round-tripped."

	InvalidMapFloat32ToString map[float32]string // want "field InvalidMapFloat32ToString map key should not use a float value because they cannot be reliably round-tripped."

	InvalidMapFloat64ToString map[float64]string // want "field InvalidMapFloat64ToString map key should not use a float value because they cannot be reliably round-tripped."

	InvalidMapFloat32PtrToString map[*float32]string // want "field InvalidMapFloat32PtrToString map key pointer should not use a float value because they cannot be reliably round-tripped."

	InvalidMapFloat64PtrToString map[*float64]string // want "field InvalidMapFloat64PtrToString map key pointer should not use a float value because they cannot be reliably round-tripped."
}

// DoNothingFloat32 is used to check that the analyser doesn't report on methods.
func (Floats) DoNothingFloat32(a float32) float32 {
	return a
}

// DoNothingFloat64 is used to check that the analyser doesn't report on methods.
func (Floats) DoNothingFloat64(a float64) float64 {
	return a
}

type Float32Alias float32 // want "type Float32Alias should not use a float value because they cannot be reliably round-tripped."

type Float64Alias float64 // want "type Float64Alias should not use a float value because they cannot be reliably round-tripped."

type Float32AliasPtr *float32 // want "type Float32AliasPtr pointer should not use a float value because they cannot be reliably round-tripped."

type Float64AliasPtr *float64 // want "type Float64AliasPtr pointer should not use a float value because they cannot be reliably round-tripped."

type Float32AliasSlice []float32 // want "type Float32AliasSlice array element should not use a float value because they cannot be reliably round-tripped."

type Float64AliasSlice []float64 // want "type Float64AliasSlice array element should not use a float value because they cannot be reliably round-tripped."

type Float32AliasPtrSlice []*float32 // want "type Float32AliasPtrSlice array element pointer should not use a float value because they cannot be reliably round-tripped."

type Float64AliasPtrSlice []*float64 // want "type Float64AliasPtrSlice array element pointer should not use a float value because they cannot be reliably round-tripped."

type MapStringToFloat32Alias map[string]float32 // want "type MapStringToFloat32Alias map value should not use a float value because they cannot be reliably round-tripped."

type MapStringToFloat64Alias map[string]float64 // want "type MapStringToFloat64Alias map value should not use a float value because they cannot be reliably round-tripped."

type MapStringToFloat32PtrAlias map[string]*float32 // want "type MapStringToFloat32PtrAlias map value pointer should not use a float value because they cannot be reliably round-tripped."

type MapStringToFloat64PtrAlias map[string]*float64 // want "type MapStringToFloat64PtrAlias map value pointer should not use a float value because they cannot be reliably round-tripped."
