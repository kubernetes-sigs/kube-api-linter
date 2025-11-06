package defaultconfigurations

// VolumeMount describes a mounting of a Volume within a container.
type VolumeMount struct {
	// This must match the Name of a Volume. // want "commentstart: godoc for field VolumeMount.Name should start with 'name ...'"
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"` // want "optionalorrequired: field VolumeMount.Name must be marked as optional or required"
	// Mounted read-only if true, read-write otherwise (false or unspecified). // want "commentstart: godoc for field VolumeMount.ReadOnly should start with 'readOnly ...'"
	// Defaults to false.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,2,opt,name=readOnly"` // want "optionalfields: field ReadOnly should be a pointer."
	// RecursiveReadOnly specifies whether read-only mounts should be handled // want "commentstart: godoc for field VolumeMount.RecursiveReadOnly should start with 'recursiveReadOnly ...'"
	// recursively.
	//
	// If ReadOnly is false, this field has no meaning and must be unspecified.
	//
	// If ReadOnly is true, and this field is set to Disabled, the mount is not made
	// recursively read-only.  If this field is set to IfPossible, the mount is made
	// recursively read-only, if it is supported by the container runtime.  If this
	// field is set to Enabled, the mount is made recursively read-only if it is
	// supported by the container runtime, otherwise the pod will not be started and
	// an error will be generated to indicate the reason.
	//
	// If this field is set to IfPossible or Enabled, MountPropagation must be set to
	// None (or be unspecified, which defaults to None).
	//
	// If this field is not specified, it is treated as an equivalent of Disabled.
	// +optional
	RecursiveReadOnly *RecursiveReadOnlyMode `json:"recursiveReadOnly,omitempty" protobuf:"bytes,7,opt,name=recursiveReadOnly,casttype=RecursiveReadOnlyMode"`
	// Path within the container at which the volume should be mounted.  Must // want "commentstart: godoc for field VolumeMount.MountPath should start with 'mountPath ...'"
	// not contain ':'.
	MountPath string `json:"mountPath" protobuf:"bytes,3,opt,name=mountPath"` // want "optionalorrequired: field VolumeMount.MountPath must be marked as optional or required"
	// Path within the volume from which the container's volume should be mounted. // want "commentstart: godoc for field VolumeMount.SubPath should start with 'subPath ...'"
	// Defaults to "" (volume's root).
	// +optional
	SubPath string `json:"subPath,omitempty" protobuf:"bytes,4,opt,name=subPath"` // want "optionalfields: field SubPath should be a pointer."
	// mountPropagation determines how mounts are propagated from the host
	// to container and the other way around.
	// When not set, MountPropagationNone is used.
	// This field is beta in 1.10.
	// When RecursiveReadOnly is set to IfPossible or to Enabled, MountPropagation must be None or unspecified
	// (which defaults to None).
	// +optional
	MountPropagation *MountPropagationMode `json:"mountPropagation,omitempty" protobuf:"bytes,5,opt,name=mountPropagation,casttype=MountPropagationMode"`
	// Expanded path within the volume from which the container's volume should be mounted. // want "commentstart: godoc for field VolumeMount.SubPathExpr should start with 'subPathExpr ...'"
	// Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.
	// Defaults to "" (volume's root).
	// SubPathExpr and SubPath are mutually exclusive.
	// +optional
	SubPathExpr string `json:"subPathExpr,omitempty" protobuf:"bytes,6,opt,name=subPathExpr"` // want "optionalfields: field SubPathExpr should be a pointer."
}

// MountPropagationMode describes mount propagation.
// +enum
type MountPropagationMode string

const (
	// MountPropagationNone means that the volume in a container will
	// not receive new mounts from the host or other containers, and filesystems
	// mounted inside the container won't be propagated to the host or other
	// containers.
	// Note that this mode corresponds to "private" in Linux terminology.
	MountPropagationNone MountPropagationMode = "None"
	// MountPropagationHostToContainer means that the volume in a container will
	// receive new mounts from the host or other containers, but filesystems
	// mounted inside the container won't be propagated to the host or other
	// containers.
	// Note that this mode is recursively applied to all mounts in the volume
	// ("rslave" in Linux terminology).
	MountPropagationHostToContainer MountPropagationMode = "HostToContainer"
	// MountPropagationBidirectional means that the volume in a container will
	// receive new mounts from the host or other containers, and its own mounts
	// will be propagated from the container to the host or other containers.
	// Note that this mode is recursively applied to all mounts in the volume
	// ("rshared" in Linux terminology).
	MountPropagationBidirectional MountPropagationMode = "Bidirectional"
)

// RecursiveReadOnlyMode describes recursive-readonly mode.
type RecursiveReadOnlyMode string

const (
	// RecursiveReadOnlyDisabled disables recursive-readonly mode.
	RecursiveReadOnlyDisabled RecursiveReadOnlyMode = "Disabled"
	// RecursiveReadOnlyIfPossible enables recursive-readonly mode if possible.
	RecursiveReadOnlyIfPossible RecursiveReadOnlyMode = "IfPossible"
	// RecursiveReadOnlyEnabled enables recursive-readonly mode, or raise an error.
	RecursiveReadOnlyEnabled RecursiveReadOnlyMode = "Enabled"
)

// volumeDevice describes a mapping of a raw block device within a container.
type VolumeDevice struct {
	// name must match the name of a persistentVolumeClaim in the pod
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"` // want "optionalorrequired: field VolumeDevice.Name must be marked as optional or required"
	// devicePath is the path inside of the container that the device will be mapped to.
	DevicePath string `json:"devicePath" protobuf:"bytes,2,opt,name=devicePath"` // want "optionalorrequired: field VolumeDevice.DevicePath must be marked as optional or required"
}

// Volume represents a named volume in a pod that may be accessed by any container in the pod.
type Volume struct {
	// name of the volume.
	// Must be a DNS_LABEL and unique within the pod.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"` // want "optionalorrequired: field Volume.Name must be marked as optional or required"
	// volumeSource represents the location and type of the mounted volume.
	// If not specified, the Volume is implied to be an EmptyDir.
	// This implied behavior is deprecated and will be removed in a future version.
	VolumeSource `json:",inline" protobuf:"bytes,2,opt,name=volumeSource"`
}

// Represents the source of a volume to mount.
// Only one of its members may be specified.
type VolumeSource struct {
	// hostPath represents a pre-existing file or directory on the host
	// machine that is directly exposed to the container. This is generally
	// used for system agents or other privileged things that are allowed
	// to see the host machine. Most containers will NOT need this.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	// ---
	// TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not
	// mount host directories as read/write.
	// +optional
	HostPath *HostPathVolumeSource `json:"hostPath,omitempty" protobuf:"bytes,1,opt,name=hostPath"`
	// emptyDir represents a temporary directory that shares a pod's lifetime.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
	// +optional
	EmptyDir *EmptyDirVolumeSource `json:"emptyDir,omitempty" protobuf:"bytes,2,opt,name=emptyDir"`
	// gcePersistentDisk represents a GCE Disk resource that is attached to a
	// kubelet's host machine and then exposed to the pod.
	// Deprecated: GCEPersistentDisk is deprecated. All operations for the in-tree
	// gcePersistentDisk type are redirected to the pd.csi.storage.gke.io CSI driver.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	// +optional
	GCEPersistentDisk *GCEPersistentDiskVolumeSource `json:"gcePersistentDisk,omitempty" protobuf:"bytes,3,opt,name=gcePersistentDisk"`
	// awsElasticBlockStore represents an AWS Disk resource that is attached to a
	// kubelet's host machine and then exposed to the pod.
	// Deprecated: AWSElasticBlockStore is deprecated. All operations for the in-tree
	// awsElasticBlockStore type are redirected to the ebs.csi.aws.com CSI driver.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	// +optional
	AWSElasticBlockStore *AWSElasticBlockStoreVolumeSource `json:"awsElasticBlockStore,omitempty" protobuf:"bytes,4,opt,name=awsElasticBlockStore"`
	// gitRepo represents a git repository at a particular revision.
	// Deprecated: GitRepo is deprecated. To provision a container with a git repo, mount an
	// EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir
	// into the Pod's container.
	// +optional
	GitRepo *GitRepoVolumeSource `json:"gitRepo,omitempty" protobuf:"bytes,5,opt,name=gitRepo"`
	// secret represents a secret that should populate this volume.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#secret
	// +optional
	Secret *SecretVolumeSource `json:"secret,omitempty" protobuf:"bytes,6,opt,name=secret"`
	// nfs represents an NFS mount on the host that shares a pod's lifetime
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	// +optional
	NFS *NFSVolumeSource `json:"nfs,omitempty" protobuf:"bytes,7,opt,name=nfs"`
	// iscsi represents an ISCSI Disk resource that is attached to a
	// kubelet's host machine and then exposed to the pod.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes/#iscsi
	// +optional
	ISCSI *ISCSIVolumeSource `json:"iscsi,omitempty" protobuf:"bytes,8,opt,name=iscsi"`
	// glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime.
	// Deprecated: Glusterfs is deprecated and the in-tree glusterfs type is no longer supported.
	// +optional
	Glusterfs *GlusterfsVolumeSource `json:"glusterfs,omitempty" protobuf:"bytes,9,opt,name=glusterfs"`
	// persistentVolumeClaimVolumeSource represents a reference to a // want "commentstart: godoc for field VolumeSource.PersistentVolumeClaim should start with 'persistentVolumeClaim ...'"
	// PersistentVolumeClaim in the same namespace.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	// +optional
	PersistentVolumeClaim *PersistentVolumeClaimVolumeSource `json:"persistentVolumeClaim,omitempty" protobuf:"bytes,10,opt,name=persistentVolumeClaim"`
	// rbd represents a Rados Block Device mount on the host that shares a pod's lifetime.
	// Deprecated: RBD is deprecated and the in-tree rbd type is no longer supported.
	// +optional
	RBD *RBDVolumeSource `json:"rbd,omitempty" protobuf:"bytes,11,opt,name=rbd"`
	// flexVolume represents a generic volume resource that is
	// provisioned/attached using an exec based plugin.
	// Deprecated: FlexVolume is deprecated. Consider using a CSIDriver instead.
	// +optional
	FlexVolume *FlexVolumeSource `json:"flexVolume,omitempty" protobuf:"bytes,12,opt,name=flexVolume"`
	// cinder represents a cinder volume attached and mounted on kubelets host machine.
	// Deprecated: Cinder is deprecated. All operations for the in-tree cinder type
	// are redirected to the cinder.csi.openstack.org CSI driver.
	// More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	// +optional
	Cinder *CinderVolumeSource `json:"cinder,omitempty" protobuf:"bytes,13,opt,name=cinder"`
	// cephFS represents a Ceph FS mount on the host that shares a pod's lifetime. // want "commentstart: godoc for field VolumeSource.CephFS should start with 'cephfs ...'"
	// Deprecated: CephFS is deprecated and the in-tree cephfs type is no longer supported.
	// +optional
	CephFS *CephFSVolumeSource `json:"cephfs,omitempty" protobuf:"bytes,14,opt,name=cephfs"`
	// flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running.
	// Deprecated: Flocker is deprecated and the in-tree flocker type is no longer supported.
	// +optional
	Flocker *FlockerVolumeSource `json:"flocker,omitempty" protobuf:"bytes,15,opt,name=flocker"`
	// downwardAPI represents downward API about the pod that should populate this volume
	// +optional
	DownwardAPI *DownwardAPIVolumeSource `json:"downwardAPI,omitempty" protobuf:"bytes,16,opt,name=downwardAPI"`
	// fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.
	// +optional
	FC *FCVolumeSource `json:"fc,omitempty" protobuf:"bytes,17,opt,name=fc"`
	// azureFile represents an Azure File Service mount on the host and bind mount to the pod.
	// Deprecated: AzureFile is deprecated. All operations for the in-tree azureFile type
	// are redirected to the file.csi.azure.com CSI driver.
	// +optional
	AzureFile *AzureFileVolumeSource `json:"azureFile,omitempty" protobuf:"bytes,18,opt,name=azureFile"`
	// configMap represents a configMap that should populate this volume
	// +optional
	ConfigMap *ConfigMapVolumeSource `json:"configMap,omitempty" protobuf:"bytes,19,opt,name=configMap"`
	// vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine.
	// Deprecated: VsphereVolume is deprecated. All operations for the in-tree vsphereVolume type
	// are redirected to the csi.vsphere.vmware.com CSI driver.
	// +optional
	VsphereVolume *VsphereVirtualDiskVolumeSource `json:"vsphereVolume,omitempty" protobuf:"bytes,20,opt,name=vsphereVolume"`
	// quobyte represents a Quobyte mount on the host that shares a pod's lifetime.
	// Deprecated: Quobyte is deprecated and the in-tree quobyte type is no longer supported.
	// +optional
	Quobyte *QuobyteVolumeSource `json:"quobyte,omitempty" protobuf:"bytes,21,opt,name=quobyte"`
	// azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
	// Deprecated: AzureDisk is deprecated. All operations for the in-tree azureDisk type
	// are redirected to the disk.csi.azure.com CSI driver.
	// +optional
	AzureDisk *AzureDiskVolumeSource `json:"azureDisk,omitempty" protobuf:"bytes,22,opt,name=azureDisk"`
	// photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine.
	// Deprecated: PhotonPersistentDisk is deprecated and the in-tree photonPersistentDisk type is no longer supported.
	PhotonPersistentDisk *PhotonPersistentDiskVolumeSource `json:"photonPersistentDisk,omitempty" protobuf:"bytes,23,opt,name=photonPersistentDisk"` // want "optionalorrequired: field VolumeSource.PhotonPersistentDisk must be marked as optional or required"
	// projected items for all in one resources secrets, configmaps, and downward API
	Projected *ProjectedVolumeSource `json:"projected,omitempty" protobuf:"bytes,26,opt,name=projected"` // want "optionalorrequired: field VolumeSource.Projected must be marked as optional or required"
	// portworxVolume represents a portworx volume attached and mounted on kubelets host machine.
	// Deprecated: PortworxVolume is deprecated. All operations for the in-tree portworxVolume type
	// are redirected to the pxd.portworx.com CSI driver when the CSIMigrationPortworx feature-gate
	// is on.
	// +optional
	PortworxVolume *PortworxVolumeSource `json:"portworxVolume,omitempty" protobuf:"bytes,24,opt,name=portworxVolume"`
	// scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.
	// Deprecated: ScaleIO is deprecated and the in-tree scaleIO type is no longer supported.
	// +optional
	ScaleIO *ScaleIOVolumeSource `json:"scaleIO,omitempty" protobuf:"bytes,25,opt,name=scaleIO"`
	// storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes. // want "commentstart: godoc for field VolumeSource.StorageOS should start with 'storageos ...'"
	// Deprecated: StorageOS is deprecated and the in-tree storageos type is no longer supported.
	// +optional
	StorageOS *StorageOSVolumeSource `json:"storageos,omitempty" protobuf:"bytes,27,opt,name=storageos"`
	// csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers.
	// +optional
	CSI *CSIVolumeSource `json:"csi,omitempty" protobuf:"bytes,28,opt,name=csi"`
	// ephemeral represents a volume that is handled by a cluster storage driver.
	// The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts,
	// and deleted when the pod is removed.
	//
	// Use this if:
	// a) the volume is only needed while the pod runs,
	// b) features of normal volumes like restoring from snapshot or capacity
	//    tracking are needed,
	// c) the storage driver is specified through a storage class, and
	// d) the storage driver supports dynamic volume provisioning through
	//    a PersistentVolumeClaim (see EphemeralVolumeSource for more
	//    information on the connection between this volume type
	//    and PersistentVolumeClaim).
	//
	// Use PersistentVolumeClaim or one of the vendor-specific
	// APIs for volumes that persist for longer than the lifecycle
	// of an individual pod.
	//
	// Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to
	// be used that way - see the documentation of the driver for
	// more information.
	//
	// A pod can use both types of ephemeral volumes and
	// persistent volumes at the same time.
	//
	// +optional
	Ephemeral *EphemeralVolumeSource `json:"ephemeral,omitempty" protobuf:"bytes,29,opt,name=ephemeral"`
	// image represents an OCI object (a container image or artifact) pulled and mounted on the kubelet's host machine.
	// The volume is resolved at pod startup depending on which PullPolicy value is provided:
	//
	// - Always: the kubelet always attempts to pull the reference. Container creation will fail If the pull fails.
	// - Never: the kubelet never pulls the reference and only uses a local image or artifact. Container creation will fail if the reference isn't present.
	// - IfNotPresent: the kubelet pulls if the reference isn't already present on disk. Container creation will fail if the reference isn't present and the pull fails.
	//
	// The volume gets re-resolved if the pod gets deleted and recreated, which means that new remote content will become available on pod recreation.
	// A failure to resolve or pull the image during pod startup will block containers from starting and may add significant latency. Failures will be retried using normal volume backoff and will be reported on the pod reason and message.
	// The types of objects that may be mounted by this volume are defined by the container runtime implementation on a host machine and at minimum must include all valid types supported by the container image field.
	// The OCI object gets mounted in a single directory (spec.containers[*].volumeMounts.mountPath) by merging the manifest layers in the same way as for container images.
	// The volume will be mounted read-only (ro) and non-executable files (noexec).
	// Sub path mounts for containers are not supported (spec.containers[*].volumeMounts.subpath) before 1.33.
	// The field spec.securityContext.fsGroupChangePolicy has no effect on this volume type.
	// +featureGate=ImageVolume
	// +optional
	Image *ImageVolumeSource `json:"image,omitempty" protobuf:"bytes,30,opt,name=image"`
}

