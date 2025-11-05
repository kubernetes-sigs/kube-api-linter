package defaultconfigurations

// An EphemeralContainer is a temporary container that you may add to an existing Pod for
// user-initiated activities such as debugging. Ephemeral containers have no resource or
// scheduling guarantees, and they will not be restarted when they exit or when a Pod is
// removed or restarted. The kubelet may evict a Pod if an ephemeral container causes the
// Pod to exceed its resource allocation.
//
// To add an ephemeral container, use the ephemeralcontainers subresource of an existing
// Pod. Ephemeral containers may not be removed or restarted.
type EphemeralContainer struct {
	// Ephemeral containers have all of the fields of Container, plus additional fields
	// specific to ephemeral containers. Fields in common with Container are in the
	// following inlined struct so than an EphemeralContainer may easily be converted
	// to a Container.
	EphemeralContainerCommon `json:",inline" protobuf:"bytes,1,req"`

	// If set, the name of the container from PodSpec that this ephemeral container targets. // want "commentstart: godoc for field EphemeralContainer.TargetContainerName should start with 'targetContainerName ...'"
	// The ephemeral container will be run in the namespaces (IPC, PID, etc) of this container.
	// If not set then the ephemeral container uses the namespaces configured in the Pod spec.
	//
	// The container runtime must implement support for this feature. If the runtime does not
	// support namespace targeting then the result of setting this field is undefined.
	// +optional
	TargetContainerName string `json:"targetContainerName,omitempty" protobuf:"bytes,2,opt,name=targetContainerName"` // want "optionalfields: field TargetContainerName should be a pointer."
}

