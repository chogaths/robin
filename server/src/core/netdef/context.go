package netdef

import (
	"github.com/codegangsta/inject"
	"log"
	"runtime"
)

type InjectHandler struct {
	handler interface{}
	stack   string
}

func NewInjectHandler(handler interface{}) *InjectHandler {
	// 构建1024的栈存储调用信息
	buf := make([]byte, 1024)
	runtime.Stack(buf, false)

	return &InjectHandler{
		handler: handler,
		stack:   string(buf),
	}
}

// 注入器
type IPacketContext interface {
	inject.Injector

	CallHandlers([]*InjectHandler)

	Next()

	CancelRoute()
}

type packetContext struct {
	inject.Injector

	index int
}

// 调用句柄
func (self *packetContext) CallHandlers(hs []*InjectHandler) {

	// 复位
	self.index = 0

	for self.index < len(hs) {

		h := hs[self.index]

		_, err := self.Invoke(h.handler)

		if err != nil {
			log.Printf("%s %s", err, h.stack)
		}

		// Invoke时, 已经cancel了, 所以这里就退出
		if self.index < 0 {
			break
		}

		self.Next()
	}

}

// 跳过下一个
func (self *packetContext) Next() {
	self.index++
}

// 取消后续所有注入, 直接返回
func (self *packetContext) CancelRoute() {
	self.index = -1
}

func NewPacketContext(pkt *Packet, ses *Session) IPacketContext {
	c := &packetContext{
		Injector: inject.New(),
		index:    0,
	}
	c.Map(pkt)
	c.Map(ses)
	c.MapTo(c, (*IPacketContext)(nil))

	return c
}