// Represents a projected volume source
type ProjectedVolumeSource struct {
	// sources is the list of volume projections. Each entry in this list
	// handles one source.
	// +optional
	// +listType=atomic
	Sources []VolumeProjection `json:"sources" protobuf:"bytes,1,rep,name=sources"` // want "optionalfields: field Sources should have the omitempty tag." "arrayofstruct: ProjectedVolumeSource.Sources is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// defaultMode are the mode bits used to set permissions on created files by default.
	// Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.
	// YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.
	// Directories within the path are not affected by this setting.
	// This might be in conflict with other options that affect the file
	// mode, like fsGroup, and the result can be other mode bits set.
	// +optional
	DefaultMode *int32 `json:"defaultMode,omitempty" protobuf:"varint,2,opt,name=defaultMode"`
}

// ImageVolumeSource represents a image volume resource.
type ImageVolumeSource struct {
	// Required: Image or artifact reference to be used. // want "commentstart: godoc for field ImageVolumeSource.Reference should start with 'reference ...'"
	// Behaves in the same way as pod.spec.containers[*].image.
	// Pull secrets will be assembled in the same way as for the container image by looking up node credentials, SA image pull secrets, and pod spec image pull secrets.
	// More info: https://kubernetes.io/docs/concepts/containers/images
	// This field is optional to allow higher level config management to default or override
	// container images in workload controllers like Deployments and StatefulSets.
	// +optional
	Reference string `json:"reference,omitempty" protobuf:"bytes,1,opt,name=reference"` // want "optionalfields: field Reference should be a pointer." "noreferences: naming convention \"reference-to-ref\": field ImageVolumeSource.Reference: field names should use 'Ref' instead of 'Reference'"

	// Policy for pulling OCI objects. Possible values are: // want "commentstart: godoc for field ImageVolumeSource.PullPolicy should start with 'pullPolicy ...'"
	// Always: the kubelet always attempts to pull the reference. Container creation will fail If the pull fails.
	// Never: the kubelet never pulls the reference and only uses a local image or artifact. Container creation will fail if the reference isn't present.
	// IfNotPresent: the kubelet pulls if the reference isn't already present on disk. Container creation will fail if the reference isn't present and the pull fails.
	// Defaults to Always if :latest tag is specified, or IfNotPresent otherwise.
	// +optional
	PullPolicy PullPolicy `json:"pullPolicy,omitempty" protobuf:"bytes,2,opt,name=pullPolicy,casttype=PullPolicy"` // want "optionalfields: field PullPolicy should be a pointer."
}

// PersistentVolumeClaimVolumeSource references the user's PVC in the same namespace.
// This volume finds the bound PV and mounts that volume for the pod. A
// PersistentVolumeClaimVolumeSource is, essentially, a wrapper around another
// type of volume that is owned by someone else (the system).
type PersistentVolumeClaimVolumeSource struct {
	// claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	ClaimName string `json:"claimName" protobuf:"bytes,1,opt,name=claimName"` // want "optionalorrequired: field PersistentVolumeClaimVolumeSource.ClaimName must be marked as optional or required"
	// readOnly Will force the ReadOnly setting in VolumeMounts.
	// Default false.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,2,opt,name=readOnly"` // want "optionalfields: field ReadOnly should be a pointer."
}

// PersistentVolumeSource is similar to VolumeSource but meant for the
// administrator who creates PVs. Exactly one of its members must be set.
type PersistentVolumeSource struct {
	// gcePersistentDisk represents a GCE Disk resource that is attached to a
	// kubelet's host machine and then exposed to the pod. Provisioned by an admin.
	// Deprecated: GCEPersistentDisk is deprecated. All operations for the in-tree
	// gcePersistentDisk type are redirected to the pd.csi.storage.gke.io CSI driver.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	// +optional
	GCEPersistentDisk *GCEPersistentDiskVolumeSource `json:"gcePersistentDisk,omitempty" protobuf:"bytes,1,opt,name=gcePersistentDisk"`
	// awsElasticBlockStore represents an AWS Disk resource that is attached to a
	// kubelet's host machine and then exposed to the pod.
	// Deprecated: AWSElasticBlockStore is deprecated. All operations for the in-tree
	// awsElasticBlockStore type are redirected to the ebs.csi.aws.com CSI driver.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	// +optional
	AWSElasticBlockStore *AWSElasticBlockStoreVolumeSource `json:"awsElasticBlockStore,omitempty" protobuf:"bytes,2,opt,name=awsElasticBlockStore"`
	// hostPath represents a directory on the host.
	// Provisioned by a developer or tester.
	// This is useful for single-node development and testing only!
	// On-host storage is not supported in any way and WILL NOT WORK in a multi-node cluster.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	// +optional
	HostPath *HostPathVolumeSource `json:"hostPath,omitempty" protobuf:"bytes,3,opt,name=hostPath"`
	// glusterfs represents a Glusterfs volume that is attached to a host and
	// exposed to the pod. Provisioned by an admin.
	// Deprecated: Glusterfs is deprecated and the in-tree glusterfs type is no longer supported.
	// More info: https://examples.k8s.io/volumes/glusterfs/README.md
	// +optional
	Glusterfs *GlusterfsPersistentVolumeSource `json:"glusterfs,omitempty" protobuf:"bytes,4,opt,name=glusterfs"`
	// nfs represents an NFS mount on the host. Provisioned by an admin.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	// +optional
	NFS *NFSVolumeSource `json:"nfs,omitempty" protobuf:"bytes,5,opt,name=nfs"`
	// rbd represents a Rados Block Device mount on the host that shares a pod's lifetime.
	// Deprecated: RBD is deprecated and the in-tree rbd type is no longer supported.
	// More info: https://examples.k8s.io/volumes/rbd/README.md
	// +optional
	RBD *RBDPersistentVolumeSource `json:"rbd,omitempty" protobuf:"bytes,6,opt,name=rbd"`
	// iscsi represents an ISCSI Disk resource that is attached to a
	// kubelet's host machine and then exposed to the pod. Provisioned by an admin.
	// +optional
	ISCSI *ISCSIPersistentVolumeSource `json:"iscsi,omitempty" protobuf:"bytes,7,opt,name=iscsi"`
	// cinder represents a cinder volume attached and mounted on kubelets host machine.
	// Deprecated: Cinder is deprecated. All operations for the in-tree cinder type
	// are redirected to the cinder.csi.openstack.org CSI driver.
	// More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	// +optional
	Cinder *CinderPersistentVolumeSource `json:"cinder,omitempty" protobuf:"bytes,8,opt,name=cinder"`
	// cephFS represents a Ceph FS mount on the host that shares a pod's lifetime. // want "commentstart: godoc for field PersistentVolumeSource.CephFS should start with 'cephfs ...'"
	// Deprecated: CephFS is deprecated and the in-tree cephfs type is no longer supported.
	// +optional
	CephFS *CephFSPersistentVolumeSource `json:"cephfs,omitempty" protobuf:"bytes,9,opt,name=cephfs"`
	// fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.
	// +optional
	FC *FCVolumeSource `json:"fc,omitempty" protobuf:"bytes,10,opt,name=fc"`
	// flocker represents a Flocker volume attached to a kubelet's host machine and exposed to the pod for its usage. This depends on the Flocker control service being running.
	// Deprecated: Flocker is deprecated and the in-tree flocker type is no longer supported.
	// +optional
	Flocker *FlockerVolumeSource `json:"flocker,omitempty" protobuf:"bytes,11,opt,name=flocker"`
	// flexVolume represents a generic volume resource that is
	// provisioned/attached using an exec based plugin.
	// Deprecated: FlexVolume is deprecated. Consider using a CSIDriver instead.
	// +optional
	FlexVolume *FlexPersistentVolumeSource `json:"flexVolume,omitempty" protobuf:"bytes,12,opt,name=flexVolume"`
	// azureFile represents an Azure File Service mount on the host and bind mount to the pod.
	// Deprecated: AzureFile is deprecated. All operations for the in-tree azureFile type
	// are redirected to the file.csi.azure.com CSI driver.
	// +optional
	AzureFile *AzureFilePersistentVolumeSource `json:"azureFile,omitempty" protobuf:"bytes,13,opt,name=azureFile"`
	// vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine.
	// Deprecated: VsphereVolume is deprecated. All operations for the in-tree vsphereVolume type
	// are redirected to the csi.vsphere.vmware.com CSI driver.
	// +optional
	VsphereVolume *VsphereVirtualDiskVolumeSource `json:"vsphereVolume,omitempty" protobuf:"bytes,14,opt,name=vsphereVolume"`
	// quobyte represents a Quobyte mount on the host that shares a pod's lifetime.
	// Deprecated: Quobyte is deprecated and the in-tree quobyte type is no longer supported.
	// +optional
	Quobyte *QuobyteVolumeSource `json:"quobyte,omitempty" protobuf:"bytes,15,opt,name=quobyte"`
	// azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
	// Deprecated: AzureDisk is deprecated. All operations for the in-tree azureDisk type
	// are redirected to the disk.csi.azure.com CSI driver.
	// +optional
	AzureDisk *AzureDiskVolumeSource `json:"azureDisk,omitempty" protobuf:"bytes,16,opt,name=azureDisk"`
	// photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine.
	// Deprecated: PhotonPersistentDisk is deprecated and the in-tree photonPersistentDisk type is no longer supported.
	PhotonPersistentDisk *PhotonPersistentDiskVolumeSource `json:"photonPersistentDisk,omitempty" protobuf:"bytes,17,opt,name=photonPersistentDisk"` // want "optionalorrequired: field PersistentVolumeSource.PhotonPersistentDisk must be marked as optional or required"
	// portworxVolume represents a portworx volume attached and mounted on kubelets host machine.
	// Deprecated: PortworxVolume is deprecated. All operations for the in-tree portworxVolume type
	// are redirected to the pxd.portworx.com CSI driver when the CSIMigrationPortworx feature-gate
	// is on.
	// +optional
	PortworxVolume *PortworxVolumeSource `json:"portworxVolume,omitempty" protobuf:"bytes,18,opt,name=portworxVolume"`
	// scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.
	// Deprecated: ScaleIO is deprecated and the in-tree scaleIO type is no longer supported.
	// +optional
	ScaleIO *ScaleIOPersistentVolumeSource `json:"scaleIO,omitempty" protobuf:"bytes,19,opt,name=scaleIO"`
	// local represents directly-attached storage with node affinity
	// +optional
	Local *LocalVolumeSource `json:"local,omitempty" protobuf:"bytes,20,opt,name=local"`
	// storageOS represents a StorageOS volume that is attached to the kubelet's host machine and mounted into the pod. // want "commentstart: godoc for field PersistentVolumeSource.StorageOS should start with 'storageos ...'"
	// Deprecated: StorageOS is deprecated and the in-tree storageos type is no longer supported.
	// More info: https://examples.k8s.io/volumes/storageos/README.md
	// +optional
	StorageOS *StorageOSPersistentVolumeSource `json:"storageos,omitempty" protobuf:"bytes,21,opt,name=storageos"`
	// csi represents storage that is handled by an external CSI driver.
	// +optional
	CSI *CSIPersistentVolumeSource `json:"csi,omitempty" protobuf:"bytes,22,opt,name=csi"`
}

