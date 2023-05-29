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
// source: GetGameplayRecommendationReq.proto

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

// CmdId: 195
// Obf: FKEKHPFJAOP
type GetGameplayRecommendationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarId uint32 `protobuf:"varint,8,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
	// Types that are assignable to Detail:
	//
	//	*GetGameplayRecommendationReq_SkillRequest
	//	*GetGameplayRecommendationReq_ReliquaryRequest
	//	*GetGameplayRecommendationReq_ElementReliquaryRequest
	Detail isGetGameplayRecommendationReq_Detail `protobuf_oneof:"detail"`
}

func (x *GetGameplayRecommendationReq) Reset() {
	*x = GetGameplayRecommendationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetGameplayRecommendationReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGameplayRecommendationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGameplayRecommendationReq) ProtoMessage() {}

func (x *GetGameplayRecommendationReq) ProtoReflect() protoreflect.Message {
	mi := &file_GetGameplayRecommendationReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGameplayRecommendationReq.ProtoReflect.Descriptor instead.
func (*GetGameplayRecommendationReq) Descriptor() ([]byte, []int) {
	return file_GetGameplayRecommendationReq_proto_rawDescGZIP(), []int{0}
}

func (x *GetGameplayRecommendationReq) GetAvatarId() uint32 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

func (m *GetGameplayRecommendationReq) GetDetail() isGetGameplayRecommendationReq_Detail {
	if m != nil {
		return m.Detail
	}
	return nil
}

func (x *GetGameplayRecommendationReq) GetSkillRequest() *GameplayRecommendationSkillRequest {
	if x, ok := x.GetDetail().(*GetGameplayRecommendationReq_SkillRequest); ok {
		return x.SkillRequest
	}
	return nil
}

func (x *GetGameplayRecommendationReq) GetReliquaryRequest() *GameplayRecommendationReliquaryRequest {
	if x, ok := x.GetDetail().(*GetGameplayRecommendationReq_ReliquaryRequest); ok {
		return x.ReliquaryRequest
	}
	return nil
}

func (x *GetGameplayRecommendationReq) GetElementReliquaryRequest() *GameplayRecommendationElementReliquaryRequest {
	if x, ok := x.GetDetail().(*GetGameplayRecommendationReq_ElementReliquaryRequest); ok {
		return x.ElementReliquaryRequest
	}
	return nil
}

type isGetGameplayRecommendationReq_Detail interface {
	isGetGameplayRecommendationReq_Detail()
}

type GetGameplayRecommendationReq_SkillRequest struct {
	SkillRequest *GameplayRecommendationSkillRequest `protobuf:"bytes,195,opt,name=skill_request,json=skillRequest,proto3,oneof"`
}

type GetGameplayRecommendationReq_ReliquaryRequest struct {
	ReliquaryRequest *GameplayRecommendationReliquaryRequest `protobuf:"bytes,293,opt,name=reliquary_request,json=reliquaryRequest,proto3,oneof"`
}

type GetGameplayRecommendationReq_ElementReliquaryRequest struct {
	ElementReliquaryRequest *GameplayRecommendationElementReliquaryRequest `protobuf:"bytes,1333,opt,name=element_reliquary_request,json=elementReliquaryRequest,proto3,oneof"`
}

func (*GetGameplayRecommendationReq_SkillRequest) isGetGameplayRecommendationReq_Detail() {}

func (*GetGameplayRecommendationReq_ReliquaryRequest) isGetGameplayRecommendationReq_Detail() {}

func (*GetGameplayRecommendationReq_ElementReliquaryRequest) isGetGameplayRecommendationReq_Detail() {
}

var File_GetGameplayRecommendationReq_proto protoreflect.FileDescriptor

var file_GetGameplayRecommendationReq_proto_rawDesc = []byte{
	0x0a, 0x22, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x47, 0x61, 0x6d, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x52, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6b, 0x69, 0x6c,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c,
	0x47, 0x61, 0x6d, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6c, 0x69, 0x71, 0x75, 0x61, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x47, 0x61,
	0x6d, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x6c, 0x69, 0x71,
	0x75, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xda, 0x02, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x70, 0x6c, 0x61,
	0x79, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x12,
	0x4b, 0x0a, 0x0d, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0xc3, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x70, 0x6c,
	0x61, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0c,
	0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x57, 0x0a, 0x11,
	0x72, 0x65, 0x6c, 0x69, 0x71, 0x75, 0x61, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0xa5, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x70,
	0x6c, 0x61, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x6c, 0x69, 0x71, 0x75, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x48, 0x00, 0x52, 0x10, 0x72, 0x65, 0x6c, 0x69, 0x71, 0x75, 0x61, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x6d, 0x0a, 0x19, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x72, 0x65, 0x6c, 0x69, 0x71, 0x75, 0x61, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0xb5, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x47, 0x61, 0x6d, 0x65,
	0x70, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x6c, 0x69, 0x71, 0x75, 0x61,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x17, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x6c, 0x69, 0x71, 0x75, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x06,
	0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetGameplayRecommendationReq_proto_rawDescOnce sync.Once
	file_GetGameplayRecommendationReq_proto_rawDescData = file_GetGameplayRecommendationReq_proto_rawDesc
)

func file_GetGameplayRecommendationReq_proto_rawDescGZIP() []byte {
	file_GetGameplayRecommendationReq_proto_rawDescOnce.Do(func() {
		file_GetGameplayRecommendationReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetGameplayRecommendationReq_proto_rawDescData)
	})
	return file_GetGameplayRecommendationReq_proto_rawDescData
}

var file_GetGameplayRecommendationReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetGameplayRecommendationReq_proto_goTypes = []interface{}{
	(*GetGameplayRecommendationReq)(nil),                  // 0: GetGameplayRecommendationReq
	(*GameplayRecommendationSkillRequest)(nil),            // 1: GameplayRecommendationSkillRequest
	(*GameplayRecommendationReliquaryRequest)(nil),        // 2: GameplayRecommendationReliquaryRequest
	(*GameplayRecommendationElementReliquaryRequest)(nil), // 3: GameplayRecommendationElementReliquaryRequest
}
var file_GetGameplayRecommendationReq_proto_depIdxs = []int32{
	1, // 0: GetGameplayRecommendationReq.skill_request:type_name -> GameplayRecommendationSkillRequest
	2, // 1: GetGameplayRecommendationReq.reliquary_request:type_name -> GameplayRecommendationReliquaryRequest
	3, // 2: GetGameplayRecommendationReq.element_reliquary_request:type_name -> GameplayRecommendationElementReliquaryRequest
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_GetGameplayRecommendationReq_proto_init() }
func file_GetGameplayRecommendationReq_proto_init() {
	if File_GetGameplayRecommendationReq_proto != nil {
		return
	}
	file_GameplayRecommendationSkillRequest_proto_init()
	file_GameplayRecommendationReliquaryRequest_proto_init()
	file_GameplayRecommendationElementReliquaryRequest_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetGameplayRecommendationReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGameplayRecommendationReq); i {
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
	file_GetGameplayRecommendationReq_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*GetGameplayRecommendationReq_SkillRequest)(nil),
		(*GetGameplayRecommendationReq_ReliquaryRequest)(nil),
		(*GetGameplayRecommendationReq_ElementReliquaryRequest)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_GetGameplayRecommendationReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetGameplayRecommendationReq_proto_goTypes,
		DependencyIndexes: file_GetGameplayRecommendationReq_proto_depIdxs,
		MessageInfos:      file_GetGameplayRecommendationReq_proto_msgTypes,
	}.Build()
	File_GetGameplayRecommendationReq_proto = out.File
	file_GetGameplayRecommendationReq_proto_rawDesc = nil
	file_GetGameplayRecommendationReq_proto_goTypes = nil
	file_GetGameplayRecommendationReq_proto_depIdxs = nil
}
