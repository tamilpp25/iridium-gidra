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
// source: IrodoriMasterGallerySettleNotify.proto

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

type IrodoriMasterGallerySettleNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SettleInfo *IrodoriMasterGallerySettleInfo `protobuf:"bytes,9,opt,name=settle_info,json=settleInfo,proto3" json:"settle_info,omitempty"`
	GalleryId  uint32                          `protobuf:"varint,11,opt,name=gallery_id,json=galleryId,proto3" json:"gallery_id,omitempty"`
}

func (x *IrodoriMasterGallerySettleNotify) Reset() {
	*x = IrodoriMasterGallerySettleNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IrodoriMasterGallerySettleNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IrodoriMasterGallerySettleNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IrodoriMasterGallerySettleNotify) ProtoMessage() {}

func (x *IrodoriMasterGallerySettleNotify) ProtoReflect() protoreflect.Message {
	mi := &file_IrodoriMasterGallerySettleNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IrodoriMasterGallerySettleNotify.ProtoReflect.Descriptor instead.
func (*IrodoriMasterGallerySettleNotify) Descriptor() ([]byte, []int) {
	return file_IrodoriMasterGallerySettleNotify_proto_rawDescGZIP(), []int{0}
}

func (x *IrodoriMasterGallerySettleNotify) GetSettleInfo() *IrodoriMasterGallerySettleInfo {
	if x != nil {
		return x.SettleInfo
	}
	return nil
}

func (x *IrodoriMasterGallerySettleNotify) GetGalleryId() uint32 {
	if x != nil {
		return x.GalleryId
	}
	return 0
}

var File_IrodoriMasterGallerySettleNotify_proto protoreflect.FileDescriptor

var file_IrodoriMasterGallerySettleNotify_proto_rawDesc = []byte{
	0x0a, 0x26, 0x49, 0x72, 0x6f, 0x64, 0x6f, 0x72, 0x69, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x47,
	0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x49, 0x72, 0x6f, 0x64, 0x6f, 0x72,
	0x69, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x65,
	0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83,
	0x01, 0x0a, 0x20, 0x49, 0x72, 0x6f, 0x64, 0x6f, 0x72, 0x69, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x12, 0x40, 0x0a, 0x0b, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x49, 0x72, 0x6f, 0x64, 0x6f,
	0x72, 0x69, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53,
	0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x73, 0x65, 0x74, 0x74, 0x6c,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x67, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x79, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_IrodoriMasterGallerySettleNotify_proto_rawDescOnce sync.Once
	file_IrodoriMasterGallerySettleNotify_proto_rawDescData = file_IrodoriMasterGallerySettleNotify_proto_rawDesc
)

func file_IrodoriMasterGallerySettleNotify_proto_rawDescGZIP() []byte {
	file_IrodoriMasterGallerySettleNotify_proto_rawDescOnce.Do(func() {
		file_IrodoriMasterGallerySettleNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_IrodoriMasterGallerySettleNotify_proto_rawDescData)
	})
	return file_IrodoriMasterGallerySettleNotify_proto_rawDescData
}

var file_IrodoriMasterGallerySettleNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_IrodoriMasterGallerySettleNotify_proto_goTypes = []interface{}{
	(*IrodoriMasterGallerySettleNotify)(nil), // 0: IrodoriMasterGallerySettleNotify
	(*IrodoriMasterGallerySettleInfo)(nil),   // 1: IrodoriMasterGallerySettleInfo
}
var file_IrodoriMasterGallerySettleNotify_proto_depIdxs = []int32{
	1, // 0: IrodoriMasterGallerySettleNotify.settle_info:type_name -> IrodoriMasterGallerySettleInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_IrodoriMasterGallerySettleNotify_proto_init() }
func file_IrodoriMasterGallerySettleNotify_proto_init() {
	if File_IrodoriMasterGallerySettleNotify_proto != nil {
		return
	}
	file_IrodoriMasterGallerySettleInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_IrodoriMasterGallerySettleNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IrodoriMasterGallerySettleNotify); i {
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
			RawDescriptor: file_IrodoriMasterGallerySettleNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_IrodoriMasterGallerySettleNotify_proto_goTypes,
		DependencyIndexes: file_IrodoriMasterGallerySettleNotify_proto_depIdxs,
		MessageInfos:      file_IrodoriMasterGallerySettleNotify_proto_msgTypes,
	}.Build()
	File_IrodoriMasterGallerySettleNotify_proto = out.File
	file_IrodoriMasterGallerySettleNotify_proto_rawDesc = nil
	file_IrodoriMasterGallerySettleNotify_proto_goTypes = nil
	file_IrodoriMasterGallerySettleNotify_proto_depIdxs = nil
}
