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
// source: PlayerEyePointStateNotify.proto

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

// CmdId: 3327
// Obf: NAIIGNAOFCC
type PlayerEyePointStateNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FixLodLevel int32   `protobuf:"varint,12,opt,name=fix_lod_level,json=fixLodLevel,proto3" json:"fix_lod_level,omitempty"`
	JHPIEDBDGGB bool    `protobuf:"varint,15,opt,name=JHPIEDBDGGB,proto3" json:"JHPIEDBDGGB,omitempty"`
	ALEEOKDNOEE uint32  `protobuf:"varint,2,opt,name=ALEEOKDNOEE,proto3" json:"ALEEOKDNOEE,omitempty"`
	KNOFPFLJPFA bool    `protobuf:"varint,14,opt,name=KNOFPFLJPFA,proto3" json:"KNOFPFLJPFA,omitempty"`
	IMKAAEABEPB uint32  `protobuf:"varint,1,opt,name=IMKAAEABEPB,proto3" json:"IMKAAEABEPB,omitempty"`
	PACCIPICIEK uint32  `protobuf:"varint,4,opt,name=PACCIPICIEK,proto3" json:"PACCIPICIEK,omitempty"`
	EyePointPos *Vector `protobuf:"bytes,9,opt,name=eye_point_pos,json=eyePointPos,proto3" json:"eye_point_pos,omitempty"`
	AKIAFBHPMDI uint32  `protobuf:"varint,6,opt,name=AKIAFBHPMDI,proto3" json:"AKIAFBHPMDI,omitempty"`
	// Types that are assignable to RegionSize:
	//
	//	*PlayerEyePointStateNotify_SphereRadius
	//	*PlayerEyePointStateNotify_CubicSize
	//	*PlayerEyePointStateNotify_CylinderSize
	//	*PlayerEyePointStateNotify_PolygonSize
	RegionSize isPlayerEyePointStateNotify_RegionSize `protobuf_oneof:"region_size"`
}

func (x *PlayerEyePointStateNotify) Reset() {
	*x = PlayerEyePointStateNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerEyePointStateNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerEyePointStateNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerEyePointStateNotify) ProtoMessage() {}

