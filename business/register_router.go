package business

import (
	"StudyZinx/protocal"
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

	protocal := protocal.NewRegisterProtocal()
	protocal.IsSuccess = true
	if protocal.IsSuccess {
		protocal.Code = "123456"
	}
	bytes, err := protocal.Serialize()
	if err != nil {
		fmt.Println("register protocal serialize error ", err)
		return
	}
	err = request.GetConnection().SendMsg(request.GetMsgID(), bytes)
	if err != nil {
		fmt.Println("write register error ", err)
	}
}
