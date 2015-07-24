package netfw

import (
	"core/netdef"
	"protos/coredef"
	"time"
)

const (
	PeerEvent_OnConnected = "OnConnected"
	PeerEvent_OnAccepted  = "OnAccepted"
	PeerEvent_OnClosed    = "OnClosed"
)

// 成员首字母大写对于组合子来说, 等效于protected
type Peer struct {
	Address       string // 因为可能connector会修改, 因此这里保存
	NotifyAddress string // 从global拷贝过来的, 在这里保存

	Index int32  // PeerIndex, 需要根据type定义index
	ID    string // 复杂定义, 参见svcid
	Ready bool
	// 配置
	Define *coredef.PeerDefine

	Codec netdef.IPacketCodec

	*EventDispatcher

	*ComponentManager

	OnBeginRecv func()
	OnEndRecv   func()
}

func (self *Peer) SetRecvCallback(beginRecv func(), endRecv func()) {
	self.OnBeginRecv = beginRecv
	self.OnEndRecv = endRecv
}

func (self *Peer) GetCodec() netdef.IPacketCodec {
	return self.Codec
}

func (self *Peer) IsReady() bool {
	return self.Ready
}

func (self *Peer) GetIndex() int32 {
	return self.Index
}

func (self *Peer) GetAddress() string {
	return self.Address
}

func (self *Peer) GetNotifyAddress() string {
	return self.NotifyAddress
}

func (self *Peer) SetAddress(v string) {
	self.Address = v
}

func (self *Peer) GetID() string {
	return self.ID
}

func (self *Peer) GetSocketWriteTimeout() time.Duration {

	return time.Duration(self.Define.GetSocketSendTimeoutMS()) * time.Millisecond
}

func (self *Peer) GetSocketReadTimeout() time.Duration {

	return time.Duration(self.Define.GetSocketRecvTimeoutMS()) * time.Millisecond
}

// 流包到Chan队列
func RecvPacketLoop(ses *netdef.Session) {

	var err error
	var pkt *netdef.Packet
	// 循环读取封包
	for {

		pkt, err = ses.PStream.Read()

		if err != nil {
			break
		}

		// 封包变事件
		ses.PostEvent(&netdef.NetEvent{Method: netdef.EventRecvPacket, Pkt: pkt, Ses: ses, Tag: nil})
	}

	ses.InternalStop(err)

}

// Chan队列消息发送
func (self *Peer) SendPacketLoop(ses *netdef.Session) {
	for {
		ev := <-ses.OutputChan

		c := netdef.NewPacketContext(ev.Pkt, ev.Ses)

		// 发出事件
		self.CallHandlers(int(ev.Method), c)

		// 注意 Send事件在独立线程, 和Recv事件不在一个线程
		switch ev.Method {
		case netdef.EventSendPacket:
			ses.PStream.Write(ev.Pkt)
		case netdef.EventEnd:
			goto exit
		}

	}
exit:
	//log.Println("send loop exit")
}

// 派发消息到回调
func (self *Peer) DispatchLoop(ses *netdef.Session) {

	for {
		ev := <-ses.InputChan

		if self.OnBeginRecv != nil {
			self.OnBeginRecv()
		}

		c := netdef.NewPacketContext(ev.Pkt, ev.Ses)

		// 派发事件
		self.CallHandlers(int(ev.Method), c)

		// 派发消息
		switch ev.Method {
		case netdef.EventRecvPacket, netdef.EventPostPacket:

			self.CallHandlers(int(ev.Pkt.MsgID), c)

		case netdef.EventEnd:
			if self.OnEndRecv != nil {
				self.OnEndRecv()
			}
			goto exit
		}

		if self.OnEndRecv != nil {
			self.OnEndRecv()
		}

	}
exit:
	//log.Println("dispatch loop exit")
}

// 被实现手动设置进来
func (self *Peer) Init(p IPeer) {

	self.ComponentManager = NewComponentManager(self, p)

	// 调试选项开启时, 注入调试句柄
	if SvcConfig.GetDebugRoute().GetEnable() {

		injectDebugHandler(self)
	}
}

func (self *Peer) GetDefine() *coredef.PeerDefine {
	return self.Define
}

func NewPeerData(def *coredef.PeerDefine, codec netdef.IPacketCodec) *Peer {

	return &Peer{
		Define:          def,
		Ready:           false,
		Codec:           codec,
		EventDispatcher: NewEventDispatcher(def.GetCapturePanic(), codec),
	}
}
