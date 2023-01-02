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
// source: ModifierAction.proto

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

type ModifierAction int32

const (
	ModifierAction_MODIFIER_ACTION_ADDED   ModifierAction = 0
	ModifierAction_MODIFIER_ACTION_REMOVED ModifierAction = 1
)

// Enum value maps for ModifierAction.
var (
	ModifierAction_name = map[int32]string{
		0: "MODIFIER_ACTION_ADDED",
		1: "MODIFIER_ACTION_REMOVED",
	}
	ModifierAction_value = map[string]int32{
		"MODIFIER_ACTION_ADDED":   0,
		"MODIFIER_ACTION_REMOVED": 1,
	}
)

func (x ModifierAction) Enum() *ModifierAction {
	p := new(ModifierAction)
	*p = x
	return p
}

func (x ModifierAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ModifierAction) Descriptor() protoreflect.EnumDescriptor {
	return file_ModifierAction_proto_enumTypes[0].Descriptor()
}

func (ModifierAction) Type() protoreflect.EnumType {
	return &file_ModifierAction_proto_enumTypes[0]
}

func (x ModifierAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ModifierAction.Descriptor instead.
func (ModifierAction) EnumDescriptor() ([]byte, []int) {
	return file_ModifierAction_proto_rawDescGZIP(), []int{0}
}

var File_ModifierAction_proto protoreflect.FileDescriptor

var file_ModifierAction_proto_rawDesc = []byte{
	0x0a, 0x14, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x48, 0x0a, 0x0e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x4f, 0x44, 0x49,
	0x46, 0x49, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x44, 0x44, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x49, 0x45, 0x52, 0x5f,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x01,
	0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ModifierAction_proto_rawDescOnce sync.Once
	file_ModifierAction_proto_rawDescData = file_ModifierAction_proto_rawDesc
)

func file_ModifierAction_proto_rawDescGZIP() []byte {
	file_ModifierAction_proto_rawDescOnce.Do(func() {
		file_ModifierAction_proto_rawDescData = protoimpl.X.CompressGZIP(file_ModifierAction_proto_rawDescData)
	})
	return file_ModifierAction_proto_rawDescData
}

var file_ModifierAction_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ModifierAction_proto_goTypes = []interface{}{
	(ModifierAction)(0), // 0: ModifierAction
}
var file_ModifierAction_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ModifierAction_proto_init() }
func file_ModifierAction_proto_init() {
	if File_ModifierAction_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ModifierAction_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ModifierAction_proto_goTypes,
		DependencyIndexes: file_ModifierAction_proto_depIdxs,
		EnumInfos:         file_ModifierAction_proto_enumTypes,
	}.Build()
	File_ModifierAction_proto = out.File
	file_ModifierAction_proto_rawDesc = nil
	file_ModifierAction_proto_goTypes = nil
	file_ModifierAction_proto_depIdxs = nil
}
