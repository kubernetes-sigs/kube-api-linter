package a

type Bools struct {
	ValidString string

	ValidMap map[string]string

	ValidInt32 int32

	ValidInt64 int64

	InvalidBool bool // want "field Bools.InvalidBool should not use a bool. Use a string type with meaningful constant values as an enum."

	InvalidBoolPtr *bool // want "field Bools.InvalidBoolPtr pointer should not use a bool. Use a string type with meaningful constant values as an enum."

	InvalidBoolSlice []bool // want "field Bools.InvalidBoolSlice array element should not use a bool. Use a string type with meaningful constant values as an enum."

	InvalidBoolPtrSlice []*bool // want "field Bools.InvalidBoolPtrSlice array element pointer should not use a bool. Use a string type with meaningful constant values as an enum."

	InvalidBoolAlias BoolAlias // want "field Bools.InvalidBoolAlias type BoolAlias should not use a bool. Use a string type with meaningful constant values as an enum."

	InvalidBoolPtrAlias *BoolAlias // want "field Bools.InvalidBoolPtrAlias pointer type BoolAlias should not use a bool. Use a string type with meaningful constant values as an enum."

	InvalidBoolSliceAlias []BoolAlias // want "field Bools.InvalidBoolSliceAlias array element type BoolAlias should not use a bool. Use a string type with meaningful constant values as an enum."

	InvalidBoolPtrSliceAlias []*BoolAlias // want "field Bools.InvalidBoolPtrSliceAlias array element pointer type BoolAlias should not use a bool. Use a string type with meaningful constant values as an enum."

	InvalidMapStringToBool map[string]bool // want "field Bools.InvalidMapStringToBool map value should not use a bool. Use a string type with meaningful constant values as an enum."

	InvalidMapStringToBoolPtr map[string]*bool // want "field Bools.InvalidMapStringToBoolPtr map value pointer should not use a bool. Use a string type with meaningful constant values as an enum."

	InvalidMapBoolToString map[bool]string // want "field Bools.InvalidMapBoolToString map key should not use a bool. Use a string type with meaningful constant values as an enum."

	InvalidMapBoolPtrToString map[*bool]string // want "field Bools.InvalidMapBoolPtrToString map key pointer should not use a bool. Use a string type with meaningful constant values as an enum."

	InvalidBoolAliasFromAnotherFile BoolAliasB // want "field Bools.InvalidBoolAliasFromAnotherFile type BoolAliasB should not use a bool. Use a string type with meaningful constant values as an enum."

	InvalidBoolPtrAliasFromAnotherFile *BoolAliasB // want "field Bools.InvalidBoolPtrAliasFromAnotherFile pointer type BoolAliasB should not use a bool. Use a string type with meaningful constant values as an enum."
}

// DoNothing is used to check that the analyser doesn't report on methods.
func (Bools) DoNothing(a bool) bool {
	return a
}

type BoolAlias bool // want "type BoolAlias should not use a bool. Use a string type with meaningful constant values as an enum."

type BoolAliasPtr *bool // want "type BoolAliasPtr pointer should not use a bool. Use a string type with meaningful constant values as an enum."

type BoolAliasSlice []bool // want "type BoolAliasSlice array element should not use a bool. Use a string type with meaningful constant values as an enum."

type BoolAliasPtrSlice []*bool // want "type BoolAliasPtrSlice array element pointer should not use a bool. Use a string type with meaningful constant values as an enum."

type MapStringToBoolAlias map[string]bool // want "type MapStringToBoolAlias map value should not use a bool. Use a string type with meaningful constant values as an enum"

type MapStringToBoolPtrAlias map[string]*bool //want "type MapStringToBoolPtrAlias map value pointer should not use a bool. Use a string type with meaningful constant values as an enum"
