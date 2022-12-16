package net

import (
	"encoding/binary"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	tcpConnection    *EventItem
	allreceiedCount  int64 = 0
	btime            int64 = time.Now().UnixMicro()
	ttime            int64 = time.Now().UnixMicro()
	testMessageCount int64 = 10000000
	testThreadCount  int   = 4
)

func onMsgForTest(qMsg *QueueMsg) bool {
	t := binary.BigEndian.Uint32(qMsg.data[4:8])
	if t == 0 { // from client
		binary.BigEndian.PutUint32(qMsg.data[4:], uint32(1))
		(*qMsg.conn).Write(qMsg.data)
	} else {
		atomic.AddInt64(&allreceiedCount, 1)
		etime := time.Now().UnixMicro()
		if etime-btime >= 3000000 {
			fmt.Printf("all: %d, avg: %d\n", allreceiedCount, allreceiedCount*1000000/(etime-ttime))
			btime = etime
		}
	}

	return true
}

func SendMessage() {
	tcpClient := NewTcpClient(onMsgForTest)
	tcpConnection = tcpClient.ConnectServer("127.0.0.1:8990")
	if tcpConnection == nil {
		fmt.Println("connect error.")
		return
	}

	data := []byte("testsdadfasdfasdfadfd")
	val := make([]byte, 12)
	binary.BigEndian.PutUint32(val[0:], uint32(12+len(data)))
	binary.BigEndian.PutUint32(val[4:], 0)
	binary.BigEndian.PutUint32(val[8:], 0)
	val = append(val, data...)
	for i := int64(0); i < testMessageCount; i++ {
		n, err := tcpClient.Send(tcpConnection, val)
		if err != nil {
			fmt.Printf("send message failed: %d, %v", n, err)
			break
		}
	}

	for {
		if allreceiedCount >= testMessageCount*int64(testThreadCount) {
			break
		}

		time.Sleep(time.Microsecond * 100)
	}

	tcpClient.Close(tcpConnection)
	fmt.Println("close success")
}

func TestTcp(t *testing.T) {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:   true,
		TimestampFormat: "2006-01-02 15:03:04",
	})

	svr := NewTcpServer(onMsgForTest)
	go svr.StartServer("127.0.0.1:8990")
	time.Sleep(time.Second * 1)

	var wg sync.WaitGroup

	for i := 0; i < testThreadCount; i++ {
		wg.Add(1)
		go SendMessage()
	}

	wg.Wait()
	svr.StopServer()
	fmt.Println("all over")
}
