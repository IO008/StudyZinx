package znet

import (
	"StudyZinx/utils"
	"StudyZinx/ziface"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
)

type Connection struct {
	TcpServer ziface.IServer // current connection belong to which server
	// current connection socket
	Conn *net.TCPConn
	// connection ID(Session ID) is global unique ID
	ConnId uint32
	// current connection close status
	isClosed bool

	MsgHandler ziface.IMsgHandler

	// connection exit channel
	ExitBuffChan chan bool

	msgChan     chan []byte
	msgBuffChan chan []byte

	property     map[string]interface{}
	propertyLock sync.RWMutex
}

func NewConnection(server ziface.IServer, conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandler) *Connection {
	c := &Connection{
		TcpServer:    server,
		Conn:         conn,
		ConnId:       connID,
		isClosed:     false,
		MsgHandler:   msgHandler,
		ExitBuffChan: make(chan bool, 1),
		msgChan:      make(chan []byte),
		msgBuffChan:  make(chan []byte, utils.GlobalObject.MaxMsgChanLen),
		property:     make(map[string]interface{}),
	}
	c.TcpServer.GetConnMgr().Add(c)
	return c
}

// Goroutine that handles conn reading data
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running")
	defer fmt.Println(c.Conn.RemoteAddr().String(), " conn reader exit!")
	defer c.Stop()

	for {
		dp := NewDataPack()

		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTcpConnection(), headData); err != nil {
			fmt.Println("recv buf err ", err)
			c.ExitBuffChan <- true
			return
		}

		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack err ", err)
			c.ExitBuffChan <- true
			return
		}

		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.GetTcpConnection(), data); err != nil {
				fmt.Println("read msg data error", err)
				c.ExitBuffChan <- true
				return
			}
		}
		msg.SetData(data)

		req := Request{
			conn: c,
			msg:  msg,
		}

		if utils.GlobalObject.WorkerPoolSize > 0 {
			c.MsgHandler.SendMsgToTaskQueue(&req)
		} else {
			go c.MsgHandler.DoMsgHandler(&req)
		}
	}
}

// start connection let current connection begin work
func (c *Connection) Start() {
	go c.StartReader()

	go c.StartWriter()

	c.TcpServer.CallOnConnStart(c)
}

// stop connection
func (c *Connection) Stop() {
	if c.isClosed {
		return
	}
	c.isClosed = true

	c.TcpServer.CallOnConnStop(c)

	c.Conn.Close()
	c.ExitBuffChan <- true

	c.TcpServer.GetConnMgr().Remove(c)

	close(c.ExitBuffChan)
	close(c.msgBuffChan)
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

func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("Connection closed when send msg")
	}

	dp := NewDataPack()
	msg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("Pack error msg id=", msgId)
		return errors.New("Pack error msg")
	}

	c.msgChan <- msg
	return nil
}

func (c *Connection) SendBuffMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("Connection closed when send buff msg")
	}

	dp := NewDataPack()
	msg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("Pack error msg id=", msgId)
		return errors.New("Pack error msg")
	}

	c.msgBuffChan <- msg
	return nil
}

func (c *Connection) StartWriter() {
	fmt.Println("[Writer Goroutine is running]")
	defer fmt.Println(c.RemoteAddr().String(), "[conn Writer exit!]")

	for {
		select {
		case data := <-c.msgChan:
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send Data error: ", err, " Conn writer exit")
				return
			}
		case data, ok := <-c.msgBuffChan:
			if ok {
				if _, err := c.Conn.Write(data); err != nil {
					fmt.Println("Send Buff Data error: ", err, " Conn writer exit")
					return
				}
			} else {
				fmt.Println("msgBuffChan is Closed")
				return
			}
		case <-c.ExitBuffChan:
			return // conn is closed
		}
	}
}

func (c *Connection) SetProperty(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	c.property[key] = value
}

func (c *Connection) GetProperty(key string) (interface{}, error) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()

	if value, ok := c.property[key]; ok {
		return value, nil
	} else {
		return nil, errors.New("no property found")
	}
}

func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	delete(c.property, key)
}
