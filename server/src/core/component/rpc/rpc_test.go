package rpc

import (
	_ "core/binarynet"
	"core/netdef"
	"core/netfw"
	"github.com/golang/protobuf/proto"
	"log"
	"protos/coredef"
	"testing"
)

func TestRPC(t *testing.T) {

	netfw.MergeConfigString(`
Channel
{
	Name:"rpclocaltest"
	Address: "127.0.0.1:9001"
}

RunAcceptor
{
	Name: "rpclocaltest"
	Component
	{
		Name: "rpc"
	} 
}

RunConnector
{
	Name: "rpclocaltest"
	Component
	{
		Name: "rpc"
	} 
}
	`)

	netfw.Start()

	connectedsignal := make(chan bool)

	// 模拟服务器跑在另外一个地方
	go func() {
		if acc := netfw.GetAcceptor("rpclocaltest"); acc != nil {
			acc.RegisterMessage("coredef.RPCEchoACK", func(msg *coredef.RPCEchoACK, resp IResponse) {

				log.Printf("[server]recv RPCEchoACK content: %s", msg.GetContent())

				resp.Feedback(&coredef.RPCEchoACK{
					Content: proto.String("this is rpc!"),
				})

				log.Println("[server]send RPCEchoACK feedback")

			})

		}

		if conn := netfw.GetConnector("rpclocaltest"); conn != nil {
			conn.RegisterEvent(netdef.EventConnected, func() {

				// 这里表示rpc链路连接ok
				connectedsignal <- true
			})

		}

	}()

	// 等连上后, 执行后面的东西
	<-connectedsignal

	// 获取rpc session
	ctx := GetInterface("rpclocaltest")

	log.Printf("[client]send rpc request")

	ret := &coredef.RPCEchoACK{}

	err := ctx.Call(&coredef.RPCEchoACK{
		Content: proto.String("hello"),
	}, ret)

	if err != nil {
		log.Println(err)
	}

	log.Printf("[client]client recv feed back :%s", ret.GetContent())

	log.Println("[client]do etc logic after rpc")
}
