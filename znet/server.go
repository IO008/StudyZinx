package znet

import (
	"StudyZinx/ziface"
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

		//3. start server network connection
		for {
			//3.1 block wait for client connection
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept error: ", err)
				continue
			}

			//3.2 TODO Server.Start() set max conn
			//3.3 TODO Server.Start() handle conn business logic, handler and conn should be a bind

			go func() {
				// while read client data
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf error: ", err)
						continue
					}

					// echo
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf error: ", err)
						continue
					}
				}
			}()
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
