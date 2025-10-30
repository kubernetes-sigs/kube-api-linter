package a

type EOFFieldTest struct {
	// Field at end of file without trailing newline
	// +kubebuilder:validation:Optional
	EndOfFileField string `json:"endOfFileField"` // want `field EndOfFileField uses marker "kubebuilder:validation:Optional", should use preferred marker "k8s:optional" instead`
}