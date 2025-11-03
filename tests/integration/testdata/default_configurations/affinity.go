package defaultconfigurations

// Affinity is a group of affinity scheduling rules.
type Affinity struct {
	// Describes node affinity scheduling rules for the pod. // want "commentstart: godoc for field NodeAffinity should start with 'nodeAffinity ...'"
	// +optional
	NodeAffinity *NodeAffinity `json:"nodeAffinity,omitempty" protobuf:"bytes,1,opt,name=nodeAffinity"`
	// Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)). // want "commentstart: godoc for field PodAffinity should start with 'podAffinity ...'"
	// +optional
	PodAffinity *PodAffinity `json:"podAffinity,omitempty" protobuf:"bytes,2,opt,name=podAffinity"`
	// Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)). // want "commentstart: godoc for field PodAntiAffinity should start with 'podAntiAffinity ...'"
	// +optional
	PodAntiAffinity *PodAntiAffinity `json:"podAntiAffinity,omitempty" protobuf:"bytes,3,opt,name=podAntiAffinity"`
}

// Pod affinity is a group of inter pod affinity scheduling rules.
type PodAffinity struct {
	// NOT YET IMPLEMENTED. TODO: Uncomment field once it is implemented.
	// If the affinity requirements specified by this field are not met at
	// scheduling time, the pod will not be scheduled onto the node.
	// If the affinity requirements specified by this field cease to be met
	// at some point during pod execution (e.g. due to a pod label update), the
	// system will try to eventually evict the pod from its node.
	// When there are multiple elements, the lists of nodes corresponding to each
	// podAffinityTerm are intersected, i.e. all terms must be satisfied.
	// +optional
	// RequiredDuringSchedulingRequiredDuringExecution []PodAffinityTerm  `json:"requiredDuringSchedulingRequiredDuringExecution,omitempty"`

	// If the affinity requirements specified by this field are not met at // want "commentstart: godoc for field RequiredDuringSchedulingIgnoredDuringExecution should start with 'requiredDuringSchedulingIgnoredDuringExecution ...'"
	// scheduling time, the pod will not be scheduled onto the node.
	// If the affinity requirements specified by this field cease to be met
	// at some point during pod execution (e.g. due to a pod label update), the
	// system may or may not try to eventually evict the pod from its node.
	// When there are multiple elements, the lists of nodes corresponding to each
	// podAffinityTerm are intersected, i.e. all terms must be satisfied.
	// +optional
	// +listType=atomic
	RequiredDuringSchedulingIgnoredDuringExecution []PodAffinityTerm `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty" protobuf:"bytes,1,rep,name=requiredDuringSchedulingIgnoredDuringExecution"` // want "arrayofstruct: PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// The scheduler will prefer to schedule pods to nodes that satisfy // want "commentstart: godoc for field PreferredDuringSchedulingIgnoredDuringExecution should start with 'preferredDuringSchedulingIgnoredDuringExecution ...'"
	// the affinity expressions specified by this field, but it may choose
	// a node that violates one or more of the expressions. The node that is
	// most preferred is the one with the greatest sum of weights, i.e.
	// for each node that meets all of the scheduling requirements (resource
	// request, requiredDuringScheduling affinity expressions, etc.),
	// compute a sum by iterating through the elements of this field and adding
	// "weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the
	// node(s) with the highest sum are the most preferred.
	// +optional
	// +listType=atomic
	PreferredDuringSchedulingIgnoredDuringExecution []WeightedPodAffinityTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty" protobuf:"bytes,2,rep,name=preferredDuringSchedulingIgnoredDuringExecution"` // want "arrayofstruct: PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
}