// Represents a host path mapped into a pod.
// Host path volumes do not support ownership management or SELinux relabeling.
type HostPathVolumeSource struct {
	// path of the directory on the host.
	// If the path is a symlink, it will follow the link to the real path.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	Path string `json:"path" protobuf:"bytes,1,opt,name=path"` // want "optionalorrequired: field HostPathVolumeSource.Path must be marked as optional or required"
	// type for HostPath Volume
	// Defaults to ""
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	// +optional
	Type *HostPathType `json:"type,omitempty" protobuf:"bytes,2,opt,name=type"`
}

// +enum
type HostPathType string

const (
	// For backwards compatible, leave it empty if unset
	HostPathUnset HostPathType = ""
	// If nothing exists at the given path, an empty directory will be created there
	// as needed with file mode 0755, having the same group and ownership with Kubelet.
	HostPathDirectoryOrCreate HostPathType = "DirectoryOrCreate"
	// A directory must exist at the given path
	HostPathDirectory HostPathType = "Directory"
	// If nothing exists at the given path, an empty file will be created there
	// as needed with file mode 0644, having the same group and ownership with Kubelet.
	HostPathFileOrCreate HostPathType = "FileOrCreate"
	// A file must exist at the given path
	HostPathFile HostPathType = "File"
	// A UNIX socket must exist at the given path
	HostPathSocket HostPathType = "Socket"
	// A character device must exist at the given path
	HostPathCharDev HostPathType = "CharDevice"
	// A block device must exist at the given path
	HostPathBlockDev HostPathType = "BlockDevice"
)

// Represents an empty directory for a pod.
// Empty directory volumes support ownership management and SELinux relabeling.
type EmptyDirVolumeSource struct {
	// medium represents what type of storage medium should back this directory.
	// The default is "" which means to use the node's default medium.
	// Must be an empty string (default) or Memory.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
	// +optional
	Medium StorageMedium `json:"medium,omitempty" protobuf:"bytes,1,opt,name=medium,casttype=StorageMedium"` // want "optionalfields: field Medium should be a pointer."
	// sizeLimit is the total amount of local storage required for this EmptyDir volume.
	// The size limit is also applicable for memory medium.
	// The maximum usage on memory medium EmptyDir would be the minimum value between
	// the SizeLimit specified here and the sum of memory limits of all containers in a pod.
	// The default is nil which means that the limit is undefined.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
	// +optional
	// SizeLimit *resource.Quantity `json:"sizeLimit,omitempty" protobuf:"bytes,2,opt,name=sizeLimit"`
}

// Represents a Glusterfs mount that lasts the lifetime of a pod.
// Glusterfs volumes do not support ownership management or SELinux relabeling.
type GlusterfsVolumeSource struct {
	// endpoints is the endpoint name that details Glusterfs topology.
	EndpointsName string `json:"endpoints" protobuf:"bytes,1,opt,name=endpoints"` // want "optionalorrequired: field GlusterfsVolumeSource.EndpointsName must be marked as optional or required"

	// path is the Glusterfs volume path.
	// More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	Path string `json:"path" protobuf:"bytes,2,opt,name=path"` // want "optionalorrequired: field GlusterfsVolumeSource.Path must be marked as optional or required"

	// readOnly here will force the Glusterfs volume to be mounted with read-only permissions.
	// Defaults to false.
	// More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly"` // want "optionalfields: field ReadOnly should be a pointer."
}

// Represents a Glusterfs mount that lasts the lifetime of a pod.
// Glusterfs volumes do not support ownership management or SELinux relabeling.
type GlusterfsPersistentVolumeSource struct {
	// endpoints is the endpoint name that details Glusterfs topology.
	// More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	EndpointsName string `json:"endpoints" protobuf:"bytes,1,opt,name=endpoints"` // want "optionalorrequired: field GlusterfsPersistentVolumeSource.EndpointsName must be marked as optional or required"

	// path is the Glusterfs volume path.
	// More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	Path string `json:"path" protobuf:"bytes,2,opt,name=path"` // want "optionalorrequired: field GlusterfsPersistentVolumeSource.Path must be marked as optional or required"

	// readOnly here will force the Glusterfs volume to be mounted with read-only permissions.
	// Defaults to false.
	// More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly"` // want "optionalfields: field ReadOnly should be a pointer."

	// endpointsNamespace is the namespace that contains Glusterfs endpoint.
	// If this field is empty, the EndpointNamespace defaults to the same namespace as the bound PVC.
	// More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	// +optional
	EndpointsNamespace *string `json:"endpointsNamespace,omitempty" protobuf:"bytes,4,opt,name=endpointsNamespace"`
}

// Represents a Rados Block Device mount that lasts the lifetime of a pod.
// RBD volumes support ownership management and SELinux relabeling.
type RBDVolumeSource struct {
	// monitors is a collection of Ceph monitors.
	// More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	// +listType=atomic
	CephMonitors []string `json:"monitors" protobuf:"bytes,1,rep,name=monitors"` // want "optionalorrequired: field RBDVolumeSource.CephMonitors must be marked as optional or required"
	// image is the rados image name.
	// More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	RBDImage string `json:"image" protobuf:"bytes,2,opt,name=image"` // want "optionalorrequired: field RBDVolumeSource.RBDImage must be marked as optional or required"
	// fsType is the filesystem type of the volume that you want to mount.
	// Tip: Ensure that the filesystem type is supported by the host operating system.
	// Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd
	// TODO: how do we prevent errors in the filesystem from compromising the machine
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,3,opt,name=fsType"` // want "optionalfields: field FSType should be a pointer."
	// pool is the rados pool name.
	// Default is rbd.
	// More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	// +optional
	// +default="rbd"
	RBDPool string `json:"pool,omitempty" protobuf:"bytes,4,opt,name=pool"` // want "optionalfields: field RBDPool should be a pointer."
	// user is the rados user name.
	// Default is admin.
	// More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	// +optional
	// +default="admin"
	RadosUser string `json:"user,omitempty" protobuf:"bytes,5,opt,name=user"` // want "optionalfields: field RadosUser should be a pointer."
	// keyring is the path to key ring for RBDUser.
	// Default is /etc/ceph/keyring.
	// More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	// +optional
	// +default="/etc/ceph/keyring"
	Keyring string `json:"keyring,omitempty" protobuf:"bytes,6,opt,name=keyring"` // want "optionalfields: field Keyring should be a pointer."
	// secretRef is name of the authentication secret for RBDUser. If provided
	// overrides keyring.
	// Default is nil.
	// More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	// +optional
	SecretRef *LocalObjectReference `json:"secretRef,omitempty" protobuf:"bytes,7,opt,name=secretRef"`
	// readOnly here will force the ReadOnly setting in VolumeMounts.
	// Defaults to false.
	// More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,8,opt,name=readOnly"` // want "optionalfields: field ReadOnly should be a pointer."
}

// Represents a Rados Block Device mount that lasts the lifetime of a pod.
// RBD volumes support ownership management and SELinux relabeling.
type RBDPersistentVolumeSource struct {
	// monitors is a collection of Ceph monitors.
	// More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	// +listType=atomic
	CephMonitors []string `json:"monitors" protobuf:"bytes,1,rep,name=monitors"` // want "optionalorrequired: field RBDPersistentVolumeSource.CephMonitors must be marked as optional or required"
	// image is the rados image name.
	// More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	RBDImage string `json:"image" protobuf:"bytes,2,opt,name=image"` // want "optionalorrequired: field RBDPersistentVolumeSource.RBDImage must be marked as optional or required"
	// fsType is the filesystem type of the volume that you want to mount.
	// Tip: Ensure that the filesystem type is supported by the host operating system.
	// Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd
	// TODO: how do we prevent errors in the filesystem from compromising the machine
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,3,opt,name=fsType"` // want "optionalfields: field FSType should be a pointer."
	// pool is the rados pool name.
	// Default is rbd.
	// More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	// +optional
	// +default="rbd"
	RBDPool string `json:"pool,omitempty" protobuf:"bytes,4,opt,name=pool"` // want "optionalfields: field RBDPool should be a pointer."
	// user is the rados user name.
	// Default is admin.
	// More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	// +optional
	// +default="admin"
	RadosUser string `json:"user,omitempty" protobuf:"bytes,5,opt,name=user"` // want "optionalfields: field RadosUser should be a pointer."
	// keyring is the path to key ring for RBDUser.
	// Default is /etc/ceph/keyring.
	// More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	// +optional
	// +default="/etc/ceph/keyring"
	Keyring string `json:"keyring,omitempty" protobuf:"bytes,6,opt,name=keyring"` // want "optionalfields: field Keyring should be a pointer."
	// secretRef is name of the authentication secret for RBDUser. If provided
	// overrides keyring.
	// Default is nil.
	// More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	// +optional
	SecretRef *SecretReference `json:"secretRef,omitempty" protobuf:"bytes,7,opt,name=secretRef"`
	// readOnly here will force the ReadOnly setting in VolumeMounts.
	// Defaults to false.
	// More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,8,opt,name=readOnly"` // want "optionalfields: field ReadOnly should be a pointer."
}

// Represents a cinder volume resource in Openstack.
// A Cinder volume must exist before mounting to a container.
// The volume must also be in the same region as the kubelet.
// Cinder volumes support ownership management and SELinux relabeling.
type CinderVolumeSource struct {
	// volumeID used to identify the volume in cinder.
	// More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	VolumeID string `json:"volumeID" protobuf:"bytes,1,opt,name=volumeID"` // want "optionalorrequired: field CinderVolumeSource.VolumeID must be marked as optional or required"
	// fsType is the filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType"` // want "optionalfields: field FSType should be a pointer."
	// readOnly defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly"` // want "optionalfields: field ReadOnly should be a pointer."
	// secretRef is optional: points to a secret object containing parameters used to connect
	// to OpenStack.
	// +optional
	SecretRef *LocalObjectReference `json:"secretRef,omitempty" protobuf:"bytes,4,opt,name=secretRef"`
}

// Represents a cinder volume resource in Openstack.
// A Cinder volume must exist before mounting to a container.
// The volume must also be in the same region as the kubelet.
// Cinder volumes support ownership management and SELinux relabeling.
type CinderPersistentVolumeSource struct {
	// volumeID used to identify the volume in cinder.
	// More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	VolumeID string `json:"volumeID" protobuf:"bytes,1,opt,name=volumeID"` // want "optionalorrequired: field CinderPersistentVolumeSource.VolumeID must be marked as optional or required"
	// fsType Filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType"` // want "optionalfields: field FSType should be a pointer."
	// readOnly is Optional: Defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly"` // want "optionalfields: field ReadOnly should be a pointer."
	// secretRef is Optional: points to a secret object containing parameters used to connect
	// to OpenStack.
	// +optional
	SecretRef *SecretReference `json:"secretRef,omitempty" protobuf:"bytes,4,opt,name=secretRef"`
}

// Represents a Ceph Filesystem mount that lasts the lifetime of a pod
// Cephfs volumes do not support ownership management or SELinux relabeling.
type CephFSVolumeSource struct {
	// monitors is Required: Monitors is a collection of Ceph monitors
	// More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	// +listType=atomic
	Monitors []string `json:"monitors" protobuf:"bytes,1,rep,name=monitors"` // want "optionalorrequired: field CephFSVolumeSource.Monitors must be marked as optional or required"
	// path is Optional: Used as the mounted root, rather than the full Ceph tree, default is /
	// +optional
	Path string `json:"path,omitempty" protobuf:"bytes,2,opt,name=path"` // want "optionalfields: field Path should be a pointer."
	// user is optional: User is the rados user name, default is admin
	// More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	// +optional
	User string `json:"user,omitempty" protobuf:"bytes,3,opt,name=user"` // want "optionalfields: field User should be a pointer."
	// secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret
	// More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	// +optional
	SecretFile string `json:"secretFile,omitempty" protobuf:"bytes,4,opt,name=secretFile"` // want "optionalfields: field SecretFile should be a pointer."
	// secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty.
	// More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	// +optional
	SecretRef *LocalObjectReference `json:"secretRef,omitempty" protobuf:"bytes,5,opt,name=secretRef"`
	// readOnly is Optional: Defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,6,opt,name=readOnly"` // want "optionalfields: field ReadOnly should be a pointer."
}

