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
// source: SalesmanStatusType.proto

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

type SalesmanStatusType int32

const (
	SalesmanStatusType_SALESMAN_STATUS_TYPE_NONE      SalesmanStatusType = 0
	SalesmanStatusType_SALESMAN_STATUS_TYPE_UNSTARTED SalesmanStatusType = 1
	SalesmanStatusType_SALESMAN_STATUS_TYPE_STARTED   SalesmanStatusType = 2
	SalesmanStatusType_SALESMAN_STATUS_TYPE_DELIVERED SalesmanStatusType = 3
)

// Enum value maps for SalesmanStatusType.
var (
	SalesmanStatusType_name = map[int32]string{
		0: "SALESMAN_STATUS_TYPE_NONE",
		1: "SALESMAN_STATUS_TYPE_UNSTARTED",
		2: "SALESMAN_STATUS_TYPE_STARTED",
		3: "SALESMAN_STATUS_TYPE_DELIVERED",
	}
	SalesmanStatusType_value = map[string]int32{
		"SALESMAN_STATUS_TYPE_NONE":      0,
		"SALESMAN_STATUS_TYPE_UNSTARTED": 1,
		"SALESMAN_STATUS_TYPE_STARTED":   2,
		"SALESMAN_STATUS_TYPE_DELIVERED": 3,
	}
)

func (x SalesmanStatusType) Enum() *SalesmanStatusType {
	p := new(SalesmanStatusType)
	*p = x
	return p
}

func (x SalesmanStatusType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SalesmanStatusType) Descriptor() protoreflect.EnumDescriptor {
	return file_SalesmanStatusType_proto_enumTypes[0].Descriptor()
}

func (SalesmanStatusType) Type() protoreflect.EnumType {
	return &file_SalesmanStatusType_proto_enumTypes[0]
}

func (x SalesmanStatusType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SalesmanStatusType.Descriptor instead.
func (SalesmanStatusType) EnumDescriptor() ([]byte, []int) {
	return file_SalesmanStatusType_proto_rawDescGZIP(), []int{0}
}

var File_SalesmanStatusType_proto protoreflect.FileDescriptor

var file_SalesmanStatusType_proto_rawDesc = []byte{
	0x0a, 0x18, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x6d, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x9d, 0x01, 0x0a, 0x12, 0x53,
	0x61, 0x6c, 0x65, 0x73, 0x6d, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x41, 0x4c, 0x45, 0x53, 0x4d, 0x41, 0x4e, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x22, 0x0a, 0x1e, 0x53, 0x41, 0x4c, 0x45, 0x53, 0x4d, 0x41, 0x4e, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x54, 0x41, 0x52, 0x54,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x41, 0x4c, 0x45, 0x53, 0x4d, 0x41, 0x4e,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41,
	0x52, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x22, 0x0a, 0x1e, 0x53, 0x41, 0x4c, 0x45, 0x53, 0x4d,
	0x41, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44,
	0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x45, 0x44, 0x10, 0x03, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67,
	0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SalesmanStatusType_proto_rawDescOnce sync.Once
	file_SalesmanStatusType_proto_rawDescData = file_SalesmanStatusType_proto_rawDesc
)

func file_SalesmanStatusType_proto_rawDescGZIP() []byte {
	file_SalesmanStatusType_proto_rawDescOnce.Do(func() {
		file_SalesmanStatusType_proto_rawDescData = protoimpl.X.CompressGZIP(file_SalesmanStatusType_proto_rawDescData)
	})
	return file_SalesmanStatusType_proto_rawDescData
}

var file_SalesmanStatusType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_SalesmanStatusType_proto_goTypes = []interface{}{
	(SalesmanStatusType)(0), // 0: SalesmanStatusType
}
var file_SalesmanStatusType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SalesmanStatusType_proto_init() }
func file_SalesmanStatusType_proto_init() {
	if File_SalesmanStatusType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_SalesmanStatusType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SalesmanStatusType_proto_goTypes,
		DependencyIndexes: file_SalesmanStatusType_proto_depIdxs,
		EnumInfos:         file_SalesmanStatusType_proto_enumTypes,
	}.Build()
	File_SalesmanStatusType_proto = out.File
	file_SalesmanStatusType_proto_rawDesc = nil
	file_SalesmanStatusType_proto_goTypes = nil
	file_SalesmanStatusType_proto_depIdxs = nil
}
