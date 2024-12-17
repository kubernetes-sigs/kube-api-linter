package a

type Integers struct {
	ValidString string

	ValidMap map[string]string

	ValidInt32 int32

	ValidInt32Ptr *int32

	ValidInt64 int64

	ValidInt64Ptr *int64

	InvalidInt int // want "field InvalidInt should not use an int, int8 or int16. Use int32 or int64 depending on bounding requirements"

	InvalidIntPtr *int // want "field InvalidIntPtr pointer should not use an int, int8 or int16. Use int32 or int64 depending on bounding requirements"

	InvalidInt8 int8 // want "field InvalidInt8 should not use an int, int8 or int16. Use int32 or int64 depending on bounding requirements"

	InvalidInt16 int16 // want "field InvalidInt16 should not use an int, int8 or int16. Use int32 or int64 depending on bounding requirements"

	InvalidUInt uint // want "field InvalidUInt should not use unsigned integers, use only int32 or int64 and apply validation to ensure the value is positive"

	InvalidUIntPtr uint // want "field InvalidUIntPtr should not use unsigned integers, use only int32 or int64 and apply validation to ensure the value is positive"

	InvalidUInt8 uint8 // want "field InvalidUInt8 should not use unsigned integers, use only int32 or int64 and apply validation to ensure the value is positive"

	InvalidUInt16 uint16 // want "field InvalidUInt16 should not use unsigned integers, use only int32 or int64 and apply validation to ensure the value is positive"

	InvalidUInt32 uint32 // want "field InvalidUInt32 should not use unsigned integers, use only int32 or int64 and apply validation to ensure the value is positive"

	InvalidUInt64 uint64 // want "field InvalidUInt64 should not use unsigned integers, use only int32 or int64 and apply validation to ensure the value is positive"

	ValidInt32Alias ValidInt32Alias

	ValidInt32AliasPtr *ValidInt32Alias

	InvalidIntAlias InvalidIntAlias // want "field InvalidIntAlias type InvalidIntAlias should not use an int, int8 or int16. Use int32 or int64 depending on bounding requirements"

	InvalidIntAliasPtr *InvalidIntAlias // want "field InvalidIntAliasPtr pointer type InvalidIntAlias should not use an int, int8 or int16. Use int32 or int64 depending on bounding requirements"

	InvalidUIntAlias InvalidUIntAlias // want "field InvalidUIntAlias type InvalidUIntAlias should not use unsigned integers, use only int32 or int64 and apply validation to ensure the value is positive"

	InvalidUIntAliasPtr *InvalidUIntAlias // want "field InvalidUIntAliasPtr pointer type InvalidUIntAlias should not use unsigned integers, use only int32 or int64 and apply validation to ensure the value is positive"

	InvalidIntAliasAlias InvalidIntAliasAlias // want "field InvalidIntAliasAlias type InvalidIntAliasAlias type InvalidIntAlias should not use an int, int8 or int16. Use int32 or int64 depending on bounding requirements"

	ValidSliceInt32 []int32

	ValidSliceInt32Ptr []*int32

	ValidSliceInt64 []int64

	ValidSliceInt64Ptr []*int64

	InvalidSliceInt []int // want "field InvalidSliceInt array element should not use an int, int8 or int16. Use int32 or int64 depending on bounding requirements"

	InvalidSliceIntPtr []*int // want "field InvalidSliceIntPtr array element pointer should not use an int, int8 or int16. Use int32 or int64 depending on bounding requirements"

	InvalidSliceUInt []uint // want "field InvalidSliceUInt array element should not use unsigned integers, use only int32 or int64 and apply validation to ensure the value is positive"

	InvalidSliceUIntPtr []*uint // want "field InvalidSliceUIntPtr array element pointer should not use unsigned integers, use only int32 or int64 and apply validation to ensure the value is positive"

	InvalidSliceIntAlias []InvalidIntAlias // want "field InvalidSliceIntAlias array element type InvalidIntAlias should not use an int, int8 or int16. Use int32 or int64 depending on bounding requirements"

	InvalidSliceIntAliasPtr []*InvalidIntAlias // want "field InvalidSliceIntAliasPtr array element pointer type InvalidIntAlias should not use an int, int8 or int16. Use int32 or int64 depending on bounding requirements"

	InvalidSliceUIntAlias []InvalidUIntAlias // want "field InvalidSliceUIntAlias array element type InvalidUIntAlias should not use unsigned integers, use only int32 or int64 and apply validation to ensure the value is positive"

	InvalidSliceUIntAliasPtr []*InvalidUIntAlias // want "field InvalidSliceUIntAliasPtr array element pointer type InvalidUIntAlias should not use unsigned integers, use only int32 or int64 and apply validation to ensure the value is positive"

	ValidMapStringToInt32 map[string]int32

	ValidMapStringToInt64 map[string]int64

	InvalidMapStringToInt map[string]int // want "field InvalidMapStringToInt map value should not use an int, int8 or int16. Use int32 or int64 depending on bounding requirements"

	InvalidMapStringToUInt map[string]uint // want "field InvalidMapStringToUInt map value should not use unsigned integers, use only int32 or int64 and apply validation to ensure the value is positive"

	ValidMapInt32ToString map[int32]string

	ValidMapInt64ToString map[int64]string

	InvalidMapIntToString map[int]string // want "field InvalidMapIntToString map key should not use an int, int8 or int16. Use int32 or int64 depending on bounding requirements"

	InvalidMapUIntToString map[uint]string // want "field InvalidMapUIntToString map key should not use unsigned integers, use only int32 or int64 and apply validation to ensure the value is positive"
}

