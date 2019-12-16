package siface

type IMessage interface {
	GetMsgId()		byte
	GetMethodName()		string
	GetNameLen()	uint16
	GetDataLen()  uint16
	GetData()		[]byte


	SetMsgId(id byte)
	SetMethodName(name string)
	SetNameLen(len uint16)
	SetData(data []byte)
	SetDataLen(len uint16)
}


