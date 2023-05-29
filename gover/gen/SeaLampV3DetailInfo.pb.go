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
// source: SeaLampV3DetailInfo.proto

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

// Obf: GNIMFGCJKMP
type SeaLampV3DetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CampInfo   *SeaLampV3CampInfo   `protobuf:"bytes,13,opt,name=camp_info,json=campInfo,proto3" json:"camp_info,omitempty"`
	ShadowInfo *SeaLampV3ShadowInfo `protobuf:"bytes,1,opt,name=shadow_info,json=shadowInfo,proto3" json:"shadow_info,omitempty"`
	RaceInfo   *SeaLampV3RaceInfo   `protobuf:"bytes,5,opt,name=race_info,json=raceInfo,proto3" json:"race_info,omitempty"`
}

func (x *SeaLampV3DetailInfo) Reset() {
	*x = SeaLampV3DetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SeaLampV3DetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeaLampV3DetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeaLampV3DetailInfo) ProtoMessage() {}

func (x *SeaLampV3DetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_SeaLampV3DetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeaLampV3DetailInfo.ProtoReflect.Descriptor instead.
func (*SeaLampV3DetailInfo) Descriptor() ([]byte, []int) {
	return file_SeaLampV3DetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SeaLampV3DetailInfo) GetCampInfo() *SeaLampV3CampInfo {
	if x != nil {
		return x.CampInfo
	}
	return nil
}

func (x *SeaLampV3DetailInfo) GetShadowInfo() *SeaLampV3ShadowInfo {
	if x != nil {
		return x.ShadowInfo
	}
	return nil
}

func (x *SeaLampV3DetailInfo) GetRaceInfo() *SeaLampV3RaceInfo {
	if x != nil {
		return x.RaceInfo
	}
	return nil
}

var File_SeaLampV3DetailInfo_proto protoreflect.FileDescriptor

var file_SeaLampV3DetailInfo_proto_rawDesc = []byte{
	0x0a, 0x19, 0x53, 0x65, 0x61, 0x4c, 0x61, 0x6d, 0x70, 0x56, 0x33, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x53, 0x65, 0x61,
	0x4c, 0x61, 0x6d, 0x70, 0x56, 0x33, 0x43, 0x61, 0x6d, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x53, 0x65, 0x61, 0x4c, 0x61, 0x6d, 0x70, 0x56, 0x33, 0x53,
	0x68, 0x61, 0x64, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x53, 0x65, 0x61, 0x4c, 0x61, 0x6d, 0x70, 0x56, 0x33, 0x52, 0x61, 0x63, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x61,
	0x4c, 0x61, 0x6d, 0x70, 0x56, 0x33, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x2f, 0x0a, 0x09, 0x63, 0x61, 0x6d, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x53, 0x65, 0x61, 0x4c, 0x61, 0x6d, 0x70, 0x56, 0x33, 0x43,
	0x61, 0x6d, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63, 0x61, 0x6d, 0x70, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x35, 0x0a, 0x0b, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x53, 0x65, 0x61, 0x4c, 0x61, 0x6d, 0x70,
	0x56, 0x33, 0x53, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x73, 0x68,
	0x61, 0x64, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2f, 0x0a, 0x09, 0x72, 0x61, 0x63, 0x65,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x53, 0x65,
	0x61, 0x4c, 0x61, 0x6d, 0x70, 0x56, 0x33, 0x52, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x08, 0x72, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SeaLampV3DetailInfo_proto_rawDescOnce sync.Once
	file_SeaLampV3DetailInfo_proto_rawDescData = file_SeaLampV3DetailInfo_proto_rawDesc
)

func file_SeaLampV3DetailInfo_proto_rawDescGZIP() []byte {
	file_SeaLampV3DetailInfo_proto_rawDescOnce.Do(func() {
		file_SeaLampV3DetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_SeaLampV3DetailInfo_proto_rawDescData)
	})
	return file_SeaLampV3DetailInfo_proto_rawDescData
}

var file_SeaLampV3DetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SeaLampV3DetailInfo_proto_goTypes = []interface{}{
	(*SeaLampV3DetailInfo)(nil), // 0: SeaLampV3DetailInfo
	(*SeaLampV3CampInfo)(nil),   // 1: SeaLampV3CampInfo
	(*SeaLampV3ShadowInfo)(nil), // 2: SeaLampV3ShadowInfo
	(*SeaLampV3RaceInfo)(nil),   // 3: SeaLampV3RaceInfo
}
var file_SeaLampV3DetailInfo_proto_depIdxs = []int32{
	1, // 0: SeaLampV3DetailInfo.camp_info:type_name -> SeaLampV3CampInfo
	2, // 1: SeaLampV3DetailInfo.shadow_info:type_name -> SeaLampV3ShadowInfo
	3, // 2: SeaLampV3DetailInfo.race_info:type_name -> SeaLampV3RaceInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_SeaLampV3DetailInfo_proto_init() }
func file_SeaLampV3DetailInfo_proto_init() {
	if File_SeaLampV3DetailInfo_proto != nil {
		return
	}
	file_SeaLampV3CampInfo_proto_init()
	file_SeaLampV3ShadowInfo_proto_init()
	file_SeaLampV3RaceInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SeaLampV3DetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeaLampV3DetailInfo); i {
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
			RawDescriptor: file_SeaLampV3DetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SeaLampV3DetailInfo_proto_goTypes,
		DependencyIndexes: file_SeaLampV3DetailInfo_proto_depIdxs,
		MessageInfos:      file_SeaLampV3DetailInfo_proto_msgTypes,
	}.Build()
	File_SeaLampV3DetailInfo_proto = out.File
	file_SeaLampV3DetailInfo_proto_rawDesc = nil
	file_SeaLampV3DetailInfo_proto_goTypes = nil
	file_SeaLampV3DetailInfo_proto_depIdxs = nil
}
