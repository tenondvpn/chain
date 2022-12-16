package net

import (
	"testing"

	"github.com/golang/protobuf/proto"
)

func TestProto(t *testing.T) {
	test := &BroadcastParam{}
	var tt uint32 = 1
	test.Type = &tt

	data, err := proto.Marshal(test)
	if err != nil {
		t.Errorf("marshaling error: ", err)
	}
	newTest := &BroadcastParam{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		t.Errorf("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if test.GetType() != newTest.GetType() {
		t.Errorf("data mismatch %q != %q", test.GetType(), newTest.GetType())
	}
}
