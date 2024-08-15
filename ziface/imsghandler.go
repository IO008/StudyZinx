package ziface

type IMsgHandler interface {
	DoMsgHandler(request IRequest) // Process messages in a non-blocking manner
	AddRouter(msgID uint32, router IRouter)
	StartWorkerPool()
	SendMsgToTaskQueue(request IRequest)
}
