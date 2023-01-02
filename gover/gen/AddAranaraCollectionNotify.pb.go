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
// source: AddAranaraCollectionNotify.proto

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

type AddAranaraCollectionNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk3300_NJLJBBMJLBK AranaraCollectionState `protobuf:"varint,10,opt,name=Unk3300_NJLJBBMJLBK,json=Unk3300NJLJBBMJLBK,proto3,enum=AranaraCollectionState" json:"Unk3300_NJLJBBMJLBK,omitempty"`
	CollectionId        uint32                 `protobuf:"varint,7,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	Unk3300_AFBIBLNKCOD AranaraCollectionState `protobuf:"varint,4,opt,name=Unk3300_AFBIBLNKCOD,json=Unk3300AFBIBLNKCOD,proto3,enum=AranaraCollectionState" json:"Unk3300_AFBIBLNKCOD,omitempty"`
	CollectionType      uint32                 `protobuf:"varint,2,opt,name=collection_type,json=collectionType,proto3" json:"collection_type,omitempty"`
}

func (x *AddAranaraCollectionNotify) Reset() {
	*x = AddAranaraCollectionNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AddAranaraCollectionNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAranaraCollectionNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAranaraCollectionNotify) ProtoMessage() {}

func (x *AddAranaraCollectionNotify) ProtoReflect() protoreflect.Message {
	mi := &file_AddAranaraCollectionNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAranaraCollectionNotify.ProtoReflect.Descriptor instead.
func (*AddAranaraCollectionNotify) Descriptor() ([]byte, []int) {
	return file_AddAranaraCollectionNotify_proto_rawDescGZIP(), []int{0}
}

func (x *AddAranaraCollectionNotify) GetUnk3300_NJLJBBMJLBK() AranaraCollectionState {
	if x != nil {
		return x.Unk3300_NJLJBBMJLBK
	}
	return AranaraCollectionState_ARANARA_COLLECTION_STATE_NONE
}

func (x *AddAranaraCollectionNotify) GetCollectionId() uint32 {
	if x != nil {
		return x.CollectionId
	}
	return 0
}

func (x *AddAranaraCollectionNotify) GetUnk3300_AFBIBLNKCOD() AranaraCollectionState {
	if x != nil {
		return x.Unk3300_AFBIBLNKCOD
	}
	return AranaraCollectionState_ARANARA_COLLECTION_STATE_NONE
}

func (x *AddAranaraCollectionNotify) GetCollectionType() uint32 {
	if x != nil {
		return x.CollectionType
	}
	return 0
}

var File_AddAranaraCollectionNotify_proto protoreflect.FileDescriptor

var file_AddAranaraCollectionNotify_proto_rawDesc = []byte{
	0x0a, 0x20, 0x41, 0x64, 0x64, 0x41, 0x72, 0x61, 0x6e, 0x61, 0x72, 0x61, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x41, 0x72, 0x61, 0x6e, 0x61, 0x72, 0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xfe, 0x01, 0x0a, 0x1a, 0x41, 0x64, 0x64, 0x41, 0x72, 0x61, 0x6e, 0x61, 0x72, 0x61, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12,
	0x48, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x4e, 0x4a, 0x4c, 0x4a, 0x42,
	0x42, 0x4d, 0x4a, 0x4c, 0x42, 0x4b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x41,
	0x72, 0x61, 0x6e, 0x61, 0x72, 0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x4e, 0x4a,
	0x4c, 0x4a, 0x42, 0x42, 0x4d, 0x4a, 0x4c, 0x42, 0x4b, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x48,
	0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x41, 0x46, 0x42, 0x49, 0x42, 0x4c,
	0x4e, 0x4b, 0x43, 0x4f, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x41, 0x72,
	0x61, 0x6e, 0x61, 0x72, 0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x41, 0x46, 0x42,
	0x49, 0x42, 0x4c, 0x4e, 0x4b, 0x43, 0x4f, 0x44, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_AddAranaraCollectionNotify_proto_rawDescOnce sync.Once
	file_AddAranaraCollectionNotify_proto_rawDescData = file_AddAranaraCollectionNotify_proto_rawDesc
)

func file_AddAranaraCollectionNotify_proto_rawDescGZIP() []byte {
	file_AddAranaraCollectionNotify_proto_rawDescOnce.Do(func() {
		file_AddAranaraCollectionNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_AddAranaraCollectionNotify_proto_rawDescData)
	})
	return file_AddAranaraCollectionNotify_proto_rawDescData
}

var file_AddAranaraCollectionNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AddAranaraCollectionNotify_proto_goTypes = []interface{}{
	(*AddAranaraCollectionNotify)(nil), // 0: AddAranaraCollectionNotify
	(AranaraCollectionState)(0),        // 1: AranaraCollectionState
}
var file_AddAranaraCollectionNotify_proto_depIdxs = []int32{
	1, // 0: AddAranaraCollectionNotify.Unk3300_NJLJBBMJLBK:type_name -> AranaraCollectionState
	1, // 1: AddAranaraCollectionNotify.Unk3300_AFBIBLNKCOD:type_name -> AranaraCollectionState
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_AddAranaraCollectionNotify_proto_init() }
func file_AddAranaraCollectionNotify_proto_init() {
	if File_AddAranaraCollectionNotify_proto != nil {
		return
	}
	file_AranaraCollectionState_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AddAranaraCollectionNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAranaraCollectionNotify); i {
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
			RawDescriptor: file_AddAranaraCollectionNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AddAranaraCollectionNotify_proto_goTypes,
		DependencyIndexes: file_AddAranaraCollectionNotify_proto_depIdxs,
		MessageInfos:      file_AddAranaraCollectionNotify_proto_msgTypes,
	}.Build()
	File_AddAranaraCollectionNotify_proto = out.File
	file_AddAranaraCollectionNotify_proto_rawDesc = nil
	file_AddAranaraCollectionNotify_proto_goTypes = nil
	file_AddAranaraCollectionNotify_proto_depIdxs = nil
}
