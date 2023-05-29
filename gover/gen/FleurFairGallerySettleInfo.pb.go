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
// source: FleurFairGallerySettleInfo.proto

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

// Obf: HGEBLCIHPJP
type FleurFairGallerySettleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess     bool             `protobuf:"varint,3,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	CHEAJDKKPDB   uint32           `protobuf:"varint,7,opt,name=CHEAJDKKPDB,proto3" json:"CHEAJDKKPDB,omitempty"`
	EnergyStatMap map[uint32]int32 `protobuf:"bytes,11,rep,name=energy_stat_map,json=energyStatMap,proto3" json:"energy_stat_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	FGIHLGCHBMN   uint32           `protobuf:"varint,15,opt,name=FGIHLGCHBMN,proto3" json:"FGIHLGCHBMN,omitempty"`
	Energy        uint32           `protobuf:"varint,14,opt,name=energy,proto3" json:"energy,omitempty"`
}

func (x *FleurFairGallerySettleInfo) Reset() {
	*x = FleurFairGallerySettleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FleurFairGallerySettleInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FleurFairGallerySettleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FleurFairGallerySettleInfo) ProtoMessage() {}

func (x *FleurFairGallerySettleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_FleurFairGallerySettleInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FleurFairGallerySettleInfo.ProtoReflect.Descriptor instead.
func (*FleurFairGallerySettleInfo) Descriptor() ([]byte, []int) {
	return file_FleurFairGallerySettleInfo_proto_rawDescGZIP(), []int{0}
}

func (x *FleurFairGallerySettleInfo) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *FleurFairGallerySettleInfo) GetCHEAJDKKPDB() uint32 {
	if x != nil {
		return x.CHEAJDKKPDB
	}
	return 0
}

func (x *FleurFairGallerySettleInfo) GetEnergyStatMap() map[uint32]int32 {
	if x != nil {
		return x.EnergyStatMap
	}
	return nil
}

func (x *FleurFairGallerySettleInfo) GetFGIHLGCHBMN() uint32 {
	if x != nil {
		return x.FGIHLGCHBMN
	}
	return 0
}

func (x *FleurFairGallerySettleInfo) GetEnergy() uint32 {
	if x != nil {
		return x.Energy
	}
	return 0
}

var File_FleurFairGallerySettleInfo_proto protoreflect.FileDescriptor

var file_FleurFairGallerySettleInfo_proto_rawDesc = []byte{
	0x0a, 0x20, 0x46, 0x6c, 0x65, 0x75, 0x72, 0x46, 0x61, 0x69, 0x72, 0x47, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x79, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb1, 0x02, 0x0a, 0x1a, 0x46, 0x6c, 0x65, 0x75, 0x72, 0x46, 0x61, 0x69, 0x72,
	0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x43, 0x48, 0x45, 0x41, 0x4a, 0x44, 0x4b, 0x4b, 0x50, 0x44, 0x42, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x48, 0x45, 0x41, 0x4a, 0x44, 0x4b, 0x4b, 0x50,
	0x44, 0x42, 0x12, 0x56, 0x0a, 0x0f, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x46, 0x6c,
	0x65, 0x75, 0x72, 0x46, 0x61, 0x69, 0x72, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x65,
	0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x65, 0x6e, 0x65,
	0x72, 0x67, 0x79, 0x53, 0x74, 0x61, 0x74, 0x4d, 0x61, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x47,
	0x49, 0x48, 0x4c, 0x47, 0x43, 0x48, 0x42, 0x4d, 0x4e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x46, 0x47, 0x49, 0x48, 0x4c, 0x47, 0x43, 0x48, 0x42, 0x4d, 0x4e, 0x12, 0x16, 0x0a, 0x06,
	0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x65, 0x6e,
	0x65, 0x72, 0x67, 0x79, 0x1a, 0x40, 0x0a, 0x12, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FleurFairGallerySettleInfo_proto_rawDescOnce sync.Once
	file_FleurFairGallerySettleInfo_proto_rawDescData = file_FleurFairGallerySettleInfo_proto_rawDesc
)

func file_FleurFairGallerySettleInfo_proto_rawDescGZIP() []byte {
	file_FleurFairGallerySettleInfo_proto_rawDescOnce.Do(func() {
		file_FleurFairGallerySettleInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_FleurFairGallerySettleInfo_proto_rawDescData)
	})
	return file_FleurFairGallerySettleInfo_proto_rawDescData
}

var file_FleurFairGallerySettleInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_FleurFairGallerySettleInfo_proto_goTypes = []interface{}{
	(*FleurFairGallerySettleInfo)(nil), // 0: FleurFairGallerySettleInfo
	nil,                                // 1: FleurFairGallerySettleInfo.EnergyStatMapEntry
}
var file_FleurFairGallerySettleInfo_proto_depIdxs = []int32{
	1, // 0: FleurFairGallerySettleInfo.energy_stat_map:type_name -> FleurFairGallerySettleInfo.EnergyStatMapEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_FleurFairGallerySettleInfo_proto_init() }
func file_FleurFairGallerySettleInfo_proto_init() {
	if File_FleurFairGallerySettleInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_FleurFairGallerySettleInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FleurFairGallerySettleInfo); i {
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
			RawDescriptor: file_FleurFairGallerySettleInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FleurFairGallerySettleInfo_proto_goTypes,
		DependencyIndexes: file_FleurFairGallerySettleInfo_proto_depIdxs,
		MessageInfos:      file_FleurFairGallerySettleInfo_proto_msgTypes,
	}.Build()
	File_FleurFairGallerySettleInfo_proto = out.File
	file_FleurFairGallerySettleInfo_proto_rawDesc = nil
	file_FleurFairGallerySettleInfo_proto_goTypes = nil
	file_FleurFairGallerySettleInfo_proto_depIdxs = nil
}
