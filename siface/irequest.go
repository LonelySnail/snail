package siface

type IRequest interface {
	GetAgent()	 IAgent
	GetData()[]byte
	GetMsgLen() uint16
	GetMsgId() byte
}