// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: opentelemetry/proto/collector/logs/v1/logs_service.proto

// NOTE: This proto is experimental and is subject to change at this point.
// Please do not use it at the moment.

package v1

import (
	v1 "github.com/influxdata/influxdb-observability/otlp/logs/v1"
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

type ExportLogsServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An array of ResourceLogs.
	// For data coming from a single resource this array will typically contain one
	// element. Intermediary nodes (such as OpenTelemetry Collector) that receive
	// data from multiple origins typically batch the data before forwarding further and
	// in that case this array will contain multiple elements.
	ResourceLogs []*v1.ResourceLogs `protobuf:"bytes,1,rep,name=resource_logs,json=resourceLogs,proto3" json:"resource_logs,omitempty"`
}

func (x *ExportLogsServiceRequest) Reset() {
	*x = ExportLogsServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opentelemetry_proto_collector_logs_v1_logs_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportLogsServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportLogsServiceRequest) ProtoMessage() {}

func (x *ExportLogsServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_opentelemetry_proto_collector_logs_v1_logs_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportLogsServiceRequest.ProtoReflect.Descriptor instead.
func (*ExportLogsServiceRequest) Descriptor() ([]byte, []int) {
	return file_opentelemetry_proto_collector_logs_v1_logs_service_proto_rawDescGZIP(), []int{0}
}

func (x *ExportLogsServiceRequest) GetResourceLogs() []*v1.ResourceLogs {
	if x != nil {
		return x.ResourceLogs
	}
	return nil
}

type ExportLogsServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ExportLogsServiceResponse) Reset() {
	*x = ExportLogsServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opentelemetry_proto_collector_logs_v1_logs_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportLogsServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportLogsServiceResponse) ProtoMessage() {}

func (x *ExportLogsServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_opentelemetry_proto_collector_logs_v1_logs_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportLogsServiceResponse.ProtoReflect.Descriptor instead.
func (*ExportLogsServiceResponse) Descriptor() ([]byte, []int) {
	return file_opentelemetry_proto_collector_logs_v1_logs_service_proto_rawDescGZIP(), []int{1}
}

var File_opentelemetry_proto_collector_logs_v1_logs_service_proto protoreflect.FileDescriptor

var file_opentelemetry_proto_collector_logs_v1_logs_service_proto_rawDesc = []byte{
	0x0a, 0x38, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f,
	0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x26, 0x6f, 0x70, 0x65, 0x6e,
	0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x73, 0x0a, 0x18, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6f, 0x67, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x57,
	0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x22, 0x1b, 0x0a, 0x19, 0x45, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x4c, 0x6f, 0x67, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0xaf, 0x01, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x9f, 0x01, 0x0a, 0x06, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x48, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x49, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x4c, 0x6f, 0x67, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x8c, 0x01, 0x0a, 0x31, 0x69, 0x6f, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x4c, 0x6f,
	0x67, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66,
	0x6c, 0x75, 0x78, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x78, 0x64, 0x62,
	0x2d, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2f, 0x6f,
	0x74, 0x6c, 0x70, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x6c, 0x6f,
	0x67, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_opentelemetry_proto_collector_logs_v1_logs_service_proto_rawDescOnce sync.Once
	file_opentelemetry_proto_collector_logs_v1_logs_service_proto_rawDescData = file_opentelemetry_proto_collector_logs_v1_logs_service_proto_rawDesc
)

func file_opentelemetry_proto_collector_logs_v1_logs_service_proto_rawDescGZIP() []byte {
	file_opentelemetry_proto_collector_logs_v1_logs_service_proto_rawDescOnce.Do(func() {
		file_opentelemetry_proto_collector_logs_v1_logs_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_opentelemetry_proto_collector_logs_v1_logs_service_proto_rawDescData)
	})
	return file_opentelemetry_proto_collector_logs_v1_logs_service_proto_rawDescData
}

var file_opentelemetry_proto_collector_logs_v1_logs_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_opentelemetry_proto_collector_logs_v1_logs_service_proto_goTypes = []interface{}{
	(*ExportLogsServiceRequest)(nil),  // 0: internal.opentelemetry.proto.collector.logs.v1.ExportLogsServiceRequest
	(*ExportLogsServiceResponse)(nil), // 1: internal.opentelemetry.proto.collector.logs.v1.ExportLogsServiceResponse
	(*v1.ResourceLogs)(nil),           // 2: internal.opentelemetry.proto.logs.v1.ResourceLogs
}
var file_opentelemetry_proto_collector_logs_v1_logs_service_proto_depIdxs = []int32{
	2, // 0: internal.opentelemetry.proto.collector.logs.v1.ExportLogsServiceRequest.resource_logs:type_name -> internal.opentelemetry.proto.logs.v1.ResourceLogs
	0, // 1: internal.opentelemetry.proto.collector.logs.v1.LogsService.Export:input_type -> internal.opentelemetry.proto.collector.logs.v1.ExportLogsServiceRequest
	1, // 2: internal.opentelemetry.proto.collector.logs.v1.LogsService.Export:output_type -> internal.opentelemetry.proto.collector.logs.v1.ExportLogsServiceResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_opentelemetry_proto_collector_logs_v1_logs_service_proto_init() }
func file_opentelemetry_proto_collector_logs_v1_logs_service_proto_init() {
	if File_opentelemetry_proto_collector_logs_v1_logs_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_opentelemetry_proto_collector_logs_v1_logs_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportLogsServiceRequest); i {
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
		file_opentelemetry_proto_collector_logs_v1_logs_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportLogsServiceResponse); i {
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
			RawDescriptor: file_opentelemetry_proto_collector_logs_v1_logs_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_opentelemetry_proto_collector_logs_v1_logs_service_proto_goTypes,
		DependencyIndexes: file_opentelemetry_proto_collector_logs_v1_logs_service_proto_depIdxs,
		MessageInfos:      file_opentelemetry_proto_collector_logs_v1_logs_service_proto_msgTypes,
	}.Build()
	File_opentelemetry_proto_collector_logs_v1_logs_service_proto = out.File
	file_opentelemetry_proto_collector_logs_v1_logs_service_proto_rawDesc = nil
	file_opentelemetry_proto_collector_logs_v1_logs_service_proto_goTypes = nil
	file_opentelemetry_proto_collector_logs_v1_logs_service_proto_depIdxs = nil
}