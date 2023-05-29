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
// source: AbilityScalarType.proto

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

// Obf: EBCOKIIELIE
type AbilityScalarType int32

const (
	AbilityScalarType_ABILITY_SCALAR_TYPE_UNKNOW  AbilityScalarType = 0
	AbilityScalarType_ABILITY_SCALAR_TYPE_FLOAT   AbilityScalarType = 1
	AbilityScalarType_ABILITY_SCALAR_TYPE_INT     AbilityScalarType = 2
	AbilityScalarType_ABILITY_SCALAR_TYPE_BOOL    AbilityScalarType = 3
	AbilityScalarType_ABILITY_SCALAR_TYPE_TRIGGER AbilityScalarType = 4
	AbilityScalarType_ABILITY_SCALAR_TYPE_STRING  AbilityScalarType = 5
	AbilityScalarType_ABILITY_SCALAR_TYPE_UINT    AbilityScalarType = 6
)

// Enum value maps for AbilityScalarType.
var (
	AbilityScalarType_name = map[int32]string{
		0: "ABILITY_SCALAR_TYPE_UNKNOW",
		1: "ABILITY_SCALAR_TYPE_FLOAT",
		2: "ABILITY_SCALAR_TYPE_INT",
		3: "ABILITY_SCALAR_TYPE_BOOL",
		4: "ABILITY_SCALAR_TYPE_TRIGGER",
		5: "ABILITY_SCALAR_TYPE_STRING",
		6: "ABILITY_SCALAR_TYPE_UINT",
	}
	AbilityScalarType_value = map[string]int32{
		"ABILITY_SCALAR_TYPE_UNKNOW":  0,
		"ABILITY_SCALAR_TYPE_FLOAT":   1,
		"ABILITY_SCALAR_TYPE_INT":     2,
		"ABILITY_SCALAR_TYPE_BOOL":    3,
		"ABILITY_SCALAR_TYPE_TRIGGER": 4,
		"ABILITY_SCALAR_TYPE_STRING":  5,
		"ABILITY_SCALAR_TYPE_UINT":    6,
	}
)

func (x AbilityScalarType) Enum() *AbilityScalarType {
	p := new(AbilityScalarType)
	*p = x
	return p
}

func (x AbilityScalarType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AbilityScalarType) Descriptor() protoreflect.EnumDescriptor {
	return file_AbilityScalarType_proto_enumTypes[0].Descriptor()
}

func (AbilityScalarType) Type() protoreflect.EnumType {
	return &file_AbilityScalarType_proto_enumTypes[0]
}

func (x AbilityScalarType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AbilityScalarType.Descriptor instead.
func (AbilityScalarType) EnumDescriptor() ([]byte, []int) {
	return file_AbilityScalarType_proto_rawDescGZIP(), []int{0}
}

var File_AbilityScalarType_proto protoreflect.FileDescriptor

var file_AbilityScalarType_proto_rawDesc = []byte{
	0x0a, 0x17, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xec, 0x01, 0x0a, 0x11, 0x41, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1e, 0x0a, 0x1a, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x53, 0x43, 0x41, 0x4c, 0x41,
	0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x10, 0x00, 0x12,
	0x1d, 0x0a, 0x19, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x53, 0x43, 0x41, 0x4c, 0x41,
	0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x10, 0x01, 0x12, 0x1b,
	0x0a, 0x17, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x53, 0x43, 0x41, 0x4c, 0x41, 0x52,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x41,
	0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x53, 0x43, 0x41, 0x4c, 0x41, 0x52, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x10, 0x03, 0x12, 0x1f, 0x0a, 0x1b, 0x41, 0x42, 0x49,
	0x4c, 0x49, 0x54, 0x59, 0x5f, 0x53, 0x43, 0x41, 0x4c, 0x41, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x54, 0x52, 0x49, 0x47, 0x47, 0x45, 0x52, 0x10, 0x04, 0x12, 0x1e, 0x0a, 0x1a, 0x41, 0x42,
	0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x53, 0x43, 0x41, 0x4c, 0x41, 0x52, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x1c, 0x0a, 0x18, 0x41, 0x42,
	0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x53, 0x43, 0x41, 0x4c, 0x41, 0x52, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x10, 0x06, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AbilityScalarType_proto_rawDescOnce sync.Once
	file_AbilityScalarType_proto_rawDescData = file_AbilityScalarType_proto_rawDesc
)

func file_AbilityScalarType_proto_rawDescGZIP() []byte {
	file_AbilityScalarType_proto_rawDescOnce.Do(func() {
		file_AbilityScalarType_proto_rawDescData = protoimpl.X.CompressGZIP(file_AbilityScalarType_proto_rawDescData)
	})
	return file_AbilityScalarType_proto_rawDescData
}

var file_AbilityScalarType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_AbilityScalarType_proto_goTypes = []interface{}{
	(AbilityScalarType)(0), // 0: AbilityScalarType
}
var file_AbilityScalarType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AbilityScalarType_proto_init() }
func file_AbilityScalarType_proto_init() {
	if File_AbilityScalarType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_AbilityScalarType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AbilityScalarType_proto_goTypes,
		DependencyIndexes: file_AbilityScalarType_proto_depIdxs,
		EnumInfos:         file_AbilityScalarType_proto_enumTypes,
	}.Build()
	File_AbilityScalarType_proto = out.File
	file_AbilityScalarType_proto_rawDesc = nil
	file_AbilityScalarType_proto_goTypes = nil
	file_AbilityScalarType_proto_depIdxs = nil
}
