package pack

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/gstones/zinx/zconf"
	"github.com/gstones/zinx/ziface"
	"github.com/gstones/zinx/zpack"
)

var defaultHeaderLen uint32 = 12

type DataPack struct {
}

// NewDataPack initializes a packing and unpacking instance
// (封包拆包实例初始化方法)
func NewDataPack() *DataPack {
	return &DataPack{}
}

// GetHeadLen returns the length of the message header
// (获取包头长度方法)
func (dp *DataPack) GetHeadLen() uint32 {
	//ID uint32(4 bytes) +  DataLen uint32(4 bytes) +  encrypt uint32(4 bytes)
	return defaultHeaderLen
}

// Pack packs the message (compresses the data)
// (封包方法,压缩数据)
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	// Create a buffer to store the bytes
	// (创建一个存放bytes字节的缓冲)
	dataBuff := bytes.NewBuffer([]byte{})

	// Write the data length
	if err := binary.Write(dataBuff, binary.BigEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	// Write the message ID
	if err := binary.Write(dataBuff, binary.BigEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}
	// Write the data
	if err := binary.Write(dataBuff, binary.BigEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

// Unpack unpacks the message (decompresses the data)
// (拆包方法,解压数据)
func (dp *DataPack) Unpack(binaryData []byte) (ziface.IMessage, error) {
	// Create an ioReader for the input binary data
	dataBuff := bytes.NewReader(binaryData)

	// Only unpack the header information to obtain the data length and message ID
	// (只解压head的信息，得到dataLen和msgID)
	msg := &zpack.Message{}
	encryptLen := uint32(0)

	// Read the data length
	if err := binary.Read(dataBuff, binary.BigEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	// Read the encrypt
	if err := binary.Read(dataBuff, binary.BigEndian, &encryptLen); err != nil {
		return nil, err
	}

	// Read the message ID
	if err := binary.Read(dataBuff, binary.BigEndian, &msg.ID); err != nil {
		return nil, err
	}

	// Determine whether the data length exceeds the maximum packet length
	if msg.DataLen > zconf.GlobalObject.MaxPacketSize {
		return nil, fmt.Errorf("too large msg data received id:%d  len:%d", msg.ID, msg.DataLen)
	}

	return msg, nil
}
