// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.22.2
// source: player.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PBPlayerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonData *PBCommonData `protobuf:"bytes,1,opt,name=CommonData" json:"CommonData"` // 公共数据
}

func (x *PBPlayerData) Reset() {
	*x = PBPlayerData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBPlayerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBPlayerData) ProtoMessage() {}

func (x *PBPlayerData) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBPlayerData.ProtoReflect.Descriptor instead.
func (*PBPlayerData) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{0}
}

func (x *PBPlayerData) GetCommonData() *PBCommonData {
	if x != nil {
		return x.CommonData
	}
	return nil
}

type PBCommonData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUniID    *int64  `protobuf:"varint,1,opt,name=UserUniID" json:"UserUniID"`       // 用户ID
	NickName     *string `protobuf:"bytes,2,opt,name=NickName" json:"NickName"`          // 昵称
	HeadImage    *string `protobuf:"bytes,3,opt,name=HeadImage" json:"HeadImage"`        // 头像
	RegisterTime *uint32 `protobuf:"varint,4,opt,name=RegisterTime" json:"RegisterTime"` // 注册时间
	PlayerLevel  *int32  `protobuf:"varint,5,opt,name=PlayerLevel" json:"PlayerLevel"`   // 等级
	GoldCoin     *int32  `protobuf:"varint,6,opt,name=GoldCoin" json:"GoldCoin"`         // 金币
	Diamond      *int32  `protobuf:"varint,7,opt,name=Diamond" json:"Diamond"`           // 钻石
	Strength     *int32  `protobuf:"varint,8,opt,name=Strength" json:"Strength"`         // 体力
	Experience   *int32  `protobuf:"varint,9,opt,name=Experience" json:"Experience"`     // 经验
}

func (x *PBCommonData) Reset() {
	*x = PBCommonData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBCommonData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBCommonData) ProtoMessage() {}

func (x *PBCommonData) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBCommonData.ProtoReflect.Descriptor instead.
func (*PBCommonData) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{1}
}

func (x *PBCommonData) GetUserUniID() int64 {
	if x != nil && x.UserUniID != nil {
		return *x.UserUniID
	}
	return 0
}

func (x *PBCommonData) GetNickName() string {
	if x != nil && x.NickName != nil {
		return *x.NickName
	}
	return ""
}

func (x *PBCommonData) GetHeadImage() string {
	if x != nil && x.HeadImage != nil {
		return *x.HeadImage
	}
	return ""
}

func (x *PBCommonData) GetRegisterTime() uint32 {
	if x != nil && x.RegisterTime != nil {
		return *x.RegisterTime
	}
	return 0
}

func (x *PBCommonData) GetPlayerLevel() int32 {
	if x != nil && x.PlayerLevel != nil {
		return *x.PlayerLevel
	}
	return 0
}

func (x *PBCommonData) GetGoldCoin() int32 {
	if x != nil && x.GoldCoin != nil {
		return *x.GoldCoin
	}
	return 0
}

func (x *PBCommonData) GetDiamond() int32 {
	if x != nil && x.Diamond != nil {
		return *x.Diamond
	}
	return 0
}

func (x *PBCommonData) GetStrength() int32 {
	if x != nil && x.Strength != nil {
		return *x.Strength
	}
	return 0
}

func (x *PBCommonData) GetExperience() int32 {
	if x != nil && x.Experience != nil {
		return *x.Experience
	}
	return 0
}

// 玩家登录请求
type PlayerLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoginType   *string `protobuf:"bytes,1,opt,name=LoginType" json:"LoginType"`     // 登录类型(quick)
	Account     *string `protobuf:"bytes,2,opt,name=Account" json:"Account"`         // 用户名
	PassWord    *string `protobuf:"bytes,3,opt,name=PassWord" json:"PassWord"`       // 密码
	ChannelType *string `protobuf:"bytes,4,opt,name=ChannelType" json:"ChannelType"` // 渠道标识
}

func (x *PlayerLoginRequest) Reset() {
	*x = PlayerLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerLoginRequest) ProtoMessage() {}

func (x *PlayerLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerLoginRequest.ProtoReflect.Descriptor instead.
func (*PlayerLoginRequest) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{2}
}

