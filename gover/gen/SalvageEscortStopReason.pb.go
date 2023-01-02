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
// source: SalvageEscortStopReason.proto

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

type SalvageEscortStopReason int32

const (
	SalvageEscortStopReason_SALVAGE_ESCORT_STOP_REASON_NONE      SalvageEscortStopReason = 0
	SalvageEscortStopReason_SALVAGE_ESCORT_STOP_REASON_SUCCESS   SalvageEscortStopReason = 1
	SalvageEscortStopReason_SALVAGE_ESCORT_STOP_REASON_DUMP      SalvageEscortStopReason = 2
	SalvageEscortStopReason_SALVAGE_ESCORT_STOP_REASON_TIME      SalvageEscortStopReason = 3
	SalvageEscortStopReason_SALVAGE_ESCORT_STOP_REASON_INTERRUPT SalvageEscortStopReason = 4
	SalvageEscortStopReason_SALVAGE_ESCORT_STOP_REASON_LEAVE     SalvageEscortStopReason = 5
	SalvageEscortStopReason_SALVAGE_ESCORT_STOP_REASON_FULL      SalvageEscortStopReason = 6
)

// Enum value maps for SalvageEscortStopReason.
var (
	SalvageEscortStopReason_name = map[int32]string{
		0: "SALVAGE_ESCORT_STOP_REASON_NONE",
		1: "SALVAGE_ESCORT_STOP_REASON_SUCCESS",
		2: "SALVAGE_ESCORT_STOP_REASON_DUMP",
		3: "SALVAGE_ESCORT_STOP_REASON_TIME",
		4: "SALVAGE_ESCORT_STOP_REASON_INTERRUPT",
		5: "SALVAGE_ESCORT_STOP_REASON_LEAVE",
		6: "SALVAGE_ESCORT_STOP_REASON_FULL",
	}
	SalvageEscortStopReason_value = map[string]int32{
		"SALVAGE_ESCORT_STOP_REASON_NONE":      0,
		"SALVAGE_ESCORT_STOP_REASON_SUCCESS":   1,
		"SALVAGE_ESCORT_STOP_REASON_DUMP":      2,
		"SALVAGE_ESCORT_STOP_REASON_TIME":      3,
		"SALVAGE_ESCORT_STOP_REASON_INTERRUPT": 4,
		"SALVAGE_ESCORT_STOP_REASON_LEAVE":     5,
		"SALVAGE_ESCORT_STOP_REASON_FULL":      6,
	}
)

func (x SalvageEscortStopReason) Enum() *SalvageEscortStopReason {
	p := new(SalvageEscortStopReason)
	*p = x
	return p
}

func (x SalvageEscortStopReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SalvageEscortStopReason) Descriptor() protoreflect.EnumDescriptor {
	return file_SalvageEscortStopReason_proto_enumTypes[0].Descriptor()
}

func (SalvageEscortStopReason) Type() protoreflect.EnumType {
	return &file_SalvageEscortStopReason_proto_enumTypes[0]
}

func (x SalvageEscortStopReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SalvageEscortStopReason.Descriptor instead.
func (SalvageEscortStopReason) EnumDescriptor() ([]byte, []int) {
	return file_SalvageEscortStopReason_proto_rawDescGZIP(), []int{0}
}

var File_SalvageEscortStopReason_proto protoreflect.FileDescriptor

var file_SalvageEscortStopReason_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x53, 0x61, 0x6c, 0x76, 0x61, 0x67, 0x65, 0x45, 0x73, 0x63, 0x6f, 0x72, 0x74, 0x53,
	0x74, 0x6f, 0x70, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a,
	0xa5, 0x02, 0x0a, 0x17, 0x53, 0x61, 0x6c, 0x76, 0x61, 0x67, 0x65, 0x45, 0x73, 0x63, 0x6f, 0x72,
	0x74, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x1f, 0x53,
	0x41, 0x4c, 0x56, 0x41, 0x47, 0x45, 0x5f, 0x45, 0x53, 0x43, 0x4f, 0x52, 0x54, 0x5f, 0x53, 0x54,
	0x4f, 0x50, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x26, 0x0a, 0x22, 0x53, 0x41, 0x4c, 0x56, 0x41, 0x47, 0x45, 0x5f, 0x45, 0x53, 0x43, 0x4f,
	0x52, 0x54, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x53,
	0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x23, 0x0a, 0x1f, 0x53, 0x41, 0x4c, 0x56,
	0x41, 0x47, 0x45, 0x5f, 0x45, 0x53, 0x43, 0x4f, 0x52, 0x54, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x5f,
	0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x44, 0x55, 0x4d, 0x50, 0x10, 0x02, 0x12, 0x23, 0x0a,
	0x1f, 0x53, 0x41, 0x4c, 0x56, 0x41, 0x47, 0x45, 0x5f, 0x45, 0x53, 0x43, 0x4f, 0x52, 0x54, 0x5f,
	0x53, 0x54, 0x4f, 0x50, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x54, 0x49, 0x4d, 0x45,
	0x10, 0x03, 0x12, 0x28, 0x0a, 0x24, 0x53, 0x41, 0x4c, 0x56, 0x41, 0x47, 0x45, 0x5f, 0x45, 0x53,
	0x43, 0x4f, 0x52, 0x54, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e,
	0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x52, 0x55, 0x50, 0x54, 0x10, 0x04, 0x12, 0x24, 0x0a, 0x20,
	0x53, 0x41, 0x4c, 0x56, 0x41, 0x47, 0x45, 0x5f, 0x45, 0x53, 0x43, 0x4f, 0x52, 0x54, 0x5f, 0x53,
	0x54, 0x4f, 0x50, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x4c, 0x45, 0x41, 0x56, 0x45,
	0x10, 0x05, 0x12, 0x23, 0x0a, 0x1f, 0x53, 0x41, 0x4c, 0x56, 0x41, 0x47, 0x45, 0x5f, 0x45, 0x53,
	0x43, 0x4f, 0x52, 0x54, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e,
	0x5f, 0x46, 0x55, 0x4c, 0x4c, 0x10, 0x06, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SalvageEscortStopReason_proto_rawDescOnce sync.Once
	file_SalvageEscortStopReason_proto_rawDescData = file_SalvageEscortStopReason_proto_rawDesc
)

func file_SalvageEscortStopReason_proto_rawDescGZIP() []byte {
	file_SalvageEscortStopReason_proto_rawDescOnce.Do(func() {
		file_SalvageEscortStopReason_proto_rawDescData = protoimpl.X.CompressGZIP(file_SalvageEscortStopReason_proto_rawDescData)
	})
	return file_SalvageEscortStopReason_proto_rawDescData
}

var file_SalvageEscortStopReason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_SalvageEscortStopReason_proto_goTypes = []interface{}{
	(SalvageEscortStopReason)(0), // 0: SalvageEscortStopReason
}
var file_SalvageEscortStopReason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SalvageEscortStopReason_proto_init() }
func file_SalvageEscortStopReason_proto_init() {
	if File_SalvageEscortStopReason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_SalvageEscortStopReason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SalvageEscortStopReason_proto_goTypes,
		DependencyIndexes: file_SalvageEscortStopReason_proto_depIdxs,
		EnumInfos:         file_SalvageEscortStopReason_proto_enumTypes,
	}.Build()
	File_SalvageEscortStopReason_proto = out.File
	file_SalvageEscortStopReason_proto_rawDesc = nil
	file_SalvageEscortStopReason_proto_goTypes = nil
	file_SalvageEscortStopReason_proto_depIdxs = nil
}
