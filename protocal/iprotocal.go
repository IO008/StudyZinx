package protocal

type IProtocal interface {
	Serialize() []byte
	Deserialize(data []byte)
}
