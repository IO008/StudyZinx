package mock_client

import (
	"StudyZinx/business"
	"StudyZinx/protocal"
	"StudyZinx/znet"
	"fmt"
	"io"
	"net"
	"strconv"
	"sync"
)

type MockClient struct {
	conn         net.Conn
	ui           *UI
	commandChan  chan uint32
	router       chan string
	apiId        uint32
	apiIdMutex   sync.Mutex
	dp           *znet.DataPack
	sendMsg      bool
	sendMsgMutex sync.Mutex
}

func NewMockClient(conn net.Conn) *MockClient {
	return &MockClient{
		conn:        conn,
		ui:          NewUI(),
		commandChan: make(chan uint32),
		router:      make(chan string),
		apiId:       business.Unknown,
		dp:          znet.NewDataPack(),
		sendMsg:     false,
	}
}

func (mc *MockClient) ShowLaunchUI() {
	mc.ui.ShowLaunchUI()
}

func (mc *MockClient) LoopInput() {
	for {
		result := mc.ui.ReadInput()
		if !mc.getSendMsg() {
			mc.commandChan <- mc.checkApiId(result)
		} else {
			mc.router <- result
		}
	}
}

func (mc *MockClient) isApiId() bool {
	return mc.getApiId() != business.SendMsg
}

func (mc *MockClient) Block() {
	for {
		select {
		case result := <-mc.commandChan:
			mc.setApiId(result)
			switch result {
			case business.Exit:
				fmt.Println("user exit")
				return
			case business.Register:
				mc.ui.ShowRegister()
				mc.setSendMsg(true)
			case business.VerificationCode:
				mc.ui.ShowVerificationCode()
				mc.setSendMsg(true)
			default:
				fmt.Println("unknown command", result)
				mc.setApiId(business.Unknown)
			}

		case result := <-mc.router:
			mc.sendMessage(mc.getApiId(), result)
		}

	}
}

func (mc *MockClient) checkApiId(value string) uint32 {
	apiId, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		fmt.Println("input command number error", err)
		return 0
	}
	return uint32(apiId)
}

func (mc *MockClient) sendMessage(apiId uint32, result string) {
	msg, err := mc.dp.Pack(znet.NewMsgPackage(apiId, []byte(result)))
	if err != nil {
		fmt.Println("Pack error msg id=", apiId)
		return
	}
	fmt.Printf("==> Send Msg: ID= %d len = %d data = % x  \n", apiId, len(msg), msg)
	mc.conn.Write(msg)

}

func (mc *MockClient) LoopReceived() {

	for {
		headData := make([]byte, mc.dp.GetHeadLen())
		_, err := io.ReadFull(mc.conn, headData)
		if err != nil {
			fmt.Println("read head error")
			return
		}

		msgHead, err := mc.dp.Unpack(headData)
		if err != nil {
			fmt.Println("recv unpack err:", err)
			return
		}

		if msgHead.GetDataLen() > 0 {
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetDataLen())

			_, err := io.ReadFull(mc.conn, msg.Data)
			if err != nil {
				fmt.Println("read recv data err:", err)
				return
			}
			fmt.Printf("==> Recv Msg: ID= %d len = %d data = % x \n", msg.Id, msg.DataLen, msg.Data)
			mc.handleRouter(msg)
		}
	}
}

func (mc *MockClient) handleRouter(msg *znet.Message) {
	switch msg.Id {
	case business.Register:
		mc.handleRegister(msg)
		mc.setApiId(business.ApiMode)
	default:
		fmt.Println("unknown msg id=", msg.Id)
	}
}

func (mc *MockClient) handleRegister(msg *znet.Message) {
	prototal := protocal.NewRegisterProtocal()
	err := prototal.Deserialize(msg.GetData())
	if err != nil {
		return
	}

	if prototal.IsExsit {
		// TODO show chat list
	} else {

		mc.commandChan <- business.VerificationCode
	}
}

func (mc *MockClient) setApiId(apiId uint32) {
	mc.apiIdMutex.Lock()
	defer mc.apiIdMutex.Unlock()
	mc.apiId = apiId
}

func (mc *MockClient) getApiId() uint32 {
	mc.apiIdMutex.Lock()
	defer mc.apiIdMutex.Unlock()
	return mc.apiId
}

func (mc *MockClient) setSendMsg(isSendMsg bool) {
	mc.sendMsgMutex.Lock()
	defer mc.sendMsgMutex.Unlock()
	mc.sendMsg = isSendMsg
}

func (mc *MockClient) getSendMsg() bool {
	mc.sendMsgMutex.Lock()
	defer mc.sendMsgMutex.Unlock()
	return mc.sendMsg
}
