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
// source: ServerLogNotify.proto

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

// CmdId: 84
// Obf: PAIOIAPJEEM
type ServerLogNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FHJIOJOGJHJ ServerLogType  `protobuf:"varint,1,opt,name=FHJIOJOGJHJ,proto3,enum=ServerLogType" json:"FHJIOJOGJHJ,omitempty"`
	ServerLog   string         `protobuf:"bytes,4,opt,name=server_log,json=serverLog,proto3" json:"server_log,omitempty"`
	DBILBLAGLPC ServerLogLevel `protobuf:"varint,9,opt,name=DBILBLAGLPC,proto3,enum=ServerLogLevel" json:"DBILBLAGLPC,omitempty"`
}

func (x *ServerLogNotify) Reset() {
	*x = ServerLogNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ServerLogNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerLogNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerLogNotify) ProtoMessage() {}

func (x *ServerLogNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ServerLogNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerLogNotify.ProtoReflect.Descriptor instead.
func (*ServerLogNotify) Descriptor() ([]byte, []int) {
	return file_ServerLogNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ServerLogNotify) GetFHJIOJOGJHJ() ServerLogType {
	if x != nil {
		return x.FHJIOJOGJHJ
	}
	return ServerLogType_SERVER_LOG_NONE
}

func (x *ServerLogNotify) GetServerLog() string {
	if x != nil {
		return x.ServerLog
	}
	return ""
}

func (x *ServerLogNotify) GetDBILBLAGLPC() ServerLogLevel {
	if x != nil {
		return x.DBILBLAGLPC
	}
	return ServerLogLevel_LOG_LEVEL_NONE
}

var File_ServerLogNotify_proto protoreflect.FileDescriptor

var file_ServerLogNotify_proto_rawDesc = []byte{
	0x0a, 0x15, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x67,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x30, 0x0a, 0x0b, 0x46, 0x48, 0x4a, 0x49, 0x4f, 0x4a,
	0x4f, 0x47, 0x4a, 0x48, 0x4a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x46, 0x48, 0x4a,
	0x49, 0x4f, 0x4a, 0x4f, 0x47, 0x4a, 0x48, 0x4a, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x12, 0x31, 0x0a, 0x0b, 0x44, 0x42, 0x49, 0x4c, 0x42,
	0x4c, 0x41, 0x47, 0x4c, 0x50, 0x43, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x0b, 0x44,
	0x42, 0x49, 0x4c, 0x42, 0x4c, 0x41, 0x47, 0x4c, 0x50, 0x43, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67,
	0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ServerLogNotify_proto_rawDescOnce sync.Once
	file_ServerLogNotify_proto_rawDescData = file_ServerLogNotify_proto_rawDesc
)

func file_ServerLogNotify_proto_rawDescGZIP() []byte {
	file_ServerLogNotify_proto_rawDescOnce.Do(func() {
		file_ServerLogNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ServerLogNotify_proto_rawDescData)
	})
	return file_ServerLogNotify_proto_rawDescData
}

var file_ServerLogNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ServerLogNotify_proto_goTypes = []interface{}{
	(*ServerLogNotify)(nil), // 0: ServerLogNotify
	(ServerLogType)(0),      // 1: ServerLogType
	(ServerLogLevel)(0),     // 2: ServerLogLevel
}
var file_ServerLogNotify_proto_depIdxs = []int32{
	1, // 0: ServerLogNotify.FHJIOJOGJHJ:type_name -> ServerLogType
	2, // 1: ServerLogNotify.DBILBLAGLPC:type_name -> ServerLogLevel
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ServerLogNotify_proto_init() }
func file_ServerLogNotify_proto_init() {
	if File_ServerLogNotify_proto != nil {
		return
	}
	file_ServerLogType_proto_init()
	file_ServerLogLevel_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ServerLogNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerLogNotify); i {
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
			RawDescriptor: file_ServerLogNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ServerLogNotify_proto_goTypes,
		DependencyIndexes: file_ServerLogNotify_proto_depIdxs,
		MessageInfos:      file_ServerLogNotify_proto_msgTypes,
	}.Build()
	File_ServerLogNotify_proto = out.File
	file_ServerLogNotify_proto_rawDesc = nil
	file_ServerLogNotify_proto_goTypes = nil
	file_ServerLogNotify_proto_depIdxs = nil
}
