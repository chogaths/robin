package binarynet

import (
	"core/netdef"
	"core/netfw"
	"log"
	"net"
	"time"
)

type binaryConnector struct {
	*netfw.Peer

	Session *netdef.Session

	address string
}

func (self *binaryConnector) Start() {

	// 已经连接, 就不重复连接
	if self.Ready {
		return
	}

	if self.Session == nil {
		self.Session = netdef.NewSession()
		self.Session.Codec = self.Codec
	}

	// 防止线程阻塞
	go self.connect()

}

func (self *binaryConnector) GetType() string {
	return "connector"
}

func (self *binaryConnector) GetSession() *netdef.Session {
	return self.Session
}

func (self *binaryConnector) connect() {

	var errReported bool

	// 建立连接循环
	for {

		err := self.createSession()
		if err != nil {

			// 错误只报一次
			if !errReported {
				log.Printf(err.Error())
				errReported = true
			}

		} else {

			// 连上了
			break
		}

		// 没有重连设置, 退出
		if !self.Define.GetAutoReconnect() {
			break
		}

		// 等待一段时间后再试
		time.Sleep((time.Duration)(self.Define.GetFailedWaitSec()) * time.Second)

	}

	if self.Ready {
		ses := self.Session

		ses.OnInternalStop = func(e error) {
			self.Ready = false
			log.Printf("[%s] lost connection", self.ID)

			// 断开连接时, 如果有自动重连设置, 连接
			if self.Define.GetAutoReconnect() {

				// 等待1秒后, 开始重连
				time.AfterFunc(time.Second, func() {
					go self.connect()
				})

			}
		}

		// 启动派发线程(逻辑线程)
		go self.DispatchLoop(ses)

		// 启动发送线程
		go self.SendPacketLoop(ses)

		// 投递连接事件
		ses.PostEvent(&netdef.NetEvent{Method: netdef.EventConnected, Ses: ses, Tag: self})

		// 启动接收线程
		go netfw.RecvPacketLoop(ses)

	}
}

func (self *binaryConnector) createSession() error {

	var err error
	var conn net.Conn

	conn, err = net.Dial("tcp", self.Address)
	if err != nil {
		return err
	}

	log.Printf("%s connected to %s\n", self.ID, self.Address)

	// 创建会话
	self.Session.PStream = NewPacketStream(conn, self.GetSocketReadTimeout(), self.GetSocketWriteTimeout())

	self.Ready = true

	return nil
}
