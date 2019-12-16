package packet

import (
	"bytes"
	"encoding/binary"
	"github.com/snail/siface"
	"io"
)

/*
	msgId ： byte   nameLen：uint16  name string   dataLen： uint16  data :[]byte
*/
type Packet struct {

}

func NewPacket()  siface.IPack{
	return &Packet{}
}

func (pack *Packet) GetHeaderLen() uint32 {
	// id uint8 + nameLen unit16
	return 3
}

func (pack *Packet) GetDataLen() uint16 {
	// id uint8 + nameLen unit16
	return 2
}

func (pack *Packet) Pack(msg siface.IMessage) ([]byte,error) {
	buf :=bytes.NewBuffer([]byte{})
	//写msgID
	if err := binary.Write(buf, binary.BigEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}
	//  方法名长度
	if err := binary.Write(buf, binary.BigEndian,msg.GetNameLen() ); err != nil {
		return nil, err
	}
	// 方法名
	if err := binary.Write(buf, binary.BigEndian, []byte(msg.GetMethodName())); err != nil {
		return nil, err
	}
	//写dataLen
	if err := binary.Write(buf, binary.BigEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	//写data数据
	if err := binary.Write(buf, binary.BigEndian, msg.GetData()); err != nil {
		return nil ,err
	}
	return buf.Bytes(), nil
}

func (pack *Packet) UnPack(r io.Reader) (siface.IMessage,error) {
	head := make([]byte,pack.GetHeaderLen())
	if _,err := io.ReadFull(r,head);err != nil {
		return  nil,err
	}

	buf := bytes.NewReader(head)
	msg := new(Message)

	if err := binary.Read(buf,binary.BigEndian,&msg.Id);err != nil {
		return nil,err
	}

	if err := binary.Read(buf,binary.BigEndian,&msg.NameLen);err != nil {
		return nil,err
	}
	name := make([]byte,msg.NameLen)
	if _,err := io.ReadFull(r,name);err != nil {
		return nil,err
	}
	msg.SetMethodName(string(name))

	len := make([]byte,pack.GetDataLen())
	if _,err := io.ReadFull(r,len);err != nil {
		return  nil,err
	}
	buf = bytes.NewReader(len)
	if err := binary.Read(buf,binary.BigEndian,&msg.Len);err != nil {
		return nil,err
	}

	data := make([]byte,msg.GetDataLen())
	if _,err := io.ReadFull(r,data);err != nil {
		return  nil,err
	}
	msg.SetData(data)
	return msg,nil
}
