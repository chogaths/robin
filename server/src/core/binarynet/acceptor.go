package binarynet

import (
	"core/netdef"
	"core/netfw"
	"log"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

type binaryAcceptor struct {
	*netfw.Peer

	sesMap      map[int64]*netdef.Session
	sesIDAcc    int64
	sesMapGuard sync.Mutex
}

func (self *binaryAcceptor) Start() {

	ln, err := net.Listen("tcp", self.Address)

	if err != nil {
		log.Printf("error listen %s", err.Error())
		return
	}

	// Accept线程
	go func() {

		for {
			conn, err := ln.Accept()

			if err != nil {
				continue
			}

			// 启动session线程
			go func() {
				ses := netdef.NewSession()
				ses.Codec = self.Codec
				ses.PStream = NewPacketStream(conn, self.GetSocketReadTimeout(), self.GetSocketWriteTimeout())

				self.addSession(ses)

				ses.OnInternalStop = func(e error) {

					ses.EventChan <- true
					self.removeSession(ses)

				}

				// 启动派发线程(逻辑线程)
				go self.DispatchLoop(ses)

				// 启动发送线程
				go self.SendPacketLoop(ses)

				// 心跳
				go self.SessionEventLoop(ses)

				// 启动接收线程
				go netfw.RecvPacketLoop(ses)

				// 投递接受的消息
				ses.PostEvent(&netdef.NetEvent{Method: netdef.EventAccepted, Ses: ses, Tag: self})
			}()

		}

	}()

	self.Ready = true

}

func (self *binaryAcceptor) GetType() string {
	return "acceptor"
}

// 获得一个连接
func (self *binaryAcceptor) GetSession(sessionID int64) *netdef.Session {
	self.sesMapGuard.Lock()

	defer self.sesMapGuard.Unlock()
	v, ok := self.sesMap[sessionID]
	if ok {
		return v
	}

	return nil
}

// 广播到所有连接
func (self *binaryAcceptor) Broardcast(msg interface{}) {
	self.sesMapGuard.Lock()
	defer self.sesMapGuard.Unlock()

	for _, ses := range self.sesMap {
		ses.Send(msg)
	}

}

// 广播到所有连接
func (self *binaryAcceptor) BroardcastRaw(pkt *netdef.Packet) {
	self.sesMapGuard.Lock()
	defer self.sesMapGuard.Unlock()

	for _, ses := range self.sesMap {
		ses.SendRaw(pkt)
	}
}

func (self *binaryAcceptor) Iterate(callback func(*netdef.Session) bool) {
	self.sesMapGuard.Lock()
	defer self.sesMapGuard.Unlock()

	for _, ses := range self.sesMap {
		if !callback(ses) {
			break
		}
	}

}

func (self *binaryAcceptor) addSession(ses *netdef.Session) {

	ses.SessionID = atomic.AddInt64(&self.sesIDAcc, 1)

	self.sesMapGuard.Lock()
	self.sesMap[self.sesIDAcc] = ses
	self.sesMapGuard.Unlock()

}

func (self *binaryAcceptor) removeSession(ses *netdef.Session) {
	self.sesMapGuard.Lock()
	delete(self.sesMap, ses.SessionID)
	self.sesMapGuard.Unlock()
}

func (self *binaryAcceptor) SessionEventLoop(ses *netdef.Session) {

	cfgms := self.Define.GetSessionHeartBeatMS()
	if cfgms == 0 {

		<-ses.EventChan

	} else {

		ms := time.Duration(cfgms) * time.Millisecond

		for {

			select {
			case <-time.After(ms):
				ses.PostEvent(&netdef.NetEvent{Method: netdef.EventHeartBeat, Pkt: nil, Ses: ses, Tag: nil})
			case <-ses.EventChan:
				goto exit
			}
		}

	}

exit:
	//log.Println("session event loop exit")
}
