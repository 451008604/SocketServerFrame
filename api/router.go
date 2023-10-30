package api

import (
	"github.com/451008604/socketServerFrame/network"
	pb "github.com/451008604/socketServerFrame/proto/bin"
	"google.golang.org/protobuf/proto"
)

func RegisterRouter() {
	msgHandler := network.GetInstanceMsgHandler()
	msgHandler.SetFilter(MsgFilter)
	msgHandler.AddRouter(pb.MSgID_Heartbeat_Req, func() proto.Message { return new(pb.HeartbeatRequest) }, HeartBeatHandler)
	msgHandler.AddRouter(pb.MSgID_PlayerLogin_Req, func() proto.Message { return new(pb.PlayerLoginRequest) }, LoginHandler)
	msgHandler.AddRouter(pb.MSgID_Broadcast_Req, func() proto.Message { return new(pb.BroadcastRequest) }, BroadcastHandler)
}

func RegisterRouterClient() {
	msgHandler := network.GetInstanceMsgHandler()
	msgHandler.AddRouter(pb.MSgID_Heartbeat_Res, func() proto.Message { return new(pb.HeartbeatResponse) }, nil)
	msgHandler.AddRouter(pb.MSgID_PlayerLogin_Res, func() proto.Message { return new(pb.PlayerLoginResponse) }, nil)
	msgHandler.AddRouter(pb.MSgID_Broadcast_Res, func() proto.Message { return new(pb.BroadcastResponse) }, nil)
}
