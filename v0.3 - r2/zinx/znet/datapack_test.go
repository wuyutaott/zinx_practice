package znet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

func TestDataPack(t *testing.T) {
	fmt.Println("测试封包解包")

	// 数据包1
	msg1 := &Message{
		DataLen: 5,
		ID: 1,
		Data: []byte("hello"),
	}
	dp := &DataPack{}
	data1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("pack msg1 err:", err)
		return
	}

	// 数据包2
	msg2 := &Message{
		DataLen: 3,
		ID: 2,
		Data: []byte("ktm"),
	}
	dp2 := &DataPack{}
	data2, err := dp2.Pack(msg2)
	if err != nil {
		fmt.Println("pack msg2 err:", err)
		return
	}

	data1 = append(data1, data2...)

	fmt.Println("封包成功", len(data1))

	// 测试解包
	reader := bytes.NewReader(data1)
	for {
		msgHeader := make([]byte, 8)
		if err := binary.Read(reader, binary.BigEndian, msgHeader); err != nil {
			fmt.Println("read err:", err)
			return
		}
		dp3 := &DataPack{}
		msg3, err := dp3.UnPack(msgHeader)
		if err != nil {
			fmt.Println("unpack dp3 err:", err)
			return
		}
		msgBody := make([]byte, msg3.GetDataLength())
		if err := binary.Read(reader, binary.BigEndian, msgBody); err != nil {
			fmt.Println("unpack dp3 body err:", err)
		}
		fmt.Printf("msg3 --> ID = %d, DataLen = %d, Data = %s \n", msg3.GetID(), msg3.GetDataLength(), string(msgBody))
	}

}
