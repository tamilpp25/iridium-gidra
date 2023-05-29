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
// source: GetScenePointRsp.proto

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

// CmdId: 294
// Obf: OKEJNDKCNKG
type GetScenePointRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnhidePointList             []uint32 `protobuf:"varint,12,rep,packed,name=unhide_point_list,json=unhidePointList,proto3" json:"unhide_point_list,omitempty"`
	GPMHJGJLAFH                 bool     `protobuf:"varint,2,opt,name=GPMHJGJLAFH,proto3" json:"GPMHJGJLAFH,omitempty"`
	UnlockedPointList           []uint32 `protobuf:"varint,3,rep,packed,name=unlocked_point_list,json=unlockedPointList,proto3" json:"unlocked_point_list,omitempty"`
	ToBeExploreDungeonEntryList []uint32 `protobuf:"varint,10,rep,packed,name=toBeExploreDungeonEntryList,proto3" json:"toBeExploreDungeonEntryList,omitempty"`
	SceneId                     uint32   `protobuf:"varint,5,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	HidePointList               []uint32 `protobuf:"varint,4,rep,packed,name=hide_point_list,json=hidePointList,proto3" json:"hide_point_list,omitempty"`
	LockedPointList             []uint32 `protobuf:"varint,8,rep,packed,name=locked_point_list,json=lockedPointList,proto3" json:"locked_point_list,omitempty"`
	GroupUnlimitPointList       []uint32 `protobuf:"varint,15,rep,packed,name=groupUnlimitPointList,proto3" json:"groupUnlimitPointList,omitempty"`
	BelongUid                   uint32   `protobuf:"varint,11,opt,name=belong_uid,json=belongUid,proto3" json:"belong_uid,omitempty"`
	NotInteractDungeonEntryList []uint32 `protobuf:"varint,1,rep,packed,name=notInteractDungeonEntryList,proto3" json:"notInteractDungeonEntryList,omitempty"`
	Retcode                     int32    `protobuf:"varint,7,opt,name=retcode,proto3" json:"retcode,omitempty"`
	NotExploredDungeonEntryList []uint32 `protobuf:"varint,6,rep,packed,name=notExploredDungeonEntryList,proto3" json:"notExploredDungeonEntryList,omitempty"`
	UnlockAreaList              []uint32 `protobuf:"varint,14,rep,packed,name=unlockAreaList,proto3" json:"unlockAreaList,omitempty"`
}

func (x *GetScenePointRsp) Reset() {
	*x = GetScenePointRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetScenePointRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetScenePointRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScenePointRsp) ProtoMessage() {}

func (x *GetScenePointRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetScenePointRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetScenePointRsp.ProtoReflect.Descriptor instead.
func (*GetScenePointRsp) Descriptor() ([]byte, []int) {
	return file_GetScenePointRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetScenePointRsp) GetUnhidePointList() []uint32 {
	if x != nil {
		return x.UnhidePointList
	}
	return nil
}

func (x *GetScenePointRsp) GetGPMHJGJLAFH() bool {
	if x != nil {
		return x.GPMHJGJLAFH
	}
	return false
}

func (x *GetScenePointRsp) GetUnlockedPointList() []uint32 {
	if x != nil {
		return x.UnlockedPointList
	}
	return nil
}

func (x *GetScenePointRsp) GetToBeExploreDungeonEntryList() []uint32 {
	if x != nil {
		return x.ToBeExploreDungeonEntryList
	}
	return nil
}

func (x *GetScenePointRsp) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *GetScenePointRsp) GetHidePointList() []uint32 {
	if x != nil {
		return x.HidePointList
	}
	return nil
}

func (x *GetScenePointRsp) GetLockedPointList() []uint32 {
	if x != nil {
		return x.LockedPointList
	}
	return nil
}

func (x *GetScenePointRsp) GetGroupUnlimitPointList() []uint32 {
	if x != nil {
		return x.GroupUnlimitPointList
	}
	return nil
}

func (x *GetScenePointRsp) GetBelongUid() uint32 {
	if x != nil {
		return x.BelongUid
	}
	return 0
}

func (x *GetScenePointRsp) GetNotInteractDungeonEntryList() []uint32 {
	if x != nil {
		return x.NotInteractDungeonEntryList
	}
	return nil
}

func (x *GetScenePointRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetScenePointRsp) GetNotExploredDungeonEntryList() []uint32 {
	if x != nil {
		return x.NotExploredDungeonEntryList
	}
	return nil
}

func (x *GetScenePointRsp) GetUnlockAreaList() []uint32 {
	if x != nil {
		return x.UnlockAreaList
	}
	return nil
}

var File_GetScenePointRsp_proto protoreflect.FileDescriptor

var file_GetScenePointRsp_proto_rawDesc = []byte{
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x04, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x53, 0x63, 0x65, 0x6e, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x12, 0x2a, 0x0a,
	0x11, 0x75, 0x6e, 0x68, 0x69, 0x64, 0x65, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0f, 0x75, 0x6e, 0x68, 0x69, 0x64, 0x65,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x50, 0x4d,
	0x48, 0x4a, 0x47, 0x4a, 0x4c, 0x41, 0x46, 0x48, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x47, 0x50, 0x4d, 0x48, 0x4a, 0x47, 0x4a, 0x4c, 0x41, 0x46, 0x48, 0x12, 0x2e, 0x0a, 0x13, 0x75,
	0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x11, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b,
	0x65, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x1b, 0x74,
	0x6f, 0x42, 0x65, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f,
	0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x1b, 0x74, 0x6f, 0x42, 0x65, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x44, 0x75, 0x6e,
	0x67, 0x65, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x68, 0x69, 0x64, 0x65,
	0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x0d, 0x68, 0x69, 0x64, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0f, 0x6c, 0x6f, 0x63,
	0x6b, 0x65, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x15,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x6e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x15, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x55, 0x6e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x75, 0x69, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x62, 0x65, 0x6c, 0x6f, 0x6e, 0x67, 0x55, 0x69,
	0x64, 0x12, 0x40, 0x0a, 0x1b, 0x6e, 0x6f, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x1b, 0x6e, 0x6f, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x40, 0x0a,
	0x1b, 0x6e, 0x6f, 0x74, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x64, 0x44, 0x75, 0x6e, 0x67,
	0x65, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x1b, 0x6e, 0x6f, 0x74, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x64, 0x44,
	0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x26, 0x0a, 0x0e, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x72, 0x65, 0x61, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0e, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x41,
	0x72, 0x65, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetScenePointRsp_proto_rawDescOnce sync.Once
	file_GetScenePointRsp_proto_rawDescData = file_GetScenePointRsp_proto_rawDesc
)

func file_GetScenePointRsp_proto_rawDescGZIP() []byte {
	file_GetScenePointRsp_proto_rawDescOnce.Do(func() {
		file_GetScenePointRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetScenePointRsp_proto_rawDescData)
	})
	return file_GetScenePointRsp_proto_rawDescData
}

var file_GetScenePointRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetScenePointRsp_proto_goTypes = []interface{}{
	(*GetScenePointRsp)(nil), // 0: GetScenePointRsp
}
var file_GetScenePointRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetScenePointRsp_proto_init() }
func file_GetScenePointRsp_proto_init() {
	if File_GetScenePointRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetScenePointRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetScenePointRsp); i {
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
			RawDescriptor: file_GetScenePointRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetScenePointRsp_proto_goTypes,
		DependencyIndexes: file_GetScenePointRsp_proto_depIdxs,
		MessageInfos:      file_GetScenePointRsp_proto_msgTypes,
	}.Build()
	File_GetScenePointRsp_proto = out.File
	file_GetScenePointRsp_proto_rawDesc = nil
	file_GetScenePointRsp_proto_goTypes = nil
	file_GetScenePointRsp_proto_depIdxs = nil
}
