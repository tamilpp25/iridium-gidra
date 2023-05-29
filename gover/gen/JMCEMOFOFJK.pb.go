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
// source: JMCEMOFOFJK.proto

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

// CmdId: 9172
type JMCEMOFOFJK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DungeonGuid uint64 `protobuf:"varint,15,opt,name=dungeon_guid,json=dungeonGuid,proto3" json:"dungeon_guid,omitempty"`
	RoomId      uint32 `protobuf:"varint,12,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	IsAdd       bool   `protobuf:"varint,8,opt,name=is_add,json=isAdd,proto3" json:"is_add,omitempty"`
}

func (x *JMCEMOFOFJK) Reset() {
	*x = JMCEMOFOFJK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_JMCEMOFOFJK_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JMCEMOFOFJK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JMCEMOFOFJK) ProtoMessage() {}

func (x *JMCEMOFOFJK) ProtoReflect() protoreflect.Message {
	mi := &file_JMCEMOFOFJK_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JMCEMOFOFJK.ProtoReflect.Descriptor instead.
func (*JMCEMOFOFJK) Descriptor() ([]byte, []int) {
	return file_JMCEMOFOFJK_proto_rawDescGZIP(), []int{0}
}

func (x *JMCEMOFOFJK) GetDungeonGuid() uint64 {
	if x != nil {
		return x.DungeonGuid
	}
	return 0
}

func (x *JMCEMOFOFJK) GetRoomId() uint32 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *JMCEMOFOFJK) GetIsAdd() bool {
	if x != nil {
		return x.IsAdd
	}
	return false
}

var File_JMCEMOFOFJK_proto protoreflect.FileDescriptor

var file_JMCEMOFOFJK_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4a, 0x4d, 0x43, 0x45, 0x4d, 0x4f, 0x46, 0x4f, 0x46, 0x4a, 0x4b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x0b, 0x4a, 0x4d, 0x43, 0x45, 0x4d, 0x4f, 0x46, 0x4f, 0x46,
	0x4a, 0x4b, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x67, 0x75,
	0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f,
	0x6e, 0x47, 0x75, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x15,
	0x0a, 0x06, 0x69, 0x73, 0x5f, 0x61, 0x64, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x69, 0x73, 0x41, 0x64, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_JMCEMOFOFJK_proto_rawDescOnce sync.Once
	file_JMCEMOFOFJK_proto_rawDescData = file_JMCEMOFOFJK_proto_rawDesc
)

func file_JMCEMOFOFJK_proto_rawDescGZIP() []byte {
	file_JMCEMOFOFJK_proto_rawDescOnce.Do(func() {
		file_JMCEMOFOFJK_proto_rawDescData = protoimpl.X.CompressGZIP(file_JMCEMOFOFJK_proto_rawDescData)
	})
	return file_JMCEMOFOFJK_proto_rawDescData
}

var file_JMCEMOFOFJK_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_JMCEMOFOFJK_proto_goTypes = []interface{}{
	(*JMCEMOFOFJK)(nil), // 0: JMCEMOFOFJK
}
var file_JMCEMOFOFJK_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_JMCEMOFOFJK_proto_init() }
func file_JMCEMOFOFJK_proto_init() {
	if File_JMCEMOFOFJK_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_JMCEMOFOFJK_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JMCEMOFOFJK); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_JMCEMOFOFJK_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_JMCEMOFOFJK_proto_goTypes,
		DependencyIndexes: file_JMCEMOFOFJK_proto_depIdxs,
		MessageInfos:      file_JMCEMOFOFJK_proto_msgTypes,
	}.Build()
	File_JMCEMOFOFJK_proto = out.File
	file_JMCEMOFOFJK_proto_rawDesc = nil
	file_JMCEMOFOFJK_proto_goTypes = nil
	file_JMCEMOFOFJK_proto_depIdxs = nil
}
