package netdef

import (
	"time"
)

type Packet struct {
	MsgID uint32 // 消息ID
	Size  uint16 // Data的大小
	Data  []byte
}

// 封包流
type IPacketStream interface {
	Read() (*Packet, error)
	Write(pkt *Packet) error
	Close() error
	GetDataStream() IDataStream
}

// 数据流, 可以抽象封包或者文件等
type IDataStream interface {
	Read(b []byte) (n int, err error)
	Write(b []byte) (n int, err error)
	Close() error

	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
}
