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
// source: QueryPathReq.proto

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

type QueryPathReq_OptionType int32

const (
	QueryPathReq_OPTION_TYPE_NONE         QueryPathReq_OptionType = 0
	QueryPathReq_OPTION_TYPE_NORMAL       QueryPathReq_OptionType = 1
	QueryPathReq_OPTION_TYPE_FIRST_CAN_GO QueryPathReq_OptionType = 2
)

// Enum value maps for QueryPathReq_OptionType.
var (
	QueryPathReq_OptionType_name = map[int32]string{
		0: "OPTION_TYPE_NONE",
		1: "OPTION_TYPE_NORMAL",
		2: "OPTION_TYPE_FIRST_CAN_GO",
	}
	QueryPathReq_OptionType_value = map[string]int32{
		"OPTION_TYPE_NONE":         0,
		"OPTION_TYPE_NORMAL":       1,
		"OPTION_TYPE_FIRST_CAN_GO": 2,
	}
)

func (x QueryPathReq_OptionType) Enum() *QueryPathReq_OptionType {
	p := new(QueryPathReq_OptionType)
	*p = x
	return p
}

func (x QueryPathReq_OptionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QueryPathReq_OptionType) Descriptor() protoreflect.EnumDescriptor {
	return file_QueryPathReq_proto_enumTypes[0].Descriptor()
}

func (QueryPathReq_OptionType) Type() protoreflect.EnumType {
	return &file_QueryPathReq_proto_enumTypes[0]
}

func (x QueryPathReq_OptionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QueryPathReq_OptionType.Descriptor instead.
func (QueryPathReq_OptionType) EnumDescriptor() ([]byte, []int) {
	return file_QueryPathReq_proto_rawDescGZIP(), []int{0, 0}
}

type QueryPathReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter              *QueryFilter            `protobuf:"bytes,7,opt,name=filter,proto3" json:"filter,omitempty"`
	QueryId             int32                   `protobuf:"varint,3,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
	DestinationPos      []*Vector               `protobuf:"bytes,6,rep,name=destination_pos,json=destinationPos,proto3" json:"destination_pos,omitempty"`
	QueryType           QueryPathReq_OptionType `protobuf:"varint,11,opt,name=query_type,json=queryType,proto3,enum=QueryPathReq_OptionType" json:"query_type,omitempty"`
	Unk3300_LHNGPJFOMIK *Vector3Int             `protobuf:"bytes,4,opt,name=Unk3300_LHNGPJFOMIK,json=Unk3300LHNGPJFOMIK,proto3" json:"Unk3300_LHNGPJFOMIK,omitempty"`
	SceneId             uint32                  `protobuf:"varint,2,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	SourcePos           *Vector                 `protobuf:"bytes,8,opt,name=source_pos,json=sourcePos,proto3" json:"source_pos,omitempty"`
	Unk3300_CLGJBBJDOLN *Vector3Int             `protobuf:"bytes,5,opt,name=Unk3300_CLGJBBJDOLN,json=Unk3300CLGJBBJDOLN,proto3" json:"Unk3300_CLGJBBJDOLN,omitempty"`
}

func (x *QueryPathReq) Reset() {
	*x = QueryPathReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_QueryPathReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPathReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPathReq) ProtoMessage() {}

func (x *QueryPathReq) ProtoReflect() protoreflect.Message {
	mi := &file_QueryPathReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPathReq.ProtoReflect.Descriptor instead.
func (*QueryPathReq) Descriptor() ([]byte, []int) {
	return file_QueryPathReq_proto_rawDescGZIP(), []int{0}
}

func (x *QueryPathReq) GetFilter() *QueryFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *QueryPathReq) GetQueryId() int32 {
	if x != nil {
		return x.QueryId
	}
	return 0
}

func (x *QueryPathReq) GetDestinationPos() []*Vector {
	if x != nil {
		return x.DestinationPos
	}
	return nil
}

func (x *QueryPathReq) GetQueryType() QueryPathReq_OptionType {
	if x != nil {
		return x.QueryType
	}
	return QueryPathReq_OPTION_TYPE_NONE
}

func (x *QueryPathReq) GetUnk3300_LHNGPJFOMIK() *Vector3Int {
	if x != nil {
		return x.Unk3300_LHNGPJFOMIK
	}
	return nil
}

