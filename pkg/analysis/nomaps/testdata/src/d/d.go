package d

type (
	MapStringComponent    map[string]Component  // want "type MapStringComponent should not use a map type, use a list type with a unique name/identifier instead"
	PtrMapStringComponent *map[string]Component // want "type PtrMapStringComponent pointer should not use a map type, use a list type with a unique name/identifier instead"
	MapStringInt          map[string]int
	MapIntString          map[int]string
)

type (
	MapStringComponentAlias           = map[string]Component  // want "type MapStringComponentAlias should not use a map type, use a list type with a unique name/identifier instead"
	MapStringPtrComponentAlias        = *map[string]Component // want "type MapStringPtrComponentAlias pointer should not use a map type, use a list type with a unique name/identifier instead"
	MapStringIntAlias                 = map[string]int
	DefinedMapStringComponentAlias    = MapStringComponent  // want "type DefinedMapStringComponentAlias type MapStringComponent should not use a map type, use a list type with a unique name/identifier instead"
	DefinedMapStringComponentPtrAlias = *MapStringComponent // want "type DefinedMapStringComponentPtrAlias pointer type MapStringComponent should not use a map type, use a list type with a unique name/identifier instead"
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
	MapComponents    map[string]Component  `json:"mapComponents"` // want "field NoMapsTestStruct.MapComponents should not use a map type, use a list type with a unique name/identifier instead"
	PtrMapComponents *map[string]Component `json:"mapComponents"` // want "field NoMapsTestStruct.PtrMapComponents pointer should not use a map type, use a list type with a unique name/identifier instead"
	MapStringInt     map[string]int        `json:"mapStringInt"`
	Labels           map[string]string     `json:"specialCase"`
}

type NoMapsTestStructWithDefiningType struct {
	MapStringComponent    MapStringComponent    `json:"mapStringComponent"`    // want "field NoMapsTestStructWithDefiningType.MapStringComponent type MapStringComponent should not use a map type, use a list type with a unique name/identifier instead"
	PtrMapStringComponent PtrMapStringComponent `json:"ptrMapStringComponent"` // want "field NoMapsTestStructWithDefiningType.PtrMapStringComponent type PtrMapStringComponent pointer should not use a map type, use a list type with a unique name/identifier instead"
	MapStringInt          MapStringInt          `json:"mapStringInt"`
	MapIntString          MapIntString          `json:"mapIntString"`
}

type NoMapsTestStructWithAlias struct {
	MapStringComponentAlias           MapStringComponentAlias           `json:"mapStringComponentAlias"`    // want "field NoMapsTestStructWithAlias.MapStringComponentAlias type MapStringComponentAlias should not use a map type, use a list type with a unique name/identifier instead"
	MapStringPtrComponentAlias        MapStringPtrComponentAlias        `json:"mapStringPtrComponentAlias"` // want "field NoMapsTestStructWithAlias.MapStringPtrComponentAlias type MapStringPtrComponentAlias pointer should not use a map type, use a list type with a unique name/identifier instead"
	MapStringIntAlias                 MapStringIntAlias                 `json:"mapStringIntAlias"`
	DefinedMapStringComponentAlias    DefinedMapStringComponentAlias    `json:"definedMapStringComponentAlias"`    // want "field NoMapsTestStructWithAlias.DefinedMapStringComponentAlias type DefinedMapStringComponentAlias type MapStringComponent should not use a map type, use a list type with a unique name/identifier instead"
	DefinedMapStringComponentPtrAlias DefinedMapStringComponentPtrAlias `json:"definedMapStringComponentPtrAlias"` // want "field NoMapsTestStructWithAlias.DefinedMapStringComponentPtrAlias type DefinedMapStringComponentPtrAlias pointer type MapStringComponent should not use a map type, use a list type with a unique name/identifier instead"
}

type NoMapsTestStructWithGenerics[K comparable, V any] struct {
	MapStringGenerics      MapStringGenerics[V]      `json:"mapStringGenerics"`      // want "field NoMapsTestStructWithGenerics.MapStringGenerics type MapStringGenerics should not use a map type, use a list type with a unique name/identifier instead"
	MapIntGenerics         MapIntGenerics[V]         `json:"mapIntGenerics"`         // want "field NoMapsTestStructWithGenerics.MapIntGenerics type MapIntGenerics should not use a map type, use a list type with a unique name/identifier instead"
	MapComparableKeyString MapComparableKeyString[K] `json:"mapComparableKeyString"` // want "field NoMapsTestStructWithGenerics.MapComparableKeyString type MapComparableKeyString should not use a map type, use a list type with a unique name/identifier instead"
	MapComparableKeyInt    MapComparableKeyInt[K]    `json:"mapComparableKeyInt"`    // want "field NoMapsTestStructWithGenerics.MapComparableKeyInt type MapComparableKeyInt should not use a map type, use a list type with a unique name/identifier instead"
}

type NoMapsTestCompositeLiteral struct {
	MapStringArray     map[string][5]string         `json:"arrayMap"`     // want "field NoMapsTestCompositeLiteral.MapStringArray should not use a map type, use a list type with a unique name/identifier instead"
	MapStringStruct    map[string]struct{}          `json:"structMap"`    // want "field NoMapsTestCompositeLiteral.MapStringStruct should not use a map type, use a list type with a unique name/identifier instead"
	MapStringPointer   map[string]*string           `json:"pointerMap"`   // want "field NoMapsTestCompositeLiteral.MapStringPointer should not use a map type, use a list type with a unique name/identifier instead"
	MapStringFunc      map[string]func()            `json:"funcMap"`      // want "field NoMapsTestCompositeLiteral.MapStringFunc should not use a map type, use a list type with a unique name/identifier instead"
	MapStringInterface map[string]any               `json:"interfaceMap"` // want "field NoMapsTestCompositeLiteral.MapStringInterface should not use a map type, use a list type with a unique name/identifier instead"
	MapStringSlice     map[string][]string          `json:"sliceMap"`     // want "field NoMapsTestCompositeLiteral.MapStringSlice should not use a map type, use a list type with a unique name/identifier instead"
	MapStringMap       map[string]map[string]string `json:"mapMap"`       // want "field NoMapsTestCompositeLiteral.MapStringMap should not use a map type, use a list type with a unique name/identifier instead"
	MapStringChan      map[string]chan string       `json:"chanMap"`      // want "field NoMapsTestCompositeLiteral.MapStringChan should not use a map type, use a list type with a unique name/identifier instead"
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
