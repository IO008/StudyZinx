package ziface

type IRouter interface {
	// hook method before processing conn business logic
	PreHandle(request IRequest)
	// processing conn business logic
	Handle(request IRequest)
	// hook method after processing conn business logic
	PostHandle(request IRequest)
}
