package defaultconfigurations

// PodSpec is a description of a pod.
type PodSpec struct {
	// List of volumes that can be mounted by containers belonging to the pod. // want "commentstart: godoc for field PodSpec.Volumes should start with 'volumes ...'"
	// More info: https://kubernetes.io/docs/concepts/storage/volumes
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge,retainKeys
	// +listType=map
	// +listMapKey=name
	Volumes []Volume `json:"volumes,omitempty" patchStrategy:"merge,retainKeys" patchMergeKey:"name" protobuf:"bytes,1,rep,name=volumes"` // want "arrayofstruct: PodSpec.Volumes is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// List of initialization containers belonging to the pod. // want "commentstart: godoc for field PodSpec.InitContainers should start with 'initContainers ...'"
	// Init containers are executed in order prior to containers being started. If any
	// init container fails, the pod is considered to have failed and is handled according
	// to its restartPolicy. The name for an init container or normal container must be
	// unique among all containers.
	// Init containers may not have Lifecycle actions, Readiness probes, Liveness probes, or Startup probes.
	// The resourceRequirements of an init container are taken into account during scheduling
	// by finding the highest request/limit for each resource type, and then using the max of
	// that value or the sum of the normal containers. Limits are applied to init containers
	// in a similar fashion.
	// Init containers cannot currently be added or removed.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/
	// +patchMergeKey=name
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=name
	InitContainers []Container `json:"initContainers,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,20,rep,name=initContainers"` // want "optionalorrequired: field PodSpec.InitContainers must be marked as optional or required" "arrayofstruct: PodSpec.InitContainers is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// List of containers belonging to the pod. // want "commentstart: godoc for field PodSpec.Containers should start with 'containers ...'"
	// Containers cannot currently be added or removed.
	// There must be at least one container in a Pod.
	// Cannot be updated.
	// +patchMergeKey=name
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=name
	Containers []Container `json:"containers" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,2,rep,name=containers"` // want "optionalorrequired: field PodSpec.Containers must be marked as optional or required" "arrayofstruct: PodSpec.Containers is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// List of ephemeral containers run in this pod. Ephemeral containers may be run in an existing // want "commentstart: godoc for field PodSpec.EphemeralContainers should start with 'ephemeralContainers ...'"
	// pod to perform user-initiated actions such as debugging. This list cannot be specified when
	// creating a pod, and it cannot be modified by updating the pod spec. In order to add an
	// ephemeral container to an existing pod, use the pod's ephemeralcontainers subresource.
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=name
	EphemeralContainers []EphemeralContainer `json:"ephemeralContainers,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,34,rep,name=ephemeralContainers"` // want "arrayofstruct: PodSpec.EphemeralContainers is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// Restart policy for all containers within the pod. // want "commentstart: godoc for field PodSpec.RestartPolicy should start with 'restartPolicy ...'"
	// One of Always, OnFailure, Never. In some contexts, only a subset of those values may be permitted.
	// Default to Always.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	// +optional
	RestartPolicy RestartPolicy `json:"restartPolicy,omitempty" protobuf:"bytes,3,opt,name=restartPolicy,casttype=RestartPolicy"` // want "optionalfields: field RestartPolicy should be a pointer."
	// Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request. // want "commentstart: godoc for field PodSpec.TerminationGracePeriodSeconds should start with 'terminationGracePeriodSeconds ...'"
	// Value must be non-negative integer. The value zero indicates stop immediately via
	// the kill signal (no opportunity to shut down).
	// If this value is nil, the default grace period will be used instead.
	// The grace period is the duration in seconds after the processes running in the pod are sent
	// a termination signal and the time when the processes are forcibly halted with a kill signal.
	// Set this value longer than the expected cleanup time for your process.
	// Defaults to 30 seconds.
	// +optional
	TerminationGracePeriodSeconds *int64 `json:"terminationGracePeriodSeconds,omitempty" protobuf:"varint,4,opt,name=terminationGracePeriodSeconds"`
	// Optional duration in seconds the pod may be active on the node relative to // want "commentstart: godoc for field PodSpec.ActiveDeadlineSeconds should start with 'activeDeadlineSeconds ...'"
	// StartTime before the system will actively try to mark it failed and kill associated containers.
	// Value must be a positive integer.
	// +optional
	ActiveDeadlineSeconds *int64 `json:"activeDeadlineSeconds,omitempty" protobuf:"varint,5,opt,name=activeDeadlineSeconds"`
	// Set DNS policy for the pod. // want "commentstart: godoc for field PodSpec.DNSPolicy should start with 'dnsPolicy ...'"
	// Defaults to "ClusterFirst".
	// Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'.
	// DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy.
	// To have DNS options set along with hostNetwork, you have to specify DNS policy
	// explicitly to 'ClusterFirstWithHostNet'.
	// +optional
	DNSPolicy DNSPolicy `json:"dnsPolicy,omitempty" protobuf:"bytes,6,opt,name=dnsPolicy,casttype=DNSPolicy"` // want "optionalfields: field DNSPolicy should be a pointer."
	// NodeSelector is a selector which must be true for the pod to fit on a node. // want "commentstart: godoc for field PodSpec.NodeSelector should start with 'nodeSelector ...'"
	// Selector which must match a node's labels for the pod to be scheduled on that node.
	// More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	// +optional
	// +mapType=atomic
	NodeSelector map[string]string `json:"nodeSelector,omitempty" protobuf:"bytes,7,rep,name=nodeSelector"`

	// ServiceAccountName is the name of the ServiceAccount to use to run this pod. // want "commentstart: godoc for field PodSpec.ServiceAccountName should start with 'serviceAccountName ...'"
	// More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	// +optional
	ServiceAccountName string `json:"serviceAccountName,omitempty" protobuf:"bytes,8,opt,name=serviceAccountName"` // want "optionalfields: field ServiceAccountName should be a pointer."
	// DeprecatedServiceAccount is a deprecated alias for ServiceAccountName. // want "commentstart: godoc for field PodSpec.DeprecatedServiceAccount should start with 'serviceAccount ...'"
	// Deprecated: Use serviceAccountName instead.
	// +k8s:conversion-gen=false
	// +optional
	DeprecatedServiceAccount string `json:"serviceAccount,omitempty" protobuf:"bytes,9,opt,name=serviceAccount"` // want "optionalfields: field DeprecatedServiceAccount should be a pointer."
	// AutomountServiceAccountToken indicates whether a service account token should be automatically mounted. // want "commentstart: godoc for field PodSpec.AutomountServiceAccountToken should start with 'automountServiceAccountToken ...'"
	// +optional
	AutomountServiceAccountToken *bool `json:"automountServiceAccountToken,omitempty" protobuf:"varint,21,opt,name=automountServiceAccountToken"`

	// NodeName indicates in which node this pod is scheduled. // want "commentstart: godoc for field PodSpec.NodeName should start with 'nodeName ...'"
	// If empty, this pod is a candidate for scheduling by the scheduler defined in schedulerName.
	// Once this field is set, the kubelet for this node becomes responsible for the lifecycle of this pod.
	// This field should not be used to express a desire for the pod to be scheduled on a specific node.
	// https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodename
	// +optional
	NodeName string `json:"nodeName,omitempty" protobuf:"bytes,10,opt,name=nodeName"` // want "optionalfields: field NodeName should be a pointer."
	// Host networking requested for this pod. Use the host's network namespace. // want "commentstart: godoc for field PodSpec.HostNetwork should start with 'hostNetwork ...'"
	// When using HostNetwork you should specify ports so the scheduler is aware.
	// When `hostNetwork` is true, specified `hostPort` fields in port definitions must match `containerPort`,
	// and unspecified `hostPort` fields in port definitions are defaulted to match `containerPort`.
	// Default to false.
	// +k8s:conversion-gen=false
	// +optional
	HostNetwork bool `json:"hostNetwork,omitempty" protobuf:"varint,11,opt,name=hostNetwork"` // want "optionalfields: field HostNetwork should be a pointer."
	// Use the host's pid namespace. // want "commentstart: godoc for field PodSpec.HostPID should start with 'hostPID ...'"
	// Optional: Default to false.
	// +k8s:conversion-gen=false
	// +optional
	HostPID bool `json:"hostPID,omitempty" protobuf:"varint,12,opt,name=hostPID"` // want "optionalfields: field HostPID should be a pointer."
	// Use the host's ipc namespace. // want "commentstart: godoc for field PodSpec.HostIPC should start with 'hostIPC ...'"
	// Optional: Default to false.
	// +k8s:conversion-gen=false
	// +optional
	HostIPC bool `json:"hostIPC,omitempty" protobuf:"varint,13,opt,name=hostIPC"` // want "optionalfields: field HostIPC should be a pointer."
	// Share a single process namespace between all of the containers in a pod. // want "commentstart: godoc for field PodSpec.ShareProcessNamespace should start with 'shareProcessNamespace ...'"
	// When this is set containers will be able to view and signal processes from other containers
	// in the same pod, and the first process in each container will not be assigned PID 1.
	// HostPID and ShareProcessNamespace cannot both be set.
	// Optional: Default to false.
	// +k8s:conversion-gen=false
	// +optional
	ShareProcessNamespace *bool `json:"shareProcessNamespace,omitempty" protobuf:"varint,27,opt,name=shareProcessNamespace"`
	// SecurityContext holds pod-level security attributes and common container settings. // want "commentstart: godoc for field PodSpec.SecurityContext should start with 'securityContext ...'"
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	SecurityContext *PodSecurityContext `json:"securityContext,omitempty" protobuf:"bytes,14,opt,name=securityContext"`
	// ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. // want "commentstart: godoc for field PodSpec.ImagePullSecrets should start with 'imagePullSecrets ...'"
	// If specified, these secrets will be passed to individual puller implementations for them to use.
	// More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=name
	ImagePullSecrets []LocalObjectReference `json:"imagePullSecrets,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,15,rep,name=imagePullSecrets"` // want "arrayofstruct: PodSpec.ImagePullSecrets is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// Specifies the hostname of the Pod // want "commentstart: godoc for field PodSpec.Hostname should start with 'hostname ...'"
	// If not specified, the pod's hostname will be set to a system-defined value.
	// +optional
	Hostname string `json:"hostname,omitempty" protobuf:"bytes,16,opt,name=hostname"` // want "optionalfields: field Hostname should be a pointer."
	// If specified, the fully qualified Pod hostname will be "<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>". // want "commentstart: godoc for field PodSpec.Subdomain should start with 'subdomain ...'"
	// If not specified, the pod will not have a domainname at all.
	// +optional
	Subdomain string `json:"subdomain,omitempty" protobuf:"bytes,17,opt,name=subdomain"` // want "optionalfields: field Subdomain should be a pointer."
	// If specified, the pod's scheduling constraints // want "commentstart: godoc for field PodSpec.Affinity should start with 'affinity ...'"
	// +optional
	Affinity *Affinity `json:"affinity,omitempty" protobuf:"bytes,18,opt,name=affinity"`
	// If specified, the pod will be dispatched by specified scheduler. // want "commentstart: godoc for field PodSpec.SchedulerName should start with 'schedulerName ...'"
	// If not specified, the pod will be dispatched by default scheduler.
	// +optional
	SchedulerName string `json:"schedulerName,omitempty" protobuf:"bytes,19,opt,name=schedulerName"` // want "optionalfields: field SchedulerName should be a pointer."
	// If specified, the pod's tolerations. // want "commentstart: godoc for field PodSpec.Tolerations should start with 'tolerations ...'"
	// +optional
	// +listType=atomic
	Tolerations []Toleration `json:"tolerations,omitempty" protobuf:"bytes,22,opt,name=tolerations"` // want "arrayofstruct: PodSpec.Tolerations is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// HostAliases is an optional list of hosts and IPs that will be injected into the pod's hosts // want "commentstart: godoc for field PodSpec.HostAliases should start with 'hostAliases ...'"
	// file if specified.
	// +optional
	// +patchMergeKey=ip
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=ip
	HostAliases []HostAlias `json:"hostAliases,omitempty" patchStrategy:"merge" patchMergeKey:"ip" protobuf:"bytes,23,rep,name=hostAliases"`
	// If specified, indicates the pod's priority. "system-node-critical" and // want "commentstart: godoc for field PodSpec.PriorityClassName should start with 'priorityClassName ...'"
	// "system-cluster-critical" are two special keywords which indicate the
	// highest priorities with the former being the highest priority. Any other
	// name must be defined by creating a PriorityClass object with that name.
	// If not specified, the pod priority will be default or zero if there is no
	// default.
	// +optional
	PriorityClassName string `json:"priorityClassName,omitempty" protobuf:"bytes,24,opt,name=priorityClassName"` // want "optionalfields: field PriorityClassName should be a pointer."
	// The priority value. Various system components use this field to find the // want "commentstart: godoc for field PodSpec.Priority should start with 'priority ...'"
	// priority of the pod. When Priority Admission Controller is enabled, it
	// prevents users from setting this field. The admission controller populates
	// this field from PriorityClassName.
	// The higher the value, the higher the priority.
	// +optional
	Priority *int32 `json:"priority,omitempty" protobuf:"bytes,25,opt,name=priority"`
	// Specifies the DNS parameters of a pod. // want "commentstart: godoc for field PodSpec.DNSConfig should start with 'dnsConfig ...'"
	// Parameters specified here will be merged to the generated DNS
	// configuration based on DNSPolicy.
	// +optional
	DNSConfig *PodDNSConfig `json:"dnsConfig,omitempty" protobuf:"bytes,26,opt,name=dnsConfig"`
	// If specified, all readiness gates will be evaluated for pod readiness. // want "commentstart: godoc for field PodSpec.ReadinessGates should start with 'readinessGates ...'"
	// A pod is ready when all its containers are ready AND
	// all conditions specified in the readiness gates have status equal to "True"
	// More info: https://git.k8s.io/enhancements/keps/sig-network/580-pod-readiness-gates
	// +optional
	// +listType=atomic
	ReadinessGates []PodReadinessGate `json:"readinessGates,omitempty" protobuf:"bytes,28,opt,name=readinessGates"` // want "arrayofstruct: PodSpec.ReadinessGates is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group, which should be used // want "commentstart: godoc for field PodSpec.RuntimeClassName should start with 'runtimeClassName ...'"
	// to run this pod.  If no RuntimeClass resource matches the named class, the pod will not be run.
	// If unset or empty, the "legacy" RuntimeClass will be used, which is an implicit class with an
	// empty definition that uses the default runtime handler.
	// More info: https://git.k8s.io/enhancements/keps/sig-node/585-runtime-class
	// +optional
	RuntimeClassName *string `json:"runtimeClassName,omitempty" protobuf:"bytes,29,opt,name=runtimeClassName"`
	// EnableServiceLinks indicates whether information about services should be injected into pod's // want "commentstart: godoc for field PodSpec.EnableServiceLinks should start with 'enableServiceLinks ...'"
	// environment variables, matching the syntax of Docker links.
	// Optional: Defaults to true.
	// +optional
	EnableServiceLinks *bool `json:"enableServiceLinks,omitempty" protobuf:"varint,30,opt,name=enableServiceLinks"`
	// PreemptionPolicy is the Policy for preempting pods with lower priority. // want "commentstart: godoc for field PodSpec.PreemptionPolicy should start with 'preemptionPolicy ...'"
	// One of Never, PreemptLowerPriority.
	// Defaults to PreemptLowerPriority if unset.
	// +optional
	PreemptionPolicy *PreemptionPolicy `json:"preemptionPolicy,omitempty" protobuf:"bytes,31,opt,name=preemptionPolicy"`
	// Overhead represents the resource overhead associated with running a pod for a given RuntimeClass.
	// This field will be autopopulated at admission time by the RuntimeClass admission controller. If
	// the RuntimeClass admission controller is enabled, overhead must not be set in Pod create requests.
	// The RuntimeClass admission controller will reject Pod create requests which have the overhead already
	// set. If RuntimeClass is configured and selected in the PodSpec, Overhead will be set to the value
	// defined in the corresponding RuntimeClass, otherwise it will remain unset and treated as zero.
	// More info: https://git.k8s.io/enhancements/keps/sig-node/688-pod-overhead/README.md
	// +optional
	// Overhead ResourceList `json:"overhead,omitempty" protobuf:"bytes,32,opt,name=overhead"`

	// TopologySpreadConstraints describes how a group of pods ought to spread across topology // want "commentstart: godoc for field PodSpec.TopologySpreadConstraints should start with 'topologySpreadConstraints ...'"
	// domains. Scheduler will schedule pods in a way which abides by the constraints.
	// All topologySpreadConstraints are ANDed.
	// +optional
	// +patchMergeKey=topologyKey
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=topologyKey
	// +listMapKey=whenUnsatisfiable
	TopologySpreadConstraints []TopologySpreadConstraint `json:"topologySpreadConstraints,omitempty" patchStrategy:"merge" patchMergeKey:"topologyKey" protobuf:"bytes,33,opt,name=topologySpreadConstraints"` // want "arrayofstruct: PodSpec.TopologySpreadConstraints is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// If true the pod's hostname will be configured as the pod's FQDN, rather than the leaf name (the default). // want "commentstart: godoc for field PodSpec.SetHostnameAsFQDN should start with 'setHostnameAsFQDN ...'"
	// In Linux containers, this means setting the FQDN in the hostname field of the kernel (the nodename field of struct utsname).
	// In Windows containers, this means setting the registry value of hostname for the registry key HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Services\\Tcpip\\Parameters to FQDN.
	// If a pod does not have FQDN, this has no effect.
	// Default to false.
	// +optional
	SetHostnameAsFQDN *bool `json:"setHostnameAsFQDN,omitempty" protobuf:"varint,35,opt,name=setHostnameAsFQDN"`
	// Specifies the OS of the containers in the pod. // want "commentstart: godoc for field PodSpec.OS should start with 'os ...'"
	// Some pod and container fields are restricted if this is set.
	//
	// If the OS field is set to linux, the following fields must be unset:
	// -securityContext.windowsOptions
	//
	// If the OS field is set to windows, following fields must be unset:
	// - spec.hostPID
	// - spec.hostIPC
	// - spec.hostUsers
	// - spec.resources
	// - spec.securityContext.appArmorProfile
	// - spec.securityContext.seLinuxOptions
	// - spec.securityContext.seccompProfile
	// - spec.securityContext.fsGroup
	// - spec.securityContext.fsGroupChangePolicy
	// - spec.securityContext.sysctls
	// - spec.shareProcessNamespace
	// - spec.securityContext.runAsUser
	// - spec.securityContext.runAsGroup
	// - spec.securityContext.supplementalGroups
	// - spec.securityContext.supplementalGroupsPolicy
	// - spec.containers[*].securityContext.appArmorProfile
	// - spec.containers[*].securityContext.seLinuxOptions
	// - spec.containers[*].securityContext.seccompProfile
	// - spec.containers[*].securityContext.capabilities
	// - spec.containers[*].securityContext.readOnlyRootFilesystem
	// - spec.containers[*].securityContext.privileged
	// - spec.containers[*].securityContext.allowPrivilegeEscalation
	// - spec.containers[*].securityContext.procMount
	// - spec.containers[*].securityContext.runAsUser
	// - spec.containers[*].securityContext.runAsGroup
	// +optional
	OS *PodOS `json:"os,omitempty" protobuf:"bytes,36,opt,name=os"`

	// Use the host's user namespace. // want "commentstart: godoc for field PodSpec.HostUsers should start with 'hostUsers ...'"
	// Optional: Default to true.
	// If set to true or not present, the pod will be run in the host user namespace, useful
	// for when the pod needs a feature only available to the host user namespace, such as
	// loading a kernel module with CAP_SYS_MODULE.
	// When set to false, a new userns is created for the pod. Setting false is useful for
	// mitigating container breakout vulnerabilities even allowing users to run their
	// containers as root without actually having root privileges on the host.
	// This field is alpha-level and is only honored by servers that enable the UserNamespacesSupport feature.
	// +k8s:conversion-gen=false
	// +optional
	HostUsers *bool `json:"hostUsers,omitempty" protobuf:"bytes,37,opt,name=hostUsers"`

	// SchedulingGates is an opaque list of values that if specified will block scheduling the pod. // want "commentstart: godoc for field PodSpec.SchedulingGates should start with 'schedulingGates ...'"
	// If schedulingGates is not empty, the pod will stay in the SchedulingGated state and the
	// scheduler will not attempt to schedule the pod.
	//
	// SchedulingGates can only be set at pod creation time, and be removed only afterwards.
	//
	// +patchMergeKey=name
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=name
	// +optional
	SchedulingGates []PodSchedulingGate `json:"schedulingGates,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,38,opt,name=schedulingGates"` // want "arrayofstruct: PodSpec.SchedulingGates is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// ResourceClaims defines which ResourceClaims must be allocated // want "commentstart: godoc for field PodSpec.ResourceClaims should start with 'resourceClaims ...'"
	// and reserved before the Pod is allowed to start. The resources
	// will be made available to those containers which consume them
	// by name.
	//
	// This is a stable field but requires that the
	// DynamicResourceAllocation feature gate is enabled.
	//
	// This field is immutable.
	//
	// +patchMergeKey=name
	// +patchStrategy=merge,retainKeys
	// +listType=map
	// +listMapKey=name
	// +featureGate=DynamicResourceAllocation
	// +optional
	ResourceClaims []PodResourceClaim `json:"resourceClaims,omitempty" patchStrategy:"merge,retainKeys" patchMergeKey:"name" protobuf:"bytes,39,rep,name=resourceClaims"` // want "arrayofstruct: PodSpec.ResourceClaims is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
	// Resources is the total amount of CPU and Memory resources required by all // want "commentstart: godoc for field PodSpec.Resources should start with 'resources ...'"
	// containers in the pod. It supports specifying Requests and Limits for
	// "cpu", "memory" and "hugepages-" resource names only. ResourceClaims are not supported.
	//
	// This field enables fine-grained control over resource allocation for the
	// entire pod, allowing resource sharing among containers in a pod.
	// TODO: For beta graduation, expand this comment with a detailed explanation.
	//
	// This is an alpha field and requires enabling the PodLevelResources feature
	// gate.
	//
	// +featureGate=PodLevelResources
	// +optional
	Resources *ResourceRequirements `json:"resources,omitempty" protobuf:"bytes,40,opt,name=resources"`
	// HostnameOverride specifies an explicit override for the pod's hostname as perceived by the pod. // want "commentstart: godoc for field PodSpec.HostnameOverride should start with 'hostnameOverride ...'"
	// This field only specifies the pod's hostname and does not affect its DNS records.
	// When this field is set to a non-empty string:
	// - It takes precedence over the values set in `hostname` and `subdomain`.
	// - The Pod's hostname will be set to this value.
	// - `setHostnameAsFQDN` must be nil or set to false.
	// - `hostNetwork` must be set to false.
	//
	// This field must be a valid DNS subdomain as defined in RFC 1123 and contain at most 64 characters.
	// Requires the HostnameOverride feature gate to be enabled.
	//
	// +featureGate=HostnameOverride
	// +optional
	HostnameOverride *string `json:"hostnameOverride,omitempty" protobuf:"bytes,41,opt,name=hostnameOverride"`
}

