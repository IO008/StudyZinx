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
}

// Define an interface for unified processing of connection services
type HandFunc func(*net.TCPConn, []byte, int) error
