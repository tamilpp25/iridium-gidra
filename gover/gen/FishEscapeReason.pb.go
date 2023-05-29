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
// source: FishEscapeReason.proto

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

// Obf: DGDCCMHODEN
type FishEscapeReason int32

const (
	FishEscapeReason_FISN_ESCAPE_NONE    FishEscapeReason = 0
	FishEscapeReason_FISH_ESCAPE_SHOCKED FishEscapeReason = 1
	FishEscapeReason_FISH_ESCAPE_UNHOOK  FishEscapeReason = 2
)

// Enum value maps for FishEscapeReason.
var (
	FishEscapeReason_name = map[int32]string{
		0: "FISN_ESCAPE_NONE",
		1: "FISH_ESCAPE_SHOCKED",
		2: "FISH_ESCAPE_UNHOOK",
	}
	FishEscapeReason_value = map[string]int32{
		"FISN_ESCAPE_NONE":    0,
		"FISH_ESCAPE_SHOCKED": 1,
		"FISH_ESCAPE_UNHOOK":  2,
	}
)

func (x FishEscapeReason) Enum() *FishEscapeReason {
	p := new(FishEscapeReason)
	*p = x
	return p
}

func (x FishEscapeReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FishEscapeReason) Descriptor() protoreflect.EnumDescriptor {
	return file_FishEscapeReason_proto_enumTypes[0].Descriptor()
}

func (FishEscapeReason) Type() protoreflect.EnumType {
	return &file_FishEscapeReason_proto_enumTypes[0]
}

func (x FishEscapeReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FishEscapeReason.Descriptor instead.
func (FishEscapeReason) EnumDescriptor() ([]byte, []int) {
	return file_FishEscapeReason_proto_rawDescGZIP(), []int{0}
}

var File_FishEscapeReason_proto protoreflect.FileDescriptor

var file_FishEscapeReason_proto_rawDesc = []byte{
	0x0a, 0x16, 0x46, 0x69, 0x73, 0x68, 0x45, 0x73, 0x63, 0x61, 0x70, 0x65, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x59, 0x0a, 0x10, 0x46, 0x69, 0x73, 0x68,
	0x45, 0x73, 0x63, 0x61, 0x70, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x10,
	0x46, 0x49, 0x53, 0x4e, 0x5f, 0x45, 0x53, 0x43, 0x41, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x49, 0x53, 0x48, 0x5f, 0x45, 0x53, 0x43, 0x41, 0x50,
	0x45, 0x5f, 0x53, 0x48, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x46,
	0x49, 0x53, 0x48, 0x5f, 0x45, 0x53, 0x43, 0x41, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x48, 0x4f, 0x4f,
	0x4b, 0x10, 0x02, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_FishEscapeReason_proto_rawDescOnce sync.Once
	file_FishEscapeReason_proto_rawDescData = file_FishEscapeReason_proto_rawDesc
)

func file_FishEscapeReason_proto_rawDescGZIP() []byte {
	file_FishEscapeReason_proto_rawDescOnce.Do(func() {
		file_FishEscapeReason_proto_rawDescData = protoimpl.X.CompressGZIP(file_FishEscapeReason_proto_rawDescData)
	})
	return file_FishEscapeReason_proto_rawDescData
}

var file_FishEscapeReason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_FishEscapeReason_proto_goTypes = []interface{}{
	(FishEscapeReason)(0), // 0: FishEscapeReason
}
var file_FishEscapeReason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_FishEscapeReason_proto_init() }
func file_FishEscapeReason_proto_init() {
	if File_FishEscapeReason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_FishEscapeReason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FishEscapeReason_proto_goTypes,
		DependencyIndexes: file_FishEscapeReason_proto_depIdxs,
		EnumInfos:         file_FishEscapeReason_proto_enumTypes,
	}.Build()
	File_FishEscapeReason_proto = out.File
	file_FishEscapeReason_proto_rawDesc = nil
	file_FishEscapeReason_proto_goTypes = nil
	file_FishEscapeReason_proto_depIdxs = nil
}
