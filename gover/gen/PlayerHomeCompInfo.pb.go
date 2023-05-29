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
// source: PlayerHomeCompInfo.proto

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

// Obf: KOGGBOKJHGO
type PlayerHomeCompInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FriendEnterHomeOption     FriendEnterHomeOption `protobuf:"varint,14,opt,name=friend_enter_home_option,json=friendEnterHomeOption,proto3,enum=FriendEnterHomeOption" json:"friend_enter_home_option,omitempty"`
	LevelupRewardGotLevelList []uint32              `protobuf:"varint,15,rep,packed,name=levelup_reward_got_level_list,json=levelupRewardGotLevelList,proto3" json:"levelup_reward_got_level_list,omitempty"`
	SeenModuleIdList          []uint32              `protobuf:"varint,2,rep,packed,name=seen_module_id_list,json=seenModuleIdList,proto3" json:"seen_module_id_list,omitempty"`
	UnlockedModuleIdList      []uint32              `protobuf:"varint,1,rep,packed,name=unlocked_module_id_list,json=unlockedModuleIdList,proto3" json:"unlocked_module_id_list,omitempty"`
}

func (x *PlayerHomeCompInfo) Reset() {
	*x = PlayerHomeCompInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerHomeCompInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerHomeCompInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerHomeCompInfo) ProtoMessage() {}

func (x *PlayerHomeCompInfo) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerHomeCompInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerHomeCompInfo.ProtoReflect.Descriptor instead.
func (*PlayerHomeCompInfo) Descriptor() ([]byte, []int) {
	return file_PlayerHomeCompInfo_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerHomeCompInfo) GetFriendEnterHomeOption() FriendEnterHomeOption {
	if x != nil {
		return x.FriendEnterHomeOption
	}
	return FriendEnterHomeOption_FRIEND_ENTER_HOME_OPTION_NEED_CONFIRM
}

func (x *PlayerHomeCompInfo) GetLevelupRewardGotLevelList() []uint32 {
	if x != nil {
		return x.LevelupRewardGotLevelList
	}
	return nil
}

func (x *PlayerHomeCompInfo) GetSeenModuleIdList() []uint32 {
	if x != nil {
		return x.SeenModuleIdList
	}
	return nil
}

func (x *PlayerHomeCompInfo) GetUnlockedModuleIdList() []uint32 {
	if x != nil {
		return x.UnlockedModuleIdList
	}
	return nil
}

var File_PlayerHomeCompInfo_proto protoreflect.FileDescriptor

var file_PlayerHomeCompInfo_proto_rawDesc = []byte{
	0x0a, 0x18, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x48, 0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x48, 0x6f, 0x6d, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x02, 0x0a, 0x12, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x48, 0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4f,
	0x0a, 0x18, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x68,
	0x6f, 0x6d, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x16, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x48, 0x6f,
	0x6d, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x15, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x45, 0x6e, 0x74, 0x65, 0x72, 0x48, 0x6f, 0x6d, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x40, 0x0a, 0x1d, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x75, 0x70, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x5f, 0x67, 0x6f, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x0f, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x19, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x75, 0x70, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x47, 0x6f, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x2d, 0x0a, 0x13, 0x73, 0x65, 0x65, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x10,
	0x73, 0x65, 0x65, 0x6e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x35, 0x0a, 0x17, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x14, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerHomeCompInfo_proto_rawDescOnce sync.Once
	file_PlayerHomeCompInfo_proto_rawDescData = file_PlayerHomeCompInfo_proto_rawDesc
)

func file_PlayerHomeCompInfo_proto_rawDescGZIP() []byte {
	file_PlayerHomeCompInfo_proto_rawDescOnce.Do(func() {
		file_PlayerHomeCompInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerHomeCompInfo_proto_rawDescData)
	})
	return file_PlayerHomeCompInfo_proto_rawDescData
}

var file_PlayerHomeCompInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerHomeCompInfo_proto_goTypes = []interface{}{
	(*PlayerHomeCompInfo)(nil), // 0: PlayerHomeCompInfo
	(FriendEnterHomeOption)(0), // 1: FriendEnterHomeOption
}
var file_PlayerHomeCompInfo_proto_depIdxs = []int32{
	1, // 0: PlayerHomeCompInfo.friend_enter_home_option:type_name -> FriendEnterHomeOption
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PlayerHomeCompInfo_proto_init() }
func file_PlayerHomeCompInfo_proto_init() {
	if File_PlayerHomeCompInfo_proto != nil {
		return
	}
	file_FriendEnterHomeOption_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerHomeCompInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerHomeCompInfo); i {
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
			RawDescriptor: file_PlayerHomeCompInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerHomeCompInfo_proto_goTypes,
		DependencyIndexes: file_PlayerHomeCompInfo_proto_depIdxs,
		MessageInfos:      file_PlayerHomeCompInfo_proto_msgTypes,
	}.Build()
	File_PlayerHomeCompInfo_proto = out.File
	file_PlayerHomeCompInfo_proto_rawDesc = nil
	file_PlayerHomeCompInfo_proto_goTypes = nil
	file_PlayerHomeCompInfo_proto_depIdxs = nil
}
