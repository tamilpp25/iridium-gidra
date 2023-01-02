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
// source: Unk3300_PPKPCOCOMDH.proto

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

type Unk3300_PPKPCOCOMDH struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ControllerId uint32 `protobuf:"varint,12,opt,name=controller_id,json=controllerId,proto3" json:"controller_id,omitempty"`
	OpSeq        uint32 `protobuf:"varint,13,opt,name=op_seq,json=opSeq,proto3" json:"op_seq,omitempty"`
	// Types that are assignable to Detail:
	//
	//	*Unk3300_PPKPCOCOMDH_GmMsg
	//	*Unk3300_PPKPCOCOMDH_Duel
	Detail isUnk3300_PPKPCOCOMDH_Detail `protobuf_oneof:"detail"`
}

func (x *Unk3300_PPKPCOCOMDH) Reset() {
	*x = Unk3300_PPKPCOCOMDH{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk3300_PPKPCOCOMDH_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk3300_PPKPCOCOMDH) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk3300_PPKPCOCOMDH) ProtoMessage() {}

func (x *Unk3300_PPKPCOCOMDH) ProtoReflect() protoreflect.Message {
	mi := &file_Unk3300_PPKPCOCOMDH_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk3300_PPKPCOCOMDH.ProtoReflect.Descriptor instead.
func (*Unk3300_PPKPCOCOMDH) Descriptor() ([]byte, []int) {
	return file_Unk3300_PPKPCOCOMDH_proto_rawDescGZIP(), []int{0}
}

func (x *Unk3300_PPKPCOCOMDH) GetControllerId() uint32 {
	if x != nil {
		return x.ControllerId
	}
	return 0
}

func (x *Unk3300_PPKPCOCOMDH) GetOpSeq() uint32 {
	if x != nil {
		return x.OpSeq
	}
	return 0
}

func (m *Unk3300_PPKPCOCOMDH) GetDetail() isUnk3300_PPKPCOCOMDH_Detail {
	if m != nil {
		return m.Detail
	}
	return nil
}

func (x *Unk3300_PPKPCOCOMDH) GetGmMsg() string {
	if x, ok := x.GetDetail().(*Unk3300_PPKPCOCOMDH_GmMsg); ok {
		return x.GmMsg
	}
	return ""
}

func (x *Unk3300_PPKPCOCOMDH) GetDuel() *GCGDuel {
	if x, ok := x.GetDetail().(*Unk3300_PPKPCOCOMDH_Duel); ok {
		return x.Duel
	}
	return nil
}

type isUnk3300_PPKPCOCOMDH_Detail interface {
	isUnk3300_PPKPCOCOMDH_Detail()
}

type Unk3300_PPKPCOCOMDH_GmMsg struct {
	GmMsg string `protobuf:"bytes,2,opt,name=gm_msg,json=gmMsg,proto3,oneof"`
}

type Unk3300_PPKPCOCOMDH_Duel struct {
	Duel *GCGDuel `protobuf:"bytes,14,opt,name=duel,proto3,oneof"`
}

func (*Unk3300_PPKPCOCOMDH_GmMsg) isUnk3300_PPKPCOCOMDH_Detail() {}

func (*Unk3300_PPKPCOCOMDH_Duel) isUnk3300_PPKPCOCOMDH_Detail() {}

var File_Unk3300_PPKPCOCOMDH_proto protoreflect.FileDescriptor

var file_Unk3300_PPKPCOCOMDH_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x50, 0x50, 0x4b, 0x50, 0x43, 0x4f,
	0x43, 0x4f, 0x4d, 0x44, 0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x47, 0x43, 0x47,
	0x44, 0x75, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x13, 0x55,
	0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x50, 0x50, 0x4b, 0x50, 0x43, 0x4f, 0x43, 0x4f, 0x4d,
	0x44, 0x48, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x70, 0x5f, 0x73, 0x65,
	0x71, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6f, 0x70, 0x53, 0x65, 0x71, 0x12, 0x17,
	0x0a, 0x06, 0x67, 0x6d, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x05, 0x67, 0x6d, 0x4d, 0x73, 0x67, 0x12, 0x1e, 0x0a, 0x04, 0x64, 0x75, 0x65, 0x6c, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x47, 0x43, 0x47, 0x44, 0x75, 0x65, 0x6c, 0x48,
	0x00, 0x52, 0x04, 0x64, 0x75, 0x65, 0x6c, 0x42, 0x08, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_Unk3300_PPKPCOCOMDH_proto_rawDescOnce sync.Once
	file_Unk3300_PPKPCOCOMDH_proto_rawDescData = file_Unk3300_PPKPCOCOMDH_proto_rawDesc
)

func file_Unk3300_PPKPCOCOMDH_proto_rawDescGZIP() []byte {
	file_Unk3300_PPKPCOCOMDH_proto_rawDescOnce.Do(func() {
		file_Unk3300_PPKPCOCOMDH_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk3300_PPKPCOCOMDH_proto_rawDescData)
	})
	return file_Unk3300_PPKPCOCOMDH_proto_rawDescData
}

var file_Unk3300_PPKPCOCOMDH_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Unk3300_PPKPCOCOMDH_proto_goTypes = []interface{}{
	(*Unk3300_PPKPCOCOMDH)(nil), // 0: Unk3300_PPKPCOCOMDH
	(*GCGDuel)(nil),             // 1: GCGDuel
}
var file_Unk3300_PPKPCOCOMDH_proto_depIdxs = []int32{
	1, // 0: Unk3300_PPKPCOCOMDH.duel:type_name -> GCGDuel
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Unk3300_PPKPCOCOMDH_proto_init() }
func file_Unk3300_PPKPCOCOMDH_proto_init() {
	if File_Unk3300_PPKPCOCOMDH_proto != nil {
		return
	}
	file_GCGDuel_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Unk3300_PPKPCOCOMDH_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk3300_PPKPCOCOMDH); i {
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
	file_Unk3300_PPKPCOCOMDH_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Unk3300_PPKPCOCOMDH_GmMsg)(nil),
		(*Unk3300_PPKPCOCOMDH_Duel)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Unk3300_PPKPCOCOMDH_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk3300_PPKPCOCOMDH_proto_goTypes,
		DependencyIndexes: file_Unk3300_PPKPCOCOMDH_proto_depIdxs,
		MessageInfos:      file_Unk3300_PPKPCOCOMDH_proto_msgTypes,
	}.Build()
	File_Unk3300_PPKPCOCOMDH_proto = out.File
	file_Unk3300_PPKPCOCOMDH_proto_rawDesc = nil
	file_Unk3300_PPKPCOCOMDH_proto_goTypes = nil
	file_Unk3300_PPKPCOCOMDH_proto_depIdxs = nil
}
