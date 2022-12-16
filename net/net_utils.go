package net

import (
	"net"
)

type EventItem struct {
	conn       net.Conn
	fd         int
	msgBuf     *MsgBuffer
	removed    bool
	packageLen int
	parentConn *Connection
}