// Pod anti affinity is a group of inter pod anti affinity scheduling rules.
type PodAntiAffinity struct {
	// NOT YET IMPLEMENTED. TODO: Uncomment field once it is implemented.
	// If the anti-affinity requirements specified by this field are not met at
	// scheduling time, the pod will not be scheduled onto the node.
	// If the anti-affinity requirements specified by this field cease to be met
	// at some point during pod execution (e.g. due to a pod label update), the
	// system will try to eventually evict the pod from its node.
	// When there are multiple elements, the lists of nodes corresponding to each
	// podAffinityTerm are intersected, i.e. all terms must be satisfied.
	// +optional
	// RequiredDuringSchedulingRequiredDuringExecution []PodAffinityTerm  `json:"requiredDuringSchedulingRequiredDuringExecution,omitempty"`

	// If the anti-affinity requirements specified by this field are not met at // want "commentstart: godoc for field RequiredDuringSchedulingIgnoredDuringExecution should start with 'requiredDuringSchedulingIgnoredDuringExecution ...'"
	// scheduling time, the pod will not be scheduled onto the node.
	// If the anti-affinity requirements specified by this field cease to be met
	// at some point during pod execution (e.g. due to a pod label update), the
	// system may or may not try to eventually evict the pod from its node.
	// When there are multiple elements, the lists of nodes corresponding to each
	// podAffinityTerm are intersected, i.e. all terms must be satisfied.
	// +optional
	// +listType=atomic
	RequiredDuringSchedulingIgnoredDuringExecution []PodAffinityTerm `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty" protobuf:"bytes,1,rep,name=requiredDuringSchedulingIgnoredDuringExecution"` // want "arrayofstruct: PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// The scheduler will prefer to schedule pods to nodes that satisfy // want "commentstart: godoc for field PreferredDuringSchedulingIgnoredDuringExecution should start with 'preferredDuringSchedulingIgnoredDuringExecution ...'"
	// the anti-affinity expressions specified by this field, but it may choose
	// a node that violates one or more of the expressions. The node that is
	// most preferred is the one with the greatest sum of weights, i.e.
	// for each node that meets all of the scheduling requirements (resource
	// request, requiredDuringScheduling anti-affinity expressions, etc.),
	// compute a sum by iterating through the elements of this field and subtracting
	// "weight" from the sum if the node has pods which matches the corresponding podAffinityTerm; the
	// node(s) with the highest sum are the most preferred.
	// +optional
	// +listType=atomic
	PreferredDuringSchedulingIgnoredDuringExecution []WeightedPodAffinityTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty" protobuf:"bytes,2,rep,name=preferredDuringSchedulingIgnoredDuringExecution"` // want "arrayofstruct: PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
}

// The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s)
type WeightedPodAffinityTerm struct {
	// weight associated with matching the corresponding podAffinityTerm,
	// in the range 1-100.
	Weight int32 `json:"weight" protobuf:"varint,1,opt,name=weight"` // want "optionalorrequired: field Weight must be marked as optional or required"
	// Required. A pod affinity term, associated with the corresponding weight. // want "commentstart: godoc for field PodAffinityTerm should start with 'podAffinityTerm ...'"
	PodAffinityTerm PodAffinityTerm `json:"podAffinityTerm" protobuf:"bytes,2,opt,name=podAffinityTerm"` // want "optionalorrequired: field PodAffinityTerm must be marked as optional or required"
}

