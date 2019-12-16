package snet

import (
	"github.com/snail/logger"
	"github.com/snail/siface"
	agent2 "github.com/snail/snet/agent"
	"github.com/snail/util"
	"go.uber.org/zap"
	"log"
	"net"
	"os"
	"os/signal"
)

type Server struct {
	Name  		string
	Addr 		string
	MaxConn			int
	MAxPacketSize 	int
	msgHandler    siface.IMsgHandler
	agentManage    siface.IAgentManage
}

func NewServer(name,addr string) siface.IServer{
	conf := util.LoadConfig("")
	s := &Server{
		Name:          conf.Name,
		Addr:          conf.Addr,
		MaxConn:       conf.MaxConn,
		MAxPacketSize: conf.MAxPacketSize,
		msgHandler:    NewMsgHandler(),
		agentManage:   agent2.NewAgentManage(),
	}


	return s
}
func  (s *Server)Server()  {
	s.Start()
	ch := make(chan os.Signal)
	signal.Notify(ch)
	<- ch
	s.Stop()
}

func (s *Server)Start()  {
	logger.ZapLog.Info("server starting",zap.String("addr",s.Addr))
	listener, err := net.Listen("tcp",s.Addr )
	if err != nil {

	}

	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}

		go s.handler(conn)
	}
	logger.ZapLog.Info("server stop",zap.String("addr",s.Addr))
}

func (s *Server)Stop()  {
	
}
//  单个链接的生存周期
func (s *Server)handler(conn net.Conn)  {
	agent := agent2.NewAgent(conn,s.msgHandler)
	s.agentManage.Add(agent)
	agent.Start()
	agent.Stop()
	s.agentManage.Remove(agent.GetAgentId())
}

func (s *Server)AddRouter(id byte,router siface.IRouter)  {
	s.msgHandler.AddRouter(id,router)
}

func (s *Server)GetRouter(id byte) (siface.IRouter,bool) {
	return s.msgHandler.GetRouter(id)
}

func (s *Server)GetAgentManage() siface.IAgentManage {
	return  s.agentManage
}