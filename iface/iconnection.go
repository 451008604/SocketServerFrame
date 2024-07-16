package iface

import (
	"container/list"
	"google.golang.org/protobuf/proto"
)

type IConnection interface {
	// 启动连接(通过connmanager调用)
	Start(readerHandler func() bool, writerHandler func(data []byte) bool)
	// 停止连接(通过connmanager调用)
	Stop()

	// 启动接收消息协程
	StartReader() bool
	// 启动发送消息协程
	StartWriter(data []byte) bool

	// 获取当前连接Id
	GetConnId() int
	// 获取当前工作池Id
	GetWorkId() int
	// 获取接收到的消息队列
	GetMsgQueue() *list.List
	// 获取客户端地址信息
	RemoteAddrStr() string
	// 获取连接是否已关闭
	IsClose() bool

	// 发送消息给客户端
	SendMsg(msgId int32, msgData proto.Message)

	// 设置连接属性
	SetProperty(key IConnProperty, value any)
	// 获取连接属性
	GetProperty(key IConnProperty) (value any)
	// 删除连接属性
	RemoveProperty(key IConnProperty)

	// 限流控制
	FlowControl() bool

	// 序列化
	ProtocolToByte(str proto.Message) []byte
	// 反序列化
	ByteToProtocol(byte []byte, target proto.Message) error
}

// 连接挂载的属性
type IConnProperty string

const (
	SysPropertyConnOpened IConnProperty = "SysPropertyConnOpened" // 连接建立时
	SysPropertyConnClosed IConnProperty = "SysPropertyConnClosed" // 连接关闭时
	SysPropertyLimit      IConnProperty = "SysPropertyLimit"      // 触发限流时
)
