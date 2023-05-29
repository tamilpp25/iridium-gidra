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
// source: StartRogueDiaryPlayRsp.proto

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

// CmdId: 8460
// Obf: EOLEDCMNJMF
type StartRogueDiaryPlayRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrialAvatarList []*RogueDiaryAvatar `protobuf:"bytes,2,rep,name=trial_avatar_list,json=trialAvatarList,proto3" json:"trial_avatar_list,omitempty"`
	Retcode         int32               `protobuf:"varint,7,opt,name=retcode,proto3" json:"retcode,omitempty"`
	AvatarList      []*RogueDiaryAvatar `protobuf:"bytes,12,rep,name=avatar_list,json=avatarList,proto3" json:"avatar_list,omitempty"`
	RandCardList    []uint32            `protobuf:"varint,3,rep,packed,name=rand_card_list,json=randCardList,proto3" json:"rand_card_list,omitempty"`
}

func (x *StartRogueDiaryPlayRsp) Reset() {
	*x = StartRogueDiaryPlayRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_StartRogueDiaryPlayRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRogueDiaryPlayRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRogueDiaryPlayRsp) ProtoMessage() {}

func (x *StartRogueDiaryPlayRsp) ProtoReflect() protoreflect.Message {
	mi := &file_StartRogueDiaryPlayRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRogueDiaryPlayRsp.ProtoReflect.Descriptor instead.
func (*StartRogueDiaryPlayRsp) Descriptor() ([]byte, []int) {
	return file_StartRogueDiaryPlayRsp_proto_rawDescGZIP(), []int{0}
}

func (x *StartRogueDiaryPlayRsp) GetTrialAvatarList() []*RogueDiaryAvatar {
	if x != nil {
		return x.TrialAvatarList
	}
	return nil
}

func (x *StartRogueDiaryPlayRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *StartRogueDiaryPlayRsp) GetAvatarList() []*RogueDiaryAvatar {
	if x != nil {
		return x.AvatarList
	}
	return nil
}

func (x *StartRogueDiaryPlayRsp) GetRandCardList() []uint32 {
	if x != nil {
		return x.RandCardList
	}
	return nil
}

var File_StartRogueDiaryPlayRsp_proto protoreflect.FileDescriptor

var file_StartRogueDiaryPlayRsp_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61, 0x72,
	0x79, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61, 0x72, 0x79, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcb, 0x01, 0x0a, 0x16, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61, 0x72, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x73,
	0x70, 0x12, 0x3d, 0x0a, 0x11, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61, 0x72, 0x79, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52,
	0x0f, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a, 0x0b, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61, 0x72, 0x79, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24,
	0x0a, 0x0e, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x72, 0x61, 0x6e, 0x64, 0x43, 0x61, 0x72, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_StartRogueDiaryPlayRsp_proto_rawDescOnce sync.Once
	file_StartRogueDiaryPlayRsp_proto_rawDescData = file_StartRogueDiaryPlayRsp_proto_rawDesc
)

func file_StartRogueDiaryPlayRsp_proto_rawDescGZIP() []byte {
	file_StartRogueDiaryPlayRsp_proto_rawDescOnce.Do(func() {
		file_StartRogueDiaryPlayRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_StartRogueDiaryPlayRsp_proto_rawDescData)
	})
	return file_StartRogueDiaryPlayRsp_proto_rawDescData
}

var file_StartRogueDiaryPlayRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_StartRogueDiaryPlayRsp_proto_goTypes = []interface{}{
	(*StartRogueDiaryPlayRsp)(nil), // 0: StartRogueDiaryPlayRsp
	(*RogueDiaryAvatar)(nil),       // 1: RogueDiaryAvatar
}
var file_StartRogueDiaryPlayRsp_proto_depIdxs = []int32{
	1, // 0: StartRogueDiaryPlayRsp.trial_avatar_list:type_name -> RogueDiaryAvatar
	1, // 1: StartRogueDiaryPlayRsp.avatar_list:type_name -> RogueDiaryAvatar
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_StartRogueDiaryPlayRsp_proto_init() }
func file_StartRogueDiaryPlayRsp_proto_init() {
	if File_StartRogueDiaryPlayRsp_proto != nil {
		return
	}
	file_RogueDiaryAvatar_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_StartRogueDiaryPlayRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRogueDiaryPlayRsp); i {
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
			RawDescriptor: file_StartRogueDiaryPlayRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_StartRogueDiaryPlayRsp_proto_goTypes,
		DependencyIndexes: file_StartRogueDiaryPlayRsp_proto_depIdxs,
		MessageInfos:      file_StartRogueDiaryPlayRsp_proto_msgTypes,
	}.Build()
	File_StartRogueDiaryPlayRsp_proto = out.File
	file_StartRogueDiaryPlayRsp_proto_rawDesc = nil
	file_StartRogueDiaryPlayRsp_proto_goTypes = nil
	file_StartRogueDiaryPlayRsp_proto_depIdxs = nil
}
