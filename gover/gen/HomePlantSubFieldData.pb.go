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
// source: HomePlantSubFieldData.proto

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

// Obf: JICIBIJIHAB
type HomePlantSubFieldData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MMNMLFLBMCF  HomePlantFieldStatus `protobuf:"varint,3,opt,name=MMNMLFLBMCF,proto3,enum=HomePlantFieldStatus" json:"MMNMLFLBMCF,omitempty"`
	JBCOCICMEJE  uint32               `protobuf:"varint,4,opt,name=JBCOCICMEJE,proto3" json:"JBCOCICMEJE,omitempty"`
	EndTime      uint32               `protobuf:"fixed32,15,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	EntityIdList []uint32             `protobuf:"varint,2,rep,packed,name=entity_id_list,json=entityIdList,proto3" json:"entity_id_list,omitempty"`
	FDCENJDIOCF  uint32               `protobuf:"varint,9,opt,name=FDCENJDIOCF,proto3" json:"FDCENJDIOCF,omitempty"`
}

func (x *HomePlantSubFieldData) Reset() {
	*x = HomePlantSubFieldData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HomePlantSubFieldData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomePlantSubFieldData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomePlantSubFieldData) ProtoMessage() {}

func (x *HomePlantSubFieldData) ProtoReflect() protoreflect.Message {
	mi := &file_HomePlantSubFieldData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomePlantSubFieldData.ProtoReflect.Descriptor instead.
func (*HomePlantSubFieldData) Descriptor() ([]byte, []int) {
	return file_HomePlantSubFieldData_proto_rawDescGZIP(), []int{0}
}

func (x *HomePlantSubFieldData) GetMMNMLFLBMCF() HomePlantFieldStatus {
	if x != nil {
		return x.MMNMLFLBMCF
	}
	return HomePlantFieldStatus_HOME_FIELD_STATUE_NONE
}

func (x *HomePlantSubFieldData) GetJBCOCICMEJE() uint32 {
	if x != nil {
		return x.JBCOCICMEJE
	}
	return 0
}

func (x *HomePlantSubFieldData) GetEndTime() uint32 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *HomePlantSubFieldData) GetEntityIdList() []uint32 {
	if x != nil {
		return x.EntityIdList
	}
	return nil
}

func (x *HomePlantSubFieldData) GetFDCENJDIOCF() uint32 {
	if x != nil {
		return x.FDCENJDIOCF
	}
	return 0
}

var File_HomePlantSubFieldData_proto protoreflect.FileDescriptor

var file_HomePlantSubFieldData_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x48, 0x6f, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x48,
	0x6f, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x01, 0x0a, 0x15, 0x48, 0x6f,
	0x6d, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x37, 0x0a, 0x0b, 0x4d, 0x4d, 0x4e, 0x4d, 0x4c, 0x46, 0x4c, 0x42, 0x4d,
	0x43, 0x46, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x0b, 0x4d, 0x4d, 0x4e, 0x4d, 0x4c, 0x46, 0x4c, 0x42, 0x4d, 0x43, 0x46, 0x12, 0x20, 0x0a, 0x0b,
	0x4a, 0x42, 0x43, 0x4f, 0x43, 0x49, 0x43, 0x4d, 0x45, 0x4a, 0x45, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x4a, 0x42, 0x43, 0x4f, 0x43, 0x49, 0x43, 0x4d, 0x45, 0x4a, 0x45, 0x12, 0x19,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x07,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x0c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x46, 0x44, 0x43, 0x45, 0x4e, 0x4a, 0x44, 0x49, 0x4f, 0x43, 0x46, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x46, 0x44, 0x43, 0x45, 0x4e, 0x4a, 0x44, 0x49, 0x4f, 0x43,
	0x46, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_HomePlantSubFieldData_proto_rawDescOnce sync.Once
	file_HomePlantSubFieldData_proto_rawDescData = file_HomePlantSubFieldData_proto_rawDesc
)

func file_HomePlantSubFieldData_proto_rawDescGZIP() []byte {
	file_HomePlantSubFieldData_proto_rawDescOnce.Do(func() {
		file_HomePlantSubFieldData_proto_rawDescData = protoimpl.X.CompressGZIP(file_HomePlantSubFieldData_proto_rawDescData)
	})
	return file_HomePlantSubFieldData_proto_rawDescData
}

var file_HomePlantSubFieldData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HomePlantSubFieldData_proto_goTypes = []interface{}{
	(*HomePlantSubFieldData)(nil), // 0: HomePlantSubFieldData
	(HomePlantFieldStatus)(0),     // 1: HomePlantFieldStatus
}
var file_HomePlantSubFieldData_proto_depIdxs = []int32{
	1, // 0: HomePlantSubFieldData.MMNMLFLBMCF:type_name -> HomePlantFieldStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_HomePlantSubFieldData_proto_init() }
func file_HomePlantSubFieldData_proto_init() {
	if File_HomePlantSubFieldData_proto != nil {
		return
	}
	file_HomePlantFieldStatus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HomePlantSubFieldData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomePlantSubFieldData); i {
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
			RawDescriptor: file_HomePlantSubFieldData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HomePlantSubFieldData_proto_goTypes,
		DependencyIndexes: file_HomePlantSubFieldData_proto_depIdxs,
		MessageInfos:      file_HomePlantSubFieldData_proto_msgTypes,
	}.Build()
	File_HomePlantSubFieldData_proto = out.File
	file_HomePlantSubFieldData_proto_rawDesc = nil
	file_HomePlantSubFieldData_proto_goTypes = nil
	file_HomePlantSubFieldData_proto_depIdxs = nil
}
