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
// source: GCGMsgDiceRoll.proto

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

type GCGMsgDiceRoll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ControllerId uint32            `protobuf:"varint,9,opt,name=controller_id,json=controllerId,proto3" json:"controller_id,omitempty"`
	DiceNum      uint32            `protobuf:"varint,3,opt,name=dice_num,json=diceNum,proto3" json:"dice_num,omitempty"`
	DiceSideList []GCGDiceSideType `protobuf:"varint,14,rep,packed,name=dice_side_list,json=diceSideList,proto3,enum=GCGDiceSideType" json:"dice_side_list,omitempty"`
}

func (x *GCGMsgDiceRoll) Reset() {
	*x = GCGMsgDiceRoll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GCGMsgDiceRoll_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GCGMsgDiceRoll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GCGMsgDiceRoll) ProtoMessage() {}

func (x *GCGMsgDiceRoll) ProtoReflect() protoreflect.Message {
	mi := &file_GCGMsgDiceRoll_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GCGMsgDiceRoll.ProtoReflect.Descriptor instead.
func (*GCGMsgDiceRoll) Descriptor() ([]byte, []int) {
	return file_GCGMsgDiceRoll_proto_rawDescGZIP(), []int{0}
}

func (x *GCGMsgDiceRoll) GetControllerId() uint32 {
	if x != nil {
		return x.ControllerId
	}
	return 0
}

func (x *GCGMsgDiceRoll) GetDiceNum() uint32 {
	if x != nil {
		return x.DiceNum
	}
	return 0
}

func (x *GCGMsgDiceRoll) GetDiceSideList() []GCGDiceSideType {
	if x != nil {
		return x.DiceSideList
	}
	return nil
}

var File_GCGMsgDiceRoll_proto protoreflect.FileDescriptor

var file_GCGMsgDiceRoll_proto_rawDesc = []byte{
	0x0a, 0x14, 0x47, 0x43, 0x47, 0x4d, 0x73, 0x67, 0x44, 0x69, 0x63, 0x65, 0x52, 0x6f, 0x6c, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x47, 0x43, 0x47, 0x44, 0x69, 0x63, 0x65, 0x53,
	0x69, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01,
	0x0a, 0x0e, 0x47, 0x43, 0x47, 0x4d, 0x73, 0x67, 0x44, 0x69, 0x63, 0x65, 0x52, 0x6f, 0x6c, 0x6c,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x75,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x64, 0x69, 0x63, 0x65, 0x4e, 0x75, 0x6d,
	0x12, 0x36, 0x0a, 0x0e, 0x64, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x47, 0x43, 0x47, 0x44, 0x69,
	0x63, 0x65, 0x53, 0x69, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x64, 0x69, 0x63, 0x65,
	0x53, 0x69, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GCGMsgDiceRoll_proto_rawDescOnce sync.Once
	file_GCGMsgDiceRoll_proto_rawDescData = file_GCGMsgDiceRoll_proto_rawDesc
)

func file_GCGMsgDiceRoll_proto_rawDescGZIP() []byte {
	file_GCGMsgDiceRoll_proto_rawDescOnce.Do(func() {
		file_GCGMsgDiceRoll_proto_rawDescData = protoimpl.X.CompressGZIP(file_GCGMsgDiceRoll_proto_rawDescData)
	})
	return file_GCGMsgDiceRoll_proto_rawDescData
}

var file_GCGMsgDiceRoll_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GCGMsgDiceRoll_proto_goTypes = []interface{}{
	(*GCGMsgDiceRoll)(nil), // 0: GCGMsgDiceRoll
	(GCGDiceSideType)(0),   // 1: GCGDiceSideType
}
var file_GCGMsgDiceRoll_proto_depIdxs = []int32{
	1, // 0: GCGMsgDiceRoll.dice_side_list:type_name -> GCGDiceSideType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GCGMsgDiceRoll_proto_init() }
func file_GCGMsgDiceRoll_proto_init() {
	if File_GCGMsgDiceRoll_proto != nil {
		return
	}
	file_GCGDiceSideType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GCGMsgDiceRoll_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GCGMsgDiceRoll); i {
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
			RawDescriptor: file_GCGMsgDiceRoll_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GCGMsgDiceRoll_proto_goTypes,
		DependencyIndexes: file_GCGMsgDiceRoll_proto_depIdxs,
		MessageInfos:      file_GCGMsgDiceRoll_proto_msgTypes,
	}.Build()
	File_GCGMsgDiceRoll_proto = out.File
	file_GCGMsgDiceRoll_proto_rawDesc = nil
	file_GCGMsgDiceRoll_proto_goTypes = nil
	file_GCGMsgDiceRoll_proto_depIdxs = nil
}
