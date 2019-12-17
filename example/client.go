package main

import (
	"fmt"
	"net"
	"time"
	"github.com/snail/packet"
)

/*
   模拟客户端
*/
func main() {

	fmt.Println("Client Test ... start")
	//3秒之后发起测试请求，给服务端开启服务的机会
	time.Sleep(3 * time.Second)

	conn,err := net.Dial("tcp", "192.168.1.225:9999")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	for {
		//发封包message消息
		dp := packet.NewPacket()
		msg, err := dp.Pack(packet.NewMsg(0,"a",[]byte("Client0 Test Message")))
		if err != nil {
			fmt.Println(err)
		}

		_, err = conn.Write(msg)
		if err !=nil {
			fmt.Println("write error err ", err)
			return
		}

		dp = packet.NewPacket()
		pk,err := dp.UnPack(conn)
		if err != nil {
			continue


		}
		fmt.Println("Msg: ID=", pk.GetMsgId(), ", len=", pk.GetDataLen(), ", data=",string(pk.GetData()) )
		time.Sleep(1*time.Second)
	}
}