package a

type (
	MapStringComponentB    map[string]Component  // want "type MapStringComponentB should not use a map type, use a list type with a unique name/identifier instead"
	PtrMapStringComponentB *map[string]Component // want "type PtrMapStringComponentB pointer should not use a map type, use a list type with a unique name/identifier instead"
	MapStringIntB          map[string]int        // want "type MapStringIntB should not use a map type, use a list type with a unique name/identifier instead"
	MapIntStringB          map[int]string        // want "type MapIntStringB should not use a map type, use a list type with a unique name/identifier instead"
)