// PodDNSConfig defines the DNS parameters of a pod in addition to
// those generated from DNSPolicy.
type PodDNSConfig struct {
	// A list of DNS name server IP addresses. // want "commentstart: godoc for field PodDNSConfig.Nameservers should start with 'nameservers ...'"
	// This will be appended to the base nameservers generated from DNSPolicy.
	// Duplicated nameservers will be removed.
	// +optional
	// +listType=atomic
	Nameservers []string `json:"nameservers,omitempty" protobuf:"bytes,1,rep,name=nameservers"`
	// A list of DNS search domains for host-name lookup. // want "commentstart: godoc for field PodDNSConfig.Searches should start with 'searches ...'"
	// This will be appended to the base search paths generated from DNSPolicy.
	// Duplicated search paths will be removed.
	// +optional
	// +listType=atomic
	Searches []string `json:"searches,omitempty" protobuf:"bytes,2,rep,name=searches"`
	// A list of DNS resolver options. // want "commentstart: godoc for field PodDNSConfig.Options should start with 'options ...'"
	// This will be merged with the base options generated from DNSPolicy.
	// Duplicated entries will be removed. Resolution options given in Options
	// will override those that appear in the base DNSPolicy.
	// +optional
	// +listType=atomic
	Options []PodDNSConfigOption `json:"options,omitempty" protobuf:"bytes,3,rep,name=options"` // want "arrayofstruct: PodDNSConfig.Options is an array of structs, but the struct has no required fields. At least one field should be marked as required to prevent ambiguous YAML configurations"
}