func (x *PlayerEyePointStateNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerEyePointStateNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerEyePointStateNotify.ProtoReflect.Descriptor instead.
func (*PlayerEyePointStateNotify) Descriptor() ([]byte, []int) {
	return file_PlayerEyePointStateNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerEyePointStateNotify) GetFixLodLevel() int32 {
	if x != nil {
		return x.FixLodLevel
	}
	return 0
}

func (x *PlayerEyePointStateNotify) GetJHPIEDBDGGB() bool {
	if x != nil {
		return x.JHPIEDBDGGB
	}
	return false
}

func (x *PlayerEyePointStateNotify) GetALEEOKDNOEE() uint32 {
	if x != nil {
		return x.ALEEOKDNOEE
	}
	return 0
}

func (x *PlayerEyePointStateNotify) GetKNOFPFLJPFA() bool {
	if x != nil {
		return x.KNOFPFLJPFA
	}
	return false
}

func (x *PlayerEyePointStateNotify) GetIMKAAEABEPB() uint32 {
	if x != nil {
		return x.IMKAAEABEPB
	}
	return 0
}

func (x *PlayerEyePointStateNotify) GetPACCIPICIEK() uint32 {
	if x != nil {
		return x.PACCIPICIEK
	}
	return 0
}

func (x *PlayerEyePointStateNotify) GetEyePointPos() *Vector {
	if x != nil {
		return x.EyePointPos
	}
	return nil
}

func (x *PlayerEyePointStateNotify) GetAKIAFBHPMDI() uint32 {
	if x != nil {
		return x.AKIAFBHPMDI
	}
	return 0
}

func (m *PlayerEyePointStateNotify) GetRegionSize() isPlayerEyePointStateNotify_RegionSize {
	if m != nil {
		return m.RegionSize
	}
	return nil
}

func (x *PlayerEyePointStateNotify) GetSphereRadius() float32 {
	if x, ok := x.GetRegionSize().(*PlayerEyePointStateNotify_SphereRadius); ok {
		return x.SphereRadius
	}
	return 0
}

func (x *PlayerEyePointStateNotify) GetCubicSize() *Vector {
	if x, ok := x.GetRegionSize().(*PlayerEyePointStateNotify_CubicSize); ok {
		return x.CubicSize
	}
	return nil
}

func (x *PlayerEyePointStateNotify) GetCylinderSize() *CylinderRegionSize {
	if x, ok := x.GetRegionSize().(*PlayerEyePointStateNotify_CylinderSize); ok {
		return x.CylinderSize
	}
	return nil
}

func (x *PlayerEyePointStateNotify) GetPolygonSize() *PolygonRegionSize {
	if x, ok := x.GetRegionSize().(*PlayerEyePointStateNotify_PolygonSize); ok {
		return x.PolygonSize
	}
	return nil
}

type isPlayerEyePointStateNotify_RegionSize interface {
	isPlayerEyePointStateNotify_RegionSize()
}

type PlayerEyePointStateNotify_SphereRadius struct {
	SphereRadius float32 `protobuf:"fixed32,1935,opt,name=sphere_radius,json=sphereRadius,proto3,oneof"`
}

type PlayerEyePointStateNotify_CubicSize struct {
	CubicSize *Vector `protobuf:"bytes,181,opt,name=cubic_size,json=cubicSize,proto3,oneof"`
}

type PlayerEyePointStateNotify_CylinderSize struct {
	CylinderSize *CylinderRegionSize `protobuf:"bytes,976,opt,name=cylinder_size,json=cylinderSize,proto3,oneof"`
}

type PlayerEyePointStateNotify_PolygonSize struct {
	PolygonSize *PolygonRegionSize `protobuf:"bytes,1589,opt,name=polygon_size,json=polygonSize,proto3,oneof"`
}

func (*PlayerEyePointStateNotify_SphereRadius) isPlayerEyePointStateNotify_RegionSize() {}

func (*PlayerEyePointStateNotify_CubicSize) isPlayerEyePointStateNotify_RegionSize() {}

func (*PlayerEyePointStateNotify_CylinderSize) isPlayerEyePointStateNotify_RegionSize() {}

func (*PlayerEyePointStateNotify_PolygonSize) isPlayerEyePointStateNotify_RegionSize() {}

var File_PlayerEyePointStateNotify_proto protoreflect.FileDescriptor

var file_PlayerEyePointStateNotify_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x79, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x18, 0x43, 0x79, 0x6c, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x53,
	0x69, 0x7a, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x50, 0x6f, 0x6c, 0x79, 0x67,
	0x6f, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x7a, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x91, 0x04, 0x0a, 0x19, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x79, 0x65,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x22, 0x0a, 0x0d, 0x66, 0x69, 0x78, 0x5f, 0x6c, 0x6f, 0x64, 0x5f, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x66, 0x69, 0x78, 0x4c, 0x6f, 0x64, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x48, 0x50, 0x49, 0x45, 0x44, 0x42, 0x44,
	0x47, 0x47, 0x42, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4a, 0x48, 0x50, 0x49, 0x45,
	0x44, 0x42, 0x44, 0x47, 0x47, 0x42, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x4c, 0x45, 0x45, 0x4f, 0x4b,
	0x44, 0x4e, 0x4f, 0x45, 0x45, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x41, 0x4c, 0x45,
	0x45, 0x4f, 0x4b, 0x44, 0x4e, 0x4f, 0x45, 0x45, 0x12, 0x20, 0x0a, 0x0b, 0x4b, 0x4e, 0x4f, 0x46,
	0x50, 0x46, 0x4c, 0x4a, 0x50, 0x46, 0x41, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4b,
	0x4e, 0x4f, 0x46, 0x50, 0x46, 0x4c, 0x4a, 0x50, 0x46, 0x41, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x4d,
	0x4b, 0x41, 0x41, 0x45, 0x41, 0x42, 0x45, 0x50, 0x42, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x49, 0x4d, 0x4b, 0x41, 0x41, 0x45, 0x41, 0x42, 0x45, 0x50, 0x42, 0x12, 0x20, 0x0a, 0x0b,
	0x50, 0x41, 0x43, 0x43, 0x49, 0x50, 0x49, 0x43, 0x49, 0x45, 0x4b, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x50, 0x41, 0x43, 0x43, 0x49, 0x50, 0x49, 0x43, 0x49, 0x45, 0x4b, 0x12, 0x2b,
	0x0a, 0x0d, 0x65, 0x79, 0x65, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x70, 0x6f, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0b,
	0x65, 0x79, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x41,
	0x4b, 0x49, 0x41, 0x46, 0x42, 0x48, 0x50, 0x4d, 0x44, 0x49, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x41, 0x4b, 0x49, 0x41, 0x46, 0x42, 0x48, 0x50, 0x4d, 0x44, 0x49, 0x12, 0x26, 0x0a,
	0x0d, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x5f, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x8f,
	0x0f, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x52,
	0x61, 0x64, 0x69, 0x75, 0x73, 0x12, 0x29, 0x0a, 0x0a, 0x63, 0x75, 0x62, 0x69, 0x63, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0xb5, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x09, 0x63, 0x75, 0x62, 0x69, 0x63, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x3b, 0x0a, 0x0d, 0x63, 0x79, 0x6c, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0xd0, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x43, 0x79, 0x6c, 0x69, 0x6e,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x7a, 0x65, 0x48, 0x00, 0x52,
	0x0c, 0x63, 0x79, 0x6c, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x38, 0x0a,
	0x0c, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0xb5, 0x0c,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x7a, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x6f, 0x6c, 0x79,
	0x67, 0x6f, 0x6e, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerEyePointStateNotify_proto_rawDescOnce sync.Once
	file_PlayerEyePointStateNotify_proto_rawDescData = file_PlayerEyePointStateNotify_proto_rawDesc
)

