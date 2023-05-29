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
// source: HomeGroupPlayerInfo.proto

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

// Obf: IJHCHNPKNGD
type HomeGroupPlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerLevel    uint32          `protobuf:"varint,7,opt,name=player_level,json=playerLevel,proto3" json:"player_level,omitempty"`
	Uid            uint32          `protobuf:"varint,11,opt,name=uid,proto3" json:"uid,omitempty"`
	PsnId          string          `protobuf:"bytes,5,opt,name=psn_id,json=psnId,proto3" json:"psn_id,omitempty"`
	Nickname       string          `protobuf:"bytes,9,opt,name=nickname,proto3" json:"nickname,omitempty"`
	OnlineId       string          `protobuf:"bytes,8,opt,name=online_id,json=onlineId,proto3" json:"online_id,omitempty"`
	ProfilePicture *ProfilePicture `protobuf:"bytes,3,opt,name=profile_picture,json=profilePicture,proto3" json:"profile_picture,omitempty"`
}

func (x *HomeGroupPlayerInfo) Reset() {
	*x = HomeGroupPlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HomeGroupPlayerInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeGroupPlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeGroupPlayerInfo) ProtoMessage() {}

func (x *HomeGroupPlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_HomeGroupPlayerInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeGroupPlayerInfo.ProtoReflect.Descriptor instead.
func (*HomeGroupPlayerInfo) Descriptor() ([]byte, []int) {
	return file_HomeGroupPlayerInfo_proto_rawDescGZIP(), []int{0}
}

func (x *HomeGroupPlayerInfo) GetPlayerLevel() uint32 {
	if x != nil {
		return x.PlayerLevel
	}
	return 0
}

func (x *HomeGroupPlayerInfo) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *HomeGroupPlayerInfo) GetPsnId() string {
	if x != nil {
		return x.PsnId
	}
	return ""
}

func (x *HomeGroupPlayerInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *HomeGroupPlayerInfo) GetOnlineId() string {
	if x != nil {
		return x.OnlineId
	}
	return ""
}

func (x *HomeGroupPlayerInfo) GetProfilePicture() *ProfilePicture {
	if x != nil {
		return x.ProfilePicture
	}
	return nil
}

var File_HomeGroupPlayerInfo_proto protoreflect.FileDescriptor

var file_HomeGroupPlayerInfo_proto_rawDesc = []byte{
	0x0a, 0x19, 0x48, 0x6f, 0x6d, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd4, 0x01, 0x0a, 0x13, 0x48, 0x6f, 0x6d, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x15,
	0x0a, 0x06, 0x70, 0x73, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x73, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x38,
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HomeGroupPlayerInfo_proto_rawDescOnce sync.Once
	file_HomeGroupPlayerInfo_proto_rawDescData = file_HomeGroupPlayerInfo_proto_rawDesc
)

func file_HomeGroupPlayerInfo_proto_rawDescGZIP() []byte {
	file_HomeGroupPlayerInfo_proto_rawDescOnce.Do(func() {
		file_HomeGroupPlayerInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_HomeGroupPlayerInfo_proto_rawDescData)
	})
	return file_HomeGroupPlayerInfo_proto_rawDescData
}

var file_HomeGroupPlayerInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HomeGroupPlayerInfo_proto_goTypes = []interface{}{
	(*HomeGroupPlayerInfo)(nil), // 0: HomeGroupPlayerInfo
	(*ProfilePicture)(nil),      // 1: ProfilePicture
}
var file_HomeGroupPlayerInfo_proto_depIdxs = []int32{
	1, // 0: HomeGroupPlayerInfo.profile_picture:type_name -> ProfilePicture
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_HomeGroupPlayerInfo_proto_init() }
func file_HomeGroupPlayerInfo_proto_init() {
	if File_HomeGroupPlayerInfo_proto != nil {
		return
	}
	file_ProfilePicture_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HomeGroupPlayerInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeGroupPlayerInfo); i {
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
			RawDescriptor: file_HomeGroupPlayerInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HomeGroupPlayerInfo_proto_goTypes,
		DependencyIndexes: file_HomeGroupPlayerInfo_proto_depIdxs,
		MessageInfos:      file_HomeGroupPlayerInfo_proto_msgTypes,
	}.Build()
	File_HomeGroupPlayerInfo_proto = out.File
	file_HomeGroupPlayerInfo_proto_rawDesc = nil
	file_HomeGroupPlayerInfo_proto_goTypes = nil
	file_HomeGroupPlayerInfo_proto_depIdxs = nil
}
