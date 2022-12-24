package ziface

type IDataPack interface {
	Pack(msg IMessage) ([]byte, error)
	UnPack(data []byte) (IMessage, error)
}
