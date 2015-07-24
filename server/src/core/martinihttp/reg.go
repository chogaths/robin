package martinihttp

import (
	"core/netfw"
)

func init() {
	netfw.RegisterPeer("martinihttp", func(peerType string, p *netfw.Peer) netfw.IPeerStarter {
		switch peerType {
		case "acceptor":

			acc := &martiniAcceptor{
				Peer: p,
			}

			acc.initSettings()

			acc.Init(acc)

			return acc
		}

		return nil
	})
}
