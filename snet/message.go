package snet

import "github.com/snail/siface"

type Message struct {
	Id 		byte
	Len 	uint16
	Data     []byte
}

func NewMsg(id byte,data []byte) siface.IMessage {
	return &Message{
		Id:id,
		Len:uint16(len(data)),
		Data:data,
	}
}

func (msg *Message)GetDataLen() uint16 {
	return  msg.Len
}

func (msg *Message)GetData() []byte  {
	return  msg.Data
}

func (msg *Message)GetMsgId() byte  {
	return  msg.Id
}

func (msg *Message)SetMsgId(id byte)  {
	msg.Id = id
}

func (msg *Message)SetDataLen(len uint16)  {
	msg.Len = len
}

func (msg *Message)SetData(data []byte)  {
	msg.Data = data
}