// PodDNSConfigOption defines DNS resolver options of a pod.
type PodDNSConfigOption struct {
	// Name is this DNS resolver option's name. // want "commentstart: godoc for field PodDNSConfigOption.Name should start with 'name ...'"
	// Required.
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"` // want "optionalorrequired: field PodDNSConfigOption.Name must be marked as optional or required"
	// Value is this DNS resolver option's value. // want "commentstart: godoc for field PodDNSConfigOption.Value should start with 'value ...'"
	// +optional
	Value *string `json:"value,omitempty" protobuf:"bytes,2,opt,name=value"`
}

// PodReadinessGate contains the reference to a pod condition
type PodReadinessGate struct {
	// ConditionType refers to a condition in the pod's condition list with matching type. // want "commentstart: godoc for field PodReadinessGate.ConditionType should start with 'conditionType ...'"
	ConditionType PodConditionType `json:"conditionType" protobuf:"bytes,1,opt,name=conditionType,casttype=PodConditionType"` // want "optionalorrequired: field PodReadinessGate.ConditionType must be marked as optional or required"
}

// PodConditionType is a valid value for PodCondition.Type
// +kubebuilder:validation:Enum=ContainersReady;Initialized;Ready;PodScheduled;DisruptionTarget;PodReadyToStartContainers;PodResizePending;PodResizeInProgress
type PodConditionType string