// Represents a Ceph Filesystem mount that lasts the lifetime of a pod
// Cephfs volumes do not support ownership management or SELinux relabeling.
type CephFSPersistentVolumeSource struct {
	// monitors is Required: Monitors is a collection of Ceph monitors
	// More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	// +listType=atomic
	Monitors []string `json:"monitors" protobuf:"bytes,1,rep,name=monitors"` // want "optionalorrequired: field CephFSPersistentVolumeSource.Monitors must be marked as optional or required"
	// path is Optional: Used as the mounted root, rather than the full Ceph tree, default is /
	// +optional
	Path string `json:"path,omitempty" protobuf:"bytes,2,opt,name=path"` // want "optionalfields: field Path should be a pointer."
	// user is Optional: User is the rados user name, default is admin
	// More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	// +optional
	User string `json:"user,omitempty" protobuf:"bytes,3,opt,name=user"` // want "optionalfields: field User should be a pointer."
	// secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret
	// More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	// +optional
	SecretFile string `json:"secretFile,omitempty" protobuf:"bytes,4,opt,name=secretFile"` // want "optionalfields: field SecretFile should be a pointer."
	// secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty.
	// More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	// +optional
	SecretRef *SecretReference `json:"secretRef,omitempty" protobuf:"bytes,5,opt,name=secretRef"`
	// readOnly is Optional: Defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,6,opt,name=readOnly"` // want "optionalfields: field ReadOnly should be a pointer."
}

// Represents a Flocker volume mounted by the Flocker agent.
// One and only one of datasetName and datasetUUID should be set.
// Flocker volumes do not support ownership management or SELinux relabeling.
type FlockerVolumeSource struct {
	// datasetName is Name of the dataset stored as metadata -> name on the dataset for Flocker
	// should be considered as deprecated
	// +optional
	DatasetName string `json:"datasetName,omitempty" protobuf:"bytes,1,opt,name=datasetName"` // want "optionalfields: field DatasetName should be a pointer."
	// datasetUUID is the UUID of the dataset. This is unique identifier of a Flocker dataset
	// +optional
	DatasetUUID string `json:"datasetUUID,omitempty" protobuf:"bytes,2,opt,name=datasetUUID"` // want "optionalfields: field DatasetUUID should be a pointer."
}

// StorageMedium defines ways that storage can be allocated to a volume.
type StorageMedium string

const (
	StorageMediumDefault         StorageMedium = ""           // use whatever the default is for the node, assume anything we don't explicitly handle is this
	StorageMediumMemory          StorageMedium = "Memory"     // use memory (e.g. tmpfs on linux)
	StorageMediumHugePages       StorageMedium = "HugePages"  // use hugepages
	StorageMediumHugePagesPrefix StorageMedium = "HugePages-" // prefix for full medium notation HugePages-<size>
)

// Represents a Persistent Disk resource in Google Compute Engine.
//
// A GCE PD must exist before mounting to a container. The disk must
// also be in the same GCE project and zone as the kubelet. A GCE PD
// can only be mounted as read/write once or read-only many times. GCE
// PDs support ownership management and SELinux relabeling.
type GCEPersistentDiskVolumeSource struct {
	// pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	PDName string `json:"pdName" protobuf:"bytes,1,opt,name=pdName"` // want "optionalorrequired: field GCEPersistentDiskVolumeSource.PDName must be marked as optional or required"
	// fsType is filesystem type of the volume that you want to mount.
	// Tip: Ensure that the filesystem type is supported by the host operating system.
	// Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	// TODO: how do we prevent errors in the filesystem from compromising the machine
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType"` // want "optionalfields: field FSType should be a pointer."
	// partition is the partition in the volume that you want to mount.
	// If omitted, the default is to mount by volume name.
	// Examples: For volume /dev/sda1, you specify the partition as "1".
	// Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	// +optional
	Partition int32 `json:"partition,omitempty" protobuf:"varint,3,opt,name=partition"` // want "optionalfields: field Partition should be a pointer."
	// readOnly here will force the ReadOnly setting in VolumeMounts.
	// Defaults to false.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,4,opt,name=readOnly"` // want "optionalfields: field ReadOnly should be a pointer."
}

// Represents a Quobyte mount that lasts the lifetime of a pod.
// Quobyte volumes do not support ownership management or SELinux relabeling.
type QuobyteVolumeSource struct {
	// registry represents a single or multiple Quobyte Registry services
	// specified as a string as host:port pair (multiple entries are separated with commas)
	// which acts as the central registry for volumes
	Registry string `json:"registry" protobuf:"bytes,1,opt,name=registry"` // want "optionalorrequired: field QuobyteVolumeSource.Registry must be marked as optional or required"

	// volume is a string that references an already created Quobyte volume by name.
	Volume string `json:"volume" protobuf:"bytes,2,opt,name=volume"` // want "optionalorrequired: field QuobyteVolumeSource.Volume must be marked as optional or required"

	// readOnly here will force the Quobyte volume to be mounted with read-only permissions.
	// Defaults to false.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly"` // want "optionalfields: field ReadOnly should be a pointer."

	// user to map volume access to
	// Defaults to serivceaccount user
	// +optional
	User string `json:"user,omitempty" protobuf:"bytes,4,opt,name=user"` // want "optionalfields: field User should be a pointer."

	// group to map volume access to
	// Default is no group
	// +optional
	Group string `json:"group,omitempty" protobuf:"bytes,5,opt,name=group"` // want "optionalfields: field Group should be a pointer."

	// tenant owning the given Quobyte volume in the Backend
	// Used with dynamically provisioned Quobyte volumes, value is set by the plugin
	// +optional
	Tenant string `json:"tenant,omitempty" protobuf:"bytes,6,opt,name=tenant"` // want "optionalfields: field Tenant should be a pointer."
}

// FlexPersistentVolumeSource represents a generic persistent volume resource that is
// provisioned/attached using an exec based plugin.
type FlexPersistentVolumeSource struct {
	// driver is the name of the driver to use for this volume.
	Driver string `json:"driver" protobuf:"bytes,1,opt,name=driver"` // want "optionalorrequired: field FlexPersistentVolumeSource.Driver must be marked as optional or required"
	// fsType is the Filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType"` // want "optionalfields: field FSType should be a pointer."
	// secretRef is Optional: SecretRef is reference to the secret object containing
	// sensitive information to pass to the plugin scripts. This may be
	// empty if no secret object is specified. If the secret object
	// contains more than one secret, all secrets are passed to the plugin
	// scripts.
	// +optional
	SecretRef *SecretReference `json:"secretRef,omitempty" protobuf:"bytes,3,opt,name=secretRef"`
	// readOnly is Optional: defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,4,opt,name=readOnly"` // want "optionalfields: field ReadOnly should be a pointer."
	// options is Optional: this field holds extra command options if any.
	// +optional
	Options map[string]string `json:"options,omitempty" protobuf:"bytes,5,rep,name=options"`
}

// FlexVolume represents a generic volume resource that is
// provisioned/attached using an exec based plugin.
type FlexVolumeSource struct {
	// driver is the name of the driver to use for this volume.
	Driver string `json:"driver" protobuf:"bytes,1,opt,name=driver"` // want "optionalorrequired: field FlexVolumeSource.Driver must be marked as optional or required"
	// fsType is the filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType"` // want "optionalfields: field FSType should be a pointer."
	// secretRef is Optional: secretRef is reference to the secret object containing
	// sensitive information to pass to the plugin scripts. This may be
	// empty if no secret object is specified. If the secret object
	// contains more than one secret, all secrets are passed to the plugin
	// scripts.
	// +optional
	SecretRef *LocalObjectReference `json:"secretRef,omitempty" protobuf:"bytes,3,opt,name=secretRef"`
	// readOnly is Optional: defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,4,opt,name=readOnly"` // want "optionalfields: field ReadOnly should be a pointer."
	// options is Optional: this field holds extra command options if any.
	// +optional
	Options map[string]string `json:"options,omitempty" protobuf:"bytes,5,rep,name=options"`
}

// Represents a Persistent Disk resource in AWS.
//
// An AWS EBS disk must exist before mounting to a container. The disk
// must also be in the same AWS zone as the kubelet. An AWS EBS disk
// can only be mounted as read/write once. AWS EBS volumes support
// ownership management and SELinux relabeling.
type AWSElasticBlockStoreVolumeSource struct {
	// volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume).
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	VolumeID string `json:"volumeID" protobuf:"bytes,1,opt,name=volumeID"` // want "optionalorrequired: field AWSElasticBlockStoreVolumeSource.VolumeID must be marked as optional or required"
	// fsType is the filesystem type of the volume that you want to mount.
	// Tip: Ensure that the filesystem type is supported by the host operating system.
	// Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	// TODO: how do we prevent errors in the filesystem from compromising the machine
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType"` // want "optionalfields: field FSType should be a pointer."
	// partition is the partition in the volume that you want to mount.
	// If omitted, the default is to mount by volume name.
	// Examples: For volume /dev/sda1, you specify the partition as "1".
	// Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).
	// +optional
	Partition int32 `json:"partition,omitempty" protobuf:"varint,3,opt,name=partition"` // want "optionalfields: field Partition should be a pointer."
	// readOnly value true will force the readOnly setting in VolumeMounts.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,4,opt,name=readOnly"` // want "optionalfields: field ReadOnly should be a pointer."
}

// Represents a volume that is populated with the contents of a git repository.
// Git repo volumes do not support ownership management.
// Git repo volumes support SELinux relabeling.
//
// DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an
// EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir
// into the Pod's container.
type GitRepoVolumeSource struct {
	// repository is the URL
	Repository string `json:"repository" protobuf:"bytes,1,opt,name=repository"` // want "optionalorrequired: field GitRepoVolumeSource.Repository must be marked as optional or required"
	// revision is the commit hash for the specified revision.
	// +optional
	Revision string `json:"revision,omitempty" protobuf:"bytes,2,opt,name=revision"` // want "optionalfields: field Revision should be a pointer."
	// directory is the target directory name.
	// Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the
	// git repository.  Otherwise, if specified, the volume will contain the git repository in
	// the subdirectory with the given name.
	// +optional
	Directory string `json:"directory,omitempty" protobuf:"bytes,3,opt,name=directory"` // want "optionalfields: field Directory should be a pointer."
}

// Adapts a Secret into a volume.
//
// The contents of the target Secret's Data field will be presented in a volume
// as files using the keys in the Data field as the file names.
// Secret volumes support ownership management and SELinux relabeling.
type SecretVolumeSource struct {
	// secretName is the name of the secret in the pod's namespace to use.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#secret
	// +optional
	SecretName string `json:"secretName,omitempty" protobuf:"bytes,1,opt,name=secretName"` // want "optionalfields: field SecretName should be a pointer."
	// items If unspecified, each key-value pair in the Data field of the referenced
	// Secret will be projected into the volume as a file whose name is the
	// key and content is the value. If specified, the listed keys will be
	// projected into the specified paths, and unlisted keys will not be
	// present. If a key is specified which is not present in the Secret,
	// the volume setup will error unless it is marked optional. Paths must be
	// relative and may not contain the '..' path or start with '..'.
	// +optional
	// +listType=atomic
	Items []KeyToPath `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"` // want "arrayofstruct: SecretVolumeSource.Items is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// defaultMode is Optional: mode bits used to set permissions on created files by default.
	// Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.
	// YAML accepts both octal and decimal values, JSON requires decimal values
	// for mode bits. Defaults to 0644.
	// Directories within the path are not affected by this setting.
	// This might be in conflict with other options that affect the file
	// mode, like fsGroup, and the result can be other mode bits set.
	// +optional
	DefaultMode *int32 `json:"defaultMode,omitempty" protobuf:"bytes,3,opt,name=defaultMode"`
	// optional field specify whether the Secret or its keys must be defined
	// +optional
	Optional *bool `json:"optional,omitempty" protobuf:"varint,4,opt,name=optional"`
}

const (
	SecretVolumeSourceDefaultMode int32 = 0644
)

// Adapts a secret into a projected volume.
//
// The contents of the target Secret's Data field will be presented in a
// projected volume as files using the keys in the Data field as the file names.
// Note that this is identical to a secret volume source without the default
// mode.
type SecretProjection struct {
	LocalObjectReference `json:",inline" protobuf:"bytes,1,opt,name=localObjectReference"`
	// items if unspecified, each key-value pair in the Data field of the referenced
	// Secret will be projected into the volume as a file whose name is the
	// key and content is the value. If specified, the listed keys will be
	// projected into the specified paths, and unlisted keys will not be
	// present. If a key is specified which is not present in the Secret,
	// the volume setup will error unless it is marked optional. Paths must be
	// relative and may not contain the '..' path or start with '..'.
	// +optional
	// +listType=atomic
	Items []KeyToPath `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"` // want "arrayofstruct: SecretProjection.Items is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// optional field specify whether the Secret or its key must be defined
	// +optional
	Optional *bool `json:"optional,omitempty" protobuf:"varint,4,opt,name=optional"`
}

// Represents an NFS mount that lasts the lifetime of a pod.
// NFS volumes do not support ownership management or SELinux relabeling.
type NFSVolumeSource struct {
	// server is the hostname or IP address of the NFS server.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	Server string `json:"server" protobuf:"bytes,1,opt,name=server"` // want "optionalorrequired: field NFSVolumeSource.Server must be marked as optional or required"

	// path that is exported by the NFS server.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	Path string `json:"path" protobuf:"bytes,2,opt,name=path"` // want "optionalorrequired: field NFSVolumeSource.Path must be marked as optional or required"

	// readOnly here will force the NFS export to be mounted with read-only permissions.
	// Defaults to false.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly"` // want "optionalfields: field ReadOnly should be a pointer."
}

