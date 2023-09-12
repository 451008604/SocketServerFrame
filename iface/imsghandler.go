package iface

import (
	pb "github.com/451008604/socketServerFrame/proto/bin"
)

// 消息管理抽象层
type IMsgHandler interface {
	// 异步处理消息
	DoMsgHandler(request IRequest)
	// 为消息添加具体的处理逻辑
	AddRouter(msgId pb.MSgID, msg INewMsgStructTemplate, handler IReceiveMsgHandler)
	// 将消息推入TaskQueue，等待Worker处理
	SendMsgToTaskQueue(request IRequest)
}
