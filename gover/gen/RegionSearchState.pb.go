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
// source: RegionSearchState.proto

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

// Obf: JEONFIOIHJC
type RegionSearchState int32

const (
	RegionSearchState_REGION_SEARCH_NONE        RegionSearchState = 0
	RegionSearchState_REGION_SEARCH_UNSTARTED   RegionSearchState = 1
	RegionSearchState_REGION_SEARCH_STARTED     RegionSearchState = 2
	RegionSearchState_REGION_SEARCH_WAIT_REWARD RegionSearchState = 3
	RegionSearchState_REGION_SEARCH_FINISHED    RegionSearchState = 4
)

// Enum value maps for RegionSearchState.
var (
	RegionSearchState_name = map[int32]string{
		0: "REGION_SEARCH_NONE",
		1: "REGION_SEARCH_UNSTARTED",
		2: "REGION_SEARCH_STARTED",
		3: "REGION_SEARCH_WAIT_REWARD",
		4: "REGION_SEARCH_FINISHED",
	}
	RegionSearchState_value = map[string]int32{
		"REGION_SEARCH_NONE":        0,
		"REGION_SEARCH_UNSTARTED":   1,
		"REGION_SEARCH_STARTED":     2,
		"REGION_SEARCH_WAIT_REWARD": 3,
		"REGION_SEARCH_FINISHED":    4,
	}
)

func (x RegionSearchState) Enum() *RegionSearchState {
	p := new(RegionSearchState)
	*p = x
	return p
}

func (x RegionSearchState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RegionSearchState) Descriptor() protoreflect.EnumDescriptor {
	return file_RegionSearchState_proto_enumTypes[0].Descriptor()
}

func (RegionSearchState) Type() protoreflect.EnumType {
	return &file_RegionSearchState_proto_enumTypes[0]
}

func (x RegionSearchState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RegionSearchState.Descriptor instead.
func (RegionSearchState) EnumDescriptor() ([]byte, []int) {
	return file_RegionSearchState_proto_rawDescGZIP(), []int{0}
}

var File_RegionSearchState_proto protoreflect.FileDescriptor

var file_RegionSearchState_proto_rawDesc = []byte{
	0x0a, 0x17, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x9e, 0x01, 0x0a, 0x11, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x16, 0x0a, 0x12, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48,
	0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x45, 0x47, 0x49, 0x4f,
	0x4e, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x55, 0x4e, 0x53, 0x54, 0x41, 0x52, 0x54,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x53,
	0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12,
	0x1d, 0x0a, 0x19, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48,
	0x5f, 0x57, 0x41, 0x49, 0x54, 0x5f, 0x52, 0x45, 0x57, 0x41, 0x52, 0x44, 0x10, 0x03, 0x12, 0x1a,
	0x0a, 0x16, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f,
	0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x04, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67,
	0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RegionSearchState_proto_rawDescOnce sync.Once
	file_RegionSearchState_proto_rawDescData = file_RegionSearchState_proto_rawDesc
)

func file_RegionSearchState_proto_rawDescGZIP() []byte {
	file_RegionSearchState_proto_rawDescOnce.Do(func() {
		file_RegionSearchState_proto_rawDescData = protoimpl.X.CompressGZIP(file_RegionSearchState_proto_rawDescData)
	})
	return file_RegionSearchState_proto_rawDescData
}

var file_RegionSearchState_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_RegionSearchState_proto_goTypes = []interface{}{
	(RegionSearchState)(0), // 0: RegionSearchState
}
var file_RegionSearchState_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RegionSearchState_proto_init() }
func file_RegionSearchState_proto_init() {
	if File_RegionSearchState_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_RegionSearchState_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RegionSearchState_proto_goTypes,
		DependencyIndexes: file_RegionSearchState_proto_depIdxs,
		EnumInfos:         file_RegionSearchState_proto_enumTypes,
	}.Build()
	File_RegionSearchState_proto = out.File
	file_RegionSearchState_proto_rawDesc = nil
	file_RegionSearchState_proto_goTypes = nil
	file_RegionSearchState_proto_depIdxs = nil
}