// Defines a set of pods (namely those matching the labelSelector
// relative to the given namespace(s)) that this pod should be
// co-located (affinity) or not co-located (anti-affinity) with,
// where co-located is defined as running on a node whose value of
// the label with key <topologyKey> matches that of any node on which
// a pod of the set of pods is running
type PodAffinityTerm struct {
	// A label query over a set of resources, in this case pods.
	// If it's null, this PodAffinityTerm matches with no Pods.
	// +optional
	// LabelSelector *metav1.LabelSelector `json:"labelSelector,omitempty" protobuf:"bytes,1,opt,name=labelSelector"`

	// namespaces specifies a static list of namespace names that the term applies to.
	// The term is applied to the union of the namespaces listed in this field
	// and the ones selected by namespaceSelector.
	// null or empty namespaces list and null namespaceSelector means "this pod's namespace".
	// +optional
	// +listType=atomic
	Namespaces []string `json:"namespaces,omitempty" protobuf:"bytes,2,rep,name=namespaces"`
	// This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching // want "commentstart: godoc for field TopologyKey should start with 'topologyKey ...'"
	// the labelSelector in the specified namespaces, where co-located is defined as running on a node
	// whose value of the label with key topologyKey matches that of any node on which any of the
	// selected pods is running.
	// Empty topologyKey is not allowed.
	TopologyKey string `json:"topologyKey" protobuf:"bytes,3,opt,name=topologyKey"` // want "optionalorrequired: field TopologyKey must be marked as optional or required"
	// A label query over the set of namespaces that the term applies to.
	// The term is applied to the union of the namespaces selected by this field
	// and the ones listed in the namespaces field.
	// null selector and null or empty namespaces list means "this pod's namespace".
	// An empty selector ({}) matches all namespaces.
	// +optional
	// NamespaceSelector *metav1.LabelSelector `json:"namespaceSelector,omitempty" protobuf:"bytes,4,opt,name=namespaceSelector"`

	// MatchLabelKeys is a set of pod label keys to select which pods will // want "commentstart: godoc for field MatchLabelKeys should start with 'matchLabelKeys ...'"
	// be taken into consideration. The keys are used to lookup values from the
	// incoming pod labels, those key-value labels are merged with `labelSelector` as `key in (value)`
	// to select the group of existing pods which pods will be taken into consideration
	// for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming
	// pod labels will be ignored. The default value is empty.
	// The same key is forbidden to exist in both matchLabelKeys and labelSelector.
	// Also, matchLabelKeys cannot be set when labelSelector isn't set.
	//
	// +listType=atomic
	// +optional
	MatchLabelKeys []string `json:"matchLabelKeys,omitempty" protobuf:"bytes,5,opt,name=matchLabelKeys"`
	// MismatchLabelKeys is a set of pod label keys to select which pods will // want "commentstart: godoc for field MismatchLabelKeys should start with 'mismatchLabelKeys ...'"
	// be taken into consideration. The keys are used to lookup values from the
	// incoming pod labels, those key-value labels are merged with `labelSelector` as `key notin (value)`
	// to select the group of existing pods which pods will be taken into consideration
	// for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming
	// pod labels will be ignored. The default value is empty.
	// The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.
	// Also, mismatchLabelKeys cannot be set when labelSelector isn't set.
	//
	// +listType=atomic
	// +optional
	MismatchLabelKeys []string `json:"mismatchLabelKeys,omitempty" protobuf:"bytes,6,opt,name=mismatchLabelKeys"`
}

// Node affinity is a group of node affinity scheduling rules.
type NodeAffinity struct {
	// NOT YET IMPLEMENTED. TODO: Uncomment field once it is implemented.
	// If the affinity requirements specified by this field are not met at
	// scheduling time, the pod will not be scheduled onto the node.
	// If the affinity requirements specified by this field cease to be met
	// at some point during pod execution (e.g. due to an update), the system
	// will try to eventually evict the pod from its node.
	// +optional
	// RequiredDuringSchedulingRequiredDuringExecution *NodeSelector `json:"requiredDuringSchedulingRequiredDuringExecution,omitempty"`

	// If the affinity requirements specified by this field are not met at // want "commentstart: godoc for field RequiredDuringSchedulingIgnoredDuringExecution should start with 'requiredDuringSchedulingIgnoredDuringExecution ...'"
	// scheduling time, the pod will not be scheduled onto the node.
	// If the affinity requirements specified by this field cease to be met
	// at some point during pod execution (e.g. due to an update), the system
	// may or may not try to eventually evict the pod from its node.
	// +optional
	RequiredDuringSchedulingIgnoredDuringExecution *NodeSelector `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty" protobuf:"bytes,1,opt,name=requiredDuringSchedulingIgnoredDuringExecution"`
	// The scheduler will prefer to schedule pods to nodes that satisfy // want "commentstart: godoc for field PreferredDuringSchedulingIgnoredDuringExecution should start with 'preferredDuringSchedulingIgnoredDuringExecution ...'"
	// the affinity expressions specified by this field, but it may choose
	// a node that violates one or more of the expressions. The node that is
	// most preferred is the one with the greatest sum of weights, i.e.
	// for each node that meets all of the scheduling requirements (resource
	// request, requiredDuringScheduling affinity expressions, etc.),
	// compute a sum by iterating through the elements of this field and adding
	// "weight" to the sum if the node matches the corresponding matchExpressions; the
	// node(s) with the highest sum are the most preferred.
	// +optional
	// +listType=atomic
	PreferredDuringSchedulingIgnoredDuringExecution []PreferredSchedulingTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty" protobuf:"bytes,2,rep,name=preferredDuringSchedulingIgnoredDuringExecution"` // want "arrayofstruct: NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
}

