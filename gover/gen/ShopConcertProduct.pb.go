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
// source: ShopConcertProduct.proto

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

// Obf: CKIBKCAKBNI
type ShopConcertProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId   string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	PriceTier   string `protobuf:"bytes,2,opt,name=price_tier,json=priceTier,proto3" json:"price_tier,omitempty"`
	ObtainCount uint32 `protobuf:"varint,3,opt,name=obtain_count,json=obtainCount,proto3" json:"obtain_count,omitempty"`
	ObtainLimit uint32 `protobuf:"varint,4,opt,name=obtain_limit,json=obtainLimit,proto3" json:"obtain_limit,omitempty"`
	BeginTime   uint32 `protobuf:"varint,5,opt,name=begin_time,json=beginTime,proto3" json:"begin_time,omitempty"`
	EndTime     uint32 `protobuf:"varint,6,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	BuyTimes    uint32 `protobuf:"varint,7,opt,name=buy_times,json=buyTimes,proto3" json:"buy_times,omitempty"`
}

func (x *ShopConcertProduct) Reset() {
	*x = ShopConcertProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ShopConcertProduct_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopConcertProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopConcertProduct) ProtoMessage() {}

func (x *ShopConcertProduct) ProtoReflect() protoreflect.Message {
	mi := &file_ShopConcertProduct_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopConcertProduct.ProtoReflect.Descriptor instead.
func (*ShopConcertProduct) Descriptor() ([]byte, []int) {
	return file_ShopConcertProduct_proto_rawDescGZIP(), []int{0}
}

func (x *ShopConcertProduct) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *ShopConcertProduct) GetPriceTier() string {
	if x != nil {
		return x.PriceTier
	}
	return ""
}

func (x *ShopConcertProduct) GetObtainCount() uint32 {
	if x != nil {
		return x.ObtainCount
	}
	return 0
}

func (x *ShopConcertProduct) GetObtainLimit() uint32 {
	if x != nil {
		return x.ObtainLimit
	}
	return 0
}

func (x *ShopConcertProduct) GetBeginTime() uint32 {
	if x != nil {
		return x.BeginTime
	}
	return 0
}

func (x *ShopConcertProduct) GetEndTime() uint32 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *ShopConcertProduct) GetBuyTimes() uint32 {
	if x != nil {
		return x.BuyTimes
	}
	return 0
}

var File_ShopConcertProduct_proto protoreflect.FileDescriptor

var file_ShopConcertProduct_proto_rawDesc = []byte{
	0x0a, 0x18, 0x53, 0x68, 0x6f, 0x70, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x72, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x01, 0x0a, 0x12, 0x53,
	0x68, 0x6f, 0x70, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x69, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x54, 0x69, 0x65, 0x72, 0x12,
	0x21, 0x0a, 0x0c, 0x6f, 0x62, 0x74, 0x61, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6f, 0x62, 0x74, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x62, 0x74, 0x61, 0x69, 0x6e, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6f, 0x62, 0x74, 0x61, 0x69, 0x6e,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x62, 0x75, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x62, 0x75, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x42, 0x06, 0x5a, 0x04,
	0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ShopConcertProduct_proto_rawDescOnce sync.Once
	file_ShopConcertProduct_proto_rawDescData = file_ShopConcertProduct_proto_rawDesc
)

func file_ShopConcertProduct_proto_rawDescGZIP() []byte {
	file_ShopConcertProduct_proto_rawDescOnce.Do(func() {
		file_ShopConcertProduct_proto_rawDescData = protoimpl.X.CompressGZIP(file_ShopConcertProduct_proto_rawDescData)
	})
	return file_ShopConcertProduct_proto_rawDescData
}

var file_ShopConcertProduct_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ShopConcertProduct_proto_goTypes = []interface{}{
	(*ShopConcertProduct)(nil), // 0: ShopConcertProduct
}
var file_ShopConcertProduct_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ShopConcertProduct_proto_init() }
func file_ShopConcertProduct_proto_init() {
	if File_ShopConcertProduct_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ShopConcertProduct_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShopConcertProduct); i {
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
			RawDescriptor: file_ShopConcertProduct_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ShopConcertProduct_proto_goTypes,
		DependencyIndexes: file_ShopConcertProduct_proto_depIdxs,
		MessageInfos:      file_ShopConcertProduct_proto_msgTypes,
	}.Build()
	File_ShopConcertProduct_proto = out.File
	file_ShopConcertProduct_proto_rawDesc = nil
	file_ShopConcertProduct_proto_goTypes = nil
	file_ShopConcertProduct_proto_depIdxs = nil
}
