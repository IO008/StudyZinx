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

	// connection handle API
	handleAPI ziface.HandFunc

	// connection exit channel
	ExitBuffChan chan bool
}

func NewConnection(conn *net.TCPConn, connID uint32, callback_api ziface.HandFunc) *Connection {
	return &Connection{
		Conn:         conn,
		ConnId:       connID,
		isClosed:     false,
		handleAPI:    callback_api,
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
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err ", err)
			c.ExitBuffChan <- true
			continue
		}

		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("connID ", c.ConnId, " handle is error")
			c.ExitBuffChan <- true
			return
		}

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
