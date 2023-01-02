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
// source: SummerTimeFloatSignalUpdateNotify.proto

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

type SummerTimeFloatSignalUpdateNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsTransferAnchor bool    `protobuf:"varint,1,opt,name=is_transfer_anchor,json=isTransferAnchor,proto3" json:"is_transfer_anchor,omitempty"`
	Position         *Vector `protobuf:"bytes,7,opt,name=position,proto3" json:"position,omitempty"`
	FloatSignalId    uint32  `protobuf:"varint,12,opt,name=float_signal_id,json=floatSignalId,proto3" json:"float_signal_id,omitempty"`
}

func (x *SummerTimeFloatSignalUpdateNotify) Reset() {
	*x = SummerTimeFloatSignalUpdateNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SummerTimeFloatSignalUpdateNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SummerTimeFloatSignalUpdateNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummerTimeFloatSignalUpdateNotify) ProtoMessage() {}

func (x *SummerTimeFloatSignalUpdateNotify) ProtoReflect() protoreflect.Message {
	mi := &file_SummerTimeFloatSignalUpdateNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummerTimeFloatSignalUpdateNotify.ProtoReflect.Descriptor instead.
func (*SummerTimeFloatSignalUpdateNotify) Descriptor() ([]byte, []int) {
	return file_SummerTimeFloatSignalUpdateNotify_proto_rawDescGZIP(), []int{0}
}

func (x *SummerTimeFloatSignalUpdateNotify) GetIsTransferAnchor() bool {
	if x != nil {
		return x.IsTransferAnchor
	}
	return false
}

func (x *SummerTimeFloatSignalUpdateNotify) GetPosition() *Vector {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *SummerTimeFloatSignalUpdateNotify) GetFloatSignalId() uint32 {
	if x != nil {
		return x.FloatSignalId
	}
	return 0
}

var File_SummerTimeFloatSignalUpdateNotify_proto protoreflect.FileDescriptor

var file_SummerTimeFloatSignalUpdateNotify_proto_rawDesc = []byte{
	0x0a, 0x27, 0x53, 0x75, 0x6d, 0x6d, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x6c, 0x6f, 0x61,
	0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01, 0x0a, 0x21, 0x53, 0x75, 0x6d, 0x6d,
	0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x2c, 0x0a,
	0x12, 0x69, 0x73, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x61, 0x6e, 0x63,
	0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x12, 0x23, 0x0a, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x26, 0x0a, 0x0f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SummerTimeFloatSignalUpdateNotify_proto_rawDescOnce sync.Once
	file_SummerTimeFloatSignalUpdateNotify_proto_rawDescData = file_SummerTimeFloatSignalUpdateNotify_proto_rawDesc
)

func file_SummerTimeFloatSignalUpdateNotify_proto_rawDescGZIP() []byte {
	file_SummerTimeFloatSignalUpdateNotify_proto_rawDescOnce.Do(func() {
		file_SummerTimeFloatSignalUpdateNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_SummerTimeFloatSignalUpdateNotify_proto_rawDescData)
	})
	return file_SummerTimeFloatSignalUpdateNotify_proto_rawDescData
}

var file_SummerTimeFloatSignalUpdateNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SummerTimeFloatSignalUpdateNotify_proto_goTypes = []interface{}{
	(*SummerTimeFloatSignalUpdateNotify)(nil), // 0: SummerTimeFloatSignalUpdateNotify
	(*Vector)(nil), // 1: Vector
}
var file_SummerTimeFloatSignalUpdateNotify_proto_depIdxs = []int32{
	1, // 0: SummerTimeFloatSignalUpdateNotify.position:type_name -> Vector
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SummerTimeFloatSignalUpdateNotify_proto_init() }
func file_SummerTimeFloatSignalUpdateNotify_proto_init() {
	if File_SummerTimeFloatSignalUpdateNotify_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SummerTimeFloatSignalUpdateNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SummerTimeFloatSignalUpdateNotify); i {
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
			RawDescriptor: file_SummerTimeFloatSignalUpdateNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SummerTimeFloatSignalUpdateNotify_proto_goTypes,
		DependencyIndexes: file_SummerTimeFloatSignalUpdateNotify_proto_depIdxs,
		MessageInfos:      file_SummerTimeFloatSignalUpdateNotify_proto_msgTypes,
	}.Build()
	File_SummerTimeFloatSignalUpdateNotify_proto = out.File
	file_SummerTimeFloatSignalUpdateNotify_proto_rawDesc = nil
	file_SummerTimeFloatSignalUpdateNotify_proto_goTypes = nil
	file_SummerTimeFloatSignalUpdateNotify_proto_depIdxs = nil
}
