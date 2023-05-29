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
// source: RogueDiaryActivityDetailInfo.proto

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

// Obf: FEKLEGPDNJI
type RogueDiaryActivityDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StageList       []*RogueDiaryStageInfo `protobuf:"bytes,11,rep,name=stage_list,json=stageList,proto3" json:"stage_list,omitempty"`
	CurProgress     *RogueDiaryProgress    `protobuf:"bytes,6,opt,name=cur_progress,json=curProgress,proto3" json:"cur_progress,omitempty"`
	IsHaveProgress  bool                   `protobuf:"varint,1,opt,name=is_have_progress,json=isHaveProgress,proto3" json:"is_have_progress,omitempty"`
	IsContentClosed bool                   `protobuf:"varint,8,opt,name=is_content_closed,json=isContentClosed,proto3" json:"is_content_closed,omitempty"`
}

func (x *RogueDiaryActivityDetailInfo) Reset() {
	*x = RogueDiaryActivityDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueDiaryActivityDetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueDiaryActivityDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueDiaryActivityDetailInfo) ProtoMessage() {}

func (x *RogueDiaryActivityDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_RogueDiaryActivityDetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueDiaryActivityDetailInfo.ProtoReflect.Descriptor instead.
func (*RogueDiaryActivityDetailInfo) Descriptor() ([]byte, []int) {
	return file_RogueDiaryActivityDetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *RogueDiaryActivityDetailInfo) GetStageList() []*RogueDiaryStageInfo {
	if x != nil {
		return x.StageList
	}
	return nil
}

func (x *RogueDiaryActivityDetailInfo) GetCurProgress() *RogueDiaryProgress {
	if x != nil {
		return x.CurProgress
	}
	return nil
}

func (x *RogueDiaryActivityDetailInfo) GetIsHaveProgress() bool {
	if x != nil {
		return x.IsHaveProgress
	}
	return false
}

func (x *RogueDiaryActivityDetailInfo) GetIsContentClosed() bool {
	if x != nil {
		return x.IsContentClosed
	}
	return false
}

var File_RogueDiaryActivityDetailInfo_proto protoreflect.FileDescriptor

var file_RogueDiaryActivityDetailInfo_proto_rawDesc = []byte{
	0x0a, 0x22, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61, 0x72, 0x79, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61, 0x72, 0x79,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x18, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x01, 0x0a, 0x1c, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x44, 0x69, 0x61, 0x72, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x33, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x73, 0x74, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x36, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61,
	0x72, 0x79, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x73, 0x5f, 0x68, 0x61,
	0x76, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0e, 0x69, 0x73, 0x48, 0x61, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f,
	0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x42, 0x06, 0x5a,
	0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueDiaryActivityDetailInfo_proto_rawDescOnce sync.Once
	file_RogueDiaryActivityDetailInfo_proto_rawDescData = file_RogueDiaryActivityDetailInfo_proto_rawDesc
)

func file_RogueDiaryActivityDetailInfo_proto_rawDescGZIP() []byte {
	file_RogueDiaryActivityDetailInfo_proto_rawDescOnce.Do(func() {
		file_RogueDiaryActivityDetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueDiaryActivityDetailInfo_proto_rawDescData)
	})
	return file_RogueDiaryActivityDetailInfo_proto_rawDescData
}

var file_RogueDiaryActivityDetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueDiaryActivityDetailInfo_proto_goTypes = []interface{}{
	(*RogueDiaryActivityDetailInfo)(nil), // 0: RogueDiaryActivityDetailInfo
	(*RogueDiaryStageInfo)(nil),          // 1: RogueDiaryStageInfo
	(*RogueDiaryProgress)(nil),           // 2: RogueDiaryProgress
}
var file_RogueDiaryActivityDetailInfo_proto_depIdxs = []int32{
	1, // 0: RogueDiaryActivityDetailInfo.stage_list:type_name -> RogueDiaryStageInfo
	2, // 1: RogueDiaryActivityDetailInfo.cur_progress:type_name -> RogueDiaryProgress
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_RogueDiaryActivityDetailInfo_proto_init() }
func file_RogueDiaryActivityDetailInfo_proto_init() {
	if File_RogueDiaryActivityDetailInfo_proto != nil {
		return
	}
	file_RogueDiaryStageInfo_proto_init()
	file_RogueDiaryProgress_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueDiaryActivityDetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueDiaryActivityDetailInfo); i {
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
			RawDescriptor: file_RogueDiaryActivityDetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueDiaryActivityDetailInfo_proto_goTypes,
		DependencyIndexes: file_RogueDiaryActivityDetailInfo_proto_depIdxs,
		MessageInfos:      file_RogueDiaryActivityDetailInfo_proto_msgTypes,
	}.Build()
	File_RogueDiaryActivityDetailInfo_proto = out.File
	file_RogueDiaryActivityDetailInfo_proto_rawDesc = nil
	file_RogueDiaryActivityDetailInfo_proto_goTypes = nil
	file_RogueDiaryActivityDetailInfo_proto_depIdxs = nil
}
