package a

type Integers struct {
	String string // want "field String is a string"

	Map map[string]string

	Int32 int32

	Int64 int64

	Bool bool

	StringPtr *string // want "field StringPtr pointer is a string"

	StringSlice []string // want "field StringSlice array element is a string"

	StringPtrSlice []*string // want "field StringPtrSlice array element pointer is a string"

	StringAlias StringAlias // want "field StringAlias type StringAlias is a string"

	StringAliasPtr *StringAlias // want "field StringAliasPtr pointer type StringAlias is a string"

	StringAliasSlice []StringAlias // want "field StringAliasSlice array element type StringAlias is a string"

	StringAliasPtrSlice []*StringAlias // want "field StringAliasPtrSlice array element pointer type StringAlias is a string"
}

type StringAlias string // want "type StringAlias is a string"

type StringAliasPtr *string // want "type StringAliasPtr pointer is a string"

type StringAliasSlice []string // want "type StringAliasSlice array element is a string"

type StringAliasPtrSlice []*string // want "type StringAliasPtrSlice array element pointer is a string"
