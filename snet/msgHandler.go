package snet

import (
	"fmt"
	"github.com/snail/siface"
	"github.com/snail/snet/method"
)

type MsgHandler struct {
	Apis map[byte]siface.IRouter
}

func NewMsgHandler()  siface.IMsgHandler{
	return &MsgHandler{
		Apis: map[byte]siface.IRouter{},
	}
}

func (h *MsgHandler) DoMsgHandler(agent siface.IAgent,msg siface.IMessage) error  {
		router,ok := h.GetRouter(msg.GetMsgId())
		if !ok {
			return fmt.Errorf("msgId:%d is wrong",msg.GetMsgId())
		}
		m,ok :=router.GetMethod(msg.GetMethodName())
		if !ok {
			return fmt.Errorf("funcName:%s is wrong",msg.GetNameLen())
		}
		args := &method.Args{
			MsgId:msg.GetMsgId(),
			Name:msg.GetMethodName(),
			Data: msg.GetData(),
		}
		m.CallFunc(agent,args)
		return  nil
}

func (h *MsgHandler)AddRouter(id byte,router siface.IRouter)  {
	router.RegisterMethods()
	if _,ok := h.Apis[id]; ok {
		panic("msg Id is already register")
	}
	h.Apis[id] = router
}

func (h *MsgHandler)GetRouter(id byte) (siface.IRouter,bool) {
	 router,ok := h.Apis[id]
	return router,ok
}