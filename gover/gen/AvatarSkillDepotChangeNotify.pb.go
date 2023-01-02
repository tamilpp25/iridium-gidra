// Sorapointa - A server software re-implementation for a certain anime game, and avoid sorapointa.
// Copyright (C) 2022  Sorapointa Team
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
// 	protoc        v3.11.3
// source: AvatarSkillDepotChangeNotify.proto

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

type AvatarSkillDepotChangeNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProudSkillExtraLevelMap map[uint32]uint32 `protobuf:"bytes,9,rep,name=proud_skill_extra_level_map,json=proudSkillExtraLevelMap,proto3" json:"proud_skill_extra_level_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	ProudSkillList          []uint32          `protobuf:"varint,3,rep,packed,name=proud_skill_list,json=proudSkillList,proto3" json:"proud_skill_list,omitempty"`
	EntityId                uint32            `protobuf:"varint,5,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	AvatarGuid              uint64            `protobuf:"varint,7,opt,name=avatar_guid,json=avatarGuid,proto3" json:"avatar_guid,omitempty"`
	SkillLevelMap           map[uint32]uint32 `protobuf:"bytes,1,rep,name=skill_level_map,json=skillLevelMap,proto3" json:"skill_level_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	SkillDepotId            uint32            `protobuf:"varint,13,opt,name=skill_depot_id,json=skillDepotId,proto3" json:"skill_depot_id,omitempty"`
	TalentIdList            []uint32          `protobuf:"varint,14,rep,packed,name=talent_id_list,json=talentIdList,proto3" json:"talent_id_list,omitempty"`
	CoreProudSkillLevel     uint32            `protobuf:"varint,15,opt,name=core_proud_skill_level,json=coreProudSkillLevel,proto3" json:"core_proud_skill_level,omitempty"`
}

func (x *AvatarSkillDepotChangeNotify) Reset() {
	*x = AvatarSkillDepotChangeNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarSkillDepotChangeNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarSkillDepotChangeNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarSkillDepotChangeNotify) ProtoMessage() {}

func (x *AvatarSkillDepotChangeNotify) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarSkillDepotChangeNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarSkillDepotChangeNotify.ProtoReflect.Descriptor instead.
func (*AvatarSkillDepotChangeNotify) Descriptor() ([]byte, []int) {
	return file_AvatarSkillDepotChangeNotify_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarSkillDepotChangeNotify) GetProudSkillExtraLevelMap() map[uint32]uint32 {
	if x != nil {
		return x.ProudSkillExtraLevelMap
	}
	return nil
}

func (x *AvatarSkillDepotChangeNotify) GetProudSkillList() []uint32 {
	if x != nil {
		return x.ProudSkillList
	}
	return nil
}

func (x *AvatarSkillDepotChangeNotify) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *AvatarSkillDepotChangeNotify) GetAvatarGuid() uint64 {
	if x != nil {
		return x.AvatarGuid
	}
	return 0
}

func (x *AvatarSkillDepotChangeNotify) GetSkillLevelMap() map[uint32]uint32 {
	if x != nil {
		return x.SkillLevelMap
	}
	return nil
}

func (x *AvatarSkillDepotChangeNotify) GetSkillDepotId() uint32 {
	if x != nil {
		return x.SkillDepotId
	}
	return 0
}

func (x *AvatarSkillDepotChangeNotify) GetTalentIdList() []uint32 {
	if x != nil {
		return x.TalentIdList
	}
	return nil
}

func (x *AvatarSkillDepotChangeNotify) GetCoreProudSkillLevel() uint32 {
	if x != nil {
		return x.CoreProudSkillLevel
	}
	return 0
}

var File_AvatarSkillDepotChangeNotify_proto protoreflect.FileDescriptor

var file_AvatarSkillDepotChangeNotify_proto_rawDesc = []byte{
	0x0a, 0x22, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x44, 0x65, 0x70,
	0x6f, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x04, 0x0a, 0x1c, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x44, 0x65, 0x70, 0x6f, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x78, 0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x75, 0x64, 0x5f, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x5f, 0x6d, 0x61, 0x70, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x44, 0x65, 0x70, 0x6f, 0x74, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x75, 0x64, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x45, 0x78, 0x74, 0x72, 0x61, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x17, 0x70, 0x72, 0x6f, 0x75, 0x64, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x45, 0x78, 0x74, 0x72, 0x61, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x12,
	0x28, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x75, 0x64, 0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x75, 0x64,
	0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x47, 0x75, 0x69, 0x64, 0x12, 0x58, 0x0a, 0x0f, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x30, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x44, 0x65,
	0x70, 0x6f, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e,
	0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0d, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61,
	0x70, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x64, 0x65, 0x70, 0x6f, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x44, 0x65, 0x70, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x61, 0x6c, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x0c, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33, 0x0a,
	0x16, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x75, 0x64, 0x5f, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x63,
	0x6f, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x75, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x1a, 0x4a, 0x0a, 0x1c, 0x50, 0x72, 0x6f, 0x75, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x40,
	0x0a, 0x12, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AvatarSkillDepotChangeNotify_proto_rawDescOnce sync.Once
	file_AvatarSkillDepotChangeNotify_proto_rawDescData = file_AvatarSkillDepotChangeNotify_proto_rawDesc
)

func file_AvatarSkillDepotChangeNotify_proto_rawDescGZIP() []byte {
	file_AvatarSkillDepotChangeNotify_proto_rawDescOnce.Do(func() {
		file_AvatarSkillDepotChangeNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarSkillDepotChangeNotify_proto_rawDescData)
	})
	return file_AvatarSkillDepotChangeNotify_proto_rawDescData
}

var file_AvatarSkillDepotChangeNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_AvatarSkillDepotChangeNotify_proto_goTypes = []interface{}{
	(*AvatarSkillDepotChangeNotify)(nil), // 0: AvatarSkillDepotChangeNotify
	nil,                                  // 1: AvatarSkillDepotChangeNotify.ProudSkillExtraLevelMapEntry
	nil,                                  // 2: AvatarSkillDepotChangeNotify.SkillLevelMapEntry
}
var file_AvatarSkillDepotChangeNotify_proto_depIdxs = []int32{
	1, // 0: AvatarSkillDepotChangeNotify.proud_skill_extra_level_map:type_name -> AvatarSkillDepotChangeNotify.ProudSkillExtraLevelMapEntry
	2, // 1: AvatarSkillDepotChangeNotify.skill_level_map:type_name -> AvatarSkillDepotChangeNotify.SkillLevelMapEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_AvatarSkillDepotChangeNotify_proto_init() }
func file_AvatarSkillDepotChangeNotify_proto_init() {
	if File_AvatarSkillDepotChangeNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AvatarSkillDepotChangeNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarSkillDepotChangeNotify); i {
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
			RawDescriptor: file_AvatarSkillDepotChangeNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarSkillDepotChangeNotify_proto_goTypes,
		DependencyIndexes: file_AvatarSkillDepotChangeNotify_proto_depIdxs,
		MessageInfos:      file_AvatarSkillDepotChangeNotify_proto_msgTypes,
	}.Build()
	File_AvatarSkillDepotChangeNotify_proto = out.File
	file_AvatarSkillDepotChangeNotify_proto_rawDesc = nil
	file_AvatarSkillDepotChangeNotify_proto_goTypes = nil
	file_AvatarSkillDepotChangeNotify_proto_depIdxs = nil
}
