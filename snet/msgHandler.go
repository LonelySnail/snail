package snet

import (
	"github.com/snail/siface"
)

type MsgHandler struct {
	Apis map[byte]siface.IRouter
}

func NewMsgHandler()  siface.IMsgHandler{
	return &MsgHandler{
		Apis: map[byte]siface.IRouter{},
	}
}

func (h *MsgHandler) DoMsgHandler(request siface.IRequest)  {
		router,_ := h.GetRouter(request.GetMsgId())
		m,_ :=router.GetMethod("a")
		m.Call(request)

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