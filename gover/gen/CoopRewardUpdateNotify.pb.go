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
// source: CoopRewardUpdateNotify.proto

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

// CmdId: 1966
// Obf: CKLPHKLCLJF
type CoopRewardUpdateNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RewardList []*CoopReward `protobuf:"bytes,8,rep,name=reward_list,json=rewardList,proto3" json:"reward_list,omitempty"`
}

func (x *CoopRewardUpdateNotify) Reset() {
	*x = CoopRewardUpdateNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CoopRewardUpdateNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoopRewardUpdateNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoopRewardUpdateNotify) ProtoMessage() {}

func (x *CoopRewardUpdateNotify) ProtoReflect() protoreflect.Message {
	mi := &file_CoopRewardUpdateNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoopRewardUpdateNotify.ProtoReflect.Descriptor instead.
func (*CoopRewardUpdateNotify) Descriptor() ([]byte, []int) {
	return file_CoopRewardUpdateNotify_proto_rawDescGZIP(), []int{0}
}

func (x *CoopRewardUpdateNotify) GetRewardList() []*CoopReward {
	if x != nil {
		return x.RewardList
	}
	return nil
}

var File_CoopRewardUpdateNotify_proto protoreflect.FileDescriptor

var file_CoopRewardUpdateNotify_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x43, 0x6f, 0x6f, 0x70, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10,
	0x43, 0x6f, 0x6f, 0x70, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x46, 0x0a, 0x16, 0x43, 0x6f, 0x6f, 0x70, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x2c, 0x0a, 0x0b, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x43, 0x6f, 0x6f, 0x70, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x0a, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CoopRewardUpdateNotify_proto_rawDescOnce sync.Once
	file_CoopRewardUpdateNotify_proto_rawDescData = file_CoopRewardUpdateNotify_proto_rawDesc
)

func file_CoopRewardUpdateNotify_proto_rawDescGZIP() []byte {
	file_CoopRewardUpdateNotify_proto_rawDescOnce.Do(func() {
		file_CoopRewardUpdateNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_CoopRewardUpdateNotify_proto_rawDescData)
	})
	return file_CoopRewardUpdateNotify_proto_rawDescData
}

var file_CoopRewardUpdateNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CoopRewardUpdateNotify_proto_goTypes = []interface{}{
	(*CoopRewardUpdateNotify)(nil), // 0: CoopRewardUpdateNotify
	(*CoopReward)(nil),             // 1: CoopReward
}
var file_CoopRewardUpdateNotify_proto_depIdxs = []int32{
	1, // 0: CoopRewardUpdateNotify.reward_list:type_name -> CoopReward
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CoopRewardUpdateNotify_proto_init() }
func file_CoopRewardUpdateNotify_proto_init() {
	if File_CoopRewardUpdateNotify_proto != nil {
		return
	}
	file_CoopReward_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CoopRewardUpdateNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoopRewardUpdateNotify); i {
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
			RawDescriptor: file_CoopRewardUpdateNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CoopRewardUpdateNotify_proto_goTypes,
		DependencyIndexes: file_CoopRewardUpdateNotify_proto_depIdxs,
		MessageInfos:      file_CoopRewardUpdateNotify_proto_msgTypes,
	}.Build()
	File_CoopRewardUpdateNotify_proto = out.File
	file_CoopRewardUpdateNotify_proto_rawDesc = nil
	file_CoopRewardUpdateNotify_proto_goTypes = nil
	file_CoopRewardUpdateNotify_proto_depIdxs = nil
}
