package znet

import "StudyZinx/ziface"

type Request struct {
	conn ziface.IConnection // client connected to the server
	data []byte             // client request data
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
