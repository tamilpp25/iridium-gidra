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
// source: GCGReason.proto

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

type GCGReason int32

const (
	GCGReason_GCG_REASON_DEFAULT          GCGReason = 0
	GCGReason_GCG_REASON_EFFECT           GCGReason = 1
	GCGReason_GCG_REASON_COST             GCGReason = 2
	GCGReason_GCG_REASON_GM               GCGReason = 3
	GCGReason_GCG_REASON_ATTACK           GCGReason = 4
	GCGReason_GCG_REASON_REBOOT           GCGReason = 5
	GCGReason_GCG_REASON_PLAY_CARD        GCGReason = 6
	GCGReason_GCG_REASON_QUICKLY_ONSTAGE  GCGReason = 7
	GCGReason_GCG_REASON_REMOVE_AFTER_DIE GCGReason = 8
	GCGReason_GCG_REASON_INIT             GCGReason = 9
	GCGReason_GCG_REASON_EFFECT_DAMAGE    GCGReason = 10
	GCGReason_GCG_REASON_EFFECT_HEAL      GCGReason = 11
	GCGReason_GCG_REASON_EFFECT_REVIVE    GCGReason = 12
)

// Enum value maps for GCGReason.
var (
	GCGReason_name = map[int32]string{
		0:  "GCG_REASON_DEFAULT",
		1:  "GCG_REASON_EFFECT",
		2:  "GCG_REASON_COST",
		3:  "GCG_REASON_GM",
		4:  "GCG_REASON_ATTACK",
		5:  "GCG_REASON_REBOOT",
		6:  "GCG_REASON_PLAY_CARD",
		7:  "GCG_REASON_QUICKLY_ONSTAGE",
		8:  "GCG_REASON_REMOVE_AFTER_DIE",
		9:  "GCG_REASON_INIT",
		10: "GCG_REASON_EFFECT_DAMAGE",
		11: "GCG_REASON_EFFECT_HEAL",
		12: "GCG_REASON_EFFECT_REVIVE",
	}
	GCGReason_value = map[string]int32{
		"GCG_REASON_DEFAULT":          0,
		"GCG_REASON_EFFECT":           1,
		"GCG_REASON_COST":             2,
		"GCG_REASON_GM":               3,
		"GCG_REASON_ATTACK":           4,
		"GCG_REASON_REBOOT":           5,
		"GCG_REASON_PLAY_CARD":        6,
		"GCG_REASON_QUICKLY_ONSTAGE":  7,
		"GCG_REASON_REMOVE_AFTER_DIE": 8,
		"GCG_REASON_INIT":             9,
		"GCG_REASON_EFFECT_DAMAGE":    10,
		"GCG_REASON_EFFECT_HEAL":      11,
		"GCG_REASON_EFFECT_REVIVE":    12,
	}
)

func (x GCGReason) Enum() *GCGReason {
	p := new(GCGReason)
	*p = x
	return p
}

func (x GCGReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GCGReason) Descriptor() protoreflect.EnumDescriptor {
	return file_GCGReason_proto_enumTypes[0].Descriptor()
}

func (GCGReason) Type() protoreflect.EnumType {
	return &file_GCGReason_proto_enumTypes[0]
}

func (x GCGReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GCGReason.Descriptor instead.
func (GCGReason) EnumDescriptor() ([]byte, []int) {
	return file_GCGReason_proto_rawDescGZIP(), []int{0}
}

var File_GCGReason_proto protoreflect.FileDescriptor

var file_GCGReason_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x47, 0x43, 0x47, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2a, 0xd8, 0x02, 0x0a, 0x09, 0x47, 0x43, 0x47, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x12, 0x47, 0x43, 0x47, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x44, 0x45,
	0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x47, 0x43, 0x47, 0x5f, 0x52,
	0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x45, 0x46, 0x46, 0x45, 0x43, 0x54, 0x10, 0x01, 0x12, 0x13,
	0x0a, 0x0f, 0x47, 0x43, 0x47, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x53,
	0x54, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x47, 0x43, 0x47, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f,
	0x4e, 0x5f, 0x47, 0x4d, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x47, 0x43, 0x47, 0x5f, 0x52, 0x45,
	0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x41, 0x54, 0x54, 0x41, 0x43, 0x4b, 0x10, 0x04, 0x12, 0x15, 0x0a,
	0x11, 0x47, 0x43, 0x47, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x42, 0x4f,
	0x4f, 0x54, 0x10, 0x05, 0x12, 0x18, 0x0a, 0x14, 0x47, 0x43, 0x47, 0x5f, 0x52, 0x45, 0x41, 0x53,
	0x4f, 0x4e, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0x06, 0x12, 0x1e,
	0x0a, 0x1a, 0x47, 0x43, 0x47, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x51, 0x55, 0x49,
	0x43, 0x4b, 0x4c, 0x59, 0x5f, 0x4f, 0x4e, 0x53, 0x54, 0x41, 0x47, 0x45, 0x10, 0x07, 0x12, 0x1f,
	0x0a, 0x1b, 0x47, 0x43, 0x47, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x4d,
	0x4f, 0x56, 0x45, 0x5f, 0x41, 0x46, 0x54, 0x45, 0x52, 0x5f, 0x44, 0x49, 0x45, 0x10, 0x08, 0x12,
	0x13, 0x0a, 0x0f, 0x47, 0x43, 0x47, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x49, 0x4e,
	0x49, 0x54, 0x10, 0x09, 0x12, 0x1c, 0x0a, 0x18, 0x47, 0x43, 0x47, 0x5f, 0x52, 0x45, 0x41, 0x53,
	0x4f, 0x4e, 0x5f, 0x45, 0x46, 0x46, 0x45, 0x43, 0x54, 0x5f, 0x44, 0x41, 0x4d, 0x41, 0x47, 0x45,
	0x10, 0x0a, 0x12, 0x1a, 0x0a, 0x16, 0x47, 0x43, 0x47, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e,
	0x5f, 0x45, 0x46, 0x46, 0x45, 0x43, 0x54, 0x5f, 0x48, 0x45, 0x41, 0x4c, 0x10, 0x0b, 0x12, 0x1c,
	0x0a, 0x18, 0x47, 0x43, 0x47, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x45, 0x46, 0x46,
	0x45, 0x43, 0x54, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x56, 0x45, 0x10, 0x0c, 0x42, 0x06, 0x5a, 0x04,
	0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GCGReason_proto_rawDescOnce sync.Once
	file_GCGReason_proto_rawDescData = file_GCGReason_proto_rawDesc
)

func file_GCGReason_proto_rawDescGZIP() []byte {
	file_GCGReason_proto_rawDescOnce.Do(func() {
		file_GCGReason_proto_rawDescData = protoimpl.X.CompressGZIP(file_GCGReason_proto_rawDescData)
	})
	return file_GCGReason_proto_rawDescData
}

var file_GCGReason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GCGReason_proto_goTypes = []interface{}{
	(GCGReason)(0), // 0: GCGReason
}
var file_GCGReason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GCGReason_proto_init() }
func file_GCGReason_proto_init() {
	if File_GCGReason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_GCGReason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GCGReason_proto_goTypes,
		DependencyIndexes: file_GCGReason_proto_depIdxs,
		EnumInfos:         file_GCGReason_proto_enumTypes,
	}.Build()
	File_GCGReason_proto = out.File
	file_GCGReason_proto_rawDesc = nil
	file_GCGReason_proto_goTypes = nil
	file_GCGReason_proto_depIdxs = nil
}
