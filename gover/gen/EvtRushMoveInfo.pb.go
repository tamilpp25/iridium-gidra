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
// source: EvtRushMoveInfo.proto

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

type EvtRushMoveInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Velocity            *Vector `protobuf:"bytes,8,opt,name=velocity,proto3" json:"velocity,omitempty"`
	StateNameHash       int32   `protobuf:"varint,4,opt,name=state_name_hash,json=stateNameHash,proto3" json:"state_name_hash,omitempty"`
	EntityId            uint32  `protobuf:"varint,12,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	Pos                 *Vector `protobuf:"bytes,7,opt,name=pos,proto3" json:"pos,omitempty"`
	FaceAngleCompact    int32   `protobuf:"varint,1,opt,name=face_angle_compact,json=faceAngleCompact,proto3" json:"face_angle_compact,omitempty"`
	TimeRange           float32 `protobuf:"fixed32,13,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"`
	Unk3300_NEBMDDGPBON *Vector `protobuf:"bytes,15,opt,name=Unk3300_NEBMDDGPBON,json=Unk3300NEBMDDGPBON,proto3" json:"Unk3300_NEBMDDGPBON,omitempty"`
	Unk3300_FGAPBJIPJFG *Vector `protobuf:"bytes,5,opt,name=Unk3300_FGAPBJIPJFG,json=Unk3300FGAPBJIPJFG,proto3" json:"Unk3300_FGAPBJIPJFG,omitempty"`
}

func (x *EvtRushMoveInfo) Reset() {
	*x = EvtRushMoveInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EvtRushMoveInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvtRushMoveInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvtRushMoveInfo) ProtoMessage() {}

func (x *EvtRushMoveInfo) ProtoReflect() protoreflect.Message {
	mi := &file_EvtRushMoveInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvtRushMoveInfo.ProtoReflect.Descriptor instead.
func (*EvtRushMoveInfo) Descriptor() ([]byte, []int) {
	return file_EvtRushMoveInfo_proto_rawDescGZIP(), []int{0}
}

func (x *EvtRushMoveInfo) GetVelocity() *Vector {
	if x != nil {
		return x.Velocity
	}
	return nil
}

func (x *EvtRushMoveInfo) GetStateNameHash() int32 {
	if x != nil {
		return x.StateNameHash
	}
	return 0
}

func (x *EvtRushMoveInfo) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *EvtRushMoveInfo) GetPos() *Vector {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *EvtRushMoveInfo) GetFaceAngleCompact() int32 {
	if x != nil {
		return x.FaceAngleCompact
	}
	return 0
}

func (x *EvtRushMoveInfo) GetTimeRange() float32 {
	if x != nil {
		return x.TimeRange
	}
	return 0
}

func (x *EvtRushMoveInfo) GetUnk3300_NEBMDDGPBON() *Vector {
	if x != nil {
		return x.Unk3300_NEBMDDGPBON
	}
	return nil
}

func (x *EvtRushMoveInfo) GetUnk3300_FGAPBJIPJFG() *Vector {
	if x != nil {
		return x.Unk3300_FGAPBJIPJFG
	}
	return nil
}

var File_EvtRushMoveInfo_proto protoreflect.FileDescriptor

var file_EvtRushMoveInfo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x45, 0x76, 0x74, 0x52, 0x75, 0x73, 0x68, 0x4d, 0x6f, 0x76, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x02, 0x0a, 0x0f, 0x45, 0x76, 0x74, 0x52, 0x75, 0x73,
	0x68, 0x4d, 0x6f, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x08, 0x76, 0x65, 0x6c,
	0x6f, 0x63, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x12, 0x26,
	0x0a, 0x0f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12, 0x2c,
	0x0a, 0x12, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x66, 0x61, 0x63, 0x65,
	0x41, 0x6e, 0x67, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x13, 0x55,
	0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x4e, 0x45, 0x42, 0x4d, 0x44, 0x44, 0x47, 0x50, 0x42,
	0x4f, 0x4e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x4e, 0x45, 0x42, 0x4d, 0x44, 0x44,
	0x47, 0x50, 0x42, 0x4f, 0x4e, 0x12, 0x38, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30,
	0x5f, 0x46, 0x47, 0x41, 0x50, 0x42, 0x4a, 0x49, 0x50, 0x4a, 0x46, 0x47, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x12, 0x55, 0x6e, 0x6b,
	0x33, 0x33, 0x30, 0x30, 0x46, 0x47, 0x41, 0x50, 0x42, 0x4a, 0x49, 0x50, 0x4a, 0x46, 0x47, 0x42,
	0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EvtRushMoveInfo_proto_rawDescOnce sync.Once
	file_EvtRushMoveInfo_proto_rawDescData = file_EvtRushMoveInfo_proto_rawDesc
)

func file_EvtRushMoveInfo_proto_rawDescGZIP() []byte {
	file_EvtRushMoveInfo_proto_rawDescOnce.Do(func() {
		file_EvtRushMoveInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_EvtRushMoveInfo_proto_rawDescData)
	})
	return file_EvtRushMoveInfo_proto_rawDescData
}

var file_EvtRushMoveInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EvtRushMoveInfo_proto_goTypes = []interface{}{
	(*EvtRushMoveInfo)(nil), // 0: EvtRushMoveInfo
	(*Vector)(nil),          // 1: Vector
}
var file_EvtRushMoveInfo_proto_depIdxs = []int32{
	1, // 0: EvtRushMoveInfo.velocity:type_name -> Vector
	1, // 1: EvtRushMoveInfo.pos:type_name -> Vector
	1, // 2: EvtRushMoveInfo.Unk3300_NEBMDDGPBON:type_name -> Vector
	1, // 3: EvtRushMoveInfo.Unk3300_FGAPBJIPJFG:type_name -> Vector
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_EvtRushMoveInfo_proto_init() }
func file_EvtRushMoveInfo_proto_init() {
	if File_EvtRushMoveInfo_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EvtRushMoveInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvtRushMoveInfo); i {
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
			RawDescriptor: file_EvtRushMoveInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EvtRushMoveInfo_proto_goTypes,
		DependencyIndexes: file_EvtRushMoveInfo_proto_depIdxs,
		MessageInfos:      file_EvtRushMoveInfo_proto_msgTypes,
	}.Build()
	File_EvtRushMoveInfo_proto = out.File
	file_EvtRushMoveInfo_proto_rawDesc = nil
	file_EvtRushMoveInfo_proto_goTypes = nil
	file_EvtRushMoveInfo_proto_depIdxs = nil
}
