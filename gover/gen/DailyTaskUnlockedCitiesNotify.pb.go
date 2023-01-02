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
// source: DailyTaskUnlockedCitiesNotify.proto

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

type DailyTaskUnlockedCitiesNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnlockedCityList []uint32 `protobuf:"varint,15,rep,packed,name=unlocked_city_list,json=unlockedCityList,proto3" json:"unlocked_city_list,omitempty"`
}

func (x *DailyTaskUnlockedCitiesNotify) Reset() {
	*x = DailyTaskUnlockedCitiesNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DailyTaskUnlockedCitiesNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DailyTaskUnlockedCitiesNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DailyTaskUnlockedCitiesNotify) ProtoMessage() {}

func (x *DailyTaskUnlockedCitiesNotify) ProtoReflect() protoreflect.Message {
	mi := &file_DailyTaskUnlockedCitiesNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DailyTaskUnlockedCitiesNotify.ProtoReflect.Descriptor instead.
func (*DailyTaskUnlockedCitiesNotify) Descriptor() ([]byte, []int) {
	return file_DailyTaskUnlockedCitiesNotify_proto_rawDescGZIP(), []int{0}
}

func (x *DailyTaskUnlockedCitiesNotify) GetUnlockedCityList() []uint32 {
	if x != nil {
		return x.UnlockedCityList
	}
	return nil
}

var File_DailyTaskUnlockedCitiesNotify_proto protoreflect.FileDescriptor

var file_DailyTaskUnlockedCitiesNotify_proto_rawDesc = []byte{
	0x0a, 0x23, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x55, 0x6e, 0x6c, 0x6f, 0x63,
	0x6b, 0x65, 0x64, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x1d, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x54, 0x61,
	0x73, 0x6b, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x2c, 0x0a, 0x12, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b,
	0x65, 0x64, 0x5f, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x10, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x43, 0x69, 0x74, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DailyTaskUnlockedCitiesNotify_proto_rawDescOnce sync.Once
	file_DailyTaskUnlockedCitiesNotify_proto_rawDescData = file_DailyTaskUnlockedCitiesNotify_proto_rawDesc
)

func file_DailyTaskUnlockedCitiesNotify_proto_rawDescGZIP() []byte {
	file_DailyTaskUnlockedCitiesNotify_proto_rawDescOnce.Do(func() {
		file_DailyTaskUnlockedCitiesNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_DailyTaskUnlockedCitiesNotify_proto_rawDescData)
	})
	return file_DailyTaskUnlockedCitiesNotify_proto_rawDescData
}

var file_DailyTaskUnlockedCitiesNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DailyTaskUnlockedCitiesNotify_proto_goTypes = []interface{}{
	(*DailyTaskUnlockedCitiesNotify)(nil), // 0: DailyTaskUnlockedCitiesNotify
}
var file_DailyTaskUnlockedCitiesNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_DailyTaskUnlockedCitiesNotify_proto_init() }
func file_DailyTaskUnlockedCitiesNotify_proto_init() {
	if File_DailyTaskUnlockedCitiesNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_DailyTaskUnlockedCitiesNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DailyTaskUnlockedCitiesNotify); i {
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
			RawDescriptor: file_DailyTaskUnlockedCitiesNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DailyTaskUnlockedCitiesNotify_proto_goTypes,
		DependencyIndexes: file_DailyTaskUnlockedCitiesNotify_proto_depIdxs,
		MessageInfos:      file_DailyTaskUnlockedCitiesNotify_proto_msgTypes,
	}.Build()
	File_DailyTaskUnlockedCitiesNotify_proto = out.File
	file_DailyTaskUnlockedCitiesNotify_proto_rawDesc = nil
	file_DailyTaskUnlockedCitiesNotify_proto_goTypes = nil
	file_DailyTaskUnlockedCitiesNotify_proto_depIdxs = nil
}
