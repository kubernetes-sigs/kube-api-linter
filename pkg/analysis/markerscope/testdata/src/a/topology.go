/*
Copyright 2025 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package a

// Valid: listType on array type
// +listType=map
type ItemList []Item

// Valid: listMapKey on array type
// +listType=map
// +listMapKey=name
type ItemListWithKey []Item

// Valid: mapType on map type
// +mapType=granular
type ConfigMap map[string]string

// Valid: structType on struct type
// +structType=atomic
type AtomicStruct struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
}

// Invalid: listType on non-array type
// +listType=map // want `marker "listType": type object is not allowed \(expected one of: \[array\]\)`
type InvalidListTypeOnStruct struct {
	Field string `json:"field"`
}

// Invalid: mapType on non-map type
// +mapType=granular // want `marker "mapType": type array is not allowed \(expected one of: \[object\]\)`
type InvalidMapTypeOnArray []string

// Invalid: structType on non-struct type
// +structType=atomic // want `marker "structType": type string is not allowed \(expected one of: \[object\]\)`
type InvalidStructTypeOnString string

type TopologyMarkersFieldTest struct {
	// Valid: listType marker on array field
	// +listType=map
	ValidListType []Item `json:"validListType"`

	// Valid: listMapKey marker on array field
	// +listType=map
	// +listMapKey=name
	ValidListMapKey []Item `json:"validListMapKey"`

	// Valid: mapType marker on map field
	// +mapType=granular
	ValidMapType map[string]string `json:"validMapType"`

	// Valid: structType marker on struct field
	// +structType=atomic
	ValidStructType EmbeddedStruct `json:"validStructType"`

	// Valid: Using named type with listType
	ValidItemListTyped ItemList `json:"validItemListTyped"`

	// Valid: Using named type with listMapKey
	ValidItemListWithKeyTyped ItemListWithKey `json:"validItemListWithKeyTyped"`

	// Valid: Using named type with mapType
	ValidConfigMapTyped ConfigMap `json:"validConfigMapTyped"`

	// Valid: Using named type with structType
	ValidAtomicStructTyped AtomicStruct `json:"validAtomicStructTyped"`

	// Invalid: listType marker on named type
	// +listType=map // want `marker "listType": marker should be declared on the type definition of ItemList instead of the field`
	InvalidListTypeOnItemList ItemList `json:"invalidListTypeOnItemList"`

	// Invalid: listMapKey marker on named type
	// +listMapKey=name // want `marker "listMapKey": marker should be declared on the type definition of ItemList instead of the field`
	InvalidListMapKeyOnItemList ItemList `json:"invalidListMapKeyOnItemList"`

	// Invalid: mapType marker on named type
	// +mapType=granular // want `marker "mapType": marker should be declared on the type definition of ConfigMap instead of the field`
	InvalidMapTypeOnConfigMap ConfigMap `json:"invalidMapTypeOnConfigMap"`

	// Invalid: structType marker on named type
	// +structType=atomic
	InvalidStructTypeOnAtomicStruct AtomicStruct `json:"invalidStructTypeOnAtomicStruct"`

	// Invalid: listType marker on string field
	// +listType=map // want `marker "listType": type string is not allowed \(expected one of: \[array\]\)`
	InvalidListTypeOnString string `json:"invalidListTypeOnString"`

	// Invalid: mapType marker on array field
	// +mapType=granular // want `marker "mapType": type array is not allowed \(expected one of: \[object\]\)`
	InvalidMapTypeOnArray []string `json:"invalidMapTypeOnArray"`

	// Invalid: structType marker on integer field
	// +structType=atomic // want `marker "structType": type integer is not allowed \(expected one of: \[object\]\)`
	InvalidStructTypeOnInt int32 `json:"invalidStructTypeOnInt"`

	// Invalid: Using invalid named type
	// +listType=map // want `marker "listType": type object is not allowed \(expected one of: \[array\]\)`
	InvalidListTypeOnStructTyped InvalidListTypeOnStruct `json:"invalidListTypeOnStructTyped"`

	// Invalid: Using invalid named type
	InvalidMapTypeOnArrayTyped InvalidMapTypeOnArray `json:"invalidMapTypeOnArrayTyped"`

	// Invalid: Using invalid named type
	InvalidStructTypeOnStringTyped InvalidStructTypeOnString `json:"invalidStructTypeOnStringTyped"`
}

type EmbeddedStruct struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
}

type Item struct {
	Name string `json:"name"`
}
