package net

import (
	"testing"

	"github.com/golang/protobuf/proto"
)

func TestProto(t *testing.T) {
	test := &BroadcastParam{}
	var tt uint32 = 1
	test.NeighborCount = &tt

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
	if test.GetNeighborCount() != newTest.GetNeighborCount() {
		t.Errorf("data mismatch %q != %q", test.GetNeighborCount(), newTest.GetNeighborCount())
	}
}
