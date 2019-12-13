package snet

import "github.com/snail/siface"

type Request struct {
	agent    siface.IAgent
	msg 	siface.IMessage
}

func (r *Request)GetAgent()siface.IAgent  {
	return  r.agent
}

func (r *Request)GetData()[]byte  {
	return r.msg.GetData()
}

func (r *Request)GetMsgId() byte  {
	return r.msg.GetMsgId()
}

func (r *Request)GetMsgLen() uint16  {
	return r.msg.GetDataLen()
}