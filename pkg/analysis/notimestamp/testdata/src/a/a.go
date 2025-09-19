package a

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NoTimeStampTestStruct struct {
	// +optional
	TimeStamp *time.Time `json:"timeStamp,omitempty"` // want `naming convention "notimestamp": prefer use of the term 'time' over 'timestamp'`

	// +optional
	Timestamp *time.Time `json:"timestamp,omitempty"` // want `naming convention "notimestamp": prefer use of the term 'time' over 'timestamp'`

	// +optional
	FooTimeStamp *time.Time `json:"fooTimeStamp,omitempty"` // want `naming convention "notimestamp": prefer use of the term 'time' over 'timestamp'`

	// +optional
	FootimeStamp *time.Time `json:"footimeStamp,omitempty"` // want `naming convention "notimestamp": prefer use of the term 'time' over 'timestamp'`

	// +optional
	BarTimestamp *time.Time `json:"barTimestamp,omitempty"` // want `naming convention "notimestamp": prefer use of the term 'time' over 'timestamp'`

	// +optional
	FootimestampBar *time.Time `json:"fooTimestampBar,omitempty"` // want `naming convention "notimestamp": prefer use of the term 'time' over 'timestamp'`

	// +optional
	FooTimestampBarTimeStamp *time.Time `json:"fooTimestampBarTimeStamp,omitempty"` // want `naming convention "notimestamp": prefer use of the term 'time' over 'timestamp'`

	// +optional
	MetaTimeStamp *metav1.Time `json:"metaTimeStamp,omitempty"` // want `naming convention "notimestamp": prefer use of the term 'time' over 'timestamp'`
}

// DoNothing is used to check that the analyser doesn't report on methods.
func (NoTimeStampTestStruct) DoNothing() {}

type NoSubTimeStampTestStruct struct {
	// +optional
	FooTimeStamp *time.Time `json:"fooTimeStamp,omitempty"` // want `naming convention "notimestamp": prefer use of the term 'time' over 'timestamp'`
}

type SerializedTimeStampTestStruct struct {
	// +optional
	FooTime *time.Time `json:"fooTime,omitempty"`
}
