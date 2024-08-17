package protocal

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type RegisterProtocal struct {
	IsSuccess bool
	Code      string
}

func NewRegisterProtocal() *RegisterProtocal {
	return &RegisterProtocal{}
}

func (rp *RegisterProtocal) Serialize() []byte {
	dataBuff := bytes.NewBuffer([]byte{})
	if err := binary.Write(dataBuff, binary.LittleEndian, rp.IsSuccess); err != nil {
		fmt.Println("register protocal serialize isSuccess error ", err)
	}

	if rp.IsSuccess {
		if err := binary.Write(dataBuff, binary.LittleEndian, []byte(rp.Code)); err != nil {
			fmt.Println("register protocal serialize code error ", err)
		}
	}

	return dataBuff.Bytes()
}

func (rp *RegisterProtocal) Deserialize(data []byte) {
	dataBuff := bytes.NewReader(data)
	if err := binary.Read(dataBuff, binary.LittleEndian, &rp.IsSuccess); err != nil {
		fmt.Println("register protocal deserialize error ", err)
	}

	if rp.IsSuccess {
		code := make([]byte, len(data)-1)
		if err := binary.Read(dataBuff, binary.LittleEndian, code); err != nil {
			fmt.Println("register protocal deserialize code error ", err)
		}
		rp.Code = string(code)
	}

	fmt.Println("register protocal deserialize isSuccess=", rp.IsSuccess, ", code=", rp.Code)
}
