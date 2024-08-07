package znet

import "StudyZinx/ziface"

type Request struct {
	conn ziface.IConnection // client connected to the server
	msg  ziface.IMessage    // client request data
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgId()
}
