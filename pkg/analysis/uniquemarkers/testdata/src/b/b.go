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
	
	// +custom:MultiMarker:fruit=apple,color=blue,country='US'
	// +custom:MultiMarker:fruit=apple,color=blue,country='UK'
	// +custom:MultiMarker:fruit=apple,color=green,country='US'
	// +custom:MultiMarker:fruit=orange,color=blue,country='US'
	// +custom:MultiMarker:fruit=orange,color=blue,country='UK'
	// +custom:MultiMarker:fruit=orange,color=green,country='US'
	UniqueMultiAttributeMarker string

	// +custom:MultiMarker:fruit=apple,color=blue,country='US',state="NY"
	// +custom:MultiMarker:fruit=apple,color=blue,country='US',state="NC"
	NonUniqueMultiAttributeMarker string // want "field NonUniqueMultiAttributeMarker has multiple definitions of marker custom:MultiMarker:fruit=apple,color=blue,country='US' when only a single definition should exist"

	// +custom:MultiMarker:fruit=apple,color=blue
	// +custom:MultiMarker:fruit=apple,color=blue,country="UK"
	// +custom:MultiMarker:fruit=apple,country="UK"
	// +custom:MultiMarker:color=blue,country="UK"
	UniqueMultiAttributeMarkerFromMissingAttribute string

	// +custom:MultiMarker:fruit=apple,color=blue
	// +custom:MultiMarker:fruit=apple,color=blue
	NonUniqueMultiAttributeMarkerFromMissingAttribute string // want "field NonUniqueMultiAttributeMarkerFromMissingAttribute has multiple definitions of marker custom:MultiMarker:fruit=apple,color=blue,country= when only a single definition should exist"
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
