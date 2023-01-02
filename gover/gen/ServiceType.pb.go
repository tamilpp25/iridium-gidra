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
// source: ServiceType.proto

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

type ServiceType int32

const (
	ServiceType_SERVICE_TYPE_NONE        ServiceType = 0
	ServiceType_SERVICE_TYPE_CLIENT      ServiceType = 1
	ServiceType_SERVICE_TYPE_GATE        ServiceType = 2
	ServiceType_SERVICE_TYPE_GAME        ServiceType = 3
	ServiceType_SERVICE_TYPE_NODE        ServiceType = 4
	ServiceType_SERVICE_TYPE_DB          ServiceType = 5
	ServiceType_SERVICE_TYPE_SNS         ServiceType = 6
	ServiceType_SERVICE_TYPE_DISPATCH    ServiceType = 7
	ServiceType_SERVICE_TYPE_MUIP        ServiceType = 8
	ServiceType_SERVICE_TYPE_OFFLINE_MSG ServiceType = 9
	ServiceType_SERVICE_TYPE_MAIL        ServiceType = 10
	ServiceType_SERVICE_TYPE_MP          ServiceType = 11
	ServiceType_SERVICE_TYPE_HTTPPROXY   ServiceType = 12
	ServiceType_SERVICE_TYPE_ACTIVITY    ServiceType = 13
	ServiceType_SERVICE_TYPE_PATHFINDING ServiceType = 14
	ServiceType_SERVICE_TYPE_SOCIAL      ServiceType = 15
	ServiceType_SERVICE_TYPE_OA          ServiceType = 16
	ServiceType_SERVICE_TYPE_MATCH       ServiceType = 17
	ServiceType_SERVICE_TYPE_OFFLINE_OP  ServiceType = 18
	ServiceType_SERVICE_TYPE_TOTHEMOON   ServiceType = 19
	ServiceType_SERVICE_TYPE_GCG         ServiceType = 20
)

// Enum value maps for ServiceType.
var (
	ServiceType_name = map[int32]string{
		0:  "SERVICE_TYPE_NONE",
		1:  "SERVICE_TYPE_CLIENT",
		2:  "SERVICE_TYPE_GATE",
		3:  "SERVICE_TYPE_GAME",
		4:  "SERVICE_TYPE_NODE",
		5:  "SERVICE_TYPE_DB",
		6:  "SERVICE_TYPE_SNS",
		7:  "SERVICE_TYPE_DISPATCH",
		8:  "SERVICE_TYPE_MUIP",
		9:  "SERVICE_TYPE_OFFLINE_MSG",
		10: "SERVICE_TYPE_MAIL",
		11: "SERVICE_TYPE_MP",
		12: "SERVICE_TYPE_HTTPPROXY",
		13: "SERVICE_TYPE_ACTIVITY",
		14: "SERVICE_TYPE_PATHFINDING",
		15: "SERVICE_TYPE_SOCIAL",
		16: "SERVICE_TYPE_OA",
		17: "SERVICE_TYPE_MATCH",
		18: "SERVICE_TYPE_OFFLINE_OP",
		19: "SERVICE_TYPE_TOTHEMOON",
		20: "SERVICE_TYPE_GCG",
	}
	ServiceType_value = map[string]int32{
		"SERVICE_TYPE_NONE":        0,
		"SERVICE_TYPE_CLIENT":      1,
		"SERVICE_TYPE_GATE":        2,
		"SERVICE_TYPE_GAME":        3,
		"SERVICE_TYPE_NODE":        4,
		"SERVICE_TYPE_DB":          5,
		"SERVICE_TYPE_SNS":         6,
		"SERVICE_TYPE_DISPATCH":    7,
		"SERVICE_TYPE_MUIP":        8,
		"SERVICE_TYPE_OFFLINE_MSG": 9,
		"SERVICE_TYPE_MAIL":        10,
		"SERVICE_TYPE_MP":          11,
		"SERVICE_TYPE_HTTPPROXY":   12,
		"SERVICE_TYPE_ACTIVITY":    13,
		"SERVICE_TYPE_PATHFINDING": 14,
		"SERVICE_TYPE_SOCIAL":      15,
		"SERVICE_TYPE_OA":          16,
		"SERVICE_TYPE_MATCH":       17,
		"SERVICE_TYPE_OFFLINE_OP":  18,
		"SERVICE_TYPE_TOTHEMOON":   19,
		"SERVICE_TYPE_GCG":         20,
	}
)

