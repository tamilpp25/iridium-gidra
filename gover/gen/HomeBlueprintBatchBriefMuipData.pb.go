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
// source: HomeBlueprintBatchBriefMuipData.proto

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

type HomeBlueprintBatchBriefMuipData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BriefList             []*HomeBlueprintBriefMuipData `protobuf:"bytes,1,rep,name=brief_list,json=briefList,proto3" json:"brief_list,omitempty"`
	NotExistShareCodeList []string                      `protobuf:"bytes,2,rep,name=not_exist_share_code_list,json=notExistShareCodeList,proto3" json:"not_exist_share_code_list,omitempty"`
}

func (x *HomeBlueprintBatchBriefMuipData) Reset() {
	*x = HomeBlueprintBatchBriefMuipData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HomeBlueprintBatchBriefMuipData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeBlueprintBatchBriefMuipData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeBlueprintBatchBriefMuipData) ProtoMessage() {}

func (x *HomeBlueprintBatchBriefMuipData) ProtoReflect() protoreflect.Message {
	mi := &file_HomeBlueprintBatchBriefMuipData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeBlueprintBatchBriefMuipData.ProtoReflect.Descriptor instead.
func (*HomeBlueprintBatchBriefMuipData) Descriptor() ([]byte, []int) {
	return file_HomeBlueprintBatchBriefMuipData_proto_rawDescGZIP(), []int{0}
}

func (x *HomeBlueprintBatchBriefMuipData) GetBriefList() []*HomeBlueprintBriefMuipData {
	if x != nil {
		return x.BriefList
	}
	return nil
}

func (x *HomeBlueprintBatchBriefMuipData) GetNotExistShareCodeList() []string {
	if x != nil {
		return x.NotExistShareCodeList
	}
	return nil
}

var File_HomeBlueprintBatchBriefMuipData_proto protoreflect.FileDescriptor

var file_HomeBlueprintBatchBriefMuipData_proto_rawDesc = []byte{
	0x0a, 0x25, 0x48, 0x6f, 0x6d, 0x65, 0x42, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x42, 0x72, 0x69, 0x65, 0x66, 0x4d, 0x75, 0x69, 0x70, 0x44, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x48, 0x6f, 0x6d, 0x65, 0x42, 0x6c, 0x75,
	0x65, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x42, 0x72, 0x69, 0x65, 0x66, 0x4d, 0x75, 0x69, 0x70, 0x44,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x1f, 0x48, 0x6f,
	0x6d, 0x65, 0x42, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x42, 0x72, 0x69, 0x65, 0x66, 0x4d, 0x75, 0x69, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3a, 0x0a,
	0x0a, 0x62, 0x72, 0x69, 0x65, 0x66, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x42, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x42, 0x72, 0x69, 0x65, 0x66, 0x4d, 0x75, 0x69, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09,
	0x62, 0x72, 0x69, 0x65, 0x66, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x19, 0x6e, 0x6f, 0x74,
	0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x15, 0x6e, 0x6f,
	0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_HomeBlueprintBatchBriefMuipData_proto_rawDescOnce sync.Once
	file_HomeBlueprintBatchBriefMuipData_proto_rawDescData = file_HomeBlueprintBatchBriefMuipData_proto_rawDesc
)

func file_HomeBlueprintBatchBriefMuipData_proto_rawDescGZIP() []byte {
	file_HomeBlueprintBatchBriefMuipData_proto_rawDescOnce.Do(func() {
		file_HomeBlueprintBatchBriefMuipData_proto_rawDescData = protoimpl.X.CompressGZIP(file_HomeBlueprintBatchBriefMuipData_proto_rawDescData)
	})
	return file_HomeBlueprintBatchBriefMuipData_proto_rawDescData
}

var file_HomeBlueprintBatchBriefMuipData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HomeBlueprintBatchBriefMuipData_proto_goTypes = []interface{}{
	(*HomeBlueprintBatchBriefMuipData)(nil), // 0: HomeBlueprintBatchBriefMuipData
	(*HomeBlueprintBriefMuipData)(nil),      // 1: HomeBlueprintBriefMuipData
}
var file_HomeBlueprintBatchBriefMuipData_proto_depIdxs = []int32{
	1, // 0: HomeBlueprintBatchBriefMuipData.brief_list:type_name -> HomeBlueprintBriefMuipData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_HomeBlueprintBatchBriefMuipData_proto_init() }
func file_HomeBlueprintBatchBriefMuipData_proto_init() {
	if File_HomeBlueprintBatchBriefMuipData_proto != nil {
		return
	}
	file_HomeBlueprintBriefMuipData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HomeBlueprintBatchBriefMuipData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeBlueprintBatchBriefMuipData); i {
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
			RawDescriptor: file_HomeBlueprintBatchBriefMuipData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HomeBlueprintBatchBriefMuipData_proto_goTypes,
		DependencyIndexes: file_HomeBlueprintBatchBriefMuipData_proto_depIdxs,
		MessageInfos:      file_HomeBlueprintBatchBriefMuipData_proto_msgTypes,
	}.Build()
	File_HomeBlueprintBatchBriefMuipData_proto = out.File
	file_HomeBlueprintBatchBriefMuipData_proto_rawDesc = nil
	file_HomeBlueprintBatchBriefMuipData_proto_goTypes = nil
	file_HomeBlueprintBatchBriefMuipData_proto_depIdxs = nil
}
