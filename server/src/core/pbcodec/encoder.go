package pbcodec

import (
	"core/netdef"
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"log"
	"reflect"
)

// 通过反射取得消息名
func getMessageName(pbmsg proto.Message) string {
	pt := reflect.TypeOf(pbmsg)

	e := pt.Elem()

	if e == nil {
		e = pt
	}

	return e.String()
}

// 消息到封包
func BuildPacket(msg proto.Message) *netdef.Packet {

	data, err := proto.Marshal(msg)

	if err != nil {
		log.Fatal(err)
	}

	meta := getProtoMetaByName(getMessageName(msg))

	if meta == nil {
		log.Fatal("msg not registed")
		return nil
	}

	return &netdef.Packet{MsgID: meta.ID, Data: data, Size: uint16(binary.Size(data))}
}
