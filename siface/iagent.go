package siface

import "net"

type IAgent interface {
	Start()
	Stop()
	GetConn() net.Conn
	GetAgentId() uint32
	RemoteAddr()  net.Addr
	SendMsg(id byte,data []byte) error
	SetProperty(key string,value interface{})
	GetProperty(key string)(interface{},bool)
	DelProperty(key string)
}