// EphemeralContainerCommon is a copy of all fields in Container to be inlined in
// EphemeralContainer. This separate type allows easy conversion from EphemeralContainer
// to Container and allows separate documentation for the fields of EphemeralContainer.
// When a new field is added to Container it must be added here as well.
type EphemeralContainerCommon struct {
	// Name of the ephemeral container specified as a DNS_LABEL. // want "commentstart: godoc for field EphemeralContainerCommon.Name should start with 'name ...'"
	// This name must be unique among all containers, init containers and ephemeral containers.
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"` // want "optionalorrequired: field EphemeralContainerCommon.Name must be marked as optional or required"
	// Container image name. // want "commentstart: godoc for field EphemeralContainerCommon.Image should start with 'image ...'"
	// More info: https://kubernetes.io/docs/concepts/containers/images
	Image string `json:"image,omitempty" protobuf:"bytes,2,opt,name=image"` // want "optionalorrequired: field EphemeralContainerCommon.Image must be marked as optional or required"
	// Entrypoint array. Not executed within a shell. // want "commentstart: godoc for field EphemeralContainerCommon.Command should start with 'command ...'"
	// The image's ENTRYPOINT is used if this is not provided.
	// Variable references $(VAR_NAME) are expanded using the container's environment. If a variable
	// cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced
	// to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will
	// produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless
	// of whether the variable exists or not. Cannot be updated.
	// More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	// +optional
	// +listType=atomic
	Command []string `json:"command,omitempty" protobuf:"bytes,3,rep,name=command"`
	// Arguments to the entrypoint. // want "commentstart: godoc for field EphemeralContainerCommon.Args should start with 'args ...'"
	// The image's CMD is used if this is not provided.
	// Variable references $(VAR_NAME) are expanded using the container's environment. If a variable
	// cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced
	// to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will
	// produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless
	// of whether the variable exists or not. Cannot be updated.
	// More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	// +optional
	// +listType=atomic
	Args []string `json:"args,omitempty" protobuf:"bytes,4,rep,name=args"`
	// Container's working directory. // want "commentstart: godoc for field EphemeralContainerCommon.WorkingDir should start with 'workingDir ...'"
	// If not specified, the container runtime's default will be used, which
	// might be configured in the container image.
	// Cannot be updated.
	// +optional
	WorkingDir string `json:"workingDir,omitempty" protobuf:"bytes,5,opt,name=workingDir"` // want "optionalfields: field WorkingDir should be a pointer."
	// Ports are not allowed for ephemeral containers. // want "commentstart: godoc for field EphemeralContainerCommon.Ports should start with 'ports ...'"
	// +optional
	// +patchMergeKey=containerPort
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=containerPort
	// +listMapKey=protocol
	Ports []ContainerPort `json:"ports,omitempty" patchStrategy:"merge" patchMergeKey:"containerPort" protobuf:"bytes,6,rep,name=ports"` // want "arrayofstruct: EphemeralContainerCommon.Ports is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// List of sources to populate environment variables in the container. // want "commentstart: godoc for field EphemeralContainerCommon.EnvFrom should start with 'envFrom ...'"
	// The keys defined within a source may consist of any printable ASCII characters except '='.
	// When a key exists in multiple
	// sources, the value associated with the last source will take precedence.
	// Values defined by an Env with a duplicate key will take precedence.
	// Cannot be updated.
	// +optional
	// +listType=atomic
	EnvFrom []EnvFromSource `json:"envFrom,omitempty" protobuf:"bytes,19,rep,name=envFrom"` // want "arrayofstruct: EphemeralContainerCommon.EnvFrom is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// List of environment variables to set in the container. // want "commentstart: godoc for field EphemeralContainerCommon.Env should start with 'env ...'"
	// Cannot be updated.
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=name
	Env []EnvVar `json:"env,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,7,rep,name=env"` // want "arrayofstruct: EphemeralContainerCommon.Env is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources // want "commentstart: godoc for field EphemeralContainerCommon.Resources should start with 'resources ...'"
	// already allocated to the pod.
	// +optional
	Resources ResourceRequirements `json:"resources,omitempty" protobuf:"bytes,8,opt,name=resources"` // want "optionalfields: field Resources should be a pointer."
	// Resources resize policy for the container. // want "commentstart: godoc for field EphemeralContainerCommon.ResizePolicy should start with 'resizePolicy ...'"
	// +featureGate=InPlacePodVerticalScaling
	// +optional
	// +listType=atomic
	ResizePolicy []ContainerResizePolicy `json:"resizePolicy,omitempty" protobuf:"bytes,23,rep,name=resizePolicy"` // want "arrayofstruct: EphemeralContainerCommon.ResizePolicy is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// Restart policy for the container to manage the restart behavior of each // want "commentstart: godoc for field EphemeralContainerCommon.RestartPolicy should start with 'restartPolicy ...'"
	// container within a pod.
	// You cannot set this field on ephemeral containers.
	// +optional
	RestartPolicy *ContainerRestartPolicy `json:"restartPolicy,omitempty" protobuf:"bytes,24,opt,name=restartPolicy,casttype=ContainerRestartPolicy"`
	// Represents a list of rules to be checked to determine if the // want "commentstart: godoc for field EphemeralContainerCommon.RestartPolicyRules should start with 'restartPolicyRules ...'"
	// container should be restarted on exit. You cannot set this field on
	// ephemeral containers.
	// +featureGate=ContainerRestartRules
	// +optional
	// +listType=atomic
	RestartPolicyRules []ContainerRestartRule `json:"restartPolicyRules,omitempty" protobuf:"bytes,25,rep,name=restartPolicyRules"`
	// Pod volumes to mount into the container's filesystem. Subpath mounts are not allowed for ephemeral containers. // want "commentstart: godoc for field EphemeralContainerCommon.VolumeMounts should start with 'volumeMounts ...'"
	// Cannot be updated.
	// +optional
	// +patchMergeKey=mountPath
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=mountPath
	VolumeMounts []VolumeMount `json:"volumeMounts,omitempty" patchStrategy:"merge" patchMergeKey:"mountPath" protobuf:"bytes,9,rep,name=volumeMounts"` // want "arrayofstruct: EphemeralContainerCommon.VolumeMounts is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// volumeDevices is the list of block devices to be used by the container.
	// +patchMergeKey=devicePath
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=devicePath
	// +optional
	VolumeDevices []VolumeDevice `json:"volumeDevices,omitempty" patchStrategy:"merge" patchMergeKey:"devicePath" protobuf:"bytes,21,rep,name=volumeDevices"` // want "arrayofstruct: EphemeralContainerCommon.VolumeDevices is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// Probes are not allowed for ephemeral containers. // want "commentstart: godoc for field EphemeralContainerCommon.LivenessProbe should start with 'livenessProbe ...'"
	// +optional
	LivenessProbe *Probe `json:"livenessProbe,omitempty" protobuf:"bytes,10,opt,name=livenessProbe"`
	// Probes are not allowed for ephemeral containers. // want "commentstart: godoc for field EphemeralContainerCommon.ReadinessProbe should start with 'readinessProbe ...'"
	// +optional
	ReadinessProbe *Probe `json:"readinessProbe,omitempty" protobuf:"bytes,11,opt,name=readinessProbe"`
	// Probes are not allowed for ephemeral containers. // want "commentstart: godoc for field EphemeralContainerCommon.StartupProbe should start with 'startupProbe ...'"
	// +optional
	StartupProbe *Probe `json:"startupProbe,omitempty" protobuf:"bytes,22,opt,name=startupProbe"`
	// Lifecycle is not allowed for ephemeral containers. // want "commentstart: godoc for field EphemeralContainerCommon.Lifecycle should start with 'lifecycle ...'"
	// +optional
	Lifecycle *Lifecycle `json:"lifecycle,omitempty" protobuf:"bytes,12,opt,name=lifecycle"`
	// Optional: Path at which the file to which the container's termination message // want "commentstart: godoc for field EphemeralContainerCommon.TerminationMessagePath should start with 'terminationMessagePath ...'"
	// will be written is mounted into the container's filesystem.
	// Message written is intended to be brief final status, such as an assertion failure message.
	// Will be truncated by the node if greater than 4096 bytes. The total message length across
	// all containers will be limited to 12kb.
	// Defaults to /dev/termination-log.
	// Cannot be updated.
	// +optional
	TerminationMessagePath string `json:"terminationMessagePath,omitempty" protobuf:"bytes,13,opt,name=terminationMessagePath"` // want "optionalfields: field TerminationMessagePath should be a pointer."
	// Indicate how the termination message should be populated. File will use the contents of // want "commentstart: godoc for field EphemeralContainerCommon.TerminationMessagePolicy should start with 'terminationMessagePolicy ...'"
	// terminationMessagePath to populate the container status message on both success and failure.
	// FallbackToLogsOnError will use the last chunk of container log output if the termination
	// message file is empty and the container exited with an error.
	// The log output is limited to 2048 bytes or 80 lines, whichever is smaller.
	// Defaults to File.
	// Cannot be updated.
	// +optional
	TerminationMessagePolicy TerminationMessagePolicy `json:"terminationMessagePolicy,omitempty" protobuf:"bytes,20,opt,name=terminationMessagePolicy,casttype=TerminationMessagePolicy"` // want "optionalfields: field TerminationMessagePolicy should be a pointer."
	// Image pull policy. // want "commentstart: godoc for field EphemeralContainerCommon.ImagePullPolicy should start with 'imagePullPolicy ...'"
	// One of Always, Never, IfNotPresent.
	// Defaults to Always if :latest tag is specified, or IfNotPresent otherwise.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/containers/images#updating-images
	// +optional
	ImagePullPolicy PullPolicy `json:"imagePullPolicy,omitempty" protobuf:"bytes,14,opt,name=imagePullPolicy,casttype=PullPolicy"` // want "optionalfields: field ImagePullPolicy should be a pointer."
	// Optional: SecurityContext defines the security options the ephemeral container should be run with. // want "commentstart: godoc for field EphemeralContainerCommon.SecurityContext should start with 'securityContext ...'"
	// If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext.
	// +optional
	SecurityContext *SecurityContext `json:"securityContext,omitempty" protobuf:"bytes,15,opt,name=securityContext"`

	// Variables for interactive containers, these have very specialized use-cases (e.g. debugging)
	// and shouldn't be used for general purpose containers.

	// Whether this container should allocate a buffer for stdin in the container runtime. If this // want "commentstart: godoc for field EphemeralContainerCommon.Stdin should start with 'stdin ...'"
	// is not set, reads from stdin in the container will always result in EOF.
	// Default is false.
	// +optional
	Stdin bool `json:"stdin,omitempty" protobuf:"varint,16,opt,name=stdin"` // want "optionalfields: field Stdin should be a pointer."
	// Whether the container runtime should close the stdin channel after it has been opened by // want "commentstart: godoc for field EphemeralContainerCommon.StdinOnce should start with 'stdinOnce ...'"
	// a single attach. When stdin is true the stdin stream will remain open across multiple attach
	// sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the
	// first client attaches to stdin, and then remains open and accepts data until the client disconnects,
	// at which time stdin is closed and remains closed until the container is restarted. If this
	// flag is false, a container processes that reads from stdin will never receive an EOF.
	// Default is false
	// +optional
	StdinOnce bool `json:"stdinOnce,omitempty" protobuf:"varint,17,opt,name=stdinOnce"` // want "optionalfields: field StdinOnce should be a pointer."
	// Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. // want "commentstart: godoc for field EphemeralContainerCommon.TTY should start with 'tty ...'"
	// Default is false.
	// +optional
	TTY bool `json:"tty,omitempty" protobuf:"varint,18,opt,name=tty"` // want "optionalfields: field TTY should be a pointer."
}
