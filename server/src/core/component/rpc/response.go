package rpc

import (
	"core/netdef"
	"github.com/golang/protobuf/proto"
	"protos/coredef"
)

type IResponse interface {
	// 反馈给请求方的消息
	Feedback(interface{})
}

type rpcResponse struct {
	req   *coredef.RemoteCallREQ
	ses   *netdef.Session
	codec netdef.IPacketCodec
}

func (self *rpcResponse) Feedback(msg interface{}) {

	ses := self.ses

	pkt := self.codec.EncodeMessage(msg)

	ses.Send(&coredef.RemoteCallACK{
		UserMsgID: proto.Uint32(pkt.MsgID),
		UserMsg:   pkt.Data,
		CallID:    proto.Int64(self.req.GetCallID()),
	})
}

func injectResponseHandler(c netdef.IPacketContext, codec netdef.IPacketCodec, ses *netdef.Session, msg *coredef.RemoteCallREQ) {

	c.MapTo(&rpcResponse{
		codec: codec,
		req:   msg,
		ses:   ses,
	}, (*IResponse)(nil))
}
