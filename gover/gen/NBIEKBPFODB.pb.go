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
// source: NBIEKBPFODB.proto

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

type NBIEKBPFODB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HABECPAHOGI uint32         `protobuf:"varint,5,opt,name=HABECPAHOGI,proto3" json:"HABECPAHOGI,omitempty"`
	EIOJAKIDLJI uint32         `protobuf:"varint,8,opt,name=EIOJAKIDLJI,proto3" json:"EIOJAKIDLJI,omitempty"`
	DHCMGPFOEON []uint32       `protobuf:"varint,6,rep,packed,name=DHCMGPFOEON,proto3" json:"DHCMGPFOEON,omitempty"`
	Round       uint32         `protobuf:"varint,7,opt,name=round,proto3" json:"round,omitempty"`
	EFHBHDPAKLC []*PIINAODJAGC `protobuf:"bytes,2,rep,name=EFHBHDPAKLC,proto3" json:"EFHBHDPAKLC,omitempty"`
	ExpireTime  uint32         `protobuf:"fixed32,15,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
}

func (x *NBIEKBPFODB) Reset() {
	*x = NBIEKBPFODB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NBIEKBPFODB_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NBIEKBPFODB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NBIEKBPFODB) ProtoMessage() {}

func (x *NBIEKBPFODB) ProtoReflect() protoreflect.Message {
	mi := &file_NBIEKBPFODB_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NBIEKBPFODB.ProtoReflect.Descriptor instead.
func (*NBIEKBPFODB) Descriptor() ([]byte, []int) {
	return file_NBIEKBPFODB_proto_rawDescGZIP(), []int{0}
}

func (x *NBIEKBPFODB) GetHABECPAHOGI() uint32 {
	if x != nil {
		return x.HABECPAHOGI
	}
	return 0
}

func (x *NBIEKBPFODB) GetEIOJAKIDLJI() uint32 {
	if x != nil {
		return x.EIOJAKIDLJI
	}
	return 0
}

func (x *NBIEKBPFODB) GetDHCMGPFOEON() []uint32 {
	if x != nil {
		return x.DHCMGPFOEON
	}
	return nil
}

func (x *NBIEKBPFODB) GetRound() uint32 {
	if x != nil {
		return x.Round
	}
	return 0
}

func (x *NBIEKBPFODB) GetEFHBHDPAKLC() []*PIINAODJAGC {
	if x != nil {
		return x.EFHBHDPAKLC
	}
	return nil
}

func (x *NBIEKBPFODB) GetExpireTime() uint32 {
	if x != nil {
		return x.ExpireTime
	}
	return 0
}

var File_NBIEKBPFODB_proto protoreflect.FileDescriptor

var file_NBIEKBPFODB_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4e, 0x42, 0x49, 0x45, 0x4b, 0x42, 0x50, 0x46, 0x4f, 0x44, 0x42, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x50, 0x49, 0x49, 0x4e, 0x41, 0x4f, 0x44, 0x4a, 0x41, 0x47, 0x43,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x01, 0x0a, 0x0b, 0x4e, 0x42, 0x49, 0x45, 0x4b,
	0x42, 0x50, 0x46, 0x4f, 0x44, 0x42, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x41, 0x42, 0x45, 0x43, 0x50,
	0x41, 0x48, 0x4f, 0x47, 0x49, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x48, 0x41, 0x42,
	0x45, 0x43, 0x50, 0x41, 0x48, 0x4f, 0x47, 0x49, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x49, 0x4f, 0x4a,
	0x41, 0x4b, 0x49, 0x44, 0x4c, 0x4a, 0x49, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x45,
	0x49, 0x4f, 0x4a, 0x41, 0x4b, 0x49, 0x44, 0x4c, 0x4a, 0x49, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x48,
	0x43, 0x4d, 0x47, 0x50, 0x46, 0x4f, 0x45, 0x4f, 0x4e, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x0b, 0x44, 0x48, 0x43, 0x4d, 0x47, 0x50, 0x46, 0x4f, 0x45, 0x4f, 0x4e, 0x12, 0x14, 0x0a, 0x05,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x12, 0x2e, 0x0a, 0x0b, 0x45, 0x46, 0x48, 0x42, 0x48, 0x44, 0x50, 0x41, 0x4b, 0x4c,
	0x43, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x49, 0x49, 0x4e, 0x41, 0x4f,
	0x44, 0x4a, 0x41, 0x47, 0x43, 0x52, 0x0b, 0x45, 0x46, 0x48, 0x42, 0x48, 0x44, 0x50, 0x41, 0x4b,
	0x4c, 0x43, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x07, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_NBIEKBPFODB_proto_rawDescOnce sync.Once
	file_NBIEKBPFODB_proto_rawDescData = file_NBIEKBPFODB_proto_rawDesc
)

func file_NBIEKBPFODB_proto_rawDescGZIP() []byte {
	file_NBIEKBPFODB_proto_rawDescOnce.Do(func() {
		file_NBIEKBPFODB_proto_rawDescData = protoimpl.X.CompressGZIP(file_NBIEKBPFODB_proto_rawDescData)
	})
	return file_NBIEKBPFODB_proto_rawDescData
}

var file_NBIEKBPFODB_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_NBIEKBPFODB_proto_goTypes = []interface{}{
	(*NBIEKBPFODB)(nil), // 0: NBIEKBPFODB
	(*PIINAODJAGC)(nil), // 1: PIINAODJAGC
}
var file_NBIEKBPFODB_proto_depIdxs = []int32{
	1, // 0: NBIEKBPFODB.EFHBHDPAKLC:type_name -> PIINAODJAGC
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_NBIEKBPFODB_proto_init() }
func file_NBIEKBPFODB_proto_init() {
	if File_NBIEKBPFODB_proto != nil {
		return
	}
	file_PIINAODJAGC_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_NBIEKBPFODB_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NBIEKBPFODB); i {
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
			RawDescriptor: file_NBIEKBPFODB_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_NBIEKBPFODB_proto_goTypes,
		DependencyIndexes: file_NBIEKBPFODB_proto_depIdxs,
		MessageInfos:      file_NBIEKBPFODB_proto_msgTypes,
	}.Build()
	File_NBIEKBPFODB_proto = out.File
	file_NBIEKBPFODB_proto_rawDesc = nil
	file_NBIEKBPFODB_proto_goTypes = nil
	file_NBIEKBPFODB_proto_depIdxs = nil
}
