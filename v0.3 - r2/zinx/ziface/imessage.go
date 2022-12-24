package ziface

type IMessage interface {
	GetDataLength() uint32
	GetID() uint32
	GetData() []byte
}