func (x *QueryPathReq) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *QueryPathReq) GetSourcePos() *Vector {
	if x != nil {
		return x.SourcePos
	}
	return nil
}

func (x *QueryPathReq) GetUnk3300_CLGJBBJDOLN() *Vector3Int {
	if x != nil {
		return x.Unk3300_CLGJBBJDOLN
	}
	return nil
}

var File_QueryPathReq_proto protoreflect.FileDescriptor

var file_QueryPathReq_proto_rawDesc = []byte{
	0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x49, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x03, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x19,
	0x0a, 0x08, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x79, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x0f, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0e, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x73, 0x12, 0x37, 0x0a, 0x0a, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x18, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x2e, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x3c, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f,
	0x4c, 0x48, 0x4e, 0x47, 0x50, 0x4a, 0x46, 0x4f, 0x4d, 0x49, 0x4b, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x49, 0x6e, 0x74, 0x52, 0x12,
	0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x4c, 0x48, 0x4e, 0x47, 0x50, 0x4a, 0x46, 0x4f, 0x4d,
	0x49, 0x4b, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x26, 0x0a,
	0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x50, 0x6f, 0x73, 0x12, 0x3c, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30,
	0x5f, 0x43, 0x4c, 0x47, 0x4a, 0x42, 0x42, 0x4a, 0x44, 0x4f, 0x4c, 0x4e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x49, 0x6e, 0x74, 0x52,
	0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x43, 0x4c, 0x47, 0x4a, 0x42, 0x42, 0x4a, 0x44,
	0x4f, 0x4c, 0x4e, 0x22, 0x58, 0x0a, 0x0a, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x10, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x4f, 0x50, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x01, 0x12,
	0x1c, 0x0a, 0x18, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46,
	0x49, 0x52, 0x53, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x5f, 0x47, 0x4f, 0x10, 0x02, 0x42, 0x06, 0x5a,
	0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_QueryPathReq_proto_rawDescOnce sync.Once
	file_QueryPathReq_proto_rawDescData = file_QueryPathReq_proto_rawDesc
)

func file_QueryPathReq_proto_rawDescGZIP() []byte {
	file_QueryPathReq_proto_rawDescOnce.Do(func() {
		file_QueryPathReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_QueryPathReq_proto_rawDescData)
	})
	return file_QueryPathReq_proto_rawDescData
}

var file_QueryPathReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_QueryPathReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_QueryPathReq_proto_goTypes = []interface{}{
	(QueryPathReq_OptionType)(0), // 0: QueryPathReq.OptionType
	(*QueryPathReq)(nil),         // 1: QueryPathReq
	(*QueryFilter)(nil),          // 2: QueryFilter
	(*Vector)(nil),               // 3: Vector
	(*Vector3Int)(nil),           // 4: Vector3Int
}
var file_QueryPathReq_proto_depIdxs = []int32{
	2, // 0: QueryPathReq.filter:type_name -> QueryFilter
	3, // 1: QueryPathReq.destination_pos:type_name -> Vector
	0, // 2: QueryPathReq.query_type:type_name -> QueryPathReq.OptionType
	4, // 3: QueryPathReq.Unk3300_LHNGPJFOMIK:type_name -> Vector3Int
	3, // 4: QueryPathReq.source_pos:type_name -> Vector
	4, // 5: QueryPathReq.Unk3300_CLGJBBJDOLN:type_name -> Vector3Int
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_QueryPathReq_proto_init() }
func file_QueryPathReq_proto_init() {
	if File_QueryPathReq_proto != nil {
		return
	}
	file_QueryFilter_proto_init()
	file_Vector_proto_init()
	file_Vector3Int_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_QueryPathReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPathReq); i {
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
			RawDescriptor: file_QueryPathReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_QueryPathReq_proto_goTypes,
		DependencyIndexes: file_QueryPathReq_proto_depIdxs,
		EnumInfos:         file_QueryPathReq_proto_enumTypes,
		MessageInfos:      file_QueryPathReq_proto_msgTypes,
	}.Build()
	File_QueryPathReq_proto = out.File
	file_QueryPathReq_proto_rawDesc = nil
	file_QueryPathReq_proto_goTypes = nil
	file_QueryPathReq_proto_depIdxs = nil
}
