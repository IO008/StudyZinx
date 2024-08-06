package ziface

type Iserver interface {
	// start service
	Start()

	// stop service
	Stop()

	// start business services
	Serve()

	//router function: current servier registers a router business method, for client connection processing
	AddRouter(router IRouter)
}
