package rpc

import (
	"core/netdef"
	"core/netfw"
	"errors"
	"github.com/golang/protobuf/proto"
	"log"
	"protos/coredef"
)

var (
	lostCallDataError      error = errors.New("RPC lost call data")
	ackUnmarshalingError   error = errors.New("[RemoteCallACK] unmarshaling error")
	replyUnmarshalingError error = errors.New("Reply data unmarshaling error")
)

func (self *rpcComponent) Start(p netfw.IPeer) {

	// 应答方
	p.RegisterMessage("coredef.RemoteCallREQ", func(msg *coredef.RemoteCallREQ, ses *netdef.Session) {

		userpkt := &netdef.Packet{
			MsgID: msg.GetUserMsgID(),
			Size:  uint16(len(msg.GetUserMsg())),
			Data:  msg.GetUserMsg(),
		}

		userc := netdef.NewPacketContext(userpkt, ses)

		injectResponseHandler(userc, p.GetCodec(), ses, msg)

		p.CallHandlers(int(userpkt.MsgID), userc)

	})

	p.RegisterMessage("coredef.RemoteCallACK", func(msg *coredef.RemoteCallACK) {

		// 获得callid对应的调用上下文
		rcd := self.getCallData(msg.GetCallID())

		defer self.removeCallData(msg.GetCallID())

		if rcd == nil {
			rcd.fail(lostCallDataError)
			return
		}

		err := proto.Unmarshal(msg.GetUserMsg(), rcd.Reply.(proto.Message))
		if err != nil {
			rcd.fail(replyUnmarshalingError)
			return
		}

		// 回应请求方, 请求方解除阻塞
		rcd.done()

	})
}

// 从sess获取一个可用的绑定信息
func GetInterface(channelname string) IRemoteCall {

	p := netfw.FindPeer(channelname)

	if p == nil {
		log.Printf("channel not found: %s", channelname)
		return nil
	}

	com := p.GetComponent(componentName)

	if com == nil {
		log.Printf("rpc component not found on channel: %s", channelname)
		return nil
	}

	return com.(IRemoteCall)
}

const componentName string = "rpc"

func init() {
	// 取得消息对应的消息ID

	netfw.RegisterComponent(componentName, func(p netfw.IPeer) netfw.IComponent {
		c := &rpcComponent{
			callDataMap:     make(map[int64]*Call),
			rpcIDacc:        1,
			peerImplementor: p,
		}

		c.SetTimeout(3000)
		return c
	})
}
