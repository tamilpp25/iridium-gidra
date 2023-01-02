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
// source: ChallengeRecord.proto

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

type ChallengeRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChallengeRecordType uint32 `protobuf:"varint,2,opt,name=challenge_record_type,json=challengeRecordType,proto3" json:"challenge_record_type,omitempty"`
	ChallengeIndex      uint32 `protobuf:"varint,3,opt,name=challenge_index,json=challengeIndex,proto3" json:"challenge_index,omitempty"`
	ChallengeId         uint32 `protobuf:"varint,9,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
	BestValue           uint32 `protobuf:"varint,4,opt,name=best_value,json=bestValue,proto3" json:"best_value,omitempty"`
}

func (x *ChallengeRecord) Reset() {
	*x = ChallengeRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChallengeRecord_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChallengeRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChallengeRecord) ProtoMessage() {}

func (x *ChallengeRecord) ProtoReflect() protoreflect.Message {
	mi := &file_ChallengeRecord_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChallengeRecord.ProtoReflect.Descriptor instead.
func (*ChallengeRecord) Descriptor() ([]byte, []int) {
	return file_ChallengeRecord_proto_rawDescGZIP(), []int{0}
}

func (x *ChallengeRecord) GetChallengeRecordType() uint32 {
	if x != nil {
		return x.ChallengeRecordType
	}
	return 0
}

func (x *ChallengeRecord) GetChallengeIndex() uint32 {
	if x != nil {
		return x.ChallengeIndex
	}
	return 0
}

func (x *ChallengeRecord) GetChallengeId() uint32 {
	if x != nil {
		return x.ChallengeId
	}
	return 0
}

func (x *ChallengeRecord) GetBestValue() uint32 {
	if x != nil {
		return x.BestValue
	}
	return 0
}

var File_ChallengeRecord_proto protoreflect.FileDescriptor

var file_ChallengeRecord_proto_rawDesc = []byte{
	0x0a, 0x15, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x0f, 0x43, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x63,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x63, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x27, 0x0a, 0x0f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x65, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x62, 0x65, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67,
	0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChallengeRecord_proto_rawDescOnce sync.Once
	file_ChallengeRecord_proto_rawDescData = file_ChallengeRecord_proto_rawDesc
)

func file_ChallengeRecord_proto_rawDescGZIP() []byte {
	file_ChallengeRecord_proto_rawDescOnce.Do(func() {
		file_ChallengeRecord_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChallengeRecord_proto_rawDescData)
	})
	return file_ChallengeRecord_proto_rawDescData
}

var file_ChallengeRecord_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChallengeRecord_proto_goTypes = []interface{}{
	(*ChallengeRecord)(nil), // 0: ChallengeRecord
}
var file_ChallengeRecord_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ChallengeRecord_proto_init() }
func file_ChallengeRecord_proto_init() {
	if File_ChallengeRecord_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ChallengeRecord_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChallengeRecord); i {
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
			RawDescriptor: file_ChallengeRecord_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChallengeRecord_proto_goTypes,
		DependencyIndexes: file_ChallengeRecord_proto_depIdxs,
		MessageInfos:      file_ChallengeRecord_proto_msgTypes,
	}.Build()
	File_ChallengeRecord_proto = out.File
	file_ChallengeRecord_proto_rawDesc = nil
	file_ChallengeRecord_proto_goTypes = nil
	file_ChallengeRecord_proto_depIdxs = nil
}
