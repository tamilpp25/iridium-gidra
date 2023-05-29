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
// source: GetCityReputationMapInfoRsp.proto

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

// CmdId: 2821
// Obf: FOCHOIHCBHJ
type GetCityReputationMapInfoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NNJLDLAOPIE bool     `protobuf:"varint,10,opt,name=NNJLDLAOPIE,proto3" json:"NNJLDLAOPIE,omitempty"`
	Retcode     int32    `protobuf:"varint,4,opt,name=retcode,proto3" json:"retcode,omitempty"`
	BNOMCHCNPFP []uint32 `protobuf:"varint,15,rep,packed,name=BNOMCHCNPFP,proto3" json:"BNOMCHCNPFP,omitempty"`
	JCBEMEPIHGH []uint32 `protobuf:"varint,14,rep,packed,name=JCBEMEPIHGH,proto3" json:"JCBEMEPIHGH,omitempty"`
	JGAFDIAMBKO bool     `protobuf:"varint,8,opt,name=JGAFDIAMBKO,proto3" json:"JGAFDIAMBKO,omitempty"`
}

func (x *GetCityReputationMapInfoRsp) Reset() {
	*x = GetCityReputationMapInfoRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetCityReputationMapInfoRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCityReputationMapInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCityReputationMapInfoRsp) ProtoMessage() {}

func (x *GetCityReputationMapInfoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetCityReputationMapInfoRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCityReputationMapInfoRsp.ProtoReflect.Descriptor instead.
func (*GetCityReputationMapInfoRsp) Descriptor() ([]byte, []int) {
	return file_GetCityReputationMapInfoRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetCityReputationMapInfoRsp) GetNNJLDLAOPIE() bool {
	if x != nil {
		return x.NNJLDLAOPIE
	}
	return false
}

func (x *GetCityReputationMapInfoRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetCityReputationMapInfoRsp) GetBNOMCHCNPFP() []uint32 {
	if x != nil {
		return x.BNOMCHCNPFP
	}
	return nil
}

func (x *GetCityReputationMapInfoRsp) GetJCBEMEPIHGH() []uint32 {
	if x != nil {
		return x.JCBEMEPIHGH
	}
	return nil
}

func (x *GetCityReputationMapInfoRsp) GetJGAFDIAMBKO() bool {
	if x != nil {
		return x.JGAFDIAMBKO
	}
	return false
}

var File_GetCityReputationMapInfoRsp_proto protoreflect.FileDescriptor

var file_GetCityReputationMapInfoRsp_proto_rawDesc = []byte{
	0x0a, 0x21, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x4e, 0x4a, 0x4c, 0x44, 0x4c, 0x41, 0x4f, 0x50,
	0x49, 0x45, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4e, 0x4e, 0x4a, 0x4c, 0x44, 0x4c,
	0x41, 0x4f, 0x50, 0x49, 0x45, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x42, 0x4e, 0x4f, 0x4d, 0x43, 0x48, 0x43, 0x4e, 0x50, 0x46, 0x50, 0x18, 0x0f,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x42, 0x4e, 0x4f, 0x4d, 0x43, 0x48, 0x43, 0x4e, 0x50, 0x46,
	0x50, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x43, 0x42, 0x45, 0x4d, 0x45, 0x50, 0x49, 0x48, 0x47, 0x48,
	0x18, 0x0e, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x4a, 0x43, 0x42, 0x45, 0x4d, 0x45, 0x50, 0x49,
	0x48, 0x47, 0x48, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x47, 0x41, 0x46, 0x44, 0x49, 0x41, 0x4d, 0x42,
	0x4b, 0x4f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4a, 0x47, 0x41, 0x46, 0x44, 0x49,
	0x41, 0x4d, 0x42, 0x4b, 0x4f, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetCityReputationMapInfoRsp_proto_rawDescOnce sync.Once
	file_GetCityReputationMapInfoRsp_proto_rawDescData = file_GetCityReputationMapInfoRsp_proto_rawDesc
)

func file_GetCityReputationMapInfoRsp_proto_rawDescGZIP() []byte {
	file_GetCityReputationMapInfoRsp_proto_rawDescOnce.Do(func() {
		file_GetCityReputationMapInfoRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetCityReputationMapInfoRsp_proto_rawDescData)
	})
	return file_GetCityReputationMapInfoRsp_proto_rawDescData
}

var file_GetCityReputationMapInfoRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetCityReputationMapInfoRsp_proto_goTypes = []interface{}{
	(*GetCityReputationMapInfoRsp)(nil), // 0: GetCityReputationMapInfoRsp
}
var file_GetCityReputationMapInfoRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetCityReputationMapInfoRsp_proto_init() }
func file_GetCityReputationMapInfoRsp_proto_init() {
	if File_GetCityReputationMapInfoRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetCityReputationMapInfoRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCityReputationMapInfoRsp); i {
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
			RawDescriptor: file_GetCityReputationMapInfoRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetCityReputationMapInfoRsp_proto_goTypes,
		DependencyIndexes: file_GetCityReputationMapInfoRsp_proto_depIdxs,
		MessageInfos:      file_GetCityReputationMapInfoRsp_proto_msgTypes,
	}.Build()
	File_GetCityReputationMapInfoRsp_proto = out.File
	file_GetCityReputationMapInfoRsp_proto_rawDesc = nil
	file_GetCityReputationMapInfoRsp_proto_goTypes = nil
	file_GetCityReputationMapInfoRsp_proto_depIdxs = nil
}
