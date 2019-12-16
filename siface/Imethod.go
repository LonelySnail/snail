package siface

type IMethod interface {
	CallFunc(request IRequest)
	Call(request IRequest) //  等待返回值
	CallGo(request IRequest) // 不需要返回值
}

type Args interface {
	GetMsgId() byte
	ToInt64() int64
	ToString() string
	ToMap() (map[string]interface{},error)
}