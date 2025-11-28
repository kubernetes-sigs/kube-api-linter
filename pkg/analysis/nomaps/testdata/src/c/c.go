package c

type (
	MapStringComponent    map[string]Component  // want "type MapStringComponent should not use a map type, use a list type with a unique name/identifier instead"
	PtrMapStringComponent *map[string]Component // want "type PtrMapStringComponent pointer should not use a map type, use a list type with a unique name/identifier instead"
	MapStringInt          map[string]int        // want "type MapStringInt should not use a map type, use a list type with a unique name/identifier instead"
	MapIntString          map[int]string        // want "type MapIntString should not use a map type, use a list type with a unique name/identifier instead"
)

type (
	MapStringComponentAlias           = map[string]Component  // want "type MapStringComponentAlias should not use a map type, use a list type with a unique name/identifier instead"
	MapStringPtrComponentAlias        = *map[string]Component // want "type MapStringPtrComponentAlias pointer should not use a map type, use a list type with a unique name/identifier instead"
	MapStringIntAlias                 = map[string]int        // want "type MapStringIntAlias should not use a map type, use a list type with a unique name/identifier instead"
	DefinedMapStringComponentAlias    = MapStringComponent    // want "type DefinedMapStringComponentAlias type MapStringComponent should not use a map type, use a list type with a unique name/identifier instead"
	DefinedMapStringComponentPtrAlias = *MapStringComponent   // want "type DefinedMapStringComponentPtrAlias pointer type MapStringComponent should not use a map type, use a list type with a unique name/identifier instead"
)

type (
	MapStringGenerics[V any]             map[string]V // want "type MapStringGenerics should not use a map type, use a list type with a unique name/identifier instead"
	MapIntGenerics[V any]                map[int]V    // want "type MapIntGenerics should not use a map type, use a list type with a unique name/identifier instead"
	MapComparableKeyString[K comparable] map[K]string // want "type MapComparableKeyString should not use a map type, use a list type with a unique name/identifier instead"
	MapComparableKeyInt[K comparable]    map[K]int    // want "type MapComparableKeyInt should not use a map type, use a list type with a unique name/identifier instead"
)

type NoMapsTestStruct struct {
	Primitive        int32                 `json:"primitive"`
	Components       []Component           `json:"components"`
	MapComponents    map[string]Component  `json:"mapComponents"`    // want "NoMapsTestStruct.MapComponents should not use a map type, use a list type with a unique name/identifier instead"
	PtrMapComponents *map[string]Component `json:"ptrMapComponents"` // want "NoMapsTestStruct.PtrMapComponents pointer should not use a map type, use a list type with a unique name/identifier instead"
	MapStringInt     map[string]int        `json:"mapStringInt"`     // want "NoMapsTestStruct.MapStringInt should not use a map type, use a list type with a unique name/identifier instead"
	Labels           map[string]string     `json:"specialCase"`
}

type NoMapsTestStructWithDefiningType struct {
	MapStringComponent    MapStringComponent    `json:"mapStringComponent"`    // want "NoMapsTestStructWithDefiningType.MapStringComponent type MapStringComponent should not use a map type, use a list type with a unique name/identifier instead"
	PtrMapStringComponent PtrMapStringComponent `json:"ptrMapStringComponent"` // want "NoMapsTestStructWithDefiningType.PtrMapStringComponent type PtrMapStringComponent pointer should not use a map type, use a list type with a unique name/identifier instead"
	MapStringInt          MapStringInt          `json:"mapStringInt"`          // want "NoMapsTestStructWithDefiningType.MapStringInt type MapStringInt should not use a map type, use a list type with a unique name/identifier instead"
	MapIntString          MapIntString          `json:"mapIntString"`          // want "NoMapsTestStructWithDefiningType.MapIntString type MapIntString should not use a map type, use a list type with a unique name/identifier instead"
}

type NoMapsTestStructWithAlias struct {
	MapStringComponentAlias           MapStringComponentAlias           `json:"mapStringComponentAlias"`           // want "NoMapsTestStructWithAlias.MapStringComponentAlias type MapStringComponentAlias should not use a map type, use a list type with a unique name/identifier instead"
	MapStringPtrComponentAlias        MapStringPtrComponentAlias        `json:"mapStringPtrComponentAlias"`        // want "NoMapsTestStructWithAlias.MapStringPtrComponentAlias type MapStringPtrComponentAlias pointer should not use a map type, use a list type with a unique name/identifier instead"
	MapStringIntAlias                 MapStringIntAlias                 `json:"mapStringIntAlias"`                 // want "NoMapsTestStructWithAlias.MapStringIntAlias type MapStringIntAlias should not use a map type, use a list type with a unique name/identifier instead"
	DefinedMapStringComponentAlias    DefinedMapStringComponentAlias    `json:"definedMapStringComponentAlias"`    // want "NoMapsTestStructWithAlias.DefinedMapStringComponentAlias type DefinedMapStringComponentAlias type MapStringComponent should not use a map type, use a list type with a unique name/identifier instead"
	DefinedMapStringComponentPtrAlias DefinedMapStringComponentPtrAlias `json:"definedMapStringComponentPtrAlias"` // want "NoMapsTestStructWithAlias.DefinedMapStringComponentPtrAlias type DefinedMapStringComponentPtrAlias pointer type MapStringComponent should not use a map type, use a list type with a unique name/identifier instead"
}

type NoMapsTestStructWithGenerics[K comparable, V any] struct {
	MapStringGenerics      MapStringGenerics[V]      `json:"mapStringGenerics"`      // want "NoMapsTestStructWithGenerics.MapStringGenerics type MapStringGenerics should not use a map type, use a list type with a unique name/identifier instead"
	MapIntGenerics         MapIntGenerics[V]         `json:"mapIntGenerics"`         // want "NoMapsTestStructWithGenerics.MapIntGenerics type MapIntGenerics should not use a map type, use a list type with a unique name/identifier instead"
	MapComparableKeyString MapComparableKeyString[K] `json:"mapComparableKeyString"` // want "NoMapsTestStructWithGenerics.MapComparableKeyString type MapComparableKeyString should not use a map type, use a list type with a unique name/identifier instead"
	MapComparableKeyInt    MapComparableKeyInt[K]    `json:"mapComparableKeyInt"`    // want "NoMapsTestStructWithGenerics.MapComparableKeyInt type MapComparableKeyInt should not use a map type, use a list type with a unique name/identifier instead"
}

type NoMapsTestStructWithEmbedded struct {
	NoMapsTestStruct
	NoMapsTestStructWithDefiningType
	NoMapsTestStructWithGenerics[string, Component]
	NoMapsTestStructWithAlias
}

type Component struct {
	Key   string `json:"key"`
	Value int32  `json:"value"`
}

type StringBasedType string

type MapWithStringBasedTypes struct {
	MapWithStringBasedElement    map[string]StringBasedType          `json:"stringBasedMapElem,omitempty"`
	MapWithStringBasedKey        map[StringBasedType]string          `json:"stringBasedMapKey,omitempty"`
	MapWithStringBasedKeyAndElem map[StringBasedType]StringBasedType `json:"stringBasedMapKey,omitempty"`
}
