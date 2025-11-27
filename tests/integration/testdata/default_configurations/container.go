package defaultconfigurations

// A single application container that you want to run within a pod.
type Container struct {
	// Name of the container specified as a DNS_LABEL. // want "commentstart: godoc for field Container.Name should start with 'name ...'"
	// Each container in a pod must have a unique name (DNS_LABEL).
	// Cannot be updated.
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"` // want "optionalorrequired: field Container.Name must be marked as optional or required"
	// Container image name. // want "commentstart: godoc for field Container.Image should start with 'image ...'"
	// More info: https://kubernetes.io/docs/concepts/containers/images
	// This field is optional to allow higher level config management to default or override
	// container images in workload controllers like Deployments and StatefulSets.
	// +optional
	Image string `json:"image,omitempty" protobuf:"bytes,2,opt,name=image"` // want "optionalfields: field Image should be a pointer."
	// Entrypoint array. Not executed within a shell. // want "commentstart: godoc for field Container.Command should start with 'command ...'"
	// The container image's ENTRYPOINT is used if this is not provided.
	// Variable references $(VAR_NAME) are expanded using the container's environment. If a variable
	// cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced
	// to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will
	// produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless
	// of whether the variable exists or not. Cannot be updated.
	// More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	// +optional
	// +listType=atomic
	Command []string `json:"command,omitempty" protobuf:"bytes,3,rep,name=command"`
	// Arguments to the entrypoint. // want "commentstart: godoc for field Container.Args should start with 'args ...'"
	// The container image's CMD is used if this is not provided.
	// Variable references $(VAR_NAME) are expanded using the container's environment. If a variable
	// cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced
	// to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will
	// produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless
	// of whether the variable exists or not. Cannot be updated.
	// More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	// +optional
	// +listType=atomic
	Args []string `json:"args,omitempty" protobuf:"bytes,4,rep,name=args"`
	// Container's working directory. // want "commentstart: godoc for field Container.WorkingDir should start with 'workingDir ...'"
	// If not specified, the container runtime's default will be used, which
	// might be configured in the container image.
	// Cannot be updated.
	// +optional
	WorkingDir string `json:"workingDir,omitempty" protobuf:"bytes,5,opt,name=workingDir"` // want "optionalfields: field WorkingDir should be a pointer."
	// List of ports to expose from the container. Not specifying a port here // want "commentstart: godoc for field Container.Ports should start with 'ports ...'"
	// DOES NOT prevent that port from being exposed. Any port which is
	// listening on the default "0.0.0.0" address inside a container will be
	// accessible from the network.
	// Modifying this array with strategic merge patch may corrupt the data.
	// For more information See https://github.com/kubernetes/kubernetes/issues/108255.
	// Cannot be updated.
	// +optional
	// +patchMergeKey=containerPort
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=containerPort
	// +listMapKey=protocol
	Ports []ContainerPort `json:"ports,omitempty" patchStrategy:"merge" patchMergeKey:"containerPort" protobuf:"bytes,6,rep,name=ports"` // want "arrayofstruct: Container.Ports is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// List of sources to populate environment variables in the container. // want "commentstart: godoc for field Container.EnvFrom should start with 'envFrom ...'"
	// The keys defined within a source may consist of any printable ASCII characters except '='.
	// When a key exists in multiple
	// sources, the value associated with the last source will take precedence.
	// Values defined by an Env with a duplicate key will take precedence.
	// Cannot be updated.
	// +optional
	// +listType=atomic
	EnvFrom []EnvFromSource `json:"envFrom,omitempty" protobuf:"bytes,19,rep,name=envFrom"` // want "arrayofstruct: Container.EnvFrom is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// List of environment variables to set in the container. // want "commentstart: godoc for field Container.Env should start with 'env ...'"
	// Cannot be updated.
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=name
	Env []EnvVar `json:"env,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,7,rep,name=env"` // want "arrayofstruct: Container.Env is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// Compute Resources required by this container. // want "commentstart: godoc for field Container.Resources should start with 'resources ...'"
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	// +optional
	Resources ResourceRequirements `json:"resources,omitempty" protobuf:"bytes,8,opt,name=resources"` // want "optionalfields: field Resources should be a pointer."
	// Resources resize policy for the container. // want "commentstart: godoc for field Container.ResizePolicy should start with 'resizePolicy ...'"
	// This field cannot be set on ephemeral containers.
	// +featureGate=InPlacePodVerticalScaling
	// +optional
	// +listType=atomic
	ResizePolicy []ContainerResizePolicy `json:"resizePolicy,omitempty" protobuf:"bytes,23,rep,name=resizePolicy"` // want "arrayofstruct: Container.ResizePolicy is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// RestartPolicy defines the restart behavior of individual containers in a pod. // want "commentstart: godoc for field Container.RestartPolicy should start with 'restartPolicy ...'"
	// This overrides the pod-level restart policy. When this field is not specified,
	// the restart behavior is defined by the Pod's restart policy and the container type.
	// Additionally, setting the RestartPolicy as "Always" for the init container will
	// have the following effect:
	// this init container will be continually restarted on
	// exit until all regular containers have terminated. Once all regular
	// containers have completed, all init containers with restartPolicy "Always"
	// will be shut down. This lifecycle differs from normal init containers and
	// is often referred to as a "sidecar" container. Although this init
	// container still starts in the init container sequence, it does not wait
	// for the container to complete before proceeding to the next init
	// container. Instead, the next init container starts immediately after this
	// init container is started, or after any startupProbe has successfully
	// completed.
	// +optional
	RestartPolicy *ContainerRestartPolicy `json:"restartPolicy,omitempty" protobuf:"bytes,24,opt,name=restartPolicy,casttype=ContainerRestartPolicy"`
	// Represents a list of rules to be checked to determine if the // want "commentstart: godoc for field Container.RestartPolicyRules should start with 'restartPolicyRules ...'"
	// container should be restarted on exit. The rules are evaluated in
	// order. Once a rule matches a container exit condition, the remaining
	// rules are ignored. If no rule matches the container exit condition,
	// the Container-level restart policy determines the whether the container
	// is restarted or not. Constraints on the rules:
	// - At most 20 rules are allowed.
	// - Rules can have the same action.
	// - Identical rules are not forbidden in validations.
	// When rules are specified, container MUST set RestartPolicy explicitly
	// even it if matches the Pod's RestartPolicy.
	// +featureGate=ContainerRestartRules
	// +optional
	// +listType=atomic
	RestartPolicyRules []ContainerRestartRule `json:"restartPolicyRules,omitempty" protobuf:"bytes,25,rep,name=restartPolicyRules"`
	// Pod volumes to mount into the container's filesystem. // want "commentstart: godoc for field Container.VolumeMounts should start with 'volumeMounts ...'"
	// Cannot be updated.
	// +optional
	// +patchMergeKey=mountPath
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=mountPath
	VolumeMounts []VolumeMount `json:"volumeMounts,omitempty" patchStrategy:"merge" patchMergeKey:"mountPath" protobuf:"bytes,9,rep,name=volumeMounts"` // want "arrayofstruct: Container.VolumeMounts is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// volumeDevices is the list of block devices to be used by the container.
	// +patchMergeKey=devicePath
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=devicePath
	// +optional
	VolumeDevices []VolumeDevice `json:"volumeDevices,omitempty" patchStrategy:"merge" patchMergeKey:"devicePath" protobuf:"bytes,21,rep,name=volumeDevices"` // want "arrayofstruct: Container.VolumeDevices is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// Periodic probe of container liveness. // want "commentstart: godoc for field Container.LivenessProbe should start with 'livenessProbe ...'"
	// Container will be restarted if the probe fails.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	LivenessProbe *Probe `json:"livenessProbe,omitempty" protobuf:"bytes,10,opt,name=livenessProbe"`
	// Periodic probe of container service readiness. // want "commentstart: godoc for field Container.ReadinessProbe should start with 'readinessProbe ...'"
	// Container will be removed from service endpoints if the probe fails.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	ReadinessProbe *Probe `json:"readinessProbe,omitempty" protobuf:"bytes,11,opt,name=readinessProbe"`
	// StartupProbe indicates that the Pod has successfully initialized. // want "commentstart: godoc for field Container.StartupProbe should start with 'startupProbe ...'"
	// If specified, no other probes are executed until this completes successfully.
	// If this probe fails, the Pod will be restarted, just as if the livenessProbe failed.
	// This can be used to provide different probe parameters at the beginning of a Pod's lifecycle,
	// when it might take a long time to load data or warm a cache, than during steady-state operation.
	// This cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	StartupProbe *Probe `json:"startupProbe,omitempty" protobuf:"bytes,22,opt,name=startupProbe"`
	// Actions that the management system should take in response to container lifecycle events. // want "commentstart: godoc for field Container.Lifecycle should start with 'lifecycle ...'"
	// Cannot be updated.
	// +optional
	Lifecycle *Lifecycle `json:"lifecycle,omitempty" protobuf:"bytes,12,opt,name=lifecycle"`
	// Optional: Path at which the file to which the container's termination message // want "commentstart: godoc for field Container.TerminationMessagePath should start with 'terminationMessagePath ...'"
	// will be written is mounted into the container's filesystem.
	// Message written is intended to be brief final status, such as an assertion failure message.
	// Will be truncated by the node if greater than 4096 bytes. The total message length across
	// all containers will be limited to 12kb.
	// Defaults to /dev/termination-log.
	// Cannot be updated.
	// +optional
	TerminationMessagePath string `json:"terminationMessagePath,omitempty" protobuf:"bytes,13,opt,name=terminationMessagePath"` // want "optionalfields: field TerminationMessagePath should be a pointer."
	// Indicate how the termination message should be populated. File will use the contents of // want "commentstart: godoc for field Container.TerminationMessagePolicy should start with 'terminationMessagePolicy ...'"
	// terminationMessagePath to populate the container status message on both success and failure.
	// FallbackToLogsOnError will use the last chunk of container log output if the termination
	// message file is empty and the container exited with an error.
	// The log output is limited to 2048 bytes or 80 lines, whichever is smaller.
	// Defaults to File.
	// Cannot be updated.
	// +optional
	TerminationMessagePolicy TerminationMessagePolicy `json:"terminationMessagePolicy,omitempty" protobuf:"bytes,20,opt,name=terminationMessagePolicy,casttype=TerminationMessagePolicy"` // want "optionalfields: field TerminationMessagePolicy should be a pointer."
	// Image pull policy. // want "commentstart: godoc for field Container.ImagePullPolicy should start with 'imagePullPolicy ...'"
	// One of Always, Never, IfNotPresent.
	// Defaults to Always if :latest tag is specified, or IfNotPresent otherwise.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/containers/images#updating-images
	// +optional
	ImagePullPolicy PullPolicy `json:"imagePullPolicy,omitempty" protobuf:"bytes,14,opt,name=imagePullPolicy,casttype=PullPolicy"` // want "optionalfields: field ImagePullPolicy should be a pointer."
	// SecurityContext defines the security options the container should be run with. // want "commentstart: godoc for field Container.SecurityContext should start with 'securityContext ...'"
	// If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext.
	// More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	// +optional
	SecurityContext *SecurityContext `json:"securityContext,omitempty" protobuf:"bytes,15,opt,name=securityContext"`

	// Variables for interactive containers, these have very specialized use-cases (e.g. debugging)
	// and shouldn't be used for general purpose containers.

	// Whether this container should allocate a buffer for stdin in the container runtime. If this // want "commentstart: godoc for field Container.Stdin should start with 'stdin ...'"
	// is not set, reads from stdin in the container will always result in EOF.
	// Default is false.
	// +optional
	Stdin bool `json:"stdin,omitempty" protobuf:"varint,16,opt,name=stdin"` // want "optionalfields: field Stdin should be a pointer."
	// Whether the container runtime should close the stdin channel after it has been opened by // want "commentstart: godoc for field Container.StdinOnce should start with 'stdinOnce ...'"
	// a single attach. When stdin is true the stdin stream will remain open across multiple attach
	// sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the
	// first client attaches to stdin, and then remains open and accepts data until the client disconnects,
	// at which time stdin is closed and remains closed until the container is restarted. If this
	// flag is false, a container processes that reads from stdin will never receive an EOF.
	// Default is false
	// +optional
	StdinOnce bool `json:"stdinOnce,omitempty" protobuf:"varint,17,opt,name=stdinOnce"` // want "optionalfields: field StdinOnce should be a pointer."
	// Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. // want "commentstart: godoc for field Container.TTY should start with 'tty ...'"
	// Default is false.
	// +optional
	TTY bool `json:"tty,omitempty" protobuf:"varint,18,opt,name=tty"` // want "optionalfields: field TTY should be a pointer."
}

