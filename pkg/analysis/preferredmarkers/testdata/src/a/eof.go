package a

// Type at end of file without trailing newline
// +kubebuilder:validation:Optional
type EndOfFileType string // want `type EndOfFileType uses marker "kubebuilder:validation:Optional", should use preferred marker "k8s:optional" instead`