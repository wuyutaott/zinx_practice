package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"zinx_practice/zinx/ziface"
)

type DataPack struct {
}

func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(buffer, binary.BigEndian, msg.GetDataLength()); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, msg.GetID()); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (dp *DataPack) UnPack(data []byte) (ziface.IMessage, error) {
	msg := &Message{}
	buffer := bytes.NewBuffer(data)
	if err := binary.Read(buffer, binary.BigEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	if err := binary.Read(buffer, binary.BigEndian, &msg.ID); err != nil {
		return nil, err
	}
	if msg.DataLen > 4096 {
		return nil, errors.New(fmt.Sprintf("DataLen[%d] is to large", msg.DataLen))
	}
	return msg, nil
}

func NewDataPack() ziface.IDataPack {
	return &DataPack{}
}
