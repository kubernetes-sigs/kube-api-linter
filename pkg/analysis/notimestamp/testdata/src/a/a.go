package a

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NoTimeStampTestStruct struct {
	// +optional
	TimeStamp *time.Time `json:"timeStamp,omitempty"` // want "field TimeStamp: prefer use of the term time over timestamp"

	// +optional
	Timestamp *time.Time `json:"timestamp,omitempty"` // want "field Timestamp: prefer use of the term time over timestamp"

	// +optional
	FooTimeStamp *time.Time `json:"fooTimeStamp,omitempty"` // want "field FooTimeStamp: prefer use of the term time over timestamp"

	// +optional
	FootimeStamp *time.Time `json:"footimeStamp,omitempty"` // want "field FootimeStamp: prefer use of the term time over timestamp"

	// +optional
	BarTimestamp *time.Time `json:"barTimestamp,omitempty"` // want "field BarTimestamp: prefer use of the term time over timestamp"

	// +optional
	FootimestampBar *time.Time `json:"fooTimestampBar,omitempty"` // want "field FootimestampBar: prefer use of the term time over timestamp"

	// +optional
	FooTimestampBarTimeStamp *time.Time `json:"fooTimestampBarTimeStamp,omitempty"` // want "field FooTimestampBarTimeStamp: prefer use of the term time over timestamp"

	// +optional
	MetaTimeStamp *metav1.Time `json:"metaTimeStamp,omitempty"` // want "field MetaTimeStamp: prefer use of the term time over timestamp"
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
