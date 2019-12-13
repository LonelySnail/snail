package siface

type IMsgHandler interface {
	DoMsgHandler(request IRequest)
	AddRouter(id byte,router IRouter)
	GetRouter(id byte) (IRouter,bool)
}
