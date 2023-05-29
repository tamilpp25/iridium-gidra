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
// source: BigTalentPointConvertReq.proto

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

// CmdId: 1055
// Obf: MNCGEGPGLJE
type BigTalentPointConvertReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemGuidList []uint64 `protobuf:"varint,12,rep,packed,name=item_guid_list,json=itemGuidList,proto3" json:"item_guid_list,omitempty"`
	AvatarGuid   uint64   `protobuf:"varint,11,opt,name=avatar_guid,json=avatarGuid,proto3" json:"avatar_guid,omitempty"`
}

func (x *BigTalentPointConvertReq) Reset() {
	*x = BigTalentPointConvertReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BigTalentPointConvertReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BigTalentPointConvertReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigTalentPointConvertReq) ProtoMessage() {}

func (x *BigTalentPointConvertReq) ProtoReflect() protoreflect.Message {
	mi := &file_BigTalentPointConvertReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigTalentPointConvertReq.ProtoReflect.Descriptor instead.
func (*BigTalentPointConvertReq) Descriptor() ([]byte, []int) {
	return file_BigTalentPointConvertReq_proto_rawDescGZIP(), []int{0}
}

func (x *BigTalentPointConvertReq) GetItemGuidList() []uint64 {
	if x != nil {
		return x.ItemGuidList
	}
	return nil
}

func (x *BigTalentPointConvertReq) GetAvatarGuid() uint64 {
	if x != nil {
		return x.AvatarGuid
	}
	return 0
}

var File_BigTalentPointConvertReq_proto protoreflect.FileDescriptor

var file_BigTalentPointConvertReq_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x42, 0x69, 0x67, 0x54, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x61, 0x0a, 0x18, 0x42, 0x69, 0x67, 0x54, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x0e,
	0x69, 0x74, 0x65, 0x6d, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c,
	0x20, 0x03, 0x28, 0x04, 0x52, 0x0c, 0x69, 0x74, 0x65, 0x6d, 0x47, 0x75, 0x69, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x67, 0x75, 0x69,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x47,
	0x75, 0x69, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_BigTalentPointConvertReq_proto_rawDescOnce sync.Once
	file_BigTalentPointConvertReq_proto_rawDescData = file_BigTalentPointConvertReq_proto_rawDesc
)

func file_BigTalentPointConvertReq_proto_rawDescGZIP() []byte {
	file_BigTalentPointConvertReq_proto_rawDescOnce.Do(func() {
		file_BigTalentPointConvertReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_BigTalentPointConvertReq_proto_rawDescData)
	})
	return file_BigTalentPointConvertReq_proto_rawDescData
}

var file_BigTalentPointConvertReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BigTalentPointConvertReq_proto_goTypes = []interface{}{
	(*BigTalentPointConvertReq)(nil), // 0: BigTalentPointConvertReq
}
var file_BigTalentPointConvertReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_BigTalentPointConvertReq_proto_init() }
func file_BigTalentPointConvertReq_proto_init() {
	if File_BigTalentPointConvertReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_BigTalentPointConvertReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BigTalentPointConvertReq); i {
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
			RawDescriptor: file_BigTalentPointConvertReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BigTalentPointConvertReq_proto_goTypes,
		DependencyIndexes: file_BigTalentPointConvertReq_proto_depIdxs,
		MessageInfos:      file_BigTalentPointConvertReq_proto_msgTypes,
	}.Build()
	File_BigTalentPointConvertReq_proto = out.File
	file_BigTalentPointConvertReq_proto_rawDesc = nil
	file_BigTalentPointConvertReq_proto_goTypes = nil
	file_BigTalentPointConvertReq_proto_depIdxs = nil
}
