package ziface

type Iserver interface {
	// start service
	Start()

	// stop service
	Stop()

	// start business services
	Serve()
}
