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
// source: WinterCampEditSnowmanCombinationReq.proto

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

// CmdId: 8234
// Obf: DBOMOCKKAHL
type WinterCampEditSnowmanCombinationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId        uint32                `protobuf:"varint,5,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	CombinationInfo *CustomGadgetTreeInfo `protobuf:"bytes,3,opt,name=combination_info,json=combinationInfo,proto3" json:"combination_info,omitempty"`
}

func (x *WinterCampEditSnowmanCombinationReq) Reset() {
	*x = WinterCampEditSnowmanCombinationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WinterCampEditSnowmanCombinationReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WinterCampEditSnowmanCombinationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WinterCampEditSnowmanCombinationReq) ProtoMessage() {}

func (x *WinterCampEditSnowmanCombinationReq) ProtoReflect() protoreflect.Message {
	mi := &file_WinterCampEditSnowmanCombinationReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WinterCampEditSnowmanCombinationReq.ProtoReflect.Descriptor instead.
func (*WinterCampEditSnowmanCombinationReq) Descriptor() ([]byte, []int) {
	return file_WinterCampEditSnowmanCombinationReq_proto_rawDescGZIP(), []int{0}
}

func (x *WinterCampEditSnowmanCombinationReq) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *WinterCampEditSnowmanCombinationReq) GetCombinationInfo() *CustomGadgetTreeInfo {
	if x != nil {
		return x.CombinationInfo
	}
	return nil
}

var File_WinterCampEditSnowmanCombinationReq_proto protoreflect.FileDescriptor

var file_WinterCampEditSnowmanCombinationReq_proto_rawDesc = []byte{
	0x0a, 0x29, 0x57, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x61, 0x6d, 0x70, 0x45, 0x64, 0x69, 0x74,
	0x53, 0x6e, 0x6f, 0x77, 0x6d, 0x61, 0x6e, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x54, 0x72, 0x65, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x23, 0x57, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x43, 0x61, 0x6d, 0x70, 0x45, 0x64, 0x69, 0x74, 0x53, 0x6e, 0x6f, 0x77, 0x6d, 0x61,
	0x6e, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12,
	0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x10,
	0x63, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x47,
	0x61, 0x64, 0x67, 0x65, 0x74, 0x54, 0x72, 0x65, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x63,
	0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x06,
	0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_WinterCampEditSnowmanCombinationReq_proto_rawDescOnce sync.Once
	file_WinterCampEditSnowmanCombinationReq_proto_rawDescData = file_WinterCampEditSnowmanCombinationReq_proto_rawDesc
)

func file_WinterCampEditSnowmanCombinationReq_proto_rawDescGZIP() []byte {
	file_WinterCampEditSnowmanCombinationReq_proto_rawDescOnce.Do(func() {
		file_WinterCampEditSnowmanCombinationReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_WinterCampEditSnowmanCombinationReq_proto_rawDescData)
	})
	return file_WinterCampEditSnowmanCombinationReq_proto_rawDescData
}

var file_WinterCampEditSnowmanCombinationReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WinterCampEditSnowmanCombinationReq_proto_goTypes = []interface{}{
	(*WinterCampEditSnowmanCombinationReq)(nil), // 0: WinterCampEditSnowmanCombinationReq
	(*CustomGadgetTreeInfo)(nil),                // 1: CustomGadgetTreeInfo
}
var file_WinterCampEditSnowmanCombinationReq_proto_depIdxs = []int32{
	1, // 0: WinterCampEditSnowmanCombinationReq.combination_info:type_name -> CustomGadgetTreeInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_WinterCampEditSnowmanCombinationReq_proto_init() }
func file_WinterCampEditSnowmanCombinationReq_proto_init() {
	if File_WinterCampEditSnowmanCombinationReq_proto != nil {
		return
	}
	file_CustomGadgetTreeInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_WinterCampEditSnowmanCombinationReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WinterCampEditSnowmanCombinationReq); i {
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
			RawDescriptor: file_WinterCampEditSnowmanCombinationReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WinterCampEditSnowmanCombinationReq_proto_goTypes,
		DependencyIndexes: file_WinterCampEditSnowmanCombinationReq_proto_depIdxs,
		MessageInfos:      file_WinterCampEditSnowmanCombinationReq_proto_msgTypes,
	}.Build()
	File_WinterCampEditSnowmanCombinationReq_proto = out.File
	file_WinterCampEditSnowmanCombinationReq_proto_rawDesc = nil
	file_WinterCampEditSnowmanCombinationReq_proto_goTypes = nil
	file_WinterCampEditSnowmanCombinationReq_proto_depIdxs = nil
}
