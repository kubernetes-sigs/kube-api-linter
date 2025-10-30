package defaultconfigurations

import "k8s.io/apimachinery/pkg/types"

// ObjectReference contains enough information to let you inspect or modify the referred object.
// ---
// New uses of this type are discouraged because of difficulty describing its usage when embedded in APIs.
//  1. Ignored fields.  It includes many fields which are not generally honored.  For instance, ResourceVersion and FieldPath are both very rarely valid in actual usage.
//  2. Invalid usage help.  It is impossible to add specific help for individual usage.  In most embedded usages, there are particular
//     restrictions like, "must refer only to types A and B" or "UID not honored" or "name must be restricted".
//     Those cannot be well described when embedded.
//  3. Inconsistent validation.  Because the usages are different, the validation rules are different by usage, which makes it hard for users to predict what will happen.
//  4. The fields are both imprecise and overly precise.  Kind is not a precise mapping to a URL. This can produce ambiguity
//     during interpretation and require a REST mapping.  In most cases, the dependency is on the group,resource tuple
//     and the version of the actual struct is irrelevant.
//  5. We cannot easily change it.  Because this type is embedded in many locations, updates to this type
//     will affect numerous schemas.  Don't make new APIs embed an underspecified API type they do not control.
//
// Instead of using this type, create a locally provided and used type that is well-focused on your reference.
// For example, ServiceReferences for admission registration: https://github.com/kubernetes/api/blob/release-1.17/admissionregistration/v1/types.go#L533 .
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +structType=atomic
type ObjectReference struct {
	// Kind of the referent. // want "commentstart: godoc for field Kind should start with 'kind ...'"
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// +optional
	Kind string `json:"kind,omitempty" protobuf:"bytes,1,opt,name=kind"` // want "optionalfields: field Kind should be a pointer."
	// Namespace of the referent. // want "commentstart: godoc for field Namespace should start with 'namespace ...'"
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	// +optional
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,2,opt,name=namespace"` // want "optionalfields: field Namespace should be a pointer."
	// Name of the referent. // want "commentstart: godoc for field Name should start with 'name ...'"
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	// +optional
	Name string `json:"name,omitempty" protobuf:"bytes,3,opt,name=name"` // want "optionalfields: field Name should be a pointer."
	// UID of the referent. // want "commentstart: godoc for field UID should start with 'uid ...'"
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids
	// +optional
	UID types.UID `json:"uid,omitempty" protobuf:"bytes,4,opt,name=uid,casttype=k8s.io/apimachinery/pkg/types.UID"` // want "optionalfields: field UID should be a pointer."
	// API version of the referent. // want "commentstart: godoc for field APIVersion should start with 'apiVersion ...'"
	// +optional
	APIVersion string `json:"apiVersion,omitempty" protobuf:"bytes,5,opt,name=apiVersion"` // want "optionalfields: field APIVersion should be a pointer."
	// Specific resourceVersion to which this reference is made, if any. // want "commentstart: godoc for field ResourceVersion should start with 'resourceVersion ...'"
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
	// +optional
	ResourceVersion string `json:"resourceVersion,omitempty" protobuf:"bytes,6,opt,name=resourceVersion"` // want "optionalfields: field ResourceVersion should be a pointer."

	// If referring to a piece of an object instead of an entire object, this string // want "commentstart: godoc for field FieldPath should start with 'fieldPath ...'"
	// should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].
	// For example, if the object reference is to a container within a pod, this would take on a value like:
	// "spec.containers{name}" (where "name" refers to the name of the container that triggered
	// the event) or if no container name is specified "spec.containers[2]" (container with
	// index 2 in this pod). This syntax is chosen only to have some well-defined way of
	// referencing a part of an object.
	// TODO: this design is not final and this field is subject to change in the future.
	// +optional
	FieldPath string `json:"fieldPath,omitempty" protobuf:"bytes,7,opt,name=fieldPath"` // want "optionalfields: field FieldPath should be a pointer."
}

// LocalObjectReference contains enough information to let you locate the
// referenced object inside the same namespace.
// ---
// New uses of this type are discouraged because of difficulty describing its usage when embedded in APIs.
//  1. Invalid usage help.  It is impossible to add specific help for individual usage.  In most embedded usages, there are particular
//     restrictions like, "must refer only to types A and B" or "UID not honored" or "name must be restricted".
//     Those cannot be well described when embedded.
//  2. Inconsistent validation.  Because the usages are different, the validation rules are different by usage, which makes it hard for users to predict what will happen.
//  3. We cannot easily change it.  Because this type is embedded in many locations, updates to this type
//     will affect numerous schemas.  Don't make new APIs embed an underspecified API type they do not control.
//
// Instead of using this type, create a locally provided and used type that is well-focused on your reference.
// For example, ServiceReferences for admission registration: https://github.com/kubernetes/api/blob/release-1.17/admissionregistration/v1/types.go#L533 .
// +structType=atomic
type LocalObjectReference struct {
	// Name of the referent. // want "commentstart: godoc for field Name should start with 'name ...'"
	// This field is effectively required, but due to backwards compatibility is
	// allowed to be empty. Instances of this type with an empty value here are
	// almost certainly wrong.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	// +optional
	// +default=""
	// +kubebuilder:default=""
	// TODO: Drop `kubebuilder:default` when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"` // want "optionalfields: field Name should be a pointer."
}

