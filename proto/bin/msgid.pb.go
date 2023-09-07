// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.22.2
// source: msgid.proto

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

type MsgID int32

const (
	MsgID_PlayerLogin_Req MsgID = 1001 // PlayerLoginReq 玩家登录请求
	MsgID_PlayerLogin_Res MsgID = 1002 // PlayerLoginRes 玩家登录回应
)

// Enum value maps for MsgID.
var (
	MsgID_name = map[int32]string{
		1001: "PlayerLogin_Req",
		1002: "PlayerLogin_Res",
	}
	MsgID_value = map[string]int32{
		"PlayerLogin_Req": 1001,
		"PlayerLogin_Res": 1002,
	}
)

func (x MsgID) Enum() *MsgID {
	p := new(MsgID)
	*p = x
	return p
}

func (x MsgID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgID) Descriptor() protoreflect.EnumDescriptor {
	return file_msgid_proto_enumTypes[0].Descriptor()
}

func (MsgID) Type() protoreflect.EnumType {
	return &file_msgid_proto_enumTypes[0]
}

func (x MsgID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *MsgID) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = MsgID(num)
	return nil
}

// Deprecated: Use MsgID.Descriptor instead.
func (MsgID) EnumDescriptor() ([]byte, []int) {
	return file_msgid_proto_rawDescGZIP(), []int{0}
}

var File_msgid_proto protoreflect.FileDescriptor

var file_msgid_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x73, 0x67, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x2a, 0x33, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x0f, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x52, 0x65, 0x71, 0x10, 0xe9, 0x07,
	0x12, 0x14, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x5f,
	0x52, 0x65, 0x73, 0x10, 0xea, 0x07, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62,
}

var (
	file_msgid_proto_rawDescOnce sync.Once
	file_msgid_proto_rawDescData = file_msgid_proto_rawDesc
)

func file_msgid_proto_rawDescGZIP() []byte {
	file_msgid_proto_rawDescOnce.Do(func() {
		file_msgid_proto_rawDescData = protoimpl.X.CompressGZIP(file_msgid_proto_rawDescData)
	})
	return file_msgid_proto_rawDescData
}

var file_msgid_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_msgid_proto_goTypes = []interface{}{
	(MsgID)(0), // 0: pb.MsgID
}
var file_msgid_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_msgid_proto_init() }
func file_msgid_proto_init() {
	if File_msgid_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msgid_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msgid_proto_goTypes,
		DependencyIndexes: file_msgid_proto_depIdxs,
		EnumInfos:         file_msgid_proto_enumTypes,
	}.Build()
	File_msgid_proto = out.File
	file_msgid_proto_rawDesc = nil
	file_msgid_proto_goTypes = nil
	file_msgid_proto_depIdxs = nil
}
