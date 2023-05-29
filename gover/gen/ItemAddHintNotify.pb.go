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
// source: ItemAddHintNotify.proto

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

// CmdId: 655
// Obf: FOCMGEBMEOH
type ItemAddHintNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GCNFBOCLHPE                 bool        `protobuf:"varint,7,opt,name=GCNFBOCLHPE,proto3" json:"GCNFBOCLHPE,omitempty"`
	PJLFNPEEBNP                 bool        `protobuf:"varint,5,opt,name=PJLFNPEEBNP,proto3" json:"PJLFNPEEBNP,omitempty"`
	OverflowTransformedItemList []*ItemHint `protobuf:"bytes,1,rep,name=overflow_transformed_item_list,json=overflowTransformedItemList,proto3" json:"overflow_transformed_item_list,omitempty"`
	ItemList                    []*ItemHint `protobuf:"bytes,13,rep,name=item_list,json=itemList,proto3" json:"item_list,omitempty"`
	QuestId                     uint32      `protobuf:"varint,15,opt,name=quest_id,json=questId,proto3" json:"quest_id,omitempty"`
	Position                    *Vector     `protobuf:"bytes,8,opt,name=position,proto3" json:"position,omitempty"`
	MMHKDGBFKAC                 bool        `protobuf:"varint,9,opt,name=MMHKDGBFKAC,proto3" json:"MMHKDGBFKAC,omitempty"`
	Reason                      uint32      `protobuf:"varint,4,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *ItemAddHintNotify) Reset() {
	*x = ItemAddHintNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ItemAddHintNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemAddHintNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemAddHintNotify) ProtoMessage() {}

func (x *ItemAddHintNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ItemAddHintNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemAddHintNotify.ProtoReflect.Descriptor instead.
func (*ItemAddHintNotify) Descriptor() ([]byte, []int) {
	return file_ItemAddHintNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ItemAddHintNotify) GetGCNFBOCLHPE() bool {
	if x != nil {
		return x.GCNFBOCLHPE
	}
	return false
}

func (x *ItemAddHintNotify) GetPJLFNPEEBNP() bool {
	if x != nil {
		return x.PJLFNPEEBNP
	}
	return false
}

func (x *ItemAddHintNotify) GetOverflowTransformedItemList() []*ItemHint {
	if x != nil {
		return x.OverflowTransformedItemList
	}
	return nil
}

func (x *ItemAddHintNotify) GetItemList() []*ItemHint {
	if x != nil {
		return x.ItemList
	}
	return nil
}

func (x *ItemAddHintNotify) GetQuestId() uint32 {
	if x != nil {
		return x.QuestId
	}
	return 0
}

func (x *ItemAddHintNotify) GetPosition() *Vector {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *ItemAddHintNotify) GetMMHKDGBFKAC() bool {
	if x != nil {
		return x.MMHKDGBFKAC
	}
	return false
}

func (x *ItemAddHintNotify) GetReason() uint32 {
	if x != nil {
		return x.Reason
	}
	return 0
}

var File_ItemAddHintNotify_proto protoreflect.FileDescriptor

var file_ItemAddHintNotify_proto_rawDesc = []byte{
	0x0a, 0x17, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x64, 0x64, 0x48, 0x69, 0x6e, 0x74, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x49, 0x74, 0x65, 0x6d, 0x48,
	0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x02, 0x0a, 0x11, 0x49, 0x74, 0x65, 0x6d,
	0x41, 0x64, 0x64, 0x48, 0x69, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x20, 0x0a,
	0x0b, 0x47, 0x43, 0x4e, 0x46, 0x42, 0x4f, 0x43, 0x4c, 0x48, 0x50, 0x45, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x47, 0x43, 0x4e, 0x46, 0x42, 0x4f, 0x43, 0x4c, 0x48, 0x50, 0x45, 0x12,
	0x20, 0x0a, 0x0b, 0x50, 0x4a, 0x4c, 0x46, 0x4e, 0x50, 0x45, 0x45, 0x42, 0x4e, 0x50, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x50, 0x4a, 0x4c, 0x46, 0x4e, 0x50, 0x45, 0x45, 0x42, 0x4e,
	0x50, 0x12, 0x4e, 0x0a, 0x1e, 0x6f, 0x76, 0x65, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x48, 0x69, 0x6e, 0x74, 0x52, 0x1b, 0x6f, 0x76, 0x65, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x26, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x48, 0x69, 0x6e, 0x74, 0x52,
	0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x4d, 0x48,
	0x4b, 0x44, 0x47, 0x42, 0x46, 0x4b, 0x41, 0x43, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x4d, 0x4d, 0x48, 0x4b, 0x44, 0x47, 0x42, 0x46, 0x4b, 0x41, 0x43, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_ItemAddHintNotify_proto_rawDescOnce sync.Once
	file_ItemAddHintNotify_proto_rawDescData = file_ItemAddHintNotify_proto_rawDesc
)

func file_ItemAddHintNotify_proto_rawDescGZIP() []byte {
	file_ItemAddHintNotify_proto_rawDescOnce.Do(func() {
		file_ItemAddHintNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ItemAddHintNotify_proto_rawDescData)
	})
	return file_ItemAddHintNotify_proto_rawDescData
}

var file_ItemAddHintNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ItemAddHintNotify_proto_goTypes = []interface{}{
	(*ItemAddHintNotify)(nil), // 0: ItemAddHintNotify
	(*ItemHint)(nil),          // 1: ItemHint
	(*Vector)(nil),            // 2: Vector
}
var file_ItemAddHintNotify_proto_depIdxs = []int32{
	1, // 0: ItemAddHintNotify.overflow_transformed_item_list:type_name -> ItemHint
	1, // 1: ItemAddHintNotify.item_list:type_name -> ItemHint
	2, // 2: ItemAddHintNotify.position:type_name -> Vector
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ItemAddHintNotify_proto_init() }
func file_ItemAddHintNotify_proto_init() {
	if File_ItemAddHintNotify_proto != nil {
		return
	}
	file_ItemHint_proto_init()
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ItemAddHintNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemAddHintNotify); i {
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
			RawDescriptor: file_ItemAddHintNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ItemAddHintNotify_proto_goTypes,
		DependencyIndexes: file_ItemAddHintNotify_proto_depIdxs,
		MessageInfos:      file_ItemAddHintNotify_proto_msgTypes,
	}.Build()
	File_ItemAddHintNotify_proto = out.File
	file_ItemAddHintNotify_proto_rawDesc = nil
	file_ItemAddHintNotify_proto_goTypes = nil
	file_ItemAddHintNotify_proto_depIdxs = nil
}
