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
// source: SceneGalleryMuqadasPotionInfo.proto

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

// Obf: ILLHNCAJNLB
type SceneGalleryMuqadasPotionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GKFJJMPJCMK          uint32 `protobuf:"varint,11,opt,name=GKFJJMPJCMK,proto3" json:"GKFJJMPJCMK,omitempty"`
	Score                uint32 `protobuf:"varint,12,opt,name=score,proto3" json:"score,omitempty"`
	CaptureWeaknessCount uint32 `protobuf:"varint,14,opt,name=capture_weakness_count,json=captureWeaknessCount,proto3" json:"capture_weakness_count,omitempty"`
	FNAEKILFJLD          uint32 `protobuf:"varint,7,opt,name=FNAEKILFJLD,proto3" json:"FNAEKILFJLD,omitempty"`
}

func (x *SceneGalleryMuqadasPotionInfo) Reset() {
	*x = SceneGalleryMuqadasPotionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneGalleryMuqadasPotionInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneGalleryMuqadasPotionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneGalleryMuqadasPotionInfo) ProtoMessage() {}

func (x *SceneGalleryMuqadasPotionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_SceneGalleryMuqadasPotionInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneGalleryMuqadasPotionInfo.ProtoReflect.Descriptor instead.
func (*SceneGalleryMuqadasPotionInfo) Descriptor() ([]byte, []int) {
	return file_SceneGalleryMuqadasPotionInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SceneGalleryMuqadasPotionInfo) GetGKFJJMPJCMK() uint32 {
	if x != nil {
		return x.GKFJJMPJCMK
	}
	return 0
}

func (x *SceneGalleryMuqadasPotionInfo) GetScore() uint32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *SceneGalleryMuqadasPotionInfo) GetCaptureWeaknessCount() uint32 {
	if x != nil {
		return x.CaptureWeaknessCount
	}
	return 0
}

func (x *SceneGalleryMuqadasPotionInfo) GetFNAEKILFJLD() uint32 {
	if x != nil {
		return x.FNAEKILFJLD
	}
	return 0
}

var File_SceneGalleryMuqadasPotionInfo_proto protoreflect.FileDescriptor

var file_SceneGalleryMuqadasPotionInfo_proto_rawDesc = []byte{
	0x0a, 0x23, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x4d, 0x75,
	0x71, 0x61, 0x64, 0x61, 0x73, 0x50, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x01, 0x0a, 0x1d, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x47,
	0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x4d, 0x75, 0x71, 0x61, 0x64, 0x61, 0x73, 0x50, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x4b, 0x46, 0x4a, 0x4a,
	0x4d, 0x50, 0x4a, 0x43, 0x4d, 0x4b, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x47, 0x4b,
	0x46, 0x4a, 0x4a, 0x4d, 0x50, 0x4a, 0x43, 0x4d, 0x4b, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12,
	0x34, 0x0a, 0x16, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x77, 0x65, 0x61, 0x6b, 0x6e,
	0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x14, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x57, 0x65, 0x61, 0x6b, 0x6e, 0x65, 0x73, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x4e, 0x41, 0x45, 0x4b, 0x49, 0x4c,
	0x46, 0x4a, 0x4c, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x46, 0x4e, 0x41, 0x45,
	0x4b, 0x49, 0x4c, 0x46, 0x4a, 0x4c, 0x44, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneGalleryMuqadasPotionInfo_proto_rawDescOnce sync.Once
	file_SceneGalleryMuqadasPotionInfo_proto_rawDescData = file_SceneGalleryMuqadasPotionInfo_proto_rawDesc
)

func file_SceneGalleryMuqadasPotionInfo_proto_rawDescGZIP() []byte {
	file_SceneGalleryMuqadasPotionInfo_proto_rawDescOnce.Do(func() {
		file_SceneGalleryMuqadasPotionInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneGalleryMuqadasPotionInfo_proto_rawDescData)
	})
	return file_SceneGalleryMuqadasPotionInfo_proto_rawDescData
}

var file_SceneGalleryMuqadasPotionInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneGalleryMuqadasPotionInfo_proto_goTypes = []interface{}{
	(*SceneGalleryMuqadasPotionInfo)(nil), // 0: SceneGalleryMuqadasPotionInfo
}
var file_SceneGalleryMuqadasPotionInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SceneGalleryMuqadasPotionInfo_proto_init() }
func file_SceneGalleryMuqadasPotionInfo_proto_init() {
	if File_SceneGalleryMuqadasPotionInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SceneGalleryMuqadasPotionInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneGalleryMuqadasPotionInfo); i {
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
			RawDescriptor: file_SceneGalleryMuqadasPotionInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneGalleryMuqadasPotionInfo_proto_goTypes,
		DependencyIndexes: file_SceneGalleryMuqadasPotionInfo_proto_depIdxs,
		MessageInfos:      file_SceneGalleryMuqadasPotionInfo_proto_msgTypes,
	}.Build()
	File_SceneGalleryMuqadasPotionInfo_proto = out.File
	file_SceneGalleryMuqadasPotionInfo_proto_rawDesc = nil
	file_SceneGalleryMuqadasPotionInfo_proto_goTypes = nil
	file_SceneGalleryMuqadasPotionInfo_proto_depIdxs = nil
}
