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
// source: GCGActionType.proto

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

type GCGActionType int32

const (
	GCGActionType_GCG_ACTION_TYPE_NONE                  GCGActionType = 0
	GCGActionType_GCG_ACTION_TYPE_SPECIAL_PHASE         GCGActionType = 1
	GCGActionType_GCG_ACTION_TYPE_NEXT_PHASE            GCGActionType = 2
	GCGActionType_GCG_ACTION_TYPE_DRAW                  GCGActionType = 3
	GCGActionType_GCG_ACTION_TYPE_REDRAW                GCGActionType = 4
	GCGActionType_GCG_ACTION_TYPE_SELECT_ONSTAGE        GCGActionType = 5
	GCGActionType_GCG_ACTION_TYPE_ROLL                  GCGActionType = 6
	GCGActionType_GCG_ACTION_TYPE_REROLL                GCGActionType = 7
	GCGActionType_GCG_ACTION_TYPE_ATTACK                GCGActionType = 8
	GCGActionType_GCG_ACTION_TYPE_PLAY_CARD             GCGActionType = 9
	GCGActionType_GCG_ACTION_TYPE_PASS                  GCGActionType = 10
	GCGActionType_GCG_ACTION_TYPE_REBOOT                GCGActionType = 11
	GCGActionType_GCG_ACTION_TYPE_GAME_OVER             GCGActionType = 12
	GCGActionType_GCG_ACTION_TYPE_TRIGGER               GCGActionType = 13
	GCGActionType_GCG_ACTION_TYPE_PHASE_EXIT            GCGActionType = 14
	GCGActionType_GCG_ACTION_TYPE_CUSTOM                GCGActionType = 15
	GCGActionType_GCG_ACTION_TYPE_NOTIFY_COST           GCGActionType = 16
	GCGActionType_GCG_ACTION_TYPE_AFTER_OPERATION       GCGActionType = 17
	GCGActionType_GCG_ACTION_TYPE_USE_SKILL             GCGActionType = 18
	GCGActionType_GCG_ACTION_TYPE_NOTIFY_SKILL_PREVIEW  GCGActionType = 19
	GCGActionType_GCG_ACTION_TYPE_PREVIEW_ATTACK        GCGActionType = 20
	GCGActionType_GCG_ACTION_TYPE_PREVIEW_AFTER_ATTACK  GCGActionType = 21
	GCGActionType_GCG_ACTION_TYPE_SEND_MESSAGE          GCGActionType = 22
	GCGActionType_GCG_ACTION_TYPE_WAITING_CHARACTER     GCGActionType = 23
	GCGActionType_GCG_ACTION_TYPE_TRIGGER_SKILL         GCGActionType = 24
	GCGActionType_GCG_ACTION_TYPE_BEFORE_NEXT_OPERATION GCGActionType = 25
)

