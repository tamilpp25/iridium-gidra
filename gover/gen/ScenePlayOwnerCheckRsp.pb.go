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
// source: ScenePlayOwnerCheckRsp.proto

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

// CmdId: 4352
// Obf: MNHNFNBFDMG
type ScenePlayOwnerCheckRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayId      uint32   `protobuf:"varint,12,opt,name=play_id,json=playId,proto3" json:"play_id,omitempty"`
	IsSkipMatch bool     `protobuf:"varint,2,opt,name=is_skip_match,json=isSkipMatch,proto3" json:"is_skip_match,omitempty"`
	Retcode     int32    `protobuf:"varint,10,opt,name=retcode,proto3" json:"retcode,omitempty"`
	WrongUid    uint32   `protobuf:"varint,8,opt,name=wrong_uid,json=wrongUid,proto3" json:"wrong_uid,omitempty"`
	ParamList   []uint32 `protobuf:"varint,4,rep,packed,name=param_list,json=paramList,proto3" json:"param_list,omitempty"`
}

func (x *ScenePlayOwnerCheckRsp) Reset() {
	*x = ScenePlayOwnerCheckRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ScenePlayOwnerCheckRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScenePlayOwnerCheckRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScenePlayOwnerCheckRsp) ProtoMessage() {}

func (x *ScenePlayOwnerCheckRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ScenePlayOwnerCheckRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScenePlayOwnerCheckRsp.ProtoReflect.Descriptor instead.
func (*ScenePlayOwnerCheckRsp) Descriptor() ([]byte, []int) {
	return file_ScenePlayOwnerCheckRsp_proto_rawDescGZIP(), []int{0}
}

func (x *ScenePlayOwnerCheckRsp) GetPlayId() uint32 {
	if x != nil {
		return x.PlayId
	}
	return 0
}

func (x *ScenePlayOwnerCheckRsp) GetIsSkipMatch() bool {
	if x != nil {
		return x.IsSkipMatch
	}
	return false
}

func (x *ScenePlayOwnerCheckRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *ScenePlayOwnerCheckRsp) GetWrongUid() uint32 {
	if x != nil {
		return x.WrongUid
	}
	return 0
}

func (x *ScenePlayOwnerCheckRsp) GetParamList() []uint32 {
	if x != nil {
		return x.ParamList
	}
	return nil
}

var File_ScenePlayOwnerCheckRsp_proto protoreflect.FileDescriptor

var file_ScenePlayOwnerCheckRsp_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab,
	0x01, 0x0a, 0x16, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x73, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79,
	0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x53, 0x6b, 0x69,
	0x70, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x77, 0x72, 0x6f, 0x6e, 0x67, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x77, 0x72, 0x6f, 0x6e, 0x67, 0x55, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04,
	0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ScenePlayOwnerCheckRsp_proto_rawDescOnce sync.Once
	file_ScenePlayOwnerCheckRsp_proto_rawDescData = file_ScenePlayOwnerCheckRsp_proto_rawDesc
)

func file_ScenePlayOwnerCheckRsp_proto_rawDescGZIP() []byte {
	file_ScenePlayOwnerCheckRsp_proto_rawDescOnce.Do(func() {
		file_ScenePlayOwnerCheckRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_ScenePlayOwnerCheckRsp_proto_rawDescData)
	})
	return file_ScenePlayOwnerCheckRsp_proto_rawDescData
}

var file_ScenePlayOwnerCheckRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ScenePlayOwnerCheckRsp_proto_goTypes = []interface{}{
	(*ScenePlayOwnerCheckRsp)(nil), // 0: ScenePlayOwnerCheckRsp
}
var file_ScenePlayOwnerCheckRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ScenePlayOwnerCheckRsp_proto_init() }
func file_ScenePlayOwnerCheckRsp_proto_init() {
	if File_ScenePlayOwnerCheckRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ScenePlayOwnerCheckRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScenePlayOwnerCheckRsp); i {
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
			RawDescriptor: file_ScenePlayOwnerCheckRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ScenePlayOwnerCheckRsp_proto_goTypes,
		DependencyIndexes: file_ScenePlayOwnerCheckRsp_proto_depIdxs,
		MessageInfos:      file_ScenePlayOwnerCheckRsp_proto_msgTypes,
	}.Build()
	File_ScenePlayOwnerCheckRsp_proto = out.File
	file_ScenePlayOwnerCheckRsp_proto_rawDesc = nil
	file_ScenePlayOwnerCheckRsp_proto_goTypes = nil
	file_ScenePlayOwnerCheckRsp_proto_depIdxs = nil
}
