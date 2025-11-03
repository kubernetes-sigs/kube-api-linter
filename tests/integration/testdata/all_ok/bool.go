package a

type TestBools struct {
	// bool is a pointer to a bool.
	// +required
	Bool *bool `json:"bool,omitempty"`
}
