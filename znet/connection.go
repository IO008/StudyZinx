package znet

import (
	"StudyZinx/ziface"
	"fmt"
	"net"
)

type Connection struct {
	// current connection socket
	Conn *net.TCPConn
	// connection ID(Session ID) is global unique ID
	ConnId uint32
	// current connection close status
	isClosed bool

	// connection handler method router
	Router ziface.IRouter

	// connection exit channel
	ExitBuffChan chan bool
}

func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	return &Connection{
		Conn:         conn,
		ConnId:       connID,
		isClosed:     false,
		Router:       router,
		ExitBuffChan: make(chan bool, 1),
	}
}

// Goroutine that handles conn reading data
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running")
	defer fmt.Println(c.Conn.RemoteAddr().String(), " conn reader exit!")
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err ", err)
			c.ExitBuffChan <- true
			continue
		}

		req := Request{
			conn: c,
			data: buf,
		}

		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}

// start connection let current connection begin work
func (c *Connection) Start() {
	go c.StartReader()

	for {
		select {
		case <-c.ExitBuffChan:
			return // No more blocking when get exit signal
		}
	}
}

// stop connection
func (c *Connection) Stop() {
	if c.isClosed {
		return
	}
	c.isClosed = true

	// TODO callback func when connection is closed

	c.Conn.Close()

	close(c.ExitBuffChan)
}

func (c *Connection) GetTcpConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnId
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
