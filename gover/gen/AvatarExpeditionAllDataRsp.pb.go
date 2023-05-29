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
// source: AvatarExpeditionAllDataRsp.proto

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

// CmdId: 1722
// Obf: FIOCIHFCAPH
type AvatarExpeditionAllDataRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExpeditionInfoMap    map[uint64]*AvatarExpeditionInfo `protobuf:"bytes,14,rep,name=expedition_info_map,json=expeditionInfoMap,proto3" json:"expedition_info_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ExpeditionCountLimit uint32                           `protobuf:"varint,10,opt,name=expedition_count_limit,json=expeditionCountLimit,proto3" json:"expedition_count_limit,omitempty"`
	Retcode              int32                            `protobuf:"varint,13,opt,name=retcode,proto3" json:"retcode,omitempty"`
	OpenExpeditionList   []uint32                         `protobuf:"varint,5,rep,packed,name=open_expedition_list,json=openExpeditionList,proto3" json:"open_expedition_list,omitempty"`
}

func (x *AvatarExpeditionAllDataRsp) Reset() {
	*x = AvatarExpeditionAllDataRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarExpeditionAllDataRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarExpeditionAllDataRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarExpeditionAllDataRsp) ProtoMessage() {}

func (x *AvatarExpeditionAllDataRsp) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarExpeditionAllDataRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarExpeditionAllDataRsp.ProtoReflect.Descriptor instead.
func (*AvatarExpeditionAllDataRsp) Descriptor() ([]byte, []int) {
	return file_AvatarExpeditionAllDataRsp_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarExpeditionAllDataRsp) GetExpeditionInfoMap() map[uint64]*AvatarExpeditionInfo {
	if x != nil {
		return x.ExpeditionInfoMap
	}
	return nil
}

func (x *AvatarExpeditionAllDataRsp) GetExpeditionCountLimit() uint32 {
	if x != nil {
		return x.ExpeditionCountLimit
	}
	return 0
}

func (x *AvatarExpeditionAllDataRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *AvatarExpeditionAllDataRsp) GetOpenExpeditionList() []uint32 {
	if x != nil {
		return x.OpenExpeditionList
	}
	return nil
}

var File_AvatarExpeditionAllDataRsp_proto protoreflect.FileDescriptor

var file_AvatarExpeditionAllDataRsp_proto_rawDesc = []byte{
	0x0a, 0x20, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1a, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf,
	0x02, 0x0a, 0x1a, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x73, 0x70, 0x12, 0x62, 0x0a,
	0x13, 0x65, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x5f, 0x6d, 0x61, 0x70, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x6c,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x73, 0x70, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x11,
	0x65, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61,
	0x70, 0x12, 0x34, 0x0a, 0x16, 0x65, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x14, 0x65, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x30, 0x0a, 0x14, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x12, 0x6f, 0x70, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x1a, 0x5b, 0x0a, 0x16, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AvatarExpeditionAllDataRsp_proto_rawDescOnce sync.Once
	file_AvatarExpeditionAllDataRsp_proto_rawDescData = file_AvatarExpeditionAllDataRsp_proto_rawDesc
)

func file_AvatarExpeditionAllDataRsp_proto_rawDescGZIP() []byte {
	file_AvatarExpeditionAllDataRsp_proto_rawDescOnce.Do(func() {
		file_AvatarExpeditionAllDataRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarExpeditionAllDataRsp_proto_rawDescData)
	})
	return file_AvatarExpeditionAllDataRsp_proto_rawDescData
}

var file_AvatarExpeditionAllDataRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_AvatarExpeditionAllDataRsp_proto_goTypes = []interface{}{
	(*AvatarExpeditionAllDataRsp)(nil), // 0: AvatarExpeditionAllDataRsp
	nil,                                // 1: AvatarExpeditionAllDataRsp.ExpeditionInfoMapEntry
	(*AvatarExpeditionInfo)(nil),       // 2: AvatarExpeditionInfo
}
var file_AvatarExpeditionAllDataRsp_proto_depIdxs = []int32{
	1, // 0: AvatarExpeditionAllDataRsp.expedition_info_map:type_name -> AvatarExpeditionAllDataRsp.ExpeditionInfoMapEntry
	2, // 1: AvatarExpeditionAllDataRsp.ExpeditionInfoMapEntry.value:type_name -> AvatarExpeditionInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_AvatarExpeditionAllDataRsp_proto_init() }
func file_AvatarExpeditionAllDataRsp_proto_init() {
	if File_AvatarExpeditionAllDataRsp_proto != nil {
		return
	}
	file_AvatarExpeditionInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AvatarExpeditionAllDataRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarExpeditionAllDataRsp); i {
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
			RawDescriptor: file_AvatarExpeditionAllDataRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarExpeditionAllDataRsp_proto_goTypes,
		DependencyIndexes: file_AvatarExpeditionAllDataRsp_proto_depIdxs,
		MessageInfos:      file_AvatarExpeditionAllDataRsp_proto_msgTypes,
	}.Build()
	File_AvatarExpeditionAllDataRsp_proto = out.File
	file_AvatarExpeditionAllDataRsp_proto_rawDesc = nil
	file_AvatarExpeditionAllDataRsp_proto_goTypes = nil
	file_AvatarExpeditionAllDataRsp_proto_depIdxs = nil
}
