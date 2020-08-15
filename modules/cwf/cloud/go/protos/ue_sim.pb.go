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
// source: ue_sim.proto

package protos

import (
	context "context"
	protos "github.com/go-magma/magma/lib/go/protos"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// --------------------------------------------------------------------------
// UE config
// --------------------------------------------------------------------------
type UEConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier for the UE.
	Imsi string `protobuf:"bytes,1,opt,name=imsi,proto3" json:"imsi,omitempty"`
	// Authentication key (k).
	AuthKey []byte `protobuf:"bytes,2,opt,name=auth_key,json=authKey,proto3" json:"auth_key,omitempty"`
	// Operator configuration field (Op) signed with authentication key (k).
	AuthOpc []byte `protobuf:"bytes,3,opt,name=auth_opc,json=authOpc,proto3" json:"auth_opc,omitempty"`
	// Sequence Number (SEQ).
	Seq uint64 `protobuf:"varint,4,opt,name=seq,proto3" json:"seq,omitempty"`
}

func (x *UEConfig) Reset() {
	*x = UEConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ue_sim_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UEConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UEConfig) ProtoMessage() {}

func (x *UEConfig) ProtoReflect() protoreflect.Message {
	mi := &file_ue_sim_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UEConfig.ProtoReflect.Descriptor instead.
func (*UEConfig) Descriptor() ([]byte, []int) {
	return file_ue_sim_proto_rawDescGZIP(), []int{0}
}

func (x *UEConfig) GetImsi() string {
	if x != nil {
		return x.Imsi
	}
	return ""
}

func (x *UEConfig) GetAuthKey() []byte {
	if x != nil {
		return x.AuthKey
	}
	return nil
}

func (x *UEConfig) GetAuthOpc() []byte {
	if x != nil {
		return x.AuthOpc
	}
	return nil
}

func (x *UEConfig) GetSeq() uint64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

type AuthenticateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Imsi            string `protobuf:"bytes,1,opt,name=imsi,proto3" json:"imsi,omitempty"`
	CalledStationID string `protobuf:"bytes,2,opt,name=calledStationID,proto3" json:"calledStationID,omitempty"`
}

func (x *AuthenticateRequest) Reset() {
	*x = AuthenticateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ue_sim_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateRequest) ProtoMessage() {}

func (x *AuthenticateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ue_sim_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateRequest) Descriptor() ([]byte, []int) {
	return file_ue_sim_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticateRequest) GetImsi() string {
	if x != nil {
		return x.Imsi
	}
	return ""
}

func (x *AuthenticateRequest) GetCalledStationID() string {
	if x != nil {
		return x.CalledStationID
	}
	return ""
}

type AuthenticateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RadiusPacket []byte `protobuf:"bytes,1,opt,name=radiusPacket,proto3" json:"radiusPacket,omitempty"`
}

func (x *AuthenticateResponse) Reset() {
	*x = AuthenticateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ue_sim_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateResponse) ProtoMessage() {}

func (x *AuthenticateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ue_sim_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateResponse.ProtoReflect.Descriptor instead.
func (*AuthenticateResponse) Descriptor() ([]byte, []int) {
	return file_ue_sim_proto_rawDescGZIP(), []int{2}
}

func (x *AuthenticateResponse) GetRadiusPacket() []byte {
	if x != nil {
		return x.RadiusPacket
	}
	return nil
}

type DisconnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Imsi            string `protobuf:"bytes,1,opt,name=imsi,proto3" json:"imsi,omitempty"`
	CalledStationID string `protobuf:"bytes,2,opt,name=calledStationID,proto3" json:"calledStationID,omitempty"`
}

func (x *DisconnectRequest) Reset() {
	*x = DisconnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ue_sim_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectRequest) ProtoMessage() {}

func (x *DisconnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ue_sim_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectRequest.ProtoReflect.Descriptor instead.
func (*DisconnectRequest) Descriptor() ([]byte, []int) {
	return file_ue_sim_proto_rawDescGZIP(), []int{3}
}

func (x *DisconnectRequest) GetImsi() string {
	if x != nil {
		return x.Imsi
	}
	return ""
}

func (x *DisconnectRequest) GetCalledStationID() string {
	if x != nil {
		return x.CalledStationID
	}
	return ""
}

type DisconnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RadiusPacket []byte `protobuf:"bytes,1,opt,name=radiusPacket,proto3" json:"radiusPacket,omitempty"`
}

func (x *DisconnectResponse) Reset() {
	*x = DisconnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ue_sim_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectResponse) ProtoMessage() {}

