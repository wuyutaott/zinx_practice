package znet

type Message struct {
	DataLen uint32
	ID      uint32
	Data    []byte
}

func (m *Message) GetID() uint32 {
	return m.ID
}

func (m *Message) GetDataLength() uint32 {
	return m.DataLen
}

func (m *Message) GetData() []byte {
	return m.Data
}
