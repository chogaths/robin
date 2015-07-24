package netdef

// 1个封包处理适配器扩展信息
type IPacketMeta interface {
	// 取名字
	GetName() string

	// 取ID
	GetID() uint32

	// 取适配器( 将封包转为对应编码的适配函数 )
	GetAdapter() interface{}
}

// 封包编码接口
type IPacketCodec interface {

	// 将消息做成封包
	EncodeMessage(interface{}) *Packet

	// 根据消息名, 取得消息定义信息
	GetMetaByName(string) IPacketMeta

	// 根据消息ID, 去的消息定义信息
	GetMetaByID(uint32) IPacketMeta

	// 封包转字符串
	PacketToString(pkt *Packet) string
}