// Protocol defines network protocols supported for things like container ports.
// +kubebuilder:validation:Enum
type Protocol string

const (
	// ProtocolTCP is the TCP protocol.
	ProtocolTCP Protocol = "TCP"
	// ProtocolUDP is the UDP protocol.
	ProtocolUDP Protocol = "UDP"
	// ProtocolSCTP is the SCTP protocol.
	ProtocolSCTP Protocol = "SCTP"
)

// ContainerPort represents a network port in a single container.
type ContainerPort struct {
	// If specified, this must be an IANA_SVC_NAME and unique within the pod. Each // want "commentstart: godoc for field ContainerPort.Name should start with 'name ...'"
	// named port in a pod must have a unique name. Name for the port that can be
	// referred to by services.
	// +optional
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"` // want "optionalfields: field Name should be a pointer."
	// Number of port to expose on the host. // want "commentstart: godoc for field ContainerPort.HostPort should start with 'hostPort ...'"
	// If specified, this must be a valid port number, 0 < x < 65536.
	// If HostNetwork is specified, this must match ContainerPort.
	// Most containers do not need this.
	// +optional
	HostPort int32 `json:"hostPort,omitempty" protobuf:"varint,2,opt,name=hostPort"` // want "optionalfields: field HostPort should be a pointer."
	// Number of port to expose on the pod's IP address. // want "commentstart: godoc for field ContainerPort.ContainerPort should start with 'containerPort ...'"
	// This must be a valid port number, 0 < x < 65536.
	ContainerPort int32 `json:"containerPort" protobuf:"varint,3,opt,name=containerPort"` // want "optionalorrequired: field ContainerPort.ContainerPort must be marked as optional or required"
	// Protocol for port. Must be UDP, TCP, or SCTP. // want "commentstart: godoc for field ContainerPort.Protocol should start with 'protocol ...'"
	// Defaults to "TCP".
	// +optional
	// +default="TCP"
	Protocol Protocol `json:"protocol,omitempty" protobuf:"bytes,4,opt,name=protocol,casttype=Protocol"` // want "optionalfields: field Protocol should be a pointer."
	// What host IP to bind the external port to. // want "commentstart: godoc for field ContainerPort.HostIP should start with 'hostIP ...'"
	// +optional
	HostIP string `json:"hostIP,omitempty" protobuf:"bytes,5,opt,name=hostIP"` // want "optionalfields: field HostIP should be a pointer."
}

// EnvVar represents an environment variable present in a Container.
type EnvVar struct {
	// Name of the environment variable. // want "commentstart: godoc for field EnvVar.Name should start with 'name ...'"
	// May consist of any printable ASCII characters except '='.
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"` // want "optionalorrequired: field EnvVar.Name must be marked as optional or required"

	// Optional: no more than one of the following may be specified.

	// Variable references $(VAR_NAME) are expanded // want "commentstart: godoc for field EnvVar.Value should start with 'value ...'"
	// using the previously defined environment variables in the container and
	// any service environment variables. If a variable cannot be resolved,
	// the reference in the input string will be unchanged. Double $$ are reduced
	// to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.
	// "$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)".
	// Escaped references will never be expanded, regardless of whether the variable
	// exists or not.
	// Defaults to "".
	// +optional
	Value string `json:"value,omitempty" protobuf:"bytes,2,opt,name=value"` // want "optionalfields: field Value should be a pointer."
	// Source for the environment variable's value. Cannot be used if value is not empty. // want "commentstart: godoc for field EnvVar.ValueFrom should start with 'valueFrom ...'"
	// +optional
	ValueFrom *EnvVarSource `json:"valueFrom,omitempty" protobuf:"bytes,3,opt,name=valueFrom"`
}

