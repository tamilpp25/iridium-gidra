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
// source: GCGMsgMoveCard.proto

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

type GCGMsgMoveCard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To           GCGZoneType `protobuf:"varint,9,opt,name=to,proto3,enum=GCGZoneType" json:"to,omitempty"`
	FailGuidList []uint32    `protobuf:"varint,2,rep,packed,name=fail_guid_list,json=failGuidList,proto3" json:"fail_guid_list,omitempty"`
	From         GCGZoneType `protobuf:"varint,14,opt,name=from,proto3,enum=GCGZoneType" json:"from,omitempty"`
	ControllerId uint32      `protobuf:"varint,4,opt,name=controller_id,json=controllerId,proto3" json:"controller_id,omitempty"`
	Reason       GCGReason   `protobuf:"varint,13,opt,name=reason,proto3,enum=GCGReason" json:"reason,omitempty"`
	CardGuidList []uint32    `protobuf:"varint,5,rep,packed,name=card_guid_list,json=cardGuidList,proto3" json:"card_guid_list,omitempty"`
}

func (x *GCGMsgMoveCard) Reset() {
	*x = GCGMsgMoveCard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GCGMsgMoveCard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GCGMsgMoveCard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GCGMsgMoveCard) ProtoMessage() {}

func (x *GCGMsgMoveCard) ProtoReflect() protoreflect.Message {
	mi := &file_GCGMsgMoveCard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GCGMsgMoveCard.ProtoReflect.Descriptor instead.
func (*GCGMsgMoveCard) Descriptor() ([]byte, []int) {
	return file_GCGMsgMoveCard_proto_rawDescGZIP(), []int{0}
}

func (x *GCGMsgMoveCard) GetTo() GCGZoneType {
	if x != nil {
		return x.To
	}
	return GCGZoneType_GCG_ZONE_TYPE_INVALID
}

func (x *GCGMsgMoveCard) GetFailGuidList() []uint32 {
	if x != nil {
		return x.FailGuidList
	}
	return nil
}

func (x *GCGMsgMoveCard) GetFrom() GCGZoneType {
	if x != nil {
		return x.From
	}
	return GCGZoneType_GCG_ZONE_TYPE_INVALID
}

func (x *GCGMsgMoveCard) GetControllerId() uint32 {
	if x != nil {
		return x.ControllerId
	}
	return 0
}

func (x *GCGMsgMoveCard) GetReason() GCGReason {
	if x != nil {
		return x.Reason
	}
	return GCGReason_GCG_REASON_DEFAULT
}

func (x *GCGMsgMoveCard) GetCardGuidList() []uint32 {
	if x != nil {
		return x.CardGuidList
	}
	return nil
}

var File_GCGMsgMoveCard_proto protoreflect.FileDescriptor

var file_GCGMsgMoveCard_proto_rawDesc = []byte{
	0x0a, 0x14, 0x47, 0x43, 0x47, 0x4d, 0x73, 0x67, 0x4d, 0x6f, 0x76, 0x65, 0x43, 0x61, 0x72, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x47, 0x43, 0x47, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x47, 0x43, 0x47, 0x5a, 0x6f, 0x6e, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x01, 0x0a, 0x0e, 0x47,
	0x43, 0x47, 0x4d, 0x73, 0x67, 0x4d, 0x6f, 0x76, 0x65, 0x43, 0x61, 0x72, 0x64, 0x12, 0x1c, 0x0a,
	0x02, 0x74, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x47, 0x43, 0x47, 0x5a,
	0x6f, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x24, 0x0a, 0x0e, 0x66,
	0x61, 0x69, 0x6c, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x0c, 0x66, 0x61, 0x69, 0x6c, 0x47, 0x75, 0x69, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0c, 0x2e, 0x47, 0x43, 0x47, 0x5a, 0x6f, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x47, 0x43, 0x47, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0e,
	0x63, 0x61, 0x72, 0x64, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x61, 0x72, 0x64, 0x47, 0x75, 0x69, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_GCGMsgMoveCard_proto_rawDescOnce sync.Once
	file_GCGMsgMoveCard_proto_rawDescData = file_GCGMsgMoveCard_proto_rawDesc
)

func file_GCGMsgMoveCard_proto_rawDescGZIP() []byte {
	file_GCGMsgMoveCard_proto_rawDescOnce.Do(func() {
		file_GCGMsgMoveCard_proto_rawDescData = protoimpl.X.CompressGZIP(file_GCGMsgMoveCard_proto_rawDescData)
	})
	return file_GCGMsgMoveCard_proto_rawDescData
}

var file_GCGMsgMoveCard_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GCGMsgMoveCard_proto_goTypes = []interface{}{
	(*GCGMsgMoveCard)(nil), // 0: GCGMsgMoveCard
	(GCGZoneType)(0),       // 1: GCGZoneType
	(GCGReason)(0),         // 2: GCGReason
}
var file_GCGMsgMoveCard_proto_depIdxs = []int32{
	1, // 0: GCGMsgMoveCard.to:type_name -> GCGZoneType
	1, // 1: GCGMsgMoveCard.from:type_name -> GCGZoneType
	2, // 2: GCGMsgMoveCard.reason:type_name -> GCGReason
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_GCGMsgMoveCard_proto_init() }
func file_GCGMsgMoveCard_proto_init() {
	if File_GCGMsgMoveCard_proto != nil {
		return
	}
	file_GCGReason_proto_init()
	file_GCGZoneType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GCGMsgMoveCard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GCGMsgMoveCard); i {
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
			RawDescriptor: file_GCGMsgMoveCard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GCGMsgMoveCard_proto_goTypes,
		DependencyIndexes: file_GCGMsgMoveCard_proto_depIdxs,
		MessageInfos:      file_GCGMsgMoveCard_proto_msgTypes,
	}.Build()
	File_GCGMsgMoveCard_proto = out.File
	file_GCGMsgMoveCard_proto_rawDesc = nil
	file_GCGMsgMoveCard_proto_goTypes = nil
	file_GCGMsgMoveCard_proto_depIdxs = nil
}
