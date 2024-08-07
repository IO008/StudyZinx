package main

import (
	"StudyZinx/ziface"
	"StudyZinx/znet"
	"fmt"
)

type PingRouter struct {
	znet.BaseRouter
}

// Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle")
	fmt.Println("recv from client: msgID=", request.GetMsgID(), ", data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(1, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println("echo write error ", err)
	}
}

func main() {
	s := znet.NewServer("[zinx v0.5]")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
