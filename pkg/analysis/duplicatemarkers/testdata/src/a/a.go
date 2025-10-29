package a

// It must be ignored since it is not a type
// +kubebuilder:validation:Enum=foo;bar;baz
// +kubebuilder:validation:Enum=foo;bar;baz
var Variable string

// +kubebuilder:validation:Enum=foo;bar;baz
// +kubebuilder:validation:Enum=foo;bar;baz
type Enum string // want "Enum has duplicated markers kubebuilder:validation:Enum"

// +kubebuilder:validation:MaxLength=10
// +kubebuilder:validation:MaxLength=11
type MaxLength int

// +kubebuilder:validation:MaxLength=10
// +kubebuilder:validation:MaxLength=10
type DuplicatedMaxLength int // want "DuplicatedMaxLength has duplicated markers kubebuilder:validation:MaxLength=10"

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:object:root=true
type DuplicateMarkerSpec struct { // want "DuplicateMarkerSpec has duplicated markers kubebuilder:object:root"
	// +kubebuilder:validation:Required
	// should be ignored since it only has single marker
	Required string `json:"required"`

	// +listType=map
	// +listMapKey=primaryKey
	// +listMapKey=secondaryKey
	// +required
	// should be ignored since listMapKey is allowed to have different values
	Map Map `json:"map"`

	// +optional
	// +kubebuilder:validation:XValidation:rule="self >= 1 && self <= 3",message="must be 1 to 5"
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="replicas must be immutable"
	// should be ignored since XValidation is allowed to have different values
	Replicas *int `json:"replicas"`

	// +kubebuilder:validation:MaxLength=11
	// +kubebuilder:validation:MaxLength=10
	// should be ignored since MaxLength is allowed to have different values
	Maxlength int `json:"maxlength"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Required
	DuplicatedRequired string `json:"duplicatedRequired"` // want "DuplicatedRequired has duplicated markers kubebuilder:validation:Required"

	// +kubebuilder:validation:Enum=foo;bar;baz
	// +kubebuilder:validation:Enum=foo;bar;baz
	DuplicatedEnum string `json:"duplicatedEnum"` // want "DuplicatedEnum has duplicated markers kubebuilder:validation:Enum"

	// +kubebuilder:validation:MaxLength=10
	// +kubebuilder:validation:MaxLength=10
	DuplicatedMaxLength int `json:"duplicatedMaxLength"` // want "DuplicatedMaxLength has duplicated markers kubebuilder:validation:MaxLength=10"

	// +kubebuilder:validation:MaxLength=10
	DuplicatedMaxLengthIncludingTypeMarker MaxLength `json:"duplicatedMaxLengthIncludingTypeMarker"` // want "DuplicatedMaxLengthIncludingTypeMarker has duplicated markers kubebuilder:validation:MaxLength=10"

	// +listType=map
	// +listMapKey=primaryKey
	// +listMapKey=secondaryKey
	// +listType=map
	// +required
	DuplicatedListTypeMap Map `json:"duplicatedListTypeMap"` // want "DuplicatedListTypeMap has duplicated markers listType=map"

	// +optional
	// +kubebuilder:validation:XValidation:rule="self >= 1 && self <= 3",message="must be 1 to 5"
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="replicas must be immutable"
	// +kubebuilder:validation:XValidation:rule="self >= 1 && self <= 3",message="must be 1 to 5"
	DuplicatedReplicas *int `json:"duplicatedReplicas"` // want "DuplicatedReplicas has duplicated markers kubebuilder:validation:XValidation:rule=\"self >= 1 && self <= 3\",message=\"must be 1 to 5\""

	// +optional
	// +kubebuilder:validation:XValidation:rule="self >= 1 && self <= 3",message="must be 1 to 5"
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="replicas must be immutable"
	// +kubebuilder:validation:XValidation:message="must be 1 to 5",rule="self >= 1 && self <= 3"
	DuplicatedUnorderedValidationReplicas *int `json:"duplicatedUnorderedValidationReplicas"` // want "DuplicatedUnorderedValidationReplicas has duplicated markers kubebuilder:validation:XValidation:message=\"must be 1 to 5\",rule=\"self >= 1 && self <= 3\""

	StringFromAnotherFile StringFromAnotherFile `json:"stringFromAnotherFile"`

	// +kubebuilder:validation:MaxLength=10
	StringFromAnotherFileWithMaxLength StringFromAnotherFile `json:"stringFromAnotherFileWithMaxLength"` // want "StringFromAnotherFileWithMaxLength has duplicated markers kubebuilder:validation:MaxLength=10"
}

type Map struct {
	// +required
	PrimaryKey string `json:"primaryKey"`
	// +required
	SecondaryKey string `json:"secondaryKey"`
	// +required
	Value string `json:"value"`
	// +-------+-------+-------+
	// | zone1 | zone2 | zone3 |
	// +-------+-------+-------+
	// |  P P  |  P P  |   P   |
	// +-------+-------+-------+
	// should be ignored because we allow markdown tables as they are not actual markers
	MarkdownTable string `json:"markdownTable"`
}

// +enum
type UnsatisfiableConstraintAction string

type DuplicateMarkerBug struct {
	// MaxSkew describes the degree to which pods may be unevenly distributed.
	// When `whenUnsatisfiable=DoNotSchedule`, it is the maximum permitted difference
	// between the number of matching pods in the target topology and the global minimum.
	// The global minimum is the minimum number of matching pods in an eligible domain
	// or zero if the number of eligible domains is less than MinDomains.
	// For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same
	// labelSelector spread as 2/2/1:
	// In this case, the global minimum is 1.
	// +-------+-------+-------+
	// | zone1 | zone2 | zone3 |
	// +-------+-------+-------+
	// |  P P  |  P P  |   P   |
	// +-------+-------+-------+
	// - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2;
	// scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2)
	// violate MaxSkew(1).
	// - if MaxSkew is 2, incoming pod can be scheduled onto any zone.
	// When `whenUnsatisfiable=ScheduleAnyway`, it is used to give higher precedence
	// to topologies that satisfy it.
	// It's a required field. Default value is 1 and 0 is not allowed.
	MaxSkew int32 `json:"maxSkew" protobuf:"varint,1,opt,name=maxSkew"`

	// WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy
	// the spread constraint.
	// - DoNotSchedule (default) tells the scheduler not to schedule it.
	// - ScheduleAnyway tells the scheduler to schedule the pod in any location,
	//   but giving higher precedence to topologies that would help reduce the
	//   skew.
	// A constraint is considered "Unsatisfiable" for an incoming pod
	// if and only if every possible node assignment for that pod would violate
	// "MaxSkew" on some topology.
	// For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same
	// labelSelector spread as 3/1/1:
	// +-------+-------+-------+
	// | zone1 | zone2 | zone3 |
	// +-------+-------+-------+
	// | P P P |   P   |   P   |
	// +-------+-------+-------+
	// If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled
	// to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies
	// MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler
	// won't make it *more* imbalanced.
	// It's a required field.
	WhenUnsatisfiable UnsatisfiableConstraintAction `json:"whenUnsatisfiable" protobuf:"bytes,3,opt,name=whenUnsatisfiable,casttype=UnsatisfiableConstraintAction"`

	// For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the same
	// labelSelector spread as 2/2/2:
	// +-------+-------+-------+
	// | zone1 | zone2 | zone3 |
	// +-------+-------+-------+
	// |  P P  |  P P  |  P P  |
	// +-------+-------+-------+
	// The number of domains is less than 5(MinDomains), so "global minimum" is treated as 0.
	// In this situation, new pod with the same labelSelector cannot be scheduled,
	// because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones,
	// it will violate MaxSkew.
	// +optional
	MinDomains *int32 `json:"minDomains,omitempty" protobuf:"varint,5,opt,name=minDomains"`
}
