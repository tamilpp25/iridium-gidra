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
// source: EntityJumpNotify.proto

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

type EntityJumpNotify_Type int32

const (
	EntityJumpNotify_TYPE_NULL    EntityJumpNotify_Type = 0
	EntityJumpNotify_TYPE_ACTIVE  EntityJumpNotify_Type = 1
	EntityJumpNotify_TYPE_PASSIVE EntityJumpNotify_Type = 2
)

// Enum value maps for EntityJumpNotify_Type.
var (
	EntityJumpNotify_Type_name = map[int32]string{
		0: "TYPE_NULL",
		1: "TYPE_ACTIVE",
		2: "TYPE_PASSIVE",
	}
	EntityJumpNotify_Type_value = map[string]int32{
		"TYPE_NULL":    0,
		"TYPE_ACTIVE":  1,
		"TYPE_PASSIVE": 2,
	}
)

func (x EntityJumpNotify_Type) Enum() *EntityJumpNotify_Type {
	p := new(EntityJumpNotify_Type)
	*p = x
	return p
}

func (x EntityJumpNotify_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntityJumpNotify_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_EntityJumpNotify_proto_enumTypes[0].Descriptor()
}

func (EntityJumpNotify_Type) Type() protoreflect.EnumType {
	return &file_EntityJumpNotify_proto_enumTypes[0]
}

func (x EntityJumpNotify_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntityJumpNotify_Type.Descriptor instead.
func (EntityJumpNotify_Type) EnumDescriptor() ([]byte, []int) {
	return file_EntityJumpNotify_proto_rawDescGZIP(), []int{0, 0}
}

type EntityJumpNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pos      *Vector               `protobuf:"bytes,6,opt,name=pos,proto3" json:"pos,omitempty"`
	JumpType EntityJumpNotify_Type `protobuf:"varint,2,opt,name=jump_type,json=jumpType,proto3,enum=EntityJumpNotify_Type" json:"jump_type,omitempty"`
	EntityId uint32                `protobuf:"varint,4,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	Rot      *Vector               `protobuf:"bytes,9,opt,name=rot,proto3" json:"rot,omitempty"`
}

func (x *EntityJumpNotify) Reset() {
	*x = EntityJumpNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EntityJumpNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityJumpNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityJumpNotify) ProtoMessage() {}

func (x *EntityJumpNotify) ProtoReflect() protoreflect.Message {
	mi := &file_EntityJumpNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityJumpNotify.ProtoReflect.Descriptor instead.
func (*EntityJumpNotify) Descriptor() ([]byte, []int) {
	return file_EntityJumpNotify_proto_rawDescGZIP(), []int{0}
}

func (x *EntityJumpNotify) GetPos() *Vector {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *EntityJumpNotify) GetJumpType() EntityJumpNotify_Type {
	if x != nil {
		return x.JumpType
	}
	return EntityJumpNotify_TYPE_NULL
}

func (x *EntityJumpNotify) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *EntityJumpNotify) GetRot() *Vector {
	if x != nil {
		return x.Rot
	}
	return nil
}

var File_EntityJumpNotify_proto protoreflect.FileDescriptor

var file_EntityJumpNotify_proto_rawDesc = []byte{
	0x0a, 0x16, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4a, 0x75, 0x6d, 0x70, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x01, 0x0a, 0x10, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x4a, 0x75, 0x6d, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x19, 0x0a, 0x03, 0x70,
	0x6f, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12, 0x33, 0x0a, 0x09, 0x6a, 0x75, 0x6d, 0x70, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x4a, 0x75, 0x6d, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x08, 0x6a, 0x75, 0x6d, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x03, 0x72, 0x6f, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x03,
	0x72, 0x6f, 0x74, 0x22, 0x38, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x49, 0x56, 0x45, 0x10, 0x02, 0x42, 0x06, 0x5a,
	0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EntityJumpNotify_proto_rawDescOnce sync.Once
	file_EntityJumpNotify_proto_rawDescData = file_EntityJumpNotify_proto_rawDesc
)

func file_EntityJumpNotify_proto_rawDescGZIP() []byte {
	file_EntityJumpNotify_proto_rawDescOnce.Do(func() {
		file_EntityJumpNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_EntityJumpNotify_proto_rawDescData)
	})
	return file_EntityJumpNotify_proto_rawDescData
}

var file_EntityJumpNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_EntityJumpNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EntityJumpNotify_proto_goTypes = []interface{}{
	(EntityJumpNotify_Type)(0), // 0: EntityJumpNotify.Type
	(*EntityJumpNotify)(nil),   // 1: EntityJumpNotify
	(*Vector)(nil),             // 2: Vector
}
var file_EntityJumpNotify_proto_depIdxs = []int32{
	2, // 0: EntityJumpNotify.pos:type_name -> Vector
	0, // 1: EntityJumpNotify.jump_type:type_name -> EntityJumpNotify.Type
	2, // 2: EntityJumpNotify.rot:type_name -> Vector
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_EntityJumpNotify_proto_init() }
func file_EntityJumpNotify_proto_init() {
	if File_EntityJumpNotify_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EntityJumpNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityJumpNotify); i {
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
			RawDescriptor: file_EntityJumpNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EntityJumpNotify_proto_goTypes,
		DependencyIndexes: file_EntityJumpNotify_proto_depIdxs,
		EnumInfos:         file_EntityJumpNotify_proto_enumTypes,
		MessageInfos:      file_EntityJumpNotify_proto_msgTypes,
	}.Build()
	File_EntityJumpNotify_proto = out.File
	file_EntityJumpNotify_proto_rawDesc = nil
	file_EntityJumpNotify_proto_goTypes = nil
	file_EntityJumpNotify_proto_depIdxs = nil
}
