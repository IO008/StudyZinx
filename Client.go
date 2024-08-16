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

	mc.ShowLaunchUI()

	go mc.LoopInput()

	mc.Block()

	fmt.Println("Client end")

	/* var count = 0
	for {
		dp := znet.NewDataPack()
		msg, _ := dp.Pack(znet.NewMsgPackage(0, []byte("ZinxV0.5 Client Test Message")))
		msg1, _ := dp.Pack(znet.NewMsgPackage(1, []byte("ZinxV0.6 Client Test Message")))

		if count%2 == 0 {
			msg = msg1
		}
		_, err := conn.Write(msg)
		if err != nil {
			fmt.Println("write conn err ", err)
			return
		}

		headData := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, headData)
		if err != nil {
			fmt.Println("read head error")
			break
		}
		msgHead, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("server unpack err:", err)
			return
		}

		if msgHead.GetDataLen() > 0 {
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetDataLen())

			_, err := io.ReadFull(conn, msg.Data)
			if err != nil {
				fmt.Println("server unpack data err:", err)
				return
			}
			fmt.Println("==> Recv Msg: ID=", msg.Id, ", len =", msg.DataLen, ", data =", string(msg.Data))
		}
		count++
		time.Sleep(1 * time.Second)
	} */
}
