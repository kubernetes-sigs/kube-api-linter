package a

// +custom:forbidden
type ForbiddenMarkerType string // want `type ForbiddenMarkerType has forbidden marker "custom:forbidden"`

// +custom:AttrNoValues:fruit=apple
type ForbiddenMarkerWithAttrType string // want `type ForbiddenMarkerWithAttrType has forbidden marker "custom:AttrNoValues:fruit=apple"`

// +custom:AttrsNoValues:fruit=apple,color=blue
type ForbiddenMarkerWithMultipleAttrsType string // want `type ForbiddenMarkerWithMultipleAttrsType has forbidden marker "custom:AttrsNoValues:fruit=apple,color=blue"`

// +custom:AttrValues:fruit=orange
type ForbiddenMarkerWithAttrWithValueType string // want `type ForbiddenMarkerWithAttrWithValueType has forbidden marker "custom:AttrValues:fruit=orange"`

// +custom:AttrsValues:fruit=orange,color=blue
type ForbiddenMarkerWithAttrsWithValueType string // want `type ForbiddenMarkerWithAttrsWithValueType has forbidden marker "custom:AttrsValues:fruit=orange,color=blue"`

// +custom:MultiRuleSet:fruit=apple,color=blue
// +custom:MultiRuleSet:fruit=apple,color=green
// +custom:MultiRuleSet:fruit=orange,color=green
// +custom:MultiRuleSet:fruit=orange,color=blue
// +custom:MultiRuleSet:fruit=orange,color=red
// +custom:MultiRuleSet:fruit=banana,color=yellow
type ForbiddenMarkerWithMultipleRuleSetsType string // want `type ForbiddenMarkerWithMultipleRuleSetsType has forbidden marker "custom:MultiRuleSet:fruit=apple,color=blue"` `type ForbiddenMarkerWithMultipleRuleSetsType has forbidden marker "custom:MultiRuleSet:fruit=apple,color=green"` `type ForbiddenMarkerWithMultipleRuleSetsType has forbidden marker "custom:MultiRuleSet:fruit=orange,color=green"` `type ForbiddenMarkerWithMultipleRuleSetsType has forbidden marker "custom:MultiRuleSet:fruit=orange,color=blue"` `type ForbiddenMarkerWithMultipleRuleSetsType has forbidden marker "custom:MultiRuleSet:fruit=orange,color=red"` `type ForbiddenMarkerWithMultipleRuleSetsType has forbidden marker "custom:MultiRuleSet:fruit=banana,color=yellow"`

// +allowed
type AllowedMarkerType string

// Allowed because the configuration only forbids when the fruit attributes are specified
// +custom:AttrNoValues:color=blue
type AllowedMarkerWithAttrType string

// Allowed because the configuration only forbids when both fruit and color attributes are specified
// +custom:AttrsNoValues:fruit=apple
type AllowedMarkerWithMultipleAttrsType string

// Allowed because the configuration only forbids when the fruit attribute is one of apple, banana, or orange
// +custom:AttrValues:fruit=cherry
type AllowedMarkerWithAttrWithValueType string

// Allowed because the configuration only forbids when the fruit attribute is one of apple, banana, or orange
// and the color attribute is one of blue, red, or green
// +custom:AttrsValues:fruit=cherry,color=blue
type AllowedMarkerWithAttrsWithValueType string

// +custom:MultiRuleSet:fruit=cherry,color=blue
// +custom:MultiRuleSet:fruit=apple,color=red
// +custom:MultiRuleSet:fruit=apple,color=cyan
// +custom:MultiRuleSet:fruit=orange,color=orange
// +custom:MultiRuleSet:fruit=orange,color=purple
type AllowedMarkerWithMultipleRuleSetsType string

