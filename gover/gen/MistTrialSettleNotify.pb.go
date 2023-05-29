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
// source: MistTrialSettleNotify.proto

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

// CmdId: 8153
// Obf: ICIKIBNLBCG
type MistTrialSettleNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DungeonSceneId uint32                 `protobuf:"varint,3,opt,name=dungeon_scene_id,json=dungeonSceneId,proto3" json:"dungeon_scene_id,omitempty"`
	ECFLBCKABPE    map[uint32]uint32      `protobuf:"bytes,6,rep,name=ECFLBCKABPE,proto3" json:"ECFLBCKABPE,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	BestAvatarList []*MistTrialBestAvatar `protobuf:"bytes,5,rep,name=best_avatar_list,json=bestAvatarList,proto3" json:"best_avatar_list,omitempty"`
	BestHitAvatar  *MistTrialBestAvatar   `protobuf:"bytes,10,opt,name=best_hit_avatar,json=bestHitAvatar,proto3" json:"best_hit_avatar,omitempty"`
	MLPGIMJCPBJ    map[uint32]uint32      `protobuf:"bytes,1,rep,name=MLPGIMJCPBJ,proto3" json:"MLPGIMJCPBJ,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	FirstPassTime  uint32                 `protobuf:"varint,13,opt,name=first_pass_time,json=firstPassTime,proto3" json:"first_pass_time,omitempty"`
}

func (x *MistTrialSettleNotify) Reset() {
	*x = MistTrialSettleNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MistTrialSettleNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MistTrialSettleNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MistTrialSettleNotify) ProtoMessage() {}

func (x *MistTrialSettleNotify) ProtoReflect() protoreflect.Message {
	mi := &file_MistTrialSettleNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MistTrialSettleNotify.ProtoReflect.Descriptor instead.
func (*MistTrialSettleNotify) Descriptor() ([]byte, []int) {
	return file_MistTrialSettleNotify_proto_rawDescGZIP(), []int{0}
}

func (x *MistTrialSettleNotify) GetDungeonSceneId() uint32 {
	if x != nil {
		return x.DungeonSceneId
	}
	return 0
}

func (x *MistTrialSettleNotify) GetECFLBCKABPE() map[uint32]uint32 {
	if x != nil {
		return x.ECFLBCKABPE
	}
	return nil
}

func (x *MistTrialSettleNotify) GetBestAvatarList() []*MistTrialBestAvatar {
	if x != nil {
		return x.BestAvatarList
	}
	return nil
}

func (x *MistTrialSettleNotify) GetBestHitAvatar() *MistTrialBestAvatar {
	if x != nil {
		return x.BestHitAvatar
	}
	return nil
}

func (x *MistTrialSettleNotify) GetMLPGIMJCPBJ() map[uint32]uint32 {
	if x != nil {
		return x.MLPGIMJCPBJ
	}
	return nil
}

func (x *MistTrialSettleNotify) GetFirstPassTime() uint32 {
	if x != nil {
		return x.FirstPassTime
	}
	return 0
}

var File_MistTrialSettleNotify_proto protoreflect.FileDescriptor

var file_MistTrialSettleNotify_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x4d, 0x69, 0x73, 0x74, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x6c,
	0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x4d,
	0x69, 0x73, 0x74, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x42, 0x65, 0x73, 0x74, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x03, 0x0a, 0x15, 0x4d, 0x69, 0x73,
	0x74, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x12, 0x28, 0x0a, 0x10, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x73, 0x63,
	0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x64, 0x75,
	0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x49, 0x0a, 0x0b,
	0x45, 0x43, 0x46, 0x4c, 0x42, 0x43, 0x4b, 0x41, 0x42, 0x50, 0x45, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x4d, 0x69, 0x73, 0x74, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x65, 0x74,
	0x74, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x45, 0x43, 0x46, 0x4c, 0x42, 0x43,
	0x4b, 0x41, 0x42, 0x50, 0x45, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x45, 0x43, 0x46, 0x4c,
	0x42, 0x43, 0x4b, 0x41, 0x42, 0x50, 0x45, 0x12, 0x3e, 0x0a, 0x10, 0x62, 0x65, 0x73, 0x74, 0x5f,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x4d, 0x69, 0x73, 0x74, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x42, 0x65, 0x73,
	0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x0e, 0x62, 0x65, 0x73, 0x74, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0f, 0x62, 0x65, 0x73, 0x74, 0x5f,
	0x68, 0x69, 0x74, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x4d, 0x69, 0x73, 0x74, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x42, 0x65, 0x73, 0x74,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x0d, 0x62, 0x65, 0x73, 0x74, 0x48, 0x69, 0x74, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x49, 0x0a, 0x0b, 0x4d, 0x4c, 0x50, 0x47, 0x49, 0x4d, 0x4a,
	0x43, 0x50, 0x42, 0x4a, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x4d, 0x69, 0x73,
	0x74, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x4d, 0x4c, 0x50, 0x47, 0x49, 0x4d, 0x4a, 0x43, 0x50, 0x42, 0x4a, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0b, 0x4d, 0x4c, 0x50, 0x47, 0x49, 0x4d, 0x4a, 0x43, 0x50, 0x42, 0x4a,
	0x12, 0x26, 0x0a, 0x0f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x50, 0x61, 0x73, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x1a, 0x3e, 0x0a, 0x10, 0x45, 0x43, 0x46, 0x4c,
	0x42, 0x43, 0x4b, 0x41, 0x42, 0x50, 0x45, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x4d, 0x4c, 0x50, 0x47,
	0x49, 0x4d, 0x4a, 0x43, 0x50, 0x42, 0x4a, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MistTrialSettleNotify_proto_rawDescOnce sync.Once
	file_MistTrialSettleNotify_proto_rawDescData = file_MistTrialSettleNotify_proto_rawDesc
)

