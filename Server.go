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
	_, err := request.GetConnection().GetTcpConnection().Write([]byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

func main() {
	s := znet.NewServer("[zinx v0.3]")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
