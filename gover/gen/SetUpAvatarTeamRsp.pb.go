// https://github.com/SlushinPS/beach-simulator
// Copyright (C) 2023 Slushy Team
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: SetUpAvatarTeamRsp.proto

package gen

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

// CmdId: 1613
// Obf: EBLCDMCNFJI
type SetUpAvatarTeamRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurAvatarGuid      uint64   `protobuf:"varint,5,opt,name=cur_avatar_guid,json=curAvatarGuid,proto3" json:"cur_avatar_guid,omitempty"`
	TeamId             uint32   `protobuf:"varint,2,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	Retcode            int32    `protobuf:"varint,8,opt,name=retcode,proto3" json:"retcode,omitempty"`
	AvatarTeamGuidList []uint64 `protobuf:"varint,12,rep,packed,name=avatar_team_guid_list,json=avatarTeamGuidList,proto3" json:"avatar_team_guid_list,omitempty"`
}

func (x *SetUpAvatarTeamRsp) Reset() {
	*x = SetUpAvatarTeamRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SetUpAvatarTeamRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetUpAvatarTeamRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUpAvatarTeamRsp) ProtoMessage() {}

func (x *SetUpAvatarTeamRsp) ProtoReflect() protoreflect.Message {
	mi := &file_SetUpAvatarTeamRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUpAvatarTeamRsp.ProtoReflect.Descriptor instead.
func (*SetUpAvatarTeamRsp) Descriptor() ([]byte, []int) {
	return file_SetUpAvatarTeamRsp_proto_rawDescGZIP(), []int{0}
}

func (x *SetUpAvatarTeamRsp) GetCurAvatarGuid() uint64 {
	if x != nil {
		return x.CurAvatarGuid
	}
	return 0
}

func (x *SetUpAvatarTeamRsp) GetTeamId() uint32 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *SetUpAvatarTeamRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *SetUpAvatarTeamRsp) GetAvatarTeamGuidList() []uint64 {
	if x != nil {
		return x.AvatarTeamGuidList
	}
	return nil
}

var File_SetUpAvatarTeamRsp_proto protoreflect.FileDescriptor

var file_SetUpAvatarTeamRsp_proto_rawDesc = []byte{
	0x0a, 0x18, 0x53, 0x65, 0x74, 0x55, 0x70, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x65, 0x61,
	0x6d, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x12, 0x53,
	0x65, 0x74, 0x55, 0x70, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x73,
	0x70, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x75, 0x72, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f,
	0x67, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x47, 0x75, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x31, 0x0a, 0x15,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x67, 0x75, 0x69, 0x64,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x04, 0x52, 0x12, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x54, 0x65, 0x61, 0x6d, 0x47, 0x75, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SetUpAvatarTeamRsp_proto_rawDescOnce sync.Once
	file_SetUpAvatarTeamRsp_proto_rawDescData = file_SetUpAvatarTeamRsp_proto_rawDesc
)

func file_SetUpAvatarTeamRsp_proto_rawDescGZIP() []byte {
	file_SetUpAvatarTeamRsp_proto_rawDescOnce.Do(func() {
		file_SetUpAvatarTeamRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_SetUpAvatarTeamRsp_proto_rawDescData)
	})
	return file_SetUpAvatarTeamRsp_proto_rawDescData
}

var file_SetUpAvatarTeamRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SetUpAvatarTeamRsp_proto_goTypes = []interface{}{
	(*SetUpAvatarTeamRsp)(nil), // 0: SetUpAvatarTeamRsp
}
var file_SetUpAvatarTeamRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SetUpAvatarTeamRsp_proto_init() }
func file_SetUpAvatarTeamRsp_proto_init() {
	if File_SetUpAvatarTeamRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SetUpAvatarTeamRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetUpAvatarTeamRsp); i {
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
			RawDescriptor: file_SetUpAvatarTeamRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SetUpAvatarTeamRsp_proto_goTypes,
		DependencyIndexes: file_SetUpAvatarTeamRsp_proto_depIdxs,
		MessageInfos:      file_SetUpAvatarTeamRsp_proto_msgTypes,
	}.Build()
	File_SetUpAvatarTeamRsp_proto = out.File
	file_SetUpAvatarTeamRsp_proto_rawDesc = nil
	file_SetUpAvatarTeamRsp_proto_goTypes = nil
	file_SetUpAvatarTeamRsp_proto_depIdxs = nil
}
