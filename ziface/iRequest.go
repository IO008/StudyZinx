package ziface

type IRequest interface {
	GetConnection() IConnection // get current connection information
	GetData() []byte            // get request data
}
