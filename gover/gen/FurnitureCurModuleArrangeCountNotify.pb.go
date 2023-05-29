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
// source: FurnitureCurModuleArrangeCountNotify.proto

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

// CmdId: 4856
// Obf: MCDJNNGJBIO
type FurnitureCurModuleArrangeCountNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FurnitureArrangeCountList []*Uint32Pair `protobuf:"bytes,14,rep,name=furniture_arrange_count_list,json=furnitureArrangeCountList,proto3" json:"furniture_arrange_count_list,omitempty"`
}

func (x *FurnitureCurModuleArrangeCountNotify) Reset() {
	*x = FurnitureCurModuleArrangeCountNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FurnitureCurModuleArrangeCountNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FurnitureCurModuleArrangeCountNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FurnitureCurModuleArrangeCountNotify) ProtoMessage() {}

func (x *FurnitureCurModuleArrangeCountNotify) ProtoReflect() protoreflect.Message {
	mi := &file_FurnitureCurModuleArrangeCountNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FurnitureCurModuleArrangeCountNotify.ProtoReflect.Descriptor instead.
func (*FurnitureCurModuleArrangeCountNotify) Descriptor() ([]byte, []int) {
	return file_FurnitureCurModuleArrangeCountNotify_proto_rawDescGZIP(), []int{0}
}

func (x *FurnitureCurModuleArrangeCountNotify) GetFurnitureArrangeCountList() []*Uint32Pair {
	if x != nil {
		return x.FurnitureArrangeCountList
	}
	return nil
}

var File_FurnitureCurModuleArrangeCountNotify_proto protoreflect.FileDescriptor

var file_FurnitureCurModuleArrangeCountNotify_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x43, 0x75, 0x72, 0x4d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x41, 0x72, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x55, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x50, 0x61, 0x69, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74,
	0x0a, 0x24, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x43, 0x75, 0x72, 0x4d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x41, 0x72, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x4c, 0x0a, 0x1c, 0x66, 0x75, 0x72, 0x6e, 0x69, 0x74,
	0x75, 0x72, 0x65, 0x5f, 0x61, 0x72, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x55,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x50, 0x61, 0x69, 0x72, 0x52, 0x19, 0x66, 0x75, 0x72, 0x6e, 0x69,
	0x74, 0x75, 0x72, 0x65, 0x41, 0x72, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FurnitureCurModuleArrangeCountNotify_proto_rawDescOnce sync.Once
	file_FurnitureCurModuleArrangeCountNotify_proto_rawDescData = file_FurnitureCurModuleArrangeCountNotify_proto_rawDesc
)

func file_FurnitureCurModuleArrangeCountNotify_proto_rawDescGZIP() []byte {
	file_FurnitureCurModuleArrangeCountNotify_proto_rawDescOnce.Do(func() {
		file_FurnitureCurModuleArrangeCountNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_FurnitureCurModuleArrangeCountNotify_proto_rawDescData)
	})
	return file_FurnitureCurModuleArrangeCountNotify_proto_rawDescData
}

var file_FurnitureCurModuleArrangeCountNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FurnitureCurModuleArrangeCountNotify_proto_goTypes = []interface{}{
	(*FurnitureCurModuleArrangeCountNotify)(nil), // 0: FurnitureCurModuleArrangeCountNotify
	(*Uint32Pair)(nil),                           // 1: Uint32Pair
}
var file_FurnitureCurModuleArrangeCountNotify_proto_depIdxs = []int32{
	1, // 0: FurnitureCurModuleArrangeCountNotify.furniture_arrange_count_list:type_name -> Uint32Pair
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_FurnitureCurModuleArrangeCountNotify_proto_init() }
func file_FurnitureCurModuleArrangeCountNotify_proto_init() {
	if File_FurnitureCurModuleArrangeCountNotify_proto != nil {
		return
	}
	file_Uint32Pair_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FurnitureCurModuleArrangeCountNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FurnitureCurModuleArrangeCountNotify); i {
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
			RawDescriptor: file_FurnitureCurModuleArrangeCountNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FurnitureCurModuleArrangeCountNotify_proto_goTypes,
		DependencyIndexes: file_FurnitureCurModuleArrangeCountNotify_proto_depIdxs,
		MessageInfos:      file_FurnitureCurModuleArrangeCountNotify_proto_msgTypes,
	}.Build()
	File_FurnitureCurModuleArrangeCountNotify_proto = out.File
	file_FurnitureCurModuleArrangeCountNotify_proto_rawDesc = nil
	file_FurnitureCurModuleArrangeCountNotify_proto_goTypes = nil
	file_FurnitureCurModuleArrangeCountNotify_proto_depIdxs = nil
}
