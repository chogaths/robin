package netfw

import (
	"core/netdef"
	"log"
	"protos/coredef"
)

func validForThisPeer(name string, config *coredef.DebugRouteDefine) bool {

	for _, libname := range config.GetBlockChannelName() {

		if libname == name {
			return false
		}
	}

	return true
}

func validForThisMsg(name string, config *coredef.DebugRouteDefine) bool {

	for _, libname := range config.GetBlockMsgName() {

		if libname == name {
			return false
		}
	}

	return true
}

// 调试信息输出
func registerDebugRouteHandler(peer *Peer, method int) {

	peer.RegisterEvent(method, func(ses *netdef.Session, pkt *netdef.Packet) {

		if ses == nil {
			return
		}

		cname := peer.GetDefine().GetName()

		// 检查和这个peer是否匹配
		if !validForThisPeer(cname, SvcConfig.GetDebugRoute()) {
			return
		}

		sid := ses.SessionID

		var msgname string = ""

		if pkt != nil {
			meta := peer.Codec.GetMetaByID(pkt.MsgID)

			if meta != nil {
				msgname = meta.GetName()
			}
		}

		// 消息有效时, 检查是否有屏蔽的消息
		if msgname != "" && !validForThisMsg(msgname, SvcConfig.GetDebugRoute()) {
			return
		}

		msginfo := peer.Codec.PacketToString(pkt)

		switch method {
		case netdef.EventRecvPacket:
			log.Printf("#recvpacket [%s] sid: %d %s", cname, sid, msginfo)
		case netdef.EventPostPacket:
			log.Printf("#postpacket [%s] sid: %d %s", cname, sid, msginfo)
		case netdef.EventConnected:
			log.Printf("#connected [%s] sid: %d", cname, sid)
		case netdef.EventAccepted:
			log.Printf("#accepted [%s] sid: %d", cname, sid)
		case netdef.EventClosed:
			log.Printf("#closed [%s] sid: %d", cname, sid)
		case netdef.EventSendPacket:
			log.Printf("#sendpacket [%s] sid: %d %s", cname, sid, msginfo)
		}

	})
}

func injectDebugHandler(peer *Peer) {
	registerDebugRouteHandler(peer, netdef.EventAccepted)
	registerDebugRouteHandler(peer, netdef.EventRecvPacket)
	registerDebugRouteHandler(peer, netdef.EventClosed)
	registerDebugRouteHandler(peer, netdef.EventConnected)
	registerDebugRouteHandler(peer, netdef.EventSendPacket)
	registerDebugRouteHandler(peer, netdef.EventPostPacket)
}
