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
	MapComponents    map[string]Component  `json:"mapComponents"`    // want "NoMapsTestStruct.MapComponents should not use a map type, use a list type with a unique name/identifier instead"
	PtrMapComponents *map[string]Component `json:"ptrMapComponents"` // want "NoMapsTestStruct.PtrMapComponents should not use a map type, use a list type with a unique name/identifier instead"
	MapStringInt     map[string]int        `json:"mapStringInt"`     // want "NoMapsTestStruct.MapStringInt should not use a map type, use a list type with a unique name/identifier instead"
	Labels           map[string]string     `json:"specialCase"`
}

type NoMapsTestStructWithDefiningType struct {
	MapStringComponent    MapStringComponent    `json:"mapStringComponent"`    // want "NoMapsTestStructWithDefiningType.MapStringComponent should not use a map type, use a list type with a unique name/identifier instead"
	PtrMapStringComponent PtrMapStringComponent `json:"ptrMapStringComponent"` // want "NoMapsTestStructWithDefiningType.PtrMapStringComponent should not use a map type, use a list type with a unique name/identifier instead"
	MapStringInt          MapStringInt          `json:"mapStringInt"`          // want "NoMapsTestStructWithDefiningType.MapStringInt should not use a map type, use a list type with a unique name/identifier instead"
	MapIntString          MapIntString          `json:"mapIntString"`          // want "NoMapsTestStructWithDefiningType.MapIntString should not use a map type, use a list type with a unique name/identifier instead"
}

type NoMapsTestStructWithDefiningTypeAcrossFiles struct {
	MapStringComponent    MapStringComponentB    `json:"mapStringComponent"`    // want "NoMapsTestStructWithDefiningTypeAcrossFiles.MapStringComponent should not use a map type, use a list type with a unique name/identifier instead"
	PtrMapStringComponent PtrMapStringComponentB `json:"ptrMapStringComponent"` // want "NoMapsTestStructWithDefiningTypeAcrossFiles.PtrMapStringComponent should not use a map type, use a list type with a unique name/identifier instead"
	MapStringInt          MapStringIntB          `json:"mapStringInt"`          // want "NoMapsTestStructWithDefiningTypeAcrossFiles.MapStringInt should not use a map type, use a list type with a unique name/identifier instead"
	MapIntString          MapIntStringB          `json:"mapIntString"`          // want "NoMapsTestStructWithDefiningTypeAcrossFiles.MapIntString should not use a map type, use a list type with a unique name/identifier instead"
}

type NoMapsTestStructWithAlias struct {
	MapStringComponentAlias           MapStringComponentAlias           `json:"mapStringComponentAlias"`           // want "NoMapsTestStructWithAlias.MapStringComponentAlias should not use a map type, use a list type with a unique name/identifier instead"
	MapStringPtrComponentAlias        MapStringPtrComponentAlias        `json:"mapStringPtrComponentAlias"`        // want "NoMapsTestStructWithAlias.MapStringPtrComponentAlias should not use a map type, use a list type with a unique name/identifier instead"
	MapStringIntAlias                 MapStringIntAlias                 `json:"mapStringIntAlias"`                 // want "NoMapsTestStructWithAlias.MapStringIntAlias should not use a map type, use a list type with a unique name/identifier instead"
	DefinedMapStringComponentAlias    DefinedMapStringComponentAlias    `json:"definedMapStringComponentAlias"`    // want "NoMapsTestStructWithAlias.DefinedMapStringComponentAlias should not use a map type, use a list type with a unique name/identifier instead"
	DefinedMapStringComponentPtrAlias DefinedMapStringComponentPtrAlias `json:"definedMapStringComponentPtrAlias"` // want "NoMapsTestStructWithAlias.DefinedMapStringComponentPtrAlias should not use a map type, use a list type with a unique name/identifier instead"
}

type NoMapsTestStructWithGenerics[K comparable, V any] struct {
	MapStringGenerics      MapStringGenerics[V]      `json:"mapStringGenerics"`      // want "NoMapsTestStructWithGenerics.MapStringGenerics should not use a map type, use a list type with a unique name/identifier instead"
	MapIntGenerics         MapIntGenerics[V]         `json:"mapIntGenerics"`         // want "NoMapsTestStructWithGenerics.MapIntGenerics should not use a map type, use a list type with a unique name/identifier instead"
	MapComparableKeyString MapComparableKeyString[K] `json:"mapComparableKeyString"` // want "NoMapsTestStructWithGenerics.MapComparableKeyString should not use a map type, use a list type with a unique name/identifier instead"
	MapComparableKeyInt    MapComparableKeyInt[K]    `json:"mapComparableKeyInt"`    // want "NoMapsTestStructWithGenerics.MapComparableKeyInt should not use a map type, use a list type with a unique name/identifier instead"
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
