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
// source: CheckSegmentCRCReq.proto

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

// CmdId: 68
// Obf: IGMEGBIDJBE
type CheckSegmentCRCReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InfoList []*SegmentCRCInfo `protobuf:"bytes,2,rep,name=info_list,json=infoList,proto3" json:"info_list,omitempty"`
}

func (x *CheckSegmentCRCReq) Reset() {
	*x = CheckSegmentCRCReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CheckSegmentCRCReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckSegmentCRCReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckSegmentCRCReq) ProtoMessage() {}

func (x *CheckSegmentCRCReq) ProtoReflect() protoreflect.Message {
	mi := &file_CheckSegmentCRCReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckSegmentCRCReq.ProtoReflect.Descriptor instead.
func (*CheckSegmentCRCReq) Descriptor() ([]byte, []int) {
	return file_CheckSegmentCRCReq_proto_rawDescGZIP(), []int{0}
}

func (x *CheckSegmentCRCReq) GetInfoList() []*SegmentCRCInfo {
	if x != nil {
		return x.InfoList
	}
	return nil
}

var File_CheckSegmentCRCReq_proto protoreflect.FileDescriptor

var file_CheckSegmentCRCReq_proto_rawDesc = []byte{
	0x0a, 0x18, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x52,
	0x43, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x53, 0x65, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x43, 0x52, 0x43, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x42, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x43, 0x52, 0x43, 0x52, 0x65, 0x71, 0x12, 0x2c, 0x0a, 0x09, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x53, 0x65, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x43, 0x52, 0x43, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x69, 0x6e, 0x66, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CheckSegmentCRCReq_proto_rawDescOnce sync.Once
	file_CheckSegmentCRCReq_proto_rawDescData = file_CheckSegmentCRCReq_proto_rawDesc
)

func file_CheckSegmentCRCReq_proto_rawDescGZIP() []byte {
	file_CheckSegmentCRCReq_proto_rawDescOnce.Do(func() {
		file_CheckSegmentCRCReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_CheckSegmentCRCReq_proto_rawDescData)
	})
	return file_CheckSegmentCRCReq_proto_rawDescData
}

var file_CheckSegmentCRCReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CheckSegmentCRCReq_proto_goTypes = []interface{}{
	(*CheckSegmentCRCReq)(nil), // 0: CheckSegmentCRCReq
	(*SegmentCRCInfo)(nil),     // 1: SegmentCRCInfo
}
var file_CheckSegmentCRCReq_proto_depIdxs = []int32{
	1, // 0: CheckSegmentCRCReq.info_list:type_name -> SegmentCRCInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CheckSegmentCRCReq_proto_init() }
func file_CheckSegmentCRCReq_proto_init() {
	if File_CheckSegmentCRCReq_proto != nil {
		return
	}
	file_SegmentCRCInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CheckSegmentCRCReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckSegmentCRCReq); i {
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
			RawDescriptor: file_CheckSegmentCRCReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CheckSegmentCRCReq_proto_goTypes,
		DependencyIndexes: file_CheckSegmentCRCReq_proto_depIdxs,
		MessageInfos:      file_CheckSegmentCRCReq_proto_msgTypes,
	}.Build()
	File_CheckSegmentCRCReq_proto = out.File
	file_CheckSegmentCRCReq_proto_rawDesc = nil
	file_CheckSegmentCRCReq_proto_goTypes = nil
	file_CheckSegmentCRCReq_proto_depIdxs = nil
}
