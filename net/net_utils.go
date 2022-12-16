package net

import (
	"net"
	"reflect"
	"time"
	"unsafe"

	queue "github.com/smallnest/queue"
)

var (
	MaxPackageSize = 1024 * 1024 * 2
	IoThreadCount  = 4
	IoPoolSize     = 32
	BuffPoolSize   = IoThreadCount * (IoPoolSize + 1)
	msgHandler     *MessageHandler
)

type EndPoint struct {
	ip   string
	port uint16
}

type Connection struct {
	endpoint          *EndPoint
	cliEv             *EventItem
	timeout           int64
	rxId              uint16
	ackedIndex        uint64
	prevHeatbeatTimer int64
	connectValid      bool
	seqAckedIndex     uint64
	conn              *net.Conn
	syncedValid       bool
	seqIdSyncedValid  bool
}

type EventItem struct {
	conn       net.Conn
	fd         int
	msgBuf     *MsgBuffer
	removed    bool
	packageLen int
	parentConn *Connection
}

type QueueMsg struct {
	queueId  uint16
	msgType  uint16
	data     []byte
	conn     *net.Conn
	msgBuf   *MsgBuffer
	bufQueue *queue.LKQueue
}

type ConfigMsg struct {
	queueId   uint16
	msgType   uint16
	txCluster string
	data      []byte
}

type TcpCallback func(qMsg *QueueMsg) bool
type ReceiverCallback func(msgId uint64, data []byte) int

func TimeStampMilli() int64 {
	return time.Now().UnixNano() / (1000 * 1000)
}

func Bytes2Str(slice []byte) string {
	return *(*string)(unsafe.Pointer(&slice))
}

func Str2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
