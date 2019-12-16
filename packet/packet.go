package packet

import (
	"bytes"
	"encoding/binary"
	"github.com/snail/siface"
)

type Packet struct {

}

func NewPacket()  siface.IPack{
	return &Packet{}
}

func (pack *Packet) GetHeaderLen() uint32 {
	// id uint8 + dataLen unit16
	return 3
}


func (pack *Packet) Pack(msg siface.IMessage) ([]byte,error) {
	buf :=bytes.NewBuffer([]byte{})
	//写dataLen
	if err := binary.Write(buf, binary.BigEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	//写msgID
	if err := binary.Write(buf, binary.BigEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	//写data数据
	if err := binary.Write(buf, binary.BigEndian, msg.GetData()); err != nil {
		return nil ,err
	}
	return buf.Bytes(), nil
}

func (pack *Packet) UnPack(payLoad []byte) (siface.IMessage,error) {
	buf := bytes.NewReader(payLoad)
	msg := new(Message)
	if err :=	binary.Read(buf,binary.BigEndian,&msg.Len);err != nil {
		return nil,err
	}

	if err := binary.Read(buf,binary.BigEndian,&msg.Id);err != nil {
		return nil,err
	}
	//if err :=	binary.Read(buf,binary.BigEndian,msg.Data);err != nil {
	//	return nil,err
	//}

	return msg,nil
}