// Enum value maps for GCGActionType.
var (
	GCGActionType_name = map[int32]string{
		0:  "GCG_ACTION_TYPE_NONE",
		1:  "GCG_ACTION_TYPE_SPECIAL_PHASE",
		2:  "GCG_ACTION_TYPE_NEXT_PHASE",
		3:  "GCG_ACTION_TYPE_DRAW",
		4:  "GCG_ACTION_TYPE_REDRAW",
		5:  "GCG_ACTION_TYPE_SELECT_ONSTAGE",
		6:  "GCG_ACTION_TYPE_ROLL",
		7:  "GCG_ACTION_TYPE_REROLL",
		8:  "GCG_ACTION_TYPE_ATTACK",
		9:  "GCG_ACTION_TYPE_PLAY_CARD",
		10: "GCG_ACTION_TYPE_PASS",
		11: "GCG_ACTION_TYPE_REBOOT",
		12: "GCG_ACTION_TYPE_GAME_OVER",
		13: "GCG_ACTION_TYPE_TRIGGER",
		14: "GCG_ACTION_TYPE_PHASE_EXIT",
		15: "GCG_ACTION_TYPE_CUSTOM",
		16: "GCG_ACTION_TYPE_NOTIFY_COST",
		17: "GCG_ACTION_TYPE_AFTER_OPERATION",
		18: "GCG_ACTION_TYPE_USE_SKILL",
		19: "GCG_ACTION_TYPE_NOTIFY_SKILL_PREVIEW",
		20: "GCG_ACTION_TYPE_PREVIEW_ATTACK",
		21: "GCG_ACTION_TYPE_PREVIEW_AFTER_ATTACK",
		22: "GCG_ACTION_TYPE_SEND_MESSAGE",
		23: "GCG_ACTION_TYPE_WAITING_CHARACTER",
		24: "GCG_ACTION_TYPE_TRIGGER_SKILL",
		25: "GCG_ACTION_TYPE_BEFORE_NEXT_OPERATION",
	}
	GCGActionType_value = map[string]int32{
		"GCG_ACTION_TYPE_NONE":                  0,
		"GCG_ACTION_TYPE_SPECIAL_PHASE":         1,
		"GCG_ACTION_TYPE_NEXT_PHASE":            2,
		"GCG_ACTION_TYPE_DRAW":                  3,
		"GCG_ACTION_TYPE_REDRAW":                4,
		"GCG_ACTION_TYPE_SELECT_ONSTAGE":        5,
		"GCG_ACTION_TYPE_ROLL":                  6,
		"GCG_ACTION_TYPE_REROLL":                7,
		"GCG_ACTION_TYPE_ATTACK":                8,
		"GCG_ACTION_TYPE_PLAY_CARD":             9,
		"GCG_ACTION_TYPE_PASS":                  10,
		"GCG_ACTION_TYPE_REBOOT":                11,
		"GCG_ACTION_TYPE_GAME_OVER":             12,
		"GCG_ACTION_TYPE_TRIGGER":               13,
		"GCG_ACTION_TYPE_PHASE_EXIT":            14,
		"GCG_ACTION_TYPE_CUSTOM":                15,
		"GCG_ACTION_TYPE_NOTIFY_COST":           16,
		"GCG_ACTION_TYPE_AFTER_OPERATION":       17,
		"GCG_ACTION_TYPE_USE_SKILL":             18,
		"GCG_ACTION_TYPE_NOTIFY_SKILL_PREVIEW":  19,
		"GCG_ACTION_TYPE_PREVIEW_ATTACK":        20,
		"GCG_ACTION_TYPE_PREVIEW_AFTER_ATTACK":  21,
		"GCG_ACTION_TYPE_SEND_MESSAGE":          22,
		"GCG_ACTION_TYPE_WAITING_CHARACTER":     23,
		"GCG_ACTION_TYPE_TRIGGER_SKILL":         24,
		"GCG_ACTION_TYPE_BEFORE_NEXT_OPERATION": 25,
	}
)

func (x GCGActionType) Enum() *GCGActionType {
	p := new(GCGActionType)
	*p = x
	return p
}

func (x GCGActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GCGActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_GCGActionType_proto_enumTypes[0].Descriptor()
}

func (GCGActionType) Type() protoreflect.EnumType {
	return &file_GCGActionType_proto_enumTypes[0]
}

