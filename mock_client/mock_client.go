package mock_client

import (
	"StudyZinx/znet"
	"fmt"
	"io"
	"net"
	"strconv"
	"sync"
)

type MockClient struct {
	conn          net.Conn
	ui            *UI
	commandChan   chan string
	roter         chan string
	currentCmmand string
	mu            sync.Mutex
	dp            *znet.DataPack
}

func NewMockClient(conn net.Conn) *MockClient {
	return &MockClient{
		conn:          conn,
		ui:            NewUI(),
		commandChan:   make(chan string),
		roter:         make(chan string),
		currentCmmand: "",
		dp:            znet.NewDataPack(),
	}
}

func (mc *MockClient) ShowLaunchUI() {
	mc.ui.ShowLaunchUI()
}

func (mc *MockClient) LoopInput() {
	for {
		result := mc.ui.ReadInput()
		if mc.getCurrentCommand() == "" {
			mc.commandChan <- result
		} else {
			mc.roter <- result
		}
	}
}

func (mc *MockClient) Block() {
	for {
		select {
		case result := <-mc.commandChan:
			mc.setCurrentCommand(result)
			switch result {
			case exit:
				fmt.Println("user exit")
				return
			case register:
				mc.ui.ShowRegister()
			default:
				fmt.Println("unknown command")
				mc.setCurrentCommand("")
			}

		case result := <-mc.roter:
			mc.handleRouter(result)
		}

	}
}

func (mc *MockClient) handleRouter(result string) {
	fmt.Println("handle router ", mc.getCurrentCommand(), result)
	switch mc.getCurrentCommand() {
	case register:
		fmt.Println("user register", result)
		mc.sendRegister(result)
	}
}

func (mc *MockClient) sendRegister(result string) {
	num, err := strconv.Atoi(register)
	if err != nil {
		fmt.Println("register error", result)
		return

	}
	msg, err := mc.dp.Pack(znet.NewMsgPackage(uint32(num), []byte(result)))
	if err != nil {
		fmt.Println("Pack error msg id=", register)
		return
	}
	mc.conn.Write(msg)

	mc.ReceivedRegister()
}

func (mc *MockClient) ReceivedRegister() {
	headData := make([]byte, mc.dp.GetHeadLen())
	_, err := io.ReadFull(mc.conn, headData)
	if err != nil {
		fmt.Println("read head error")
		return
	}

	msgHead, err := mc.dp.Unpack(headData)
	if err != nil {
		fmt.Println("register unpack err:", err)
		return
	}

	if msgHead.GetDataLen() > 0 {
		msg := msgHead.(*znet.Message)
		msg.Data = make([]byte, msg.GetDataLen())

		_, err := io.ReadFull(mc.conn, msg.Data)
		if err != nil {
			fmt.Println("read register data err:", err)
			return
		}
		fmt.Println("==> Recv Msg: ID=", msg.Id, ", len =", msg.DataLen, ", data =", msg.Data)
	}
}

func (mc *MockClient) setCurrentCommand(command string) {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	mc.currentCmmand = command
}

func (mc *MockClient) getCurrentCommand() string {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	return mc.currentCmmand
}
