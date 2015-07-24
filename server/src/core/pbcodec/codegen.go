package pbcodec

import (
	"fmt"
	//	"log"
)

// PB消息的扩展信息
type pbCodecMeta struct {
	Name    string      // 消息名称
	ID      uint32      // 消息ID
	Adapter interface{} // 消息适配器
}

var protoName2MetaMap = map[string]*pbCodecMeta{}
var protoID2MetaMap = map[uint32]*pbCodecMeta{}

func (self *pbCodecMeta) GetID() uint32 {
	return self.ID
}

func (self *pbCodecMeta) GetName() string {
	return self.Name
}

func (self *pbCodecMeta) GetAdapter() interface{} {
	return self.Adapter
}

// Code generation uti
func RegisterProto(name string, id uint32, adpater interface{}) {

	meta := &pbCodecMeta{
		ID:      id,
		Name:    name,
		Adapter: adpater,
	}

	// 名称重复检查
	if libmeta := getProtoMetaByName(name); libmeta != nil && libmeta.ID == meta.ID {
		panic(fmt.Sprintf("Name duplicate proto name: %s id: 0x%x", name, meta.ID))
	}

	// ID重复检查
	if libmeta := getProtoMetaByID(meta.ID); libmeta != nil {
		panic(fmt.Sprintf("ID duplicate proto name: %s id: 0x%x", name, meta.ID))
	}

	protoName2MetaMap[name] = meta
	protoID2MetaMap[meta.ID] = meta

}

func getProtoMetaByName(name string) *pbCodecMeta {
	if meta, ok := protoName2MetaMap[name]; ok {

		return meta
	}

	return nil
}

func getProtoMetaByID(id uint32) *pbCodecMeta {
	if meta, ok := protoID2MetaMap[id]; ok {
		return meta
	}

	return nil
}