// TypedLocalObjectReference contains enough information to let you locate the
// typed referenced object inside the same namespace.
// ---
// New uses of this type are discouraged because of difficulty describing its usage when embedded in APIs.
//  1. Invalid usage help.  It is impossible to add specific help for individual usage.  In most embedded usages, there are particular
//     restrictions like, "must refer only to types A and B" or "UID not honored" or "name must be restricted".
//     Those cannot be well described when embedded.
//  2. Inconsistent validation.  Because the usages are different, the validation rules are different by usage, which makes it hard for users to predict what will happen.
//  3. The fields are both imprecise and overly precise.  Kind is not a precise mapping to a URL. This can produce ambiguity
//     during interpretation and require a REST mapping.  In most cases, the dependency is on the group,resource tuple
//     and the version of the actual struct is irrelevant.
//  4. We cannot easily change it.  Because this type is embedded in many locations, updates to this type
//     will affect numerous schemas.  Don't make new APIs embed an underspecified API type they do not control.
//
// Instead of using this type, create a locally provided and used type that is well-focused on your reference.
// For example, ServiceReferences for admission registration: https://github.com/kubernetes/api/blob/release-1.17/admissionregistration/v1/types.go#L533 .
// +structType=atomic
type TypedLocalObjectReference struct {
	// APIGroup is the group for the resource being referenced. // want "commentstart: godoc for field APIGroup should start with 'apiGroup ...'"
	// If APIGroup is not specified, the specified Kind must be in the core API group.
	// For any other third-party types, APIGroup is required.
	// +optional
	APIGroup *string `json:"apiGroup" protobuf:"bytes,1,opt,name=apiGroup"` // want "optionalfields: field APIGroup should have the omitempty tag."
	// Kind is the type of resource being referenced // want "commentstart: godoc for field Kind should start with 'kind ...'"
	Kind string `json:"kind" protobuf:"bytes,2,opt,name=kind"` // want "optionalorrequired: field Kind must be marked as optional or required"
	// Name is the name of resource being referenced // want "commentstart: godoc for field Name should start with 'name ...'"
	Name string `json:"name" protobuf:"bytes,3,opt,name=name"` // want "optionalorrequired: field Name must be marked as optional or required"
}

// SecretReference represents a Secret Reference. It has enough information to retrieve secret
// in any namespace
// +structType=atomic
type SecretReference struct {
	// name is unique within a namespace to reference a secret resource.
	// +optional
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"` // want "optionalfields: field Name should be a pointer."
	// namespace defines the space within which the secret name must be unique.
	// +optional
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,2,opt,name=namespace"` // want "optionalfields: field Namespace should be a pointer."
}

// TypedObjectReference contains enough information to let you locate the typed referenced object
type TypedObjectReference struct {
	// APIGroup is the group for the resource being referenced. // want "commentstart: godoc for field APIGroup should start with 'apiGroup ...'"
	// If APIGroup is not specified, the specified Kind must be in the core API group.
	// For any other third-party types, APIGroup is required.
	// +optional
	APIGroup *string `json:"apiGroup" protobuf:"bytes,1,opt,name=apiGroup"` // want "optionalfields: field APIGroup should have the omitempty tag."
	// Kind is the type of resource being referenced // want "commentstart: godoc for field Kind should start with 'kind ...'"
	Kind string `json:"kind" protobuf:"bytes,2,opt,name=kind"` // want "optionalorrequired: field Kind must be marked as optional or required"
	// Name is the name of resource being referenced // want "commentstart: godoc for field Name should start with 'name ...'"
	Name string `json:"name" protobuf:"bytes,3,opt,name=name"` // want "optionalorrequired: field Name must be marked as optional or required"
	// Namespace is the namespace of resource being referenced // want "commentstart: godoc for field Namespace should start with 'namespace ...'"
	// Note that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.
	// (Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.
	// +featureGate=CrossNamespaceVolumeDataSource
	// +optional
	Namespace *string `json:"namespace,omitempty" protobuf:"bytes,4,opt,name=namespace"`
}
