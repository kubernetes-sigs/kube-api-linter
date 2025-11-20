package a

type Integers struct {
	String string // want "field Integers.String is a string"

	Map map[string]string // want "field Integers.Map map key is a string" "field Integers.Map map value is a string"

	MapStringToStringAlias map[string]StringAlias // want "field Integers.MapStringToStringAlias map key is a string" "field Integers.MapStringToStringAlias map value type StringAlias is a string"

	Int32 int32

	Int64 int64

	Bool bool

	StringPtr *string // want "field Integers.StringPtr pointer is a string"

	StringSlice []string // want "field Integers.StringSlice array element is a string"

	StringPtrSlice []*string // want "field Integers.StringPtrSlice array element pointer is a string"

	StringAlias StringAlias // want "field Integers.StringAlias type StringAlias is a string"

	StringAliasPtr *StringAlias // want "field Integers.StringAliasPtr pointer type StringAlias is a string"

	StringAliasSlice []StringAlias // want "field Integers.StringAliasSlice array element type StringAlias is a string"

	StringAliasPtrSlice []*StringAlias // want "field Integers.StringAliasPtrSlice array element pointer type StringAlias is a string"

	StringAliasFromAnotherFile StringAliasB // want "field Integers.StringAliasFromAnotherFile type StringAliasB is a string"
}

type StringAlias string // want "type StringAlias is a string"

type StringAliasPtr *string // want "type StringAliasPtr pointer is a string"

type StringAliasSlice []string // want "type StringAliasSlice array element is a string"

type StringAliasPtrSlice []*string // want "type StringAliasPtrSlice array element pointer is a string"
