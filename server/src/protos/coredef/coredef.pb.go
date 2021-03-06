// Code generated by protoc-gen-go.
// source: coredef.proto
// DO NOT EDIT!

/*
Package coredef is a generated protocol buffer package.

It is generated from these files:
	coredef.proto

It has these top-level messages:
	ChannelDefine
	PeerDefine
	DBDefine
	DebugRouteDefine
	OperateAccount
	OperateDB
	OperateConfig
	ProfDefine
	ServiceConfig
	LogDefine
	RemoteCallREQ
	RemoteCallACK
	RPCEchoACK
	RelayMessageACK
	GameSvcConfig
	WorldSvcConfig
	OprSvcConfig
*/
package coredef

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type RemoteCallREQ_MsgID int32

const (
	RemoteCallREQ_ID RemoteCallREQ_MsgID = 0
)

var RemoteCallREQ_MsgID_name = map[int32]string{
	0: "ID",
}
var RemoteCallREQ_MsgID_value = map[string]int32{
	"ID": 0,
}

func (x RemoteCallREQ_MsgID) Enum() *RemoteCallREQ_MsgID {
	p := new(RemoteCallREQ_MsgID)
	*p = x
	return p
}
func (x RemoteCallREQ_MsgID) String() string {
	return proto.EnumName(RemoteCallREQ_MsgID_name, int32(x))
}
func (x *RemoteCallREQ_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RemoteCallREQ_MsgID_value, data, "RemoteCallREQ_MsgID")
	if err != nil {
		return err
	}
	*x = RemoteCallREQ_MsgID(value)
	return nil
}

type RemoteCallACK_MsgID int32

const (
	RemoteCallACK_ID RemoteCallACK_MsgID = 0
)

var RemoteCallACK_MsgID_name = map[int32]string{
	0: "ID",
}
var RemoteCallACK_MsgID_value = map[string]int32{
	"ID": 0,
}

func (x RemoteCallACK_MsgID) Enum() *RemoteCallACK_MsgID {
	p := new(RemoteCallACK_MsgID)
	*p = x
	return p
}
func (x RemoteCallACK_MsgID) String() string {
	return proto.EnumName(RemoteCallACK_MsgID_name, int32(x))
}
func (x *RemoteCallACK_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RemoteCallACK_MsgID_value, data, "RemoteCallACK_MsgID")
	if err != nil {
		return err
	}
	*x = RemoteCallACK_MsgID(value)
	return nil
}

type RPCEchoACK_MsgID int32

const (
	RPCEchoACK_ID RPCEchoACK_MsgID = 0
)

var RPCEchoACK_MsgID_name = map[int32]string{
	0: "ID",
}
var RPCEchoACK_MsgID_value = map[string]int32{
	"ID": 0,
}

func (x RPCEchoACK_MsgID) Enum() *RPCEchoACK_MsgID {
	p := new(RPCEchoACK_MsgID)
	*p = x
	return p
}
func (x RPCEchoACK_MsgID) String() string {
	return proto.EnumName(RPCEchoACK_MsgID_name, int32(x))
}
func (x *RPCEchoACK_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RPCEchoACK_MsgID_value, data, "RPCEchoACK_MsgID")
	if err != nil {
		return err
	}
	*x = RPCEchoACK_MsgID(value)
	return nil
}

type RelayMessageACK_MsgID int32

const (
	RelayMessageACK_ID RelayMessageACK_MsgID = 0
)

var RelayMessageACK_MsgID_name = map[int32]string{
	0: "ID",
}
var RelayMessageACK_MsgID_value = map[string]int32{
	"ID": 0,
}

func (x RelayMessageACK_MsgID) Enum() *RelayMessageACK_MsgID {
	p := new(RelayMessageACK_MsgID)
	*p = x
	return p
}
func (x RelayMessageACK_MsgID) String() string {
	return proto.EnumName(RelayMessageACK_MsgID_name, int32(x))
}
func (x *RelayMessageACK_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RelayMessageACK_MsgID_value, data, "RelayMessageACK_MsgID")
	if err != nil {
		return err
	}
	*x = RelayMessageACK_MsgID(value)
	return nil
}

