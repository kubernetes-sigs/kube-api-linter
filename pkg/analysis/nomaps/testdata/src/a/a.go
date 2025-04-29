package a

type (
	MapStringComponent    map[string]Component
	PtrMapStringComponent *map[string]Component
	MapStringInt          map[string]int
	MapIntString          map[int]string
)

type (
	MapStringComponentAlias           = map[string]Component
	MapStringPtrComponentAlias        = *map[string]Component
	MapStringIntAlias                 = map[string]int
	DefinedMapStringComponentAlias    = MapStringComponent
	DefinedMapStringComponentPtrAlias = *MapStringComponent
)

type (
	MapStringGenerics[V any]             map[string]V
	MapIntGenerics[V any]                map[int]V
	MapComparableKeyString[K comparable] map[K]string
	MapComparableKeyInt[K comparable]    map[K]int
)

type NoMapsTestStruct struct {
	Primitive        int32                 `json:"primitive"`
	Components       []Component           `json:"components"`
	MapComponents    map[string]Component  `json:"mapComponents"` // want "MapComponents should not use a map type, use a list type with a unique name/identifier instead"
	PtrMapComponents *map[string]Component `json:"mapComponents"` // want "MapComponents should not use a map type, use a list type with a unique name/identifier instead"
	MapStringInt     map[string]int        `json:"mapStringInt"`  // want "MapStringInt should not use a map type, use a list type with a unique name/identifier instead"
	Labels           map[string]string     `json:"specialCase"`
}

type NoMapsTestStructWithDefiningType struct {
	MapStringComponent    MapStringComponent    `json:"mapStringComponent"`    // want "MapStringComponent should not use a map type, use a list type with a unique name/identifier instead"
	PtrMapStringComponent PtrMapStringComponent `json:"ptrMapStringComponent"` // want "PtrMapStringComponent should not use a map type, use a list type with a unique name/identifier instead"
	MapStringInt          MapStringInt          `json:"mapStringInt"`          // want "MapStringInt should not use a map type, use a list type with a unique name/identifier instead"
	MapIntString          MapIntString          `json:"mapIntString"`          // want "MapIntString should not use a map type, use a list type with a unique name/identifier instead"
}

type NoMapsTestStructWithDefiningTypeAcrossFiles struct {
	MapStringComponent    MapStringComponentB    `json:"mapStringComponent"`    // want "MapStringComponent should not use a map type, use a list type with a unique name/identifier instead"
	PtrMapStringComponent PtrMapStringComponentB `json:"ptrMapStringComponent"` // want "PtrMapStringComponent should not use a map type, use a list type with a unique name/identifier instead"
	MapStringInt          MapStringIntB          `json:"mapStringInt"`          // want "MapStringInt should not use a map type, use a list type with a unique name/identifier instead"
	MapIntString          MapIntStringB          `json:"mapIntString"`          // want "MapIntString should not use a map type, use a list type with a unique name/identifier instead"
}

type NoMapsTestStructWithAlias struct {
	MapStringComponentAlias           MapStringComponentAlias           `json:"mapStringComponentAlias"`           // want "MapStringComponentAlias should not use a map type, use a list type with a unique name/identifier instead"
	MapStringPtrComponentAlias        MapStringPtrComponentAlias        `json:"mapStringPtrComponentAlias"`        // want "MapStringPtrComponentAlias should not use a map type, use a list type with a unique name/identifier instead"
	MapStringIntAlias                 MapStringIntAlias                 `json:"mapStringIntAlias"`                 // want "MapStringIntAlias should not use a map type, use a list type with a unique name/identifier instead"
	DefinedMapStringComponentAlias    DefinedMapStringComponentAlias    `json:"definedMapStringComponentAlias"`    // want "DefinedMapStringComponentAlias should not use a map type, use a list type with a unique name/identifier instead"
	DefinedMapStringComponentPtrAlias DefinedMapStringComponentPtrAlias `json:"definedMapStringComponentPtrAlias"` // want "DefinedMapStringComponentPtrAlias should not use a map type, use a list type with a unique name/identifier instead"
}

type NoMapsTestStructWithGenerics[K comparable, V any] struct {
	MapStringGenerics      MapStringGenerics[V]      `json:"mapStringGenerics"`      // want "MapStringGenerics should not use a map type, use a list type with a unique name/identifier instead"
	MapIntGenerics         MapIntGenerics[V]         `json:"mapIntGenerics"`         // want "MapIntGenerics should not use a map type, use a list type with a unique name/identifier instead"
	MapComparableKeyString MapComparableKeyString[K] `json:"mapComparableKeyString"` // want "MapComparableKeyString should not use a map type, use a list type with a unique name/identifier instead"
	MapComparableKeyInt    MapComparableKeyInt[K]    `json:"mapComparableKeyInt"`    // want "MapComparableKeyInt should not use a map type, use a list type with a unique name/identifier instead"
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
