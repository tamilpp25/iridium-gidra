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
// source: CADIFCLFOFF.proto

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

type CADIFCLFOFF int32

const (
	CADIFCLFOFF_CADIFCLFOFF_None     CADIFCLFOFF = 0
	CADIFCLFOFF_CADIFCLFOFF_Starred  CADIFCLFOFF = 1
	CADIFCLFOFF_CADIFCLFOFF_Official CADIFCLFOFF = 2
	CADIFCLFOFF_CADIFCLFOFF_Template CADIFCLFOFF = 3
)

// Enum value maps for CADIFCLFOFF.
var (
	CADIFCLFOFF_name = map[int32]string{
		0: "CADIFCLFOFF_None",
		1: "CADIFCLFOFF_Starred",
		2: "CADIFCLFOFF_Official",
		3: "CADIFCLFOFF_Template",
	}
	CADIFCLFOFF_value = map[string]int32{
		"CADIFCLFOFF_None":     0,
		"CADIFCLFOFF_Starred":  1,
		"CADIFCLFOFF_Official": 2,
		"CADIFCLFOFF_Template": 3,
	}
)

func (x CADIFCLFOFF) Enum() *CADIFCLFOFF {
	p := new(CADIFCLFOFF)
	*p = x
	return p
}

func (x CADIFCLFOFF) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CADIFCLFOFF) Descriptor() protoreflect.EnumDescriptor {
	return file_CADIFCLFOFF_proto_enumTypes[0].Descriptor()
}

func (CADIFCLFOFF) Type() protoreflect.EnumType {
	return &file_CADIFCLFOFF_proto_enumTypes[0]
}

func (x CADIFCLFOFF) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CADIFCLFOFF.Descriptor instead.
func (CADIFCLFOFF) EnumDescriptor() ([]byte, []int) {
	return file_CADIFCLFOFF_proto_rawDescGZIP(), []int{0}
}

var File_CADIFCLFOFF_proto protoreflect.FileDescriptor

var file_CADIFCLFOFF_proto_rawDesc = []byte{
	0x0a, 0x11, 0x43, 0x41, 0x44, 0x49, 0x46, 0x43, 0x4c, 0x46, 0x4f, 0x46, 0x46, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0x70, 0x0a, 0x0b, 0x43, 0x41, 0x44, 0x49, 0x46, 0x43, 0x4c, 0x46, 0x4f,
	0x46, 0x46, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x41, 0x44, 0x49, 0x46, 0x43, 0x4c, 0x46, 0x4f, 0x46,
	0x46, 0x5f, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x41, 0x44, 0x49,
	0x46, 0x43, 0x4c, 0x46, 0x4f, 0x46, 0x46, 0x5f, 0x53, 0x74, 0x61, 0x72, 0x72, 0x65, 0x64, 0x10,
	0x01, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x41, 0x44, 0x49, 0x46, 0x43, 0x4c, 0x46, 0x4f, 0x46, 0x46,
	0x5f, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x43,
	0x41, 0x44, 0x49, 0x46, 0x43, 0x4c, 0x46, 0x4f, 0x46, 0x46, 0x5f, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x10, 0x03, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CADIFCLFOFF_proto_rawDescOnce sync.Once
	file_CADIFCLFOFF_proto_rawDescData = file_CADIFCLFOFF_proto_rawDesc
)

func file_CADIFCLFOFF_proto_rawDescGZIP() []byte {
	file_CADIFCLFOFF_proto_rawDescOnce.Do(func() {
		file_CADIFCLFOFF_proto_rawDescData = protoimpl.X.CompressGZIP(file_CADIFCLFOFF_proto_rawDescData)
	})
	return file_CADIFCLFOFF_proto_rawDescData
}

var file_CADIFCLFOFF_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_CADIFCLFOFF_proto_goTypes = []interface{}{
	(CADIFCLFOFF)(0), // 0: CADIFCLFOFF
}
var file_CADIFCLFOFF_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_CADIFCLFOFF_proto_init() }
func file_CADIFCLFOFF_proto_init() {
	if File_CADIFCLFOFF_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_CADIFCLFOFF_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CADIFCLFOFF_proto_goTypes,
		DependencyIndexes: file_CADIFCLFOFF_proto_depIdxs,
		EnumInfos:         file_CADIFCLFOFF_proto_enumTypes,
	}.Build()
	File_CADIFCLFOFF_proto = out.File
	file_CADIFCLFOFF_proto_rawDesc = nil
	file_CADIFCLFOFF_proto_goTypes = nil
	file_CADIFCLFOFF_proto_depIdxs = nil
}
