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
// source: hss_service.proto

package protos

import (
	context "context"
	protos1 "github.com/go-magma/magma/lib/go/protos"
	protos "github.com/go-magma/magma/modules/lte/cloud/go/protos"
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

var File_hss_service_proto protoreflect.FileDescriptor

var file_hss_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x68, 0x73, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x66, 0x65, 0x67, 0x1a, 0x0c,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xe9, 0x02, 0x0a, 0x0f, 0x48, 0x53, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x3f, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74,
	0x65, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x1a, 0x11, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x56,
	0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x6d, 0x61, 0x67, 0x6d,
	0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x49, 0x44, 0x1a, 0x11, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x6d, 0x61,
	0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x11, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f,
	0x72, 0x63, 0x38, 0x72, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x17, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x19, 0x2e, 0x6d, 0x61, 0x67, 0x6d,
	0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x14, 0x44, 0x65, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x12, 0x17,
	0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x11, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e,
	0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x61,
	0x67, 0x6d, 0x61, 0x2f, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x73, 0x2f, 0x66, 0x65, 0x67, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_hss_service_proto_goTypes = []interface{}{
	(*protos.SubscriberData)(nil), // 0: magma.lte.SubscriberData
	(*protos.SubscriberID)(nil),   // 1: magma.lte.SubscriberID
	(*protos1.Void)(nil),          // 2: magma.orc8r.Void
}
var file_hss_service_proto_depIdxs = []int32{
	0, // 0: magma.feg.HSSConfigurator.AddSubscriber:input_type -> magma.lte.SubscriberData
	1, // 1: magma.feg.HSSConfigurator.DeleteSubscriber:input_type -> magma.lte.SubscriberID
	0, // 2: magma.feg.HSSConfigurator.UpdateSubscriber:input_type -> magma.lte.SubscriberData
	1, // 3: magma.feg.HSSConfigurator.GetSubscriberData:input_type -> magma.lte.SubscriberID
	1, // 4: magma.feg.HSSConfigurator.DeregisterSubscriber:input_type -> magma.lte.SubscriberID
	2, // 5: magma.feg.HSSConfigurator.AddSubscriber:output_type -> magma.orc8r.Void
	2, // 6: magma.feg.HSSConfigurator.DeleteSubscriber:output_type -> magma.orc8r.Void
	2, // 7: magma.feg.HSSConfigurator.UpdateSubscriber:output_type -> magma.orc8r.Void
	0, // 8: magma.feg.HSSConfigurator.GetSubscriberData:output_type -> magma.lte.SubscriberData
	2, // 9: magma.feg.HSSConfigurator.DeregisterSubscriber:output_type -> magma.orc8r.Void
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hss_service_proto_init() }
func file_hss_service_proto_init() {
	if File_hss_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hss_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hss_service_proto_goTypes,
		DependencyIndexes: file_hss_service_proto_depIdxs,
	}.Build()
	File_hss_service_proto = out.File
	file_hss_service_proto_rawDesc = nil
	file_hss_service_proto_goTypes = nil
	file_hss_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HSSConfiguratorClient is the client API for HSSConfigurator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HSSConfiguratorClient interface {
	// Adds a new subscriber to the store.
	// Throws ALREADY_EXISTS if the subscriber already exists.
	//
	AddSubscriber(ctx context.Context, in *protos.SubscriberData, opts ...grpc.CallOption) (*protos1.Void, error)
	// Deletes an existing subscriber.
	// If the subscriber is not already present, this request is ignored.
	//
	DeleteSubscriber(ctx context.Context, in *protos.SubscriberID, opts ...grpc.CallOption) (*protos1.Void, error)
	// Updates an existing subscriber.
	// Throws NOT_FOUND if the subscriber is missing.
	//
	UpdateSubscriber(ctx context.Context, in *protos.SubscriberData, opts ...grpc.CallOption) (*protos1.Void, error)
	// Returns the SubscriberData for a subscriber.
	// Throws NOT_FOUND if the subscriber is missing.
	//
	GetSubscriberData(ctx context.Context, in *protos.SubscriberID, opts ...grpc.CallOption) (*protos.SubscriberData, error)
	// De-register an authenticated subscriber
	DeregisterSubscriber(ctx context.Context, in *protos.SubscriberID, opts ...grpc.CallOption) (*protos1.Void, error)
}

type hSSConfiguratorClient struct {
	cc grpc.ClientConnInterface
}

func NewHSSConfiguratorClient(cc grpc.ClientConnInterface) HSSConfiguratorClient {
	return &hSSConfiguratorClient{cc}
}

func (c *hSSConfiguratorClient) AddSubscriber(ctx context.Context, in *protos.SubscriberData, opts ...grpc.CallOption) (*protos1.Void, error) {
	out := new(protos1.Void)
	err := c.cc.Invoke(ctx, "/magma.feg.HSSConfigurator/AddSubscriber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hSSConfiguratorClient) DeleteSubscriber(ctx context.Context, in *protos.SubscriberID, opts ...grpc.CallOption) (*protos1.Void, error) {
	out := new(protos1.Void)
	err := c.cc.Invoke(ctx, "/magma.feg.HSSConfigurator/DeleteSubscriber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hSSConfiguratorClient) UpdateSubscriber(ctx context.Context, in *protos.SubscriberData, opts ...grpc.CallOption) (*protos1.Void, error) {
	out := new(protos1.Void)
	err := c.cc.Invoke(ctx, "/magma.feg.HSSConfigurator/UpdateSubscriber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hSSConfiguratorClient) GetSubscriberData(ctx context.Context, in *protos.SubscriberID, opts ...grpc.CallOption) (*protos.SubscriberData, error) {
	out := new(protos.SubscriberData)
	err := c.cc.Invoke(ctx, "/magma.feg.HSSConfigurator/GetSubscriberData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hSSConfiguratorClient) DeregisterSubscriber(ctx context.Context, in *protos.SubscriberID, opts ...grpc.CallOption) (*protos1.Void, error) {
	out := new(protos1.Void)
	err := c.cc.Invoke(ctx, "/magma.feg.HSSConfigurator/DeregisterSubscriber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HSSConfiguratorServer is the server API for HSSConfigurator service.
type HSSConfiguratorServer interface {
	// Adds a new subscriber to the store.
	// Throws ALREADY_EXISTS if the subscriber already exists.
	//
	AddSubscriber(context.Context, *protos.SubscriberData) (*protos1.Void, error)
	// Deletes an existing subscriber.
	// If the subscriber is not already present, this request is ignored.
	//
	DeleteSubscriber(context.Context, *protos.SubscriberID) (*protos1.Void, error)
	// Updates an existing subscriber.
	// Throws NOT_FOUND if the subscriber is missing.
	//
	UpdateSubscriber(context.Context, *protos.SubscriberData) (*protos1.Void, error)
	// Returns the SubscriberData for a subscriber.
	// Throws NOT_FOUND if the subscriber is missing.
	//
	GetSubscriberData(context.Context, *protos.SubscriberID) (*protos.SubscriberData, error)
	// De-register an authenticated subscriber
	DeregisterSubscriber(context.Context, *protos.SubscriberID) (*protos1.Void, error)
}

// UnimplementedHSSConfiguratorServer can be embedded to have forward compatible implementations.
type UnimplementedHSSConfiguratorServer struct {
}

func (*UnimplementedHSSConfiguratorServer) AddSubscriber(context.Context, *protos.SubscriberData) (*protos1.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSubscriber not implemented")
}
func (*UnimplementedHSSConfiguratorServer) DeleteSubscriber(context.Context, *protos.SubscriberID) (*protos1.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubscriber not implemented")
}
func (*UnimplementedHSSConfiguratorServer) UpdateSubscriber(context.Context, *protos.SubscriberData) (*protos1.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSubscriber not implemented")
}
func (*UnimplementedHSSConfiguratorServer) GetSubscriberData(context.Context, *protos.SubscriberID) (*protos.SubscriberData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriberData not implemented")
}
func (*UnimplementedHSSConfiguratorServer) DeregisterSubscriber(context.Context, *protos.SubscriberID) (*protos1.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeregisterSubscriber not implemented")
}

func RegisterHSSConfiguratorServer(s *grpc.Server, srv HSSConfiguratorServer) {
	s.RegisterService(&_HSSConfigurator_serviceDesc, srv)
}

func _HSSConfigurator_AddSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.SubscriberData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HSSConfiguratorServer).AddSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.HSSConfigurator/AddSubscriber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HSSConfiguratorServer).AddSubscriber(ctx, req.(*protos.SubscriberData))
	}
	return interceptor(ctx, in, info, handler)
}

func _HSSConfigurator_DeleteSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.SubscriberID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HSSConfiguratorServer).DeleteSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.HSSConfigurator/DeleteSubscriber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HSSConfiguratorServer).DeleteSubscriber(ctx, req.(*protos.SubscriberID))
	}
	return interceptor(ctx, in, info, handler)
}