func file_PlayerEyePointStateNotify_proto_rawDescGZIP() []byte {
	file_PlayerEyePointStateNotify_proto_rawDescOnce.Do(func() {
		file_PlayerEyePointStateNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerEyePointStateNotify_proto_rawDescData)
	})
	return file_PlayerEyePointStateNotify_proto_rawDescData
}

var file_PlayerEyePointStateNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerEyePointStateNotify_proto_goTypes = []interface{}{
	(*PlayerEyePointStateNotify)(nil), // 0: PlayerEyePointStateNotify
	(*Vector)(nil),                    // 1: Vector
	(*CylinderRegionSize)(nil),        // 2: CylinderRegionSize
	(*PolygonRegionSize)(nil),         // 3: PolygonRegionSize
}
var file_PlayerEyePointStateNotify_proto_depIdxs = []int32{
	1, // 0: PlayerEyePointStateNotify.eye_point_pos:type_name -> Vector
	1, // 1: PlayerEyePointStateNotify.cubic_size:type_name -> Vector
	2, // 2: PlayerEyePointStateNotify.cylinder_size:type_name -> CylinderRegionSize
	3, // 3: PlayerEyePointStateNotify.polygon_size:type_name -> PolygonRegionSize
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_PlayerEyePointStateNotify_proto_init() }
func file_PlayerEyePointStateNotify_proto_init() {
	if File_PlayerEyePointStateNotify_proto != nil {
		return
	}
	file_Vector_proto_init()
	file_CylinderRegionSize_proto_init()
	file_PolygonRegionSize_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerEyePointStateNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerEyePointStateNotify); i {
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
	file_PlayerEyePointStateNotify_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PlayerEyePointStateNotify_SphereRadius)(nil),
		(*PlayerEyePointStateNotify_CubicSize)(nil),
		(*PlayerEyePointStateNotify_CylinderSize)(nil),
		(*PlayerEyePointStateNotify_PolygonSize)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_PlayerEyePointStateNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerEyePointStateNotify_proto_goTypes,
		DependencyIndexes: file_PlayerEyePointStateNotify_proto_depIdxs,
		MessageInfos:      file_PlayerEyePointStateNotify_proto_msgTypes,
	}.Build()
	File_PlayerEyePointStateNotify_proto = out.File
	file_PlayerEyePointStateNotify_proto_rawDesc = nil
	file_PlayerEyePointStateNotify_proto_goTypes = nil
	file_PlayerEyePointStateNotify_proto_depIdxs = nil
}
