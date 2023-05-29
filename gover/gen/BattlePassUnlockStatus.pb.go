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
// source: BattlePassUnlockStatus.proto

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

// Obf: GIJOJHBPCLK
type BattlePassUnlockStatus int32

const (
	BattlePassUnlockStatus_BATTLE_PASS_UNLOCK_STATUS_INVALID BattlePassUnlockStatus = 0
	BattlePassUnlockStatus_BATTLE_PASS_UNLOCK_STATUS_FREE    BattlePassUnlockStatus = 1
	BattlePassUnlockStatus_BATTLE_PASS_UNLOCK_STATUS_PAID    BattlePassUnlockStatus = 2
)

// Enum value maps for BattlePassUnlockStatus.
var (
	BattlePassUnlockStatus_name = map[int32]string{
		0: "BATTLE_PASS_UNLOCK_STATUS_INVALID",
		1: "BATTLE_PASS_UNLOCK_STATUS_FREE",
		2: "BATTLE_PASS_UNLOCK_STATUS_PAID",
	}
	BattlePassUnlockStatus_value = map[string]int32{
		"BATTLE_PASS_UNLOCK_STATUS_INVALID": 0,
		"BATTLE_PASS_UNLOCK_STATUS_FREE":    1,
		"BATTLE_PASS_UNLOCK_STATUS_PAID":    2,
	}
)

func (x BattlePassUnlockStatus) Enum() *BattlePassUnlockStatus {
	p := new(BattlePassUnlockStatus)
	*p = x
	return p
}

func (x BattlePassUnlockStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BattlePassUnlockStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_BattlePassUnlockStatus_proto_enumTypes[0].Descriptor()
}

func (BattlePassUnlockStatus) Type() protoreflect.EnumType {
	return &file_BattlePassUnlockStatus_proto_enumTypes[0]
}

func (x BattlePassUnlockStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BattlePassUnlockStatus.Descriptor instead.
func (BattlePassUnlockStatus) EnumDescriptor() ([]byte, []int) {
	return file_BattlePassUnlockStatus_proto_rawDescGZIP(), []int{0}
}

var File_BattlePassUnlockStatus_proto protoreflect.FileDescriptor

var file_BattlePassUnlockStatus_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x55, 0x6e, 0x6c, 0x6f,
	0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x87,
	0x01, 0x0a, 0x16, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x55, 0x6e, 0x6c,
	0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x21, 0x42, 0x41, 0x54,
	0x54, 0x4c, 0x45, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x5f, 0x55, 0x4e, 0x4c, 0x4f, 0x43, 0x4b, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00,
	0x12, 0x22, 0x0a, 0x1e, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x5f,
	0x55, 0x4e, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x52,
	0x45, 0x45, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x50,
	0x41, 0x53, 0x53, 0x5f, 0x55, 0x4e, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x50, 0x41, 0x49, 0x44, 0x10, 0x02, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BattlePassUnlockStatus_proto_rawDescOnce sync.Once
	file_BattlePassUnlockStatus_proto_rawDescData = file_BattlePassUnlockStatus_proto_rawDesc
)

func file_BattlePassUnlockStatus_proto_rawDescGZIP() []byte {
	file_BattlePassUnlockStatus_proto_rawDescOnce.Do(func() {
		file_BattlePassUnlockStatus_proto_rawDescData = protoimpl.X.CompressGZIP(file_BattlePassUnlockStatus_proto_rawDescData)
	})
	return file_BattlePassUnlockStatus_proto_rawDescData
}

var file_BattlePassUnlockStatus_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_BattlePassUnlockStatus_proto_goTypes = []interface{}{
	(BattlePassUnlockStatus)(0), // 0: BattlePassUnlockStatus
}
var file_BattlePassUnlockStatus_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_BattlePassUnlockStatus_proto_init() }
func file_BattlePassUnlockStatus_proto_init() {
	if File_BattlePassUnlockStatus_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_BattlePassUnlockStatus_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BattlePassUnlockStatus_proto_goTypes,
		DependencyIndexes: file_BattlePassUnlockStatus_proto_depIdxs,
		EnumInfos:         file_BattlePassUnlockStatus_proto_enumTypes,
	}.Build()
	File_BattlePassUnlockStatus_proto = out.File
	file_BattlePassUnlockStatus_proto_rawDesc = nil
	file_BattlePassUnlockStatus_proto_goTypes = nil
	file_BattlePassUnlockStatus_proto_depIdxs = nil
}
