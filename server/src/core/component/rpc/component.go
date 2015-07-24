package rpc

import (
	"core/netdef"
	"core/netfw"
	"errors"
	"github.com/golang/protobuf/proto"
	"protos/coredef"
	"sync"
	"time"
)

type IRemoteCall interface {
	// 参数, 返回结构体, 结构体中必须携带MsgID字段, 以方便服务端做消息处理
	Call(args interface{}, reply interface{}) error

	// 手动控制进程
	Go(interface{}, interface{}) (*Call, error)

	// RPC超时, 默认3000毫秒
	SetTimeout(ms int32)
}

// rpc每次调用上下文
type Call struct {
	Done   chan bool
	Reply  interface{} // 用户返回结构体
	callid int64
	e      error
}

func (self *Call) done() {
	self.Done <- true
}

func (self *Call) fail(e error) {
	self.e = e
	self.Done <- false
}

// 1个Component绑定在1个Peer(IConnector)上
// 1个IConnector只有1个session
type rpcComponent struct {
	callDataMap      map[int64]*Call
	rpcIDacc         int64
	callDataMapGuard sync.Mutex

	ses    *netdef.Session
	inited bool

	peerImplementor netfw.IPeer

	timeOut time.Duration
}

const rpcComponentTag = "RPCComponent"

var (
	errRequestTimeout error = errors.New("RPC reqest time out")
	errSessionNotInit error = errors.New("RPC session not inited")
)

func (self *rpcComponent) SetTimeout(ms int32) {
	self.timeOut = time.Millisecond * time.Duration(ms)
}

// 远程调用
func (self *rpcComponent) Call(args interface{}, reply interface{}) error {
	c, err := self.Go(args, reply)

	if err != nil {
		return err
	}

	select {
	// 等待异步响应
	case <-c.Done:
		return nil
	case <-time.After(self.timeOut):
		self.removeCallData(c.callid)
		return errRequestTimeout
	}

	return nil
}

func (self *rpcComponent) Go(args interface{}, reply interface{}) (*Call, error) {

	// 本地调用的上下文
	rcd := &Call{Done: make(chan bool), Reply: reply}

	// 生成本次的调用id
	self.callDataMapGuard.Lock()

	rcd.callid = self.rpcIDacc
	self.rpcIDacc = self.rpcIDacc + 1
	self.callDataMap[rcd.callid] = rcd

	self.callDataMapGuard.Unlock()

	if self.ses == nil {

		// 延迟再取session, 避免session没初始化时就GetInterface
		self.ses = self.peerImplementor.(netfw.IConnector).GetSession()

		if self.ses == nil {
			return nil, errSessionNotInit
		}
	}

	// 处理用户封包
	pkt := self.peerImplementor.GetCodec().EncodeMessage(args)

	// 整合调用封包
	self.ses.Send(&coredef.RemoteCallREQ{
		UserMsgID: proto.Uint32(pkt.MsgID),
		UserMsg:   pkt.Data,
		CallID:    proto.Int64(rcd.callid),
	})

	return rcd, nil
}

// 根据调用号, 获取调用信息
func (self *rpcComponent) getCallData(callid int64) *Call {

	self.callDataMapGuard.Lock()

	defer self.callDataMapGuard.Unlock()

	if v, ok := self.callDataMap[callid]; ok {
		return v
	}

	return nil
}

// 移除调用信息
func (self *rpcComponent) removeCallData(callid int64) {
	self.callDataMapGuard.Lock()

	delete(self.callDataMap, callid)

	self.callDataMapGuard.Unlock()
}

func (self *rpcComponent) Ready() bool {
	return self.inited
}
