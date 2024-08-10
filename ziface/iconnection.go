package ziface

import "net"

type IConnection interface {
	// start connection
	Start()

	// stop connection
	Stop()

	// get origin socket TCPConn from current connection
	GetTcpConnection() *net.TCPConn

	// get connection ID
	GetConnID() uint32

	// get client remote addr
	RemoteAddr() net.Addr

	SendMsg(msgId uint32, data []byte) error

	SendBuffMsg(msgId uint32, data []byte) error

	SetProperty(key string, value interface{})
	GetProperty(key string) (interface{}, error)
	RemoveProperty(key string)
}