// An empty preferred scheduling term matches all objects with implicit weight 0
// (i.e. it's a no-op). A null preferred scheduling term matches no objects (i.e. is also a no-op).
type PreferredSchedulingTerm struct {
	// Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100. // want "commentstart: godoc for field Weight should start with 'weight ...'"
	Weight int32 `json:"weight" protobuf:"varint,1,opt,name=weight"` // want "optionalorrequired: field Weight must be marked as optional or required"
	// A node selector term, associated with the corresponding weight. // want "commentstart: godoc for field Preference should start with 'preference ...'"
	Preference NodeSelectorTerm `json:"preference" protobuf:"bytes,2,opt,name=preference"` // want "optionalorrequired: field Preference must be marked as optional or required"
}

// The node this Taint is attached to has the "effect" on
// any pod that does not tolerate the Taint.
type Taint struct {
	// Required. The taint key to be applied to a node. // want "commentstart: godoc for field Key should start with 'key ...'"
	Key string `json:"key" protobuf:"bytes,1,opt,name=key"` // want "optionalorrequired: field Key must be marked as optional or required"
	// The taint value corresponding to the taint key. // want "commentstart: godoc for field Value should start with 'value ...'"
	// +optional
	Value string `json:"value,omitempty" protobuf:"bytes,2,opt,name=value"` // want "optionalfields: field Value should be a pointer."
	// Required. The effect of the taint on pods // want "commentstart: godoc for field Effect should start with 'effect ...'"
	// that do not tolerate the taint.
	// Valid effects are NoSchedule, PreferNoSchedule and NoExecute.
	Effect TaintEffect `json:"effect" protobuf:"bytes,3,opt,name=effect,casttype=TaintEffect"` // want "optionalorrequired: field Effect must be marked as optional or required"

	// TimeAdded represents the time at which the taint was added.
	// +optional
	// TimeAdded *metav1.Time `json:"timeAdded,omitempty" protobuf:"bytes,4,opt,name=timeAdded"`

}

// +enum
type TaintEffect string

const (
	// Do not allow new pods to schedule onto the node unless they tolerate the taint,
	// but allow all pods submitted to Kubelet without going through the scheduler
	// to start, and allow all already-running pods to continue running.
	// Enforced by the scheduler.
	TaintEffectNoSchedule TaintEffect = "NoSchedule"
	// Like TaintEffectNoSchedule, but the scheduler tries not to schedule
	// new pods onto the node, rather than prohibiting new pods from scheduling
	// onto the node entirely. Enforced by the scheduler.
	TaintEffectPreferNoSchedule TaintEffect = "PreferNoSchedule"
	// NOT YET IMPLEMENTED. TODO: Uncomment field once it is implemented.
	// Like TaintEffectNoSchedule, but additionally do not allow pods submitted to
	// Kubelet without going through the scheduler to start.
	// Enforced by Kubelet and the scheduler.
	// TaintEffectNoScheduleNoAdmit TaintEffect = "NoScheduleNoAdmit"

	// Evict any already-running pods that do not tolerate the taint.
	// Currently enforced by NodeController.
	TaintEffectNoExecute TaintEffect = "NoExecute"
)

