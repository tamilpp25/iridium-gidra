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
// source: AJNGNNFKHGA.proto

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

// CmdId: 7596
type AJNGNNFKHGA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode     int32        `protobuf:"varint,7,opt,name=retcode,proto3" json:"retcode,omitempty"`
	DAJBJLAPDPF *FCHFBEOPIMF `protobuf:"bytes,15,opt,name=DAJBJLAPDPF,proto3" json:"DAJBJLAPDPF,omitempty"`
}

func (x *AJNGNNFKHGA) Reset() {
	*x = AJNGNNFKHGA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AJNGNNFKHGA_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AJNGNNFKHGA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AJNGNNFKHGA) ProtoMessage() {}

func (x *AJNGNNFKHGA) ProtoReflect() protoreflect.Message {
	mi := &file_AJNGNNFKHGA_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AJNGNNFKHGA.ProtoReflect.Descriptor instead.
func (*AJNGNNFKHGA) Descriptor() ([]byte, []int) {
	return file_AJNGNNFKHGA_proto_rawDescGZIP(), []int{0}
}

func (x *AJNGNNFKHGA) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *AJNGNNFKHGA) GetDAJBJLAPDPF() *FCHFBEOPIMF {
	if x != nil {
		return x.DAJBJLAPDPF
	}
	return nil
}

var File_AJNGNNFKHGA_proto protoreflect.FileDescriptor

var file_AJNGNNFKHGA_proto_rawDesc = []byte{
	0x0a, 0x11, 0x41, 0x4a, 0x4e, 0x47, 0x4e, 0x4e, 0x46, 0x4b, 0x48, 0x47, 0x41, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x43, 0x48, 0x46, 0x42, 0x45, 0x4f, 0x50, 0x49, 0x4d, 0x46,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x0b, 0x41, 0x4a, 0x4e, 0x47, 0x4e, 0x4e,
	0x46, 0x4b, 0x48, 0x47, 0x41, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x2e, 0x0a, 0x0b, 0x44, 0x41, 0x4a, 0x42, 0x4a, 0x4c, 0x41, 0x50, 0x44, 0x50, 0x46, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x46, 0x43, 0x48, 0x46, 0x42, 0x45, 0x4f, 0x50, 0x49,
	0x4d, 0x46, 0x52, 0x0b, 0x44, 0x41, 0x4a, 0x42, 0x4a, 0x4c, 0x41, 0x50, 0x44, 0x50, 0x46, 0x42,
	0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AJNGNNFKHGA_proto_rawDescOnce sync.Once
	file_AJNGNNFKHGA_proto_rawDescData = file_AJNGNNFKHGA_proto_rawDesc
)

func file_AJNGNNFKHGA_proto_rawDescGZIP() []byte {
	file_AJNGNNFKHGA_proto_rawDescOnce.Do(func() {
		file_AJNGNNFKHGA_proto_rawDescData = protoimpl.X.CompressGZIP(file_AJNGNNFKHGA_proto_rawDescData)
	})
	return file_AJNGNNFKHGA_proto_rawDescData
}

var file_AJNGNNFKHGA_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AJNGNNFKHGA_proto_goTypes = []interface{}{
	(*AJNGNNFKHGA)(nil), // 0: AJNGNNFKHGA
	(*FCHFBEOPIMF)(nil), // 1: FCHFBEOPIMF
}
var file_AJNGNNFKHGA_proto_depIdxs = []int32{
	1, // 0: AJNGNNFKHGA.DAJBJLAPDPF:type_name -> FCHFBEOPIMF
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_AJNGNNFKHGA_proto_init() }
func file_AJNGNNFKHGA_proto_init() {
	if File_AJNGNNFKHGA_proto != nil {
		return
	}
	file_FCHFBEOPIMF_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AJNGNNFKHGA_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AJNGNNFKHGA); i {
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
			RawDescriptor: file_AJNGNNFKHGA_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AJNGNNFKHGA_proto_goTypes,
		DependencyIndexes: file_AJNGNNFKHGA_proto_depIdxs,
		MessageInfos:      file_AJNGNNFKHGA_proto_msgTypes,
	}.Build()
	File_AJNGNNFKHGA_proto = out.File
	file_AJNGNNFKHGA_proto_rawDesc = nil
	file_AJNGNNFKHGA_proto_goTypes = nil
	file_AJNGNNFKHGA_proto_depIdxs = nil
}