// EnvVarSource represents a source for the value of an EnvVar.
type EnvVarSource struct {
	// Selects a field of the pod: supports metadata.name, metadata.namespace, `metadata.labels['<KEY>']`, `metadata.annotations['<KEY>']`, // want "commentstart: godoc for field EnvVarSource.FieldRef should start with 'fieldRef ...'"
	// spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.
	// +optional
	FieldRef *ObjectFieldSelector `json:"fieldRef,omitempty" protobuf:"bytes,1,opt,name=fieldRef"`
	// Selects a resource of the container: only resources limits and requests // want "commentstart: godoc for field EnvVarSource.ResourceFieldRef should start with 'resourceFieldRef ...'"
	// (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.
	// +optional
	ResourceFieldRef *ResourceFieldSelector `json:"resourceFieldRef,omitempty" protobuf:"bytes,2,opt,name=resourceFieldRef"`
	// Selects a key of a ConfigMap. // want "commentstart: godoc for field EnvVarSource.ConfigMapKeyRef should start with 'configMapKeyRef ...'"
	// +optional
	ConfigMapKeyRef *ConfigMapKeySelector `json:"configMapKeyRef,omitempty" protobuf:"bytes,3,opt,name=configMapKeyRef"`
	// Selects a key of a secret in the pod's namespace // want "commentstart: godoc for field EnvVarSource.SecretKeyRef should start with 'secretKeyRef ...'"
	// +optional
	SecretKeyRef *SecretKeySelector `json:"secretKeyRef,omitempty" protobuf:"bytes,4,opt,name=secretKeyRef"`
	// FileKeyRef selects a key of the env file. // want "commentstart: godoc for field EnvVarSource.FileKeyRef should start with 'fileKeyRef ...'"
	// Requires the EnvFiles feature gate to be enabled.
	//
	// +featureGate=EnvFiles
	// +optional
	FileKeyRef *FileKeySelector `json:"fileKeyRef,omitempty" protobuf:"bytes,5,opt,name=fileKeyRef"`
}

// FileKeySelector selects a key of the env file.
// +structType=atomic
type FileKeySelector struct {
	// The name of the volume mount containing the env file. // want "commentstart: godoc for field FileKeySelector.VolumeName should start with 'volumeName ...'"
	// +required
	VolumeName string `json:"volumeName" protobuf:"bytes,1,opt,name=volumeName"` // want "requiredfields: field VolumeName should have the omitempty tag." "requiredfields: field VolumeName has a valid zero value \\(\\\"\\\"\\), but the validation is not complete \\(e.g. minimum length\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."
	// The path within the volume from which to select the file. // want "commentstart: godoc for field FileKeySelector.Path should start with 'path ...'"
	// Must be relative and may not contain the '..' path or start with '..'.
	// +required
	Path string `json:"path" protobuf:"bytes,2,opt,name=path"` // want "requiredfields: field Path should have the omitempty tag." "requiredfields: field Path has a valid zero value \\(\\\"\\\"\\), but the validation is not complete \\(e.g. minimum length\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."
	// The key within the env file. An invalid key will prevent the pod from starting. // want "commentstart: godoc for field FileKeySelector.Key should start with 'key ...'"
	// The keys defined within a source may consist of any printable ASCII characters except '='.
	// During Alpha stage of the EnvFiles feature gate, the key size is limited to 128 characters.
	// +required
	Key string `json:"key" protobuf:"bytes,3,opt,name=key"` // want "requiredfields: field Key should have the omitempty tag." "requiredfields: field Key has a valid zero value \\(\\\"\\\"\\), but the validation is not complete \\(e.g. minimum length\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."
	// Specify whether the file or its key must be defined. If the file or key // want "commentstart: godoc for field FileKeySelector.Optional should start with 'optional ...'"
	// does not exist, then the env var is not published.
	// If optional is set to true and the specified key does not exist,
	// the environment variable will not be set in the Pod's containers.
	//
	// If optional is set to false and the specified key does not exist,
	// an error will be returned during Pod creation.
	// +optional
	// +default=false
	Optional *bool `json:"optional,omitempty" protobuf:"varint,4,opt,name=optional"`
}

// ObjectFieldSelector selects an APIVersioned field of an object.
// +structType=atomic
type ObjectFieldSelector struct {
	// Version of the schema the FieldPath is written in terms of, defaults to "v1". // want "commentstart: godoc for field ObjectFieldSelector.APIVersion should start with 'apiVersion ...'"
	// +optional
	APIVersion string `json:"apiVersion,omitempty" protobuf:"bytes,1,opt,name=apiVersion"` // want "optionalfields: field APIVersion should be a pointer."
	// Path of the field to select in the specified API version. // want "commentstart: godoc for field ObjectFieldSelector.FieldPath should start with 'fieldPath ...'"
	FieldPath string `json:"fieldPath" protobuf:"bytes,2,opt,name=fieldPath"` // want "optionalorrequired: field ObjectFieldSelector.FieldPath must be marked as optional or required"
}

// ResourceFieldSelector represents container resources (cpu, memory) and their output format
// +structType=atomic
type ResourceFieldSelector struct {
	// Container name: required for volumes, optional for env vars // want "commentstart: godoc for field ResourceFieldSelector.ContainerName should start with 'containerName ...'"
	// +optional
	ContainerName string `json:"containerName,omitempty" protobuf:"bytes,1,opt,name=containerName"` // want "optionalfields: field ContainerName should be a pointer."
	// Required: resource to select // want "commentstart: godoc for field ResourceFieldSelector.Resource should start with 'resource ...'"
	Resource string `json:"resource" protobuf:"bytes,2,opt,name=resource"` // want "optionalorrequired: field ResourceFieldSelector.Resource must be marked as optional or required"
	// Specifies the output format of the exposed resources, defaults to "1"
	// +optional
	// Divisor resource.Quantity `json:"divisor,omitempty" protobuf:"bytes,3,opt,name=divisor"`
}

// Selects a key from a ConfigMap.
// +structType=atomic
type ConfigMapKeySelector struct {
	// The ConfigMap to select from.
	LocalObjectReference `json:",inline" protobuf:"bytes,1,opt,name=localObjectReference"`
	// The key to select. // want "commentstart: godoc for field ConfigMapKeySelector.Key should start with 'key ...'"
	Key string `json:"key" protobuf:"bytes,2,opt,name=key"` // want "optionalorrequired: field ConfigMapKeySelector.Key must be marked as optional or required"
	// Specify whether the ConfigMap or its key must be defined // want "commentstart: godoc for field ConfigMapKeySelector.Optional should start with 'optional ...'"
	// +optional
	Optional *bool `json:"optional,omitempty" protobuf:"varint,3,opt,name=optional"`
}

// SecretKeySelector selects a key of a Secret.
// +structType=atomic
type SecretKeySelector struct {
	// The name of the secret in the pod's namespace to select from.
	LocalObjectReference `json:",inline" protobuf:"bytes,1,opt,name=localObjectReference"`
	// The key of the secret to select from.  Must be a valid secret key. // want "commentstart: godoc for field SecretKeySelector.Key should start with 'key ...'"
	Key string `json:"key" protobuf:"bytes,2,opt,name=key"` // want "optionalorrequired: field SecretKeySelector.Key must be marked as optional or required"
	// Specify whether the Secret or its key must be defined // want "commentstart: godoc for field SecretKeySelector.Optional should start with 'optional ...'"
	// +optional
	Optional *bool `json:"optional,omitempty" protobuf:"varint,3,opt,name=optional"`
}

// EnvFromSource represents the source of a set of ConfigMaps or Secrets
type EnvFromSource struct {
	// Optional text to prepend to the name of each environment variable. // want "commentstart: godoc for field EnvFromSource.Prefix should start with 'prefix ...'"
	// May consist of any printable ASCII characters except '='.
	// +optional
	Prefix string `json:"prefix,omitempty" protobuf:"bytes,1,opt,name=prefix"` // want "optionalfields: field Prefix should be a pointer."
	// The ConfigMap to select from // want "commentstart: godoc for field EnvFromSource.ConfigMapRef should start with 'configMapRef ...'"
	// +optional
	ConfigMapRef *ConfigMapEnvSource `json:"configMapRef,omitempty" protobuf:"bytes,2,opt,name=configMapRef"`
	// The Secret to select from // want "commentstart: godoc for field EnvFromSource.SecretRef should start with 'secretRef ...'"
	// +optional
	SecretRef *SecretEnvSource `json:"secretRef,omitempty" protobuf:"bytes,3,opt,name=secretRef"`
}

// ConfigMapEnvSource selects a ConfigMap to populate the environment
// variables with.
//
// The contents of the target ConfigMap's Data field will represent the
// key-value pairs as environment variables.
type ConfigMapEnvSource struct {
	// The ConfigMap to select from.
	LocalObjectReference `json:",inline" protobuf:"bytes,1,opt,name=localObjectReference"`
	// Specify whether the ConfigMap must be defined // want "commentstart: godoc for field ConfigMapEnvSource.Optional should start with 'optional ...'"
	// +optional
	Optional *bool `json:"optional,omitempty" protobuf:"varint,2,opt,name=optional"`
}

