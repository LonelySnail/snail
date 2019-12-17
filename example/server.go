
package main

import (
	"fmt"
	"github.com/snail/siface"
	"github.com/snail/snet"
)

//ping test 自定义路由
type PingRouter struct {
	*snet.BaseRouter
}

func (this *PingRouter)RegisterMethods(){
	this.BaseRouter = snet.NewRouter()
	this.Register("a",this.Handlecc)
}

//Ping Handle
func (this *PingRouter) Handlecc(session siface.IAgent,args siface.IArgs) {
	fmt.Println("Call PingRouter Handle",args.ToString())
	err := session.SendMsg(args.GetMsgId(),args.GetMethodName(), []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}



func main() {
	//创建一个server句柄
	s := snet.NewServer("test","192.168.1.225:9999")

	//配置路由
	s.AddRouter(0, &PingRouter{})


	//开启服务
	s.Server()
}

