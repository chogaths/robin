package netdef

import (
	"net"
	"sync"
)

type Session struct {
	SessionID int64 // 实例id, 流水号

	PStream IPacketStream // 封包流
	Codec   IPacketCodec

	InputChan  chan *NetEvent // 收到事件
	OutputChan chan *NetEvent // 发出事件
	EventChan  chan bool

	tagMap      map[string]interface{}
	tagMapGuard sync.Mutex

	OnInternalStop func(e error)
}

// 获取玩家的IP
func (self *Session) GetRemoteAddr() string {

	addri := self.PStream.GetDataStream().(net.Conn)

	if addri == nil {
		return ""
	}

	return addri.RemoteAddr().String()
}

// 获取本机的IP
func (self *Session) GetLocalAddr() string {

	addri := self.PStream.GetDataStream().(net.Conn)

	if addri == nil {
		return ""
	}

	return addri.LocalAddr().String()
}

// 发送消息
func (self *Session) Send(msg interface{}) {

	self.SendRaw(self.Codec.EncodeMessage(msg))
}

func (self *Session) SendRaw(pkt *Packet) {

	self.OutputChan <- &NetEvent{
		Method: EventSendPacket,
		Pkt:    pkt,
		Ses:    self,
	}
}

// 派发消息到逻辑线程
func (self *Session) Post(msg interface{}) {

	self.InputChan <- &NetEvent{
		Method: EventPostPacket,
		Pkt:    self.Codec.EncodeMessage(msg),
		Ses:    self,
	}
}

// 获取名字对应的Tag
func (self *Session) GetTag(name string) interface{} {
	self.tagMapGuard.Lock()

	defer self.tagMapGuard.Unlock()

	if v, ok := self.tagMap[name]; ok {
		return v
	}

	return nil
}

func (self *Session) HasTag(name string) bool {
	self.tagMapGuard.Lock()

	defer self.tagMapGuard.Unlock()

	if _, ok := self.tagMap[name]; ok {
		return true
	}

	return false
}

// 设置名字对应的Tag, nil时删除
func (self *Session) SetTag(name string, tag interface{}) {
	self.tagMapGuard.Lock()

	defer self.tagMapGuard.Unlock()

	if tag == nil {
		delete(self.tagMap, name)
	} else {
		self.tagMap[name] = tag
	}
}

func (self *Session) InternalStop(e error) {

	if self.OnInternalStop != nil {
		self.OnInternalStop(e)
	}
	// 连接关闭
	self.PostEvent(&NetEvent{Method: EventClosed, Ses: self, Tag: nil, Pkt: &Packet{Data: []byte(e.Error())}})

	self.PStream.Close()

	// 结束其他线程, 同步线程
	endSignal := &NetEvent{Method: EventEnd}

	self.InputChan <- endSignal

	self.OutputChan <- endSignal
}

// 内部逻辑使用
func (self *Session) PostEvent(ev *NetEvent) {
	self.InputChan <- ev
}

const recvQueueSize = 8
const sendQueueSize = 8

func NewSession() *Session {
	return &Session{
		InputChan:  make(chan *NetEvent, recvQueueSize),
		OutputChan: make(chan *NetEvent, sendQueueSize),
		EventChan:  make(chan bool),
		tagMap:     make(map[string]interface{}),
	}
}
