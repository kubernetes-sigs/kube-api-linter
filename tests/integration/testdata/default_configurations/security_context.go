package defaultconfigurations

// PodSecurityContext holds pod-level security attributes and common container settings.
// Some fields are also present in container.securityContext.  Field values of
// container.securityContext take precedence over field values of PodSecurityContext.
type PodSecurityContext struct {
	// The SELinux context to be applied to all containers. // want "commentstart: godoc for field SELinuxOptions should start with 'seLinuxOptions ...'"
	// If unspecified, the container runtime will allocate a random SELinux context for each
	// container.  May also be set in SecurityContext.  If set in
	// both SecurityContext and PodSecurityContext, the value specified in SecurityContext
	// takes precedence for that container.
	// Note that this field cannot be set when spec.os.name is windows.
	// +optional
	SELinuxOptions *SELinuxOptions `json:"seLinuxOptions,omitempty" protobuf:"bytes,1,opt,name=seLinuxOptions"`
	// The Windows specific settings applied to all containers. // want "commentstart: godoc for field WindowsOptions should start with 'windowsOptions ...'"
	// If unspecified, the options within a container's SecurityContext will be used.
	// If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	// Note that this field cannot be set when spec.os.name is linux.
	// +optional
	WindowsOptions *WindowsSecurityContextOptions `json:"windowsOptions,omitempty" protobuf:"bytes,8,opt,name=windowsOptions"`
	// The UID to run the entrypoint of the container process. // want "commentstart: godoc for field RunAsUser should start with 'runAsUser ...'"
	// Defaults to user specified in image metadata if unspecified.
	// May also be set in SecurityContext.  If set in both SecurityContext and
	// PodSecurityContext, the value specified in SecurityContext takes precedence
	// for that container.
	// Note that this field cannot be set when spec.os.name is windows.
	// +optional
	RunAsUser *int64 `json:"runAsUser,omitempty" protobuf:"varint,2,opt,name=runAsUser"`
	// The GID to run the entrypoint of the container process. // want "commentstart: godoc for field RunAsGroup should start with 'runAsGroup ...'"
	// Uses runtime default if unset.
	// May also be set in SecurityContext.  If set in both SecurityContext and
	// PodSecurityContext, the value specified in SecurityContext takes precedence
	// for that container.
	// Note that this field cannot be set when spec.os.name is windows.
	// +optional
	RunAsGroup *int64 `json:"runAsGroup,omitempty" protobuf:"varint,6,opt,name=runAsGroup"`
	// Indicates that the container must run as a non-root user. // want "commentstart: godoc for field RunAsNonRoot should start with 'runAsNonRoot ...'"
	// If true, the Kubelet will validate the image at runtime to ensure that it
	// does not run as UID 0 (root) and fail to start the container if it does.
	// If unset or false, no such validation will be performed.
	// May also be set in SecurityContext.  If set in both SecurityContext and
	// PodSecurityContext, the value specified in SecurityContext takes precedence.
	// +optional
	RunAsNonRoot *bool `json:"runAsNonRoot,omitempty" protobuf:"varint,3,opt,name=runAsNonRoot"`
	// A list of groups applied to the first process run in each container, in // want "commentstart: godoc for field SupplementalGroups should start with 'supplementalGroups ...'"
	// addition to the container's primary GID and fsGroup (if specified).  If
	// the SupplementalGroupsPolicy feature is enabled, the
	// supplementalGroupsPolicy field determines whether these are in addition
	// to or instead of any group memberships defined in the container image.
	// If unspecified, no additional groups are added, though group memberships
	// defined in the container image may still be used, depending on the
	// supplementalGroupsPolicy field.
	// Note that this field cannot be set when spec.os.name is windows.
	// +optional
	// +listType=atomic
	SupplementalGroups []int64 `json:"supplementalGroups,omitempty" protobuf:"varint,4,rep,name=supplementalGroups"`
	// Defines how supplemental groups of the first container processes are calculated. // want "commentstart: godoc for field SupplementalGroupsPolicy should start with 'supplementalGroupsPolicy ...'"
	// Valid values are "Merge" and "Strict". If not specified, "Merge" is used.
	// (Alpha) Using the field requires the SupplementalGroupsPolicy feature gate to be enabled
	// and the container runtime must implement support for this feature.
	// Note that this field cannot be set when spec.os.name is windows.
	// TODO: update the default value to "Merge" when spec.os.name is not windows in v1.34
	// +featureGate=SupplementalGroupsPolicy
	// +optional
	SupplementalGroupsPolicy *SupplementalGroupsPolicy `json:"supplementalGroupsPolicy,omitempty" protobuf:"bytes,12,opt,name=supplementalGroupsPolicy"`
	// A special supplemental group that applies to all containers in a pod. // want "commentstart: godoc for field FSGroup should start with 'fsGroup ...'"
	// Some volume types allow the Kubelet to change the ownership of that volume
	// to be owned by the pod:
	//
	// 1. The owning GID will be the FSGroup
	// 2. The setgid bit is set (new files created in the volume will be owned by FSGroup)
	// 3. The permission bits are OR'd with rw-rw----
	//
	// If unset, the Kubelet will not modify the ownership and permissions of any volume.
	// Note that this field cannot be set when spec.os.name is windows.
	// +optional
	FSGroup *int64 `json:"fsGroup,omitempty" protobuf:"varint,5,opt,name=fsGroup"`
	// Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported // want "commentstart: godoc for field Sysctls should start with 'sysctls ...'"
	// sysctls (by the container runtime) might fail to launch.
	// Note that this field cannot be set when spec.os.name is windows.
	// +optional
	// +listType=atomic
	Sysctls []Sysctl `json:"sysctls,omitempty" protobuf:"bytes,7,rep,name=sysctls"` // want "arrayofstruct: PodSecurityContext.Sysctls is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// fsGroupChangePolicy defines behavior of changing ownership and permission of the volume
	// before being exposed inside Pod. This field will only apply to
	// volume types which support fsGroup based ownership(and permissions).
	// It will have no effect on ephemeral volume types such as: secret, configmaps
	// and emptydir.
	// Valid values are "OnRootMismatch" and "Always". If not specified, "Always" is used.
	// Note that this field cannot be set when spec.os.name is windows.
	// +optional
	FSGroupChangePolicy *PodFSGroupChangePolicy `json:"fsGroupChangePolicy,omitempty" protobuf:"bytes,9,opt,name=fsGroupChangePolicy"`
	// The seccomp options to use by the containers in this pod. // want "commentstart: godoc for field SeccompProfile should start with 'seccompProfile ...'"
	// Note that this field cannot be set when spec.os.name is windows.
	// +optional
	SeccompProfile *SeccompProfile `json:"seccompProfile,omitempty" protobuf:"bytes,10,opt,name=seccompProfile"`
	// appArmorProfile is the AppArmor options to use by the containers in this pod.
	// Note that this field cannot be set when spec.os.name is windows.
	// +optional
	AppArmorProfile *AppArmorProfile `json:"appArmorProfile,omitempty" protobuf:"bytes,11,opt,name=appArmorProfile"`
	// seLinuxChangePolicy defines how the container's SELinux label is applied to all volumes used by the Pod.
	// It has no effect on nodes that do not support SELinux or to volumes does not support SELinux.
	// Valid values are "MountOption" and "Recursive".
	//
	// "Recursive" means relabeling of all files on all Pod volumes by the container runtime.
	// This may be slow for large volumes, but allows mixing privileged and unprivileged Pods sharing the same volume on the same node.
	//
	// "MountOption" mounts all eligible Pod volumes with `-o context` mount option.
	// This requires all Pods that share the same volume to use the same SELinux label.
	// It is not possible to share the same volume among privileged and unprivileged Pods.
	// Eligible volumes are in-tree FibreChannel and iSCSI volumes, and all CSI volumes
	// whose CSI driver announces SELinux support by setting spec.seLinuxMount: true in their
	// CSIDriver instance. Other volumes are always re-labelled recursively.
	// "MountOption" value is allowed only when SELinuxMount feature gate is enabled.
	//
	// If not specified and SELinuxMount feature gate is enabled, "MountOption" is used.
	// If not specified and SELinuxMount feature gate is disabled, "MountOption" is used for ReadWriteOncePod volumes
	// and "Recursive" for all other volumes.
	//
	// This field affects only Pods that have SELinux label set, either in PodSecurityContext or in SecurityContext of all containers.
	//
	// All Pods that use the same volume should use the same seLinuxChangePolicy, otherwise some pods can get stuck in ContainerCreating state.
	// Note that this field cannot be set when spec.os.name is windows.
	// +featureGate=SELinuxChangePolicy
	// +optional
	SELinuxChangePolicy *PodSELinuxChangePolicy `json:"seLinuxChangePolicy,omitempty" protobuf:"bytes,13,opt,name=seLinuxChangePolicy"`
}

