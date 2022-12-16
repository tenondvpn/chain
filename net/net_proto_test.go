package net

import (
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	test := &BroadcastParam{
		Type:              1,
		NeighborCount:     2,
		StopTimes:         3,
		HopLimit:          4,
		LayerLeft:         4,
		LayerRight:        4,
		Overlap:           4,
		HopToLayer:        4,
		Header:            "",
		Body:              "",
		NetCrossed:        false,
		Bloomfilter:       []uint64{1, 2, 3},
		EvilRate:          4,
		IgnBloomfilterHop: 4,
	}

	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newTest := &BroadcastParam{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if test.GetType() != newTest.GetType() {
		log.Fatalf("data mismatch %q != %q", test.GetType(), newTest.GetType())
	}
}
