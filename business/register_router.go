package business

import (
	"StudyZinx/ziface"
	"StudyZinx/znet"
	"fmt"
)

type RegisterRouter struct {
	znet.BaseRouter
}

func (r *RegisterRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Ping Handle")
	fmt.Println("recv from client: msgID=", request.GetMsgID(), ", data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(request.GetMsgID(), []byte{r.WriteBool(false)})
	if err != nil {
		fmt.Println("write register error ", err)
	}
}
