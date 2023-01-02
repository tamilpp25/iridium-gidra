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
// source: PlantFlowerAcceptFlowerResultInfo.proto

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

type PlantFlowerAcceptFlowerResultInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid                 uint32            `protobuf:"varint,9,opt,name=uid,proto3" json:"uid,omitempty"`
	Unk3300_PALEJEDIAJP map[uint32]uint32 `protobuf:"bytes,6,rep,name=Unk3300_PALEJEDIAJP,json=Unk3300PALEJEDIAJP,proto3" json:"Unk3300_PALEJEDIAJP,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Unk3300_ILAADDCOPKE map[uint32]uint32 `protobuf:"bytes,7,rep,name=Unk3300_ILAADDCOPKE,json=Unk3300ILAADDCOPKE,proto3" json:"Unk3300_ILAADDCOPKE,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *PlantFlowerAcceptFlowerResultInfo) Reset() {
	*x = PlantFlowerAcceptFlowerResultInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlantFlowerAcceptFlowerResultInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlantFlowerAcceptFlowerResultInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlantFlowerAcceptFlowerResultInfo) ProtoMessage() {}

func (x *PlantFlowerAcceptFlowerResultInfo) ProtoReflect() protoreflect.Message {
	mi := &file_PlantFlowerAcceptFlowerResultInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlantFlowerAcceptFlowerResultInfo.ProtoReflect.Descriptor instead.
func (*PlantFlowerAcceptFlowerResultInfo) Descriptor() ([]byte, []int) {
	return file_PlantFlowerAcceptFlowerResultInfo_proto_rawDescGZIP(), []int{0}
}

func (x *PlantFlowerAcceptFlowerResultInfo) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *PlantFlowerAcceptFlowerResultInfo) GetUnk3300_PALEJEDIAJP() map[uint32]uint32 {
	if x != nil {
		return x.Unk3300_PALEJEDIAJP
	}
	return nil
}

func (x *PlantFlowerAcceptFlowerResultInfo) GetUnk3300_ILAADDCOPKE() map[uint32]uint32 {
	if x != nil {
		return x.Unk3300_ILAADDCOPKE
	}
	return nil
}

var File_PlantFlowerAcceptFlowerResultInfo_proto protoreflect.FileDescriptor

var file_PlantFlowerAcceptFlowerResultInfo_proto_rawDesc = []byte{
	0x0a, 0x27, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x41, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x03, 0x0a, 0x21, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x46,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x12, 0x6b, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x50, 0x41, 0x4c,
	0x45, 0x4a, 0x45, 0x44, 0x49, 0x41, 0x4a, 0x50, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a,
	0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x50, 0x41, 0x4c, 0x45, 0x4a, 0x45,
	0x44, 0x49, 0x41, 0x4a, 0x50, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33,
	0x33, 0x30, 0x30, 0x50, 0x41, 0x4c, 0x45, 0x4a, 0x45, 0x44, 0x49, 0x41, 0x4a, 0x50, 0x12, 0x6b,
	0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x49, 0x4c, 0x41, 0x41, 0x44, 0x44,
	0x43, 0x4f, 0x50, 0x4b, 0x45, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x46,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x49, 0x4c, 0x41, 0x41, 0x44, 0x44, 0x43, 0x4f, 0x50,
	0x4b, 0x45, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30,
	0x49, 0x4c, 0x41, 0x41, 0x44, 0x44, 0x43, 0x4f, 0x50, 0x4b, 0x45, 0x1a, 0x45, 0x0a, 0x17, 0x55,
	0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x50, 0x41, 0x4c, 0x45, 0x4a, 0x45, 0x44, 0x49, 0x41, 0x4a,
	0x50, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x45, 0x0a, 0x17, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x49, 0x4c, 0x41,
	0x41, 0x44, 0x44, 0x43, 0x4f, 0x50, 0x4b, 0x45, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlantFlowerAcceptFlowerResultInfo_proto_rawDescOnce sync.Once
	file_PlantFlowerAcceptFlowerResultInfo_proto_rawDescData = file_PlantFlowerAcceptFlowerResultInfo_proto_rawDesc
)

func file_PlantFlowerAcceptFlowerResultInfo_proto_rawDescGZIP() []byte {
	file_PlantFlowerAcceptFlowerResultInfo_proto_rawDescOnce.Do(func() {
		file_PlantFlowerAcceptFlowerResultInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlantFlowerAcceptFlowerResultInfo_proto_rawDescData)
	})
	return file_PlantFlowerAcceptFlowerResultInfo_proto_rawDescData
}

var file_PlantFlowerAcceptFlowerResultInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_PlantFlowerAcceptFlowerResultInfo_proto_goTypes = []interface{}{
	(*PlantFlowerAcceptFlowerResultInfo)(nil), // 0: PlantFlowerAcceptFlowerResultInfo
	nil, // 1: PlantFlowerAcceptFlowerResultInfo.Unk3300PALEJEDIAJPEntry
	nil, // 2: PlantFlowerAcceptFlowerResultInfo.Unk3300ILAADDCOPKEEntry
}
var file_PlantFlowerAcceptFlowerResultInfo_proto_depIdxs = []int32{
	1, // 0: PlantFlowerAcceptFlowerResultInfo.Unk3300_PALEJEDIAJP:type_name -> PlantFlowerAcceptFlowerResultInfo.Unk3300PALEJEDIAJPEntry
	2, // 1: PlantFlowerAcceptFlowerResultInfo.Unk3300_ILAADDCOPKE:type_name -> PlantFlowerAcceptFlowerResultInfo.Unk3300ILAADDCOPKEEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_PlantFlowerAcceptFlowerResultInfo_proto_init() }
func file_PlantFlowerAcceptFlowerResultInfo_proto_init() {
	if File_PlantFlowerAcceptFlowerResultInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PlantFlowerAcceptFlowerResultInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlantFlowerAcceptFlowerResultInfo); i {
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
			RawDescriptor: file_PlantFlowerAcceptFlowerResultInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlantFlowerAcceptFlowerResultInfo_proto_goTypes,
		DependencyIndexes: file_PlantFlowerAcceptFlowerResultInfo_proto_depIdxs,
		MessageInfos:      file_PlantFlowerAcceptFlowerResultInfo_proto_msgTypes,
	}.Build()
	File_PlantFlowerAcceptFlowerResultInfo_proto = out.File
	file_PlantFlowerAcceptFlowerResultInfo_proto_rawDesc = nil
	file_PlantFlowerAcceptFlowerResultInfo_proto_goTypes = nil
	file_PlantFlowerAcceptFlowerResultInfo_proto_depIdxs = nil
}
