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
// source: VintageMarketDealInfo.proto

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

// Obf: CKOJLAKBAGH
type VintageMarketDealInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraderItemInfoMap map[uint32]*VintageMarketTraderInfo `protobuf:"bytes,10,rep,name=trader_item_info_map,json=traderItemInfoMap,proto3" json:"trader_item_info_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *VintageMarketDealInfo) Reset() {
	*x = VintageMarketDealInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_VintageMarketDealInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VintageMarketDealInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VintageMarketDealInfo) ProtoMessage() {}

func (x *VintageMarketDealInfo) ProtoReflect() protoreflect.Message {
	mi := &file_VintageMarketDealInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VintageMarketDealInfo.ProtoReflect.Descriptor instead.
func (*VintageMarketDealInfo) Descriptor() ([]byte, []int) {
	return file_VintageMarketDealInfo_proto_rawDescGZIP(), []int{0}
}

func (x *VintageMarketDealInfo) GetTraderItemInfoMap() map[uint32]*VintageMarketTraderInfo {
	if x != nil {
		return x.TraderItemInfoMap
	}
	return nil
}

var File_VintageMarketDealInfo_proto protoreflect.FileDescriptor

var file_VintageMarketDealInfo_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x56, 0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x44,
	0x65, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x56,
	0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x54, 0x72, 0x61, 0x64,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x01, 0x0a,
	0x15, 0x56, 0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x44, 0x65,
	0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x5e, 0x0a, 0x14, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x56, 0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x44, 0x65, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x11, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x49,
	0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x1a, 0x5e, 0x0a, 0x16, 0x54, 0x72, 0x61, 0x64, 0x65, 0x72,
	0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x56, 0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_VintageMarketDealInfo_proto_rawDescOnce sync.Once
	file_VintageMarketDealInfo_proto_rawDescData = file_VintageMarketDealInfo_proto_rawDesc
)

func file_VintageMarketDealInfo_proto_rawDescGZIP() []byte {
	file_VintageMarketDealInfo_proto_rawDescOnce.Do(func() {
		file_VintageMarketDealInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_VintageMarketDealInfo_proto_rawDescData)
	})
	return file_VintageMarketDealInfo_proto_rawDescData
}

var file_VintageMarketDealInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_VintageMarketDealInfo_proto_goTypes = []interface{}{
	(*VintageMarketDealInfo)(nil),   // 0: VintageMarketDealInfo
	nil,                             // 1: VintageMarketDealInfo.TraderItemInfoMapEntry
	(*VintageMarketTraderInfo)(nil), // 2: VintageMarketTraderInfo
}
var file_VintageMarketDealInfo_proto_depIdxs = []int32{
	1, // 0: VintageMarketDealInfo.trader_item_info_map:type_name -> VintageMarketDealInfo.TraderItemInfoMapEntry
	2, // 1: VintageMarketDealInfo.TraderItemInfoMapEntry.value:type_name -> VintageMarketTraderInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_VintageMarketDealInfo_proto_init() }
func file_VintageMarketDealInfo_proto_init() {
	if File_VintageMarketDealInfo_proto != nil {
		return
	}
	file_VintageMarketTraderInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_VintageMarketDealInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VintageMarketDealInfo); i {
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
			RawDescriptor: file_VintageMarketDealInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_VintageMarketDealInfo_proto_goTypes,
		DependencyIndexes: file_VintageMarketDealInfo_proto_depIdxs,
		MessageInfos:      file_VintageMarketDealInfo_proto_msgTypes,
	}.Build()
	File_VintageMarketDealInfo_proto = out.File
	file_VintageMarketDealInfo_proto_rawDesc = nil
	file_VintageMarketDealInfo_proto_goTypes = nil
	file_VintageMarketDealInfo_proto_depIdxs = nil
}
