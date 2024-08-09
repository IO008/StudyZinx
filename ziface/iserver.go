package ziface

type IServer interface {
	// start service
	Start()

	// stop service
	Stop()

	// start business services
	Serve()

	//router function: current servier registers a router business method, for client connection processing
	AddRouter(msgId uint32, router IRouter)

	GetConnMgr() IConnManager

	SetOnConnStart(func(IConnection))
	SetOnConnStop(func(IConnection))
	CallOnConnStart(IConnection)
	CallOnConnStop(IConnection)
}