// These are built-in conditions of pod. An application may use a custom condition not listed here.
const (
	// ContainersReady indicates whether all containers in the pod are ready.
	ContainersReady PodConditionType = "ContainersReady"
	// PodInitialized means that all init containers in the pod have started successfully.
	PodInitialized PodConditionType = "Initialized"
	// PodReady means the pod is able to service requests and should be added to the
	// load balancing pools of all matching services.
	PodReady PodConditionType = "Ready"
	// PodScheduled represents status of the scheduling process for this pod.
	PodScheduled PodConditionType = "PodScheduled"
	// DisruptionTarget indicates the pod is about to be terminated due to a
	// disruption (such as preemption, eviction API or garbage-collection).
	DisruptionTarget PodConditionType = "DisruptionTarget"
	// PodReadyToStartContainers pod sandbox is successfully configured and
	// the pod is ready to launch containers.
	PodReadyToStartContainers PodConditionType = "PodReadyToStartContainers"
	// PodResizePending indicates that the pod has been resized, but kubelet has not
	// yet allocated the resources. If both PodResizePending and PodResizeInProgress
	// are set, it means that a new resize was requested in the middle of a previous
	// pod resize that is still in progress.
	PodResizePending PodConditionType = "PodResizePending"
	// PodResizeInProgress indicates that a resize is in progress, and is present whenever
	// the Kubelet has allocated resources for the resize, but has not yet actuated all of
	// the required changes.
	// If both PodResizePending and PodResizeInProgress are set, it means that a new resize was
	// requested in the middle of a previous pod resize that is still in progress.
	PodResizeInProgress PodConditionType = "PodResizeInProgress"
)

