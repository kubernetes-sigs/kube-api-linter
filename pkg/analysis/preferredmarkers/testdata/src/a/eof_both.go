package a

// Type at end with both preferred and equivalent (no trailing newline)
// +k8s:optional
// +kubebuilder:validation:Optional
type EOFTypeWithBoth string // want `type EOFTypeWithBoth uses marker "kubebuilder:validation:Optional", should use preferred marker "k8s:optional" instead`