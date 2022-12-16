package net

import (
	"github.com/sirupsen/logrus"
	queue "github.com/smallnest/queue"
)

type MessageHandler struct {
	threadCount int
}

func NewMessageHandler(threadCount int) *MessageHandler {
	handler := new(MessageHandler)
	handler.threadCount = threadCount
	return handler
}

func (h *MessageHandler) handleMessage(
	items []*EventItem,
	epoller *epoll,
	cb TcpCallback,
	bufQueue *queue.LKQueue) {
	for _, item := range items {
		if item == nil {
			break
		}

		for {
			n, buf, err := item.msgBuf.readFromReader()
			if err != nil {
				if err := epoller.Remove(item); err != nil {
					logrus.Errorf("failed to remove %v", err)
				}

				item.conn.Close()
				break
			}

			if n == 0 {
				break
			}

			if n > 0 {
				qMsg := &CallbackMessage{buf, &item.conn, item.msgBuf, bufQueue}
				hold := cb(qMsg)
				data := bufQueue.Dequeue()
				if data == nil {
					panic("get data from queue failed!")
				}
				tmpBuf, ok := data.(*MsgBuffer)
				if !ok {
					panic("transfer data msg buffer failed!")
				}

				tmpBuf.pkgLen = 0
				tmpBuf.start = 0
				tmpBuf.reader = item.msgBuf.reader
				exchangeBuf := item.msgBuf
				item.msgBuf = tmpBuf
				if hold {
					bufQueue.Enqueue(exchangeBuf)
				}

				break
			}
		}
	}
}