func (x *PlayerLoginRequest) GetLoginType() string {
	if x != nil && x.LoginType != nil {
		return *x.LoginType
	}
	return ""
}

func (x *PlayerLoginRequest) GetAccount() string {
	if x != nil && x.Account != nil {
		return *x.Account
	}
	return ""
}

func (x *PlayerLoginRequest) GetPassWord() string {
	if x != nil && x.PassWord != nil {
		return *x.PassWord
	}
	return ""
}

func (x *PlayerLoginRequest) GetChannelType() string {
	if x != nil && x.ChannelType != nil {
		return *x.ChannelType
	}
	return ""
}

// 玩家登录回应
type PlayerLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result       *int32              `protobuf:"varint,1,opt,name=Result" json:"Result"`             // 错误码
	ReqData      *PlayerLoginRequest `protobuf:"bytes,2,opt,name=ReqData" json:"ReqData"`            // 请求数据
	UserUniID    *int64              `protobuf:"varint,3,opt,name=UserUniID" json:"UserUniID"`       // 用户ID
	Account      *string             `protobuf:"bytes,4,opt,name=Account" json:"Account"`            // 账号
	Password     *string             `protobuf:"bytes,5,opt,name=Password" json:"Password"`          // 密码
	Register     *uint32             `protobuf:"varint,6,opt,name=Register" json:"Register"`         // 是否新玩家
	RegisterTime *uint32             `protobuf:"varint,7,opt,name=RegisterTime" json:"RegisterTime"` // 注册时间
	RandomSeed   *uint32             `protobuf:"varint,8,opt,name=RandomSeed" json:"RandomSeed"`     // 随机种子
}

func (x *PlayerLoginResponse) Reset() {
	*x = PlayerLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerLoginResponse) ProtoMessage() {}

func (x *PlayerLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerLoginResponse.ProtoReflect.Descriptor instead.
func (*PlayerLoginResponse) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{3}
}

func (x *PlayerLoginResponse) GetResult() int32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *PlayerLoginResponse) GetReqData() *PlayerLoginRequest {
	if x != nil {
		return x.ReqData
	}
	return nil
}

func (x *PlayerLoginResponse) GetUserUniID() int64 {
	if x != nil && x.UserUniID != nil {
		return *x.UserUniID
	}
	return 0
}

func (x *PlayerLoginResponse) GetAccount() string {
	if x != nil && x.Account != nil {
		return *x.Account
	}
	return ""
}

func (x *PlayerLoginResponse) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *PlayerLoginResponse) GetRegister() uint32 {
	if x != nil && x.Register != nil {
		return *x.Register
	}
	return 0
}

func (x *PlayerLoginResponse) GetRegisterTime() uint32 {
	if x != nil && x.RegisterTime != nil {
		return *x.RegisterTime
	}
	return 0
}

func (x *PlayerLoginResponse) GetRandomSeed() uint32 {
	if x != nil && x.RandomSeed != nil {
		return *x.RandomSeed
	}
	return 0
}

// 心跳请求
type HeartbeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HeartbeatRequest) Reset() {
	*x = HeartbeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatRequest) ProtoMessage() {}

func (x *HeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatRequest.ProtoReflect.Descriptor instead.
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{4}
}

// 心跳响应
type HeartbeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerTime *uint32 `protobuf:"varint,1,opt,name=ServerTime" json:"ServerTime"` // 服务器时间（时间戳，秒）
}

func (x *HeartbeatResponse) Reset() {
	*x = HeartbeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatResponse) ProtoMessage() {}

func (x *HeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatResponse.ProtoReflect.Descriptor instead.
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{5}
}

func (x *HeartbeatResponse) GetServerTime() uint32 {
	if x != nil && x.ServerTime != nil {
		return *x.ServerTime
	}
	return 0
}

type BroadcastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Str *string `protobuf:"bytes,1,opt,name=Str" json:"Str"`
}

func (x *BroadcastRequest) Reset() {
	*x = BroadcastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastRequest) ProtoMessage() {}

func (x *BroadcastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastRequest.ProtoReflect.Descriptor instead.
func (*BroadcastRequest) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{6}
}