// Represents an ISCSI disk.
// ISCSI volumes can only be mounted as read/write once.
// ISCSI volumes support ownership management and SELinux relabeling.
type ISCSIVolumeSource struct {
	// targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port
	// is other than default (typically TCP ports 860 and 3260).
	TargetPortal string `json:"targetPortal" protobuf:"bytes,1,opt,name=targetPortal"` // want "optionalorrequired: field ISCSIVolumeSource.TargetPortal must be marked as optional or required"
	// iqn is the target iSCSI Qualified Name.
	IQN string `json:"iqn" protobuf:"bytes,2,opt,name=iqn"` // want "optionalorrequired: field ISCSIVolumeSource.IQN must be marked as optional or required"
	// lun represents iSCSI Target Lun number.
	Lun int32 `json:"lun" protobuf:"varint,3,opt,name=lun"` // want "optionalorrequired: field ISCSIVolumeSource.Lun must be marked as optional or required"
	// iscsiInterface is the interface Name that uses an iSCSI transport.
	// Defaults to 'default' (tcp).
	// +optional
	// +default="default"
	ISCSIInterface string `json:"iscsiInterface,omitempty" protobuf:"bytes,4,opt,name=iscsiInterface"` // want "optionalfields: field ISCSIInterface should be a pointer."
	// fsType is the filesystem type of the volume that you want to mount.
	// Tip: Ensure that the filesystem type is supported by the host operating system.
	// Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi
	// TODO: how do we prevent errors in the filesystem from compromising the machine
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,5,opt,name=fsType"` // want "optionalfields: field FSType should be a pointer."
	// readOnly here will force the ReadOnly setting in VolumeMounts.
	// Defaults to false.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,6,opt,name=readOnly"` // want "optionalfields: field ReadOnly should be a pointer."
	// portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port
	// is other than default (typically TCP ports 860 and 3260).
	// +optional
	// +listType=atomic
	Portals []string `json:"portals,omitempty" protobuf:"bytes,7,opt,name=portals"`
	// chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication
	// +optional
	DiscoveryCHAPAuth bool `json:"chapAuthDiscovery,omitempty" protobuf:"varint,8,opt,name=chapAuthDiscovery"` // want "optionalfields: field DiscoveryCHAPAuth should be a pointer."
	// chapAuthSession defines whether support iSCSI Session CHAP authentication
	// +optional
	SessionCHAPAuth bool `json:"chapAuthSession,omitempty" protobuf:"varint,11,opt,name=chapAuthSession"` // want "optionalfields: field SessionCHAPAuth should be a pointer."
	// secretRef is the CHAP Secret for iSCSI target and initiator authentication
	// +optional
	SecretRef *LocalObjectReference `json:"secretRef,omitempty" protobuf:"bytes,10,opt,name=secretRef"`
	// initiatorName is the custom iSCSI Initiator Name.
	// If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface
	// <target portal>:<volume name> will be created for the connection.
	// +optional
	InitiatorName *string `json:"initiatorName,omitempty" protobuf:"bytes,12,opt,name=initiatorName"`
}

// ISCSIPersistentVolumeSource represents an ISCSI disk.
// ISCSI volumes can only be mounted as read/write once.
// ISCSI volumes support ownership management and SELinux relabeling.
type ISCSIPersistentVolumeSource struct {
	// targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port
	// is other than default (typically TCP ports 860 and 3260).
	TargetPortal string `json:"targetPortal" protobuf:"bytes,1,opt,name=targetPortal"` // want "optionalorrequired: field ISCSIPersistentVolumeSource.TargetPortal must be marked as optional or required"
	// iqn is Target iSCSI Qualified Name.
	IQN string `json:"iqn" protobuf:"bytes,2,opt,name=iqn"` // want "optionalorrequired: field ISCSIPersistentVolumeSource.IQN must be marked as optional or required"
	// lun is iSCSI Target Lun number.
	Lun int32 `json:"lun" protobuf:"varint,3,opt,name=lun"` // want "optionalorrequired: field ISCSIPersistentVolumeSource.Lun must be marked as optional or required"
	// iscsiInterface is the interface Name that uses an iSCSI transport.
	// Defaults to 'default' (tcp).
	// +optional
	// +default="default"
	ISCSIInterface string `json:"iscsiInterface,omitempty" protobuf:"bytes,4,opt,name=iscsiInterface"` // want "optionalfields: field ISCSIInterface should be a pointer."
	// fsType is the filesystem type of the volume that you want to mount.
	// Tip: Ensure that the filesystem type is supported by the host operating system.
	// Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi
	// TODO: how do we prevent errors in the filesystem from compromising the machine
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,5,opt,name=fsType"` // want "optionalfields: field FSType should be a pointer."
	// readOnly here will force the ReadOnly setting in VolumeMounts.
	// Defaults to false.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,6,opt,name=readOnly"` // want "optionalfields: field ReadOnly should be a pointer."
	// portals is the iSCSI Target Portal List. The Portal is either an IP or ip_addr:port if the port
	// is other than default (typically TCP ports 860 and 3260).
	// +optional
	// +listType=atomic
	Portals []string `json:"portals,omitempty" protobuf:"bytes,7,opt,name=portals"`
	// chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication
	// +optional
	DiscoveryCHAPAuth bool `json:"chapAuthDiscovery,omitempty" protobuf:"varint,8,opt,name=chapAuthDiscovery"` // want "optionalfields: field DiscoveryCHAPAuth should be a pointer."
	// chapAuthSession defines whether support iSCSI Session CHAP authentication
	// +optional
	SessionCHAPAuth bool `json:"chapAuthSession,omitempty" protobuf:"varint,11,opt,name=chapAuthSession"` // want "optionalfields: field SessionCHAPAuth should be a pointer."
	// secretRef is the CHAP Secret for iSCSI target and initiator authentication
	// +optional
	SecretRef *SecretReference `json:"secretRef,omitempty" protobuf:"bytes,10,opt,name=secretRef"`
	// initiatorName is the custom iSCSI Initiator Name.
	// If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface
	// <target portal>:<volume name> will be created for the connection.
	// +optional
	InitiatorName *string `json:"initiatorName,omitempty" protobuf:"bytes,12,opt,name=initiatorName"`
}

// Represents a Fibre Channel volume.
// Fibre Channel volumes can only be mounted as read/write once.
// Fibre Channel volumes support ownership management and SELinux relabeling.
type FCVolumeSource struct {
	// targetWWNs is Optional: FC target worldwide names (WWNs)
	// +optional
	// +listType=atomic
	TargetWWNs []string `json:"targetWWNs,omitempty" protobuf:"bytes,1,rep,name=targetWWNs"`
	// lun is Optional: FC target lun number
	// +optional
	Lun *int32 `json:"lun,omitempty" protobuf:"varint,2,opt,name=lun"`
	// fsType is the filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// TODO: how do we prevent errors in the filesystem from compromising the machine
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,3,opt,name=fsType"` // want "optionalfields: field FSType should be a pointer."
	// readOnly is Optional: Defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,4,opt,name=readOnly"` // want "optionalfields: field ReadOnly should be a pointer."
	// wwids Optional: FC volume world wide identifiers (wwids)
	// Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.
	// +optional
	// +listType=atomic
	WWIDs []string `json:"wwids,omitempty" protobuf:"bytes,5,rep,name=wwids"`
}

// AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
type AzureFileVolumeSource struct {
	// secretName is the  name of secret that contains Azure Storage Account Name and Key
	SecretName string `json:"secretName" protobuf:"bytes,1,opt,name=secretName"` // want "optionalorrequired: field AzureFileVolumeSource.SecretName must be marked as optional or required"
	// shareName is the azure share Name
	ShareName string `json:"shareName" protobuf:"bytes,2,opt,name=shareName"` // want "optionalorrequired: field AzureFileVolumeSource.ShareName must be marked as optional or required"
	// readOnly defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly"` // want "optionalfields: field ReadOnly should be a pointer."
}

// AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
type AzureFilePersistentVolumeSource struct {
	// secretName is the name of secret that contains Azure Storage Account Name and Key
	SecretName string `json:"secretName" protobuf:"bytes,1,opt,name=secretName"` // want "optionalorrequired: field AzureFilePersistentVolumeSource.SecretName must be marked as optional or required"
	// shareName is the azure Share Name
	ShareName string `json:"shareName" protobuf:"bytes,2,opt,name=shareName"` // want "optionalorrequired: field AzureFilePersistentVolumeSource.ShareName must be marked as optional or required"
	// readOnly defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly"` // want "optionalfields: field ReadOnly should be a pointer."
	// secretNamespace is the namespace of the secret that contains Azure Storage Account Name and Key
	// default is the same as the Pod
	// +optional
	SecretNamespace *string `json:"secretNamespace" protobuf:"bytes,4,opt,name=secretNamespace"` // want "optionalfields: field SecretNamespace should have the omitempty tag."
}

// Represents a vSphere volume resource.
type VsphereVirtualDiskVolumeSource struct {
	// volumePath is the path that identifies vSphere volume vmdk
	VolumePath string `json:"volumePath" protobuf:"bytes,1,opt,name=volumePath"` // want "optionalorrequired: field VsphereVirtualDiskVolumeSource.VolumePath must be marked as optional or required"
	// fsType is filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType"` // want "optionalfields: field FSType should be a pointer."
	// storagePolicyName is the storage Policy Based Management (SPBM) profile name.
	// +optional
	StoragePolicyName string `json:"storagePolicyName,omitempty" protobuf:"bytes,3,opt,name=storagePolicyName"` // want "optionalfields: field StoragePolicyName should be a pointer."
	// storagePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.
	// +optional
	StoragePolicyID string `json:"storagePolicyID,omitempty" protobuf:"bytes,4,opt,name=storagePolicyID"` // want "optionalfields: field StoragePolicyID should be a pointer."
}

// Represents a Photon Controller persistent disk resource.
type PhotonPersistentDiskVolumeSource struct {
	// pdID is the ID that identifies Photon Controller persistent disk
	PdID string `json:"pdID" protobuf:"bytes,1,opt,name=pdID"` // want "optionalorrequired: field PhotonPersistentDiskVolumeSource.PdID must be marked as optional or required"
	// fsType is the filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	FSType string `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType"` // want "optionalorrequired: field PhotonPersistentDiskVolumeSource.FSType must be marked as optional or required"
}

// +enum
type AzureDataDiskCachingMode string

// +enum
type AzureDataDiskKind string

const (
	AzureDataDiskCachingNone      AzureDataDiskCachingMode = "None"
	AzureDataDiskCachingReadOnly  AzureDataDiskCachingMode = "ReadOnly"
	AzureDataDiskCachingReadWrite AzureDataDiskCachingMode = "ReadWrite"

	AzureSharedBlobDisk    AzureDataDiskKind = "Shared"
	AzureDedicatedBlobDisk AzureDataDiskKind = "Dedicated"
	AzureManagedDisk       AzureDataDiskKind = "Managed"
)

// AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
type AzureDiskVolumeSource struct {
	// diskName is the Name of the data disk in the blob storage
	DiskName string `json:"diskName" protobuf:"bytes,1,opt,name=diskName"` // want "optionalorrequired: field AzureDiskVolumeSource.DiskName must be marked as optional or required"
	// diskURI is the URI of data disk in the blob storage
	DataDiskURI string `json:"diskURI" protobuf:"bytes,2,opt,name=diskURI"` // want "optionalorrequired: field AzureDiskVolumeSource.DataDiskURI must be marked as optional or required"
	// cachingMode is the Host Caching mode: None, Read Only, Read Write.
	// +optional
	// +default=ref(AzureDataDiskCachingReadWrite)
	CachingMode *AzureDataDiskCachingMode `json:"cachingMode,omitempty" protobuf:"bytes,3,opt,name=cachingMode,casttype=AzureDataDiskCachingMode"`
	// fsType is Filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// +optional
	// +default="ext4"
	FSType *string `json:"fsType,omitempty" protobuf:"bytes,4,opt,name=fsType"`
	// readOnly Defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// +optional
	// +default=false
	ReadOnly *bool `json:"readOnly,omitempty" protobuf:"varint,5,opt,name=readOnly"`
	// kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared
	// +default=ref(AzureSharedBlobDisk)
	Kind *AzureDataDiskKind `json:"kind,omitempty" protobuf:"bytes,6,opt,name=kind,casttype=AzureDataDiskKind"` // want "optionalorrequired: field AzureDiskVolumeSource.Kind must be marked as optional or required"
}

// PortworxVolumeSource represents a Portworx volume resource.
type PortworxVolumeSource struct {
	// volumeID uniquely identifies a Portworx volume
	VolumeID string `json:"volumeID" protobuf:"bytes,1,opt,name=volumeID"` // want "optionalorrequired: field PortworxVolumeSource.VolumeID must be marked as optional or required"
	// fSType represents the filesystem type to mount // want "commentstart: godoc for field PortworxVolumeSource.FSType should start with 'fsType ...'"
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs". Implicitly inferred to be "ext4" if unspecified.
	FSType string `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType"` // want "optionalorrequired: field PortworxVolumeSource.FSType must be marked as optional or required"
	// readOnly defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly"` // want "optionalfields: field ReadOnly should be a pointer."
}

