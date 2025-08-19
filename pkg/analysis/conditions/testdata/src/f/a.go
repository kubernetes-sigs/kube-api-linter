package a

import (
	"go/ast"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ValidConditions struct {
	// conditions is an accurate representation of the desired state of a conditions object.
	// +listType=map
	// +listMapKey=type
	// +patchStrategy=merge
	// +patchMergeKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"` // want "Conditions field in ValidConditions has incorrect tags, should be: `json:\"conditions,omitempty\" protobuf:\"bytes,1,rep,name=conditions\"`" "Conditions field in ValidConditions has the following additional markers: patchStrategy=merge, patchMergeKey=type"

	// other fields
	OtherField string `json:"otherField,omitempty"`
}

// DoNothing is used to check that the analyser doesn't report on methods.
func (ValidConditions) DoNothing() {}

type ConditionsNotFirst struct {
	// other fields
	OtherField string `json:"otherField,omitempty"`

	// conditions is an accurate representation of the desired state of a conditions object.
	// +listType=map
	// +listMapKey=type
	// +patchStrategy=merge
	// +patchMergeKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"` // want "Conditions field in ConditionsNotFirst must be the first field in the struct" "Conditions field in ConditionsNotFirst has incorrect tags, should be: `json:\"conditions,omitempty\" protobuf:\"bytes,1,rep,name=conditions\"`" "Conditions field in ConditionsNotFirst has the following additional markers: patchStrategy=merge, patchMergeKey=type"
}

type ConditionsIncorrectType struct {
	// conditions has an incorrect type.
	Conditions map[string]metav1.Condition // want "Conditions field in ConditionsIncorrectType must be a slice of metav1.Condition"
}

type ConditionsIncorrectSliceElement struct {
	// conditions has an incorrect type.
	Conditions []string // want "Conditions field in ConditionsIncorrectSliceElement must be a slice of metav1.Condition"
}

type ConditionsIncorrectImportedSliceElement struct {
	// conditions has an incorrect type.
	Conditions []metav1.Time // want "Conditions field in ConditionsIncorrectImportedSliceElement must be a slice of metav1.Condition"
}

type ConditionsIncorrectImportedPackage struct {
	// conditions has an incorrect type.
	Conditions []ast.Node // want "Conditions field in ConditionsIncorrectImportedPackage must be a slice of metav1.Condition"
}

type MissingAllMarkers struct {
	// conditions is missing all markers.
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"` // want "Conditions field in MissingAllMarkers is missing the following markers: listType=map, listMapKey=type, optional" "Conditions field in MissingAllMarkers has incorrect tags, should be: `json:\"conditions,omitempty\" protobuf:\"bytes,1,rep,name=conditions\"`"
}

type MissingListMarkers struct {
	// conditions is missing list markers.
	// +patchStrategy=merge
	// +patchMergeKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"` // want "Conditions field in MissingListMarkers is missing the following markers: listType=map, listMapKey=type" "Conditions field in MissingListMarkers has incorrect tags, should be: `json:\"conditions,omitempty\" protobuf:\"bytes,1,rep,name=conditions\"`" "Conditions field in MissingListMarkers has the following additional markers: patchStrategy=merge, patchMergeKey=type"
}

type MissingPatchMarkers struct {
	// conditions is missing patch markers.
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"` // want "Conditions field in MissingPatchMarkers has incorrect tags, should be: `json:\"conditions,omitempty\" protobuf:\"bytes,1,rep,name=conditions\"`"
}

type MissingFieldTag struct {
	// conditions is missing the field tag.
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition // want "Conditions field in MissingFieldTag is missing tags, should be: `json:\"conditions,omitempty\" protobuf:\"bytes,1,rep,name=conditions\"`"
}

type IncorrectFieldTag struct {
	// conditions has an incorrect field tag.
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions"  patchMergeKey:"type" protobuf:"bytes,3,rep,name=conditions"` // want "Conditions field in IncorrectFieldTag has incorrect tags, should be: `json:\"conditions,omitempty\" protobuf:\"bytes,1,rep,name=conditions\"`"
}
