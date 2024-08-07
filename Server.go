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
	fmt.Println("Call Ping Handle")
	fmt.Println("recv from client: msgID=", request.GetMsgID(), ", data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(0, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println("echo write error ", err)
	}
}

type HelloZinxRouter struct {
	znet.BaseRouter
}

func (this *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HelloZinxRouter Handle")
	fmt.Println("recv from client: msgID=", request.GetMsgID(), ", data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(1, []byte("Hello Zinx Router V0.6"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	s := znet.NewServer()

	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})

	s.Serve()
}
