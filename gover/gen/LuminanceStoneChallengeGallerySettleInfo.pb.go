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
// source: LuminanceStoneChallengeGallerySettleInfo.proto

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

// Obf: IOFOPNAPOLG
type LuminanceStoneChallengeGallerySettleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CleanMudCount           uint32      `protobuf:"varint,6,opt,name=clean_mud_count,json=cleanMudCount,proto3" json:"clean_mud_count,omitempty"`
	GalleryId               uint32      `protobuf:"varint,12,opt,name=gallery_id,json=galleryId,proto3" json:"gallery_id,omitempty"`
	Reason                  HDDANIDHCMI `protobuf:"varint,3,opt,name=reason,proto3,enum=HDDANIDHCMI" json:"reason,omitempty"`
	FinalScore              uint32      `protobuf:"varint,14,opt,name=final_score,json=finalScore,proto3" json:"final_score,omitempty"`
	KillMonsterCount        uint32      `protobuf:"varint,9,opt,name=kill_monster_count,json=killMonsterCount,proto3" json:"kill_monster_count,omitempty"`
	KillSpecialMonsterCount uint32      `protobuf:"varint,5,opt,name=kill_special_monster_count,json=killSpecialMonsterCount,proto3" json:"kill_special_monster_count,omitempty"`
}

func (x *LuminanceStoneChallengeGallerySettleInfo) Reset() {
	*x = LuminanceStoneChallengeGallerySettleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LuminanceStoneChallengeGallerySettleInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LuminanceStoneChallengeGallerySettleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LuminanceStoneChallengeGallerySettleInfo) ProtoMessage() {}

func (x *LuminanceStoneChallengeGallerySettleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_LuminanceStoneChallengeGallerySettleInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LuminanceStoneChallengeGallerySettleInfo.ProtoReflect.Descriptor instead.
func (*LuminanceStoneChallengeGallerySettleInfo) Descriptor() ([]byte, []int) {
	return file_LuminanceStoneChallengeGallerySettleInfo_proto_rawDescGZIP(), []int{0}
}

func (x *LuminanceStoneChallengeGallerySettleInfo) GetCleanMudCount() uint32 {
	if x != nil {
		return x.CleanMudCount
	}
	return 0
}

func (x *LuminanceStoneChallengeGallerySettleInfo) GetGalleryId() uint32 {
	if x != nil {
		return x.GalleryId
	}
	return 0
}

func (x *LuminanceStoneChallengeGallerySettleInfo) GetReason() HDDANIDHCMI {
	if x != nil {
		return x.Reason
	}
	return HDDANIDHCMI_HDDANIDHCMI_GalleryStopNone
}

func (x *LuminanceStoneChallengeGallerySettleInfo) GetFinalScore() uint32 {
	if x != nil {
		return x.FinalScore
	}
	return 0
}

func (x *LuminanceStoneChallengeGallerySettleInfo) GetKillMonsterCount() uint32 {
	if x != nil {
		return x.KillMonsterCount
	}
	return 0
}

func (x *LuminanceStoneChallengeGallerySettleInfo) GetKillSpecialMonsterCount() uint32 {
	if x != nil {
		return x.KillSpecialMonsterCount
	}
	return 0
}

var File_LuminanceStoneChallengeGallerySettleInfo_proto protoreflect.FileDescriptor

var file_LuminanceStoneChallengeGallerySettleInfo_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x4c, 0x75, 0x6d, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x6f, 0x6e, 0x65,
	0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79,
	0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x48, 0x44, 0x44, 0x41, 0x4e, 0x49, 0x44, 0x48, 0x43, 0x4d, 0x49, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x02, 0x0a, 0x28, 0x4c, 0x75, 0x6d, 0x69, 0x6e, 0x61, 0x6e, 0x63,
	0x65, 0x53, 0x74, 0x6f, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x47,
	0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x26, 0x0a, 0x0f, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x5f, 0x6d, 0x75, 0x64, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x63, 0x6c, 0x65, 0x61, 0x6e,
	0x4d, 0x75, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x61, 0x6c, 0x6c,
	0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x67, 0x61,
	0x6c, 0x6c, 0x65, 0x72, 0x79, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x48, 0x44, 0x44, 0x41, 0x4e, 0x49,
	0x44, 0x48, 0x43, 0x4d, 0x49, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1f, 0x0a,
	0x0b, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x2c,
	0x0a, 0x12, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x6b, 0x69, 0x6c, 0x6c,
	0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x1a,
	0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x17, 0x6b, 0x69, 0x6c, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x4d, 0x6f, 0x6e,
	0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LuminanceStoneChallengeGallerySettleInfo_proto_rawDescOnce sync.Once
	file_LuminanceStoneChallengeGallerySettleInfo_proto_rawDescData = file_LuminanceStoneChallengeGallerySettleInfo_proto_rawDesc
)

func file_LuminanceStoneChallengeGallerySettleInfo_proto_rawDescGZIP() []byte {
	file_LuminanceStoneChallengeGallerySettleInfo_proto_rawDescOnce.Do(func() {
		file_LuminanceStoneChallengeGallerySettleInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_LuminanceStoneChallengeGallerySettleInfo_proto_rawDescData)
	})
	return file_LuminanceStoneChallengeGallerySettleInfo_proto_rawDescData
}

var file_LuminanceStoneChallengeGallerySettleInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_LuminanceStoneChallengeGallerySettleInfo_proto_goTypes = []interface{}{
	(*LuminanceStoneChallengeGallerySettleInfo)(nil), // 0: LuminanceStoneChallengeGallerySettleInfo
	(HDDANIDHCMI)(0), // 1: HDDANIDHCMI
}
var file_LuminanceStoneChallengeGallerySettleInfo_proto_depIdxs = []int32{
	1, // 0: LuminanceStoneChallengeGallerySettleInfo.reason:type_name -> HDDANIDHCMI
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_LuminanceStoneChallengeGallerySettleInfo_proto_init() }
func file_LuminanceStoneChallengeGallerySettleInfo_proto_init() {
	if File_LuminanceStoneChallengeGallerySettleInfo_proto != nil {
		return
	}
	file_HDDANIDHCMI_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_LuminanceStoneChallengeGallerySettleInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LuminanceStoneChallengeGallerySettleInfo); i {
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
			RawDescriptor: file_LuminanceStoneChallengeGallerySettleInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LuminanceStoneChallengeGallerySettleInfo_proto_goTypes,
		DependencyIndexes: file_LuminanceStoneChallengeGallerySettleInfo_proto_depIdxs,
		MessageInfos:      file_LuminanceStoneChallengeGallerySettleInfo_proto_msgTypes,
	}.Build()
	File_LuminanceStoneChallengeGallerySettleInfo_proto = out.File
	file_LuminanceStoneChallengeGallerySettleInfo_proto_rawDesc = nil
	file_LuminanceStoneChallengeGallerySettleInfo_proto_goTypes = nil
	file_LuminanceStoneChallengeGallerySettleInfo_proto_depIdxs = nil
}
