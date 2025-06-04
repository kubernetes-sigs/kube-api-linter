package b

type B struct {
	// +custom:SomeCustomMarker:=value
	UniqueCustomMarker string

	// +custom:SomeCustomMarker:=value
	// +custom:SomeCustomMarker:=diffvalue
	NonUniqueCustomMarker string // want "field NonUniqueCustomMarker has multiple definitions of marker custom:SomeCustomMarker when only a single definition should exist"

	// +custom:SomeCustomMarker:=value
	NonUniqueCustomMarkerFromAliasWithCustomMarker UniqueCustomMarkerAlias // want "field NonUniqueCustomMarkerFromAliasWithCustomMarker has multiple definitions of marker custom:SomeCustomMarker when only a single definition should exist"

	NonUniqueCustomMarkerOnlyFromAliasWithCustomMarker NonUniqueCustomMarkerAlias // want "field NonUniqueCustomMarkerOnlyFromAliasWithCustomMarker has multiple definitions of marker custom:SomeCustomMarker when only a single definition should exist"

	// +custom:SomeCustomMarker:=value
	// +custom:SomeCustomMarker:=value
	NonUniqueSameValueCustomMarker string // want "field NonUniqueSameValueCustomMarker has multiple definitions of marker custom:SomeCustomMarker when only a single definition should exist"

	// +custom:OtherMarker:attribute=apple,otherAttribute=orange
	UniqueCustomMarkerWithAttribute string

	// +custom:OtherMarker:attribute=apple,otherAttribute=orange
	// +custom:OtherMarker:attribute=apple,otherAttribute=banana
	NonUniqueCustomMarkerWithAttribute string // want "field NonUniqueCustomMarkerWithAttribute has multiple definitions of marker custom:OtherMarker:attribute=apple when only a single definition should exist"

	// +custom:OtherMarker:attribute=apple,otherAttribute=orange
	NonUniqueCustomMarkerWithAttributeFromAliasWithCustomMarkerWithAttribute UniqueCustomMarkerWithAttributeAlias // want "field NonUniqueCustomMarkerWithAttributeFromAliasWithCustomMarkerWithAttribute has multiple definitions of marker custom:OtherMarker:attribute=apple when only a single definition should exist"

	NonUniqueCustomMarkerWithAttributeOnlyFromAliasWithCustomMarkerWithAttribute NonUniqueCustomMarkerWithAttributeAlias // want "field NonUniqueCustomMarkerWithAttributeOnlyFromAliasWithCustomMarkerWithAttribute has multiple definitions of marker custom:OtherMarker:attribute=apple when only a single definition should exist"

	// +custom:OtherMarker:attribute=apple,otherAttribute=orange
	// +custom:OtherMarker:attribute=orange,otherAttribute=apple
	MultipleUniqueCustomMarkerWithAttribute string
}

// +custom:SomeCustomMarker:=diffvalue
type UniqueCustomMarkerAlias string

// +custom:SomeCustomMarker:=value
// +custom:SomeCustomMarker:=diffvalue
type NonUniqueCustomMarkerAlias string // want "type NonUniqueCustomMarkerAlias has multiple definitions of marker custom:SomeCustomMarker when only a single definition should exist"

// +custom:OtherMarker:attribute=apple,otherAttribute=banana
type UniqueCustomMarkerWithAttributeAlias string

// +custom:OtherMarker:attribute=apple,otherAttribute=orange
// +custom:OtherMarker:attribute=apple,otherAttribute=banana
type NonUniqueCustomMarkerWithAttributeAlias string // want "type NonUniqueCustomMarkerWithAttributeAlias has multiple definitions of marker custom:OtherMarker:attribute=apple when only a single definition should exist"
