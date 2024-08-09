package znet

import (
	"StudyZinx/utils"
	"StudyZinx/ziface"
	"fmt"
	"net"
)

type Server struct {
	Name       string
	IPVersion  string
	IP         string
	Port       int
	msgHandler ziface.IMsgHandler
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listener at IP: %s, Port: %d, is starting\n", s.IP, s.Port)
	fmt.Printf("[Zinx] version: %s, MaxConn: %d, MaxPacketSize: %d\n", utils.GlobalObject.Version, utils.GlobalObject.MaxConn, utils.GlobalObject.MaxPacktSize)

	// start linster
	go func() {

		s.msgHandler.StartWorkerPool()

		// 1.get a tcp addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("Resolve TCP address error: ", err)
			return
		}

		// 2. listen server addr
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("Listen ", s.IPVersion, " error: ", err)
			return

		}
		fmt.Println("start Zinx server success, ", s.Name, " is listening...")

		// TODO server.go should have a ID by auto generate
		var cid uint32 = 0

		//3. start server network connection
		for {
			//3.1 block wait for client connection
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept error: ", err)
				continue
			}

			//3.2 TODO Server.Start() set max conn if more than max conn, close the conn

			dealConn := NewConnection(conn, cid, s.msgHandler)
			cid++

			// start current connection business
			go dealConn.Start()
		}

	}()
}

func (s *Server) Stop() {
	fmt.Println("[Stop] Server, name ", s.Name)
	// TODO Server.Stop() stop server, clean up the connection, resource, etc.
}

func (s *Server) Serve() {
	s.Start()

	// TODO Server.Serve() start server, and do some other things, like handle signal, etc.

	// block
	select {}
}

func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.msgHandler.AddRouter(msgId, router)
	fmt.Println("Add Router success for msgID: ", msgId)
}

func NewServer() ziface.Iserver {
	utils.GlobalObject.Reload()
	s := &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		msgHandler: NewMsgHandle(),
	}
	return s
}
