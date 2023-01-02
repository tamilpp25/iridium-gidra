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
// source: GearActivityFinishPlayGearReq.proto

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

type GearActivityFinishPlayGearReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UseClue            bool              `protobuf:"varint,10,opt,name=use_clue,json=useClue,proto3" json:"use_clue,omitempty"`
	LevelId            uint32            `protobuf:"varint,3,opt,name=level_id,json=levelId,proto3" json:"level_id,omitempty"`
	IsSuccess          bool              `protobuf:"varint,8,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	GearColumnInfoList []*GearColumnInfo `protobuf:"bytes,12,rep,name=gear_column_info_list,json=gearColumnInfoList,proto3" json:"gear_column_info_list,omitempty"`
}

func (x *GearActivityFinishPlayGearReq) Reset() {
	*x = GearActivityFinishPlayGearReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GearActivityFinishPlayGearReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GearActivityFinishPlayGearReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GearActivityFinishPlayGearReq) ProtoMessage() {}

func (x *GearActivityFinishPlayGearReq) ProtoReflect() protoreflect.Message {
	mi := &file_GearActivityFinishPlayGearReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GearActivityFinishPlayGearReq.ProtoReflect.Descriptor instead.
func (*GearActivityFinishPlayGearReq) Descriptor() ([]byte, []int) {
	return file_GearActivityFinishPlayGearReq_proto_rawDescGZIP(), []int{0}
}

func (x *GearActivityFinishPlayGearReq) GetUseClue() bool {
	if x != nil {
		return x.UseClue
	}
	return false
}

func (x *GearActivityFinishPlayGearReq) GetLevelId() uint32 {
	if x != nil {
		return x.LevelId
	}
	return 0
}

func (x *GearActivityFinishPlayGearReq) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *GearActivityFinishPlayGearReq) GetGearColumnInfoList() []*GearColumnInfo {
	if x != nil {
		return x.GearColumnInfoList
	}
	return nil
}

var File_GearActivityFinishPlayGearReq_proto protoreflect.FileDescriptor

var file_GearActivityFinishPlayGearReq_proto_rawDesc = []byte{
	0x0a, 0x23, 0x47, 0x65, 0x61, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x46, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x50, 0x6c, 0x61, 0x79, 0x47, 0x65, 0x61, 0x72, 0x52, 0x65, 0x71, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x47, 0x65, 0x61, 0x72, 0x43, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x1d,
	0x47, 0x65, 0x61, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x46, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x50, 0x6c, 0x61, 0x79, 0x47, 0x65, 0x61, 0x72, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x6c, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x75, 0x73, 0x65, 0x43, 0x6c, 0x75, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x42, 0x0a, 0x15, 0x67, 0x65, 0x61, 0x72, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x47, 0x65, 0x61, 0x72, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x12, 0x67, 0x65, 0x61, 0x72, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GearActivityFinishPlayGearReq_proto_rawDescOnce sync.Once
	file_GearActivityFinishPlayGearReq_proto_rawDescData = file_GearActivityFinishPlayGearReq_proto_rawDesc
)

func file_GearActivityFinishPlayGearReq_proto_rawDescGZIP() []byte {
	file_GearActivityFinishPlayGearReq_proto_rawDescOnce.Do(func() {
		file_GearActivityFinishPlayGearReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_GearActivityFinishPlayGearReq_proto_rawDescData)
	})
	return file_GearActivityFinishPlayGearReq_proto_rawDescData
}

var file_GearActivityFinishPlayGearReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GearActivityFinishPlayGearReq_proto_goTypes = []interface{}{
	(*GearActivityFinishPlayGearReq)(nil), // 0: GearActivityFinishPlayGearReq
	(*GearColumnInfo)(nil),                // 1: GearColumnInfo
}
var file_GearActivityFinishPlayGearReq_proto_depIdxs = []int32{
	1, // 0: GearActivityFinishPlayGearReq.gear_column_info_list:type_name -> GearColumnInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GearActivityFinishPlayGearReq_proto_init() }
func file_GearActivityFinishPlayGearReq_proto_init() {
	if File_GearActivityFinishPlayGearReq_proto != nil {
		return
	}
	file_GearColumnInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GearActivityFinishPlayGearReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GearActivityFinishPlayGearReq); i {
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
			RawDescriptor: file_GearActivityFinishPlayGearReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GearActivityFinishPlayGearReq_proto_goTypes,
		DependencyIndexes: file_GearActivityFinishPlayGearReq_proto_depIdxs,
		MessageInfos:      file_GearActivityFinishPlayGearReq_proto_msgTypes,
	}.Build()
	File_GearActivityFinishPlayGearReq_proto = out.File
	file_GearActivityFinishPlayGearReq_proto_rawDesc = nil
	file_GearActivityFinishPlayGearReq_proto_goTypes = nil
	file_GearActivityFinishPlayGearReq_proto_depIdxs = nil
}
