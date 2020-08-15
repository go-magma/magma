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
// 	protoc-gen-go v1.23.0
// 	protoc        (unknown)
// source: sync_rpc_service.proto

package protos

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GatewayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GwId string `protobuf:"bytes,1,opt,name=gwId,proto3" json:"gwId,omitempty"`
	// sync_rpc_service leverages the fact that grpc is built on top of http/2.
	// As pseudo header :method will always be POST, and :scheme will always be http,
	// they are excluded from the message definition. grpc uses pseudo header :authority
	// and :path to route the rpc calls to the corresponding services
	Authority string `protobuf:"bytes,2,opt,name=authority,proto3" json:"authority,omitempty"`
	Path      string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	// non-pseudo headers
	Headers map[string]string `protobuf:"bytes,5,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Payload []byte            `protobuf:"bytes,6,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *GatewayRequest) Reset() {
	*x = GatewayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_rpc_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayRequest) ProtoMessage() {}

func (x *GatewayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sync_rpc_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayRequest.ProtoReflect.Descriptor instead.
func (*GatewayRequest) Descriptor() ([]byte, []int) {
	return file_sync_rpc_service_proto_rawDescGZIP(), []int{0}
}

func (x *GatewayRequest) GetGwId() string {
	if x != nil {
		return x.GwId
	}
	return ""
}

func (x *GatewayRequest) GetAuthority() string {
	if x != nil {
		return x.Authority
	}
	return ""
}

func (x *GatewayRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *GatewayRequest) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *GatewayRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type GatewayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pseudo header :status
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// non-pseudo headers
	Headers map[string]string `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// response body
	Payload []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	// err message to return to the caller, if any
	Err string `protobuf:"bytes,4,opt,name=err,proto3" json:"err,omitempty"`
	// keepConnActive is used to tell the client that the connection is still active, even if no response has
	// been received for a while
	KeepConnActive bool `protobuf:"varint,5,opt,name=keepConnActive,proto3" json:"keepConnActive,omitempty"`
}

func (x *GatewayResponse) Reset() {
	*x = GatewayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_rpc_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayResponse) ProtoMessage() {}

func (x *GatewayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sync_rpc_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayResponse.ProtoReflect.Descriptor instead.
func (*GatewayResponse) Descriptor() ([]byte, []int) {
	return file_sync_rpc_service_proto_rawDescGZIP(), []int{1}
}

func (x *GatewayResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GatewayResponse) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *GatewayResponse) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *GatewayResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

func (x *GatewayResponse) GetKeepConnActive() bool {
	if x != nil {
		return x.KeepConnActive
	}
	return false
}

// SyncRPCRequest is sent down to gateway from cloud
type SyncRPCRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// unique request Id passed in from cloud
	ReqId   uint32          `protobuf:"varint,1,opt,name=reqId,proto3" json:"reqId,omitempty"`
	ReqBody *GatewayRequest `protobuf:"bytes,2,opt,name=reqBody,proto3" json:"reqBody,omitempty"`
	// cloud will send a heartBeat every minute, if no other requests are sent
	// down to the gateway
	HeartBeat bool `protobuf:"varint,3,opt,name=heartBeat,proto3" json:"heartBeat,omitempty"`
	// connClosed is set to true when the client closes the connection
	ConnClosed bool `protobuf:"varint,4,opt,name=connClosed,proto3" json:"connClosed,omitempty"`
}

func (x *SyncRPCRequest) Reset() {
	*x = SyncRPCRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_rpc_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncRPCRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncRPCRequest) ProtoMessage() {}

func (x *SyncRPCRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sync_rpc_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncRPCRequest.ProtoReflect.Descriptor instead.
func (*SyncRPCRequest) Descriptor() ([]byte, []int) {
	return file_sync_rpc_service_proto_rawDescGZIP(), []int{2}
}

func (x *SyncRPCRequest) GetReqId() uint32 {
	if x != nil {
		return x.ReqId
	}
	return 0
}

func (x *SyncRPCRequest) GetReqBody() *GatewayRequest {
	if x != nil {
		return x.ReqBody
	}
	return nil
}

func (x *SyncRPCRequest) GetHeartBeat() bool {
	if x != nil {
		return x.HeartBeat
	}
	return false
}

func (x *SyncRPCRequest) GetConnClosed() bool {
	if x != nil {
		return x.ConnClosed
	}
	return false
}

// SyncRPCResponse is sent from gateway to cloud
type SyncRPCResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqId    uint32           `protobuf:"varint,1,opt,name=reqId,proto3" json:"reqId,omitempty"`
	RespBody *GatewayResponse `protobuf:"bytes,2,opt,name=respBody,proto3" json:"respBody,omitempty"`
	// gateway will send a heartBeat if it hasn't received SyncRPCRequests from cloud for a while.
	// If it's a heartbeat, reqId and respBody will be ignored.
	HeartBeat bool `protobuf:"varint,3,opt,name=heartBeat,proto3" json:"heartBeat,omitempty"`
}

func (x *SyncRPCResponse) Reset() {
	*x = SyncRPCResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_rpc_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncRPCResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncRPCResponse) ProtoMessage() {}

func (x *SyncRPCResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sync_rpc_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncRPCResponse.ProtoReflect.Descriptor instead.
func (*SyncRPCResponse) Descriptor() ([]byte, []int) {
	return file_sync_rpc_service_proto_rawDescGZIP(), []int{3}
}

func (x *SyncRPCResponse) GetReqId() uint32 {
	if x != nil {
		return x.ReqId
	}
	return 0
}

func (x *SyncRPCResponse) GetRespBody() *GatewayResponse {
	if x != nil {
		return x.RespBody
	}
	return nil
}

func (x *SyncRPCResponse) GetHeartBeat() bool {
	if x != nil {
		return x.HeartBeat
	}
	return false
}

var File_sync_rpc_service_proto protoreflect.FileDescriptor

var file_sync_rpc_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e,
	0x6f, 0x72, 0x63, 0x38, 0x72, 0x22, 0xf0, 0x01, 0x0a, 0x0e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x77, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x77, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x42,
	0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x3a, 0x0a, 0x0c,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xfe, 0x01, 0x0a, 0x0f, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x43, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72,
	0x63, 0x38, 0x72, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x65, 0x72, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x6b, 0x65, 0x65, 0x70, 0x43, 0x6f, 0x6e,
	0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x6b,
	0x65, 0x65, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x1a, 0x3a, 0x0a,
	0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9b, 0x01, 0x0a, 0x0e, 0x53, 0x79,
	0x6e, 0x63, 0x52, 0x50, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x72, 0x65, 0x71, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x72, 0x65, 0x71,
	0x49, 0x64, 0x12, 0x35, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38,
	0x72, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x07, 0x72, 0x65, 0x71, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x65, 0x61,
	0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x68, 0x65,
	0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x43,
	0x6c, 0x6f, 0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x63, 0x6f, 0x6e,
	0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x22, 0x7f, 0x0a, 0x0f, 0x53, 0x79, 0x6e, 0x63, 0x52,
	0x50, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65,
	0x71, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64,
	0x12, 0x38, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72,
	0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x65,
	0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x68,
	0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x32, 0xb7, 0x01, 0x0a, 0x0e, 0x53, 0x79, 0x6e,
	0x63, 0x52, 0x50, 0x43, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x16, 0x45,
	0x73, 0x74, 0x61, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x50, 0x43, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1c, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72,
	0x63, 0x38, 0x72, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x50, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x1a, 0x1b, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38,
	0x72, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x50, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x4a, 0x0a, 0x07, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x50,
	0x43, 0x12, 0x1c, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e,
	0x53, 0x79, 0x6e, 0x63, 0x52, 0x50, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a,
	0x1b, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x53, 0x79,
	0x6e, 0x63, 0x52, 0x50, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f,
	0x6c, 0x69, 0x62, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sync_rpc_service_proto_rawDescOnce sync.Once
	file_sync_rpc_service_proto_rawDescData = file_sync_rpc_service_proto_rawDesc
)

func file_sync_rpc_service_proto_rawDescGZIP() []byte {
	file_sync_rpc_service_proto_rawDescOnce.Do(func() {
		file_sync_rpc_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_sync_rpc_service_proto_rawDescData)
	})
	return file_sync_rpc_service_proto_rawDescData
}

var file_sync_rpc_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sync_rpc_service_proto_goTypes = []interface{}{
	(*GatewayRequest)(nil),  // 0: magma.orc8r.GatewayRequest
	(*GatewayResponse)(nil), // 1: magma.orc8r.GatewayResponse
	(*SyncRPCRequest)(nil),  // 2: magma.orc8r.SyncRPCRequest
	(*SyncRPCResponse)(nil), // 3: magma.orc8r.SyncRPCResponse
	nil,                     // 4: magma.orc8r.GatewayRequest.HeadersEntry
	nil,                     // 5: magma.orc8r.GatewayResponse.HeadersEntry
}
var file_sync_rpc_service_proto_depIdxs = []int32{
	4, // 0: magma.orc8r.GatewayRequest.headers:type_name -> magma.orc8r.GatewayRequest.HeadersEntry
	5, // 1: magma.orc8r.GatewayResponse.headers:type_name -> magma.orc8r.GatewayResponse.HeadersEntry
	0, // 2: magma.orc8r.SyncRPCRequest.reqBody:type_name -> magma.orc8r.GatewayRequest
	1, // 3: magma.orc8r.SyncRPCResponse.respBody:type_name -> magma.orc8r.GatewayResponse
	3, // 4: magma.orc8r.SyncRPCService.EstablishSyncRPCStream:input_type -> magma.orc8r.SyncRPCResponse
	3, // 5: magma.orc8r.SyncRPCService.SyncRPC:input_type -> magma.orc8r.SyncRPCResponse
	2, // 6: magma.orc8r.SyncRPCService.EstablishSyncRPCStream:output_type -> magma.orc8r.SyncRPCRequest
	2, // 7: magma.orc8r.SyncRPCService.SyncRPC:output_type -> magma.orc8r.SyncRPCRequest
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_sync_rpc_service_proto_init() }
func file_sync_rpc_service_proto_init() {
	if File_sync_rpc_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sync_rpc_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayRequest); i {
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
		file_sync_rpc_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayResponse); i {
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
		file_sync_rpc_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncRPCRequest); i {
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
		file_sync_rpc_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncRPCResponse); i {
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
			RawDescriptor: file_sync_rpc_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sync_rpc_service_proto_goTypes,
		DependencyIndexes: file_sync_rpc_service_proto_depIdxs,
		MessageInfos:      file_sync_rpc_service_proto_msgTypes,
	}.Build()
	File_sync_rpc_service_proto = out.File
	file_sync_rpc_service_proto_rawDesc = nil
	file_sync_rpc_service_proto_goTypes = nil
	file_sync_rpc_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SyncRPCServiceClient is the client API for SyncRPCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SyncRPCServiceClient interface {
	// creates a bidirectional stream from gateway to cloud
	// so cloud can send in SyncRPCRequest, and wait for SyncRPCResponse.
	// This will be the underlying service for Synchronous RPC from the cloud.
	EstablishSyncRPCStream(ctx context.Context, opts ...grpc.CallOption) (SyncRPCService_EstablishSyncRPCStreamClient, error)
	// same method as EstablishSyncRPCStream, but named differently for backward compatibility
	SyncRPC(ctx context.Context, opts ...grpc.CallOption) (SyncRPCService_SyncRPCClient, error)
}

type syncRPCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSyncRPCServiceClient(cc grpc.ClientConnInterface) SyncRPCServiceClient {
	return &syncRPCServiceClient{cc}
}

func (c *syncRPCServiceClient) EstablishSyncRPCStream(ctx context.Context, opts ...grpc.CallOption) (SyncRPCService_EstablishSyncRPCStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SyncRPCService_serviceDesc.Streams[0], "/magma.orc8r.SyncRPCService/EstablishSyncRPCStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &syncRPCServiceEstablishSyncRPCStreamClient{stream}
	return x, nil
}

type SyncRPCService_EstablishSyncRPCStreamClient interface {
	Send(*SyncRPCResponse) error
	Recv() (*SyncRPCRequest, error)
	grpc.ClientStream
}

type syncRPCServiceEstablishSyncRPCStreamClient struct {
	grpc.ClientStream
}

func (x *syncRPCServiceEstablishSyncRPCStreamClient) Send(m *SyncRPCResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *syncRPCServiceEstablishSyncRPCStreamClient) Recv() (*SyncRPCRequest, error) {
	m := new(SyncRPCRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *syncRPCServiceClient) SyncRPC(ctx context.Context, opts ...grpc.CallOption) (SyncRPCService_SyncRPCClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SyncRPCService_serviceDesc.Streams[1], "/magma.orc8r.SyncRPCService/SyncRPC", opts...)
	if err != nil {
		return nil, err
	}
	x := &syncRPCServiceSyncRPCClient{stream}
	return x, nil
}

type SyncRPCService_SyncRPCClient interface {
	Send(*SyncRPCResponse) error
	Recv() (*SyncRPCRequest, error)
	grpc.ClientStream
}

type syncRPCServiceSyncRPCClient struct {
	grpc.ClientStream
}

func (x *syncRPCServiceSyncRPCClient) Send(m *SyncRPCResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *syncRPCServiceSyncRPCClient) Recv() (*SyncRPCRequest, error) {
	m := new(SyncRPCRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SyncRPCServiceServer is the server API for SyncRPCService service.
type SyncRPCServiceServer interface {
	// creates a bidirectional stream from gateway to cloud
	// so cloud can send in SyncRPCRequest, and wait for SyncRPCResponse.
	// This will be the underlying service for Synchronous RPC from the cloud.
	EstablishSyncRPCStream(SyncRPCService_EstablishSyncRPCStreamServer) error
	// same method as EstablishSyncRPCStream, but named differently for backward compatibility
	SyncRPC(SyncRPCService_SyncRPCServer) error
}

// UnimplementedSyncRPCServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSyncRPCServiceServer struct {
}

func (*UnimplementedSyncRPCServiceServer) EstablishSyncRPCStream(SyncRPCService_EstablishSyncRPCStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method EstablishSyncRPCStream not implemented")
}
func (*UnimplementedSyncRPCServiceServer) SyncRPC(SyncRPCService_SyncRPCServer) error {
	return status.Errorf(codes.Unimplemented, "method SyncRPC not implemented")
}

func RegisterSyncRPCServiceServer(s *grpc.Server, srv SyncRPCServiceServer) {
	s.RegisterService(&_SyncRPCService_serviceDesc, srv)
}

func _SyncRPCService_EstablishSyncRPCStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SyncRPCServiceServer).EstablishSyncRPCStream(&syncRPCServiceEstablishSyncRPCStreamServer{stream})
}

type SyncRPCService_EstablishSyncRPCStreamServer interface {
	Send(*SyncRPCRequest) error
	Recv() (*SyncRPCResponse, error)
	grpc.ServerStream
}

type syncRPCServiceEstablishSyncRPCStreamServer struct {
	grpc.ServerStream
}

func (x *syncRPCServiceEstablishSyncRPCStreamServer) Send(m *SyncRPCRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *syncRPCServiceEstablishSyncRPCStreamServer) Recv() (*SyncRPCResponse, error) {
	m := new(SyncRPCResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SyncRPCService_SyncRPC_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SyncRPCServiceServer).SyncRPC(&syncRPCServiceSyncRPCServer{stream})
}

type SyncRPCService_SyncRPCServer interface {
	Send(*SyncRPCRequest) error
	Recv() (*SyncRPCResponse, error)
	grpc.ServerStream
}

type syncRPCServiceSyncRPCServer struct {
	grpc.ServerStream
}

func (x *syncRPCServiceSyncRPCServer) Send(m *SyncRPCRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *syncRPCServiceSyncRPCServer) Recv() (*SyncRPCResponse, error) {
	m := new(SyncRPCResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SyncRPCService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.SyncRPCService",
	HandlerType: (*SyncRPCServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EstablishSyncRPCStream",
			Handler:       _SyncRPCService_EstablishSyncRPCStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SyncRPC",
			Handler:       _SyncRPCService_SyncRPC_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "sync_rpc_service.proto",
}
