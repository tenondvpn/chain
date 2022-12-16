package net

import (
	"sync"

	queue "github.com/smallnest/queue"
)

type MessageFilter struct {
	uniqueMessageSet map[uint64]bool
	addedQueue       *queue.CQueue
	lockMutex        sync.Mutex
	maxSavedCount    int
}

func NewMessageFilter(maxCount int) *MessageFilter {
	handler := new(MessageFilter)
	handler.uniqueMessageSet = make(map[uint64]bool)
	handler.addedQueue = queue.NewCQueue()
	handler.maxSavedCount = maxCount
	return handler
}

func (f *MessageFilter) CheckUnique(msgHash uint64) bool {
	f.lockMutex.Lock()
	defer f.lockMutex.Unlock()
	if _, ok := f.uniqueMessageSet[msgHash]; ok {
		return false
	}

	f.uniqueMessageSet[msgHash] = true
	f.addedQueue.Enqueue(msgHash)
	if len(f.uniqueMessageSet) >= f.maxSavedCount {
		headMsgHash := f.addedQueue.Dequeue()
		tmpheadMsgHash, ok := headMsgHash.(uint64)
		if !ok {
			panic("transfer data msg buffer failed!")
		}

		delete(f.uniqueMessageSet, tmpheadMsgHash)
	}

	return true
}