// SecretEnvSource selects a Secret to populate the environment
// variables with.
//
// The contents of the target Secret's Data field will represent the
// key-value pairs as environment variables.
type SecretEnvSource struct {
	// The Secret to select from.
	LocalObjectReference `json:",inline" protobuf:"bytes,1,opt,name=localObjectReference"`
	// Specify whether the Secret must be defined // want "commentstart: godoc for field SecretEnvSource.Optional should start with 'optional ...'"
	// +optional
	Optional *bool `json:"optional,omitempty" protobuf:"varint,2,opt,name=optional"`
}

// ResourceRequirements describes the compute resource requirements.
type ResourceRequirements struct {
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

	// Claims lists the names of resources, defined in spec.resourceClaims, // want "commentstart: godoc for field ResourceRequirements.Claims should start with 'claims ...'"
	// that are used by this container.
	//
	// This field depends on the
	// DynamicResourceAllocation feature gate.
	//
	// This field is immutable. It can only be set for containers.
	//
	// +listType=map
	// +listMapKey=name
	// +featureGate=DynamicResourceAllocation
	// +optional
	Claims []ResourceClaim `json:"claims,omitempty" protobuf:"bytes,3,opt,name=claims"` // want "arrayofstruct: ResourceRequirements.Claims is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
}

// ResourceClaim references one entry in PodSpec.ResourceClaims.
type ResourceClaim struct {
	// Name must match the name of one entry in pod.spec.resourceClaims of // want "commentstart: godoc for field ResourceClaim.Name should start with 'name ...'"
	// the Pod where this field is used. It makes that resource available
	// inside a container.
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"` // want "optionalorrequired: field ResourceClaim.Name must be marked as optional or required"

	// Request is the name chosen for a request in the referenced claim. // want "commentstart: godoc for field ResourceClaim.Request should start with 'request ...'"
	// If empty, everything from the claim is made available, otherwise
	// only the result of this request.
	//
	// +optional
	Request string `json:"request,omitempty" protobuf:"bytes,2,opt,name=request"` // want "optionalfields: field Request should be a pointer."
}

// ResourceResizeRestartPolicy specifies how to handle container resource resize.
// +kubebuilder:validation:Enum=NotRequired;RestartContainer
type ResourceResizeRestartPolicy string

// These are the valid resource resize restart policy values:
const (
	// 'NotRequired' means Kubernetes will try to resize the container
	// without restarting it, if possible. Kubernetes may however choose to
	// restart the container if it is unable to actuate resize without a
	// restart. For e.g. the runtime doesn't support restart-free resizing.
	NotRequired ResourceResizeRestartPolicy = "NotRequired"
	// 'RestartContainer' means Kubernetes will resize the container in-place
	// by stopping and starting the container when new resources are applied.
	// This is needed for legacy applications. For e.g. java apps using the
	// -xmxN flag which are unable to use resized memory without restarting.
	RestartContainer ResourceResizeRestartPolicy = "RestartContainer"
)

// ResourceName is the name identifying various resources in a ResourceList.
type ResourceName string

// ContainerResizePolicy represents resource resize policy for the container.
type ContainerResizePolicy struct {
	// Name of the resource to which this resource resize policy applies. // want "commentstart: godoc for field ContainerResizePolicy.ResourceName should start with 'resourceName ...'"
	// Supported values: cpu, memory.
	ResourceName ResourceName `json:"resourceName" protobuf:"bytes,1,opt,name=resourceName,casttype=ResourceName"` // want "optionalorrequired: field ContainerResizePolicy.ResourceName must be marked as optional or required"
	// Restart policy to apply when specified resource is resized. // want "commentstart: godoc for field ContainerResizePolicy.RestartPolicy should start with 'restartPolicy ...'"
	// If not specified, it defaults to NotRequired.
	RestartPolicy ResourceResizeRestartPolicy `json:"restartPolicy" protobuf:"bytes,2,opt,name=restartPolicy,casttype=ResourceResizeRestartPolicy"` // want "optionalorrequired: field ContainerResizePolicy.RestartPolicy must be marked as optional or required"
}

// ContainerRestartPolicy is the restart policy for a single container.
// The only allowed values are "Always", "Never", and "OnFailure".
// +kubebuilder:validation:Enum=Always;Never;OnFailure
type ContainerRestartPolicy string

const (
	ContainerRestartPolicyAlways    ContainerRestartPolicy = "Always"
	ContainerRestartPolicyNever     ContainerRestartPolicy = "Never"
	ContainerRestartPolicyOnFailure ContainerRestartPolicy = "OnFailure"
)

// ContainerRestartRule describes how a container exit is handled.
type ContainerRestartRule struct {
	// Specifies the action taken on a container exit if the requirements // want "commentstart: godoc for field ContainerRestartRule.Action should start with 'action ...'"
	// are satisfied. The only possible value is "Restart" to restart the
	// container.
	// +required
	Action ContainerRestartRuleAction `json:"action,omitempty" proto:"bytes,1,opt,name=action" protobuf:"bytes,1,opt,name=action,casttype=ContainerRestartRuleAction"`

	// Represents the exit codes to check on container exits. // want "commentstart: godoc for field ContainerRestartRule.ExitCodes should start with 'exitCodes ...'"
	// +optional
	// +oneOf=when
	ExitCodes *ContainerRestartRuleOnExitCodes `json:"exitCodes,omitempty" proto:"bytes,2,opt,name=exitCodes" protobuf:"bytes,2,opt,name=exitCodes"`
}

// ContainerRestartRuleAction describes the action to take when the
// container exits.
// +kubebuilder:validation:Enum=Restart
type ContainerRestartRuleAction string

// The only valid action is Restart.
const (
	ContainerRestartRuleActionRestart ContainerRestartRuleAction = "Restart"
)

// ContainerRestartRuleOnExitCodes describes the condition
// for handling an exited container based on its exit codes.
type ContainerRestartRuleOnExitCodes struct {
	// Represents the relationship between the container exit code(s) and the // want "commentstart: godoc for field ContainerRestartRuleOnExitCodes.Operator should start with 'operator ...'"
	// specified values. Possible values are:
	// - In: the requirement is satisfied if the container exit code is in the
	//   set of specified values.
	// - NotIn: the requirement is satisfied if the container exit code is
	//   not in the set of specified values.
	// +required
	Operator ContainerRestartRuleOnExitCodesOperator `json:"operator,omitempty" proto:"bytes,1,opt,name=operator" protobuf:"bytes,1,opt,name=operator,casttype=ContainerRestartRuleOnExitCodesOperator"`

	// Specifies the set of values to check for container exit codes. // want "commentstart: godoc for field ContainerRestartRuleOnExitCodes.Values should start with 'values ...'"
	// At most 255 elements are allowed.
	// +optional
	// +listType=set
	Values []int32 `json:"values,omitempty" proto:"varint,2,rep,name=values" protobuf:"varint,2,rep,name=values"`
}

// ContainerRestartRuleOnExitCodesOperator describes the operator
// to take for the exit codes.
// +kubebuilder:validation:Enum
type ContainerRestartRuleOnExitCodesOperator string

const (
	ContainerRestartRuleOnExitCodesOpIn    ContainerRestartRuleOnExitCodesOperator = "In"
	ContainerRestartRuleOnExitCodesOpNotIn ContainerRestartRuleOnExitCodesOperator = "NotIn"
)

