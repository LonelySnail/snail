package siface

type IMethod interface {
	Call(request IRequest) //  等待返回值
	CallGo(request IRequest) // 不需要返回值
}
