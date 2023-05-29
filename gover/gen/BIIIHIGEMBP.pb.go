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
// source: BIIIHIGEMBP.proto

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

// CmdId: 2317
type BIIIHIGEMBP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter         *QueryFilter `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	SceneId        uint32       `protobuf:"varint,4,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	DestinationPos *Vector      `protobuf:"bytes,1,opt,name=destination_pos,json=destinationPos,proto3" json:"destination_pos,omitempty"`
	QueryId        int32        `protobuf:"varint,12,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
	Uid            int32        `protobuf:"varint,5,opt,name=uid,proto3" json:"uid,omitempty"`
	OCDLHELPCKN    *Vector3Int  `protobuf:"bytes,13,opt,name=OCDLHELPCKN,proto3" json:"OCDLHELPCKN,omitempty"`
	SourcePos      *Vector      `protobuf:"bytes,9,opt,name=source_pos,json=sourcePos,proto3" json:"source_pos,omitempty"`
	COBIGJHHLJJ    *Vector3Int  `protobuf:"bytes,15,opt,name=COBIGJHHLJJ,proto3" json:"COBIGJHHLJJ,omitempty"`
}

func (x *BIIIHIGEMBP) Reset() {
	*x = BIIIHIGEMBP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BIIIHIGEMBP_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BIIIHIGEMBP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BIIIHIGEMBP) ProtoMessage() {}

func (x *BIIIHIGEMBP) ProtoReflect() protoreflect.Message {
	mi := &file_BIIIHIGEMBP_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BIIIHIGEMBP.ProtoReflect.Descriptor instead.
func (*BIIIHIGEMBP) Descriptor() ([]byte, []int) {
	return file_BIIIHIGEMBP_proto_rawDescGZIP(), []int{0}
}

func (x *BIIIHIGEMBP) GetFilter() *QueryFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *BIIIHIGEMBP) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *BIIIHIGEMBP) GetDestinationPos() *Vector {
	if x != nil {
		return x.DestinationPos
	}
	return nil
}

func (x *BIIIHIGEMBP) GetQueryId() int32 {
	if x != nil {
		return x.QueryId
	}
	return 0
}

func (x *BIIIHIGEMBP) GetUid() int32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *BIIIHIGEMBP) GetOCDLHELPCKN() *Vector3Int {
	if x != nil {
		return x.OCDLHELPCKN
	}
	return nil
}

func (x *BIIIHIGEMBP) GetSourcePos() *Vector {
	if x != nil {
		return x.SourcePos
	}
	return nil
}

func (x *BIIIHIGEMBP) GetCOBIGJHHLJJ() *Vector3Int {
	if x != nil {
		return x.COBIGJHHLJJ
	}
	return nil
}

var File_BIIIHIGEMBP_proto protoreflect.FileDescriptor

var file_BIIIHIGEMBP_proto_rawDesc = []byte{
	0x0a, 0x11, 0x42, 0x49, 0x49, 0x49, 0x48, 0x49, 0x47, 0x45, 0x4d, 0x42, 0x50, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x49, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x02, 0x0a, 0x0b, 0x42, 0x49, 0x49, 0x49, 0x48,
	0x49, 0x47, 0x45, 0x4d, 0x42, 0x50, 0x12, 0x24, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08,
	0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x73, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x0f, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x0b, 0x4f, 0x43, 0x44, 0x4c, 0x48, 0x45,
	0x4c, 0x50, 0x43, 0x4b, 0x4e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x33, 0x49, 0x6e, 0x74, 0x52, 0x0b, 0x4f, 0x43, 0x44, 0x4c, 0x48, 0x45,
	0x4c, 0x50, 0x43, 0x4b, 0x4e, 0x12, 0x26, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x70, 0x6f, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x73, 0x12, 0x2d, 0x0a,
	0x0b, 0x43, 0x4f, 0x42, 0x49, 0x47, 0x4a, 0x48, 0x48, 0x4c, 0x4a, 0x4a, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x49, 0x6e, 0x74, 0x52,
	0x0b, 0x43, 0x4f, 0x42, 0x49, 0x47, 0x4a, 0x48, 0x48, 0x4c, 0x4a, 0x4a, 0x42, 0x06, 0x5a, 0x04,
	0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BIIIHIGEMBP_proto_rawDescOnce sync.Once
	file_BIIIHIGEMBP_proto_rawDescData = file_BIIIHIGEMBP_proto_rawDesc
)

func file_BIIIHIGEMBP_proto_rawDescGZIP() []byte {
	file_BIIIHIGEMBP_proto_rawDescOnce.Do(func() {
		file_BIIIHIGEMBP_proto_rawDescData = protoimpl.X.CompressGZIP(file_BIIIHIGEMBP_proto_rawDescData)
	})
	return file_BIIIHIGEMBP_proto_rawDescData
}

var file_BIIIHIGEMBP_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BIIIHIGEMBP_proto_goTypes = []interface{}{
	(*BIIIHIGEMBP)(nil), // 0: BIIIHIGEMBP
	(*QueryFilter)(nil), // 1: QueryFilter
	(*Vector)(nil),      // 2: Vector
	(*Vector3Int)(nil),  // 3: Vector3Int
}
var file_BIIIHIGEMBP_proto_depIdxs = []int32{
	1, // 0: BIIIHIGEMBP.filter:type_name -> QueryFilter
	2, // 1: BIIIHIGEMBP.destination_pos:type_name -> Vector
	3, // 2: BIIIHIGEMBP.OCDLHELPCKN:type_name -> Vector3Int
	2, // 3: BIIIHIGEMBP.source_pos:type_name -> Vector
	3, // 4: BIIIHIGEMBP.COBIGJHHLJJ:type_name -> Vector3Int
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_BIIIHIGEMBP_proto_init() }
func file_BIIIHIGEMBP_proto_init() {
	if File_BIIIHIGEMBP_proto != nil {
		return
	}
	file_QueryFilter_proto_init()
	file_Vector_proto_init()
	file_Vector3Int_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BIIIHIGEMBP_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BIIIHIGEMBP); i {
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
			RawDescriptor: file_BIIIHIGEMBP_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BIIIHIGEMBP_proto_goTypes,
		DependencyIndexes: file_BIIIHIGEMBP_proto_depIdxs,
		MessageInfos:      file_BIIIHIGEMBP_proto_msgTypes,
	}.Build()
	File_BIIIHIGEMBP_proto = out.File
	file_BIIIHIGEMBP_proto_rawDesc = nil
	file_BIIIHIGEMBP_proto_goTypes = nil
	file_BIIIHIGEMBP_proto_depIdxs = nil
}