// Probe describes a health check to be performed against a container to determine whether it is
// alive or ready to receive traffic.
type Probe struct {
	// The action taken to determine the health of a container
	ProbeHandler `json:",inline" protobuf:"bytes,1,opt,name=handler"`
	// Number of seconds after the container has started before liveness probes are initiated. // want "commentstart: godoc for field Probe.InitialDelaySeconds should start with 'initialDelaySeconds ...'"
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	InitialDelaySeconds int32 `json:"initialDelaySeconds,omitempty" protobuf:"varint,2,opt,name=initialDelaySeconds"` // want "optionalfields: field InitialDelaySeconds should be a pointer."
	// Number of seconds after which the probe times out. // want "commentstart: godoc for field Probe.TimeoutSeconds should start with 'timeoutSeconds ...'"
	// Defaults to 1 second. Minimum value is 1.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	TimeoutSeconds int32 `json:"timeoutSeconds,omitempty" protobuf:"varint,3,opt,name=timeoutSeconds"` // want "optionalfields: field TimeoutSeconds should be a pointer."
	// How often (in seconds) to perform the probe. // want "commentstart: godoc for field Probe.PeriodSeconds should start with 'periodSeconds ...'"
	// Default to 10 seconds. Minimum value is 1.
	// +optional
	PeriodSeconds int32 `json:"periodSeconds,omitempty" protobuf:"varint,4,opt,name=periodSeconds"` // want "optionalfields: field PeriodSeconds should be a pointer."
	// Minimum consecutive successes for the probe to be considered successful after having failed. // want "commentstart: godoc for field Probe.SuccessThreshold should start with 'successThreshold ...'"
	// Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.
	// +optional
	SuccessThreshold int32 `json:"successThreshold,omitempty" protobuf:"varint,5,opt,name=successThreshold"` // want "optionalfields: field SuccessThreshold should be a pointer."
	// Minimum consecutive failures for the probe to be considered failed after having succeeded. // want "commentstart: godoc for field Probe.FailureThreshold should start with 'failureThreshold ...'"
	// Defaults to 3. Minimum value is 1.
	// +optional
	FailureThreshold int32 `json:"failureThreshold,omitempty" protobuf:"varint,6,opt,name=failureThreshold"` // want "optionalfields: field FailureThreshold should be a pointer."
	// Optional duration in seconds the pod needs to terminate gracefully upon probe failure. // want "commentstart: godoc for field Probe.TerminationGracePeriodSeconds should start with 'terminationGracePeriodSeconds ...'"
	// The grace period is the duration in seconds after the processes running in the pod are sent
	// a termination signal and the time when the processes are forcibly halted with a kill signal.
	// Set this value longer than the expected cleanup time for your process.
	// If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this
	// value overrides the value provided by the pod spec.
	// Value must be non-negative integer. The value zero indicates stop immediately via
	// the kill signal (no opportunity to shut down).
	// This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate.
	// Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.
	// +optional
	TerminationGracePeriodSeconds *int64 `json:"terminationGracePeriodSeconds,omitempty" protobuf:"varint,7,opt,name=terminationGracePeriodSeconds"`
}

// ProbeHandler defines a specific action that should be taken in a probe.
// One and only one of the fields must be specified.
type ProbeHandler struct {
	// Exec specifies a command to execute in the container. // want "commentstart: godoc for field ProbeHandler.Exec should start with 'exec ...'"
	// +optional
	Exec *ExecAction `json:"exec,omitempty" protobuf:"bytes,1,opt,name=exec"`
	// HTTPGet specifies an HTTP GET request to perform. // want "commentstart: godoc for field ProbeHandler.HTTPGet should start with 'httpGet ...'"
	// +optional
	HTTPGet *HTTPGetAction `json:"httpGet,omitempty" protobuf:"bytes,2,opt,name=httpGet"`
	// TCPSocket specifies a connection to a TCP port. // want "commentstart: godoc for field ProbeHandler.TCPSocket should start with 'tcpSocket ...'"
	// +optional
	TCPSocket *TCPSocketAction `json:"tcpSocket,omitempty" protobuf:"bytes,3,opt,name=tcpSocket"`
	// GRPC specifies a GRPC HealthCheckRequest. // want "commentstart: godoc for field ProbeHandler.GRPC should start with 'grpc ...'"
	// +optional
	GRPC *GRPCAction `json:"grpc,omitempty" protobuf:"bytes,4,opt,name=grpc"`
}

// HTTPGetAction describes an action based on HTTP Get requests.
type HTTPGetAction struct {
	// Path to access on the HTTP server. // want "commentstart: godoc for field HTTPGetAction.Path should start with 'path ...'"
	// +optional
	Path string `json:"path,omitempty" protobuf:"bytes,1,opt,name=path"` // want "optionalfields: field Path should be a pointer."
	// Name or number of the port to access on the container.
	// Number must be in the range 1 to 65535.
	// Name must be an IANA_SVC_NAME.
	// Port intstr.IntOrString `json:"port" protobuf:"bytes,2,opt,name=port"`

	// Host name to connect to, defaults to the pod IP. You probably want to set // want "commentstart: godoc for field HTTPGetAction.Host should start with 'host ...'"
	// "Host" in httpHeaders instead.
	// +optional
	Host string `json:"host,omitempty" protobuf:"bytes,3,opt,name=host"` // want "optionalfields: field Host should be a pointer."
	// Scheme to use for connecting to the host. // want "commentstart: godoc for field HTTPGetAction.Scheme should start with 'scheme ...'"
	// Defaults to HTTP.
	// +optional
	Scheme URIScheme `json:"scheme,omitempty" protobuf:"bytes,4,opt,name=scheme,casttype=URIScheme"` // want "optionalfields: field Scheme should be a pointer."
	// Custom headers to set in the request. HTTP allows repeated headers. // want "commentstart: godoc for field HTTPGetAction.HTTPHeaders should start with 'httpHeaders ...'"
	// +optional
	// +listType=atomic
	HTTPHeaders []HTTPHeader `json:"httpHeaders,omitempty" protobuf:"bytes,5,rep,name=httpHeaders"` // want "arrayofstruct: HTTPGetAction.HTTPHeaders is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
}

// HTTPHeader describes a custom header to be used in HTTP probes
type HTTPHeader struct {
	// The header field name. // want "commentstart: godoc for field HTTPHeader.Name should start with 'name ...'"
	// This will be canonicalized upon output, so case-variant names will be understood as the same header.
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"` // want "optionalorrequired: field HTTPHeader.Name must be marked as optional or required"
	// The header field value // want "commentstart: godoc for field HTTPHeader.Value should start with 'value ...'"
	Value string `json:"value" protobuf:"bytes,2,opt,name=value"` // want "optionalorrequired: field HTTPHeader.Value must be marked as optional or required"
}

// URIScheme identifies the scheme used for connection to a host for Get actions
// +kubebuilder:validation:Enum
type URIScheme string

const (
	// URISchemeHTTP means that the scheme used will be http://
	URISchemeHTTP URIScheme = "HTTP"
	// URISchemeHTTPS means that the scheme used will be https://
	URISchemeHTTPS URIScheme = "HTTPS"
)

// TCPSocketAction describes an action based on opening a socket
type TCPSocketAction struct {
	// Number or name of the port to access on the container.
	// Number must be in the range 1 to 65535.
	// Name must be an IANA_SVC_NAME.
	// Port intstr.IntOrString `json:"port" protobuf:"bytes,1,opt,name=port"`

	// Optional: Host name to connect to, defaults to the pod IP. // want "commentstart: godoc for field TCPSocketAction.Host should start with 'host ...'"
	// +optional
	Host string `json:"host,omitempty" protobuf:"bytes,2,opt,name=host"` // want "optionalfields: field Host should be a pointer."
}

// GRPCAction specifies an action involving a GRPC service.
type GRPCAction struct {
	// Port number of the gRPC service. Number must be in the range 1 to 65535. // want "commentstart: godoc for field GRPCAction.Port should start with 'port ...'"
	Port int32 `json:"port" protobuf:"bytes,1,opt,name=port"` // want "optionalorrequired: field GRPCAction.Port must be marked as optional or required"

	// Service is the name of the service to place in the gRPC HealthCheckRequest // want "commentstart: godoc for field GRPCAction.Service should start with 'service ...'"
	// (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).
	//
	// If this is not specified, the default behavior is defined by gRPC.
	// +optional
	// +default=""
	Service *string `json:"service" protobuf:"bytes,2,opt,name=service"` // want "optionalfields: field Service should have the omitempty tag."
}

// ExecAction describes a "run in container" action.
type ExecAction struct {
	// Command is the command line to execute inside the container, the working directory for the // want "commentstart: godoc for field ExecAction.Command should start with 'command ...'"
	// command  is root ('/') in the container's filesystem. The command is simply exec'd, it is
	// not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use
	// a shell, you need to explicitly call out to that shell.
	// Exit status of 0 is treated as live/healthy and non-zero is unhealthy.
	// +optional
	// +listType=atomic
	Command []string `json:"command,omitempty" protobuf:"bytes,1,rep,name=command"`
}

// SleepAction describes a "sleep" action.
type SleepAction struct {
	// Seconds is the number of seconds to sleep. // want "commentstart: godoc for field SleepAction.Seconds should start with 'seconds ...'"
	Seconds int64 `json:"seconds" protobuf:"bytes,1,opt,name=seconds"` // want "optionalorrequired: field SleepAction.Seconds must be marked as optional or required"
}