func (x *DisconnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ue_sim_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectResponse.ProtoReflect.Descriptor instead.
func (*DisconnectResponse) Descriptor() ([]byte, []int) {
	return file_ue_sim_proto_rawDescGZIP(), []int{4}
}

func (x *DisconnectResponse) GetRadiusPacket() []byte {
	if x != nil {
		return x.RadiusPacket
	}
	return nil
}

type GenTrafficRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Imsi                    string                `protobuf:"bytes,1,opt,name=imsi,proto3" json:"imsi,omitempty"`
	Volume                  *wrappers.StringValue `protobuf:"bytes,2,opt,name=volume,proto3" json:"volume,omitempty"`
	Bitrate                 *wrappers.StringValue `protobuf:"bytes,3,opt,name=bitrate,proto3" json:"bitrate,omitempty"`
	TimeInSecs              uint64                `protobuf:"varint,4,opt,name=timeInSecs,proto3" json:"timeInSecs,omitempty"`
	ReportingIntervalInSecs uint64                `protobuf:"varint,5,opt,name=reportingIntervalInSecs,proto3" json:"reportingIntervalInSecs,omitempty"`
	ReverseMode             bool                  `protobuf:"varint,6,opt,name=reverseMode,proto3" json:"reverseMode,omitempty"`
	// Configure a max timeout iperf client will run
	Timeout uint32 `protobuf:"varint,7,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// Enables/Disable a function to check if server is alive before sending any traffic
	DisableServerReachabilityCheck bool `protobuf:"varint,8,opt,name=disableServerReachabilityCheck,proto3" json:"disableServerReachabilityCheck,omitempty"`
}

func (x *GenTrafficRequest) Reset() {
	*x = GenTrafficRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ue_sim_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenTrafficRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenTrafficRequest) ProtoMessage() {}

func (x *GenTrafficRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ue_sim_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenTrafficRequest.ProtoReflect.Descriptor instead.
func (*GenTrafficRequest) Descriptor() ([]byte, []int) {
	return file_ue_sim_proto_rawDescGZIP(), []int{5}
}

func (x *GenTrafficRequest) GetImsi() string {
	if x != nil {
		return x.Imsi
	}
	return ""
}

func (x *GenTrafficRequest) GetVolume() *wrappers.StringValue {
	if x != nil {
		return x.Volume
	}
	return nil
}

func (x *GenTrafficRequest) GetBitrate() *wrappers.StringValue {
	if x != nil {
		return x.Bitrate
	}
	return nil
}

func (x *GenTrafficRequest) GetTimeInSecs() uint64 {
	if x != nil {
		return x.TimeInSecs
	}
	return 0
}

func (x *GenTrafficRequest) GetReportingIntervalInSecs() uint64 {
	if x != nil {
		return x.ReportingIntervalInSecs
	}
	return 0
}

func (x *GenTrafficRequest) GetReverseMode() bool {
	if x != nil {
		return x.ReverseMode
	}
	return false
}

func (x *GenTrafficRequest) GetTimeout() uint32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *GenTrafficRequest) GetDisableServerReachabilityCheck() bool {
	if x != nil {
		return x.DisableServerReachabilityCheck
	}
	return false
}

type GenTrafficResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output []byte `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *GenTrafficResponse) Reset() {
	*x = GenTrafficResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ue_sim_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenTrafficResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenTrafficResponse) ProtoMessage() {}

