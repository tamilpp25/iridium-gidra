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
// source: PlatformInfo.proto

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

type PlatformInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouteId            uint32             `protobuf:"varint,1,opt,name=route_id,json=routeId,proto3" json:"route_id,omitempty"`
	StartIndex         int32              `protobuf:"varint,2,opt,name=start_index,json=startIndex,proto3" json:"start_index,omitempty"`
	StartRouteTime     uint32             `protobuf:"varint,3,opt,name=start_route_time,json=startRouteTime,proto3" json:"start_route_time,omitempty"`
	StartSceneTime     uint32             `protobuf:"varint,4,opt,name=start_scene_time,json=startSceneTime,proto3" json:"start_scene_time,omitempty"`
	StartPos           *Vector            `protobuf:"bytes,7,opt,name=start_pos,json=startPos,proto3" json:"start_pos,omitempty"`
	IsStarted          bool               `protobuf:"varint,8,opt,name=is_started,json=isStarted,proto3" json:"is_started,omitempty"`
	StartRot           *MathQuaternion    `protobuf:"bytes,9,opt,name=start_rot,json=startRot,proto3" json:"start_rot,omitempty"`
	StopSceneTime      uint32             `protobuf:"varint,10,opt,name=stop_scene_time,json=stopSceneTime,proto3" json:"stop_scene_time,omitempty"`
	PosOffset          *Vector            `protobuf:"bytes,11,opt,name=pos_offset,json=posOffset,proto3" json:"pos_offset,omitempty"`
	RotOffset          *MathQuaternion    `protobuf:"bytes,12,opt,name=rot_offset,json=rotOffset,proto3" json:"rot_offset,omitempty"`
	MovingPlatformType MovingPlatformType `protobuf:"varint,13,opt,name=moving_platform_type,json=movingPlatformType,proto3,enum=MovingPlatformType" json:"moving_platform_type,omitempty"`
	IsActive           bool               `protobuf:"varint,14,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Route              *Route             `protobuf:"bytes,15,opt,name=route,proto3" json:"route,omitempty"`
	PointId            uint32             `protobuf:"varint,16,opt,name=point_id,json=pointId,proto3" json:"point_id,omitempty"`
}

func (x *PlatformInfo) Reset() {
	*x = PlatformInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlatformInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatformInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformInfo) ProtoMessage() {}

func (x *PlatformInfo) ProtoReflect() protoreflect.Message {
	mi := &file_PlatformInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformInfo.ProtoReflect.Descriptor instead.
func (*PlatformInfo) Descriptor() ([]byte, []int) {
	return file_PlatformInfo_proto_rawDescGZIP(), []int{0}
}

func (x *PlatformInfo) GetRouteId() uint32 {
	if x != nil {
		return x.RouteId
	}
	return 0
}

func (x *PlatformInfo) GetStartIndex() int32 {
	if x != nil {
		return x.StartIndex
	}
	return 0
}

func (x *PlatformInfo) GetStartRouteTime() uint32 {
	if x != nil {
		return x.StartRouteTime
	}
	return 0
}

func (x *PlatformInfo) GetStartSceneTime() uint32 {
	if x != nil {
		return x.StartSceneTime
	}
	return 0
}

func (x *PlatformInfo) GetStartPos() *Vector {
	if x != nil {
		return x.StartPos
	}
	return nil
}

func (x *PlatformInfo) GetIsStarted() bool {
	if x != nil {
		return x.IsStarted
	}
	return false
}

func (x *PlatformInfo) GetStartRot() *MathQuaternion {
	if x != nil {
		return x.StartRot
	}
	return nil
}

func (x *PlatformInfo) GetStopSceneTime() uint32 {
	if x != nil {
		return x.StopSceneTime
	}
	return 0
}

func (x *PlatformInfo) GetPosOffset() *Vector {
	if x != nil {
		return x.PosOffset
	}
	return nil
}

func (x *PlatformInfo) GetRotOffset() *MathQuaternion {
	if x != nil {
		return x.RotOffset
	}
	return nil
}

func (x *PlatformInfo) GetMovingPlatformType() MovingPlatformType {
	if x != nil {
		return x.MovingPlatformType
	}
	return MovingPlatformType_MOVING_PLATFORM_TYPE_NONE
}

func (x *PlatformInfo) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *PlatformInfo) GetRoute() *Route {
	if x != nil {
		return x.Route
	}
	return nil
}

func (x *PlatformInfo) GetPointId() uint32 {
	if x != nil {
		return x.PointId
	}
	return 0
}

var File_PlatformInfo_proto protoreflect.FileDescriptor

var file_PlatformInfo_proto_rawDesc = []byte{
	0x0a, 0x12, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x4d, 0x61, 0x74, 0x68, 0x51, 0x75, 0x61, 0x74, 0x65, 0x72,
	0x6e, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x4d, 0x6f, 0x76, 0x69,
	0x6e, 0x67, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xae, 0x04, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x28, 0x0a, 0x10,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x24, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x50, 0x6f, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x72,
	0x6f, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4d, 0x61, 0x74, 0x68, 0x51,
	0x75, 0x61, 0x74, 0x65, 0x72, 0x6e, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x52, 0x6f, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x73, 0x63, 0x65, 0x6e,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x73, 0x74,
	0x6f, 0x70, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0a, 0x70,
	0x6f, 0x73, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x2e, 0x0a, 0x0a, 0x72, 0x6f, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4d, 0x61, 0x74, 0x68, 0x51, 0x75,
	0x61, 0x74, 0x65, 0x72, 0x6e, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x72, 0x6f, 0x74, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x45, 0x0a, 0x14, 0x6d, 0x6f, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x13, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x12, 0x6d, 0x6f, 0x76, 0x69, 0x6e, 0x67, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x05,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64,
	0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlatformInfo_proto_rawDescOnce sync.Once
	file_PlatformInfo_proto_rawDescData = file_PlatformInfo_proto_rawDesc
)

func file_PlatformInfo_proto_rawDescGZIP() []byte {
	file_PlatformInfo_proto_rawDescOnce.Do(func() {
		file_PlatformInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlatformInfo_proto_rawDescData)
	})
	return file_PlatformInfo_proto_rawDescData
}

var file_PlatformInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlatformInfo_proto_goTypes = []interface{}{
	(*PlatformInfo)(nil),    // 0: PlatformInfo
	(*Vector)(nil),          // 1: Vector
	(*MathQuaternion)(nil),  // 2: MathQuaternion
	(MovingPlatformType)(0), // 3: MovingPlatformType
	(*Route)(nil),           // 4: Route
}
var file_PlatformInfo_proto_depIdxs = []int32{
	1, // 0: PlatformInfo.start_pos:type_name -> Vector
	2, // 1: PlatformInfo.start_rot:type_name -> MathQuaternion
	1, // 2: PlatformInfo.pos_offset:type_name -> Vector
	2, // 3: PlatformInfo.rot_offset:type_name -> MathQuaternion
	3, // 4: PlatformInfo.moving_platform_type:type_name -> MovingPlatformType
	4, // 5: PlatformInfo.route:type_name -> Route
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_PlatformInfo_proto_init() }
func file_PlatformInfo_proto_init() {
	if File_PlatformInfo_proto != nil {
		return
	}
	file_MathQuaternion_proto_init()
	file_MovingPlatformType_proto_init()
	file_Route_proto_init()
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlatformInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatformInfo); i {
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
			RawDescriptor: file_PlatformInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlatformInfo_proto_goTypes,
		DependencyIndexes: file_PlatformInfo_proto_depIdxs,
		MessageInfos:      file_PlatformInfo_proto_msgTypes,
	}.Build()
	File_PlatformInfo_proto = out.File
	file_PlatformInfo_proto_rawDesc = nil
	file_PlatformInfo_proto_goTypes = nil
	file_PlatformInfo_proto_depIdxs = nil
}
