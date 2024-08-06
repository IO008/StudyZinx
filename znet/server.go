package znet

import (
	"StudyZinx/ziface"
	"errors"
	"fmt"
	"net"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listener at IP: %s, Port: %d, is starting\n", s.IP, s.Port)

	// start linster
	go func() {
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

			//3.3 handle conn business logic, handler and conn should be a bind
			dealConn := NewConnection(conn, cid, CallBackToClient)
			cid++

			//3.4 start current connection business
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

func NewServer(name string) ziface.Iserver {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      7777,
	}
	return s
}

// handle api for current client connection
func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[Conn Handle] CallBackToClient")
	// echo business
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err ", err)
		return errors.New("CallBackToClient error")
	}
	return nil
}
