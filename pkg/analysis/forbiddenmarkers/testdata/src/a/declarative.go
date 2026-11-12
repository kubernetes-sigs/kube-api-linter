package a

// +k8s:forbidden
type DeclarativeForbiddenMarkerType string // want `type DeclarativeForbiddenMarkerType has forbidden marker "k8s:forbidden"`

// +k8s:alpha(since: "1.36")=+k8s:forbidden
type DeclarativeNestedAlphaForbiddenMarkerType string // want `type DeclarativeNestedAlphaForbiddenMarkerType has forbidden marker "k8s:alpha\(since: \\"1.36\\"\)=\+k8s:forbidden"`

// +k8s:ifEnabled("MyGate")=+k8s:forbidden
type DeclarativeNestedIfEnabledForbiddenMarkerType string // want `type DeclarativeNestedIfEnabledForbiddenMarkerType has forbidden marker "k8s:ifEnabled\(\\"MyGate\\"\)=\+k8s:forbidden"`

// +k8s:alpha(since: "1.36")=+k8s:ifEnabled("MyGate")=+k8s:forbidden
type DeclarativeDeeplyNestedForbiddenMarkerType string // want `type DeclarativeDeeplyNestedForbiddenMarkerType has forbidden marker "k8s:alpha\(since: \\"1.36\\"\)=\+k8s:ifEnabled\(\\"MyGate\\"\)=\+k8s:forbidden"`

// +k8s:alpha(since: "1.36")=+k8s:forbiddenWithAttr(fruit: "apple")
type DeclarativeNestedForbiddenMarkerWithAttrType string // want `type DeclarativeNestedForbiddenMarkerWithAttrType has forbidden marker "k8s:alpha\(since: \\"1.36\\"\)=\+k8s:forbiddenWithAttr\(fruit: \\"apple\\"\)"`

// Allowed because the configuration only forbids when the fruit attribute is apple
// +k8s:alpha(since: "1.36")=+k8s:forbiddenWithAttr(fruit: "cherry")
type DeclarativeNestedAllowedMarkerWithAttrType string

// Allowed because neither the outer nor the inner marker is forbidden
// +k8s:alpha(since: "1.36")=+k8s:maxLength=10
type DeclarativeNestedAllowedMarkerType string

type DeclarativeTest struct {
	// +k8s:forbidden
	DeclarativeForbiddenMarkerField string `json:"declarativeForbiddenMarkerField"` // want `field DeclarativeTest.DeclarativeForbiddenMarkerField has forbidden marker "k8s:forbidden"`

	// +k8s:alpha(since: "1.36")=+k8s:forbidden
	DeclarativeNestedAlphaForbiddenMarkerField string `json:"declarativeNestedAlphaForbiddenMarkerField"` // want `field DeclarativeTest.DeclarativeNestedAlphaForbiddenMarkerField has forbidden marker "k8s:alpha\(since: \\"1.36\\"\)=\+k8s:forbidden"`

	// +k8s:ifEnabled("MyGate")=+k8s:forbidden
	DeclarativeNestedIfEnabledForbiddenMarkerField string `json:"declarativeNestedIfEnabledForbiddenMarkerField"` // want `field DeclarativeTest.DeclarativeNestedIfEnabledForbiddenMarkerField has forbidden marker "k8s:ifEnabled\(\\"MyGate\\"\)=\+k8s:forbidden"`

	// +k8s:alpha(since: "1.36")=+k8s:ifEnabled("MyGate")=+k8s:forbidden
	DeclarativeDeeplyNestedForbiddenMarkerField string `json:"declarativeDeeplyNestedForbiddenMarkerField"` // want `field DeclarativeTest.DeclarativeDeeplyNestedForbiddenMarkerField has forbidden marker "k8s:alpha\(since: \\"1.36\\"\)=\+k8s:ifEnabled\(\\"MyGate\\"\)=\+k8s:forbidden"`

	// +k8s:alpha(since: "1.36")=+k8s:forbiddenWithAttr(fruit: "apple")
	DeclarativeNestedForbiddenMarkerWithAttrField string `json:"declarativeNestedForbiddenMarkerWithAttrField"` // want `field DeclarativeTest.DeclarativeNestedForbiddenMarkerWithAttrField has forbidden marker "k8s:alpha\(since: \\"1.36\\"\)=\+k8s:forbiddenWithAttr\(fruit: \\"apple\\"\)"`

	DeclarativeNestedAlphaForbiddenMarkerFieldTypeAlias DeclarativeNestedAlphaForbiddenMarkerType `json:"declarativeNestedAlphaForbiddenMarkerFieldTypeAlias"` // want `field DeclarativeTest.DeclarativeNestedAlphaForbiddenMarkerFieldTypeAlias has forbidden marker "k8s:alpha\(since: \\"1.36\\"\)=\+k8s:forbidden"`

	DeclarativeDeeplyNestedForbiddenMarkerFieldTypeAlias DeclarativeDeeplyNestedForbiddenMarkerType `json:"declarativeDeeplyNestedForbiddenMarkerFieldTypeAlias"` // want `field DeclarativeTest.DeclarativeDeeplyNestedForbiddenMarkerFieldTypeAlias has forbidden marker "k8s:alpha\(since: \\"1.36\\"\)=\+k8s:ifEnabled\(\\"MyGate\\"\)=\+k8s:forbidden"`

	// Allowed because the configuration only forbids when the fruit attribute is apple
	// +k8s:alpha(since: "1.36")=+k8s:forbiddenWithAttr(fruit: "cherry")
	DeclarativeNestedAllowedMarkerWithAttrField string `json:"declarativeNestedAllowedMarkerWithAttrField"`

	// Allowed because neither the outer nor the inner marker is forbidden
	// +k8s:alpha(since: "1.36")=+k8s:maxLength=10
	DeclarativeNestedAllowedMarkerField string `json:"declarativeNestedAllowedMarkerField"`

	DeclarativeNestedAllowedMarkerFieldTypeAlias DeclarativeNestedAllowedMarkerType `json:"declarativeNestedAllowedMarkerFieldTypeAlias"`
}
