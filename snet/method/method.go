package method

import (
	"github.com/snail/siface"
	"github.com/snail/logger"
	"reflect"
	"runtime"
)

type Method struct{
	Go   bool
	Func  interface{}
}

type Args struct {
	MsgId    byte
	Name     string
	Data 	 []byte
}

func (m *Method)CallFunc(agent siface.IAgent,args siface.IArgs)  {
	if m.Go{
		m.CallGo(agent,args)
		return
	}
	m.Call(agent,args)
}

func (m *Method)Call(agent siface.IAgent,args siface.IArgs)  {
	defer func() {
		if err := recover();err != nil {
			buf := make([]byte, 1024)
			l := runtime.Stack(buf, false)
			logger.ZapLog.Error( string(buf[:l]))
		}
	}()
	f := reflect.ValueOf(m.Func)
	in := make([]reflect.Value,0)
	in = append(in,reflect.ValueOf(agent))
	in = append(in,reflect.ValueOf(args))

	f.Call(in)
}

func (m *Method)CallGo(agent siface.IAgent,args siface.IArgs)  {
	defer func() {
		if err := recover();err != nil {
			buf := make([]byte, 1024)
			l := runtime.Stack(buf, false)
			logger.ZapLog.Error( string(buf[:l]))
		}
	}()
	f := reflect.ValueOf(m.Func)
	in := make([]reflect.Value,0)

	in = append(in,reflect.ValueOf(agent))
	in = append(in,reflect.ValueOf(args))

	go f.Call(in)
}