// SupplementalGroupsPolicy defines how supplemental groups
// of the first container processes are calculated.
// +enum
type SupplementalGroupsPolicy string

const (
	// SupplementalGroupsPolicyMerge means that the container's provided
	// SupplementalGroups and FsGroup (specified in SecurityContext) will be
	// merged with the primary user's groups as defined in the container image
	// (in /etc/group).
	SupplementalGroupsPolicyMerge SupplementalGroupsPolicy = "Merge"
	// SupplementalGroupsPolicyStrict means that the container's provided
	// SupplementalGroups and FsGroup (specified in SecurityContext) will be
	// used instead of any groups defined in the container image.
	SupplementalGroupsPolicyStrict SupplementalGroupsPolicy = "Strict"
)

// Sysctl defines a kernel parameter to be set
type Sysctl struct {
	// Name of a property to set // want "commentstart: godoc for field Name should start with 'name ...'"
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"` // want "optionalorrequired: field Name must be marked as optional or required"
	// Value of a property to set // want "commentstart: godoc for field Value should start with 'value ...'"
	Value string `json:"value" protobuf:"bytes,2,opt,name=value"` // want "optionalorrequired: field Value must be marked as optional or required"
}

// PodFSGroupChangePolicy holds policies that will be used for applying fsGroup to a volume
// when volume is mounted.
// +enum
type PodFSGroupChangePolicy string

