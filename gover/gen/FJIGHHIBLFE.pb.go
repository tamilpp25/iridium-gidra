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
// source: FJIGHHIBLFE.proto

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

// CmdId: 7458
type FJIGHHIBLFE struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Duel        *GCGDuel          `protobuf:"bytes,11,opt,name=duel,proto3" json:"duel,omitempty"`
	MsgPackList []*GCGMessagePack `protobuf:"bytes,5,rep,name=msg_pack_list,json=msgPackList,proto3" json:"msg_pack_list,omitempty"`
	Retcode     int32             `protobuf:"varint,10,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *FJIGHHIBLFE) Reset() {
	*x = FJIGHHIBLFE{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FJIGHHIBLFE_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FJIGHHIBLFE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FJIGHHIBLFE) ProtoMessage() {}

func (x *FJIGHHIBLFE) ProtoReflect() protoreflect.Message {
	mi := &file_FJIGHHIBLFE_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FJIGHHIBLFE.ProtoReflect.Descriptor instead.
func (*FJIGHHIBLFE) Descriptor() ([]byte, []int) {
	return file_FJIGHHIBLFE_proto_rawDescGZIP(), []int{0}
}

func (x *FJIGHHIBLFE) GetDuel() *GCGDuel {
	if x != nil {
		return x.Duel
	}
	return nil
}

func (x *FJIGHHIBLFE) GetMsgPackList() []*GCGMessagePack {
	if x != nil {
		return x.MsgPackList
	}
	return nil
}

func (x *FJIGHHIBLFE) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_FJIGHHIBLFE_proto protoreflect.FileDescriptor

var file_FJIGHHIBLFE_proto_rawDesc = []byte{
	0x0a, 0x11, 0x46, 0x4a, 0x49, 0x47, 0x48, 0x48, 0x49, 0x42, 0x4c, 0x46, 0x45, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x47, 0x43, 0x47, 0x44, 0x75, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x14, 0x47, 0x43, 0x47, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x61,
	0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x0b, 0x46, 0x4a, 0x49, 0x47,
	0x48, 0x48, 0x49, 0x42, 0x4c, 0x46, 0x45, 0x12, 0x1c, 0x0a, 0x04, 0x64, 0x75, 0x65, 0x6c, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x47, 0x43, 0x47, 0x44, 0x75, 0x65, 0x6c, 0x52,
	0x04, 0x64, 0x75, 0x65, 0x6c, 0x12, 0x33, 0x0a, 0x0d, 0x6d, 0x73, 0x67, 0x5f, 0x70, 0x61, 0x63,
	0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x47,
	0x43, 0x47, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x0b, 0x6d,
	0x73, 0x67, 0x50, 0x61, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FJIGHHIBLFE_proto_rawDescOnce sync.Once
	file_FJIGHHIBLFE_proto_rawDescData = file_FJIGHHIBLFE_proto_rawDesc
)

func file_FJIGHHIBLFE_proto_rawDescGZIP() []byte {
	file_FJIGHHIBLFE_proto_rawDescOnce.Do(func() {
		file_FJIGHHIBLFE_proto_rawDescData = protoimpl.X.CompressGZIP(file_FJIGHHIBLFE_proto_rawDescData)
	})
	return file_FJIGHHIBLFE_proto_rawDescData
}

var file_FJIGHHIBLFE_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FJIGHHIBLFE_proto_goTypes = []interface{}{
	(*FJIGHHIBLFE)(nil),    // 0: FJIGHHIBLFE
	(*GCGDuel)(nil),        // 1: GCGDuel
	(*GCGMessagePack)(nil), // 2: GCGMessagePack
}
var file_FJIGHHIBLFE_proto_depIdxs = []int32{
	1, // 0: FJIGHHIBLFE.duel:type_name -> GCGDuel
	2, // 1: FJIGHHIBLFE.msg_pack_list:type_name -> GCGMessagePack
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_FJIGHHIBLFE_proto_init() }
func file_FJIGHHIBLFE_proto_init() {
	if File_FJIGHHIBLFE_proto != nil {
		return
	}
	file_GCGDuel_proto_init()
	file_GCGMessagePack_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FJIGHHIBLFE_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FJIGHHIBLFE); i {
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
			RawDescriptor: file_FJIGHHIBLFE_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FJIGHHIBLFE_proto_goTypes,
		DependencyIndexes: file_FJIGHHIBLFE_proto_depIdxs,
		MessageInfos:      file_FJIGHHIBLFE_proto_msgTypes,
	}.Build()
	File_FJIGHHIBLFE_proto = out.File
	file_FJIGHHIBLFE_proto_rawDesc = nil
	file_FJIGHHIBLFE_proto_goTypes = nil
	file_FJIGHHIBLFE_proto_depIdxs = nil
}
