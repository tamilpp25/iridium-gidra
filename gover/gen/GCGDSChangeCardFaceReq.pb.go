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
// source: GCGDSChangeCardFaceReq.proto

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

// CmdId: 7536
// Obf: CLOJNHDBJAP
type GCGDSChangeCardFaceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardId   uint32 `protobuf:"varint,4,opt,name=card_id,json=cardId,proto3" json:"card_id,omitempty"`
	FaceType uint32 `protobuf:"varint,1,opt,name=face_type,json=faceType,proto3" json:"face_type,omitempty"`
}

func (x *GCGDSChangeCardFaceReq) Reset() {
	*x = GCGDSChangeCardFaceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GCGDSChangeCardFaceReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GCGDSChangeCardFaceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GCGDSChangeCardFaceReq) ProtoMessage() {}

func (x *GCGDSChangeCardFaceReq) ProtoReflect() protoreflect.Message {
	mi := &file_GCGDSChangeCardFaceReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GCGDSChangeCardFaceReq.ProtoReflect.Descriptor instead.
func (*GCGDSChangeCardFaceReq) Descriptor() ([]byte, []int) {
	return file_GCGDSChangeCardFaceReq_proto_rawDescGZIP(), []int{0}
}

func (x *GCGDSChangeCardFaceReq) GetCardId() uint32 {
	if x != nil {
		return x.CardId
	}
	return 0
}

func (x *GCGDSChangeCardFaceReq) GetFaceType() uint32 {
	if x != nil {
		return x.FaceType
	}
	return 0
}

var File_GCGDSChangeCardFaceReq_proto protoreflect.FileDescriptor

var file_GCGDSChangeCardFaceReq_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x47, 0x43, 0x47, 0x44, 0x53, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x61, 0x72,
	0x64, 0x46, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e,
	0x0a, 0x16, 0x47, 0x43, 0x47, 0x44, 0x53, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x61, 0x72,
	0x64, 0x46, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x61, 0x72, 0x64, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x66, 0x61, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06,
	0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GCGDSChangeCardFaceReq_proto_rawDescOnce sync.Once
	file_GCGDSChangeCardFaceReq_proto_rawDescData = file_GCGDSChangeCardFaceReq_proto_rawDesc
)

func file_GCGDSChangeCardFaceReq_proto_rawDescGZIP() []byte {
	file_GCGDSChangeCardFaceReq_proto_rawDescOnce.Do(func() {
		file_GCGDSChangeCardFaceReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_GCGDSChangeCardFaceReq_proto_rawDescData)
	})
	return file_GCGDSChangeCardFaceReq_proto_rawDescData
}

var file_GCGDSChangeCardFaceReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GCGDSChangeCardFaceReq_proto_goTypes = []interface{}{
	(*GCGDSChangeCardFaceReq)(nil), // 0: GCGDSChangeCardFaceReq
}
var file_GCGDSChangeCardFaceReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GCGDSChangeCardFaceReq_proto_init() }
func file_GCGDSChangeCardFaceReq_proto_init() {
	if File_GCGDSChangeCardFaceReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GCGDSChangeCardFaceReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GCGDSChangeCardFaceReq); i {
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
			RawDescriptor: file_GCGDSChangeCardFaceReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GCGDSChangeCardFaceReq_proto_goTypes,
		DependencyIndexes: file_GCGDSChangeCardFaceReq_proto_depIdxs,
		MessageInfos:      file_GCGDSChangeCardFaceReq_proto_msgTypes,
	}.Build()
	File_GCGDSChangeCardFaceReq_proto = out.File
	file_GCGDSChangeCardFaceReq_proto_rawDesc = nil
	file_GCGDSChangeCardFaceReq_proto_goTypes = nil
	file_GCGDSChangeCardFaceReq_proto_depIdxs = nil
}