func (x ServiceType) Enum() *ServiceType {
	p := new(ServiceType)
	*p = x
	return p
}

func (x ServiceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceType) Descriptor() protoreflect.EnumDescriptor {
	return file_ServiceType_proto_enumTypes[0].Descriptor()
}

func (ServiceType) Type() protoreflect.EnumType {
	return &file_ServiceType_proto_enumTypes[0]
}

func (x ServiceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceType.Descriptor instead.
func (ServiceType) EnumDescriptor() ([]byte, []int) {
	return file_ServiceType_proto_rawDescGZIP(), []int{0}
}

var File_ServiceType_proto protoreflect.FileDescriptor

var file_ServiceType_proto_rawDesc = []byte{
	0x0a, 0x11, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0x93, 0x04, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x45,
	0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e,
	0x54, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x47, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x45,
	0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x10,
	0x03, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45, 0x52, 0x56,
	0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x42, 0x10, 0x05, 0x12, 0x14, 0x0a,
	0x10, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x4e,
	0x53, 0x10, 0x06, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x44, 0x49, 0x53, 0x50, 0x41, 0x54, 0x43, 0x48, 0x10, 0x07, 0x12, 0x15,
	0x0a, 0x11, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d,
	0x55, 0x49, 0x50, 0x10, 0x08, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x46, 0x46, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x4d, 0x53,
	0x47, 0x10, 0x09, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x0a, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45,
	0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x50, 0x10, 0x0b, 0x12,
	0x1a, 0x0a, 0x16, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x48, 0x54, 0x54, 0x50, 0x50, 0x52, 0x4f, 0x58, 0x59, 0x10, 0x0c, 0x12, 0x19, 0x0a, 0x15, 0x53,
	0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x56, 0x49, 0x54, 0x59, 0x10, 0x0d, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x41, 0x54, 0x48, 0x46, 0x49, 0x4e, 0x44, 0x49,
	0x4e, 0x47, 0x10, 0x0e, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x4f, 0x43, 0x49, 0x41, 0x4c, 0x10, 0x0f, 0x12, 0x13, 0x0a,
	0x0f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x41,
	0x10, 0x10, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x10, 0x11, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x45,
	0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x46, 0x46, 0x4c, 0x49,
	0x4e, 0x45, 0x5f, 0x4f, 0x50, 0x10, 0x12, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x45, 0x52, 0x56, 0x49,
	0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x4f, 0x54, 0x48, 0x45, 0x4d, 0x4f, 0x4f,
	0x4e, 0x10, 0x13, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x47, 0x43, 0x47, 0x10, 0x14, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ServiceType_proto_rawDescOnce sync.Once
	file_ServiceType_proto_rawDescData = file_ServiceType_proto_rawDesc
)

func file_ServiceType_proto_rawDescGZIP() []byte {
	file_ServiceType_proto_rawDescOnce.Do(func() {
		file_ServiceType_proto_rawDescData = protoimpl.X.CompressGZIP(file_ServiceType_proto_rawDescData)
	})
	return file_ServiceType_proto_rawDescData
}

var file_ServiceType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ServiceType_proto_goTypes = []interface{}{
	(ServiceType)(0), // 0: ServiceType
}
var file_ServiceType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ServiceType_proto_init() }
func file_ServiceType_proto_init() {
	if File_ServiceType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ServiceType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ServiceType_proto_goTypes,
		DependencyIndexes: file_ServiceType_proto_depIdxs,
		EnumInfos:         file_ServiceType_proto_enumTypes,
	}.Build()
	File_ServiceType_proto = out.File
	file_ServiceType_proto_rawDesc = nil
	file_ServiceType_proto_goTypes = nil
	file_ServiceType_proto_depIdxs = nil
}
