package main

import (
	"StudyZinx/business"
	"StudyZinx/ziface"
	"StudyZinx/znet"
	"fmt"
)

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

func DoConnectionBegin(conn ziface.IConnection) {
	fmt.Println("DoConnectionBegin is Called...")
	conn.SetProperty("Name", "studyZinx")
	conn.SetProperty("Home", "https://github.com/IO008/StudyZinx")
	if err := conn.SendMsg(2, []byte("DoConnection BEGIN")); err != nil {
		fmt.Println(err)
	}
}

func DoConnectionLost(conn ziface.IConnection) {
	if name, err := conn.GetProperty("Name"); err == nil {
		fmt.Println("Conn Property Name = ", name)
	}

	if name, err := conn.GetProperty("Home"); err == nil {
		fmt.Println("Conn Property Home = ", name)
	}
	fmt.Printf("DoConnectionLost(%d) is Called... \n", conn.GetConnID())
}

func main() {
	s := znet.NewServer()

	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)

	s.AddRouter(0, &business.RegisterRouter{})
	s.AddRouter(1, &HelloZinxRouter{})

	s.Serve()
}