// Signal defines the stop signal of containers
// +kubebuilder:validation:Enum=SIGABRT;SIGALRM;SIGBUS;SIGCHLD;SIGCLD;SIGCONT;SIGFPE;SIGHUP;SIGILL;SIGINT;SIGIO;SIGIOT;SIGKILL;SIGPIPE;SIGPOLL;SIGPROF;SIGPWR;SIGQUIT;SIGSEGV;SIGSTKFLT;SIGSTOP;SIGSYS;SIGTERM;SIGTRAP;SIGTSTP;SIGTTIN;SIGTTOU;SIGURG;SIGUSR1;SIGUSR2;SIGVTALRM;SIGWINCH;SIGXCPU;SIGXFSZ;SIGRTMIN;SIGRTMIN+1;SIGRTMIN+2;SIGRTMIN+3;SIGRTMIN+4;SIGRTMIN+5;SIGRTMIN+6;SIGRTMIN+7;SIGRTMIN+8;SIGRTMIN+9;SIGRTMIN+10;SIGRTMIN+11;SIGRTMIN+12;SIGRTMIN+13;SIGRTMIN+14;SIGRTMIN+15;SIGRTMAX-14;SIGRTMAX-13;SIGRTMAX-12;SIGRTMAX-11;SIGRTMAX-10;SIGRTMAX-9;SIGRTMAX-8;SIGRTMAX-7;SIGRTMAX-6;SIGRTMAX-5;SIGRTMAX-4;SIGRTMAX-3;SIGRTMAX-2;SIGRTMAX-1;SIGRTMAX
type Signal string

const (
	SIGABRT         Signal = "SIGABRT"
	SIGALRM         Signal = "SIGALRM"
	SIGBUS          Signal = "SIGBUS"
	SIGCHLD         Signal = "SIGCHLD"
	SIGCLD          Signal = "SIGCLD"
	SIGCONT         Signal = "SIGCONT"
	SIGFPE          Signal = "SIGFPE"
	SIGHUP          Signal = "SIGHUP"
	SIGILL          Signal = "SIGILL"
	SIGINT          Signal = "SIGINT"
	SIGIO           Signal = "SIGIO"
	SIGIOT          Signal = "SIGIOT"
	SIGKILL         Signal = "SIGKILL"
	SIGPIPE         Signal = "SIGPIPE"
	SIGPOLL         Signal = "SIGPOLL"
	SIGPROF         Signal = "SIGPROF"
	SIGPWR          Signal = "SIGPWR"
	SIGQUIT         Signal = "SIGQUIT"
	SIGSEGV         Signal = "SIGSEGV"
	SIGSTKFLT       Signal = "SIGSTKFLT"
	SIGSTOP         Signal = "SIGSTOP"
	SIGSYS          Signal = "SIGSYS"
	SIGTERM         Signal = "SIGTERM"
	SIGTRAP         Signal = "SIGTRAP"
	SIGTSTP         Signal = "SIGTSTP"
	SIGTTIN         Signal = "SIGTTIN"
	SIGTTOU         Signal = "SIGTTOU"
	SIGURG          Signal = "SIGURG"
	SIGUSR1         Signal = "SIGUSR1"
	SIGUSR2         Signal = "SIGUSR2"
	SIGVTALRM       Signal = "SIGVTALRM"
	SIGWINCH        Signal = "SIGWINCH"
	SIGXCPU         Signal = "SIGXCPU"
	SIGXFSZ         Signal = "SIGXFSZ"
	SIGRTMIN        Signal = "SIGRTMIN"
	SIGRTMINPLUS1   Signal = "SIGRTMIN+1"
	SIGRTMINPLUS2   Signal = "SIGRTMIN+2"
	SIGRTMINPLUS3   Signal = "SIGRTMIN+3"
	SIGRTMINPLUS4   Signal = "SIGRTMIN+4"
	SIGRTMINPLUS5   Signal = "SIGRTMIN+5"
	SIGRTMINPLUS6   Signal = "SIGRTMIN+6"
	SIGRTMINPLUS7   Signal = "SIGRTMIN+7"
	SIGRTMINPLUS8   Signal = "SIGRTMIN+8"
	SIGRTMINPLUS9   Signal = "SIGRTMIN+9"
	SIGRTMINPLUS10  Signal = "SIGRTMIN+10"
	SIGRTMINPLUS11  Signal = "SIGRTMIN+11"
	SIGRTMINPLUS12  Signal = "SIGRTMIN+12"
	SIGRTMINPLUS13  Signal = "SIGRTMIN+13"
	SIGRTMINPLUS14  Signal = "SIGRTMIN+14"
	SIGRTMINPLUS15  Signal = "SIGRTMIN+15"
	SIGRTMAXMINUS14 Signal = "SIGRTMAX-14"
	SIGRTMAXMINUS13 Signal = "SIGRTMAX-13"
	SIGRTMAXMINUS12 Signal = "SIGRTMAX-12"
	SIGRTMAXMINUS11 Signal = "SIGRTMAX-11"
	SIGRTMAXMINUS10 Signal = "SIGRTMAX-10"
	SIGRTMAXMINUS9  Signal = "SIGRTMAX-9"
	SIGRTMAXMINUS8  Signal = "SIGRTMAX-8"
	SIGRTMAXMINUS7  Signal = "SIGRTMAX-7"
	SIGRTMAXMINUS6  Signal = "SIGRTMAX-6"
	SIGRTMAXMINUS5  Signal = "SIGRTMAX-5"
	SIGRTMAXMINUS4  Signal = "SIGRTMAX-4"
	SIGRTMAXMINUS3  Signal = "SIGRTMAX-3"
	SIGRTMAXMINUS2  Signal = "SIGRTMAX-2"
	SIGRTMAXMINUS1  Signal = "SIGRTMAX-1"
	SIGRTMAX        Signal = "SIGRTMAX"
)

// Lifecycle describes actions that the management system should take in response to container lifecycle
// events. For the PostStart and PreStop lifecycle handlers, management of the container blocks
// until the action is complete, unless the container process fails, in which case the handler is aborted.
type Lifecycle struct {
	// PostStart is called immediately after a container is created. If the handler fails, // want "commentstart: godoc for field Lifecycle.PostStart should start with 'postStart ...'"
	// the container is terminated and restarted according to its restart policy.
	// Other management of the container blocks until the hook completes.
	// More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks
	// +optional
	PostStart *LifecycleHandler `json:"postStart,omitempty" protobuf:"bytes,1,opt,name=postStart"`
	// PreStop is called immediately before a container is terminated due to an // want "commentstart: godoc for field Lifecycle.PreStop should start with 'preStop ...'"
	// API request or management event such as liveness/startup probe failure,
	// preemption, resource contention, etc. The handler is not called if the
	// container crashes or exits. The Pod's termination grace period countdown begins before the
	// PreStop hook is executed. Regardless of the outcome of the handler, the
	// container will eventually terminate within the Pod's termination grace
	// period (unless delayed by finalizers). Other management of the container blocks until the hook completes
	// or until the termination grace period is reached.
	// More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks
	// +optional
	PreStop *LifecycleHandler `json:"preStop,omitempty" protobuf:"bytes,2,opt,name=preStop"`
	// StopSignal defines which signal will be sent to a container when it is being stopped. // want "commentstart: godoc for field Lifecycle.StopSignal should start with 'stopSignal ...'"
	// If not specified, the default is defined by the container runtime in use.
	// StopSignal can only be set for Pods with a non-empty .spec.os.name
	// +optional
	StopSignal *Signal `json:"stopSignal,omitempty" protobuf:"bytes,3,opt,name=stopSignal"`
}

// LifecycleHandler defines a specific action that should be taken in a lifecycle
// hook. One and only one of the fields, except TCPSocket must be specified.
type LifecycleHandler struct {
	// Exec specifies a command to execute in the container. // want "commentstart: godoc for field LifecycleHandler.Exec should start with 'exec ...'"
	// +optional
	Exec *ExecAction `json:"exec,omitempty" protobuf:"bytes,1,opt,name=exec"`
	// HTTPGet specifies an HTTP GET request to perform. // want "commentstart: godoc for field LifecycleHandler.HTTPGet should start with 'httpGet ...'"
	// +optional
	HTTPGet *HTTPGetAction `json:"httpGet,omitempty" protobuf:"bytes,2,opt,name=httpGet"`
	// Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for // want "commentstart: godoc for field LifecycleHandler.TCPSocket should start with 'tcpSocket ...'"
	// for backward compatibility. There is no validation of this field and
	// lifecycle hooks will fail at runtime when it is specified.
	// +optional
	TCPSocket *TCPSocketAction `json:"tcpSocket,omitempty" protobuf:"bytes,3,opt,name=tcpSocket"`
	// Sleep represents a duration that the container should sleep. // want "commentstart: godoc for field LifecycleHandler.Sleep should start with 'sleep ...'"
	// +optional
	Sleep *SleepAction `json:"sleep,omitempty" protobuf:"bytes,4,opt,name=sleep"`
}

// TerminationMessagePolicy describes how termination messages are retrieved from a container.
// +kubebuilder:validation:Enum
type TerminationMessagePolicy string

