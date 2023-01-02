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
// source: BartenderCompleteOrderReq.proto

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

type BartenderCompleteOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestId             uint32       `protobuf:"varint,14,opt,name=quest_id,json=questId,proto3" json:"quest_id,omitempty"`
	OptionalOrderList   []uint32     `protobuf:"varint,15,rep,packed,name=optional_order_list,json=optionalOrderList,proto3" json:"optional_order_list,omitempty"`
	Unk3300_NKMHHHHGHIF uint32       `protobuf:"varint,6,opt,name=Unk3300_NKMHHHHGHIF,json=Unk3300NKMHHHHGHIF,proto3" json:"Unk3300_NKMHHHHGHIF,omitempty"`
	Unk3300_LONHKJFDOND uint32       `protobuf:"varint,5,opt,name=Unk3300_LONHKJFDOND,json=Unk3300LONHKJFDOND,proto3" json:"Unk3300_LONHKJFDOND,omitempty"`
	Unk3300_POEGPFJLNGB uint32       `protobuf:"varint,11,opt,name=Unk3300_POEGPFJLNGB,json=Unk3300POEGPFJLNGB,proto3" json:"Unk3300_POEGPFJLNGB,omitempty"`
	IsViewFormula       bool         `protobuf:"varint,8,opt,name=is_view_formula,json=isViewFormula,proto3" json:"is_view_formula,omitempty"`
	ItemList            []*ItemParam `protobuf:"bytes,12,rep,name=item_list,json=itemList,proto3" json:"item_list,omitempty"`
}

func (x *BartenderCompleteOrderReq) Reset() {
	*x = BartenderCompleteOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BartenderCompleteOrderReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BartenderCompleteOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BartenderCompleteOrderReq) ProtoMessage() {}

func (x *BartenderCompleteOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_BartenderCompleteOrderReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BartenderCompleteOrderReq.ProtoReflect.Descriptor instead.
func (*BartenderCompleteOrderReq) Descriptor() ([]byte, []int) {
	return file_BartenderCompleteOrderReq_proto_rawDescGZIP(), []int{0}
}

func (x *BartenderCompleteOrderReq) GetQuestId() uint32 {
	if x != nil {
		return x.QuestId
	}
	return 0
}

func (x *BartenderCompleteOrderReq) GetOptionalOrderList() []uint32 {
	if x != nil {
		return x.OptionalOrderList
	}
	return nil
}

func (x *BartenderCompleteOrderReq) GetUnk3300_NKMHHHHGHIF() uint32 {
	if x != nil {
		return x.Unk3300_NKMHHHHGHIF
	}
	return 0
}

func (x *BartenderCompleteOrderReq) GetUnk3300_LONHKJFDOND() uint32 {
	if x != nil {
		return x.Unk3300_LONHKJFDOND
	}
	return 0
}

func (x *BartenderCompleteOrderReq) GetUnk3300_POEGPFJLNGB() uint32 {
	if x != nil {
		return x.Unk3300_POEGPFJLNGB
	}
	return 0
}

func (x *BartenderCompleteOrderReq) GetIsViewFormula() bool {
	if x != nil {
		return x.IsViewFormula
	}
	return false
}

func (x *BartenderCompleteOrderReq) GetItemList() []*ItemParam {
	if x != nil {
		return x.ItemList
	}
	return nil
}

var File_BartenderCompleteOrderReq_proto protoreflect.FileDescriptor

var file_BartenderCompleteOrderReq_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x42, 0x61, 0x72, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0f, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xca, 0x02, 0x0a, 0x19, 0x42, 0x61, 0x72, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x11, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x55,
	0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x4e, 0x4b, 0x4d, 0x48, 0x48, 0x48, 0x48, 0x47, 0x48,
	0x49, 0x46, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30,
	0x30, 0x4e, 0x4b, 0x4d, 0x48, 0x48, 0x48, 0x48, 0x47, 0x48, 0x49, 0x46, 0x12, 0x2f, 0x0a, 0x13,
	0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x4c, 0x4f, 0x4e, 0x48, 0x4b, 0x4a, 0x46, 0x44,
	0x4f, 0x4e, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33,
	0x30, 0x30, 0x4c, 0x4f, 0x4e, 0x48, 0x4b, 0x4a, 0x46, 0x44, 0x4f, 0x4e, 0x44, 0x12, 0x2f, 0x0a,
	0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x50, 0x4f, 0x45, 0x47, 0x50, 0x46, 0x4a,
	0x4c, 0x4e, 0x47, 0x42, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33,
	0x33, 0x30, 0x30, 0x50, 0x4f, 0x45, 0x47, 0x50, 0x46, 0x4a, 0x4c, 0x4e, 0x47, 0x42, 0x12, 0x26,
	0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x75, 0x6c,
	0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x56, 0x69, 0x65, 0x77, 0x46,
	0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x12, 0x27, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BartenderCompleteOrderReq_proto_rawDescOnce sync.Once
	file_BartenderCompleteOrderReq_proto_rawDescData = file_BartenderCompleteOrderReq_proto_rawDesc
)

func file_BartenderCompleteOrderReq_proto_rawDescGZIP() []byte {
	file_BartenderCompleteOrderReq_proto_rawDescOnce.Do(func() {
		file_BartenderCompleteOrderReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_BartenderCompleteOrderReq_proto_rawDescData)
	})
	return file_BartenderCompleteOrderReq_proto_rawDescData
}

var file_BartenderCompleteOrderReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BartenderCompleteOrderReq_proto_goTypes = []interface{}{
	(*BartenderCompleteOrderReq)(nil), // 0: BartenderCompleteOrderReq
	(*ItemParam)(nil),                 // 1: ItemParam
}
var file_BartenderCompleteOrderReq_proto_depIdxs = []int32{
	1, // 0: BartenderCompleteOrderReq.item_list:type_name -> ItemParam
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_BartenderCompleteOrderReq_proto_init() }
func file_BartenderCompleteOrderReq_proto_init() {
	if File_BartenderCompleteOrderReq_proto != nil {
		return
	}
	file_ItemParam_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BartenderCompleteOrderReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BartenderCompleteOrderReq); i {
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
			RawDescriptor: file_BartenderCompleteOrderReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BartenderCompleteOrderReq_proto_goTypes,
		DependencyIndexes: file_BartenderCompleteOrderReq_proto_depIdxs,
		MessageInfos:      file_BartenderCompleteOrderReq_proto_msgTypes,
	}.Build()
	File_BartenderCompleteOrderReq_proto = out.File
	file_BartenderCompleteOrderReq_proto_rawDesc = nil
	file_BartenderCompleteOrderReq_proto_goTypes = nil
	file_BartenderCompleteOrderReq_proto_depIdxs = nil
}