func (x *GenTrafficResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ue_sim_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenTrafficResponse.ProtoReflect.Descriptor instead.
func (*GenTrafficResponse) Descriptor() ([]byte, []int) {
	return file_ue_sim_proto_rawDescGZIP(), []int{6}
}

func (x *GenTrafficResponse) GetOutput() []byte {
	if x != nil {
		return x.Output
	}
	return nil
}

var File_ue_sim_proto protoreflect.FileDescriptor

var file_ue_sim_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x75, 0x65, 0x5f, 0x73, 0x69, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x63, 0x77, 0x66, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x08, 0x55, 0x45, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6d, 0x73, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x69, 0x6d, 0x73, 0x69, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x4b,
	0x65, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6f, 0x70, 0x63, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x4f, 0x70, 0x63, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x65, 0x71, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x73, 0x65, 0x71, 0x22,
	0x53, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6d, 0x73, 0x69, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6d, 0x73, 0x69, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x61,
	0x6c, 0x6c, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x22, 0x3a, 0x0a, 0x14, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0c, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x22, 0x51, 0x0a, 0x11, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6d, 0x73, 0x69, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6d, 0x73, 0x69, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x61, 0x6c,
	0x6c, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x22, 0x38, 0x0a, 0x12, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x61, 0x64,
	0x69, 0x75, 0x73, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0c, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x22, 0xf3, 0x02,
	0x0a, 0x11, 0x47, 0x65, 0x6e, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6d, 0x73, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x69, 0x6d, 0x73, 0x69, 0x12, 0x34, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x36, 0x0a,
	0x07, 0x62, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x62, 0x69,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x53,
	0x65, 0x63, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x49,
	0x6e, 0x53, 0x65, 0x63, 0x73, 0x12, 0x38, 0x0a, 0x17, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x49, 0x6e, 0x53, 0x65, 0x63, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x17, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x49, 0x6e, 0x53, 0x65, 0x63, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x4d, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x46, 0x0a, 0x1e, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x61, 0x63,
	0x68, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x1e, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x61, 0x63, 0x68, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x22, 0x2c, 0x0a, 0x12, 0x47, 0x65, 0x6e, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x32, 0xa7, 0x02, 0x0a, 0x05, 0x55, 0x45, 0x53, 0x69, 0x6d, 0x12, 0x31, 0x0a, 0x05, 0x41,
	0x64, 0x64, 0x55, 0x45, 0x12, 0x13, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x63, 0x77, 0x66,
	0x2e, 0x55, 0x45, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x11, 0x2e, 0x6d, 0x61, 0x67, 0x6d,
	0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x51,
	0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x1e,
	0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x63, 0x77, 0x66, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x63, 0x77, 0x66, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4b, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12,
	0x1c, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x63, 0x77, 0x66, 0x2e, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x63, 0x77, 0x66, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b,
	0x0a, 0x0a, 0x47, 0x65, 0x6e, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x12, 0x1c, 0x2e, 0x6d,
	0x61, 0x67, 0x6d, 0x61, 0x2e, 0x63, 0x77, 0x66, 0x2e, 0x47, 0x65, 0x6e, 0x54, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x61, 0x67,
	0x6d, 0x61, 0x2e, 0x63, 0x77, 0x66, 0x2e, 0x47, 0x65, 0x6e, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x61, 0x67,
	0x6d, 0x61, 0x2f, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f, 0x63, 0x77, 0x66, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ue_sim_proto_rawDescOnce sync.Once
	file_ue_sim_proto_rawDescData = file_ue_sim_proto_rawDesc
)

func file_ue_sim_proto_rawDescGZIP() []byte {
	file_ue_sim_proto_rawDescOnce.Do(func() {
		file_ue_sim_proto_rawDescData = protoimpl.X.CompressGZIP(file_ue_sim_proto_rawDescData)
	})
	return file_ue_sim_proto_rawDescData
}

