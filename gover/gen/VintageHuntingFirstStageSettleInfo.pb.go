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
// source: VintageHuntingFirstStageSettleInfo.proto

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

type VintageHuntingFirstStageSettleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk3300_HIKADHBLPEB uint32 `protobuf:"varint,2,opt,name=Unk3300_HIKADHBLPEB,json=Unk3300HIKADHBLPEB,proto3" json:"Unk3300_HIKADHBLPEB,omitempty"`
	Score               uint32 `protobuf:"varint,1,opt,name=score,proto3" json:"score,omitempty"`
	Unk3300_NILGNHOPGLD uint32 `protobuf:"varint,3,opt,name=Unk3300_NILGNHOPGLD,json=Unk3300NILGNHOPGLD,proto3" json:"Unk3300_NILGNHOPGLD,omitempty"`
	Unk3300_IPAJLKOBDKC uint32 `protobuf:"varint,11,opt,name=Unk3300_IPAJLKOBDKC,json=Unk3300IPAJLKOBDKC,proto3" json:"Unk3300_IPAJLKOBDKC,omitempty"`
}

func (x *VintageHuntingFirstStageSettleInfo) Reset() {
	*x = VintageHuntingFirstStageSettleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_VintageHuntingFirstStageSettleInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VintageHuntingFirstStageSettleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VintageHuntingFirstStageSettleInfo) ProtoMessage() {}

func (x *VintageHuntingFirstStageSettleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_VintageHuntingFirstStageSettleInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VintageHuntingFirstStageSettleInfo.ProtoReflect.Descriptor instead.
func (*VintageHuntingFirstStageSettleInfo) Descriptor() ([]byte, []int) {
	return file_VintageHuntingFirstStageSettleInfo_proto_rawDescGZIP(), []int{0}
}

func (x *VintageHuntingFirstStageSettleInfo) GetUnk3300_HIKADHBLPEB() uint32 {
	if x != nil {
		return x.Unk3300_HIKADHBLPEB
	}
	return 0
}

func (x *VintageHuntingFirstStageSettleInfo) GetScore() uint32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *VintageHuntingFirstStageSettleInfo) GetUnk3300_NILGNHOPGLD() uint32 {
	if x != nil {
		return x.Unk3300_NILGNHOPGLD
	}
	return 0
}

func (x *VintageHuntingFirstStageSettleInfo) GetUnk3300_IPAJLKOBDKC() uint32 {
	if x != nil {
		return x.Unk3300_IPAJLKOBDKC
	}
	return 0
}

var File_VintageHuntingFirstStageSettleInfo_proto protoreflect.FileDescriptor

var file_VintageHuntingFirstStageSettleInfo_proto_rawDesc = []byte{
	0x0a, 0x28, 0x56, 0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x48, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67,
	0x46, 0x69, 0x72, 0x73, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x22, 0x56,
	0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x48, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x72,
	0x73, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x48, 0x49, 0x4b,
	0x41, 0x44, 0x48, 0x42, 0x4c, 0x50, 0x45, 0x42, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12,
	0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x48, 0x49, 0x4b, 0x41, 0x44, 0x48, 0x42, 0x4c, 0x50,
	0x45, 0x42, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33,
	0x33, 0x30, 0x30, 0x5f, 0x4e, 0x49, 0x4c, 0x47, 0x4e, 0x48, 0x4f, 0x50, 0x47, 0x4c, 0x44, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x4e, 0x49,
	0x4c, 0x47, 0x4e, 0x48, 0x4f, 0x50, 0x47, 0x4c, 0x44, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b,
	0x33, 0x33, 0x30, 0x30, 0x5f, 0x49, 0x50, 0x41, 0x4a, 0x4c, 0x4b, 0x4f, 0x42, 0x44, 0x4b, 0x43,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x49,
	0x50, 0x41, 0x4a, 0x4c, 0x4b, 0x4f, 0x42, 0x44, 0x4b, 0x43, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67,
	0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_VintageHuntingFirstStageSettleInfo_proto_rawDescOnce sync.Once
	file_VintageHuntingFirstStageSettleInfo_proto_rawDescData = file_VintageHuntingFirstStageSettleInfo_proto_rawDesc
)

func file_VintageHuntingFirstStageSettleInfo_proto_rawDescGZIP() []byte {
	file_VintageHuntingFirstStageSettleInfo_proto_rawDescOnce.Do(func() {
		file_VintageHuntingFirstStageSettleInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_VintageHuntingFirstStageSettleInfo_proto_rawDescData)
	})
	return file_VintageHuntingFirstStageSettleInfo_proto_rawDescData
}

var file_VintageHuntingFirstStageSettleInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_VintageHuntingFirstStageSettleInfo_proto_goTypes = []interface{}{
	(*VintageHuntingFirstStageSettleInfo)(nil), // 0: VintageHuntingFirstStageSettleInfo
}
var file_VintageHuntingFirstStageSettleInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_VintageHuntingFirstStageSettleInfo_proto_init() }
func file_VintageHuntingFirstStageSettleInfo_proto_init() {
	if File_VintageHuntingFirstStageSettleInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_VintageHuntingFirstStageSettleInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VintageHuntingFirstStageSettleInfo); i {
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
			RawDescriptor: file_VintageHuntingFirstStageSettleInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_VintageHuntingFirstStageSettleInfo_proto_goTypes,
		DependencyIndexes: file_VintageHuntingFirstStageSettleInfo_proto_depIdxs,
		MessageInfos:      file_VintageHuntingFirstStageSettleInfo_proto_msgTypes,
	}.Build()
	File_VintageHuntingFirstStageSettleInfo_proto = out.File
	file_VintageHuntingFirstStageSettleInfo_proto_rawDesc = nil
	file_VintageHuntingFirstStageSettleInfo_proto_goTypes = nil
	file_VintageHuntingFirstStageSettleInfo_proto_depIdxs = nil
}