// The pod this Toleration is attached to tolerates any taint that matches
// the triple <key,value,effect> using the matching operator <operator>.
type Toleration struct {
	// Key is the taint key that the toleration applies to. Empty means match all taint keys. // want "commentstart: godoc for field Key should start with 'key ...'"
	// If the key is empty, operator must be Exists; this combination means to match all values and all keys.
	// +optional
	Key string `json:"key,omitempty" protobuf:"bytes,1,opt,name=key"` // want "optionalfields: field Key should be a pointer."
	// Operator represents a key's relationship to the value. // want "commentstart: godoc for field Operator should start with 'operator ...'"
	// Valid operators are Exists and Equal. Defaults to Equal.
	// Exists is equivalent to wildcard for value, so that a pod can
	// tolerate all taints of a particular category.
	// +optional
	Operator TolerationOperator `json:"operator,omitempty" protobuf:"bytes,2,opt,name=operator,casttype=TolerationOperator"` // want "optionalfields: field Operator should be a pointer."
	// Value is the taint value the toleration matches to. // want "commentstart: godoc for field Value should start with 'value ...'"
	// If the operator is Exists, the value should be empty, otherwise just a regular string.
	// +optional
	Value string `json:"value,omitempty" protobuf:"bytes,3,opt,name=value"` // want "optionalfields: field Value should be a pointer."
	// Effect indicates the taint effect to match. Empty means match all taint effects. // want "commentstart: godoc for field Effect should start with 'effect ...'"
	// When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.
	// +optional
	Effect TaintEffect `json:"effect,omitempty" protobuf:"bytes,4,opt,name=effect,casttype=TaintEffect"` // want "optionalfields: field Effect should be a pointer."
	// TolerationSeconds represents the period of time the toleration (which must be // want "commentstart: godoc for field TolerationSeconds should start with 'tolerationSeconds ...'"
	// of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,
	// it is not set, which means tolerate the taint forever (do not evict). Zero and
	// negative values will be treated as 0 (evict immediately) by the system.
	// +optional
	TolerationSeconds *int64 `json:"tolerationSeconds,omitempty" protobuf:"varint,5,opt,name=tolerationSeconds"`
}

// A toleration operator is the set of operators that can be used in a toleration.
// +enum
type TolerationOperator string

const (
	TolerationOpExists TolerationOperator = "Exists"
	TolerationOpEqual  TolerationOperator = "Equal"
)

// A node selector represents the union of the results of one or more label queries
// over a set of nodes; that is, it represents the OR of the selectors represented
// by the node selector terms.
// +structType=atomic
type NodeSelector struct {
	// Required. A list of node selector terms. The terms are ORed. // want "commentstart: godoc for field NodeSelectorTerms should start with 'nodeSelectorTerms ...'"
	// +listType=atomic
	NodeSelectorTerms []NodeSelectorTerm `json:"nodeSelectorTerms" protobuf:"bytes,1,rep,name=nodeSelectorTerms"` // want "optionalorrequired: field NodeSelectorTerms must be marked as optional or required" "arrayofstruct: NodeSelector.NodeSelectorTerms is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
}

// A null or empty node selector term matches no objects. The requirements of
// them are ANDed.
// The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.
// +structType=atomic
type NodeSelectorTerm struct {
	// A list of node selector requirements by node's labels. // want "commentstart: godoc for field MatchExpressions should start with 'matchExpressions ...'"
	// +optional
	// +listType=atomic
	MatchExpressions []NodeSelectorRequirement `json:"matchExpressions,omitempty" protobuf:"bytes,1,rep,name=matchExpressions"` // want "arrayofstruct: NodeSelectorTerm.MatchExpressions is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// A list of node selector requirements by node's fields. // want "commentstart: godoc for field MatchFields should start with 'matchFields ...'"
	// +optional
	// +listType=atomic
	MatchFields []NodeSelectorRequirement `json:"matchFields,omitempty" protobuf:"bytes,2,rep,name=matchFields"` // want "arrayofstruct: NodeSelectorTerm.MatchFields is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
}

