package siface

type IRouter interface {
	BeforeHandler(request IRequest)
	Handler(request IRequest)
	AfterHandler(request IRequest)
	RegisterMethods()
	GetMethod(name string)(IMethod,bool)
}
