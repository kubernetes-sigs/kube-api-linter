package a

// +custom:forbidden
type ForbiddenMarkerType string // want `type ForbiddenMarkerType has forbidden marker "custom:forbidden"`

// +custom:AttrNoValues:fruit=apple
type ForbiddenMarkerWithAttrType string // want `type ForbiddenMarkerWithAttrType has forbidden marker "forbidden:AttrNoValues"`

// +allowed
type AllowedMarkerType string

// +custom:AttrNoValues:color=blue
type AllowedMarkerWithAttrType string

type Test struct {
	// +custom:forbidden
	ForbiddenMarkerField string `json:"forbiddenMarkerField"`// want `field ForbiddenMarkerField has forbidden marker "custom:forbidden"`

	ForbiddenMarkerFieldTypeAlias ForbiddenMarkerType `json:"forbiddenMarkerFieldTypeAlias"` // want `field ForbiddenMarkerFieldTypeAlias has forbidden marker "forbidden"`

	// +allowed
	AllowedMarkerField string `json:"allowedMarkerField"`

	AllowedMarkerFieldTypeAlias AllowedMarkerType `json:"AllowedMarkerFieldTypeAlias"`
}
