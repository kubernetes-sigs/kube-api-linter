package a

import "time"

type NoPhaseTestStruct struct {
	// +optional
	Phase *string `json:"phase,omitempty"` // want "field Phase: phase fields are deprecated and discouraged. conditions should be used instead."
}

// DoNothing is used to check that the analyser doesn't report on methods.
func (NoPhaseTestStruct) DoNothing() {}

type NoSubPhaseTestStruct struct {
	// +optional
	FooPhase *string `json:"fooPhase,omitempty"` // want "field FooPhase: phase fields are deprecated and discouraged. conditions should be used instead."
}

type SerializedPhaseTeststruct struct {
	// +optional
	FooField *string `json:"fooPhase,omitempty"`
}

type NoTimeStampTestStruct struct {
	// +optional
	TimeStamp *time.Time `json:"timeStamp,omitempty"` // want "field TimeStamp: prefer use of the term 'time' over 'timestamp'"
}

// DoNothing is used to check that the analyser doesn't report on methods.
func (NoTimeStampTestStruct) DoNothing() {}

type NoSubTimeStampTestStruct struct {
	// +optional
	FooTimeStamp *time.Time `json:"fooTimeStamp,omitempty"` // want "field FooTimeStamp: prefer use of the term 'time' over 'timestamp'"
}

type SerializedTimeStampTestStruct struct {
	// +optional
	FooTime *time.Time `json:"fooTime,omitempty"`
}

type NoReferenceTestStruct struct {
	// +optional
	Reference *string `json:"reference,omitempty"` // want "field Reference: prefer use of the term 'ref' over 'reference'"
}

// DoNothing is used to check that the analyser doesn't report on methods.
func (NoReferenceTestStruct) DoNothing() {}

type NoReferenceSuffixTestStruct struct {
	// +optional
	FooReference *string `json:"fooReference,omitempty"` // want "field FooReference: prefer use of the term 'ref' over 'reference'"
}
