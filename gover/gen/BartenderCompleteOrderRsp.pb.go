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
// source: BartenderCompleteOrderRsp.proto

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

// CmdId: 8388
// Obf: DFNFGBPBAPO
type BartenderCompleteOrderRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode       int32    `protobuf:"varint,8,opt,name=retcode,proto3" json:"retcode,omitempty"`
	AffixList     []uint32 `protobuf:"varint,9,rep,packed,name=affix_list,json=affixList,proto3" json:"affix_list,omitempty"`
	IsNew         bool     `protobuf:"varint,5,opt,name=is_new,json=isNew,proto3" json:"is_new,omitempty"`
	FinishOrderId uint32   `protobuf:"varint,12,opt,name=finish_order_id,json=finishOrderId,proto3" json:"finish_order_id,omitempty"`
	QuestId       uint32   `protobuf:"varint,11,opt,name=quest_id,json=questId,proto3" json:"quest_id,omitempty"`
	FormulaId     uint32   `protobuf:"varint,4,opt,name=formula_id,json=formulaId,proto3" json:"formula_id,omitempty"`
}

func (x *BartenderCompleteOrderRsp) Reset() {
	*x = BartenderCompleteOrderRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BartenderCompleteOrderRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BartenderCompleteOrderRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BartenderCompleteOrderRsp) ProtoMessage() {}

func (x *BartenderCompleteOrderRsp) ProtoReflect() protoreflect.Message {
	mi := &file_BartenderCompleteOrderRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BartenderCompleteOrderRsp.ProtoReflect.Descriptor instead.
func (*BartenderCompleteOrderRsp) Descriptor() ([]byte, []int) {
	return file_BartenderCompleteOrderRsp_proto_rawDescGZIP(), []int{0}
}

func (x *BartenderCompleteOrderRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *BartenderCompleteOrderRsp) GetAffixList() []uint32 {
	if x != nil {
		return x.AffixList
	}
	return nil
}

func (x *BartenderCompleteOrderRsp) GetIsNew() bool {
	if x != nil {
		return x.IsNew
	}
	return false
}

func (x *BartenderCompleteOrderRsp) GetFinishOrderId() uint32 {
	if x != nil {
		return x.FinishOrderId
	}
	return 0
}

func (x *BartenderCompleteOrderRsp) GetQuestId() uint32 {
	if x != nil {
		return x.QuestId
	}
	return 0
}

func (x *BartenderCompleteOrderRsp) GetFormulaId() uint32 {
	if x != nil {
		return x.FormulaId
	}
	return 0
}

var File_BartenderCompleteOrderRsp_proto protoreflect.FileDescriptor

var file_BartenderCompleteOrderRsp_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x42, 0x61, 0x72, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x19, 0x42, 0x61, 0x72, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x73, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x66, 0x66,
	0x69, 0x78, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x09, 0x61,
	0x66, 0x66, 0x69, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x6e,
	0x65, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x4e, 0x65, 0x77, 0x12,
	0x26, 0x0a, 0x0f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x66, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x49,
	0x64, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_BartenderCompleteOrderRsp_proto_rawDescOnce sync.Once
	file_BartenderCompleteOrderRsp_proto_rawDescData = file_BartenderCompleteOrderRsp_proto_rawDesc
)

func file_BartenderCompleteOrderRsp_proto_rawDescGZIP() []byte {
	file_BartenderCompleteOrderRsp_proto_rawDescOnce.Do(func() {
		file_BartenderCompleteOrderRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_BartenderCompleteOrderRsp_proto_rawDescData)
	})
	return file_BartenderCompleteOrderRsp_proto_rawDescData
}

var file_BartenderCompleteOrderRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BartenderCompleteOrderRsp_proto_goTypes = []interface{}{
	(*BartenderCompleteOrderRsp)(nil), // 0: BartenderCompleteOrderRsp
}
var file_BartenderCompleteOrderRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_BartenderCompleteOrderRsp_proto_init() }
func file_BartenderCompleteOrderRsp_proto_init() {
	if File_BartenderCompleteOrderRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_BartenderCompleteOrderRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BartenderCompleteOrderRsp); i {
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
			RawDescriptor: file_BartenderCompleteOrderRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BartenderCompleteOrderRsp_proto_goTypes,
		DependencyIndexes: file_BartenderCompleteOrderRsp_proto_depIdxs,
		MessageInfos:      file_BartenderCompleteOrderRsp_proto_msgTypes,
	}.Build()
	File_BartenderCompleteOrderRsp_proto = out.File
	file_BartenderCompleteOrderRsp_proto_rawDesc = nil
	file_BartenderCompleteOrderRsp_proto_goTypes = nil
	file_BartenderCompleteOrderRsp_proto_depIdxs = nil
}