const (
	// FSGroupChangeOnRootMismatch indicates that volume's ownership and permissions will be changed
	// only when permission and ownership of root directory does not match with expected
	// permissions on the volume. This can help shorten the time it takes to change
	// ownership and permissions of a volume.
	FSGroupChangeOnRootMismatch PodFSGroupChangePolicy = "OnRootMismatch"
	// FSGroupChangeAlways indicates that volume's ownership and permissions
	// should always be changed whenever volume is mounted inside a Pod. This the default
	// behavior.
	FSGroupChangeAlways PodFSGroupChangePolicy = "Always"
)

// PodSELinuxChangePolicy defines how the container's SELinux label is applied to all volumes used by the Pod.
type PodSELinuxChangePolicy string

const (
	// Recursive relabeling of all Pod volumes by the container runtime.
	// This may be slow for large volumes, but allows mixing privileged and unprivileged Pods sharing the same volume on the same node.
	SELinuxChangePolicyRecursive PodSELinuxChangePolicy = "Recursive"
	// MountOption mounts all eligible Pod volumes with `-o context` mount option.
	// This requires all Pods that share the same volume to use the same SELinux label.
	// It is not possible to share the same volume among privileged and unprivileged Pods.
	// Eligible volumes are in-tree FibreChannel and iSCSI volumes, and all CSI volumes
	// whose CSI driver announces SELinux support by setting spec.seLinuxMount: true in their
	// CSIDriver instance. Other volumes are always re-labelled recursively.
	SELinuxChangePolicyMountOption PodSELinuxChangePolicy = "MountOption"
)
