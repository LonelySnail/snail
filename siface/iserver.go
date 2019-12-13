package siface

type IServer interface {
	Start()
	Stop()
	Server()
	AddRouter(id byte,router IRouter)
	GetAgentManage() IAgentManage
}