// ScaleIOVolumeSource represents a persistent ScaleIO volume
type ScaleIOVolumeSource struct {
	// gateway is the host address of the ScaleIO API Gateway.
	Gateway string `json:"gateway" protobuf:"bytes,1,opt,name=gateway"` // want "optionalorrequired: field ScaleIOVolumeSource.Gateway must be marked as optional or required"
	// system is the name of the storage system as configured in ScaleIO.
	System string `json:"system" protobuf:"bytes,2,opt,name=system"` // want "optionalorrequired: field ScaleIOVolumeSource.System must be marked as optional or required"
	// secretRef references to the secret for ScaleIO user and other
	// sensitive information. If this is not provided, Login operation will fail.
	SecretRef *LocalObjectReference `json:"secretRef" protobuf:"bytes,3,opt,name=secretRef"` // want "optionalorrequired: field ScaleIOVolumeSource.SecretRef must be marked as optional or required"
	// sslEnabled Flag enable/disable SSL communication with Gateway, default false
	// +optional
	SSLEnabled bool `json:"sslEnabled,omitempty" protobuf:"varint,4,opt,name=sslEnabled"` // want "optionalfields: field SSLEnabled should be a pointer."
	// protectionDomain is the name of the ScaleIO Protection Domain for the configured storage.
	// +optional
	ProtectionDomain string `json:"protectionDomain,omitempty" protobuf:"bytes,5,opt,name=protectionDomain"` // want "optionalfields: field ProtectionDomain should be a pointer."
	// storagePool is the ScaleIO Storage Pool associated with the protection domain.
	// +optional
	StoragePool string `json:"storagePool,omitempty" protobuf:"bytes,6,opt,name=storagePool"` // want "optionalfields: field StoragePool should be a pointer."
	// storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned.
	// Default is ThinProvisioned.
	// +optional
	// +default="ThinProvisioned"
	StorageMode string `json:"storageMode,omitempty" protobuf:"bytes,7,opt,name=storageMode"` // want "optionalfields: field StorageMode should be a pointer."
	// volumeName is the name of a volume already created in the ScaleIO system
	// that is associated with this volume source.
	VolumeName string `json:"volumeName,omitempty" protobuf:"bytes,8,opt,name=volumeName"` // want "optionalorrequired: field ScaleIOVolumeSource.VolumeName must be marked as optional or required"
	// fsType is the filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs".
	// Default is "xfs".
	// +optional
	// +default="xfs"
	FSType string `json:"fsType,omitempty" protobuf:"bytes,9,opt,name=fsType"` // want "optionalfields: field FSType should be a pointer."
	// readOnly Defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,10,opt,name=readOnly"` // want "optionalfields: field ReadOnly should be a pointer."
}

// ScaleIOPersistentVolumeSource represents a persistent ScaleIO volume
type ScaleIOPersistentVolumeSource struct {
	// gateway is the host address of the ScaleIO API Gateway.
	Gateway string `json:"gateway" protobuf:"bytes,1,opt,name=gateway"` // want "optionalorrequired: field ScaleIOPersistentVolumeSource.Gateway must be marked as optional or required"
	// system is the name of the storage system as configured in ScaleIO.
	System string `json:"system" protobuf:"bytes,2,opt,name=system"` // want "optionalorrequired: field ScaleIOPersistentVolumeSource.System must be marked as optional or required"
	// secretRef references to the secret for ScaleIO user and other
	// sensitive information. If this is not provided, Login operation will fail.
	SecretRef *SecretReference `json:"secretRef" protobuf:"bytes,3,opt,name=secretRef"` // want "optionalorrequired: field ScaleIOPersistentVolumeSource.SecretRef must be marked as optional or required"
	// sslEnabled is the flag to enable/disable SSL communication with Gateway, default false
	// +optional
	SSLEnabled bool `json:"sslEnabled,omitempty" protobuf:"varint,4,opt,name=sslEnabled"` // want "optionalfields: field SSLEnabled should be a pointer."
	// protectionDomain is the name of the ScaleIO Protection Domain for the configured storage.
	// +optional
	ProtectionDomain string `json:"protectionDomain,omitempty" protobuf:"bytes,5,opt,name=protectionDomain"` // want "optionalfields: field ProtectionDomain should be a pointer."
	// storagePool is the ScaleIO Storage Pool associated with the protection domain.
	// +optional
	StoragePool string `json:"storagePool,omitempty" protobuf:"bytes,6,opt,name=storagePool"` // want "optionalfields: field StoragePool should be a pointer."
	// storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned.
	// Default is ThinProvisioned.
	// +optional
	// +default="ThinProvisioned"
	StorageMode string `json:"storageMode,omitempty" protobuf:"bytes,7,opt,name=storageMode"` // want "optionalfields: field StorageMode should be a pointer."
	// volumeName is the name of a volume already created in the ScaleIO system
	// that is associated with this volume source.
	VolumeName string `json:"volumeName,omitempty" protobuf:"bytes,8,opt,name=volumeName"` // want "optionalorrequired: field ScaleIOPersistentVolumeSource.VolumeName must be marked as optional or required"
	// fsType is the filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs".
	// Default is "xfs"
	// +optional
	// +default="xfs"
	FSType string `json:"fsType,omitempty" protobuf:"bytes,9,opt,name=fsType"` // want "optionalfields: field FSType should be a pointer."
	// readOnly defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,10,opt,name=readOnly"` // want "optionalfields: field ReadOnly should be a pointer."
}

// Represents a StorageOS persistent volume resource.
type StorageOSVolumeSource struct {
	// volumeName is the human-readable name of the StorageOS volume.  Volume
	// names are only unique within a namespace.
	VolumeName string `json:"volumeName,omitempty" protobuf:"bytes,1,opt,name=volumeName"` // want "optionalorrequired: field StorageOSVolumeSource.VolumeName must be marked as optional or required"
	// volumeNamespace specifies the scope of the volume within StorageOS.  If no
	// namespace is specified then the Pod's namespace will be used.  This allows the
	// Kubernetes name scoping to be mirrored within StorageOS for tighter integration.
	// Set VolumeName to any name to override the default behaviour.
	// Set to "default" if you are not using namespaces within StorageOS.
	// Namespaces that do not pre-exist within StorageOS will be created.
	// +optional
	VolumeNamespace string `json:"volumeNamespace,omitempty" protobuf:"bytes,2,opt,name=volumeNamespace"` // want "optionalfields: field VolumeNamespace should be a pointer."
	// fsType is the filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,3,opt,name=fsType"` // want "optionalfields: field FSType should be a pointer."
	// readOnly defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,4,opt,name=readOnly"` // want "optionalfields: field ReadOnly should be a pointer."
	// secretRef specifies the secret to use for obtaining the StorageOS API
	// credentials.  If not specified, default values will be attempted.
	// +optional
	SecretRef *LocalObjectReference `json:"secretRef,omitempty" protobuf:"bytes,5,opt,name=secretRef"`
}

// Represents a StorageOS persistent volume resource.
type StorageOSPersistentVolumeSource struct {
	// volumeName is the human-readable name of the StorageOS volume.  Volume
	// names are only unique within a namespace.
	VolumeName string `json:"volumeName,omitempty" protobuf:"bytes,1,opt,name=volumeName"` // want "optionalorrequired: field StorageOSPersistentVolumeSource.VolumeName must be marked as optional or required"
	// volumeNamespace specifies the scope of the volume within StorageOS.  If no
	// namespace is specified then the Pod's namespace will be used.  This allows the
	// Kubernetes name scoping to be mirrored within StorageOS for tighter integration.
	// Set VolumeName to any name to override the default behaviour.
	// Set to "default" if you are not using namespaces within StorageOS.
	// Namespaces that do not pre-exist within StorageOS will be created.
	// +optional
	VolumeNamespace string `json:"volumeNamespace,omitempty" protobuf:"bytes,2,opt,name=volumeNamespace"` // want "optionalfields: field VolumeNamespace should be a pointer."
	// fsType is the filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,3,opt,name=fsType"` // want "optionalfields: field FSType should be a pointer."
	// readOnly defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,4,opt,name=readOnly"` // want "optionalfields: field ReadOnly should be a pointer."
	// secretRef specifies the secret to use for obtaining the StorageOS API
	// credentials.  If not specified, default values will be attempted.
	// +optional
	SecretRef *ObjectReference `json:"secretRef,omitempty" protobuf:"bytes,5,opt,name=secretRef"`
}

// Adapts a ConfigMap into a volume.
//
// The contents of the target ConfigMap's Data field will be presented in a
// volume as files using the keys in the Data field as the file names, unless
// the items element is populated with specific mappings of keys to paths.
// ConfigMap volumes support ownership management and SELinux relabeling.
type ConfigMapVolumeSource struct {
	LocalObjectReference `json:",inline" protobuf:"bytes,1,opt,name=localObjectReference"`
	// items if unspecified, each key-value pair in the Data field of the referenced
	// ConfigMap will be projected into the volume as a file whose name is the
	// key and content is the value. If specified, the listed keys will be
	// projected into the specified paths, and unlisted keys will not be
	// present. If a key is specified which is not present in the ConfigMap,
	// the volume setup will error unless it is marked optional. Paths must be
	// relative and may not contain the '..' path or start with '..'.
	// +optional
	// +listType=atomic
	Items []KeyToPath `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"` // want "arrayofstruct: ConfigMapVolumeSource.Items is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// defaultMode is optional: mode bits used to set permissions on created files by default.
	// Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.
	// YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.
	// Defaults to 0644.
	// Directories within the path are not affected by this setting.
	// This might be in conflict with other options that affect the file
	// mode, like fsGroup, and the result can be other mode bits set.
	// +optional
	DefaultMode *int32 `json:"defaultMode,omitempty" protobuf:"varint,3,opt,name=defaultMode"`
	// optional specify whether the ConfigMap or its keys must be defined
	// +optional
	Optional *bool `json:"optional,omitempty" protobuf:"varint,4,opt,name=optional"`
}

// Projection that may be projected along with other supported volume types.
// Exactly one of these fields must be set.
type VolumeProjection struct {
	// secret information about the secret data to project
	// +optional
	Secret *SecretProjection `json:"secret,omitempty" protobuf:"bytes,1,opt,name=secret"`
	// downwardAPI information about the downwardAPI data to project
	// +optional
	DownwardAPI *DownwardAPIProjection `json:"downwardAPI,omitempty" protobuf:"bytes,2,opt,name=downwardAPI"`
	// configMap information about the configMap data to project
	// +optional
	ConfigMap *ConfigMapProjection `json:"configMap,omitempty" protobuf:"bytes,3,opt,name=configMap"`
	// serviceAccountToken is information about the serviceAccountToken data to project
	// +optional
	ServiceAccountToken *ServiceAccountTokenProjection `json:"serviceAccountToken,omitempty" protobuf:"bytes,4,opt,name=serviceAccountToken"`

	// ClusterTrustBundle allows a pod to access the `.spec.trustBundle` field // want "commentstart: godoc for field VolumeProjection.ClusterTrustBundle should start with 'clusterTrustBundle ...'"
	// of ClusterTrustBundle objects in an auto-updating file.
	//
	// Alpha, gated by the ClusterTrustBundleProjection feature gate.
	//
	// ClusterTrustBundle objects can either be selected by name, or by the
	// combination of signer name and a label selector.
	//
	// Kubelet performs aggressive normalization of the PEM contents written
	// into the pod filesystem.  Esoteric PEM features such as inter-block
	// comments and block headers are stripped.  Certificates are deduplicated.
	// The ordering of certificates within the file is arbitrary, and Kubelet
	// may change the order over time.
	//
	// +featureGate=ClusterTrustBundleProjection
	// +optional
	ClusterTrustBundle *ClusterTrustBundleProjection `json:"clusterTrustBundle,omitempty" protobuf:"bytes,5,opt,name=clusterTrustBundle"`

	// Projects an auto-rotating credential bundle (private key and certificate // want "commentstart: godoc for field VolumeProjection.PodCertificate should start with 'podCertificate ...'"
	// chain) that the pod can use either as a TLS client or server.
	//
	// Kubelet generates a private key and uses it to send a
	// PodCertificateRequest to the named signer.  Once the signer approves the
	// request and issues a certificate chain, Kubelet writes the key and
	// certificate chain to the pod filesystem.  The pod does not start until
	// certificates have been issued for each podCertificate projected volume
	// source in its spec.
	//
	// Kubelet will begin trying to rotate the certificate at the time indicated
	// by the signer using the PodCertificateRequest.Status.BeginRefreshAt
	// timestamp.
	//
	// Kubelet can write a single file, indicated by the credentialBundlePath
	// field, or separate files, indicated by the keyPath and
	// certificateChainPath fields.
	//
	// The credential bundle is a single file in PEM format.  The first PEM
	// entry is the private key (in PKCS#8 format), and the remaining PEM
	// entries are the certificate chain issued by the signer (typically,
	// signers will return their certificate chain in leaf-to-root order).
	//
	// Prefer using the credential bundle format, since your application code
	// can read it atomically.  If you use keyPath and certificateChainPath,
	// your application must make two separate file reads. If these coincide
	// with a certificate rotation, it is possible that the private key and leaf
	// certificate you read may not correspond to each other.  Your application
	// will need to check for this condition, and re-read until they are
	// consistent.
	//
	// The named signer controls chooses the format of the certificate it
	// issues; consult the signer implementation's documentation to learn how to
	// use the certificates it issues.
	//
	// +featureGate=PodCertificateProjection
	// +optional
	PodCertificate *PodCertificateProjection `json:"podCertificate,omitempty" protobuf:"bytes,6,opt,name=podCertificate"`
}

