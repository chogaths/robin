package pbcodec

import (
	"core/netdef"
	"fmt"
	"github.com/golang/protobuf/proto"
)

type pbCodec struct {
}

// 消息转封包
func (self *pbCodec) EncodeMessage(msg interface{}) *netdef.Packet {

	return BuildPacket(msg.(proto.Message))
}

func (self *pbCodec) GetMetaByName(name string) netdef.IPacketMeta {

	m := getProtoMetaByName(name)

	// 这是golang的大坑, 不能直接返回
	if m == nil {
		return nil
	}

	return m
}

func (self *pbCodec) GetMetaByID(id uint32) netdef.IPacketMeta {

	m := getProtoMetaByID(id)

	// 这是golang的大坑, 不能直接返回
	if m == nil {
		return nil
	}

	return m

}

// 封包转字符串, 调试用
func (self *pbCodec) PacketToString(pkt *netdef.Packet) string {

	if pkt == nil {
		return ""
	}

	meta := getProtoMetaByID(pkt.MsgID)
	if meta == nil {
		return fmt.Sprintf("[0x%x] size: %d", pkt.MsgID, pkt.Size)
	}

	c := netdef.NewPacketContext(pkt, nil)

	// 注入消息适配器
	c.Invoke(meta.Adapter)

	// 使用消息适配器翻译消息
	var final string
	c.Invoke(func(msg proto.Message) {

		final = msg.String()

	})

	return fmt.Sprintf("%s size: %d|%s", meta.GetName(), pkt.Size, final)
}

var instance pbCodec

func GetInterface() netdef.IPacketCodec {
	return &instance
}
