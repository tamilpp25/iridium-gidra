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
// source: MPLevelEntityInfo.proto

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

// Obf: DBEDDNAPIOA
type MPLevelEntityInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AbilityInfo     *AbilitySyncStateInfo `protobuf:"bytes,11,opt,name=ability_info,json=abilityInfo,proto3" json:"ability_info,omitempty"`
	EntityId        uint32                `protobuf:"varint,14,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	AuthorityPeerId uint32                `protobuf:"varint,8,opt,name=authority_peer_id,json=authorityPeerId,proto3" json:"authority_peer_id,omitempty"`
}

func (x *MPLevelEntityInfo) Reset() {
	*x = MPLevelEntityInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MPLevelEntityInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MPLevelEntityInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MPLevelEntityInfo) ProtoMessage() {}

func (x *MPLevelEntityInfo) ProtoReflect() protoreflect.Message {
	mi := &file_MPLevelEntityInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MPLevelEntityInfo.ProtoReflect.Descriptor instead.
func (*MPLevelEntityInfo) Descriptor() ([]byte, []int) {
	return file_MPLevelEntityInfo_proto_rawDescGZIP(), []int{0}
}

func (x *MPLevelEntityInfo) GetAbilityInfo() *AbilitySyncStateInfo {
	if x != nil {
		return x.AbilityInfo
	}
	return nil
}

func (x *MPLevelEntityInfo) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *MPLevelEntityInfo) GetAuthorityPeerId() uint32 {
	if x != nil {
		return x.AuthorityPeerId
	}
	return 0
}

var File_MPLevelEntityInfo_proto protoreflect.FileDescriptor

var file_MPLevelEntityInfo_proto_rawDesc = []byte{
	0x0a, 0x17, 0x4d, 0x50, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x41, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x01, 0x0a, 0x11, 0x4d, 0x50, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x38, 0x0a, 0x0c, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x79, 0x6e, 0x63, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f,
	0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x50, 0x65, 0x65, 0x72, 0x49, 0x64, 0x42, 0x06,
	0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MPLevelEntityInfo_proto_rawDescOnce sync.Once
	file_MPLevelEntityInfo_proto_rawDescData = file_MPLevelEntityInfo_proto_rawDesc
)

func file_MPLevelEntityInfo_proto_rawDescGZIP() []byte {
	file_MPLevelEntityInfo_proto_rawDescOnce.Do(func() {
		file_MPLevelEntityInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_MPLevelEntityInfo_proto_rawDescData)
	})
	return file_MPLevelEntityInfo_proto_rawDescData
}

var file_MPLevelEntityInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MPLevelEntityInfo_proto_goTypes = []interface{}{
	(*MPLevelEntityInfo)(nil),    // 0: MPLevelEntityInfo
	(*AbilitySyncStateInfo)(nil), // 1: AbilitySyncStateInfo
}
var file_MPLevelEntityInfo_proto_depIdxs = []int32{
	1, // 0: MPLevelEntityInfo.ability_info:type_name -> AbilitySyncStateInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_MPLevelEntityInfo_proto_init() }
func file_MPLevelEntityInfo_proto_init() {
	if File_MPLevelEntityInfo_proto != nil {
		return
	}
	file_AbilitySyncStateInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MPLevelEntityInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MPLevelEntityInfo); i {
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
			RawDescriptor: file_MPLevelEntityInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MPLevelEntityInfo_proto_goTypes,
		DependencyIndexes: file_MPLevelEntityInfo_proto_depIdxs,
		MessageInfos:      file_MPLevelEntityInfo_proto_msgTypes,
	}.Build()
	File_MPLevelEntityInfo_proto = out.File
	file_MPLevelEntityInfo_proto_rawDesc = nil
	file_MPLevelEntityInfo_proto_goTypes = nil
	file_MPLevelEntityInfo_proto_depIdxs = nil
}