// Maps a string key to a path within a volume.
type KeyToPath struct {
	// key is the key to project.
	Key string `json:"key" protobuf:"bytes,1,opt,name=key"` // want "optionalorrequired: field KeyToPath.Key must be marked as optional or required"

	// path is the relative path of the file to map the key to.
	// May not be an absolute path.
	// May not contain the path element '..'.
	// May not start with the string '..'.
	Path string `json:"path" protobuf:"bytes,2,opt,name=path"` // want "optionalorrequired: field KeyToPath.Path must be marked as optional or required"
	// mode is Optional: mode bits used to set permissions on this file.
	// Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.
	// YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.
	// If not specified, the volume defaultMode will be used.
	// This might be in conflict with other options that affect the file
	// mode, like fsGroup, and the result can be other mode bits set.
	// +optional
	Mode *int32 `json:"mode,omitempty" protobuf:"varint,3,opt,name=mode"`
}

// DownwardAPIVolumeSource represents a volume containing downward API info.
// Downward API volumes support ownership management and SELinux relabeling.
type DownwardAPIVolumeSource struct {
	// Items is a list of downward API volume file // want "commentstart: godoc for field DownwardAPIVolumeSource.Items should start with 'items ...'"
	// +optional
	// +listType=atomic
	Items []DownwardAPIVolumeFile `json:"items,omitempty" protobuf:"bytes,1,rep,name=items"` // want "arrayofstruct: DownwardAPIVolumeSource.Items is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// Optional: mode bits to use on created files by default. Must be a // want "commentstart: godoc for field DownwardAPIVolumeSource.DefaultMode should start with 'defaultMode ...'"
	// Optional: mode bits used to set permissions on created files by default.
	// Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.
	// YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.
	// Defaults to 0644.
	// Directories within the path are not affected by this setting.
	// This might be in conflict with other options that affect the file
	// mode, like fsGroup, and the result can be other mode bits set.
	// +optional
	DefaultMode *int32 `json:"defaultMode,omitempty" protobuf:"varint,2,opt,name=defaultMode"`
}

// DownwardAPIVolumeFile represents information to create the file containing the pod field
type DownwardAPIVolumeFile struct {
	// Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..' // want "commentstart: godoc for field DownwardAPIVolumeFile.Path should start with 'path ...'"
	Path string `json:"path" protobuf:"bytes,1,opt,name=path"` // want "optionalorrequired: field DownwardAPIVolumeFile.Path must be marked as optional or required"
	// Required: Selects a field of the pod: only annotations, labels, name, namespace and uid are supported. // want "commentstart: godoc for field DownwardAPIVolumeFile.FieldRef should start with 'fieldRef ...'"
	// +optional
	FieldRef *ObjectFieldSelector `json:"fieldRef,omitempty" protobuf:"bytes,2,opt,name=fieldRef"`
	// Selects a resource of the container: only resources limits and requests // want "commentstart: godoc for field DownwardAPIVolumeFile.ResourceFieldRef should start with 'resourceFieldRef ...'"
	// (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.
	// +optional
	ResourceFieldRef *ResourceFieldSelector `json:"resourceFieldRef,omitempty" protobuf:"bytes,3,opt,name=resourceFieldRef"`
	// Optional: mode bits used to set permissions on this file, must be an octal value // want "commentstart: godoc for field DownwardAPIVolumeFile.Mode should start with 'mode ...'"
	// between 0000 and 0777 or a decimal value between 0 and 511.
	// YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.
	// If not specified, the volume defaultMode will be used.
	// This might be in conflict with other options that affect the file
	// mode, like fsGroup, and the result can be other mode bits set.
	// +optional
	Mode *int32 `json:"mode,omitempty" protobuf:"varint,4,opt,name=mode"`
}

// Represents downward API info for projecting into a projected volume.
// Note that this is identical to a downwardAPI volume source without the default
// mode.
type DownwardAPIProjection struct {
	// Items is a list of DownwardAPIVolume file // want "commentstart: godoc for field DownwardAPIProjection.Items should start with 'items ...'"
	// +optional
	// +listType=atomic
	Items []DownwardAPIVolumeFile `json:"items,omitempty" protobuf:"bytes,1,rep,name=items"` // want "arrayofstruct: DownwardAPIProjection.Items is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
}

// Represents a source location of a volume to mount, managed by an external CSI driver
type CSIVolumeSource struct {
	// driver is the name of the CSI driver that handles this volume.
	// Consult with your admin for the correct name as registered in the cluster.
	Driver string `json:"driver" protobuf:"bytes,1,opt,name=driver"` // want "optionalorrequired: field CSIVolumeSource.Driver must be marked as optional or required"

	// readOnly specifies a read-only configuration for the volume.
	// Defaults to false (read/write).
	// +optional
	ReadOnly *bool `json:"readOnly,omitempty" protobuf:"varint,2,opt,name=readOnly"`

	// fsType to mount. Ex. "ext4", "xfs", "ntfs".
	// If not provided, the empty value is passed to the associated CSI driver
	// which will determine the default filesystem to apply.
	// +optional
	FSType *string `json:"fsType,omitempty" protobuf:"bytes,3,opt,name=fsType"`

	// volumeAttributes stores driver-specific properties that are passed to the CSI
	// driver. Consult your driver's documentation for supported values.
	// +optional
	VolumeAttributes map[string]string `json:"volumeAttributes,omitempty" protobuf:"bytes,4,rep,name=volumeAttributes"`

	// nodePublishSecretRef is a reference to the secret object containing
	// sensitive information to pass to the CSI driver to complete the CSI
	// NodePublishVolume and NodeUnpublishVolume calls.
	// This field is optional, and  may be empty if no secret is required. If the
	// secret object contains more than one secret, all secret references are passed.
	// +optional
	NodePublishSecretRef *LocalObjectReference `json:"nodePublishSecretRef,omitempty" protobuf:"bytes,5,opt,name=nodePublishSecretRef"`
}

// Represents an ephemeral volume that is handled by a normal storage driver.
type EphemeralVolumeSource struct {
	// Will be used to create a stand-alone PVC to provision the volume. // want "commentstart: godoc for field EphemeralVolumeSource.VolumeClaimTemplate should start with 'volumeClaimTemplate ...'"
	// The pod in which this EphemeralVolumeSource is embedded will be the
	// owner of the PVC, i.e. the PVC will be deleted together with the
	// pod.  The name of the PVC will be `<pod name>-<volume name>` where
	// `<volume name>` is the name from the `PodSpec.Volumes` array
	// entry. Pod validation will reject the pod if the concatenated name
	// is not valid for a PVC (for example, too long).
	//
	// An existing PVC with that name that is not owned by the pod
	// will *not* be used for the pod to avoid using an unrelated
	// volume by mistake. Starting the pod is then blocked until
	// the unrelated PVC is removed. If such a pre-created PVC is
	// meant to be used by the pod, the PVC has to updated with an
	// owner reference to the pod once the pod exists. Normally
	// this should not be necessary, but it may be useful when
	// manually reconstructing a broken cluster.
	//
	// This field is read-only and no changes will be made by Kubernetes
	// to the PVC after it has been created.
	//
	// Required, must not be nil.
	VolumeClaimTemplate *PersistentVolumeClaimTemplate `json:"volumeClaimTemplate,omitempty" protobuf:"bytes,1,opt,name=volumeClaimTemplate"` // want "optionalorrequired: field EphemeralVolumeSource.VolumeClaimTemplate must be marked as optional or required"

	// ReadOnly is tombstoned to show why 2 is a reserved protobuf tag.
	// ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,2,opt,name=readOnly"`
}

// PersistentVolumeClaimTemplate is used to produce
// PersistentVolumeClaim objects as part of an EphemeralVolumeSource.
type PersistentVolumeClaimTemplate struct {
	// May contain labels and annotations that will be copied into the PVC
	// when creating it. No other fields are allowed and will be rejected during
	// validation.
	//
	// +optional
	// metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// The specification for the PersistentVolumeClaim. The entire content is // want "commentstart: godoc for field PersistentVolumeClaimTemplate.Spec should start with 'spec ...'"
	// copied unchanged into the PVC that gets created from this
	// template. The same fields as in a PersistentVolumeClaim
	// are also valid here.
	Spec PersistentVolumeClaimSpec `json:"spec" protobuf:"bytes,2,name=spec"` // want "optionalorrequired: field PersistentVolumeClaimTemplate.Spec must be marked as optional or required"
}

// Adapts a ConfigMap into a projected volume.
//
// The contents of the target ConfigMap's Data field will be presented in a
// projected volume as files using the keys in the Data field as the file names,
// unless the items element is populated with specific mappings of keys to paths.
// Note that this is identical to a configmap volume source without the default
// mode.
type ConfigMapProjection struct {
	LocalObjectReference `json:",inline" protobuf:"bytes,1,opt,name=localObjectReference"`
	// items if unspecified, each key-value pair in the Data field of the referenced
	// ConfigMap will be projected into the volume as a file whose name is the
	// key and content is the value. If specified, the listed keys will be
	// projected into the specified paths, and unlisted keys will not be
	// present. If a key is specified which is not present in the ConfigMap,
	// the volume setup will error unless it is marked optional. Paths must be
	// relative and may not contain the '..' path or start with '..'.
	// +optional
	// +listType=atomic
	Items []KeyToPath `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"` // want "arrayofstruct: ConfigMapProjection.Items is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// optional specify whether the ConfigMap or its keys must be defined
	// +optional
	Optional *bool `json:"optional,omitempty" protobuf:"varint,4,opt,name=optional"`
}

// ServiceAccountTokenProjection represents a projected service account token
// volume. This projection can be used to insert a service account token into
// the pods runtime filesystem for use against APIs (Kubernetes API Server or
// otherwise).
type ServiceAccountTokenProjection struct {
	// audience is the intended audience of the token. A recipient of a token
	// must identify itself with an identifier specified in the audience of the
	// token, and otherwise should reject the token. The audience defaults to the
	// identifier of the apiserver.
	// +optional
	Audience string `json:"audience,omitempty" protobuf:"bytes,1,rep,name=audience"` // want "optionalfields: field Audience should be a pointer."
	// expirationSeconds is the requested duration of validity of the service
	// account token. As the token approaches expiration, the kubelet volume
	// plugin will proactively rotate the service account token. The kubelet will
	// start trying to rotate the token if the token is older than 80 percent of
	// its time to live or if the token is older than 24 hours.Defaults to 1 hour
	// and must be at least 10 minutes.
	// +optional
	ExpirationSeconds *int64 `json:"expirationSeconds,omitempty" protobuf:"varint,2,opt,name=expirationSeconds"`
	// path is the path relative to the mount point of the file to project the
	// token into.
	Path string `json:"path" protobuf:"bytes,3,opt,name=path"` // want "optionalorrequired: field ServiceAccountTokenProjection.Path must be marked as optional or required"
}

// ClusterTrustBundleProjection describes how to select a set of
// ClusterTrustBundle objects and project their contents into the pod
// filesystem.
type ClusterTrustBundleProjection struct {
	// Select a single ClusterTrustBundle by object name.  Mutually-exclusive with // want "commentstart: godoc for field ClusterTrustBundleProjection.Name should start with 'name ...'"
	// with signerName and labelSelector.
	// +optional
	Name *string `json:"name,omitempty" protobuf:"bytes,1,rep,name=name"`

	// Select all ClusterTrustBundles that match this signer name. // want "commentstart: godoc for field ClusterTrustBundleProjection.SignerName should start with 'signerName ...'"
	// Mutually-exclusive with name.  The contents of all selected
	// ClusterTrustBundles will be unified and deduplicated.
	// +optional
	SignerName *string `json:"signerName,omitempty" protobuf:"bytes,2,rep,name=signerName"`

	// Select all ClusterTrustBundles that match this label selector.  Only has
	// effect if signerName is set.  Mutually-exclusive with name.  If unset,
	// interpreted as "match nothing".  If set but empty, interpreted as "match
	// everything".
	// +optional
	// LabelSelector *metav1.LabelSelector `json:"labelSelector,omitempty" protobuf:"bytes,3,rep,name=labelSelector"`

	// If true, don't block pod startup if the referenced ClusterTrustBundle(s) // want "commentstart: godoc for field ClusterTrustBundleProjection.Optional should start with 'optional ...'"
	// aren't available.  If using name, then the named ClusterTrustBundle is
	// allowed not to exist.  If using signerName, then the combination of
	// signerName and labelSelector is allowed to match zero
	// ClusterTrustBundles.
	// +optional
	Optional *bool `json:"optional,omitempty" protobuf:"varint,5,opt,name=optional"`

	// Relative path from the volume root to write the bundle. // want "commentstart: godoc for field ClusterTrustBundleProjection.Path should start with 'path ...'"
	Path string `json:"path" protobuf:"bytes,4,rep,name=path"` // want "optionalorrequired: field ClusterTrustBundleProjection.Path must be marked as optional or required"
}

