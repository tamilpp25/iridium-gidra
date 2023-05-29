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
// source: Shop.proto

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

// Obf: JBIFADEKOFE
type Shop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextRefreshTime     uint32                `protobuf:"varint,2,opt,name=next_refresh_time,json=nextRefreshTime,proto3" json:"next_refresh_time,omitempty"`
	CardProductList     []*ShopCardProduct    `protobuf:"bytes,1,rep,name=card_product_list,json=cardProductList,proto3" json:"card_product_list,omitempty"`
	ShopType            uint32                `protobuf:"varint,5,opt,name=shop_type,json=shopType,proto3" json:"shop_type,omitempty"`
	McoinProductList    []*ShopMcoinProduct   `protobuf:"bytes,14,rep,name=mcoin_product_list,json=mcoinProductList,proto3" json:"mcoin_product_list,omitempty"`
	CityId              uint32                `protobuf:"varint,8,opt,name=city_id,json=cityId,proto3" json:"city_id,omitempty"`
	GoodsList           []*ShopGoods          `protobuf:"bytes,7,rep,name=goods_list,json=goodsList,proto3" json:"goods_list,omitempty"`
	CityReputationLevel uint32                `protobuf:"varint,10,opt,name=city_reputation_level,json=cityReputationLevel,proto3" json:"city_reputation_level,omitempty"`
	ConcertProductList  []*ShopConcertProduct `protobuf:"bytes,6,rep,name=concert_product_list,json=concertProductList,proto3" json:"concert_product_list,omitempty"`
}

func (x *Shop) Reset() {
	*x = Shop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Shop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shop) ProtoMessage() {}

func (x *Shop) ProtoReflect() protoreflect.Message {
	mi := &file_Shop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shop.ProtoReflect.Descriptor instead.
func (*Shop) Descriptor() ([]byte, []int) {
	return file_Shop_proto_rawDescGZIP(), []int{0}
}

func (x *Shop) GetNextRefreshTime() uint32 {
	if x != nil {
		return x.NextRefreshTime
	}
	return 0
}

func (x *Shop) GetCardProductList() []*ShopCardProduct {
	if x != nil {
		return x.CardProductList
	}
	return nil
}

func (x *Shop) GetShopType() uint32 {
	if x != nil {
		return x.ShopType
	}
	return 0
}

func (x *Shop) GetMcoinProductList() []*ShopMcoinProduct {
	if x != nil {
		return x.McoinProductList
	}
	return nil
}

func (x *Shop) GetCityId() uint32 {
	if x != nil {
		return x.CityId
	}
	return 0
}

func (x *Shop) GetGoodsList() []*ShopGoods {
	if x != nil {
		return x.GoodsList
	}
	return nil
}

func (x *Shop) GetCityReputationLevel() uint32 {
	if x != nil {
		return x.CityReputationLevel
	}
	return 0
}

func (x *Shop) GetConcertProductList() []*ShopConcertProduct {
	if x != nil {
		return x.ConcertProductList
	}
	return nil
}

var File_Shop_proto protoreflect.FileDescriptor

var file_Shop_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x53, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x53, 0x68,
	0x6f, 0x70, 0x43, 0x61, 0x72, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x53, 0x68, 0x6f, 0x70, 0x4d, 0x63, 0x6f, 0x69, 0x6e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x53, 0x68, 0x6f,
	0x70, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x53, 0x68,
	0x6f, 0x70, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x03, 0x0a, 0x04, 0x53, 0x68, 0x6f, 0x70, 0x12,
	0x2a, 0x0a, 0x11, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6e, 0x65, 0x78, 0x74,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x11, 0x63,
	0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x43, 0x61, 0x72,
	0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x0f, 0x63, 0x61, 0x72, 0x64, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f,
	0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x68,
	0x6f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3f, 0x0a, 0x12, 0x6d, 0x63, 0x6f, 0x69, 0x6e, 0x5f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x4d, 0x63, 0x6f, 0x69, 0x6e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x10, 0x6d, 0x63, 0x6f, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x69, 0x74, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x12, 0x29, 0x0a, 0x0a, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x73,
	0x52, 0x09, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x63,
	0x69, 0x74, 0x79, 0x5f, 0x72, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x63, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x45, 0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x53, 0x68, 0x6f, 0x70, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Shop_proto_rawDescOnce sync.Once
	file_Shop_proto_rawDescData = file_Shop_proto_rawDesc
)

func file_Shop_proto_rawDescGZIP() []byte {
	file_Shop_proto_rawDescOnce.Do(func() {
		file_Shop_proto_rawDescData = protoimpl.X.CompressGZIP(file_Shop_proto_rawDescData)
	})
	return file_Shop_proto_rawDescData
}

var file_Shop_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Shop_proto_goTypes = []interface{}{
	(*Shop)(nil),               // 0: Shop
	(*ShopCardProduct)(nil),    // 1: ShopCardProduct
	(*ShopMcoinProduct)(nil),   // 2: ShopMcoinProduct
	(*ShopGoods)(nil),          // 3: ShopGoods
	(*ShopConcertProduct)(nil), // 4: ShopConcertProduct
}
var file_Shop_proto_depIdxs = []int32{
	1, // 0: Shop.card_product_list:type_name -> ShopCardProduct
	2, // 1: Shop.mcoin_product_list:type_name -> ShopMcoinProduct
	3, // 2: Shop.goods_list:type_name -> ShopGoods
	4, // 3: Shop.concert_product_list:type_name -> ShopConcertProduct
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_Shop_proto_init() }
func file_Shop_proto_init() {
	if File_Shop_proto != nil {
		return
	}
	file_ShopCardProduct_proto_init()
	file_ShopMcoinProduct_proto_init()
	file_ShopGoods_proto_init()
	file_ShopConcertProduct_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Shop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shop); i {
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
			RawDescriptor: file_Shop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Shop_proto_goTypes,
		DependencyIndexes: file_Shop_proto_depIdxs,
		MessageInfos:      file_Shop_proto_msgTypes,
	}.Build()
	File_Shop_proto = out.File
	file_Shop_proto_rawDesc = nil
	file_Shop_proto_goTypes = nil
	file_Shop_proto_depIdxs = nil
}