type ChannelDefine struct {
	Name             *string `protobuf:"bytes,1,opt" json:"Name,omitempty"`
	Address          *string `protobuf:"bytes,2,opt" json:"Address,omitempty"`
	NotifyAddress    *string `protobuf:"bytes,3,opt" json:"NotifyAddress,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ChannelDefine) Reset()         { *m = ChannelDefine{} }
func (m *ChannelDefine) String() string { return proto.CompactTextString(m) }
func (*ChannelDefine) ProtoMessage()    {}

func (m *ChannelDefine) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ChannelDefine) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

func (m *ChannelDefine) GetNotifyAddress() string {
	if m != nil && m.NotifyAddress != nil {
		return *m.NotifyAddress
	}
	return ""
}

// Channel实例化后的参数
type PeerDefine struct {
	Name                *string  `protobuf:"bytes,1,opt" json:"Name,omitempty"`
	ManualStart         *bool    `protobuf:"varint,2,opt" json:"ManualStart,omitempty"`
	Type                *string  `protobuf:"bytes,3,opt" json:"Type,omitempty"`
	Implementor         *string  `protobuf:"bytes,4,opt,def=binarynet" json:"Implementor,omitempty"`
	Component           []string `protobuf:"bytes,5,rep" json:"Component,omitempty"`
	CapturePanic        *bool    `protobuf:"varint,6,opt" json:"CapturePanic,omitempty"`
	SocketRecvTimeoutMS *int32   `protobuf:"varint,7,opt" json:"SocketRecvTimeoutMS,omitempty"`
	SocketSendTimeoutMS *int32   `protobuf:"varint,8,opt" json:"SocketSendTimeoutMS,omitempty"`
	Address             *string  `protobuf:"bytes,9,opt" json:"Address,omitempty"`
	NotifyAddress       *string  `protobuf:"bytes,10,opt" json:"NotifyAddress,omitempty"`
	// Acceptor
	SessionHeartBeatMS *int32 `protobuf:"varint,51,opt" json:"SessionHeartBeatMS,omitempty"`
	// Connector
	AutoReconnect *bool  `protobuf:"varint,81,opt" json:"AutoReconnect,omitempty"`
	FailedWaitSec *int32 `protobuf:"varint,82,opt,def=2" json:"FailedWaitSec,omitempty"`
	PeerIndex     *int32 `protobuf:"varint,83,opt,def=1" json:"PeerIndex,omitempty"`
	PeerCount     *int32 `protobuf:"varint,84,opt,def=1" json:"PeerCount,omitempty"`
	// HttpAcceptor
	TemplateDir      *string `protobuf:"bytes,121,opt" json:"TemplateDir,omitempty"`
	StaticFileDir    *string `protobuf:"bytes,122,opt" json:"StaticFileDir,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PeerDefine) Reset()         { *m = PeerDefine{} }
func (m *PeerDefine) String() string { return proto.CompactTextString(m) }
func (*PeerDefine) ProtoMessage()    {}

const Default_PeerDefine_Implementor string = "binarynet"
const Default_PeerDefine_FailedWaitSec int32 = 2
const Default_PeerDefine_PeerIndex int32 = 1
const Default_PeerDefine_PeerCount int32 = 1

func (m *PeerDefine) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *PeerDefine) GetManualStart() bool {
	if m != nil && m.ManualStart != nil {
		return *m.ManualStart
	}
	return false
}

func (m *PeerDefine) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *PeerDefine) GetImplementor() string {
	if m != nil && m.Implementor != nil {
		return *m.Implementor
	}
	return Default_PeerDefine_Implementor
}

func (m *PeerDefine) GetComponent() []string {
	if m != nil {
		return m.Component
	}
	return nil
}

func (m *PeerDefine) GetCapturePanic() bool {
	if m != nil && m.CapturePanic != nil {
		return *m.CapturePanic
	}
	return false
}

func (m *PeerDefine) GetSocketRecvTimeoutMS() int32 {
	if m != nil && m.SocketRecvTimeoutMS != nil {
		return *m.SocketRecvTimeoutMS
	}
	return 0
}

func (m *PeerDefine) GetSocketSendTimeoutMS() int32 {
	if m != nil && m.SocketSendTimeoutMS != nil {
		return *m.SocketSendTimeoutMS
	}
	return 0
}

func (m *PeerDefine) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