func (x GCGActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GCGActionType.Descriptor instead.
func (GCGActionType) EnumDescriptor() ([]byte, []int) {
	return file_GCGActionType_proto_rawDescGZIP(), []int{0}
}

var File_GCGActionType_proto protoreflect.FileDescriptor

var file_GCGActionType_proto_rawDesc = []byte{
	0x0a, 0x13, 0x47, 0x43, 0x47, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xd9, 0x06, 0x0a, 0x0d, 0x47, 0x43, 0x47, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x47, 0x43, 0x47, 0x5f, 0x41,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10,
	0x00, 0x12, 0x21, 0x0a, 0x1d, 0x47, 0x43, 0x47, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x41, 0x4c, 0x5f, 0x50, 0x48, 0x41,
	0x53, 0x45, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x47, 0x43, 0x47, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x45, 0x58, 0x54, 0x5f, 0x50, 0x48, 0x41,
	0x53, 0x45, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x47, 0x43, 0x47, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x52, 0x41, 0x57, 0x10, 0x03, 0x12, 0x1a,
	0x0a, 0x16, 0x47, 0x43, 0x47, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x52, 0x45, 0x44, 0x52, 0x41, 0x57, 0x10, 0x04, 0x12, 0x22, 0x0a, 0x1e, 0x47, 0x43,
	0x47, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x45,
	0x4c, 0x45, 0x43, 0x54, 0x5f, 0x4f, 0x4e, 0x53, 0x54, 0x41, 0x47, 0x45, 0x10, 0x05, 0x12, 0x18,
	0x0a, 0x14, 0x47, 0x43, 0x47, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x52, 0x4f, 0x4c, 0x4c, 0x10, 0x06, 0x12, 0x1a, 0x0a, 0x16, 0x47, 0x43, 0x47, 0x5f,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x52, 0x4f,
	0x4c, 0x4c, 0x10, 0x07, 0x12, 0x1a, 0x0a, 0x16, 0x47, 0x43, 0x47, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x54, 0x54, 0x41, 0x43, 0x4b, 0x10, 0x08,
	0x12, 0x1d, 0x0a, 0x19, 0x47, 0x43, 0x47, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0x09, 0x12,
	0x18, 0x0a, 0x14, 0x47, 0x43, 0x47, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x10, 0x0a, 0x12, 0x1a, 0x0a, 0x16, 0x47, 0x43, 0x47,
	0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x42,
	0x4f, 0x4f, 0x54, 0x10, 0x0b, 0x12, 0x1d, 0x0a, 0x19, 0x47, 0x43, 0x47, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x4f, 0x56,
	0x45, 0x52, 0x10, 0x0c, 0x12, 0x1b, 0x0a, 0x17, 0x47, 0x43, 0x47, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x49, 0x47, 0x47, 0x45, 0x52, 0x10,
	0x0d, 0x12, 0x1e, 0x0a, 0x1a, 0x47, 0x43, 0x47, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x45, 0x58, 0x49, 0x54, 0x10,
	0x0e, 0x12, 0x1a, 0x0a, 0x16, 0x47, 0x43, 0x47, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x10, 0x0f, 0x12, 0x1f, 0x0a,
	0x1b, 0x47, 0x43, 0x47, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x43, 0x4f, 0x53, 0x54, 0x10, 0x10, 0x12, 0x23,
	0x0a, 0x1f, 0x47, 0x43, 0x47, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x41, 0x46, 0x54, 0x45, 0x52, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x10, 0x11, 0x12, 0x1d, 0x0a, 0x19, 0x47, 0x43, 0x47, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x5f, 0x53, 0x4b, 0x49, 0x4c, 0x4c,
	0x10, 0x12, 0x12, 0x28, 0x0a, 0x24, 0x47, 0x43, 0x47, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x53, 0x4b, 0x49,
	0x4c, 0x4c, 0x5f, 0x50, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x13, 0x12, 0x22, 0x0a, 0x1e,
	0x47, 0x43, 0x47, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x50, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x41, 0x54, 0x54, 0x41, 0x43, 0x4b, 0x10, 0x14,
	0x12, 0x28, 0x0a, 0x24, 0x47, 0x43, 0x47, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x50, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x41, 0x46, 0x54, 0x45,
	0x52, 0x5f, 0x41, 0x54, 0x54, 0x41, 0x43, 0x4b, 0x10, 0x15, 0x12, 0x20, 0x0a, 0x1c, 0x47, 0x43,
	0x47, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x45,
	0x4e, 0x44, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x16, 0x12, 0x25, 0x0a, 0x21,
	0x47, 0x43, 0x47, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x48, 0x41, 0x52, 0x41, 0x43, 0x54, 0x45,
	0x52, 0x10, 0x17, 0x12, 0x21, 0x0a, 0x1d, 0x47, 0x43, 0x47, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x49, 0x47, 0x47, 0x45, 0x52, 0x5f, 0x53,
	0x4b, 0x49, 0x4c, 0x4c, 0x10, 0x18, 0x12, 0x29, 0x0a, 0x25, 0x47, 0x43, 0x47, 0x5f, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x45, 0x46, 0x4f, 0x52, 0x45,
	0x5f, 0x4e, 0x45, 0x58, 0x54, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x19, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_GCGActionType_proto_rawDescOnce sync.Once
	file_GCGActionType_proto_rawDescData = file_GCGActionType_proto_rawDesc
)

func file_GCGActionType_proto_rawDescGZIP() []byte {
	file_GCGActionType_proto_rawDescOnce.Do(func() {
		file_GCGActionType_proto_rawDescData = protoimpl.X.CompressGZIP(file_GCGActionType_proto_rawDescData)
	})
	return file_GCGActionType_proto_rawDescData
}

var file_GCGActionType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GCGActionType_proto_goTypes = []interface{}{
	(GCGActionType)(0), // 0: GCGActionType
}
var file_GCGActionType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GCGActionType_proto_init() }
func file_GCGActionType_proto_init() {
	if File_GCGActionType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_GCGActionType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GCGActionType_proto_goTypes,
		DependencyIndexes: file_GCGActionType_proto_depIdxs,
		EnumInfos:         file_GCGActionType_proto_enumTypes,
	}.Build()
	File_GCGActionType_proto = out.File
	file_GCGActionType_proto_rawDesc = nil
	file_GCGActionType_proto_goTypes = nil
	file_GCGActionType_proto_depIdxs = nil
}