const (
	// TerminationMessageReadFile is the default behavior and will set the container status message to
	// the contents of the container's terminationMessagePath when the container exits.
	TerminationMessageReadFile TerminationMessagePolicy = "File"
	// TerminationMessageFallbackToLogsOnError will read the most recent contents of the container logs
	// for the container status message when the container exits with an error and the
	// terminationMessagePath has no contents.
	TerminationMessageFallbackToLogsOnError TerminationMessagePolicy = "FallbackToLogsOnError"
)

// PullPolicy describes a policy for if/when to pull a container image
// +kubebuilder:validation:Enum
type PullPolicy string

const (
	// PullAlways means that kubelet always attempts to pull the latest image. Container will fail If the pull fails.
	PullAlways PullPolicy = "Always"
	// PullNever means that kubelet never pulls an image, but only uses a local image. Container will fail if the image isn't present
	PullNever PullPolicy = "Never"
	// PullIfNotPresent means that kubelet pulls if the image isn't present on disk. Container will fail if the image isn't present and the pull fails.
	PullIfNotPresent PullPolicy = "IfNotPresent"
)

// SecurityContext holds security configuration that will be applied to a container.
// Some fields are present in both SecurityContext and PodSecurityContext.  When both
// are set, the values in SecurityContext take precedence.
type SecurityContext struct {
	// The capabilities to add/drop when running containers. // want "commentstart: godoc for field SecurityContext.Capabilities should start with 'capabilities ...'"
	// Defaults to the default set of capabilities granted by the container runtime.
	// Note that this field cannot be set when spec.os.name is windows.
	// +optional
	Capabilities *Capabilities `json:"capabilities,omitempty" protobuf:"bytes,1,opt,name=capabilities"`
	// Run container in privileged mode. // want "commentstart: godoc for field SecurityContext.Privileged should start with 'privileged ...'"
	// Processes in privileged containers are essentially equivalent to root on the host.
	// Defaults to false.
	// Note that this field cannot be set when spec.os.name is windows.
	// +optional
	Privileged *bool `json:"privileged,omitempty" protobuf:"varint,2,opt,name=privileged"`
	// The SELinux context to be applied to the container. // want "commentstart: godoc for field SecurityContext.SELinuxOptions should start with 'seLinuxOptions ...'"
	// If unspecified, the container runtime will allocate a random SELinux context for each
	// container.  May also be set in PodSecurityContext.  If set in both SecurityContext and
	// PodSecurityContext, the value specified in SecurityContext takes precedence.
	// Note that this field cannot be set when spec.os.name is windows.
	// +optional
	SELinuxOptions *SELinuxOptions `json:"seLinuxOptions,omitempty" protobuf:"bytes,3,opt,name=seLinuxOptions"`
	// The Windows specific settings applied to all containers. // want "commentstart: godoc for field SecurityContext.WindowsOptions should start with 'windowsOptions ...'"
	// If unspecified, the options from the PodSecurityContext will be used.
	// If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	// Note that this field cannot be set when spec.os.name is linux.
	// +optional
	WindowsOptions *WindowsSecurityContextOptions `json:"windowsOptions,omitempty" protobuf:"bytes,10,opt,name=windowsOptions"`
	// The UID to run the entrypoint of the container process. // want "commentstart: godoc for field SecurityContext.RunAsUser should start with 'runAsUser ...'"
	// Defaults to user specified in image metadata if unspecified.
	// May also be set in PodSecurityContext.  If set in both SecurityContext and
	// PodSecurityContext, the value specified in SecurityContext takes precedence.
	// Note that this field cannot be set when spec.os.name is windows.
	// +optional
	RunAsUser *int64 `json:"runAsUser,omitempty" protobuf:"varint,4,opt,name=runAsUser"`
	// The GID to run the entrypoint of the container process. // want "commentstart: godoc for field SecurityContext.RunAsGroup should start with 'runAsGroup ...'"
	// Uses runtime default if unset.
	// May also be set in PodSecurityContext.  If set in both SecurityContext and
	// PodSecurityContext, the value specified in SecurityContext takes precedence.
	// Note that this field cannot be set when spec.os.name is windows.
	// +optional
	RunAsGroup *int64 `json:"runAsGroup,omitempty" protobuf:"varint,8,opt,name=runAsGroup"`
	// Indicates that the container must run as a non-root user. // want "commentstart: godoc for field SecurityContext.RunAsNonRoot should start with 'runAsNonRoot ...'"
	// If true, the Kubelet will validate the image at runtime to ensure that it
	// does not run as UID 0 (root) and fail to start the container if it does.
	// If unset or false, no such validation will be performed.
	// May also be set in PodSecurityContext.  If set in both SecurityContext and
	// PodSecurityContext, the value specified in SecurityContext takes precedence.
	// +optional
	RunAsNonRoot *bool `json:"runAsNonRoot,omitempty" protobuf:"varint,5,opt,name=runAsNonRoot"`
	// Whether this container has a read-only root filesystem. // want "commentstart: godoc for field SecurityContext.ReadOnlyRootFilesystem should start with 'readOnlyRootFilesystem ...'"
	// Default is false.
	// Note that this field cannot be set when spec.os.name is windows.
	// +optional
	ReadOnlyRootFilesystem *bool `json:"readOnlyRootFilesystem,omitempty" protobuf:"varint,6,opt,name=readOnlyRootFilesystem"`
	// AllowPrivilegeEscalation controls whether a process can gain more // want "commentstart: godoc for field SecurityContext.AllowPrivilegeEscalation should start with 'allowPrivilegeEscalation ...'"
	// privileges than its parent process. This bool directly controls if
	// the no_new_privs flag will be set on the container process.
	// AllowPrivilegeEscalation is true always when the container is:
	// 1) run as Privileged
	// 2) has CAP_SYS_ADMIN
	// Note that this field cannot be set when spec.os.name is windows.
	// +optional
	AllowPrivilegeEscalation *bool `json:"allowPrivilegeEscalation,omitempty" protobuf:"varint,7,opt,name=allowPrivilegeEscalation"`
	// procMount denotes the type of proc mount to use for the containers.
	// The default value is Default which uses the container runtime defaults for
	// readonly paths and masked paths.
	// This requires the ProcMountType feature flag to be enabled.
	// Note that this field cannot be set when spec.os.name is windows.
	// +optional
	ProcMount *ProcMountType `json:"procMount,omitempty" protobuf:"bytes,9,opt,name=procMount"`
	// The seccomp options to use by this container. If seccomp options are // want "commentstart: godoc for field SecurityContext.SeccompProfile should start with 'seccompProfile ...'"
	// provided at both the pod & container level, the container options
	// override the pod options.
	// Note that this field cannot be set when spec.os.name is windows.
	// +optional
	SeccompProfile *SeccompProfile `json:"seccompProfile,omitempty" protobuf:"bytes,11,opt,name=seccompProfile"`
	// appArmorProfile is the AppArmor options to use by this container. If set, this profile
	// overrides the pod's appArmorProfile.
	// Note that this field cannot be set when spec.os.name is windows.
	// +optional
	AppArmorProfile *AppArmorProfile `json:"appArmorProfile,omitempty" protobuf:"bytes,12,opt,name=appArmorProfile"`
}

// +kubebuilder:validation:Enum
type ProcMountType string

const (
	// DefaultProcMount uses the container runtime defaults for readonly and masked
	// paths for /proc.  Most container runtimes mask certain paths in /proc to avoid
	// accidental security exposure of special devices or information.
	DefaultProcMount ProcMountType = "Default"

	// UnmaskedProcMount bypasses the default masking behavior of the container
	// runtime and ensures the newly created /proc the container stays in tact with
	// no modifications.
	UnmaskedProcMount ProcMountType = "Unmasked"
)

// Capability represent POSIX capabilities type
type Capability string

// Adds and removes POSIX capabilities from running containers.
type Capabilities struct {
	// Added capabilities // want "commentstart: godoc for field Capabilities.Add should start with 'add ...'"
	// +optional
	// +listType=atomic
	Add []Capability `json:"add,omitempty" protobuf:"bytes,1,rep,name=add,casttype=Capability"`
	// Removed capabilities // want "commentstart: godoc for field Capabilities.Drop should start with 'drop ...'"
	// +optional
	// +listType=atomic
	Drop []Capability `json:"drop,omitempty" protobuf:"bytes,2,rep,name=drop,casttype=Capability"`
}

