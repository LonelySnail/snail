package siface

type IMsgHandler interface {
	DoMsgHandler(agent  IAgent,msg IMessage) error
	AddRouter(id byte,router IRouter)
	GetRouter(id byte) (IRouter,bool)
}
