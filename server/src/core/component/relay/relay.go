package relay

import (
	"core/netdef"
	"core/netfw"
	"errors"
	"github.com/golang/protobuf/proto"
	"log"
	"protos/coredef"
)

type IRelaySender interface {

	// 发送一个消息到gameSession, 由其转发到sessiond的客户端
	Send(relaySession *netdef.Session, sessionid int64, msg interface{}) error

	// 广播一个消息
	Broardcast(msg interface{})

	// 发送一个裸封包到relaySession, 由其转发到sessiond的客户端
	SendRaw(relaySession *netdef.Session, sessionid int64, pkt *netdef.Packet) error

	// 广播一个裸封包
	BroardcastRaw(pkt *netdef.Packet)
}

type IRelayReceiver interface {

	// 收到来自sessionid通过relaySession转发过来的pkt消息
	RegisterCallback(func(relaySession *netdef.Session, sessionid int64, pkt *netdef.Packet, broardCast bool))
}

type relayComponent struct {
	peerImplementor netfw.IPeer
	callback        func(*netdef.Session, int64, *netdef.Packet, bool)
}

var (
	errConnectorNotFound error = errors.New("peer is not connector")
	errSessionNotInit    error = errors.New("relay session not inited or need 'peerses' component ")
)

func (self *relayComponent) SendRaw(relaySession *netdef.Session, sessionid int64, pkt *netdef.Packet) error {

	if relaySession == nil {
		return nil
	}

	// 整合调用封包
	relaySession.Send(&coredef.RelayMessageACK{
		UserMsgID: proto.Uint32(pkt.MsgID),
		UserMsg:   pkt.Data,
		SessionID: proto.Int64(sessionid),
	})

	return nil

}

func (self *relayComponent) Send(relaySession *netdef.Session, sessionid int64, msg interface{}) error {

	return self.SendRaw(relaySession, sessionid, self.peerImplementor.GetCodec().EncodeMessage(msg))
}

func (self *relayComponent) Broardcast(msg interface{}) {

	self.BroardcastRaw(self.peerImplementor.GetCodec().EncodeMessage(msg))

}

func (self *relayComponent) BroardcastRaw(pkt *netdef.Packet) {

	self.peerImplementor.(netfw.IAcceptor).Broardcast(&coredef.RelayMessageACK{
		UserMsgID:  proto.Uint32(pkt.MsgID),
		UserMsg:    pkt.Data,
		SessionID:  proto.Int64(0),
		BroardCast: proto.Bool(true),
	})

}

func (self *relayComponent) RegisterCallback(callback func(relaySession *netdef.Session, sessionid int64, pkt *netdef.Packet, broardCast bool)) {
	self.callback = callback
}

func (self *relayComponent) Start(p netfw.IPeer) {

	// 接收方
	p.RegisterMessage("coredef.RelayMessageACK", func(msg *coredef.RelayMessageACK, relaySession *netdef.Session) {

		userpkt := &netdef.Packet{
			MsgID: msg.GetUserMsgID(),
			Size:  uint16(len(msg.GetUserMsg())),
			Data:  msg.GetUserMsg(),
		}

		if self.callback == nil {
			return
		}

		self.callback(relaySession, msg.GetSessionID(), userpkt, msg.GetBroardCast())

	})
}

func getComponentInterface(channelname string) netfw.IComponent {
	p := netfw.FindPeer(channelname)

	if p == nil {
		return nil
	}

	com := p.GetComponent(componentName)

	if com == nil {
		log.Printf("rpc component not found on channel: %s", channelname)
		return nil
	}

	return com
}

// 发送方接口
func GetSender(channelname string) IRelaySender {

	com := getComponentInterface(channelname)
	if com == nil {
		return nil
	}

	return com.(IRelaySender)
}

// 接收方接口
func GetReceiver(channelname string) IRelayReceiver {

	com := getComponentInterface(channelname)
	if com == nil {
		return nil
	}

	return com.(IRelayReceiver)
}

const componentName string = "relay"

func init() {

	netfw.RegisterComponent(componentName, func(p netfw.IPeer) netfw.IComponent {
		return &relayComponent{
			peerImplementor: p,
		}
	})
}
