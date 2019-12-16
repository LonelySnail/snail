package siface

type IMethod interface {
	CallFunc(agent  IAgent,msg IArgs)
	Call(agent  IAgent,msg IArgs) //  等待返回值
	CallGo(agent  IAgent,msg IArgs) // 不需要返回值
}

type IArgs interface {
	GetMsgId() byte
	GetMethodName() string
	ToInt64() int64
	ToString() string
	ToMap() (map[string]interface{},error)
}