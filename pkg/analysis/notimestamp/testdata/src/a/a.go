package a

import "time"

type NoTimeStampTestStruct struct {
	// +optional
	TimeStamp *time.Time `json:"timeStamp,omitempty"` // want "field TimeStamp: prefer use of the term time over timestamp"
}

// DoNothing is used to check that the analyser doesn't report on methods.
func (NoTimeStampTestStruct) DoNothing() {}

type NoSubTimeStampTestStruct struct {
	// +optional
	FooTimeStamp *time.Time `json:"fooTimeStamp,omitempty"` // want "field FooTimeStamp: prefer use of the term time over timestamp"
}

type SerializedTimeStampTestStruct struct {
	// +optional
	FooTime *time.Time `json:"fooTime,omitempty"`
}
