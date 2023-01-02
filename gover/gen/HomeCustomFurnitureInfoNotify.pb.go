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
// source: HomeCustomFurnitureInfoNotify.proto

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

type HomeCustomFurnitureInfoNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeleteCustomFurnitureList []uint32                   `protobuf:"varint,13,rep,packed,name=delete_custom_furniture_list,json=deleteCustomFurnitureList,proto3" json:"delete_custom_furniture_list,omitempty"`
	CustomFurnitureInfoList   []*HomeCustomFurnitureInfo `protobuf:"bytes,15,rep,name=custom_furniture_info_list,json=customFurnitureInfoList,proto3" json:"custom_furniture_info_list,omitempty"`
	UsedSubFurnitureCountMap  map[uint32]uint32          `protobuf:"bytes,3,rep,name=used_sub_furniture_count_map,json=usedSubFurnitureCountMap,proto3" json:"used_sub_furniture_count_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *HomeCustomFurnitureInfoNotify) Reset() {
	*x = HomeCustomFurnitureInfoNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HomeCustomFurnitureInfoNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeCustomFurnitureInfoNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeCustomFurnitureInfoNotify) ProtoMessage() {}

func (x *HomeCustomFurnitureInfoNotify) ProtoReflect() protoreflect.Message {
	mi := &file_HomeCustomFurnitureInfoNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeCustomFurnitureInfoNotify.ProtoReflect.Descriptor instead.
func (*HomeCustomFurnitureInfoNotify) Descriptor() ([]byte, []int) {
	return file_HomeCustomFurnitureInfoNotify_proto_rawDescGZIP(), []int{0}
}

func (x *HomeCustomFurnitureInfoNotify) GetDeleteCustomFurnitureList() []uint32 {
	if x != nil {
		return x.DeleteCustomFurnitureList
	}
	return nil
}

func (x *HomeCustomFurnitureInfoNotify) GetCustomFurnitureInfoList() []*HomeCustomFurnitureInfo {
	if x != nil {
		return x.CustomFurnitureInfoList
	}
	return nil
}

func (x *HomeCustomFurnitureInfoNotify) GetUsedSubFurnitureCountMap() map[uint32]uint32 {
	if x != nil {
		return x.UsedSubFurnitureCountMap
	}
	return nil
}

var File_HomeCustomFurnitureInfoNotify_proto protoreflect.FileDescriptor

var file_HomeCustomFurnitureInfoNotify_proto_rawDesc = []byte{
	0x0a, 0x23, 0x48, 0x6f, 0x6d, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x46, 0x75, 0x72, 0x6e,
	0x69, 0x74, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x48, 0x6f, 0x6d, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x03, 0x0a, 0x1d, 0x48, 0x6f, 0x6d, 0x65, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x3f, 0x0a, 0x1c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x66, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72,
	0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x19, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74,
	0x75, 0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x55, 0x0a, 0x1a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x5f, 0x66, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x48, 0x6f,
	0x6d, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x17, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x46, 0x75, 0x72,
	0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x7c,
	0x0a, 0x1c, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x73, 0x75, 0x62, 0x5f, 0x66, 0x75, 0x72, 0x6e, 0x69,
	0x74, 0x75, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x64, 0x53, 0x75, 0x62, 0x46, 0x75, 0x72, 0x6e,
	0x69, 0x74, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x18, 0x75, 0x73, 0x65, 0x64, 0x53, 0x75, 0x62, 0x46, 0x75, 0x72, 0x6e, 0x69,
	0x74, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x1a, 0x4b, 0x0a, 0x1d,
	0x55, 0x73, 0x65, 0x64, 0x53, 0x75, 0x62, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HomeCustomFurnitureInfoNotify_proto_rawDescOnce sync.Once
	file_HomeCustomFurnitureInfoNotify_proto_rawDescData = file_HomeCustomFurnitureInfoNotify_proto_rawDesc
)

func file_HomeCustomFurnitureInfoNotify_proto_rawDescGZIP() []byte {
	file_HomeCustomFurnitureInfoNotify_proto_rawDescOnce.Do(func() {
		file_HomeCustomFurnitureInfoNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_HomeCustomFurnitureInfoNotify_proto_rawDescData)
	})
	return file_HomeCustomFurnitureInfoNotify_proto_rawDescData
}

var file_HomeCustomFurnitureInfoNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_HomeCustomFurnitureInfoNotify_proto_goTypes = []interface{}{
	(*HomeCustomFurnitureInfoNotify)(nil), // 0: HomeCustomFurnitureInfoNotify
	nil,                                   // 1: HomeCustomFurnitureInfoNotify.UsedSubFurnitureCountMapEntry
	(*HomeCustomFurnitureInfo)(nil),       // 2: HomeCustomFurnitureInfo
}
var file_HomeCustomFurnitureInfoNotify_proto_depIdxs = []int32{
	2, // 0: HomeCustomFurnitureInfoNotify.custom_furniture_info_list:type_name -> HomeCustomFurnitureInfo
	1, // 1: HomeCustomFurnitureInfoNotify.used_sub_furniture_count_map:type_name -> HomeCustomFurnitureInfoNotify.UsedSubFurnitureCountMapEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_HomeCustomFurnitureInfoNotify_proto_init() }
func file_HomeCustomFurnitureInfoNotify_proto_init() {
	if File_HomeCustomFurnitureInfoNotify_proto != nil {
		return
	}
	file_HomeCustomFurnitureInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HomeCustomFurnitureInfoNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeCustomFurnitureInfoNotify); i {
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
			RawDescriptor: file_HomeCustomFurnitureInfoNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HomeCustomFurnitureInfoNotify_proto_goTypes,
		DependencyIndexes: file_HomeCustomFurnitureInfoNotify_proto_depIdxs,
		MessageInfos:      file_HomeCustomFurnitureInfoNotify_proto_msgTypes,
	}.Build()
	File_HomeCustomFurnitureInfoNotify_proto = out.File
	file_HomeCustomFurnitureInfoNotify_proto_rawDesc = nil
	file_HomeCustomFurnitureInfoNotify_proto_goTypes = nil
	file_HomeCustomFurnitureInfoNotify_proto_depIdxs = nil
}
