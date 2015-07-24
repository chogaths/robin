package netfw

import (
	"core/netdef"
	"log"
	"protos/coredef"
)

// 消息派发器
type IEventDispatcher interface {
	// 注册消息
	RegisterMessage(name string, handlers ...interface{})

	// 注册事件, 事件参考core/netdef/event.go
	RegisterEvent(eventid int, handlers ...interface{})

	// 调用消息/事件id
	CallHandlers(id int, c netdef.IPacketContext)
}

// 端
type IPeer interface {

	// 修改Peer运行地址
	SetAddress(v string)

	// 获取Peer侦听/连接地址
	GetAddress() string

	// 获取下发客户端地址
	GetNotifyAddress() string

	// 获取Peer类型
	GetType() string

	// 获取Peer序号
	GetIndex() int32

	// 获取全局唯一的号
	GetID() string

	// 获取频道实例参数
	GetDefine() *coredef.PeerDefine

	// 获取组件, 如果没有组件, 自动安装
	GetComponent(name string) IComponent

	// 是否有某组件
	HasComponent(name string) bool

	// 获取编码接口
	GetCodec() netdef.IPacketCodec

	// 是否可用
	IsReady() bool

	// 接收回调, 提供给消息加锁使用
	SetRecvCallback(beginRecv func(), endRecv func())

	IEventDispatcher
}

type IPeerStarter interface {
	// 启动Peer
	Start()
}

// 连接器
type IConnector interface {
	GetSession() *netdef.Session
}

// 长连接接收器
type IAcceptor interface {

	// 广播
	Broardcast(msg interface{})

	// 广播裸封包
	BroardcastRaw(pkt *netdef.Packet)

	// 获取连接信息
	GetSession(sessionid int64) *netdef.Session

	//遍历
	Iterate(callback func(*netdef.Session) bool)
}

// Peer挂接及创建
type PeerCreator func(peerType string, p *Peer) IPeerStarter

var peerCreatorMap map[string]PeerCreator = make(map[string]PeerCreator)

// 注册Peer类型创建器, 定制Peer
func RegisterPeer(name string, creator PeerCreator) {

	if _, ok := peerCreatorMap[name]; ok {
		panic("duplicate register peer type: %s")
	}

	peerCreatorMap[name] = creator
}

func CreatePeer(name string, peertype string, p *Peer) IPeerStarter {

	if v, ok := peerCreatorMap[name]; ok {
		return v(peertype, p)
	}

	log.Fatalf("can not found Acceptor creator: '%s'", name)

	return nil
}