func file_MistTrialSettleNotify_proto_rawDescGZIP() []byte {
	file_MistTrialSettleNotify_proto_rawDescOnce.Do(func() {
		file_MistTrialSettleNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_MistTrialSettleNotify_proto_rawDescData)
	})
	return file_MistTrialSettleNotify_proto_rawDescData
}

var file_MistTrialSettleNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_MistTrialSettleNotify_proto_goTypes = []interface{}{
	(*MistTrialSettleNotify)(nil), // 0: MistTrialSettleNotify
	nil,                           // 1: MistTrialSettleNotify.ECFLBCKABPEEntry
	nil,                           // 2: MistTrialSettleNotify.MLPGIMJCPBJEntry
	(*MistTrialBestAvatar)(nil),   // 3: MistTrialBestAvatar
}
var file_MistTrialSettleNotify_proto_depIdxs = []int32{
	1, // 0: MistTrialSettleNotify.ECFLBCKABPE:type_name -> MistTrialSettleNotify.ECFLBCKABPEEntry
	3, // 1: MistTrialSettleNotify.best_avatar_list:type_name -> MistTrialBestAvatar
	3, // 2: MistTrialSettleNotify.best_hit_avatar:type_name -> MistTrialBestAvatar
	2, // 3: MistTrialSettleNotify.MLPGIMJCPBJ:type_name -> MistTrialSettleNotify.MLPGIMJCPBJEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_MistTrialSettleNotify_proto_init() }
func file_MistTrialSettleNotify_proto_init() {
	if File_MistTrialSettleNotify_proto != nil {
		return
	}
	file_MistTrialBestAvatar_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MistTrialSettleNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MistTrialSettleNotify); i {
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
			RawDescriptor: file_MistTrialSettleNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MistTrialSettleNotify_proto_goTypes,
		DependencyIndexes: file_MistTrialSettleNotify_proto_depIdxs,
		MessageInfos:      file_MistTrialSettleNotify_proto_msgTypes,
	}.Build()
	File_MistTrialSettleNotify_proto = out.File
	file_MistTrialSettleNotify_proto_rawDesc = nil
	file_MistTrialSettleNotify_proto_goTypes = nil
	file_MistTrialSettleNotify_proto_depIdxs = nil
}