// +kubebuilder:validation:Enum=linux;windows
// OSName is the set of OS'es that can be used in OS.
type OSName string

// These are valid values for OSName
const (
	Linux   OSName = "linux"
	Windows OSName = "windows"
)

// PodOS defines the OS parameters of a pod.
type PodOS struct {
	// Name is the name of the operating system. The currently supported values are linux and windows. // want "commentstart: godoc for field PodOS.Name should start with 'name ...'"
	// Additional value may be defined in future and can be one of:
	// https://github.com/opencontainers/runtime-spec/blob/master/config.md#platform-specific-configuration
	// Clients should expect to handle additional values and treat unrecognized values in this field as os: null
	Name OSName `json:"name" protobuf:"bytes,1,opt,name=name"` // want "optionalorrequired: field PodOS.Name must be marked as optional or required"
}

// PodSchedulingGate is associated to a Pod to guard its scheduling.
type PodSchedulingGate struct {
	// Name of the scheduling gate. // want "commentstart: godoc for field PodSchedulingGate.Name should start with 'name ...'"
	// Each scheduling gate must have a unique name field.
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"` // want "optionalorrequired: field PodSchedulingGate.Name must be marked as optional or required"
}

// PodResourceClaim references exactly one ResourceClaim, either directly
// or by naming a ResourceClaimTemplate which is then turned into a ResourceClaim
// for the pod.
//
// It adds a name to it that uniquely identifies the ResourceClaim inside the Pod.
// Containers that need access to the ResourceClaim reference it with this name.
type PodResourceClaim struct {
	// Name uniquely identifies this resource claim inside the pod. // want "commentstart: godoc for field PodResourceClaim.Name should start with 'name ...'"
	// This must be a DNS_LABEL.
	Name string `json:"name" protobuf:"bytes,1,name=name"` // want "optionalorrequired: field PodResourceClaim.Name must be marked as optional or required"

	// Source is tombstoned since Kubernetes 1.31 where it got replaced with
	// the inlined fields below.
	//
	// Source ClaimSource `json:"source,omitempty" protobuf:"bytes,2,name=source"`

	// ResourceClaimName is the name of a ResourceClaim object in the same // want "commentstart: godoc for field PodResourceClaim.ResourceClaimName should start with 'resourceClaimName ...'"
	// namespace as this pod.
	//
	// Exactly one of ResourceClaimName and ResourceClaimTemplateName must
	// be set.
	ResourceClaimName *string `json:"resourceClaimName,omitempty" protobuf:"bytes,3,opt,name=resourceClaimName"` // want "optionalorrequired: field PodResourceClaim.ResourceClaimName must be marked as optional or required"

	// ResourceClaimTemplateName is the name of a ResourceClaimTemplate // want "commentstart: godoc for field PodResourceClaim.ResourceClaimTemplateName should start with 'resourceClaimTemplateName ...'"
	// object in the same namespace as this pod.
	//
	// The template will be used to create a new ResourceClaim, which will
	// be bound to this pod. When this pod is deleted, the ResourceClaim
	// will also be deleted. The pod name and resource name, along with a
	// generated component, will be used to form a unique name for the
	// ResourceClaim, which will be recorded in pod.status.resourceClaimStatuses.
	//
	// This field is immutable and no changes will be made to the
	// corresponding ResourceClaim by the control plane after creating the
	// ResourceClaim.
	//
	// Exactly one of ResourceClaimName and ResourceClaimTemplateName must
	// be set.
	ResourceClaimTemplateName *string `json:"resourceClaimTemplateName,omitempty" protobuf:"bytes,4,opt,name=resourceClaimTemplateName"` // want "optionalorrequired: field PodResourceClaim.ResourceClaimTemplateName must be marked as optional or required"
}
