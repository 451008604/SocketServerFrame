package logic

import (
	"github.com/451008604/socketServerFrame/common"
	"github.com/451008604/socketServerFrame/dao/sqlmodel"
	"github.com/451008604/socketServerFrame/iface"
	pb "github.com/451008604/socketServerFrame/proto/bin"
	"google.golang.org/protobuf/proto"
)

type Player struct {
	Conn          iface.IConnection
	Data          *pb.PBPlayerData
	Random        *common.Random // 随机数工具
	itemSpaceSize [2]uint32      // 道具空间大小 [宽, 高]
}

// 初始化玩家默认数据结构
func (p *Player) InitializationSaveData() *pb.PBPlayerData {
	// 初始化临时变量
	p.itemSpaceSize = [2]uint32{common.ItemSpaceWidth, common.ItemSpaceHeight}

	// 初始化缓存变量
	p.Data = &pb.PBPlayerData{
		CommonData: &pb.PBCommonData{},
	}

	return p.Data
}

func (p *Player) SetPlayerData(user *sqlmodel.HouseUser) {
	p.Data.CommonData.UserUniID = proto.Int64(user.UniID)
	p.Data.CommonData.NickName = proto.String(user.Nickname)
	p.Data.CommonData.HeadImage = proto.String(user.HeadImage)
	p.Data.CommonData.RegisterTime = proto.Uint32(uint32(user.RegisterTime))
}

// 项目空间大小
func (p *Player) ItemSpaceSize() uint32 {
	return p.itemSpaceSize[0] * p.itemSpaceSize[1]
}