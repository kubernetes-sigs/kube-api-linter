package b

type TestObject struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type SimpleObject struct {
	Name string `json:"name"`
}

type String string
type Int int
type Object TestObject
type StringArray []string

type StringAlias = string
type IntAlias = int
type ObjectAlias = TestObject
type StringArrayAlias = []string

type SSATagsTestSpec struct {
	// Valid atomic list - should pass
	// +kubebuilder:listType=atomic
	AtomicStringList []string `json:"atomicStringList,omitempty"`

	// Valid atomic object list - should pass
	// +kubebuilder:listType=atomic
	AtomicObjectList []TestObject `json:"atomicObjectList,omitempty"`

	// Valid set for primitive list - should pass (no warning for primitive lists)
	// +kubebuilder:listType=set
	SetPrimitiveList []string `json:"setPrimitiveList,omitempty"`

	// Set for object list - no warning when ListTypeSetUsage=Ignore
	// +kubebuilder:listType=set
	SetObjectList []TestObject `json:"setObjectList,omitempty"`

	// Invalid map on primitive list - should error even when ListTypeSetUsage=Ignore
	// +kubebuilder:listType=map
	// +kubebuilder:listMapKey=name
	MapPrimitiveList []string `json:"mapPrimitiveList,omitempty"` // want "MapPrimitiveList with listType=map can only be used for object lists, not primitive lists"

	// Valid map on object list with proper listMapKey - should pass
	// +kubebuilder:listType=map
	// +kubebuilder:listMapKey=name
	MapObjectList []TestObject `json:"mapObjectList,omitempty"`

	// Map on object list without listMapKey - should error even when ListTypeSetUsage=Ignore
	// +kubebuilder:listType=map
	MapObjectListNoKey []TestObject `json:"mapObjectListNoKey,omitempty"` // want "MapObjectListNoKey with listType=map must have at least one listMapKey marker"

	// Invalid listType value - should error regardless of config
	// +kubebuilder:listType=invalid
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
	// +kubebuilder:listType=atomic
	StringAtomicList []String `json:"stringAtomicList,omitempty"`

	// +kubebuilder:listType=set
	StringSetList []String `json:"stringSetList,omitempty"`

	// Defined type to object should behave as object list - no set warning when ListTypeSetUsage=Ignore
	// +kubebuilder:listType=atomic
	ObjectAtomicList []Object `json:"objectAtomicList,omitempty"`

	// +kubebuilder:listType=set
	ObjectSetList []Object `json:"objectSetList,omitempty"`

	// Missing listType on defined type to basic type - should warn
	StringNoMarker []String `json:"stringNoMarker,omitempty"` // want "StringNoMarker should have a listType marker for proper Server-Side Apply behavior \\(atomic, set, or map\\)"

	// Missing listType on defined type to object - should warn
	ObjectNoMarker []Object `json:"objectNoMarker,omitempty"` // want "ObjectNoMarker should have a listType marker for proper Server-Side Apply behavior \\(atomic, set, or map\\)"

	// Pointer to defined type - should behave same as defined type
	// +kubebuilder:listType=atomic
	PointerToString []*String `json:"pointerToString,omitempty"`

	// +kubebuilder:listType=set
	PointerToObject []*Object `json:"pointerToObject,omitempty"`

	// Type alias tests - aliases to basic types should behave as primitive lists
	// +kubebuilder:listType=atomic
	StringAliasAtomicList []StringAlias `json:"stringAliasAtomicList,omitempty"`

	// +kubebuilder:listType=set
	StringAliasSetList []StringAlias `json:"stringAliasSetList,omitempty"`

	// Type alias to object should behave as object list - no set warning when ListTypeSetUsage=Ignore
	// +kubebuilder:listType=atomic
	ObjectAliasAtomicList []ObjectAlias `json:"objectAliasAtomicList,omitempty"`

	// +kubebuilder:listType=set
	ObjectAliasSetList []ObjectAlias `json:"objectAliasSetList,omitempty"`

	// Missing listType on alias to basic type - should warn
	StringAliasNoMarker []StringAlias `json:"stringAliasNoMarker,omitempty"` // want "StringAliasNoMarker should have a listType marker for proper Server-Side Apply behavior \\(atomic, set, or map\\)"

	// Missing listType on alias to object - should warn
	ObjectAliasNoMarker []ObjectAlias `json:"objectAliasNoMarker,omitempty"` // want "ObjectAliasNoMarker should have a listType marker for proper Server-Side Apply behavior \\(atomic, set, or map\\)"

	// Pointer to alias - should behave same as alias
	// +kubebuilder:listType=atomic
	PointerToStringAlias []*StringAlias `json:"pointerToStringAlias,omitempty"`

	// +kubebuilder:listType=set
	PointerToObjectAlias []*ObjectAlias `json:"pointerToObjectAlias,omitempty"`

	// Multiple pointer levels
	PointerToPointerObject []*(*TestObject) `json:"pointerToPointerObject,omitempty"` // want "PointerToPointerObject should have a listType marker for proper Server-Side Apply behavior \\(atomic, set, or map\\)"

	// Array defined type tests - defined types to array types should behave as array lists
	// +kubebuilder:listType=atomic
	StringArrayAtomicList StringArray `json:"stringArrayAtomicList,omitempty"`

	// +kubebuilder:listType=set
	StringArraySetList StringArray `json:"stringArraySetList,omitempty"`

	// Missing listType on array defined type - should warn
	StringArrayNoMarker StringArray `json:"stringArrayNoMarker,omitempty"` // want "StringArrayNoMarker should have a listType marker for proper Server-Side Apply behavior \\(atomic, set, or map\\)"

	// Array type alias tests - aliases to array types should behave as array lists
	// +kubebuilder:listType=atomic
	StringArrayAliasAtomicList StringArrayAlias `json:"stringArrayAliasAtomicList,omitempty"`

	// +kubebuilder:listType=set
	StringArrayAliasSetList StringArrayAlias `json:"stringArrayAliasSetList,omitempty"`

	// Missing listType on array alias - should warn
	StringArrayAliasNoMarker StringArrayAlias `json:"stringArrayAliasNoMarker,omitempty"` // want "StringArrayAliasNoMarker should have a listType marker for proper Server-Side Apply behavior \\(atomic, set, or map\\)"

	// Valid map with correct JSON tag name - should pass
	// +kubebuilder:listType=map
	// +kubebuilder:listMapKey=name
	MapObjectListValid []SimpleObject `json:"mapObjectListValid,omitempty"`

	// Invalid map with non-existent JSON tag name - should error
	// +kubebuilder:listType=map
	// +kubebuilder:listMapKey=invalid_name
	MapObjectListInvalidName []SimpleObject `json:"mapObjectListInvalidName,omitempty"` // want "MapObjectListInvalidName listMapKey \"invalid_name\" does not exist as a field in the struct"
}
