package snet

import (
	"github.com/snail/siface"
	"reflect"
)

type method struct{
	Go   bool
	Func  interface{}
}

func (m *method)Call(request siface.IRequest)  {
	f := reflect.ValueOf(m.Func)
	v := reflect.ValueOf(request)
	f.Call([]reflect.Value{v})
}

func (m *method)CallGo(request siface.IRequest)  {
	f := reflect.ValueOf(m.Func)
	v := reflect.ValueOf(request)
	go f.Call([]reflect.Value{v})
}