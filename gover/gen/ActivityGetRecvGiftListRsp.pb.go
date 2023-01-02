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
// source: ActivityGetRecvGiftListRsp.proto

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

type ActivityGetRecvGiftListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScheduleId   uint32                  `protobuf:"varint,7,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
	Retcode      int32                   `protobuf:"varint,9,opt,name=retcode,proto3" json:"retcode,omitempty"`
	RecvGiftList []*ActivityRecvGiftData `protobuf:"bytes,10,rep,name=recv_gift_list,json=recvGiftList,proto3" json:"recv_gift_list,omitempty"`
}

func (x *ActivityGetRecvGiftListRsp) Reset() {
	*x = ActivityGetRecvGiftListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ActivityGetRecvGiftListRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityGetRecvGiftListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityGetRecvGiftListRsp) ProtoMessage() {}

func (x *ActivityGetRecvGiftListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ActivityGetRecvGiftListRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityGetRecvGiftListRsp.ProtoReflect.Descriptor instead.
func (*ActivityGetRecvGiftListRsp) Descriptor() ([]byte, []int) {
	return file_ActivityGetRecvGiftListRsp_proto_rawDescGZIP(), []int{0}
}

func (x *ActivityGetRecvGiftListRsp) GetScheduleId() uint32 {
	if x != nil {
		return x.ScheduleId
	}
	return 0
}

func (x *ActivityGetRecvGiftListRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *ActivityGetRecvGiftListRsp) GetRecvGiftList() []*ActivityRecvGiftData {
	if x != nil {
		return x.RecvGiftList
	}
	return nil
}

var File_ActivityGetRecvGiftListRsp_proto protoreflect.FileDescriptor

var file_ActivityGetRecvGiftListRsp_proto_rawDesc = []byte{
	0x0a, 0x20, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63,
	0x76, 0x47, 0x69, 0x66, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x63, 0x76,
	0x47, 0x69, 0x66, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94,
	0x01, 0x0a, 0x1a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x63, 0x76, 0x47, 0x69, 0x66, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x3b, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x76,
	0x5f, 0x67, 0x69, 0x66, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x63, 0x76, 0x47,
	0x69, 0x66, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x76, 0x47, 0x69, 0x66,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ActivityGetRecvGiftListRsp_proto_rawDescOnce sync.Once
	file_ActivityGetRecvGiftListRsp_proto_rawDescData = file_ActivityGetRecvGiftListRsp_proto_rawDesc
)

func file_ActivityGetRecvGiftListRsp_proto_rawDescGZIP() []byte {
	file_ActivityGetRecvGiftListRsp_proto_rawDescOnce.Do(func() {
		file_ActivityGetRecvGiftListRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_ActivityGetRecvGiftListRsp_proto_rawDescData)
	})
	return file_ActivityGetRecvGiftListRsp_proto_rawDescData
}

var file_ActivityGetRecvGiftListRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ActivityGetRecvGiftListRsp_proto_goTypes = []interface{}{
	(*ActivityGetRecvGiftListRsp)(nil), // 0: ActivityGetRecvGiftListRsp
	(*ActivityRecvGiftData)(nil),       // 1: ActivityRecvGiftData
}
var file_ActivityGetRecvGiftListRsp_proto_depIdxs = []int32{
	1, // 0: ActivityGetRecvGiftListRsp.recv_gift_list:type_name -> ActivityRecvGiftData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ActivityGetRecvGiftListRsp_proto_init() }
func file_ActivityGetRecvGiftListRsp_proto_init() {
	if File_ActivityGetRecvGiftListRsp_proto != nil {
		return
	}
	file_ActivityRecvGiftData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ActivityGetRecvGiftListRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityGetRecvGiftListRsp); i {
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
			RawDescriptor: file_ActivityGetRecvGiftListRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ActivityGetRecvGiftListRsp_proto_goTypes,
		DependencyIndexes: file_ActivityGetRecvGiftListRsp_proto_depIdxs,
		MessageInfos:      file_ActivityGetRecvGiftListRsp_proto_msgTypes,
	}.Build()
	File_ActivityGetRecvGiftListRsp_proto = out.File
	file_ActivityGetRecvGiftListRsp_proto_rawDesc = nil
	file_ActivityGetRecvGiftListRsp_proto_goTypes = nil
	file_ActivityGetRecvGiftListRsp_proto_depIdxs = nil
}
