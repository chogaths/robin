package binarynet

import (
	"core/netdef"
	"core/netfw"
)

func init() {
	netfw.RegisterPeer("binarynet", func(peerType string, p *netfw.Peer) netfw.IPeerStarter {
		switch peerType {
		case "connector":
			conn := &binaryConnector{
				Peer: p,
			}

			conn.Init(conn)

			return conn
		case "acceptor":

			acc := &binaryAcceptor{
				Peer:   p,
				sesMap: make(map[int64]*netdef.Session),
			}

			acc.Init(acc)

			return acc
		}

		return nil
	})

}
