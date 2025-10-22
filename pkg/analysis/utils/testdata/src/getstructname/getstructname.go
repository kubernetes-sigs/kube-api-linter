package getstructname

type MyStruct1 struct {
	Field1 string // want "field Field1 is in struct MyStruct1"
	Field2 int    // want "field Field2 is in struct MyStruct1"
}

type MyStruct2 struct {
	AnotherField bool // want "field AnotherField is in struct MyStruct2"
}
