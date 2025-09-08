/*
This is a copy of the minimum amount of the original file to be able to test the nodurations linter.
*/

package v1

import "time"

// Duration is a wrapper around time.Duration which supports correct
// marshaling to YAML and JSON. In particular, it marshals into strings, which
// can be used as map keys in json.
type Duration struct {
	time.Duration `protobuf:"varint,1,opt,name=duration,casttype=time.Duration"`
}
