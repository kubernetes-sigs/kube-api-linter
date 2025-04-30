package a

type BoolAliasB bool // want "type BoolAliasB should not use a bool. Use a string type with meaningful constant values as an enum."

type BoolAliasPtrB *bool // want "type BoolAliasPtrB pointer should not use a bool. Use a string type with meaningful constant values as an enum."
