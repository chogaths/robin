package binarynet

import (
	"bytes"
	"core/netdef"
	"encoding/binary"
	"errors"
	"io"
	"sync"
	"time"
)

const (
	PackageHeaderSize = 8
	MaxPacketSize     = 1024 * 8
)

type pktStream struct {
	recvtag      uint16
	sendtag      uint16
	ds           netdef.IDataStream
	sendtagGuard sync.RWMutex

	readTimeout  time.Duration
	writeTimeout time.Duration
}

var (
	packageTagNotMatch     = errors.New("ReadPacket: package tag not match")
	packageDataSizeInvalid = errors.New("ReadPacket: package crack, invalid size")
	packageTooBig          = errors.New("ReadPacket: package too big")
)

func (self *pktStream) GetDataStream() netdef.IDataStream {
	return self.ds
}

// 参考hub_client.go
// Read a packet from a datastream interface , return packet struct
func (self *pktStream) Read() (p *netdef.Packet, err error) {

	headdata := make([]byte, PackageHeaderSize)

	//	if self.readTimeout != 0 {
	//		self.ds.SetReadDeadline(time.Now().Add(self.readTimeout))
	//	}

	if _, err = io.ReadFull(self.ds, headdata); err != nil {
		return nil, err
	}

	p = &netdef.Packet{}

	// 读取包头
	headbuf := bytes.NewReader(headdata)
	if err = binary.Read(headbuf, binary.LittleEndian, &p.MsgID); err != nil {
		return nil, err
	}

	// 读取tag
	var tag uint16
	if err = binary.Read(headbuf, binary.LittleEndian, &tag); err != nil {
		return nil, err
	}

	// 读取整包大小
	var fullsize uint16
	if err = binary.Read(headbuf, binary.LittleEndian, &fullsize); err != nil {
		return nil, err
	}

	// 封包太大
	if fullsize > MaxPacketSize {
		return nil, packageTooBig
	}

	// tag不匹配
	if self.recvtag != tag {
		return nil, packageTagNotMatch
	}

	p.Size = fullsize - PackageHeaderSize
	if p.Size < 0 {
		return nil, packageDataSizeInvalid
	}

	// 读取数据
	p.Data = make([]byte, p.Size)
	if _, err = io.ReadFull(self.ds, p.Data); err != nil {
		return nil, err
	}

	// 增加序列号值
	self.recvtag++

	return
}

// Write a packet to datastream interface
func (self *pktStream) Write(pkt *netdef.Packet) (err error) {

	outbuff := bytes.NewBuffer([]byte{})

	// 防止将Send放在go内造成的多线程冲突问题
	self.sendtagGuard.Lock()
	defer self.sendtagGuard.Unlock()

	//	if self.writeTimeout != 0 {
	//		self.ds.SetWriteDeadline(time.Now().Add(self.writeTimeout))
	//	}

	// 发消息ID
	if err = binary.Write(outbuff, binary.LittleEndian, pkt.MsgID); err != nil {
		return
	}

	// 发序列号
	if err = binary.Write(outbuff, binary.LittleEndian, self.sendtag); err != nil {
		return
	}

	// 发包大小
	if err = binary.Write(outbuff, binary.LittleEndian, uint16(pkt.Size+PackageHeaderSize)); err != nil {
		return
	}

	// 发包头
	if _, err = self.ds.Write(outbuff.Bytes()); err != nil {
		return
	}

	// 发包内容
	if _, err = self.ds.Write(pkt.Data); err != nil {
		return
	}

	// 增加序列号值

	self.sendtag++

	return
}

func (self *pktStream) Close() error {
	return self.ds.Close()
}

func NewPacketStream(inds netdef.IDataStream,
	readTimeout time.Duration,
	writeTimeout time.Duration) netdef.IPacketStream {
	return &pktStream{
		ds:           inds,
		recvtag:      1,
		sendtag:      1,
		readTimeout:  readTimeout,
		writeTimeout: writeTimeout,
	}
}
