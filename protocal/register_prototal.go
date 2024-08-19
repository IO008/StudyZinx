package protocal

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type RegisterProtocal struct {
	IsExsit bool
	codeLen uint8
	Code    string
}

func NewRegisterProtocal() *RegisterProtocal {
	return &RegisterProtocal{}
}

func (rp *RegisterProtocal) Serialize() ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})
	if err := binary.Write(dataBuff, binary.LittleEndian, rp.IsExsit); err != nil {
		fmt.Println("register protocal serialize isSuccess error ", err)
		return nil, err
	}

	rp.codeLen = uint8(len(rp.Code))
	if err := binary.Write(dataBuff, binary.LittleEndian, rp.codeLen); err != nil {
		fmt.Println("register protocal serialize codeLen error ", err)
		return nil, err
	}
	if err := binary.Write(dataBuff, binary.LittleEndian, []byte(rp.Code)); err != nil {
		fmt.Println("register protocal serialize code error ", err)
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

func (rp *RegisterProtocal) Deserialize(data []byte) error {
	dataBuff := bytes.NewReader(data)
	if err := binary.Read(dataBuff, binary.LittleEndian, &rp.IsExsit); err != nil {
		fmt.Println("register protocal deserialize error ", err)
		return err
	}

	if err := binary.Read(dataBuff, binary.LittleEndian, &rp.codeLen); err != nil {
		fmt.Println("register protocal deserialize codeLen error ", err)
		return err
	}
	code := make([]byte, int(rp.codeLen))
	if err := binary.Read(dataBuff, binary.LittleEndian, code); err != nil {
		fmt.Println("register protocal deserialize code error ", err)
		return err
	}
	rp.Code = string(code)

	fmt.Println("register protocal deserialize isSuccess=", rp.IsExsit, ", code=", rp.Code)
	return nil
}
