package net

import (
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	test := &BroadcastParam{}
	test.Type = 1

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
