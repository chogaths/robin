package netfw

import (
	"core/netdef"
	"log"
	"reflect"
)

type EventDispatcher struct {
	// 事件处理器
	eventHandlers map[int][]*netdef.InjectHandler
	capturePanic  bool
	codec         netdef.IPacketCodec
}

func validateHandler(handler interface{}) {
	if reflect.TypeOf(handler).Kind() != reflect.Func {
		panic("register handler must be a callable func")
	}
}

func (self *EventDispatcher) addHandlers(eventid int, handlers ...interface{}) {

	// 事件
	em, ok := self.eventHandlers[eventid]

	// 新建
	if !ok {

		em = make([]*netdef.InjectHandler, 0)

	}

	// 添加用户注入句柄
	for _, h := range handlers {

		validateHandler(h)

		em = append(em, netdef.NewInjectHandler(h))
	}

	self.eventHandlers[eventid] = em

}

func (self *EventDispatcher) RegisterMessage(name string, handlers ...interface{}) {
	meta := self.codec.GetMetaByName(name)

	if meta == nil {
		log.Printf("message not found: %s", name)
		return
	}

	// 添加消息回调
	self.addHandlers(int(meta.GetID()), meta.GetAdapter())

	// 添加用户回调
	self.addHandlers(int(meta.GetID()), handlers...)
}

// 注册事件
func (self *EventDispatcher) RegisterEvent(eventid int, handlers ...interface{}) {
	self.addHandlers(eventid, handlers...)
}

func (self *EventDispatcher) CallHandlers(id int, c netdef.IPacketContext) {

	if self.capturePanic {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("id: %x panic: %v", id, err) // 这里的err其实就是panic传入的内容，55
			}

		}()
	}

	em, ok := self.eventHandlers[id]

	if ok && len(em) > 0 {
		c.CallHandlers(em)
	}
}

func NewEventDispatcher(capturePanic bool, codec netdef.IPacketCodec) *EventDispatcher {
	return &EventDispatcher{
		eventHandlers: make(map[int][]*netdef.InjectHandler),
		codec:         codec,
		capturePanic:  capturePanic,
	}
}