// A node selector requirement is a selector that contains values, a key, and an operator
// that relates the key and values.
type NodeSelectorRequirement struct {
	// The label key that the selector applies to. // want "commentstart: godoc for field Key should start with 'key ...'"
	Key string `json:"key" protobuf:"bytes,1,opt,name=key"` // want "optionalorrequired: field Key must be marked as optional or required"
	// Represents a key's relationship to a set of values. // want "commentstart: godoc for field Operator should start with 'operator ...'"
	// Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.
	Operator NodeSelectorOperator `json:"operator" protobuf:"bytes,2,opt,name=operator,casttype=NodeSelectorOperator"` // want "optionalorrequired: field Operator must be marked as optional or required"
	// An array of string values. If the operator is In or NotIn, // want "commentstart: godoc for field Values should start with 'values ...'"
	// the values array must be non-empty. If the operator is Exists or DoesNotExist,
	// the values array must be empty. If the operator is Gt or Lt, the values
	// array must have a single element, which will be interpreted as an integer.
	// This array is replaced during a strategic merge patch.
	// +optional
	// +listType=atomic
	Values []string `json:"values,omitempty" protobuf:"bytes,3,rep,name=values"`
}

// A node selector operator is the set of operators that can be used in
// a node selector requirement.
// +enum
type NodeSelectorOperator string

const (
	NodeSelectorOpIn           NodeSelectorOperator = "In"
	NodeSelectorOpNotIn        NodeSelectorOperator = "NotIn"
	NodeSelectorOpExists       NodeSelectorOperator = "Exists"
	NodeSelectorOpDoesNotExist NodeSelectorOperator = "DoesNotExist"
	NodeSelectorOpGt           NodeSelectorOperator = "Gt"
	NodeSelectorOpLt           NodeSelectorOperator = "Lt"
)

