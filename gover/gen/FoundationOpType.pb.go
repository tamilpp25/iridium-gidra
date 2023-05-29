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
// source: FoundationOpType.proto

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

// Obf: MDNDNFJFKDJ
type FoundationOpType int32

const (
	FoundationOpType_FOUNDATION_OP_NONE       FoundationOpType = 0
	FoundationOpType_FOUNDATION_OP_BUILD      FoundationOpType = 1
	FoundationOpType_FOUNDATION_OP_DEMOLITION FoundationOpType = 2
	FoundationOpType_FOUNDATION_OP_REBUILD    FoundationOpType = 3
	FoundationOpType_FOUNDATION_OP_ROTATE     FoundationOpType = 4
	FoundationOpType_FOUNDATION_OP_LOCK       FoundationOpType = 5
	FoundationOpType_FOUNDATION_OP_UNLOCK     FoundationOpType = 6
)

// Enum value maps for FoundationOpType.
var (
	FoundationOpType_name = map[int32]string{
		0: "FOUNDATION_OP_NONE",
		1: "FOUNDATION_OP_BUILD",
		2: "FOUNDATION_OP_DEMOLITION",
		3: "FOUNDATION_OP_REBUILD",
		4: "FOUNDATION_OP_ROTATE",
		5: "FOUNDATION_OP_LOCK",
		6: "FOUNDATION_OP_UNLOCK",
	}
	FoundationOpType_value = map[string]int32{
		"FOUNDATION_OP_NONE":       0,
		"FOUNDATION_OP_BUILD":      1,
		"FOUNDATION_OP_DEMOLITION": 2,
		"FOUNDATION_OP_REBUILD":    3,
		"FOUNDATION_OP_ROTATE":     4,
		"FOUNDATION_OP_LOCK":       5,
		"FOUNDATION_OP_UNLOCK":     6,
	}
)

func (x FoundationOpType) Enum() *FoundationOpType {
	p := new(FoundationOpType)
	*p = x
	return p
}

func (x FoundationOpType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FoundationOpType) Descriptor() protoreflect.EnumDescriptor {
	return file_FoundationOpType_proto_enumTypes[0].Descriptor()
}

func (FoundationOpType) Type() protoreflect.EnumType {
	return &file_FoundationOpType_proto_enumTypes[0]
}

func (x FoundationOpType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FoundationOpType.Descriptor instead.
func (FoundationOpType) EnumDescriptor() ([]byte, []int) {
	return file_FoundationOpType_proto_rawDescGZIP(), []int{0}
}

var File_FoundationOpType_proto protoreflect.FileDescriptor

var file_FoundationOpType_proto_rawDesc = []byte{
	0x0a, 0x16, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x54, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xc8, 0x01, 0x0a, 0x10, 0x46, 0x6f, 0x75,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x12, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x50, 0x5f, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x50, 0x5f, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x10, 0x01, 0x12, 0x1c,
	0x0a, 0x18, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x50, 0x5f,
	0x44, 0x45, 0x4d, 0x4f, 0x4c, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15,
	0x46, 0x4f, 0x55, 0x4e, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x50, 0x5f, 0x52, 0x45,
	0x42, 0x55, 0x49, 0x4c, 0x44, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x46, 0x4f, 0x55, 0x4e, 0x44,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x50, 0x5f, 0x52, 0x4f, 0x54, 0x41, 0x54, 0x45, 0x10,
	0x04, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x4f, 0x50, 0x5f, 0x4c, 0x4f, 0x43, 0x4b, 0x10, 0x05, 0x12, 0x18, 0x0a, 0x14, 0x46, 0x4f, 0x55,
	0x4e, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x50, 0x5f, 0x55, 0x4e, 0x4c, 0x4f, 0x43,
	0x4b, 0x10, 0x06, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_FoundationOpType_proto_rawDescOnce sync.Once
	file_FoundationOpType_proto_rawDescData = file_FoundationOpType_proto_rawDesc
)

func file_FoundationOpType_proto_rawDescGZIP() []byte {
	file_FoundationOpType_proto_rawDescOnce.Do(func() {
		file_FoundationOpType_proto_rawDescData = protoimpl.X.CompressGZIP(file_FoundationOpType_proto_rawDescData)
	})
	return file_FoundationOpType_proto_rawDescData
}

var file_FoundationOpType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_FoundationOpType_proto_goTypes = []interface{}{
	(FoundationOpType)(0), // 0: FoundationOpType
}
var file_FoundationOpType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_FoundationOpType_proto_init() }
func file_FoundationOpType_proto_init() {
	if File_FoundationOpType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_FoundationOpType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FoundationOpType_proto_goTypes,
		DependencyIndexes: file_FoundationOpType_proto_depIdxs,
		EnumInfos:         file_FoundationOpType_proto_enumTypes,
	}.Build()
	File_FoundationOpType_proto = out.File
	file_FoundationOpType_proto_rawDesc = nil
	file_FoundationOpType_proto_goTypes = nil
	file_FoundationOpType_proto_depIdxs = nil
}
