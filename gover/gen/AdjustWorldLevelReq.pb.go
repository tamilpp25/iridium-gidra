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
// source: AdjustWorldLevelReq.proto

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

type AdjustWorldLevelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk3300_DNLKIMGLBLP uint32 `protobuf:"varint,14,opt,name=Unk3300_DNLKIMGLBLP,json=Unk3300DNLKIMGLBLP,proto3" json:"Unk3300_DNLKIMGLBLP,omitempty"`
	Unk3300_DJBKBPDCBFH uint32 `protobuf:"varint,3,opt,name=Unk3300_DJBKBPDCBFH,json=Unk3300DJBKBPDCBFH,proto3" json:"Unk3300_DJBKBPDCBFH,omitempty"`
}

func (x *AdjustWorldLevelReq) Reset() {
	*x = AdjustWorldLevelReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AdjustWorldLevelReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdjustWorldLevelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdjustWorldLevelReq) ProtoMessage() {}

func (x *AdjustWorldLevelReq) ProtoReflect() protoreflect.Message {
	mi := &file_AdjustWorldLevelReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdjustWorldLevelReq.ProtoReflect.Descriptor instead.
func (*AdjustWorldLevelReq) Descriptor() ([]byte, []int) {
	return file_AdjustWorldLevelReq_proto_rawDescGZIP(), []int{0}
}

func (x *AdjustWorldLevelReq) GetUnk3300_DNLKIMGLBLP() uint32 {
	if x != nil {
		return x.Unk3300_DNLKIMGLBLP
	}
	return 0
}

func (x *AdjustWorldLevelReq) GetUnk3300_DJBKBPDCBFH() uint32 {
	if x != nil {
		return x.Unk3300_DJBKBPDCBFH
	}
	return 0
}

var File_AdjustWorldLevelReq_proto protoreflect.FileDescriptor

var file_AdjustWorldLevelReq_proto_rawDesc = []byte{
	0x0a, 0x19, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x77, 0x0a, 0x13, 0x41,
	0x64, 0x6a, 0x75, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52,
	0x65, 0x71, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x44, 0x4e,
	0x4c, 0x4b, 0x49, 0x4d, 0x47, 0x4c, 0x42, 0x4c, 0x50, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x44, 0x4e, 0x4c, 0x4b, 0x49, 0x4d, 0x47, 0x4c,
	0x42, 0x4c, 0x50, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x44,
	0x4a, 0x42, 0x4b, 0x42, 0x50, 0x44, 0x43, 0x42, 0x46, 0x48, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x44, 0x4a, 0x42, 0x4b, 0x42, 0x50, 0x44,
	0x43, 0x42, 0x46, 0x48, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AdjustWorldLevelReq_proto_rawDescOnce sync.Once
	file_AdjustWorldLevelReq_proto_rawDescData = file_AdjustWorldLevelReq_proto_rawDesc
)

func file_AdjustWorldLevelReq_proto_rawDescGZIP() []byte {
	file_AdjustWorldLevelReq_proto_rawDescOnce.Do(func() {
		file_AdjustWorldLevelReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_AdjustWorldLevelReq_proto_rawDescData)
	})
	return file_AdjustWorldLevelReq_proto_rawDescData
}

var file_AdjustWorldLevelReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AdjustWorldLevelReq_proto_goTypes = []interface{}{
	(*AdjustWorldLevelReq)(nil), // 0: AdjustWorldLevelReq
}
var file_AdjustWorldLevelReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AdjustWorldLevelReq_proto_init() }
func file_AdjustWorldLevelReq_proto_init() {
	if File_AdjustWorldLevelReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AdjustWorldLevelReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdjustWorldLevelReq); i {
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
			RawDescriptor: file_AdjustWorldLevelReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AdjustWorldLevelReq_proto_goTypes,
		DependencyIndexes: file_AdjustWorldLevelReq_proto_depIdxs,
		MessageInfos:      file_AdjustWorldLevelReq_proto_msgTypes,
	}.Build()
	File_AdjustWorldLevelReq_proto = out.File
	file_AdjustWorldLevelReq_proto_rawDesc = nil
	file_AdjustWorldLevelReq_proto_goTypes = nil
	file_AdjustWorldLevelReq_proto_depIdxs = nil
}