var file_ue_sim_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_ue_sim_proto_goTypes = []interface{}{
	(*UEConfig)(nil),             // 0: magma.cwf.UEConfig
	(*AuthenticateRequest)(nil),  // 1: magma.cwf.AuthenticateRequest
	(*AuthenticateResponse)(nil), // 2: magma.cwf.AuthenticateResponse
	(*DisconnectRequest)(nil),    // 3: magma.cwf.DisconnectRequest
	(*DisconnectResponse)(nil),   // 4: magma.cwf.DisconnectResponse
	(*GenTrafficRequest)(nil),    // 5: magma.cwf.GenTrafficRequest
	(*GenTrafficResponse)(nil),   // 6: magma.cwf.GenTrafficResponse
	(*wrappers.StringValue)(nil), // 7: google.protobuf.StringValue
	(*protos.Void)(nil),          // 8: magma.orc8r.Void
}
var file_ue_sim_proto_depIdxs = []int32{
	7, // 0: magma.cwf.GenTrafficRequest.volume:type_name -> google.protobuf.StringValue
	7, // 1: magma.cwf.GenTrafficRequest.bitrate:type_name -> google.protobuf.StringValue
	0, // 2: magma.cwf.UESim.AddUE:input_type -> magma.cwf.UEConfig
	1, // 3: magma.cwf.UESim.Authenticate:input_type -> magma.cwf.AuthenticateRequest
	3, // 4: magma.cwf.UESim.Disconnect:input_type -> magma.cwf.DisconnectRequest
	5, // 5: magma.cwf.UESim.GenTraffic:input_type -> magma.cwf.GenTrafficRequest
	8, // 6: magma.cwf.UESim.AddUE:output_type -> magma.orc8r.Void
	2, // 7: magma.cwf.UESim.Authenticate:output_type -> magma.cwf.AuthenticateResponse
	4, // 8: magma.cwf.UESim.Disconnect:output_type -> magma.cwf.DisconnectResponse
	6, // 9: magma.cwf.UESim.GenTraffic:output_type -> magma.cwf.GenTrafficResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ue_sim_proto_init() }
func file_ue_sim_proto_init() {
	if File_ue_sim_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ue_sim_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UEConfig); i {
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
		file_ue_sim_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateRequest); i {
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
		file_ue_sim_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateResponse); i {
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
		file_ue_sim_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectRequest); i {
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
		file_ue_sim_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectResponse); i {
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
		file_ue_sim_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenTrafficRequest); i {
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
		file_ue_sim_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenTrafficResponse); i {
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
			RawDescriptor: file_ue_sim_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ue_sim_proto_goTypes,
		DependencyIndexes: file_ue_sim_proto_depIdxs,
		MessageInfos:      file_ue_sim_proto_msgTypes,
	}.Build()
	File_ue_sim_proto = out.File
	file_ue_sim_proto_rawDesc = nil
	file_ue_sim_proto_goTypes = nil
	file_ue_sim_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UESimClient is the client API for UESim service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UESimClient interface {
	// Adds a new UE to the store.
	//
	AddUE(ctx context.Context, in *UEConfig, opts ...grpc.CallOption) (*protos.Void, error)
	// Triggers an authentication for the UE with the specified imsi.
	//
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
	Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error)
	// Triggers iperf traffic towards the CWAG
	GenTraffic(ctx context.Context, in *GenTrafficRequest, opts ...grpc.CallOption) (*GenTrafficResponse, error)
}

type uESimClient struct {
	cc grpc.ClientConnInterface
}

func NewUESimClient(cc grpc.ClientConnInterface) UESimClient {
	return &uESimClient{cc}
}

func (c *uESimClient) AddUE(ctx context.Context, in *UEConfig, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.cwf.UESim/AddUE", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uESimClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := c.cc.Invoke(ctx, "/magma.cwf.UESim/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uESimClient) Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error) {
	out := new(DisconnectResponse)
	err := c.cc.Invoke(ctx, "/magma.cwf.UESim/Disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uESimClient) GenTraffic(ctx context.Context, in *GenTrafficRequest, opts ...grpc.CallOption) (*GenTrafficResponse, error) {
	out := new(GenTrafficResponse)
	err := c.cc.Invoke(ctx, "/magma.cwf.UESim/GenTraffic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UESimServer is the server API for UESim service.
type UESimServer interface {
	// Adds a new UE to the store.
	//
	AddUE(context.Context, *UEConfig) (*protos.Void, error)
	// Triggers an authentication for the UE with the specified imsi.
	//
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
	Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error)
	// Triggers iperf traffic towards the CWAG
	GenTraffic(context.Context, *GenTrafficRequest) (*GenTrafficResponse, error)
}

// UnimplementedUESimServer can be embedded to have forward compatible implementations.
type UnimplementedUESimServer struct {
}

func (*UnimplementedUESimServer) AddUE(context.Context, *UEConfig) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUE not implemented")
}
func (*UnimplementedUESimServer) Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (*UnimplementedUESimServer) Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}
func (*UnimplementedUESimServer) GenTraffic(context.Context, *GenTrafficRequest) (*GenTrafficResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenTraffic not implemented")
}

func RegisterUESimServer(s *grpc.Server, srv UESimServer) {
	s.RegisterService(&_UESim_serviceDesc, srv)
}

func _UESim_AddUE_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UEConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UESimServer).AddUE(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.cwf.UESim/AddUE",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UESimServer).AddUE(ctx, req.(*UEConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _UESim_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UESimServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.cwf.UESim/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UESimServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UESim_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UESimServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.cwf.UESim/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UESimServer).Disconnect(ctx, req.(*DisconnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UESim_GenTraffic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenTrafficRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UESimServer).GenTraffic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.cwf.UESim/GenTraffic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UESimServer).GenTraffic(ctx, req.(*GenTrafficRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UESim_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.cwf.UESim",
	HandlerType: (*UESimServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUE",
			Handler:    _UESim_AddUE_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _UESim_Authenticate_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _UESim_Disconnect_Handler,
		},
		{
			MethodName: "GenTraffic",
			Handler:    _UESim_GenTraffic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ue_sim.proto",
}
