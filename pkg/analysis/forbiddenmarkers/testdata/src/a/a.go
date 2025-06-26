package a

// +forbidden
type ForbiddenMarkerType string // want `type ForbiddenMarkerType has forbidden marker "forbidden"`

// +allowed
type AllowedMarkerType string

type Test struct {
	// +forbidden
	ForbiddenMarkerField string `json:"forbiddenMarkerField"`// want `field ForbiddenMarkerField has forbidden marker "forbidden"`

	ForbiddenMarkerFieldTypeAlias ForbiddenMarkerType `json:"forbiddenMarkerFieldTypeAlias"` // want `field ForbiddenMarkerFieldTypeAlias has forbidden marker "forbidden"`

	// +allowed
	AllowedMarkerField string `json:"allowedMarkerField"`

	AllowedMarkerFieldTypeAlias AllowedMarkerType `json:"AllowedMarkerFieldTypeAlias"`
}

