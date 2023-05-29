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
// source: DraftInviteFailReason.proto

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

// Obf: KGMFIEKKAPP
type DraftInviteFailReason int32

const (
	DraftInviteFailReason_DRAFT_FAIL_UNKNOWN           DraftInviteFailReason = 0
	DraftInviteFailReason_DRAFT_ACTIVITY_NOT_OPEN      DraftInviteFailReason = 1
	DraftInviteFailReason_DRAFT_ACTIVITY_PLAY_NOT_OPEN DraftInviteFailReason = 2
	DraftInviteFailReason_DRAFT_SCENE_NOT_MEET         DraftInviteFailReason = 3
	DraftInviteFailReason_DRAFT_WORLD_NOT_MEET         DraftInviteFailReason = 4
	DraftInviteFailReason_DRAFT_PLAY_LIMIT_NOT_MEET    DraftInviteFailReason = 5
)

// Enum value maps for DraftInviteFailReason.
var (
	DraftInviteFailReason_name = map[int32]string{
		0: "DRAFT_FAIL_UNKNOWN",
		1: "DRAFT_ACTIVITY_NOT_OPEN",
		2: "DRAFT_ACTIVITY_PLAY_NOT_OPEN",
		3: "DRAFT_SCENE_NOT_MEET",
		4: "DRAFT_WORLD_NOT_MEET",
		5: "DRAFT_PLAY_LIMIT_NOT_MEET",
	}
	DraftInviteFailReason_value = map[string]int32{
		"DRAFT_FAIL_UNKNOWN":           0,
		"DRAFT_ACTIVITY_NOT_OPEN":      1,
		"DRAFT_ACTIVITY_PLAY_NOT_OPEN": 2,
		"DRAFT_SCENE_NOT_MEET":         3,
		"DRAFT_WORLD_NOT_MEET":         4,
		"DRAFT_PLAY_LIMIT_NOT_MEET":    5,
	}
)

func (x DraftInviteFailReason) Enum() *DraftInviteFailReason {
	p := new(DraftInviteFailReason)
	*p = x
	return p
}

func (x DraftInviteFailReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DraftInviteFailReason) Descriptor() protoreflect.EnumDescriptor {
	return file_DraftInviteFailReason_proto_enumTypes[0].Descriptor()
}

func (DraftInviteFailReason) Type() protoreflect.EnumType {
	return &file_DraftInviteFailReason_proto_enumTypes[0]
}

func (x DraftInviteFailReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DraftInviteFailReason.Descriptor instead.
func (DraftInviteFailReason) EnumDescriptor() ([]byte, []int) {
	return file_DraftInviteFailReason_proto_rawDescGZIP(), []int{0}
}

var File_DraftInviteFailReason_proto protoreflect.FileDescriptor

var file_DraftInviteFailReason_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x44, 0x72, 0x61, 0x66, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x46, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xc1, 0x01,
	0x0a, 0x15, 0x44, 0x72, 0x61, 0x66, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x46, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x52, 0x41, 0x46, 0x54,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x1b, 0x0a, 0x17, 0x44, 0x52, 0x41, 0x46, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54,
	0x59, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c,
	0x44, 0x52, 0x41, 0x46, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x50,
	0x4c, 0x41, 0x59, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x18,
	0x0a, 0x14, 0x44, 0x52, 0x41, 0x46, 0x54, 0x5f, 0x53, 0x43, 0x45, 0x4e, 0x45, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x4d, 0x45, 0x45, 0x54, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x52, 0x41, 0x46,
	0x54, 0x5f, 0x57, 0x4f, 0x52, 0x4c, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x45, 0x45, 0x54,
	0x10, 0x04, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x52, 0x41, 0x46, 0x54, 0x5f, 0x50, 0x4c, 0x41, 0x59,
	0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x45, 0x45, 0x54, 0x10,
	0x05, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_DraftInviteFailReason_proto_rawDescOnce sync.Once
	file_DraftInviteFailReason_proto_rawDescData = file_DraftInviteFailReason_proto_rawDesc
)

func file_DraftInviteFailReason_proto_rawDescGZIP() []byte {
	file_DraftInviteFailReason_proto_rawDescOnce.Do(func() {
		file_DraftInviteFailReason_proto_rawDescData = protoimpl.X.CompressGZIP(file_DraftInviteFailReason_proto_rawDescData)
	})
	return file_DraftInviteFailReason_proto_rawDescData
}

var file_DraftInviteFailReason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_DraftInviteFailReason_proto_goTypes = []interface{}{
	(DraftInviteFailReason)(0), // 0: DraftInviteFailReason
}
var file_DraftInviteFailReason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_DraftInviteFailReason_proto_init() }
func file_DraftInviteFailReason_proto_init() {
	if File_DraftInviteFailReason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_DraftInviteFailReason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DraftInviteFailReason_proto_goTypes,
		DependencyIndexes: file_DraftInviteFailReason_proto_depIdxs,
		EnumInfos:         file_DraftInviteFailReason_proto_enumTypes,
	}.Build()
	File_DraftInviteFailReason_proto = out.File
	file_DraftInviteFailReason_proto_rawDesc = nil
	file_DraftInviteFailReason_proto_goTypes = nil
	file_DraftInviteFailReason_proto_depIdxs = nil
}
