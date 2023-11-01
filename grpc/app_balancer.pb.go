// Copyright 2015 The gRPC Authors
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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: grpc/app_balancer.proto

package grpc

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

// Points are represented as latitude-longitude pairs in the E7 representation
// (degrees multiplied by 10**7 and rounded to the nearest integer).
// Latitudes should be in the range +/- 90 degrees and longitude should be in
// the range +/- 180 degrees (inclusive).
type UpdatePodSpeed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Speed float64 `protobuf:"fixed64,2,opt,name=speed,proto3" json:"speed,omitempty"`
}

func (x *UpdatePodSpeed) Reset() {
	*x = UpdatePodSpeed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_app_balancer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePodSpeed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePodSpeed) ProtoMessage() {}

func (x *UpdatePodSpeed) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_app_balancer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePodSpeed.ProtoReflect.Descriptor instead.
func (*UpdatePodSpeed) Descriptor() ([]byte, []int) {
	return file_grpc_app_balancer_proto_rawDescGZIP(), []int{0}
}

func (x *UpdatePodSpeed) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePodSpeed) GetSpeed() float64 {
	if x != nil {
		return x.Speed
	}
	return 0
}

type CreatePodApp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreatePodApp) Reset() {
	*x = CreatePodApp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_app_balancer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePodApp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePodApp) ProtoMessage() {}

func (x *CreatePodApp) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_app_balancer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePodApp.ProtoReflect.Descriptor instead.
func (*CreatePodApp) Descriptor() ([]byte, []int) {
	return file_grpc_app_balancer_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePodApp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_app_balancer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_app_balancer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_grpc_app_balancer_proto_rawDescGZIP(), []int{2}
}

func (x *Result) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_grpc_app_balancer_proto protoreflect.FileDescriptor

var file_grpc_app_balancer_proto_rawDesc = []byte{
	0x0a, 0x17, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x70, 0x70, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x70, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x22, 0x36, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x64, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x22, 0x1e,
	0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x64, 0x41, 0x70, 0x70, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x22,
	0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0x8f, 0x01, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x72, 0x12, 0x41, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65, 0x65,
	0x64, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x70, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x64, 0x53, 0x70, 0x65, 0x65, 0x64, 0x1a, 0x13,
	0x2e, 0x61, 0x70, 0x70, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x70, 0x70, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x70, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x64, 0x41, 0x70, 0x70, 0x1a, 0x13, 0x2e,
	0x61, 0x70, 0x70, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x00, 0x42, 0x4e, 0x0a, 0x10, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x70, 0x70,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x42, 0x10, 0x41, 0x70, 0x70, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x6d, 0x63, 0x69, 0x6b, 0x6a,
	0x61, 0x6e, 0x2f, 0x61, 0x70, 0x70, 0x2d, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_app_balancer_proto_rawDescOnce sync.Once
	file_grpc_app_balancer_proto_rawDescData = file_grpc_app_balancer_proto_rawDesc
)

func file_grpc_app_balancer_proto_rawDescGZIP() []byte {
	file_grpc_app_balancer_proto_rawDescOnce.Do(func() {
		file_grpc_app_balancer_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_app_balancer_proto_rawDescData)
	})
	return file_grpc_app_balancer_proto_rawDescData
}

var file_grpc_app_balancer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_grpc_app_balancer_proto_goTypes = []interface{}{
	(*UpdatePodSpeed)(nil), // 0: appbalancer.UpdatePodSpeed
	(*CreatePodApp)(nil),   // 1: appbalancer.CreatePodApp
	(*Result)(nil),         // 2: appbalancer.Result
}
var file_grpc_app_balancer_proto_depIdxs = []int32{
	0, // 0: appbalancer.AppBalancer.UpdateSpeed:input_type -> appbalancer.UpdatePodSpeed
	1, // 1: appbalancer.AppBalancer.CreateApp:input_type -> appbalancer.CreatePodApp
	2, // 2: appbalancer.AppBalancer.UpdateSpeed:output_type -> appbalancer.Result
	2, // 3: appbalancer.AppBalancer.CreateApp:output_type -> appbalancer.Result
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_app_balancer_proto_init() }
func file_grpc_app_balancer_proto_init() {
	if File_grpc_app_balancer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_app_balancer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePodSpeed); i {
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
		file_grpc_app_balancer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePodApp); i {
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
		file_grpc_app_balancer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
			RawDescriptor: file_grpc_app_balancer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_app_balancer_proto_goTypes,
		DependencyIndexes: file_grpc_app_balancer_proto_depIdxs,
		MessageInfos:      file_grpc_app_balancer_proto_msgTypes,
	}.Build()
	File_grpc_app_balancer_proto = out.File
	file_grpc_app_balancer_proto_rawDesc = nil
	file_grpc_app_balancer_proto_goTypes = nil
	file_grpc_app_balancer_proto_depIdxs = nil
}