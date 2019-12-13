package siface

type IMessage interface {
	GetDataLen()  uint16
	GetMsgId()		byte
	GetData()		[]byte

	SetMsgId(id byte)
	SetData(data []byte)
	SetDataLen(len uint16)
}
