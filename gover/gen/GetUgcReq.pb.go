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
// source: GetUgcReq.proto

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

// CmdId: 6328
// Obf: AFLHDFGBBED
type GetUgcReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UgcType        UgcType     `protobuf:"varint,6,opt,name=ugc_type,json=ugcType,proto3,enum=UgcType" json:"ugc_type,omitempty"`
	UgcGuid        uint64      `protobuf:"varint,11,opt,name=ugc_guid,json=ugcGuid,proto3" json:"ugc_guid,omitempty"`
	UgcRecordUsage RecordUsage `protobuf:"varint,12,opt,name=ugc_record_usage,json=ugcRecordUsage,proto3,enum=RecordUsage" json:"ugc_record_usage,omitempty"`
	GetUgcType_    GetUgcType  `protobuf:"varint,3,opt,name=get_ugc_type,json=getUgcType,proto3,enum=GetUgcType" json:"get_ugc_type,omitempty"`
	IsRequireBrief bool        `protobuf:"varint,10,opt,name=is_require_brief,json=isRequireBrief,proto3" json:"is_require_brief,omitempty"`
	ScheduleId     uint32      `protobuf:"varint,14,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
}

func (x *GetUgcReq) Reset() {
	*x = GetUgcReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetUgcReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUgcReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUgcReq) ProtoMessage() {}

func (x *GetUgcReq) ProtoReflect() protoreflect.Message {
	mi := &file_GetUgcReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUgcReq.ProtoReflect.Descriptor instead.
func (*GetUgcReq) Descriptor() ([]byte, []int) {
	return file_GetUgcReq_proto_rawDescGZIP(), []int{0}
}

func (x *GetUgcReq) GetUgcType() UgcType {
	if x != nil {
		return x.UgcType
	}
	return UgcType_UGC_TYPE_NONE
}

func (x *GetUgcReq) GetUgcGuid() uint64 {
	if x != nil {
		return x.UgcGuid
	}
	return 0
}

func (x *GetUgcReq) GetUgcRecordUsage() RecordUsage {
	if x != nil {
		return x.UgcRecordUsage
	}
	return RecordUsage_UGC_RECORD_USAGE_NONE
}

func (x *GetUgcReq) GetGetUgcType_() GetUgcType {
	if x != nil {
		return x.GetUgcType_
	}
	return GetUgcType_GET_UGC_NONE
}

func (x *GetUgcReq) GetIsRequireBrief() bool {
	if x != nil {
		return x.IsRequireBrief
	}
	return false
}

func (x *GetUgcReq) GetScheduleId() uint32 {
	if x != nil {
		return x.ScheduleId
	}
	return 0
}

var File_GetUgcReq_proto protoreflect.FileDescriptor

var file_GetUgcReq_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x67, 0x63, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0d, 0x55, 0x67, 0x63, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x55, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x67, 0x63, 0x54, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x55, 0x67, 0x63,
	0x52, 0x65, 0x71, 0x12, 0x23, 0x0a, 0x08, 0x75, 0x67, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x08, 0x2e, 0x55, 0x67, 0x63, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x07, 0x75, 0x67, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x67, 0x63, 0x5f,
	0x67, 0x75, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x75, 0x67, 0x63, 0x47,
	0x75, 0x69, 0x64, 0x12, 0x36, 0x0a, 0x10, 0x75, 0x67, 0x63, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0e, 0x75, 0x67, 0x63,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x0c, 0x67,
	0x65, 0x74, 0x5f, 0x75, 0x67, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0b, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x67, 0x63, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a,
	0x67, 0x65, 0x74, 0x55, 0x67, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x73,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x5f, 0x62, 0x72, 0x69, 0x65, 0x66, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x42,
	0x72, 0x69, 0x65, 0x66, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetUgcReq_proto_rawDescOnce sync.Once
	file_GetUgcReq_proto_rawDescData = file_GetUgcReq_proto_rawDesc
)

func file_GetUgcReq_proto_rawDescGZIP() []byte {
	file_GetUgcReq_proto_rawDescOnce.Do(func() {
		file_GetUgcReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetUgcReq_proto_rawDescData)
	})
	return file_GetUgcReq_proto_rawDescData
}

var file_GetUgcReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetUgcReq_proto_goTypes = []interface{}{
	(*GetUgcReq)(nil), // 0: GetUgcReq
	(UgcType)(0),      // 1: UgcType
	(RecordUsage)(0),  // 2: RecordUsage
	(GetUgcType)(0),   // 3: GetUgcType
}
var file_GetUgcReq_proto_depIdxs = []int32{
	1, // 0: GetUgcReq.ugc_type:type_name -> UgcType
	2, // 1: GetUgcReq.ugc_record_usage:type_name -> RecordUsage
	3, // 2: GetUgcReq.get_ugc_type:type_name -> GetUgcType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_GetUgcReq_proto_init() }
func file_GetUgcReq_proto_init() {
	if File_GetUgcReq_proto != nil {
		return
	}
	file_UgcType_proto_init()
	file_RecordUsage_proto_init()
	file_GetUgcType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetUgcReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUgcReq); i {
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
			RawDescriptor: file_GetUgcReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetUgcReq_proto_goTypes,
		DependencyIndexes: file_GetUgcReq_proto_depIdxs,
		MessageInfos:      file_GetUgcReq_proto_msgTypes,
	}.Build()
	File_GetUgcReq_proto = out.File
	file_GetUgcReq_proto_rawDesc = nil
	file_GetUgcReq_proto_goTypes = nil
	file_GetUgcReq_proto_depIdxs = nil
}
