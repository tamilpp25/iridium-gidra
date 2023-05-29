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
// source: QuickUseWidgetRsp.proto

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

// CmdId: 4253
// Obf: PIBILHMMMKA
type QuickUseWidgetRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaterialId uint32 `protobuf:"varint,9,opt,name=material_id,json=materialId,proto3" json:"material_id,omitempty"`
	Retcode    int32  `protobuf:"varint,5,opt,name=retcode,proto3" json:"retcode,omitempty"`
	// Types that are assignable to Param:
	//
	//	*QuickUseWidgetRsp_DetectorData
	//	*QuickUseWidgetRsp_ClientCollectorData
	//	*QuickUseWidgetRsp_SkyCrystalDetectorQuickUseResult
	Param isQuickUseWidgetRsp_Param `protobuf_oneof:"param"`
}

func (x *QuickUseWidgetRsp) Reset() {
	*x = QuickUseWidgetRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_QuickUseWidgetRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuickUseWidgetRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuickUseWidgetRsp) ProtoMessage() {}

func (x *QuickUseWidgetRsp) ProtoReflect() protoreflect.Message {
	mi := &file_QuickUseWidgetRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuickUseWidgetRsp.ProtoReflect.Descriptor instead.
func (*QuickUseWidgetRsp) Descriptor() ([]byte, []int) {
	return file_QuickUseWidgetRsp_proto_rawDescGZIP(), []int{0}
}

func (x *QuickUseWidgetRsp) GetMaterialId() uint32 {
	if x != nil {
		return x.MaterialId
	}
	return 0
}

func (x *QuickUseWidgetRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (m *QuickUseWidgetRsp) GetParam() isQuickUseWidgetRsp_Param {
	if m != nil {
		return m.Param
	}
	return nil
}

func (x *QuickUseWidgetRsp) GetDetectorData() *OneoffGatherPointDetectorData {
	if x, ok := x.GetParam().(*QuickUseWidgetRsp_DetectorData); ok {
		return x.DetectorData
	}
	return nil
}

func (x *QuickUseWidgetRsp) GetClientCollectorData() *ClientCollectorData {
	if x, ok := x.GetParam().(*QuickUseWidgetRsp_ClientCollectorData); ok {
		return x.ClientCollectorData
	}
	return nil
}

func (x *QuickUseWidgetRsp) GetSkyCrystalDetectorQuickUseResult() *SkyCrystalDetectorQuickUseResult {
	if x, ok := x.GetParam().(*QuickUseWidgetRsp_SkyCrystalDetectorQuickUseResult); ok {
		return x.SkyCrystalDetectorQuickUseResult
	}
	return nil
}

type isQuickUseWidgetRsp_Param interface {
	isQuickUseWidgetRsp_Param()
}

type QuickUseWidgetRsp_DetectorData struct {
	DetectorData *OneoffGatherPointDetectorData `protobuf:"bytes,4,opt,name=detector_data,json=detectorData,proto3,oneof"`
}

type QuickUseWidgetRsp_ClientCollectorData struct {
	ClientCollectorData *ClientCollectorData `protobuf:"bytes,11,opt,name=client_collector_data,json=clientCollectorData,proto3,oneof"`
}

type QuickUseWidgetRsp_SkyCrystalDetectorQuickUseResult struct {
	SkyCrystalDetectorQuickUseResult *SkyCrystalDetectorQuickUseResult `protobuf:"bytes,150503,opt,name=sky_crystal_detector_quick_use_result,json=skyCrystalDetectorQuickUseResult,proto3,oneof"`
}

func (*QuickUseWidgetRsp_DetectorData) isQuickUseWidgetRsp_Param() {}

func (*QuickUseWidgetRsp_ClientCollectorData) isQuickUseWidgetRsp_Param() {}

func (*QuickUseWidgetRsp_SkyCrystalDetectorQuickUseResult) isQuickUseWidgetRsp_Param() {}

var File_QuickUseWidgetRsp_proto protoreflect.FileDescriptor

var file_QuickUseWidgetRsp_proto_rawDesc = []byte{
	0x0a, 0x17, 0x51, 0x75, 0x69, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74,
	0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x4f, 0x6e, 0x65, 0x6f, 0x66,
	0x66, 0x47, 0x61, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x53, 0x6b, 0x79, 0x43, 0x72,
	0x79, 0x73, 0x74, 0x61, 0x6c, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x51, 0x75, 0x69,
	0x63, 0x6b, 0x55, 0x73, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe2, 0x02, 0x0a, 0x11, 0x51, 0x75, 0x69, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x57, 0x69,
	0x64, 0x67, 0x65, 0x74, 0x52, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x45, 0x0a, 0x0d, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x4f, 0x6e, 0x65, 0x6f,
	0x66, 0x66, 0x47, 0x61, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x0c, 0x64, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x4a, 0x0a, 0x15, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00,
	0x52, 0x13, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x76, 0x0a, 0x25, 0x73, 0x6b, 0x79, 0x5f, 0x63, 0x72, 0x79,
	0x73, 0x74, 0x61, 0x6c, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x71, 0x75,
	0x69, 0x63, 0x6b, 0x5f, 0x75, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0xe7,
	0x97, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x53, 0x6b, 0x79, 0x43, 0x72, 0x79, 0x73,
	0x74, 0x61, 0x6c, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x51, 0x75, 0x69, 0x63, 0x6b,
	0x55, 0x73, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x20, 0x73, 0x6b, 0x79,
	0x43, 0x72, 0x79, 0x73, 0x74, 0x61, 0x6c, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x51,
	0x75, 0x69, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x07, 0x0a,
	0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_QuickUseWidgetRsp_proto_rawDescOnce sync.Once
	file_QuickUseWidgetRsp_proto_rawDescData = file_QuickUseWidgetRsp_proto_rawDesc
)

func file_QuickUseWidgetRsp_proto_rawDescGZIP() []byte {
	file_QuickUseWidgetRsp_proto_rawDescOnce.Do(func() {
		file_QuickUseWidgetRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_QuickUseWidgetRsp_proto_rawDescData)
	})
	return file_QuickUseWidgetRsp_proto_rawDescData
}

var file_QuickUseWidgetRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_QuickUseWidgetRsp_proto_goTypes = []interface{}{
	(*QuickUseWidgetRsp)(nil),                // 0: QuickUseWidgetRsp
	(*OneoffGatherPointDetectorData)(nil),    // 1: OneoffGatherPointDetectorData
	(*ClientCollectorData)(nil),              // 2: ClientCollectorData
	(*SkyCrystalDetectorQuickUseResult)(nil), // 3: SkyCrystalDetectorQuickUseResult
}
var file_QuickUseWidgetRsp_proto_depIdxs = []int32{
	1, // 0: QuickUseWidgetRsp.detector_data:type_name -> OneoffGatherPointDetectorData
	2, // 1: QuickUseWidgetRsp.client_collector_data:type_name -> ClientCollectorData
	3, // 2: QuickUseWidgetRsp.sky_crystal_detector_quick_use_result:type_name -> SkyCrystalDetectorQuickUseResult
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_QuickUseWidgetRsp_proto_init() }
func file_QuickUseWidgetRsp_proto_init() {
	if File_QuickUseWidgetRsp_proto != nil {
		return
	}
	file_OneoffGatherPointDetectorData_proto_init()
	file_ClientCollectorData_proto_init()
	file_SkyCrystalDetectorQuickUseResult_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_QuickUseWidgetRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuickUseWidgetRsp); i {
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
	file_QuickUseWidgetRsp_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*QuickUseWidgetRsp_DetectorData)(nil),
		(*QuickUseWidgetRsp_ClientCollectorData)(nil),
		(*QuickUseWidgetRsp_SkyCrystalDetectorQuickUseResult)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_QuickUseWidgetRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_QuickUseWidgetRsp_proto_goTypes,
		DependencyIndexes: file_QuickUseWidgetRsp_proto_depIdxs,
		MessageInfos:      file_QuickUseWidgetRsp_proto_msgTypes,
	}.Build()
	File_QuickUseWidgetRsp_proto = out.File
	file_QuickUseWidgetRsp_proto_rawDesc = nil
	file_QuickUseWidgetRsp_proto_goTypes = nil
	file_QuickUseWidgetRsp_proto_depIdxs = nil
}
