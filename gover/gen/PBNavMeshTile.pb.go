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
// source: PBNavMeshTile.proto

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

type PBNavMeshTile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vecs  []*Vector        `protobuf:"bytes,12,rep,name=vecs,proto3" json:"vecs,omitempty"`
	Polys []*PBNavMeshPoly `protobuf:"bytes,13,rep,name=polys,proto3" json:"polys,omitempty"`
}

func (x *PBNavMeshTile) Reset() {
	*x = PBNavMeshTile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PBNavMeshTile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBNavMeshTile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBNavMeshTile) ProtoMessage() {}

func (x *PBNavMeshTile) ProtoReflect() protoreflect.Message {
	mi := &file_PBNavMeshTile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBNavMeshTile.ProtoReflect.Descriptor instead.
func (*PBNavMeshTile) Descriptor() ([]byte, []int) {
	return file_PBNavMeshTile_proto_rawDescGZIP(), []int{0}
}

func (x *PBNavMeshTile) GetVecs() []*Vector {
	if x != nil {
		return x.Vecs
	}
	return nil
}

func (x *PBNavMeshTile) GetPolys() []*PBNavMeshPoly {
	if x != nil {
		return x.Polys
	}
	return nil
}

var File_PBNavMeshTile_proto protoreflect.FileDescriptor

var file_PBNavMeshTile_proto_rawDesc = []byte{
	0x0a, 0x13, 0x50, 0x42, 0x4e, 0x61, 0x76, 0x4d, 0x65, 0x73, 0x68, 0x54, 0x69, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x50, 0x42, 0x4e, 0x61, 0x76, 0x4d, 0x65, 0x73, 0x68,
	0x50, 0x6f, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x0d, 0x50, 0x42, 0x4e, 0x61,
	0x76, 0x4d, 0x65, 0x73, 0x68, 0x54, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x76, 0x65, 0x63,
	0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x04, 0x76, 0x65, 0x63, 0x73, 0x12, 0x24, 0x0a, 0x05, 0x70, 0x6f, 0x6c, 0x79, 0x73, 0x18,
	0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x50, 0x42, 0x4e, 0x61, 0x76, 0x4d, 0x65, 0x73,
	0x68, 0x50, 0x6f, 0x6c, 0x79, 0x52, 0x05, 0x70, 0x6f, 0x6c, 0x79, 0x73, 0x42, 0x06, 0x5a, 0x04,
	0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PBNavMeshTile_proto_rawDescOnce sync.Once
	file_PBNavMeshTile_proto_rawDescData = file_PBNavMeshTile_proto_rawDesc
)

func file_PBNavMeshTile_proto_rawDescGZIP() []byte {
	file_PBNavMeshTile_proto_rawDescOnce.Do(func() {
		file_PBNavMeshTile_proto_rawDescData = protoimpl.X.CompressGZIP(file_PBNavMeshTile_proto_rawDescData)
	})
	return file_PBNavMeshTile_proto_rawDescData
}

var file_PBNavMeshTile_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PBNavMeshTile_proto_goTypes = []interface{}{
	(*PBNavMeshTile)(nil), // 0: PBNavMeshTile
	(*Vector)(nil),        // 1: Vector
	(*PBNavMeshPoly)(nil), // 2: PBNavMeshPoly
}
var file_PBNavMeshTile_proto_depIdxs = []int32{
	1, // 0: PBNavMeshTile.vecs:type_name -> Vector
	2, // 1: PBNavMeshTile.polys:type_name -> PBNavMeshPoly
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_PBNavMeshTile_proto_init() }
func file_PBNavMeshTile_proto_init() {
	if File_PBNavMeshTile_proto != nil {
		return
	}
	file_PBNavMeshPoly_proto_init()
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PBNavMeshTile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBNavMeshTile); i {
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
			RawDescriptor: file_PBNavMeshTile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PBNavMeshTile_proto_goTypes,
		DependencyIndexes: file_PBNavMeshTile_proto_depIdxs,
		MessageInfos:      file_PBNavMeshTile_proto_msgTypes,
	}.Build()
	File_PBNavMeshTile_proto = out.File
	file_PBNavMeshTile_proto_rawDesc = nil
	file_PBNavMeshTile_proto_goTypes = nil
	file_PBNavMeshTile_proto_depIdxs = nil
}