func (x *BroadcastRequest) GetStr() string {
	if x != nil && x.Str != nil {
		return *x.Str
	}
	return ""
}

type BroadcastResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *uint32 `protobuf:"varint,1,opt,name=Result" json:"Result"`
	Str    *string `protobuf:"bytes,2,opt,name=Str" json:"Str"`
}

func (x *BroadcastResponse) Reset() {
	*x = BroadcastResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastResponse) ProtoMessage() {}

func (x *BroadcastResponse) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastResponse.ProtoReflect.Descriptor instead.
func (*BroadcastResponse) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{7}
}

func (x *BroadcastResponse) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *BroadcastResponse) GetStr() string {
	if x != nil && x.Str != nil {
		return *x.Str
	}
	return ""
}

var File_player_proto protoreflect.FileDescriptor

var file_player_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x40, 0x0a, 0x0c, 0x50, 0x42, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x30, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x42, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x9e, 0x02, 0x0a, 0x0c, 0x50, 0x42, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x55, 0x6e, 0x69,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x55, 0x73, 0x65, 0x72, 0x55, 0x6e,
	0x69, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x48, 0x65, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x47, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x69, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x47, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x69, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x44, 0x69, 0x61, 0x6d, 0x6f, 0x6e, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x44, 0x69, 0x61, 0x6d, 0x6f, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x74, 0x72,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x53, 0x74, 0x72,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x45, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x8a, 0x01, 0x0a, 0x12, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x93, 0x02, 0x0a, 0x13, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x30, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x52, 0x65, 0x71,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x55, 0x6e, 0x69, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x55, 0x73, 0x65, 0x72, 0x55, 0x6e, 0x69,
	0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x61, 0x6e, 0x64,
	0x6f, 0x6d, 0x53, 0x65, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x52, 0x61,
	0x6e, 0x64, 0x6f, 0x6d, 0x53, 0x65, 0x65, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x33, 0x0a, 0x11,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x24, 0x0a, 0x10, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x53, 0x74, 0x72, 0x22, 0x3d, 0x0a, 0x11, 0x42, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x74, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x53, 0x74, 0x72, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62,
}

var (
	file_player_proto_rawDescOnce sync.Once
	file_player_proto_rawDescData = file_player_proto_rawDesc
)

func file_player_proto_rawDescGZIP() []byte {
	file_player_proto_rawDescOnce.Do(func() {
		file_player_proto_rawDescData = protoimpl.X.CompressGZIP(file_player_proto_rawDescData)
	})
	return file_player_proto_rawDescData
}

var file_player_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_player_proto_goTypes = []interface{}{
	(*PBPlayerData)(nil),        // 0: pb.PBPlayerData
	(*PBCommonData)(nil),        // 1: pb.PBCommonData
	(*PlayerLoginRequest)(nil),  // 2: pb.PlayerLoginRequest
	(*PlayerLoginResponse)(nil), // 3: pb.PlayerLoginResponse
	(*HeartbeatRequest)(nil),    // 4: pb.HeartbeatRequest
	(*HeartbeatResponse)(nil),   // 5: pb.HeartbeatResponse
	(*BroadcastRequest)(nil),    // 6: pb.BroadcastRequest
	(*BroadcastResponse)(nil),   // 7: pb.BroadcastResponse
}
var file_player_proto_depIdxs = []int32{
	1, // 0: pb.PBPlayerData.CommonData:type_name -> pb.PBCommonData
	2, // 1: pb.PlayerLoginResponse.ReqData:type_name -> pb.PlayerLoginRequest
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_player_proto_init() }
func file_player_proto_init() {
	if File_player_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_player_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBPlayerData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_player_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBCommonData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_player_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerLoginRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_player_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerLoginResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_player_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_player_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_player_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_player_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_player_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_player_proto_goTypes,
		DependencyIndexes: file_player_proto_depIdxs,
		MessageInfos:      file_player_proto_msgTypes,
	}.Build()
	File_player_proto = out.File
	file_player_proto_rawDesc = nil
	file_player_proto_goTypes = nil
	file_player_proto_depIdxs = nil
}