func (m *PeerDefine) GetNotifyAddress() string {
	if m != nil && m.NotifyAddress != nil {
		return *m.NotifyAddress
	}
	return ""
}

func (m *PeerDefine) GetSessionHeartBeatMS() int32 {
	if m != nil && m.SessionHeartBeatMS != nil {
		return *m.SessionHeartBeatMS
	}
	return 0
}

func (m *PeerDefine) GetAutoReconnect() bool {
	if m != nil && m.AutoReconnect != nil {
		return *m.AutoReconnect
	}
	return false
}

func (m *PeerDefine) GetFailedWaitSec() int32 {
	if m != nil && m.FailedWaitSec != nil {
		return *m.FailedWaitSec
	}
	return Default_PeerDefine_FailedWaitSec
}

func (m *PeerDefine) GetPeerIndex() int32 {
	if m != nil && m.PeerIndex != nil {
		return *m.PeerIndex
	}
	return Default_PeerDefine_PeerIndex
}

func (m *PeerDefine) GetPeerCount() int32 {
	if m != nil && m.PeerCount != nil {
		return *m.PeerCount
	}
	return Default_PeerDefine_PeerCount
}

func (m *PeerDefine) GetTemplateDir() string {
	if m != nil && m.TemplateDir != nil {
		return *m.TemplateDir
	}
	return ""
}

func (m *PeerDefine) GetStaticFileDir() string {
	if m != nil && m.StaticFileDir != nil {
		return *m.StaticFileDir
	}
	return ""
}

