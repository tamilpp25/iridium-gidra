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
// source: LuaShellType.proto

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

// Obf: OEHKPJCBOEA
type LuaShellType int32

const (
	LuaShellType_LUASHELL_NONE       LuaShellType = 0
	LuaShellType_LUASHELL_NORMAL     LuaShellType = 1
	LuaShellType_LUASHELL_SECURITY   LuaShellType = 2
	LuaShellType_LUASHELL_SHELL_CODE LuaShellType = 3
)

// Enum value maps for LuaShellType.
var (
	LuaShellType_name = map[int32]string{
		0: "LUASHELL_NONE",
		1: "LUASHELL_NORMAL",
		2: "LUASHELL_SECURITY",
		3: "LUASHELL_SHELL_CODE",
	}
	LuaShellType_value = map[string]int32{
		"LUASHELL_NONE":       0,
		"LUASHELL_NORMAL":     1,
		"LUASHELL_SECURITY":   2,
		"LUASHELL_SHELL_CODE": 3,
	}
)

func (x LuaShellType) Enum() *LuaShellType {
	p := new(LuaShellType)
	*p = x
	return p
}

func (x LuaShellType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LuaShellType) Descriptor() protoreflect.EnumDescriptor {
	return file_LuaShellType_proto_enumTypes[0].Descriptor()
}

func (LuaShellType) Type() protoreflect.EnumType {
	return &file_LuaShellType_proto_enumTypes[0]
}

func (x LuaShellType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LuaShellType.Descriptor instead.
func (LuaShellType) EnumDescriptor() ([]byte, []int) {
	return file_LuaShellType_proto_rawDescGZIP(), []int{0}
}

var File_LuaShellType_proto protoreflect.FileDescriptor

var file_LuaShellType_proto_rawDesc = []byte{
	0x0a, 0x12, 0x4c, 0x75, 0x61, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x66, 0x0a, 0x0c, 0x4c, 0x75, 0x61, 0x53, 0x68, 0x65, 0x6c, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x55, 0x41, 0x53, 0x48, 0x45, 0x4c, 0x4c,
	0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x55, 0x41, 0x53, 0x48,
	0x45, 0x4c, 0x4c, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11,
	0x4c, 0x55, 0x41, 0x53, 0x48, 0x45, 0x4c, 0x4c, 0x5f, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54,
	0x59, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x4c, 0x55, 0x41, 0x53, 0x48, 0x45, 0x4c, 0x4c, 0x5f,
	0x53, 0x48, 0x45, 0x4c, 0x4c, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x03, 0x42, 0x06, 0x5a, 0x04,
	0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LuaShellType_proto_rawDescOnce sync.Once
	file_LuaShellType_proto_rawDescData = file_LuaShellType_proto_rawDesc
)

func file_LuaShellType_proto_rawDescGZIP() []byte {
	file_LuaShellType_proto_rawDescOnce.Do(func() {
		file_LuaShellType_proto_rawDescData = protoimpl.X.CompressGZIP(file_LuaShellType_proto_rawDescData)
	})
	return file_LuaShellType_proto_rawDescData
}

var file_LuaShellType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_LuaShellType_proto_goTypes = []interface{}{
	(LuaShellType)(0), // 0: LuaShellType
}
var file_LuaShellType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_LuaShellType_proto_init() }
func file_LuaShellType_proto_init() {
	if File_LuaShellType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_LuaShellType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LuaShellType_proto_goTypes,
		DependencyIndexes: file_LuaShellType_proto_depIdxs,
		EnumInfos:         file_LuaShellType_proto_enumTypes,
	}.Build()
	File_LuaShellType_proto = out.File
	file_LuaShellType_proto_rawDesc = nil
	file_LuaShellType_proto_goTypes = nil
	file_LuaShellType_proto_depIdxs = nil
}
