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
}

// +custom:SomeCustomMarker:=diffvalue
type UniqueCustomMarkerAlias string

// +custom:SomeCustomMarker:=value
// +custom:SomeCustomMarker:=diffvalue
type NonUniqueCustomMarkerAlias string // want "type NonUniqueCustomMarkerAlias has multiple definitions of marker custom:SomeCustomMarker when only a single definition should exist"
