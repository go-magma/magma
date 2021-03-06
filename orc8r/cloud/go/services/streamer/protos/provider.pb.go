//
//Copyright 2020 The Magma Authors.
//
//This source code is licensed under the BSD-style license found in the
//LICENSE file in the root directory of this source tree.
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: provider.proto

package protos

import (
	context "context"
	protos "github.com/go-magma/magma/lib/go/protos"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
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

var File_provider_proto protoreflect.FileDescriptor

var file_provider_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x1a, 0x0e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x5a, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f,
	0x72, 0x63, 0x38, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x22, 0x00, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f,
	0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_provider_proto_goTypes = []interface{}{
	(*protos.StreamRequest)(nil),   // 0: magma.orc8r.StreamRequest
	(*protos.DataUpdateBatch)(nil), // 1: magma.orc8r.DataUpdateBatch
}
var file_provider_proto_depIdxs = []int32{
	0, // 0: magma.orc8r.streamer.StreamProvider.GetUpdates:input_type -> magma.orc8r.StreamRequest
	1, // 1: magma.orc8r.streamer.StreamProvider.GetUpdates:output_type -> magma.orc8r.DataUpdateBatch
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_provider_proto_init() }
func file_provider_proto_init() {
	if File_provider_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_provider_proto_goTypes,
		DependencyIndexes: file_provider_proto_depIdxs,
	}.Build()
	File_provider_proto = out.File
	file_provider_proto_rawDesc = nil
	file_provider_proto_goTypes = nil
	file_provider_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StreamProviderClient is the client API for StreamProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamProviderClient interface {
	GetUpdates(ctx context.Context, in *protos.StreamRequest, opts ...grpc.CallOption) (*protos.DataUpdateBatch, error)
}

type streamProviderClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamProviderClient(cc grpc.ClientConnInterface) StreamProviderClient {
	return &streamProviderClient{cc}
}

func (c *streamProviderClient) GetUpdates(ctx context.Context, in *protos.StreamRequest, opts ...grpc.CallOption) (*protos.DataUpdateBatch, error) {
	out := new(protos.DataUpdateBatch)
	err := c.cc.Invoke(ctx, "/magma.orc8r.streamer.StreamProvider/GetUpdates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamProviderServer is the server API for StreamProvider service.
type StreamProviderServer interface {
	GetUpdates(context.Context, *protos.StreamRequest) (*protos.DataUpdateBatch, error)
}

// UnimplementedStreamProviderServer can be embedded to have forward compatible implementations.
type UnimplementedStreamProviderServer struct {
}

func (*UnimplementedStreamProviderServer) GetUpdates(context.Context, *protos.StreamRequest) (*protos.DataUpdateBatch, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUpdates not implemented")
}

func RegisterStreamProviderServer(s *grpc.Server, srv StreamProviderServer) {
	s.RegisterService(&_StreamProvider_serviceDesc, srv)
}

func _StreamProvider_GetUpdates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.StreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamProviderServer).GetUpdates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.streamer.StreamProvider/GetUpdates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamProviderServer).GetUpdates(ctx, req.(*protos.StreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StreamProvider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.streamer.StreamProvider",
	HandlerType: (*StreamProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUpdates",
			Handler:    _StreamProvider_GetUpdates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provider.proto",
}
