package protocal

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type VerificationCodeProtocal struct {
	Len  uint8
	Code string
}

func NewVerificationCodeProtocal() *VerificationCodeProtocal {
	return &VerificationCodeProtocal{}
}

func (vcp *VerificationCodeProtocal) Serialize() ([]byte, error) {
	buff := bytes.NewBuffer([]byte{})
	err := binary.Write(buff, binary.LittleEndian, vcp.Len)
	if err != nil {
		fmt.Println("verification code protocal serialize len error ", err)
		return nil, err
	}
	err = binary.Write(buff, binary.LittleEndian, []byte(vcp.Code))
	if err != nil {
		fmt.Println("verification code protocal serialize code error ", err)
		return nil, err

	}
	return buff.Bytes(), nil
}

func (vcp *VerificationCodeProtocal) Deserialize(data []byte) error {
	buff := bytes.NewReader(data)
	if err := binary.Read(buff, binary.LittleEndian, &vcp.Len); err != nil {
		fmt.Println("verification code protocal deserialize len error ", err)
		return err

	}
	code := make([]byte, int(vcp.Len))
	if err := binary.Read(buff, binary.LittleEndian, code); err != nil {
		fmt.Println("verification code protocal deserialize code error ", err)
		return err
	}
	vcp.Code = string(code)
	return nil
}
