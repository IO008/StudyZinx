package znet

import "StudyZinx/ziface"

type BaseRouter struct{}

func (br *BaseRouter) PreHandle(request ziface.IRequest)  {}
func (br *BaseRouter) Handle(request ziface.IRequest)     {}
func (br *BaseRouter) PostHandle(request ziface.IRequest) {}

func (br *BaseRouter) WriteBool(b bool) byte {
	if b {
		return 1
	}
	return 0
}

func (br *BaseRouter) ReadBool(b byte) bool {
	return b == 1
}
