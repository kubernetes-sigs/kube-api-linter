package a

import (
	"a/aa"
)

type TestObject struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type SimpleObject struct {
	Name string `json:"name"`
}

// ObjectWithEmbedded is an object with an embedded struct.
// As there's no json tag, the embedded struct is inlined.
type ObjectWithEmbedded struct {
	TestObject
	Embedded string `json:"embedded"`
}

type ObjectWithInlineEmbedded struct {
	TestObject `json:",inline"`
	Embedded   string `json:"embedded"`
}

type ObjectWithNamedEmbedded struct {
	TestObject `json:"testObject"`
	Embedded   string `json:"embedded"`
}

type String string
type Int int
type Object TestObject
type StringArray []string

type StringAlias = string
type IntAlias = int
type ObjectAlias = TestObject
type StringArrayAlias = []string

type ByteArray []byte
type ByteArrayAlias = []byte

// +listType=atomic
type StringArrayAtomic StringArray

type SSATagsTestSpec struct {
	// Valid atomic list - should pass
	// +listType=atomic
	AtomicStringList []string `json:"atomicStringList,omitempty"`

	// Valid atomic object list - should pass
	StringArrayAtomic StringArrayAtomic `json:"stringArrayAtomic,omitempty"`

	// Valid atomic object list from another file - should pass
	OtherFileStringArrayAtomic OtherFileStringArrayAtomic `json:"otherFileStringArrayAtomic,omitempty"`

	// Valid atomic object list from another package - should pass
	OtherPackageStringArrayAtomic aa.OtherPackageStringArrayAtomic `json:"otherPackageStringArrayAtomic,omitempty"`

	// Valid atomic object list - should pass
	// +listType=atomic
	AtomicObjectList []TestObject `json:"atomicObjectList,omitempty"`

	// Valid set for primitive list - should pass (no warning for primitive lists)
	// +listType=set
	SetPrimitiveList []string `json:"setPrimitiveList,omitempty"`

	// Invalid set for object list - should warn about listType=set compatibility issues
	// +listType=set
	SetObjectList []TestObject `json:"setObjectList,omitempty"` // want "SetObjectList with listType=set is not recommended due to Server-Side Apply compatibility issues. Consider using listType=atomic or listType=map instead"

	// Invalid map on primitive list - should error
	// +listType=map
	// +listMapKey=name
	MapPrimitiveList []string `json:"mapPrimitiveList,omitempty"` // want "MapPrimitiveList with listType=map can only be used for object lists, not primitive lists"

	// Valid map on object list with proper listMapKey - should pass
	// +listType=map
	// +listMapKey=name
	MapObjectList []TestObject `json:"mapObjectList,omitempty"`

	// Invalid map on object list without listMapKey - should error
	// +listType=map
	MapObjectListNoKey []TestObject `json:"mapObjectListNoKey,omitempty"` // want "MapObjectListNoKey with listType=map must have at least one listMapKey marker"

	// Invalid map on object list with non-existent listMapKey - should error
	// +listType=map
	// +listMapKey=nonexistent
	MapObjectListInvalidKey []TestObject `json:"mapObjectListInvalidKey,omitempty"` // want "MapObjectListInvalidKey listMapKey \"nonexistent\" does not exist as a field in the struct"

	// Valid map with correct JSON tag name - should pass
	// +listType=map
	// +listMapKey=name
	MapObjectListValid []SimpleObject `json:"mapObjectListValid,omitempty"`

	// Invalid map with non-existent JSON tag name - should error
	// +listType=map
	// +listMapKey=invalid_name
	MapObjectListInvalidName []SimpleObject `json:"mapObjectListInvalidName,omitempty"` // want "MapObjectListInvalidName listMapKey \"invalid_name\" does not exist as a field in the struct"

	// Invalid listType value - should error
	// +listType=invalid
	InvalidListType []string `json:"invalidListType,omitempty"` // want "InvalidListType has invalid listType \"invalid\", must be one of: atomic, set, map"

	// Missing listType on primitive array - should warn
	PrimitiveArrayNoMarker []string `json:"primitiveArrayNoMarker,omitempty"` // want "PrimitiveArrayNoMarker should have a listType marker for proper Server-Side Apply behavior \\(atomic, set, or map\\)"

	// Missing listType on object array - should warn
	ObjectArrayNoMarker []TestObject `json:"objectArrayNoMarker,omitempty"` // want "ObjectArrayNoMarker should have a listType marker for proper Server-Side Apply behavior \\(atomic, set, or map\\)"

	// Non-array field - should be ignored
	SingleValue string `json:"singleValue,omitempty"`

	// Pointer array field - should behave same as non-pointer
	PointerArrayNoMarker []*TestObject `json:"pointerArrayNoMarker,omitempty"` // want "PointerArrayNoMarker should have a listType marker for proper Server-Side Apply behavior \\(atomic, set, or map\\)"

	// Defined type tests - defined types should behave as their underlying types
	// +listType=atomic
	StringAtomicList []String `json:"stringAtomicList,omitempty"`

	// +listType=set
	StringSetList []String `json:"stringSetList,omitempty"`

	// Defined type to object should behave as object list
	// +listType=atomic
	ObjectAtomicList []Object `json:"objectAtomicList,omitempty"`

	// +listType=set
	ObjectSetList []Object `json:"objectSetList,omitempty"` // want "ObjectSetList with listType=set is not recommended due to Server-Side Apply compatibility issues. Consider using listType=atomic or listType=map instead"

	// Missing listType on defined type to basic type - should warn
	StringNoMarker []String `json:"stringNoMarker,omitempty"` // want "StringNoMarker should have a listType marker for proper Server-Side Apply behavior \\(atomic, set, or map\\)"

	// Missing listType on defined type to object - should warn
	ObjectNoMarker []Object `json:"objectNoMarker,omitempty"` // want "ObjectNoMarker should have a listType marker for proper Server-Side Apply behavior \\(atomic, set, or map\\)"

	// Pointer to defined type - should behave same as defined type
	// +listType=atomic
	PointerToString []*String `json:"pointerToString,omitempty"`

	// +listType=set
	PointerToObject []*Object `json:"pointerToObject,omitempty"` // want "PointerToObject with listType=set is not recommended due to Server-Side Apply compatibility issues. Consider using listType=atomic or listType=map instead"

	// Type alias tests - aliases to basic types should behave as primitive lists
	// +listType=atomic
	StringAliasAtomicList []StringAlias `json:"stringAliasAtomicList,omitempty"`

	// +listType=set
	StringAliasSetList []StringAlias `json:"stringAliasSetList,omitempty"`

	// Type alias to object should behave as object list
	// +listType=atomic
	ObjectAliasAtomicList []ObjectAlias `json:"objectAliasAtomicList,omitempty"`

	// +listType=set
	ObjectAliasSetList []ObjectAlias `json:"objectAliasSetList,omitempty"` // want "ObjectAliasSetList with listType=set is not recommended due to Server-Side Apply compatibility issues. Consider using listType=atomic or listType=map instead"

	// Missing listType on alias to basic type - should warn
	StringAliasNoMarker []StringAlias `json:"stringAliasNoMarker,omitempty"` // want "StringAliasNoMarker should have a listType marker for proper Server-Side Apply behavior \\(atomic, set, or map\\)"

	// Missing listType on alias to object - should warn
	ObjectAliasNoMarker []ObjectAlias `json:"objectAliasNoMarker,omitempty"` // want "ObjectAliasNoMarker should have a listType marker for proper Server-Side Apply behavior \\(atomic, set, or map\\)"

	// Pointer to alias - should behave same as alias
	// +listType=atomic
	PointerToStringAlias []*StringAlias `json:"pointerToStringAlias,omitempty"`

	// +listType=set
	PointerToObjectAlias []*ObjectAlias `json:"pointerToObjectAlias,omitempty"` // want "PointerToObjectAlias with listType=set is not recommended due to Server-Side Apply compatibility issues. Consider using listType=atomic or listType=map instead"

	// Multiple pointer levels
	PointerToPointerObject []*(*TestObject) `json:"pointerToPointerObject,omitempty"` // want "PointerToPointerObject should have a listType marker for proper Server-Side Apply behavior \\(atomic, set, or map\\)"

	// Array defined type tests - defined types to array types should behave as array lists
	// +listType=atomic
	StringArrayAtomicList StringArray `json:"stringArrayAtomicList,omitempty"`

	// +listType=set
	StringArraySetList StringArray `json:"stringArraySetList,omitempty"`

	// Missing listType on array defined type - should warn
	StringArrayNoMarker StringArray `json:"stringArrayNoMarker,omitempty"` // want "StringArrayNoMarker should have a listType marker for proper Server-Side Apply behavior \\(atomic, set, or map\\)"

	// Array type alias tests - aliases to array types should behave as array lists
	// +listType=atomic
	StringArrayAliasAtomicList StringArrayAlias `json:"stringArrayAliasAtomicList,omitempty"`

	// +listType=set
	StringArrayAliasSetList StringArrayAlias `json:"stringArrayAliasSetList,omitempty"`

	// Missing listType on array alias - should warn
	StringArrayAliasNoMarker StringArrayAlias `json:"stringArrayAliasNoMarker,omitempty"` // want "StringArrayAliasNoMarker should have a listType marker for proper Server-Side Apply behavior \\(atomic, set, or map\\)"

	// Byte array tests - these should be skipped by IsByteArray condition
	// Direct byte array - should be ignored (no listType required)
	ByteArrayDirect []byte `json:"byteArrayDirect,omitempty"`

	// Defined type byte array - should be ignored (no listType required)
	ByteArrayDefinedType ByteArray `json:"byteArrayDefinedType,omitempty"`

	// Aliased byte array - should be ignored (no listType required)
	ByteArrayAliased ByteArrayAlias `json:"byteArrayAliased,omitempty"`

	// Byte array with markers - should be ignored even if markers are present
	// +listType=atomic
	ByteArrayWithMarker []byte `json:"byteArrayWithMarker,omitempty"` // want "ByteArrayWithMarker is a byte array, which does not support the listType marker. Remove the listType marker"

	// +optional
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	AllOf []JSONSchemaProps `json:"allOf,omitempty"`

	// Object with embedded struct - list map key from embedded struct should be allowed
	// Without a json tag, the default is to inline the embedded struct.
	// +listType=map
	// +listMapKey=name
	ObjectWithEmbeddedStruct []ObjectWithEmbedded `json:"objectWithEmbeddedStruct,omitempty"`

	// Object with inline embedded struct - list map key from embedded struct should be allowed
	// With a json tag, the embedded struct is not inlined.
	// +listType=map
	// +listMapKey=name
	ObjectWithInlineEmbeddedStruct []ObjectWithInlineEmbedded `json:"objectWithInlineEmbeddedStruct,omitempty"`

	// Object with named embedded struct - list map key from embedded struct should be allowed
	// With a json tag, the embedded struct is not inlined.
	// +listType=map
	// +listMapKey=name
	ObjectWithNamedEmbeddedStruct []ObjectWithNamedEmbedded `json:"objectWithNamedEmbeddedStruct,omitempty"` // want "ObjectWithNamedEmbeddedStruct listMapKey \"name\" does not exist as a field in the struct"
}

// JSONSchemaProps is a placeholder for the JSON schema properties.
type JSONSchemaProps struct{}
