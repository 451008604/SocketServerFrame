package iface

import (
	pb "github.com/451008604/socketServerFrame/proto/bin"
)

/*
	IRequest 接口：
	实际上是把客户端请求的连接信息 和 请求的数据 包装到了 Request里
*/
type IRequest interface {
	// 获取请求连接信息
	GetConnection() IConnection
	// 获取请求消息的数据
	GetData() []byte
	// 获取请求消息的ID
	GetMsgID() pb.MSgID
}
