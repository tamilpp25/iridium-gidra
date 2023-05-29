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
// source: GMShowNavMeshRsp.proto

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

// CmdId: 2309
// Obf: CKEFNCPEEKL
type GMShowNavMeshRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode int32            `protobuf:"varint,9,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Tiles   []*PBNavMeshTile `protobuf:"bytes,15,rep,name=tiles,proto3" json:"tiles,omitempty"`
}

func (x *GMShowNavMeshRsp) Reset() {
	*x = GMShowNavMeshRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GMShowNavMeshRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GMShowNavMeshRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GMShowNavMeshRsp) ProtoMessage() {}

func (x *GMShowNavMeshRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GMShowNavMeshRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GMShowNavMeshRsp.ProtoReflect.Descriptor instead.
func (*GMShowNavMeshRsp) Descriptor() ([]byte, []int) {
	return file_GMShowNavMeshRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GMShowNavMeshRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GMShowNavMeshRsp) GetTiles() []*PBNavMeshTile {
	if x != nil {
		return x.Tiles
	}
	return nil
}

var File_GMShowNavMeshRsp_proto protoreflect.FileDescriptor

var file_GMShowNavMeshRsp_proto_rawDesc = []byte{
	0x0a, 0x16, 0x47, 0x4d, 0x53, 0x68, 0x6f, 0x77, 0x4e, 0x61, 0x76, 0x4d, 0x65, 0x73, 0x68, 0x52,
	0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x50, 0x42, 0x4e, 0x61, 0x76, 0x4d,
	0x65, 0x73, 0x68, 0x54, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a,
	0x10, 0x47, 0x4d, 0x53, 0x68, 0x6f, 0x77, 0x4e, 0x61, 0x76, 0x4d, 0x65, 0x73, 0x68, 0x52, 0x73,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x74,
	0x69, 0x6c, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x50, 0x42, 0x4e,
	0x61, 0x76, 0x4d, 0x65, 0x73, 0x68, 0x54, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x74, 0x69, 0x6c, 0x65,
	0x73, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_GMShowNavMeshRsp_proto_rawDescOnce sync.Once
	file_GMShowNavMeshRsp_proto_rawDescData = file_GMShowNavMeshRsp_proto_rawDesc
)

func file_GMShowNavMeshRsp_proto_rawDescGZIP() []byte {
	file_GMShowNavMeshRsp_proto_rawDescOnce.Do(func() {
		file_GMShowNavMeshRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GMShowNavMeshRsp_proto_rawDescData)
	})
	return file_GMShowNavMeshRsp_proto_rawDescData
}

var file_GMShowNavMeshRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GMShowNavMeshRsp_proto_goTypes = []interface{}{
	(*GMShowNavMeshRsp)(nil), // 0: GMShowNavMeshRsp
	(*PBNavMeshTile)(nil),    // 1: PBNavMeshTile
}
var file_GMShowNavMeshRsp_proto_depIdxs = []int32{
	1, // 0: GMShowNavMeshRsp.tiles:type_name -> PBNavMeshTile
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GMShowNavMeshRsp_proto_init() }
func file_GMShowNavMeshRsp_proto_init() {
	if File_GMShowNavMeshRsp_proto != nil {
		return
	}
	file_PBNavMeshTile_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GMShowNavMeshRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GMShowNavMeshRsp); i {
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
			RawDescriptor: file_GMShowNavMeshRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GMShowNavMeshRsp_proto_goTypes,
		DependencyIndexes: file_GMShowNavMeshRsp_proto_depIdxs,
		MessageInfos:      file_GMShowNavMeshRsp_proto_msgTypes,
	}.Build()
	File_GMShowNavMeshRsp_proto = out.File
	file_GMShowNavMeshRsp_proto_rawDesc = nil
	file_GMShowNavMeshRsp_proto_goTypes = nil
	file_GMShowNavMeshRsp_proto_depIdxs = nil
}