// PodCertificateProjection provides a private key and X.509 certificate in the
// pod filesystem.
type PodCertificateProjection struct {
	// Kubelet's generated CSRs will be addressed to this signer. // want "commentstart: godoc for field PodCertificateProjection.SignerName should start with 'signerName ...'"
	//
	// +required
	SignerName string `json:"signerName,omitempty" protobuf:"bytes,1,rep,name=signerName"` // want "requiredfields: field SignerName has a valid zero value \\(\\\"\\\"\\), but the validation is not complete \\(e.g. minimum length\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// The type of keypair Kubelet will generate for the pod. // want "commentstart: godoc for field PodCertificateProjection.KeyType should start with 'keyType ...'"
	//
	// Valid values are "RSA3072", "RSA4096", "ECDSAP256", "ECDSAP384",
	// "ECDSAP521", and "ED25519".
	//
	// +required
	KeyType string `json:"keyType,omitempty" protobuf:"bytes,2,rep,name=keyType"` // want "requiredfields: field KeyType has a valid zero value \\(\\\"\\\"\\), but the validation is not complete \\(e.g. minimum length\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."

	// maxExpirationSeconds is the maximum lifetime permitted for the
	// certificate.
	//
	// Kubelet copies this value verbatim into the PodCertificateRequests it
	// generates for this projection.
	//
	// If omitted, kube-apiserver will set it to 86400(24 hours). kube-apiserver
	// will reject values shorter than 3600 (1 hour).  The maximum allowable
	// value is 7862400 (91 days).
	//
	// The signer implementation is then free to issue a certificate with any
	// lifetime *shorter* than MaxExpirationSeconds, but no shorter than 3600
	// seconds (1 hour).  This constraint is enforced by kube-apiserver.
	// `kubernetes.io` signers will never issue certificates with a lifetime
	// longer than 24 hours.
	//
	// +optional
	MaxExpirationSeconds *int32 `json:"maxExpirationSeconds,omitempty" protobuf:"varint,3,opt,name=maxExpirationSeconds"`

	// Write the credential bundle at this path in the projected volume. // want "commentstart: godoc for field PodCertificateProjection.CredentialBundlePath should start with 'credentialBundlePath ...'"
	//
	// The credential bundle is a single file that contains multiple PEM blocks.
	// The first PEM block is a PRIVATE KEY block, containing a PKCS#8 private
	// key.
	//
	// The remaining blocks are CERTIFICATE blocks, containing the issued
	// certificate chain from the signer (leaf and any intermediates).
	//
	// Using credentialBundlePath lets your Pod's application code make a single
	// atomic read that retrieves a consistent key and certificate chain.  If you
	// project them to separate files, your application code will need to
	// additionally check that the leaf certificate was issued to the key.
	//
	// +optional
	CredentialBundlePath string `json:"credentialBundlePath,omitempty" protobuf:"bytes,4,rep,name=credentialBundlePath"` // want "optionalfields: field CredentialBundlePath should be a pointer."

	// Write the key at this path in the projected volume. // want "commentstart: godoc for field PodCertificateProjection.KeyPath should start with 'keyPath ...'"
	//
	// Most applications should use credentialBundlePath.  When using keyPath
	// and certificateChainPath, your application needs to check that the key
	// and leaf certificate are consistent, because it is possible to read the
	// files mid-rotation.
	//
	// +optional
	KeyPath string `json:"keyPath,omitempty" protobuf:"bytes,5,rep,name=keyPath"` // want "optionalfields: field KeyPath should be a pointer."

	// Write the certificate chain at this path in the projected volume. // want "commentstart: godoc for field PodCertificateProjection.CertificateChainPath should start with 'certificateChainPath ...'"
	//
	// Most applications should use credentialBundlePath.  When using keyPath
	// and certificateChainPath, your application needs to check that the key
	// and leaf certificate are consistent, because it is possible to read the
	// files mid-rotation.
	//
	// +optional
	CertificateChainPath string `json:"certificateChainPath,omitempty" protobuf:"bytes,6,rep,name=certificateChainPath"` // want "optionalfields: field CertificateChainPath should be a pointer."
}

// PersistentVolumeClaimSpec describes the common attributes of storage devices
// and allows a Source for provider-specific attributes
type PersistentVolumeClaimSpec struct {
	// accessModes contains the desired access modes the volume should have.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
	// +optional
	// +listType=atomic
	AccessModes []PersistentVolumeAccessMode `json:"accessModes,omitempty" protobuf:"bytes,1,rep,name=accessModes,casttype=PersistentVolumeAccessMode"`
	// selector is a label query over volumes to consider for binding.
	// +optional
	// Selector *metav1.LabelSelector `json:"selector,omitempty" protobuf:"bytes,4,opt,name=selector"`

	// resources represents the minimum resources the volume should have.
	// Users are allowed to specify resource requirements
	// that are lower than previous value but must still be higher than capacity recorded in the
	// status field of the claim.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources
	// +optional
	Resources VolumeResourceRequirements `json:"resources,omitempty" protobuf:"bytes,2,opt,name=resources"` // want "optionalfields: field Resources should be a pointer."
	// volumeName is the binding reference to the PersistentVolume backing this claim.
	// +optional
	VolumeName string `json:"volumeName,omitempty" protobuf:"bytes,3,opt,name=volumeName"` // want "optionalfields: field VolumeName should be a pointer."
	// storageClassName is the name of the StorageClass required by the claim.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1
	// +optional
	StorageClassName *string `json:"storageClassName,omitempty" protobuf:"bytes,5,opt,name=storageClassName"`
	// volumeMode defines what type of volume is required by the claim.
	// Value of Filesystem is implied when not included in claim spec.
	// +optional
	VolumeMode *PersistentVolumeMode `json:"volumeMode,omitempty" protobuf:"bytes,6,opt,name=volumeMode,casttype=PersistentVolumeMode"`
	// dataSource field can be used to specify either:
	// * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)
	// * An existing PVC (PersistentVolumeClaim)
	// If the provisioner or an external controller can support the specified data source,
	// it will create a new volume based on the contents of the specified data source.
	// When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,
	// and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.
	// If the namespace is specified, then dataSourceRef will not be copied to dataSource.
	// +optional
	DataSource *TypedLocalObjectReference `json:"dataSource,omitempty" protobuf:"bytes,7,opt,name=dataSource"`
	// dataSourceRef specifies the object from which to populate the volume with data, if a non-empty
	// volume is desired. This may be any object from a non-empty API group (non
	// core object) or a PersistentVolumeClaim object.
	// When this field is specified, volume binding will only succeed if the type of
	// the specified object matches some installed volume populator or dynamic
	// provisioner.
	// This field will replace the functionality of the dataSource field and as such
	// if both fields are non-empty, they must have the same value. For backwards
	// compatibility, when namespace isn't specified in dataSourceRef,
	// both fields (dataSource and dataSourceRef) will be set to the same
	// value automatically if one of them is empty and the other is non-empty.
	// When namespace is specified in dataSourceRef,
	// dataSource isn't set to the same value and must be empty.
	// There are three important differences between dataSource and dataSourceRef:
	// * While dataSource only allows two specific types of objects, dataSourceRef
	//   allows any non-core object, as well as PersistentVolumeClaim objects.
	// * While dataSource ignores disallowed values (dropping them), dataSourceRef
	//   preserves all values, and generates an error if a disallowed value is
	//   specified.
	// * While dataSource only allows local objects, dataSourceRef allows objects
	//   in any namespaces.
	// (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.
	// (Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.
	// +optional
	DataSourceRef *TypedObjectReference `json:"dataSourceRef,omitempty" protobuf:"bytes,8,opt,name=dataSourceRef"`
	// volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim.
	// If specified, the CSI driver will create or update the volume with the attributes defined
	// in the corresponding VolumeAttributesClass. This has a different purpose than storageClassName,
	// it can be changed after the claim is created. An empty string or nil value indicates that no
	// VolumeAttributesClass will be applied to the claim. If the claim enters an Infeasible error state,
	// this field can be reset to its previous value (including nil) to cancel the modification.
	// If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will be
	// set to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resource
	// exists.
	// More info: https://kubernetes.io/docs/concepts/storage/volume-attributes-classes/
	// +featureGate=VolumeAttributesClass
	// +optional
	VolumeAttributesClassName *string `json:"volumeAttributesClassName,omitempty" protobuf:"bytes,9,opt,name=volumeAttributesClassName"`
}

// +enum
type PersistentVolumeAccessMode string

const (
	// can be mounted in read/write mode to exactly 1 host
	ReadWriteOnce PersistentVolumeAccessMode = "ReadWriteOnce"
	// can be mounted in read-only mode to many hosts
	ReadOnlyMany PersistentVolumeAccessMode = "ReadOnlyMany"
	// can be mounted in read/write mode to many hosts
	ReadWriteMany PersistentVolumeAccessMode = "ReadWriteMany"
	// can be mounted in read/write mode to exactly 1 pod
	// cannot be used in combination with other access modes
	ReadWriteOncePod PersistentVolumeAccessMode = "ReadWriteOncePod"
)

// PersistentVolumeMode describes how a volume is intended to be consumed, either Block or Filesystem.
// +enum
type PersistentVolumeMode string

const (
	// PersistentVolumeBlock means the volume will not be formatted with a filesystem and will remain a raw block device.
	PersistentVolumeBlock PersistentVolumeMode = "Block"
	// PersistentVolumeFilesystem means the volume will be or is formatted with a filesystem.
	PersistentVolumeFilesystem PersistentVolumeMode = "Filesystem"
)

// VolumeResourceRequirements describes the storage resource requirements for a volume.
type VolumeResourceRequirements struct {
	// Limits describes the maximum amount of compute resources allowed.
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	// +optional
	// Limits ResourceList `json:"limits,omitempty" protobuf:"bytes,1,rep,name=limits,casttype=ResourceList,castkey=ResourceName"`

	// Requests describes the minimum amount of compute resources required.
	// If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,
	// otherwise to an implementation-defined value. Requests cannot exceed Limits.
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	// +optional
	// Requests ResourceList `json:"requests,omitempty" protobuf:"bytes,2,rep,name=requests,casttype=ResourceList,castkey=ResourceName"`

	// Claims got added by accident when volumes shared the ResourceRequirements struct
	// with containers. Stripping the field got added in 1.27 and was backported to 1.26.
	// Starting with Kubernetes 1.28, this field is not part of the volume API anymore.
	//
	// Future extensions must not use "claims" or field number 3.
	// Claims []ResourceClaim `json:"claims,omitempty" protobuf:"bytes,3,opt,name=claims"`
}

// Local represents directly-attached storage with node affinity
type LocalVolumeSource struct {
	// path of the full path to the volume on the node.
	// It can be either a directory or block device (disk, partition, ...).
	Path string `json:"path" protobuf:"bytes,1,opt,name=path"` // want "optionalorrequired: field LocalVolumeSource.Path must be marked as optional or required"

	// fsType is the filesystem type to mount.
	// It applies only when the Path is a block device.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs". The default value is to auto-select a filesystem if unspecified.
	// +optional
	FSType *string `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType"`
}

// Represents storage that is managed by an external CSI volume driver
type CSIPersistentVolumeSource struct {
	// driver is the name of the driver to use for this volume.
	// Required.
	Driver string `json:"driver" protobuf:"bytes,1,opt,name=driver"` // want "optionalorrequired: field CSIPersistentVolumeSource.Driver must be marked as optional or required"

	// volumeHandle is the unique volume name returned by the CSI volume
	// plugin’s CreateVolume to refer to the volume on all subsequent calls.
	// Required.
	VolumeHandle string `json:"volumeHandle" protobuf:"bytes,2,opt,name=volumeHandle"` // want "optionalorrequired: field CSIPersistentVolumeSource.VolumeHandle must be marked as optional or required"

	// readOnly value to pass to ControllerPublishVolumeRequest.
	// Defaults to false (read/write).
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly"` // want "optionalfields: field ReadOnly should be a pointer."

	// fsType to mount. Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs".
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,4,opt,name=fsType"` // want "optionalfields: field FSType should be a pointer."

	// volumeAttributes of the volume to publish.
	// +optional
	VolumeAttributes map[string]string `json:"volumeAttributes,omitempty" protobuf:"bytes,5,rep,name=volumeAttributes"`

	// controllerPublishSecretRef is a reference to the secret object containing
	// sensitive information to pass to the CSI driver to complete the CSI
	// ControllerPublishVolume and ControllerUnpublishVolume calls.
	// This field is optional, and may be empty if no secret is required. If the
	// secret object contains more than one secret, all secrets are passed.
	// +optional
	ControllerPublishSecretRef *SecretReference `json:"controllerPublishSecretRef,omitempty" protobuf:"bytes,6,opt,name=controllerPublishSecretRef"`

	// nodeStageSecretRef is a reference to the secret object containing sensitive
	// information to pass to the CSI driver to complete the CSI NodeStageVolume
	// and NodeStageVolume and NodeUnstageVolume calls.
	// This field is optional, and may be empty if no secret is required. If the
	// secret object contains more than one secret, all secrets are passed.
	// +optional
	NodeStageSecretRef *SecretReference `json:"nodeStageSecretRef,omitempty" protobuf:"bytes,7,opt,name=nodeStageSecretRef"`

	// nodePublishSecretRef is a reference to the secret object containing
	// sensitive information to pass to the CSI driver to complete the CSI
	// NodePublishVolume and NodeUnpublishVolume calls.
	// This field is optional, and may be empty if no secret is required. If the
	// secret object contains more than one secret, all secrets are passed.
	// +optional
	NodePublishSecretRef *SecretReference `json:"nodePublishSecretRef,omitempty" protobuf:"bytes,8,opt,name=nodePublishSecretRef"`

	// controllerExpandSecretRef is a reference to the secret object containing
	// sensitive information to pass to the CSI driver to complete the CSI
	// ControllerExpandVolume call.
	// This field is optional, and may be empty if no secret is required. If the
	// secret object contains more than one secret, all secrets are passed.
	// +optional
	ControllerExpandSecretRef *SecretReference `json:"controllerExpandSecretRef,omitempty" protobuf:"bytes,9,opt,name=controllerExpandSecretRef"`

	// nodeExpandSecretRef is a reference to the secret object containing
	// sensitive information to pass to the CSI driver to complete the CSI
	// NodeExpandVolume call.
	// This field is optional, may be omitted if no secret is required. If the
	// secret object contains more than one secret, all secrets are passed.
	// +optional
	NodeExpandSecretRef *SecretReference `json:"nodeExpandSecretRef,omitempty" protobuf:"bytes,10,opt,name=nodeExpandSecretRef"`
}
