package net

import (
	"fmt"
	"testing"
)

func TestMessageFilter(t *testing.T) {
	msgFilter := NewMessageFilter(10)
	for i := uint64(0); i < 1024; i++ {
		msgFilter.CheckUnique(i)
		if len(msgFilter.uniqueMessageSet) >= 100 {
			t.Errorf("invalid check unique exceeded max count!")
		}
	}

	if msgFilter.CheckUnique(1023) {
		t.Errorf("invalid check unique 1023!")
	}

	fmt.Printf("test message filter success: %d\n", len(msgFilter.uniqueMessageSet))
}
