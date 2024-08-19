package protocal

type IProtocal interface {
	Serialize() ([]byte, error)
	Deserialize(data []byte) error
}
