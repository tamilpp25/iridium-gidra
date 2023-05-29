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
// source: MassivePropParam.proto

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

// Obf: EEBMHDBBEEH
type MassivePropParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type             int32     `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	ReactionInfoList []uint32  `protobuf:"varint,2,rep,packed,name=reaction_info_list,json=reactionInfoList,proto3" json:"reaction_info_list,omitempty"`
	ParamList        []float32 `protobuf:"fixed32,3,rep,packed,name=param_list,json=paramList,proto3" json:"param_list,omitempty"`
	SyncFlag         uint32    `protobuf:"varint,4,opt,name=sync_flag,json=syncFlag,proto3" json:"sync_flag,omitempty"`
}

func (x *MassivePropParam) Reset() {
	*x = MassivePropParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MassivePropParam_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MassivePropParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MassivePropParam) ProtoMessage() {}

func (x *MassivePropParam) ProtoReflect() protoreflect.Message {
	mi := &file_MassivePropParam_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MassivePropParam.ProtoReflect.Descriptor instead.
func (*MassivePropParam) Descriptor() ([]byte, []int) {
	return file_MassivePropParam_proto_rawDescGZIP(), []int{0}
}

func (x *MassivePropParam) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *MassivePropParam) GetReactionInfoList() []uint32 {
	if x != nil {
		return x.ReactionInfoList
	}
	return nil
}

func (x *MassivePropParam) GetParamList() []float32 {
	if x != nil {
		return x.ParamList
	}
	return nil
}

func (x *MassivePropParam) GetSyncFlag() uint32 {
	if x != nil {
		return x.SyncFlag
	}
	return 0
}

var File_MassivePropParam_proto protoreflect.FileDescriptor

var file_MassivePropParam_proto_rawDesc = []byte{
	0x0a, 0x16, 0x4d, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x01, 0x0a, 0x10, 0x4d, 0x61, 0x73,
	0x73, 0x69, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x10, 0x72,
	0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x02, 0x52, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x73, 0x79, 0x6e, 0x63, 0x46, 0x6c, 0x61, 0x67, 0x42, 0x06, 0x5a, 0x04, 0x2f,
	0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MassivePropParam_proto_rawDescOnce sync.Once
	file_MassivePropParam_proto_rawDescData = file_MassivePropParam_proto_rawDesc
)

func file_MassivePropParam_proto_rawDescGZIP() []byte {
	file_MassivePropParam_proto_rawDescOnce.Do(func() {
		file_MassivePropParam_proto_rawDescData = protoimpl.X.CompressGZIP(file_MassivePropParam_proto_rawDescData)
	})
	return file_MassivePropParam_proto_rawDescData
}

var file_MassivePropParam_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MassivePropParam_proto_goTypes = []interface{}{
	(*MassivePropParam)(nil), // 0: MassivePropParam
}
var file_MassivePropParam_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MassivePropParam_proto_init() }
func file_MassivePropParam_proto_init() {
	if File_MassivePropParam_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MassivePropParam_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MassivePropParam); i {
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
			RawDescriptor: file_MassivePropParam_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MassivePropParam_proto_goTypes,
		DependencyIndexes: file_MassivePropParam_proto_depIdxs,
		MessageInfos:      file_MassivePropParam_proto_msgTypes,
	}.Build()
	File_MassivePropParam_proto = out.File
	file_MassivePropParam_proto_rawDesc = nil
	file_MassivePropParam_proto_goTypes = nil
	file_MassivePropParam_proto_depIdxs = nil
}