// DB参数
type DBDefine struct {
	DSN              *string `protobuf:"bytes,1,opt" json:"DSN,omitempty"`
	Enable           *bool   `protobuf:"varint,2,opt" json:"Enable,omitempty"`
	GroupID          *int32  `protobuf:"varint,3,opt" json:"GroupID,omitempty"`
	ShowOperate      *bool   `protobuf:"varint,4,opt" json:"ShowOperate,omitempty"`
	ConnCount        *int32  `protobuf:"varint,5,opt,def=1" json:"ConnCount,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DBDefine) Reset()         { *m = DBDefine{} }
func (m *DBDefine) String() string { return proto.CompactTextString(m) }
func (*DBDefine) ProtoMessage()    {}

const Default_DBDefine_ConnCount int32 = 1

func (m *DBDefine) GetDSN() string {
	if m != nil && m.DSN != nil {
		return *m.DSN
	}
	return ""
}

func (m *DBDefine) GetEnable() bool {
	if m != nil && m.Enable != nil {
		return *m.Enable
	}
	return false
}

func (m *DBDefine) GetGroupID() int32 {
	if m != nil && m.GroupID != nil {
		return *m.GroupID
	}
	return 0
}

func (m *DBDefine) GetShowOperate() bool {
	if m != nil && m.ShowOperate != nil {
		return *m.ShowOperate
	}
	return false
}

func (m *DBDefine) GetConnCount() int32 {
	if m != nil && m.ConnCount != nil {
		return *m.ConnCount
	}
	return Default_DBDefine_ConnCount
}

type DebugRouteDefine struct {
	Enable           *bool    `protobuf:"varint,1,opt" json:"Enable,omitempty"`
	BlockChannelName []string `protobuf:"bytes,2,rep" json:"BlockChannelName,omitempty"`
	BlockMsgName     []string `protobuf:"bytes,3,rep" json:"BlockMsgName,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DebugRouteDefine) Reset()         { *m = DebugRouteDefine{} }
func (m *DebugRouteDefine) String() string { return proto.CompactTextString(m) }
func (*DebugRouteDefine) ProtoMessage()    {}

func (m *DebugRouteDefine) GetEnable() bool {
	if m != nil && m.Enable != nil {
		return *m.Enable
	}
	return false
}

func (m *DebugRouteDefine) GetBlockChannelName() []string {
	if m != nil {
		return m.BlockChannelName
	}
	return nil
}

func (m *DebugRouteDefine) GetBlockMsgName() []string {
	if m != nil {
		return m.BlockMsgName
	}
	return nil
}

// GM工具参数
type OperateAccount struct {
	Account          *string `protobuf:"bytes,1,opt" json:"Account,omitempty"`
	Password         *string `protobuf:"bytes,2,opt" json:"Password,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *OperateAccount) Reset()         { *m = OperateAccount{} }
func (m *OperateAccount) String() string { return proto.CompactTextString(m) }
func (*OperateAccount) ProtoMessage()    {}

func (m *OperateAccount) GetAccount() string {
	if m != nil && m.Account != nil {
		return *m.Account
	}
	return ""
}

func (m *OperateAccount) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

type OperateDB struct {
	Name             *string `protobuf:"bytes,1,opt" json:"Name,omitempty"`
	Addr             *string `protobuf:"bytes,2,opt" json:"Addr,omitempty"`
	ShowName         *string `protobuf:"bytes,3,opt" json:"ShowName,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *OperateDB) Reset()         { *m = OperateDB{} }
func (m *OperateDB) String() string { return proto.CompactTextString(m) }
func (*OperateDB) ProtoMessage()    {}

func (m *OperateDB) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *OperateDB) GetAddr() string {
	if m != nil && m.Addr != nil {
		return *m.Addr
	}
	return ""
}

func (m *OperateDB) GetShowName() string {
	if m != nil && m.ShowName != nil {
		return *m.ShowName
	}
	return ""
}

type OperateConfig struct {
	DB               []*OperateDB      `protobuf:"bytes,1,rep" json:"DB,omitempty"`
	Account          []*OperateAccount `protobuf:"bytes,2,rep" json:"Account,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *OperateConfig) Reset()         { *m = OperateConfig{} }
func (m *OperateConfig) String() string { return proto.CompactTextString(m) }
func (*OperateConfig) ProtoMessage()    {}

func (m *OperateConfig) GetDB() []*OperateDB {
	if m != nil {
		return m.DB
	}
	return nil
}

func (m *OperateConfig) GetAccount() []*OperateAccount {
	if m != nil {
		return m.Account
	}
	return nil
}

// 性能跟踪
type ProfDefine struct {
	CPU              *bool  `protobuf:"varint,1,opt" json:"CPU,omitempty"`
	Mem              *bool  `protobuf:"varint,2,opt" json:"Mem,omitempty"`
	Block            *bool  `protobuf:"varint,3,opt" json:"Block,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ProfDefine) Reset()         { *m = ProfDefine{} }
func (m *ProfDefine) String() string { return proto.CompactTextString(m) }
func (*ProfDefine) ProtoMessage()    {}

func (m *ProfDefine) GetCPU() bool {
	if m != nil && m.CPU != nil {
		return *m.CPU
	}
	return false
}

func (m *ProfDefine) GetMem() bool {
	if m != nil && m.Mem != nil {
		return *m.Mem
	}
	return false
}

func (m *ProfDefine) GetBlock() bool {
	if m != nil && m.Block != nil {
		return *m.Block
	}
	return false
}

// 服务配置
type ServiceConfig struct {
	Channel          []*ChannelDefine  `protobuf:"bytes,1,rep" json:"Channel,omitempty"`
	Peer             []*PeerDefine     `protobuf:"bytes,2,rep" json:"Peer,omitempty"`
	DB               *DBDefine         `protobuf:"bytes,3,opt" json:"DB,omitempty"`
	DebugRoute       *DebugRouteDefine `protobuf:"bytes,4,opt" json:"DebugRoute,omitempty"`
	Log              *LogDefine        `protobuf:"bytes,5,opt" json:"Log,omitempty"`
	Prof             *ProfDefine       `protobuf:"bytes,7,opt" json:"Prof,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *ServiceConfig) Reset()         { *m = ServiceConfig{} }
func (m *ServiceConfig) String() string { return proto.CompactTextString(m) }
func (*ServiceConfig) ProtoMessage()    {}

func (m *ServiceConfig) GetChannel() []*ChannelDefine {
	if m != nil {
		return m.Channel
	}
	return nil
}

func (m *ServiceConfig) GetPeer() []*PeerDefine {
	if m != nil {
		return m.Peer
	}
	return nil
}

func (m *ServiceConfig) GetDB() *DBDefine {
	if m != nil {
		return m.DB
	}
	return nil
}

func (m *ServiceConfig) GetDebugRoute() *DebugRouteDefine {
	if m != nil {
		return m.DebugRoute
	}
	return nil
}

func (m *ServiceConfig) GetLog() *LogDefine {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *ServiceConfig) GetProf() *ProfDefine {
	if m != nil {
		return m.Prof
	}
	return nil
}

type LogDefine struct {
	FileName         *string `protobuf:"bytes,1,opt" json:"FileName,omitempty"`
	Enable           *bool   `protobuf:"varint,2,opt" json:"Enable,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LogDefine) Reset()         { *m = LogDefine{} }
func (m *LogDefine) String() string { return proto.CompactTextString(m) }
func (*LogDefine) ProtoMessage()    {}

func (m *LogDefine) GetFileName() string {
	if m != nil && m.FileName != nil {
		return *m.FileName
	}
	return ""
}

func (m *LogDefine) GetEnable() bool {
	if m != nil && m.Enable != nil {
		return *m.Enable
	}
	return false
}

type RemoteCallREQ struct {
	UserMsg          []byte  `protobuf:"bytes,1,req" json:"UserMsg,omitempty"`
	UserMsgID        *uint32 `protobuf:"varint,2,req" json:"UserMsgID,omitempty"`
	CallID           *int64  `protobuf:"varint,3,req" json:"CallID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RemoteCallREQ) Reset()         { *m = RemoteCallREQ{} }
func (m *RemoteCallREQ) String() string { return proto.CompactTextString(m) }
func (*RemoteCallREQ) ProtoMessage()    {}

func (m *RemoteCallREQ) GetUserMsg() []byte {
	if m != nil {
		return m.UserMsg
	}
	return nil
}

func (m *RemoteCallREQ) GetUserMsgID() uint32 {
	if m != nil && m.UserMsgID != nil {
		return *m.UserMsgID
	}
	return 0
}

func (m *RemoteCallREQ) GetCallID() int64 {
	if m != nil && m.CallID != nil {
		return *m.CallID
	}
	return 0
}

type RemoteCallACK struct {
	UserMsg          []byte  `protobuf:"bytes,1,req" json:"UserMsg,omitempty"`
	UserMsgID        *uint32 `protobuf:"varint,2,req" json:"UserMsgID,omitempty"`
	CallID           *int64  `protobuf:"varint,3,req" json:"CallID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RemoteCallACK) Reset()         { *m = RemoteCallACK{} }
func (m *RemoteCallACK) String() string { return proto.CompactTextString(m) }
func (*RemoteCallACK) ProtoMessage()    {}

func (m *RemoteCallACK) GetUserMsg() []byte {
	if m != nil {
		return m.UserMsg
	}
	return nil
}

func (m *RemoteCallACK) GetUserMsgID() uint32 {
	if m != nil && m.UserMsgID != nil {
		return *m.UserMsgID
	}
	return 0
}

func (m *RemoteCallACK) GetCallID() int64 {
	if m != nil && m.CallID != nil {
		return *m.CallID
	}
	return 0
}

// 测试用消息
type RPCEchoACK struct {
	Content          *string `protobuf:"bytes,1,req" json:"Content,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RPCEchoACK) Reset()         { *m = RPCEchoACK{} }
func (m *RPCEchoACK) String() string { return proto.CompactTextString(m) }
func (*RPCEchoACK) ProtoMessage()    {}

func (m *RPCEchoACK) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

// ////////////////////////////////////////
// relay 组件 服务器转发消息
// ////////////////////////////////////////
type RelayMessageACK struct {
	UserMsg          []byte  `protobuf:"bytes,1,req" json:"UserMsg,omitempty"`
	UserMsgID        *uint32 `protobuf:"varint,2,req" json:"UserMsgID,omitempty"`
	SessionID        *int64  `protobuf:"varint,3,req" json:"SessionID,omitempty"`
	BroardCast       *bool   `protobuf:"varint,4,opt" json:"BroardCast,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RelayMessageACK) Reset()         { *m = RelayMessageACK{} }
func (m *RelayMessageACK) String() string { return proto.CompactTextString(m) }
func (*RelayMessageACK) ProtoMessage()    {}

func (m *RelayMessageACK) GetUserMsg() []byte {
	if m != nil {
		return m.UserMsg
	}
	return nil
}

func (m *RelayMessageACK) GetUserMsgID() uint32 {
	if m != nil && m.UserMsgID != nil {
		return *m.UserMsgID
	}
	return 0
}

func (m *RelayMessageACK) GetSessionID() int64 {
	if m != nil && m.SessionID != nil {
		return *m.SessionID
	}
	return 0
}

func (m *RelayMessageACK) GetBroardCast() bool {
	if m != nil && m.BroardCast != nil {
		return *m.BroardCast
	}
	return false
}

// gs配置
type GameSvcConfig struct {
	UserCount                   *int32  `protobuf:"varint,1,opt" json:"UserCount,omitempty"`
	GMOperationAllowedOn        *bool   `protobuf:"varint,2,opt" json:"GMOperationAllowedOn,omitempty"`
	StatusReportIntervalSeconds *int32  `protobuf:"varint,3,opt" json:"StatusReportIntervalSeconds,omitempty"`
	CharNameBlockFile           *string `protobuf:"bytes,4,opt" json:"CharNameBlockFile,omitempty"`
	XXX_unrecognized            []byte  `json:"-"`
}

func (m *GameSvcConfig) Reset()         { *m = GameSvcConfig{} }
func (m *GameSvcConfig) String() string { return proto.CompactTextString(m) }
func (*GameSvcConfig) ProtoMessage()    {}

func (m *GameSvcConfig) GetUserCount() int32 {
	if m != nil && m.UserCount != nil {
		return *m.UserCount
	}
	return 0
}

func (m *GameSvcConfig) GetGMOperationAllowedOn() bool {
	if m != nil && m.GMOperationAllowedOn != nil {
		return *m.GMOperationAllowedOn
	}
	return false
}

func (m *GameSvcConfig) GetStatusReportIntervalSeconds() int32 {
	if m != nil && m.StatusReportIntervalSeconds != nil {
		return *m.StatusReportIntervalSeconds
	}
	return 0
}

func (m *GameSvcConfig) GetCharNameBlockFile() string {
	if m != nil && m.CharNameBlockFile != nil {
		return *m.CharNameBlockFile
	}
	return ""
}

// world配置
type WorldSvcConfig struct {
	CrossLogicDayHour *int32 `protobuf:"varint,1,opt" json:"CrossLogicDayHour,omitempty"`
	XXX_unrecognized  []byte `json:"-"`
}

func (m *WorldSvcConfig) Reset()         { *m = WorldSvcConfig{} }
func (m *WorldSvcConfig) String() string { return proto.CompactTextString(m) }
func (*WorldSvcConfig) ProtoMessage()    {}

func (m *WorldSvcConfig) GetCrossLogicDayHour() int32 {
	if m != nil && m.CrossLogicDayHour != nil {
		return *m.CrossLogicDayHour
	}
	return 0
}

// opr配置
type OprSvcConfig struct {
	GetList          []string `protobuf:"bytes,8,rep" json:"GetList,omitempty"`
	PostList         []string `protobuf:"bytes,9,rep" json:"PostList,omitempty"`
	AuthorityList    []string `protobuf:"bytes,10,rep" json:"AuthorityList,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *OprSvcConfig) Reset()         { *m = OprSvcConfig{} }
func (m *OprSvcConfig) String() string { return proto.CompactTextString(m) }
func (*OprSvcConfig) ProtoMessage()    {}

func (m *OprSvcConfig) GetGetList() []string {
	if m != nil {
		return m.GetList
	}
	return nil
}

func (m *OprSvcConfig) GetPostList() []string {
	if m != nil {
		return m.PostList
	}
	return nil
}

func (m *OprSvcConfig) GetAuthorityList() []string {
	if m != nil {
		return m.AuthorityList
	}
	return nil
}

func init() {
	proto.RegisterEnum("coredef.RemoteCallREQ_MsgID", RemoteCallREQ_MsgID_name, RemoteCallREQ_MsgID_value)
	proto.RegisterEnum("coredef.RemoteCallACK_MsgID", RemoteCallACK_MsgID_name, RemoteCallACK_MsgID_value)
	proto.RegisterEnum("coredef.RPCEchoACK_MsgID", RPCEchoACK_MsgID_name, RPCEchoACK_MsgID_value)
	proto.RegisterEnum("coredef.RelayMessageACK_MsgID", RelayMessageACK_MsgID_name, RelayMessageACK_MsgID_value)
}