func _HSSConfigurator_UpdateSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.SubscriberData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HSSConfiguratorServer).UpdateSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.HSSConfigurator/UpdateSubscriber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HSSConfiguratorServer).UpdateSubscriber(ctx, req.(*protos.SubscriberData))
	}
	return interceptor(ctx, in, info, handler)
}

func _HSSConfigurator_GetSubscriberData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.SubscriberID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HSSConfiguratorServer).GetSubscriberData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.HSSConfigurator/GetSubscriberData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HSSConfiguratorServer).GetSubscriberData(ctx, req.(*protos.SubscriberID))
	}
	return interceptor(ctx, in, info, handler)
}

func _HSSConfigurator_DeregisterSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.SubscriberID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HSSConfiguratorServer).DeregisterSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.HSSConfigurator/DeregisterSubscriber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HSSConfiguratorServer).DeregisterSubscriber(ctx, req.(*protos.SubscriberID))
	}
	return interceptor(ctx, in, info, handler)
}

var _HSSConfigurator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.feg.HSSConfigurator",
	HandlerType: (*HSSConfiguratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddSubscriber",
			Handler:    _HSSConfigurator_AddSubscriber_Handler,
		},
		{
			MethodName: "DeleteSubscriber",
			Handler:    _HSSConfigurator_DeleteSubscriber_Handler,
		},
		{
			MethodName: "UpdateSubscriber",
			Handler:    _HSSConfigurator_UpdateSubscriber_Handler,
		},
		{
			MethodName: "GetSubscriberData",
			Handler:    _HSSConfigurator_GetSubscriberData_Handler,
		},
		{
			MethodName: "DeregisterSubscriber",
			Handler:    _HSSConfigurator_DeregisterSubscriber_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hss_service.proto",
}