type ValidInt32Alias int32

type ValidInt32PtrAlias *int32

type ValidInt64Alias int64

type ValidInt64PtrAlias *int64

type InvalidIntAlias int // want "type InvalidIntAlias should not use an int, int8 or int16. Use int32 or int64 depending on bounding requirements"

type InvalidIntPtrAlias *int // want "type InvalidIntPtrAlias pointer should not use an int, int8 or int16. Use int32 or int64 depending on bounding requirements"

type InvalidInt8Alias int8 // want "type InvalidInt8Alias should not use an int, int8 or int16. Use int32 or int64 depending on bounding requirements"

type InvalidInt16Alias int16 // want "type InvalidInt16Alias should not use an int, int8 or int16. Use int32 or int64 depending on bounding requirements"

type InvalidUIntAlias uint // want "type InvalidUIntAlias should not use unsigned integers, use only int32 or int64 and apply validation to ensure the value is positive"

type InvalidUIntPtrAlias *uint // want "type InvalidUIntPtrAlias pointer should not use unsigned integers, use only int32 or int64 and apply validation to ensure the value is positive"

type InvalidUInt8Alias uint8 // want "type InvalidUInt8Alias should not use unsigned integers, use only int32 or int64 and apply validation to ensure the value is positive"

type InvalidUInt16Alias uint16 // want "type InvalidUInt16Alias should not use unsigned integers, use only int32 or int64 and apply validation to ensure the value is positive"

type InvalidUInt32Alias uint32 // want "type InvalidUInt32Alias should not use unsigned integers, use only int32 or int64 and apply validation to ensure the value is positive"

type InvalidUInt64Alias uint64 // want "type InvalidUInt64Alias should not use unsigned integers, use only int32 or int64 and apply validation to ensure the value is positive"

type InvalidIntAliasAlias InvalidIntAlias // want "type InvalidIntAliasAlias type InvalidIntAlias should not use an int, int8 or int16. Use int32 or int64 depending on bounding requirements"

type InvalidSliceIntAlias []int // want "type InvalidSliceIntAlias array element should not use an int, int8 or int16. Use int32 or int64 depending on bounding requirements"

type InvalidSliceIntPtrAlias []*int // want "type InvalidSliceIntPtrAlias array element pointer should not use an int, int8 or int16. Use int32 or int64 depending on bounding requirements"

type InvalidSliceIntAliasAlias []InvalidIntAlias // want "type InvalidSliceIntAliasAlias array element type InvalidIntAlias should not use an int, int8 or int16. Use int32 or int64 depending on bounding requirements"

type InvalidSliceIntAliasPtrAlias []*InvalidIntAlias // want "type InvalidSliceIntAliasPtrAlias array element pointer type InvalidIntAlias should not use an int, int8 or int16. Use int32 or int64 depending on bounding requirements"
