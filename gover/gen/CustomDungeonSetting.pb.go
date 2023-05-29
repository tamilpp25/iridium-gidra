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
// source: CustomDungeonSetting.proto

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

// Obf: MAHLNNJLPLC
type CustomDungeonSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsForbidSkill  bool     `protobuf:"varint,2,opt,name=is_forbid_skill,json=isForbidSkill,proto3" json:"is_forbid_skill,omitempty"`
	LifeNum        uint32   `protobuf:"varint,9,opt,name=life_num,json=lifeNum,proto3" json:"life_num,omitempty"`
	CoinLimit      uint32   `protobuf:"varint,13,opt,name=coin_limit,json=coinLimit,proto3" json:"coin_limit,omitempty"`
	TimeLimit      uint32   `protobuf:"varint,10,opt,name=time_limit,json=timeLimit,proto3" json:"time_limit,omitempty"`
	StartRoomId    uint32   `protobuf:"varint,5,opt,name=start_room_id,json=startRoomId,proto3" json:"start_room_id,omitempty"`
	IsArriveFinish bool     `protobuf:"varint,14,opt,name=is_arrive_finish,json=isArriveFinish,proto3" json:"is_arrive_finish,omitempty"`
	OpenRoomList   []uint32 `protobuf:"varint,11,rep,packed,name=open_room_list,json=openRoomList,proto3" json:"open_room_list,omitempty"`
}

func (x *CustomDungeonSetting) Reset() {
	*x = CustomDungeonSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CustomDungeonSetting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomDungeonSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomDungeonSetting) ProtoMessage() {}

func (x *CustomDungeonSetting) ProtoReflect() protoreflect.Message {
	mi := &file_CustomDungeonSetting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomDungeonSetting.ProtoReflect.Descriptor instead.
func (*CustomDungeonSetting) Descriptor() ([]byte, []int) {
	return file_CustomDungeonSetting_proto_rawDescGZIP(), []int{0}
}

func (x *CustomDungeonSetting) GetIsForbidSkill() bool {
	if x != nil {
		return x.IsForbidSkill
	}
	return false
}

func (x *CustomDungeonSetting) GetLifeNum() uint32 {
	if x != nil {
		return x.LifeNum
	}
	return 0
}

func (x *CustomDungeonSetting) GetCoinLimit() uint32 {
	if x != nil {
		return x.CoinLimit
	}
	return 0
}

func (x *CustomDungeonSetting) GetTimeLimit() uint32 {
	if x != nil {
		return x.TimeLimit
	}
	return 0
}

func (x *CustomDungeonSetting) GetStartRoomId() uint32 {
	if x != nil {
		return x.StartRoomId
	}
	return 0
}

func (x *CustomDungeonSetting) GetIsArriveFinish() bool {
	if x != nil {
		return x.IsArriveFinish
	}
	return false
}

func (x *CustomDungeonSetting) GetOpenRoomList() []uint32 {
	if x != nil {
		return x.OpenRoomList
	}
	return nil
}

var File_CustomDungeonSetting_proto protoreflect.FileDescriptor

var file_CustomDungeonSetting_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x02, 0x0a,
	0x14, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x72, 0x62,
	0x69, 0x64, 0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x69, 0x73, 0x46, 0x6f, 0x72, 0x62, 0x69, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x19, 0x0a,
	0x08, 0x6c, 0x69, 0x66, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x6c, 0x69, 0x66, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x69, 0x6e,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f,
	0x69, 0x6e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x73,
	0x5f, 0x61, 0x72, 0x72, 0x69, 0x76, 0x65, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x41, 0x72, 0x72, 0x69, 0x76, 0x65, 0x46, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x12, 0x24, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x72, 0x6f, 0x6f,
	0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x6f, 0x70,
	0x65, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67,
	0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CustomDungeonSetting_proto_rawDescOnce sync.Once
	file_CustomDungeonSetting_proto_rawDescData = file_CustomDungeonSetting_proto_rawDesc
)

func file_CustomDungeonSetting_proto_rawDescGZIP() []byte {
	file_CustomDungeonSetting_proto_rawDescOnce.Do(func() {
		file_CustomDungeonSetting_proto_rawDescData = protoimpl.X.CompressGZIP(file_CustomDungeonSetting_proto_rawDescData)
	})
	return file_CustomDungeonSetting_proto_rawDescData
}

var file_CustomDungeonSetting_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CustomDungeonSetting_proto_goTypes = []interface{}{
	(*CustomDungeonSetting)(nil), // 0: CustomDungeonSetting
}
var file_CustomDungeonSetting_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_CustomDungeonSetting_proto_init() }
func file_CustomDungeonSetting_proto_init() {
	if File_CustomDungeonSetting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CustomDungeonSetting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomDungeonSetting); i {
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
			RawDescriptor: file_CustomDungeonSetting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CustomDungeonSetting_proto_goTypes,
		DependencyIndexes: file_CustomDungeonSetting_proto_depIdxs,
		MessageInfos:      file_CustomDungeonSetting_proto_msgTypes,
	}.Build()
	File_CustomDungeonSetting_proto = out.File
	file_CustomDungeonSetting_proto_rawDesc = nil
	file_CustomDungeonSetting_proto_goTypes = nil
	file_CustomDungeonSetting_proto_depIdxs = nil
}
