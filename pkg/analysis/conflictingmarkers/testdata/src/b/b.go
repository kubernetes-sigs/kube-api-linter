package b

type TestStruct struct {
	// No conflict: +listType=atomic and +k8s:listType=atomic match
	// +listType=atomic
	// +k8s:listType=atomic
	MatchingAtomicValues []string `json:"matchingAtomicValues"`

	// No conflict: +listType=set and +k8s:listType=set match
	// +listType=set
	// +k8s:listType=set
	MatchingSetValues []string `json:"matchingSetValues"`

	// No conflict: +listType=map and +k8s:listType=map match
	// +listType=map
	// +k8s:listType=map
	MatchingMapValues []string `json:"matchingMapValues"`

	// No conflict: only +listType present, no +k8s:listType
	// +listType=atomic
	OnlyListType []string `json:"onlyListType"`

	// No conflict: only +k8s:listType present, no +listType
	// +k8s:listType=atomic
	OnlyK8sListType []string `json:"onlyK8sListType"`

	// Conflict: +listType=atomic but +k8s:listType=set
	// +listType=atomic
	// +k8s:listType=set
	ConflictAtomicVsSet []string `json:"conflictAtomicVsSet"` // want "field TestStruct.ConflictAtomicVsSet has conflicting markers: listType_atomic_consistency: \\{\\[listType=atomic\\], \\[k8s:listType=set\\]\\}. when \\+k8s:listType is present, its value must match \\+listType"

	// Conflict: +listType=atomic but +k8s:listType=map
	// +listType=atomic
	// +k8s:listType=map
	ConflictAtomicVsMap []string `json:"conflictAtomicVsMap"` // want "field TestStruct.ConflictAtomicVsMap has conflicting markers: listType_atomic_consistency: \\{\\[listType=atomic\\], \\[k8s:listType=map\\]\\}. when \\+k8s:listType is present, its value must match \\+listType"

	// Conflict: +listType=set but +k8s:listType=atomic
	// +listType=set
	// +k8s:listType=atomic
	ConflictSetVsAtomic []string `json:"conflictSetVsAtomic"` // want "field TestStruct.ConflictSetVsAtomic has conflicting markers: listType_set_consistency: \\{\\[listType=set\\], \\[k8s:listType=atomic\\]\\}. when \\+k8s:listType is present, its value must match \\+listType"

	// Conflict: +listType=set but +k8s:listType=map
	// +listType=set
	// +k8s:listType=map
	ConflictSetVsMap []string `json:"conflictSetVsMap"` // want "field TestStruct.ConflictSetVsMap has conflicting markers: listType_set_consistency: \\{\\[listType=set\\], \\[k8s:listType=map\\]\\}. when \\+k8s:listType is present, its value must match \\+listType"

	// Conflict: +listType=map but +k8s:listType=atomic
	// +listType=map
	// +k8s:listType=atomic
	ConflictMapVsAtomic []string `json:"conflictMapVsAtomic"` // want "field TestStruct.ConflictMapVsAtomic has conflicting markers: listType_map_consistency: \\{\\[listType=map\\], \\[k8s:listType=atomic\\]\\}. when \\+k8s:listType is present, its value must match \\+listType"

	// Conflict: +listType=map but +k8s:listType=set
	// +listType=map
	// +k8s:listType=set
	ConflictMapVsSet []string `json:"conflictMapVsSet"` // want "field TestStruct.ConflictMapVsSet has conflicting markers: listType_map_consistency: \\{\\[listType=map\\], \\[k8s:listType=set\\]\\}. when \\+k8s:listType is present, its value must match \\+listType"

	// Conflict: using := separator, +listType:=atomic but +k8s:listType=set
	// +listType:=atomic
	// +k8s:listType=set
	ConflictColonEquals []string `json:"conflictColonEquals"` // want "field TestStruct.ConflictColonEquals has conflicting markers: listType_atomic_consistency: \\{\\[listType=atomic\\], \\[k8s:listType=set\\]\\}. when \\+k8s:listType is present, its value must match \\+listType"
}
