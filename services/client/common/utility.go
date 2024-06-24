package common

import (
	"fmt"
	"io"
	"net"

	"github.com/gstones/zinx/ziface"
	"github.com/gstones/zinx/zpack"
	"google.golang.org/protobuf/proto"

	cpb "github.com/moke-game/game/api/gen/common"
	"github.com/moke-game/game/services/common/pack"
)

func SendMsg(conn net.Conn, msgId cpb.C2S_EVENT, data proto.Message) error {
	if conn == nil {
		return fmt.Errorf("conn is nil")
	}
	dp := pack.NewCliDataPack()
	msg, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("marshal msg err:", err)
		return err
	}
	pkg := zpack.NewMsgPackage(uint32(msgId), msg)
	d, err := dp.Pack(pkg)
	if err != nil {
		return err
	}
	_, err = conn.Write(d)
	if err != nil {
		return err
	}
	return nil
}

func RecvMsg(conn io.Reader) (ziface.IMessage, error) {
	dp := pack.NewCliDataPack()
	//先读出流中的head部分
	headData := make([]byte, dp.GetHeadLen())
	_, err := conn.Read(headData)
	if err != nil {
		return nil, err
	}
	//将headData字节流 拆包到msg中
	msg, err := dp.Unpack(headData)
	if err != nil {
		return nil, err
	}
	if msg.GetDataLen() > 0 {
		//msg是有data数据的，需要再次读取data数据
		data := make([]byte, msg.GetDataLen())
		//根据dataLen从io中读取字节流
		_, err := io.ReadFull(conn, data)
		if err != nil {
			return nil, err
		}
		msg.SetData(data)
	}
	return msg, nil
}
