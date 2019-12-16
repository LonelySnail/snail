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
	Data 	 []byte
}

func (m *Method)CallFunc(request siface.IRequest)  {
	if m.Go{
		m.CallGo(request)
		return
	}
	m.Call(request)
}

func (m *Method)Call(request siface.IRequest)  {
	defer func() {
		if err := recover();err != nil {
			buf := make([]byte, 1024)
			l := runtime.Stack(buf, false)
			logger.ZapLog.Error( string(buf[:l]))
		}
	}()
	f := reflect.ValueOf(m.Func)
	in := make([]reflect.Value,0)
	args := &Args{
		MsgId:request.GetMsgId(),
		Data: request.GetData(),
	}
	in = append(in,reflect.ValueOf(request.GetAgent()))
	in = append(in,reflect.ValueOf(args))

	f.Call(in)
}

func (m *Method)CallGo(request siface.IRequest)  {
	defer func() {
		if err := recover();err != nil {
			buf := make([]byte, 1024)
			l := runtime.Stack(buf, false)
			logger.ZapLog.Error( string(buf[:l]))
		}
	}()
	f := reflect.ValueOf(m.Func)
	in := make([]reflect.Value,0)
	args := &Args{
		MsgId:request.GetMsgId(),
		Data: request.GetData(),
	}
	in = append(in,reflect.ValueOf(request.GetAgent()))
	in = append(in,reflect.ValueOf(args))

	go f.Call(in)
}