// TopologySpreadConstraint specifies how to spread matching pods among the given topology.
type TopologySpreadConstraint struct {
	// MaxSkew describes the degree to which pods may be unevenly distributed. // want "commentstart: godoc for field MaxSkew should start with 'maxSkew ...'"
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
	MaxSkew int32 `json:"maxSkew" protobuf:"varint,1,opt,name=maxSkew"` // want  "optionalorrequired: field MaxSkew must be marked as optional or required"
	// TopologyKey is the key of node labels. Nodes that have a label with this key // want "commentstart: godoc for field TopologyKey should start with 'topologyKey ...'"
	// and identical values are considered to be in the same topology.
	// We consider each <key, value> as a "bucket", and try to put balanced number
	// of pods into each bucket.
	// We define a domain as a particular instance of a topology.
	// Also, we define an eligible domain as a domain whose nodes meet the requirements of
	// nodeAffinityPolicy and nodeTaintsPolicy.
	// e.g. If TopologyKey is "kubernetes.io/hostname", each Node is a domain of that topology.
	// And, if TopologyKey is "topology.kubernetes.io/zone", each zone is a domain of that topology.
	// It's a required field.
	TopologyKey string `json:"topologyKey" protobuf:"bytes,2,opt,name=topologyKey"` // want "optionalorrequired: field TopologyKey must be marked as optional or required"

	// WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy // want "commentstart: godoc for field WhenUnsatisfiable should start with 'whenUnsatisfiable ...'"
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
	WhenUnsatisfiable UnsatisfiableConstraintAction `json:"whenUnsatisfiable" protobuf:"bytes,3,opt,name=whenUnsatisfiable,casttype=UnsatisfiableConstraintAction"` // want "optionalorrequired: field WhenUnsatisfiable must be marked as optional or required"

	// LabelSelector is used to find matching pods.
	// Pods that match this label selector are counted to determine the number of pods
	// in their corresponding topology domain.
	// +optional
	// LabelSelector *metav1.LabelSelector `json:"labelSelector,omitempty" protobuf:"bytes,4,opt,name=labelSelector"`

	// MinDomains indicates a minimum number of eligible domains. // want "commentstart: godoc for field MinDomains should start with 'minDomains ...'"
	// When the number of eligible domains with matching topology keys is less than minDomains,
	// Pod Topology Spread treats "global minimum" as 0, and then the calculation of Skew is performed.
	// And when the number of eligible domains with matching topology keys equals or greater than minDomains,
	// this value has no effect on scheduling.
	// As a result, when the number of eligible domains is less than minDomains,
	// scheduler won't schedule more than maxSkew Pods to those domains.
	// If value is nil, the constraint behaves as if MinDomains is equal to 1.
	// Valid values are integers greater than 0.
	// When value is not nil, WhenUnsatisfiable must be DoNotSchedule.
	//
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
	// NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelector // want "commentstart: godoc for field NodeAffinityPolicy should start with 'nodeAffinityPolicy ...'"
	// when calculating pod topology spread skew. Options are:
	// - Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations.
	// - Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations.
	//
	// If this value is nil, the behavior is equivalent to the Honor policy.
	// +optional
	NodeAffinityPolicy *NodeInclusionPolicy `json:"nodeAffinityPolicy,omitempty" protobuf:"bytes,6,opt,name=nodeAffinityPolicy"`
	// NodeTaintsPolicy indicates how we will treat node taints when calculating // want "commentstart: godoc for field NodeTaintsPolicy should start with 'nodeTaintsPolicy ...'"
	// pod topology spread skew. Options are:
	// - Honor: nodes without taints, along with tainted nodes for which the incoming pod
	// has a toleration, are included.
	// - Ignore: node taints are ignored. All nodes are included.
	//
	// If this value is nil, the behavior is equivalent to the Ignore policy.
	// +optional
	NodeTaintsPolicy *NodeInclusionPolicy `json:"nodeTaintsPolicy,omitempty" protobuf:"bytes,7,opt,name=nodeTaintsPolicy"`
	// MatchLabelKeys is a set of pod label keys to select the pods over which // want "commentstart: godoc for field MatchLabelKeys should start with 'matchLabelKeys ...'"
	// spreading will be calculated. The keys are used to lookup values from the
	// incoming pod labels, those key-value labels are ANDed with labelSelector
	// to select the group of existing pods over which spreading will be calculated
	// for the incoming pod. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.
	// MatchLabelKeys cannot be set when LabelSelector isn't set.
	// Keys that don't exist in the incoming pod labels will
	// be ignored. A null or empty list means only match against labelSelector.
	//
	// This is a beta field and requires the MatchLabelKeysInPodTopologySpread feature gate to be enabled (enabled by default).
	// +listType=atomic
	// +optional
	MatchLabelKeys []string `json:"matchLabelKeys,omitempty" protobuf:"bytes,8,opt,name=matchLabelKeys"`
}

// +enum
type UnsatisfiableConstraintAction string

const (
	// DoNotSchedule instructs the scheduler not to schedule the pod
	// when constraints are not satisfied.
	DoNotSchedule UnsatisfiableConstraintAction = "DoNotSchedule"
	// ScheduleAnyway instructs the scheduler to schedule the pod
	// even if constraints are not satisfied.
	ScheduleAnyway UnsatisfiableConstraintAction = "ScheduleAnyway"
)

// NodeInclusionPolicy defines the type of node inclusion policy
// +enum
type NodeInclusionPolicy string

const (
	// NodeInclusionPolicyIgnore means ignore this scheduling directive when calculating pod topology spread skew.
	NodeInclusionPolicyIgnore NodeInclusionPolicy = "Ignore"
	// NodeInclusionPolicyHonor means use this scheduling directive when calculating pod topology spread skew.
	NodeInclusionPolicyHonor NodeInclusionPolicy = "Honor"
)

// PreemptionPolicy describes a policy for if/when to preempt a pod.
// +enum
type PreemptionPolicy string

const (
	// PreemptLowerPriority means that pod can preempt other pods with lower priority.
	PreemptLowerPriority PreemptionPolicy = "PreemptLowerPriority"
	// PreemptNever means that pod never preempts other pods with lower priority.
	PreemptNever PreemptionPolicy = "Never"
)