type Test struct {
	// +custom:forbidden
	ForbiddenMarkerField string `json:"forbiddenMarkerField"` // want `field ForbiddenMarkerField has forbidden marker "custom:forbidden"`

	ForbiddenMarkerFieldTypeAlias ForbiddenMarkerType `json:"forbiddenMarkerFieldTypeAlias"` // want `field ForbiddenMarkerFieldTypeAlias has forbidden marker "custom:forbidden"`

	// +custom:AttrNoValues:fruit=apple
	ForbiddenMarkerWithAttrField string `json:"forbiddenMarkerWithAttrField"` // want `field ForbiddenMarkerWithAttrField has forbidden marker "custom:AttrNoValues:fruit=apple"`

	ForbiddenMarkerWithAttrFieldTypeAlias ForbiddenMarkerWithAttrType `json:"forbiddenMarkerWithAttrFieldTypeAlias"` // want `field ForbiddenMarkerWithAttrFieldTypeAlias has forbidden marker "custom:AttrNoValues:fruit=apple"`

	// +custom:AttrsNoValues:fruit=apple,color=blue
	ForbiddenMarkerWithMultipleAttrsField string `json:"forbiddenMarkerWithMultipleAttrsField"` // want `field ForbiddenMarkerWithMultipleAttrsField has forbidden marker "custom:AttrsNoValues:fruit=apple,color=blue"`

	ForbiddenMarkerWithMutlipleAttrsFieldTypeAlias ForbiddenMarkerWithMultipleAttrsType `json:"forbiddenMarkerWithMultipleAttrsFieldTypeAlias"` // want `field ForbiddenMarkerWithMutlipleAttrsFieldTypeAlias has forbidden marker "custom:AttrsNoValues:fruit=apple,color=blue"`

	// +custom:AttrValues:fruit=orange
	ForbiddenMarkerWithAttrWithValueField string `json:"forbiddenMarkerWithAttrWithValueField"` // want `field ForbiddenMarkerWithAttrWithValueField has forbidden marker "custom:AttrValues:fruit=orange"`

	ForbiddenMarkerWithAttrWithValueFieldTypeAlias ForbiddenMarkerWithAttrWithValueType `json:"forbiddenMarkerWithAttrWithValueFieldTypeAlias"` // want `field ForbiddenMarkerWithAttrWithValueFieldTypeAlias has forbidden marker "custom:AttrValues:fruit=orange"`

	// +custom:AttrsValues:fruit=orange,color=blue
	ForbiddenMarkerWithAttrsWithValueField string `json:"forbiddenMarkerWithAttrsWithValueField"` // want `field ForbiddenMarkerWithAttrsWithValueField has forbidden marker "custom:AttrsValues:fruit=orange,color=blue"`

	ForbiddenMarkerWithAttrsWithValueFieldTypeAlias ForbiddenMarkerWithAttrsWithValueType `json:"forbiddenMarkerWithAttrsWithValueFieldTypeAlias"` // want `field ForbiddenMarkerWithAttrsWithValueFieldTypeAlias has forbidden marker "custom:AttrsValues:fruit=orange,color=blue"`

	// +allowed
	AllowedMarkerField string `json:"allowedMarkerField"`

	AllowedMarkerFieldTypeAlias AllowedMarkerType `json:"allowedMarkerFieldTypeAlias"`

	// Allowed because the configuration only forbids when the fruit attributes are specified
	// +custom:AttrNoValues:color=blue
	AllowedMarkerWithAttrField string `json:"allowedMarkerWithAttrField"`

	AllowedMarkerWithAttrFieldTypeAlias AllowedMarkerWithAttrType `json:"allowedMarkerWithAttrFieldTypeAlias"`

	// Allowed because the configuration only forbids when both fruit and color attributes are specified
	// +custom:AttrsNoValues:fruit=apple
	AllowedMarkerWithMultipleAttrsField string `json:"allowedMarkerWithMultipleAttrsField"`

	AllowedMarkerWithMultipleAttrsFieldTypeAlias AllowedMarkerWithMultipleAttrsType `json:"allowedMarkerWithMultipleAttrsFieldTypeAlias"`

	// Allowed because the configuration only forbids when the fruit attribute is one of apple, banana, or orange
	// +custom:AttrsNoValues:fruit=cherry
	AllowedMarkerWithAttrWithValueField string `json:"allowedMarkerWithAttrWithValueField"`

	AllowedMarkerWithAttrWithValueFieldTypeAlias AllowedMarkerWithAttrWithValueType `json:"allowedMarkerWithAttrWithValueFieldTypeAlias"`

	// Allowed because the configuration only forbids when the fruit attribute is one of apple, banana, or orange
	// and the color attribute is one of blue, red, or green
	// +custom:AttrsValues:fruit=cherry,color=blue
	AllowedMarkerWithAttrsWithValueField string `json:"allowedMarkerWithAttrsWithValueField"`

	AllowedMarkerWithAttrsWithValueFieldTypeAlias AllowedMarkerWithAttrsWithValueType `json:"allowedMarkerWithAttrsWithValueFieldTypeAlias"`
}
