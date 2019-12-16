package packet

import "github.com/snail/siface"

type Message struct {
	Id 			byte      	//  消息类型
	NameLen		uint16		//  方法长度
	MethodName 	string		//  方法名字
	Len 		uint16		//  数据长度
	Data     	[]byte		//  数据内容
}

func NewMsg(id byte,name string,data []byte) siface.IMessage {
	return &Message{
		Id:id,
		MethodName:name,
		NameLen:uint16(len([]byte(name))),
		Len:uint16(len(data)),
		Data:data,
	}
}

func (msg *Message)GetMethodName()	string{
	return msg.MethodName
}
func (msg *Message)GetNameLen()	uint16{
	return msg.NameLen
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


func (msg *Message)SetMethodName(name string){
	msg.MethodName = name
}
func (msg *Message)SetNameLen(len uint16){
	msg.NameLen = len
}
func (msg *Message)SetDataLen(len uint16)  {
	msg.Len = len
}

func (msg *Message)SetData(data []byte)  {
	msg.Data = data
}