// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: pkg/pb/export.proto

package pb

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ExposureKeyExport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Time window of keys based on arrival to server, in UTC
	StartTimestamp int64 `protobuf:"varint,1,opt,name=startTimestamp,proto3" json:"startTimestamp,omitempty"`
	EndTimestamp   int64 `protobuf:"varint,2,opt,name=endTimestamp,proto3" json:"endTimestamp,omitempty"`
	// Region for which these keys came from (e.g., country)
	Region string                           `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	Keys   []*ExposureKeyExport_ExposureKey `protobuf:"bytes,4,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *ExposureKeyExport) Reset() {
	*x = ExposureKeyExport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_export_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExposureKeyExport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExposureKeyExport) ProtoMessage() {}

func (x *ExposureKeyExport) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_export_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExposureKeyExport.ProtoReflect.Descriptor instead.
func (*ExposureKeyExport) Descriptor() ([]byte, []int) {
	return file_pkg_pb_export_proto_rawDescGZIP(), []int{0}
}

func (x *ExposureKeyExport) GetStartTimestamp() int64 {
	if x != nil {
		return x.StartTimestamp
	}
	return 0
}

func (x *ExposureKeyExport) GetEndTimestamp() int64 {
	if x != nil {
		return x.EndTimestamp
	}
	return 0
}

func (x *ExposureKeyExport) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *ExposureKeyExport) GetKeys() []*ExposureKeyExport_ExposureKey {
	if x != nil {
		return x.Keys
	}
	return nil
}

type ExposureKeyExport_ExposureKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key of infected user
	ExposureKey []byte `protobuf:"bytes,1,opt,name=exposureKey,proto3" json:"exposureKey,omitempty"`
	// Together, next two fields specify window of validity of key
	// E.g., assuming 10 min intervals a key from 2020-04-19 and
	// only good until 1 pm might have intervalStart = 2645424
	// and intervalCount = 78
	IntervalNumber int32 `protobuf:"varint,2,opt,name=intervalNumber,proto3" json:"intervalNumber,omitempty"`
	IntervalCount  int32 `protobuf:"varint,3,opt,name=intervalCount,proto3" json:"intervalCount,omitempty"`
}

func (x *ExposureKeyExport_ExposureKey) Reset() {
	*x = ExposureKeyExport_ExposureKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_export_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExposureKeyExport_ExposureKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExposureKeyExport_ExposureKey) ProtoMessage() {}

func (x *ExposureKeyExport_ExposureKey) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_export_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExposureKeyExport_ExposureKey.ProtoReflect.Descriptor instead.
func (*ExposureKeyExport_ExposureKey) Descriptor() ([]byte, []int) {
	return file_pkg_pb_export_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ExposureKeyExport_ExposureKey) GetExposureKey() []byte {
	if x != nil {
		return x.ExposureKey
	}
	return nil
}

func (x *ExposureKeyExport_ExposureKey) GetIntervalNumber() int32 {
	if x != nil {
		return x.IntervalNumber
	}
	return 0
}

func (x *ExposureKeyExport_ExposureKey) GetIntervalCount() int32 {
	if x != nil {
		return x.IntervalCount
	}
	return 0
}

var File_pkg_pb_export_proto protoreflect.FileDescriptor

var file_pkg_pb_export_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x02, 0x0a, 0x11, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x75,
	0x72, 0x65, 0x4b, 0x65, 0x79, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x65, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12,
	0x32, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x45, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x45, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x04, 0x6b,
	0x65, 0x79, 0x73, 0x1a, 0x7d, 0x0a, 0x0b, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x4b,
	0x65, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x4b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72,
	0x65, 0x4b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_export_proto_rawDescOnce sync.Once
	file_pkg_pb_export_proto_rawDescData = file_pkg_pb_export_proto_rawDesc
)

func file_pkg_pb_export_proto_rawDescGZIP() []byte {
	file_pkg_pb_export_proto_rawDescOnce.Do(func() {
		file_pkg_pb_export_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_export_proto_rawDescData)
	})
	return file_pkg_pb_export_proto_rawDescData
}

var file_pkg_pb_export_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_pb_export_proto_goTypes = []interface{}{
	(*ExposureKeyExport)(nil),             // 0: ExposureKeyExport
	(*ExposureKeyExport_ExposureKey)(nil), // 1: ExposureKeyExport.ExposureKey
}
var file_pkg_pb_export_proto_depIdxs = []int32{
	1, // 0: ExposureKeyExport.keys:type_name -> ExposureKeyExport.ExposureKey
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_pb_export_proto_init() }
func file_pkg_pb_export_proto_init() {
	if File_pkg_pb_export_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_export_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExposureKeyExport); i {
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
		file_pkg_pb_export_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExposureKeyExport_ExposureKey); i {
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
			RawDescriptor: file_pkg_pb_export_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_pb_export_proto_goTypes,
		DependencyIndexes: file_pkg_pb_export_proto_depIdxs,
		MessageInfos:      file_pkg_pb_export_proto_msgTypes,
	}.Build()
	File_pkg_pb_export_proto = out.File
	file_pkg_pb_export_proto_rawDesc = nil
	file_pkg_pb_export_proto_goTypes = nil
	file_pkg_pb_export_proto_depIdxs = nil
}
