package main

import (
	"fmt"
	"net"

	"StudyZinx/mock_client"
)

func StartClient(conn net.Conn) {
}

func main() {
	fmt.Println("Client Test ... start")

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start err, exit! ", err)
		return
	}

	mc := mock_client.NewMockClient(conn)

	go mc.LoopReceived()

	mc.ShowLaunchUI()

	go mc.LoopInput()

	mc.Block()

	fmt.Println("Client end")

}
