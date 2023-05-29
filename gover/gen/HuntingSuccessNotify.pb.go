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
// source: HuntingSuccessNotify.proto

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

// CmdId: 4316
// Obf: DOIEGEDKPLC
type HuntingSuccessNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HuntingPair *HuntingPair `protobuf:"bytes,15,opt,name=hunting_pair,json=huntingPair,proto3" json:"hunting_pair,omitempty"`
}

func (x *HuntingSuccessNotify) Reset() {
	*x = HuntingSuccessNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HuntingSuccessNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HuntingSuccessNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HuntingSuccessNotify) ProtoMessage() {}

func (x *HuntingSuccessNotify) ProtoReflect() protoreflect.Message {
	mi := &file_HuntingSuccessNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HuntingSuccessNotify.ProtoReflect.Descriptor instead.
func (*HuntingSuccessNotify) Descriptor() ([]byte, []int) {
	return file_HuntingSuccessNotify_proto_rawDescGZIP(), []int{0}
}

func (x *HuntingSuccessNotify) GetHuntingPair() *HuntingPair {
	if x != nil {
		return x.HuntingPair
	}
	return nil
}

var File_HuntingSuccessNotify_proto protoreflect.FileDescriptor

var file_HuntingSuccessNotify_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x48, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x48, 0x75,
	0x6e, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x69, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x47, 0x0a, 0x14, 0x48, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x2f, 0x0a, 0x0c, 0x68, 0x75, 0x6e, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x48, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x69, 0x72, 0x52, 0x0b, 0x68, 0x75, 0x6e,
	0x74, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x69, 0x72, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HuntingSuccessNotify_proto_rawDescOnce sync.Once
	file_HuntingSuccessNotify_proto_rawDescData = file_HuntingSuccessNotify_proto_rawDesc
)

func file_HuntingSuccessNotify_proto_rawDescGZIP() []byte {
	file_HuntingSuccessNotify_proto_rawDescOnce.Do(func() {
		file_HuntingSuccessNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_HuntingSuccessNotify_proto_rawDescData)
	})
	return file_HuntingSuccessNotify_proto_rawDescData
}

var file_HuntingSuccessNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HuntingSuccessNotify_proto_goTypes = []interface{}{
	(*HuntingSuccessNotify)(nil), // 0: HuntingSuccessNotify
	(*HuntingPair)(nil),          // 1: HuntingPair
}
var file_HuntingSuccessNotify_proto_depIdxs = []int32{
	1, // 0: HuntingSuccessNotify.hunting_pair:type_name -> HuntingPair
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_HuntingSuccessNotify_proto_init() }
func file_HuntingSuccessNotify_proto_init() {
	if File_HuntingSuccessNotify_proto != nil {
		return
	}
	file_HuntingPair_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HuntingSuccessNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HuntingSuccessNotify); i {
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
			RawDescriptor: file_HuntingSuccessNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HuntingSuccessNotify_proto_goTypes,
		DependencyIndexes: file_HuntingSuccessNotify_proto_depIdxs,
		MessageInfos:      file_HuntingSuccessNotify_proto_msgTypes,
	}.Build()
	File_HuntingSuccessNotify_proto = out.File
	file_HuntingSuccessNotify_proto_rawDesc = nil
	file_HuntingSuccessNotify_proto_goTypes = nil
	file_HuntingSuccessNotify_proto_depIdxs = nil
}
