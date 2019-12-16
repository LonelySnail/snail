package snet

import (
	"fmt"
	"github.com/snail/siface"
	"github.com/snail/snet/method"
)

type BaseRouter struct {
	methods  map[string]siface.IMethod
}



func NewRouter() *BaseRouter {
	return &BaseRouter{
		methods: map[string]siface.IMethod{},
	}
}

func (r *BaseRouter)BeforeHandler()  {

}

func (r *BaseRouter)Handler()  {

}

func (r *BaseRouter)AfterHandler()  {

}

func (r *BaseRouter)Register(name string,function interface{})  {
		if _,ok := r.methods[name];ok {
			panic(fmt.Sprintf("name:%s is  existing",name))
		}
		m := &method.Method{Go: false,Func:function}
		r.methods[name]= m
}

func (r *BaseRouter)RegisterGo(name string,function interface{})  {
	if _,ok := r.methods[name];ok {
		panic(fmt.Sprintf("name:%s is  existing",name))
	}
	m := &method.Method{Go: true,Func:function}
	r.methods[name]= m
}

func (r *BaseRouter)GetMethod(name string)(siface.IMethod,bool)  {
	m,ok := r.methods[name]
	return m,ok
}