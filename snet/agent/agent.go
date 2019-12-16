package agent

import (
	"bufio"
	"errors"
	"github.com/snail/logger"
	"github.com/snail/packet"
	"github.com/snail/siface"
	"go.uber.org/zap"
	"net"
	"runtime"
	"sync"
)

type Agent struct {
	conn   net.Conn
	id 		uint32
	isClose  bool
	remoteAddr  net.Addr
	router  	siface.IRouter
	r           *bufio.Reader
	w           *bufio.Writer
	handler     siface.IMsgHandler
	property     sync.Map
	msgChan      chan []byte
	writeNum 	int
	closeChan  chan  bool
}

func NewAgent(conn net.Conn,handler siface.IMsgHandler) siface.IAgent  {
	agent := new(Agent)
	agent.conn   = conn
	agent.remoteAddr = conn.RemoteAddr()
	agent.r  = bufio.NewReaderSize(conn,128)
	agent.w  = bufio.NewWriterSize(conn,128)
	agent.handler = handler
	agent.msgChan = make(chan []byte,10)
	agent.closeChan = make(chan bool,1)
	return  agent
}

func (agent *Agent) Start()  {
	agent.ListenAndLoop()
	<- agent.closeChan
}

func (agent *Agent) Stop()  {
	if agent.isClose {
		return
	}
	agent.isClose = true
	agent.conn.Close()
	close(agent.msgChan)
	close(agent.closeChan)

}

func (agent *Agent) GetConn() net.Conn {
	return agent.conn
}

func (agent *Agent) GetAgentId()  uint32{
	return agent.id
}

func (agent *Agent) RemoteAddr() net.Addr  {
	return agent.remoteAddr
}

func (agent *Agent)ListenAndLoop()  {
	defer func() {
		if err := recover();err != nil {
			buf := make([]byte, 1024)
			l := runtime.Stack(buf, false)
			logger.ZapLog.Error( string(buf[:l]))
		}
	}()
	go agent.writeLoop()
	go agent.readLoop()

	return
}

func (agent *Agent)readLoop ()  {
	defer func() {
		if err := recover();err != nil {
			buf := make([]byte, 1024)
			l := runtime.Stack(buf, false)
			logger.ZapLog.Error( string(buf[:l]))
		}
	}()
loop:
	for !agent.isClose {
		pack := packet.NewPacket()

		msg,err := pack.UnPack(agent.r)
		if err != nil {
			break loop
		}

		logger.ZapLog.Info("request",zap.String("req",string(msg.GetData())))
		err = agent.handler.DoMsgHandler(agent,msg)
		if err != nil {
			logger.ZapLog.Info(err.Error())
		}
	}
	if !agent.isClose {
		agent.closeChan <- true
	}

}

func (agent *Agent) writeLoop(){
	defer func() {
		if err := recover();err != nil {
			buf := make([]byte, 1024)
			l := runtime.Stack(buf, false)
			logger.ZapLog.Error( string(buf[:l]))
		}
	}()
loop:
	for  !agent.isClose{
		select {
		case msg,ok := <- agent.msgChan:
			if !ok {
				break loop
			}
			err := agent.WriteMsg(msg)
			if err != nil {
				logger.ZapLog.Info(err.Error())
			}
		}
	}
	if !agent.isClose {
		agent.closeChan <- true
	}

	return
}

func (agent *Agent)SendMsg(id byte,name string,data []byte)error  {
	defer func() {
		if err := recover();err != nil {
			buf := make([]byte, 1024)
			l := runtime.Stack(buf, false)
			logger.ZapLog.Error( string(buf[:l]))
		}
	}()
	if agent.isClose {
		return errors.New("connect is closed")
	}
	pack := packet.NewPacket()
	msg,err := pack.Pack(packet.NewMsg(id,name,data))
	if err != nil {
		return  err
	}
	agent.msgChan <- msg
	return  nil
}

func (agent *Agent)WriteMsg(msg []byte)  error {
	agent.writeNum ++
	_,err := agent.w.Write(msg)
	if err != nil {
		logger.ZapLog.Info(err.Error())
		return err
	}
	return agent.w.Flush()
}

func (agent *Agent)SetProperty(key string,value interface{})  {
	agent.property.Store(key,value)
}

func (agent *Agent)GetProperty(key string)(interface{},bool)  {
	value,ok := agent.property.Load(key)
	return  value,ok
}

func (agent *Agent)DelProperty(key string)  {
	agent.property.Delete(key)
}