// SELinuxOptions are the labels to be applied to the container
type SELinuxOptions struct {
	// User is a SELinux user label that applies to the container. // want "commentstart: godoc for field SELinuxOptions.User should start with 'user ...'"
	// +optional
	User string `json:"user,omitempty" protobuf:"bytes,1,opt,name=user"` // want "optionalfields: field User should be a pointer."
	// Role is a SELinux role label that applies to the container. // want "commentstart: godoc for field SELinuxOptions.Role should start with 'role ...'"
	// +optional
	Role string `json:"role,omitempty" protobuf:"bytes,2,opt,name=role"` // want "optionalfields: field Role should be a pointer."
	// Type is a SELinux type label that applies to the container. // want "commentstart: godoc for field SELinuxOptions.Type should start with 'type ...'"
	// +optional
	Type string `json:"type,omitempty" protobuf:"bytes,3,opt,name=type"` // want "optionalfields: field Type should be a pointer."
	// Level is SELinux level label that applies to the container. // want "commentstart: godoc for field SELinuxOptions.Level should start with 'level ...'"
	// +optional
	Level string `json:"level,omitempty" protobuf:"bytes,4,opt,name=level"` // want "optionalfields: field Level should be a pointer."
}

// WindowsSecurityContextOptions contain Windows-specific options and credentials.
type WindowsSecurityContextOptions struct {
	// GMSACredentialSpecName is the name of the GMSA credential spec to use. // want "commentstart: godoc for field WindowsSecurityContextOptions.GMSACredentialSpecName should start with 'gmsaCredentialSpecName ...'"
	// +optional
	GMSACredentialSpecName *string `json:"gmsaCredentialSpecName,omitempty" protobuf:"bytes,1,opt,name=gmsaCredentialSpecName"`

	// GMSACredentialSpec is where the GMSA admission webhook // want "commentstart: godoc for field WindowsSecurityContextOptions.GMSACredentialSpec should start with 'gmsaCredentialSpec ...'"
	// (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the
	// GMSA credential spec named by the GMSACredentialSpecName field.
	// +optional
	GMSACredentialSpec *string `json:"gmsaCredentialSpec,omitempty" protobuf:"bytes,2,opt,name=gmsaCredentialSpec"`

	// The UserName in Windows to run the entrypoint of the container process. // want "commentstart: godoc for field WindowsSecurityContextOptions.RunAsUserName should start with 'runAsUserName ...'"
	// Defaults to the user specified in image metadata if unspecified.
	// May also be set in PodSecurityContext. If set in both SecurityContext and
	// PodSecurityContext, the value specified in SecurityContext takes precedence.
	// +optional
	RunAsUserName *string `json:"runAsUserName,omitempty" protobuf:"bytes,3,opt,name=runAsUserName"`

	// HostProcess determines if a container should be run as a 'Host Process' container. // want "commentstart: godoc for field WindowsSecurityContextOptions.HostProcess should start with 'hostProcess ...'"
	// All of a Pod's containers must have the same effective HostProcess value
	// (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).
	// In addition, if HostProcess is true then HostNetwork must also be set to true.
	// +optional
	HostProcess *bool `json:"hostProcess,omitempty" protobuf:"bytes,4,opt,name=hostProcess"`
}

// SeccompProfile defines a pod/container's seccomp profile settings.
// Only one profile source may be set.
// +union
type SeccompProfile struct {
	// type indicates which kind of seccomp profile will be applied.
	// Valid options are:
	//
	// Localhost - a profile defined in a file on the node should be used.
	// RuntimeDefault - the container runtime default profile should be used.
	// Unconfined - no profile should be applied.
	// +unionDiscriminator
	Type SeccompProfileType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=SeccompProfileType"` // want "optionalorrequired: field SeccompProfile.Type must be marked as optional or required"
	// localhostProfile indicates a profile defined in a file on the node should be used.
	// The profile must be preconfigured on the node to work.
	// Must be a descending path, relative to the kubelet's configured seccomp profile location.
	// Must be set if type is "Localhost". Must NOT be set for any other type.
	// +optional
	LocalhostProfile *string `json:"localhostProfile,omitempty" protobuf:"bytes,2,opt,name=localhostProfile"`
}

// SeccompProfileType defines the supported seccomp profile types.
// +kubebuilder:validation:Enum
type SeccompProfileType string

const (
	// SeccompProfileTypeUnconfined indicates no seccomp profile is applied (A.K.A. unconfined).
	SeccompProfileTypeUnconfined SeccompProfileType = "Unconfined"
	// SeccompProfileTypeRuntimeDefault represents the default container runtime seccomp profile.
	SeccompProfileTypeRuntimeDefault SeccompProfileType = "RuntimeDefault"
	// SeccompProfileTypeLocalhost indicates a profile defined in a file on the node should be used.
	// The file's location relative to <kubelet-root-dir>/seccomp.
	SeccompProfileTypeLocalhost SeccompProfileType = "Localhost"
)

// AppArmorProfile defines a pod or container's AppArmor settings.
// +union
type AppArmorProfile struct {
	// type indicates which kind of AppArmor profile will be applied.
	// Valid options are:
	//   Localhost - a profile pre-loaded on the node.
	//   RuntimeDefault - the container runtime's default profile.
	//   Unconfined - no AppArmor enforcement.
	// +unionDiscriminator
	Type AppArmorProfileType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=AppArmorProfileType"` // want "optionalorrequired: field AppArmorProfile.Type must be marked as optional or required"

	// localhostProfile indicates a profile loaded on the node that should be used.
	// The profile must be preconfigured on the node to work.
	// Must match the loaded name of the profile.
	// Must be set if and only if type is "Localhost".
	// +optional
	LocalhostProfile *string `json:"localhostProfile,omitempty" protobuf:"bytes,2,opt,name=localhostProfile"`
}

// +kubebuilder:validation:Enum
type AppArmorProfileType string

const (
	// AppArmorProfileTypeUnconfined indicates that no AppArmor profile should be enforced.
	AppArmorProfileTypeUnconfined AppArmorProfileType = "Unconfined"
	// AppArmorProfileTypeRuntimeDefault indicates that the container runtime's default AppArmor
	// profile should be used.
	AppArmorProfileTypeRuntimeDefault AppArmorProfileType = "RuntimeDefault"
	// AppArmorProfileTypeLocalhost indicates that a profile pre-loaded on the node should be used.
	AppArmorProfileTypeLocalhost AppArmorProfileType = "Localhost"
)

// HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the
// pod's hosts file.
type HostAlias struct {
	// IP address of the host file entry. // want "commentstart: godoc for field HostAlias.IP should start with 'ip ...'"
	// +required
	IP string `json:"ip" protobuf:"bytes,1,opt,name=ip"` // want "requiredfields: field IP should have the omitempty tag." "requiredfields: field IP has a valid zero value \\(\\\"\\\"\\), but the validation is not complete \\(e.g. minimum length\\). The field should be a pointer to allow the zero value to be set. If the zero value is not a valid use case, complete the validation and remove the pointer."
	// Hostnames for the above IP address. // want "commentstart: godoc for field HostAlias.Hostnames should start with 'hostnames ...'"
	// +listType=atomic
	Hostnames []string `json:"hostnames,omitempty" protobuf:"bytes,2,rep,name=hostnames"` // want "optionalorrequired: field HostAlias.Hostnames must be marked as optional or required"
}

// RestartPolicy describes how the container should be restarted.
// Only one of the following restart policies may be specified.
// If none of the following policies is specified, the default one
// is RestartPolicyAlways.
// +kubebuilder:validation:Enum
type RestartPolicy string

const (
	RestartPolicyAlways    RestartPolicy = "Always"
	RestartPolicyOnFailure RestartPolicy = "OnFailure"
	RestartPolicyNever     RestartPolicy = "Never"
)

// DNSPolicy defines how a pod's DNS will be configured.
// +kubebuilder:validation:Enum
type DNSPolicy string

const (
	// DNSClusterFirstWithHostNet indicates that the pod should use cluster DNS
	// first, if it is available, then fall back on the default
	// (as determined by kubelet) DNS settings.
	DNSClusterFirstWithHostNet DNSPolicy = "ClusterFirstWithHostNet"

	// DNSClusterFirst indicates that the pod should use cluster DNS
	// first unless hostNetwork is true, if it is available, then
	// fall back on the default (as determined by kubelet) DNS settings.
	DNSClusterFirst DNSPolicy = "ClusterFirst"

	// DNSDefault indicates that the pod should use the default (as
	// determined by kubelet) DNS settings.
	DNSDefault DNSPolicy = "Default"

	// DNSNone indicates that the pod should use empty DNS settings. DNS
	// parameters such as nameservers and search paths should be defined via
	// DNSConfig.
	DNSNone DNSPolicy = "None"
)
