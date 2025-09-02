package a

// +nullable
type NullableMarkerType string // want `type NullableMarkerType has forbidden marker "nullable"`

// +allowed
type AllowedMarkerType string

type Test struct {
	// +nullable
	NullableMarkerField string `json:"nullableMarkerField"`// want `field NullableMarkerField has forbidden marker "nullable"`

	NullableMarkerFieldTypeAlias NullableMarkerType `json:"nullableMarkerFieldTypeAlias"` // want `field NullableMarkerFieldTypeAlias has forbidden marker "nullable"`

	// +allowed
	AllowedMarkerField string `json:"allowedMarkerField"`

	AllowedMarkerFieldTypeAlias AllowedMarkerType `json:"AllowedMarkerFieldTypeAlias"`
}

