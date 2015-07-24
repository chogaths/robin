package netdef

const (
	EventRecvPacket = 1 // 接收封包
	EventClosed     = 2 // socket关闭
	EventConnected  = 3 // 连接上
	EventAccepted   = 4 // 获取一个连接
	EventSendPacket = 5 // 发包
	EventPostPacket = 6 // 内部投递
	EventHeartBeat  = 7 // 心跳
	EventEnd        = 8 // 系统停止
)

type NetEvent struct {
	Method int32       // 见上面常量表
	Pkt    *Packet     // 封包
	Ses    *Session    // 回话
	Tag    interface{} // 上下文绑定, 例如: acceptor/connector
}
