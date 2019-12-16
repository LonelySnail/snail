package siface

type IRouter interface {
	BeforeHandler()
	Handler()
	AfterHandler()
	RegisterMethods()
	GetMethod(name string)(IMethod,bool)
}
