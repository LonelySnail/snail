package snet

import (
	"fmt"
	"github.com/snail/siface"
)

type BaseRouter struct {
	methods  map[string]*method
}



func NewRouter() *BaseRouter {
	return &BaseRouter{
		methods: map[string]*method{},
	}
}

func (r *BaseRouter)BeforeHandler(request siface.IRequest)  {

}

func (r *BaseRouter)Handler(request siface.IRequest)  {

}

func (r *BaseRouter)AfterHandler(request siface.IRequest)  {

}

func (r *BaseRouter)Register(name string,function interface{})  {
		if _,ok := r.methods[name];ok {
			panic(fmt.Sprintf("name:%s is  existing",name))
		}
		m := &method{Go:false,Func:function}
		r.methods[name]= m
}

func (r *BaseRouter)RegisterGo(name string,function interface{})  {
	if _,ok := r.methods[name];ok {
		panic(fmt.Sprintf("name:%s is  existing",name))
	}
	m := &method{Go:true,Func:function}
	r.methods[name]= m
}

func (r *BaseRouter)GetMethod(name string)(siface.IMethod,bool)  {
	m,ok := r.methods[name]
	return m,ok
}