package net

import (
	"net"
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
