package a

type K8s struct {
	// +k8s:format:=date-time
	UniqueFormat string

	// +k8s:format:=date-time
	// +k8s:format:=password
	NonUniqueFormat string // want "field NonUniqueFormat has multiple definitions of marker k8s:format when only a single definition should exist"

	// +k8s:minLength:=10
	UniqueMinLength string

	// +k8s:minLength:=10
	// +k8s:minLength:=20
	NonUniqueMinLength string // want "field NonUniqueMinLength has multiple definitions of marker k8s:minLength when only a single definition should exist"

	// +k8s:maxLength:=100
	UniqueMaxLength string

	// +k8s:maxLength:=100
	// +k8s:maxLength:=200
	NonUniqueMaxLength string // want "field NonUniqueMaxLength has multiple definitions of marker k8s:maxLength when only a single definition should exist"

	// +k8s:minItems:=10
	UniqueMinItems []string

	// +k8s:minItems:=10
	// +k8s:minItems:=20
	NonUniqueMinItems []string // want "field NonUniqueMinItems has multiple definitions of marker k8s:minItems when only a single definition should exist"

	// +k8s:maxItems:=100
	UniqueMaxItems []string

	// +k8s:maxItems:=100
	// +k8s:maxItems:=200
	NonUniqueMaxItems []string // want "field NonUniqueMaxItems has multiple definitions of marker k8s:maxItems when only a single definition should exist"

	// +k8s:listType:=map
	UniqueListType []string

	// +k8s:listType:=map
	// +k8s:listType:=atomic
	NonUniqueListType []string // want "field NonUniqueListType has multiple definitions of marker k8s:listType when only a single definition should exist"
}
