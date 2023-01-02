// Sorapointa - A server software re-implementation for a certain anime game, and avoid sorapointa.
// Copyright (C) 2022  Sorapointa Team
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
// 	protoc        v3.11.3
// source: PlantFlowerGetFriendFlowerWishListRsp.proto

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

type PlantFlowerGetFriendFlowerWishListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScheduleId           uint32                             `protobuf:"varint,14,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
	Retcode              int32                              `protobuf:"varint,7,opt,name=retcode,proto3" json:"retcode,omitempty"`
	FriendFlowerWishList []*PlantFlowerFriendFlowerWishData `protobuf:"bytes,13,rep,name=friend_flower_wish_list,json=friendFlowerWishList,proto3" json:"friend_flower_wish_list,omitempty"`
}

func (x *PlantFlowerGetFriendFlowerWishListRsp) Reset() {
	*x = PlantFlowerGetFriendFlowerWishListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlantFlowerGetFriendFlowerWishListRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlantFlowerGetFriendFlowerWishListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlantFlowerGetFriendFlowerWishListRsp) ProtoMessage() {}

func (x *PlantFlowerGetFriendFlowerWishListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_PlantFlowerGetFriendFlowerWishListRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlantFlowerGetFriendFlowerWishListRsp.ProtoReflect.Descriptor instead.
func (*PlantFlowerGetFriendFlowerWishListRsp) Descriptor() ([]byte, []int) {
	return file_PlantFlowerGetFriendFlowerWishListRsp_proto_rawDescGZIP(), []int{0}
}

func (x *PlantFlowerGetFriendFlowerWishListRsp) GetScheduleId() uint32 {
	if x != nil {
		return x.ScheduleId
	}
	return 0
}

func (x *PlantFlowerGetFriendFlowerWishListRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *PlantFlowerGetFriendFlowerWishListRsp) GetFriendFlowerWishList() []*PlantFlowerFriendFlowerWishData {
	if x != nil {
		return x.FriendFlowerWishList
	}
	return nil
}

var File_PlantFlowerGetFriendFlowerWishListRsp_proto protoreflect.FileDescriptor

var file_PlantFlowerGetFriendFlowerWishListRsp_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x47, 0x65, 0x74,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x46, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x57, 0x69, 0x73, 0x68,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x46, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x57, 0x69, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x01, 0x0a, 0x25, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x46, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x46, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x57, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x57, 0x0a, 0x17, 0x66, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x77, 0x69, 0x73, 0x68, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x46, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x57, 0x69, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x14, 0x66, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x46, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x57, 0x69, 0x73, 0x68, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_PlantFlowerGetFriendFlowerWishListRsp_proto_rawDescOnce sync.Once
	file_PlantFlowerGetFriendFlowerWishListRsp_proto_rawDescData = file_PlantFlowerGetFriendFlowerWishListRsp_proto_rawDesc
)

func file_PlantFlowerGetFriendFlowerWishListRsp_proto_rawDescGZIP() []byte {
	file_PlantFlowerGetFriendFlowerWishListRsp_proto_rawDescOnce.Do(func() {
		file_PlantFlowerGetFriendFlowerWishListRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlantFlowerGetFriendFlowerWishListRsp_proto_rawDescData)
	})
	return file_PlantFlowerGetFriendFlowerWishListRsp_proto_rawDescData
}

var file_PlantFlowerGetFriendFlowerWishListRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlantFlowerGetFriendFlowerWishListRsp_proto_goTypes = []interface{}{
	(*PlantFlowerGetFriendFlowerWishListRsp)(nil), // 0: PlantFlowerGetFriendFlowerWishListRsp
	(*PlantFlowerFriendFlowerWishData)(nil),       // 1: PlantFlowerFriendFlowerWishData
}
var file_PlantFlowerGetFriendFlowerWishListRsp_proto_depIdxs = []int32{
	1, // 0: PlantFlowerGetFriendFlowerWishListRsp.friend_flower_wish_list:type_name -> PlantFlowerFriendFlowerWishData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PlantFlowerGetFriendFlowerWishListRsp_proto_init() }
func file_PlantFlowerGetFriendFlowerWishListRsp_proto_init() {
	if File_PlantFlowerGetFriendFlowerWishListRsp_proto != nil {
		return
	}
	file_PlantFlowerFriendFlowerWishData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlantFlowerGetFriendFlowerWishListRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlantFlowerGetFriendFlowerWishListRsp); i {
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
			RawDescriptor: file_PlantFlowerGetFriendFlowerWishListRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlantFlowerGetFriendFlowerWishListRsp_proto_goTypes,
		DependencyIndexes: file_PlantFlowerGetFriendFlowerWishListRsp_proto_depIdxs,
		MessageInfos:      file_PlantFlowerGetFriendFlowerWishListRsp_proto_msgTypes,
	}.Build()
	File_PlantFlowerGetFriendFlowerWishListRsp_proto = out.File
	file_PlantFlowerGetFriendFlowerWishListRsp_proto_rawDesc = nil
	file_PlantFlowerGetFriendFlowerWishListRsp_proto_goTypes = nil
	file_PlantFlowerGetFriendFlowerWishListRsp_proto_depIdxs = nil
}
