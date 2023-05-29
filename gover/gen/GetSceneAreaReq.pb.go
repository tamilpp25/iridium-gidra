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
// source: GetSceneAreaReq.proto

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

// CmdId: 233
// Obf: HJJMIDADBAP
type GetSceneAreaReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BelongUid uint32 `protobuf:"varint,11,opt,name=belong_uid,json=belongUid,proto3" json:"belong_uid,omitempty"`
	SceneId   uint32 `protobuf:"varint,6,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
}

func (x *GetSceneAreaReq) Reset() {
	*x = GetSceneAreaReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetSceneAreaReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSceneAreaReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSceneAreaReq) ProtoMessage() {}

func (x *GetSceneAreaReq) ProtoReflect() protoreflect.Message {
	mi := &file_GetSceneAreaReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSceneAreaReq.ProtoReflect.Descriptor instead.
func (*GetSceneAreaReq) Descriptor() ([]byte, []int) {
	return file_GetSceneAreaReq_proto_rawDescGZIP(), []int{0}
}

func (x *GetSceneAreaReq) GetBelongUid() uint32 {
	if x != nil {
		return x.BelongUid
	}
	return 0
}

func (x *GetSceneAreaReq) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

var File_GetSceneAreaReq_proto protoreflect.FileDescriptor

var file_GetSceneAreaReq_proto_rawDesc = []byte{
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65,
	0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x63,
	0x65, 0x6e, 0x65, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65,
	0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x62, 0x65, 0x6c, 0x6f, 0x6e, 0x67, 0x55, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x65,
	0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x65,
	0x6e, 0x65, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetSceneAreaReq_proto_rawDescOnce sync.Once
	file_GetSceneAreaReq_proto_rawDescData = file_GetSceneAreaReq_proto_rawDesc
)

func file_GetSceneAreaReq_proto_rawDescGZIP() []byte {
	file_GetSceneAreaReq_proto_rawDescOnce.Do(func() {
		file_GetSceneAreaReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetSceneAreaReq_proto_rawDescData)
	})
	return file_GetSceneAreaReq_proto_rawDescData
}

var file_GetSceneAreaReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetSceneAreaReq_proto_goTypes = []interface{}{
	(*GetSceneAreaReq)(nil), // 0: GetSceneAreaReq
}
var file_GetSceneAreaReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetSceneAreaReq_proto_init() }
func file_GetSceneAreaReq_proto_init() {
	if File_GetSceneAreaReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetSceneAreaReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSceneAreaReq); i {
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
			RawDescriptor: file_GetSceneAreaReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetSceneAreaReq_proto_goTypes,
		DependencyIndexes: file_GetSceneAreaReq_proto_depIdxs,
		MessageInfos:      file_GetSceneAreaReq_proto_msgTypes,
	}.Build()
	File_GetSceneAreaReq_proto = out.File
	file_GetSceneAreaReq_proto_rawDesc = nil
	file_GetSceneAreaReq_proto_goTypes = nil
	file_GetSceneAreaReq_proto_depIdxs = nil
}
