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
// source: LanternRiteEndFireworksReformRsp.proto

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

// CmdId: 8301
// Obf: NCJKMIJMNOG
type LanternRiteEndFireworksReformRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GHNMLHGPOOE bool   `protobuf:"varint,3,opt,name=GHNMLHGPOOE,proto3" json:"GHNMLHGPOOE,omitempty"`
	LBCHAOJNILB bool   `protobuf:"varint,13,opt,name=LBCHAOJNILB,proto3" json:"LBCHAOJNILB,omitempty"`
	APCMMCILEOJ bool   `protobuf:"varint,15,opt,name=APCMMCILEOJ,proto3" json:"APCMMCILEOJ,omitempty"`
	IsNewRecord bool   `protobuf:"varint,8,opt,name=is_new_record,json=isNewRecord,proto3" json:"is_new_record,omitempty"`
	ChallengeId uint32 `protobuf:"varint,12,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
	StageId     uint32 `protobuf:"varint,9,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
	FinalScore  uint32 `protobuf:"varint,14,opt,name=final_score,json=finalScore,proto3" json:"final_score,omitempty"`
	IsFullScore bool   `protobuf:"varint,10,opt,name=is_full_score,json=isFullScore,proto3" json:"is_full_score,omitempty"`
	Retcode     int32  `protobuf:"varint,5,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *LanternRiteEndFireworksReformRsp) Reset() {
	*x = LanternRiteEndFireworksReformRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LanternRiteEndFireworksReformRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LanternRiteEndFireworksReformRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LanternRiteEndFireworksReformRsp) ProtoMessage() {}

func (x *LanternRiteEndFireworksReformRsp) ProtoReflect() protoreflect.Message {
	mi := &file_LanternRiteEndFireworksReformRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LanternRiteEndFireworksReformRsp.ProtoReflect.Descriptor instead.
func (*LanternRiteEndFireworksReformRsp) Descriptor() ([]byte, []int) {
	return file_LanternRiteEndFireworksReformRsp_proto_rawDescGZIP(), []int{0}
}

func (x *LanternRiteEndFireworksReformRsp) GetGHNMLHGPOOE() bool {
	if x != nil {
		return x.GHNMLHGPOOE
	}
	return false
}

func (x *LanternRiteEndFireworksReformRsp) GetLBCHAOJNILB() bool {
	if x != nil {
		return x.LBCHAOJNILB
	}
	return false
}

func (x *LanternRiteEndFireworksReformRsp) GetAPCMMCILEOJ() bool {
	if x != nil {
		return x.APCMMCILEOJ
	}
	return false
}

func (x *LanternRiteEndFireworksReformRsp) GetIsNewRecord() bool {
	if x != nil {
		return x.IsNewRecord
	}
	return false
}

func (x *LanternRiteEndFireworksReformRsp) GetChallengeId() uint32 {
	if x != nil {
		return x.ChallengeId
	}
	return 0
}

func (x *LanternRiteEndFireworksReformRsp) GetStageId() uint32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *LanternRiteEndFireworksReformRsp) GetFinalScore() uint32 {
	if x != nil {
		return x.FinalScore
	}
	return 0
}

func (x *LanternRiteEndFireworksReformRsp) GetIsFullScore() bool {
	if x != nil {
		return x.IsFullScore
	}
	return false
}

func (x *LanternRiteEndFireworksReformRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_LanternRiteEndFireworksReformRsp_proto protoreflect.FileDescriptor

var file_LanternRiteEndFireworksReformRsp_proto_rawDesc = []byte{
	0x0a, 0x26, 0x4c, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x52, 0x69, 0x74, 0x65, 0x45, 0x6e, 0x64,
	0x46, 0x69, 0x72, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x52,
	0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x02, 0x0a, 0x20, 0x4c, 0x61, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x52, 0x69, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x46, 0x69, 0x72, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x73, 0x70, 0x12, 0x20, 0x0a,
	0x0b, 0x47, 0x48, 0x4e, 0x4d, 0x4c, 0x48, 0x47, 0x50, 0x4f, 0x4f, 0x45, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x47, 0x48, 0x4e, 0x4d, 0x4c, 0x48, 0x47, 0x50, 0x4f, 0x4f, 0x45, 0x12,
	0x20, 0x0a, 0x0b, 0x4c, 0x42, 0x43, 0x48, 0x41, 0x4f, 0x4a, 0x4e, 0x49, 0x4c, 0x42, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4c, 0x42, 0x43, 0x48, 0x41, 0x4f, 0x4a, 0x4e, 0x49, 0x4c,
	0x42, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x50, 0x43, 0x4d, 0x4d, 0x43, 0x49, 0x4c, 0x45, 0x4f, 0x4a,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x41, 0x50, 0x43, 0x4d, 0x4d, 0x43, 0x49, 0x4c,
	0x45, 0x4f, 0x4a, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x4e, 0x65,
	0x77, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6c, 0x6c,
	0x65, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x61,
	0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x66, 0x75, 0x6c,
	0x6c, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69,
	0x73, 0x46, 0x75, 0x6c, 0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LanternRiteEndFireworksReformRsp_proto_rawDescOnce sync.Once
	file_LanternRiteEndFireworksReformRsp_proto_rawDescData = file_LanternRiteEndFireworksReformRsp_proto_rawDesc
)

func file_LanternRiteEndFireworksReformRsp_proto_rawDescGZIP() []byte {
	file_LanternRiteEndFireworksReformRsp_proto_rawDescOnce.Do(func() {
		file_LanternRiteEndFireworksReformRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_LanternRiteEndFireworksReformRsp_proto_rawDescData)
	})
	return file_LanternRiteEndFireworksReformRsp_proto_rawDescData
}

var file_LanternRiteEndFireworksReformRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_LanternRiteEndFireworksReformRsp_proto_goTypes = []interface{}{
	(*LanternRiteEndFireworksReformRsp)(nil), // 0: LanternRiteEndFireworksReformRsp
}
var file_LanternRiteEndFireworksReformRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_LanternRiteEndFireworksReformRsp_proto_init() }
func file_LanternRiteEndFireworksReformRsp_proto_init() {
	if File_LanternRiteEndFireworksReformRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_LanternRiteEndFireworksReformRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LanternRiteEndFireworksReformRsp); i {
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
			RawDescriptor: file_LanternRiteEndFireworksReformRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LanternRiteEndFireworksReformRsp_proto_goTypes,
		DependencyIndexes: file_LanternRiteEndFireworksReformRsp_proto_depIdxs,
		MessageInfos:      file_LanternRiteEndFireworksReformRsp_proto_msgTypes,
	}.Build()
	File_LanternRiteEndFireworksReformRsp_proto = out.File
	file_LanternRiteEndFireworksReformRsp_proto_rawDesc = nil
	file_LanternRiteEndFireworksReformRsp_proto_goTypes = nil
	file_LanternRiteEndFireworksReformRsp_proto_depIdxs = nil
}
